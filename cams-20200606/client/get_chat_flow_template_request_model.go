// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetChatFlowTemplateRequest
	GetBizCode() *string
	SetId(v int64) *GetChatFlowTemplateRequest
	GetId() *int64
	SetOwnerId(v int64) *GetChatFlowTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetChatFlowTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatFlowTemplateRequest
	GetResourceOwnerId() *int64
}

type GetChatFlowTemplateRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Template ID
	//
	// example:
	//
	// 2
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetChatFlowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetChatFlowTemplateRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetChatFlowTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *GetChatFlowTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatFlowTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatFlowTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatFlowTemplateRequest) SetBizCode(v string) *GetChatFlowTemplateRequest {
	s.BizCode = &v
	return s
}

func (s *GetChatFlowTemplateRequest) SetId(v int64) *GetChatFlowTemplateRequest {
	s.Id = &v
	return s
}

func (s *GetChatFlowTemplateRequest) SetOwnerId(v int64) *GetChatFlowTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatFlowTemplateRequest) SetResourceOwnerAccount(v string) *GetChatFlowTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatFlowTemplateRequest) SetResourceOwnerId(v int64) *GetChatFlowTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatFlowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
