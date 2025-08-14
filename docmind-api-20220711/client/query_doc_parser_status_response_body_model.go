// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDocParserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDocParserStatusResponseBody
	GetCode() *string
	SetData(v *QueryDocParserStatusResponseBodyData) *QueryDocParserStatusResponseBody
	GetData() *QueryDocParserStatusResponseBodyData
	SetMessage(v string) *QueryDocParserStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDocParserStatusResponseBody
	GetRequestId() *string
}

type QueryDocParserStatusResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryDocParserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDocParserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDocParserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDocParserStatusResponseBody) GetData() *QueryDocParserStatusResponseBodyData {
	return s.Data
}

func (s *QueryDocParserStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDocParserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDocParserStatusResponseBody) SetCode(v string) *QueryDocParserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDocParserStatusResponseBody) SetData(v *QueryDocParserStatusResponseBodyData) *QueryDocParserStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryDocParserStatusResponseBody) SetMessage(v string) *QueryDocParserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDocParserStatusResponseBody) SetRequestId(v string) *QueryDocParserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDocParserStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDocParserStatusResponseBodyData struct {
	ImageCount                *int32   `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	NumberOfSuccessfulParsing *int32   `json:"NumberOfSuccessfulParsing,omitempty" xml:"NumberOfSuccessfulParsing,omitempty"`
	PageCountEstimate         *int32   `json:"PageCountEstimate,omitempty" xml:"PageCountEstimate,omitempty"`
	ParagraphCount            *int32   `json:"ParagraphCount,omitempty" xml:"ParagraphCount,omitempty"`
	Processing                *float32 `json:"Processing,omitempty" xml:"Processing,omitempty"`
	Status                    *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TableCount                *int32   `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	Tokens                    *int64   `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
}

func (s QueryDocParserStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDocParserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponseBodyData) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *QueryDocParserStatusResponseBodyData) GetNumberOfSuccessfulParsing() *int32 {
	return s.NumberOfSuccessfulParsing
}

func (s *QueryDocParserStatusResponseBodyData) GetPageCountEstimate() *int32 {
	return s.PageCountEstimate
}

func (s *QueryDocParserStatusResponseBodyData) GetParagraphCount() *int32 {
	return s.ParagraphCount
}

func (s *QueryDocParserStatusResponseBodyData) GetProcessing() *float32 {
	return s.Processing
}

func (s *QueryDocParserStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryDocParserStatusResponseBodyData) GetTableCount() *int32 {
	return s.TableCount
}

func (s *QueryDocParserStatusResponseBodyData) GetTokens() *int64 {
	return s.Tokens
}

func (s *QueryDocParserStatusResponseBodyData) SetImageCount(v int32) *QueryDocParserStatusResponseBodyData {
	s.ImageCount = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetNumberOfSuccessfulParsing(v int32) *QueryDocParserStatusResponseBodyData {
	s.NumberOfSuccessfulParsing = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetPageCountEstimate(v int32) *QueryDocParserStatusResponseBodyData {
	s.PageCountEstimate = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetParagraphCount(v int32) *QueryDocParserStatusResponseBodyData {
	s.ParagraphCount = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetProcessing(v float32) *QueryDocParserStatusResponseBodyData {
	s.Processing = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetStatus(v string) *QueryDocParserStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetTableCount(v int32) *QueryDocParserStatusResponseBodyData {
	s.TableCount = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetTokens(v int64) *QueryDocParserStatusResponseBodyData {
	s.Tokens = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
