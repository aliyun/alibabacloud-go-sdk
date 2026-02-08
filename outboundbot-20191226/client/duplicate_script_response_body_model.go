// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDuplicateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DuplicateScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DuplicateScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DuplicateScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *DuplicateScriptResponseBody
	GetRequestId() *string
	SetScriptId(v string) *DuplicateScriptResponseBody
	GetScriptId() *string
	SetSuccess(v bool) *DuplicateScriptResponseBody
	GetSuccess() *bool
}

type DuplicateScriptResponseBody struct {
	// Status code
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
	// Request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 6114e7e8-4140-48d9-b46d-65ea29f13fe8
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DuplicateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DuplicateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DuplicateScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *DuplicateScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DuplicateScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DuplicateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DuplicateScriptResponseBody) GetScriptId() *string {
	return s.ScriptId
}

func (s *DuplicateScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DuplicateScriptResponseBody) SetCode(v string) *DuplicateScriptResponseBody {
	s.Code = &v
	return s
}

func (s *DuplicateScriptResponseBody) SetHttpStatusCode(v int32) *DuplicateScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DuplicateScriptResponseBody) SetMessage(v string) *DuplicateScriptResponseBody {
	s.Message = &v
	return s
}

func (s *DuplicateScriptResponseBody) SetRequestId(v string) *DuplicateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DuplicateScriptResponseBody) SetScriptId(v string) *DuplicateScriptResponseBody {
	s.ScriptId = &v
	return s
}

func (s *DuplicateScriptResponseBody) SetSuccess(v bool) *DuplicateScriptResponseBody {
	s.Success = &v
	return s
}

func (s *DuplicateScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
