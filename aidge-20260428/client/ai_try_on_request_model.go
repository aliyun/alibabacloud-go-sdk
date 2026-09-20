// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTryOnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothImageUrl(v string) *AiTryOnRequest
	GetClothImageUrl() *string
	SetClothType(v string) *AiTryOnRequest
	GetClothType() *string
	SetModelImageUrl(v string) *AiTryOnRequest
	GetModelImageUrl() *string
	SetResolution(v string) *AiTryOnRequest
	GetResolution() *string
}

type AiTryOnRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/cloth.png
	ClothImageUrl *string `json:"ClothImageUrl,omitempty" xml:"ClothImageUrl,omitempty"`
	// example:
	//
	// tops
	ClothType *string `json:"ClothType,omitempty" xml:"ClothType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/model.png
	ModelImageUrl *string `json:"ModelImageUrl,omitempty" xml:"ModelImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s AiTryOnRequest) String() string {
	return dara.Prettify(s)
}

func (s AiTryOnRequest) GoString() string {
	return s.String()
}

func (s *AiTryOnRequest) GetClothImageUrl() *string {
	return s.ClothImageUrl
}

func (s *AiTryOnRequest) GetClothType() *string {
	return s.ClothType
}

func (s *AiTryOnRequest) GetModelImageUrl() *string {
	return s.ModelImageUrl
}

func (s *AiTryOnRequest) GetResolution() *string {
	return s.Resolution
}

func (s *AiTryOnRequest) SetClothImageUrl(v string) *AiTryOnRequest {
	s.ClothImageUrl = &v
	return s
}

func (s *AiTryOnRequest) SetClothType(v string) *AiTryOnRequest {
	s.ClothType = &v
	return s
}

func (s *AiTryOnRequest) SetModelImageUrl(v string) *AiTryOnRequest {
	s.ModelImageUrl = &v
	return s
}

func (s *AiTryOnRequest) SetResolution(v string) *AiTryOnRequest {
	s.Resolution = &v
	return s
}

func (s *AiTryOnRequest) Validate() error {
	return dara.Validate(s)
}
