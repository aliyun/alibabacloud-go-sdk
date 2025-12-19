// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppInstanceShrinkRequest
	GetAppName() *string
	SetAppType(v string) *CreateAppInstanceShrinkRequest
	GetAppType() *string
	SetClientToken(v string) *CreateAppInstanceShrinkRequest
	GetClientToken() *string
	SetDBInstanceConfigShrink(v string) *CreateAppInstanceShrinkRequest
	GetDBInstanceConfigShrink() *string
	SetDBInstanceName(v string) *CreateAppInstanceShrinkRequest
	GetDBInstanceName() *string
	SetDashboardPassword(v string) *CreateAppInstanceShrinkRequest
	GetDashboardPassword() *string
	SetDashboardUsername(v string) *CreateAppInstanceShrinkRequest
	GetDashboardUsername() *string
	SetDatabasePassword(v string) *CreateAppInstanceShrinkRequest
	GetDatabasePassword() *string
	SetInitializeWithExistingData(v bool) *CreateAppInstanceShrinkRequest
	GetInitializeWithExistingData() *bool
	SetInstanceClass(v string) *CreateAppInstanceShrinkRequest
	GetInstanceClass() *string
	SetPublicEndpointEnabled(v bool) *CreateAppInstanceShrinkRequest
	GetPublicEndpointEnabled() *bool
	SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceShrinkRequest
	GetPublicNetworkAccessEnabled() *bool
	SetRAGEnabled(v bool) *CreateAppInstanceShrinkRequest
	GetRAGEnabled() *bool
	SetRegionId(v string) *CreateAppInstanceShrinkRequest
	GetRegionId() *string
	SetVSwitchId(v string) *CreateAppInstanceShrinkRequest
	GetVSwitchId() *string
}

type CreateAppInstanceShrinkRequest struct {
	// example:
	//
	// test-supabase
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceConfigShrink *string `json:"DBInstanceConfig,omitempty" xml:"DBInstanceConfig,omitempty"`
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// test_Password
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// example:
	//
	// supabase
	DashboardUsername *string `json:"DashboardUsername,omitempty" xml:"DashboardUsername,omitempty"`
	// example:
	//
	// test_Password
	DatabasePassword           *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	InitializeWithExistingData *bool   `json:"InitializeWithExistingData,omitempty" xml:"InitializeWithExistingData,omitempty"`
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass         *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	PublicEndpointEnabled *bool   `json:"PublicEndpointEnabled,omitempty" xml:"PublicEndpointEnabled,omitempty"`
	// example:
	//
	// false
	PublicNetworkAccessEnabled *bool `json:"PublicNetworkAccessEnabled,omitempty" xml:"PublicNetworkAccessEnabled,omitempty"`
	RAGEnabled                 *bool `json:"RAGEnabled,omitempty" xml:"RAGEnabled,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vsw-9dp2hkpm22gxscfgy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppInstanceShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppInstanceShrinkRequest) GetDBInstanceConfigShrink() *string {
	return s.DBInstanceConfigShrink
}

func (s *CreateAppInstanceShrinkRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateAppInstanceShrinkRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *CreateAppInstanceShrinkRequest) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *CreateAppInstanceShrinkRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *CreateAppInstanceShrinkRequest) GetInitializeWithExistingData() *bool {
	return s.InitializeWithExistingData
}

func (s *CreateAppInstanceShrinkRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateAppInstanceShrinkRequest) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *CreateAppInstanceShrinkRequest) GetPublicNetworkAccessEnabled() *bool {
	return s.PublicNetworkAccessEnabled
}

func (s *CreateAppInstanceShrinkRequest) GetRAGEnabled() *bool {
	return s.RAGEnabled
}

func (s *CreateAppInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAppInstanceShrinkRequest) SetAppName(v string) *CreateAppInstanceShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetAppType(v string) *CreateAppInstanceShrinkRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetClientToken(v string) *CreateAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDBInstanceConfigShrink(v string) *CreateAppInstanceShrinkRequest {
	s.DBInstanceConfigShrink = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDBInstanceName(v string) *CreateAppInstanceShrinkRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDashboardPassword(v string) *CreateAppInstanceShrinkRequest {
	s.DashboardPassword = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDashboardUsername(v string) *CreateAppInstanceShrinkRequest {
	s.DashboardUsername = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDatabasePassword(v string) *CreateAppInstanceShrinkRequest {
	s.DatabasePassword = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetInitializeWithExistingData(v bool) *CreateAppInstanceShrinkRequest {
	s.InitializeWithExistingData = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetInstanceClass(v string) *CreateAppInstanceShrinkRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetPublicEndpointEnabled(v bool) *CreateAppInstanceShrinkRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceShrinkRequest {
	s.PublicNetworkAccessEnabled = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetRAGEnabled(v bool) *CreateAppInstanceShrinkRequest {
	s.RAGEnabled = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetRegionId(v string) *CreateAppInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetVSwitchId(v string) *CreateAppInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
