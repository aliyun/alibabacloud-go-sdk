// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatbotInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListChatbotInstancesRequest
	GetInstanceId() *string
	SetNluServiceParamsJson(v string) *ListChatbotInstancesRequest
	GetNluServiceParamsJson() *string
	SetNluServiceType(v string) *ListChatbotInstancesRequest
	GetNluServiceType() *string
	SetPageNumber(v int32) *ListChatbotInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListChatbotInstancesRequest
	GetPageSize() *int32
	SetUnionSource(v string) *ListChatbotInstancesRequest
	GetUnionSource() *string
}

type ListChatbotInstancesRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	NluServiceType       *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UnionSource *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListChatbotInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatbotInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChatbotInstancesRequest) GetNluServiceParamsJson() *string {
	return s.NluServiceParamsJson
}

func (s *ListChatbotInstancesRequest) GetNluServiceType() *string {
	return s.NluServiceType
}

func (s *ListChatbotInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListChatbotInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatbotInstancesRequest) GetUnionSource() *string {
	return s.UnionSource
}

func (s *ListChatbotInstancesRequest) SetInstanceId(v string) *ListChatbotInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetNluServiceParamsJson(v string) *ListChatbotInstancesRequest {
	s.NluServiceParamsJson = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetNluServiceType(v string) *ListChatbotInstancesRequest {
	s.NluServiceType = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageNumber(v int32) *ListChatbotInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageSize(v int32) *ListChatbotInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetUnionSource(v string) *ListChatbotInstancesRequest {
	s.UnionSource = &v
	return s
}

func (s *ListChatbotInstancesRequest) Validate() error {
	return dara.Validate(s)
}
