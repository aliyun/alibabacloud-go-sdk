// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchTranslateResponseBody
	GetCode() *string
	SetData(v *BatchTranslateResponseBodyData) *BatchTranslateResponseBody
	GetData() *BatchTranslateResponseBodyData
	SetHttpStatusCode(v string) *BatchTranslateResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *BatchTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchTranslateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchTranslateResponseBody
	GetSuccess() *bool
}

type BatchTranslateResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s BatchTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchTranslateResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchTranslateResponseBody) GetData() *BatchTranslateResponseBodyData {
	return s.Data
}

func (s *BatchTranslateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *BatchTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchTranslateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchTranslateResponseBody) SetCode(v string) *BatchTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *BatchTranslateResponseBody) SetData(v *BatchTranslateResponseBodyData) *BatchTranslateResponseBody {
	s.Data = v
	return s
}

func (s *BatchTranslateResponseBody) SetHttpStatusCode(v string) *BatchTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchTranslateResponseBody) SetMessage(v string) *BatchTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *BatchTranslateResponseBody) SetRequestId(v string) *BatchTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchTranslateResponseBody) SetSuccess(v bool) *BatchTranslateResponseBody {
	s.Success = &v
	return s
}

func (s *BatchTranslateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchTranslateResponseBodyData struct {
	TranslationList []*BatchTranslateResponseBodyDataTranslationList `json:"translationList,omitempty" xml:"translationList,omitempty" type:"Repeated"`
}

func (s BatchTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchTranslateResponseBodyData) GetTranslationList() []*BatchTranslateResponseBodyDataTranslationList {
	return s.TranslationList
}

func (s *BatchTranslateResponseBodyData) SetTranslationList(v []*BatchTranslateResponseBodyDataTranslationList) *BatchTranslateResponseBodyData {
	s.TranslationList = v
	return s
}

func (s *BatchTranslateResponseBodyData) Validate() error {
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

type BatchTranslateResponseBodyDataTranslationList struct {
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
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// Featured Double Color Ball experts: Liu Ke and A Wang both hit the second prize, winning 1.43 million!
	Translation *string                                             `json:"translation,omitempty" xml:"translation,omitempty"`
	Usage       *BatchTranslateResponseBodyDataTranslationListUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s BatchTranslateResponseBodyDataTranslationList) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateResponseBodyDataTranslationList) GoString() string {
	return s.String()
}

func (s *BatchTranslateResponseBodyDataTranslationList) GetCode() *int64 {
	return s.Code
}

func (s *BatchTranslateResponseBodyDataTranslationList) GetIndex() *string {
	return s.Index
}

func (s *BatchTranslateResponseBodyDataTranslationList) GetMessage() *string {
	return s.Message
}

func (s *BatchTranslateResponseBodyDataTranslationList) GetTranslation() *string {
	return s.Translation
}

func (s *BatchTranslateResponseBodyDataTranslationList) GetUsage() *BatchTranslateResponseBodyDataTranslationListUsage {
	return s.Usage
}

func (s *BatchTranslateResponseBodyDataTranslationList) SetCode(v int64) *BatchTranslateResponseBodyDataTranslationList {
	s.Code = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationList) SetIndex(v string) *BatchTranslateResponseBodyDataTranslationList {
	s.Index = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationList) SetMessage(v string) *BatchTranslateResponseBodyDataTranslationList {
	s.Message = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationList) SetTranslation(v string) *BatchTranslateResponseBodyDataTranslationList {
	s.Translation = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationList) SetUsage(v *BatchTranslateResponseBodyDataTranslationListUsage) *BatchTranslateResponseBodyDataTranslationList {
	s.Usage = v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationList) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchTranslateResponseBodyDataTranslationListUsage struct {
	// example:
	//
	// 480
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 520
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1000
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s BatchTranslateResponseBodyDataTranslationListUsage) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateResponseBodyDataTranslationListUsage) GoString() string {
	return s.String()
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) SetInputTokens(v int64) *BatchTranslateResponseBodyDataTranslationListUsage {
	s.InputTokens = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) SetOutputTokens(v int64) *BatchTranslateResponseBodyDataTranslationListUsage {
	s.OutputTokens = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) SetTotalTokens(v int64) *BatchTranslateResponseBodyDataTranslationListUsage {
	s.TotalTokens = &v
	return s
}

func (s *BatchTranslateResponseBodyDataTranslationListUsage) Validate() error {
	return dara.Validate(s)
}
