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
}

type ListProductsRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

func (s *ListProductsRequest) SetLanguage(v string) *ListProductsRequest {
	s.Language = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
