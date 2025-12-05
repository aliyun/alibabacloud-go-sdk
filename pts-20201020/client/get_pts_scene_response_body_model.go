// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPtsSceneResponseBody
	GetRequestId() *string
	SetScene(v *GetPtsSceneResponseBodyScene) *GetPtsSceneResponseBody
	GetScene() *GetPtsSceneResponseBodyScene
	SetSuccess(v bool) *GetPtsSceneResponseBody
	GetSuccess() *bool
}

type GetPtsSceneResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, N/A is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC4E3177-6745-4925-B423-4E89VV34221A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The structure of the scenario.
	Scene *GetPtsSceneResponseBodyScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Struct"`
	// Indicates whether the operation is successful.
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsSceneResponseBody) GetScene() *GetPtsSceneResponseBodyScene {
	return s.Scene
}

func (s *GetPtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsSceneResponseBody) SetCode(v string) *GetPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetMessage(v string) *GetPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetRequestId(v string) *GetPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetScene(v *GetPtsSceneResponseBodyScene) *GetPtsSceneResponseBody {
	s.Scene = v
	return s
}

func (s *GetPtsSceneResponseBody) SetSuccess(v bool) *GetPtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsSceneResponseBody) Validate() error {
	if s.Scene != nil {
		if err := s.Scene.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPtsSceneResponseBodyScene struct {
	// The advanced settings.
	AdvanceSetting *GetPtsSceneResponseBodySceneAdvanceSetting `json:"AdvanceSetting,omitempty" xml:"AdvanceSetting,omitempty" type:"Struct"`
	// The creation time of the scenario.
	//
	// example:
	//
	// 2021-02-26 15:30:30
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The file parameters.
	FileParameterList []*GetPtsSceneResponseBodySceneFileParameterList `json:"FileParameterList,omitempty" xml:"FileParameterList,omitempty" type:"Repeated"`
	// Global parameters
	GlobalParameterList []*GetPtsSceneResponseBodySceneGlobalParameterList `json:"GlobalParameterList,omitempty" xml:"GlobalParameterList,omitempty" type:"Repeated"`
	// The global headers for the scenario.
	Headers []*GetPtsSceneResponseBodySceneHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The load settings.
	LoadConfig *GetPtsSceneResponseBodySceneLoadConfig `json:"LoadConfig,omitempty" xml:"LoadConfig,omitempty" type:"Struct"`
	// The last modification time of the scenario.
	//
	// example:
	//
	// 2021-03-26 15:30:30
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The sessions.
	RelationList []*GetPtsSceneResponseBodySceneRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	// The ID of the scenario.
	//
	// example:
	//
	// BGFJ7GV
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The name of the scenario
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the scenario.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPtsSceneResponseBodyScene) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodyScene) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodyScene) GetAdvanceSetting() *GetPtsSceneResponseBodySceneAdvanceSetting {
	return s.AdvanceSetting
}

func (s *GetPtsSceneResponseBodyScene) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPtsSceneResponseBodyScene) GetFileParameterList() []*GetPtsSceneResponseBodySceneFileParameterList {
	return s.FileParameterList
}

func (s *GetPtsSceneResponseBodyScene) GetGlobalParameterList() []*GetPtsSceneResponseBodySceneGlobalParameterList {
	return s.GlobalParameterList
}

func (s *GetPtsSceneResponseBodyScene) GetHeaders() []*GetPtsSceneResponseBodySceneHeaders {
	return s.Headers
}

func (s *GetPtsSceneResponseBodyScene) GetLoadConfig() *GetPtsSceneResponseBodySceneLoadConfig {
	return s.LoadConfig
}

func (s *GetPtsSceneResponseBodyScene) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetPtsSceneResponseBodyScene) GetRelationList() []*GetPtsSceneResponseBodySceneRelationList {
	return s.RelationList
}

func (s *GetPtsSceneResponseBodyScene) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsSceneResponseBodyScene) GetSceneName() *string {
	return s.SceneName
}

func (s *GetPtsSceneResponseBodyScene) GetStatus() *string {
	return s.Status
}

func (s *GetPtsSceneResponseBodyScene) SetAdvanceSetting(v *GetPtsSceneResponseBodySceneAdvanceSetting) *GetPtsSceneResponseBodyScene {
	s.AdvanceSetting = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetCreateTime(v string) *GetPtsSceneResponseBodyScene {
	s.CreateTime = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetFileParameterList(v []*GetPtsSceneResponseBodySceneFileParameterList) *GetPtsSceneResponseBodyScene {
	s.FileParameterList = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetGlobalParameterList(v []*GetPtsSceneResponseBodySceneGlobalParameterList) *GetPtsSceneResponseBodyScene {
	s.GlobalParameterList = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetHeaders(v []*GetPtsSceneResponseBodySceneHeaders) *GetPtsSceneResponseBodyScene {
	s.Headers = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetLoadConfig(v *GetPtsSceneResponseBodySceneLoadConfig) *GetPtsSceneResponseBodyScene {
	s.LoadConfig = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetModifiedTime(v string) *GetPtsSceneResponseBodyScene {
	s.ModifiedTime = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetRelationList(v []*GetPtsSceneResponseBodySceneRelationList) *GetPtsSceneResponseBodyScene {
	s.RelationList = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetSceneId(v string) *GetPtsSceneResponseBodyScene {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetSceneName(v string) *GetPtsSceneResponseBodyScene {
	s.SceneName = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetStatus(v string) *GetPtsSceneResponseBodyScene {
	s.Status = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) Validate() error {
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
	if s.Headers != nil {
		for _, item := range s.Headers {
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

type GetPtsSceneResponseBodySceneAdvanceSetting struct {
	// The timeout period of the scenario. Unit: seconds.
	//
	// example:
	//
	// 5
	ConnectionTimeoutInSecond *int32 `json:"ConnectionTimeoutInSecond,omitempty" xml:"ConnectionTimeoutInSecond,omitempty"`
	// The IP-domain name bindings.
	DomainBindingList []*GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList `json:"DomainBindingList,omitempty" xml:"DomainBindingList,omitempty" type:"Repeated"`
	// The log sampling rate.
	//
	// example:
	//
	// 1
	LogRate *int32 `json:"LogRate,omitempty" xml:"LogRate,omitempty"`
	// The custom success code.
	//
	// example:
	//
	// 429,304
	SuccessCode *string `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty"`
}

func (s GetPtsSceneResponseBodySceneAdvanceSetting) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneAdvanceSetting) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) GetConnectionTimeoutInSecond() *int32 {
	return s.ConnectionTimeoutInSecond
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) GetDomainBindingList() []*GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList {
	return s.DomainBindingList
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) GetLogRate() *int32 {
	return s.LogRate
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) GetSuccessCode() *string {
	return s.SuccessCode
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetConnectionTimeoutInSecond(v int32) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.ConnectionTimeoutInSecond = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetDomainBindingList(v []*GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.DomainBindingList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetLogRate(v int32) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.LogRate = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetSuccessCode(v string) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.SuccessCode = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) Validate() error {
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

type GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList struct {
	// The domain name.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IPs bound to the domain name.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) GetDomain() *string {
	return s.Domain
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) GetIps() []*string {
	return s.Ips
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) SetDomain(v string) *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList {
	s.Domain = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) SetIps(v []*string) *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList {
	s.Ips = v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneFileParameterList struct {
	// The file name.
	//
	// example:
	//
	// city.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS address of the file. Make sure that the address is accessible from the Internet.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/json.jar
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s GetPtsSceneResponseBodySceneFileParameterList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneFileParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) GetFileName() *string {
	return s.FileName
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) SetFileName(v string) *GetPtsSceneResponseBodySceneFileParameterList {
	s.FileName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) SetFileOssAddress(v string) *GetPtsSceneResponseBodySceneFileParameterList {
	s.FileOssAddress = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneGlobalParameterList struct {
	// The name of the parameter.
	//
	// example:
	//
	// userName
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// lisi
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s GetPtsSceneResponseBodySceneGlobalParameterList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneGlobalParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) GetParamName() *string {
	return s.ParamName
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) GetParamValue() *string {
	return s.ParamValue
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) SetParamName(v string) *GetPtsSceneResponseBodySceneGlobalParameterList {
	s.ParamName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) SetParamValue(v string) *GetPtsSceneResponseBodySceneGlobalParameterList {
	s.ParamValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneHeaders struct {
	// The name of the header.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the header.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPtsSceneResponseBodySceneHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneHeaders) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneHeaders) GetName() *string {
	return s.Name
}

func (s *GetPtsSceneResponseBodySceneHeaders) GetValue() *string {
	return s.Value
}

func (s *GetPtsSceneResponseBodySceneHeaders) SetName(v string) *GetPtsSceneResponseBodySceneHeaders {
	s.Name = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneHeaders) SetValue(v string) *GetPtsSceneResponseBodySceneHeaders {
	s.Value = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneHeaders) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneLoadConfig struct {
	// The number of load generators.
	//
	// example:
	//
	// 1
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The API request load settings.
	ApiLoadConfigList []*GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList `json:"ApiLoadConfigList,omitempty" xml:"ApiLoadConfigList,omitempty" type:"Repeated"`
	// Indicates whether the load is automatically incremented.
	//
	// example:
	//
	// false
	AutoStep *bool `json:"AutoStep,omitempty" xml:"AutoStep,omitempty"`
	// The concurrency and RPS settings of the scenario.
	Configuration *GetPtsSceneResponseBodySceneLoadConfigConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// The increment percentage. The valid values are 10 to 100, in increments of 10. This parameter is returned only if you set testMode to concurrency_mode and set autoStep to true.
	//
	// example:
	//
	// 10
	Increment *int32 `json:"Increment,omitempty" xml:"Increment,omitempty"`
	// The duration during which a specific load level is applied. The duration is less than the value of maxRunningTime. Unit: minutes.
	//
	// example:
	//
	// 2
	KeepTime *int32 `json:"KeepTime,omitempty" xml:"KeepTime,omitempty"`
	// The maximum duration of load application. Unit: minutes.
	//
	// example:
	//
	// 2
	MaxRunningTime *int32 `json:"MaxRunningTime,omitempty" xml:"MaxRunningTime,omitempty"`
	// The session load settings.
	RelationLoadConfigList []*GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList `json:"RelationLoadConfigList,omitempty" xml:"RelationLoadConfigList,omitempty" type:"Repeated"`
	// The load application mode. Transactions per second (TPS) indicates the RPS mode.
	//
	// >  The load application mode is CONCURRENCY/TPS.
	//
	// example:
	//
	// TPS
	TestMode *string `json:"TestMode,omitempty" xml:"TestMode,omitempty"`
	// The virtual private cloud (VPC) settings. This information is returned only if you set the testing mode to VPC.
	VpcLoadConfig *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig `json:"VpcLoadConfig,omitempty" xml:"VpcLoadConfig,omitempty" type:"Struct"`
}

func (s GetPtsSceneResponseBodySceneLoadConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfig) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetApiLoadConfigList() []*GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	return s.ApiLoadConfigList
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetAutoStep() *bool {
	return s.AutoStep
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetConfiguration() *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	return s.Configuration
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetIncrement() *int32 {
	return s.Increment
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetKeepTime() *int32 {
	return s.KeepTime
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetMaxRunningTime() *int32 {
	return s.MaxRunningTime
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetRelationLoadConfigList() []*GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	return s.RelationLoadConfigList
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetTestMode() *string {
	return s.TestMode
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) GetVpcLoadConfig() *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig {
	return s.VpcLoadConfig
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetAgentCount(v int32) *GetPtsSceneResponseBodySceneLoadConfig {
	s.AgentCount = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetApiLoadConfigList(v []*GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) *GetPtsSceneResponseBodySceneLoadConfig {
	s.ApiLoadConfigList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetAutoStep(v bool) *GetPtsSceneResponseBodySceneLoadConfig {
	s.AutoStep = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetConfiguration(v *GetPtsSceneResponseBodySceneLoadConfigConfiguration) *GetPtsSceneResponseBodySceneLoadConfig {
	s.Configuration = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetIncrement(v int32) *GetPtsSceneResponseBodySceneLoadConfig {
	s.Increment = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetKeepTime(v int32) *GetPtsSceneResponseBodySceneLoadConfig {
	s.KeepTime = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetMaxRunningTime(v int32) *GetPtsSceneResponseBodySceneLoadConfig {
	s.MaxRunningTime = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetRelationLoadConfigList(v []*GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) *GetPtsSceneResponseBodySceneLoadConfig {
	s.RelationLoadConfigList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetTestMode(v string) *GetPtsSceneResponseBodySceneLoadConfig {
	s.TestMode = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetVpcLoadConfig(v *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) *GetPtsSceneResponseBodySceneLoadConfig {
	s.VpcLoadConfig = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) Validate() error {
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

type GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList struct {
	// The API ID. You can track an API by its ID in sessions.
	//
	// example:
	//
	// GBFDCV8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The starting requests per second (RPS).
	//
	// example:
	//
	// 10
	RpsBegin *int32 `json:"RpsBegin,omitempty" xml:"RpsBegin,omitempty"`
	// The maximum RPS.
	//
	// example:
	//
	// 20
	RpsLimit *int32 `json:"RpsLimit,omitempty" xml:"RpsLimit,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) GetApiId() *string {
	return s.ApiId
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) GetRpsBegin() *int32 {
	return s.RpsBegin
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) GetRpsLimit() *int32 {
	return s.RpsLimit
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) SetApiId(v string) *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	s.ApiId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) SetRpsBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	s.RpsBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) SetRpsLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	s.RpsLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneLoadConfigConfiguration struct {
	// The starting number of concurrent sessions.
	//
	// example:
	//
	// 100
	AllConcurrencyBegin *int32 `json:"AllConcurrencyBegin,omitempty" xml:"AllConcurrencyBegin,omitempty"`
	// The maximum number of concurrent sessions.
	//
	// example:
	//
	// 200
	AllConcurrencyLimit *int32 `json:"AllConcurrencyLimit,omitempty" xml:"AllConcurrencyLimit,omitempty"`
	// The starting RPS.
	//
	// example:
	//
	// 100
	AllRpsBegin *int32 `json:"AllRpsBegin,omitempty" xml:"AllRpsBegin,omitempty"`
	// The maximum RPS.
	//
	// example:
	//
	// 200
	AllRpsLimit *int32 `json:"AllRpsLimit,omitempty" xml:"AllRpsLimit,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigConfiguration) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) GetAllConcurrencyBegin() *int32 {
	return s.AllConcurrencyBegin
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) GetAllConcurrencyLimit() *int32 {
	return s.AllConcurrencyLimit
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) GetAllRpsBegin() *int32 {
	return s.AllRpsBegin
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) GetAllRpsLimit() *int32 {
	return s.AllRpsLimit
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllConcurrencyBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllConcurrencyBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllConcurrencyLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllConcurrencyLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllRpsBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllRpsBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllRpsLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllRpsLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList struct {
	// The starting number of concurrent sessions.
	//
	// example:
	//
	// 10
	ConcurrencyBegin *int32 `json:"ConcurrencyBegin,omitempty" xml:"ConcurrencyBegin,omitempty"`
	// The maximum number of concurrent sessions.
	//
	// example:
	//
	// 20
	ConcurrencyLimit *int32 `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
	// The session ID.
	//
	// example:
	//
	// HNBGS7M
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) GetConcurrencyBegin() *int32 {
	return s.ConcurrencyBegin
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) GetConcurrencyLimit() *int32 {
	return s.ConcurrencyLimit
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) GetRelationId() *string {
	return s.RelationId
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) SetConcurrencyBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	s.ConcurrencyBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) SetConcurrencyLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	s.ConcurrencyLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) SetRelationId(v string) *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	s.RelationId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-jkasgfieiajidsjakjscb
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-skjfhlahsljkhsfalkjdoiw
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-akjhsdajgjsfggahjkga
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) SetRegionId(v string) *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig {
	s.RegionId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) SetSecurityGroupId(v string) *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) SetVSwitchId(v string) *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig {
	s.VSwitchId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) SetVpcId(v string) *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig {
	s.VpcId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigVpcLoadConfig) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneRelationList struct {
	// The APIs.
	ApiList []*GetPtsSceneResponseBodySceneRelationListApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	// The file parameters.
	FileParameterExplainList []*GetPtsSceneResponseBodySceneRelationListFileParameterExplainList `json:"FileParameterExplainList,omitempty" xml:"FileParameterExplainList,omitempty" type:"Repeated"`
	// The session ID.
	//
	// example:
	//
	// HNBGS7M
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The session name.
	RelationName *string `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationList) GetApiList() []*GetPtsSceneResponseBodySceneRelationListApiList {
	return s.ApiList
}

func (s *GetPtsSceneResponseBodySceneRelationList) GetFileParameterExplainList() []*GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	return s.FileParameterExplainList
}

func (s *GetPtsSceneResponseBodySceneRelationList) GetRelationId() *string {
	return s.RelationId
}

func (s *GetPtsSceneResponseBodySceneRelationList) GetRelationName() *string {
	return s.RelationName
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetApiList(v []*GetPtsSceneResponseBodySceneRelationListApiList) *GetPtsSceneResponseBodySceneRelationList {
	s.ApiList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetFileParameterExplainList(v []*GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) *GetPtsSceneResponseBodySceneRelationList {
	s.FileParameterExplainList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetRelationId(v string) *GetPtsSceneResponseBodySceneRelationList {
	s.RelationId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetRelationName(v string) *GetPtsSceneResponseBodySceneRelationList {
	s.RelationName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) Validate() error {
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

type GetPtsSceneResponseBodySceneRelationListApiList struct {
	// The API ID. You can track an API by its ID in sessions.
	//
	// example:
	//
	// GBFDCV8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The request body.
	Body *GetPtsSceneResponseBodySceneRelationListApiListBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The checkpoints.
	CheckPointList []*GetPtsSceneResponseBodySceneRelationListApiListCheckPointList `json:"CheckPointList,omitempty" xml:"CheckPointList,omitempty" type:"Repeated"`
	// The exported parameters.
	ExportList []*GetPtsSceneResponseBodySceneRelationListApiListExportList `json:"ExportList,omitempty" xml:"ExportList,omitempty" type:"Repeated"`
	// The headers used in the API request.
	HeaderList []*GetPtsSceneResponseBodySceneRelationListApiListHeaderList `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Repeated"`
	// The request method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The number of redirections.
	//
	// example:
	//
	// 5
	RedirectCountLimit *int32 `json:"RedirectCountLimit,omitempty" xml:"RedirectCountLimit,omitempty"`
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 5
	TimeoutInSecond *int32 `json:"TimeoutInSecond,omitempty" xml:"TimeoutInSecond,omitempty"`
	// The URL to which the request is sent.
	//
	// example:
	//
	// https://www.aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetApiId() *string {
	return s.ApiId
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetApiName() *string {
	return s.ApiName
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetBody() *GetPtsSceneResponseBodySceneRelationListApiListBody {
	return s.Body
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetCheckPointList() []*GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	return s.CheckPointList
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetExportList() []*GetPtsSceneResponseBodySceneRelationListApiListExportList {
	return s.ExportList
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetHeaderList() []*GetPtsSceneResponseBodySceneRelationListApiListHeaderList {
	return s.HeaderList
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetMethod() *string {
	return s.Method
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetRedirectCountLimit() *int32 {
	return s.RedirectCountLimit
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetTimeoutInSecond() *int32 {
	return s.TimeoutInSecond
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) GetUrl() *string {
	return s.Url
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetApiId(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.ApiId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetApiName(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.ApiName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetBody(v *GetPtsSceneResponseBodySceneRelationListApiListBody) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.Body = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetCheckPointList(v []*GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.CheckPointList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetExportList(v []*GetPtsSceneResponseBodySceneRelationListApiListExportList) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.ExportList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetHeaderList(v []*GetPtsSceneResponseBodySceneRelationListApiListHeaderList) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.HeaderList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetMethod(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.Method = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetRedirectCountLimit(v int32) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.RedirectCountLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetTimeoutInSecond(v int32) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.TimeoutInSecond = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetUrl(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.Url = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) Validate() error {
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

type GetPtsSceneResponseBodySceneRelationListApiListBody struct {
	// The body value.
	//
	// example:
	//
	// {\\"key1\\":\\"111\\",\\"key2\\":\\"222\\"}
	BodyValue *string `json:"BodyValue,omitempty" xml:"BodyValue,omitempty"`
	// The body type.
	//
	// example:
	//
	// application/x-www-form-urlencoded
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) GetBodyValue() *string {
	return s.BodyValue
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) GetContentType() *string {
	return s.ContentType
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) SetBodyValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListBody {
	s.BodyValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) SetContentType(v string) *GetPtsSceneResponseBodySceneRelationListApiListBody {
	s.ContentType = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneRelationListApiListCheckPointList struct {
	// The checked parameter.
	//
	// example:
	//
	// userId
	CheckPoint *string `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	// The check type.
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
	// The check operator.
	//
	// example:
	//
	// ctn
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) GetCheckPoint() *string {
	return s.CheckPoint
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) GetCheckType() *string {
	return s.CheckType
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) GetOperator() *string {
	return s.Operator
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetCheckPoint(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.CheckPoint = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetCheckType(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.CheckType = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetExpectValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.ExpectValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetOperator(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.Operator = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneRelationListApiListExportList struct {
	// The number of items or entries related to the export operation.
	//
	// example:
	//
	// 0
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The path where the exported value can be found.
	//
	// example:
	//
	// data.username
	ExportName *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	// The format in which data is exported.
	//
	// example:
	//
	// BODY_JSON
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The parameter that is exported.
	//
	// example:
	//
	// username
	ExportValue *string `json:"ExportValue,omitempty" xml:"ExportValue,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListExportList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListExportList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) GetCount() *string {
	return s.Count
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) GetExportName() *string {
	return s.ExportName
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) GetExportType() *string {
	return s.ExportType
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) GetExportValue() *string {
	return s.ExportValue
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetCount(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.Count = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetExportName(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.ExportName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetExportType(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.ExportType = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetExportValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.ExportValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneRelationListApiListHeaderList struct {
	// The header name.
	//
	// example:
	//
	// userId
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The header value.
	//
	// example:
	//
	// 1111
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListHeaderList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListHeaderList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) GetHeaderName() *string {
	return s.HeaderName
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) SetHeaderName(v string) *GetPtsSceneResponseBodySceneRelationListApiListHeaderList {
	s.HeaderName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) SetHeaderValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneResponseBodySceneRelationListFileParameterExplainList struct {
	// Indicates whether the file serves as the primary dataset for the test.
	//
	// example:
	//
	// true
	BaseFile *bool `json:"BaseFile,omitempty" xml:"BaseFile,omitempty"`
	// Indicates whether the parameters are used for a single test execution.
	//
	// example:
	//
	// true
	CycleOnce *bool `json:"CycleOnce,omitempty" xml:"CycleOnce,omitempty"`
	// The file name.
	//
	// example:
	//
	// city.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The parameter names in the file.
	//
	// example:
	//
	// userName,age
	FileParamName *string `json:"FileParamName,omitempty" xml:"FileParamName,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) GetBaseFile() *bool {
	return s.BaseFile
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) GetCycleOnce() *bool {
	return s.CycleOnce
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) GetFileName() *string {
	return s.FileName
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) GetFileParamName() *string {
	return s.FileParamName
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetBaseFile(v bool) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.BaseFile = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetCycleOnce(v bool) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.CycleOnce = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetFileName(v string) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.FileName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetFileParamName(v string) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.FileParamName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) Validate() error {
	return dara.Validate(s)
}
