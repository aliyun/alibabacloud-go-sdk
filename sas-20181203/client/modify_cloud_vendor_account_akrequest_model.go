// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudVendorAccountAKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthIds(v string) *ModifyCloudVendorAccountAKRequest
	GetAuthIds() *string
	SetAuthModules(v []*string) *ModifyCloudVendorAccountAKRequest
	GetAuthModules() []*string
	SetCtdrCloudUserId(v string) *ModifyCloudVendorAccountAKRequest
	GetCtdrCloudUserId() *string
	SetDomain(v string) *ModifyCloudVendorAccountAKRequest
	GetDomain() *string
	SetExtendInfo(v string) *ModifyCloudVendorAccountAKRequest
	GetExtendInfo() *string
	SetLang(v string) *ModifyCloudVendorAccountAKRequest
	GetLang() *string
	SetRegions(v []*string) *ModifyCloudVendorAccountAKRequest
	GetRegions() []*string
	SetSecretId(v string) *ModifyCloudVendorAccountAKRequest
	GetSecretId() *string
	SetSecretKey(v string) *ModifyCloudVendorAccountAKRequest
	GetSecretKey() *string
	SetStatus(v int32) *ModifyCloudVendorAccountAKRequest
	GetStatus() *int32
	SetSubscriptionIds(v []*string) *ModifyCloudVendorAccountAKRequest
	GetSubscriptionIds() []*string
	SetTenantId(v string) *ModifyCloudVendorAccountAKRequest
	GetTenantId() *string
	SetVendorAuthAlias(v string) *ModifyCloudVendorAccountAKRequest
	GetVendorAuthAlias() *string
}

type ModifyCloudVendorAccountAKRequest struct {
	// The unique ID of the AccessKey pair.
	//
	// >  You can call the [DescribeCloudVendorAccountAKList](~~DescribeCloudVendorAccountAKList~~) operation to query the unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2832
	AuthIds *string `json:"AuthIds,omitempty" xml:"AuthIds,omitempty"`
	// The modules that are associated with the AccessKey pair. Valid values:
	//
	// 	- **HOST**: host.
	//
	// 	- **CSPM**: configuration assessment.
	//
	// 	- **SIEM**: Cloud Threat Detection and Response (CTDR).
	//
	// 	- **TRIAL**: log audit.
	//
	// >  You can call the [GetSupportedModules](~~GetSupportedModules~~) operation to query the supported modules.
	AuthModules []*string `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	// Account ID.
	//
	// > The account ID of the connected cloud vendor, required when the permission description includes threat analysis and response.
	//
	// example:
	//
	// azure_demo_1
	CtdrCloudUserId *string `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// Access account domain. Values:
	//
	// -  **china**: China
	//
	// -  **global**: Global
	//
	// -  **europe**: Huawei Europe
	//
	// > This parameter is only valid and required for **Vendor*	- being **HUAWEICLOUD**, **Azure**, **AWS**, or **VOLCENGINE**.
	//
	// example:
	//
	// global
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Extended information.
	//
	// > Used to record extended information from different vendors.
	//
	// > For Google Cloud, which is accessed through a service account, ExtendInfo stores a JSON-formatted service key file, excluding the private_key_id and zprivate_key fields. The file includes the following fields: type, project_id, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url, universe_domain.
	//
	// example:
	//
	// {\\"product\\":\\"webFirewall\\",\\"remark\\":\\"remark\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions that are examined during AccessKey pair authentication.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// ID of the AK parameter. Values:
	//
	// 1. When AkType is primary:
	//
	// - **Tencent**: AccessKeyId of the main account
	//
	// - **HUAWEICLOUD**: AccessKeyId of the main account
	//
	// - **Azure**: ClientId
	//
	// - **AWS**: AccessKeyId of the main account
	//
	// - **VOLCENGINE**: AccessKeyId of the main account
	//
	// 2. When AkType is sub:
	//
	// - **Tencent**: AccessKeyId of the sub-account
	//
	// - **HUAWEICLOUD**: AccessKeyId of the sub-account
	//
	// - **Azure**: ClientId
	//
	// - **AWS**: AccessKeyId of the sub-account
	//
	// - **VOLCENGINE**: AccessKeyId of the sub-account
	//
	// - **google**: private_key_id
	//
	// > If AkType is **primary**, this value is the SecretID of the main account from another cloud. If AkType is **sub**, this value is the Access Key ID of the sub-account from another cloud. For **Azure**, there is no distinction, and this value is the **appId*	- of the authentication information. Google Cloud is accessed through a service account, with AkType defaulting to sub, and this value is taken from the private_key_id attribute in the JSON format service key file.
	//
	// example:
	//
	// S3D6c4O***
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The AccessKey secret.
	//
	// >  If AkType is set to **primary**, you must set SecretKey to the AccessKey secret of the third-party master account. If AkType is set to **sub**, you must set SecretKey to the AccessKey secret of the third-party sub-account. This parameter value does not change for a **Microsoft Azure account**. For an Azure account, set this parameter to the **password*	- that is used for authentication.
	//
	// example:
	//
	// AE6SLd****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- **0**: enabled.
	//
	// 	- **1**: disabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of subscriptions.
	//
	// >  This parameter takes effect only when Vendor is set to Azure.
	SubscriptionIds []*string `json:"SubscriptionIds,omitempty" xml:"SubscriptionIds,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// >  This parameter takes effect only when Vendor is set to Azure.
	//
	// example:
	//
	// 95304a97-339b-4de5-9a7d-cdbffaf****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the AccessKey pair.
	//
	// >  The account information of the third-party cloud servers.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
}

func (s ModifyCloudVendorAccountAKRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorAccountAKRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorAccountAKRequest) GetAuthIds() *string {
	return s.AuthIds
}

func (s *ModifyCloudVendorAccountAKRequest) GetAuthModules() []*string {
	return s.AuthModules
}

func (s *ModifyCloudVendorAccountAKRequest) GetCtdrCloudUserId() *string {
	return s.CtdrCloudUserId
}

func (s *ModifyCloudVendorAccountAKRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyCloudVendorAccountAKRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ModifyCloudVendorAccountAKRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyCloudVendorAccountAKRequest) GetRegions() []*string {
	return s.Regions
}

func (s *ModifyCloudVendorAccountAKRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *ModifyCloudVendorAccountAKRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *ModifyCloudVendorAccountAKRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyCloudVendorAccountAKRequest) GetSubscriptionIds() []*string {
	return s.SubscriptionIds
}

func (s *ModifyCloudVendorAccountAKRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ModifyCloudVendorAccountAKRequest) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *ModifyCloudVendorAccountAKRequest) SetAuthIds(v string) *ModifyCloudVendorAccountAKRequest {
	s.AuthIds = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetAuthModules(v []*string) *ModifyCloudVendorAccountAKRequest {
	s.AuthModules = v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetCtdrCloudUserId(v string) *ModifyCloudVendorAccountAKRequest {
	s.CtdrCloudUserId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetDomain(v string) *ModifyCloudVendorAccountAKRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetExtendInfo(v string) *ModifyCloudVendorAccountAKRequest {
	s.ExtendInfo = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetLang(v string) *ModifyCloudVendorAccountAKRequest {
	s.Lang = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetRegions(v []*string) *ModifyCloudVendorAccountAKRequest {
	s.Regions = v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetSecretId(v string) *ModifyCloudVendorAccountAKRequest {
	s.SecretId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetSecretKey(v string) *ModifyCloudVendorAccountAKRequest {
	s.SecretKey = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetStatus(v int32) *ModifyCloudVendorAccountAKRequest {
	s.Status = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetSubscriptionIds(v []*string) *ModifyCloudVendorAccountAKRequest {
	s.SubscriptionIds = v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetTenantId(v string) *ModifyCloudVendorAccountAKRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) SetVendorAuthAlias(v string) *ModifyCloudVendorAccountAKRequest {
	s.VendorAuthAlias = &v
	return s
}

func (s *ModifyCloudVendorAccountAKRequest) Validate() error {
	return dara.Validate(s)
}
