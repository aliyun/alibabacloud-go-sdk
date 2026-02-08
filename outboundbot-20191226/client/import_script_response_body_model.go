// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ImportScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportScriptResponseBody
	GetRequestId() *string
	SetScriptId(v string) *ImportScriptResponseBody
	GetScriptId() *string
	SetSuccess(v bool) *ImportScriptResponseBody
	GetSuccess() *bool
}

type ImportScriptResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Outbound scenario ID
	//
	// example:
	//
	// d7fbd0a0-27bc-49c4-a456-ecb75e79122b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Indicates whether the invocation succeeded.
	//
	// - **true**: The invocation succeeded.
	//
	// - **false**: Failed to invoke.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ImportScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportScriptResponseBody) GetScriptId() *string {
	return s.ScriptId
}

func (s *ImportScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportScriptResponseBody) SetCode(v string) *ImportScriptResponseBody {
	s.Code = &v
	return s
}

func (s *ImportScriptResponseBody) SetHttpStatusCode(v int32) *ImportScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportScriptResponseBody) SetMessage(v string) *ImportScriptResponseBody {
	s.Message = &v
	return s
}

func (s *ImportScriptResponseBody) SetRequestId(v string) *ImportScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportScriptResponseBody) SetScriptId(v string) *ImportScriptResponseBody {
	s.ScriptId = &v
	return s
}

func (s *ImportScriptResponseBody) SetSuccess(v bool) *ImportScriptResponseBody {
	s.Success = &v
	return s
}

func (s *ImportScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
