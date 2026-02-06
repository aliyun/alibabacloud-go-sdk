// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PublishScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishScriptResponseBody
	GetSuccess() *bool
}

type PublishScriptResponseBody struct {
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

func (s PublishScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptResponseBody) GoString() string {
	return s.String()
}

func (s *PublishScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishScriptResponseBody) SetCode(v string) *PublishScriptResponseBody {
	s.Code = &v
	return s
}

func (s *PublishScriptResponseBody) SetHttpStatusCode(v int32) *PublishScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishScriptResponseBody) SetMessage(v string) *PublishScriptResponseBody {
	s.Message = &v
	return s
}

func (s *PublishScriptResponseBody) SetRequestId(v string) *PublishScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishScriptResponseBody) SetSuccess(v bool) *PublishScriptResponseBody {
	s.Success = &v
	return s
}

func (s *PublishScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
