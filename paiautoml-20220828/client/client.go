// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AutofeExperimentConfiguration struct {
	OdpsConfig *AutofeExperimentConfigurationOdpsConfig `json:"odps_config,omitempty" xml:"odps_config,omitempty" type:"Struct"`
	OssConfig  *AutofeExperimentConfigurationOssConfig  `json:"oss_config,omitempty" xml:"oss_config,omitempty" type:"Struct"`
	YmlConfig  *AutofeExperimentConfigurationYmlConfig  `json:"yml_config,omitempty" xml:"yml_config,omitempty" type:"Struct"`
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
	OdpsAccessId    *string `json:"odps_access_id,omitempty" xml:"odps_access_id,omitempty"`
	OdpsAccessKey   *string `json:"odps_access_key,omitempty" xml:"odps_access_key,omitempty"`
	OdpsEndpoint    *string `json:"odps_endpoint,omitempty" xml:"odps_endpoint,omitempty"`
	OdpsProjectName *string `json:"odps_project_name,omitempty" xml:"odps_project_name,omitempty"`
	OdpsRegionId    *string `json:"odps_region_id,omitempty" xml:"odps_region_id,omitempty"`
	OdpsRoleArn     *string `json:"odps_role_arn,omitempty" xml:"odps_role_arn,omitempty"`
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
	OssBucket    *string `json:"oss_bucket,omitempty" xml:"oss_bucket,omitempty"`
	OssEndpoint  *string `json:"oss_endpoint,omitempty" xml:"oss_endpoint,omitempty"`
	OssRoleArn   *string `json:"oss_role_arn,omitempty" xml:"oss_role_arn,omitempty"`
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
	Action             *string `json:"action,omitempty" xml:"action,omitempty"`
	AggregateOnly      *string `json:"aggregate_only,omitempty" xml:"aggregate_only,omitempty"`
	AnalyzeExpId       *string `json:"analyze_exp_id,omitempty" xml:"analyze_exp_id,omitempty"`
	Cpu                *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	DataPartition      *string `json:"data_partition,omitempty" xml:"data_partition,omitempty"`
	DataSource         *string `json:"data_source,omitempty" xml:"data_source,omitempty"`
	DataType           *string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	DebugMode          *string `json:"debug_mode,omitempty" xml:"debug_mode,omitempty"`
	ExcludeColumns     *string `json:"exclude_columns,omitempty" xml:"exclude_columns,omitempty"`
	FeatureSelection   *string `json:"feature_selection,omitempty" xml:"feature_selection,omitempty"`
	FilterThresh       *string `json:"filter_thresh,omitempty" xml:"filter_thresh,omitempty"`
	IvThresh           *string `json:"iv_thresh,omitempty" xml:"iv_thresh,omitempty"`
	Label              *string `json:"label,omitempty" xml:"label,omitempty"`
	Memory             *string `json:"memory,omitempty" xml:"memory,omitempty"`
	OutputConfigOssDir *string `json:"output_config_oss_dir,omitempty" xml:"output_config_oss_dir,omitempty"`
	PipelineExpId      *string `json:"pipeline_exp_id,omitempty" xml:"pipeline_exp_id,omitempty"`
	ReuseResults       *string `json:"reuse_results,omitempty" xml:"reuse_results,omitempty"`
	SampleRatio        *string `json:"sample_ratio,omitempty" xml:"sample_ratio,omitempty"`
	SampleSize         *string `json:"sample_size,omitempty" xml:"sample_size,omitempty"`
	SelectionExpId     *string `json:"selection_exp_id,omitempty" xml:"selection_exp_id,omitempty"`
	SkipSelect         *string `json:"skip_select,omitempty" xml:"skip_select,omitempty"`
	Workers            *string `json:"workers,omitempty" xml:"workers,omitempty"`
	WorkspaceName      *string `json:"workspace_name,omitempty" xml:"workspace_name,omitempty"`
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
	DlcConfig      *HpoExperimentConfigDlcConfig      `json:"dlc_config,omitempty" xml:"dlc_config,omitempty" type:"Struct"`
	K8sConfig      *HpoExperimentConfigK8sConfig      `json:"k8s_config,omitempty" xml:"k8s_config,omitempty" type:"Struct"`
	MetricConfig   *HpoExperimentConfigMetricConfig   `json:"metric_config,omitempty" xml:"metric_config,omitempty" type:"Struct"`
	MonitorConfig  *HpoExperimentConfigMonitorConfig  `json:"monitor_config,omitempty" xml:"monitor_config,omitempty" type:"Struct"`
	OdpsConfig     *HpoExperimentConfigOdpsConfig     `json:"odps_config,omitempty" xml:"odps_config,omitempty" type:"Struct"`
	OssConfig      *HpoExperimentConfigOssConfig      `json:"oss_config,omitempty" xml:"oss_config,omitempty" type:"Struct"`
	OutputConfig   *HpoExperimentConfigOutputConfig   `json:"output_config,omitempty" xml:"output_config,omitempty" type:"Struct"`
	PaiflowConfig  *HpoExperimentConfigPaiflowConfig  `json:"paiflow_config,omitempty" xml:"paiflow_config,omitempty" type:"Struct"`
	ParamsConfig   *HpoExperimentConfigParamsConfig   `json:"params_config,omitempty" xml:"params_config,omitempty" type:"Struct"`
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
	FinalMode           *string                `json:"final_mode,omitempty" xml:"final_mode,omitempty"`
	MetricDict          map[string]interface{} `json:"metric_dict,omitempty" xml:"metric_dict,omitempty"`
	MetricSource        []*string              `json:"metric_source,omitempty" xml:"metric_source,omitempty" type:"Repeated"`
	MetricType          *string                `json:"metric_type,omitempty" xml:"metric_type,omitempty"`
	SourceListFinalMode *string                `json:"source_list_final_mode,omitempty" xml:"source_list_final_mode,omitempty"`
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
	Cmd    []*string `json:"cmd,omitempty" xml:"cmd,omitempty" type:"Repeated"`
	Name   *string   `json:"name,omitempty" xml:"name,omitempty"`
	Resume *string   `json:"resume,omitempty" xml:"resume,omitempty"`
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

type CreateHpoExperimentRequest struct {
	Accessibility              *string              `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description                *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	HpoExperimentConfiguration *HpoExperimentConfig `json:"HpoExperimentConfiguration,omitempty" xml:"HpoExperimentConfiguration,omitempty"`
	Name                       *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkspaceId                *string              `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	Code         *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	Detail       map[string]*string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	ExperimentId *string            `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Message      *string            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
