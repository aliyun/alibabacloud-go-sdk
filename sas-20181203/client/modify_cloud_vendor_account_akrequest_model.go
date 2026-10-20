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
	// The unique ID of the AccessKey (AK).
	//
	// >You can call the [DescribeCloudVendorAccountAKList](~~DescribeCloudVendorAccountAKList~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2832
	AuthIds *string `json:"AuthIds,omitempty" xml:"AuthIds,omitempty"`
	// The list of codes for the modules associated with the AK. Valid values:
	//
	// - **HOST**: host
	//
	// - **CSPM**: cloud product configuration check
	//
	// - **SIEM**: threat analysis and response
	//
	// - **TRIAL**: log audit
	//
	// >You can call the [GetSupportedModules](~~GetSupportedModules~~) operation to obtain the supported modules.
	AuthModules []*string `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	// The account ID.
	//
	// >The account ID of the connected cloud vendor. This parameter is required when the permission description includes threat analysis and response.
	//
	// example:
	//
	// azure_demo_1
	CtdrCloudUserId *string `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The domain to which the connected account belongs. Valid values:
	//
	// - **china**: China
	//
	// - **global**: global
	//
	// - **europe**: Huawei Cloud Europe site
	//
	// >This parameter is valid and required only when **Vendor*	- is set to **HUAWEICLOUD**, **Azure**, **AWS**, **VOLCENGINE**, **KingsoftCloud**, **UCloud**, **BaiduCloud**, or **OracleCloud**. For **KingsoftCloud*	- and **BaiduCloud**, this parameter is fixed to **china**. For **OracleCloud**, this parameter is fixed to **global**.
	//
	// example:
	//
	// global
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The extended information.
	//
	// >This parameter is used to record the extended information of different vendors.
	//
	// Google Cloud is accessed through a service account. ExtendInfo stores the JSON-formatted service key file, excluding the private_key_id and zprivate_key fields. The file contains the following fields: type, project_id, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url, and universe_domain. If the vendor is **OracleCloud**, this parameter contains the **fingerprint*	- key, which indicates the fingerprint.
	//
	// example:
	//
	// {\\"product\\":\\"webFirewall\\",\\"remark\\":\\"remark\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The language type of the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of regions used for AK information verification. This parameter is valid only when **Vendor*	- is set to **AWS*	- or **OracleCloud**.
	//
	// >You can call the [ListCloudVendorRegions](~~ListCloudVendorRegions~~) operation to obtain this parameter.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the AK parameter. Valid values:
	//
	// - If AkType is set to primary:
	//
	// - - **Tencent**: the AccessKey ID of the primary account
	//
	// - - **HUAWEICLOUD**: the AccessKey ID of the primary account
	//
	// - - **Azure**: the client ID
	//
	// - - **AWS**: the AccessKey ID of the primary account
	//
	// - - **VOLCENGINE**: the AccessKey ID of the primary account
	//
	// - - **KingsoftCloud**: the AccessKey ID of the primary account
	//
	// - - **UCloud**: the AccessKey ID of the primary account
	//
	// - If AkType is set to sub:
	//
	// - - **Tencent**: the AccessKey ID of the sub-account
	//
	// - - **HUAWEICLOUD**: the AccessKey ID of the sub-account
	//
	// - - **Azure**: the client ID
	//
	// - - **AWS**: the AccessKey ID of the sub-account
	//
	// - - **VOLCENGINE**: the AccessKey ID of the sub-account
	//
	// - - **google**: the private key ID
	//
	// - - **OracleCloud**: the Oracle Cloud Identifier (OCID) of the sub-user
	//
	// - - **BaiduCloud**: the AccessKey ID of the sub-account
	//
	// - - **KingsoftCloud**: the AccessKey ID of the sub-account
	//
	// - - **UCloud**: the AccessKey ID of the sub-account
	//
	// - - **StateCloud**: the AccessKey ID of the sub-account
	//
	// >If AkType is set to **primary**, this value is the SecretID of the primary account of the third-party cloud. If AkType is set to **sub**, this value is the AccessKey ID of the sub-account of the third-party cloud. For **Azure**, no distinction is made, and this value is the **appId*	- of the authentication information. Google Cloud is accessed through a service account. By default, AkType is set to sub, and this value is obtained from the private_key_id property of the JSON-formatted service key file. If the vendor is **OracleCloud**, this value is the OCID of the sub-user.
	//
	// example:
	//
	// S3D6c4O***
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The secret of the AK parameter. Valid values:
	//
	// - If AkType is set to primary:
	//
	// - - **Tencent**: the SecretAccessKey of the primary account
	//
	// - - **HUAWEICLOUD**: the SecretAccessKey of the primary account
	//
	// - - **Azure**: the client secret
	//
	// - - **AWS**: the SecretAccessKey of the primary account
	//
	// - - **VOLCENGINE**: the SecretAccessKey of the primary account
	//
	// - - **KingsoftCloud**: the SecretAccessKey of the primary account
	//
	// - - **UCloud**: the SecretAccessKey of the primary account
	//
	// - If AkType is set to sub:
	//
	// - - **Tencent**: the SecretAccessKey of the sub-account
	//
	// - - **HUAWEICLOUD**: the SecretAccessKey of the sub-account
	//
	// - - **Azure**: the client secret
	//
	// - - **AWS**: the SecretAccessKey of the sub-account
	//
	// - - **google**: the private key
	//
	// - - **OracleCloud**: the private key of the sub-user
	//
	// - - **BaiduCloud**: the SecretAccessKey of the sub-account
	//
	// - - **KingsoftCloud**: the SecretAccessKey of the sub-account
	//
	// - - **UCloud**: the SecretAccessKey of the sub-account
	//
	// - - **StateCloud**: the SecretAccessKey of the sub-account
	//
	// >If AkType is set to **primary**, this value is the Secret Access Key of the primary account of the third-party cloud. If AkType is set to **sub**, this value is the Secret Access Key of the sub-account of the third-party cloud. For **Azure**, no distinction is made, and this value is the **password*	- of the authentication information. Google Cloud is accessed through a service account. By default, AkType is set to sub, and this value is obtained from the private_key property of the JSON-formatted service key file. For **OracleCloud**, this value is obtained from the content of the corresponding private key file.
	//
	// example:
	//
	// AE6SLd****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The usage status of the AK. Valid values:
	//
	// - **0**: enabled
	//
	// - **1**: disabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of subscription IDs.
	//
	// > **[Deprecated]*	- This parameter is no longer valid.
	SubscriptionIds []*string `json:"SubscriptionIds,omitempty" xml:"SubscriptionIds,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// >This parameter is valid only when Vendor is set to **Azure*	- or **OracleCloud**.
	//
	// example:
	//
	// 95304a97-339b-4de5-9a7d-cdbffaf****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the AK account.
	//
	// >This parameter is used to identify the account information to which the third-party host assets belong.
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
