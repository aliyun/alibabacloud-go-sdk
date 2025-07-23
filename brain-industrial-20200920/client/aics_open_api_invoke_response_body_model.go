// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAicsOpenApiInvokeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AicsOpenApiInvokeResponseBody
	GetCode() *string
	SetData(v interface{}) *AicsOpenApiInvokeResponseBody
	GetData() interface{}
	SetMessage(v string) *AicsOpenApiInvokeResponseBody
	GetMessage() *string
	SetRequestId(v string) *AicsOpenApiInvokeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AicsOpenApiInvokeResponseBody
	GetSuccess() *string
}

type AicsOpenApiInvokeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"c":2}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AicsOpenApiInvokeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AicsOpenApiInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *AicsOpenApiInvokeResponseBody) GetCode() *string {
	return s.Code
}

func (s *AicsOpenApiInvokeResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AicsOpenApiInvokeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AicsOpenApiInvokeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AicsOpenApiInvokeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AicsOpenApiInvokeResponseBody) SetCode(v string) *AicsOpenApiInvokeResponseBody {
	s.Code = &v
	return s
}

func (s *AicsOpenApiInvokeResponseBody) SetData(v interface{}) *AicsOpenApiInvokeResponseBody {
	s.Data = v
	return s
}

func (s *AicsOpenApiInvokeResponseBody) SetMessage(v string) *AicsOpenApiInvokeResponseBody {
	s.Message = &v
	return s
}

func (s *AicsOpenApiInvokeResponseBody) SetRequestId(v string) *AicsOpenApiInvokeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AicsOpenApiInvokeResponseBody) SetSuccess(v string) *AicsOpenApiInvokeResponseBody {
	s.Success = &v
	return s
}

func (s *AicsOpenApiInvokeResponseBody) Validate() error {
	return dara.Validate(s)
}
