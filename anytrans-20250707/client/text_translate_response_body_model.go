// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TextTranslateResponseBody
	GetCode() *string
	SetData(v *TextTranslateResponseBodyData) *TextTranslateResponseBody
	GetData() *TextTranslateResponseBodyData
	SetHttpStatusCode(v string) *TextTranslateResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *TextTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextTranslateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TextTranslateResponseBody
	GetSuccess() *bool
}

type TextTranslateResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *TextTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 299C57B2-EBB4-57E2-A6FE-723B874ACB74
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TextTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *TextTranslateResponseBody) GetCode() *string {
	return s.Code
}

func (s *TextTranslateResponseBody) GetData() *TextTranslateResponseBodyData {
	return s.Data
}

func (s *TextTranslateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *TextTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextTranslateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TextTranslateResponseBody) SetCode(v string) *TextTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *TextTranslateResponseBody) SetData(v *TextTranslateResponseBodyData) *TextTranslateResponseBody {
	s.Data = v
	return s
}

func (s *TextTranslateResponseBody) SetHttpStatusCode(v string) *TextTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TextTranslateResponseBody) SetMessage(v string) *TextTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *TextTranslateResponseBody) SetRequestId(v string) *TextTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextTranslateResponseBody) SetSuccess(v bool) *TextTranslateResponseBody {
	s.Success = &v
	return s
}

func (s *TextTranslateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextTranslateResponseBodyData struct {
	// example:
	//
	// How does Mogujie solve the data annotation challenge by building a platform?
	Translation *string                             `json:"translation,omitempty" xml:"translation,omitempty"`
	Usage       *TextTranslateResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s TextTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextTranslateResponseBodyData) GetTranslation() *string {
	return s.Translation
}

func (s *TextTranslateResponseBodyData) GetUsage() *TextTranslateResponseBodyDataUsage {
	return s.Usage
}

func (s *TextTranslateResponseBodyData) SetTranslation(v string) *TextTranslateResponseBodyData {
	s.Translation = &v
	return s
}

func (s *TextTranslateResponseBodyData) SetUsage(v *TextTranslateResponseBodyDataUsage) *TextTranslateResponseBodyData {
	s.Usage = v
	return s
}

func (s *TextTranslateResponseBodyData) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextTranslateResponseBodyDataUsage struct {
	// example:
	//
	// 491
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 400
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 891
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s TextTranslateResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *TextTranslateResponseBodyDataUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *TextTranslateResponseBodyDataUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *TextTranslateResponseBodyDataUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *TextTranslateResponseBodyDataUsage) SetInputTokens(v int64) *TextTranslateResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *TextTranslateResponseBodyDataUsage) SetOutputTokens(v int64) *TextTranslateResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *TextTranslateResponseBodyDataUsage) SetTotalTokens(v int64) *TextTranslateResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *TextTranslateResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
