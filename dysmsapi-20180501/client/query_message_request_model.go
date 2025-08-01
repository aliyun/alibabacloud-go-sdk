// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *QueryMessageRequest
	GetMessageId() *string
}

type QueryMessageRequest struct {
	// The ID of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1008030xxx3003
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMessageRequest) SetMessageId(v string) *QueryMessageRequest {
	s.MessageId = &v
	return s
}

func (s *QueryMessageRequest) Validate() error {
	return dara.Validate(s)
}
