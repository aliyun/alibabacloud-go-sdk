// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAuditResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSmartAuditResultResponseBody
	GetCode() *string
	SetData(v *GetSmartAuditResultResponseBodyData) *GetSmartAuditResultResponseBody
	GetData() *GetSmartAuditResultResponseBodyData
	SetHttpStatusCode(v int32) *GetSmartAuditResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSmartAuditResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmartAuditResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSmartAuditResultResponseBody
	GetSuccess() *bool
}

type GetSmartAuditResultResponseBody struct {
	// The error code.
	//
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The review result.
	Data *GetSmartAuditResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSmartAuditResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmartAuditResultResponseBody) GetData() *GetSmartAuditResultResponseBodyData {
	return s.Data
}

func (s *GetSmartAuditResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSmartAuditResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmartAuditResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartAuditResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSmartAuditResultResponseBody) SetCode(v string) *GetSmartAuditResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetData(v *GetSmartAuditResultResponseBodyData) *GetSmartAuditResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetHttpStatusCode(v int32) *GetSmartAuditResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetMessage(v string) *GetSmartAuditResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetRequestId(v string) *GetSmartAuditResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetSuccess(v bool) *GetSmartAuditResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmartAuditResultResponseBodyData struct {
	// The list of review error details.
	ErrorItemDetails []*GetSmartAuditResultResponseBodyDataErrorItemDetails `json:"ErrorItemDetails,omitempty" xml:"ErrorItemDetails,omitempty" type:"Repeated"`
	// If the final status is not SUCCESSED, read this error message to identify the fault.
	//
	// example:
	//
	// 审核被取消
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task execution status. Valid values: PENDING, RUNNING, SUCCESSED, SUSPENDED, FAILED, and CANCELLED.
	//
	// example:
	//
	// SUCCESSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSmartAuditResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponseBodyData) GetErrorItemDetails() []*GetSmartAuditResultResponseBodyDataErrorItemDetails {
	return s.ErrorItemDetails
}

func (s *GetSmartAuditResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSmartAuditResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSmartAuditResultResponseBodyData) SetErrorItemDetails(v []*GetSmartAuditResultResponseBodyDataErrorItemDetails) *GetSmartAuditResultResponseBodyData {
	s.ErrorItemDetails = v
	return s
}

func (s *GetSmartAuditResultResponseBodyData) SetErrorMessage(v string) *GetSmartAuditResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyData) SetStatus(v string) *GetSmartAuditResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyData) Validate() error {
	if s.ErrorItemDetails != nil {
		for _, item := range s.ErrorItemDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSmartAuditResultResponseBodyDataErrorItemDetails struct {
	// The unique ID of the review item.
	//
	// example:
	//
	// 审核项唯一标识。
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The original text segment.
	//
	// example:
	//
	// 原文片段
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The offset index of the incorrect word within the context.
	//
	// example:
	//
	// 0
	ContextOffset *int32 `json:"ContextOffset,omitempty" xml:"ContextOffset,omitempty"`
	// The error level. 1: critical, 2: warning, 3: notice, 4: suggestion.
	//
	// example:
	//
	// 2
	ErrorLevel *int32 `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	// The incorrect word.
	//
	// example:
	//
	// ”xxx“
	ErrorWord *string `json:"ErrorWord,omitempty" xml:"ErrorWord,omitempty"`
	// The primary error code.
	//
	// example:
	//
	// ContentAccuracy
	MajorCode *string `json:"MajorCode,omitempty" xml:"MajorCode,omitempty"`
	// The description of the primary error.
	//
	// example:
	//
	// 内容准确性
	MajorCodeDesc *string `json:"MajorCodeDesc,omitempty" xml:"MajorCodeDesc,omitempty"`
	// The offset index of the incorrect word in the full text.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The error description.
	//
	// example:
	//
	// 中文双引号应成对正确使用，先左双引号，后右双引号
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The suggested correction.
	//
	// example:
	//
	// “xxx”
	RightWord *string `json:"RightWord,omitempty" xml:"RightWord,omitempty"`
	// The sub-error code.
	//
	// example:
	//
	// PunctuationError
	SubClassCode *string `json:"SubClassCode,omitempty" xml:"SubClassCode,omitempty"`
	// The description of the sub-error.
	//
	// example:
	//
	// 标点符号错误
	SubClassDesc *string `json:"SubClassDesc,omitempty" xml:"SubClassDesc,omitempty"`
	// In an image review scenario, this is the public URL of the image that triggered the review.
	//
	// example:
	//
	// http://www.example.com/xxxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSmartAuditResultResponseBodyDataErrorItemDetails) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponseBodyDataErrorItemDetails) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetCheckId() *string {
	return s.CheckId
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetContext() *string {
	return s.Context
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetContextOffset() *int32 {
	return s.ContextOffset
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetErrorLevel() *int32 {
	return s.ErrorLevel
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetErrorWord() *string {
	return s.ErrorWord
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetMajorCode() *string {
	return s.MajorCode
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetMajorCodeDesc() *string {
	return s.MajorCodeDesc
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetOffset() *int32 {
	return s.Offset
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetReason() *string {
	return s.Reason
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetRightWord() *string {
	return s.RightWord
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetSubClassCode() *string {
	return s.SubClassCode
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetSubClassDesc() *string {
	return s.SubClassDesc
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetUrl() *string {
	return s.Url
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetCheckId(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.CheckId = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetContext(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Context = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetContextOffset(v int32) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.ContextOffset = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetErrorLevel(v int32) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.ErrorLevel = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetErrorWord(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.ErrorWord = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetMajorCode(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.MajorCode = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetMajorCodeDesc(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.MajorCodeDesc = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetOffset(v int32) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Offset = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetReason(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Reason = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetRightWord(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.RightWord = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetSubClassCode(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.SubClassCode = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetSubClassDesc(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.SubClassDesc = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetUrl(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Url = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) Validate() error {
	return dara.Validate(s)
}
