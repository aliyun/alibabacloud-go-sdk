// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDSEntityValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListDSEntityValueRequest
	GetAgentKey() *string
	SetEntityId(v int64) *ListDSEntityValueRequest
	GetEntityId() *int64
	SetEntityValueId(v int64) *ListDSEntityValueRequest
	GetEntityValueId() *int64
	SetInstanceId(v string) *ListDSEntityValueRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListDSEntityValueRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListDSEntityValueRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDSEntityValueRequest
	GetPageSize() *int32
}

type ListDSEntityValueRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 234
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDSEntityValueRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListDSEntityValueRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListDSEntityValueRequest) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *ListDSEntityValueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDSEntityValueRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDSEntityValueRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDSEntityValueRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDSEntityValueRequest) SetAgentKey(v string) *ListDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDSEntityValueRequest) SetEntityId(v int64) *ListDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *ListDSEntityValueRequest) SetEntityValueId(v int64) *ListDSEntityValueRequest {
	s.EntityValueId = &v
	return s
}

func (s *ListDSEntityValueRequest) SetInstanceId(v string) *ListDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDSEntityValueRequest) SetKeyword(v string) *ListDSEntityValueRequest {
	s.Keyword = &v
	return s
}

func (s *ListDSEntityValueRequest) SetPageNumber(v int32) *ListDSEntityValueRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityValueRequest) SetPageSize(v int32) *ListDSEntityValueRequest {
	s.PageSize = &v
	return s
}

func (s *ListDSEntityValueRequest) Validate() error {
	return dara.Validate(s)
}
