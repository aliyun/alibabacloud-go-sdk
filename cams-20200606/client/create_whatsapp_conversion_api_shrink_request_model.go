// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhatsappConversionApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWhatsappConversionApiShrinkRequest
	GetCode() *string
	SetInstanceId(v string) *CreateWhatsappConversionApiShrinkRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *CreateWhatsappConversionApiShrinkRequest
	GetOwnerId() *int64
	SetPermissionsShrink(v string) *CreateWhatsappConversionApiShrinkRequest
	GetPermissionsShrink() *string
	SetResourceOwnerAccount(v string) *CreateWhatsappConversionApiShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateWhatsappConversionApiShrinkRequest
	GetResourceOwnerId() *int64
}

type CreateWhatsappConversionApiShrinkRequest struct {
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 131
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PermissionsShrink    *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateWhatsappConversionApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWhatsappConversionApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWhatsappConversionApiShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *CreateWhatsappConversionApiShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWhatsappConversionApiShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateWhatsappConversionApiShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *CreateWhatsappConversionApiShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateWhatsappConversionApiShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateWhatsappConversionApiShrinkRequest) SetCode(v string) *CreateWhatsappConversionApiShrinkRequest {
	s.Code = &v
	return s
}

func (s *CreateWhatsappConversionApiShrinkRequest) SetInstanceId(v string) *CreateWhatsappConversionApiShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWhatsappConversionApiShrinkRequest) SetOwnerId(v int64) *CreateWhatsappConversionApiShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateWhatsappConversionApiShrinkRequest) SetPermissionsShrink(v string) *CreateWhatsappConversionApiShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *CreateWhatsappConversionApiShrinkRequest) SetResourceOwnerAccount(v string) *CreateWhatsappConversionApiShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateWhatsappConversionApiShrinkRequest) SetResourceOwnerId(v int64) *CreateWhatsappConversionApiShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateWhatsappConversionApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
