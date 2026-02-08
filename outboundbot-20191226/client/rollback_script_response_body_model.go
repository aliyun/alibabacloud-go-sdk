// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RollbackScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RollbackScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RollbackScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *RollbackScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RollbackScriptResponseBody
	GetSuccess() *bool
}

type RollbackScriptResponseBody struct {
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
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackScriptResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *RollbackScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RollbackScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RollbackScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackScriptResponseBody) SetCode(v string) *RollbackScriptResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackScriptResponseBody) SetHttpStatusCode(v int32) *RollbackScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RollbackScriptResponseBody) SetMessage(v string) *RollbackScriptResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackScriptResponseBody) SetRequestId(v string) *RollbackScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackScriptResponseBody) SetSuccess(v bool) *RollbackScriptResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
