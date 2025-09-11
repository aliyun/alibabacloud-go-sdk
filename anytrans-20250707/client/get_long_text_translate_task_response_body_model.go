// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLongTextTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLongTextTranslateTaskResponseBody
	GetCode() *string
	SetData(v *GetLongTextTranslateTaskResponseBodyData) *GetLongTextTranslateTaskResponseBody
	GetData() *GetLongTextTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *GetLongTextTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetLongTextTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLongTextTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLongTextTranslateTaskResponseBody
	GetSuccess() *bool
}

type GetLongTextTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetLongTextTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetLongTextTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLongTextTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetLongTextTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLongTextTranslateTaskResponseBody) GetData() *GetLongTextTranslateTaskResponseBodyData {
	return s.Data
}

func (s *GetLongTextTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetLongTextTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLongTextTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLongTextTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLongTextTranslateTaskResponseBody) SetCode(v string) *GetLongTextTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBody) SetData(v *GetLongTextTranslateTaskResponseBodyData) *GetLongTextTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetLongTextTranslateTaskResponseBody) SetHttpStatusCode(v string) *GetLongTextTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBody) SetMessage(v string) *GetLongTextTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBody) SetRequestId(v string) *GetLongTextTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBody) SetSuccess(v bool) *GetLongTextTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLongTextTranslateTaskResponseBodyData struct {
	// example:
	//
	// Featured Double Color Ball experts: Liu Ke and A Wang both hit the second prize, winning 1.43 million!
	Translation *string                                        `json:"translation,omitempty" xml:"translation,omitempty"`
	Usage       *GetLongTextTranslateTaskResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetLongTextTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLongTextTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLongTextTranslateTaskResponseBodyData) GetTranslation() *string {
	return s.Translation
}

func (s *GetLongTextTranslateTaskResponseBodyData) GetUsage() *GetLongTextTranslateTaskResponseBodyDataUsage {
	return s.Usage
}

func (s *GetLongTextTranslateTaskResponseBodyData) SetTranslation(v string) *GetLongTextTranslateTaskResponseBodyData {
	s.Translation = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBodyData) SetUsage(v *GetLongTextTranslateTaskResponseBodyDataUsage) *GetLongTextTranslateTaskResponseBodyData {
	s.Usage = v
	return s
}

func (s *GetLongTextTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetLongTextTranslateTaskResponseBodyDataUsage struct {
	// example:
	//
	// 495
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 515
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1010
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetLongTextTranslateTaskResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GetLongTextTranslateTaskResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) SetInputTokens(v int64) *GetLongTextTranslateTaskResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) SetOutputTokens(v int64) *GetLongTextTranslateTaskResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) SetTotalTokens(v int64) *GetLongTextTranslateTaskResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetLongTextTranslateTaskResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
