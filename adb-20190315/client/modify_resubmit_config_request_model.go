// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResubmitConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyResubmitConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyResubmitConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyResubmitConfigRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyResubmitConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyResubmitConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyResubmitConfigRequest
	GetResourceOwnerId() *int64
	SetRules(v []*ModifyResubmitConfigRequestRules) *ModifyResubmitConfigRequest
	GetRules() []*ModifyResubmitConfigRequestRules
}

type ModifyResubmitConfigRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6wjk5xxxxxxxxxx
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The job resubmission rules.
	//
	// This parameter is required.
	Rules []*ModifyResubmitConfigRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyResubmitConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResubmitConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyResubmitConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyResubmitConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyResubmitConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyResubmitConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyResubmitConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyResubmitConfigRequest) GetRules() []*ModifyResubmitConfigRequestRules {
	return s.Rules
}

func (s *ModifyResubmitConfigRequest) SetDBClusterId(v string) *ModifyResubmitConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetOwnerAccount(v string) *ModifyResubmitConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetOwnerId(v int64) *ModifyResubmitConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetResourceGroupId(v string) *ModifyResubmitConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetResourceOwnerAccount(v string) *ModifyResubmitConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetResourceOwnerId(v int64) *ModifyResubmitConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetRules(v []*ModifyResubmitConfigRequestRules) *ModifyResubmitConfigRequest {
	s.Rules = v
	return s
}

func (s *ModifyResubmitConfigRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyResubmitConfigRequestRules struct {
	// Specifies whether to configure out-of-memory (OOM) check.
	//
	// example:
	//
	// false
	ExceedMemoryException *bool `json:"ExceedMemoryException,omitempty" xml:"ExceedMemoryException,omitempty"`
	// The name of the source resource group.
	//
	// example:
	//
	// test2
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The peak memory usage.
	//
	// example:
	//
	// 32
	PeakMemory *string `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The duration of the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 300
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The name of the destination resource group.
	//
	// example:
	//
	// test_target_group
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s ModifyResubmitConfigRequestRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyResubmitConfigRequestRules) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigRequestRules) GetExceedMemoryException() *bool {
	return s.ExceedMemoryException
}

func (s *ModifyResubmitConfigRequestRules) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyResubmitConfigRequestRules) GetPeakMemory() *string {
	return s.PeakMemory
}

func (s *ModifyResubmitConfigRequestRules) GetQueryTime() *string {
	return s.QueryTime
}

func (s *ModifyResubmitConfigRequestRules) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *ModifyResubmitConfigRequestRules) SetExceedMemoryException(v bool) *ModifyResubmitConfigRequestRules {
	s.ExceedMemoryException = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetGroupName(v string) *ModifyResubmitConfigRequestRules {
	s.GroupName = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetPeakMemory(v string) *ModifyResubmitConfigRequestRules {
	s.PeakMemory = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetQueryTime(v string) *ModifyResubmitConfigRequestRules {
	s.QueryTime = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetTargetGroupName(v string) *ModifyResubmitConfigRequestRules {
	s.TargetGroupName = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) Validate() error {
	return dara.Validate(s)
}
