// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslateImageBatchResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTranslateImageBatchResultResponseBody
	GetCode() *int32
	SetData(v *GetTranslateImageBatchResultResponseBodyData) *GetTranslateImageBatchResultResponseBody
	GetData() *GetTranslateImageBatchResultResponseBodyData
	SetMessage(v string) *GetTranslateImageBatchResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTranslateImageBatchResultResponseBody
	GetRequestId() *string
}

type GetTranslateImageBatchResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTranslateImageBatchResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DACD263C-C068-5116-83EC-117245AA35CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranslateImageBatchResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateImageBatchResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTranslateImageBatchResultResponseBody) GetData() *GetTranslateImageBatchResultResponseBodyData {
	return s.Data
}

func (s *GetTranslateImageBatchResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTranslateImageBatchResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranslateImageBatchResultResponseBody) SetCode(v int32) *GetTranslateImageBatchResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) SetData(v *GetTranslateImageBatchResultResponseBodyData) *GetTranslateImageBatchResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) SetMessage(v string) *GetTranslateImageBatchResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) SetRequestId(v string) *GetTranslateImageBatchResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTranslateImageBatchResultResponseBodyData struct {
	Result []*GetTranslateImageBatchResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTranslateImageBatchResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateImageBatchResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponseBodyData) GetResult() []*GetTranslateImageBatchResultResponseBodyDataResult {
	return s.Result
}

func (s *GetTranslateImageBatchResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTranslateImageBatchResultResponseBodyData) SetResult(v []*GetTranslateImageBatchResultResponseBodyDataResult) *GetTranslateImageBatchResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyData) SetStatus(v string) *GetTranslateImageBatchResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTranslateImageBatchResultResponseBodyDataResult struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	FinalImageUrl *string `json:"FinalImageUrl,omitempty" xml:"FinalImageUrl,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	InPaintingUrl *string `json:"InPaintingUrl,omitempty" xml:"InPaintingUrl,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	SourceImageUrl *string `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// {"TemplateJson": "Editor Template Json String	"}
	TemplateJson *string `json:"TemplateJson,omitempty" xml:"TemplateJson,omitempty"`
}

func (s GetTranslateImageBatchResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateImageBatchResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetCode() *int32 {
	return s.Code
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetFinalImageUrl() *string {
	return s.FinalImageUrl
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetInPaintingUrl() *string {
	return s.InPaintingUrl
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetMessage() *string {
	return s.Message
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetSourceImageUrl() *string {
	return s.SourceImageUrl
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) GetTemplateJson() *string {
	return s.TemplateJson
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetCode(v int32) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.Code = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetFinalImageUrl(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.FinalImageUrl = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetInPaintingUrl(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.InPaintingUrl = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetMessage(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.Message = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetSourceImageUrl(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.SourceImageUrl = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetSuccess(v bool) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.Success = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetTemplateJson(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.TemplateJson = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
