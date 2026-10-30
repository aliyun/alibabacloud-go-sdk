// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyShareImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAccount(v []*int64) *ModifyShareImageRequest
	GetAddAccount() []*int64
	SetBizType(v int32) *ModifyShareImageRequest
	GetBizType() *int32
	SetImageId(v string) *ModifyShareImageRequest
	GetImageId() *string
	SetProductType(v string) *ModifyShareImageRequest
	GetProductType() *string
	SetRemoveAccount(v []*int64) *ModifyShareImageRequest
	GetRemoveAccount() []*int64
}

type ModifyShareImageRequest struct {
	// The IDs of the Alibaba Cloud accounts to share the image with.
	AddAccount []*int64 `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
	// The business line. Only the Enterprise Edition scenario is supported. Default value: 1.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The ID of the source image.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The product type, such as Cloud Desktop, cloud application, or workstation.
	//
	// example:
	//
	// CloudDesktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The IDs of the Alibaba Cloud accounts to remove from the image sharing list.
	RemoveAccount []*int64 `json:"RemoveAccount,omitempty" xml:"RemoveAccount,omitempty" type:"Repeated"`
}

func (s ModifyShareImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyShareImageRequest) GoString() string {
	return s.String()
}

func (s *ModifyShareImageRequest) GetAddAccount() []*int64 {
	return s.AddAccount
}

func (s *ModifyShareImageRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ModifyShareImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyShareImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyShareImageRequest) GetRemoveAccount() []*int64 {
	return s.RemoveAccount
}

func (s *ModifyShareImageRequest) SetAddAccount(v []*int64) *ModifyShareImageRequest {
	s.AddAccount = v
	return s
}

func (s *ModifyShareImageRequest) SetBizType(v int32) *ModifyShareImageRequest {
	s.BizType = &v
	return s
}

func (s *ModifyShareImageRequest) SetImageId(v string) *ModifyShareImageRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyShareImageRequest) SetProductType(v string) *ModifyShareImageRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyShareImageRequest) SetRemoveAccount(v []*int64) *ModifyShareImageRequest {
	s.RemoveAccount = v
	return s
}

func (s *ModifyShareImageRequest) Validate() error {
	return dara.Validate(s)
}
