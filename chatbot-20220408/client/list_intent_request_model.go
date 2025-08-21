// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListIntentRequest
	GetAgentKey() *string
	SetInstanceId(v string) *ListIntentRequest
	GetInstanceId() *string
	SetIntentName(v string) *ListIntentRequest
	GetIntentName() *string
	SetPageNumber(v int32) *ListIntentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIntentRequest
	GetPageSize() *int32
}

type ListIntentRequest struct {
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
	// example:
	//
	// 查天气
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntentRequest) GoString() string {
	return s.String()
}

func (s *ListIntentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntentRequest) GetIntentName() *string {
	return s.IntentName
}

func (s *ListIntentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIntentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentRequest) SetAgentKey(v string) *ListIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *ListIntentRequest) SetInstanceId(v string) *ListIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntentRequest) SetIntentName(v string) *ListIntentRequest {
	s.IntentName = &v
	return s
}

func (s *ListIntentRequest) SetPageNumber(v int32) *ListIntentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIntentRequest) SetPageSize(v int32) *ListIntentRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntentRequest) Validate() error {
	return dara.Validate(s)
}
