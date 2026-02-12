// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessageGetByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsDLQMessageGetByIdRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsDLQMessageGetByIdRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsDLQMessageGetByIdRequest
	GetMsgId() *string
}

type OnsDLQMessageGetByIdRequest struct {
	// The ID of the consumer group whose dead-letter message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the dead-letter message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0BC16699165C03B925DB8A404E2D****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s OnsDLQMessageGetByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageGetByIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsDLQMessageGetByIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsDLQMessageGetByIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsDLQMessageGetByIdRequest) SetGroupId(v string) *OnsDLQMessageGetByIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) SetInstanceId(v string) *OnsDLQMessageGetByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) SetMsgId(v string) *OnsDLQMessageGetByIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) Validate() error {
	return dara.Validate(s)
}
