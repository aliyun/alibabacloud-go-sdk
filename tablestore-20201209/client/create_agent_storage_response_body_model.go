// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentStorageResponseBody
	GetCode() *string
	SetMessage(v string) *CreateAgentStorageResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentStorageResponseBody
	GetRequestId() *string
}

type CreateAgentStorageResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// request id
	//
	// example:
	//
	// 18DD77BF-F967-576D-80D1-79121399AB53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentStorageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentStorageResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentStorageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentStorageResponseBody) SetCode(v string) *CreateAgentStorageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentStorageResponseBody) SetMessage(v string) *CreateAgentStorageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentStorageResponseBody) SetRequestId(v string) *CreateAgentStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
