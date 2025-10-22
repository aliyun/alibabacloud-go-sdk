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
}

type DescribeNodeGroupsRequest struct {
	// example:
	//
	// c-718fb04c7112****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FE
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// null
	InstanceId   *string   `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NodeGroupIds []*string `json:"nodeGroupIds,omitempty" xml:"nodeGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// ng_1
	NodeGroupName *string `json:"nodeGroupName,omitempty" xml:"nodeGroupName,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *DescribeNodeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
