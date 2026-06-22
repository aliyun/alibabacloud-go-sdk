// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprecateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *DeprecateFlowRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *DeprecateFlowRequest
	GetFlowId() *string
	SetOwnerId(v int64) *DeprecateFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeprecateFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeprecateFlowRequest
	GetResourceOwnerId() *int64
}

type DeprecateFlowRequest struct {
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

func (s DeprecateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeprecateFlowRequest) GoString() string {
	return s.String()
}

func (s *DeprecateFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeprecateFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DeprecateFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeprecateFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeprecateFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeprecateFlowRequest) SetCustSpaceId(v string) *DeprecateFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeprecateFlowRequest) SetFlowId(v string) *DeprecateFlowRequest {
	s.FlowId = &v
	return s
}

func (s *DeprecateFlowRequest) SetOwnerId(v int64) *DeprecateFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeprecateFlowRequest) SetResourceOwnerAccount(v string) *DeprecateFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeprecateFlowRequest) SetResourceOwnerId(v int64) *DeprecateFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeprecateFlowRequest) Validate() error {
	return dara.Validate(s)
}
