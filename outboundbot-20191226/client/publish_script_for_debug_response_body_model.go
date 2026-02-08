// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptForDebugResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishScriptForDebugResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PublishScriptForDebugResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishScriptForDebugResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishScriptForDebugResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishScriptForDebugResponseBody
	GetSuccess() *bool
}

type PublishScriptForDebugResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5a865b03-d2b9-4ef9-be98-f21fa0d93744
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishScriptForDebugResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptForDebugResponseBody) GoString() string {
	return s.String()
}

func (s *PublishScriptForDebugResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishScriptForDebugResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishScriptForDebugResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishScriptForDebugResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishScriptForDebugResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishScriptForDebugResponseBody) SetCode(v string) *PublishScriptForDebugResponseBody {
	s.Code = &v
	return s
}

func (s *PublishScriptForDebugResponseBody) SetHttpStatusCode(v int32) *PublishScriptForDebugResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishScriptForDebugResponseBody) SetMessage(v string) *PublishScriptForDebugResponseBody {
	s.Message = &v
	return s
}

func (s *PublishScriptForDebugResponseBody) SetRequestId(v string) *PublishScriptForDebugResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishScriptForDebugResponseBody) SetSuccess(v bool) *PublishScriptForDebugResponseBody {
	s.Success = &v
	return s
}

func (s *PublishScriptForDebugResponseBody) Validate() error {
	return dara.Validate(s)
}
