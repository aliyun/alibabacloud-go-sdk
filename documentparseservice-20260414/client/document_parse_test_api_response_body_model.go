// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseTestApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocumentParseTestApiResponseBody
	GetCode() *string
	SetData(v *DocumentParseTestApiResponseBodyData) *DocumentParseTestApiResponseBody
	GetData() *DocumentParseTestApiResponseBodyData
	SetMessage(v string) *DocumentParseTestApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocumentParseTestApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DocumentParseTestApiResponseBody
	GetSuccess() *bool
}

type DocumentParseTestApiResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DocumentParseTestApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DocumentParseTestApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseTestApiResponseBody) GoString() string {
	return s.String()
}

func (s *DocumentParseTestApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocumentParseTestApiResponseBody) GetData() *DocumentParseTestApiResponseBodyData {
	return s.Data
}

func (s *DocumentParseTestApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocumentParseTestApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocumentParseTestApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DocumentParseTestApiResponseBody) SetCode(v string) *DocumentParseTestApiResponseBody {
	s.Code = &v
	return s
}

func (s *DocumentParseTestApiResponseBody) SetData(v *DocumentParseTestApiResponseBodyData) *DocumentParseTestApiResponseBody {
	s.Data = v
	return s
}

func (s *DocumentParseTestApiResponseBody) SetMessage(v string) *DocumentParseTestApiResponseBody {
	s.Message = &v
	return s
}

func (s *DocumentParseTestApiResponseBody) SetRequestId(v string) *DocumentParseTestApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocumentParseTestApiResponseBody) SetSuccess(v bool) *DocumentParseTestApiResponseBody {
	s.Success = &v
	return s
}

func (s *DocumentParseTestApiResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocumentParseTestApiResponseBodyData struct {
	// example:
	//
	// [\\"LTAI5tQcqrTt1jrRbZf7mHFs\\"]
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// BAA69D16-9453-5A8F-AE3C-498D956A0A73
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Usage     *DocumentParseTestApiResponseBodyDataUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s DocumentParseTestApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseTestApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *DocumentParseTestApiResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *DocumentParseTestApiResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *DocumentParseTestApiResponseBodyData) GetUsage() *DocumentParseTestApiResponseBodyDataUsage {
	return s.Usage
}

func (s *DocumentParseTestApiResponseBodyData) SetContent(v string) *DocumentParseTestApiResponseBodyData {
	s.Content = &v
	return s
}

func (s *DocumentParseTestApiResponseBodyData) SetRequestId(v string) *DocumentParseTestApiResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DocumentParseTestApiResponseBodyData) SetUsage(v *DocumentParseTestApiResponseBodyDataUsage) *DocumentParseTestApiResponseBodyData {
	s.Usage = v
	return s
}

func (s *DocumentParseTestApiResponseBodyData) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocumentParseTestApiResponseBodyDataUsage struct {
	// example:
	//
	// 1000
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 2000
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 3000
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s DocumentParseTestApiResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseTestApiResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *DocumentParseTestApiResponseBodyDataUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *DocumentParseTestApiResponseBodyDataUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *DocumentParseTestApiResponseBodyDataUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *DocumentParseTestApiResponseBodyDataUsage) SetInputTokens(v int32) *DocumentParseTestApiResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *DocumentParseTestApiResponseBodyDataUsage) SetOutputTokens(v int32) *DocumentParseTestApiResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *DocumentParseTestApiResponseBodyDataUsage) SetTotalTokens(v int32) *DocumentParseTestApiResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *DocumentParseTestApiResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
