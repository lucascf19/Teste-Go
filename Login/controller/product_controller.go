package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController{
	return ProductController{
		ProductUsecase: usecase,
		}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUsecase.GetProducts()
	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context){
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if (id == "") {
		response := model.Response{
			Massage: "Id do produto não pode ser nulo" ,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if (err != nil) {
		response := model.Response{
			Massage: "Id do produto precisa ser um numero" ,
	}
	ctx.JSON(http.StatusBadRequest, response)
	return
}


	product, err := p.ProductUsecase.GetProductById(productId)
	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}


	if product == nil {
		response := model.Response{
			Massage: "Produto não foi encontrado no banco de dados" ,
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
