// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDialogueTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *ListDialogueTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDialogueTemplateResponseBody
	GetRequestId() *string
	SetResult(v []*ListDialogueTemplateResponseBodyResult) *ListDialogueTemplateResponseBody
	GetResult() []*ListDialogueTemplateResponseBodyResult
}

type ListDialogueTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDialogueTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDialogueTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDialogueTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDialogueTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDialogueTemplateResponseBody) GetResult() []*ListDialogueTemplateResponseBodyResult {
	return s.Result
}

func (s *ListDialogueTemplateResponseBody) SetCode(v int32) *ListDialogueTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialogueTemplateResponseBody) SetMessage(v string) *ListDialogueTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialogueTemplateResponseBody) SetRequestId(v string) *ListDialogueTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialogueTemplateResponseBody) SetResult(v []*ListDialogueTemplateResponseBodyResult) *ListDialogueTemplateResponseBody {
	s.Result = v
	return s
}

func (s *ListDialogueTemplateResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDialogueTemplateResponseBodyResult struct {
	TemplateDetail *ListDialogueTemplateResponseBodyResultTemplateDetail `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty" type:"Struct"`
	// example:
	//
	// 4
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 物品多轮模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDialogueTemplateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResult) GetTemplateDetail() *ListDialogueTemplateResponseBodyResultTemplateDetail {
	return s.TemplateDetail
}

func (s *ListDialogueTemplateResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListDialogueTemplateResponseBodyResult) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListDialogueTemplateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDialogueTemplateResponseBodyResult) SetTemplateDetail(v *ListDialogueTemplateResponseBodyResultTemplateDetail) *ListDialogueTemplateResponseBodyResult {
	s.TemplateDetail = v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) SetTemplateId(v int64) *ListDialogueTemplateResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) SetTemplateName(v string) *ListDialogueTemplateResponseBodyResult {
	s.TemplateName = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) SetType(v string) *ListDialogueTemplateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResult) Validate() error {
	if s.TemplateDetail != nil {
		if err := s.TemplateDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDialogueTemplateResponseBodyResultTemplateDetail struct {
	FirstDialogueTemplate  *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate  `json:"FirstDialogueTemplate,omitempty" xml:"FirstDialogueTemplate,omitempty" type:"Struct"`
	SecondDialogueTemplate *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate `json:"SecondDialogueTemplate,omitempty" xml:"SecondDialogueTemplate,omitempty" type:"Struct"`
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetail) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetail) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) GetFirstDialogueTemplate() *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	return s.FirstDialogueTemplate
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) GetSecondDialogueTemplate() *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate {
	return s.SecondDialogueTemplate
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) SetFirstDialogueTemplate(v *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) *ListDialogueTemplateResponseBodyResultTemplateDetail {
	s.FirstDialogueTemplate = v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) SetSecondDialogueTemplate(v *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) *ListDialogueTemplateResponseBodyResultTemplateDetail {
	s.SecondDialogueTemplate = v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetail) Validate() error {
	if s.FirstDialogueTemplate != nil {
		if err := s.FirstDialogueTemplate.Validate(); err != nil {
			return err
		}
	}
	if s.SecondDialogueTemplate != nil {
		if err := s.SecondDialogueTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate struct {
	// example:
	//
	// ${goodsName}${price}元，请问需要服务员送来吗？
	NonzeroPriceYesAnswer *string `json:"NonzeroPriceYesAnswer,omitempty" xml:"NonzeroPriceYesAnswer,omitempty"`
	// example:
	//
	// 对不起，暂时不提供此物品。
	ZeroPriceNoAnswer *string `json:"ZeroPriceNoAnswer,omitempty" xml:"ZeroPriceNoAnswer,omitempty"`
	// example:
	//
	// 好的，服务员会尽快送来。
	ZeroPriceYesAnswer *string `json:"ZeroPriceYesAnswer,omitempty" xml:"ZeroPriceYesAnswer,omitempty"`
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) GetNonzeroPriceYesAnswer() *string {
	return s.NonzeroPriceYesAnswer
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) GetZeroPriceNoAnswer() *string {
	return s.ZeroPriceNoAnswer
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) GetZeroPriceYesAnswer() *string {
	return s.ZeroPriceYesAnswer
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) SetNonzeroPriceYesAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	s.NonzeroPriceYesAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) SetZeroPriceNoAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	s.ZeroPriceNoAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) SetZeroPriceYesAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate {
	s.ZeroPriceYesAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailFirstDialogueTemplate) Validate() error {
	return dara.Validate(s)
}

type ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate struct {
	// example:
	//
	// 好的，已取消。
	NonzeroPriceNoAnswer *string `json:"NonzeroPriceNoAnswer,omitempty" xml:"NonzeroPriceNoAnswer,omitempty"`
	// example:
	//
	// 好的，服务员会尽快送来${goodsName}
	NonzeroPriceYesAnswer *string `json:"NonzeroPriceYesAnswer,omitempty" xml:"NonzeroPriceYesAnswer,omitempty"`
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) GetNonzeroPriceNoAnswer() *string {
	return s.NonzeroPriceNoAnswer
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) GetNonzeroPriceYesAnswer() *string {
	return s.NonzeroPriceYesAnswer
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) SetNonzeroPriceNoAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate {
	s.NonzeroPriceNoAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) SetNonzeroPriceYesAnswer(v string) *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate {
	s.NonzeroPriceYesAnswer = &v
	return s
}

func (s *ListDialogueTemplateResponseBodyResultTemplateDetailSecondDialogueTemplate) Validate() error {
	return dara.Validate(s)
}
