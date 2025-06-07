package usecase

import (
	"go-api-1/model"
	"go-api-1/repository"
)

type ProductUseCase struct {
	//Repository
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{repository: repo}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()

}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil

}

func (pu *ProductUseCase) GetProductByID(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductByID(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
