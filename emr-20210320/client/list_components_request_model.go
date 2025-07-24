// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationNames(v []*string) *ListComponentsRequest
	GetApplicationNames() []*string
	SetClusterId(v string) *ListComponentsRequest
	GetClusterId() *string
	SetComponentNames(v []*string) *ListComponentsRequest
	GetComponentNames() []*string
	SetComponentStates(v []*string) *ListComponentsRequest
	GetComponentStates() []*string
	SetMaxResults(v int32) *ListComponentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListComponentsRequest
	GetRegionId() *string
}

type ListComponentsRequest struct {
	// 应用名称列表。
	//
	// example:
	//
	// 20
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
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 地域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) GetApplicationNames() []*string {
	return s.ApplicationNames
}

func (s *ListComponentsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListComponentsRequest) GetComponentNames() []*string {
	return s.ComponentNames
}

func (s *ListComponentsRequest) GetComponentStates() []*string {
	return s.ComponentStates
}

func (s *ListComponentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComponentsRequest) SetApplicationNames(v []*string) *ListComponentsRequest {
	s.ApplicationNames = v
	return s
}

func (s *ListComponentsRequest) SetClusterId(v string) *ListComponentsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListComponentsRequest) SetComponentNames(v []*string) *ListComponentsRequest {
	s.ComponentNames = v
	return s
}

func (s *ListComponentsRequest) SetComponentStates(v []*string) *ListComponentsRequest {
	s.ComponentStates = v
	return s
}

func (s *ListComponentsRequest) SetMaxResults(v int32) *ListComponentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComponentsRequest) SetNextToken(v string) *ListComponentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListComponentsRequest) SetRegionId(v string) *ListComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListComponentsRequest) Validate() error {
	return dara.Validate(s)
}
