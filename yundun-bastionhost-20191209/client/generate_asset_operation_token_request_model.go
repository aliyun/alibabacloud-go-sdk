// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAssetOperationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetAccountId(v string) *GenerateAssetOperationTokenRequest
	GetAssetAccountId() *string
	SetAssetAccountName(v string) *GenerateAssetOperationTokenRequest
	GetAssetAccountName() *string
	SetAssetAccountPassword(v string) *GenerateAssetOperationTokenRequest
	GetAssetAccountPassword() *string
	SetAssetAccountProtocolName(v string) *GenerateAssetOperationTokenRequest
	GetAssetAccountProtocolName() *string
	SetAssetId(v string) *GenerateAssetOperationTokenRequest
	GetAssetId() *string
	SetAssetType(v string) *GenerateAssetOperationTokenRequest
	GetAssetType() *string
	SetDatabaseSchema(v string) *GenerateAssetOperationTokenRequest
	GetDatabaseSchema() *string
	SetInstanceId(v string) *GenerateAssetOperationTokenRequest
	GetInstanceId() *string
	SetLoginAttribute(v string) *GenerateAssetOperationTokenRequest
	GetLoginAttribute() *string
	SetOperationMode(v string) *GenerateAssetOperationTokenRequest
	GetOperationMode() *string
	SetOperationNote(v string) *GenerateAssetOperationTokenRequest
	GetOperationNote() *string
	SetRegionId(v string) *GenerateAssetOperationTokenRequest
	GetRegionId() *string
	SetSsoClient(v string) *GenerateAssetOperationTokenRequest
	GetSsoClient() *string
}

type GenerateAssetOperationTokenRequest struct {
	// The ID of the asset account for which you want to apply for an O&M token.
	//
	// > You must specify at least one of AssetAccountId and AssetAccountName. If both are specified, AssetAccountId takes precedence.
	//
	// example:
	//
	// 2
	AssetAccountId *string `json:"AssetAccountId,omitempty" xml:"AssetAccountId,omitempty"`
	// The name of the host account. If the account is a custom account, specify the actual account name.
	//
	// > If both AssetAccountId and AssetAccountName are specified, AssetAccountId takes precedence.
	//
	// example:
	//
	// root
	AssetAccountName *string `json:"AssetAccountName,omitempty" xml:"AssetAccountName,omitempty"`
	// The Base64-encoded password. This parameter is required when you apply for an O&M token for a custom account.
	//
	// example:
	//
	// dGVzdHBhc3N3b3Jk
	AssetAccountPassword *string `json:"AssetAccountPassword,omitempty" xml:"AssetAccountPassword,omitempty"`
	// The name of the protocol. Valid values:
	//
	// - **SSH**
	//
	// - **RDP**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **MySQL**
	//
	// - **SQLServer**
	//
	// > This parameter is required when you apply for an O&M token for a custom account.
	//
	// example:
	//
	// SSH
	AssetAccountProtocolName *string `json:"AssetAccountProtocolName,omitempty" xml:"AssetAccountProtocolName,omitempty"`
	// The ID of the asset for which you want to apply for an O&M token.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The Asset Type for which you want to apply for a token. Valid values:
	//
	// - **Host**: host
	//
	// - **Database**: database
	//
	// This parameter is required.
	//
	// example:
	//
	// Host
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The name of the database. This parameter is required when OperationMode is set to Sso, AssetAccountProtocolName is set to PostgreSQL or Oracle, and the account is a custom account.
	//
	// > This parameter is supported only in Bastionhost V3.2.44 and later.
	//
	// example:
	//
	// orcl
	DatabaseSchema *string `json:"DatabaseSchema,omitempty" xml:"DatabaseSchema,omitempty"`
	// The instance ID of the Bastionhost instance for which you want to apply for an O&M token.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The logon attribute. This parameter is required when OperationMode is set to Sso and the account is a custom Oracle account. Valid values:
	//
	// - **SERVICENAME**
	//
	// - **SID**
	//
	// > This parameter is supported only in Bastionhost V3.2.44 and later.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// SID
	LoginAttribute *string `json:"LoginAttribute,omitempty" xml:"LoginAttribute,omitempty"`
	// The O&M logon method. Valid values:
	//
	// - **WebToken**: obtain an O&M token
	//
	// - **Sso**: log on by using a local client
	//
	// > This parameter is supported only in Bastionhost V3.2.44 and later. If this parameter is left empty, the default value WebToken is used.
	//
	// example:
	//
	// Sso
	OperationMode *string `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	// The logon note. This parameter is required when the administrator enables the logon note feature in the control policy.
	//
	// example:
	//
	// comment
	OperationNote *string `json:"OperationNote,omitempty" xml:"OperationNote,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the local client to be opened for O&M operations on Linux assets. This parameter is required when OperationMode is set to Sso and the asset protocol is SSH. Valid values:
	//
	// - **ssh**: open an SSH client to perform O&M operations with SSH permissions
	//
	// - **sftp**: open an SFTP client to perform O&M operations with SFTP permissions
	//
	// > This parameter is supported only in Bastionhost V3.2.44 and later.
	//
	// example:
	//
	// ssh
	SsoClient *string `json:"SsoClient,omitempty" xml:"SsoClient,omitempty"`
}

func (s GenerateAssetOperationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAssetOperationTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateAssetOperationTokenRequest) GetAssetAccountId() *string {
	return s.AssetAccountId
}

func (s *GenerateAssetOperationTokenRequest) GetAssetAccountName() *string {
	return s.AssetAccountName
}

func (s *GenerateAssetOperationTokenRequest) GetAssetAccountPassword() *string {
	return s.AssetAccountPassword
}

func (s *GenerateAssetOperationTokenRequest) GetAssetAccountProtocolName() *string {
	return s.AssetAccountProtocolName
}

func (s *GenerateAssetOperationTokenRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *GenerateAssetOperationTokenRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *GenerateAssetOperationTokenRequest) GetDatabaseSchema() *string {
	return s.DatabaseSchema
}

func (s *GenerateAssetOperationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateAssetOperationTokenRequest) GetLoginAttribute() *string {
	return s.LoginAttribute
}

func (s *GenerateAssetOperationTokenRequest) GetOperationMode() *string {
	return s.OperationMode
}

func (s *GenerateAssetOperationTokenRequest) GetOperationNote() *string {
	return s.OperationNote
}

func (s *GenerateAssetOperationTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateAssetOperationTokenRequest) GetSsoClient() *string {
	return s.SsoClient
}

func (s *GenerateAssetOperationTokenRequest) SetAssetAccountId(v string) *GenerateAssetOperationTokenRequest {
	s.AssetAccountId = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetAssetAccountName(v string) *GenerateAssetOperationTokenRequest {
	s.AssetAccountName = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetAssetAccountPassword(v string) *GenerateAssetOperationTokenRequest {
	s.AssetAccountPassword = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetAssetAccountProtocolName(v string) *GenerateAssetOperationTokenRequest {
	s.AssetAccountProtocolName = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetAssetId(v string) *GenerateAssetOperationTokenRequest {
	s.AssetId = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetAssetType(v string) *GenerateAssetOperationTokenRequest {
	s.AssetType = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetDatabaseSchema(v string) *GenerateAssetOperationTokenRequest {
	s.DatabaseSchema = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetInstanceId(v string) *GenerateAssetOperationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetLoginAttribute(v string) *GenerateAssetOperationTokenRequest {
	s.LoginAttribute = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetOperationMode(v string) *GenerateAssetOperationTokenRequest {
	s.OperationMode = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetOperationNote(v string) *GenerateAssetOperationTokenRequest {
	s.OperationNote = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetRegionId(v string) *GenerateAssetOperationTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) SetSsoClient(v string) *GenerateAssetOperationTokenRequest {
	s.SsoClient = &v
	return s
}

func (s *GenerateAssetOperationTokenRequest) Validate() error {
	return dara.Validate(s)
}
