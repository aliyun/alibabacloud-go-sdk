// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *PublishFlowRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *PublishFlowRequest
	GetFlowId() *string
	SetOwnerId(v int64) *PublishFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PublishFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PublishFlowRequest
	GetResourceOwnerId() *int64
}

type PublishFlowRequest struct {
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PublishFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowRequest) GoString() string {
	return s.String()
}

func (s *PublishFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *PublishFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *PublishFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PublishFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PublishFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PublishFlowRequest) SetCustSpaceId(v string) *PublishFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *PublishFlowRequest) SetFlowId(v string) *PublishFlowRequest {
	s.FlowId = &v
	return s
}

func (s *PublishFlowRequest) SetOwnerId(v int64) *PublishFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishFlowRequest) SetResourceOwnerAccount(v string) *PublishFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PublishFlowRequest) SetResourceOwnerId(v int64) *PublishFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PublishFlowRequest) Validate() error {
	return dara.Validate(s)
}
