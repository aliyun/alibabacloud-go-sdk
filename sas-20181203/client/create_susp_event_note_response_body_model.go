// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSuspEventNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSuspEventNoteResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateSuspEventNoteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSuspEventNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSuspEventNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSuspEventNoteResponseBody
	GetSuccess() *bool
}

type CreateSuspEventNoteResponseBody struct {
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
	// The error message returned.
	//
	// example:
	//
	// There was an error with your request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether exceptions are handled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSuspEventNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSuspEventNoteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSuspEventNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSuspEventNoteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSuspEventNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSuspEventNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSuspEventNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSuspEventNoteResponseBody) SetCode(v string) *CreateSuspEventNoteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSuspEventNoteResponseBody) SetHttpStatusCode(v int32) *CreateSuspEventNoteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSuspEventNoteResponseBody) SetMessage(v string) *CreateSuspEventNoteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSuspEventNoteResponseBody) SetRequestId(v string) *CreateSuspEventNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSuspEventNoteResponseBody) SetSuccess(v bool) *CreateSuspEventNoteResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSuspEventNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
