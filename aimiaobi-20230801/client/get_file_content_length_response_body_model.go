// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileContentLengthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFileContentLengthResponseBody
	GetCode() *string
	SetData(v *GetFileContentLengthResponseBodyData) *GetFileContentLengthResponseBody
	GetData() *GetFileContentLengthResponseBodyData
	SetHttpStatusCode(v int32) *GetFileContentLengthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetFileContentLengthResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFileContentLengthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFileContentLengthResponseBody
	GetSuccess() *bool
}

type GetFileContentLengthResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetFileContentLengthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// DD656AF9-0839-521A-A3D2-F320009F9C87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileContentLengthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentLengthResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileContentLengthResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFileContentLengthResponseBody) GetData() *GetFileContentLengthResponseBodyData {
	return s.Data
}

func (s *GetFileContentLengthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFileContentLengthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFileContentLengthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileContentLengthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileContentLengthResponseBody) SetCode(v string) *GetFileContentLengthResponseBody {
	s.Code = &v
	return s
}

func (s *GetFileContentLengthResponseBody) SetData(v *GetFileContentLengthResponseBodyData) *GetFileContentLengthResponseBody {
	s.Data = v
	return s
}

func (s *GetFileContentLengthResponseBody) SetHttpStatusCode(v int32) *GetFileContentLengthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFileContentLengthResponseBody) SetMessage(v string) *GetFileContentLengthResponseBody {
	s.Message = &v
	return s
}

func (s *GetFileContentLengthResponseBody) SetRequestId(v string) *GetFileContentLengthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileContentLengthResponseBody) SetSuccess(v bool) *GetFileContentLengthResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileContentLengthResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileContentLengthResponseBodyData struct {
	// example:
	//
	// 1024
	WordNum *int64 `json:"WordNum,omitempty" xml:"WordNum,omitempty"`
}

func (s GetFileContentLengthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentLengthResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileContentLengthResponseBodyData) GetWordNum() *int64 {
	return s.WordNum
}

func (s *GetFileContentLengthResponseBodyData) SetWordNum(v int64) *GetFileContentLengthResponseBodyData {
	s.WordNum = &v
	return s
}

func (s *GetFileContentLengthResponseBodyData) Validate() error {
	return dara.Validate(s)
}
