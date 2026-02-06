// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScriptResponseBody
	GetSuccess() *bool
}

type DeleteScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScriptResponseBody) SetCode(v string) *DeleteScriptResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScriptResponseBody) SetHttpStatusCode(v int32) *DeleteScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteScriptResponseBody) SetMessage(v string) *DeleteScriptResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScriptResponseBody) SetRequestId(v string) *DeleteScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScriptResponseBody) SetSuccess(v bool) *DeleteScriptResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
