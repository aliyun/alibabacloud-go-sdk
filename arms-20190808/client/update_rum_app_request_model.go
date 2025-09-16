// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRumAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppConfig(v string) *UpdateRumAppRequest
	GetAppConfig() *string
	SetAutoRestart(v bool) *UpdateRumAppRequest
	GetAutoRestart() *bool
	SetBackendServiceTraceRegion(v string) *UpdateRumAppRequest
	GetBackendServiceTraceRegion() *string
	SetBonreeSDKConfigJson(v string) *UpdateRumAppRequest
	GetBonreeSDKConfigJson() *string
	SetDescription(v string) *UpdateRumAppRequest
	GetDescription() *string
	SetIsSubscribe(v bool) *UpdateRumAppRequest
	GetIsSubscribe() *bool
	SetNickname(v string) *UpdateRumAppRequest
	GetNickname() *string
	SetPid(v string) *UpdateRumAppRequest
	GetPid() *string
	SetRealRegionId(v string) *UpdateRumAppRequest
	GetRealRegionId() *string
	SetRegionId(v string) *UpdateRumAppRequest
	GetRegionId() *string
	SetRestart(v bool) *UpdateRumAppRequest
	GetRestart() *bool
	SetServiceDomainOperationJson(v string) *UpdateRumAppRequest
	GetServiceDomainOperationJson() *string
	SetStop(v bool) *UpdateRumAppRequest
	GetStop() *bool
	SetWebSDKConfigJson(v string) *UpdateRumAppRequest
	GetWebSDKConfigJson() *string
}

type UpdateRumAppRequest struct {
	// The application configurations in the JSON format. This parameter is deprecated.
	//
	// example:
	//
	// {"apiRequestOfH5":300,"apiRequestOfOriginal":500,"coldStart":5000,"hotStart":3000,"staticResourceLoad":300,"stutter":1000,"viewLoadOfH5":1000,"viewLoadOfOriginal":2000}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// Specifies whether to restart the application the next day. Valid values: true and false.
	//
	// example:
	//
	// true
	AutoRestart *bool `json:"AutoRestart,omitempty" xml:"AutoRestart,omitempty"`
	// The region where the backend application is deployed. This parameter is used in end-to-end tracing scenarios.
	//
	// example:
	//
	// {\\"moduleConfig\\":{\\"enable\\":true,\\"defaultConfig\\":{\\"network\\":{\\"enable\\":true},\\"h5\\":{\\"enable\\":true},\\"routechange\\":{\\"enable\\":true},\\"crash\\":{\\"enable\\":true},\\"view\\":{\\"enable\\":true},\\"coollaunch\\":{\\"enable\\":true},\\"hotlaunch\\":{\\"enable\\":true},\\"action\\":{\\"enable\\":true},\\"lagstuck\\":{\\"enable\\":true},\\"lagfps\\":{\\"enable\\":true},\\"statechange\\":{\\"enable\\":true},\\"anr\\":{\\"enable\\":true},\\"customlog\\":{\\"enable\\":true},\\"customevent\\":{\\"enable\\":true},\\"custommetric\\":{\\"enable\\":true}},\\"versionConfigs\\":{\\"1.1.0\\":{\\"useCustom\\":true,\\"customConfig\\":{\\"network\\":{\\"enable\\":true},\\"h5\\":{\\"enable\\":true},\\"routechange\\":{\\"enable\\":true},\\"crash\\":{\\"enable\\":true},\\"view\\":{\\"enable\\":true},\\"coollaunch\\":{\\"enable\\":true},\\"hotlaunch\\":{\\"enable\\":true},\\"action\\":{\\"enable\\":true},\\"lagstuck\\":{\\"enable\\":false},\\"lagfps\\":{\\"enable\\":false},\\"statechange\\":{\\"enable\\":true},\\"anr\\":{\\"enable\\":true},\\"customlog\\":{\\"enable\\":true},\\"customevent\\":{\\"enable\\":true},\\"custommetric\\":{\\"enable\\":true}}},\\"1.2.0\\":{\\"useCustom\\":false,\\"customConfig\\":{}}}}}
	BackendServiceTraceRegion *string `json:"BackendServiceTraceRegion,omitempty" xml:"BackendServiceTraceRegion,omitempty"`
	// The collection configurations of the mobile SDK. You can enable or disable collection configurations based on the app version.
	//
	// example:
	//
	// {\\"moduleConfig\\":{\\"enable\\":true,\\"defaultConfig\\":{\\"network\\":{\\"enable\\":true},\\"h5\\":{\\"enable\\":true},\\"routechange\\":{\\"enable\\":true},\\"crash\\":{\\"enable\\":true},\\"view\\":{\\"enable\\":true},\\"coollaunch\\":{\\"enable\\":true},\\"hotlaunch\\":{\\"enable\\":true},\\"action\\":{\\"enable\\":true},\\"lagstuck\\":{\\"enable\\":true},\\"lagfps\\":{\\"enable\\":true},\\"statechange\\":{\\"enable\\":true},\\"anr\\":{\\"enable\\":true},\\"customlog\\":{\\"enable\\":true},\\"customevent\\":{\\"enable\\":true},\\"custommetric\\":{\\"enable\\":true}},\\"versionConfigs\\":{\\"1.1.0\\":{\\"useCustom\\":true,\\"customConfig\\":{\\"network\\":{\\"enable\\":true},\\"h5\\":{\\"enable\\":true},\\"routechange\\":{\\"enable\\":true},\\"crash\\":{\\"enable\\":true},\\"view\\":{\\"enable\\":true},\\"coollaunch\\":{\\"enable\\":true},\\"hotlaunch\\":{\\"enable\\":true},\\"action\\":{\\"enable\\":true},\\"lagstuck\\":{\\"enable\\":false},\\"lagfps\\":{\\"enable\\":false},\\"statechange\\":{\\"enable\\":true},\\"anr\\":{\\"enable\\":true},\\"customlog\\":{\\"enable\\":true},\\"customevent\\":{\\"enable\\":true},\\"custommetric\\":{\\"enable\\":true}}},\\"1.2.0\\":{\\"useCustom\\":false,\\"customConfig\\":{}}}}}
	BonreeSDKConfigJson *string `json:"BonreeSDKConfigJson,omitempty" xml:"BonreeSDKConfigJson,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether you want to subscribe to the application. Valid values: true and false.
	//
	// example:
	//
	// true
	IsSubscribe *bool `json:"IsSubscribe,omitempty" xml:"IsSubscribe,omitempty"`
	// The alias of the application.
	//
	// example:
	//
	// Android Test
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5xxxxs@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region where the application resides. You can leave this parameter empty or set it to China East 2 Finance.
	//
	// example:
	//
	// cn-hangzhou
	RealRegionId *string `json:"RealRegionId,omitempty" xml:"RealRegionId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to restart the application. Valid values: true and false.
	//
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// The service domain name of the application. You can create, modify, and delete service domain name configurations.
	//
	// example:
	//
	// {\\"Op\\":\\"Update\\",\\"Domain\\":\\"example.com\\",\\"Config\\":{\\"Description\\":\\"message\\",\\"Tracing\\":\\"true\\",\\"PropagatorTypes\\":[\\"sw8\\"]}}
	ServiceDomainOperationJson *string `json:"ServiceDomainOperationJson,omitempty" xml:"ServiceDomainOperationJson,omitempty"`
	// Specifies whether to stop the application. Valid values: true and false.
	//
	// example:
	//
	// true
	Stop             *bool   `json:"Stop,omitempty" xml:"Stop,omitempty"`
	WebSDKConfigJson *string `json:"WebSDKConfigJson,omitempty" xml:"WebSDKConfigJson,omitempty"`
}

func (s UpdateRumAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateRumAppRequest) GetAppConfig() *string {
	return s.AppConfig
}

func (s *UpdateRumAppRequest) GetAutoRestart() *bool {
	return s.AutoRestart
}

func (s *UpdateRumAppRequest) GetBackendServiceTraceRegion() *string {
	return s.BackendServiceTraceRegion
}

func (s *UpdateRumAppRequest) GetBonreeSDKConfigJson() *string {
	return s.BonreeSDKConfigJson
}

func (s *UpdateRumAppRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRumAppRequest) GetIsSubscribe() *bool {
	return s.IsSubscribe
}

func (s *UpdateRumAppRequest) GetNickname() *string {
	return s.Nickname
}

func (s *UpdateRumAppRequest) GetPid() *string {
	return s.Pid
}

func (s *UpdateRumAppRequest) GetRealRegionId() *string {
	return s.RealRegionId
}

func (s *UpdateRumAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateRumAppRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdateRumAppRequest) GetServiceDomainOperationJson() *string {
	return s.ServiceDomainOperationJson
}

func (s *UpdateRumAppRequest) GetStop() *bool {
	return s.Stop
}

func (s *UpdateRumAppRequest) GetWebSDKConfigJson() *string {
	return s.WebSDKConfigJson
}

func (s *UpdateRumAppRequest) SetAppConfig(v string) *UpdateRumAppRequest {
	s.AppConfig = &v
	return s
}

func (s *UpdateRumAppRequest) SetAutoRestart(v bool) *UpdateRumAppRequest {
	s.AutoRestart = &v
	return s
}

func (s *UpdateRumAppRequest) SetBackendServiceTraceRegion(v string) *UpdateRumAppRequest {
	s.BackendServiceTraceRegion = &v
	return s
}

func (s *UpdateRumAppRequest) SetBonreeSDKConfigJson(v string) *UpdateRumAppRequest {
	s.BonreeSDKConfigJson = &v
	return s
}

func (s *UpdateRumAppRequest) SetDescription(v string) *UpdateRumAppRequest {
	s.Description = &v
	return s
}

func (s *UpdateRumAppRequest) SetIsSubscribe(v bool) *UpdateRumAppRequest {
	s.IsSubscribe = &v
	return s
}

func (s *UpdateRumAppRequest) SetNickname(v string) *UpdateRumAppRequest {
	s.Nickname = &v
	return s
}

func (s *UpdateRumAppRequest) SetPid(v string) *UpdateRumAppRequest {
	s.Pid = &v
	return s
}

func (s *UpdateRumAppRequest) SetRealRegionId(v string) *UpdateRumAppRequest {
	s.RealRegionId = &v
	return s
}

func (s *UpdateRumAppRequest) SetRegionId(v string) *UpdateRumAppRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateRumAppRequest) SetRestart(v bool) *UpdateRumAppRequest {
	s.Restart = &v
	return s
}

func (s *UpdateRumAppRequest) SetServiceDomainOperationJson(v string) *UpdateRumAppRequest {
	s.ServiceDomainOperationJson = &v
	return s
}

func (s *UpdateRumAppRequest) SetStop(v bool) *UpdateRumAppRequest {
	s.Stop = &v
	return s
}

func (s *UpdateRumAppRequest) SetWebSDKConfigJson(v string) *UpdateRumAppRequest {
	s.WebSDKConfigJson = &v
	return s
}

func (s *UpdateRumAppRequest) Validate() error {
	return dara.Validate(s)
}
