// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomerInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *ModifyCustomerInfoRequest
	GetBiz() *string
	SetCustomerCategory(v string) *ModifyCustomerInfoRequest
	GetCustomerCategory() *string
	SetCustomerSubCategory(v string) *ModifyCustomerInfoRequest
	GetCustomerSubCategory() *string
	SetUserId(v int64) *ModifyCustomerInfoRequest
	GetUserId() *int64
	SetWebsite(v string) *ModifyCustomerInfoRequest
	GetWebsite() *string
}

type ModifyCustomerInfoRequest struct {
	Biz                 *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	CustomerCategory    *string `json:"CustomerCategory,omitempty" xml:"CustomerCategory,omitempty"`
	CustomerSubCategory *string `json:"CustomerSubCategory,omitempty" xml:"CustomerSubCategory,omitempty"`
	// This parameter is required.
	UserId  *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Website *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s ModifyCustomerInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomerInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomerInfoRequest) GetBiz() *string {
	return s.Biz
}

func (s *ModifyCustomerInfoRequest) GetCustomerCategory() *string {
	return s.CustomerCategory
}

func (s *ModifyCustomerInfoRequest) GetCustomerSubCategory() *string {
	return s.CustomerSubCategory
}

func (s *ModifyCustomerInfoRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ModifyCustomerInfoRequest) GetWebsite() *string {
	return s.Website
}

func (s *ModifyCustomerInfoRequest) SetBiz(v string) *ModifyCustomerInfoRequest {
	s.Biz = &v
	return s
}

func (s *ModifyCustomerInfoRequest) SetCustomerCategory(v string) *ModifyCustomerInfoRequest {
	s.CustomerCategory = &v
	return s
}

func (s *ModifyCustomerInfoRequest) SetCustomerSubCategory(v string) *ModifyCustomerInfoRequest {
	s.CustomerSubCategory = &v
	return s
}

func (s *ModifyCustomerInfoRequest) SetUserId(v int64) *ModifyCustomerInfoRequest {
	s.UserId = &v
	return s
}

func (s *ModifyCustomerInfoRequest) SetWebsite(v string) *ModifyCustomerInfoRequest {
	s.Website = &v
	return s
}

func (s *ModifyCustomerInfoRequest) Validate() error {
	return dara.Validate(s)
}
