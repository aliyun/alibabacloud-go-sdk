// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomerInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *UpdateCustomerInformationRequest
	GetBiz() *string
	SetCustomerCategory(v string) *UpdateCustomerInformationRequest
	GetCustomerCategory() *string
	SetCustomerSubCategory(v string) *UpdateCustomerInformationRequest
	GetCustomerSubCategory() *string
	SetUserId(v int64) *UpdateCustomerInformationRequest
	GetUserId() *int64
	SetWebsite(v string) *UpdateCustomerInformationRequest
	GetWebsite() *string
}

type UpdateCustomerInformationRequest struct {
	Biz                 *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	CustomerCategory    *string `json:"CustomerCategory,omitempty" xml:"CustomerCategory,omitempty"`
	CustomerSubCategory *string `json:"CustomerSubCategory,omitempty" xml:"CustomerSubCategory,omitempty"`
	// This parameter is required.
	UserId  *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Website *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s UpdateCustomerInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomerInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomerInformationRequest) GetBiz() *string {
	return s.Biz
}

func (s *UpdateCustomerInformationRequest) GetCustomerCategory() *string {
	return s.CustomerCategory
}

func (s *UpdateCustomerInformationRequest) GetCustomerSubCategory() *string {
	return s.CustomerSubCategory
}

func (s *UpdateCustomerInformationRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *UpdateCustomerInformationRequest) GetWebsite() *string {
	return s.Website
}

func (s *UpdateCustomerInformationRequest) SetBiz(v string) *UpdateCustomerInformationRequest {
	s.Biz = &v
	return s
}

func (s *UpdateCustomerInformationRequest) SetCustomerCategory(v string) *UpdateCustomerInformationRequest {
	s.CustomerCategory = &v
	return s
}

func (s *UpdateCustomerInformationRequest) SetCustomerSubCategory(v string) *UpdateCustomerInformationRequest {
	s.CustomerSubCategory = &v
	return s
}

func (s *UpdateCustomerInformationRequest) SetUserId(v int64) *UpdateCustomerInformationRequest {
	s.UserId = &v
	return s
}

func (s *UpdateCustomerInformationRequest) SetWebsite(v string) *UpdateCustomerInformationRequest {
	s.Website = &v
	return s
}

func (s *UpdateCustomerInformationRequest) Validate() error {
	return dara.Validate(s)
}
