// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureConsistencyCheckJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompareFeature(v bool) *CreateFeatureConsistencyCheckJobConfigRequest
	GetCompareFeature() *bool
	SetDatasetId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetDatasetId() *string
	SetDatasetMountPath(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetDatasetMountPath() *string
	SetDatasetName(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetDatasetName() *string
	SetDatasetType(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetDatasetType() *string
	SetDatasetUri(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetDatasetUri() *string
	SetDefaultRoute(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetDefaultRoute() *string
	SetEasServiceName(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetEasServiceName() *string
	SetEasyRecPackagePath(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetEasyRecPackagePath() *string
	SetEasyRecVersion(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetEasyRecVersion() *string
	SetFeatureDisplayExclude(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureDisplayExclude() *string
	SetFeatureLandingResourceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureLandingResourceId() *string
	SetFeaturePriority(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeaturePriority() *string
	SetFeatureStoreItemId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreItemId() *string
	SetFeatureStoreModelId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreModelId() *string
	SetFeatureStoreProjectId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreProjectId() *string
	SetFeatureStoreProjectName(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreProjectName() *string
	SetFeatureStoreSeqFeatureView(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreSeqFeatureView() *string
	SetFeatureStoreUserId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreUserId() *string
	SetFgJarVersion(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFgJarVersion() *string
	SetFgJsonFileName(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetFgJsonFileName() *string
	SetGenerateZip(v bool) *CreateFeatureConsistencyCheckJobConfigRequest
	GetGenerateZip() *bool
	SetInstanceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetInstanceId() *string
	SetItemIdField(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetItemIdField() *string
	SetItemTable(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetItemTable() *string
	SetItemTablePartitionField(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetItemTablePartitionField() *string
	SetItemTablePartitionFieldFormat(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetItemTablePartitionFieldFormat() *string
	SetName(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetName() *string
	SetOssResourceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetOssResourceId() *string
	SetPredictWorkerCount(v int32) *CreateFeatureConsistencyCheckJobConfigRequest
	GetPredictWorkerCount() *int32
	SetPredictWorkerCpu(v int32) *CreateFeatureConsistencyCheckJobConfigRequest
	GetPredictWorkerCpu() *int32
	SetPredictWorkerMemory(v int32) *CreateFeatureConsistencyCheckJobConfigRequest
	GetPredictWorkerMemory() *int32
	SetResourceConfig(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetResourceConfig() *string
	SetSampleRate(v float64) *CreateFeatureConsistencyCheckJobConfigRequest
	GetSampleRate() *float64
	SetSceneId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetSceneId() *string
	SetSecurityGroupId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetSecurityGroupId() *string
	SetServiceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetServiceId() *string
	SetSwitchId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetSwitchId() *string
	SetUseFeatureStore(v bool) *CreateFeatureConsistencyCheckJobConfigRequest
	GetUseFeatureStore() *bool
	SetUserIdField(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetUserIdField() *string
	SetUserTable(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetUserTable() *string
	SetUserTablePartitionField(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetUserTablePartitionField() *string
	SetUserTablePartitionFieldFormat(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetUserTablePartitionFieldFormat() *string
	SetVpcId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetVpcId() *string
	SetWorkflowName(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetWorkflowName() *string
	SetWorkspaceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest
	GetWorkspaceId() *string
}

type CreateFeatureConsistencyCheckJobConfigRequest struct {
	// This parameter is required.
	//
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
	// This parameter is required.
	//
	// example:
	//
	// service_123
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
	// This parameter is required.
	//
	// example:
	//
	// reso-********
	FeatureLandingResourceId *string `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// file.json
	FgJsonFileName *string `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	GenerateZip *bool `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// feature_consistency_check1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reso-********
	OssResourceId       *string `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	PredictWorkerCount  *int32  `json:"PredictWorkerCount,omitempty" xml:"PredictWorkerCount,omitempty"`
	PredictWorkerCpu    *int32  `json:"PredictWorkerCpu,omitempty" xml:"PredictWorkerCpu,omitempty"`
	PredictWorkerMemory *int32  `json:"PredictWorkerMemory,omitempty" xml:"PredictWorkerMemory,omitempty"`
	ResourceConfig      *string `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.89
	SampleRate *float64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId         *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SwitchId  *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// This parameter is required.
	UseFeatureStore *bool `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
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

func (s CreateFeatureConsistencyCheckJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetCompareFeature() *bool {
	return s.CompareFeature
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetDatasetMountPath() *string {
	return s.DatasetMountPath
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetDatasetUri() *string {
	return s.DatasetUri
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetEasyRecPackagePath() *string {
	return s.EasyRecPackagePath
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetEasyRecVersion() *string {
	return s.EasyRecVersion
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureDisplayExclude() *string {
	return s.FeatureDisplayExclude
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureLandingResourceId() *string {
	return s.FeatureLandingResourceId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeaturePriority() *string {
	return s.FeaturePriority
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreItemId() *string {
	return s.FeatureStoreItemId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreModelId() *string {
	return s.FeatureStoreModelId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreProjectId() *string {
	return s.FeatureStoreProjectId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreProjectName() *string {
	return s.FeatureStoreProjectName
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreSeqFeatureView() *string {
	return s.FeatureStoreSeqFeatureView
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreUserId() *string {
	return s.FeatureStoreUserId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFgJarVersion() *string {
	return s.FgJarVersion
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetFgJsonFileName() *string {
	return s.FgJsonFileName
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetGenerateZip() *bool {
	return s.GenerateZip
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetItemTable() *string {
	return s.ItemTable
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetItemTablePartitionField() *string {
	return s.ItemTablePartitionField
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetItemTablePartitionFieldFormat() *string {
	return s.ItemTablePartitionFieldFormat
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetName() *string {
	return s.Name
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetOssResourceId() *string {
	return s.OssResourceId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetPredictWorkerCount() *int32 {
	return s.PredictWorkerCount
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetPredictWorkerCpu() *int32 {
	return s.PredictWorkerCpu
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetPredictWorkerMemory() *int32 {
	return s.PredictWorkerMemory
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetResourceConfig() *string {
	return s.ResourceConfig
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetSampleRate() *float64 {
	return s.SampleRate
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetSwitchId() *string {
	return s.SwitchId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetUseFeatureStore() *bool {
	return s.UseFeatureStore
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetUserIdField() *string {
	return s.UserIdField
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetUserTable() *string {
	return s.UserTable
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetUserTablePartitionField() *string {
	return s.UserTablePartitionField
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetUserTablePartitionFieldFormat() *string {
	return s.UserTablePartitionFieldFormat
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetCompareFeature(v bool) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.CompareFeature = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetDatasetId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetDatasetMountPath(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetMountPath = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetDatasetName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetDatasetType(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetType = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetDatasetUri(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetUri = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetDefaultRoute(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.DefaultRoute = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetEasServiceName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.EasServiceName = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetEasyRecPackagePath(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.EasyRecPackagePath = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetEasyRecVersion(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.EasyRecVersion = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureDisplayExclude(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureDisplayExclude = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureLandingResourceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureLandingResourceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeaturePriority(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeaturePriority = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreItemId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreItemId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreModelId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreModelId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreProjectId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreProjectId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreProjectName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreProjectName = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreSeqFeatureView(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreSeqFeatureView = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreUserId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreUserId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFgJarVersion(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FgJarVersion = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetFgJsonFileName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.FgJsonFileName = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetGenerateZip(v bool) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.GenerateZip = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetInstanceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetItemIdField(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ItemIdField = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetItemTable(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ItemTable = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetItemTablePartitionField(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ItemTablePartitionField = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetItemTablePartitionFieldFormat(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ItemTablePartitionFieldFormat = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.Name = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetOssResourceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.OssResourceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetPredictWorkerCount(v int32) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.PredictWorkerCount = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetPredictWorkerCpu(v int32) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.PredictWorkerCpu = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetPredictWorkerMemory(v int32) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.PredictWorkerMemory = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetResourceConfig(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ResourceConfig = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetSampleRate(v float64) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.SampleRate = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetSceneId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.SceneId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetSecurityGroupId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetServiceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetSwitchId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.SwitchId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetUseFeatureStore(v bool) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.UseFeatureStore = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetUserIdField(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.UserIdField = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetUserTable(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.UserTable = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetUserTablePartitionField(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.UserTablePartitionField = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetUserTablePartitionFieldFormat(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.UserTablePartitionFieldFormat = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetVpcId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.VpcId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetWorkflowName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.WorkflowName = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetWorkspaceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
