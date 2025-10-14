// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHtmlTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHtmlTranslateTaskResponseBody
	GetCode() *string
	SetData(v *GetHtmlTranslateTaskResponseBodyData) *GetHtmlTranslateTaskResponseBody
	GetData() *GetHtmlTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *GetHtmlTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetHtmlTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHtmlTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHtmlTranslateTaskResponseBody
	GetSuccess() *bool
}

type GetHtmlTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHtmlTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 64191A87-C480-53AD-AEA2-2E847D4DFA66
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHtmlTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHtmlTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetHtmlTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHtmlTranslateTaskResponseBody) GetData() *GetHtmlTranslateTaskResponseBodyData {
	return s.Data
}

func (s *GetHtmlTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetHtmlTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHtmlTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHtmlTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHtmlTranslateTaskResponseBody) SetCode(v string) *GetHtmlTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBody) SetData(v *GetHtmlTranslateTaskResponseBodyData) *GetHtmlTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetHtmlTranslateTaskResponseBody) SetHttpStatusCode(v string) *GetHtmlTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBody) SetMessage(v string) *GetHtmlTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBody) SetRequestId(v string) *GetHtmlTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBody) SetSuccess(v bool) *GetHtmlTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHtmlTranslateTaskResponseBodyData struct {
	// example:
	//
	// <!DOCTYPE html>
	//
	// <html lang="zh-CN">
	//
	// <head>
	//
	// <meta charset="utf-8"/>
	//
	// <meta content="width=device-width, initial-scale=1.0" name="viewport"/>
	//
	// <title>My First Webpage</title>
	//
	// </head>
	//
	// <body>
	//
	// <h1>Welcome to my webpage!</h1>
	//
	// <p>This is a simple HTML page.</p>
	//
	// <p>Learning HTML is the first step to entering web development.</p>
	//
	// <a href="https://www.example.com">Click here to visit the sample website</a>
	//
	// </body>
	//
	// </html>
	Translation *string                                    `json:"translation,omitempty" xml:"translation,omitempty"`
	Usage       *GetHtmlTranslateTaskResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetHtmlTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHtmlTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHtmlTranslateTaskResponseBodyData) GetTranslation() *string {
	return s.Translation
}

func (s *GetHtmlTranslateTaskResponseBodyData) GetUsage() *GetHtmlTranslateTaskResponseBodyDataUsage {
	return s.Usage
}

func (s *GetHtmlTranslateTaskResponseBodyData) SetTranslation(v string) *GetHtmlTranslateTaskResponseBodyData {
	s.Translation = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBodyData) SetUsage(v *GetHtmlTranslateTaskResponseBodyDataUsage) *GetHtmlTranslateTaskResponseBodyData {
	s.Usage = v
	return s
}

func (s *GetHtmlTranslateTaskResponseBodyData) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHtmlTranslateTaskResponseBodyDataUsage struct {
	// example:
	//
	// 495
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 444
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 939
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetHtmlTranslateTaskResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GetHtmlTranslateTaskResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) SetInputTokens(v int64) *GetHtmlTranslateTaskResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) SetOutputTokens(v int64) *GetHtmlTranslateTaskResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) SetTotalTokens(v int64) *GetHtmlTranslateTaskResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetHtmlTranslateTaskResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
