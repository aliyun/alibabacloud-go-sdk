// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOpenJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenJMeterScene(v *SaveOpenJMeterSceneRequestOpenJMeterScene) *SaveOpenJMeterSceneRequest
	GetOpenJMeterScene() *SaveOpenJMeterSceneRequestOpenJMeterScene
}

type SaveOpenJMeterSceneRequest struct {
	// The details of the scenario.
	//
	// This parameter is required.
	OpenJMeterScene *SaveOpenJMeterSceneRequestOpenJMeterScene `json:"OpenJMeterScene,omitempty" xml:"OpenJMeterScene,omitempty" type:"Struct"`
}

func (s SaveOpenJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequest) GetOpenJMeterScene() *SaveOpenJMeterSceneRequestOpenJMeterScene {
	return s.OpenJMeterScene
}

func (s *SaveOpenJMeterSceneRequest) SetOpenJMeterScene(v *SaveOpenJMeterSceneRequestOpenJMeterScene) *SaveOpenJMeterSceneRequest {
	s.OpenJMeterScene = v
	return s
}

func (s *SaveOpenJMeterSceneRequest) Validate() error {
	if s.OpenJMeterScene != nil {
		if err := s.OpenJMeterScene.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveOpenJMeterSceneRequestOpenJMeterScene struct {
	// The number of stress testers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The maximum concurrency that is determined by the resource plans of users. You must configure this parameter when the mode is set to CONCURRENCY.
	//
	// example:
	//
	// 1000
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The type of the synchronization timer with fixed throughput in JMeter.
	//
	// example:
	//
	// GLOBAL
	ConstantThroughputTimerType *string `json:"ConstantThroughputTimerType,omitempty" xml:"ConstantThroughputTimerType,omitempty"`
	// The settings of Domain Name System (DNS).
	DnsCacheConfig *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig `json:"DnsCacheConfig,omitempty" xml:"DnsCacheConfig,omitempty" type:"Struct"`
	// The stress testing duration. The maximum stress testing duration is no more than one day, Unit: seconds. Valid values: 60 to 86400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 600
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the environment associated with the scenario.
	//
	// example:
	//
	// I8PZIH
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The test files.
	//
	// This parameter is required.
	FileList []*SaveOpenJMeterSceneRequestOpenJMeterSceneFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// Specifies whether the test is a virtual private cloud (VPC) test. The default value is false, which indicates a public network test. VPC-related settings take effect only when you set this parameter to true.
	//
	// example:
	//
	// true
	IsVpcTest *bool `json:"IsVpcTest,omitempty" xml:"IsVpcTest,omitempty"`
	// The JMeter properties.
	JMeterProperties []*SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties `json:"JMeterProperties,omitempty" xml:"JMeterProperties,omitempty" type:"Repeated"`
	// The JMeter plug-in tag.
	//
	// example:
	//
	// test
	JmeterPluginLabel *string `json:"JmeterPluginLabel,omitempty" xml:"JmeterPluginLabel,omitempty"`
	// The maximum RPS that takes effect in RPS mode.
	//
	// example:
	//
	// 100
	MaxRps *int32 `json:"MaxRps,omitempty" xml:"MaxRps,omitempty"`
	// The stress model.
	//
	// This parameter is required.
	//
	// example:
	//
	// CONCURRENCY
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ramp-up period. Unit: seconds.
	//
	// example:
	//
	// 600
	RampUp *int32 `json:"RampUp,omitempty" xml:"RampUp,omitempty"`
	// The region ID that is specified in the VPC test.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The requirements for the regions of the stress testers.
	RegionalCondition []*SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition `json:"RegionalCondition,omitempty" xml:"RegionalCondition,omitempty" type:"Repeated"`
	// The scenario ID. If you do not configure this parameter, the system creates a scenario. If you configure this parameter, the system updates the scenario corresponding to the ID specified by this parameter.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scenario name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The security group ID that is specified in the VPC test.
	//
	// example:
	//
	// sg-2zeid0dd7bhahsgdahspaly
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The initial concurrency that takes effect in concurrency mode. You must configure this parameter when the mode is set to CONCURRENCY.
	//
	// example:
	//
	// 10
	StartConcurrency *int32 `json:"StartConcurrency,omitempty" xml:"StartConcurrency,omitempty"`
	// The initial RPS that takes effect in RPS mode.
	//
	// example:
	//
	// 10
	StartRps *int32 `json:"StartRps,omitempty" xml:"StartRps,omitempty"`
	// The number of ramp-up steps.
	//
	// example:
	//
	// 3
	Steps *int32 `json:"Steps,omitempty" xml:"Steps,omitempty"`
	// The type of the synchronization timer in JMeter.
	//
	// example:
	//
	// GLOBAL
	SyncTimerType *string `json:"SyncTimerType,omitempty" xml:"SyncTimerType,omitempty"`
	// The test file.
	//
	// This parameter is required.
	//
	// example:
	//
	// baidu.jmx
	TestFile *string `json:"TestFile,omitempty" xml:"TestFile,omitempty"`
	// The vSwitch ID that is specified in the VPC test.
	//
	// example:
	//
	// vsw-2zehsgdhsahw1r
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID that is specified in the VPC test.
	//
	// example:
	//
	// vpc-2ze2sahjdgahsebjkqhf4pyj
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterScene) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterScene) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetConstantThroughputTimerType() *string {
	return s.ConstantThroughputTimerType
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetDnsCacheConfig() *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	return s.DnsCacheConfig
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetDuration() *int32 {
	return s.Duration
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetFileList() []*SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	return s.FileList
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetIsVpcTest() *bool {
	return s.IsVpcTest
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetJMeterProperties() []*SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties {
	return s.JMeterProperties
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetJmeterPluginLabel() *string {
	return s.JmeterPluginLabel
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetMaxRps() *int32 {
	return s.MaxRps
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetMode() *string {
	return s.Mode
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetRampUp() *int32 {
	return s.RampUp
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetRegionId() *string {
	return s.RegionId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetRegionalCondition() []*SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition {
	return s.RegionalCondition
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetSceneId() *string {
	return s.SceneId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetSceneName() *string {
	return s.SceneName
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetStartConcurrency() *int32 {
	return s.StartConcurrency
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetStartRps() *int32 {
	return s.StartRps
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetSteps() *int32 {
	return s.Steps
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetSyncTimerType() *string {
	return s.SyncTimerType
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetTestFile() *string {
	return s.TestFile
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) GetVpcId() *string {
	return s.VpcId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetAgentCount(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.AgentCount = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetConcurrency(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Concurrency = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetConstantThroughputTimerType(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.ConstantThroughputTimerType = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetDnsCacheConfig(v *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.DnsCacheConfig = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetDuration(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Duration = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetEnvironmentId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.EnvironmentId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetFileList(v []*SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.FileList = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetIsVpcTest(v bool) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.IsVpcTest = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetJMeterProperties(v []*SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.JMeterProperties = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetJmeterPluginLabel(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.JmeterPluginLabel = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetMaxRps(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.MaxRps = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetMode(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Mode = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetRampUp(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.RampUp = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetRegionId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.RegionId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetRegionalCondition(v []*SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.RegionalCondition = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSceneId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SceneId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSceneName(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SceneName = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSecurityGroupId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SecurityGroupId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetStartConcurrency(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.StartConcurrency = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetStartRps(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.StartRps = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSteps(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Steps = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSyncTimerType(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SyncTimerType = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetTestFile(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.TestFile = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetVSwitchId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.VSwitchId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetVpcId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.VpcId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) Validate() error {
	if s.DnsCacheConfig != nil {
		if err := s.DnsCacheConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FileList != nil {
		for _, item := range s.FileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.JMeterProperties != nil {
		for _, item := range s.JMeterProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionalCondition != nil {
		for _, item := range s.RegionalCondition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig struct {
	// Specifies whether to clear the cache in each iteration.
	//
	// example:
	//
	// true
	ClearCacheEachIteration *bool `json:"ClearCacheEachIteration,omitempty" xml:"ClearCacheEachIteration,omitempty"`
	// The DNS servers.
	DnsServers []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	// The table that contains bound domain names.
	HostTable map[string]*string `json:"HostTable,omitempty" xml:"HostTable,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) GetClearCacheEachIteration() *bool {
	return s.ClearCacheEachIteration
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) GetDnsServers() []*string {
	return s.DnsServers
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) GetHostTable() map[string]*string {
	return s.HostTable
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) SetClearCacheEachIteration(v bool) *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	s.ClearCacheEachIteration = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) SetDnsServers(v []*string) *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	s.DnsServers = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) SetHostTable(v map[string]*string) *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	s.HostTable = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) Validate() error {
	return dara.Validate(s)
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneFileList struct {
	// The file ID.
	//
	// example:
	//
	// 61232
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the test file.
	//
	// This parameter is required.
	//
	// example:
	//
	// baidu.jmx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The Object Storage Service (OSS) path that is used to access the test file over the Internet.
	//
	// >  Only test files in the China (Shanghai) region can be accessed.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://test.cn-shanghai.aliyuncs.com/baidu.jmx
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
	// The file size. The total file size cannot exceed 500 MB. Unit: bytes.
	//
	// example:
	//
	// 28880
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The MD5 hash of the test file.
	//
	// example:
	//
	// DA70F97A74D76B6A3BEF9CC8AE0D89EB
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// Specifies whether to split the test file. This parameter is valid only for CSV files.
	//
	// example:
	//
	// false
	SplitCsv *bool `json:"SplitCsv,omitempty" xml:"SplitCsv,omitempty"`
	// The file tag.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetFileId() *int64 {
	return s.FileId
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetFileName() *string {
	return s.FileName
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetMd5() *string {
	return s.Md5
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetSplitCsv() *bool {
	return s.SplitCsv
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GetTags() *string {
	return s.Tags
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileId(v int64) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileName(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileName = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileOssAddress(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileOssAddress = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileSize(v int64) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileSize = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetMd5(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.Md5 = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetSplitCsv(v bool) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.SplitCsv = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetTags(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.Tags = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) Validate() error {
	return dara.Validate(s)
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties struct {
	// The property name.
	//
	// example:
	//
	// https.sessioncontext.shared
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the property.
	//
	// example:
	//
	// false
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) GetName() *string {
	return s.Name
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) GetValue() *string {
	return s.Value
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) SetName(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties {
	s.Name = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) SetValue(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties {
	s.Value = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) Validate() error {
	return dara.Validate(s)
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition struct {
	// The number of stress tests. The sum of the number of stress tests in all regions must be equal to the AgentCount value of a specified scenario.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) GetAmount() *int32 {
	return s.Amount
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) GetRegion() *string {
	return s.Region
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) SetAmount(v int32) *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition {
	s.Amount = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) SetRegion(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition {
	s.Region = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneRegionalCondition) Validate() error {
	return dara.Validate(s)
}
