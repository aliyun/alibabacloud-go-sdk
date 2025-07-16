// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageCount(v int32) *GetDocTranslateTaskResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *GetDocTranslateTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDocTranslateTaskResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetDocTranslateTaskResponseBody
	GetTaskId() *string
	SetTranslateErrorCode(v string) *GetDocTranslateTaskResponseBody
	GetTranslateErrorCode() *string
	SetTranslateErrorMessage(v string) *GetDocTranslateTaskResponseBody
	GetTranslateErrorMessage() *string
	SetTranslateFileUrl(v string) *GetDocTranslateTaskResponseBody
	GetTranslateFileUrl() *string
}

type GetDocTranslateTaskResponseBody struct {
	// example:
	//
	// 20
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 12952E92-FDF3-4D3C-88E3-242D72BA7150
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// translated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0586df512c8b4bb382d7d9a6a01b5854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Error
	TranslateErrorCode *string `json:"TranslateErrorCode,omitempty" xml:"TranslateErrorCode,omitempty"`
	// example:
	//
	// Fail to get the page number of document.
	TranslateErrorMessage *string `json:"TranslateErrorMessage,omitempty" xml:"TranslateErrorMessage,omitempty"`
	// example:
	//
	// http://translateFileUrl
	TranslateFileUrl *string `json:"TranslateFileUrl,omitempty" xml:"TranslateFileUrl,omitempty"`
}

func (s GetDocTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *GetDocTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocTranslateTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDocTranslateTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocTranslateTaskResponseBody) GetTranslateErrorCode() *string {
	return s.TranslateErrorCode
}

func (s *GetDocTranslateTaskResponseBody) GetTranslateErrorMessage() *string {
	return s.TranslateErrorMessage
}

func (s *GetDocTranslateTaskResponseBody) GetTranslateFileUrl() *string {
	return s.TranslateFileUrl
}

func (s *GetDocTranslateTaskResponseBody) SetPageCount(v int32) *GetDocTranslateTaskResponseBody {
	s.PageCount = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetRequestId(v string) *GetDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetStatus(v string) *GetDocTranslateTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTaskId(v string) *GetDocTranslateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateErrorCode(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateErrorCode = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateErrorMessage(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateErrorMessage = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateFileUrl(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateFileUrl = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
