// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SendMessageToGroupRequest
	GetAppId() *string
	SetData(v string) *SendMessageToGroupRequest
	GetData() *string
	SetGroupId(v string) *SendMessageToGroupRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *SendMessageToGroupRequest
	GetOperatorUserId() *string
	SetSkipAudit(v bool) *SendMessageToGroupRequest
	GetSkipAudit() *bool
	SetType(v int32) *SendMessageToGroupRequest
	GetType() *int32
}

type SendMessageToGroupRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The message body. The value is a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user who performed the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// Specifies whether the message requires Alibaba Cloud content moderation. Valid values:
	//
	// - **true**: does not require content moderation.
	//
	// - **false**: requires content moderation. This is the default value.
	//
	// example:
	//
	// true
	SkipAudit *bool `json:"SkipAudit,omitempty" xml:"SkipAudit,omitempty"`
	// The type of the message. A value that is less than or equal to 10000 specifies a system message. A value that is greater than 10000 specifies a custom message.
	//
	// example:
	//
	// 12000
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *SendMessageToGroupRequest) GetData() *string {
	return s.Data
}

func (s *SendMessageToGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SendMessageToGroupRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *SendMessageToGroupRequest) GetSkipAudit() *bool {
	return s.SkipAudit
}

func (s *SendMessageToGroupRequest) GetType() *int32 {
	return s.Type
}

func (s *SendMessageToGroupRequest) SetAppId(v string) *SendMessageToGroupRequest {
	s.AppId = &v
	return s
}

func (s *SendMessageToGroupRequest) SetData(v string) *SendMessageToGroupRequest {
	s.Data = &v
	return s
}

func (s *SendMessageToGroupRequest) SetGroupId(v string) *SendMessageToGroupRequest {
	s.GroupId = &v
	return s
}

func (s *SendMessageToGroupRequest) SetOperatorUserId(v string) *SendMessageToGroupRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SendMessageToGroupRequest) SetSkipAudit(v bool) *SendMessageToGroupRequest {
	s.SkipAudit = &v
	return s
}

func (s *SendMessageToGroupRequest) SetType(v int32) *SendMessageToGroupRequest {
	s.Type = &v
	return s
}

func (s *SendMessageToGroupRequest) Validate() error {
	return dara.Validate(s)
}
