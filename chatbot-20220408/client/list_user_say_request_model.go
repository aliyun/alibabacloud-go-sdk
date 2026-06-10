// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListUserSayRequest
	GetAgentKey() *string
	SetContent(v string) *ListUserSayRequest
	GetContent() *string
	SetInstanceId(v string) *ListUserSayRequest
	GetInstanceId() *string
	SetIntentId(v int64) *ListUserSayRequest
	GetIntentId() *int64
	SetPageNumber(v int32) *ListUserSayRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserSayRequest
	GetPageSize() *int32
}

type ListUserSayRequest struct {
	// The key of the business space. If you omit this parameter, the default business space is used. You can obtain the key from the Business Management page of your Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The keywords used to filter user says. The query returns only user says that contain these keywords.
	//
	// example:
	//
	// 您做核酸了嘛
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the bot.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the intent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 232
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The default value is 10. The maximum value is 1000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserSayRequest) GoString() string {
	return s.String()
}

func (s *ListUserSayRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListUserSayRequest) GetContent() *string {
	return s.Content
}

func (s *ListUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserSayRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListUserSayRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserSayRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSayRequest) SetAgentKey(v string) *ListUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *ListUserSayRequest) SetContent(v string) *ListUserSayRequest {
	s.Content = &v
	return s
}

func (s *ListUserSayRequest) SetInstanceId(v string) *ListUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserSayRequest) SetIntentId(v int64) *ListUserSayRequest {
	s.IntentId = &v
	return s
}

func (s *ListUserSayRequest) SetPageNumber(v int32) *ListUserSayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserSayRequest) SetPageSize(v int32) *ListUserSayRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserSayRequest) Validate() error {
	return dara.Validate(s)
}
