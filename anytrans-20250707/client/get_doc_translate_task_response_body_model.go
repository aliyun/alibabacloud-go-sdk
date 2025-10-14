// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocTranslateTaskResponseBody
	GetCode() *string
	SetData(v *GetDocTranslateTaskResponseBodyData) *GetDocTranslateTaskResponseBody
	GetData() *GetDocTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *GetDocTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetDocTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocTranslateTaskResponseBody
	GetSuccess() *bool
}

type GetDocTranslateTaskResponseBody struct {
	// example:
	//
	// success
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetDocTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// AC642EEB-C29D-54DF-8F52-622565BBB78A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDocTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocTranslateTaskResponseBody) GetData() *GetDocTranslateTaskResponseBodyData {
	return s.Data
}

func (s *GetDocTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetDocTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocTranslateTaskResponseBody) SetCode(v string) *GetDocTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetData(v *GetDocTranslateTaskResponseBodyData) *GetDocTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetHttpStatusCode(v string) *GetDocTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetMessage(v string) *GetDocTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetRequestId(v string) *GetDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetSuccess(v bool) *GetDocTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocTranslateTaskResponseBodyData struct {
	// example:
	//
	// 4
	CharactersCount *int32 `json:"charactersCount,omitempty" xml:"charactersCount,omitempty"`
	// example:
	//
	// 2
	PageCount *int32 `json:"pageCount,omitempty" xml:"pageCount,omitempty"`
	// example:
	//
	// translated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// d3a2397bc2c14ab4a2e40a4f5b46241b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// http://translate-ai-data-dev.oss-cn-hangzhou.aliyuncs.com/tongyiTranslate/123456789/a7630164ce894c799cca0f0822c36f84_merge.md?Expires=1756700753&OSSAccessKeyId=LTAI5tRmTwxU5YoHUyrF93Jv&Signature=qB03ldVmRa%2FRwWgJ2wSi7ylayMA%3D
	TranslateFileUrl *string `json:"translateFileUrl,omitempty" xml:"translateFileUrl,omitempty"`
}

func (s GetDocTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponseBodyData) GetCharactersCount() *int32 {
	return s.CharactersCount
}

func (s *GetDocTranslateTaskResponseBodyData) GetPageCount() *int32 {
	return s.PageCount
}

func (s *GetDocTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDocTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocTranslateTaskResponseBodyData) GetTranslateFileUrl() *string {
	return s.TranslateFileUrl
}

func (s *GetDocTranslateTaskResponseBodyData) SetCharactersCount(v int32) *GetDocTranslateTaskResponseBodyData {
	s.CharactersCount = &v
	return s
}

func (s *GetDocTranslateTaskResponseBodyData) SetPageCount(v int32) *GetDocTranslateTaskResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *GetDocTranslateTaskResponseBodyData) SetStatus(v string) *GetDocTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDocTranslateTaskResponseBodyData) SetTaskId(v string) *GetDocTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBodyData) SetTranslateFileUrl(v string) *GetDocTranslateTaskResponseBodyData {
	s.TranslateFileUrl = &v
	return s
}

func (s *GetDocTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
