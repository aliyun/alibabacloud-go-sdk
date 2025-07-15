// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateUploadTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ValidateUploadTemplateResponseBody
	GetCode() *string
	SetData(v *ValidateUploadTemplateResponseBodyData) *ValidateUploadTemplateResponseBody
	GetData() *ValidateUploadTemplateResponseBodyData
	SetHttpStatusCode(v int32) *ValidateUploadTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ValidateUploadTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValidateUploadTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateUploadTemplateResponseBody
	GetSuccess() *bool
}

type ValidateUploadTemplateResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ValidateUploadTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ValidateUploadTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateUploadTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateUploadTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValidateUploadTemplateResponseBody) GetData() *ValidateUploadTemplateResponseBodyData {
	return s.Data
}

func (s *ValidateUploadTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ValidateUploadTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValidateUploadTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateUploadTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateUploadTemplateResponseBody) SetCode(v string) *ValidateUploadTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateUploadTemplateResponseBody) SetData(v *ValidateUploadTemplateResponseBodyData) *ValidateUploadTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ValidateUploadTemplateResponseBody) SetHttpStatusCode(v int32) *ValidateUploadTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ValidateUploadTemplateResponseBody) SetMessage(v string) *ValidateUploadTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateUploadTemplateResponseBody) SetRequestId(v string) *ValidateUploadTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateUploadTemplateResponseBody) SetSuccess(v bool) *ValidateUploadTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateUploadTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ValidateUploadTemplateResponseBodyData struct {
	// example:
	//
	// 50
	CommentCount *int32 `json:"CommentCount,omitempty" xml:"CommentCount,omitempty"`
	// example:
	//
	// 50
	DialogueCount *int32 `json:"DialogueCount,omitempty" xml:"DialogueCount,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ValidateUploadTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ValidateUploadTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateUploadTemplateResponseBodyData) GetCommentCount() *int32 {
	return s.CommentCount
}

func (s *ValidateUploadTemplateResponseBodyData) GetDialogueCount() *int32 {
	return s.DialogueCount
}

func (s *ValidateUploadTemplateResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ValidateUploadTemplateResponseBodyData) SetCommentCount(v int32) *ValidateUploadTemplateResponseBodyData {
	s.CommentCount = &v
	return s
}

func (s *ValidateUploadTemplateResponseBodyData) SetDialogueCount(v int32) *ValidateUploadTemplateResponseBodyData {
	s.DialogueCount = &v
	return s
}

func (s *ValidateUploadTemplateResponseBodyData) SetTotalCount(v int32) *ValidateUploadTemplateResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ValidateUploadTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
