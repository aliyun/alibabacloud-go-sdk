// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemPropertyTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSystemPropertyTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSystemPropertyTemplatesResponseBody
	GetRequestId() *string
	SetSystemPropertyTemplateModel(v []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) *DescribeSystemPropertyTemplatesResponseBody
	GetSystemPropertyTemplateModel() []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel
	SetTotalCount(v int32) *DescribeSystemPropertyTemplatesResponseBody
	GetTotalCount() *int32
}

type DescribeSystemPropertyTemplatesResponseBody struct {
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId                   *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemPropertyTemplateModel []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel `json:"SystemPropertyTemplateModel,omitempty" xml:"SystemPropertyTemplateModel,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSystemPropertyTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemPropertyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemPropertyTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSystemPropertyTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemPropertyTemplatesResponseBody) GetSystemPropertyTemplateModel() []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	return s.SystemPropertyTemplateModel
}

func (s *DescribeSystemPropertyTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSystemPropertyTemplatesResponseBody) SetNextToken(v string) *DescribeSystemPropertyTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBody) SetRequestId(v string) *DescribeSystemPropertyTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBody) SetSystemPropertyTemplateModel(v []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) *DescribeSystemPropertyTemplatesResponseBody {
	s.SystemPropertyTemplateModel = v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBody) SetTotalCount(v int32) *DescribeSystemPropertyTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBody) Validate() error {
	if s.SystemPropertyTemplateModel != nil {
		for _, item := range s.SystemPropertyTemplateModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel struct {
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// example:
	//
	// https://filepath****.com
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// init
	Status             *string                                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemPropertyInfo *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty" type:"Struct"`
	// example:
	//
	// ap-0caoenwutkkx****
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GoString() string {
	return s.String()
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GetEnableAuto() *bool {
	return s.EnableAuto
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GetSystemPropertyInfo() *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo {
	return s.SystemPropertyInfo
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) SetEnableAuto(v bool) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	s.EnableAuto = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) SetFilePath(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	s.FilePath = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) SetStatus(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	s.Status = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) SetSystemPropertyInfo(v *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	s.SystemPropertyInfo = v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) SetTemplateId(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	s.TemplateId = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) SetTemplateName(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel {
	s.TemplateName = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModel) Validate() error {
	if s.SystemPropertyInfo != nil {
		if err := s.SystemPropertyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo struct {
	CustomPropertyInfos []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos `json:"CustomPropertyInfos,omitempty" xml:"CustomPropertyInfos,omitempty" type:"Repeated"`
	// example:
	//
	// null
	RoProductDevice *string `json:"RoProductDevice,omitempty" xml:"RoProductDevice,omitempty"`
}

func (s DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) GoString() string {
	return s.String()
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) GetCustomPropertyInfos() []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos {
	return s.CustomPropertyInfos
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) GetRoProductDevice() *string {
	return s.RoProductDevice
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) SetCustomPropertyInfos(v []*DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo {
	s.CustomPropertyInfos = v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) SetRoProductDevice(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo {
	s.RoProductDevice = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfo) Validate() error {
	if s.CustomPropertyInfos != nil {
		for _, item := range s.CustomPropertyInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos struct {
	// example:
	//
	// propKey
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// example:
	//
	// propValue
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) GoString() string {
	return s.String()
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) GetPropertyName() *string {
	return s.PropertyName
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) SetPropertyName(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos {
	s.PropertyName = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) SetPropertyValue(v string) *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos {
	s.PropertyValue = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponseBodySystemPropertyTemplateModelSystemPropertyInfoCustomPropertyInfos) Validate() error {
	return dara.Validate(s)
}
