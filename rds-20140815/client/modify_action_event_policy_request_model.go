// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActionEventPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableEventLog(v string) *ModifyActionEventPolicyRequest
	GetEnableEventLog() *string
	SetOwnerId(v int64) *ModifyActionEventPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyActionEventPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyActionEventPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActionEventPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyActionEventPolicyRequest struct {
	// Specifies whether to enable the event history feature. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// This parameter is required.
	//
	// example:
	//
	// True
	EnableEventLog *string `json:"EnableEventLog,omitempty" xml:"EnableEventLog,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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

func (s ModifyActionEventPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActionEventPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyActionEventPolicyRequest) GetEnableEventLog() *string {
	return s.EnableEventLog
}

func (s *ModifyActionEventPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActionEventPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActionEventPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActionEventPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActionEventPolicyRequest) SetEnableEventLog(v string) *ModifyActionEventPolicyRequest {
	s.EnableEventLog = &v
	return s
}

func (s *ModifyActionEventPolicyRequest) SetOwnerId(v int64) *ModifyActionEventPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActionEventPolicyRequest) SetRegionId(v string) *ModifyActionEventPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActionEventPolicyRequest) SetResourceOwnerAccount(v string) *ModifyActionEventPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActionEventPolicyRequest) SetResourceOwnerId(v int64) *ModifyActionEventPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActionEventPolicyRequest) Validate() error {
	return dara.Validate(s)
}
