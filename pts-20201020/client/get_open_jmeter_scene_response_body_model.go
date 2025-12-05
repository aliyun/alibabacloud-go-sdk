// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpenJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetOpenJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOpenJMeterSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpenJMeterSceneResponseBody
	GetRequestId() *string
	SetScene(v *GetOpenJMeterSceneResponseBodyScene) *GetOpenJMeterSceneResponseBody
	GetScene() *GetOpenJMeterSceneResponseBodyScene
	SetSuccess(v bool) *GetOpenJMeterSceneResponseBody
	GetSuccess() *bool
}

type GetOpenJMeterSceneResponseBody struct {
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
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the scenario.
	Scene *GetOpenJMeterSceneResponseBodyScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Struct"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpenJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpenJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOpenJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpenJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpenJMeterSceneResponseBody) GetScene() *GetOpenJMeterSceneResponseBodyScene {
	return s.Scene
}

func (s *GetOpenJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpenJMeterSceneResponseBody) SetCode(v string) *GetOpenJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetHttpStatusCode(v int32) *GetOpenJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetMessage(v string) *GetOpenJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetRequestId(v string) *GetOpenJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetScene(v *GetOpenJMeterSceneResponseBodyScene) *GetOpenJMeterSceneResponseBody {
	s.Scene = v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetSuccess(v bool) *GetOpenJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) Validate() error {
	if s.Scene != nil {
		if err := s.Scene.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOpenJMeterSceneResponseBodyScene struct {
	// The number of load generators. A load generator supports up to 500 concurrent virtual users.
	//
	// example:
	//
	// 2
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The basic information.
	BaseInfo *GetOpenJMeterSceneResponseBodySceneBaseInfo `json:"BaseInfo,omitempty" xml:"BaseInfo,omitempty" type:"Struct"`
	// The maximum number of concurrent virtual users.
	//
	// example:
	//
	// 1000
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The type of the constant throughput timer.
	//
	// example:
	//
	// STAND_ALONE
	ConstantThroughputTimerType *string `json:"ConstantThroughputTimerType,omitempty" xml:"ConstantThroughputTimerType,omitempty"`
	// The DNS settings.
	DnsCacheConfig *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig `json:"DnsCacheConfig,omitempty" xml:"DnsCacheConfig,omitempty" type:"Struct"`
	// The duration of the performance testing. Unit: seconds.
	//
	// example:
	//
	// 600
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// EEDT7
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The files.
	FileList []*GetOpenJMeterSceneResponseBodySceneFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// Indicates whether the load is from a virtual private cloud (VPC).
	//
	// example:
	//
	// false
	IsVpcTest *bool `json:"IsVpcTest,omitempty" xml:"IsVpcTest,omitempty"`
	// The maximum RPS. This parameter is returned if you set Mode to tps_mode.
	//
	// example:
	//
	// true
	MaxRps *int32 `json:"MaxRps,omitempty" xml:"MaxRps,omitempty"`
	// The load application mode. Valid values: concurrency_mode and tps_mode.
	//
	// example:
	//
	// concurrency_mode
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The origin of the load. "" indicates the Internet and intranet-vpc indicates the VPC.
	//
	// example:
	//
	// VPC
	Pool *string `json:"Pool,omitempty" xml:"Pool,omitempty"`
	// The period of time during which the load is gradually increased to the desired level. Unit: seconds.
	//
	// example:
	//
	// 100
	RampUp *int32 `json:"RampUp,omitempty" xml:"RampUp,omitempty"`
	// The region ID. This parameter is returned if the load is from a VPC.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Customized load generator settings for regions
	RegionalCondition []*GetOpenJMeterSceneResponseBodySceneRegionalCondition `json:"RegionalCondition,omitempty" xml:"RegionalCondition,omitempty" type:"Repeated"`
	// The ID of the scenario.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The name of the scenario.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The ID of the security group. This parameter is returned if the load is from a VPC.
	//
	// example:
	//
	// sg-2zeid0dd7bhahsgdahspaly
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The start number of concurrent virtual users.
	//
	// example:
	//
	// true
	StartConcurrency *int32 `json:"StartConcurrency,omitempty" xml:"StartConcurrency,omitempty"`
	// The start requests per second (RPS). This parameter is returned if you set Mode to tps_mode.
	//
	// example:
	//
	// true
	StartRps *int32 `json:"StartRps,omitempty" xml:"StartRps,omitempty"`
	// The number of incremented users per step. If RampUp or Steps is not specified, the fixed load is used. If RampUp is specified but Steps is not specified, the load increases uniformly based on the value of RampUp. If RampUp and Steps are specified and Steps is less than RampUp, the load increases based on the value of Steps. You cannot specify Steps without specifying RampUp. If you do so, the fixed load is used.
	//
	// example:
	//
	// 3
	Steps *int32 `json:"Steps,omitempty" xml:"Steps,omitempty"`
	// The type of the synchronization timer.
	//
	// example:
	//
	// GLOBAL
	SyncTimerType *string `json:"SyncTimerType,omitempty" xml:"SyncTimerType,omitempty"`
	// The test file.
	//
	// example:
	//
	// baidu.jmx
	TestFile *string `json:"TestFile,omitempty" xml:"TestFile,omitempty"`
	// The ID of the vSwitch. This parameter is returned if the load is from a VPC.
	//
	// example:
	//
	// vsw-2zehsgdhsahw1r
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC. This parameter is returned if the load is from a VPC.
	//
	// example:
	//
	// vpc-2ze2sahjdgahsebjkqhf4pyj
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodyScene) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodyScene) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetBaseInfo() *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	return s.BaseInfo
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetConstantThroughputTimerType() *string {
	return s.ConstantThroughputTimerType
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetDnsCacheConfig() *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	return s.DnsCacheConfig
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetDuration() *int32 {
	return s.Duration
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetFileList() []*GetOpenJMeterSceneResponseBodySceneFileList {
	return s.FileList
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetIsVpcTest() *bool {
	return s.IsVpcTest
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetMaxRps() *int32 {
	return s.MaxRps
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetMode() *string {
	return s.Mode
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetPool() *string {
	return s.Pool
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetRampUp() *int32 {
	return s.RampUp
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetRegionalCondition() []*GetOpenJMeterSceneResponseBodySceneRegionalCondition {
	return s.RegionalCondition
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetSceneId() *string {
	return s.SceneId
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetSceneName() *string {
	return s.SceneName
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetStartConcurrency() *int32 {
	return s.StartConcurrency
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetStartRps() *int32 {
	return s.StartRps
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetSteps() *int32 {
	return s.Steps
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetSyncTimerType() *string {
	return s.SyncTimerType
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetTestFile() *string {
	return s.TestFile
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetOpenJMeterSceneResponseBodyScene) GetVpcId() *string {
	return s.VpcId
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetAgentCount(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.AgentCount = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetBaseInfo(v *GetOpenJMeterSceneResponseBodySceneBaseInfo) *GetOpenJMeterSceneResponseBodyScene {
	s.BaseInfo = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetConcurrency(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.Concurrency = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetConstantThroughputTimerType(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.ConstantThroughputTimerType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetDnsCacheConfig(v *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) *GetOpenJMeterSceneResponseBodyScene {
	s.DnsCacheConfig = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetDuration(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.Duration = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetEnvironmentId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.EnvironmentId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetFileList(v []*GetOpenJMeterSceneResponseBodySceneFileList) *GetOpenJMeterSceneResponseBodyScene {
	s.FileList = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetIsVpcTest(v bool) *GetOpenJMeterSceneResponseBodyScene {
	s.IsVpcTest = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetMaxRps(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.MaxRps = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetMode(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.Mode = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetPool(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.Pool = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetRampUp(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.RampUp = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetRegionId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.RegionId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetRegionalCondition(v []*GetOpenJMeterSceneResponseBodySceneRegionalCondition) *GetOpenJMeterSceneResponseBodyScene {
	s.RegionalCondition = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSceneId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SceneId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSceneName(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SceneName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSecurityGroupId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SecurityGroupId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetStartConcurrency(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.StartConcurrency = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetStartRps(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.StartRps = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSteps(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.Steps = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSyncTimerType(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SyncTimerType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetTestFile(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.TestFile = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetVSwitchId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.VSwitchId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetVpcId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.VpcId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) Validate() error {
	if s.BaseInfo != nil {
		if err := s.BaseInfo.Validate(); err != nil {
			return err
		}
	}
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

type GetOpenJMeterSceneResponseBodySceneBaseInfo struct {
	// The name of the creator.
	//
	// example:
	//
	// 张三
	CreateName *string `json:"CreateName,omitempty" xml:"CreateName,omitempty"`
	// The name of the modifier.
	//
	// example:
	//
	// 里斯
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The type of the operation.
	//
	// example:
	//
	// 保存去压测
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The person who takes charge of the performance testing.
	//
	// example:
	//
	// test-person
	Principal *string `json:"Principal,omitempty" xml:"Principal,omitempty"`
	// The comment.
	//
	// example:
	//
	// 小心压测
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The origin of the scenario.
	//
	// example:
	//
	// create
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneBaseInfo) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) GetCreateName() *string {
	return s.CreateName
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) GetModifyName() *string {
	return s.ModifyName
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) GetOperateType() *string {
	return s.OperateType
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) GetPrincipal() *string {
	return s.Principal
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) GetRemark() *string {
	return s.Remark
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) GetResource() *string {
	return s.Resource
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetCreateName(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.CreateName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetModifyName(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.ModifyName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetOperateType(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.OperateType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetPrincipal(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.Principal = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetRemark(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.Remark = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetResource(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.Resource = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) Validate() error {
	return dara.Validate(s)
}

type GetOpenJMeterSceneResponseBodySceneDnsCacheConfig struct {
	// Indicates whether the cache is cleared.
	//
	// example:
	//
	// false
	ClearCacheEachIteration *bool `json:"ClearCacheEachIteration,omitempty" xml:"ClearCacheEachIteration,omitempty"`
	// The DNS servers
	DnsServers []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	// The domain name and its bounded IP address.
	//
	// example:
	//
	// {"server.com":"6.6.6.6"}
	HostTable map[string]interface{} `json:"HostTable,omitempty" xml:"HostTable,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) GetClearCacheEachIteration() *bool {
	return s.ClearCacheEachIteration
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) GetDnsServers() []*string {
	return s.DnsServers
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) GetHostTable() map[string]interface{} {
	return s.HostTable
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) SetClearCacheEachIteration(v bool) *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	s.ClearCacheEachIteration = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) SetDnsServers(v []*string) *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	s.DnsServers = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) SetHostTable(v map[string]interface{}) *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	s.HostTable = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) Validate() error {
	return dara.Validate(s)
}

type GetOpenJMeterSceneResponseBodySceneFileList struct {
	// The name of the file.
	//
	// example:
	//
	// json.jar
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The Object Storage Service (OSS) URL of the file.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/json.jar
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 700
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The type of the file.
	//
	// example:
	//
	// jar
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 61660
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The MD5 value of the JAR package.
	//
	// example:
	//
	// 43B584026CE5E570F3DE638FA7EEF9E0
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// Indicates whether the file is split.
	//
	// example:
	//
	// false
	SplitCsv *bool `json:"SplitCsv,omitempty" xml:"SplitCsv,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneFileList) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneFileList) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetFileName() *string {
	return s.FileName
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetFileType() *string {
	return s.FileType
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetId() *int64 {
	return s.Id
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetMd5() *string {
	return s.Md5
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) GetSplitCsv() *bool {
	return s.SplitCsv
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileName(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileOssAddress(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileOssAddress = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileSize(v int64) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileSize = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileType(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetId(v int64) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.Id = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetMd5(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.Md5 = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetSplitCsv(v bool) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.SplitCsv = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) Validate() error {
	return dara.Validate(s)
}

type GetOpenJMeterSceneResponseBodySceneRegionalCondition struct {
	// The number of load generators.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneRegionalCondition) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneRegionalCondition) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneRegionalCondition) GetAmount() *int32 {
	return s.Amount
}

func (s *GetOpenJMeterSceneResponseBodySceneRegionalCondition) GetRegion() *string {
	return s.Region
}

func (s *GetOpenJMeterSceneResponseBodySceneRegionalCondition) SetAmount(v int32) *GetOpenJMeterSceneResponseBodySceneRegionalCondition {
	s.Amount = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneRegionalCondition) SetRegion(v string) *GetOpenJMeterSceneResponseBodySceneRegionalCondition {
	s.Region = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneRegionalCondition) Validate() error {
	return dara.Validate(s)
}
