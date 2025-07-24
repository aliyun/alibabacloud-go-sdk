// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationNames(v []*string) *ListComponentInstancesRequest
	GetApplicationNames() []*string
	SetClusterId(v string) *ListComponentInstancesRequest
	GetClusterId() *string
	SetComponentNames(v []*string) *ListComponentInstancesRequest
	GetComponentNames() []*string
	SetComponentStates(v []*string) *ListComponentInstancesRequest
	GetComponentStates() []*string
	SetMaxResults(v int32) *ListComponentInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentInstancesRequest
	GetNextToken() *string
	SetNodeIds(v []*string) *ListComponentInstancesRequest
	GetNodeIds() []*string
	SetNodeNames(v []*string) *ListComponentInstancesRequest
	GetNodeNames() []*string
	SetRegionId(v string) *ListComponentInstancesRequest
	GetRegionId() *string
	SetZoneId(v string) *ListComponentInstancesRequest
	GetZoneId() *string
}

type ListComponentInstancesRequest struct {
	// 应用名称列表。
	//
	// example:
	//
	// c-b933c5aac8fe****
	ApplicationNames []*string `json:"ApplicationNames,omitempty" xml:"ApplicationNames,omitempty" type:"Repeated"`
	// 集群ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// C-8CFEBCCFFEF5****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 组件名称列表。
	//
	// example:
	//
	// ["HDFS"]
	ComponentNames []*string `json:"ComponentNames,omitempty" xml:"ComponentNames,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ComponentStates []*string `json:"ComponentStates,omitempty" xml:"ComponentStates,omitempty" type:"Repeated"`
	// 一次获取的最大记录数。取值范围：1~100。
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// “”
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 节点ID列表。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// 节点名称列表。
	//
	// example:
	//
	// 20
	NodeNames []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	// 地域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListComponentInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListComponentInstancesRequest) GetApplicationNames() []*string {
	return s.ApplicationNames
}

func (s *ListComponentInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListComponentInstancesRequest) GetComponentNames() []*string {
	return s.ComponentNames
}

func (s *ListComponentInstancesRequest) GetComponentStates() []*string {
	return s.ComponentStates
}

func (s *ListComponentInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentInstancesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *ListComponentInstancesRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *ListComponentInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComponentInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListComponentInstancesRequest) SetApplicationNames(v []*string) *ListComponentInstancesRequest {
	s.ApplicationNames = v
	return s
}

func (s *ListComponentInstancesRequest) SetClusterId(v string) *ListComponentInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListComponentInstancesRequest) SetComponentNames(v []*string) *ListComponentInstancesRequest {
	s.ComponentNames = v
	return s
}

func (s *ListComponentInstancesRequest) SetComponentStates(v []*string) *ListComponentInstancesRequest {
	s.ComponentStates = v
	return s
}

func (s *ListComponentInstancesRequest) SetMaxResults(v int32) *ListComponentInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComponentInstancesRequest) SetNextToken(v string) *ListComponentInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListComponentInstancesRequest) SetNodeIds(v []*string) *ListComponentInstancesRequest {
	s.NodeIds = v
	return s
}

func (s *ListComponentInstancesRequest) SetNodeNames(v []*string) *ListComponentInstancesRequest {
	s.NodeNames = v
	return s
}

func (s *ListComponentInstancesRequest) SetRegionId(v string) *ListComponentInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListComponentInstancesRequest) SetZoneId(v string) *ListComponentInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListComponentInstancesRequest) Validate() error {
	return dara.Validate(s)
}
