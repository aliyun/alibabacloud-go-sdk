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

type ErrorResponse struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s ErrorResponse) GoString() string {
	return s.String()
}

func (s *ErrorResponse) SetCode(v string) *ErrorResponse {
	s.Code = &v
	return s
}

func (s *ErrorResponse) SetMessage(v string) *ErrorResponse {
	s.Message = &v
	return s
}

func (s *ErrorResponse) SetRequestId(v string) *ErrorResponse {
	s.RequestId = &v
	return s
}

type VariablesValue struct {
	// Specifies whether the variable cannot be modified.
	DisableModify *bool `json:"disableModify,omitempty" xml:"disableModify,omitempty"`
	// Specifies whether the variable is modified.
	IsModify *bool `json:"isModify,omitempty" xml:"isModify,omitempty"`
	// The value of the variable.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// The description about the variable.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The value of the template.
	TemplateValue *string `json:"templateValue,omitempty" xml:"templateValue,omitempty"`
	// The type of the variable. Valid values:
	//
	// *   NORMAL: a normal variable
	// *   FUNCTION: a function variable
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The function variable.
	FuncValue *VariablesValueFuncValue `json:"funcValue,omitempty" xml:"funcValue,omitempty" type:"Struct"`
}

func (s VariablesValue) String() string {
	return tea.Prettify(s)
}

func (s VariablesValue) GoString() string {
	return s.String()
}

func (s *VariablesValue) SetDisableModify(v bool) *VariablesValue {
	s.DisableModify = &v
	return s
}

func (s *VariablesValue) SetIsModify(v bool) *VariablesValue {
	s.IsModify = &v
	return s
}

func (s *VariablesValue) SetValue(v string) *VariablesValue {
	s.Value = &v
	return s
}

func (s *VariablesValue) SetDescription(v string) *VariablesValue {
	s.Description = &v
	return s
}

func (s *VariablesValue) SetTemplateValue(v string) *VariablesValue {
	s.TemplateValue = &v
	return s
}

func (s *VariablesValue) SetType(v string) *VariablesValue {
	s.Type = &v
	return s
}

func (s *VariablesValue) SetFuncValue(v *VariablesValueFuncValue) *VariablesValue {
	s.FuncValue = v
	return s
}

type VariablesValueFuncValue struct {
	// The class name.
	FuncClassName *string `json:"funcClassName,omitempty" xml:"funcClassName,omitempty"`
	// The template of the variable.
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s VariablesValueFuncValue) String() string {
	return tea.Prettify(s)
}

func (s VariablesValueFuncValue) GoString() string {
	return s.String()
}

func (s *VariablesValueFuncValue) SetFuncClassName(v string) *VariablesValueFuncValue {
	s.FuncClassName = &v
	return s
}

func (s *VariablesValueFuncValue) SetTemplate(v string) *VariablesValueFuncValue {
	s.Template = &v
	return s
}

type BuildIndexRequest struct {
	// The mode in which reindexing is performed.
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The name of the data source.
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The type of the data source.
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// The timestamp in seconds. This parameter is required if you import data from the data source by calling API operations.
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data center in which the data source resides.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the generation.
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The data partition. This parameter is required if the dataSourceType parameter is set to odps.
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s BuildIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s BuildIndexRequest) GoString() string {
	return s.String()
}

func (s *BuildIndexRequest) SetBuildMode(v string) *BuildIndexRequest {
	s.BuildMode = &v
	return s
}

func (s *BuildIndexRequest) SetDataSourceName(v string) *BuildIndexRequest {
	s.DataSourceName = &v
	return s
}

func (s *BuildIndexRequest) SetDataSourceType(v string) *BuildIndexRequest {
	s.DataSourceType = &v
	return s
}

func (s *BuildIndexRequest) SetDataTimeSec(v int32) *BuildIndexRequest {
	s.DataTimeSec = &v
	return s
}

func (s *BuildIndexRequest) SetDomain(v string) *BuildIndexRequest {
	s.Domain = &v
	return s
}

func (s *BuildIndexRequest) SetGeneration(v int64) *BuildIndexRequest {
	s.Generation = &v
	return s
}

func (s *BuildIndexRequest) SetPartition(v string) *BuildIndexRequest {
	s.Partition = &v
	return s
}

type BuildIndexResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BuildIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuildIndexResponseBody) GoString() string {
	return s.String()
}

func (s *BuildIndexResponseBody) SetRequestId(v string) *BuildIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildIndexResponseBody) SetResult(v map[string]interface{}) *BuildIndexResponseBody {
	s.Result = v
	return s
}

type BuildIndexResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BuildIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BuildIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s BuildIndexResponse) GoString() string {
	return s.String()
}

func (s *BuildIndexResponse) SetHeaders(v map[string]*string) *BuildIndexResponse {
	s.Headers = v
	return s
}

func (s *BuildIndexResponse) SetStatusCode(v int32) *BuildIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildIndexResponse) SetBody(v *BuildIndexResponseBody) *BuildIndexResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// Specifies whether to automatically balance the load between QRS workers.
	AutoLoad *bool `json:"autoLoad,omitempty" xml:"autoLoad,omitempty"`
	// The information about Searcher workers.
	DataNode *CreateClusterRequestDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the cluster.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The information about Query Result Searcher (QRS) workers.
	QueryNode *CreateClusterRequestQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetAutoLoad(v bool) *CreateClusterRequest {
	s.AutoLoad = &v
	return s
}

func (s *CreateClusterRequest) SetDataNode(v *CreateClusterRequestDataNode) *CreateClusterRequest {
	s.DataNode = v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetQueryNode(v *CreateClusterRequestQueryNode) *CreateClusterRequest {
	s.QueryNode = v
	return s
}

type CreateClusterRequestDataNode struct {
	// The number of Searcher workers.
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
}

func (s CreateClusterRequestDataNode) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestDataNode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestDataNode) SetNumber(v int32) *CreateClusterRequestDataNode {
	s.Number = &v
	return s
}

type CreateClusterRequestQueryNode struct {
	// The number of QRS workers.
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
}

func (s CreateClusterRequestQueryNode) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestQueryNode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestQueryNode) SetNumber(v int32) *CreateClusterRequestQueryNode {
	s.Number = &v
	return s
}

type CreateClusterResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetResult(v map[string]interface{}) *CreateClusterResponseBody {
	s.Result = v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	AutoBuildIndex *bool                              `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config         *CreateDataSourceRequestConfig     `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	Domain         *string                            `json:"domain,omitempty" xml:"domain,omitempty"`
	Name           *string                            `json:"name,omitempty" xml:"name,omitempty"`
	SaroConfig     *CreateDataSourceRequestSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	Type           *string                            `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetAutoBuildIndex(v bool) *CreateDataSourceRequest {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateDataSourceRequest) SetConfig(v *CreateDataSourceRequestConfig) *CreateDataSourceRequest {
	s.Config = v
	return s
}

func (s *CreateDataSourceRequest) SetDomain(v string) *CreateDataSourceRequest {
	s.Domain = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetSaroConfig(v *CreateDataSourceRequestSaroConfig) *CreateDataSourceRequest {
	s.SaroConfig = v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequest) SetDryRun(v bool) *CreateDataSourceRequest {
	s.DryRun = &v
	return s
}

type CreateDataSourceRequestConfig struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Bucket       *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Namespace    *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OssPath      *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition    *string `json:"partition,omitempty" xml:"partition,omitempty"`
	Path         *string `json:"path,omitempty" xml:"path,omitempty"`
	Project      *string `json:"project,omitempty" xml:"project,omitempty"`
	Table        *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateDataSourceRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestConfig) SetAccessKey(v string) *CreateDataSourceRequestConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetAccessSecret(v string) *CreateDataSourceRequestConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetBucket(v string) *CreateDataSourceRequestConfig {
	s.Bucket = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetEndpoint(v string) *CreateDataSourceRequestConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetNamespace(v string) *CreateDataSourceRequestConfig {
	s.Namespace = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetOssPath(v string) *CreateDataSourceRequestConfig {
	s.OssPath = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetPartition(v string) *CreateDataSourceRequestConfig {
	s.Partition = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetPath(v string) *CreateDataSourceRequestConfig {
	s.Path = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetProject(v string) *CreateDataSourceRequestConfig {
	s.Project = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetTable(v string) *CreateDataSourceRequestConfig {
	s.Table = &v
	return s
}

type CreateDataSourceRequestSaroConfig struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s CreateDataSourceRequestSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestSaroConfig) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestSaroConfig) SetNamespace(v string) *CreateDataSourceRequestSaroConfig {
	s.Namespace = &v
	return s
}

func (s *CreateDataSourceRequestSaroConfig) SetTableName(v string) *CreateDataSourceRequestSaroConfig {
	s.TableName = &v
	return s
}

type CreateDataSourceResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetResult(v map[string]interface{}) *CreateDataSourceResponseBody {
	s.Result = v
	return s
}

type CreateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetStatusCode(v int32) *CreateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type CreateIndexRequest struct {
	// The content of the index.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Optional. The data source, which can be MaxCompute, Message Service (MNS), Realtime Compute for Apache Flink, or StreamCompute.
	DataSource     *string                           `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	DataSourceInfo *CreateIndexRequestDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The data center in which the data source resides.
	Domain *string                `json:"domain,omitempty" xml:"domain,omitempty"`
	Extend map[string]interface{} `json:"extend,omitempty" xml:"extend,omitempty"`
	// The name of the index.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data partition.
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	DryRun    *bool  `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) SetContent(v string) *CreateIndexRequest {
	s.Content = &v
	return s
}

func (s *CreateIndexRequest) SetDataSource(v string) *CreateIndexRequest {
	s.DataSource = &v
	return s
}

func (s *CreateIndexRequest) SetDataSourceInfo(v *CreateIndexRequestDataSourceInfo) *CreateIndexRequest {
	s.DataSourceInfo = v
	return s
}

func (s *CreateIndexRequest) SetDomain(v string) *CreateIndexRequest {
	s.Domain = &v
	return s
}

func (s *CreateIndexRequest) SetExtend(v map[string]interface{}) *CreateIndexRequest {
	s.Extend = v
	return s
}

func (s *CreateIndexRequest) SetName(v string) *CreateIndexRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexRequest) SetPartition(v int32) *CreateIndexRequest {
	s.Partition = &v
	return s
}

func (s *CreateIndexRequest) SetDryRun(v bool) *CreateIndexRequest {
	s.DryRun = &v
	return s
}

type CreateIndexRequestDataSourceInfo struct {
	AutoBuildIndex        *bool                                   `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config                *CreateIndexRequestDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	ProcessPartitionCount *int32                                  `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	Type                  *string                                 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateIndexRequestDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfo) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfo) SetAutoBuildIndex(v bool) *CreateIndexRequestDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetConfig(v *CreateIndexRequestDataSourceInfoConfig) *CreateIndexRequestDataSourceInfo {
	s.Config = v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetProcessPartitionCount(v int32) *CreateIndexRequestDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetType(v string) *CreateIndexRequestDataSourceInfo {
	s.Type = &v
	return s
}

type CreateIndexRequestDataSourceInfoConfig struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Partition    *string `json:"partition,omitempty" xml:"partition,omitempty"`
	Project      *string `json:"project,omitempty" xml:"project,omitempty"`
	Table        *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateIndexRequestDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetAccessKey(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetAccessSecret(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetEndpoint(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetPartition(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetProject(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetTable(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Table = &v
	return s
}

type CreateIndexResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetResult(v map[string]interface{}) *CreateIndexResponseBody {
	s.Result = v
	return s
}

type CreateIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndexResponse) SetBody(v *CreateIndexResponseBody) *CreateIndexResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// The billing method of the instance. Valid values: PREPAY and POSTPAY. PREPAY: subscription. If you set this parameter to PREPAY, make sure that your Alibaba Cloud account supports balance payment or credit payment. Otherwise, the system returns the InvalidPayMethod error message. In addition, you must specify the paymentInfo parameter. POSTPAY: pay-as-you-go. This billing method is not supported.
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The specifications of the instance.
	Components []*CreateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The information about billing.
	Order *CreateInstanceRequestOrder `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest {
	s.Components = v
	return s
}

func (s *CreateInstanceRequest) SetOrder(v *CreateInstanceRequestOrder) *CreateInstanceRequest {
	s.Order = v
	return s
}

type CreateInstanceRequestComponents struct {
	// The name of the specification. The value must be the same as the name of a parameter on the buy page.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The value of the specification.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestComponents) SetCode(v string) *CreateInstanceRequestComponents {
	s.Code = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetValue(v string) *CreateInstanceRequestComponents {
	s.Value = &v
	return s
}

type CreateInstanceRequestOrder struct {
	// Specifies whether to enable auto-renewal. Valid values: true and false.
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The billing cycle. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, and 12.
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The unit of the billing cycle. Valid values: Month and Year.
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s CreateInstanceRequestOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestOrder) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestOrder) SetAutoRenew(v bool) *CreateInstanceRequestOrder {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequestOrder) SetDuration(v int64) *CreateInstanceRequestOrder {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequestOrder) SetPricingCycle(v string) *CreateInstanceRequestOrder {
	s.PricingCycle = &v
	return s
}

type CreateInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result *CreateInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody {
	s.Result = v
	return s
}

type CreateInstanceResponseBodyResult struct {
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyResult) SetInstanceId(v string) *CreateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DeleteAdvanceConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdvanceConfigResponseBody) SetRequestId(v string) *DeleteAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *DeleteAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type DeleteAdvanceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdvanceConfigResponse) SetHeaders(v map[string]*string) *DeleteAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdvanceConfigResponse) SetStatusCode(v int32) *DeleteAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdvanceConfigResponse) SetBody(v *DeleteAdvanceConfigResponseBody) *DeleteAdvanceConfigResponse {
	s.Body = v
	return s
}

type DeleteDataSourceResponseBody struct {
	// The ID of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetResult(v map[string]interface{}) *DeleteDataSourceResponseBody {
	s.Result = v
	return s
}

type DeleteDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type DeleteIndexRequest struct {
	DataSource       *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	DeleteDataSource *bool   `json:"deleteDataSource,omitempty" xml:"deleteDataSource,omitempty"`
}

func (s DeleteIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexRequest) SetDataSource(v string) *DeleteIndexRequest {
	s.DataSource = &v
	return s
}

func (s *DeleteIndexRequest) SetDeleteDataSource(v bool) *DeleteIndexRequest {
	s.DeleteDataSource = &v
	return s
}

type DeleteIndexResponseBody struct {
	// id of request
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponseBody) SetRequestId(v string) *DeleteIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexResponseBody) SetResult(v map[string]interface{}) *DeleteIndexResponseBody {
	s.Result = v
	return s
}

type DeleteIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponse) SetHeaders(v map[string]*string) *DeleteIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexResponse) SetStatusCode(v int32) *DeleteIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexResponse) SetBody(v *DeleteIndexResponseBody) *DeleteIndexResponse {
	s.Body = v
	return s
}

type DeleteIndexVersionResponseBody struct {
	// id of request
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexVersionResponseBody) SetRequestId(v string) *DeleteIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexVersionResponseBody) SetResult(v map[string]interface{}) *DeleteIndexVersionResponseBody {
	s.Result = v
	return s
}

type DeleteIndexVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexVersionResponse) SetHeaders(v map[string]*string) *DeleteIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexVersionResponse) SetStatusCode(v int32) *DeleteIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexVersionResponse) SetBody(v *DeleteIndexVersionResponseBody) *DeleteIndexVersionResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetResult(v map[string]interface{}) *DeleteInstanceResponseBody {
	s.Result = v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type ForceSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index information.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ForceSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ForceSwitchResponseBody) SetRequestId(v string) *ForceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceSwitchResponseBody) SetResult(v map[string]interface{}) *ForceSwitchResponseBody {
	s.Result = v
	return s
}

type ForceSwitchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ForceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ForceSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ForceSwitchResponse) GoString() string {
	return s.String()
}

func (s *ForceSwitchResponse) SetHeaders(v map[string]*string) *ForceSwitchResponse {
	s.Headers = v
	return s
}

func (s *ForceSwitchResponse) SetStatusCode(v int32) *ForceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceSwitchResponse) SetBody(v *ForceSwitchResponseBody) *ForceSwitchResponse {
	s.Body = v
	return s
}

type GetAdvanceConfigRequest struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAdvanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigRequest) SetType(v string) *GetAdvanceConfigRequest {
	s.Type = &v
	return s
}

type GetAdvanceConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result *GetAdvanceConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBody) SetRequestId(v string) *GetAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvanceConfigResponseBody) SetResult(v *GetAdvanceConfigResponseBodyResult) *GetAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type GetAdvanceConfigResponseBodyResult struct {
	// The content of the configuration that is returned.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The type of the configuration content. Valid values: FILE, GIT, HTTP, and ODPS.
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The description.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The information about files.
	Files []*GetAdvanceConfigResponseBodyResultFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAdvanceConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBodyResult) SetContent(v string) *GetAdvanceConfigResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetContentType(v string) *GetAdvanceConfigResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetDesc(v string) *GetAdvanceConfigResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetFiles(v []*GetAdvanceConfigResponseBodyResultFiles) *GetAdvanceConfigResponseBodyResult {
	s.Files = v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetName(v string) *GetAdvanceConfigResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetStatus(v string) *GetAdvanceConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetUpdateTime(v int64) *GetAdvanceConfigResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type GetAdvanceConfigResponseBodyResultFiles struct {
	// The name of the file path.
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether it is a directory.
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether it is a template.
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAdvanceConfigResponseBodyResultFiles) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponseBodyResultFiles) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetFullPathName(v string) *GetAdvanceConfigResponseBodyResultFiles {
	s.FullPathName = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetIsDir(v bool) *GetAdvanceConfigResponseBodyResultFiles {
	s.IsDir = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetIsTemplate(v bool) *GetAdvanceConfigResponseBodyResultFiles {
	s.IsTemplate = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetName(v string) *GetAdvanceConfigResponseBodyResultFiles {
	s.Name = &v
	return s
}

type GetAdvanceConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponse) SetHeaders(v map[string]*string) *GetAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAdvanceConfigResponse) SetStatusCode(v int32) *GetAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvanceConfigResponse) SetBody(v *GetAdvanceConfigResponseBody) *GetAdvanceConfigResponse {
	s.Body = v
	return s
}

type GetAdvanceConfigFileRequest struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetAdvanceConfigFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileRequest) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileRequest) SetFileName(v string) *GetAdvanceConfigFileRequest {
	s.FileName = &v
	return s
}

type GetAdvanceConfigFileResponseBody struct {
	// id of request
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAdvanceConfigFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdvanceConfigFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponseBody) SetRequestId(v string) *GetAdvanceConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvanceConfigFileResponseBody) SetResult(v *GetAdvanceConfigFileResponseBodyResult) *GetAdvanceConfigFileResponseBody {
	s.Result = v
	return s
}

type GetAdvanceConfigFileResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetAdvanceConfigFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponseBodyResult) SetContent(v string) *GetAdvanceConfigFileResponseBodyResult {
	s.Content = &v
	return s
}

type GetAdvanceConfigFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAdvanceConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAdvanceConfigFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileResponse) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponse) SetHeaders(v map[string]*string) *GetAdvanceConfigFileResponse {
	s.Headers = v
	return s
}

func (s *GetAdvanceConfigFileResponse) SetStatusCode(v int32) *GetAdvanceConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvanceConfigFileResponse) SetBody(v *GetAdvanceConfigFileResponseBody) *GetAdvanceConfigFileResponse {
	s.Body = v
	return s
}

type GetClusterResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of the cluster details.
	Result *GetClusterResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetResult(v *GetClusterResponseBodyResult) *GetClusterResponseBody {
	s.Result = v
	return s
}

type GetClusterResponseBodyResult struct {
	// The time when the cluster was updated.
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// The effective advanced configuration version.
	CurrentAdvanceConfigVersion *string `json:"currentAdvanceConfigVersion,omitempty" xml:"currentAdvanceConfigVersion,omitempty"`
	// The effective online configuration version.
	CurrentOnlineConfigVersion *string `json:"currentOnlineConfigVersion,omitempty" xml:"currentOnlineConfigVersion,omitempty"`
	// The specifications of the data node.
	DataNode *GetClusterResponseBodyResultDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The latest advanced configuration version.
	LatestAdvanceConfigVersion *string `json:"latestAdvanceConfigVersion,omitempty" xml:"latestAdvanceConfigVersion,omitempty"`
	// The latest online configuration version.
	LatestOnlineConfigVersion *string `json:"latestOnlineConfigVersion,omitempty" xml:"latestOnlineConfigVersion,omitempty"`
	// The name of the cluster.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The specifications of the query node.
	QueryNode *GetClusterResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
	// The creation status of the cluster. Valid values: NEW and PUBLISH. NEW indicates that the cluster is being created. PUBLISH indicates that the cluster is created.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetClusterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResult) SetConfigUpdateTime(v string) *GetClusterResponseBodyResult {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCurrentAdvanceConfigVersion(v string) *GetClusterResponseBodyResult {
	s.CurrentAdvanceConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCurrentOnlineConfigVersion(v string) *GetClusterResponseBodyResult {
	s.CurrentOnlineConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetDataNode(v *GetClusterResponseBodyResultDataNode) *GetClusterResponseBodyResult {
	s.DataNode = v
	return s
}

func (s *GetClusterResponseBodyResult) SetDescription(v string) *GetClusterResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetLatestAdvanceConfigVersion(v string) *GetClusterResponseBodyResult {
	s.LatestAdvanceConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetLatestOnlineConfigVersion(v string) *GetClusterResponseBodyResult {
	s.LatestOnlineConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetName(v string) *GetClusterResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetQueryNode(v *GetClusterResponseBodyResultQueryNode) *GetClusterResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *GetClusterResponseBodyResult) SetStatus(v string) *GetClusterResponseBodyResult {
	s.Status = &v
	return s
}

type GetClusterResponseBodyResultDataNode struct {
	// The name of the node.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of replicas.
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of partitions.
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetClusterResponseBodyResultDataNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyResultDataNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResultDataNode) SetName(v string) *GetClusterResponseBodyResultDataNode {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) SetNumber(v int32) *GetClusterResponseBodyResultDataNode {
	s.Number = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) SetPartition(v int32) *GetClusterResponseBodyResultDataNode {
	s.Partition = &v
	return s
}

type GetClusterResponseBodyResultQueryNode struct {
	// The name of the node.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of nodes.
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of replicas.
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetClusterResponseBodyResultQueryNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResultQueryNode) SetName(v string) *GetClusterResponseBodyResultQueryNode {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) SetNumber(v int32) *GetClusterResponseBodyResultQueryNode {
	s.Number = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) SetPartition(v int32) *GetClusterResponseBodyResultQueryNode {
	s.Partition = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetClusterRunTimeInfoResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The configuration progress. Unit: percentage.
	Result []*GetClusterRunTimeInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetClusterRunTimeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBody) SetRequestId(v string) *GetClusterRunTimeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBody) SetResult(v []*GetClusterRunTimeInfoResponseBodyResult) *GetClusterRunTimeInfoResponseBody {
	s.Result = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResult struct {
	// The name of the cluster
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// dataNodes
	DataNodes []*GetClusterRunTimeInfoResponseBodyResultDataNodes `json:"dataNodes,omitempty" xml:"dataNodes,omitempty" type:"Repeated"`
	// The specifications of the query node.
	QueryNode *GetClusterRunTimeInfoResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetClusterName(v string) *GetClusterRunTimeInfoResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetDataNodes(v []*GetClusterRunTimeInfoResponseBodyResultDataNodes) *GetClusterRunTimeInfoResponseBodyResult {
	s.DataNodes = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetQueryNode(v *GetClusterRunTimeInfoResponseBodyResultQueryNode) *GetClusterRunTimeInfoResponseBodyResult {
	s.QueryNode = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodes struct {
	ConfigStatusList []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList `json:"configStatusList,omitempty" xml:"configStatusList,omitempty" type:"Repeated"`
	DataStatusList   []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList   `json:"dataStatusList,omitempty" xml:"dataStatusList,omitempty" type:"Repeated"`
	ServiceStatus    *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus      `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodes) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodes) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetConfigStatusList(v []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.ConfigStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetDataStatusList(v []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.DataStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetServiceStatus(v *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.ServiceStatus = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList struct {
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	DonePercent      *int32  `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	DoneSize         *int32  `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	TotalSize        *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetConfigUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList struct {
	AdvanceConfigInfo  *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo `json:"advanceConfigInfo,omitempty" xml:"advanceConfigInfo,omitempty" type:"Struct"`
	DeployFailedWorker []*string                                                                        `json:"deployFailedWorker,omitempty" xml:"deployFailedWorker,omitempty" type:"Repeated"`
	DocSize            *int32                                                                           `json:"docSize,omitempty" xml:"docSize,omitempty"`
	DonePercent        *int32                                                                           `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	DoneSize           *int32                                                                           `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	ErrorMsg           *string                                                                          `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	FullUpdateTime     *string                                                                          `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	FullVersion        *int64                                                                           `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	IncUpdateTime      *string                                                                          `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	IncVersion         *int64                                                                           `json:"incVersion,omitempty" xml:"incVersion,omitempty"`
	IndexConfigInfo    *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo   `json:"indexConfigInfo,omitempty" xml:"indexConfigInfo,omitempty" type:"Struct"`
	IndexSize          *int64                                                                           `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	LackDiskWorker     []*string                                                                        `json:"lackDiskWorker,omitempty" xml:"lackDiskWorker,omitempty" type:"Repeated"`
	LackMemWorker      []*string                                                                        `json:"lackMemWorker,omitempty" xml:"lackMemWorker,omitempty" type:"Repeated"`
	Name               *string                                                                          `json:"name,omitempty" xml:"name,omitempty"`
	TotalSize          *int32                                                                           `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetAdvanceConfigInfo(v *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.AdvanceConfigInfo = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDeployFailedWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DeployFailedWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDocSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DocSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetErrorMsg(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.ErrorMsg = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetFullUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.FullUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetFullVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.FullVersion = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIncUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IncUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIncVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IncVersion = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIndexConfigInfo(v *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IndexConfigInfo = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIndexSize(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IndexSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetLackDiskWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.LackDiskWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetLackMemWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.LackMemWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo struct {
	ConfigMetaName *string `json:"configMetaName,omitempty" xml:"configMetaName,omitempty"`
	Version        *int64  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) SetConfigMetaName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	s.ConfigMetaName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) SetVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	s.Version = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo struct {
	ConfigMetaName *string `json:"configMetaName,omitempty" xml:"configMetaName,omitempty"`
	Version        *int64  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) SetConfigMetaName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	s.ConfigMetaName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) SetVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	s.Version = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus struct {
	DonePercent *int32  `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	DoneSize    *int32  `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	TotalSize   *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultQueryNode struct {
	// configStatusList
	ConfigStatusList []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList `json:"configStatusList,omitempty" xml:"configStatusList,omitempty" type:"Repeated"`
	// serviceStatus
	ServiceStatus *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) SetConfigStatusList(v []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	s.ConfigStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) SetServiceStatus(v *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	s.ServiceStatus = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList struct {
	// configUpdateTime
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// donePercent
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// doneSize
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// totalSize
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetConfigUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus struct {
	// donePercent
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// doneSize
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The name of the cluster.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// totalSize
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClusterRunTimeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterRunTimeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponse) SetHeaders(v map[string]*string) *GetClusterRunTimeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetClusterRunTimeInfoResponse) SetStatusCode(v int32) *GetClusterRunTimeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterRunTimeInfoResponse) SetBody(v *GetClusterRunTimeInfoResponseBody) *GetClusterRunTimeInfoResponse {
	s.Body = v
	return s
}

type GetDataSourceResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the data source.
	Result *GetDataSourceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetResult(v *GetDataSourceResponseBodyResult) *GetDataSourceResponseBody {
	s.Result = v
	return s
}

type GetDataSourceResponseBodyResult struct {
	Domain      *string   `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes     []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	LastFulTime *int64    `json:"lastFulTime,omitempty" xml:"lastFulTime,omitempty"`
	Name        *string   `json:"name,omitempty" xml:"name,omitempty"`
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the data source.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDataSourceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBodyResult) SetDomain(v string) *GetDataSourceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetIndexes(v []*string) *GetDataSourceResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetLastFulTime(v int64) *GetDataSourceResponseBodyResult {
	s.LastFulTime = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetName(v string) *GetDataSourceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetStatus(v string) *GetDataSourceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetType(v string) *GetDataSourceResponseBodyResult {
	s.Type = &v
	return s
}

type GetDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponse) SetHeaders(v map[string]*string) *GetDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceResponse) SetStatusCode(v int32) *GetDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceResponse) SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse {
	s.Body = v
	return s
}

type GetDataSourceDeployResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetDataSourceDeployResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBody) SetRequestId(v string) *GetDataSourceDeployResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceDeployResponseBody) SetResult(v *GetDataSourceDeployResponseBodyResult) *GetDataSourceDeployResponseBody {
	s.Result = v
	return s
}

type GetDataSourceDeployResponseBodyResult struct {
	AutoBuildIndex *bool                                           `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Extend         *GetDataSourceDeployResponseBodyResultExtend    `json:"extend,omitempty" xml:"extend,omitempty" type:"Struct"`
	Processor      *GetDataSourceDeployResponseBodyResultProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	Storage        *GetDataSourceDeployResponseBodyResultStorage   `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	Swift          *GetDataSourceDeployResponseBodyResultSwift     `json:"swift,omitempty" xml:"swift,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResult) SetAutoBuildIndex(v bool) *GetDataSourceDeployResponseBodyResult {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetExtend(v *GetDataSourceDeployResponseBodyResultExtend) *GetDataSourceDeployResponseBodyResult {
	s.Extend = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetProcessor(v *GetDataSourceDeployResponseBodyResultProcessor) *GetDataSourceDeployResponseBodyResult {
	s.Processor = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetStorage(v *GetDataSourceDeployResponseBodyResultStorage) *GetDataSourceDeployResponseBodyResult {
	s.Storage = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetSwift(v *GetDataSourceDeployResponseBodyResultSwift) *GetDataSourceDeployResponseBodyResult {
	s.Swift = v
	return s
}

type GetDataSourceDeployResponseBodyResultExtend struct {
	Hdfs *GetDataSourceDeployResponseBodyResultExtendHdfs `json:"hdfs,omitempty" xml:"hdfs,omitempty" type:"Struct"`
	Odps *GetDataSourceDeployResponseBodyResultExtendOdps `json:"odps,omitempty" xml:"odps,omitempty" type:"Struct"`
	Oss  *GetDataSourceDeployResponseBodyResultExtendOss  `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
	Saro *GetDataSourceDeployResponseBodyResultExtendSaro `json:"saro,omitempty" xml:"saro,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBodyResultExtend) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtend) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetHdfs(v *GetDataSourceDeployResponseBodyResultExtendHdfs) *GetDataSourceDeployResponseBodyResultExtend {
	s.Hdfs = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetOdps(v *GetDataSourceDeployResponseBodyResultExtendOdps) *GetDataSourceDeployResponseBodyResultExtend {
	s.Odps = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetOss(v *GetDataSourceDeployResponseBodyResultExtendOss) *GetDataSourceDeployResponseBodyResultExtend {
	s.Oss = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetSaro(v *GetDataSourceDeployResponseBodyResultExtendSaro) *GetDataSourceDeployResponseBodyResultExtend {
	s.Saro = v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendHdfs struct {
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendHdfs) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendHdfs) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendHdfs) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendHdfs {
	s.Path = &v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendOdps struct {
	Partitions map[string]*string `json:"partitions,omitempty" xml:"partitions,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendOdps) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendOdps) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendOdps) SetPartitions(v map[string]*string) *GetDataSourceDeployResponseBodyResultExtendOdps {
	s.Partitions = v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendOss struct {
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendOss) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendOss) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendOss) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendOss {
	s.Path = &v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendSaro struct {
	Path    *string `json:"path,omitempty" xml:"path,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendSaro) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendSaro) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendSaro {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) SetVersion(v string) *GetDataSourceDeployResponseBodyResultExtendSaro {
	s.Version = &v
	return s
}

type GetDataSourceDeployResponseBodyResultProcessor struct {
	Args     *string `json:"args,omitempty" xml:"args,omitempty"`
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultProcessor) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultProcessor) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) SetArgs(v string) *GetDataSourceDeployResponseBodyResultProcessor {
	s.Args = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) SetResource(v string) *GetDataSourceDeployResponseBodyResultProcessor {
	s.Resource = &v
	return s
}

type GetDataSourceDeployResponseBodyResultStorage struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Bucket       *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Namespace    *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OssPath      *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition    *string `json:"partition,omitempty" xml:"partition,omitempty"`
	Path         *string `json:"path,omitempty" xml:"path,omitempty"`
	Project      *string `json:"project,omitempty" xml:"project,omitempty"`
	Table        *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultStorage) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultStorage) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetAccessKey(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.AccessKey = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetAccessSecret(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.AccessSecret = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetBucket(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Bucket = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetEndpoint(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Endpoint = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetNamespace(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Namespace = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetOssPath(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.OssPath = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetPartition(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Partition = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetPath(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetProject(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Project = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetTable(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Table = &v
	return s
}

type GetDataSourceDeployResponseBodyResultSwift struct {
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	Zk    *string `json:"zk,omitempty" xml:"zk,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultSwift) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultSwift) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultSwift) SetTopic(v string) *GetDataSourceDeployResponseBodyResultSwift {
	s.Topic = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultSwift) SetZk(v string) *GetDataSourceDeployResponseBodyResultSwift {
	s.Zk = &v
	return s
}

type GetDataSourceDeployResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDataSourceDeployResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataSourceDeployResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponse) SetHeaders(v map[string]*string) *GetDataSourceDeployResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceDeployResponse) SetStatusCode(v int32) *GetDataSourceDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceDeployResponse) SetBody(v *GetDataSourceDeployResponseBody) *GetDataSourceDeployResponse {
	s.Body = v
	return s
}

type GetDeployGraphResponseBody struct {
	// Id of the request
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetDeployGraphResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDeployGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBody) SetRequestId(v string) *GetDeployGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeployGraphResponseBody) SetResult(v *GetDeployGraphResponseBodyResult) *GetDeployGraphResponseBody {
	s.Result = v
	return s
}

type GetDeployGraphResponseBodyResult struct {
	Graph *GetDeployGraphResponseBodyResultGraph `json:"graph,omitempty" xml:"graph,omitempty" type:"Struct"`
}

func (s GetDeployGraphResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResult) SetGraph(v *GetDeployGraphResponseBodyResultGraph) *GetDeployGraphResponseBodyResult {
	s.Graph = v
	return s
}

type GetDeployGraphResponseBodyResultGraph struct {
	IndexMetas         []*GetDeployGraphResponseBodyResultGraphIndexMetas   `json:"indexMetas,omitempty" xml:"indexMetas,omitempty" type:"Repeated"`
	OnlineMaster       []*GetDeployGraphResponseBodyResultGraphOnlineMaster `json:"onlineMaster,omitempty" xml:"onlineMaster,omitempty" type:"Repeated"`
	TableIndexRelation map[string][]*string                                 `json:"tableIndexRelation,omitempty" xml:"tableIndexRelation,omitempty"`
	TableMetas         []*GetDeployGraphResponseBodyResultGraphTableMetas   `json:"tableMetas,omitempty" xml:"tableMetas,omitempty" type:"Repeated"`
	ZoneIndexRelation  map[string][]*string                                 `json:"zoneIndexRelation,omitempty" xml:"zoneIndexRelation,omitempty"`
	ZoneMetas          []*GetDeployGraphResponseBodyResultGraphZoneMetas    `json:"zoneMetas,omitempty" xml:"zoneMetas,omitempty" type:"Repeated"`
}

func (s GetDeployGraphResponseBodyResultGraph) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraph) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraph) SetIndexMetas(v []*GetDeployGraphResponseBodyResultGraphIndexMetas) *GetDeployGraphResponseBodyResultGraph {
	s.IndexMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetOnlineMaster(v []*GetDeployGraphResponseBodyResultGraphOnlineMaster) *GetDeployGraphResponseBodyResultGraph {
	s.OnlineMaster = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetTableIndexRelation(v map[string][]*string) *GetDeployGraphResponseBodyResultGraph {
	s.TableIndexRelation = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetTableMetas(v []*GetDeployGraphResponseBodyResultGraphTableMetas) *GetDeployGraphResponseBodyResultGraph {
	s.TableMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetZoneIndexRelation(v map[string][]*string) *GetDeployGraphResponseBodyResultGraph {
	s.ZoneIndexRelation = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetZoneMetas(v []*GetDeployGraphResponseBodyResultGraphZoneMetas) *GetDeployGraphResponseBodyResultGraph {
	s.ZoneMetas = v
	return s
}

type GetDeployGraphResponseBodyResultGraphIndexMetas struct {
	DomainName    *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	TableDeployId *int64  `json:"tableDeployId,omitempty" xml:"tableDeployId,omitempty"`
	TableName     *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Tag           *string `json:"tag,omitempty" xml:"tag,omitempty"`
	ZoneName      *string `json:"zoneName,omitempty" xml:"zoneName,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphIndexMetas) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphIndexMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTableDeployId(v int64) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.TableDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTableName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.TableName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetZoneName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.ZoneName = &v
	return s
}

type GetDeployGraphResponseBodyResultGraphOnlineMaster struct {
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	HippoId    *string `json:"hippoId,omitempty" xml:"hippoId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphOnlineMaster) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphOnlineMaster) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetHippoId(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.HippoId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetId(v int64) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.Id = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetName(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.Name = &v
	return s
}

type GetDeployGraphResponseBodyResultGraphTableMetas struct {
	BuildDeployId *int64  `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	DomainName    *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	TableDeployId *int64  `json:"tableDeployId,omitempty" xml:"tableDeployId,omitempty"`
	Tag           *string `json:"tag,omitempty" xml:"tag,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphTableMetas) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphTableMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetBuildDeployId(v int64) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.BuildDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetTableDeployId(v int64) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.TableDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetType(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Type = &v
	return s
}

type GetDeployGraphResponseBodyResultGraphZoneMetas struct {
	DomainInfo    *string `json:"domainInfo,omitempty" xml:"domainInfo,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	SuezAdminName *string `json:"suezAdminName,omitempty" xml:"suezAdminName,omitempty"`
	Tag           *string `json:"tag,omitempty" xml:"tag,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphZoneMetas) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphZoneMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetDomainInfo(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.DomainInfo = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetSuezAdminName(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.SuezAdminName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetType(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Type = &v
	return s
}

type GetDeployGraphResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeployGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeployGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponse) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponse) SetHeaders(v map[string]*string) *GetDeployGraphResponse {
	s.Headers = v
	return s
}

func (s *GetDeployGraphResponse) SetStatusCode(v int32) *GetDeployGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeployGraphResponse) SetBody(v *GetDeployGraphResponseBody) *GetDeployGraphResponse {
	s.Body = v
	return s
}

type GetFileRequest struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) SetFileName(v string) *GetFileRequest {
	s.FileName = &v
	return s
}

type GetFileResponseBody struct {
	// id of request
	RequestId *string                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileResponseBody) SetRequestId(v string) *GetFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileResponseBody) SetResult(v *GetFileResponseBodyResult) *GetFileResponseBody {
	s.Result = v
	return s
}

type GetFileResponseBodyResult struct {
	Content      *string `json:"content,omitempty" xml:"content,omitempty"`
	DataSource   *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	IsDir        *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Partition    *int64  `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyResult) SetContent(v string) *GetFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFileResponseBodyResult) SetDataSource(v string) *GetFileResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *GetFileResponseBodyResult) SetFullPathName(v string) *GetFileResponseBodyResult {
	s.FullPathName = &v
	return s
}

func (s *GetFileResponseBodyResult) SetIsDir(v bool) *GetFileResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *GetFileResponseBodyResult) SetName(v string) *GetFileResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetFileResponseBodyResult) SetPartition(v int64) *GetFileResponseBodyResult {
	s.Partition = &v
	return s
}

type GetFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponse) GoString() string {
	return s.String()
}

func (s *GetFileResponse) SetHeaders(v map[string]*string) *GetFileResponse {
	s.Headers = v
	return s
}

func (s *GetFileResponse) SetStatusCode(v int32) *GetFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileResponse) SetBody(v *GetFileResponseBody) *GetFileResponse {
	s.Body = v
	return s
}

type GetIndexResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index.
	Result *GetIndexResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBody) SetRequestId(v string) *GetIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexResponseBody) SetResult(v *GetIndexResponseBodyResult) *GetIndexResponseBody {
	s.Result = v
	return s
}

type GetIndexResponseBodyResult struct {
	// The content of the index.
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source.
	DataSourceInfo *GetIndexResponseBodyResultDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The remarks.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The last time when full data in the index was updated.
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The version of the data.
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The last time when incremental data in the index was updated.
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The index size.
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The status of the index. Valid values: NEW, PUBLISH, IN_USE, NOT_USE, STOP_USE, and RESTORE_USE. After a Retrieval Engine Edition instance is created, it enters the IN_USE state.
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of shards.
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The information about the versions.
	Versions []*GetIndexResponseBodyResultVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetIndexResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResult) SetContent(v string) *GetIndexResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDataSource(v string) *GetIndexResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDataSourceInfo(v *GetIndexResponseBodyResultDataSourceInfo) *GetIndexResponseBodyResult {
	s.DataSourceInfo = v
	return s
}

func (s *GetIndexResponseBodyResult) SetDescription(v string) *GetIndexResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDomain(v string) *GetIndexResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetFullUpdateTime(v string) *GetIndexResponseBodyResult {
	s.FullUpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetFullVersion(v int64) *GetIndexResponseBodyResult {
	s.FullVersion = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIncUpdateTime(v string) *GetIndexResponseBodyResult {
	s.IncUpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIndexSize(v int64) *GetIndexResponseBodyResult {
	s.IndexSize = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIndexStatus(v string) *GetIndexResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetName(v string) *GetIndexResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetPartition(v int32) *GetIndexResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetVersions(v []*GetIndexResponseBodyResultVersions) *GetIndexResponseBodyResult {
	s.Versions = v
	return s
}

type GetIndexResponseBodyResultDataSourceInfo struct {
	// Indicates whether the automatic full indexing feature is enabled.
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configuration of MaxCompute data sources.
	Config *GetIndexResponseBodyResultDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The offline deployment name of the data source.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The name of the data source.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of resources used for data update.
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// The configuration of SARO data sources.
	SaroConfig *GetIndexResponseBodyResultDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfo) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetAutoBuildIndex(v bool) *GetIndexResponseBodyResultDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetConfig(v *GetIndexResponseBodyResultDataSourceInfoConfig) *GetIndexResponseBodyResultDataSourceInfo {
	s.Config = v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetDomain(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetName(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetProcessPartitionCount(v int32) *GetIndexResponseBodyResultDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetSaroConfig(v *GetIndexResponseBodyResultDataSourceInfoSaroConfig) *GetIndexResponseBodyResultDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetType(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Type = &v
	return s
}

type GetIndexResponseBodyResultDataSourceInfoConfig struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Bucket       *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// A parameter related to MaxCompute.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// A parameter related to SARO.
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// A parameter related to OSS.
	OssPath   *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// A parameter related to Apsara File Storage for HDFS.
	Path    *string `json:"path,omitempty" xml:"path,omitempty"`
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// A parameter related to SARO and MaxCompute.
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetAccessKey(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetAccessSecret(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetBucket(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetEndpoint(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetNamespace(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetOssPath(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetPartition(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetPath(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetProject(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetTable(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Table = &v
	return s
}

type GetIndexResponseBodyResultDataSourceInfoSaroConfig struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfoSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) SetNamespace(v string) *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) SetTableName(v string) *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

type GetIndexResponseBodyResultVersions struct {
	// The description of the version.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The information about the files.
	Files []*GetIndexResponseBodyResultVersionsFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the version.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the version.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last time when the version was updated.
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the version.
	VersionId *int32 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetIndexResponseBodyResultVersions) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultVersions) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultVersions) SetDesc(v string) *GetIndexResponseBodyResultVersions {
	s.Desc = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetFiles(v []*GetIndexResponseBodyResultVersionsFiles) *GetIndexResponseBodyResultVersions {
	s.Files = v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetName(v string) *GetIndexResponseBodyResultVersions {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetStatus(v string) *GetIndexResponseBodyResultVersions {
	s.Status = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetUpdateTime(v int64) *GetIndexResponseBodyResultVersions {
	s.UpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetVersionId(v int32) *GetIndexResponseBodyResultVersions {
	s.VersionId = &v
	return s
}

type GetIndexResponseBodyResultVersionsFiles struct {
	// The full path of the file.
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory.
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a template.
	IsTemplate *bool   `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetIndexResponseBodyResultVersionsFiles) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultVersionsFiles) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetFullPathName(v string) *GetIndexResponseBodyResultVersionsFiles {
	s.FullPathName = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetIsDir(v bool) *GetIndexResponseBodyResultVersionsFiles {
	s.IsDir = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetIsTemplate(v bool) *GetIndexResponseBodyResultVersionsFiles {
	s.IsTemplate = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetName(v string) *GetIndexResponseBodyResultVersionsFiles {
	s.Name = &v
	return s
}

type GetIndexResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponse) GoString() string {
	return s.String()
}

func (s *GetIndexResponse) SetHeaders(v map[string]*string) *GetIndexResponse {
	s.Headers = v
	return s
}

func (s *GetIndexResponse) SetStatusCode(v int32) *GetIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexResponse) SetBody(v *GetIndexResponseBody) *GetIndexResponse {
	s.Body = v
	return s
}

type GetIndexVersionResponseBody struct {
	// id of request
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetIndexVersionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBody) SetRequestId(v string) *GetIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexVersionResponseBody) SetResult(v *GetIndexVersionResponseBodyResult) *GetIndexVersionResponseBody {
	s.Result = v
	return s
}

type GetIndexVersionResponseBodyResult struct {
	Cluster       *string                                           `json:"cluster,omitempty" xml:"cluster,omitempty"`
	IndexVersions []*GetIndexVersionResponseBodyResultIndexVersions `json:"indexVersions,omitempty" xml:"indexVersions,omitempty" type:"Repeated"`
}

func (s GetIndexVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBodyResult) SetCluster(v string) *GetIndexVersionResponseBodyResult {
	s.Cluster = &v
	return s
}

func (s *GetIndexVersionResponseBodyResult) SetIndexVersions(v []*GetIndexVersionResponseBodyResultIndexVersions) *GetIndexVersionResponseBodyResult {
	s.IndexVersions = v
	return s
}

type GetIndexVersionResponseBodyResultIndexVersions struct {
	BuildDeployId  *string  `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	CurrentVersion *int64   `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	IndexName      *string  `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Versions       []*int64 `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetIndexVersionResponseBodyResultIndexVersions) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponseBodyResultIndexVersions) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetBuildDeployId(v string) *GetIndexVersionResponseBodyResultIndexVersions {
	s.BuildDeployId = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetCurrentVersion(v int64) *GetIndexVersionResponseBodyResultIndexVersions {
	s.CurrentVersion = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetIndexName(v string) *GetIndexVersionResponseBodyResultIndexVersions {
	s.IndexName = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetVersions(v []*int64) *GetIndexVersionResponseBodyResultIndexVersions {
	s.Versions = v
	return s
}

type GetIndexVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponse) SetHeaders(v map[string]*string) *GetIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *GetIndexVersionResponse) SetStatusCode(v int32) *GetIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexVersionResponse) SetBody(v *GetIndexVersionResponseBody) *GetIndexVersionResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The description of the instance.
	Result *GetInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResult(v *GetInstanceResponseBodyResult) *GetInstanceResponseBody {
	s.Result = v
	return s
}

type GetInstanceResponseBodyResult struct {
	// 付费类型
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 商品code
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the request.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// WB01240825
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// 是否欠费
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// 代表资源一级ID的资源属性字段
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 锁定状态
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// ### Sample responses
	//
	// **Sample success responses**
	//
	//     {
	//       "requestId": "90D6B8F5-FE97-4509-9AAB-367836C51818",
	//       "result":
	//       {
	//         "instanceId":"fadsfsafs",
	//         "inDebt":true,
	//         "lockMode":"Unlock",
	//         "expiredTime":"asdfas",
	//         "updateTime":"dfasf",
	//         "createTime":"dfasf",
	//         "resourceGroupId":"resourceGroupID",
	//         "commodityCode":"commodityCode",
	//         "chargeType":"POSYPAY",
	//         "description":"this is description",
	//         "apiVersion": "tisplus/v1",
	//         "network": {
	//           "vSwitchId": "vswitch_id_xxx",
	//           "vpcId": "vpc_id_xxx",
	//         },
	//         "userName": "user",
	//         "spec": {
	//           "searchResource": {
	//             "disk": 50,
	//             "mem": 8,
	//             "cpu": 2,
	//             "nodeCount": 2
	//           },
	//           "qrsResource": {
	//             "disk": 50,
	//             "mem": 8,
	//             "cpu": 2,
	//             "nodeCount": 2
	//           }
	//         },
	//        "status": "INIT",
	//       }
	//     }
	//
	// **Sample error responses**
	//
	//     {
	//       "requestId": "BD1EA715-DF6F-06C2-004C-C1FA0D3A9820",
	//       "httpCode": 404,
	//       "code": "App.NotFound",
	//       "message": "App not found"
	//     }
	ResourceGroupId *string                              `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Status          *string                              `json:"status,omitempty" xml:"status,omitempty"`
	Tags            []*GetInstanceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResult) SetChargeType(v string) *GetInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetCommodityCode(v string) *GetInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetCreateTime(v string) *GetInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetDescription(v string) *GetInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetExpiredTime(v string) *GetInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetInDebt(v bool) *GetInstanceResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetInstanceId(v string) *GetInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetLockMode(v string) *GetInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetResourceGroupId(v string) *GetInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetStatus(v string) *GetInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetTags(v []*GetInstanceResponseBodyResultTags) *GetInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyResult) SetUpdateTime(v string) *GetInstanceResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type GetInstanceResponseBodyResultTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetInstanceResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultTags) SetKey(v string) *GetInstanceResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyResultTags) SetValue(v string) *GetInstanceResponseBodyResultTags {
	s.Value = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetNodeConfigRequest struct {
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetNodeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNodeConfigRequest) SetClusterName(v string) *GetNodeConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *GetNodeConfigRequest) SetName(v string) *GetNodeConfigRequest {
	s.Name = &v
	return s
}

func (s *GetNodeConfigRequest) SetType(v string) *GetNodeConfigRequest {
	s.Type = &v
	return s
}

type GetNodeConfigResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetNodeConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetNodeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponseBody) SetRequestId(v string) *GetNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeConfigResponseBody) SetResult(v *GetNodeConfigResponseBodyResult) *GetNodeConfigResponseBody {
	s.Result = v
	return s
}

type GetNodeConfigResponseBodyResult struct {
	Active              *bool  `json:"active,omitempty" xml:"active,omitempty"`
	DataDuplicateNumber *int32 `json:"dataDuplicateNumber,omitempty" xml:"dataDuplicateNumber,omitempty"`
	DataFragmentNumber  *int32 `json:"dataFragmentNumber,omitempty" xml:"dataFragmentNumber,omitempty"`
	MinServicePercent   *int32 `json:"minServicePercent,omitempty" xml:"minServicePercent,omitempty"`
	Published           *bool  `json:"published,omitempty" xml:"published,omitempty"`
}

func (s GetNodeConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponseBodyResult) SetActive(v bool) *GetNodeConfigResponseBodyResult {
	s.Active = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetDataDuplicateNumber(v int32) *GetNodeConfigResponseBodyResult {
	s.DataDuplicateNumber = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetDataFragmentNumber(v int32) *GetNodeConfigResponseBodyResult {
	s.DataFragmentNumber = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetMinServicePercent(v int32) *GetNodeConfigResponseBodyResult {
	s.MinServicePercent = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetPublished(v bool) *GetNodeConfigResponseBodyResult {
	s.Published = &v
	return s
}

type GetNodeConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponse) SetHeaders(v map[string]*string) *GetNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNodeConfigResponse) SetStatusCode(v int32) *GetNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeConfigResponse) SetBody(v *GetNodeConfigResponseBody) *GetNodeConfigResponse {
	s.Body = v
	return s
}

type ListAdvanceConfigDirRequest struct {
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
}

func (s ListAdvanceConfigDirRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirRequest) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirRequest) SetDirName(v string) *ListAdvanceConfigDirRequest {
	s.DirName = &v
	return s
}

type ListAdvanceConfigDirResponseBody struct {
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListAdvanceConfigDirResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAdvanceConfigDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponseBody) SetRequestId(v string) *ListAdvanceConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBody) SetResult(v []*ListAdvanceConfigDirResponseBodyResult) *ListAdvanceConfigDirResponseBody {
	s.Result = v
	return s
}

type ListAdvanceConfigDirResponseBodyResult struct {
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	IsDir        *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	IsTemplate   *bool   `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAdvanceConfigDirResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetFullPathName(v string) *ListAdvanceConfigDirResponseBodyResult {
	s.FullPathName = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetIsDir(v bool) *ListAdvanceConfigDirResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetIsTemplate(v bool) *ListAdvanceConfigDirResponseBodyResult {
	s.IsTemplate = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetName(v string) *ListAdvanceConfigDirResponseBodyResult {
	s.Name = &v
	return s
}

type ListAdvanceConfigDirResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAdvanceConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAdvanceConfigDirResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirResponse) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponse) SetHeaders(v map[string]*string) *ListAdvanceConfigDirResponse {
	s.Headers = v
	return s
}

func (s *ListAdvanceConfigDirResponse) SetStatusCode(v int32) *ListAdvanceConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvanceConfigDirResponse) SetBody(v *ListAdvanceConfigDirResponseBody) *ListAdvanceConfigDirResponse {
	s.Body = v
	return s
}

type ListAdvanceConfigsRequest struct {
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	IndexName      *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAdvanceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsRequest) SetDataSourceName(v string) *ListAdvanceConfigsRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetIndexName(v string) *ListAdvanceConfigsRequest {
	s.IndexName = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetType(v string) *ListAdvanceConfigsRequest {
	s.Type = &v
	return s
}

type ListAdvanceConfigsResponseBody struct {
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListAdvanceConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAdvanceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBody) SetRequestId(v string) *ListAdvanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvanceConfigsResponseBody) SetResult(v []*ListAdvanceConfigsResponseBodyResult) *ListAdvanceConfigsResponseBody {
	s.Result = v
	return s
}

type ListAdvanceConfigsResponseBodyResult struct {
	// 配置内容 http，git 请求时不为空
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 配置内容的类型 (FILE, GIT, HTTP, ODPS)
	ContentType *string                                      `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Desc        *string                                      `json:"desc,omitempty" xml:"desc,omitempty"`
	Files       []*ListAdvanceConfigsResponseBodyResultFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	Name        *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	Status      *string                                      `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTime  *int64                                       `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAdvanceConfigsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBodyResult) SetContent(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetContentType(v string) *ListAdvanceConfigsResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetDesc(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetFiles(v []*ListAdvanceConfigsResponseBodyResultFiles) *ListAdvanceConfigsResponseBodyResult {
	s.Files = v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetName(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetStatus(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetUpdateTime(v int64) *ListAdvanceConfigsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type ListAdvanceConfigsResponseBodyResultFiles struct {
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	IsDir        *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	IsTemplate   *bool   `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAdvanceConfigsResponseBodyResultFiles) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponseBodyResultFiles) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetFullPathName(v string) *ListAdvanceConfigsResponseBodyResultFiles {
	s.FullPathName = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetIsDir(v bool) *ListAdvanceConfigsResponseBodyResultFiles {
	s.IsDir = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetIsTemplate(v bool) *ListAdvanceConfigsResponseBodyResultFiles {
	s.IsTemplate = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetName(v string) *ListAdvanceConfigsResponseBodyResultFiles {
	s.Name = &v
	return s
}

type ListAdvanceConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAdvanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAdvanceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponse) SetHeaders(v map[string]*string) *ListAdvanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAdvanceConfigsResponse) SetStatusCode(v int32) *ListAdvanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvanceConfigsResponse) SetBody(v *ListAdvanceConfigsResponseBody) *ListAdvanceConfigsResponse {
	s.Body = v
	return s
}

type ListClusterNamesResponseBody struct {
	// id of request
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ListClusterNamesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListClusterNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponseBody) SetRequestId(v string) *ListClusterNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterNamesResponseBody) SetResult(v *ListClusterNamesResponseBodyResult) *ListClusterNamesResponseBody {
	s.Result = v
	return s
}

type ListClusterNamesResponseBodyResult struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListClusterNamesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponseBodyResult) SetDescription(v string) *ListClusterNamesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) SetId(v int64) *ListClusterNamesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) SetName(v string) *ListClusterNamesResponseBodyResult {
	s.Name = &v
	return s
}

type ListClusterNamesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNamesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponse) SetHeaders(v map[string]*string) *ListClusterNamesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNamesResponse) SetStatusCode(v int32) *ListClusterNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNamesResponse) SetBody(v *ListClusterNamesResponseBody) *ListClusterNamesResponse {
	s.Body = v
	return s
}

type ListClusterTasksResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The date when the task was completed.
	Result []*ListClusterTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClusterTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBody) SetRequestId(v string) *ListClusterTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTasksResponseBody) SetResult(v []*ListClusterTasksResponseBodyResult) *ListClusterTasksResponseBody {
	s.Result = v
	return s
}

type ListClusterTasksResponseBodyResult struct {
	ExtraAttribute *string `json:"extraAttribute,omitempty" xml:"extraAttribute,omitempty"`
	Field3         *string `json:"field3,omitempty" xml:"field3,omitempty"`
	// fsmId
	FsmId *string `json:"fsmId,omitempty" xml:"fsmId,omitempty"`
	// ### Method
	//
	// ```java
	// GET
	// ```
	//
	// ### URI
	//
	// ```java
	// /openapi/ha3/instances/{instanceId}/cluster-tasks
	// ```
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// Displays cluster tasks .
	Name      *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	Status    *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	Tags      []*ListClusterTasksResponseBodyResultTags      `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TaskNodes []*ListClusterTasksResponseBodyResultTaskNodes `json:"taskNodes,omitempty" xml:"taskNodes,omitempty" type:"Repeated"`
	Time      *string                                        `json:"time,omitempty" xml:"time,omitempty"`
	Type      *string                                        `json:"type,omitempty" xml:"type,omitempty"`
	User      *string                                        `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ListClusterTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResult) SetExtraAttribute(v string) *ListClusterTasksResponseBodyResult {
	s.ExtraAttribute = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetField3(v string) *ListClusterTasksResponseBodyResult {
	s.Field3 = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetFsmId(v string) *ListClusterTasksResponseBodyResult {
	s.FsmId = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetGroupType(v string) *ListClusterTasksResponseBodyResult {
	s.GroupType = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetName(v string) *ListClusterTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetStatus(v string) *ListClusterTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTags(v []*ListClusterTasksResponseBodyResultTags) *ListClusterTasksResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTaskNodes(v []*ListClusterTasksResponseBodyResultTaskNodes) *ListClusterTasksResponseBodyResult {
	s.TaskNodes = v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTime(v string) *ListClusterTasksResponseBodyResult {
	s.Time = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetType(v string) *ListClusterTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetUser(v string) *ListClusterTasksResponseBodyResult {
	s.User = &v
	return s
}

type ListClusterTasksResponseBodyResultTags struct {
	Msg      *string `json:"msg,omitempty" xml:"msg,omitempty"`
	TagLevel *string `json:"tagLevel,omitempty" xml:"tagLevel,omitempty"`
}

func (s ListClusterTasksResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResultTags) SetMsg(v string) *ListClusterTasksResponseBodyResultTags {
	s.Msg = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTags) SetTagLevel(v string) *ListClusterTasksResponseBodyResultTags {
	s.TagLevel = &v
	return s
}

type ListClusterTasksResponseBodyResultTaskNodes struct {
	FinishDate *string `json:"finishDate,omitempty" xml:"finishDate,omitempty"`
	Index      *int64  `json:"index,omitempty" xml:"index,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListClusterTasksResponseBodyResultTaskNodes) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBodyResultTaskNodes) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetFinishDate(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.FinishDate = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetIndex(v int64) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Index = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetName(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Name = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetStatus(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Status = &v
	return s
}

type ListClusterTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponse) SetHeaders(v map[string]*string) *ListClusterTasksResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTasksResponse) SetStatusCode(v int32) *ListClusterTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterTasksResponse) SetBody(v *ListClusterTasksResponseBody) *ListClusterTasksResponse {
	s.Body = v
	return s
}

type ListClustersResponseBody struct {
	// id of request
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListClustersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetResult(v []*ListClustersResponseBodyResult) *ListClustersResponseBody {
	s.Result = v
	return s
}

type ListClustersResponseBodyResult struct {
	ConfigUpdateTime            *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	CurrentAdvanceConfigVersion *string `json:"currentAdvanceConfigVersion,omitempty" xml:"currentAdvanceConfigVersion,omitempty"`
	// 词典配置生效版本
	CurrentOfflineDictConfigVersion *string `json:"currentOfflineDictConfigVersion,omitempty" xml:"currentOfflineDictConfigVersion,omitempty"`
	CurrentOnlineConfigVersion      *string `json:"currentOnlineConfigVersion,omitempty" xml:"currentOnlineConfigVersion,omitempty"`
	// 查询配置生效版本
	CurrentOnlineQueryConfigVersion *string                                 `json:"currentOnlineQueryConfigVersion,omitempty" xml:"currentOnlineQueryConfigVersion,omitempty"`
	DataNode                        *ListClustersResponseBodyResultDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	Description                     *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	LatestAdvanceConfigVersion      *string                                 `json:"latestAdvanceConfigVersion,omitempty" xml:"latestAdvanceConfigVersion,omitempty"`
	// 词典配置最新版本
	LatestOfflineDictConfigVersion *string `json:"latestOfflineDictConfigVersion,omitempty" xml:"latestOfflineDictConfigVersion,omitempty"`
	LatestOnlineConfigVersion      *string `json:"latestOnlineConfigVersion,omitempty" xml:"latestOnlineConfigVersion,omitempty"`
	// 查询配置最新版本
	LatestOnlineQueryConfigVersion *string                                  `json:"latestOnlineQueryConfigVersion,omitempty" xml:"latestOnlineQueryConfigVersion,omitempty"`
	Name                           *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	QueryNode                      *ListClustersResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
	Status                         *string                                  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListClustersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResult) SetConfigUpdateTime(v string) *ListClustersResponseBodyResult {
	s.ConfigUpdateTime = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentAdvanceConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentAdvanceConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOfflineDictConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOfflineDictConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOnlineConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOnlineConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOnlineQueryConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOnlineQueryConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetDataNode(v *ListClustersResponseBodyResultDataNode) *ListClustersResponseBodyResult {
	s.DataNode = v
	return s
}

func (s *ListClustersResponseBodyResult) SetDescription(v string) *ListClustersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestAdvanceConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestAdvanceConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOfflineDictConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOfflineDictConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOnlineConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOnlineConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOnlineQueryConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOnlineQueryConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetName(v string) *ListClustersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetQueryNode(v *ListClustersResponseBodyResultQueryNode) *ListClustersResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *ListClustersResponseBodyResult) SetStatus(v string) *ListClustersResponseBodyResult {
	s.Status = &v
	return s
}

type ListClustersResponseBodyResultDataNode struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Number    *int32  `json:"number,omitempty" xml:"number,omitempty"`
	Partition *int32  `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ListClustersResponseBodyResultDataNode) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyResultDataNode) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResultDataNode) SetName(v string) *ListClustersResponseBodyResultDataNode {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) SetNumber(v int32) *ListClustersResponseBodyResultDataNode {
	s.Number = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) SetPartition(v int32) *ListClustersResponseBodyResultDataNode {
	s.Partition = &v
	return s
}

type ListClustersResponseBodyResultQueryNode struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Number    *int32  `json:"number,omitempty" xml:"number,omitempty"`
	Partition *int32  `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ListClustersResponseBodyResultQueryNode) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResultQueryNode) SetName(v string) *ListClustersResponseBodyResultQueryNode {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) SetNumber(v int32) *ListClustersResponseBodyResultQueryNode {
	s.Number = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) SetPartition(v int32) *ListClustersResponseBodyResultQueryNode {
	s.Partition = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListDataSourceSchemasResponseBody struct {
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListDataSourceSchemasResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBody) SetRequestId(v string) *ListDataSourceSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceSchemasResponseBody) SetResult(v []*ListDataSourceSchemasResponseBodyResult) *ListDataSourceSchemasResponseBody {
	s.Result = v
	return s
}

type ListDataSourceSchemasResponseBodyResult struct {
	AddIndex   *bool                                              `json:"addIndex,omitempty" xml:"addIndex,omitempty"`
	Attribute  *bool                                              `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Custom     *bool                                              `json:"custom,omitempty" xml:"custom,omitempty"`
	Name       *string                                            `json:"name,omitempty" xml:"name,omitempty"`
	PrimaryKey *ListDataSourceSchemasResponseBodyResultPrimaryKey `json:"primaryKey,omitempty" xml:"primaryKey,omitempty" type:"Struct"`
	Summary    *bool                                              `json:"summary,omitempty" xml:"summary,omitempty"`
	Type       *string                                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSourceSchemasResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBodyResult) SetAddIndex(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.AddIndex = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetAttribute(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Attribute = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetCustom(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Custom = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetName(v string) *ListDataSourceSchemasResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetPrimaryKey(v *ListDataSourceSchemasResponseBodyResultPrimaryKey) *ListDataSourceSchemasResponseBodyResult {
	s.PrimaryKey = v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetSummary(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetType(v string) *ListDataSourceSchemasResponseBodyResult {
	s.Type = &v
	return s
}

type ListDataSourceSchemasResponseBodyResultPrimaryKey struct {
	HasPrimaryKeyAttribute *bool `json:"hasPrimaryKeyAttribute,omitempty" xml:"hasPrimaryKeyAttribute,omitempty"`
	IsPrimaryKey           *bool `json:"isPrimaryKey,omitempty" xml:"isPrimaryKey,omitempty"`
	IsPrimaryKeySorted     *bool `json:"isPrimaryKeySorted,omitempty" xml:"isPrimaryKeySorted,omitempty"`
}

func (s ListDataSourceSchemasResponseBodyResultPrimaryKey) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponseBodyResultPrimaryKey) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetHasPrimaryKeyAttribute(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.HasPrimaryKeyAttribute = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetIsPrimaryKey(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.IsPrimaryKey = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetIsPrimaryKeySorted(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.IsPrimaryKeySorted = &v
	return s
}

type ListDataSourceSchemasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataSourceSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponse) SetHeaders(v map[string]*string) *ListDataSourceSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceSchemasResponse) SetStatusCode(v int32) *ListDataSourceSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceSchemasResponse) SetBody(v *ListDataSourceSchemasResponseBody) *ListDataSourceSchemasResponse {
	s.Body = v
	return s
}

type ListDataSourceTasksResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The date when the task was completed.
	Result []*ListDataSourceTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBody) SetRequestId(v string) *ListDataSourceTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTasksResponseBody) SetResult(v []*ListDataSourceTasksResponseBodyResult) *ListDataSourceTasksResponseBody {
	s.Result = v
	return s
}

type ListDataSourceTasksResponseBodyResult struct {
	ExtraAttribute *string `json:"extraAttribute,omitempty" xml:"extraAttribute,omitempty"`
	Field3         *string `json:"field3,omitempty" xml:"field3,omitempty"`
	// fsmId
	FsmId *string `json:"fsmId,omitempty" xml:"fsmId,omitempty"`
	// ### Method
	//
	// ```java
	// GET
	// ```
	//
	// ### URI
	//
	// ```java
	// /openapi/ha3/instances/{instanceId}/data-source-tasks
	// ```
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// Displays data source tasks.
	Name      *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	Status    *string                                           `json:"status,omitempty" xml:"status,omitempty"`
	Tags      []*ListDataSourceTasksResponseBodyResultTags      `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TaskNodes []*ListDataSourceTasksResponseBodyResultTaskNodes `json:"taskNodes,omitempty" xml:"taskNodes,omitempty" type:"Repeated"`
	Time      *string                                           `json:"time,omitempty" xml:"time,omitempty"`
	Type      *string                                           `json:"type,omitempty" xml:"type,omitempty"`
	User      *string                                           `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResult) SetExtraAttribute(v string) *ListDataSourceTasksResponseBodyResult {
	s.ExtraAttribute = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetField3(v string) *ListDataSourceTasksResponseBodyResult {
	s.Field3 = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetFsmId(v string) *ListDataSourceTasksResponseBodyResult {
	s.FsmId = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetGroupType(v string) *ListDataSourceTasksResponseBodyResult {
	s.GroupType = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetName(v string) *ListDataSourceTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetStatus(v string) *ListDataSourceTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTags(v []*ListDataSourceTasksResponseBodyResultTags) *ListDataSourceTasksResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTaskNodes(v []*ListDataSourceTasksResponseBodyResultTaskNodes) *ListDataSourceTasksResponseBodyResult {
	s.TaskNodes = v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTime(v string) *ListDataSourceTasksResponseBodyResult {
	s.Time = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetType(v string) *ListDataSourceTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetUser(v string) *ListDataSourceTasksResponseBodyResult {
	s.User = &v
	return s
}

type ListDataSourceTasksResponseBodyResultTags struct {
	Msg      *string `json:"msg,omitempty" xml:"msg,omitempty"`
	TagLevel *string `json:"tagLevel,omitempty" xml:"tagLevel,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResultTags) SetMsg(v string) *ListDataSourceTasksResponseBodyResultTags {
	s.Msg = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTags) SetTagLevel(v string) *ListDataSourceTasksResponseBodyResultTags {
	s.TagLevel = &v
	return s
}

type ListDataSourceTasksResponseBodyResultTaskNodes struct {
	FinishDate *string `json:"finishDate,omitempty" xml:"finishDate,omitempty"`
	Index      *int64  `json:"index,omitempty" xml:"index,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResultTaskNodes) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResultTaskNodes) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetFinishDate(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.FinishDate = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetIndex(v int64) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Index = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetName(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Name = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetStatus(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Status = &v
	return s
}

type ListDataSourceTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataSourceTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponse) SetHeaders(v map[string]*string) *ListDataSourceTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTasksResponse) SetStatusCode(v int32) *ListDataSourceTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTasksResponse) SetBody(v *ListDataSourceTasksResponseBody) *ListDataSourceTasksResponse {
	s.Body = v
	return s
}

type ListDataSourcesResponseBody struct {
	// ## Method
	//
	// `GET`
	//
	// ## URI
	//
	// `/openapi/ha3/instances/{instanceId}/data-sources`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*ListDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetResult(v []*ListDataSourcesResponseBodyResult) *ListDataSourcesResponseBody {
	s.Result = v
	return s
}

type ListDataSourcesResponseBodyResult struct {
	// The data sources deployed in offline mode.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The time when the full data of the data source was last queried.
	LastFulTime *int64 `json:"lastFulTime,omitempty" xml:"lastFulTime,omitempty"`
	// The name of the data source.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data source.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the data source.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSourcesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyResult) SetDomain(v string) *ListDataSourcesResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetIndexes(v []*string) *ListDataSourcesResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetLastFulTime(v int64) *ListDataSourcesResponseBodyResult {
	s.LastFulTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetName(v string) *ListDataSourcesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetStatus(v string) *ListDataSourcesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetType(v string) *ListDataSourcesResponseBodyResult {
	s.Type = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetStatusCode(v int32) *ListDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListDateSourceGenerationsRequest struct {
	// The data center where the data source is deployed.
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The valid state of the data source. Valid values: true and false. The default value of this parameter is true.
	//
	// 1.  true indicates that the generations that have not expired and of which the tasks have been executed are returned.
	// 2.  false indicates that all generations are returned.
	ValidStatus *bool `json:"validStatus,omitempty" xml:"validStatus,omitempty"`
}

func (s ListDateSourceGenerationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsRequest) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsRequest) SetDomainName(v string) *ListDateSourceGenerationsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDateSourceGenerationsRequest) SetValidStatus(v bool) *ListDateSourceGenerationsRequest {
	s.ValidStatus = &v
	return s
}

type ListDateSourceGenerationsResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListDateSourceGenerationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDateSourceGenerationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponseBody) SetRequestId(v string) *ListDateSourceGenerationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBody) SetResult(v []*ListDateSourceGenerationsResponseBodyResult) *ListDateSourceGenerationsResponseBody {
	s.Result = v
	return s
}

type ListDateSourceGenerationsResponseBodyResult struct {
	// buildDeployId
	BuildDeployId *int32 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The time to start index building.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The directory where the index file created by using the dump table is saved.
	DataDumpRoot *string `json:"dataDumpRoot,omitempty" xml:"dataDumpRoot,omitempty"`
	// The primary key of the generation.
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// Key indicates the name of the index. value indicates the number of shards.
	Partition map[string]*int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The status.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the offline indexing was initiated.
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListDateSourceGenerationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetBuildDeployId(v int32) *ListDateSourceGenerationsResponseBodyResult {
	s.BuildDeployId = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetCreateTime(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetDataDumpRoot(v string) *ListDateSourceGenerationsResponseBodyResult {
	s.DataDumpRoot = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetGeneration(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetPartition(v map[string]*int32) *ListDateSourceGenerationsResponseBodyResult {
	s.Partition = v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetStatus(v string) *ListDateSourceGenerationsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetTimestamp(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.Timestamp = &v
	return s
}

type ListDateSourceGenerationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDateSourceGenerationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDateSourceGenerationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsResponse) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponse) SetHeaders(v map[string]*string) *ListDateSourceGenerationsResponse {
	s.Headers = v
	return s
}

func (s *ListDateSourceGenerationsResponse) SetStatusCode(v int32) *ListDateSourceGenerationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDateSourceGenerationsResponse) SetBody(v *ListDateSourceGenerationsResponseBody) *ListDateSourceGenerationsResponse {
	s.Body = v
	return s
}

type ListIndexesRequest struct {
	NewMode *bool `json:"newMode,omitempty" xml:"newMode,omitempty"`
}

func (s ListIndexesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexesRequest) SetNewMode(v bool) *ListIndexesRequest {
	s.NewMode = &v
	return s
}

type ListIndexesResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the indexes.
	Result []*ListIndexesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBody) SetRequestId(v string) *ListIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexesResponseBody) SetResult(v []*ListIndexesResponseBodyResult) *ListIndexesResponseBody {
	s.Result = v
	return s
}

type ListIndexesResponseBodyResult struct {
	// The content of the index.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The data source.
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source.
	DataSourceInfo *ListIndexesResponseBodyResultDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The remarks.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The deployment name of the index.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The last time when full data in the index was updated.
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The version of the data.
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The last time when incremental data in the index was updated.
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The index size.
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The status of the index. Valid values: NEW and PUBLISH.
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	// The name of the index.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of shards.
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The information about the versions.
	Versions []*ListIndexesResponseBodyResultVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResult) SetContent(v string) *ListIndexesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDataSource(v string) *ListIndexesResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDataSourceInfo(v *ListIndexesResponseBodyResultDataSourceInfo) *ListIndexesResponseBodyResult {
	s.DataSourceInfo = v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDescription(v string) *ListIndexesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDomain(v string) *ListIndexesResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetFullUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.FullUpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetFullVersion(v int64) *ListIndexesResponseBodyResult {
	s.FullVersion = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIncUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.IncUpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIndexSize(v int64) *ListIndexesResponseBodyResult {
	s.IndexSize = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIndexStatus(v string) *ListIndexesResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetName(v string) *ListIndexesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetPartition(v int32) *ListIndexesResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetVersions(v []*ListIndexesResponseBodyResultVersions) *ListIndexesResponseBodyResult {
	s.Versions = v
	return s
}

type ListIndexesResponseBodyResultDataSourceInfo struct {
	// Indicates whether the automatic full indexing feature is enabled.
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configuration of MaxCompute data sources.
	Config *ListIndexesResponseBodyResultDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The offline deployment name of the data source.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The name of the data source.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of resources used for data update.
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// The configuration of SARO data sources.
	SaroConfig *ListIndexesResponseBodyResultDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetAutoBuildIndex(v bool) *ListIndexesResponseBodyResultDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetConfig(v *ListIndexesResponseBodyResultDataSourceInfoConfig) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Config = v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetDomain(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetName(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetProcessPartitionCount(v int32) *ListIndexesResponseBodyResultDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetSaroConfig(v *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) *ListIndexesResponseBodyResultDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetType(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Type = &v
	return s
}

type ListIndexesResponseBodyResultDataSourceInfoConfig struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Bucket       *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// A parameter related to MaxCompute.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// A parameter related to SARO.
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// A parameter related to OSS.
	OssPath   *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// A parameter related to Apsara File Storage for HDFS.
	Path    *string `json:"path,omitempty" xml:"path,omitempty"`
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// A parameter related to SARO and MaxCompute.
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetAccessKey(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetAccessSecret(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetBucket(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetEndpoint(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetNamespace(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetOssPath(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetPartition(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetPath(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetProject(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetTable(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Table = &v
	return s
}

type ListIndexesResponseBodyResultDataSourceInfoSaroConfig struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfoSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) SetNamespace(v string) *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) SetTableName(v string) *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

type ListIndexesResponseBodyResultVersions struct {
	// The description of the version.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The information about the files.
	Files []*ListIndexesResponseBodyResultVersionsFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the version.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the version. Valid values: drafting, used, unused, and trash.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last time when the version was updated.
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the version. The value is null for an edit version.
	VersionId *int32 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s ListIndexesResponseBodyResultVersions) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultVersions) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultVersions) SetDesc(v string) *ListIndexesResponseBodyResultVersions {
	s.Desc = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetFiles(v []*ListIndexesResponseBodyResultVersionsFiles) *ListIndexesResponseBodyResultVersions {
	s.Files = v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetName(v string) *ListIndexesResponseBodyResultVersions {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetStatus(v string) *ListIndexesResponseBodyResultVersions {
	s.Status = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetUpdateTime(v int64) *ListIndexesResponseBodyResultVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetVersionId(v int32) *ListIndexesResponseBodyResultVersions {
	s.VersionId = &v
	return s
}

type ListIndexesResponseBodyResultVersionsFiles struct {
	// The full path of the file.
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory.
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a template.
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The name of the file.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIndexesResponseBodyResultVersionsFiles) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultVersionsFiles) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetFullPathName(v string) *ListIndexesResponseBodyResultVersionsFiles {
	s.FullPathName = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetIsDir(v bool) *ListIndexesResponseBodyResultVersionsFiles {
	s.IsDir = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetIsTemplate(v bool) *ListIndexesResponseBodyResultVersionsFiles {
	s.IsTemplate = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetName(v string) *ListIndexesResponseBodyResultVersionsFiles {
	s.Name = &v
	return s
}

type ListIndexesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIndexesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponse) GoString() string {
	return s.String()
}

func (s *ListIndexesResponse) SetHeaders(v map[string]*string) *ListIndexesResponse {
	s.Headers = v
	return s
}

func (s *ListIndexesResponse) SetStatusCode(v int32) *ListIndexesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexesResponse) SetBody(v *ListIndexesResponseBody) *ListIndexesResponse {
	s.Body = v
	return s
}

type ListInstanceSpecsRequest struct {
	// The node type. Valid values: qrs, search, index, and cluster. qrs specifies an Query Result Searcher (QRS) worker, search specifies a searcher worker, index specifies an index node, and cluster specifies a cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsRequest) SetType(v string) *ListInstanceSpecsRequest {
	s.Type = &v
	return s
}

type ListInstanceSpecsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The specifications of the instances.
	Result []*ListInstanceSpecsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponseBody) SetRequestId(v string) *ListInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSpecsResponseBody) SetResult(v []*ListInstanceSpecsResponseBodyResult) *ListInstanceSpecsResponseBody {
	s.Result = v
	return s
}

type ListInstanceSpecsResponseBodyResult struct {
	// The number of CPU cores.
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The maximum storage space of a searcher worker.
	MaxDisk *int32 `json:"maxDisk,omitempty" xml:"maxDisk,omitempty"`
	// The memory size. Unit: GB.
	Mem *int32 `json:"mem,omitempty" xml:"mem,omitempty"`
	// The minimum storage space of a searcher worker.
	MinDisk *int32 `json:"minDisk,omitempty" xml:"minDisk,omitempty"`
}

func (s ListInstanceSpecsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponseBodyResult) SetCpu(v int32) *ListInstanceSpecsResponseBodyResult {
	s.Cpu = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMaxDisk(v int32) *ListInstanceSpecsResponseBodyResult {
	s.MaxDisk = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMem(v int32) *ListInstanceSpecsResponseBodyResult {
	s.Mem = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMinDisk(v int32) *ListInstanceSpecsResponseBodyResult {
	s.MinDisk = &v
	return s
}

type ListInstanceSpecsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponse) SetHeaders(v map[string]*string) *ListInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceSpecsResponse) SetStatusCode(v int32) *ListInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceSpecsResponse) SetBody(v *ListInstanceSpecsResponseBody) *ListInstanceSpecsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The description of the instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 实例类型，vector(向量索引版)，engine(召回引擎版)
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The time when the instance was created
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The status of the instance
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The description of the instance. You can use this description to filter instances. Fuzzy match is supported.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The number of the page to return. Default value: 1.
	ResourceGroupId *string                     `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*ListInstancesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetDescription(v string) *ListInstancesRequest {
	s.Description = &v
	return s
}

func (s *ListInstancesRequest) SetEdition(v string) *ListInstancesRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTags(v []*ListInstancesRequestTags) *ListInstancesRequest {
	s.Tags = v
	return s
}

type ListInstancesRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTags) SetKey(v string) *ListInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTags) SetValue(v string) *ListInstancesRequestTags {
	s.Value = &v
	return s
}

type ListInstancesShrinkRequest struct {
	// The description of the instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 实例类型，vector(向量索引版)，engine(召回引擎版)
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The time when the instance was created
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The status of the instance
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The description of the instance. You can use this description to filter instances. Fuzzy match is supported.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The number of the page to return. Default value: 1.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	TagsShrink      *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) SetDescription(v string) *ListInstancesShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetEdition(v string) *ListInstancesShrinkRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceId(v string) *ListInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int32) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int32) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagsShrink(v string) *ListInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListInstancesResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	Result     []*ListInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetResult(v []*ListInstancesResponseBodyResult) *ListInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyResult struct {
	// The ID of the resource group to which the instance belongs.
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The total number of entries returned
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// Havenask instance
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the virtual switch
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) network
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// The ID of the request
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The access point of the gateway
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Emergency test
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The lock status
	Network *ListInstancesResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 10.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The expiration time
	Status *string                                `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListInstancesResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the instance was last updated
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResult) SetChargeType(v string) *ListInstancesResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetCommodityCode(v string) *ListInstancesResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetCreateTime(v string) *ListInstancesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetDescription(v string) *ListInstancesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetExpiredTime(v string) *ListInstancesResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetInDebt(v bool) *ListInstancesResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetInstanceId(v string) *ListInstancesResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetLockMode(v string) *ListInstancesResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetNetwork(v *ListInstancesResponseBodyResultNetwork) *ListInstancesResponseBodyResult {
	s.Network = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetResourceGroupId(v string) *ListInstancesResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetStatus(v string) *ListInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetTags(v []*ListInstancesResponseBodyResultTags) *ListInstancesResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetUpdateTime(v string) *ListInstancesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type ListInstancesResponseBodyResultNetwork struct {
	// 353490
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// ### Sample responses
	//
	// **Sample success responses**
	//
	//     {
	//         "requestId": "90D6B8F5-FE97-4509-9AAB-367836C51818",
	//         "result": [
	//             {
	//                 "instanceId": "igraph-cn-xxxxxx1",
	//                 "spec": {
	//                     "password": "passwd",
	//                     "searchResource": {
	//                         "disk": 50,
	//                         "mem": 8,
	//                         "cpu": 2,
	//                         "nodeCount": 2
	//                     },
	//                     "instanceName": "testInstance",
	//                     "vSwitchId": "vswitch_id_xxx",
	//                     "vpcId": "vpc_id_xxx",
	//                     "qrsResource": {
	//                         "disk": 50,
	//                         "mem": 8,
	//                         "cpu": 2,
	//                         "nodeCount": 2
	//                     },
	//                     "region": "cn-hangzhou",
	//                     "userName": "user"
	//                 },
	//                 "status": {
	//                     "phase": "PENDING",
	//                     "instancePhase": "INIT",
	//                     "createSuccess": false
	//                 }
	//             },
	//             {
	//                 "instanceId": "igraph-cn-xxxxxx2",
	//                 "spec": {
	//                     "password": "passwd",
	//                     "searchResource": {
	//                         "disk": 50,
	//                         "mem": 8,
	//                         "cpu": 2,
	//                         "nodeCount": 2
	//                     },
	//                     "instanceName": "testInstance",
	//                     "vSwitchId": "vswitch_id_xxx",
	//                     "vpcId": "vpc_id_xxx",
	//                     "qrsResource": {
	//                         "disk": 50,
	//                         "mem": 8,
	//                         "cpu": 2,
	//                         "nodeCount": 2
	//                     },
	//                     "region": "cn-hangzhou",
	//                     "userName": "user"
	//                 },
	//                 "status": {
	//                     "phase": "PENDING",
	//                     "instancePhase": "INIT",
	//                     "createSuccess": false
	//                 }
	//             }
	//         ],
	//         "totalCount": 20
	//     }
	//
	// **Sample error responses**
	//
	//     {
	//       "requestId": "BD1EA715-DF6F-06C2-004C-C1FA0D3A9820",
	//       "httpCode": 404,
	//       "code": "App.NotFound",
	//       "message": "App not found"
	//     }
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// Queries instances.
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListInstancesResponseBodyResultNetwork) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultNetwork) SetEndpoint(v string) *ListInstancesResponseBodyResultNetwork {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetVSwitchId(v string) *ListInstancesResponseBodyResultNetwork {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetVpcId(v string) *ListInstancesResponseBodyResultNetwork {
	s.VpcId = &v
	return s
}

type ListInstancesResponseBodyResultTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultTags) SetKey(v string) *ListInstancesResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyResultTags) SetValue(v string) *ListInstancesResponseBodyResultTags {
	s.Value = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListOnlineConfigsRequest struct {
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s ListOnlineConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsRequest) SetDomain(v string) *ListOnlineConfigsRequest {
	s.Domain = &v
	return s
}

type ListOnlineConfigsResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListOnlineConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListOnlineConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponseBody) SetRequestId(v string) *ListOnlineConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnlineConfigsResponseBody) SetResult(v []*ListOnlineConfigsResponseBodyResult) *ListOnlineConfigsResponseBody {
	s.Result = v
	return s
}

type ListOnlineConfigsResponseBodyResult struct {
	Config    *string `json:"config,omitempty" xml:"config,omitempty"`
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s ListOnlineConfigsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponseBodyResult) SetConfig(v string) *ListOnlineConfigsResponseBodyResult {
	s.Config = &v
	return s
}

func (s *ListOnlineConfigsResponseBodyResult) SetIndexName(v string) *ListOnlineConfigsResponseBodyResult {
	s.IndexName = &v
	return s
}

type ListOnlineConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOnlineConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOnlineConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponse) SetHeaders(v map[string]*string) *ListOnlineConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineConfigsResponse) SetStatusCode(v int32) *ListOnlineConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineConfigsResponse) SetBody(v *ListOnlineConfigsResponseBody) *ListOnlineConfigsResponse {
	s.Body = v
	return s
}

type ListQueryResultRequest struct {
	// 353490
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	Sql   *string `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s ListQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListQueryResultRequest) SetQuery(v string) *ListQueryResultRequest {
	s.Query = &v
	return s
}

func (s *ListQueryResultRequest) SetSql(v string) *ListQueryResultRequest {
	s.Sql = &v
	return s
}

type ListQueryResultResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryResultResponseBody) SetRequestId(v string) *ListQueryResultResponseBody {
	s.RequestId = &v
	return s
}

type ListQueryResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListQueryResultResponse) SetHeaders(v map[string]*string) *ListQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListQueryResultResponse) SetStatusCode(v int32) *ListQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryResultResponse) SetBody(v *ListQueryResultResponseBody) *ListQueryResultResponse {
	s.Body = v
	return s
}

type ModifyAdvanceConfigFileRequest struct {
	// The content of the file.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The variable.
	Variables map[string]*VariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
	// The name of the file.
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ModifyAdvanceConfigFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileRequest) SetContent(v string) *ModifyAdvanceConfigFileRequest {
	s.Content = &v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) SetVariables(v map[string]*VariablesValue) *ModifyAdvanceConfigFileRequest {
	s.Variables = v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) SetFileName(v string) *ModifyAdvanceConfigFileRequest {
	s.FileName = &v
	return s
}

type ModifyAdvanceConfigFileResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyAdvanceConfigFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileResponseBody) SetRequestId(v string) *ModifyAdvanceConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvanceConfigFileResponseBody) SetResult(v map[string]interface{}) *ModifyAdvanceConfigFileResponseBody {
	s.Result = v
	return s
}

type ModifyAdvanceConfigFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAdvanceConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAdvanceConfigFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileResponse) SetHeaders(v map[string]*string) *ModifyAdvanceConfigFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) SetStatusCode(v int32) *ModifyAdvanceConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) SetBody(v *ModifyAdvanceConfigFileResponseBody) *ModifyAdvanceConfigFileResponse {
	s.Body = v
	return s
}

type ModifyClusterDescRequest struct {
	// The parameters in the request body
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDescRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDescRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescRequest) SetBody(v map[string]interface{}) *ModifyClusterDescRequest {
	s.Body = v
	return s
}

type ModifyClusterDescResponseBody struct {
	// The ID of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterDescResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDescResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescResponseBody) SetRequestId(v string) *ModifyClusterDescResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterDescResponseBody) SetResult(v map[string]interface{}) *ModifyClusterDescResponseBody {
	s.Result = v
	return s
}

type ModifyClusterDescResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyClusterDescResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterDescResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDescResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescResponse) SetHeaders(v map[string]*string) *ModifyClusterDescResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterDescResponse) SetStatusCode(v int32) *ModifyClusterDescResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterDescResponse) SetBody(v *ModifyClusterDescResponseBody) *ModifyClusterDescResponse {
	s.Body = v
	return s
}

type ModifyClusterOfflineConfigRequest struct {
	// The reindexing method. Valid values: api: API data source. indexRecover: data recovery through indexing.
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The configuration name, which is stored as a key.
	Config         map[string]*int32 `json:"config,omitempty" xml:"config,omitempty"`
	DataSourceName *string           `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The type of the data source. Valid values: odps: MaxCompute. swift: Swift. unKnow: unknown type.
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// This parameter is required if the API data source experiences full indexing.
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The domain in which the data source is deployed.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the backward data delivery.
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// This parameter is required if the MaxCompute data source experiences full indexing.
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	PushMode  *string `json:"pushMode,omitempty" xml:"pushMode,omitempty"`
}

func (s ModifyClusterOfflineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOfflineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigRequest) SetBuildMode(v string) *ModifyClusterOfflineConfigRequest {
	s.BuildMode = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetConfig(v map[string]*int32) *ModifyClusterOfflineConfigRequest {
	s.Config = v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataSourceName(v string) *ModifyClusterOfflineConfigRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataSourceType(v string) *ModifyClusterOfflineConfigRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataTimeSec(v int32) *ModifyClusterOfflineConfigRequest {
	s.DataTimeSec = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDomain(v string) *ModifyClusterOfflineConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetGeneration(v int64) *ModifyClusterOfflineConfigRequest {
	s.Generation = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetPartition(v string) *ModifyClusterOfflineConfigRequest {
	s.Partition = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetPushMode(v string) *ModifyClusterOfflineConfigRequest {
	s.PushMode = &v
	return s
}

type ModifyClusterOfflineConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result of the request.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterOfflineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOfflineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigResponseBody) SetRequestId(v string) *ModifyClusterOfflineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterOfflineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyClusterOfflineConfigResponseBody {
	s.Result = v
	return s
}

type ModifyClusterOfflineConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyClusterOfflineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterOfflineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOfflineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterOfflineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) SetStatusCode(v int32) *ModifyClusterOfflineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) SetBody(v *ModifyClusterOfflineConfigResponseBody) *ModifyClusterOfflineConfigResponse {
	s.Body = v
	return s
}

type ModifyClusterOnlineConfigRequest struct {
	// The information about the cluster
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// 配置信息
	Config map[string]*int32 `json:"config,omitempty" xml:"config,omitempty"`
}

func (s ModifyClusterOnlineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOnlineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigRequest) SetClusters(v []*string) *ModifyClusterOnlineConfigRequest {
	s.Clusters = v
	return s
}

func (s *ModifyClusterOnlineConfigRequest) SetConfig(v map[string]*int32) *ModifyClusterOnlineConfigRequest {
	s.Config = v
	return s
}

type ModifyClusterOnlineConfigResponseBody struct {
	// The ID of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterOnlineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOnlineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigResponseBody) SetRequestId(v string) *ModifyClusterOnlineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterOnlineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyClusterOnlineConfigResponseBody {
	s.Result = v
	return s
}

type ModifyClusterOnlineConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyClusterOnlineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterOnlineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOnlineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterOnlineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) SetStatusCode(v int32) *ModifyClusterOnlineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) SetBody(v *ModifyClusterOnlineConfigResponseBody) *ModifyClusterOnlineConfigResponse {
	s.Body = v
	return s
}

type ModifyDataSourceRequest struct {
	// The information about the index
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the request
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceRequest) SetBody(v map[string]interface{}) *ModifyDataSourceRequest {
	s.Body = v
	return s
}

func (s *ModifyDataSourceRequest) SetDryRun(v bool) *ModifyDataSourceRequest {
	s.DryRun = &v
	return s
}

type ModifyDataSourceResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The schema information.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBody) SetRequestId(v string) *ModifyDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceResponseBody) SetResult(v map[string]interface{}) *ModifyDataSourceResponseBody {
	s.Result = v
	return s
}

type ModifyDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponse) SetHeaders(v map[string]*string) *ModifyDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceResponse) SetStatusCode(v int32) *ModifyDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceResponse) SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse {
	s.Body = v
	return s
}

type ModifyFileRequest struct {
	// The parameters in the request body
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// auditing
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// ha-cn-tl32m2c4u01@ha-cn-tl32m2c4u01_00@bj_vpc_domain_1@automobile_vector@index_config_edit
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ModifyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileRequest) SetContent(v string) *ModifyFileRequest {
	s.Content = &v
	return s
}

func (s *ModifyFileRequest) SetPartition(v int32) *ModifyFileRequest {
	s.Partition = &v
	return s
}

func (s *ModifyFileRequest) SetFileName(v string) *ModifyFileRequest {
	s.FileName = &v
	return s
}

type ModifyFileResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileResponseBody) SetRequestId(v string) *ModifyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFileResponseBody) SetResult(v map[string]interface{}) *ModifyFileResponseBody {
	s.Result = v
	return s
}

type ModifyFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileResponse) SetHeaders(v map[string]*string) *ModifyFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileResponse) SetStatusCode(v int32) *ModifyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileResponse) SetBody(v *ModifyFileResponseBody) *ModifyFileResponse {
	s.Body = v
	return s
}

type ModifyIndexPartitionRequest struct {
	// The name of the data source.
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The information about each index.
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The name of the data center.
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The number of shards of the index.
	IndexInfos []*ModifyIndexPartitionRequestIndexInfos `json:"indexInfos,omitempty" xml:"indexInfos,omitempty" type:"Repeated"`
}

func (s ModifyIndexPartitionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionRequest) SetDataSourceName(v string) *ModifyIndexPartitionRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetDomainName(v string) *ModifyIndexPartitionRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetGeneration(v int64) *ModifyIndexPartitionRequest {
	s.Generation = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetIndexInfos(v []*ModifyIndexPartitionRequestIndexInfos) *ModifyIndexPartitionRequest {
	s.IndexInfos = v
	return s
}

type ModifyIndexPartitionRequestIndexInfos struct {
	// auditing
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The parameters in the request body.
	ParallelNum *int32 `json:"parallelNum,omitempty" xml:"parallelNum,omitempty"`
	// The number of shards of the index.
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
}

func (s ModifyIndexPartitionRequestIndexInfos) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionRequestIndexInfos) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetIndexName(v string) *ModifyIndexPartitionRequestIndexInfos {
	s.IndexName = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetParallelNum(v int32) *ModifyIndexPartitionRequestIndexInfos {
	s.ParallelNum = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetPartitionCount(v int32) *ModifyIndexPartitionRequestIndexInfos {
	s.PartitionCount = &v
	return s
}

type ModifyIndexPartitionResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexPartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionResponseBody) SetRequestId(v string) *ModifyIndexPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexPartitionResponseBody) SetResult(v map[string]interface{}) *ModifyIndexPartitionResponseBody {
	s.Result = v
	return s
}

type ModifyIndexPartitionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyIndexPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIndexPartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionResponse) SetHeaders(v map[string]*string) *ModifyIndexPartitionResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexPartitionResponse) SetStatusCode(v int32) *ModifyIndexPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexPartitionResponse) SetBody(v *ModifyIndexPartitionResponseBody) *ModifyIndexPartitionResponse {
	s.Body = v
	return s
}

type ModifyIndexVersionRequest struct {
	// The keyword used to search for a version. Fuzzy match is supported.
	Body []*ModifyIndexVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ModifyIndexVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionRequest) SetBody(v []*ModifyIndexVersionRequestBody) *ModifyIndexVersionRequest {
	s.Body = v
	return s
}

type ModifyIndexVersionRequestBody struct {
	// The ID of the index deployed in offline mode.
	BuildDeployId *string `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The name of the index.
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The version of the index.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModifyIndexVersionRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionRequestBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionRequestBody) SetBuildDeployId(v string) *ModifyIndexVersionRequestBody {
	s.BuildDeployId = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) SetIndexName(v string) *ModifyIndexVersionRequestBody {
	s.IndexName = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) SetVersion(v string) *ModifyIndexVersionRequestBody {
	s.Version = &v
	return s
}

type ModifyIndexVersionResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionResponseBody) SetRequestId(v string) *ModifyIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexVersionResponseBody) SetResult(v map[string]interface{}) *ModifyIndexVersionResponseBody {
	s.Result = v
	return s
}

type ModifyIndexVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionResponse) SetHeaders(v map[string]*string) *ModifyIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexVersionResponse) SetStatusCode(v int32) *ModifyIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexVersionResponse) SetBody(v *ModifyIndexVersionResponseBody) *ModifyIndexVersionResponse {
	s.Body = v
	return s
}

type ModifyNodeConfigRequest struct {
	Active              *bool  `json:"active,omitempty" xml:"active,omitempty"`
	DataDuplicateNumber *int32 `json:"dataDuplicateNumber,omitempty" xml:"dataDuplicateNumber,omitempty"`
	DataFragmentNumber  *int32 `json:"dataFragmentNumber,omitempty" xml:"dataFragmentNumber,omitempty"`
	MinServicePercent   *int32 `json:"minServicePercent,omitempty" xml:"minServicePercent,omitempty"`
	Published           *bool  `json:"published,omitempty" xml:"published,omitempty"`
	// The ID of the cluster.
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The parameters in the request body.
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The name of the cluster.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The original name of the node.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyNodeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigRequest) SetActive(v bool) *ModifyNodeConfigRequest {
	s.Active = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataDuplicateNumber(v int32) *ModifyNodeConfigRequest {
	s.DataDuplicateNumber = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataFragmentNumber(v int32) *ModifyNodeConfigRequest {
	s.DataFragmentNumber = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetMinServicePercent(v int32) *ModifyNodeConfigRequest {
	s.MinServicePercent = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetPublished(v bool) *ModifyNodeConfigRequest {
	s.Published = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetClusterName(v string) *ModifyNodeConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataSourceName(v string) *ModifyNodeConfigRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetName(v string) *ModifyNodeConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetType(v string) *ModifyNodeConfigRequest {
	s.Type = &v
	return s
}

type ModifyNodeConfigResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// auditing
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyNodeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigResponseBody) SetRequestId(v string) *ModifyNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeConfigResponseBody) SetResult(v map[string]interface{}) *ModifyNodeConfigResponseBody {
	s.Result = v
	return s
}

type ModifyNodeConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNodeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigResponse) SetHeaders(v map[string]*string) *ModifyNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeConfigResponse) SetStatusCode(v int32) *ModifyNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeConfigResponse) SetBody(v *ModifyNodeConfigResponseBody) *ModifyNodeConfigResponse {
	s.Body = v
	return s
}

type ModifyOnlineConfigRequest struct {
	// ashortdescriptionofstruct
	Body map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOnlineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnlineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigRequest) SetBody(v map[string]*string) *ModifyOnlineConfigRequest {
	s.Body = v
	return s
}

type ModifyOnlineConfigResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyOnlineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnlineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigResponseBody) SetRequestId(v string) *ModifyOnlineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOnlineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyOnlineConfigResponseBody {
	s.Result = v
	return s
}

type ModifyOnlineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyOnlineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOnlineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnlineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigResponse) SetHeaders(v map[string]*string) *ModifyOnlineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyOnlineConfigResponse) SetStatusCode(v int32) *ModifyOnlineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOnlineConfigResponse) SetBody(v *ModifyOnlineConfigResponseBody) *ModifyOnlineConfigResponse {
	s.Body = v
	return s
}

type ModifyPasswordRequest struct {
	// The password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ModifyPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyPasswordRequest) SetPassword(v string) *ModifyPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyPasswordRequest) SetUsername(v string) *ModifyPasswordRequest {
	s.Username = &v
	return s
}

type ModifyPasswordResponseBody struct {
	// The ID of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPasswordResponseBody) SetRequestId(v string) *ModifyPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPasswordResponseBody) SetResult(v map[string]interface{}) *ModifyPasswordResponseBody {
	s.Result = v
	return s
}

type ModifyPasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyPasswordResponse) SetHeaders(v map[string]*string) *ModifyPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyPasswordResponse) SetStatusCode(v int32) *ModifyPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPasswordResponse) SetBody(v *ModifyPasswordResponseBody) *ModifyPasswordResponse {
	s.Body = v
	return s
}

type PublishAdvanceConfigRequest struct {
	// The structure of the request
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishAdvanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequest) SetBody(v map[string]interface{}) *PublishAdvanceConfigRequest {
	s.Body = v
	return s
}

type PublishAdvanceConfigResponseBody struct {
	// The ID of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigResponseBody) SetRequestId(v string) *PublishAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *PublishAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type PublishAdvanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigResponse) SetHeaders(v map[string]*string) *PublishAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishAdvanceConfigResponse) SetStatusCode(v int32) *PublishAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAdvanceConfigResponse) SetBody(v *PublishAdvanceConfigResponseBody) *PublishAdvanceConfigResponse {
	s.Body = v
	return s
}

type PublishIndexVersionRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishIndexVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishIndexVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionRequest) SetBody(v map[string]interface{}) *PublishIndexVersionRequest {
	s.Body = v
	return s
}

type PublishIndexVersionResponseBody struct {
	// id of request
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionResponseBody) SetRequestId(v string) *PublishIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishIndexVersionResponseBody) SetResult(v map[string]interface{}) *PublishIndexVersionResponseBody {
	s.Result = v
	return s
}

type PublishIndexVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionResponse) SetHeaders(v map[string]*string) *PublishIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishIndexVersionResponse) SetStatusCode(v int32) *PublishIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishIndexVersionResponse) SetBody(v *PublishIndexVersionResponseBody) *PublishIndexVersionResponse {
	s.Body = v
	return s
}

type RecoverIndexRequest struct {
	// buildDeployId
	BuildDeployId *int32 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The name of the data source
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// generation
	Generation *string `json:"generation,omitempty" xml:"generation,omitempty"`
	// The name of the index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s RecoverIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverIndexRequest) GoString() string {
	return s.String()
}

func (s *RecoverIndexRequest) SetBuildDeployId(v int32) *RecoverIndexRequest {
	s.BuildDeployId = &v
	return s
}

func (s *RecoverIndexRequest) SetDataSourceName(v string) *RecoverIndexRequest {
	s.DataSourceName = &v
	return s
}

func (s *RecoverIndexRequest) SetGeneration(v string) *RecoverIndexRequest {
	s.Generation = &v
	return s
}

func (s *RecoverIndexRequest) SetIndexName(v string) *RecoverIndexRequest {
	s.IndexName = &v
	return s
}

type RecoverIndexResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RecoverIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverIndexResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverIndexResponseBody) SetRequestId(v string) *RecoverIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverIndexResponseBody) SetResult(v map[string]interface{}) *RecoverIndexResponseBody {
	s.Result = v
	return s
}

type RecoverIndexResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoverIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoverIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverIndexResponse) GoString() string {
	return s.String()
}

func (s *RecoverIndexResponse) SetHeaders(v map[string]*string) *RecoverIndexResponse {
	s.Headers = v
	return s
}

func (s *RecoverIndexResponse) SetStatusCode(v int32) *RecoverIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverIndexResponse) SetBody(v *RecoverIndexResponseBody) *RecoverIndexResponse {
	s.Body = v
	return s
}

type RemoveClusterResponseBody struct {
	// id of request
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterResponseBody) SetRequestId(v string) *RemoveClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClusterResponseBody) SetResult(v map[string]interface{}) *RemoveClusterResponseBody {
	s.Result = v
	return s
}

type RemoveClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterResponse) SetHeaders(v map[string]*string) *RemoveClusterResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterResponse) SetStatusCode(v int32) *RemoveClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterResponse) SetBody(v *RemoveClusterResponseBody) *RemoveClusterResponse {
	s.Body = v
	return s
}

type StopTaskResponseBody struct {
	// id of request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskResponseBody) SetResult(v map[string]interface{}) *StopTaskResponseBody {
	s.Result = v
	return s
}

type StopTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTaskResponse) SetHeaders(v map[string]*string) *StopTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTaskResponse) SetStatusCode(v int32) *StopTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTaskResponse) SetBody(v *StopTaskResponseBody) *StopTaskResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// The information about the instance type.
	Components []*UpdateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The description of the instance.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the order. Valid values: UPGRADE and DOWNGRADE. UPGRADE indicates the instance type is to be upgraded. DOWNGRADE indicates the instance type is to be downgraded.
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetComponents(v []*UpdateInstanceRequestComponents) *UpdateInstanceRequest {
	s.Components = v
	return s
}

func (s *UpdateInstanceRequest) SetDescription(v string) *UpdateInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceRequest) SetOrderType(v string) *UpdateInstanceRequest {
	s.OrderType = &v
	return s
}

type UpdateInstanceRequestComponents struct {
	// The name of the specification. The value must be the same as the name of a parameter on the buy page.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The value of the specification.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateInstanceRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestComponents) SetCode(v string) *UpdateInstanceRequestComponents {
	s.Code = &v
	return s
}

func (s *UpdateInstanceRequestComponents) SetValue(v string) *UpdateInstanceRequestComponents {
	s.Value = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result *UpdateInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody {
	s.Result = v
	return s
}

type UpdateInstanceResponseBodyResult struct {
	// The billing method of the instance.
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The service code.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the instance.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the instance expires.
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved.
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance.
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The state of the instance.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the instance was last updated.
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s UpdateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResult) SetChargeType(v string) *UpdateInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetCommodityCode(v string) *UpdateInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetCreateTime(v string) *UpdateInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDescription(v string) *UpdateInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetExpiredTime(v string) *UpdateInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInDebt(v bool) *UpdateInstanceResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInstanceId(v string) *UpdateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetLockMode(v string) *UpdateInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetResourceGroupId(v string) *UpdateInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetStatus(v string) *UpdateInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetUpdateTime(v string) *UpdateInstanceResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("searchengine"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * ## Method
 *     POST
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/actions/build-index
 *
 * @param request BuildIndexRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return BuildIndexResponse
 */
func (client *Client) BuildIndexWithOptions(instanceId *string, request *BuildIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BuildIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildMode)) {
		body["buildMode"] = request.BuildMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["dataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataTimeSec)) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BuildIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/actions/build-index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BuildIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 *     POST
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/actions/build-index
 *
 * @param request BuildIndexRequest
 * @return BuildIndexResponse
 */
func (client *Client) BuildIndex(instanceId *string, request *BuildIndexRequest) (_result *BuildIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BuildIndexResponse{}
	_body, _err := client.BuildIndexWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `POST`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/clusters`
 *
 * @param request CreateClusterRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateClusterResponse
 */
func (client *Client) CreateClusterWithOptions(instanceId *string, request *CreateClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoLoad)) {
		body["autoLoad"] = request.AutoLoad
	}

	if !tea.BoolValue(util.IsUnset(request.DataNode)) {
		body["dataNode"] = request.DataNode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.QueryNode)) {
		body["queryNode"] = request.QueryNode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `POST`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/clusters`
 *
 * @param request CreateClusterRequest
 * @return CreateClusterResponse
 */
func (client *Client) CreateCluster(instanceId *string, request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The result returned
 *
 * @param request CreateDataSourceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDataSourceResponse
 */
func (client *Client) CreateDataSourceWithOptions(instanceId *string, request *CreateDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoBuildIndex)) {
		body["autoBuildIndex"] = request.AutoBuildIndex
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SaroConfig)) {
		body["saroConfig"] = request.SaroConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The result returned
 *
 * @param request CreateDataSourceRequest
 * @return CreateDataSourceResponse
 */
func (client *Client) CreateDataSource(instanceId *string, request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * ```java
 * POST
 * ```
 * ### URI
 * ```java
 * /openapi/ha3/instances/{instanceId}/indexes
 * ```
 *
 * @param request CreateIndexRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateIndexResponse
 */
func (client *Client) CreateIndexWithOptions(instanceId *string, request *CreateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInfo)) {
		body["dataSourceInfo"] = request.DataSourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * ```java
 * POST
 * ```
 * ### URI
 * ```java
 * /openapi/ha3/instances/{instanceId}/indexes
 * ```
 *
 * @param request CreateIndexRequest
 * @return CreateIndexResponse
 */
func (client *Client) CreateIndex(instanceId *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `POST`
 * ### URI
 * `/api/instances?dryRun=false`
 *
 * @param request CreateInstanceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateInstanceResponse
 */
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		body["components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `POST`
 * ### URI
 * `/api/instances?dryRun=false`
 *
 * @param request CreateInstanceRequest
 * @return CreateInstanceResponse
 */
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Method
 *     DELETE
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAdvanceConfigResponse
 */
func (client *Client) DeleteAdvanceConfigWithOptions(instanceId *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAdvanceConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 *     DELETE
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
 *
 * @return DeleteAdvanceConfigResponse
 */
func (client *Client) DeleteAdvanceConfig(instanceId *string, configName *string) (_result *DeleteAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAdvanceConfigResponse{}
	_body, _err := client.DeleteAdvanceConfigWithOptions(instanceId, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Method
 * `DELETE`
 * ## URI
 * `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDataSourceResponse
 */
func (client *Client) DeleteDataSourceWithOptions(instanceId *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 * `DELETE`
 * ## URI
 * `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
 *
 * @return DeleteDataSourceResponse
 */
func (client *Client) DeleteDataSource(instanceId *string, dataSourceName *string) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(instanceId, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The information about the index
 *
 * @param request DeleteIndexRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteIndexResponse
 */
func (client *Client) DeleteIndexWithOptions(instanceId *string, indexName *string, request *DeleteIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		query["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteDataSource)) {
		query["deleteDataSource"] = request.DeleteDataSource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The information about the index
 *
 * @param request DeleteIndexRequest
 * @return DeleteIndexResponse
 */
func (client *Client) DeleteIndex(instanceId *string, indexName *string, request *DeleteIndexRequest) (_result *DeleteIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(instanceId, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The result
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteIndexVersionResponse
 */
func (client *Client) DeleteIndexVersionWithOptions(instanceId *string, indexName *string, versionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The result
 *
 * @return DeleteIndexVersionResponse
 */
func (client *Client) DeleteIndexVersion(instanceId *string, indexName *string, versionName *string) (_result *DeleteIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexVersionResponse{}
	_body, _err := client.DeleteIndexVersionWithOptions(instanceId, indexName, versionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The result returned
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteInstanceResponse
 */
func (client *Client) DeleteInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The result returned
 *
 * @return DeleteInstanceResponse
 */
func (client *Client) DeleteInstance(instanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * \\### Method
 * ```java
 * PUT
 * ```
 * ### URI
 * ```java
 * /openapi/ha3/instances/{instanceId}/force-switch/{fsmId}
 * ```
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ForceSwitchResponse
 */
func (client *Client) ForceSwitchWithOptions(instanceId *string, fsmId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ForceSwitchResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ForceSwitch"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/force-switch/" + tea.StringValue(openapiutil.GetEncodeParam(fsmId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ForceSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * \\### Method
 * ```java
 * PUT
 * ```
 * ### URI
 * ```java
 * /openapi/ha3/instances/{instanceId}/force-switch/{fsmId}
 * ```
 *
 * @return ForceSwitchResponse
 */
func (client *Client) ForceSwitch(instanceId *string, fsmId *string) (_result *ForceSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ForceSwitchResponse{}
	_body, _err := client.ForceSwitchWithOptions(instanceId, fsmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Method
 *     GET
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
 *
 * @param request GetAdvanceConfigRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAdvanceConfigResponse
 */
func (client *Client) GetAdvanceConfigWithOptions(instanceId *string, configName *string, request *GetAdvanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAdvanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 *     GET
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
 *
 * @param request GetAdvanceConfigRequest
 * @return GetAdvanceConfigResponse
 */
func (client *Client) GetAdvanceConfig(instanceId *string, configName *string, request *GetAdvanceConfigRequest) (_result *GetAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAdvanceConfigResponse{}
	_body, _err := client.GetAdvanceConfigWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAdvanceConfigFileWithOptions(instanceId *string, configName *string, request *GetAdvanceConfigFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAdvanceConfigFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAdvanceConfigFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/file"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdvanceConfigFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAdvanceConfigFile(instanceId *string, configName *string, request *GetAdvanceConfigFileRequest) (_result *GetAdvanceConfigFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAdvanceConfigFileResponse{}
	_body, _err := client.GetAdvanceConfigFileWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `GET`
 * ### URI
 * `/openapi/ha3/instance/{instanceId}/clusters/{clusterName}`
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetClusterResponse
 */
func (client *Client) GetClusterWithOptions(instanceId *string, clusterName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `GET`
 * ### URI
 * `/openapi/ha3/instance/{instanceId}/clusters/{clusterName}`
 *
 * @return GetClusterResponse
 */
func (client *Client) GetCluster(instanceId *string, clusterName *string) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(instanceId, clusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterRunTimeInfoWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterRunTimeInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterRunTimeInfo"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-run-time-info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterRunTimeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterRunTimeInfo(instanceId *string) (_result *GetClusterRunTimeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterRunTimeInfoResponse{}
	_body, _err := client.GetClusterRunTimeInfoWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataSourceWithOptions(instanceId *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataSource(instanceId *string, dataSourceName *string) (_result *GetDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSourceResponse{}
	_body, _err := client.GetDataSourceWithOptions(instanceId, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataSourceDeployWithOptions(instanceId *string, deployName *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSourceDeployResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSourceDeploy"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/deploys/" + tea.StringValue(openapiutil.GetEncodeParam(deployName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceDeployResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataSourceDeploy(instanceId *string, deployName *string, dataSourceName *string) (_result *GetDataSourceDeployResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSourceDeployResponse{}
	_body, _err := client.GetDataSourceDeployWithOptions(instanceId, deployName, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Sample requests
 * ```java
 * GET /openapi/ha3/instances/{instanceId}/deploy-graph
 * ```
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDeployGraphResponse
 */
func (client *Client) GetDeployGraphWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDeployGraphResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployGraph"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/deploy-graph"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeployGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Sample requests
 * ```java
 * GET /openapi/ha3/instances/{instanceId}/deploy-graph
 * ```
 *
 * @return GetDeployGraphResponse
 */
func (client *Client) GetDeployGraph(instanceId *string) (_result *GetDeployGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeployGraphResponse{}
	_body, _err := client.GetDeployGraphWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileWithOptions(instanceId *string, indexName *string, versionName *string, request *GetFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionName)) + "/file"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFile(instanceId *string, indexName *string, versionName *string, request *GetFileRequest) (_result *GetFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(instanceId, indexName, versionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIndexWithOptions(instanceId *string, indexName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIndex(instanceId *string, indexName *string) (_result *GetIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexResponse{}
	_body, _err := client.GetIndexWithOptions(instanceId, indexName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Examples
 * Sample requests
 *     GET  /openapi/ha3/instances/ha3_instance_id_1/clusters/cluster1/index-version
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetIndexVersionResponse
 */
func (client *Client) GetIndexVersionWithOptions(instanceId *string, clusterName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName)) + "/index-version"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Examples
 * Sample requests
 *     GET  /openapi/ha3/instances/ha3_instance_id_1/clusters/cluster1/index-version
 *
 * @return GetIndexVersionResponse
 */
func (client *Client) GetIndexVersion(instanceId *string, clusterName *string) (_result *GetIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexVersionResponse{}
	_body, _err := client.GetIndexVersionWithOptions(instanceId, clusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The billing method.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetInstanceResponse
 */
func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The billing method.
 *
 * @return GetInstanceResponse
 */
func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeConfigWithOptions(instanceId *string, request *GetNodeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNodeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["clusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeConfig(instanceId *string, request *GetNodeConfigRequest) (_result *GetNodeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeConfigResponse{}
	_body, _err := client.GetNodeConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Sample requests
 * `GET /openapi/ha3/instances/ose-test1/advanced-configs`
 *
 * @param request ListAdvanceConfigDirRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAdvanceConfigDirResponse
 */
func (client *Client) ListAdvanceConfigDirWithOptions(instanceId *string, configName *string, request *ListAdvanceConfigDirRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAdvanceConfigDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirName)) {
		query["dirName"] = request.DirName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdvanceConfigDir"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/dir"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdvanceConfigDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Sample requests
 * `GET /openapi/ha3/instances/ose-test1/advanced-configs`
 *
 * @param request ListAdvanceConfigDirRequest
 * @return ListAdvanceConfigDirResponse
 */
func (client *Client) ListAdvanceConfigDir(instanceId *string, configName *string, request *ListAdvanceConfigDirRequest) (_result *ListAdvanceConfigDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAdvanceConfigDirResponse{}
	_body, _err := client.ListAdvanceConfigDirWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * http
 *
 * @param request ListAdvanceConfigsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAdvanceConfigsResponse
 */
func (client *Client) ListAdvanceConfigsWithOptions(instanceId *string, request *ListAdvanceConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAdvanceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		query["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.IndexName)) {
		query["indexName"] = request.IndexName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdvanceConfigs"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdvanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * http
 *
 * @param request ListAdvanceConfigsRequest
 * @return ListAdvanceConfigsResponse
 */
func (client *Client) ListAdvanceConfigs(instanceId *string, request *ListAdvanceConfigsRequest) (_result *ListAdvanceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAdvanceConfigsResponse{}
	_body, _err := client.ListAdvanceConfigsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Sample requests
 * ```java
 * GET /openapi/ha3/instances/ha3_instance_name/cluster-names
 * ```
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListClusterNamesResponse
 */
func (client *Client) ListClusterNamesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterNamesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterNames"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/cluster-names"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Sample requests
 * ```java
 * GET /openapi/ha3/instances/ha3_instance_name/cluster-names
 * ```
 *
 * @return ListClusterNamesResponse
 */
func (client *Client) ListClusterNames() (_result *ListClusterNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterNamesResponse{}
	_body, _err := client.ListClusterNamesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterTasksWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterTasks"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterTasks(instanceId *string) (_result *ListClusterTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterTasksResponse{}
	_body, _err := client.ListClusterTasksWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * http
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListClustersResponse
 */
func (client *Client) ListClustersWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * http
 *
 * @return ListClustersResponse
 */
func (client *Client) ListClusters(instanceId *string) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Obtains the schema information of a specified data source.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDataSourceSchemasResponse
 */
func (client *Client) ListDataSourceSchemasWithOptions(instanceId *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceSchemasResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceSchemas"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/schemas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Obtains the schema information of a specified data source.
 *
 * @return ListDataSourceSchemasResponse
 */
func (client *Client) ListDataSourceSchemas(instanceId *string, dataSourceName *string) (_result *ListDataSourceSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceSchemasResponse{}
	_body, _err := client.ListDataSourceSchemasWithOptions(instanceId, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceTasksWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTasks"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-source-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSourceTasks(instanceId *string) (_result *ListDataSourceTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceTasksResponse{}
	_body, _err := client.ListDataSourceTasksWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourcesWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSources(instanceId *string) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `GET`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/generations?domainName={domainName}`
 *
 * @param request ListDateSourceGenerationsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDateSourceGenerationsResponse
 */
func (client *Client) ListDateSourceGenerationsWithOptions(instanceId *string, dataSourceName *string, request *ListDateSourceGenerationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDateSourceGenerationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ValidStatus)) {
		query["validStatus"] = request.ValidStatus
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDateSourceGenerations"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/generations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDateSourceGenerationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `GET`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/generations?domainName={domainName}`
 *
 * @param request ListDateSourceGenerationsRequest
 * @return ListDateSourceGenerationsResponse
 */
func (client *Client) ListDateSourceGenerations(instanceId *string, dataSourceName *string, request *ListDateSourceGenerationsRequest) (_result *ListDateSourceGenerationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDateSourceGenerationsResponse{}
	_body, _err := client.ListDateSourceGenerationsWithOptions(instanceId, dataSourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIndexesWithOptions(instanceId *string, request *ListIndexesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndexesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewMode)) {
		query["newMode"] = request.NewMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndexes"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndexesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIndexes(instanceId *string, request *ListIndexesRequest) (_result *ListIndexesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexesResponse{}
	_body, _err := client.ListIndexesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `GET`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/specs?type=qrs`
 *
 * @param request ListInstanceSpecsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListInstanceSpecsResponse
 */
func (client *Client) ListInstanceSpecsWithOptions(instanceId *string, request *ListInstanceSpecsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceSpecs"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/specs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `GET`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/specs?type=qrs`
 *
 * @param request ListInstanceSpecsRequest
 * @return ListInstanceSpecsResponse
 */
func (client *Client) ListInstanceSpecs(instanceId *string, request *ListInstanceSpecsRequest) (_result *ListInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceSpecsResponse{}
	_body, _err := client.ListInstanceSpecsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Indicates whether an overdue payment is involved
 *
 * @param tmpReq ListInstancesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListInstancesResponse
 */
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		query["edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Indicates whether an overdue payment is involved
 *
 * @param request ListInstancesRequest
 * @return ListInstancesResponse
 */
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * \\### Sample requests
 * ```java
 * GET  /openapi/ha3/instances/ha-test1/node/ihome_searcher/online-configs?domain=pre_ea120
 * ```
 *
 * @param request ListOnlineConfigsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListOnlineConfigsResponse
 */
func (client *Client) ListOnlineConfigsWithOptions(instanceId *string, nodeName *string, request *ListOnlineConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOnlineConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOnlineConfigs"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node/" + tea.StringValue(openapiutil.GetEncodeParam(nodeName)) + "/online-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOnlineConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * \\### Sample requests
 * ```java
 * GET  /openapi/ha3/instances/ha-test1/node/ihome_searcher/online-configs?domain=pre_ea120
 * ```
 *
 * @param request ListOnlineConfigsRequest
 * @return ListOnlineConfigsResponse
 */
func (client *Client) ListOnlineConfigs(instanceId *string, nodeName *string, request *ListOnlineConfigsRequest) (_result *ListOnlineConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOnlineConfigsResponse{}
	_body, _err := client.ListOnlineConfigsWithOptions(instanceId, nodeName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the instance
 *
 * @param request ListQueryResultRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListQueryResultResponse
 */
func (client *Client) ListQueryResultWithOptions(instanceId *string, request *ListQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		query["sql"] = request.Sql
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryResult"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the instance
 *
 * @param request ListQueryResultRequest
 * @return ListQueryResultResponse
 */
func (client *Client) ListQueryResult(instanceId *string, request *ListQueryResultRequest) (_result *ListQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryResultResponse{}
	_body, _err := client.ListQueryResultWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Method
 *     put
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
 *
 * @param request ModifyAdvanceConfigFileRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAdvanceConfigFileResponse
 */
func (client *Client) ModifyAdvanceConfigFileWithOptions(instanceId *string, configName *string, request *ModifyAdvanceConfigFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAdvanceConfigFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAdvanceConfigFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAdvanceConfigFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 *     put
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
 *
 * @param request ModifyAdvanceConfigFileRequest
 * @return ModifyAdvanceConfigFileResponse
 */
func (client *Client) ModifyAdvanceConfigFile(instanceId *string, configName *string, request *ModifyAdvanceConfigFileRequest) (_result *ModifyAdvanceConfigFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAdvanceConfigFileResponse{}
	_body, _err := client.ModifyAdvanceConfigFileWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/desc`
 *
 * @param request ModifyClusterDescRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyClusterDescResponse
 */
func (client *Client) ModifyClusterDescWithOptions(instanceId *string, clusterName *string, request *ModifyClusterDescRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterDescResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterDesc"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName)) + "/desc"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterDescResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/desc`
 *
 * @param request ModifyClusterDescRequest
 * @return ModifyClusterDescResponse
 */
func (client *Client) ModifyClusterDesc(instanceId *string, clusterName *string, request *ModifyClusterDescRequest) (_result *ModifyClusterDescResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterDescResponse{}
	_body, _err := client.ModifyClusterDescWithOptions(instanceId, clusterName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Request syntax
 *     PUT /openapi/ha3/instances/{instanceId}/cluster-offline-config
 *     ...
 *
 * @param request ModifyClusterOfflineConfigRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyClusterOfflineConfigResponse
 */
func (client *Client) ModifyClusterOfflineConfigWithOptions(instanceId *string, request *ModifyClusterOfflineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterOfflineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildMode)) {
		body["buildMode"] = request.BuildMode
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["dataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataTimeSec)) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.PushMode)) {
		body["pushMode"] = request.PushMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterOfflineConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-offline-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterOfflineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Request syntax
 *     PUT /openapi/ha3/instances/{instanceId}/cluster-offline-config
 *     ...
 *
 * @param request ModifyClusterOfflineConfigRequest
 * @return ModifyClusterOfflineConfigResponse
 */
func (client *Client) ModifyClusterOfflineConfig(instanceId *string, request *ModifyClusterOfflineConfigRequest) (_result *ModifyClusterOfflineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterOfflineConfigResponse{}
	_body, _err := client.ModifyClusterOfflineConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/cluster-online-config`
 *
 * @param request ModifyClusterOnlineConfigRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyClusterOnlineConfigResponse
 */
func (client *Client) ModifyClusterOnlineConfigWithOptions(instanceId *string, request *ModifyClusterOnlineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterOnlineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Clusters)) {
		body["clusters"] = request.Clusters
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterOnlineConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-online-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterOnlineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/cluster-online-config`
 *
 * @param request ModifyClusterOnlineConfigRequest
 * @return ModifyClusterOnlineConfigResponse
 */
func (client *Client) ModifyClusterOnlineConfig(instanceId *string, request *ModifyClusterOnlineConfigRequest) (_result *ModifyClusterOnlineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterOnlineConfigResponse{}
	_body, _err := client.ModifyClusterOnlineConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The result returned
 *
 * @param request ModifyDataSourceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDataSourceResponse
 */
func (client *Client) ModifyDataSourceWithOptions(instanceId *string, dataSourceName *string, request *ModifyDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The result returned
 *
 * @param request ModifyDataSourceRequest
 * @return ModifyDataSourceResponse
 */
func (client *Client) ModifyDataSource(instanceId *string, dataSourceName *string, request *ModifyDataSourceRequest) (_result *ModifyDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.ModifyDataSourceWithOptions(instanceId, dataSourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * ~~~
 * PUT
 * ~~~
 * ### URI
 * ~~~
 * /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
 * ~~~
 *
 * @param request ModifyFileRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyFileResponse
 */
func (client *Client) ModifyFileWithOptions(instanceId *string, indexName *string, versionName *string, request *ModifyFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionName)) + "/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * ~~~
 * PUT
 * ~~~
 * ### URI
 * ~~~
 * /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
 * ~~~
 *
 * @param request ModifyFileRequest
 * @return ModifyFileResponse
 */
func (client *Client) ModifyFile(instanceId *string, indexName *string, versionName *string, request *ModifyFileRequest) (_result *ModifyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyFileResponse{}
	_body, _err := client.ModifyFileWithOptions(instanceId, indexName, versionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The information about each index.
 *
 * @param request ModifyIndexPartitionRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyIndexPartitionResponse
 */
func (client *Client) ModifyIndexPartitionWithOptions(instanceId *string, request *ModifyIndexPartitionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyIndexPartitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.IndexInfos)) {
		body["indexInfos"] = request.IndexInfos
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIndexPartition"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/index-partition"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIndexPartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The information about each index.
 *
 * @param request ModifyIndexPartitionRequest
 * @return ModifyIndexPartitionResponse
 */
func (client *Client) ModifyIndexPartition(instanceId *string, request *ModifyIndexPartitionRequest) (_result *ModifyIndexPartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyIndexPartitionResponse{}
	_body, _err := client.ModifyIndexPartitionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Method
 *     PUT
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
 *
 * @param request ModifyIndexVersionRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyIndexVersionResponse
 */
func (client *Client) ModifyIndexVersionWithOptions(instanceId *string, clusterName *string, request *ModifyIndexVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyIndexVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName)) + "/index-version"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 *     PUT
 * ## URI
 *     /openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
 *
 * @param request ModifyIndexVersionRequest
 * @return ModifyIndexVersionResponse
 */
func (client *Client) ModifyIndexVersion(instanceId *string, clusterName *string, request *ModifyIndexVersionRequest) (_result *ModifyIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyIndexVersionResponse{}
	_body, _err := client.ModifyIndexVersionWithOptions(instanceId, clusterName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 *  ~~~
 * PUT
 * ~~~
 * ### URI
 * ~~~
 * /openapi/ha3/instances/{instanceId}/node-config?type=qrs&name=test
 * ~~~
 *
 * @param request ModifyNodeConfigRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyNodeConfigResponse
 */
func (client *Client) ModifyNodeConfigWithOptions(instanceId *string, request *ModifyNodeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyNodeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["clusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		query["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.DataDuplicateNumber)) {
		body["dataDuplicateNumber"] = request.DataDuplicateNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DataFragmentNumber)) {
		body["dataFragmentNumber"] = request.DataFragmentNumber
	}

	if !tea.BoolValue(util.IsUnset(request.MinServicePercent)) {
		body["minServicePercent"] = request.MinServicePercent
	}

	if !tea.BoolValue(util.IsUnset(request.Published)) {
		body["published"] = request.Published
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodeConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNodeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 *  ~~~
 * PUT
 * ~~~
 * ### URI
 * ~~~
 * /openapi/ha3/instances/{instanceId}/node-config?type=qrs&name=test
 * ~~~
 *
 * @param request ModifyNodeConfigRequest
 * @return ModifyNodeConfigResponse
 */
func (client *Client) ModifyNodeConfig(instanceId *string, request *ModifyNodeConfigRequest) (_result *ModifyNodeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodeConfigResponse{}
	_body, _err := client.ModifyNodeConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * ```java
 * put
 * ```
 * ### URI
 * ```java
 * /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs/{indexName}
 * ```
 *
 * @param request ModifyOnlineConfigRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyOnlineConfigResponse
 */
func (client *Client) ModifyOnlineConfigWithOptions(instanceId *string, nodeName *string, indexName *string, request *ModifyOnlineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyOnlineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOnlineConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node/" + tea.StringValue(openapiutil.GetEncodeParam(nodeName)) + "/online-configs/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOnlineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * ```java
 * put
 * ```
 * ### URI
 * ```java
 * /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs/{indexName}
 * ```
 *
 * @param request ModifyOnlineConfigRequest
 * @return ModifyOnlineConfigResponse
 */
func (client *Client) ModifyOnlineConfig(instanceId *string, nodeName *string, indexName *string, request *ModifyOnlineConfigRequest) (_result *ModifyOnlineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyOnlineConfigResponse{}
	_body, _err := client.ModifyOnlineConfigWithOptions(instanceId, nodeName, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/password`
 *
 * @param request ModifyPasswordRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyPasswordResponse
 */
func (client *Client) ModifyPasswordWithOptions(instanceId *string, request *ModifyPasswordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPassword"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/password"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/password`
 *
 * @param request ModifyPasswordRequest
 * @return ModifyPasswordResponse
 */
func (client *Client) ModifyPassword(instanceId *string, request *ModifyPasswordRequest) (_result *ModifyPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyPasswordResponse{}
	_body, _err := client.ModifyPasswordWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Method
 * ~~~
 * POST
 * ~~~
 * ## URI
 * ~~~
 * /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/actions/publish
 * ~~~
 *
 * @param request PublishAdvanceConfigRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return PublishAdvanceConfigResponse
 */
func (client *Client) PublishAdvanceConfigWithOptions(instanceId *string, configName *string, request *PublishAdvanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishAdvanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/actions/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Method
 * ~~~
 * POST
 * ~~~
 * ## URI
 * ~~~
 * /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/actions/publish
 * ~~~
 *
 * @param request PublishAdvanceConfigRequest
 * @return PublishAdvanceConfigResponse
 */
func (client *Client) PublishAdvanceConfig(instanceId *string, configName *string, request *PublishAdvanceConfigRequest) (_result *PublishAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishAdvanceConfigResponse{}
	_body, _err := client.PublishAdvanceConfigWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The information about the index
 *
 * @param request PublishIndexVersionRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return PublishIndexVersionResponse
 */
func (client *Client) PublishIndexVersionWithOptions(instanceId *string, indexName *string, request *PublishIndexVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishIndexVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/actions/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The information about the index
 *
 * @param request PublishIndexVersionRequest
 * @return PublishIndexVersionResponse
 */
func (client *Client) PublishIndexVersion(instanceId *string, indexName *string, request *PublishIndexVersionRequest) (_result *PublishIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishIndexVersionResponse{}
	_body, _err := client.PublishIndexVersionWithOptions(instanceId, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `POST`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/recover-index`
 *
 * @param request RecoverIndexRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return RecoverIndexResponse
 */
func (client *Client) RecoverIndexWithOptions(instanceId *string, request *RecoverIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecoverIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildDeployId)) {
		body["buildDeployId"] = request.BuildDeployId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.IndexName)) {
		body["indexName"] = request.IndexName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/recover-index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `POST`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}/recover-index`
 *
 * @param request RecoverIndexRequest
 * @return RecoverIndexResponse
 */
func (client *Client) RecoverIndex(instanceId *string, request *RecoverIndexRequest) (_result *RecoverIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecoverIndexResponse{}
	_body, _err := client.RecoverIndexWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The result
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveClusterResponse
 */
func (client *Client) RemoveClusterWithOptions(instanceId *string, clusterName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveClusterResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveCluster"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The result
 *
 * @return RemoveClusterResponse
 */
func (client *Client) RemoveCluster(instanceId *string, clusterName *string) (_result *RemoveClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveClusterResponse{}
	_body, _err := client.RemoveClusterWithOptions(instanceId, clusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The information about the index
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopTaskResponse
 */
func (client *Client) StopTaskWithOptions(instanceId *string, fsmId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopTask"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/stop-task/" + tea.StringValue(openapiutil.GetEncodeParam(fsmId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The information about the index
 *
 * @return StopTaskResponse
 */
func (client *Client) StopTask(instanceId *string, fsmId *string) (_result *StopTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTaskResponse{}
	_body, _err := client.StopTaskWithOptions(instanceId, fsmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}`
 *
 * @param request UpdateInstanceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateInstanceResponse
 */
func (client *Client) UpdateInstanceWithOptions(instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Components)) {
		body["components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		body["orderType"] = request.OrderType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Method
 * `PUT`
 * ### URI
 * `/openapi/ha3/instances/{instanceId}`
 *
 * @param request UpdateInstanceRequest
 * @return UpdateInstanceResponse
 */
func (client *Client) UpdateInstance(instanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
