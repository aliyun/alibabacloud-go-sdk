// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TemplateListResponseBody
	GetCode() *int64
	SetMessage(v string) *TemplateListResponseBody
	GetMessage() *string
	SetModel(v []*TemplateListResponseBodyModel) *TemplateListResponseBody
	GetModel() []*TemplateListResponseBodyModel
	SetRequestId(v string) *TemplateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TemplateListResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *TemplateListResponseBody
	GetTimestamp() *int64
}

type TemplateListResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*TemplateListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *TemplateListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TemplateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TemplateListResponseBody) GetModel() []*TemplateListResponseBodyModel {
	return s.Model
}

func (s *TemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TemplateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TemplateListResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TemplateListResponseBody) SetCode(v int64) *TemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *TemplateListResponseBody) SetMessage(v string) *TemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *TemplateListResponseBody) SetModel(v []*TemplateListResponseBodyModel) *TemplateListResponseBody {
	s.Model = v
	return s
}

func (s *TemplateListResponseBody) SetRequestId(v string) *TemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TemplateListResponseBody) SetSuccess(v bool) *TemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *TemplateListResponseBody) SetTimestamp(v int64) *TemplateListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TemplateListResponseBody) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TemplateListResponseBodyModel struct {
	// 意向标签
	IntentTags []map[string]interface{} `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 个性标签
	PersonalityTags []*string `json:"PersonalityTags,omitempty" xml:"PersonalityTags,omitempty" type:"Repeated"`
	// 话术所需参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// AI话术ID
	//
	// example:
	//
	// 59
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板名称
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 55
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s TemplateListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s TemplateListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TemplateListResponseBodyModel) GetIntentTags() []map[string]interface{} {
	return s.IntentTags
}

func (s *TemplateListResponseBodyModel) GetPersonalityTags() []*string {
	return s.PersonalityTags
}

func (s *TemplateListResponseBodyModel) GetProperties() *string {
	return s.Properties
}

func (s *TemplateListResponseBodyModel) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *TemplateListResponseBodyModel) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TemplateListResponseBodyModel) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *TemplateListResponseBodyModel) SetIntentTags(v []map[string]interface{}) *TemplateListResponseBodyModel {
	s.IntentTags = v
	return s
}

func (s *TemplateListResponseBodyModel) SetPersonalityTags(v []*string) *TemplateListResponseBodyModel {
	s.PersonalityTags = v
	return s
}

func (s *TemplateListResponseBodyModel) SetProperties(v string) *TemplateListResponseBodyModel {
	s.Properties = &v
	return s
}

func (s *TemplateListResponseBodyModel) SetTemplateId(v int64) *TemplateListResponseBodyModel {
	s.TemplateId = &v
	return s
}

func (s *TemplateListResponseBodyModel) SetTemplateName(v string) *TemplateListResponseBodyModel {
	s.TemplateName = &v
	return s
}

func (s *TemplateListResponseBodyModel) SetTemplateType(v int64) *TemplateListResponseBodyModel {
	s.TemplateType = &v
	return s
}

func (s *TemplateListResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
