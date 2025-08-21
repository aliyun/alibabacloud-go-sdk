// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListLgfRequest
	GetAgentKey() *string
	SetInstanceId(v string) *ListLgfRequest
	GetInstanceId() *string
	SetIntentId(v int64) *ListLgfRequest
	GetIntentId() *int64
	SetLgfText(v string) *ListLgfRequest
	GetLgfText() *string
	SetPageNumber(v int32) *ListLgfRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLgfRequest
	GetPageSize() *int32
}

type ListLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// .{0,5}北京天气
	LgfText *string `json:"LgfText,omitempty" xml:"LgfText,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLgfRequest) GoString() string {
	return s.String()
}

func (s *ListLgfRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLgfRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListLgfRequest) GetLgfText() *string {
	return s.LgfText
}

func (s *ListLgfRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLgfRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLgfRequest) SetAgentKey(v string) *ListLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *ListLgfRequest) SetInstanceId(v string) *ListLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLgfRequest) SetIntentId(v int64) *ListLgfRequest {
	s.IntentId = &v
	return s
}

func (s *ListLgfRequest) SetLgfText(v string) *ListLgfRequest {
	s.LgfText = &v
	return s
}

func (s *ListLgfRequest) SetPageNumber(v int32) *ListLgfRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLgfRequest) SetPageSize(v int32) *ListLgfRequest {
	s.PageSize = &v
	return s
}

func (s *ListLgfRequest) Validate() error {
	return dara.Validate(s)
}
