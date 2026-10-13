// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppInstanceRequest
	GetAppName() *string
	SetAppType(v string) *CreateAppInstanceRequest
	GetAppType() *string
	SetClientToken(v string) *CreateAppInstanceRequest
	GetClientToken() *string
	SetComponents(v []*CreateAppInstanceRequestComponents) *CreateAppInstanceRequest
	GetComponents() []*CreateAppInstanceRequestComponents
	SetDBInstanceConfig(v *CreateAppInstanceRequestDBInstanceConfig) *CreateAppInstanceRequest
	GetDBInstanceConfig() *CreateAppInstanceRequestDBInstanceConfig
	SetDBInstanceName(v string) *CreateAppInstanceRequest
	GetDBInstanceName() *string
	SetDashboardPassword(v string) *CreateAppInstanceRequest
	GetDashboardPassword() *string
	SetDashboardUsername(v string) *CreateAppInstanceRequest
	GetDashboardUsername() *string
	SetDatabase(v string) *CreateAppInstanceRequest
	GetDatabase() *string
	SetDatabasePassword(v string) *CreateAppInstanceRequest
	GetDatabasePassword() *string
	SetInitializeWithExistingData(v bool) *CreateAppInstanceRequest
	GetInitializeWithExistingData() *bool
	SetInstanceClass(v string) *CreateAppInstanceRequest
	GetInstanceClass() *string
	SetPublicEndpointEnabled(v bool) *CreateAppInstanceRequest
	GetPublicEndpointEnabled() *bool
	SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceRequest
	GetPublicNetworkAccessEnabled() *bool
	SetRAGEnabled(v bool) *CreateAppInstanceRequest
	GetRAGEnabled() *bool
	SetRegionId(v string) *CreateAppInstanceRequest
	GetRegionId() *string
	SetVSwitchId(v string) *CreateAppInstanceRequest
	GetVSwitchId() *string
}

type CreateAppInstanceRequest struct {
	// The name of the AI application to create.
	//
	// example:
	//
	// test-supabase
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Currently, only **supabase*	- is supported.
	//
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The client generates the value of this parameter to prevent duplicate requests from being submitted.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of components.
	Components []*CreateAppInstanceRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The database instance configuration.
	DBInstanceConfig *CreateAppInstanceRequestDBInstanceConfig `json:"DBInstanceConfig,omitempty" xml:"DBInstanceConfig,omitempty" type:"Struct"`
	// The instance ID of the RDS PostgreSQL database instance to which the AI application connects.
	//
	// <notice>Only **newly purchased empty RDS PostgreSQL instances*	- are supported. The major engine version must be **17**, and the minor engine version must be **20250630 or later**.
	//
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The password for the Supabase Dashboard.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// example:
	//
	// test_Password
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The username for the Supabase Dashboard.
	//
	// example:
	//
	// supabase
	DashboardUsername *string `json:"DashboardUsername,omitempty" xml:"DashboardUsername,omitempty"`
	// The database name.
	//
	// example:
	//
	// test_database_01
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The password for accessing the RDS database.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// example:
	//
	// test_Password
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// Specifies whether to restore from existing PostgreSQL data. Valid values:
	//
	// - true: restores from existing data.
	//
	// - false (default): does not restore from existing data.
	//
	// example:
	//
	// false
	InitializeWithExistingData *bool `json:"InitializeWithExistingData,omitempty" xml:"InitializeWithExistingData,omitempty"`
	// The instance type. Currently, only **rdsai.supabase.basic*	- is supported.
	//
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Specifies whether to enable the public network connection endpoint. Valid values:
	//
	// - true (default): Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	PublicEndpointEnabled *bool `json:"PublicEndpointEnabled,omitempty" xml:"PublicEndpointEnabled,omitempty"`
	// Specifies whether to enable the Internet NAT gateway. Valid values:
	//
	// - **true**: enables the Internet NAT gateway.
	//
	// - **false*	- (default): does not enable the Internet NAT gateway.
	//
	// 	Notice: If an Internet NAT gateway is already enabled for the vSwitch specified by the **VSwitchId*	- parameter settings, set this parameter to **false**.
	//
	// example:
	//
	// false
	PublicNetworkAccessEnabled *bool `json:"PublicNetworkAccessEnabled,omitempty" xml:"PublicNetworkAccessEnabled,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// false
	RAGEnabled *bool `json:"RAGEnabled,omitempty" xml:"RAGEnabled,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-9dp2hkpm22gxscfgy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppInstanceRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppInstanceRequest) GetComponents() []*CreateAppInstanceRequestComponents {
	return s.Components
}

func (s *CreateAppInstanceRequest) GetDBInstanceConfig() *CreateAppInstanceRequestDBInstanceConfig {
	return s.DBInstanceConfig
}

func (s *CreateAppInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateAppInstanceRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *CreateAppInstanceRequest) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *CreateAppInstanceRequest) GetDatabase() *string {
	return s.Database
}

func (s *CreateAppInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *CreateAppInstanceRequest) GetInitializeWithExistingData() *bool {
	return s.InitializeWithExistingData
}

func (s *CreateAppInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateAppInstanceRequest) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *CreateAppInstanceRequest) GetPublicNetworkAccessEnabled() *bool {
	return s.PublicNetworkAccessEnabled
}

func (s *CreateAppInstanceRequest) GetRAGEnabled() *bool {
	return s.RAGEnabled
}

func (s *CreateAppInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAppInstanceRequest) SetAppName(v string) *CreateAppInstanceRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppInstanceRequest) SetAppType(v string) *CreateAppInstanceRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppInstanceRequest) SetClientToken(v string) *CreateAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppInstanceRequest) SetComponents(v []*CreateAppInstanceRequestComponents) *CreateAppInstanceRequest {
	s.Components = v
	return s
}

func (s *CreateAppInstanceRequest) SetDBInstanceConfig(v *CreateAppInstanceRequestDBInstanceConfig) *CreateAppInstanceRequest {
	s.DBInstanceConfig = v
	return s
}

func (s *CreateAppInstanceRequest) SetDBInstanceName(v string) *CreateAppInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDashboardPassword(v string) *CreateAppInstanceRequest {
	s.DashboardPassword = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDashboardUsername(v string) *CreateAppInstanceRequest {
	s.DashboardUsername = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDatabase(v string) *CreateAppInstanceRequest {
	s.Database = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDatabasePassword(v string) *CreateAppInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *CreateAppInstanceRequest) SetInitializeWithExistingData(v bool) *CreateAppInstanceRequest {
	s.InitializeWithExistingData = &v
	return s
}

func (s *CreateAppInstanceRequest) SetInstanceClass(v string) *CreateAppInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateAppInstanceRequest) SetPublicEndpointEnabled(v bool) *CreateAppInstanceRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *CreateAppInstanceRequest) SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceRequest {
	s.PublicNetworkAccessEnabled = &v
	return s
}

func (s *CreateAppInstanceRequest) SetRAGEnabled(v bool) *CreateAppInstanceRequest {
	s.RAGEnabled = &v
	return s
}

func (s *CreateAppInstanceRequest) SetRegionId(v string) *CreateAppInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppInstanceRequest) SetVSwitchId(v string) *CreateAppInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAppInstanceRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DBInstanceConfig != nil {
		if err := s.DBInstanceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppInstanceRequestComponents struct {
	// The component type.
	//
	// example:
	//
	// supabase
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceRequestComponents) GetType() *string {
	return s.Type
}

func (s *CreateAppInstanceRequestComponents) SetType(v string) *CreateAppInstanceRequestComponents {
	s.Type = &v
	return s
}

func (s *CreateAppInstanceRequestComponents) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceRequestDBInstanceConfig struct {
	// The instance type of the database instance.
	//
	// example:
	//
	// pg.n2.2c.1m
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance storage capacity of the database instance.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The billing method of the database instance.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s CreateAppInstanceRequestDBInstanceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceRequestDBInstanceConfig) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceRequestDBInstanceConfig) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateAppInstanceRequestDBInstanceConfig) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateAppInstanceRequestDBInstanceConfig) GetPayType() *string {
	return s.PayType
}

func (s *CreateAppInstanceRequestDBInstanceConfig) SetDBInstanceClass(v string) *CreateAppInstanceRequestDBInstanceConfig {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateAppInstanceRequestDBInstanceConfig) SetDBInstanceStorage(v int32) *CreateAppInstanceRequestDBInstanceConfig {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateAppInstanceRequestDBInstanceConfig) SetPayType(v string) *CreateAppInstanceRequestDBInstanceConfig {
	s.PayType = &v
	return s
}

func (s *CreateAppInstanceRequestDBInstanceConfig) Validate() error {
	return dara.Validate(s)
}
