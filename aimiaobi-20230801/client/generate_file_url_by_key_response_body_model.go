// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileUrlByKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateFileUrlByKeyResponseBody
	GetCode() *string
	SetData(v string) *GenerateFileUrlByKeyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GenerateFileUrlByKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateFileUrlByKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateFileUrlByKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateFileUrlByKeyResponseBody
	GetSuccess() *bool
}

type GenerateFileUrlByKeyResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://www.example.com/a.txt
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateFileUrlByKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUrlByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateFileUrlByKeyResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateFileUrlByKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateFileUrlByKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateFileUrlByKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateFileUrlByKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateFileUrlByKeyResponseBody) SetCode(v string) *GenerateFileUrlByKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetData(v string) *GenerateFileUrlByKeyResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetHttpStatusCode(v int32) *GenerateFileUrlByKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetMessage(v string) *GenerateFileUrlByKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetRequestId(v string) *GenerateFileUrlByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetSuccess(v bool) *GenerateFileUrlByKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
