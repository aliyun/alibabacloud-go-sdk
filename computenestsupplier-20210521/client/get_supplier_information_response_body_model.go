// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplierInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcrNamespace(v string) *GetSupplierInformationResponseBody
	GetAcrNamespace() *string
	SetDeliverySettings(v *GetSupplierInformationResponseBodyDeliverySettings) *GetSupplierInformationResponseBody
	GetDeliverySettings() *GetSupplierInformationResponseBodyDeliverySettings
	SetEnableReseller(v bool) *GetSupplierInformationResponseBody
	GetEnableReseller() *bool
	SetOperationIp(v string) *GetSupplierInformationResponseBody
	GetOperationIp() *string
	SetOperationMfaPresent(v bool) *GetSupplierInformationResponseBody
	GetOperationMfaPresent() *bool
	SetRequestId(v string) *GetSupplierInformationResponseBody
	GetRequestId() *string
	SetSupplierDesc(v string) *GetSupplierInformationResponseBody
	GetSupplierDesc() *string
	SetSupplierLogo(v string) *GetSupplierInformationResponseBody
	GetSupplierLogo() *string
	SetSupplierName(v string) *GetSupplierInformationResponseBody
	GetSupplierName() *string
	SetSupplierUrl(v string) *GetSupplierInformationResponseBody
	GetSupplierUrl() *string
	SetSupportContacts(v []*GetSupplierInformationResponseBodySupportContacts) *GetSupplierInformationResponseBody
	GetSupportContacts() []*GetSupplierInformationResponseBodySupportContacts
}

type GetSupplierInformationResponseBody struct {
	// Acr container namespace
	//
	// example:
	//
	// computenest
	AcrNamespace *string `json:"AcrNamespace,omitempty" xml:"AcrNamespace,omitempty"`
	// The delivery settings.
	DeliverySettings *GetSupplierInformationResponseBodyDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// Whether to enable reseller
	//
	// example:
	//
	// true
	EnableReseller *bool `json:"EnableReseller,omitempty" xml:"EnableReseller,omitempty"`
	// The Ip of the operation.
	//
	// example:
	//
	// 10.xxx.xxx.xxx/101
	OperationIp *string `json:"OperationIp,omitempty" xml:"OperationIp,omitempty"`
	// The MFA of the operation.
	//
	// example:
	//
	// true
	OperationMfaPresent *bool `json:"OperationMfaPresent,omitempty" xml:"OperationMfaPresent,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider
	SupportContacts []*GetSupplierInformationResponseBodySupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
}

func (s GetSupplierInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupplierInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBody) GetAcrNamespace() *string {
	return s.AcrNamespace
}

func (s *GetSupplierInformationResponseBody) GetDeliverySettings() *GetSupplierInformationResponseBodyDeliverySettings {
	return s.DeliverySettings
}

func (s *GetSupplierInformationResponseBody) GetEnableReseller() *bool {
	return s.EnableReseller
}

func (s *GetSupplierInformationResponseBody) GetOperationIp() *string {
	return s.OperationIp
}

func (s *GetSupplierInformationResponseBody) GetOperationMfaPresent() *bool {
	return s.OperationMfaPresent
}

func (s *GetSupplierInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupplierInformationResponseBody) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *GetSupplierInformationResponseBody) GetSupplierLogo() *string {
	return s.SupplierLogo
}

func (s *GetSupplierInformationResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *GetSupplierInformationResponseBody) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *GetSupplierInformationResponseBody) GetSupportContacts() []*GetSupplierInformationResponseBodySupportContacts {
	return s.SupportContacts
}

func (s *GetSupplierInformationResponseBody) SetAcrNamespace(v string) *GetSupplierInformationResponseBody {
	s.AcrNamespace = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetDeliverySettings(v *GetSupplierInformationResponseBodyDeliverySettings) *GetSupplierInformationResponseBody {
	s.DeliverySettings = v
	return s
}

func (s *GetSupplierInformationResponseBody) SetEnableReseller(v bool) *GetSupplierInformationResponseBody {
	s.EnableReseller = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetOperationIp(v string) *GetSupplierInformationResponseBody {
	s.OperationIp = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetOperationMfaPresent(v bool) *GetSupplierInformationResponseBody {
	s.OperationMfaPresent = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetRequestId(v string) *GetSupplierInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierDesc(v string) *GetSupplierInformationResponseBody {
	s.SupplierDesc = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierLogo(v string) *GetSupplierInformationResponseBody {
	s.SupplierLogo = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierName(v string) *GetSupplierInformationResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierUrl(v string) *GetSupplierInformationResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupportContacts(v []*GetSupplierInformationResponseBodySupportContacts) *GetSupplierInformationResponseBody {
	s.SupportContacts = v
	return s
}

func (s *GetSupplierInformationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSupplierInformationResponseBodyDeliverySettings struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// mybucket
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Indicates whether screencast delivery to Object Storage Service (OSS) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
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

func (s GetSupplierInformationResponseBodyDeliverySettings) String() string {
	return dara.Prettify(s)
}

func (s GetSupplierInformationResponseBodyDeliverySettings) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) GetOssEnabled() *bool {
	return s.OssEnabled
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) GetOssExpirationDays() *int64 {
	return s.OssExpirationDays
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) GetOssPath() *string {
	return s.OssPath
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssBucketName(v string) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssEnabled(v bool) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssExpirationDays(v int64) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssPath(v string) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssPath = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) Validate() error {
	return dara.Validate(s)
}

type GetSupplierInformationResponseBodySupportContacts struct {
	// The type of contact information.
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of contact information.
	//
	// example:
	//
	// supplier@example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSupplierInformationResponseBodySupportContacts) String() string {
	return dara.Prettify(s)
}

func (s GetSupplierInformationResponseBodySupportContacts) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBodySupportContacts) GetType() *string {
	return s.Type
}

func (s *GetSupplierInformationResponseBodySupportContacts) GetValue() *string {
	return s.Value
}

func (s *GetSupplierInformationResponseBodySupportContacts) SetType(v string) *GetSupplierInformationResponseBodySupportContacts {
	s.Type = &v
	return s
}

func (s *GetSupplierInformationResponseBodySupportContacts) SetValue(v string) *GetSupplierInformationResponseBodySupportContacts {
	s.Value = &v
	return s
}

func (s *GetSupplierInformationResponseBodySupportContacts) Validate() error {
	return dara.Validate(s)
}
