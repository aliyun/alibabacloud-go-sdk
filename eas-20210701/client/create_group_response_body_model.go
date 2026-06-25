// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGroupResponseBody
	GetRequestId() *string
}

type CreateGroupResponseBody struct {
	// The response message.
	//
	// example:
	//
	// Create service group foo successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupResponseBody) SetMessage(v string) *CreateGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
