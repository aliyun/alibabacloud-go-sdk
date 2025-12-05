// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavePtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v *SavePtsSceneRequestScene) *SavePtsSceneRequest
	GetScene() *SavePtsSceneRequestScene
}

type SavePtsSceneRequest struct {
	// The information about the scenario.
	//
	// This parameter is required.
	Scene *SavePtsSceneRequestScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Struct"`
}

func (s SavePtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequest) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequest) GetScene() *SavePtsSceneRequestScene {
	return s.Scene
}

func (s *SavePtsSceneRequest) SetScene(v *SavePtsSceneRequestScene) *SavePtsSceneRequest {
	s.Scene = v
	return s
}

func (s *SavePtsSceneRequest) Validate() error {
	if s.Scene != nil {
		if err := s.Scene.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SavePtsSceneRequestScene struct {
	// The advanced settings.
	//
	// if can be null:
	// true
	AdvanceSetting *SavePtsSceneRequestSceneAdvanceSetting `json:"AdvanceSetting,omitempty" xml:"AdvanceSetting,omitempty" type:"Struct"`
	// The file parameters.
	//
	// if can be null:
	// true
	FileParameterList []*SavePtsSceneRequestSceneFileParameterList `json:"FileParameterList,omitempty" xml:"FileParameterList,omitempty" type:"Repeated"`
	// The global customization parameters.
	//
	// if can be null:
	// true
	GlobalParameterList []*SavePtsSceneRequestSceneGlobalParameterList `json:"GlobalParameterList,omitempty" xml:"GlobalParameterList,omitempty" type:"Repeated"`
	// The load settings.
	//
	// This parameter is required.
	LoadConfig *SavePtsSceneRequestSceneLoadConfig `json:"LoadConfig,omitempty" xml:"LoadConfig,omitempty" type:"Struct"`
	// The sessions.
	//
	// This parameter is required.
	RelationList []*SavePtsSceneRequestSceneRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	// The ID of the scenario. To save a new scenario, leave this parameter empty. To update an existing scenario, specify the ID of the scenario.
	//
	// example:
	//
	// IUYAHGJ
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The name of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s SavePtsSceneRequestScene) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestScene) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestScene) GetAdvanceSetting() *SavePtsSceneRequestSceneAdvanceSetting {
	return s.AdvanceSetting
}

func (s *SavePtsSceneRequestScene) GetFileParameterList() []*SavePtsSceneRequestSceneFileParameterList {
	return s.FileParameterList
}

func (s *SavePtsSceneRequestScene) GetGlobalParameterList() []*SavePtsSceneRequestSceneGlobalParameterList {
	return s.GlobalParameterList
}

func (s *SavePtsSceneRequestScene) GetLoadConfig() *SavePtsSceneRequestSceneLoadConfig {
	return s.LoadConfig
}

func (s *SavePtsSceneRequestScene) GetRelationList() []*SavePtsSceneRequestSceneRelationList {
	return s.RelationList
}

func (s *SavePtsSceneRequestScene) GetSceneId() *string {
	return s.SceneId
}

func (s *SavePtsSceneRequestScene) GetSceneName() *string {
	return s.SceneName
}

func (s *SavePtsSceneRequestScene) SetAdvanceSetting(v *SavePtsSceneRequestSceneAdvanceSetting) *SavePtsSceneRequestScene {
	s.AdvanceSetting = v
	return s
}

func (s *SavePtsSceneRequestScene) SetFileParameterList(v []*SavePtsSceneRequestSceneFileParameterList) *SavePtsSceneRequestScene {
	s.FileParameterList = v
	return s
}

func (s *SavePtsSceneRequestScene) SetGlobalParameterList(v []*SavePtsSceneRequestSceneGlobalParameterList) *SavePtsSceneRequestScene {
	s.GlobalParameterList = v
	return s
}

func (s *SavePtsSceneRequestScene) SetLoadConfig(v *SavePtsSceneRequestSceneLoadConfig) *SavePtsSceneRequestScene {
	s.LoadConfig = v
	return s
}

func (s *SavePtsSceneRequestScene) SetRelationList(v []*SavePtsSceneRequestSceneRelationList) *SavePtsSceneRequestScene {
	s.RelationList = v
	return s
}

func (s *SavePtsSceneRequestScene) SetSceneId(v string) *SavePtsSceneRequestScene {
	s.SceneId = &v
	return s
}

func (s *SavePtsSceneRequestScene) SetSceneName(v string) *SavePtsSceneRequestScene {
	s.SceneName = &v
	return s
}

func (s *SavePtsSceneRequestScene) Validate() error {
	if s.AdvanceSetting != nil {
		if err := s.AdvanceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.FileParameterList != nil {
		for _, item := range s.FileParameterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GlobalParameterList != nil {
		for _, item := range s.GlobalParameterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LoadConfig != nil {
		if err := s.LoadConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SavePtsSceneRequestSceneAdvanceSetting struct {
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 5
	ConnectionTimeoutInSecond *int32 `json:"ConnectionTimeoutInSecond,omitempty" xml:"ConnectionTimeoutInSecond,omitempty"`
	// The domain name-IP address binding relationships
	DomainBindingList []*SavePtsSceneRequestSceneAdvanceSettingDomainBindingList `json:"DomainBindingList,omitempty" xml:"DomainBindingList,omitempty" type:"Repeated"`
	// The log sampling rate. Valid values: 1, 10, 20, 30, 40, and 50.
	//
	// example:
	//
	// 1
	LogRate *int32 `json:"LogRate,omitempty" xml:"LogRate,omitempty"`
	// The success status code. Separate multiple status codes with commas (,).
	//
	// example:
	//
	// 205
	SuccessCode *string `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty"`
}

func (s SavePtsSceneRequestSceneAdvanceSetting) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneAdvanceSetting) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) GetConnectionTimeoutInSecond() *int32 {
	return s.ConnectionTimeoutInSecond
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) GetDomainBindingList() []*SavePtsSceneRequestSceneAdvanceSettingDomainBindingList {
	return s.DomainBindingList
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) GetLogRate() *int32 {
	return s.LogRate
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) GetSuccessCode() *string {
	return s.SuccessCode
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) SetConnectionTimeoutInSecond(v int32) *SavePtsSceneRequestSceneAdvanceSetting {
	s.ConnectionTimeoutInSecond = &v
	return s
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) SetDomainBindingList(v []*SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) *SavePtsSceneRequestSceneAdvanceSetting {
	s.DomainBindingList = v
	return s
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) SetLogRate(v int32) *SavePtsSceneRequestSceneAdvanceSetting {
	s.LogRate = &v
	return s
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) SetSuccessCode(v string) *SavePtsSceneRequestSceneAdvanceSetting {
	s.SuccessCode = &v
	return s
}

func (s *SavePtsSceneRequestSceneAdvanceSetting) Validate() error {
	if s.DomainBindingList != nil {
		for _, item := range s.DomainBindingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SavePtsSceneRequestSceneAdvanceSettingDomainBindingList struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP addresses.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) GetDomain() *string {
	return s.Domain
}

func (s *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) GetIps() []*string {
	return s.Ips
}

func (s *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) SetDomain(v string) *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList {
	s.Domain = &v
	return s
}

func (s *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) SetIps(v []*string) *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList {
	s.Ips = v
	return s
}

func (s *SavePtsSceneRequestSceneAdvanceSettingDomainBindingList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneFileParameterList struct {
	// The name of the file.
	//
	// example:
	//
	// test.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS URL of the file, which must be accessible over the Internet.
	//
	// example:
	//
	// https://jmeter-pts-testing-version.oss-cn-shanghai.aliyuncs.com/param-file.csv
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s SavePtsSceneRequestSceneFileParameterList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneFileParameterList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneFileParameterList) GetFileName() *string {
	return s.FileName
}

func (s *SavePtsSceneRequestSceneFileParameterList) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *SavePtsSceneRequestSceneFileParameterList) SetFileName(v string) *SavePtsSceneRequestSceneFileParameterList {
	s.FileName = &v
	return s
}

func (s *SavePtsSceneRequestSceneFileParameterList) SetFileOssAddress(v string) *SavePtsSceneRequestSceneFileParameterList {
	s.FileOssAddress = &v
	return s
}

func (s *SavePtsSceneRequestSceneFileParameterList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneGlobalParameterList struct {
	// The name of the parameter.
	//
	// example:
	//
	// global
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 11111
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s SavePtsSceneRequestSceneGlobalParameterList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneGlobalParameterList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneGlobalParameterList) GetParamName() *string {
	return s.ParamName
}

func (s *SavePtsSceneRequestSceneGlobalParameterList) GetParamValue() *string {
	return s.ParamValue
}

func (s *SavePtsSceneRequestSceneGlobalParameterList) SetParamName(v string) *SavePtsSceneRequestSceneGlobalParameterList {
	s.ParamName = &v
	return s
}

func (s *SavePtsSceneRequestSceneGlobalParameterList) SetParamValue(v string) *SavePtsSceneRequestSceneGlobalParameterList {
	s.ParamValue = &v
	return s
}

func (s *SavePtsSceneRequestSceneGlobalParameterList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneLoadConfig struct {
	// The number of load generators. If the number of concurrent virtual users exceeds 250 or the RPS exceeds 2,000, you can use multiple load generators. The maximum number of load generators is limited to the total number of concurrent virtual users divided by 250 or the total RPS divided by 200.
	//
	// example:
	//
	// 1
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The API request load settings.
	ApiLoadConfigList []*SavePtsSceneRequestSceneLoadConfigApiLoadConfigList `json:"ApiLoadConfigList,omitempty" xml:"ApiLoadConfigList,omitempty" type:"Repeated"`
	// Specifies whether the load is automatically incremented. This parameter takes effect only if you set `TestMode=concurrency_mode`.
	//
	// example:
	//
	// true
	AutoStep *bool `json:"AutoStep,omitempty" xml:"AutoStep,omitempty"`
	// The load level settings of the scenario.
	//
	// This parameter is required.
	Configuration *SavePtsSceneRequestSceneLoadConfigConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// The increment percentage. Valid values: 10 to 100, in increments of 10.
	//
	// This parameter takes effect only if you set `testMode=concurrency_mode`and `autoStep=true`.
	//
	// example:
	//
	// 30
	Increment *int32 `json:"Increment,omitempty" xml:"Increment,omitempty"`
	// The duration of a specific load level. Unit: minutes. The value must be less than the value of **maxRunningTime**.
	//
	// example:
	//
	// 3
	KeepTime *int32 `json:"KeepTime,omitempty" xml:"KeepTime,omitempty"`
	// The duration of load application. Unit: minutes. Valid values: 1 to 1440.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxRunningTime *int32 `json:"MaxRunningTime,omitempty" xml:"MaxRunningTime,omitempty"`
	// The session settings.
	RelationLoadConfigList []*SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList `json:"RelationLoadConfigList,omitempty" xml:"RelationLoadConfigList,omitempty" type:"Repeated"`
	// The load application mode. Valid values:
	//
	// 	- concurrency_mode: concurrency mode
	//
	// 	- tps_mode: RPS mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// concurrency_mode
	TestMode *string `json:"TestMode,omitempty" xml:"TestMode,omitempty"`
	// The VPC settings.
	//
	// if can be null:
	// true
	VpcLoadConfig *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig `json:"VpcLoadConfig,omitempty" xml:"VpcLoadConfig,omitempty" type:"Struct"`
}

func (s SavePtsSceneRequestSceneLoadConfig) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneLoadConfig) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetApiLoadConfigList() []*SavePtsSceneRequestSceneLoadConfigApiLoadConfigList {
	return s.ApiLoadConfigList
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetAutoStep() *bool {
	return s.AutoStep
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetConfiguration() *SavePtsSceneRequestSceneLoadConfigConfiguration {
	return s.Configuration
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetIncrement() *int32 {
	return s.Increment
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetKeepTime() *int32 {
	return s.KeepTime
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetMaxRunningTime() *int32 {
	return s.MaxRunningTime
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetRelationLoadConfigList() []*SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList {
	return s.RelationLoadConfigList
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetTestMode() *string {
	return s.TestMode
}

func (s *SavePtsSceneRequestSceneLoadConfig) GetVpcLoadConfig() *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig {
	return s.VpcLoadConfig
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetAgentCount(v int32) *SavePtsSceneRequestSceneLoadConfig {
	s.AgentCount = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetApiLoadConfigList(v []*SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) *SavePtsSceneRequestSceneLoadConfig {
	s.ApiLoadConfigList = v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetAutoStep(v bool) *SavePtsSceneRequestSceneLoadConfig {
	s.AutoStep = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetConfiguration(v *SavePtsSceneRequestSceneLoadConfigConfiguration) *SavePtsSceneRequestSceneLoadConfig {
	s.Configuration = v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetIncrement(v int32) *SavePtsSceneRequestSceneLoadConfig {
	s.Increment = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetKeepTime(v int32) *SavePtsSceneRequestSceneLoadConfig {
	s.KeepTime = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetMaxRunningTime(v int32) *SavePtsSceneRequestSceneLoadConfig {
	s.MaxRunningTime = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetRelationLoadConfigList(v []*SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) *SavePtsSceneRequestSceneLoadConfig {
	s.RelationLoadConfigList = v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetTestMode(v string) *SavePtsSceneRequestSceneLoadConfig {
	s.TestMode = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) SetVpcLoadConfig(v *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) *SavePtsSceneRequestSceneLoadConfig {
	s.VpcLoadConfig = v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfig) Validate() error {
	if s.ApiLoadConfigList != nil {
		for _, item := range s.ApiLoadConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.RelationLoadConfigList != nil {
		for _, item := range s.RelationLoadConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpcLoadConfig != nil {
		if err := s.VpcLoadConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SavePtsSceneRequestSceneLoadConfigApiLoadConfigList struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The starting RPS.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RpsBegin *int32 `json:"RpsBegin,omitempty" xml:"RpsBegin,omitempty"`
	// The maximum RPS.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RpsLimit *int32 `json:"RpsLimit,omitempty" xml:"RpsLimit,omitempty"`
}

func (s SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) GetApiId() *string {
	return s.ApiId
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) GetRpsBegin() *int32 {
	return s.RpsBegin
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) GetRpsLimit() *int32 {
	return s.RpsLimit
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) SetApiId(v string) *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList {
	s.ApiId = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) SetRpsBegin(v int32) *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList {
	s.RpsBegin = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) SetRpsLimit(v int32) *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList {
	s.RpsLimit = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigApiLoadConfigList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneLoadConfigConfiguration struct {
	// The starting total number of concurrent virtual users in all sessions.
	//
	// The value is evenly distributed among all sessions if you set TestMode to concurrency_mode. If you do not specify this parameter, you must configure **relationLoadConfig**.
	//
	// example:
	//
	// 100
	AllConcurrencyBegin *int32 `json:"AllConcurrencyBegin,omitempty" xml:"AllConcurrencyBegin,omitempty"`
	// The maximum total number of concurrent virtual users in all sessions.
	//
	// The value is evenly distributed among all sessions if you set TestMode to concurrency_mode. If you do not specify this parameter, you must configure **relationLoadConfig**.
	//
	// example:
	//
	// 100
	AllConcurrencyLimit *int32 `json:"AllConcurrencyLimit,omitempty" xml:"AllConcurrencyLimit,omitempty"`
	// The starting RPS for all APIs.
	//
	// The value is evenly distributed among all APIs if you set TestMode to RPS. If you do not specify this parameter, you must specify **apiLoadConfig**.
	//
	// example:
	//
	// 100
	AllRpsBegin *int32 `json:"AllRpsBegin,omitempty" xml:"AllRpsBegin,omitempty"`
	// The maximum RPS for all APIs.
	//
	// The value is evenly distributed among all APIs if you set TestMode to RPS. If you do not specify this parameter, you must specify **apiLoadConfig**.
	//
	// example:
	//
	// 100
	AllRpsLimit *int32 `json:"AllRpsLimit,omitempty" xml:"AllRpsLimit,omitempty"`
}

func (s SavePtsSceneRequestSceneLoadConfigConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneLoadConfigConfiguration) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) GetAllConcurrencyBegin() *int32 {
	return s.AllConcurrencyBegin
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) GetAllConcurrencyLimit() *int32 {
	return s.AllConcurrencyLimit
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) GetAllRpsBegin() *int32 {
	return s.AllRpsBegin
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) GetAllRpsLimit() *int32 {
	return s.AllRpsLimit
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) SetAllConcurrencyBegin(v int32) *SavePtsSceneRequestSceneLoadConfigConfiguration {
	s.AllConcurrencyBegin = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) SetAllConcurrencyLimit(v int32) *SavePtsSceneRequestSceneLoadConfigConfiguration {
	s.AllConcurrencyLimit = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) SetAllRpsBegin(v int32) *SavePtsSceneRequestSceneLoadConfigConfiguration {
	s.AllRpsBegin = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) SetAllRpsLimit(v int32) *SavePtsSceneRequestSceneLoadConfigConfiguration {
	s.AllRpsLimit = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigConfiguration) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList struct {
	// The starting number of concurrent virtual users.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ConcurrencyBegin *int32 `json:"ConcurrencyBegin,omitempty" xml:"ConcurrencyBegin,omitempty"`
	// The maximum number of concurrent virtual users.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ConcurrencyLimit *int32 `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
	// The ID of the session.
	//
	// example:
	//
	// 1
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) GetConcurrencyBegin() *int32 {
	return s.ConcurrencyBegin
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) GetConcurrencyLimit() *int32 {
	return s.ConcurrencyLimit
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) GetRelationId() *string {
	return s.RelationId
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) SetConcurrencyBegin(v int32) *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList {
	s.ConcurrencyBegin = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) SetConcurrencyLimit(v int32) *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList {
	s.ConcurrencyLimit = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) SetRelationId(v string) *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList {
	s.RelationId = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigRelationLoadConfigList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneLoadConfigVpcLoadConfig struct {
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-jkasgfieiajidsjakjscb
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-skjfhlahsljkhsfalkjdoiw
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-akjhsdajgjsfggahjkga
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) SetRegionId(v string) *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig {
	s.RegionId = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) SetSecurityGroupId(v string) *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) SetVSwitchId(v string) *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig {
	s.VSwitchId = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) SetVpcId(v string) *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig {
	s.VpcId = &v
	return s
}

func (s *SavePtsSceneRequestSceneLoadConfigVpcLoadConfig) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneRelationList struct {
	// The API operations on the session.
	//
	// This parameter is required.
	ApiList []*SavePtsSceneRequestSceneRelationListApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	// The file parameters of the session.
	FileParameterExplainList []*SavePtsSceneRequestSceneRelationListFileParameterExplainList `json:"FileParameterExplainList,omitempty" xml:"FileParameterExplainList,omitempty" type:"Repeated"`
	// The ID of the session.
	//
	// example:
	//
	// 1
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The name of the session.
	//
	// This parameter is required.
	RelationName *string `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationList) GetApiList() []*SavePtsSceneRequestSceneRelationListApiList {
	return s.ApiList
}

func (s *SavePtsSceneRequestSceneRelationList) GetFileParameterExplainList() []*SavePtsSceneRequestSceneRelationListFileParameterExplainList {
	return s.FileParameterExplainList
}

func (s *SavePtsSceneRequestSceneRelationList) GetRelationId() *string {
	return s.RelationId
}

func (s *SavePtsSceneRequestSceneRelationList) GetRelationName() *string {
	return s.RelationName
}

func (s *SavePtsSceneRequestSceneRelationList) SetApiList(v []*SavePtsSceneRequestSceneRelationListApiList) *SavePtsSceneRequestSceneRelationList {
	s.ApiList = v
	return s
}

func (s *SavePtsSceneRequestSceneRelationList) SetFileParameterExplainList(v []*SavePtsSceneRequestSceneRelationListFileParameterExplainList) *SavePtsSceneRequestSceneRelationList {
	s.FileParameterExplainList = v
	return s
}

func (s *SavePtsSceneRequestSceneRelationList) SetRelationId(v string) *SavePtsSceneRequestSceneRelationList {
	s.RelationId = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationList) SetRelationName(v string) *SavePtsSceneRequestSceneRelationList {
	s.RelationName = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationList) Validate() error {
	if s.ApiList != nil {
		for _, item := range s.ApiList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FileParameterExplainList != nil {
		for _, item := range s.FileParameterExplainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SavePtsSceneRequestSceneRelationListApiList struct {
	// The ID of the API.
	//
	// example:
	//
	// 1
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// api
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The request body.
	//
	// if can be null:
	// true
	Body *SavePtsSceneRequestSceneRelationListApiListBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The checkpoints.
	CheckPointList []*SavePtsSceneRequestSceneRelationListApiListCheckPointList `json:"CheckPointList,omitempty" xml:"CheckPointList,omitempty" type:"Repeated"`
	// The export parameters.
	ExportList []*SavePtsSceneRequestSceneRelationListApiListExportList `json:"ExportList,omitempty" xml:"ExportList,omitempty" type:"Repeated"`
	// The headers.
	HeaderList []*SavePtsSceneRequestSceneRelationListApiListHeaderList `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Repeated"`
	// The request method.
	//
	// This parameter is required.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The number of redirections. The value can be 0, which specifies that redirection is allowed, or 10, which specifies that redirection is not allowed. You can specify a value based on your business requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RedirectCountLimit *int32 `json:"RedirectCountLimit,omitempty" xml:"RedirectCountLimit,omitempty"`
	// The timeout period of the API operation. Unit: seconds. Default: 5. Valid values: 1 to 60.
	//
	// example:
	//
	// 5
	TimeoutInSecond *int32 `json:"TimeoutInSecond,omitempty" xml:"TimeoutInSecond,omitempty"`
	// The URL to which the API request is sent.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationListApiList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationListApiList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetApiId() *string {
	return s.ApiId
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetApiName() *string {
	return s.ApiName
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetBody() *SavePtsSceneRequestSceneRelationListApiListBody {
	return s.Body
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetCheckPointList() []*SavePtsSceneRequestSceneRelationListApiListCheckPointList {
	return s.CheckPointList
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetExportList() []*SavePtsSceneRequestSceneRelationListApiListExportList {
	return s.ExportList
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetHeaderList() []*SavePtsSceneRequestSceneRelationListApiListHeaderList {
	return s.HeaderList
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetMethod() *string {
	return s.Method
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetRedirectCountLimit() *int32 {
	return s.RedirectCountLimit
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetTimeoutInSecond() *int32 {
	return s.TimeoutInSecond
}

func (s *SavePtsSceneRequestSceneRelationListApiList) GetUrl() *string {
	return s.Url
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetApiId(v string) *SavePtsSceneRequestSceneRelationListApiList {
	s.ApiId = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetApiName(v string) *SavePtsSceneRequestSceneRelationListApiList {
	s.ApiName = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetBody(v *SavePtsSceneRequestSceneRelationListApiListBody) *SavePtsSceneRequestSceneRelationListApiList {
	s.Body = v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetCheckPointList(v []*SavePtsSceneRequestSceneRelationListApiListCheckPointList) *SavePtsSceneRequestSceneRelationListApiList {
	s.CheckPointList = v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetExportList(v []*SavePtsSceneRequestSceneRelationListApiListExportList) *SavePtsSceneRequestSceneRelationListApiList {
	s.ExportList = v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetHeaderList(v []*SavePtsSceneRequestSceneRelationListApiListHeaderList) *SavePtsSceneRequestSceneRelationListApiList {
	s.HeaderList = v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetMethod(v string) *SavePtsSceneRequestSceneRelationListApiList {
	s.Method = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetRedirectCountLimit(v int32) *SavePtsSceneRequestSceneRelationListApiList {
	s.RedirectCountLimit = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetTimeoutInSecond(v int32) *SavePtsSceneRequestSceneRelationListApiList {
	s.TimeoutInSecond = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) SetUrl(v string) *SavePtsSceneRequestSceneRelationListApiList {
	s.Url = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiList) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.CheckPointList != nil {
		for _, item := range s.CheckPointList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExportList != nil {
		for _, item := range s.ExportList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HeaderList != nil {
		for _, item := range s.HeaderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SavePtsSceneRequestSceneRelationListApiListBody struct {
	// The data in the body. For example, {"key1":"value2","key2":"value2"}.
	//
	// example:
	//
	// {\\"global\\":\\"${global}\\",\\"name\\":\\"${name}\\"}
	BodyValue *string `json:"BodyValue,omitempty" xml:"BodyValue,omitempty"`
	// The body type. Default: `application/x-www-form-urlencoded`.
	//
	// example:
	//
	// application/x-www-form-urlencoded
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationListApiListBody) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationListApiListBody) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationListApiListBody) GetBodyValue() *string {
	return s.BodyValue
}

func (s *SavePtsSceneRequestSceneRelationListApiListBody) GetContentType() *string {
	return s.ContentType
}

func (s *SavePtsSceneRequestSceneRelationListApiListBody) SetBodyValue(v string) *SavePtsSceneRequestSceneRelationListApiListBody {
	s.BodyValue = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListBody) SetContentType(v string) *SavePtsSceneRequestSceneRelationListApiListBody {
	s.ContentType = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListBody) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneRelationListApiListCheckPointList struct {
	// The checked item.
	//
	// This parameter specifies the fields in the header if you specify `CheckType=HEADER` or the name of the export parameter if you specify `CheckType=EXPORTED_PARAM`.
	//
	// example:
	//
	// userId
	CheckPoint *string `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	// The type of check. Valid values:
	//
	// 	- BODY_TEXT: the response body
	//
	// 	- HEADER: the response headers
	//
	// 	- STATUS_CODE: the HTTP status code returned by the API
	//
	// 	- EXPORTED_PARAM: a specific export parameter
	//
	// example:
	//
	// EXPORTED_PARAM
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 111
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// The operation or condition that is checked against the expected value.
	//
	// example:
	//
	// ctn
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationListApiListCheckPointList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationListApiListCheckPointList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) GetCheckPoint() *string {
	return s.CheckPoint
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) GetCheckType() *string {
	return s.CheckType
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) GetOperator() *string {
	return s.Operator
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) SetCheckPoint(v string) *SavePtsSceneRequestSceneRelationListApiListCheckPointList {
	s.CheckPoint = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) SetCheckType(v string) *SavePtsSceneRequestSceneRelationListApiListCheckPointList {
	s.CheckType = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) SetExpectValue(v string) *SavePtsSceneRequestSceneRelationListApiListCheckPointList {
	s.ExpectValue = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) SetOperator(v string) *SavePtsSceneRequestSceneRelationListApiListCheckPointList {
	s.Operator = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListCheckPointList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneRelationListApiListExportList struct {
	// The index of the matched item. You can specify a number or "Random". If you set ExportType to BODY_TEXT, you must specify this parameter.
	//
	// example:
	//
	// 0
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the export parameter.
	//
	// example:
	//
	// test
	ExportName *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	// The source of the export parameter. Valid values:
	//
	// 	- BODY_TEXT: the request body in the BODY_TEXT format
	//
	// 	- BODY_JSON: the request body in the BODY_JSON format.
	//
	// 	- HEADER: the request header
	//
	// 	- STATUS_CODE: the HTTP status code returned by the API
	//
	// example:
	//
	// BODY_JSON
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The actual path from which you want to extract the export parameter values.
	//
	// example:
	//
	// data.itemlist[0]
	ExportValue *string `json:"ExportValue,omitempty" xml:"ExportValue,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationListApiListExportList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationListApiListExportList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) GetCount() *string {
	return s.Count
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) GetExportName() *string {
	return s.ExportName
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) GetExportType() *string {
	return s.ExportType
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) GetExportValue() *string {
	return s.ExportValue
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) SetCount(v string) *SavePtsSceneRequestSceneRelationListApiListExportList {
	s.Count = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) SetExportName(v string) *SavePtsSceneRequestSceneRelationListApiListExportList {
	s.ExportName = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) SetExportType(v string) *SavePtsSceneRequestSceneRelationListApiListExportList {
	s.ExportType = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) SetExportValue(v string) *SavePtsSceneRequestSceneRelationListApiListExportList {
	s.ExportValue = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListExportList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneRelationListApiListHeaderList struct {
	// The name of the header.
	//
	// example:
	//
	// Accept-Encoding
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The value of the header.
	//
	// example:
	//
	// gzip, deflate, br
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationListApiListHeaderList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationListApiListHeaderList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationListApiListHeaderList) GetHeaderName() *string {
	return s.HeaderName
}

func (s *SavePtsSceneRequestSceneRelationListApiListHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *SavePtsSceneRequestSceneRelationListApiListHeaderList) SetHeaderName(v string) *SavePtsSceneRequestSceneRelationListApiListHeaderList {
	s.HeaderName = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListHeaderList) SetHeaderValue(v string) *SavePtsSceneRequestSceneRelationListApiListHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListApiListHeaderList) Validate() error {
	return dara.Validate(s)
}

type SavePtsSceneRequestSceneRelationListFileParameterExplainList struct {
	// Specifies whether the file is used as the baseline file.
	//
	// example:
	//
	// true
	BaseFile *bool `json:"BaseFile,omitempty" xml:"BaseFile,omitempty"`
	// Specifies whether the file is used for a single execution of the test.
	//
	// example:
	//
	// true
	CycleOnce *bool `json:"CycleOnce,omitempty" xml:"CycleOnce,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// fileName.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The parameter names in the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// name,uid,age
	FileParamName *string `json:"FileParamName,omitempty" xml:"FileParamName,omitempty"`
}

func (s SavePtsSceneRequestSceneRelationListFileParameterExplainList) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneRequestSceneRelationListFileParameterExplainList) GoString() string {
	return s.String()
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) GetBaseFile() *bool {
	return s.BaseFile
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) GetCycleOnce() *bool {
	return s.CycleOnce
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) GetFileName() *string {
	return s.FileName
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) GetFileParamName() *string {
	return s.FileParamName
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) SetBaseFile(v bool) *SavePtsSceneRequestSceneRelationListFileParameterExplainList {
	s.BaseFile = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) SetCycleOnce(v bool) *SavePtsSceneRequestSceneRelationListFileParameterExplainList {
	s.CycleOnce = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) SetFileName(v string) *SavePtsSceneRequestSceneRelationListFileParameterExplainList {
	s.FileName = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) SetFileParamName(v string) *SavePtsSceneRequestSceneRelationListFileParameterExplainList {
	s.FileParamName = &v
	return s
}

func (s *SavePtsSceneRequestSceneRelationListFileParameterExplainList) Validate() error {
	return dara.Validate(s)
}
