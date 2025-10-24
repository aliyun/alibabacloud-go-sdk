// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *QueryTemplateListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryTemplateListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryTemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTemplateListRequest
	GetResourceOwnerId() *int64
	SetTemplateIds(v string) *QueryTemplateListRequest
	GetTemplateIds() *string
}

type QueryTemplateListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The response parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****,88c6ca184c0e424d5w5b665e2a12****
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
}

func (s QueryTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListRequest) GoString() string {
	return s.String()
}

func (s *QueryTemplateListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryTemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTemplateListRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *QueryTemplateListRequest) SetOwnerAccount(v string) *QueryTemplateListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryTemplateListRequest) SetOwnerId(v int64) *QueryTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTemplateListRequest) SetResourceOwnerAccount(v string) *QueryTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTemplateListRequest) SetResourceOwnerId(v int64) *QueryTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTemplateListRequest) SetTemplateIds(v string) *QueryTemplateListRequest {
	s.TemplateIds = &v
	return s
}

func (s *QueryTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
