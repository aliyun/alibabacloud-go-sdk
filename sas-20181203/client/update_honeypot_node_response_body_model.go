// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHoneypotNodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateHoneypotNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateHoneypotNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHoneypotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHoneypotNodeResponseBody
	GetSuccess() *bool
}

type UpdateHoneypotNodeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 753D92B3-A062-544F-8E7B-C813AA9FA9FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHoneypotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHoneypotNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateHoneypotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHoneypotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHoneypotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHoneypotNodeResponseBody) SetCode(v string) *UpdateHoneypotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHoneypotNodeResponseBody) SetHttpStatusCode(v int32) *UpdateHoneypotNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateHoneypotNodeResponseBody) SetMessage(v string) *UpdateHoneypotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHoneypotNodeResponseBody) SetRequestId(v string) *UpdateHoneypotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHoneypotNodeResponseBody) SetSuccess(v bool) *UpdateHoneypotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHoneypotNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
