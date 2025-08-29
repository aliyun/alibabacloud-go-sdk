// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFeatureConsistencyCheckJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompareFeature(v bool) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetCompareFeature() *bool
	SetDatasetId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetDatasetId() *string
	SetDatasetMountPath(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetDatasetMountPath() *string
	SetDatasetName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetDatasetName() *string
	SetDatasetType(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetDatasetType() *string
	SetDatasetUri(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetDatasetUri() *string
	SetDefaultRoute(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetDefaultRoute() *string
	SetEasServiceName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetEasServiceName() *string
	SetEasyRecPackagePath(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetEasyRecPackagePath() *string
	SetEasyRecVersion(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetEasyRecVersion() *string
	SetFeatureDisplayExclude(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureDisplayExclude() *string
	SetFeatureLandingResourceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureLandingResourceId() *string
	SetFeaturePriority(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeaturePriority() *string
	SetFeatureStoreItemId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreItemId() *string
	SetFeatureStoreModelId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreModelId() *string
	SetFeatureStoreProjectId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreProjectId() *string
	SetFeatureStoreProjectName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreProjectName() *string
	SetFeatureStoreSeqFeatureView(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreSeqFeatureView() *string
	SetFeatureStoreUserId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFeatureStoreUserId() *string
	SetFgJarVersion(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFgJarVersion() *string
	SetFgJsonFileName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetFgJsonFileName() *string
	SetGenerateZip(v bool) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetGenerateZip() *bool
	SetInstanceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetInstanceId() *string
	SetIsUseFeatureStore(v bool) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetIsUseFeatureStore() *bool
	SetItemIdField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetItemIdField() *string
	SetItemTable(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetItemTable() *string
	SetItemTablePartitionField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetItemTablePartitionField() *string
	SetItemTablePartitionFieldFormat(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetItemTablePartitionFieldFormat() *string
	SetName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetName() *string
	SetOssResourceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetOssResourceId() *string
	SetPredictWorkerCount(v int32) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetPredictWorkerCount() *int32
	SetPredictWorkerCpu(v int32) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetPredictWorkerCpu() *int32
	SetPredictWorkerMemory(v int32) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetPredictWorkerMemory() *int32
	SetResourceConfig(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetResourceConfig() *string
	SetSampleRate(v float64) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetSampleRate() *float64
	SetSceneId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetSceneId() *string
	SetSecurityGroupId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetSecurityGroupId() *string
	SetServiceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetServiceId() *string
	SetSwitchId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetSwitchId() *string
	SetUserIdField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetUserIdField() *string
	SetUserTable(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetUserTable() *string
	SetUserTablePartitionField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetUserTablePartitionField() *string
	SetUserTablePartitionFieldFormat(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetUserTablePartitionFieldFormat() *string
	SetVpcId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetVpcId() *string
	SetWorkflowName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetWorkflowName() *string
	SetWorkspaceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest
	GetWorkspaceId() *string
}

type UpdateFeatureConsistencyCheckJobConfigRequest struct {
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
	// oss://********
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
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsUseFeatureStore *bool   `json:"IsUseFeatureStore,omitempty" xml:"IsUseFeatureStore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	ItemIdField *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_table
	ItemTable *string `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds
	ItemTablePartitionField *string `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	// This parameter is required.
	//
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
	// example:
	//
	// 4
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SwitchId  *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_id
	UserIdField *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_table
	UserTable *string `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds
	UserTablePartitionField *string `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	// This parameter is required.
	//
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

func (s UpdateFeatureConsistencyCheckJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetCompareFeature() *bool {
	return s.CompareFeature
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetDatasetMountPath() *string {
	return s.DatasetMountPath
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetDatasetUri() *string {
	return s.DatasetUri
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetEasyRecPackagePath() *string {
	return s.EasyRecPackagePath
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetEasyRecVersion() *string {
	return s.EasyRecVersion
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureDisplayExclude() *string {
	return s.FeatureDisplayExclude
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureLandingResourceId() *string {
	return s.FeatureLandingResourceId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeaturePriority() *string {
	return s.FeaturePriority
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreItemId() *string {
	return s.FeatureStoreItemId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreModelId() *string {
	return s.FeatureStoreModelId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreProjectId() *string {
	return s.FeatureStoreProjectId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreProjectName() *string {
	return s.FeatureStoreProjectName
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreSeqFeatureView() *string {
	return s.FeatureStoreSeqFeatureView
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFeatureStoreUserId() *string {
	return s.FeatureStoreUserId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFgJarVersion() *string {
	return s.FgJarVersion
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetFgJsonFileName() *string {
	return s.FgJsonFileName
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetGenerateZip() *bool {
	return s.GenerateZip
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetIsUseFeatureStore() *bool {
	return s.IsUseFeatureStore
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetItemTable() *string {
	return s.ItemTable
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetItemTablePartitionField() *string {
	return s.ItemTablePartitionField
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetItemTablePartitionFieldFormat() *string {
	return s.ItemTablePartitionFieldFormat
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetOssResourceId() *string {
	return s.OssResourceId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetPredictWorkerCount() *int32 {
	return s.PredictWorkerCount
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetPredictWorkerCpu() *int32 {
	return s.PredictWorkerCpu
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetPredictWorkerMemory() *int32 {
	return s.PredictWorkerMemory
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetResourceConfig() *string {
	return s.ResourceConfig
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetSampleRate() *float64 {
	return s.SampleRate
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetSwitchId() *string {
	return s.SwitchId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetUserIdField() *string {
	return s.UserIdField
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetUserTable() *string {
	return s.UserTable
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetUserTablePartitionField() *string {
	return s.UserTablePartitionField
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetUserTablePartitionFieldFormat() *string {
	return s.UserTablePartitionFieldFormat
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetCompareFeature(v bool) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.CompareFeature = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetDatasetId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetDatasetMountPath(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetMountPath = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetDatasetName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetDatasetType(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetType = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetDatasetUri(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.DatasetUri = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetDefaultRoute(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.DefaultRoute = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetEasServiceName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.EasServiceName = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetEasyRecPackagePath(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.EasyRecPackagePath = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetEasyRecVersion(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.EasyRecVersion = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureDisplayExclude(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureDisplayExclude = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureLandingResourceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureLandingResourceId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeaturePriority(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeaturePriority = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreItemId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreItemId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreModelId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreModelId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreProjectId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreProjectId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreProjectName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreProjectName = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreSeqFeatureView(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreSeqFeatureView = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFeatureStoreUserId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FeatureStoreUserId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFgJarVersion(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FgJarVersion = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetFgJsonFileName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.FgJsonFileName = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetGenerateZip(v bool) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.GenerateZip = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetInstanceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetIsUseFeatureStore(v bool) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.IsUseFeatureStore = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetItemIdField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ItemIdField = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetItemTable(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ItemTable = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetItemTablePartitionField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ItemTablePartitionField = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetItemTablePartitionFieldFormat(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ItemTablePartitionFieldFormat = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.Name = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetOssResourceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.OssResourceId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetPredictWorkerCount(v int32) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.PredictWorkerCount = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetPredictWorkerCpu(v int32) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.PredictWorkerCpu = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetPredictWorkerMemory(v int32) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.PredictWorkerMemory = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetResourceConfig(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ResourceConfig = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetSampleRate(v float64) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.SampleRate = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetSceneId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetSecurityGroupId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetServiceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetSwitchId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.SwitchId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetUserIdField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.UserIdField = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetUserTable(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.UserTable = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetUserTablePartitionField(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.UserTablePartitionField = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetUserTablePartitionFieldFormat(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.UserTablePartitionFieldFormat = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetVpcId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetWorkflowName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.WorkflowName = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetWorkspaceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
