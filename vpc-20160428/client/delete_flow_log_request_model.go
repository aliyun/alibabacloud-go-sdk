// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *DeleteFlowLogRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *DeleteFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteFlowLogRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteFlowLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFlowLogRequest
	GetResourceOwnerId() *int64
}

type DeleteFlowLogRequest struct {
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
	// The region ID of the flow log. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DeleteFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFlowLogRequest) SetFlowLogId(v string) *DeleteFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeleteFlowLogRequest) SetOwnerAccount(v string) *DeleteFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFlowLogRequest) SetOwnerId(v int64) *DeleteFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowLogRequest) SetRegionId(v string) *DeleteFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowLogRequest) SetResourceOwnerAccount(v string) *DeleteFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFlowLogRequest) SetResourceOwnerId(v int64) *DeleteFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
