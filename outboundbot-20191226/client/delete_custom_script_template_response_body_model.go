// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomScriptTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomScriptTemplateResponseBody
	GetCode() *string
	SetData(v string) *DeleteCustomScriptTemplateResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteCustomScriptTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomScriptTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomScriptTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomScriptTemplateResponseBody
	GetSuccess() *bool
}

type DeleteCustomScriptTemplateResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data. No business data is returned when the deletion is successful.
	//
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomScriptTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomScriptTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomScriptTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomScriptTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCustomScriptTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomScriptTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomScriptTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomScriptTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomScriptTemplateResponseBody) SetCode(v string) *DeleteCustomScriptTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponseBody) SetData(v string) *DeleteCustomScriptTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponseBody) SetHttpStatusCode(v int32) *DeleteCustomScriptTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponseBody) SetMessage(v string) *DeleteCustomScriptTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponseBody) SetRequestId(v string) *DeleteCustomScriptTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponseBody) SetSuccess(v bool) *DeleteCustomScriptTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
