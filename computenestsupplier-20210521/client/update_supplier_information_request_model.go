// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSupplierInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverySettings(v *UpdateSupplierInformationRequestDeliverySettings) *UpdateSupplierInformationRequest
	GetDeliverySettings() *UpdateSupplierInformationRequestDeliverySettings
	SetOperationIp(v string) *UpdateSupplierInformationRequest
	GetOperationIp() *string
	SetOperationMfaPresent(v bool) *UpdateSupplierInformationRequest
	GetOperationMfaPresent() *bool
	SetRegionId(v string) *UpdateSupplierInformationRequest
	GetRegionId() *string
	SetSupplierDesc(v string) *UpdateSupplierInformationRequest
	GetSupplierDesc() *string
	SetSupplierLogo(v string) *UpdateSupplierInformationRequest
	GetSupplierLogo() *string
	SetSupplierUrl(v string) *UpdateSupplierInformationRequest
	GetSupplierUrl() *string
	SetSupportContacts(v []*UpdateSupplierInformationRequestSupportContacts) *UpdateSupplierInformationRequest
	GetSupportContacts() []*UpdateSupplierInformationRequestSupportContacts
}

type UpdateSupplierInformationRequest struct {
	// The delivery settings.
	DeliverySettings *UpdateSupplierInformationRequestDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The Ip of operation.
	//
	// example:
	//
	// 192.xxx.xxx.xxx/16,192.xxx.xxx.xxx
	OperationIp *string `json:"OperationIp,omitempty" xml:"OperationIp,omitempty"`
	// The MFA of operation.
	//
	// example:
	//
	// true
	OperationMfaPresent *bool `json:"OperationMfaPresent,omitempty" xml:"OperationMfaPresent,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/34DB-4F4C-9373-003AA060****.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider
	SupportContacts []*UpdateSupplierInformationRequestSupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
}

func (s UpdateSupplierInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupplierInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationRequest) GetDeliverySettings() *UpdateSupplierInformationRequestDeliverySettings {
	return s.DeliverySettings
}

func (s *UpdateSupplierInformationRequest) GetOperationIp() *string {
	return s.OperationIp
}

func (s *UpdateSupplierInformationRequest) GetOperationMfaPresent() *bool {
	return s.OperationMfaPresent
}

func (s *UpdateSupplierInformationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSupplierInformationRequest) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *UpdateSupplierInformationRequest) GetSupplierLogo() *string {
	return s.SupplierLogo
}

func (s *UpdateSupplierInformationRequest) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *UpdateSupplierInformationRequest) GetSupportContacts() []*UpdateSupplierInformationRequestSupportContacts {
	return s.SupportContacts
}

func (s *UpdateSupplierInformationRequest) SetDeliverySettings(v *UpdateSupplierInformationRequestDeliverySettings) *UpdateSupplierInformationRequest {
	s.DeliverySettings = v
	return s
}

func (s *UpdateSupplierInformationRequest) SetOperationIp(v string) *UpdateSupplierInformationRequest {
	s.OperationIp = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetOperationMfaPresent(v bool) *UpdateSupplierInformationRequest {
	s.OperationMfaPresent = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetRegionId(v string) *UpdateSupplierInformationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupplierDesc(v string) *UpdateSupplierInformationRequest {
	s.SupplierDesc = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupplierLogo(v string) *UpdateSupplierInformationRequest {
	s.SupplierLogo = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupplierUrl(v string) *UpdateSupplierInformationRequest {
	s.SupplierUrl = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupportContacts(v []*UpdateSupplierInformationRequestSupportContacts) *UpdateSupplierInformationRequest {
	s.SupportContacts = v
	return s
}

func (s *UpdateSupplierInformationRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateSupplierInformationRequestDeliverySettings struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// mybucket
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Specifies whether to enable screencast delivery to Object Storage Service (OSS). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	OssEnabled *bool `json:"OssEnabled,omitempty" xml:"OssEnabled,omitempty"`
	// The number of days for which the screencasts are saved.
	//
	// example:
	//
	// 7
	OssExpirationDays *int64 `json:"OssExpirationDays,omitempty" xml:"OssExpirationDays,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// path1/path2/
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s UpdateSupplierInformationRequestDeliverySettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupplierInformationRequestDeliverySettings) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationRequestDeliverySettings) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateSupplierInformationRequestDeliverySettings) GetOssEnabled() *bool {
	return s.OssEnabled
}

func (s *UpdateSupplierInformationRequestDeliverySettings) GetOssExpirationDays() *int64 {
	return s.OssExpirationDays
}

func (s *UpdateSupplierInformationRequestDeliverySettings) GetOssPath() *string {
	return s.OssPath
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssBucketName(v string) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssEnabled(v bool) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssExpirationDays(v int64) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssPath(v string) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssPath = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) Validate() error {
	return dara.Validate(s)
}

type UpdateSupplierInformationRequestSupportContacts struct {
	// The type of  contact information
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of contact information
	//
	// example:
	//
	// supplier@example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSupplierInformationRequestSupportContacts) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupplierInformationRequestSupportContacts) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationRequestSupportContacts) GetType() *string {
	return s.Type
}

func (s *UpdateSupplierInformationRequestSupportContacts) GetValue() *string {
	return s.Value
}

func (s *UpdateSupplierInformationRequestSupportContacts) SetType(v string) *UpdateSupplierInformationRequestSupportContacts {
	s.Type = &v
	return s
}

func (s *UpdateSupplierInformationRequestSupportContacts) SetValue(v string) *UpdateSupplierInformationRequestSupportContacts {
	s.Value = &v
	return s
}

func (s *UpdateSupplierInformationRequestSupportContacts) Validate() error {
	return dara.Validate(s)
}
