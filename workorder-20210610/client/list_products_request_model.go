// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListProductsRequest
	GetLanguage() *string
	SetName(v string) *ListProductsRequest
	GetName() *string
}

type ListProductsRequest struct {
	// The language that you use, supporting English, Chinese, and Japanese. Valid values: en, zh, and jp, which indicate English, Chinese, and Japanese, respectively.
	//
	// example:
	//
	// cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the product. Fuzzy search is supported. This parameter is optional.
	//
	// example:
	//
	// ecs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListProductsRequest) GetName() *string {
	return s.Name
}

func (s *ListProductsRequest) SetLanguage(v string) *ListProductsRequest {
	s.Language = &v
	return s
}

func (s *ListProductsRequest) SetName(v string) *ListProductsRequest {
	s.Name = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
