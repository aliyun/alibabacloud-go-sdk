// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMaliciousNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCreateMaliciousNoteResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BatchCreateMaliciousNoteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchCreateMaliciousNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCreateMaliciousNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCreateMaliciousNoteResponseBody
	GetSuccess() *bool
}

type BatchCreateMaliciousNoteResponseBody struct {
	// The result code. A value of **200*	- indicates success, whereas other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned in the response.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateMaliciousNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMaliciousNoteResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateMaliciousNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCreateMaliciousNoteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchCreateMaliciousNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateMaliciousNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateMaliciousNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateMaliciousNoteResponseBody) SetCode(v string) *BatchCreateMaliciousNoteResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateMaliciousNoteResponseBody) SetHttpStatusCode(v int32) *BatchCreateMaliciousNoteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchCreateMaliciousNoteResponseBody) SetMessage(v string) *BatchCreateMaliciousNoteResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateMaliciousNoteResponseBody) SetRequestId(v string) *BatchCreateMaliciousNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateMaliciousNoteResponseBody) SetSuccess(v bool) *BatchCreateMaliciousNoteResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateMaliciousNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
