// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageGroupMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteLiveMessageGroupMessageResponseBody
	GetGroupId() *string
	SetMessageId(v string) *DeleteLiveMessageGroupMessageResponseBody
	GetMessageId() *string
	SetRequestId(v string) *DeleteLiveMessageGroupMessageResponseBody
	GetRequestId() *string
}

type DeleteLiveMessageGroupMessageResponseBody struct {
	// The group ID.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the deleted or non-existent message.
	//
	// example:
	//
	// a74a8fbd3cfe4b2daa8517e4e3******
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B95BE680-5A6A-1CAD-8AB1-09DFF5D6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveMessageGroupMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageGroupMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageGroupMessageResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteLiveMessageGroupMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *DeleteLiveMessageGroupMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveMessageGroupMessageResponseBody) SetGroupId(v string) *DeleteLiveMessageGroupMessageResponseBody {
	s.GroupId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageResponseBody) SetMessageId(v string) *DeleteLiveMessageGroupMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageResponseBody) SetRequestId(v string) *DeleteLiveMessageGroupMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
