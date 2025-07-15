// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePersonalNumbersFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemovePersonalNumbersFromUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemovePersonalNumbersFromUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemovePersonalNumbersFromUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemovePersonalNumbersFromUserResponseBody
	GetRequestId() *string
}

type RemovePersonalNumbersFromUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePersonalNumbersFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePersonalNumbersFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePersonalNumbersFromUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemovePersonalNumbersFromUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemovePersonalNumbersFromUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemovePersonalNumbersFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetCode(v string) *RemovePersonalNumbersFromUserResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetHttpStatusCode(v int32) *RemovePersonalNumbersFromUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetMessage(v string) *RemovePersonalNumbersFromUserResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetRequestId(v string) *RemovePersonalNumbersFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}
