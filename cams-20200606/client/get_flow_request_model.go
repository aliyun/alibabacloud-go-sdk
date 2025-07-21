// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetFlowRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *GetFlowRequest
	GetFlowId() *string
	SetOwnerId(v int64) *GetFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFlowRequest
	GetResourceOwnerId() *int64
}

type GetFlowRequest struct {
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowRequest) GoString() string {
	return s.String()
}

func (s *GetFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFlowRequest) SetCustSpaceId(v string) *GetFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetFlowRequest) SetFlowId(v string) *GetFlowRequest {
	s.FlowId = &v
	return s
}

func (s *GetFlowRequest) SetOwnerId(v int64) *GetFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFlowRequest) SetResourceOwnerAccount(v string) *GetFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFlowRequest) SetResourceOwnerId(v int64) *GetFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFlowRequest) Validate() error {
	return dara.Validate(s)
}
