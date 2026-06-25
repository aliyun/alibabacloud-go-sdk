// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSystemPropertyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAuto(v bool) *CreateSystemPropertyTemplateRequest
	GetEnableAuto() *bool
	SetFilePath(v string) *CreateSystemPropertyTemplateRequest
	GetFilePath() *string
	SetSystemPropertyInfo(v *CreateSystemPropertyTemplateRequestSystemPropertyInfo) *CreateSystemPropertyTemplateRequest
	GetSystemPropertyInfo() *CreateSystemPropertyTemplateRequestSystemPropertyInfo
	SetTemplateName(v string) *CreateSystemPropertyTemplateRequest
	GetTemplateName() *string
}

type CreateSystemPropertyTemplateRequest struct {
	// Specifies whether to automatically generate preset system properties.
	//
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// The URL of the property template file. The API parses the file synchronously. An error is returned if the file format is invalid.
	//
	// > The file must be in the following format: `{ "properties":{"key1":"value1"}}`.
	//
	// example:
	//
	// https://filepath****.com
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The information about the system property template.
	SystemPropertyInfo *CreateSystemPropertyTemplateRequestSystemPropertyInfo `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty" type:"Struct"`
	// The name of the template. The name must meet the following requirements:
	//
	// - Be 2 to 32 characters in length.
	//
	// - Start with an uppercase or lowercase letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - Contain letters, digits, colons (:), underscores (_), and hyphens (-). Periods (.) are not supported.
	//
	// example:
	//
	// Template1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateSystemPropertyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSystemPropertyTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSystemPropertyTemplateRequest) GetEnableAuto() *bool {
	return s.EnableAuto
}

func (s *CreateSystemPropertyTemplateRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateSystemPropertyTemplateRequest) GetSystemPropertyInfo() *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	return s.SystemPropertyInfo
}

func (s *CreateSystemPropertyTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSystemPropertyTemplateRequest) SetEnableAuto(v bool) *CreateSystemPropertyTemplateRequest {
	s.EnableAuto = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequest) SetFilePath(v string) *CreateSystemPropertyTemplateRequest {
	s.FilePath = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequest) SetSystemPropertyInfo(v *CreateSystemPropertyTemplateRequestSystemPropertyInfo) *CreateSystemPropertyTemplateRequest {
	s.SystemPropertyInfo = v
	return s
}

func (s *CreateSystemPropertyTemplateRequest) SetTemplateName(v string) *CreateSystemPropertyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequest) Validate() error {
	if s.SystemPropertyInfo != nil {
		if err := s.SystemPropertyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSystemPropertyTemplateRequestSystemPropertyInfo struct {
	// A list of custom properties.
	CustomPropertyInfos []*CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos `json:"CustomPropertyInfos,omitempty" xml:"CustomPropertyInfos,omitempty" type:"Repeated"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBootloader *string `json:"RoBootloader,omitempty" xml:"RoBootloader,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildDisplayId *string `json:"RoBuildDisplayId,omitempty" xml:"RoBuildDisplayId,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildFingerprint *string `json:"RoBuildFingerprint,omitempty" xml:"RoBuildFingerprint,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildHost *string `json:"RoBuildHost,omitempty" xml:"RoBuildHost,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildId *string `json:"RoBuildId,omitempty" xml:"RoBuildId,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildProduct *string `json:"RoBuildProduct,omitempty" xml:"RoBuildProduct,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildTags *string `json:"RoBuildTags,omitempty" xml:"RoBuildTags,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildType *string `json:"RoBuildType,omitempty" xml:"RoBuildType,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoBuildUser *string `json:"RoBuildUser,omitempty" xml:"RoBuildUser,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoProductBoard *string `json:"RoProductBoard,omitempty" xml:"RoProductBoard,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoProductBrand *string `json:"RoProductBrand,omitempty" xml:"RoProductBrand,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoProductDevice *string `json:"RoProductDevice,omitempty" xml:"RoProductDevice,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoProductManufacturer *string `json:"RoProductManufacturer,omitempty" xml:"RoProductManufacturer,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RoProductModel *string `json:"RoProductModel,omitempty" xml:"RoProductModel,omitempty"`
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// null
	RwRoSerialNo *string `json:"RwRoSerialNo,omitempty" xml:"RwRoSerialNo,omitempty"`
}

func (s CreateSystemPropertyTemplateRequestSystemPropertyInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateSystemPropertyTemplateRequestSystemPropertyInfo) GoString() string {
	return s.String()
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetCustomPropertyInfos() []*CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos {
	return s.CustomPropertyInfos
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBootloader() *string {
	return s.RoBootloader
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildDisplayId() *string {
	return s.RoBuildDisplayId
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildFingerprint() *string {
	return s.RoBuildFingerprint
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildHost() *string {
	return s.RoBuildHost
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildId() *string {
	return s.RoBuildId
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildProduct() *string {
	return s.RoBuildProduct
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildTags() *string {
	return s.RoBuildTags
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildType() *string {
	return s.RoBuildType
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoBuildUser() *string {
	return s.RoBuildUser
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductBoard() *string {
	return s.RoProductBoard
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductBrand() *string {
	return s.RoProductBrand
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductDevice() *string {
	return s.RoProductDevice
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductManufacturer() *string {
	return s.RoProductManufacturer
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRoProductModel() *string {
	return s.RoProductModel
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) GetRwRoSerialNo() *string {
	return s.RwRoSerialNo
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetCustomPropertyInfos(v []*CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.CustomPropertyInfos = v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBootloader(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBootloader = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildDisplayId(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildDisplayId = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildFingerprint(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildFingerprint = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildHost(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildHost = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildId(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildId = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildProduct(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildProduct = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildTags(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildTags = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildType(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildType = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoBuildUser(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoBuildUser = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductBoard(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductBoard = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductBrand(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductBrand = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductDevice(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductDevice = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductManufacturer(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductManufacturer = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRoProductModel(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RoProductModel = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) SetRwRoSerialNo(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfo {
	s.RwRoSerialNo = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfo) Validate() error {
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

type CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos struct {
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

func (s CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) GoString() string {
	return s.String()
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) GetPropertyName() *string {
	return s.PropertyName
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) SetPropertyName(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos {
	s.PropertyName = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) SetPropertyValue(v string) *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos {
	s.PropertyValue = &v
	return s
}

func (s *CreateSystemPropertyTemplateRequestSystemPropertyInfoCustomPropertyInfos) Validate() error {
	return dara.Validate(s)
}
