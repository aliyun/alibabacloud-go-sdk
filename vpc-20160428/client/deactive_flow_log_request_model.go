// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *DeactiveFlowLogRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *DeactiveFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeactiveFlowLogRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeactiveFlowLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeactiveFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeactiveFlowLogRequest
	GetResourceOwnerId() *int64
}

type DeactiveFlowLogRequest struct {
	// The ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// fl-m5e8vhz2t21sel1nq****
	FlowLogId    *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where you want to create the flow log. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeactiveFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactiveFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DeactiveFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeactiveFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeactiveFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeactiveFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeactiveFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeactiveFlowLogRequest) SetFlowLogId(v string) *DeactiveFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetOwnerAccount(v string) *DeactiveFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetOwnerId(v int64) *DeactiveFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetRegionId(v string) *DeactiveFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetResourceOwnerAccount(v string) *DeactiveFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetResourceOwnerId(v int64) *DeactiveFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeactiveFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
