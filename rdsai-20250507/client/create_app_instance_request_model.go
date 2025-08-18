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
	SetDBInstanceName(v string) *CreateAppInstanceRequest
	GetDBInstanceName() *string
	SetDashboardPassword(v string) *CreateAppInstanceRequest
	GetDashboardPassword() *string
	SetDashboardUsername(v string) *CreateAppInstanceRequest
	GetDashboardUsername() *string
	SetDatabasePassword(v string) *CreateAppInstanceRequest
	GetDatabasePassword() *string
	SetInstanceClass(v string) *CreateAppInstanceRequest
	GetInstanceClass() *string
	SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceRequest
	GetPublicNetworkAccessEnabled() *bool
	SetRegionId(v string) *CreateAppInstanceRequest
	GetRegionId() *string
	SetVSwitchId(v string) *CreateAppInstanceRequest
	GetVSwitchId() *string
}

type CreateAppInstanceRequest struct {
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// false
	PublicNetworkAccessEnabled *bool `json:"PublicNetworkAccessEnabled,omitempty" xml:"PublicNetworkAccessEnabled,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *CreateAppInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateAppInstanceRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *CreateAppInstanceRequest) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *CreateAppInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *CreateAppInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateAppInstanceRequest) GetPublicNetworkAccessEnabled() *bool {
	return s.PublicNetworkAccessEnabled
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

func (s *CreateAppInstanceRequest) SetDatabasePassword(v string) *CreateAppInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *CreateAppInstanceRequest) SetInstanceClass(v string) *CreateAppInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateAppInstanceRequest) SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceRequest {
	s.PublicNetworkAccessEnabled = &v
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
	return dara.Validate(s)
}
