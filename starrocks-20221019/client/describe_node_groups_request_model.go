// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeNodeGroupsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeNodeGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNodeGroupsRequest
	GetPageSize() *int32
	SetComponentType(v string) *DescribeNodeGroupsRequest
	GetComponentType() *string
	SetInstanceId(v string) *DescribeNodeGroupsRequest
	GetInstanceId() *string
	SetNodeGroupIds(v []*string) *DescribeNodeGroupsRequest
	GetNodeGroupIds() []*string
	SetNodeGroupName(v string) *DescribeNodeGroupsRequest
	GetNodeGroupName() *string
	SetStatus(v string) *DescribeNodeGroupsRequest
	GetStatus() *string
	SetTags(v []*DescribeNodeGroupsRequestTags) *DescribeNodeGroupsRequest
	GetTags() []*DescribeNodeGroupsRequestTags
}

type DescribeNodeGroupsRequest struct {
	// The instance ID.
	//
	// example:
	//
	// c-718fb04c7112****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number for paging. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The compute group type. Valid values:
	//
	// - FE
	//
	// - BE
	//
	// - CN
	//
	// - OBSERVER
	//
	// - AGENT
	//
	// example:
	//
	// FE
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// **[Deprecated]*	- The instance ID. This parameter is deprecated.
	//
	// example:
	//
	// null
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The list of compute group IDs to query.
	NodeGroupIds []*string `json:"nodeGroupIds,omitempty" xml:"nodeGroupIds,omitempty" type:"Repeated"`
	// The compute group name.
	//
	// example:
	//
	// ng_1
	NodeGroupName *string `json:"nodeGroupName,omitempty" xml:"nodeGroupName,omitempty"`
	// The compute group status. Valid values:
	//
	// - UNPAID: Unpaid.
	//
	// - PAID: Paid.
	//
	// - CREATING_FAILED: Creation failed.
	//
	// - CREATING: Being created.
	//
	// - RUNNING: Running.
	//
	// - MODIFYING_CONFIG: Configuration being modified.
	//
	// - MODIFYING_TIMEZONE: Time zone being modified.
	//
	// - ELASTIC_SCALING_OUT: Elastic scale-out in progress.
	//
	// - ELASTIC_SCALING_IN: Elastic scale-in in progress.
	//
	// - SCALING_OUT: Scale-out in progress.
	//
	// - RESTARTING: Restarting.
	//
	// - SCALING_IN: Scale-in in progress.
	//
	// - SCALING_UP: Upgrading specifications.
	//
	// - SCALING_DOWN: Downgrading specifications.
	//
	// - UPGRADING: Upgrading.
	//
	// - ENABLE_PUBLIC_NETWORK: Enabling public network access.
	//
	// - DISABLE_PUBLIC_NETWORK: Disabling public network access.
	//
	// - SWITCH_AZ: Switching zone.
	//
	// - DISABLE: Disabled.
	//
	// - DELETING: Being deleted.
	//
	// - DELETED: Deleted.
	//
	// example:
	//
	// RUNNING
	Status *string                          `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*DescribeNodeGroupsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s DescribeNodeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNodeGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNodeGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNodeGroupsRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeNodeGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNodeGroupsRequest) GetNodeGroupIds() []*string {
	return s.NodeGroupIds
}

func (s *DescribeNodeGroupsRequest) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeNodeGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeNodeGroupsRequest) GetTags() []*DescribeNodeGroupsRequestTags {
	return s.Tags
}

func (s *DescribeNodeGroupsRequest) SetClusterId(v string) *DescribeNodeGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetPageNumber(v int32) *DescribeNodeGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetPageSize(v int32) *DescribeNodeGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetComponentType(v string) *DescribeNodeGroupsRequest {
	s.ComponentType = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetInstanceId(v string) *DescribeNodeGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetNodeGroupIds(v []*string) *DescribeNodeGroupsRequest {
	s.NodeGroupIds = v
	return s
}

func (s *DescribeNodeGroupsRequest) SetNodeGroupName(v string) *DescribeNodeGroupsRequest {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetStatus(v string) *DescribeNodeGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeNodeGroupsRequest) SetTags(v []*DescribeNodeGroupsRequestTags) *DescribeNodeGroupsRequest {
	s.Tags = v
	return s
}

func (s *DescribeNodeGroupsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeGroupsRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeNodeGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNodeGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNodeGroupsRequestTags) SetKey(v string) *DescribeNodeGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeNodeGroupsRequestTags) SetValue(v string) *DescribeNodeGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeNodeGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
