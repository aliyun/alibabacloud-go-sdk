// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerCategoryDictionaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *GetCustomerCategoryDictionaryRequest
	GetType() *string
}

type GetCustomerCategoryDictionaryRequest struct {
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCustomerCategoryDictionaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryDictionaryRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryDictionaryRequest) GetType() *string {
	return s.Type
}

func (s *GetCustomerCategoryDictionaryRequest) SetType(v string) *GetCustomerCategoryDictionaryRequest {
	s.Type = &v
	return s
}

func (s *GetCustomerCategoryDictionaryRequest) Validate() error {
	return dara.Validate(s)
}
