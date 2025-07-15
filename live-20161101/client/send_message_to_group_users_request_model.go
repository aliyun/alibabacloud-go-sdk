// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SendMessageToGroupUsersRequest
	GetAppId() *string
	SetData(v string) *SendMessageToGroupUsersRequest
	GetData() *string
	SetGroupId(v string) *SendMessageToGroupUsersRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *SendMessageToGroupUsersRequest
	GetOperatorUserId() *string
	SetReceiverIdList(v []*string) *SendMessageToGroupUsersRequest
	GetReceiverIdList() []*string
	SetSkipAudit(v bool) *SendMessageToGroupUsersRequest
	GetSkipAudit() *bool
	SetType(v int32) *SendMessageToGroupUsersRequest
	GetType() *int32
}

type SendMessageToGroupUsersRequest struct {
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
	// The ID of the user who performs the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// The list of users to receive the message.
	ReceiverIdList []*string `json:"ReceiverIdList,omitempty" xml:"ReceiverIdList,omitempty" type:"Repeated"`
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
	// This parameter is required.
	//
	// example:
	//
	// 12000
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageToGroupUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupUsersRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *SendMessageToGroupUsersRequest) GetData() *string {
	return s.Data
}

func (s *SendMessageToGroupUsersRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SendMessageToGroupUsersRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *SendMessageToGroupUsersRequest) GetReceiverIdList() []*string {
	return s.ReceiverIdList
}

func (s *SendMessageToGroupUsersRequest) GetSkipAudit() *bool {
	return s.SkipAudit
}

func (s *SendMessageToGroupUsersRequest) GetType() *int32 {
	return s.Type
}

func (s *SendMessageToGroupUsersRequest) SetAppId(v string) *SendMessageToGroupUsersRequest {
	s.AppId = &v
	return s
}

func (s *SendMessageToGroupUsersRequest) SetData(v string) *SendMessageToGroupUsersRequest {
	s.Data = &v
	return s
}

func (s *SendMessageToGroupUsersRequest) SetGroupId(v string) *SendMessageToGroupUsersRequest {
	s.GroupId = &v
	return s
}

func (s *SendMessageToGroupUsersRequest) SetOperatorUserId(v string) *SendMessageToGroupUsersRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SendMessageToGroupUsersRequest) SetReceiverIdList(v []*string) *SendMessageToGroupUsersRequest {
	s.ReceiverIdList = v
	return s
}

func (s *SendMessageToGroupUsersRequest) SetSkipAudit(v bool) *SendMessageToGroupUsersRequest {
	s.SkipAudit = &v
	return s
}

func (s *SendMessageToGroupUsersRequest) SetType(v int32) *SendMessageToGroupUsersRequest {
	s.Type = &v
	return s
}

func (s *SendMessageToGroupUsersRequest) Validate() error {
	return dara.Validate(s)
}
