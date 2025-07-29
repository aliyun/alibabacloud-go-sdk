// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOutputFormatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateOutputFormatResponseBody
	GetCode() *string
	SetData(v *GenerateOutputFormatResponseBodyData) *GenerateOutputFormatResponseBody
	GetData() *GenerateOutputFormatResponseBodyData
	SetHttpStatusCode(v int32) *GenerateOutputFormatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateOutputFormatResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateOutputFormatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateOutputFormatResponseBody
	GetSuccess() *bool
}

type GenerateOutputFormatResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *GenerateOutputFormatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateOutputFormatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateOutputFormatResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateOutputFormatResponseBody) GetData() *GenerateOutputFormatResponseBodyData {
	return s.Data
}

func (s *GenerateOutputFormatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateOutputFormatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateOutputFormatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateOutputFormatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateOutputFormatResponseBody) SetCode(v string) *GenerateOutputFormatResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetData(v *GenerateOutputFormatResponseBodyData) *GenerateOutputFormatResponseBody {
	s.Data = v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetHttpStatusCode(v int32) *GenerateOutputFormatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetMessage(v string) *GenerateOutputFormatResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetRequestId(v string) *GenerateOutputFormatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) SetSuccess(v bool) *GenerateOutputFormatResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateOutputFormatResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateOutputFormatResponseBodyData struct {
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
}

func (s GenerateOutputFormatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateOutputFormatResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatResponseBodyData) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *GenerateOutputFormatResponseBodyData) SetOutputFormat(v string) *GenerateOutputFormatResponseBodyData {
	s.OutputFormat = &v
	return s
}

func (s *GenerateOutputFormatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
