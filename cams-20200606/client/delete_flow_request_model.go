// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *DeleteFlowRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *DeleteFlowRequest
	GetFlowId() *string
	SetOwnerId(v int64) *DeleteFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFlowRequest
	GetResourceOwnerId() *int64
}

type DeleteFlowRequest struct {
	// example:
	//
	// 示例值示例值
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

func (s DeleteFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DeleteFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFlowRequest) SetCustSpaceId(v string) *DeleteFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteFlowRequest) SetFlowId(v string) *DeleteFlowRequest {
	s.FlowId = &v
	return s
}

func (s *DeleteFlowRequest) SetOwnerId(v int64) *DeleteFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowRequest) SetResourceOwnerAccount(v string) *DeleteFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFlowRequest) SetResourceOwnerId(v int64) *DeleteFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowRequest) Validate() error {
	return dara.Validate(s)
}
