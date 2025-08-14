// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ModifyFlowLogAttributeRequest
	GetCenId() *string
	SetClientToken(v string) *ModifyFlowLogAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyFlowLogAttributeRequest
	GetDescription() *string
	SetFlowLogId(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogName() *string
	SetInterval(v int64) *ModifyFlowLogAttributeRequest
	GetInterval() *int64
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
	// The CEN instance ID.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, ClientToken is set to the value of RequestId. The value of RequestId for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the flow log.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The new name of the flow log.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The time window for collecting log data. Unit: seconds. Valid values: **60*	- or **600*	- Default value: **600**.
	//
	// example:
	//
	// 600
	Interval     *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the flow log is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s *ModifyFlowLogAttributeRequest) GetCenId() *string {
	return s.CenId
}

func (s *ModifyFlowLogAttributeRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *ModifyFlowLogAttributeRequest) GetInterval() *int64 {
	return s.Interval
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

func (s *ModifyFlowLogAttributeRequest) SetCenId(v string) *ModifyFlowLogAttributeRequest {
	s.CenId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetClientToken(v string) *ModifyFlowLogAttributeRequest {
	s.ClientToken = &v
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

func (s *ModifyFlowLogAttributeRequest) SetInterval(v int64) *ModifyFlowLogAttributeRequest {
	s.Interval = &v
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
