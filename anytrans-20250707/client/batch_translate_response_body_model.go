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
	// The status code for the overall API call.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response payload that contains the translation results.
	Data *BatchTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier for the request. Use this ID for tracing and troubleshooting.
	//
	// example:
	//
	// 3BE338D3-16B1-513F-8DD2-57C8528DEAAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call was successful.
	//
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
	// An array of translation results, one for each text provided in the request.
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
	// The status code for the individual translation within the batch.
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The language code of the detected source language.
	DetectedLang *string `json:"detectedLang,omitempty" xml:"detectedLang,omitempty"`
	// The zero-based index of this result, which corresponds to the order of the source text in the original request.
	//
	// example:
	//
	// 0
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// The status message for the individual translation.
	//
	// example:
	//
	// OK
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The translated text.
	//
	// example:
	//
	// What will the weather be like tomorrow?
	Translation *string `json:"translation,omitempty" xml:"translation,omitempty"`
	// An object detailing the token usage for this translation.
	Usage *BatchTranslateResponseBodyDataTranslationListUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
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

func (s *BatchTranslateResponseBodyDataTranslationList) GetDetectedLang() *string {
	return s.DetectedLang
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

func (s *BatchTranslateResponseBodyDataTranslationList) SetDetectedLang(v string) *BatchTranslateResponseBodyDataTranslationList {
	s.DetectedLang = &v
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
	// The number of tokens in the source text.
	//
	// example:
	//
	// 53
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The number of tokens in the generated translation.
	//
	// example:
	//
	// 8
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The total number of tokens processed for the translation (the sum of `inputTokens` and `outputTokens`).
	//
	// example:
	//
	// 61
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
