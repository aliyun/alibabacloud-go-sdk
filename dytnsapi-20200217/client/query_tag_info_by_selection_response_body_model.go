// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagInfoBySelectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTagInfoBySelectionResponseBody
	GetCode() *string
	SetData(v []*QueryTagInfoBySelectionResponseBodyData) *QueryTagInfoBySelectionResponseBody
	GetData() []*QueryTagInfoBySelectionResponseBodyData
	SetMessage(v string) *QueryTagInfoBySelectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTagInfoBySelectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTagInfoBySelectionResponseBody
	GetSuccess() *bool
}

type QueryTagInfoBySelectionResponseBody struct {
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*QueryTagInfoBySelectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1C3B8084-3A7D-570B-BC84-BF945A9CF65E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTagInfoBySelectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTagInfoBySelectionResponseBody) GetData() []*QueryTagInfoBySelectionResponseBodyData {
	return s.Data
}

func (s *QueryTagInfoBySelectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTagInfoBySelectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTagInfoBySelectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTagInfoBySelectionResponseBody) SetCode(v string) *QueryTagInfoBySelectionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetData(v []*QueryTagInfoBySelectionResponseBodyData) *QueryTagInfoBySelectionResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetMessage(v string) *QueryTagInfoBySelectionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetRequestId(v string) *QueryTagInfoBySelectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetSuccess(v bool) *QueryTagInfoBySelectionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) Validate() error {
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

type QueryTagInfoBySelectionResponseBodyData struct {
	// The list of available authorization codes.
	AuthCodeList   []*string `json:"AuthCodeList,omitempty" xml:"AuthCodeList,omitempty" type:"Repeated"`
	ComplexityType *string   `json:"ComplexityType,omitempty" xml:"ComplexityType,omitempty"`
	// The URL for the API demo.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/388997.html?spm=a2c4g.2573870.0.0.3aa921cbOrtqJz
	DemoAddress *string `json:"DemoAddress,omitempty" xml:"DemoAddress,omitempty"`
	// The URL for the API documentation.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/388997.html?spm=a2c4g.2573870.0.0.3aa921cbOrtqJz
	DocAddress *string `json:"DocAddress,omitempty" xml:"DocAddress,omitempty"`
	// The URL for the definitions of the enumerated values.
	//
	// example:
	//
	// example.aliyundoc.com
	EnumDefinitionAddress *string `json:"EnumDefinitionAddress,omitempty" xml:"EnumDefinitionAddress,omitempty"`
	// The flow name.
	//
	// example:
	//
	// process name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The industry ID.
	//
	// example:
	//
	// 83
	IndustryId *int64 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// The industry name.
	//
	// example:
	//
	// logistics
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// The list of tag parameters.
	ParamList           []*QueryTagInfoBySelectionResponseBodyDataParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	RichTextDescription *string                                             `json:"RichTextDescription,omitempty" xml:"RichTextDescription,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 41
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// General scenario
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 31
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// example:
	//
	// Number ownership
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryTagInfoBySelectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetAuthCodeList() []*string {
	return s.AuthCodeList
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetComplexityType() *string {
	return s.ComplexityType
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetDemoAddress() *string {
	return s.DemoAddress
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetDocAddress() *string {
	return s.DocAddress
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetEnumDefinitionAddress() *string {
	return s.EnumDefinitionAddress
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetFlowName() *string {
	return s.FlowName
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetIndustryId() *int64 {
	return s.IndustryId
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetIndustryName() *string {
	return s.IndustryName
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetParamList() []*QueryTagInfoBySelectionResponseBodyDataParamList {
	return s.ParamList
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetRichTextDescription() *string {
	return s.RichTextDescription
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetSceneId() *int64 {
	return s.SceneId
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryTagInfoBySelectionResponseBodyData) GetTagName() *string {
	return s.TagName
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetAuthCodeList(v []*string) *QueryTagInfoBySelectionResponseBodyData {
	s.AuthCodeList = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetComplexityType(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.ComplexityType = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetDemoAddress(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.DemoAddress = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetDocAddress(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.DocAddress = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetEnumDefinitionAddress(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.EnumDefinitionAddress = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetFlowName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetIndustryId(v int64) *QueryTagInfoBySelectionResponseBodyData {
	s.IndustryId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetIndustryName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.IndustryName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetParamList(v []*QueryTagInfoBySelectionResponseBodyDataParamList) *QueryTagInfoBySelectionResponseBodyData {
	s.ParamList = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetRichTextDescription(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.RichTextDescription = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetSceneId(v int64) *QueryTagInfoBySelectionResponseBodyData {
	s.SceneId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetSceneName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.SceneName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetTagId(v int64) *QueryTagInfoBySelectionResponseBodyData {
	s.TagId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetTagName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.TagName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) Validate() error {
	if s.ParamList != nil {
		for _, item := range s.ParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTagInfoBySelectionResponseBodyDataParamList struct {
	// The English name of the parameter.
	//
	// example:
	//
	// preame
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The input hint.
	//
	// example:
	//
	// none
	Hint *string `json:"Hint,omitempty" xml:"Hint,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// false
	Must *bool `json:"Must,omitempty" xml:"Must,omitempty"`
	// The Chinese name of the parameter.
	//
	// example:
	//
	// none
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. The code that corresponds to EnumUIWidgetTypes.
	//
	// example:
	//
	// aqzx
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The definitions of the enumerated values such as Code or Desc.
	ValueDict []*QueryTagInfoBySelectionResponseBodyDataParamListValueDict `json:"ValueDict,omitempty" xml:"ValueDict,omitempty" type:"Repeated"`
}

func (s QueryTagInfoBySelectionResponseBodyDataParamList) String() string {
	return dara.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBodyDataParamList) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) GetCode() *string {
	return s.Code
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) GetHint() *string {
	return s.Hint
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) GetMust() *bool {
	return s.Must
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) GetName() *string {
	return s.Name
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) GetType() *string {
	return s.Type
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) GetValueDict() []*QueryTagInfoBySelectionResponseBodyDataParamListValueDict {
	return s.ValueDict
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetCode(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Code = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetHint(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Hint = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetMust(v bool) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Must = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetName(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Name = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetType(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Type = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetValueDict(v []*QueryTagInfoBySelectionResponseBodyDataParamListValueDict) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.ValueDict = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) Validate() error {
	if s.ValueDict != nil {
		for _, item := range s.ValueDict {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTagInfoBySelectionResponseBodyDataParamListValueDict struct {
	// The English name.
	//
	// example:
	//
	// Aliyun
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Chinese name.
	//
	// example:
	//
	// 阿里云
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s QueryTagInfoBySelectionResponseBodyDataParamListValueDict) String() string {
	return dara.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBodyDataParamListValueDict) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) GetCode() *string {
	return s.Code
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) GetDesc() *string {
	return s.Desc
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) SetCode(v string) *QueryTagInfoBySelectionResponseBodyDataParamListValueDict {
	s.Code = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) SetDesc(v string) *QueryTagInfoBySelectionResponseBodyDataParamListValueDict {
	s.Desc = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) Validate() error {
	return dara.Validate(s)
}
