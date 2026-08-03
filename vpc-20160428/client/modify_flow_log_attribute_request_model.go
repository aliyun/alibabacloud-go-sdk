// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationInterval(v int32) *ModifyFlowLogAttributeRequest
	GetAggregationInterval() *int32
	SetDescription(v string) *ModifyFlowLogAttributeRequest
	GetDescription() *string
	SetFlowLogId(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogName() *string
	SetIpVersion(v string) *ModifyFlowLogAttributeRequest
	GetIpVersion() *string
	SetLogFormat(v string) *ModifyFlowLogAttributeRequest
	GetLogFormat() *string
	SetOwnerAccount(v string) *ModifyFlowLogAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyFlowLogAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyFlowLogAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyFlowLogAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFlowLogAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyFlowLogAttributeRequest struct {
	// The new sampling interval of the flow log. Unit: minutes. Valid values: **1**, **5**, and **10**.
	//
	// example:
	//
	// 1
	AggregationInterval *int32 `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	// The new description of the flow log.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Flowlog.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The flow log ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fl-m5e8vhz2t21sel1nq****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The new name of the flow log.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The IP version of the traffic captured by the flow log.
	//
	// example:
	//
	// IPv4
	IpVersion    *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	LogFormat    *string `json:"LogFormat,omitempty" xml:"LogFormat,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the flow log.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
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

func (s ModifyFlowLogAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeRequest) GetAggregationInterval() *int32 {
	return s.AggregationInterval
}

func (s *ModifyFlowLogAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFlowLogAttributeRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *ModifyFlowLogAttributeRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *ModifyFlowLogAttributeRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ModifyFlowLogAttributeRequest) GetLogFormat() *string {
	return s.LogFormat
}

func (s *ModifyFlowLogAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyFlowLogAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFlowLogAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyFlowLogAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFlowLogAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFlowLogAttributeRequest) SetAggregationInterval(v int32) *ModifyFlowLogAttributeRequest {
	s.AggregationInterval = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDescription(v string) *ModifyFlowLogAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogId(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogName(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetIpVersion(v string) *ModifyFlowLogAttributeRequest {
	s.IpVersion = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetLogFormat(v string) *ModifyFlowLogAttributeRequest {
	s.LogFormat = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetRegionId(v string) *ModifyFlowLogAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetResourceOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetResourceOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) Validate() error {
	return dara.Validate(s)
}
