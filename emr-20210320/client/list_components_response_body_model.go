// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*ListComponentsResponseBodyComponents) *ListComponentsResponseBody
	GetComponents() []*ListComponentsResponseBodyComponents
	SetMaxResults(v int32) *ListComponentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListComponentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListComponentsResponseBody
	GetTotalCount() *int32
}

type ListComponentsResponseBody struct {
	Components []*ListComponentsResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 返回读取到的数据位置，空代表数据已经读取完毕。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// FFAC608A-5DC3-174F-93C6-9F88CA6D5875
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) GetComponents() []*ListComponentsResponseBodyComponents {
	return s.Components
}

func (s *ListComponentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComponentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComponentsResponseBody) SetComponents(v []*ListComponentsResponseBodyComponents) *ListComponentsResponseBody {
	s.Components = v
	return s
}

func (s *ListComponentsResponseBody) SetMaxResults(v int32) *ListComponentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListComponentsResponseBody) SetNextToken(v string) *ListComponentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListComponentsResponseBody) SetRequestId(v string) *ListComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentsResponseBody) SetTotalCount(v int32) *ListComponentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListComponentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListComponentsResponseBodyComponents struct {
	// 应用名称。
	//
	// example:
	//
	// KNOX
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 属性列表。
	Attributes []*Attribute `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// 组件名称。
	//
	// example:
	//
	// KNOX
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 命名空间。
	//
	// example:
	//
	// “”
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 安装该组件的机器总数。
	//
	// example:
	//
	// 1
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
}

func (s ListComponentsResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponents) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListComponentsResponseBodyComponents) GetAttributes() []*Attribute {
	return s.Attributes
}

func (s *ListComponentsResponseBodyComponents) GetComponentName() *string {
	return s.ComponentName
}

func (s *ListComponentsResponseBodyComponents) GetNamespace() *string {
	return s.Namespace
}

func (s *ListComponentsResponseBodyComponents) GetReplica() *int32 {
	return s.Replica
}

func (s *ListComponentsResponseBodyComponents) SetApplicationName(v string) *ListComponentsResponseBodyComponents {
	s.ApplicationName = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetAttributes(v []*Attribute) *ListComponentsResponseBodyComponents {
	s.Attributes = v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetComponentName(v string) *ListComponentsResponseBodyComponents {
	s.ComponentName = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetNamespace(v string) *ListComponentsResponseBodyComponents {
	s.Namespace = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetReplica(v int32) *ListComponentsResponseBodyComponents {
	s.Replica = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) Validate() error {
	return dara.Validate(s)
}
