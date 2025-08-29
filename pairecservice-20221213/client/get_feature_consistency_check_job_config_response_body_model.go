// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConsistencyCheckJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompareFeature(v bool) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetCompareFeature() *bool
	SetDatasetId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetDatasetId() *string
	SetDatasetMountPath(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetDatasetMountPath() *string
	SetDatasetName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetDatasetName() *string
	SetDatasetType(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetDatasetType() *string
	SetDatasetUri(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetDatasetUri() *string
	SetDefaultRoute(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetDefaultRoute() *string
	SetEasServiceName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetEasServiceName() *string
	SetEasyRecPackagePath(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetEasyRecPackagePath() *string
	SetEasyRecVersion(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetEasyRecVersion() *string
	SetFeatureDisplayExclude(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureDisplayExclude() *string
	SetFeatureLandingResourceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureLandingResourceId() *string
	SetFeatureLandingResourceUri(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureLandingResourceUri() *string
	SetFeaturePriority(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeaturePriority() *string
	SetFeatureStoreItemId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureStoreItemId() *string
	SetFeatureStoreModelId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureStoreModelId() *string
	SetFeatureStoreProjectId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureStoreProjectId() *string
	SetFeatureStoreProjectName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureStoreProjectName() *string
	SetFeatureStoreSeqFeatureView(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureStoreSeqFeatureView() *string
	SetFeatureStoreUserId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureStoreUserId() *string
	SetFgJarVersion(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFgJarVersion() *string
	SetFgJsonFileName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetFgJsonFileName() *string
	SetGenerateZip(v bool) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetGenerateZip() *bool
	SetGmtCreateTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetGmtModifiedTime() *string
	SetItemIdField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetItemIdField() *string
	SetItemTable(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetItemTable() *string
	SetItemTablePartitionField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetItemTablePartitionField() *string
	SetItemTablePartitionFieldFormat(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetItemTablePartitionFieldFormat() *string
	SetLatestJobGmtSamplingEndTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetLatestJobGmtSamplingEndTime() *string
	SetLatestJobGmtSamplingStartTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetLatestJobGmtSamplingStartTime() *string
	SetLatestJobId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetLatestJobId() *string
	SetName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetName() *string
	SetOssBucket(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetOssBucket() *string
	SetOssResourceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetOssResourceId() *string
	SetPredictWorkerCount(v int32) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetPredictWorkerCount() *int32
	SetPredictWorkerCpu(v int32) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetPredictWorkerCpu() *int32
	SetPredictWorkerMemory(v int32) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetPredictWorkerMemory() *int32
	SetRequestId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetRequestId() *string
	SetResourceConfig(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetResourceConfig() *string
	SetSampleRate(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetSampleRate() *string
	SetSceneId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetSceneId() *string
	SetSceneName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetSceneName() *string
	SetSecurityGroupId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetSecurityGroupId() *string
	SetServiceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetServiceId() *string
	SetServiceName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetServiceName() *string
	SetStatus(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetStatus() *string
	SetSwitchId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetSwitchId() *string
	SetUseFeatureStore(v bool) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetUseFeatureStore() *bool
	SetUserIdField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetUserIdField() *string
	SetUserTable(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetUserTable() *string
	SetUserTablePartitionField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetUserTablePartitionField() *string
	SetUserTablePartitionFieldFormat(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetUserTablePartitionFieldFormat() *string
	SetVpcId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetVpcId() *string
	SetWorkflowName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetWorkflowName() *string
	SetWorkspaceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody
	GetWorkspaceId() *string
}

type GetFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// true
	CompareFeature   *bool   `json:"CompareFeature,omitempty" xml:"CompareFeature,omitempty"`
	DatasetId        *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetMountPath *string `json:"DatasetMountPath,omitempty" xml:"DatasetMountPath,omitempty"`
	DatasetName      *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DatasetType      *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	DatasetUri       *string `json:"DatasetUri,omitempty" xml:"DatasetUri,omitempty"`
	DefaultRoute     *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// example:
	//
	// eas_service_1
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// example:
	//
	// oss://*******
	EasyRecPackagePath *string `json:"EasyRecPackagePath,omitempty" xml:"EasyRecPackagePath,omitempty"`
	// example:
	//
	// 1.3.60
	EasyRecVersion *string `json:"EasyRecVersion,omitempty" xml:"EasyRecVersion,omitempty"`
	// example:
	//
	// feature1,feature2
	FeatureDisplayExclude *string `json:"FeatureDisplayExclude,omitempty" xml:"FeatureDisplayExclude,omitempty"`
	// example:
	//
	// reso-********
	FeatureLandingResourceId *string `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
	// example:
	//
	// mc_project_1
	FeatureLandingResourceUri *string `json:"FeatureLandingResourceUri,omitempty" xml:"FeatureLandingResourceUri,omitempty"`
	// example:
	//
	// feature1,feature2,feature3
	FeaturePriority            *string `json:"FeaturePriority,omitempty" xml:"FeaturePriority,omitempty"`
	FeatureStoreItemId         *string `json:"FeatureStoreItemId,omitempty" xml:"FeatureStoreItemId,omitempty"`
	FeatureStoreModelId        *string `json:"FeatureStoreModelId,omitempty" xml:"FeatureStoreModelId,omitempty"`
	FeatureStoreProjectId      *string `json:"FeatureStoreProjectId,omitempty" xml:"FeatureStoreProjectId,omitempty"`
	FeatureStoreProjectName    *string `json:"FeatureStoreProjectName,omitempty" xml:"FeatureStoreProjectName,omitempty"`
	FeatureStoreSeqFeatureView *string `json:"FeatureStoreSeqFeatureView,omitempty" xml:"FeatureStoreSeqFeatureView,omitempty"`
	FeatureStoreUserId         *string `json:"FeatureStoreUserId,omitempty" xml:"FeatureStoreUserId,omitempty"`
	// example:
	//
	// 1.0.0
	FgJarVersion *string `json:"FgJarVersion,omitempty" xml:"FgJarVersion,omitempty"`
	// example:
	//
	// file.json
	FgJsonFileName *string `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	// example:
	//
	// true
	GenerateZip *bool `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// item_id
	ItemIdField *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	// example:
	//
	// item_table
	ItemTable *string `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	// example:
	//
	// ds
	ItemTablePartitionField *string `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	// example:
	//
	// yyyymmdd
	ItemTablePartitionFieldFormat *string `json:"ItemTablePartitionFieldFormat,omitempty" xml:"ItemTablePartitionFieldFormat,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	LatestJobGmtSamplingEndTime *string `json:"LatestJobGmtSamplingEndTime,omitempty" xml:"LatestJobGmtSamplingEndTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	LatestJobGmtSamplingStartTime *string `json:"LatestJobGmtSamplingStartTime,omitempty" xml:"LatestJobGmtSamplingStartTime,omitempty"`
	// example:
	//
	// 3
	LatestJobId *string `json:"LatestJobId,omitempty" xml:"LatestJobId,omitempty"`
	// example:
	//
	// feature_consistency_check1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss_bucket_1
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// example:
	//
	// reso-********
	OssResourceId       *string `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	PredictWorkerCount  *int32  `json:"PredictWorkerCount,omitempty" xml:"PredictWorkerCount,omitempty"`
	PredictWorkerCpu    *int32  `json:"PredictWorkerCpu,omitempty" xml:"PredictWorkerCpu,omitempty"`
	PredictWorkerMemory *int32  `json:"PredictWorkerMemory,omitempty" xml:"PredictWorkerMemory,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceConfig *string `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	// example:
	//
	// 0.89
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// scene1
	SceneName       *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// 4
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// service1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// Editable
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SwitchId        *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	UseFeatureStore *bool   `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
	// example:
	//
	// user_id
	UserIdField *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	// example:
	//
	// user_table
	UserTable *string `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	// example:
	//
	// ds
	UserTablePartitionField *string `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	// example:
	//
	// yyyymmdd
	UserTablePartitionFieldFormat *string `json:"UserTablePartitionFieldFormat,omitempty" xml:"UserTablePartitionFieldFormat,omitempty"`
	VpcId                         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// work_flow_1
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetCompareFeature() *bool {
	return s.CompareFeature
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetDatasetMountPath() *string {
	return s.DatasetMountPath
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetDatasetType() *string {
	return s.DatasetType
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetDatasetUri() *string {
	return s.DatasetUri
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetEasyRecPackagePath() *string {
	return s.EasyRecPackagePath
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetEasyRecVersion() *string {
	return s.EasyRecVersion
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureDisplayExclude() *string {
	return s.FeatureDisplayExclude
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureLandingResourceId() *string {
	return s.FeatureLandingResourceId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureLandingResourceUri() *string {
	return s.FeatureLandingResourceUri
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeaturePriority() *string {
	return s.FeaturePriority
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureStoreItemId() *string {
	return s.FeatureStoreItemId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureStoreModelId() *string {
	return s.FeatureStoreModelId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureStoreProjectId() *string {
	return s.FeatureStoreProjectId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureStoreProjectName() *string {
	return s.FeatureStoreProjectName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureStoreSeqFeatureView() *string {
	return s.FeatureStoreSeqFeatureView
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFeatureStoreUserId() *string {
	return s.FeatureStoreUserId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFgJarVersion() *string {
	return s.FgJarVersion
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetFgJsonFileName() *string {
	return s.FgJsonFileName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetGenerateZip() *bool {
	return s.GenerateZip
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetItemTable() *string {
	return s.ItemTable
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetItemTablePartitionField() *string {
	return s.ItemTablePartitionField
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetItemTablePartitionFieldFormat() *string {
	return s.ItemTablePartitionFieldFormat
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetLatestJobGmtSamplingEndTime() *string {
	return s.LatestJobGmtSamplingEndTime
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetLatestJobGmtSamplingStartTime() *string {
	return s.LatestJobGmtSamplingStartTime
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetLatestJobId() *string {
	return s.LatestJobId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetName() *string {
	return s.Name
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetOssBucket() *string {
	return s.OssBucket
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetOssResourceId() *string {
	return s.OssResourceId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetPredictWorkerCount() *int32 {
	return s.PredictWorkerCount
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetPredictWorkerCpu() *int32 {
	return s.PredictWorkerCpu
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetPredictWorkerMemory() *int32 {
	return s.PredictWorkerMemory
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetResourceConfig() *string {
	return s.ResourceConfig
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetSceneName() *string {
	return s.SceneName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetSwitchId() *string {
	return s.SwitchId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetUseFeatureStore() *bool {
	return s.UseFeatureStore
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetUserIdField() *string {
	return s.UserIdField
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetUserTable() *string {
	return s.UserTable
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetUserTablePartitionField() *string {
	return s.UserTablePartitionField
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetUserTablePartitionFieldFormat() *string {
	return s.UserTablePartitionFieldFormat
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetCompareFeature(v bool) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.CompareFeature = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetDatasetId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetDatasetMountPath(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.DatasetMountPath = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetDatasetName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetDatasetType(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.DatasetType = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetDatasetUri(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.DatasetUri = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetDefaultRoute(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.DefaultRoute = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetEasServiceName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.EasServiceName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetEasyRecPackagePath(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.EasyRecPackagePath = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetEasyRecVersion(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.EasyRecVersion = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureDisplayExclude(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureDisplayExclude = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureLandingResourceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureLandingResourceId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureLandingResourceUri(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureLandingResourceUri = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeaturePriority(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeaturePriority = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureStoreItemId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureStoreItemId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureStoreModelId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureStoreModelId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureStoreProjectId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureStoreProjectId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureStoreProjectName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureStoreProjectName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureStoreSeqFeatureView(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureStoreSeqFeatureView = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFeatureStoreUserId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureStoreUserId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFgJarVersion(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FgJarVersion = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetFgJsonFileName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.FgJsonFileName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetGenerateZip(v bool) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.GenerateZip = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetGmtCreateTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetGmtModifiedTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetItemIdField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ItemIdField = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetItemTable(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ItemTable = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetItemTablePartitionField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ItemTablePartitionField = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetItemTablePartitionFieldFormat(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ItemTablePartitionFieldFormat = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetLatestJobGmtSamplingEndTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.LatestJobGmtSamplingEndTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetLatestJobGmtSamplingStartTime(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.LatestJobGmtSamplingStartTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetLatestJobId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.LatestJobId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.Name = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetOssBucket(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.OssBucket = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetOssResourceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.OssResourceId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetPredictWorkerCount(v int32) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.PredictWorkerCount = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetPredictWorkerCpu(v int32) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.PredictWorkerCpu = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetPredictWorkerMemory(v int32) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.PredictWorkerMemory = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetResourceConfig(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ResourceConfig = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetSampleRate(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.SampleRate = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetSceneId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetSceneName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetSecurityGroupId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetServiceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetServiceName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetStatus(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.Status = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetSwitchId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.SwitchId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetUseFeatureStore(v bool) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.UseFeatureStore = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetUserIdField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.UserIdField = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetUserTable(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.UserTable = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetUserTablePartitionField(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.UserTablePartitionField = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetUserTablePartitionFieldFormat(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.UserTablePartitionFieldFormat = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetVpcId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetWorkflowName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.WorkflowName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetWorkspaceId(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
