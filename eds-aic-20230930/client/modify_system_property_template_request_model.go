// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemPropertyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAuto(v bool) *ModifySystemPropertyTemplateRequest
	GetEnableAuto() *bool
	SetFilePath(v string) *ModifySystemPropertyTemplateRequest
	GetFilePath() *string
	SetSystemPropertyInfo(v *ModifySystemPropertyTemplateRequestSystemPropertyInfo) *ModifySystemPropertyTemplateRequest
	GetSystemPropertyInfo() *ModifySystemPropertyTemplateRequestSystemPropertyInfo
	SetTemplateId(v string) *ModifySystemPropertyTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *ModifySystemPropertyTemplateRequest
	GetTemplateName() *string
}

type ModifySystemPropertyTemplateRequest struct {
	// Specifies whether to automatically generate preset system properties.
	//
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// The URL of the property template file. The system synchronously parses the file. If the file format is invalid, a parsing error is returned.
	//
	// > File template format: `{ "properties":{"key1":"value1", "key2":"value2"}}`.
	//
	// example:
	//
	// https://filepath****.com
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The information about the system property template.
	SystemPropertyInfo *ModifySystemPropertyTemplateRequestSystemPropertyInfo `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty" type:"Struct"`
	// The ID of the property template.
	//
	// example:
	//
	// ap-angyvganxlf****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifySystemPropertyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemPropertyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifySystemPropertyTemplateRequest) GetEnableAuto() *bool {
	return s.EnableAuto
}

func (s *ModifySystemPropertyTemplateRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ModifySystemPropertyTemplateRequest) GetSystemPropertyInfo() *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	return s.SystemPropertyInfo
}

func (s *ModifySystemPropertyTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifySystemPropertyTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifySystemPropertyTemplateRequest) SetEnableAuto(v bool) *ModifySystemPropertyTemplateRequest {
	s.EnableAuto = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequest) SetFilePath(v string) *ModifySystemPropertyTemplateRequest {
	s.FilePath = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequest) SetSystemPropertyInfo(v *ModifySystemPropertyTemplateRequestSystemPropertyInfo) *ModifySystemPropertyTemplateRequest {
	s.SystemPropertyInfo = v
	return s
}

func (s *ModifySystemPropertyTemplateRequest) SetTemplateId(v string) *ModifySystemPropertyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequest) SetTemplateName(v string) *ModifySystemPropertyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequest) Validate() error {
	if s.SystemPropertyInfo != nil {
		if err := s.SystemPropertyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySystemPropertyTemplateRequestSystemPropertyInfo struct {
	// The information about custom properties.
	CustomPropertyInfos []*ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos `json:"CustomPropertyInfos,omitempty" xml:"CustomPropertyInfos,omitempty" type:"Repeated"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBootloader *string `json:"RoBootloader,omitempty" xml:"RoBootloader,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildDisplayId *string `json:"RoBuildDisplayId,omitempty" xml:"RoBuildDisplayId,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildFingerprint *string `json:"RoBuildFingerprint,omitempty" xml:"RoBuildFingerprint,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildHost *string `json:"RoBuildHost,omitempty" xml:"RoBuildHost,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildId *string `json:"RoBuildId,omitempty" xml:"RoBuildId,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildProduct *string `json:"RoBuildProduct,omitempty" xml:"RoBuildProduct,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildTags *string `json:"RoBuildTags,omitempty" xml:"RoBuildTags,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildType *string `json:"RoBuildType,omitempty" xml:"RoBuildType,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoBuildUser *string `json:"RoBuildUser,omitempty" xml:"RoBuildUser,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoProductBoard *string `json:"RoProductBoard,omitempty" xml:"RoProductBoard,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoProductBrand *string `json:"RoProductBrand,omitempty" xml:"RoProductBrand,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoProductDevice *string `json:"RoProductDevice,omitempty" xml:"RoProductDevice,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoProductManufacturer *string `json:"RoProductManufacturer,omitempty" xml:"RoProductManufacturer,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoProductModel *string `json:"RoProductModel,omitempty" xml:"RoProductModel,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RwRoSerialNo *string `json:"RwRoSerialNo,omitempty" xml:"RwRoSerialNo,omitempty"`
}

func (s ModifySystemPropertyTemplateRequestSystemPropertyInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemPropertyTemplateRequestSystemPropertyInfo) GoString() string {
	return s.String()
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetCustomPropertyInfos() []*ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos {
	return s.CustomPropertyInfos
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBootloader() *string {
	return s.RoBootloader
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildDisplayId() *string {
	return s.RoBuildDisplayId
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildFingerprint() *string {
	return s.RoBuildFingerprint
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildHost() *string {
	return s.RoBuildHost
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildId() *string {
	return s.RoBuildId
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildProduct() *string {
	return s.RoBuildProduct
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildTags() *string {
	return s.RoBuildTags
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildType() *string {
	return s.RoBuildType
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildUser() *string {
	return s.RoBuildUser
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductBoard() *string {
	return s.RoProductBoard
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductBrand() *string {
	return s.RoProductBrand
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductDevice() *string {
	return s.RoProductDevice
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductManufacturer() *string {
	return s.RoProductManufacturer
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductModel() *string {
	return s.RoProductModel
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) GetRwRoSerialNo() *string {
	return s.RwRoSerialNo
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetCustomPropertyInfos(v []*ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.CustomPropertyInfos = v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBootloader(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBootloader = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildDisplayId(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildDisplayId = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildFingerprint(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildFingerprint = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildHost(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildHost = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildId(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildId = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildProduct(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildProduct = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildTags(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildTags = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildType(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildType = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildUser(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildUser = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductBoard(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductBoard = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductBrand(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductBrand = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductDevice(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductDevice = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductManufacturer(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductManufacturer = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductModel(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductModel = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) SetRwRoSerialNo(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfo {
	s.RwRoSerialNo = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfo) Validate() error {
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

type ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos struct {
	// The key of the custom property.
	//
	// example:
	//
	// propKey
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The value of the custom property.
	//
	// example:
	//
	// propValue
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) GoString() string {
	return s.String()
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) GetPropertyName() *string {
	return s.PropertyName
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) SetPropertyName(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos {
	s.PropertyName = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) SetPropertyValue(v string) *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos {
	s.PropertyValue = &v
	return s
}

func (s *ModifySystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) Validate() error {
	return dara.Validate(s)
}
