// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *TemplateListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TemplateListRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int64) *TemplateListRequest
	GetTemplateId() *int64
}

type TemplateListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 必须空参
	//
	// example:
	//
	// 9
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s TemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s TemplateListRequest) GoString() string {
	return s.String()
}

func (s *TemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TemplateListRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *TemplateListRequest) SetOwnerId(v int64) *TemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *TemplateListRequest) SetResourceOwnerAccount(v string) *TemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TemplateListRequest) SetResourceOwnerId(v int64) *TemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TemplateListRequest) SetTemplateId(v int64) *TemplateListRequest {
	s.TemplateId = &v
	return s
}

func (s *TemplateListRequest) Validate() error {
	return dara.Validate(s)
}
