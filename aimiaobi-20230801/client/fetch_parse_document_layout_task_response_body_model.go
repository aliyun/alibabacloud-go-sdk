// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchParseDocumentLayoutTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FetchParseDocumentLayoutTaskResponseBody
	GetCode() *string
	SetData(v *FetchParseDocumentLayoutTaskResponseBodyData) *FetchParseDocumentLayoutTaskResponseBody
	GetData() *FetchParseDocumentLayoutTaskResponseBodyData
	SetHttpStatusCode(v int32) *FetchParseDocumentLayoutTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FetchParseDocumentLayoutTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchParseDocumentLayoutTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchParseDocumentLayoutTaskResponseBody
	GetSuccess() *bool
}

type FetchParseDocumentLayoutTaskResponseBody struct {
	// Status code
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data
	Data *FetchParseDocumentLayoutTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchParseDocumentLayoutTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchParseDocumentLayoutTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchParseDocumentLayoutTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FetchParseDocumentLayoutTaskResponseBody) GetData() *FetchParseDocumentLayoutTaskResponseBodyData {
	return s.Data
}

func (s *FetchParseDocumentLayoutTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FetchParseDocumentLayoutTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchParseDocumentLayoutTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchParseDocumentLayoutTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchParseDocumentLayoutTaskResponseBody) SetCode(v string) *FetchParseDocumentLayoutTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBody) SetData(v *FetchParseDocumentLayoutTaskResponseBodyData) *FetchParseDocumentLayoutTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBody) SetHttpStatusCode(v int32) *FetchParseDocumentLayoutTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBody) SetMessage(v string) *FetchParseDocumentLayoutTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBody) SetRequestId(v string) *FetchParseDocumentLayoutTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBody) SetSuccess(v bool) *FetchParseDocumentLayoutTaskResponseBody {
	s.Success = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchParseDocumentLayoutTaskResponseBodyData struct {
	// Structured content after formatting
	LayoutResult *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult `json:"LayoutResult,omitempty" xml:"LayoutResult,omitempty" type:"Struct"`
	// Task status
	//
	// example:
	//
	// PENDING-待执行、RUNNING-执行中、SUCCESSED-成功、SUSPENDED-暂停、FAILED-失败、CANCELLED-取消
	TaskStats *string `json:"TaskStats,omitempty" xml:"TaskStats,omitempty"`
}

func (s FetchParseDocumentLayoutTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchParseDocumentLayoutTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchParseDocumentLayoutTaskResponseBodyData) GetLayoutResult() *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult {
	return s.LayoutResult
}

func (s *FetchParseDocumentLayoutTaskResponseBodyData) GetTaskStats() *string {
	return s.TaskStats
}

func (s *FetchParseDocumentLayoutTaskResponseBodyData) SetLayoutResult(v *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult) *FetchParseDocumentLayoutTaskResponseBodyData {
	s.LayoutResult = v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyData) SetTaskStats(v string) *FetchParseDocumentLayoutTaskResponseBodyData {
	s.TaskStats = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyData) Validate() error {
	if s.LayoutResult != nil {
		if err := s.LayoutResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult struct {
	// Returned element data
	Elements []*FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult) String() string {
	return dara.Prettify(s)
}

func (s FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult) GoString() string {
	return s.String()
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult) GetElements() []*FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements {
	return s.Elements
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult) SetElements(v []*FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult {
	s.Elements = v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResult) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements struct {
	// Content
	//
	// example:
	//
	// 一、本月主要工作进展\\n
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// For headings, content with heading numbers removed
	//
	// example:
	//
	// 本月主要工作进展
	FormatContent *string `json:"FormatContent,omitempty" xml:"FormatContent,omitempty"`
	// Index order of each element
	//
	// example:
	//
	// 1
	Index *float32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Type
	//
	// example:
	//
	// 支持的类型如下
	//
	//     HEADING("标题"),
	//
	//     H1("一级标题"),
	//
	//     H2("二级标题"),
	//
	//     H3("三级标题"),
	//
	//     H4("四级标题"),
	//
	//     H5("五级标题"),
	//
	//     H6("六级标题"),
	//
	//     PARAGRAPH("段落"),
	//
	//     SIGNATURE("落款"),
	//
	//     FOOTNOTE("脚注"),
	//
	//     TABLE("表格"),
	//
	//     CODE_BLOCK("代码块"),
	//
	//     ATTACHMENT("附件"),
	//
	//     BLOCKQUOTE("引用");
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) String() string {
	return dara.Prettify(s)
}

func (s FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) GoString() string {
	return s.String()
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) GetContent() *string {
	return s.Content
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) GetFormatContent() *string {
	return s.FormatContent
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) GetIndex() *float32 {
	return s.Index
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) GetType() *string {
	return s.Type
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) SetContent(v string) *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements {
	s.Content = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) SetFormatContent(v string) *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements {
	s.FormatContent = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) SetIndex(v float32) *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements {
	s.Index = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) SetType(v string) *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements {
	s.Type = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponseBodyDataLayoutResultElements) Validate() error {
	return dara.Validate(s)
}
