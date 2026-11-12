// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateProductContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateProductContentResponseBody
	GetCode() *string
	SetData(v *GenerateProductContentResponseBodyData) *GenerateProductContentResponseBody
	GetData() *GenerateProductContentResponseBodyData
	SetMessage(v string) *GenerateProductContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateProductContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateProductContentResponseBody
	GetSuccess() *bool
}

type GenerateProductContentResponseBody struct {
	// The error code. This parameter is not returned for normal calls.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The processing result. Contains copy data upon success. Returns a fixed value when the business processing fails.
	Data *GenerateProductContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned for normal calls.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, used for troubleshooting. It may be empty in extreme exception cases.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the business processing is successful. Valid values:
	//
	// - true: The output passes the verification.
	//
	// - false: The business processing fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateProductContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateProductContentResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateProductContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateProductContentResponseBody) GetData() *GenerateProductContentResponseBodyData {
	return s.Data
}

func (s *GenerateProductContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateProductContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateProductContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateProductContentResponseBody) SetCode(v string) *GenerateProductContentResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateProductContentResponseBody) SetData(v *GenerateProductContentResponseBodyData) *GenerateProductContentResponseBody {
	s.Data = v
	return s
}

func (s *GenerateProductContentResponseBody) SetMessage(v string) *GenerateProductContentResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateProductContentResponseBody) SetRequestId(v string) *GenerateProductContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateProductContentResponseBody) SetSuccess(v bool) *GenerateProductContentResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateProductContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateProductContentResponseBodyData struct {
	// The actual copy type processed.
	//
	// example:
	//
	// title
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The total number of characters in data.text, including formatting characters such as spaces, punctuation marks, and line breaks.
	//
	// example:
	//
	// 54
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The code of the actual output language.
	//
	// example:
	//
	// ru
	OutputLanguage *string `json:"OutputLanguage,omitempty" xml:"OutputLanguage,omitempty"`
	// Returned only when bullet point processing is successful. The number of characters for each point, corresponding to data.points one by one.
	PointLengths []*int32 `json:"PointLengths,omitempty" xml:"PointLengths,omitempty" type:"Repeated"`
	// Returned only when bullet point processing is successful. The selling point texts arranged in output order.
	Points []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
	// The final copy. Generated keywords are separated by commas and spaces. Rewritten keywords are separated by spaces. Bullet points are separated by line breaks.
	//
	// example:
	//
	// Чашка Для Мороженого Из Нержавеющей Стали 14,5x10,8 см
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The usage information.
	UsageMap map[string]*int32 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s GenerateProductContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateProductContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateProductContentResponseBodyData) GetContentType() *string {
	return s.ContentType
}

func (s *GenerateProductContentResponseBodyData) GetLength() *int32 {
	return s.Length
}

func (s *GenerateProductContentResponseBodyData) GetOutputLanguage() *string {
	return s.OutputLanguage
}

func (s *GenerateProductContentResponseBodyData) GetPointLengths() []*int32 {
	return s.PointLengths
}

func (s *GenerateProductContentResponseBodyData) GetPoints() []*string {
	return s.Points
}

func (s *GenerateProductContentResponseBodyData) GetText() *string {
	return s.Text
}

func (s *GenerateProductContentResponseBodyData) GetUsageMap() map[string]*int32 {
	return s.UsageMap
}

func (s *GenerateProductContentResponseBodyData) SetContentType(v string) *GenerateProductContentResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *GenerateProductContentResponseBodyData) SetLength(v int32) *GenerateProductContentResponseBodyData {
	s.Length = &v
	return s
}

func (s *GenerateProductContentResponseBodyData) SetOutputLanguage(v string) *GenerateProductContentResponseBodyData {
	s.OutputLanguage = &v
	return s
}

func (s *GenerateProductContentResponseBodyData) SetPointLengths(v []*int32) *GenerateProductContentResponseBodyData {
	s.PointLengths = v
	return s
}

func (s *GenerateProductContentResponseBodyData) SetPoints(v []*string) *GenerateProductContentResponseBodyData {
	s.Points = v
	return s
}

func (s *GenerateProductContentResponseBodyData) SetText(v string) *GenerateProductContentResponseBodyData {
	s.Text = &v
	return s
}

func (s *GenerateProductContentResponseBodyData) SetUsageMap(v map[string]*int32) *GenerateProductContentResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *GenerateProductContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
