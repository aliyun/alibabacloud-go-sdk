// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessageResendByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsDLQMessageResendByIdRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsDLQMessageResendByIdRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsDLQMessageResendByIdRequest
	GetMsgId() *string
}

type OnsDLQMessageResendByIdRequest struct {
	// The ID of the consumer group in which you want to query dead-letter messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance in which the dead-letter message you want to query resides.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the dead-letter message that you want to send to a consumer group for consumption.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0BC16699343051CD9F1D798E7734****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s OnsDLQMessageResendByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageResendByIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsDLQMessageResendByIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsDLQMessageResendByIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsDLQMessageResendByIdRequest) SetGroupId(v string) *OnsDLQMessageResendByIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) SetInstanceId(v string) *OnsDLQMessageResendByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) SetMsgId(v string) *OnsDLQMessageResendByIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) Validate() error {
	return dara.Validate(s)
}
