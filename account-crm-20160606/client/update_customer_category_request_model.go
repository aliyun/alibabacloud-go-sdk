// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomerCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamList(v string) *UpdateCustomerCategoryRequest
	GetParamList() *string
	SetUserId(v int64) *UpdateCustomerCategoryRequest
	GetUserId() *int64
}

type UpdateCustomerCategoryRequest struct {
	// This parameter is required.
	ParamList *string `json:"ParamList,omitempty" xml:"ParamList,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateCustomerCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomerCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomerCategoryRequest) GetParamList() *string {
	return s.ParamList
}

func (s *UpdateCustomerCategoryRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *UpdateCustomerCategoryRequest) SetParamList(v string) *UpdateCustomerCategoryRequest {
	s.ParamList = &v
	return s
}

func (s *UpdateCustomerCategoryRequest) SetUserId(v int64) *UpdateCustomerCategoryRequest {
	s.UserId = &v
	return s
}

func (s *UpdateCustomerCategoryRequest) Validate() error {
	return dara.Validate(s)
}
