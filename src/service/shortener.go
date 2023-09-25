package service

import (
	"net/http"
	"server/src/model"
	"server/src/repository"
	"server/src/util"

	"github.com/gin-gonic/gin"
)

type ShortenerInterface interface {
	ShortenUrl(c *gin.Context)
	GetOriginalUrl(c *gin.Context)
}

type ShortenerService struct {
	Self ShortenerInterface
}

type UrlShortenRequest struct {
	Url string `json:"url"`
}

type UrlShortenerResponse struct {
	Message string
	Url     string
}

func New() ShortenerInterface {
	service := &ShortenerService{}
	service.Self = service
	return service
}

func (s *ShortenerService) ShortenUrl(gctx *gin.Context) {
	var url UrlShortenRequest
	if gctx.ShouldBind(&url) == nil {
		if url.Url == "" {
			gctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})

			return
		}

		shorten := util.RandStringRunes(5)
		repository.DB.SaveUrl(gctx, model.Url{
			Original: url.Url,
			Shorten:  shorten,
		})

		gctx.JSON(http.StatusOK, UrlShortenerResponse{
			Message: "success",
			Url:     shorten,
		})
	}
}

func (s *ShortenerService) GetOriginalUrl(gctx *gin.Context) {
	url := gctx.Param("url")
	original, err := repository.DB.GetOriginalUrl(gctx, url)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, UrlShortenerResponse{
			Message: "error",
		})

		return
	}

	gctx.JSON(http.StatusOK, UrlShortenerResponse{
		Message: "success",
		Url:     original,
	})
}
