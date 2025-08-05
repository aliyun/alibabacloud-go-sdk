// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertSetting(v string) *CreateHanaInstanceRequest
	GetAlertSetting() *string
	SetCrossAccountRoleName(v string) *CreateHanaInstanceRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateHanaInstanceRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateHanaInstanceRequest
	GetCrossAccountUserId() *int64
	SetEcsInstanceId(v string) *CreateHanaInstanceRequest
	GetEcsInstanceId() *string
	SetHanaName(v string) *CreateHanaInstanceRequest
	GetHanaName() *string
	SetHost(v string) *CreateHanaInstanceRequest
	GetHost() *string
	SetInstanceNumber(v int32) *CreateHanaInstanceRequest
	GetInstanceNumber() *int32
	SetPassword(v string) *CreateHanaInstanceRequest
	GetPassword() *string
	SetResourceGroupId(v string) *CreateHanaInstanceRequest
	GetResourceGroupId() *string
	SetSid(v string) *CreateHanaInstanceRequest
	GetSid() *string
	SetUseSsl(v bool) *CreateHanaInstanceRequest
	GetUseSsl() *bool
	SetUserName(v string) *CreateHanaInstanceRequest
	GetUserName() *string
	SetValidateCertificate(v bool) *CreateHanaInstanceRequest
	GetValidateCertificate() *bool
	SetVaultId(v string) *CreateHanaInstanceRequest
	GetVaultId() *string
}

type CreateHanaInstanceRequest struct {
	// The alert settings. Valid value: INHERITED, which indicates that the Cloud Backup client sends alert notifications by using the same method configured for the backup vault.
	//
	// example:
	//
	// INHERITED
	AlertSetting *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The backup type. Valid values:
	//
	// - **SELF_ACCOUNT**: Data is backed up within the same Alibaba Cloud account.
	//
	// - **CROSS_ACCOUNT**: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The IDs of the ECS instances that host the SAP HANA instance to be registered. Cloud Backup installs backup clients on the specified ECS instances.
	//
	// example:
	//
	// [\\"i-uf6ir9y******hvisj\\"]
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the SAP HANA instance.
	//
	// example:
	//
	// HANA-DEV
	HanaName *string `json:"HanaName,omitempty" xml:"HanaName,omitempty"`
	// The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
	//
	// example:
	//
	// 47.100.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The instance number of the SAP HANA system.
	//
	// example:
	//
	// 00
	InstanceNumber *int32 `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	// The password that is used to connect with the SAP HANA database.
	//
	// example:
	//
	// ************
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4ebtpkzx7zy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security identifier (SID) of the SAP HANA database. For more information, see [How to find sid user and instance number of HANA db?](https://answers.sap.com/questions/555192/how-to-find-sid-user-and-instance-number-of-hana-d.html?spm=a2c4g.11186623.0.0.55c34b4ftZeXNK).
	//
	// example:
	//
	// HXE
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
	//
	// example:
	//
	// true
	UseSsl *bool `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The username of the SYSTEMDB database.
	//
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Specifies whether to verify the SSL certificate of the SAP HANA database.
	//
	// example:
	//
	// false
	ValidateCertificate *bool `json:"ValidateCertificate,omitempty" xml:"ValidateCertificate,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0003v4a******gfv2
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateHanaInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateHanaInstanceRequest) GetAlertSetting() *string {
	return s.AlertSetting
}

func (s *CreateHanaInstanceRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateHanaInstanceRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateHanaInstanceRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateHanaInstanceRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *CreateHanaInstanceRequest) GetHanaName() *string {
	return s.HanaName
}

func (s *CreateHanaInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *CreateHanaInstanceRequest) GetInstanceNumber() *int32 {
	return s.InstanceNumber
}

func (s *CreateHanaInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateHanaInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateHanaInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *CreateHanaInstanceRequest) GetUseSsl() *bool {
	return s.UseSsl
}

func (s *CreateHanaInstanceRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateHanaInstanceRequest) GetValidateCertificate() *bool {
	return s.ValidateCertificate
}

func (s *CreateHanaInstanceRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateHanaInstanceRequest) SetAlertSetting(v string) *CreateHanaInstanceRequest {
	s.AlertSetting = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetCrossAccountRoleName(v string) *CreateHanaInstanceRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetCrossAccountType(v string) *CreateHanaInstanceRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetCrossAccountUserId(v int64) *CreateHanaInstanceRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetEcsInstanceId(v string) *CreateHanaInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetHanaName(v string) *CreateHanaInstanceRequest {
	s.HanaName = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetHost(v string) *CreateHanaInstanceRequest {
	s.Host = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetInstanceNumber(v int32) *CreateHanaInstanceRequest {
	s.InstanceNumber = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetPassword(v string) *CreateHanaInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetResourceGroupId(v string) *CreateHanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetSid(v string) *CreateHanaInstanceRequest {
	s.Sid = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetUseSsl(v bool) *CreateHanaInstanceRequest {
	s.UseSsl = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetUserName(v string) *CreateHanaInstanceRequest {
	s.UserName = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetValidateCertificate(v bool) *CreateHanaInstanceRequest {
	s.ValidateCertificate = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetVaultId(v string) *CreateHanaInstanceRequest {
	s.VaultId = &v
	return s
}

func (s *CreateHanaInstanceRequest) Validate() error {
	return dara.Validate(s)
}
