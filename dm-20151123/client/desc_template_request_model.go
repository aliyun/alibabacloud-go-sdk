// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromType(v int32) *DescTemplateRequest
	GetFromType() *int32
	SetOwnerId(v int64) *DescTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *DescTemplateRequest
	GetTemplateId() *int32
}

type DescTemplateRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3xxxx2
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescTemplateRequest) GetFromType() *int32 {
	return s.FromType
}

func (s *DescTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescTemplateRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescTemplateRequest) SetFromType(v int32) *DescTemplateRequest {
	s.FromType = &v
	return s
}

func (s *DescTemplateRequest) SetOwnerId(v int64) *DescTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescTemplateRequest) SetResourceOwnerAccount(v string) *DescTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescTemplateRequest) SetResourceOwnerId(v int64) *DescTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescTemplateRequest) SetTemplateId(v int32) *DescTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DescTemplateRequest) Validate() error {
	return dara.Validate(s)
}
