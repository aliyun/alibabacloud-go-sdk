// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdListJsonString(v string) *ListInstancesRequest
	GetInstanceIdListJsonString() *string
	SetName(v string) *ListInstancesRequest
	GetName() *string
	SetNluServiceTypeListJsonString(v string) *ListInstancesRequest
	GetNluServiceTypeListJsonString() *string
	SetNumber(v string) *ListInstancesRequest
	GetNumber() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
	SetUnionInstanceId(v string) *ListInstancesRequest
	GetUnionInstanceId() *string
	SetUnionSource(v string) *ListInstancesRequest
	GetUnionSource() *string
}

type ListInstancesRequest struct {
	// The list of digital employee scenario IDs.
	//
	// example:
	//
	// ["e4bebxxxxxxxxxxx"]
	InstanceIdListJsonString *string `json:"InstanceIdListJsonString,omitempty" xml:"InstanceIdListJsonString,omitempty"`
	// The scenario name used as a filter condition.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The NLU type used to filter by dialog capability source. If this parameter is not specified, all types are selected.
	//
	//  	- MANAGED: managed (deprecated).
	//
	//  	- AUTHORIZED: authorized. In the public cloud, this refers to Chatbot service.
	//
	//  	- PROVIDED: private. Configured by specifying the AccessKey ID, AccessKey Secret, and chatEndpoint in the O&M console.
	//
	// 	- CCC_AUTHORIZED: uses a Chatbot authorized by Cloud Call Center.
	//
	// 	- CCC_FUNCTION: uses Alibaba Cloud Function Compute.
	//
	// 	- SSE_FUNCTION: uses a streaming function service. Function Compute that supports SSE, used to connect to third-party large language model chatbots.
	//
	//
	// 	- PROMPTS: connects to Qwen foundation models.
	//
	// 	- LOCAL: private cloud, local Chatbot.
	//
	// example:
	//
	// ["CCC_AUTHORIZED"]
	NluServiceTypeListJsonString *string `json:"NluServiceTypeListJsonString,omitempty" xml:"NluServiceTypeListJsonString,omitempty"`
	// The inbound number used as a filter condition.
	//
	// example:
	//
	// 021xxxxxxx
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scenario status used as a filter condition. If this parameter is not specified, all statuses are selected.
	//
	// 	- DISABLED: offline.
	//
	// 	- PUBLISHED: published.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance ID.
	//
	// > When UnionSource is set to CCC, set UnionInstanceId to the instance ID of Cloud Call Center.
	//
	// example:
	//
	// zhyl
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	// The source.
	//
	// 	- CCC: Cloud Call Center.
	//
	// example:
	//
	// CCC
	UnionSource *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetInstanceIdListJsonString() *string {
	return s.InstanceIdListJsonString
}

func (s *ListInstancesRequest) GetName() *string {
	return s.Name
}

func (s *ListInstancesRequest) GetNluServiceTypeListJsonString() *string {
	return s.NluServiceTypeListJsonString
}

func (s *ListInstancesRequest) GetNumber() *string {
	return s.Number
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) GetUnionInstanceId() *string {
	return s.UnionInstanceId
}

func (s *ListInstancesRequest) GetUnionSource() *string {
	return s.UnionSource
}

func (s *ListInstancesRequest) SetInstanceIdListJsonString(v string) *ListInstancesRequest {
	s.InstanceIdListJsonString = &v
	return s
}

func (s *ListInstancesRequest) SetName(v string) *ListInstancesRequest {
	s.Name = &v
	return s
}

func (s *ListInstancesRequest) SetNluServiceTypeListJsonString(v string) *ListInstancesRequest {
	s.NluServiceTypeListJsonString = &v
	return s
}

func (s *ListInstancesRequest) SetNumber(v string) *ListInstancesRequest {
	s.Number = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetUnionInstanceId(v string) *ListInstancesRequest {
	s.UnionInstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetUnionSource(v string) *ListInstancesRequest {
	s.UnionSource = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
