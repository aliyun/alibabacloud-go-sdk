// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureConsistencyCheckConfigs(v []*ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) *ListFeatureConsistencyCheckJobConfigsResponseBody
	GetFeatureConsistencyCheckConfigs() []*ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs
	SetRequestId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListFeatureConsistencyCheckJobConfigsResponseBody
	GetTotalCount() *int64
}

type ListFeatureConsistencyCheckJobConfigsResponseBody struct {
	// The list of feature consistency check configurations.
	FeatureConsistencyCheckConfigs []*ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs `json:"FeatureConsistencyCheckConfigs,omitempty" xml:"FeatureConsistencyCheckConfigs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FCF741D8-9C30-578E-807F-B935487DB34A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of configurations.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) GetFeatureConsistencyCheckConfigs() []*ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	return s.FeatureConsistencyCheckConfigs
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) SetFeatureConsistencyCheckConfigs(v []*ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) *ListFeatureConsistencyCheckJobConfigsResponseBody {
	s.FeatureConsistencyCheckConfigs = v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) SetTotalCount(v int64) *ListFeatureConsistencyCheckJobConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBody) Validate() error {
	if s.FeatureConsistencyCheckConfigs != nil {
		for _, item := range s.FeatureConsistencyCheckConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs struct {
	// Indicates whether to enable feature comparison.
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
	// The name of the EAS service.
	//
	// example:
	//
	// eas_service_1
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// The path of the EasyRec package.
	//
	// example:
	//
	// oss://*******
	EasyRecPackagePath *string `json:"EasyRecPackagePath,omitempty" xml:"EasyRecPackagePath,omitempty"`
	// The version of EasyRec.
	//
	// example:
	//
	// 1.3.60
	EasyRecVersion *string `json:"EasyRecVersion,omitempty" xml:"EasyRecVersion,omitempty"`
	// The ID of the feature consistency check configuration.
	//
	// example:
	//
	// 3
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// The features to exclude from the results. Separate multiple features with a comma (,).
	//
	// example:
	//
	// feature1,feature2
	FeatureDisplayExclude *string `json:"FeatureDisplayExclude,omitempty" xml:"FeatureDisplayExclude,omitempty"`
	// The ID of the data source for feature landing.
	//
	// example:
	//
	// reso-********
	FeatureLandingResourceId *string `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
	// The URI of the data source for feature landing.
	//
	// example:
	//
	// mc_project_1
	FeatureLandingResourceUri *string `json:"FeatureLandingResourceUri,omitempty" xml:"FeatureLandingResourceUri,omitempty"`
	// The high-priority features to read from the user table. If a feature is not found, the system retrieves it from the behavior table. Separate multiple features with a comma (,).
	//
	// example:
	//
	// feature1,feature2,feature3
	FeaturePriority *string `json:"FeaturePriority,omitempty" xml:"FeaturePriority,omitempty"`
	// The primary key for the item side in the feature store.
	//
	// example:
	//
	// item
	FeatureStoreItemId *string `json:"FeatureStoreItemId,omitempty" xml:"FeatureStoreItemId,omitempty"`
	// The ID of the model in the feature store.
	//
	// example:
	//
	// 2
	FeatureStoreModelId *string `json:"FeatureStoreModelId,omitempty" xml:"FeatureStoreModelId,omitempty"`
	// The ID of the feature store project.
	//
	// example:
	//
	// prj-01
	FeatureStoreProjectId *string `json:"FeatureStoreProjectId,omitempty" xml:"FeatureStoreProjectId,omitempty"`
	// The name of the feature store project.
	//
	// example:
	//
	// project-1
	FeatureStoreProjectName *string `json:"FeatureStoreProjectName,omitempty" xml:"FeatureStoreProjectName,omitempty"`
	// The name of the feature view that contains item features within the sequence features.
	//
	// example:
	//
	// item-1
	FeatureStoreSeqFeatureView *string `json:"FeatureStoreSeqFeatureView,omitempty" xml:"FeatureStoreSeqFeatureView,omitempty"`
	// The primary key for the user side in the feature store.
	//
	// example:
	//
	// user
	FeatureStoreUserId *string `json:"FeatureStoreUserId,omitempty" xml:"FeatureStoreUserId,omitempty"`
	// The version of the `fg_jar` file.
	//
	// example:
	//
	// 1.0.0
	FgJarVersion *string `json:"FgJarVersion,omitempty" xml:"FgJarVersion,omitempty"`
	// The name of the `fg_json` file.
	//
	// example:
	//
	// file.json
	FgJsonFileName *string `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	// Indicates whether to generate a ZIP package.
	//
	// example:
	//
	// true
	GenerateZip *bool `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the configuration was last updated.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The name of the `item_id` field.
	//
	// example:
	//
	// item_id
	ItemIdField *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	// The name of the item table.
	//
	// example:
	//
	// item_table
	ItemTable *string `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	// The partition field of the item table.
	//
	// example:
	//
	// ds
	ItemTablePartitionField *string `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	// The format of the partition field of the item table. Valid values:
	//
	// - `yyyymmdd`
	//
	// - `yyyy-mm-dd`
	//
	// example:
	//
	// yyyymmdd
	ItemTablePartitionFieldFormat *string `json:"ItemTablePartitionFieldFormat,omitempty" xml:"ItemTablePartitionFieldFormat,omitempty"`
	// The end time of the latest job based on this configuration.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	LatestJobGmtSamplingEndTime *string `json:"LatestJobGmtSamplingEndTime,omitempty" xml:"LatestJobGmtSamplingEndTime,omitempty"`
	// The start time of the latest job based on this configuration.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	LatestJobGmtSamplingStartTime *string `json:"LatestJobGmtSamplingStartTime,omitempty" xml:"LatestJobGmtSamplingStartTime,omitempty"`
	// The ID of the most recent job created from this configuration.
	//
	// example:
	//
	// 3
	LatestJobId *string `json:"LatestJobId,omitempty" xml:"LatestJobId,omitempty"`
	// The name of the feature consistency check configuration.
	//
	// example:
	//
	// feature_consistency_check1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// oss_bucket_1
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The ID of the OSS data source.
	//
	// example:
	//
	// reso-********
	OssResourceId       *string `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	PredictWorkerCount  *int32  `json:"PredictWorkerCount,omitempty" xml:"PredictWorkerCount,omitempty"`
	PredictWorkerCpu    *int32  `json:"PredictWorkerCpu,omitempty" xml:"PredictWorkerCpu,omitempty"`
	PredictWorkerMemory *int32  `json:"PredictWorkerMemory,omitempty" xml:"PredictWorkerMemory,omitempty"`
	ResourceConfig      *string `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	// The sample rate, a value from 0 to 1.
	//
	// example:
	//
	// 0.89
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The ID of the scene.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The name of the scene.
	//
	// example:
	//
	// scene1
	SceneName       *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 4
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// service1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The status of the configuration. Valid values:
	//
	// - `Editable`: The configuration is editable.
	//
	// - `Uneditable`: The configuration is not editable.
	//
	// example:
	//
	// Editable
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// Indicates whether to use a feature store. Valid values:
	//
	// - `true`: A feature store is used. In this case, the response includes parameters such as `FeatureStoreProjectId`, `FeatureStoreProjectName`, `FeatureStoreModelId`, `FeatureStoreUserId`, and `FeatureStoreItemId`.
	//
	// - `false`: A feature store is not used. In this case, the response includes parameters such as `UserTable`, `UserIdField`, `UserTablePartitionField`, `UserTablePartitionFieldFormat`, `ItemTable`, `ItemIdField`, `ItemTablePartitionField`, and `ItemTablePartitionFieldFormat`.
	//
	// example:
	//
	// true
	UseFeatureStore *string `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
	// The name of the `user_id` field.
	//
	// example:
	//
	// user_id
	UserIdField *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	// The name of the user table.
	//
	// example:
	//
	// user_table
	UserTable *string `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	// The partition field of the user table.
	//
	// example:
	//
	// ds
	UserTablePartitionField *string `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	// The format of the partition field of the user table. Valid values:
	//
	// - `yyyymmdd`
	//
	// - `yyyy-mm-dd`
	//
	// example:
	//
	// yyyymmdd
	UserTablePartitionFieldFormat *string `json:"UserTablePartitionFieldFormat,omitempty" xml:"UserTablePartitionFieldFormat,omitempty"`
	VpcId                         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// work_flow_1
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetCompareFeature() *bool {
	return s.CompareFeature
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetDatasetMountPath() *string {
	return s.DatasetMountPath
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetDatasetType() *string {
	return s.DatasetType
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetDatasetUri() *string {
	return s.DatasetUri
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetEasyRecPackagePath() *string {
	return s.EasyRecPackagePath
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetEasyRecVersion() *string {
	return s.EasyRecVersion
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureDisplayExclude() *string {
	return s.FeatureDisplayExclude
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureLandingResourceId() *string {
	return s.FeatureLandingResourceId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureLandingResourceUri() *string {
	return s.FeatureLandingResourceUri
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeaturePriority() *string {
	return s.FeaturePriority
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureStoreItemId() *string {
	return s.FeatureStoreItemId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureStoreModelId() *string {
	return s.FeatureStoreModelId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureStoreProjectId() *string {
	return s.FeatureStoreProjectId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureStoreProjectName() *string {
	return s.FeatureStoreProjectName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureStoreSeqFeatureView() *string {
	return s.FeatureStoreSeqFeatureView
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFeatureStoreUserId() *string {
	return s.FeatureStoreUserId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFgJarVersion() *string {
	return s.FgJarVersion
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetFgJsonFileName() *string {
	return s.FgJsonFileName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetGenerateZip() *bool {
	return s.GenerateZip
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetItemTable() *string {
	return s.ItemTable
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetItemTablePartitionField() *string {
	return s.ItemTablePartitionField
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetItemTablePartitionFieldFormat() *string {
	return s.ItemTablePartitionFieldFormat
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetLatestJobGmtSamplingEndTime() *string {
	return s.LatestJobGmtSamplingEndTime
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetLatestJobGmtSamplingStartTime() *string {
	return s.LatestJobGmtSamplingStartTime
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetLatestJobId() *string {
	return s.LatestJobId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetName() *string {
	return s.Name
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetOssBucket() *string {
	return s.OssBucket
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetOssResourceId() *string {
	return s.OssResourceId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetPredictWorkerCount() *int32 {
	return s.PredictWorkerCount
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetPredictWorkerCpu() *int32 {
	return s.PredictWorkerCpu
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetPredictWorkerMemory() *int32 {
	return s.PredictWorkerMemory
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetResourceConfig() *string {
	return s.ResourceConfig
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetSampleRate() *string {
	return s.SampleRate
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetSceneId() *string {
	return s.SceneId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetSceneName() *string {
	return s.SceneName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetStatus() *string {
	return s.Status
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetUseFeatureStore() *string {
	return s.UseFeatureStore
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetUserIdField() *string {
	return s.UserIdField
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetUserTable() *string {
	return s.UserTable
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetUserTablePartitionField() *string {
	return s.UserTablePartitionField
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetUserTablePartitionFieldFormat() *string {
	return s.UserTablePartitionFieldFormat
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetVpcId() *string {
	return s.VpcId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetCompareFeature(v bool) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.CompareFeature = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetDatasetId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.DatasetId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetDatasetMountPath(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.DatasetMountPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetDatasetName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.DatasetName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetDatasetType(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.DatasetType = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetDatasetUri(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.DatasetUri = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetDefaultRoute(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.DefaultRoute = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetEasServiceName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.EasServiceName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetEasyRecPackagePath(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.EasyRecPackagePath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetEasyRecVersion(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.EasyRecVersion = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureConsistencyCheckJobConfigId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureDisplayExclude(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureDisplayExclude = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureLandingResourceId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureLandingResourceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureLandingResourceUri(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureLandingResourceUri = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeaturePriority(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeaturePriority = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureStoreItemId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureStoreItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureStoreModelId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureStoreModelId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureStoreProjectId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureStoreProjectId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureStoreProjectName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureStoreProjectName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureStoreSeqFeatureView(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureStoreSeqFeatureView = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFeatureStoreUserId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FeatureStoreUserId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFgJarVersion(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FgJarVersion = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetFgJsonFileName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.FgJsonFileName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetGenerateZip(v bool) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.GenerateZip = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetGmtCreateTime(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetGmtModifiedTime(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetItemIdField(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ItemIdField = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetItemTable(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ItemTable = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetItemTablePartitionField(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ItemTablePartitionField = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetItemTablePartitionFieldFormat(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ItemTablePartitionFieldFormat = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetLatestJobGmtSamplingEndTime(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.LatestJobGmtSamplingEndTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetLatestJobGmtSamplingStartTime(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.LatestJobGmtSamplingStartTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetLatestJobId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.LatestJobId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.Name = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetOssBucket(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.OssBucket = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetOssResourceId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.OssResourceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetPredictWorkerCount(v int32) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.PredictWorkerCount = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetPredictWorkerCpu(v int32) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.PredictWorkerCpu = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetPredictWorkerMemory(v int32) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.PredictWorkerMemory = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetResourceConfig(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ResourceConfig = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetSampleRate(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.SampleRate = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetSceneId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.SceneId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetSceneName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.SceneName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetSecurityGroupId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.SecurityGroupId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetServiceId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ServiceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetServiceName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.ServiceName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetStatus(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.Status = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetSwitchId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.SwitchId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetUseFeatureStore(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.UseFeatureStore = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetUserIdField(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.UserIdField = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetUserTable(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.UserTable = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetUserTablePartitionField(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.UserTablePartitionField = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetUserTablePartitionFieldFormat(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.UserTablePartitionFieldFormat = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetVpcId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.VpcId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetWorkflowName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.WorkflowName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetWorkspaceId(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.WorkspaceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) Validate() error {
	return dara.Validate(s)
}
