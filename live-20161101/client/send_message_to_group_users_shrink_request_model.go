// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SendMessageToGroupUsersShrinkRequest
	GetAppId() *string
	SetData(v string) *SendMessageToGroupUsersShrinkRequest
	GetData() *string
	SetGroupId(v string) *SendMessageToGroupUsersShrinkRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *SendMessageToGroupUsersShrinkRequest
	GetOperatorUserId() *string
	SetReceiverIdListShrink(v string) *SendMessageToGroupUsersShrinkRequest
	GetReceiverIdListShrink() *string
	SetSkipAudit(v bool) *SendMessageToGroupUsersShrinkRequest
	GetSkipAudit() *bool
	SetType(v int32) *SendMessageToGroupUsersShrinkRequest
	GetType() *int32
}

type SendMessageToGroupUsersShrinkRequest struct {
	// Interactive Messages application
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Message body in JSONString format.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Message group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Operator user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// User list.
	ReceiverIdListShrink *string `json:"ReceiverIdList,omitempty" xml:"ReceiverIdList,omitempty"`
	// Specifies whether the current message content requires Content Moderation by Alibaba Cloud. Valid values:
	//
	// - **true**: Content Moderation is not required.
	//
	// - **false*	- (default): Content Moderation is required.
	//
	// example:
	//
	// true
	SkipAudit *bool `json:"SkipAudit,omitempty" xml:"SkipAudit,omitempty"`
	// Message type. When the type field value is less than or equal to 10000, it indicates a system message. When the value is greater than 10000, it indicates a custom message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12000
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageToGroupUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupUsersShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *SendMessageToGroupUsersShrinkRequest) GetData() *string {
	return s.Data
}

func (s *SendMessageToGroupUsersShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SendMessageToGroupUsersShrinkRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *SendMessageToGroupUsersShrinkRequest) GetReceiverIdListShrink() *string {
	return s.ReceiverIdListShrink
}

func (s *SendMessageToGroupUsersShrinkRequest) GetSkipAudit() *bool {
	return s.SkipAudit
}

func (s *SendMessageToGroupUsersShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *SendMessageToGroupUsersShrinkRequest) SetAppId(v string) *SendMessageToGroupUsersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) SetData(v string) *SendMessageToGroupUsersShrinkRequest {
	s.Data = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) SetGroupId(v string) *SendMessageToGroupUsersShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) SetOperatorUserId(v string) *SendMessageToGroupUsersShrinkRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) SetReceiverIdListShrink(v string) *SendMessageToGroupUsersShrinkRequest {
	s.ReceiverIdListShrink = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) SetSkipAudit(v bool) *SendMessageToGroupUsersShrinkRequest {
	s.SkipAudit = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) SetType(v int32) *SendMessageToGroupUsersShrinkRequest {
	s.Type = &v
	return s
}

func (s *SendMessageToGroupUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
