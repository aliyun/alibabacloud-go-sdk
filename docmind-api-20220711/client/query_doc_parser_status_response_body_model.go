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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDocParserStatusResponseBodyData struct {
	ImageCount                *int32                                                    `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	NumberOfSuccessfulParsing *int32                                                    `json:"NumberOfSuccessfulParsing,omitempty" xml:"NumberOfSuccessfulParsing,omitempty"`
	OutputFormatResult        []*QueryDocParserStatusResponseBodyDataOutputFormatResult `json:"OutputFormatResult,omitempty" xml:"OutputFormatResult,omitempty" type:"Repeated"`
	PageCountEstimate         *int32                                                    `json:"PageCountEstimate,omitempty" xml:"PageCountEstimate,omitempty"`
	ParagraphCount            *int32                                                    `json:"ParagraphCount,omitempty" xml:"ParagraphCount,omitempty"`
	Processing                *float32                                                  `json:"Processing,omitempty" xml:"Processing,omitempty"`
	Status                    *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TableCount                *int32                                                    `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	Tokens                    *int64                                                    `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
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

func (s *QueryDocParserStatusResponseBodyData) GetOutputFormatResult() []*QueryDocParserStatusResponseBodyDataOutputFormatResult {
	return s.OutputFormatResult
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

func (s *QueryDocParserStatusResponseBodyData) SetOutputFormatResult(v []*QueryDocParserStatusResponseBodyDataOutputFormatResult) *QueryDocParserStatusResponseBodyData {
	s.OutputFormatResult = v
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
	if s.OutputFormatResult != nil {
		for _, item := range s.OutputFormatResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDocParserStatusResponseBodyDataOutputFormatResult struct {
	OutputFileUrl *string                                                        `json:"OutputFileUrl,omitempty" xml:"OutputFileUrl,omitempty"`
	OutputType    *string                                                        `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	Pages         []*QueryDocParserStatusResponseBodyDataOutputFormatResultPages `json:"Pages,omitempty" xml:"Pages,omitempty" type:"Repeated"`
}

func (s QueryDocParserStatusResponseBodyDataOutputFormatResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDocParserStatusResponseBodyDataOutputFormatResult) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) GetOutputFileUrl() *string {
	return s.OutputFileUrl
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) GetOutputType() *string {
	return s.OutputType
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) GetPages() []*QueryDocParserStatusResponseBodyDataOutputFormatResultPages {
	return s.Pages
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) SetOutputFileUrl(v string) *QueryDocParserStatusResponseBodyDataOutputFormatResult {
	s.OutputFileUrl = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) SetOutputType(v string) *QueryDocParserStatusResponseBodyDataOutputFormatResult {
	s.OutputType = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) SetPages(v []*QueryDocParserStatusResponseBodyDataOutputFormatResultPages) *QueryDocParserStatusResponseBodyDataOutputFormatResult {
	s.Pages = v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResult) Validate() error {
	if s.Pages != nil {
		for _, item := range s.Pages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDocParserStatusResponseBodyDataOutputFormatResultPages struct {
	ImageHeight   *int32  `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageUrl      *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ImageWidth    *int32  `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	PageIdAllDocs *int32  `json:"PageIdAllDocs,omitempty" xml:"PageIdAllDocs,omitempty"`
	PageIdCurDoc  *int32  `json:"PageIdCurDoc,omitempty" xml:"PageIdCurDoc,omitempty"`
}

func (s QueryDocParserStatusResponseBodyDataOutputFormatResultPages) String() string {
	return dara.Prettify(s)
}

func (s QueryDocParserStatusResponseBodyDataOutputFormatResultPages) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) GetImageHeight() *int32 {
	return s.ImageHeight
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) GetImageWidth() *int32 {
	return s.ImageWidth
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) GetPageIdAllDocs() *int32 {
	return s.PageIdAllDocs
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) GetPageIdCurDoc() *int32 {
	return s.PageIdCurDoc
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) SetImageHeight(v int32) *QueryDocParserStatusResponseBodyDataOutputFormatResultPages {
	s.ImageHeight = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) SetImageUrl(v string) *QueryDocParserStatusResponseBodyDataOutputFormatResultPages {
	s.ImageUrl = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) SetImageWidth(v int32) *QueryDocParserStatusResponseBodyDataOutputFormatResultPages {
	s.ImageWidth = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) SetPageIdAllDocs(v int32) *QueryDocParserStatusResponseBodyDataOutputFormatResultPages {
	s.PageIdAllDocs = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) SetPageIdCurDoc(v int32) *QueryDocParserStatusResponseBodyDataOutputFormatResultPages {
	s.PageIdCurDoc = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyDataOutputFormatResultPages) Validate() error {
	return dara.Validate(s)
}
