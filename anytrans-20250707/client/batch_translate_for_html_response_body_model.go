// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateForHtmlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchTranslateForHtmlResponseBody
	GetCode() *string
	SetData(v *BatchTranslateForHtmlResponseBodyData) *BatchTranslateForHtmlResponseBody
	GetData() *BatchTranslateForHtmlResponseBodyData
	SetHttpStatusCode(v string) *BatchTranslateForHtmlResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *BatchTranslateForHtmlResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchTranslateForHtmlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchTranslateForHtmlResponseBody
	GetSuccess() *bool
}

type BatchTranslateForHtmlResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchTranslateForHtmlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 3BE338D3-16B1-513F-8DD2-57C8528DEAAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchTranslateForHtmlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlResponseBody) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchTranslateForHtmlResponseBody) GetData() *BatchTranslateForHtmlResponseBodyData {
	return s.Data
}

func (s *BatchTranslateForHtmlResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *BatchTranslateForHtmlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchTranslateForHtmlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchTranslateForHtmlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchTranslateForHtmlResponseBody) SetCode(v string) *BatchTranslateForHtmlResponseBody {
	s.Code = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBody) SetData(v *BatchTranslateForHtmlResponseBodyData) *BatchTranslateForHtmlResponseBody {
	s.Data = v
	return s
}

func (s *BatchTranslateForHtmlResponseBody) SetHttpStatusCode(v string) *BatchTranslateForHtmlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBody) SetMessage(v string) *BatchTranslateForHtmlResponseBody {
	s.Message = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBody) SetRequestId(v string) *BatchTranslateForHtmlResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBody) SetSuccess(v bool) *BatchTranslateForHtmlResponseBody {
	s.Success = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchTranslateForHtmlResponseBodyData struct {
	TranslationList []*BatchTranslateForHtmlResponseBodyDataTranslationList `json:"translationList,omitempty" xml:"translationList,omitempty" type:"Repeated"`
}

func (s BatchTranslateForHtmlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlResponseBodyData) GetTranslationList() []*BatchTranslateForHtmlResponseBodyDataTranslationList {
	return s.TranslationList
}

func (s *BatchTranslateForHtmlResponseBodyData) SetTranslationList(v []*BatchTranslateForHtmlResponseBodyDataTranslationList) *BatchTranslateForHtmlResponseBodyData {
	s.TranslationList = v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyData) Validate() error {
	if s.TranslationList != nil {
		for _, item := range s.TranslationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchTranslateForHtmlResponseBodyDataTranslationList struct {
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 0
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// What will the weather be like tomorrow?
	Translation *string                                                    `json:"translation,omitempty" xml:"translation,omitempty"`
	Usage       *BatchTranslateForHtmlResponseBodyDataTranslationListUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s BatchTranslateForHtmlResponseBodyDataTranslationList) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlResponseBodyDataTranslationList) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) GetCode() *int64 {
	return s.Code
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) GetIndex() *string {
	return s.Index
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) GetMessage() *string {
	return s.Message
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) GetTranslation() *string {
	return s.Translation
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) GetUsage() *BatchTranslateForHtmlResponseBodyDataTranslationListUsage {
	return s.Usage
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) SetCode(v int64) *BatchTranslateForHtmlResponseBodyDataTranslationList {
	s.Code = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) SetIndex(v string) *BatchTranslateForHtmlResponseBodyDataTranslationList {
	s.Index = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) SetMessage(v string) *BatchTranslateForHtmlResponseBodyDataTranslationList {
	s.Message = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) SetTranslation(v string) *BatchTranslateForHtmlResponseBodyDataTranslationList {
	s.Translation = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) SetUsage(v *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) *BatchTranslateForHtmlResponseBodyDataTranslationList {
	s.Usage = v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationList) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchTranslateForHtmlResponseBodyDataTranslationListUsage struct {
	// example:
	//
	// 53
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 8
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 61
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s BatchTranslateForHtmlResponseBodyDataTranslationListUsage) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlResponseBodyDataTranslationListUsage) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) SetInputTokens(v int64) *BatchTranslateForHtmlResponseBodyDataTranslationListUsage {
	s.InputTokens = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) SetOutputTokens(v int64) *BatchTranslateForHtmlResponseBodyDataTranslationListUsage {
	s.OutputTokens = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) SetTotalTokens(v int64) *BatchTranslateForHtmlResponseBodyDataTranslationListUsage {
	s.TotalTokens = &v
	return s
}

func (s *BatchTranslateForHtmlResponseBodyDataTranslationListUsage) Validate() error {
	return dara.Validate(s)
}
