// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorAccountAKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkType(v string) *AddCloudVendorAccountAKRequest
	GetAkType() *string
	SetAuthModules(v []*string) *AddCloudVendorAccountAKRequest
	GetAuthModules() []*string
	SetCtdrCloudUserId(v string) *AddCloudVendorAccountAKRequest
	GetCtdrCloudUserId() *string
	SetDomain(v string) *AddCloudVendorAccountAKRequest
	GetDomain() *string
	SetExtendInfo(v string) *AddCloudVendorAccountAKRequest
	GetExtendInfo() *string
	SetLang(v string) *AddCloudVendorAccountAKRequest
	GetLang() *string
	SetRegions(v []*string) *AddCloudVendorAccountAKRequest
	GetRegions() []*string
	SetSecretId(v string) *AddCloudVendorAccountAKRequest
	GetSecretId() *string
	SetSecretKey(v string) *AddCloudVendorAccountAKRequest
	GetSecretKey() *string
	SetSubscriptionIds(v []*string) *AddCloudVendorAccountAKRequest
	GetSubscriptionIds() []*string
	SetTenantId(v string) *AddCloudVendorAccountAKRequest
	GetTenantId() *string
	SetVendor(v string) *AddCloudVendorAccountAKRequest
	GetVendor() *string
	SetVendorAuthAlias(v string) *AddCloudVendorAccountAKRequest
	GetVendorAuthAlias() *string
}

type AddCloudVendorAccountAKRequest struct {
	// The type of the account to which the AccessKey pair belongs. Valid values:
	//
	// 	- **primary**: a primary account
	//
	// 	- **sub**: a sub-account
	//
	// This parameter is required.
	//
	// example:
	//
	// primary
	AkType *string `json:"AkType,omitempty" xml:"AkType,omitempty"`
	// The modules that are associated with the AccessKey pair.
	AuthModules     []*string `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	CtdrCloudUserId *string   `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The Active Directory (AD) domain. This parameter takes effect only when Vendor is set to Azure. Valid values:
	//
	// 	- **china**
	//
	// 	- **global**
	//
	// example:
	//
	// global
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The language of the content in the request and response messages. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions that are examined during AccessKey pair authentication. This parameter takes effect only when Vendor is set to AWS.
	//
	// >  You can call the [ListCloudVendorRegions](~~ListCloudVendorRegions~~) operation to query regions.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The AccessKey ID. Valid values:
	//
	// 1\\. If AkType is set to primary, specify this parameter based on the following description:
	//
	// 	- **Tencent**: Enter the AccessKey ID of a primary account on Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Enter the AccessKey ID of a primary account on Huawei Cloud.
	//
	// 	- **Azure**: Enter the AccessKey ID of a primary account on Microsoft Azure.
	//
	// 	- **AWS**: Enter the AccessKey ID of a primary account on AWS.
	//
	// 2\\. If AkType is set to sub, specify this parameter based on the following description:
	//
	// 	- **Tencent**: Enter the AccessKey ID of a sub-account on Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Enter the AccessKey ID of a sub-account on Huawei Cloud.
	//
	// 	- **Azure**: Enter the AccessKey ID of a sub-account on Microsoft Azure.
	//
	// 	- **AWS**: Enter the AccessKey ID of a sub-account on AWS.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45GLRV4SOT0YFB****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The AccessKey secret. Valid values:
	//
	// 1\\. If AkType is set to primary, specify this parameter based on the following description:
	//
	// 	- **Tencent**: Enter the AccessKey secret of a primary account on Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Enter the AccessKey secret of a primary account on Huawei Cloud.
	//
	// 	- **Azure**: Enter the AccessKey secret of a primary account on Microsoft Azure.
	//
	// 	- **AWS**: Enter the AccessKey secret of a primary account on AWS.
	//
	// 2\\. If AkType is set to sub, specify this parameter based on the following description:
	//
	// 	- **Tencent**: Enter the AccessKey secret of a sub-account on Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Enter the AccessKey secret of a sub-account on Huawei Cloud.
	//
	// 	- **Azure**: Enter the AccessKey secret of a sub-account on Microsoft Azure.
	//
	// 	- **AWS**: Enter the AccessKey secret of a sub-account on AWS.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE6SLd****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The subscription IDs. This parameter takes effect only when Vendor is set to Azure.
	SubscriptionIds []*string `json:"SubscriptionIds,omitempty" xml:"SubscriptionIds,omitempty" type:"Repeated"`
	// The tenant ID. This parameter takes effect only when Vendor is set to Azure.
	//
	// example:
	//
	// 95304a97-339b-4de5-9a7d-cdbffaf****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The cloud service provider. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud
	//
	// 	- **Azure**: Microsoft Azure
	//
	// 	- **AWS**: Amazon Web Services (AWS)
	//
	// This parameter is required.
	//
	// example:
	//
	// AWS
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the AccessKey pair.
	//
	// >  The account information of the third-party cloud servers.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
}

func (s AddCloudVendorAccountAKRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorAccountAKRequest) GoString() string {
	return s.String()
}

func (s *AddCloudVendorAccountAKRequest) GetAkType() *string {
	return s.AkType
}

func (s *AddCloudVendorAccountAKRequest) GetAuthModules() []*string {
	return s.AuthModules
}

func (s *AddCloudVendorAccountAKRequest) GetCtdrCloudUserId() *string {
	return s.CtdrCloudUserId
}

func (s *AddCloudVendorAccountAKRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddCloudVendorAccountAKRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *AddCloudVendorAccountAKRequest) GetLang() *string {
	return s.Lang
}

func (s *AddCloudVendorAccountAKRequest) GetRegions() []*string {
	return s.Regions
}

func (s *AddCloudVendorAccountAKRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *AddCloudVendorAccountAKRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *AddCloudVendorAccountAKRequest) GetSubscriptionIds() []*string {
	return s.SubscriptionIds
}

func (s *AddCloudVendorAccountAKRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddCloudVendorAccountAKRequest) GetVendor() *string {
	return s.Vendor
}

func (s *AddCloudVendorAccountAKRequest) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *AddCloudVendorAccountAKRequest) SetAkType(v string) *AddCloudVendorAccountAKRequest {
	s.AkType = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetAuthModules(v []*string) *AddCloudVendorAccountAKRequest {
	s.AuthModules = v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetCtdrCloudUserId(v string) *AddCloudVendorAccountAKRequest {
	s.CtdrCloudUserId = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetDomain(v string) *AddCloudVendorAccountAKRequest {
	s.Domain = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetExtendInfo(v string) *AddCloudVendorAccountAKRequest {
	s.ExtendInfo = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetLang(v string) *AddCloudVendorAccountAKRequest {
	s.Lang = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetRegions(v []*string) *AddCloudVendorAccountAKRequest {
	s.Regions = v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetSecretId(v string) *AddCloudVendorAccountAKRequest {
	s.SecretId = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetSecretKey(v string) *AddCloudVendorAccountAKRequest {
	s.SecretKey = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetSubscriptionIds(v []*string) *AddCloudVendorAccountAKRequest {
	s.SubscriptionIds = v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetTenantId(v string) *AddCloudVendorAccountAKRequest {
	s.TenantId = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetVendor(v string) *AddCloudVendorAccountAKRequest {
	s.Vendor = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetVendorAuthAlias(v string) *AddCloudVendorAccountAKRequest {
	s.VendorAuthAlias = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) Validate() error {
	return dara.Validate(s)
}
