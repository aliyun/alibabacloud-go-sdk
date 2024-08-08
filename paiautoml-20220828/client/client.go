// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AutofeExperimentConfiguration struct {
	// This parameter is required.
	OdpsConfig *AutofeExperimentConfigurationOdpsConfig `json:"odps_config,omitempty" xml:"odps_config,omitempty" type:"Struct"`
	// This parameter is required.
	OssConfig *AutofeExperimentConfigurationOssConfig `json:"oss_config,omitempty" xml:"oss_config,omitempty" type:"Struct"`
	// This parameter is required.
	YmlConfig *AutofeExperimentConfigurationYmlConfig `json:"yml_config,omitempty" xml:"yml_config,omitempty" type:"Struct"`
}

func (s AutofeExperimentConfiguration) String() string {
	return tea.Prettify(s)
}

func (s AutofeExperimentConfiguration) GoString() string {
	return s.String()
}

func (s *AutofeExperimentConfiguration) SetOdpsConfig(v *AutofeExperimentConfigurationOdpsConfig) *AutofeExperimentConfiguration {
	s.OdpsConfig = v
	return s
}

func (s *AutofeExperimentConfiguration) SetOssConfig(v *AutofeExperimentConfigurationOssConfig) *AutofeExperimentConfiguration {
	s.OssConfig = v
	return s
}

func (s *AutofeExperimentConfiguration) SetYmlConfig(v *AutofeExperimentConfigurationYmlConfig) *AutofeExperimentConfiguration {
	s.YmlConfig = v
	return s
}

type AutofeExperimentConfigurationOdpsConfig struct {
	OdpsAccessId  *string `json:"odps_access_id,omitempty" xml:"odps_access_id,omitempty"`
	OdpsAccessKey *string `json:"odps_access_key,omitempty" xml:"odps_access_key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://service.cn-beijing.maxcompute.aliyun.com/api
	OdpsEndpoint *string `json:"odps_endpoint,omitempty" xml:"odps_endpoint,omitempty"`
	// This parameter is required.
	OdpsProjectName *string `json:"odps_project_name,omitempty" xml:"odps_project_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OdpsRegionId *string `json:"odps_region_id,omitempty" xml:"odps_region_id,omitempty"`
	OdpsRoleArn  *string `json:"odps_role_arn,omitempty" xml:"odps_role_arn,omitempty"`
}

func (s AutofeExperimentConfigurationOdpsConfig) String() string {
	return tea.Prettify(s)
}

func (s AutofeExperimentConfigurationOdpsConfig) GoString() string {
	return s.String()
}

func (s *AutofeExperimentConfigurationOdpsConfig) SetOdpsAccessId(v string) *AutofeExperimentConfigurationOdpsConfig {
	s.OdpsAccessId = &v
	return s
}

func (s *AutofeExperimentConfigurationOdpsConfig) SetOdpsAccessKey(v string) *AutofeExperimentConfigurationOdpsConfig {
	s.OdpsAccessKey = &v
	return s
}

func (s *AutofeExperimentConfigurationOdpsConfig) SetOdpsEndpoint(v string) *AutofeExperimentConfigurationOdpsConfig {
	s.OdpsEndpoint = &v
	return s
}

func (s *AutofeExperimentConfigurationOdpsConfig) SetOdpsProjectName(v string) *AutofeExperimentConfigurationOdpsConfig {
	s.OdpsProjectName = &v
	return s
}

func (s *AutofeExperimentConfigurationOdpsConfig) SetOdpsRegionId(v string) *AutofeExperimentConfigurationOdpsConfig {
	s.OdpsRegionId = &v
	return s
}

func (s *AutofeExperimentConfigurationOdpsConfig) SetOdpsRoleArn(v string) *AutofeExperimentConfigurationOdpsConfig {
	s.OdpsRoleArn = &v
	return s
}

type AutofeExperimentConfigurationOssConfig struct {
	OssAccessId  *string `json:"oss_access_id,omitempty" xml:"oss_access_id,omitempty"`
	OssAccessKey *string `json:"oss_access_key,omitempty" xml:"oss_access_key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// autofe-service
	OssBucket *string `json:"oss_bucket,omitempty" xml:"oss_bucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-beijing.aliyuncs.com
	OssEndpoint *string `json:"oss_endpoint,omitempty" xml:"oss_endpoint,omitempty"`
	OssRoleArn  *string `json:"oss_role_arn,omitempty" xml:"oss_role_arn,omitempty"`
}

func (s AutofeExperimentConfigurationOssConfig) String() string {
	return tea.Prettify(s)
}

func (s AutofeExperimentConfigurationOssConfig) GoString() string {
	return s.String()
}

func (s *AutofeExperimentConfigurationOssConfig) SetOssAccessId(v string) *AutofeExperimentConfigurationOssConfig {
	s.OssAccessId = &v
	return s
}

func (s *AutofeExperimentConfigurationOssConfig) SetOssAccessKey(v string) *AutofeExperimentConfigurationOssConfig {
	s.OssAccessKey = &v
	return s
}

func (s *AutofeExperimentConfigurationOssConfig) SetOssBucket(v string) *AutofeExperimentConfigurationOssConfig {
	s.OssBucket = &v
	return s
}

func (s *AutofeExperimentConfigurationOssConfig) SetOssEndpoint(v string) *AutofeExperimentConfigurationOssConfig {
	s.OssEndpoint = &v
	return s
}

func (s *AutofeExperimentConfigurationOssConfig) SetOssRoleArn(v string) *AutofeExperimentConfigurationOssConfig {
	s.OssRoleArn = &v
	return s
}

type AutofeExperimentConfigurationYmlConfig struct {
	// example:
	//
	// train
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// true
	AggregateOnly *string `json:"aggregate_only,omitempty" xml:"aggregate_only,omitempty"`
	AnalyzeExpId  *string `json:"analyze_exp_id,omitempty" xml:"analyze_exp_id,omitempty"`
	// example:
	//
	// 2
	Cpu           *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	DataPartition *string `json:"data_partition,omitempty" xml:"data_partition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// odps_table_name
	DataSource *string `json:"data_source,omitempty" xml:"data_source,omitempty"`
	// example:
	//
	// odps
	DataType *string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// example:
	//
	// false
	DebugMode        *string `json:"debug_mode,omitempty" xml:"debug_mode,omitempty"`
	ExcludeColumns   *string `json:"exclude_columns,omitempty" xml:"exclude_columns,omitempty"`
	FeatureSelection *string `json:"feature_selection,omitempty" xml:"feature_selection,omitempty"`
	// example:
	//
	// 10
	FilterThresh *string `json:"filter_thresh,omitempty" xml:"filter_thresh,omitempty"`
	// example:
	//
	// 0.02
	IvThresh *string `json:"iv_thresh,omitempty" xml:"iv_thresh,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// label
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// 4000
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket_name/
	OutputConfigOssDir *string `json:"output_config_oss_dir,omitempty" xml:"output_config_oss_dir,omitempty"`
	PipelineExpId      *string `json:"pipeline_exp_id,omitempty" xml:"pipeline_exp_id,omitempty"`
	// example:
	//
	// false
	ReuseResults *string `json:"reuse_results,omitempty" xml:"reuse_results,omitempty"`
	// example:
	//
	// 0
	SampleRatio *string `json:"sample_ratio,omitempty" xml:"sample_ratio,omitempty"`
	// example:
	//
	// 5000000
	SampleSize     *string `json:"sample_size,omitempty" xml:"sample_size,omitempty"`
	SelectionExpId *string `json:"selection_exp_id,omitempty" xml:"selection_exp_id,omitempty"`
	// example:
	//
	// true
	SkipSelect *string `json:"skip_select,omitempty" xml:"skip_select,omitempty"`
	// example:
	//
	// 10
	Workers *string `json:"workers,omitempty" xml:"workers,omitempty"`
	// This parameter is required.
	WorkspaceName *string `json:"workspace_name,omitempty" xml:"workspace_name,omitempty"`
}

func (s AutofeExperimentConfigurationYmlConfig) String() string {
	return tea.Prettify(s)
}

func (s AutofeExperimentConfigurationYmlConfig) GoString() string {
	return s.String()
}

func (s *AutofeExperimentConfigurationYmlConfig) SetAction(v string) *AutofeExperimentConfigurationYmlConfig {
	s.Action = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetAggregateOnly(v string) *AutofeExperimentConfigurationYmlConfig {
	s.AggregateOnly = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetAnalyzeExpId(v string) *AutofeExperimentConfigurationYmlConfig {
	s.AnalyzeExpId = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetCpu(v string) *AutofeExperimentConfigurationYmlConfig {
	s.Cpu = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetDataPartition(v string) *AutofeExperimentConfigurationYmlConfig {
	s.DataPartition = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetDataSource(v string) *AutofeExperimentConfigurationYmlConfig {
	s.DataSource = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetDataType(v string) *AutofeExperimentConfigurationYmlConfig {
	s.DataType = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetDebugMode(v string) *AutofeExperimentConfigurationYmlConfig {
	s.DebugMode = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetExcludeColumns(v string) *AutofeExperimentConfigurationYmlConfig {
	s.ExcludeColumns = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetFeatureSelection(v string) *AutofeExperimentConfigurationYmlConfig {
	s.FeatureSelection = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetFilterThresh(v string) *AutofeExperimentConfigurationYmlConfig {
	s.FilterThresh = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetIvThresh(v string) *AutofeExperimentConfigurationYmlConfig {
	s.IvThresh = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetLabel(v string) *AutofeExperimentConfigurationYmlConfig {
	s.Label = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetMemory(v string) *AutofeExperimentConfigurationYmlConfig {
	s.Memory = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetOutputConfigOssDir(v string) *AutofeExperimentConfigurationYmlConfig {
	s.OutputConfigOssDir = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetPipelineExpId(v string) *AutofeExperimentConfigurationYmlConfig {
	s.PipelineExpId = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetReuseResults(v string) *AutofeExperimentConfigurationYmlConfig {
	s.ReuseResults = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetSampleRatio(v string) *AutofeExperimentConfigurationYmlConfig {
	s.SampleRatio = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetSampleSize(v string) *AutofeExperimentConfigurationYmlConfig {
	s.SampleSize = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetSelectionExpId(v string) *AutofeExperimentConfigurationYmlConfig {
	s.SelectionExpId = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetSkipSelect(v string) *AutofeExperimentConfigurationYmlConfig {
	s.SkipSelect = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetWorkers(v string) *AutofeExperimentConfigurationYmlConfig {
	s.Workers = &v
	return s
}

func (s *AutofeExperimentConfigurationYmlConfig) SetWorkspaceName(v string) *AutofeExperimentConfigurationYmlConfig {
	s.WorkspaceName = &v
	return s
}

type HpoExperimentConfig struct {
	DlcConfig *HpoExperimentConfigDlcConfig `json:"dlc_config,omitempty" xml:"dlc_config,omitempty" type:"Struct"`
	K8sConfig *HpoExperimentConfigK8sConfig `json:"k8s_config,omitempty" xml:"k8s_config,omitempty" type:"Struct"`
	// This parameter is required.
	MetricConfig  *HpoExperimentConfigMetricConfig  `json:"metric_config,omitempty" xml:"metric_config,omitempty" type:"Struct"`
	MonitorConfig *HpoExperimentConfigMonitorConfig `json:"monitor_config,omitempty" xml:"monitor_config,omitempty" type:"Struct"`
	OdpsConfig    *HpoExperimentConfigOdpsConfig    `json:"odps_config,omitempty" xml:"odps_config,omitempty" type:"Struct"`
	OssConfig     *HpoExperimentConfigOssConfig     `json:"oss_config,omitempty" xml:"oss_config,omitempty" type:"Struct"`
	OutputConfig  *HpoExperimentConfigOutputConfig  `json:"output_config,omitempty" xml:"output_config,omitempty" type:"Struct"`
	PaiflowConfig *HpoExperimentConfigPaiflowConfig `json:"paiflow_config,omitempty" xml:"paiflow_config,omitempty" type:"Struct"`
	ParamsConfig  *HpoExperimentConfigParamsConfig  `json:"params_config,omitempty" xml:"params_config,omitempty" type:"Struct"`
	// This parameter is required.
	PlatformConfig *HpoExperimentConfigPlatformConfig `json:"platform_config,omitempty" xml:"platform_config,omitempty" type:"Struct"`
	ScheduleConfig *HpoExperimentConfigScheduleConfig `json:"schedule_config,omitempty" xml:"schedule_config,omitempty" type:"Struct"`
	SearchSpace    map[string]interface{}             `json:"search_space,omitempty" xml:"search_space,omitempty"`
	TsConfig       *HpoExperimentConfigTsConfig       `json:"ts_config,omitempty" xml:"ts_config,omitempty" type:"Struct"`
	YmlConfig      *HpoExperimentConfigYmlConfig      `json:"yml_config,omitempty" xml:"yml_config,omitempty" type:"Struct"`
}

func (s HpoExperimentConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfig) SetDlcConfig(v *HpoExperimentConfigDlcConfig) *HpoExperimentConfig {
	s.DlcConfig = v
	return s
}

func (s *HpoExperimentConfig) SetK8sConfig(v *HpoExperimentConfigK8sConfig) *HpoExperimentConfig {
	s.K8sConfig = v
	return s
}

func (s *HpoExperimentConfig) SetMetricConfig(v *HpoExperimentConfigMetricConfig) *HpoExperimentConfig {
	s.MetricConfig = v
	return s
}

func (s *HpoExperimentConfig) SetMonitorConfig(v *HpoExperimentConfigMonitorConfig) *HpoExperimentConfig {
	s.MonitorConfig = v
	return s
}

func (s *HpoExperimentConfig) SetOdpsConfig(v *HpoExperimentConfigOdpsConfig) *HpoExperimentConfig {
	s.OdpsConfig = v
	return s
}

func (s *HpoExperimentConfig) SetOssConfig(v *HpoExperimentConfigOssConfig) *HpoExperimentConfig {
	s.OssConfig = v
	return s
}

func (s *HpoExperimentConfig) SetOutputConfig(v *HpoExperimentConfigOutputConfig) *HpoExperimentConfig {
	s.OutputConfig = v
	return s
}

func (s *HpoExperimentConfig) SetPaiflowConfig(v *HpoExperimentConfigPaiflowConfig) *HpoExperimentConfig {
	s.PaiflowConfig = v
	return s
}

func (s *HpoExperimentConfig) SetParamsConfig(v *HpoExperimentConfigParamsConfig) *HpoExperimentConfig {
	s.ParamsConfig = v
	return s
}

func (s *HpoExperimentConfig) SetPlatformConfig(v *HpoExperimentConfigPlatformConfig) *HpoExperimentConfig {
	s.PlatformConfig = v
	return s
}

func (s *HpoExperimentConfig) SetScheduleConfig(v *HpoExperimentConfigScheduleConfig) *HpoExperimentConfig {
	s.ScheduleConfig = v
	return s
}

func (s *HpoExperimentConfig) SetSearchSpace(v map[string]interface{}) *HpoExperimentConfig {
	s.SearchSpace = v
	return s
}

func (s *HpoExperimentConfig) SetTsConfig(v *HpoExperimentConfigTsConfig) *HpoExperimentConfig {
	s.TsConfig = v
	return s
}

func (s *HpoExperimentConfig) SetYmlConfig(v *HpoExperimentConfigYmlConfig) *HpoExperimentConfig {
	s.YmlConfig = v
	return s
}

type HpoExperimentConfigDlcConfig struct {
	AccessId  *string `json:"access_id,omitempty" xml:"access_id,omitempty"`
	AccessKey *string `json:"access_key,omitempty" xml:"access_key,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Protocol  *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s HpoExperimentConfigDlcConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigDlcConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigDlcConfig) SetAccessId(v string) *HpoExperimentConfigDlcConfig {
	s.AccessId = &v
	return s
}

func (s *HpoExperimentConfigDlcConfig) SetAccessKey(v string) *HpoExperimentConfigDlcConfig {
	s.AccessKey = &v
	return s
}

func (s *HpoExperimentConfigDlcConfig) SetEndpoint(v string) *HpoExperimentConfigDlcConfig {
	s.Endpoint = &v
	return s
}

func (s *HpoExperimentConfigDlcConfig) SetProtocol(v string) *HpoExperimentConfigDlcConfig {
	s.Protocol = &v
	return s
}

func (s *HpoExperimentConfigDlcConfig) SetRegion(v string) *HpoExperimentConfigDlcConfig {
	s.Region = &v
	return s
}

type HpoExperimentConfigK8sConfig struct {
	NniContainerCpuLimit        *string `json:"nni_container_cpu_limit,omitempty" xml:"nni_container_cpu_limit,omitempty"`
	NniContainerMemoryLimit     *string `json:"nni_container_memory_limit,omitempty" xml:"nni_container_memory_limit,omitempty"`
	NniContainerRequestedCpu    *string `json:"nni_container_requested_cpu,omitempty" xml:"nni_container_requested_cpu,omitempty"`
	NniContainerRequestedMemory *string `json:"nni_container_requested_memory,omitempty" xml:"nni_container_requested_memory,omitempty"`
}

func (s HpoExperimentConfigK8sConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigK8sConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigK8sConfig) SetNniContainerCpuLimit(v string) *HpoExperimentConfigK8sConfig {
	s.NniContainerCpuLimit = &v
	return s
}

func (s *HpoExperimentConfigK8sConfig) SetNniContainerMemoryLimit(v string) *HpoExperimentConfigK8sConfig {
	s.NniContainerMemoryLimit = &v
	return s
}

func (s *HpoExperimentConfigK8sConfig) SetNniContainerRequestedCpu(v string) *HpoExperimentConfigK8sConfig {
	s.NniContainerRequestedCpu = &v
	return s
}

func (s *HpoExperimentConfigK8sConfig) SetNniContainerRequestedMemory(v string) *HpoExperimentConfigK8sConfig {
	s.NniContainerRequestedMemory = &v
	return s
}

type HpoExperimentConfigMetricConfig struct {
	// This parameter is required.
	FinalMode *string `json:"final_mode,omitempty" xml:"final_mode,omitempty"`
	// This parameter is required.
	MetricDict map[string]interface{} `json:"metric_dict,omitempty" xml:"metric_dict,omitempty"`
	// This parameter is required.
	MetricSource []*string `json:"metric_source,omitempty" xml:"metric_source,omitempty" type:"Repeated"`
	// This parameter is required.
	MetricType          *string `json:"metric_type,omitempty" xml:"metric_type,omitempty"`
	SourceListFinalMode *string `json:"source_list_final_mode,omitempty" xml:"source_list_final_mode,omitempty"`
}

func (s HpoExperimentConfigMetricConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigMetricConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigMetricConfig) SetFinalMode(v string) *HpoExperimentConfigMetricConfig {
	s.FinalMode = &v
	return s
}

func (s *HpoExperimentConfigMetricConfig) SetMetricDict(v map[string]interface{}) *HpoExperimentConfigMetricConfig {
	s.MetricDict = v
	return s
}

func (s *HpoExperimentConfigMetricConfig) SetMetricSource(v []*string) *HpoExperimentConfigMetricConfig {
	s.MetricSource = v
	return s
}

func (s *HpoExperimentConfigMetricConfig) SetMetricType(v string) *HpoExperimentConfigMetricConfig {
	s.MetricType = &v
	return s
}

func (s *HpoExperimentConfigMetricConfig) SetSourceListFinalMode(v string) *HpoExperimentConfigMetricConfig {
	s.SourceListFinalMode = &v
	return s
}

type HpoExperimentConfigMonitorConfig struct {
	AtMobiles *string `json:"at_mobiles,omitempty" xml:"at_mobiles,omitempty"`
	AtUserIds *string `json:"at_user_ids,omitempty" xml:"at_user_ids,omitempty"`
	IsAtAll   *string `json:"is_at_all,omitempty" xml:"is_at_all,omitempty"`
	Keyword   *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s HpoExperimentConfigMonitorConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigMonitorConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigMonitorConfig) SetAtMobiles(v string) *HpoExperimentConfigMonitorConfig {
	s.AtMobiles = &v
	return s
}

func (s *HpoExperimentConfigMonitorConfig) SetAtUserIds(v string) *HpoExperimentConfigMonitorConfig {
	s.AtUserIds = &v
	return s
}

func (s *HpoExperimentConfigMonitorConfig) SetIsAtAll(v string) *HpoExperimentConfigMonitorConfig {
	s.IsAtAll = &v
	return s
}

func (s *HpoExperimentConfigMonitorConfig) SetKeyword(v string) *HpoExperimentConfigMonitorConfig {
	s.Keyword = &v
	return s
}

func (s *HpoExperimentConfigMonitorConfig) SetUrl(v string) *HpoExperimentConfigMonitorConfig {
	s.Url = &v
	return s
}

type HpoExperimentConfigOdpsConfig struct {
	AccessId    *string `json:"access_id,omitempty" xml:"access_id,omitempty"`
	AccessKey   *string `json:"access_key,omitempty" xml:"access_key,omitempty"`
	EndPoint    *string `json:"end_point,omitempty" xml:"end_point,omitempty"`
	LogViewHost *string `json:"log_view_host,omitempty" xml:"log_view_host,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	Region      *string `json:"region,omitempty" xml:"region,omitempty"`
	RoleArn     *string `json:"role_arn,omitempty" xml:"role_arn,omitempty"`
}

func (s HpoExperimentConfigOdpsConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigOdpsConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigOdpsConfig) SetAccessId(v string) *HpoExperimentConfigOdpsConfig {
	s.AccessId = &v
	return s
}

func (s *HpoExperimentConfigOdpsConfig) SetAccessKey(v string) *HpoExperimentConfigOdpsConfig {
	s.AccessKey = &v
	return s
}

func (s *HpoExperimentConfigOdpsConfig) SetEndPoint(v string) *HpoExperimentConfigOdpsConfig {
	s.EndPoint = &v
	return s
}

func (s *HpoExperimentConfigOdpsConfig) SetLogViewHost(v string) *HpoExperimentConfigOdpsConfig {
	s.LogViewHost = &v
	return s
}

func (s *HpoExperimentConfigOdpsConfig) SetProjectName(v string) *HpoExperimentConfigOdpsConfig {
	s.ProjectName = &v
	return s
}

func (s *HpoExperimentConfigOdpsConfig) SetRegion(v string) *HpoExperimentConfigOdpsConfig {
	s.Region = &v
	return s
}

func (s *HpoExperimentConfigOdpsConfig) SetRoleArn(v string) *HpoExperimentConfigOdpsConfig {
	s.RoleArn = &v
	return s
}

type HpoExperimentConfigOssConfig struct {
	AccessKeyID     *string `json:"accessKeyID,omitempty" xml:"accessKeyID,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	RoleArn         *string `json:"role_arn,omitempty" xml:"role_arn,omitempty"`
}

func (s HpoExperimentConfigOssConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigOssConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigOssConfig) SetAccessKeyID(v string) *HpoExperimentConfigOssConfig {
	s.AccessKeyID = &v
	return s
}

func (s *HpoExperimentConfigOssConfig) SetAccessKeySecret(v string) *HpoExperimentConfigOssConfig {
	s.AccessKeySecret = &v
	return s
}

func (s *HpoExperimentConfigOssConfig) SetEndpoint(v string) *HpoExperimentConfigOssConfig {
	s.Endpoint = &v
	return s
}

func (s *HpoExperimentConfigOssConfig) SetRoleArn(v string) *HpoExperimentConfigOssConfig {
	s.RoleArn = &v
	return s
}

type HpoExperimentConfigOutputConfig struct {
	ModelPath   *string `json:"model_path,omitempty" xml:"model_path,omitempty"`
	SummaryPath *string `json:"summary_path,omitempty" xml:"summary_path,omitempty"`
}

func (s HpoExperimentConfigOutputConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigOutputConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigOutputConfig) SetModelPath(v string) *HpoExperimentConfigOutputConfig {
	s.ModelPath = &v
	return s
}

func (s *HpoExperimentConfigOutputConfig) SetSummaryPath(v string) *HpoExperimentConfigOutputConfig {
	s.SummaryPath = &v
	return s
}

type HpoExperimentConfigPaiflowConfig struct {
	AccessKeyId     *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
	AccessKeySecret *string `json:"access_key_secret,omitempty" xml:"access_key_secret,omitempty"`
	RegionId        *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	WorkspaceId     *string `json:"workspace_id,omitempty" xml:"workspace_id,omitempty"`
}

func (s HpoExperimentConfigPaiflowConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigPaiflowConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigPaiflowConfig) SetAccessKeyId(v string) *HpoExperimentConfigPaiflowConfig {
	s.AccessKeyId = &v
	return s
}

func (s *HpoExperimentConfigPaiflowConfig) SetAccessKeySecret(v string) *HpoExperimentConfigPaiflowConfig {
	s.AccessKeySecret = &v
	return s
}

func (s *HpoExperimentConfigPaiflowConfig) SetRegionId(v string) *HpoExperimentConfigPaiflowConfig {
	s.RegionId = &v
	return s
}

func (s *HpoExperimentConfigPaiflowConfig) SetWorkspaceId(v string) *HpoExperimentConfigPaiflowConfig {
	s.WorkspaceId = &v
	return s
}

type HpoExperimentConfigParamsConfig struct {
	// example:
	//
	// src_path,dst_file_path
	ParamsSrcDstFilepath []*string `json:"params_src_dst_filepath,omitempty" xml:"params_src_dst_filepath,omitempty" type:"Repeated"`
}

func (s HpoExperimentConfigParamsConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigParamsConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigParamsConfig) SetParamsSrcDstFilepath(v []*string) *HpoExperimentConfigParamsConfig {
	s.ParamsSrcDstFilepath = v
	return s
}

type HpoExperimentConfigPlatformConfig struct {
	// This parameter is required.
	Cmd []*string `json:"cmd,omitempty" xml:"cmd,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DLC, MC,paiflow, etc
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Resume *string `json:"resume,omitempty" xml:"resume,omitempty"`
}

func (s HpoExperimentConfigPlatformConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigPlatformConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigPlatformConfig) SetCmd(v []*string) *HpoExperimentConfigPlatformConfig {
	s.Cmd = v
	return s
}

func (s *HpoExperimentConfigPlatformConfig) SetName(v string) *HpoExperimentConfigPlatformConfig {
	s.Name = &v
	return s
}

func (s *HpoExperimentConfigPlatformConfig) SetResume(v string) *HpoExperimentConfigPlatformConfig {
	s.Resume = &v
	return s
}

type HpoExperimentConfigScheduleConfig struct {
	Day       *string `json:"day,omitempty" xml:"day,omitempty"`
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s HpoExperimentConfigScheduleConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigScheduleConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigScheduleConfig) SetDay(v string) *HpoExperimentConfigScheduleConfig {
	s.Day = &v
	return s
}

func (s *HpoExperimentConfigScheduleConfig) SetEndTime(v string) *HpoExperimentConfigScheduleConfig {
	s.EndTime = &v
	return s
}

func (s *HpoExperimentConfigScheduleConfig) SetStartTime(v string) *HpoExperimentConfigScheduleConfig {
	s.StartTime = &v
	return s
}

type HpoExperimentConfigTsConfig struct {
	AccessKeyId     *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
	AccessKeySecret *string `json:"access_key_secret,omitempty" xml:"access_key_secret,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	RegionId        *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
}

func (s HpoExperimentConfigTsConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigTsConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigTsConfig) SetAccessKeyId(v string) *HpoExperimentConfigTsConfig {
	s.AccessKeyId = &v
	return s
}

func (s *HpoExperimentConfigTsConfig) SetAccessKeySecret(v string) *HpoExperimentConfigTsConfig {
	s.AccessKeySecret = &v
	return s
}

func (s *HpoExperimentConfigTsConfig) SetEndpoint(v string) *HpoExperimentConfigTsConfig {
	s.Endpoint = &v
	return s
}

func (s *HpoExperimentConfigTsConfig) SetRegionId(v string) *HpoExperimentConfigTsConfig {
	s.RegionId = &v
	return s
}

type HpoExperimentConfigYmlConfig struct {
	Assessor         *HpoExperimentConfigYmlConfigAssessor `json:"assessor,omitempty" xml:"assessor,omitempty" type:"Struct"`
	Debug            *bool                                 `json:"debug,omitempty" xml:"debug,omitempty"`
	ExperimentName   *string                               `json:"experiment_name,omitempty" xml:"experiment_name,omitempty"`
	LogLevel         *string                               `json:"log_level,omitempty" xml:"log_level,omitempty"`
	MaxTrialNumber   *int32                                `json:"max_trial_number,omitempty" xml:"max_trial_number,omitempty"`
	TrialConcurrency *int32                                `json:"trial_concurrency,omitempty" xml:"trial_concurrency,omitempty"`
	Tuner            *HpoExperimentConfigYmlConfigTuner    `json:"tuner,omitempty" xml:"tuner,omitempty" type:"Struct"`
}

func (s HpoExperimentConfigYmlConfig) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigYmlConfig) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigYmlConfig) SetAssessor(v *HpoExperimentConfigYmlConfigAssessor) *HpoExperimentConfigYmlConfig {
	s.Assessor = v
	return s
}

func (s *HpoExperimentConfigYmlConfig) SetDebug(v bool) *HpoExperimentConfigYmlConfig {
	s.Debug = &v
	return s
}

func (s *HpoExperimentConfigYmlConfig) SetExperimentName(v string) *HpoExperimentConfigYmlConfig {
	s.ExperimentName = &v
	return s
}

func (s *HpoExperimentConfigYmlConfig) SetLogLevel(v string) *HpoExperimentConfigYmlConfig {
	s.LogLevel = &v
	return s
}

func (s *HpoExperimentConfigYmlConfig) SetMaxTrialNumber(v int32) *HpoExperimentConfigYmlConfig {
	s.MaxTrialNumber = &v
	return s
}

func (s *HpoExperimentConfigYmlConfig) SetTrialConcurrency(v int32) *HpoExperimentConfigYmlConfig {
	s.TrialConcurrency = &v
	return s
}

func (s *HpoExperimentConfigYmlConfig) SetTuner(v *HpoExperimentConfigYmlConfigTuner) *HpoExperimentConfigYmlConfig {
	s.Tuner = v
	return s
}

type HpoExperimentConfigYmlConfigAssessor struct {
	ClassArgs *HpoExperimentConfigYmlConfigAssessorClassArgs `json:"class_args,omitempty" xml:"class_args,omitempty" type:"Struct"`
	Name      *string                                        `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HpoExperimentConfigYmlConfigAssessor) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigYmlConfigAssessor) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigYmlConfigAssessor) SetClassArgs(v *HpoExperimentConfigYmlConfigAssessorClassArgs) *HpoExperimentConfigYmlConfigAssessor {
	s.ClassArgs = v
	return s
}

func (s *HpoExperimentConfigYmlConfigAssessor) SetName(v string) *HpoExperimentConfigYmlConfigAssessor {
	s.Name = &v
	return s
}

type HpoExperimentConfigYmlConfigAssessorClassArgs struct {
	Earlystop    *bool    `json:"earlystop,omitempty" xml:"earlystop,omitempty"`
	MovingAvg    *string  `json:"moving_avg,omitempty" xml:"moving_avg,omitempty"`
	OptimizeMode *string  `json:"optimize_mode,omitempty" xml:"optimize_mode,omitempty"`
	Proportion   *float32 `json:"proportion,omitempty" xml:"proportion,omitempty"`
	StartStep    *int32   `json:"start_step,omitempty" xml:"start_step,omitempty"`
}

func (s HpoExperimentConfigYmlConfigAssessorClassArgs) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigYmlConfigAssessorClassArgs) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigYmlConfigAssessorClassArgs) SetEarlystop(v bool) *HpoExperimentConfigYmlConfigAssessorClassArgs {
	s.Earlystop = &v
	return s
}

func (s *HpoExperimentConfigYmlConfigAssessorClassArgs) SetMovingAvg(v string) *HpoExperimentConfigYmlConfigAssessorClassArgs {
	s.MovingAvg = &v
	return s
}

func (s *HpoExperimentConfigYmlConfigAssessorClassArgs) SetOptimizeMode(v string) *HpoExperimentConfigYmlConfigAssessorClassArgs {
	s.OptimizeMode = &v
	return s
}

func (s *HpoExperimentConfigYmlConfigAssessorClassArgs) SetProportion(v float32) *HpoExperimentConfigYmlConfigAssessorClassArgs {
	s.Proportion = &v
	return s
}

func (s *HpoExperimentConfigYmlConfigAssessorClassArgs) SetStartStep(v int32) *HpoExperimentConfigYmlConfigAssessorClassArgs {
	s.StartStep = &v
	return s
}

type HpoExperimentConfigYmlConfigTuner struct {
	ClassArgs map[string]interface{} `json:"class_args,omitempty" xml:"class_args,omitempty"`
	Name      *string                `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HpoExperimentConfigYmlConfigTuner) String() string {
	return tea.Prettify(s)
}

func (s HpoExperimentConfigYmlConfigTuner) GoString() string {
	return s.String()
}

func (s *HpoExperimentConfigYmlConfigTuner) SetClassArgs(v map[string]interface{}) *HpoExperimentConfigYmlConfigTuner {
	s.ClassArgs = v
	return s
}

func (s *HpoExperimentConfigYmlConfigTuner) SetName(v string) *HpoExperimentConfigYmlConfigTuner {
	s.Name = &v
	return s
}

type CreateAutofeExperimentRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility                 *string                        `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AutofeExperimentConfiguration *AutofeExperimentConfiguration `json:"AutofeExperimentConfiguration,omitempty" xml:"AutofeExperimentConfiguration,omitempty"`
	// example:
	//
	// This is an AutoFE experiment.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my experiment x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 283301
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAutofeExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutofeExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateAutofeExperimentRequest) SetAccessibility(v string) *CreateAutofeExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateAutofeExperimentRequest) SetAutofeExperimentConfiguration(v *AutofeExperimentConfiguration) *CreateAutofeExperimentRequest {
	s.AutofeExperimentConfiguration = v
	return s
}

func (s *CreateAutofeExperimentRequest) SetDescription(v string) *CreateAutofeExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateAutofeExperimentRequest) SetName(v string) *CreateAutofeExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateAutofeExperimentRequest) SetWorkspaceId(v string) *CreateAutofeExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type CreateAutofeExperimentResponseBody struct {
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// t8cNdMO
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 22A1EC4E-15A2-51FF-BFA3-10D6735BAA69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutofeExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutofeExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutofeExperimentResponseBody) SetCode(v string) *CreateAutofeExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAutofeExperimentResponseBody) SetDetail(v map[string]interface{}) *CreateAutofeExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *CreateAutofeExperimentResponseBody) SetExperimentId(v string) *CreateAutofeExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateAutofeExperimentResponseBody) SetMessage(v string) *CreateAutofeExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAutofeExperimentResponseBody) SetRequestId(v string) *CreateAutofeExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CreateAutofeExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutofeExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutofeExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutofeExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateAutofeExperimentResponse) SetHeaders(v map[string]*string) *CreateAutofeExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateAutofeExperimentResponse) SetStatusCode(v int32) *CreateAutofeExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutofeExperimentResponse) SetBody(v *CreateAutofeExperimentResponseBody) *CreateAutofeExperimentResponse {
	s.Body = v
	return s
}

type CreateHpoExperimentRequest struct {
	// Experiment accesibility, public or private.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Experiment description.
	//
	// example:
	//
	// This is an AutoML HPO experiment.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The config object of the expriment.
	HpoExperimentConfiguration *HpoExperimentConfig `json:"HpoExperimentConfiguration,omitempty" xml:"HpoExperimentConfiguration,omitempty"`
	// Experiment Name
	//
	// example:
	//
	// my experiment x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// AI Workspace ID
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateHpoExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHpoExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateHpoExperimentRequest) SetAccessibility(v string) *CreateHpoExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateHpoExperimentRequest) SetDescription(v string) *CreateHpoExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateHpoExperimentRequest) SetHpoExperimentConfiguration(v *HpoExperimentConfig) *CreateHpoExperimentRequest {
	s.HpoExperimentConfiguration = v
	return s
}

func (s *CreateHpoExperimentRequest) SetName(v string) *CreateHpoExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateHpoExperimentRequest) SetWorkspaceId(v string) *CreateHpoExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type CreateHpoExperimentResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Detailed information of the failure.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Id of the request
	//
	// example:
	//
	// sORVEck
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHpoExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHpoExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHpoExperimentResponseBody) SetCode(v string) *CreateHpoExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHpoExperimentResponseBody) SetDetail(v map[string]*string) *CreateHpoExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *CreateHpoExperimentResponseBody) SetExperimentId(v string) *CreateHpoExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateHpoExperimentResponseBody) SetMessage(v string) *CreateHpoExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHpoExperimentResponseBody) SetRequestId(v string) *CreateHpoExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CreateHpoExperimentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHpoExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHpoExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHpoExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateHpoExperimentResponse) SetHeaders(v map[string]*string) *CreateHpoExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateHpoExperimentResponse) SetStatusCode(v int32) *CreateHpoExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHpoExperimentResponse) SetBody(v *CreateHpoExperimentResponseBody) *CreateHpoExperimentResponse {
	s.Body = v
	return s
}

type CreateServiceIdentityRoleResponseBody struct {
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Execution failure, please try again.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AliyunServiceRoleForPaiAutoml
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateServiceIdentityRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponseBody) SetCode(v string) *CreateServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetMessage(v string) *CreateServiceIdentityRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRequestId(v string) *CreateServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleName(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

type CreateServiceIdentityRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *CreateServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetStatusCode(v int32) *CreateServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetBody(v *CreateServiceIdentityRoleResponseBody) *CreateServiceIdentityRoleResponse {
	s.Body = v
	return s
}

type DeleteHpoExperimentResponseBody struct {
	// The error code when the API call is not successful.
	//
	// example:
	//
	// INVALID_USER_OR_EXP
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra info on the execution failure.
	//
	// example:
	//
	// {}
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHpoExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHpoExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHpoExperimentResponseBody) SetCode(v string) *DeleteHpoExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHpoExperimentResponseBody) SetDetail(v map[string]interface{}) *DeleteHpoExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *DeleteHpoExperimentResponseBody) SetMessage(v string) *DeleteHpoExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHpoExperimentResponseBody) SetRequestId(v string) *DeleteHpoExperimentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHpoExperimentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHpoExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHpoExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHpoExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteHpoExperimentResponse) SetHeaders(v map[string]*string) *DeleteHpoExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteHpoExperimentResponse) SetStatusCode(v int32) *DeleteHpoExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHpoExperimentResponse) SetBody(v *DeleteHpoExperimentResponseBody) *DeleteHpoExperimentResponse {
	s.Body = v
	return s
}

type GetAutofeExperimentResponseBody struct {
	// example:
	//
	// INVALID_USER_OR_EXP
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 266977839132045194
	Creator *string                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Detail  map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// tfKs0Ds
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2023-08-18T09:41:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-11-30T02:05:34.000Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// {"action": "fs_train", "request_id": "6B6D4FB6-87AA-5C0E-94AB-E0F830955D1E", "job_num": 7, "job_runs": [{"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-fqlk86bvk68", "duration": 181, "finish_at": 1719158854000, "create_time": "2024-06-23T16:04:31.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-q18k5jo66h3k", "duration": 233, "finish_at": 1719159106000, "create_time": "2024-06-23T16:07:51.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-g7b38dmqes0nm4le", "duration": 173, "finish_at": 1719159306000, "create_time": "2024-06-23T16:12:12.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-524dab97aazs3jmc", "duration": 199, "finish_at": 1719159527000, "create_time": "2024-06-23T16:15:28.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-a8z56wy2ppwmxx04", "duration": 197, "finish_at": 1719159747000, "create_time": "2024-06-23T16:19:08.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-2cgsbxzzzma1j", "duration": 210, "finish_at": 1719159982000, "create_time": "2024-06-23T16:22:51.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-3cgsvf3ien1cw7", "duration": 184, "finish_at": 1719160188000, "create_time": "2024-06-23T16:26:43.000Z", "status": "Succeeded"}], "workspace_id": "86006"}{"action": "fs_train", "request_id": "6B6D4FB6-87AA-5C0E-94AB-E0F830955D1E", "job_num": 7, "job_runs": [{"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-fqlk86betchc69vk68", "duration": 181, "finish_at": 1719158854000, "create_time": "2024-06-23T16:04:31.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-q18ijqak558jo66h3k", "duration": 233, "finish_at": 1719159106000, "create_time": "2024-06-23T16:07:51.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-g7b38df7mqes0nm4le", "duration": 173, "finish_at": 1719159306000, "create_time": "2024-06-23T16:12:12.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-524zedab97aazs3jmc", "duration": 199, "finish_at": 1719159527000, "create_time": "2024-06-23T16:15:28.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-a8z56wy2ppm6wmxx04", "duration": 197, "finish_at": 1719159747000, "create_time": "2024-06-23T16:19:08.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-2cgsbxzx9i35zzma1j", "duration": 210, "finish_at": 1719159982000, "create_time": "2024-06-23T16:22:51.000Z", "status": "Succeeded"}, {"job_url": "https://pai.console.aliyun.com/?projectId=&regionId=cn-beijing#/studio/task/detail/flow-3cgsvf3bprlien1cw7", "duration": 184, "finish_at": 1719160188000, "create_time": "2024-06-23T16:26:43.000Z", "status": "Succeeded"}], "workspace_id": "86006"}
	JobMeta *string `json:"JobMeta,omitempty" xml:"JobMeta,omitempty"`
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 071A904D-5A49-597F-9F69-81C7701D04AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// age,sex,education
	SelectedFeatures *string `json:"SelectedFeatures,omitempty" xml:"SelectedFeatures,omitempty"`
	// example:
	//
	// FINISHED, FAILED, USER_CANCELED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAutofeExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutofeExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutofeExperimentResponseBody) SetCode(v string) *GetAutofeExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetCreator(v string) *GetAutofeExperimentResponseBody {
	s.Creator = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetDetail(v map[string]interface{}) *GetAutofeExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetExperimentId(v string) *GetAutofeExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetGmtCreateTime(v string) *GetAutofeExperimentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetGmtModifiedTime(v string) *GetAutofeExperimentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetJobMeta(v string) *GetAutofeExperimentResponseBody {
	s.JobMeta = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetMessage(v string) *GetAutofeExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetName(v string) *GetAutofeExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetRequestId(v string) *GetAutofeExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetSelectedFeatures(v string) *GetAutofeExperimentResponseBody {
	s.SelectedFeatures = &v
	return s
}

func (s *GetAutofeExperimentResponseBody) SetStatus(v string) *GetAutofeExperimentResponseBody {
	s.Status = &v
	return s
}

type GetAutofeExperimentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutofeExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutofeExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutofeExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetAutofeExperimentResponse) SetHeaders(v map[string]*string) *GetAutofeExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetAutofeExperimentResponse) SetStatusCode(v int32) *GetAutofeExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutofeExperimentResponse) SetBody(v *GetAutofeExperimentResponseBody) *GetAutofeExperimentResponse {
	s.Body = v
	return s
}

type GetHpoExperimentResponseBody struct {
	// Accessibility of expriment, public or private.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Error when the API call is not success.
	//
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Experiment run environment configurations.
	//
	// example:
	//
	// [odps_config]
	//
	// access_id=foo
	//
	// access_key=bar
	//
	// end_point=http://service.odps.aliyun.com/api
	//
	// log_view_host=http://logview.odps.aliyun.com
	//
	// project_name=my_project
	ConfigIni *string `json:"ConfigIni,omitempty" xml:"ConfigIni,omitempty"`
	// HPO search config yaml.
	//
	// example:
	//
	// assessor:
	//
	//   classArgs:
	//
	//     optimize_mode: maximize
	//
	//     start_step: 1
	//
	//   name: PAIAssessor
	//
	// debug: true
	//
	// experimentName: maxcompute_kmeans_monitor
	//
	// experimentWorkingDirectory: ../expdir
	//
	// logLevel: debug
	//
	// maxTrialNumber: 10
	//
	// searchSpaceFile: search_space.json
	//
	// trainingService:
	//
	//   platform: local
	//
	// trialCommand: python3 -m hpo_tools.core.utils.run --config=./config.ini
	//
	// trialConcurrency: 2
	//
	// tuner:
	//
	//   classArgs:
	//
	//     optimize_mode: maximize
	//
	//   name: TPE
	ConfigYml *string `json:"ConfigYml,omitempty" xml:"ConfigYml,omitempty"`
	// The one who created the experiment.
	//
	// example:
	//
	// 123456789
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// If the Experiment if deleted.
	//
	// example:
	//
	// 0
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// Description of the experiment.
	//
	// example:
	//
	// This is experiment is for tune the LR of model.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Extra error message info.
	//
	// example:
	//
	// TBD
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Experiment ID.
	//
	// example:
	//
	// sX5O9Q8
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Experiment create time.
	//
	// example:
	//
	// 2024-01-01 08:30:11
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Experiment last update time.
	//
	// example:
	//
	// 2024-01-01 08:30:11
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Experiment configuration in json format.
	//
	// example:
	//
	// {
	//
	//     "output_config": {
	//
	//         "model_path": "bi_ps_${exp_id}_${trial_id}"
	//
	//     },
	//
	//     "platform_config": {
	//
	//         "cmd": [
	//
	//             "PAI -name ps_smart\\n    -project foo",
	//
	//             "PAI -name prediction\\n    -project foo",
	//
	//             "PAI -name evaluate -project foo",
	//
	//             "INSERT OVERWRITE TABLE my_table;"
	//
	//         ],
	//
	//         "name": "MaxCompute"
	//
	//     },
	//
	//     "metric_config": {
	//
	//         "metric_dict": {
	//
	//             "recall": 0.5,
	//
	//             "auc": 0.25,
	//
	//             "precision": 0.25
	//
	//         },
	//
	//         "metric_source": [
	//
	//             "select 	- from my_metrics where pt=\\"${exp_id}_${trial_id}\\";"
	//
	//         ],
	//
	//         "metric_type": "table",
	//
	//         "final_mode": "best"
	//
	//     },
	//
	//     "odps_config": {
	//
	//         "project_name": "my_project",
	//
	//         "role_arn": "acs:ram::123456789:role/aliyunserviceroleforfoo",
	//
	//         "region": "cn-shanghai",
	//
	//         "end_point": "http://service.cn.maxcompute.aliyun-inc.com/api",
	//
	//         "log_view_host": "http://logview.odps.aliyun.com"
	//
	//     },
	//
	//     "yml_config": {
	//
	//         "max_trial_number": 5,
	//
	//         "assessor": {
	//
	//             "name": "PAIAssessor",
	//
	//             "class_args": {
	//
	//                 "earlystop": true,
	//
	//                 "start_step": 5,
	//
	//                 "optimize_mode": "maximize"
	//
	//             }
	//
	//         },
	//
	//         "experiment_name": "my_exp",
	//
	//         "tuner": {
	//
	//             "name": "TPE",
	//
	//             "class_args": {
	//
	//                 "optimize_mode": "maximize"
	//
	//             }
	//
	//         },
	//
	//         "trial_concurrency": 2
	//
	//     },
	//
	//     "search_space": {
	//
	//         "${max_depth}": {
	//
	//             "_type": "choice",
	//
	//             "_value": [
	//
	//                 1,
	//
	//                 3,
	//
	//                 5
	//
	//             ]
	//
	//         },
	//
	//         "${tree_count}": {
	//
	//             "_value": [
	//
	//                 50,
	//
	//                 100,
	//
	//                 150
	//
	//             ],
	//
	//             "_type": "choice"
	//
	//         }
	//
	//     }
	//
	// }
	HpoExperimentConfiguration map[string]interface{} `json:"HpoExperimentConfiguration,omitempty" xml:"HpoExperimentConfiguration,omitempty"`
	// Experiment  Job type.
	//
	// example:
	//
	// 1
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Experiment name.
	//
	// example:
	//
	// my_hpo_exp_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// HPO hyperparameter search space.
	//
	// example:
	//
	// {"${centerCount}":{"_type":"choice","_value":[2,3,4,5]},"${distanceType}":{"_type":"choice","_value":["euclidean","cosine","cityblock"]}}
	SearchSpace *string `json:"SearchSpace,omitempty" xml:"SearchSpace,omitempty"`
	// Experiment status.
	//
	// example:
	//
	// INVALID_USER_OR_EXP, SUCCESS, or EXECUTION_FAILURE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Trials amount run till now.
	//
	// example:
	//
	// 3
	TrialCount *int32 `json:"TrialCount,omitempty" xml:"TrialCount,omitempty"`
	// Status if a trial
	TrialStatus map[string]*string `json:"TrialStatus,omitempty" xml:"TrialStatus,omitempty"`
	// AI Workspace ID.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetHpoExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHpoExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetHpoExperimentResponseBody) SetAccessibility(v string) *GetHpoExperimentResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetCode(v string) *GetHpoExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetConfigIni(v string) *GetHpoExperimentResponseBody {
	s.ConfigIni = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetConfigYml(v string) *GetHpoExperimentResponseBody {
	s.ConfigYml = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetCreator(v string) *GetHpoExperimentResponseBody {
	s.Creator = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetDeleted(v bool) *GetHpoExperimentResponseBody {
	s.Deleted = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetDescription(v string) *GetHpoExperimentResponseBody {
	s.Description = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetDetail(v map[string]interface{}) *GetHpoExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *GetHpoExperimentResponseBody) SetExperimentId(v string) *GetHpoExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetGmtCreateTime(v string) *GetHpoExperimentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetGmtModifiedTime(v string) *GetHpoExperimentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetHpoExperimentConfiguration(v map[string]interface{}) *GetHpoExperimentResponseBody {
	s.HpoExperimentConfiguration = v
	return s
}

func (s *GetHpoExperimentResponseBody) SetJobType(v string) *GetHpoExperimentResponseBody {
	s.JobType = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetMessage(v string) *GetHpoExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetName(v string) *GetHpoExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetRequestId(v string) *GetHpoExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetSearchSpace(v string) *GetHpoExperimentResponseBody {
	s.SearchSpace = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetStatus(v string) *GetHpoExperimentResponseBody {
	s.Status = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetTrialCount(v int32) *GetHpoExperimentResponseBody {
	s.TrialCount = &v
	return s
}

func (s *GetHpoExperimentResponseBody) SetTrialStatus(v map[string]*string) *GetHpoExperimentResponseBody {
	s.TrialStatus = v
	return s
}

func (s *GetHpoExperimentResponseBody) SetWorkspaceId(v string) *GetHpoExperimentResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetHpoExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHpoExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHpoExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHpoExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetHpoExperimentResponse) SetHeaders(v map[string]*string) *GetHpoExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetHpoExperimentResponse) SetStatusCode(v int32) *GetHpoExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHpoExperimentResponse) SetBody(v *GetHpoExperimentResponseBody) *GetHpoExperimentResponse {
	s.Body = v
	return s
}

type GetHpoTrialResponseBody struct {
	// Error code.
	//
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Experiment ID.
	//
	// example:
	//
	// abcde
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Final metric of the trial.
	//
	// example:
	//
	// {"default":0.087745,"type":"FINAL","val_loss=([0-9\\\\\\\\.]+)":0.087745}
	FinalMetric *string `json:"FinalMetric,omitempty" xml:"FinalMetric,omitempty"`
	// Trial create time.
	//
	// example:
	//
	// 2024-01-02 12:34:56
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Trial last update time.
	//
	// example:
	//
	// 2024-01-02 21:32:56
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Hyperparameters used to run the trial.
	//
	// example:
	//
	// {"${centerCount}": 5, "${distanceType}": "cosine"}
	Hyperparam *string `json:"Hyperparam,omitempty" xml:"Hyperparam,omitempty"`
	// trial meta infomation.
	//
	// example:
	//
	// {"nni_trial_id": "asdf", "nni_exp_id": "abcde", "nni_trial_sequence_id": 2}
	JobMeta *string `json:"JobMeta,omitempty" xml:"JobMeta,omitempty"`
	// The error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// metric details uploaded during running.
	//
	// example:
	//
	// [{"default":0.087745,"val_loss=([0-9\\\\\\\\.]+)":0.087745},
	//
	// {"default":1.085841,"val_loss=([0-9\\\\\\\\.]+)":0.085841},
	//
	// {"default":0.087745,"type":"FINAL","val_loss=([0-9\\\\\\\\.]+)":0.087745}]
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// AUC or F1, Accuracy etc.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Trial output model.
	//
	// example:
	//
	// model_table_name_foo
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Parameter id, start from 0.
	//
	// example:
	//
	// 0
	ParameterId *int32 `json:"ParameterId,omitempty" xml:"ParameterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Trial status.
	//
	// example:
	//
	// FINISHED, FAILED, USER_CANCELED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Trail ID.
	//
	// example:
	//
	// mf99W4
	TrialId *string `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	// User added comments.
	//
	// example:
	//
	// best trial till now.
	UserComment *string `json:"UserComment,omitempty" xml:"UserComment,omitempty"`
	// User updated score.
	//
	// example:
	//
	// 0
	UserScore *int32 `json:"UserScore,omitempty" xml:"UserScore,omitempty"`
}

func (s GetHpoTrialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHpoTrialResponseBody) GoString() string {
	return s.String()
}

func (s *GetHpoTrialResponseBody) SetCode(v string) *GetHpoTrialResponseBody {
	s.Code = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetDetail(v map[string]*string) *GetHpoTrialResponseBody {
	s.Detail = v
	return s
}

func (s *GetHpoTrialResponseBody) SetExperimentId(v string) *GetHpoTrialResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetFinalMetric(v string) *GetHpoTrialResponseBody {
	s.FinalMetric = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetGmtCreateTime(v string) *GetHpoTrialResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetGmtModifiedTime(v string) *GetHpoTrialResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetHyperparam(v string) *GetHpoTrialResponseBody {
	s.Hyperparam = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetJobMeta(v string) *GetHpoTrialResponseBody {
	s.JobMeta = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetMessage(v string) *GetHpoTrialResponseBody {
	s.Message = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetMetric(v string) *GetHpoTrialResponseBody {
	s.Metric = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetMetricName(v string) *GetHpoTrialResponseBody {
	s.MetricName = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetModel(v string) *GetHpoTrialResponseBody {
	s.Model = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetParameterId(v int32) *GetHpoTrialResponseBody {
	s.ParameterId = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetRequestId(v string) *GetHpoTrialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetStatus(v string) *GetHpoTrialResponseBody {
	s.Status = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetTrialId(v string) *GetHpoTrialResponseBody {
	s.TrialId = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetUserComment(v string) *GetHpoTrialResponseBody {
	s.UserComment = &v
	return s
}

func (s *GetHpoTrialResponseBody) SetUserScore(v int32) *GetHpoTrialResponseBody {
	s.UserScore = &v
	return s
}

type GetHpoTrialResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHpoTrialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHpoTrialResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHpoTrialResponse) GoString() string {
	return s.String()
}

func (s *GetHpoTrialResponse) SetHeaders(v map[string]*string) *GetHpoTrialResponse {
	s.Headers = v
	return s
}

func (s *GetHpoTrialResponse) SetStatusCode(v int32) *GetHpoTrialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHpoTrialResponse) SetBody(v *GetHpoTrialResponseBody) *GetHpoTrialResponse {
	s.Body = v
	return s
}

type GetServiceIdentityRoleResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Execution failure, please try again.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AliyunServiceRoleForPaiAutoml
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceIdentityRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponseBody) SetCode(v string) *GetServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetMessage(v string) *GetServiceIdentityRoleResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRequestId(v string) *GetServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleName(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

type GetServiceIdentityRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *GetServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetStatusCode(v int32) *GetServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetBody(v *GetServiceIdentityRoleResponseBody) *GetServiceIdentityRoleResponse {
	s.Body = v
	return s
}

type ListHpoExperimentLogsRequest struct {
	// Log name to be listed.
	//
	// example:
	//
	// nnimanager.log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 40
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHpoExperimentLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentLogsRequest) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentLogsRequest) SetLogName(v string) *ListHpoExperimentLogsRequest {
	s.LogName = &v
	return s
}

func (s *ListHpoExperimentLogsRequest) SetPageNumber(v int32) *ListHpoExperimentLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHpoExperimentLogsRequest) SetPageSize(v int32) *ListHpoExperimentLogsRequest {
	s.PageSize = &v
	return s
}

type ListHpoExperimentLogsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error message.
	//
	// example:
	//
	// {}
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The log.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 071A904D-5A49-597F-9F69-81C7701D04AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 467
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHpoExperimentLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentLogsResponseBody) SetCode(v string) *ListHpoExperimentLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHpoExperimentLogsResponseBody) SetDetail(v map[string]interface{}) *ListHpoExperimentLogsResponseBody {
	s.Detail = v
	return s
}

func (s *ListHpoExperimentLogsResponseBody) SetLogs(v []*string) *ListHpoExperimentLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListHpoExperimentLogsResponseBody) SetMessage(v string) *ListHpoExperimentLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHpoExperimentLogsResponseBody) SetRequestId(v string) *ListHpoExperimentLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHpoExperimentLogsResponseBody) SetTotalCount(v int32) *ListHpoExperimentLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHpoExperimentLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHpoExperimentLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHpoExperimentLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentLogsResponse) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentLogsResponse) SetHeaders(v map[string]*string) *ListHpoExperimentLogsResponse {
	s.Headers = v
	return s
}

func (s *ListHpoExperimentLogsResponse) SetStatusCode(v int32) *ListHpoExperimentLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHpoExperimentLogsResponse) SetBody(v *ListHpoExperimentLogsResponseBody) *ListHpoExperimentLogsResponse {
	s.Body = v
	return s
}

type ListHpoExperimentsRequest struct {
	// The accessibility of the experiments to be listed.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Who created the expriment.
	//
	// example:
	//
	// 266977839132045194
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Return experiment detailed configuration or not.
	//
	// example:
	//
	// True, False
	IncludeConfigData *string `json:"IncludeConfigData,omitempty" xml:"IncludeConfigData,omitempty"`
	// The maximum create time of the experiment.
	//
	// example:
	//
	// yyyy-mm-dd hh:mm:ss或着yyyy-mm-dd，比如2023-03-31 10:29:30
	MaxCreateTime *string `json:"MaxCreateTime,omitempty" xml:"MaxCreateTime,omitempty"`
	// The minimum create time of the experiment.
	//
	// example:
	//
	// yyyy-mm-dd hh:mm:ss或着yyyy-mm-dd，比如2023-03-31 10:29:30
	MinCreateTime *string `json:"MinCreateTime,omitempty" xml:"MinCreateTime,omitempty"`
	// Experiment name filter.
	//
	// example:
	//
	// my_hpo_exp_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// ASC, DESC.
	//
	// example:
	//
	// ASC, DESC.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Record number on each page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned experiments to be sorted by this column.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Experiment status to be listed.
	//
	// example:
	//
	// CREATED, RUNNING, FINISHED, FAILED, EARLY_STOPPED, USER_CANCELED, SYS_CANCELED, WAITING, NO_MORE_TRIAL, UNKNOWN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The AI workspace ID the experiments belongs to.
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListHpoExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentsRequest) SetAccessibility(v string) *ListHpoExperimentsRequest {
	s.Accessibility = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetCreator(v string) *ListHpoExperimentsRequest {
	s.Creator = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetIncludeConfigData(v string) *ListHpoExperimentsRequest {
	s.IncludeConfigData = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetMaxCreateTime(v string) *ListHpoExperimentsRequest {
	s.MaxCreateTime = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetMinCreateTime(v string) *ListHpoExperimentsRequest {
	s.MinCreateTime = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetName(v string) *ListHpoExperimentsRequest {
	s.Name = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetOrder(v string) *ListHpoExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetPageNumber(v int32) *ListHpoExperimentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetPageSize(v int32) *ListHpoExperimentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetSortBy(v string) *ListHpoExperimentsRequest {
	s.SortBy = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetStatus(v string) *ListHpoExperimentsRequest {
	s.Status = &v
	return s
}

func (s *ListHpoExperimentsRequest) SetWorkspaceId(v string) *ListHpoExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

type ListHpoExperimentsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// experiment array.
	Experiments []*ListHpoExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total qualified experiment count.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHpoExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentsResponseBody) SetCode(v string) *ListHpoExperimentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHpoExperimentsResponseBody) SetDetail(v map[string]*string) *ListHpoExperimentsResponseBody {
	s.Detail = v
	return s
}

func (s *ListHpoExperimentsResponseBody) SetExperiments(v []*ListHpoExperimentsResponseBodyExperiments) *ListHpoExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListHpoExperimentsResponseBody) SetMessage(v string) *ListHpoExperimentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHpoExperimentsResponseBody) SetRequestId(v string) *ListHpoExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHpoExperimentsResponseBody) SetTotalCount(v int32) *ListHpoExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHpoExperimentsResponseBodyExperiments struct {
	// Experiment accessibility, private or public.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Experiment job run environment configurations.
	//
	// example:
	//
	// [odps_config]
	//
	// key1 = value1
	//
	// key2 = value2
	ConfigIni *string `json:"ConfigIni,omitempty" xml:"ConfigIni,omitempty"`
	// HPO experiment run configuration.
	//
	// example:
	//
	// experimentName: dlc_oss_mnist
	//
	// experimentWorkingDirectory: ../expdir
	//
	// searchSpaceFile: search_space.json
	//
	// trialCommand: python3 -m hpo_tools.core.utils.run --config=./trial.ini
	//
	// trialConcurrency: 1
	//
	// maxTrialNumber: 4
	//
	// tuner:
	//
	//   name: TPE
	//
	//   classArgs:
	//
	//     optimize_mode: maximize
	//
	// trainingService:
	//
	//   platform: local
	//
	// assessor:
	//
	//   name: PAIAssessor
	//
	//   classArgs:
	//
	//     optimize_mode: maximize
	//
	//     start_step: 2
	ConfigYml *string `json:"ConfigYml,omitempty" xml:"ConfigYml,omitempty"`
	// The one who created the experiment.
	//
	// example:
	//
	// 1049310008714189
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// If the experiment is deleted.
	//
	// example:
	//
	// 1
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The description of the experiment.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Experiment ID.
	//
	// example:
	//
	// sORVEck
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Experiment Create Time.
	//
	// example:
	//
	// 2023-03-31T10:29:30Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Last modified time.
	//
	// example:
	//
	// 2022-10-27T00:55:54Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Experimetn job type.
	//
	// example:
	//
	// 1
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Experiment name
	//
	// example:
	//
	// my experiment x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Experiment hyperparameter search space.
	//
	// example:
	//
	// {\\"foo\\":\\"bar\\"}
	SearchSpace *string `json:"SearchSpace,omitempty" xml:"SearchSpace,omitempty"`
	// Experiment status.
	//
	// example:
	//
	// CREATED, RUNNING, FINISHED, FAILED, EARLY_STOPPED, USER_CANCELED, SYS_CANCELED, WAITING, NO_MORE_TRIAL, UNKNOWN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// How many trials the experiment have.
	//
	// example:
	//
	// 20
	TrialCount *int32 `json:"TrialCount,omitempty" xml:"TrialCount,omitempty"`
	// Trial status map.
	TrialStatus map[string]*string `json:"TrialStatus,omitempty" xml:"TrialStatus,omitempty"`
	// The ID of the AI workspace.
	//
	// example:
	//
	// foo
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListHpoExperimentsResponseBodyExperiments) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetAccessibility(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.Accessibility = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetConfigIni(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.ConfigIni = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetConfigYml(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.ConfigYml = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetCreator(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.Creator = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetDeleted(v bool) *ListHpoExperimentsResponseBodyExperiments {
	s.Deleted = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetDescription(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetGmtCreateTime(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.GmtCreateTime = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetGmtModifiedTime(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetJobType(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.JobType = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetName(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetSearchSpace(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.SearchSpace = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetStatus(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.Status = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetTrialCount(v int32) *ListHpoExperimentsResponseBodyExperiments {
	s.TrialCount = &v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetTrialStatus(v map[string]*string) *ListHpoExperimentsResponseBodyExperiments {
	s.TrialStatus = v
	return s
}

func (s *ListHpoExperimentsResponseBodyExperiments) SetWorkspaceId(v string) *ListHpoExperimentsResponseBodyExperiments {
	s.WorkspaceId = &v
	return s
}

type ListHpoExperimentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHpoExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHpoExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHpoExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListHpoExperimentsResponse) SetHeaders(v map[string]*string) *ListHpoExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListHpoExperimentsResponse) SetStatusCode(v int32) *ListHpoExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHpoExperimentsResponse) SetBody(v *ListHpoExperimentsResponseBody) *ListHpoExperimentsResponse {
	s.Body = v
	return s
}

type ListHpoTrialCommandsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// INVALID_USER_OR_EXP
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of commands.
	Commands []*ListHpoTrialCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 071A904D-5A49-597F-9F69-81C7701D04AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHpoTrialCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHpoTrialCommandsResponseBody) SetCode(v string) *ListHpoTrialCommandsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHpoTrialCommandsResponseBody) SetCommands(v []*ListHpoTrialCommandsResponseBodyCommands) *ListHpoTrialCommandsResponseBody {
	s.Commands = v
	return s
}

func (s *ListHpoTrialCommandsResponseBody) SetDetail(v map[string]*string) *ListHpoTrialCommandsResponseBody {
	s.Detail = v
	return s
}

func (s *ListHpoTrialCommandsResponseBody) SetMessage(v string) *ListHpoTrialCommandsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHpoTrialCommandsResponseBody) SetRequestId(v string) *ListHpoTrialCommandsResponseBody {
	s.RequestId = &v
	return s
}

type ListHpoTrialCommandsResponseBodyCommands struct {
	// The command that you want to run.
	//
	// example:
	//
	// dlc submit pytorch --name=test_nni_foo_bar
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The id of the command.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The output of the above command
	//
	// example:
	//
	// dlc job submitted.
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListHpoTrialCommandsResponseBodyCommands) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *ListHpoTrialCommandsResponseBodyCommands) SetCommand(v string) *ListHpoTrialCommandsResponseBodyCommands {
	s.Command = &v
	return s
}

func (s *ListHpoTrialCommandsResponseBodyCommands) SetId(v int32) *ListHpoTrialCommandsResponseBodyCommands {
	s.Id = &v
	return s
}

func (s *ListHpoTrialCommandsResponseBodyCommands) SetOutput(v string) *ListHpoTrialCommandsResponseBodyCommands {
	s.Output = &v
	return s
}

type ListHpoTrialCommandsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHpoTrialCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHpoTrialCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListHpoTrialCommandsResponse) SetHeaders(v map[string]*string) *ListHpoTrialCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListHpoTrialCommandsResponse) SetStatusCode(v int32) *ListHpoTrialCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHpoTrialCommandsResponse) SetBody(v *ListHpoTrialCommandsResponseBody) *ListHpoTrialCommandsResponse {
	s.Body = v
	return s
}

type ListHpoTrialLogNamesResponseBody struct {
	// Error code.
	//
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Existing log files.
	LogNames []*string `json:"LogNames,omitempty" xml:"LogNames,omitempty" type:"Repeated"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHpoTrialLogNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialLogNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHpoTrialLogNamesResponseBody) SetCode(v string) *ListHpoTrialLogNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHpoTrialLogNamesResponseBody) SetDetail(v map[string]*string) *ListHpoTrialLogNamesResponseBody {
	s.Detail = v
	return s
}

func (s *ListHpoTrialLogNamesResponseBody) SetLogNames(v []*string) *ListHpoTrialLogNamesResponseBody {
	s.LogNames = v
	return s
}

func (s *ListHpoTrialLogNamesResponseBody) SetMessage(v string) *ListHpoTrialLogNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHpoTrialLogNamesResponseBody) SetRequestId(v string) *ListHpoTrialLogNamesResponseBody {
	s.RequestId = &v
	return s
}

type ListHpoTrialLogNamesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHpoTrialLogNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHpoTrialLogNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialLogNamesResponse) GoString() string {
	return s.String()
}

func (s *ListHpoTrialLogNamesResponse) SetHeaders(v map[string]*string) *ListHpoTrialLogNamesResponse {
	s.Headers = v
	return s
}

func (s *ListHpoTrialLogNamesResponse) SetStatusCode(v int32) *ListHpoTrialLogNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHpoTrialLogNamesResponse) SetBody(v *ListHpoTrialLogNamesResponseBody) *ListHpoTrialLogNamesResponse {
	s.Body = v
	return s
}

type ListHpoTrialLogsRequest struct {
	// Log file name.
	//
	// example:
	//
	// trial.log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHpoTrialLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialLogsRequest) GoString() string {
	return s.String()
}

func (s *ListHpoTrialLogsRequest) SetLogName(v string) *ListHpoTrialLogsRequest {
	s.LogName = &v
	return s
}

func (s *ListHpoTrialLogsRequest) SetPageNumber(v int32) *ListHpoTrialLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHpoTrialLogsRequest) SetPageSize(v int32) *ListHpoTrialLogsRequest {
	s.PageSize = &v
	return s
}

type ListHpoTrialLogsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	//
	// example:
	//
	// {}
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The log.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 123
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHpoTrialLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHpoTrialLogsResponseBody) SetCode(v string) *ListHpoTrialLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHpoTrialLogsResponseBody) SetDetail(v map[string]interface{}) *ListHpoTrialLogsResponseBody {
	s.Detail = v
	return s
}

func (s *ListHpoTrialLogsResponseBody) SetLogs(v []*string) *ListHpoTrialLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListHpoTrialLogsResponseBody) SetMessage(v string) *ListHpoTrialLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHpoTrialLogsResponseBody) SetRequestId(v string) *ListHpoTrialLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHpoTrialLogsResponseBody) SetTotalCount(v int32) *ListHpoTrialLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHpoTrialLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHpoTrialLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHpoTrialLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialLogsResponse) GoString() string {
	return s.String()
}

func (s *ListHpoTrialLogsResponse) SetHeaders(v map[string]*string) *ListHpoTrialLogsResponse {
	s.Headers = v
	return s
}

func (s *ListHpoTrialLogsResponse) SetStatusCode(v int32) *ListHpoTrialLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHpoTrialLogsResponse) SetBody(v *ListHpoTrialLogsResponseBody) *ListHpoTrialLogsResponse {
	s.Body = v
	return s
}

type ListHpoTrialsRequest struct {
	// The trial results order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The results to be sorted by which column.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListHpoTrialsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialsRequest) GoString() string {
	return s.String()
}

func (s *ListHpoTrialsRequest) SetOrder(v string) *ListHpoTrialsRequest {
	s.Order = &v
	return s
}

func (s *ListHpoTrialsRequest) SetPageNumber(v int32) *ListHpoTrialsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHpoTrialsRequest) SetPageSize(v int32) *ListHpoTrialsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHpoTrialsRequest) SetSortBy(v string) *ListHpoTrialsRequest {
	s.SortBy = &v
	return s
}

type ListHpoTrialsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// INVALID_USER_OR_EXP
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries qualified.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Trial array.
	Trials []*ListHpoTrialsResponseBodyTrials `json:"Trials,omitempty" xml:"Trials,omitempty" type:"Repeated"`
}

func (s ListHpoTrialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHpoTrialsResponseBody) SetCode(v string) *ListHpoTrialsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHpoTrialsResponseBody) SetDetail(v map[string]*string) *ListHpoTrialsResponseBody {
	s.Detail = v
	return s
}

func (s *ListHpoTrialsResponseBody) SetMessage(v string) *ListHpoTrialsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHpoTrialsResponseBody) SetRequestId(v string) *ListHpoTrialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHpoTrialsResponseBody) SetTotalCount(v int32) *ListHpoTrialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHpoTrialsResponseBody) SetTrials(v []*ListHpoTrialsResponseBodyTrials) *ListHpoTrialsResponseBody {
	s.Trials = v
	return s
}

type ListHpoTrialsResponseBodyTrials struct {
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	FinalMetric     *string `json:"FinalMetric,omitempty" xml:"FinalMetric,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Hyperparam      *string `json:"Hyperparam,omitempty" xml:"Hyperparam,omitempty"`
	JobMeta         *string `json:"JobMeta,omitempty" xml:"JobMeta,omitempty"`
	Metric          *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	MetricName      *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ParameterId     *int32  `json:"ParameterId,omitempty" xml:"ParameterId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TrialId         *string `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	UserComment     *string `json:"UserComment,omitempty" xml:"UserComment,omitempty"`
	UserScore       *int32  `json:"UserScore,omitempty" xml:"UserScore,omitempty"`
}

func (s ListHpoTrialsResponseBodyTrials) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialsResponseBodyTrials) GoString() string {
	return s.String()
}

func (s *ListHpoTrialsResponseBodyTrials) SetExperimentId(v string) *ListHpoTrialsResponseBodyTrials {
	s.ExperimentId = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetFinalMetric(v string) *ListHpoTrialsResponseBodyTrials {
	s.FinalMetric = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetGmtCreateTime(v string) *ListHpoTrialsResponseBodyTrials {
	s.GmtCreateTime = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetGmtModifiedTime(v string) *ListHpoTrialsResponseBodyTrials {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetHyperparam(v string) *ListHpoTrialsResponseBodyTrials {
	s.Hyperparam = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetJobMeta(v string) *ListHpoTrialsResponseBodyTrials {
	s.JobMeta = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetMetric(v string) *ListHpoTrialsResponseBodyTrials {
	s.Metric = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetMetricName(v string) *ListHpoTrialsResponseBodyTrials {
	s.MetricName = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetModel(v string) *ListHpoTrialsResponseBodyTrials {
	s.Model = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetParameterId(v int32) *ListHpoTrialsResponseBodyTrials {
	s.ParameterId = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetStatus(v string) *ListHpoTrialsResponseBodyTrials {
	s.Status = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetTrialId(v string) *ListHpoTrialsResponseBodyTrials {
	s.TrialId = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetUserComment(v string) *ListHpoTrialsResponseBodyTrials {
	s.UserComment = &v
	return s
}

func (s *ListHpoTrialsResponseBodyTrials) SetUserScore(v int32) *ListHpoTrialsResponseBodyTrials {
	s.UserScore = &v
	return s
}

type ListHpoTrialsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHpoTrialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHpoTrialsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHpoTrialsResponse) GoString() string {
	return s.String()
}

func (s *ListHpoTrialsResponse) SetHeaders(v map[string]*string) *ListHpoTrialsResponse {
	s.Headers = v
	return s
}

func (s *ListHpoTrialsResponse) SetStatusCode(v int32) *ListHpoTrialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHpoTrialsResponse) SetBody(v *ListHpoTrialsResponseBody) *ListHpoTrialsResponse {
	s.Body = v
	return s
}

type RestartHpoTrialsRequest struct {
	// A hyperparameter combination instance.
	//
	// example:
	//
	// ‘{"${batch_size}": "32", "${lr}": "0.01"}\\"
	TrialHyperParameters *string `json:"TrialHyperParameters,omitempty" xml:"TrialHyperParameters,omitempty"`
	// Trial ID array.
	TrialIds []*string `json:"TrialIds,omitempty" xml:"TrialIds,omitempty" type:"Repeated"`
}

func (s RestartHpoTrialsRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartHpoTrialsRequest) GoString() string {
	return s.String()
}

func (s *RestartHpoTrialsRequest) SetTrialHyperParameters(v string) *RestartHpoTrialsRequest {
	s.TrialHyperParameters = &v
	return s
}

func (s *RestartHpoTrialsRequest) SetTrialIds(v []*string) *RestartHpoTrialsRequest {
	s.TrialIds = v
	return s
}

type RestartHpoTrialsResponseBody struct {
	// Response code.
	//
	// example:
	//
	// INVALID_USER_OR_EXP
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Restart status of all the trial IDs.
	//
	// example:
	//
	// {
	//
	// \\"xrYq99\\": TrialJobInfo text,
	//
	// \\"xrYq9N\\":TrialJobInfo text
	//
	// }
	Results map[string]interface{} `json:"Results,omitempty" xml:"Results,omitempty"`
}

func (s RestartHpoTrialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartHpoTrialsResponseBody) GoString() string {
	return s.String()
}

func (s *RestartHpoTrialsResponseBody) SetCode(v string) *RestartHpoTrialsResponseBody {
	s.Code = &v
	return s
}

func (s *RestartHpoTrialsResponseBody) SetDetail(v map[string]*string) *RestartHpoTrialsResponseBody {
	s.Detail = v
	return s
}

func (s *RestartHpoTrialsResponseBody) SetMessage(v string) *RestartHpoTrialsResponseBody {
	s.Message = &v
	return s
}

func (s *RestartHpoTrialsResponseBody) SetRequestId(v string) *RestartHpoTrialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartHpoTrialsResponseBody) SetResults(v map[string]interface{}) *RestartHpoTrialsResponseBody {
	s.Results = v
	return s
}

type RestartHpoTrialsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartHpoTrialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartHpoTrialsResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartHpoTrialsResponse) GoString() string {
	return s.String()
}

func (s *RestartHpoTrialsResponse) SetHeaders(v map[string]*string) *RestartHpoTrialsResponse {
	s.Headers = v
	return s
}

func (s *RestartHpoTrialsResponse) SetStatusCode(v int32) *RestartHpoTrialsResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartHpoTrialsResponse) SetBody(v *RestartHpoTrialsResponseBody) *RestartHpoTrialsResponse {
	s.Body = v
	return s
}

type StopHpoExperimentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// INVALID_INPUT_PARAMS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Experiment ID
	//
	// example:
	//
	// sHpITOH
	ExpId *string `json:"ExpId,omitempty" xml:"ExpId,omitempty"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 071A904D-5A49-597F-9F69-81C7701D04AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StopHpoExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopHpoExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *StopHpoExperimentResponseBody) SetCode(v string) *StopHpoExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *StopHpoExperimentResponseBody) SetDetail(v map[string]*string) *StopHpoExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *StopHpoExperimentResponseBody) SetExpId(v string) *StopHpoExperimentResponseBody {
	s.ExpId = &v
	return s
}

func (s *StopHpoExperimentResponseBody) SetMessage(v string) *StopHpoExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *StopHpoExperimentResponseBody) SetRequestId(v string) *StopHpoExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopHpoExperimentResponseBody) SetStatus(v string) *StopHpoExperimentResponseBody {
	s.Status = &v
	return s
}

type StopHpoExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopHpoExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopHpoExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s StopHpoExperimentResponse) GoString() string {
	return s.String()
}

func (s *StopHpoExperimentResponse) SetHeaders(v map[string]*string) *StopHpoExperimentResponse {
	s.Headers = v
	return s
}

func (s *StopHpoExperimentResponse) SetStatusCode(v int32) *StopHpoExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopHpoExperimentResponse) SetBody(v *StopHpoExperimentResponseBody) *StopHpoExperimentResponse {
	s.Body = v
	return s
}

type StopHpoTrialsRequest struct {
	// Trial Ids to be stopped.
	TrialIds []*string `json:"TrialIds,omitempty" xml:"TrialIds,omitempty" type:"Repeated"`
}

func (s StopHpoTrialsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopHpoTrialsRequest) GoString() string {
	return s.String()
}

func (s *StopHpoTrialsRequest) SetTrialIds(v []*string) *StopHpoTrialsRequest {
	s.TrialIds = v
	return s
}

type StopHpoTrialsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// INVALID_USER_OR_EXP
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	Detail map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Error message.
	//
	// example:
	//
	// failed to stop any trial
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 071A904D-5A49-597F-9F69-81C7701D04AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of trial stop.
	//
	// example:
	//
	// {
	//
	//     \\"xrYq99\\":\\"successfully stopped trial,\\",
	//
	//     \\"xrYq9N\\":\\"failed to stop trial,\\"
	//
	// }
	Results map[string]interface{} `json:"Results,omitempty" xml:"Results,omitempty"`
}

func (s StopHpoTrialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopHpoTrialsResponseBody) GoString() string {
	return s.String()
}

func (s *StopHpoTrialsResponseBody) SetCode(v string) *StopHpoTrialsResponseBody {
	s.Code = &v
	return s
}

func (s *StopHpoTrialsResponseBody) SetDetail(v map[string]*string) *StopHpoTrialsResponseBody {
	s.Detail = v
	return s
}

func (s *StopHpoTrialsResponseBody) SetMessage(v string) *StopHpoTrialsResponseBody {
	s.Message = &v
	return s
}

func (s *StopHpoTrialsResponseBody) SetRequestId(v string) *StopHpoTrialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopHpoTrialsResponseBody) SetResults(v map[string]interface{}) *StopHpoTrialsResponseBody {
	s.Results = v
	return s
}

type StopHpoTrialsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopHpoTrialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopHpoTrialsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopHpoTrialsResponse) GoString() string {
	return s.String()
}

func (s *StopHpoTrialsResponse) SetHeaders(v map[string]*string) *StopHpoTrialsResponse {
	s.Headers = v
	return s
}

func (s *StopHpoTrialsResponse) SetStatusCode(v int32) *StopHpoTrialsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopHpoTrialsResponse) SetBody(v *StopHpoTrialsResponseBody) *StopHpoTrialsResponse {
	s.Body = v
	return s
}

type UpdateHpoExperimentRequest struct {
	// Experiment accessibility, private or public.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Experiment description.
	//
	// example:
	//
	// This is an AutoML HPO experiment.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Experiment configuration.
	HpoExperimentConfiguration *HpoExperimentConfig `json:"HpoExperimentConfiguration,omitempty" xml:"HpoExperimentConfiguration,omitempty"`
	// Experiment name.
	//
	// example:
	//
	// my experiment x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Expeirment\\"s AI workspace ID.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateHpoExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHpoExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateHpoExperimentRequest) SetAccessibility(v string) *UpdateHpoExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateHpoExperimentRequest) SetDescription(v string) *UpdateHpoExperimentRequest {
	s.Description = &v
	return s
}

func (s *UpdateHpoExperimentRequest) SetHpoExperimentConfiguration(v *HpoExperimentConfig) *UpdateHpoExperimentRequest {
	s.HpoExperimentConfiguration = v
	return s
}

func (s *UpdateHpoExperimentRequest) SetName(v string) *UpdateHpoExperimentRequest {
	s.Name = &v
	return s
}

func (s *UpdateHpoExperimentRequest) SetWorkspaceId(v string) *UpdateHpoExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateHpoExperimentResponseBody struct {
	// Error code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Extra error information.
	//
	// example:
	//
	// {}
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Error message.
	//
	// example:
	//
	// Missing \\"user_id\\" in request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F190916-B3E5-5D1E-AD0C-35C0DF105F51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHpoExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHpoExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHpoExperimentResponseBody) SetCode(v string) *UpdateHpoExperimentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHpoExperimentResponseBody) SetDetail(v map[string]interface{}) *UpdateHpoExperimentResponseBody {
	s.Detail = v
	return s
}

func (s *UpdateHpoExperimentResponseBody) SetMessage(v string) *UpdateHpoExperimentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHpoExperimentResponseBody) SetRequestId(v string) *UpdateHpoExperimentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHpoExperimentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHpoExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHpoExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHpoExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateHpoExperimentResponse) SetHeaders(v map[string]*string) *UpdateHpoExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateHpoExperimentResponse) SetStatusCode(v int32) *UpdateHpoExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHpoExperimentResponse) SetBody(v *UpdateHpoExperimentResponseBody) *UpdateHpoExperimentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("paiautoml"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CreateAutofeExperiment
//
// @param request - CreateAutofeExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutofeExperimentResponse
func (client *Client) CreateAutofeExperimentWithOptions(request *CreateAutofeExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAutofeExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.AutofeExperimentConfiguration)) {
		body["AutofeExperimentConfiguration"] = request.AutofeExperimentConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutofeExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/autofe/experiment"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAutofeExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CreateAutofeExperiment
//
// @param request - CreateAutofeExperimentRequest
//
// @return CreateAutofeExperimentResponse
func (client *Client) CreateAutofeExperiment(request *CreateAutofeExperimentRequest) (_result *CreateAutofeExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAutofeExperimentResponse{}
	_body, _err := client.CreateAutofeExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create an HyperParameter Optimization experiment.
//
// @param request - CreateHpoExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHpoExperimentResponse
func (client *Client) CreateHpoExperimentWithOptions(request *CreateHpoExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHpoExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HpoExperimentConfiguration)) {
		body["HpoExperimentConfiguration"] = request.HpoExperimentConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHpoExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHpoExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an HyperParameter Optimization experiment.
//
// @param request - CreateHpoExperimentRequest
//
// @return CreateHpoExperimentResponse
func (client *Client) CreateHpoExperiment(request *CreateHpoExperimentRequest) (_result *CreateHpoExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHpoExperimentResponse{}
	_body, _err := client.CreateHpoExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CreateServiceIdentityRole
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRoleWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceIdentityRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceIdentityRole"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/serviceidentityrole"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CreateServiceIdentityRole
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRole() (_result *CreateServiceIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CreateServiceIdentityRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete an HPO experiment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHpoExperimentResponse
func (client *Client) DeleteHpoExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteHpoExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHpoExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/delete"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHpoExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an HPO experiment
//
// @return DeleteHpoExperimentResponse
func (client *Client) DeleteHpoExperiment(ExperimentId *string) (_result *DeleteHpoExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHpoExperimentResponse{}
	_body, _err := client.DeleteHpoExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get AutoFE Experiment。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutofeExperimentResponse
func (client *Client) GetAutofeExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAutofeExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutofeExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/autofe/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutofeExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get AutoFE Experiment。
//
// @return GetAutofeExperimentResponse
func (client *Client) GetAutofeExperiment(ExperimentId *string) (_result *GetAutofeExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAutofeExperimentResponse{}
	_body, _err := client.GetAutofeExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get hpo experiment by user id and exp id
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHpoExperimentResponse
func (client *Client) GetHpoExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHpoExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHpoExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHpoExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get hpo experiment by user id and exp id
//
// @return GetHpoExperimentResponse
func (client *Client) GetHpoExperiment(ExperimentId *string) (_result *GetHpoExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHpoExperimentResponse{}
	_body, _err := client.GetHpoExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get trial detail information
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHpoTrialResponse
func (client *Client) GetHpoTrialWithOptions(ExperimentId *string, TrialId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHpoTrialResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHpoTrial"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/trial/" + tea.StringValue(openapiutil.GetEncodeParam(TrialId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHpoTrialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get trial detail information
//
// @return GetHpoTrialResponse
func (client *Client) GetHpoTrial(ExperimentId *string, TrialId *string) (_result *GetHpoTrialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHpoTrialResponse{}
	_body, _err := client.GetHpoTrialWithOptions(ExperimentId, TrialId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetServiceIdentityRole, return role name if SLR exists, empty otherwise
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRoleWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceIdentityRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceIdentityRole"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/serviceidentityrole"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetServiceIdentityRole, return role name if SLR exists, empty otherwise
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRole() (_result *GetServiceIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.GetServiceIdentityRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list the content of a specified hpo experiment log
//
// @param request - ListHpoExperimentLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHpoExperimentLogsResponse
func (client *Client) ListHpoExperimentLogsWithOptions(ExperimentId *string, request *ListHpoExperimentLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHpoExperimentLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogName)) {
		query["LogName"] = request.LogName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHpoExperimentLogs"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHpoExperimentLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list the content of a specified hpo experiment log
//
// @param request - ListHpoExperimentLogsRequest
//
// @return ListHpoExperimentLogsResponse
func (client *Client) ListHpoExperimentLogs(ExperimentId *string, request *ListHpoExperimentLogsRequest) (_result *ListHpoExperimentLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHpoExperimentLogsResponse{}
	_body, _err := client.ListHpoExperimentLogsWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List HPO experiments
//
// @param request - ListHpoExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHpoExperimentsResponse
func (client *Client) ListHpoExperimentsWithOptions(request *ListHpoExperimentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHpoExperimentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeConfigData)) {
		query["IncludeConfigData"] = request.IncludeConfigData
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCreateTime)) {
		query["MaxCreateTime"] = request.MaxCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinCreateTime)) {
		query["MinCreateTime"] = request.MinCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHpoExperiments"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHpoExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List HPO experiments
//
// @param request - ListHpoExperimentsRequest
//
// @return ListHpoExperimentsResponse
func (client *Client) ListHpoExperiments(request *ListHpoExperimentsRequest) (_result *ListHpoExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHpoExperimentsResponse{}
	_body, _err := client.ListHpoExperimentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 返回一个trial所对应的任务里所有已经执行的命令
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHpoTrialCommandsResponse
func (client *Client) ListHpoTrialCommandsWithOptions(ExperimentId *string, TrialId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHpoTrialCommandsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListHpoTrialCommands"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/trial/" + tea.StringValue(openapiutil.GetEncodeParam(TrialId)) + "/commands"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHpoTrialCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 返回一个trial所对应的任务里所有已经执行的命令
//
// @return ListHpoTrialCommandsResponse
func (client *Client) ListHpoTrialCommands(ExperimentId *string, TrialId *string) (_result *ListHpoTrialCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHpoTrialCommandsResponse{}
	_body, _err := client.ListHpoTrialCommandsWithOptions(ExperimentId, TrialId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List all log file names a trial have.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHpoTrialLogNamesResponse
func (client *Client) ListHpoTrialLogNamesWithOptions(ExperimentId *string, TrialId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHpoTrialLogNamesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListHpoTrialLogNames"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/trial/" + tea.StringValue(openapiutil.GetEncodeParam(TrialId)) + "/lognames"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHpoTrialLogNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List all log file names a trial have.
//
// @return ListHpoTrialLogNamesResponse
func (client *Client) ListHpoTrialLogNames(ExperimentId *string, TrialId *string) (_result *ListHpoTrialLogNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHpoTrialLogNamesResponse{}
	_body, _err := client.ListHpoTrialLogNamesWithOptions(ExperimentId, TrialId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List Trial log lines
//
// @param request - ListHpoTrialLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHpoTrialLogsResponse
func (client *Client) ListHpoTrialLogsWithOptions(ExperimentId *string, TrialId *string, request *ListHpoTrialLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHpoTrialLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogName)) {
		query["LogName"] = request.LogName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHpoTrialLogs"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/trial/" + tea.StringValue(openapiutil.GetEncodeParam(TrialId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHpoTrialLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List Trial log lines
//
// @param request - ListHpoTrialLogsRequest
//
// @return ListHpoTrialLogsResponse
func (client *Client) ListHpoTrialLogs(ExperimentId *string, TrialId *string, request *ListHpoTrialLogsRequest) (_result *ListHpoTrialLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHpoTrialLogsResponse{}
	_body, _err := client.ListHpoTrialLogsWithOptions(ExperimentId, TrialId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List HPO trials
//
// @param request - ListHpoTrialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHpoTrialsResponse
func (client *Client) ListHpoTrialsWithOptions(ExperimentId *string, request *ListHpoTrialsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHpoTrialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHpoTrials"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/trials"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHpoTrialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List HPO trials
//
// @param request - ListHpoTrialsRequest
//
// @return ListHpoTrialsResponse
func (client *Client) ListHpoTrials(ExperimentId *string, request *ListHpoTrialsRequest) (_result *ListHpoTrialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHpoTrialsResponse{}
	_body, _err := client.ListHpoTrialsWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restart HPO trials
//
// @param request - RestartHpoTrialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartHpoTrialsResponse
func (client *Client) RestartHpoTrialsWithOptions(ExperimentId *string, request *RestartHpoTrialsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartHpoTrialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TrialHyperParameters)) {
		body["TrialHyperParameters"] = request.TrialHyperParameters
	}

	if !tea.BoolValue(util.IsUnset(request.TrialIds)) {
		body["TrialIds"] = request.TrialIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartHpoTrials"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/restart_trials"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartHpoTrialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restart HPO trials
//
// @param request - RestartHpoTrialsRequest
//
// @return RestartHpoTrialsResponse
func (client *Client) RestartHpoTrials(ExperimentId *string, request *RestartHpoTrialsRequest) (_result *RestartHpoTrialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartHpoTrialsResponse{}
	_body, _err := client.RestartHpoTrialsWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// calling hpo StopExperiment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopHpoExperimentResponse
func (client *Client) StopHpoExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopHpoExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopHpoExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopHpoExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// calling hpo StopExperiment
//
// @return StopHpoExperimentResponse
func (client *Client) StopHpoExperiment(ExperimentId *string) (_result *StopHpoExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopHpoExperimentResponse{}
	_body, _err := client.StopHpoExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stop an HPO trial.
//
// @param request - StopHpoTrialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopHpoTrialsResponse
func (client *Client) StopHpoTrialsWithOptions(ExperimentId *string, request *StopHpoTrialsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopHpoTrialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TrialIds)) {
		body["TrialIds"] = request.TrialIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopHpoTrials"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/stop_trials"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopHpoTrialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop an HPO trial.
//
// @param request - StopHpoTrialsRequest
//
// @return StopHpoTrialsResponse
func (client *Client) StopHpoTrials(ExperimentId *string, request *StopHpoTrialsRequest) (_result *StopHpoTrialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopHpoTrialsResponse{}
	_body, _err := client.StopHpoTrialsWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update a running HPO experiment\\"s configuration
//
// @param request - UpdateHpoExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHpoExperimentResponse
func (client *Client) UpdateHpoExperimentWithOptions(ExperimentId *string, request *UpdateHpoExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateHpoExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HpoExperimentConfiguration)) {
		body["HpoExperimentConfiguration"] = request.HpoExperimentConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHpoExperiment"),
		Version:     tea.String("2022-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/v1/hpo/experiment/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/update"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHpoExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a running HPO experiment\\"s configuration
//
// @param request - UpdateHpoExperimentRequest
//
// @return UpdateHpoExperimentResponse
func (client *Client) UpdateHpoExperiment(ExperimentId *string, request *UpdateHpoExperimentRequest) (_result *UpdateHpoExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHpoExperimentResponse{}
	_body, _err := client.UpdateHpoExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
