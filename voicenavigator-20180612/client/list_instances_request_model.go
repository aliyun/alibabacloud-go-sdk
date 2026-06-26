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
	// A JSON-formatted string that contains a list of digital worker instance IDs.
	//
	// example:
	//
	// ["e4bebxxxxxxxxxxx"]
	InstanceIdListJsonString *string `json:"InstanceIdListJsonString,omitempty" xml:"InstanceIdListJsonString,omitempty"`
	// The instance name. This parameter is used for filtering.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The NLU service type. This parameter is used to filter instances by the source of their conversational AI capabilities. If you do not set this parameter, instances of all types are returned.
	//
	// - `MANAGED`: managed. This value is deprecated.
	//
	// - `AUTHORIZED`: authorized. In the public cloud, this indicates the Chatbot service.
	//
	// - `PROVIDED`: private. This service is configured in the console with parameters such as `as`, `sk`, and `chatEndpoint`.
	//
	// - `CCC_AUTHORIZED`: a chatbot authorized by Cloud Connect Center (CCC).
	//
	// - `CCC_FUNCTION`: Alibaba Cloud Function Compute.
	//
	// - `SSE_FUNCTION`: a streaming function service. This refers to a Function Compute instance that supports Server-Sent Events (SSE) for integration with third-party large language model (LLM) chatbots.
	//
	// - `PROMPTS`: integration with foundational models such as Qwen.
	//
	// - `LOCAL`: a private cloud instance of Chatbot.
	//
	// example:
	//
	// ["CCC_AUTHORIZED"]
	NluServiceTypeListJsonString *string `json:"NluServiceTypeListJsonString,omitempty" xml:"NluServiceTypeListJsonString,omitempty"`
	// The inbound number. This parameter is used for filtering.
	//
	// example:
	//
	// 021xxxxxxx
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance status. This parameter is used for filtering. If you do not set this parameter, instances in all statuses are returned.
	//
	// - `DISABLED`: disabled
	//
	// - `PUBLISHED`: published
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance ID.
	//
	// > If you set `UnionSource` to `CCC`, set this parameter to the ID of your CCC instance.
	//
	// example:
	//
	// zhyl
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	// The source.
	//
	// - `CCC`: Cloud Connect Center
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
