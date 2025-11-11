// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWritingStylesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWritingStylesResponseBody
	GetCode() *string
	SetData(v []*ListWritingStylesResponseBodyData) *ListWritingStylesResponseBody
	GetData() []*ListWritingStylesResponseBodyData
	SetMaxResults(v int32) *ListWritingStylesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWritingStylesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWritingStylesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWritingStylesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListWritingStylesResponseBody
	GetSuccess() *string
	SetTotalCount(v int32) *ListWritingStylesResponseBody
	GetTotalCount() *int32
}

type ListWritingStylesResponseBody struct {
	// example:
	//
	// successful
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListWritingStylesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 58
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWritingStylesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWritingStylesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWritingStylesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWritingStylesResponseBody) GetData() []*ListWritingStylesResponseBodyData {
	return s.Data
}

func (s *ListWritingStylesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWritingStylesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWritingStylesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWritingStylesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWritingStylesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListWritingStylesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWritingStylesResponseBody) SetCode(v string) *ListWritingStylesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWritingStylesResponseBody) SetData(v []*ListWritingStylesResponseBodyData) *ListWritingStylesResponseBody {
	s.Data = v
	return s
}

func (s *ListWritingStylesResponseBody) SetMaxResults(v int32) *ListWritingStylesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWritingStylesResponseBody) SetMessage(v string) *ListWritingStylesResponseBody {
	s.Message = &v
	return s
}

func (s *ListWritingStylesResponseBody) SetNextToken(v string) *ListWritingStylesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWritingStylesResponseBody) SetRequestId(v string) *ListWritingStylesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWritingStylesResponseBody) SetSuccess(v string) *ListWritingStylesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWritingStylesResponseBody) SetTotalCount(v int32) *ListWritingStylesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWritingStylesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWritingStylesResponseBodyData struct {
	DistributeStepTemplateDefine *WritingStyleTemplateDefine `json:"DistributeStepTemplateDefine,omitempty" xml:"DistributeStepTemplateDefine,omitempty"`
	// example:
	//
	// false
	DistributeWriting *bool                       `json:"DistributeWriting,omitempty" xml:"DistributeWriting,omitempty"`
	Emoji             *string                     `json:"Emoji,omitempty" xml:"Emoji,omitempty"`
	StyleDescription  *string                     `json:"StyleDescription,omitempty" xml:"StyleDescription,omitempty"`
	StyleImage        *string                     `json:"StyleImage,omitempty" xml:"StyleImage,omitempty"`
	StyleKey          *string                     `json:"StyleKey,omitempty" xml:"StyleKey,omitempty"`
	StyleName         *string                     `json:"StyleName,omitempty" xml:"StyleName,omitempty"`
	TemplateDefine    *WritingStyleTemplateDefine `json:"TemplateDefine,omitempty" xml:"TemplateDefine,omitempty"`
}

func (s ListWritingStylesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWritingStylesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWritingStylesResponseBodyData) GetDistributeStepTemplateDefine() *WritingStyleTemplateDefine {
	return s.DistributeStepTemplateDefine
}

func (s *ListWritingStylesResponseBodyData) GetDistributeWriting() *bool {
	return s.DistributeWriting
}

func (s *ListWritingStylesResponseBodyData) GetEmoji() *string {
	return s.Emoji
}

func (s *ListWritingStylesResponseBodyData) GetStyleDescription() *string {
	return s.StyleDescription
}

func (s *ListWritingStylesResponseBodyData) GetStyleImage() *string {
	return s.StyleImage
}

func (s *ListWritingStylesResponseBodyData) GetStyleKey() *string {
	return s.StyleKey
}

func (s *ListWritingStylesResponseBodyData) GetStyleName() *string {
	return s.StyleName
}

func (s *ListWritingStylesResponseBodyData) GetTemplateDefine() *WritingStyleTemplateDefine {
	return s.TemplateDefine
}

func (s *ListWritingStylesResponseBodyData) SetDistributeStepTemplateDefine(v *WritingStyleTemplateDefine) *ListWritingStylesResponseBodyData {
	s.DistributeStepTemplateDefine = v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetDistributeWriting(v bool) *ListWritingStylesResponseBodyData {
	s.DistributeWriting = &v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetEmoji(v string) *ListWritingStylesResponseBodyData {
	s.Emoji = &v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetStyleDescription(v string) *ListWritingStylesResponseBodyData {
	s.StyleDescription = &v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetStyleImage(v string) *ListWritingStylesResponseBodyData {
	s.StyleImage = &v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetStyleKey(v string) *ListWritingStylesResponseBodyData {
	s.StyleKey = &v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetStyleName(v string) *ListWritingStylesResponseBodyData {
	s.StyleName = &v
	return s
}

func (s *ListWritingStylesResponseBodyData) SetTemplateDefine(v *WritingStyleTemplateDefine) *ListWritingStylesResponseBodyData {
	s.TemplateDefine = v
	return s
}

func (s *ListWritingStylesResponseBodyData) Validate() error {
	if s.DistributeStepTemplateDefine != nil {
		if err := s.DistributeStepTemplateDefine.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateDefine != nil {
		if err := s.TemplateDefine.Validate(); err != nil {
			return err
		}
	}
	return nil
}
