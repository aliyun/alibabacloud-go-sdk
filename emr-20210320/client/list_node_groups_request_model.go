// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodeGroupsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListNodeGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupsRequest
	GetNextToken() *string
	SetNodeGroupIds(v []*string) *ListNodeGroupsRequest
	GetNodeGroupIds() []*string
	SetNodeGroupNames(v []*string) *ListNodeGroupsRequest
	GetNodeGroupNames() []*string
	SetNodeGroupStates(v []*string) *ListNodeGroupsRequest
	GetNodeGroupStates() []*string
	SetNodeGroupTypes(v []*string) *ListNodeGroupsRequest
	GetNodeGroupTypes() []*string
	SetRegionId(v string) *ListNodeGroupsRequest
	GetRegionId() *string
	SetZoneId(v string) *ListNodeGroupsRequest
	GetZoneId() *string
}

type ListNodeGroupsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of maximum number of records to obtain at a time. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Marks the current position where reading starts. If you set this value to null, you can start from the beginning.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of node groups. Valid values of the number of array elements N: 1 to 100.
	//
	// example:
	//
	// c-b933c5aac8fe****
	NodeGroupIds []*string `json:"NodeGroupIds,omitempty" xml:"NodeGroupIds,omitempty" type:"Repeated"`
	// The list of node group names. Valid values of the number of array elements N: 1 to 100.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NodeGroupNames []*string `json:"NodeGroupNames,omitempty" xml:"NodeGroupNames,omitempty" type:"Repeated"`
	// The status of the node group. Valid values of the number of array elements N: 1 to 100.
	//
	// example:
	//
	// ["CORE"]
	NodeGroupStates []*string `json:"NodeGroupStates,omitempty" xml:"NodeGroupStates,omitempty" type:"Repeated"`
	// The list of node group types. Valid values of the number of array elements N: 1 to 100.
	//
	// example:
	//
	// 20
	NodeGroupTypes []*string `json:"NodeGroupTypes,omitempty" xml:"NodeGroupTypes,omitempty" type:"Repeated"`
	// The ID of the region in which you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupsRequest) GetNodeGroupIds() []*string {
	return s.NodeGroupIds
}

func (s *ListNodeGroupsRequest) GetNodeGroupNames() []*string {
	return s.NodeGroupNames
}

func (s *ListNodeGroupsRequest) GetNodeGroupStates() []*string {
	return s.NodeGroupStates
}

func (s *ListNodeGroupsRequest) GetNodeGroupTypes() []*string {
	return s.NodeGroupTypes
}

func (s *ListNodeGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNodeGroupsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListNodeGroupsRequest) SetClusterId(v string) *ListNodeGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetMaxResults(v int32) *ListNodeGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNextToken(v string) *ListNodeGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupIds(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupIds = v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupNames(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupNames = v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupStates(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupStates = v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupTypes(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupTypes = v
	return s
}

func (s *ListNodeGroupsRequest) SetRegionId(v string) *ListNodeGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetZoneId(v string) *ListNodeGroupsRequest {
	s.ZoneId = &v
	return s
}

func (s *ListNodeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
