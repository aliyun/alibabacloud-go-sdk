// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLocaleString(v string) *GetCustomerCategoryRequest
	GetLocaleString() *string
	SetUserId(v int64) *GetCustomerCategoryRequest
	GetUserId() *int64
}

type GetCustomerCategoryRequest struct {
	LocaleString *string `json:"LocaleString,omitempty" xml:"LocaleString,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCustomerCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryRequest) GetLocaleString() *string {
	return s.LocaleString
}

func (s *GetCustomerCategoryRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *GetCustomerCategoryRequest) SetLocaleString(v string) *GetCustomerCategoryRequest {
	s.LocaleString = &v
	return s
}

func (s *GetCustomerCategoryRequest) SetUserId(v int64) *GetCustomerCategoryRequest {
	s.UserId = &v
	return s
}

func (s *GetCustomerCategoryRequest) Validate() error {
	return dara.Validate(s)
}
