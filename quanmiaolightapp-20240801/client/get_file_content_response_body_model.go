// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFileContentResponseBody
	GetCode() *string
	SetData(v *GetFileContentResponseBodyData) *GetFileContentResponseBody
	GetData() *GetFileContentResponseBodyData
	SetHttpStatusCode(v int32) *GetFileContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetFileContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFileContentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetFileContentResponseBody
	GetSuccess() *string
}

type GetFileContentResponseBody struct {
	// example:
	//
	// successful
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetFileContentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BE603C85-90AA-50FC-A2D6-128AA9FA200D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFileContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFileContentResponseBody) GetData() *GetFileContentResponseBodyData {
	return s.Data
}

func (s *GetFileContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFileContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFileContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileContentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetFileContentResponseBody) SetCode(v string) *GetFileContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetFileContentResponseBody) SetData(v *GetFileContentResponseBodyData) *GetFileContentResponseBody {
	s.Data = v
	return s
}

func (s *GetFileContentResponseBody) SetHttpStatusCode(v int32) *GetFileContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFileContentResponseBody) SetMessage(v string) *GetFileContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetFileContentResponseBody) SetRequestId(v string) *GetFileContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileContentResponseBody) SetSuccess(v string) *GetFileContentResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileContentResponseBodyData struct {
	// example:
	//
	// xxxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetFileContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileContentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetFileContentResponseBodyData) SetContent(v string) *GetFileContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetFileContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
