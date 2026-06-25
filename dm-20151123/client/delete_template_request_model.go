// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromType(v int32) *DeleteTemplateRequest
	GetFromType() *int32
	SetOwnerId(v int64) *DeleteTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *DeleteTemplateRequest
	GetTemplateId() *int32
}

type DeleteTemplateRequest struct {
	// The channel through which the user accesses the service. Default value: 1.
	//
	// example:
	//
	// 1
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 409481
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetFromType() *int32 {
	return s.FromType
}

func (s *DeleteTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTemplateRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DeleteTemplateRequest) SetFromType(v int32) *DeleteTemplateRequest {
	s.FromType = &v
	return s
}

func (s *DeleteTemplateRequest) SetOwnerId(v int64) *DeleteTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) SetResourceOwnerAccount(v string) *DeleteTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTemplateRequest) SetResourceOwnerId(v int64) *DeleteTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateId(v int32) *DeleteTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
