// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmarttagTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteSmarttagTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSmarttagTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSmarttagTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSmarttagTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *DeleteSmarttagTemplateRequest
	GetTemplateId() *string
}

type DeleteSmarttagTemplateRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the template that you want to delete. You can obtain the template ID from the response of the [AddSmarttagTemplate](https://help.aliyun.com/document_detail/187759.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05de22f255284c7a8d2aab535dde****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteSmarttagTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmarttagTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmarttagTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSmarttagTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSmarttagTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSmarttagTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSmarttagTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteSmarttagTemplateRequest) SetOwnerAccount(v string) *DeleteSmarttagTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSmarttagTemplateRequest) SetOwnerId(v int64) *DeleteSmarttagTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmarttagTemplateRequest) SetResourceOwnerAccount(v string) *DeleteSmarttagTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmarttagTemplateRequest) SetResourceOwnerId(v int64) *DeleteSmarttagTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmarttagTemplateRequest) SetTemplateId(v string) *DeleteSmarttagTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteSmarttagTemplateRequest) Validate() error {
	return dara.Validate(s)
}
