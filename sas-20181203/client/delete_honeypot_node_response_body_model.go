// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHoneypotNodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteHoneypotNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteHoneypotNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHoneypotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHoneypotNodeResponseBody
	GetSuccess() *bool
}

type DeleteHoneypotNodeResponseBody struct {
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
	// 49FDE92F-A0B8-56CC-B7A8-23B17646CCAD
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

func (s DeleteHoneypotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHoneypotNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteHoneypotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHoneypotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHoneypotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHoneypotNodeResponseBody) SetCode(v string) *DeleteHoneypotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHoneypotNodeResponseBody) SetHttpStatusCode(v int32) *DeleteHoneypotNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteHoneypotNodeResponseBody) SetMessage(v string) *DeleteHoneypotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHoneypotNodeResponseBody) SetRequestId(v string) *DeleteHoneypotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHoneypotNodeResponseBody) SetSuccess(v bool) *DeleteHoneypotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHoneypotNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
