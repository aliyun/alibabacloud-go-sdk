// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertSetting(v string) *UpdateHanaInstanceRequest
	GetAlertSetting() *string
	SetClusterId(v string) *UpdateHanaInstanceRequest
	GetClusterId() *string
	SetHanaName(v string) *UpdateHanaInstanceRequest
	GetHanaName() *string
	SetHost(v string) *UpdateHanaInstanceRequest
	GetHost() *string
	SetInstanceNumber(v int32) *UpdateHanaInstanceRequest
	GetInstanceNumber() *int32
	SetPassword(v string) *UpdateHanaInstanceRequest
	GetPassword() *string
	SetResourceGroupId(v string) *UpdateHanaInstanceRequest
	GetResourceGroupId() *string
	SetUseSsl(v bool) *UpdateHanaInstanceRequest
	GetUseSsl() *bool
	SetUserName(v string) *UpdateHanaInstanceRequest
	GetUserName() *string
	SetValidateCertificate(v bool) *UpdateHanaInstanceRequest
	GetValidateCertificate() *bool
	SetVaultId(v string) *UpdateHanaInstanceRequest
	GetVaultId() *string
}

type UpdateHanaInstanceRequest struct {
	// The alert settings. Valid value: INHERITED, which indicates that the Cloud Backup client sends alert notifications by using the same method configured for the backup vault.
	//
	// example:
	//
	// INHERITED
	AlertSetting *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-000axjt******c6j8
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the SAP HANA instance.
	//
	// example:
	//
	// SAP-HANA-DEV
	HanaName *string `json:"HanaName,omitempty" xml:"HanaName,omitempty"`
	// The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
	//
	// example:
	//
	// 47.100.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The instance number of the SAP HANA system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00
	InstanceNumber *int32 `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	// The password that is used to connect with the SAP HANA database.
	//
	// example:
	//
	// **********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzvx7d3c4kpny
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL). Valid values:
	//
	// 	- true: The SAP HANA database is connected over SSL.
	//
	// 	- false: The SAP HANA database is not connected over SSL.
	//
	// This parameter is required.
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
	// Specifies whether to verify the SSL certificate of the SAP HANA database. Valid values:
	//
	// 	- true: The SSL certificate of the SAP HANA database is verified.
	//
	// 	- false: The SSL certificate of the SAP HANA database is not verified.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	ValidateCertificate *bool `json:"ValidateCertificate,omitempty" xml:"ValidateCertificate,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0003v4ah******9xp
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaInstanceRequest) GetAlertSetting() *string {
	return s.AlertSetting
}

func (s *UpdateHanaInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateHanaInstanceRequest) GetHanaName() *string {
	return s.HanaName
}

func (s *UpdateHanaInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *UpdateHanaInstanceRequest) GetInstanceNumber() *int32 {
	return s.InstanceNumber
}

func (s *UpdateHanaInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateHanaInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateHanaInstanceRequest) GetUseSsl() *bool {
	return s.UseSsl
}

func (s *UpdateHanaInstanceRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateHanaInstanceRequest) GetValidateCertificate() *bool {
	return s.ValidateCertificate
}

func (s *UpdateHanaInstanceRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateHanaInstanceRequest) SetAlertSetting(v string) *UpdateHanaInstanceRequest {
	s.AlertSetting = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetClusterId(v string) *UpdateHanaInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetHanaName(v string) *UpdateHanaInstanceRequest {
	s.HanaName = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetHost(v string) *UpdateHanaInstanceRequest {
	s.Host = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetInstanceNumber(v int32) *UpdateHanaInstanceRequest {
	s.InstanceNumber = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetPassword(v string) *UpdateHanaInstanceRequest {
	s.Password = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetResourceGroupId(v string) *UpdateHanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetUseSsl(v bool) *UpdateHanaInstanceRequest {
	s.UseSsl = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetUserName(v string) *UpdateHanaInstanceRequest {
	s.UserName = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetValidateCertificate(v bool) *UpdateHanaInstanceRequest {
	s.ValidateCertificate = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetVaultId(v string) *UpdateHanaInstanceRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateHanaInstanceRequest) Validate() error {
	return dara.Validate(s)
}
