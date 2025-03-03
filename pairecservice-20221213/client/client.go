// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type TrafficControlTaskTrafficInfoTargetTrafficsDataValue struct {
	Traffic    *float64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
	RecordTime *int64   `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
}

func (s TrafficControlTaskTrafficInfoTargetTrafficsDataValue) String() string {
	return tea.Prettify(s)
}

func (s TrafficControlTaskTrafficInfoTargetTrafficsDataValue) GoString() string {
	return s.String()
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) SetTraffic(v float64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	s.Traffic = &v
	return s
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) SetRecordTime(v int64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	s.RecordTime = &v
	return s
}

type TrafficControlTaskTrafficInfoTaskTrafficsValue struct {
	Traffic *float64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s TrafficControlTaskTrafficInfoTaskTrafficsValue) String() string {
	return tea.Prettify(s)
}

func (s TrafficControlTaskTrafficInfoTaskTrafficsValue) GoString() string {
	return s.String()
}

func (s *TrafficControlTaskTrafficInfoTaskTrafficsValue) SetTraffic(v float64) *TrafficControlTaskTrafficInfoTaskTrafficsValue {
	s.Traffic = &v
	return s
}

type ExperimentReportValue struct {
	// example:
	//
	// true
	Baseline      *bool                             `json:"Baseline,omitempty" xml:"Baseline,omitempty"`
	MetricResults map[string]map[string]interface{} `json:"MetricResults,omitempty" xml:"MetricResults,omitempty"`
}

func (s ExperimentReportValue) String() string {
	return tea.Prettify(s)
}

func (s ExperimentReportValue) GoString() string {
	return s.String()
}

func (s *ExperimentReportValue) SetBaseline(v bool) *ExperimentReportValue {
	s.Baseline = &v
	return s
}

func (s *ExperimentReportValue) SetMetricResults(v map[string]map[string]interface{}) *ExperimentReportValue {
	s.MetricResults = v
	return s
}

type ApplyEngineConfigRequest struct {
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ApplyEngineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *ApplyEngineConfigRequest) SetInstanceId(v string) *ApplyEngineConfigRequest {
	s.InstanceId = &v
	return s
}

type ApplyEngineConfigResponseBody struct {
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyEngineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyEngineConfigResponseBody) SetRequestId(v string) *ApplyEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

type ApplyEngineConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyEngineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *ApplyEngineConfigResponse) SetHeaders(v map[string]*string) *ApplyEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *ApplyEngineConfigResponse) SetStatusCode(v int32) *ApplyEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyEngineConfigResponse) SetBody(v *ApplyEngineConfigResponseBody) *ApplyEngineConfigResponse {
	s.Body = v
	return s
}

type BackflowFeatureConsistencyCheckJobDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\\"{\\\\\\"itemid\\\\\\":{\\\\\\"value\\\\\\":1010,\\\\\\"type\\\\\\":\\\\\\"string\\\\\\"}}\\"]
	ItemFeatures *string `json:"ItemFeatures,omitempty" xml:"ItemFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1693900981465
	LogRequestTime *int64 `json:"LogRequestTime,omitempty" xml:"LogRequestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video-feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\\"{\\\\\\"dbmtl_probs_is_valid_play\\\\\\":0.00032182207107543945,\\\\\\"dbmtl_y_play_time\\\\\\":0.0043269748210906982}\\"]
	Scores      *string `json:"Scores,omitempty" xml:"Scores,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"userid\\":{\\"value\\":1010,\\"type\\":\\"string\\"},\\"click_5_seq\\":{\\"value\\":\\"9001;9002;9003;9004;9005\\",\\"type\\":\\"string\\"}}
	UserFeatures *string `json:"UserFeatures,omitempty" xml:"UserFeatures,omitempty"`
}

func (s BackflowFeatureConsistencyCheckJobDataRequest) String() string {
	return tea.Prettify(s)
}

func (s BackflowFeatureConsistencyCheckJobDataRequest) GoString() string {
	return s.String()
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetFeatureConsistencyCheckJobConfigId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetInstanceId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.InstanceId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetItemFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.ItemFeatures = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogItemId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogItemId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogRequestId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogRequestId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogRequestTime(v int64) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogRequestTime = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetLogUserId(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.LogUserId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetSceneName(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.SceneName = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetScores(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.Scores = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetServiceName(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.ServiceName = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetUserFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.UserFeatures = &v
	return s
}

type BackflowFeatureConsistencyCheckJobDataResponseBody struct {
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BackflowFeatureConsistencyCheckJobDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BackflowFeatureConsistencyCheckJobDataResponseBody) GoString() string {
	return s.String()
}

func (s *BackflowFeatureConsistencyCheckJobDataResponseBody) SetRequestId(v string) *BackflowFeatureConsistencyCheckJobDataResponseBody {
	s.RequestId = &v
	return s
}

type BackflowFeatureConsistencyCheckJobDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackflowFeatureConsistencyCheckJobDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackflowFeatureConsistencyCheckJobDataResponse) String() string {
	return tea.Prettify(s)
}

func (s BackflowFeatureConsistencyCheckJobDataResponse) GoString() string {
	return s.String()
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) SetHeaders(v map[string]*string) *BackflowFeatureConsistencyCheckJobDataResponse {
	s.Headers = v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) SetStatusCode(v int32) *BackflowFeatureConsistencyCheckJobDataResponse {
	s.StatusCode = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) SetBody(v *BackflowFeatureConsistencyCheckJobDataResponseBody) *BackflowFeatureConsistencyCheckJobDataResponse {
	s.Body = v
	return s
}

type CheckInstanceResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CheckInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesRequest) SetType(v string) *CheckInstanceResourcesRequest {
	s.Type = &v
	return s
}

func (s *CheckInstanceResourcesRequest) SetUri(v string) *CheckInstanceResourcesRequest {
	s.Uri = &v
	return s
}

type CheckInstanceResourcesResponseBody struct {
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*CheckInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesResponseBody) SetRequestId(v string) *CheckInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceResourcesResponseBody) SetResources(v []*CheckInstanceResourcesResponseBodyResources) *CheckInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

type CheckInstanceResourcesResponseBodyResources struct {
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CheckInstanceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesResponseBodyResources) SetStatus(v string) *CheckInstanceResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *CheckInstanceResourcesResponseBodyResources) SetType(v string) *CheckInstanceResourcesResponseBodyResources {
	s.Type = &v
	return s
}

func (s *CheckInstanceResourcesResponseBodyResources) SetUri(v string) *CheckInstanceResourcesResponseBodyResources {
	s.Uri = &v
	return s
}

type CheckInstanceResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesResponse) SetHeaders(v map[string]*string) *CheckInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceResourcesResponse) SetStatusCode(v int32) *CheckInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceResourcesResponse) SetBody(v *CheckInstanceResourcesResponseBody) *CheckInstanceResourcesResponse {
	s.Body = v
	return s
}

type CheckTrafficControlTaskExpressionRequest struct {
	// This parameter is required.
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
}

func (s CheckTrafficControlTaskExpressionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckTrafficControlTaskExpressionRequest) GoString() string {
	return s.String()
}

func (s *CheckTrafficControlTaskExpressionRequest) SetExpression(v string) *CheckTrafficControlTaskExpressionRequest {
	s.Expression = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionRequest) SetInstanceId(v string) *CheckTrafficControlTaskExpressionRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionRequest) SetTableMetaId(v string) *CheckTrafficControlTaskExpressionRequest {
	s.TableMetaId = &v
	return s
}

type CheckTrafficControlTaskExpressionResponseBody struct {
	IsValie   *bool   `json:"IsValie,omitempty" xml:"IsValie,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckTrafficControlTaskExpressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckTrafficControlTaskExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTrafficControlTaskExpressionResponseBody) SetIsValie(v bool) *CheckTrafficControlTaskExpressionResponseBody {
	s.IsValie = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponseBody) SetReason(v string) *CheckTrafficControlTaskExpressionResponseBody {
	s.Reason = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponseBody) SetRequestId(v string) *CheckTrafficControlTaskExpressionResponseBody {
	s.RequestId = &v
	return s
}

type CheckTrafficControlTaskExpressionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTrafficControlTaskExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTrafficControlTaskExpressionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckTrafficControlTaskExpressionResponse) GoString() string {
	return s.String()
}

func (s *CheckTrafficControlTaskExpressionResponse) SetHeaders(v map[string]*string) *CheckTrafficControlTaskExpressionResponse {
	s.Headers = v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponse) SetStatusCode(v int32) *CheckTrafficControlTaskExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponse) SetBody(v *CheckTrafficControlTaskExpressionResponseBody) *CheckTrafficControlTaskExpressionResponse {
	s.Body = v
	return s
}

type CloneEngineConfigRequest struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneEngineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigRequest) SetConfigValue(v string) *CloneEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *CloneEngineConfigRequest) SetDescription(v string) *CloneEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *CloneEngineConfigRequest) SetEnvironment(v string) *CloneEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *CloneEngineConfigRequest) SetInstanceId(v string) *CloneEngineConfigRequest {
	s.InstanceId = &v
	return s
}

type CloneEngineConfigResponseBody struct {
	// example:
	//
	// 2
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneEngineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigResponseBody) SetEngineConfigId(v string) *CloneEngineConfigResponseBody {
	s.EngineConfigId = &v
	return s
}

func (s *CloneEngineConfigResponseBody) SetRequestId(v string) *CloneEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

type CloneEngineConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneEngineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigResponse) SetHeaders(v map[string]*string) *CloneEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *CloneEngineConfigResponse) SetStatusCode(v int32) *CloneEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneEngineConfigResponse) SetBody(v *CloneEngineConfigResponseBody) *CloneEngineConfigResponse {
	s.Body = v
	return s
}

type CloneExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneExperimentRequest) GoString() string {
	return s.String()
}

func (s *CloneExperimentRequest) SetInstanceId(v string) *CloneExperimentRequest {
	s.InstanceId = &v
	return s
}

type CloneExperimentResponseBody struct {
	// example:
	//
	// 3
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CloneExperimentResponseBody) SetExperimentId(v string) *CloneExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CloneExperimentResponseBody) SetRequestId(v string) *CloneExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CloneExperimentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneExperimentResponse) GoString() string {
	return s.String()
}

func (s *CloneExperimentResponse) SetHeaders(v map[string]*string) *CloneExperimentResponse {
	s.Headers = v
	return s
}

func (s *CloneExperimentResponse) SetStatusCode(v int32) *CloneExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneExperimentResponse) SetBody(v *CloneExperimentResponseBody) *CloneExperimentResponse {
	s.Body = v
	return s
}

type CloneExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
}

func (s CloneExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloneExperimentGroupRequest) SetEnvironment(v string) *CloneExperimentGroupRequest {
	s.Environment = &v
	return s
}

func (s *CloneExperimentGroupRequest) SetInstanceId(v string) *CloneExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneExperimentGroupRequest) SetLayerId(v string) *CloneExperimentGroupRequest {
	s.LayerId = &v
	return s
}

type CloneExperimentGroupResponseBody struct {
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 12A65C6C-AFA1-59B2-9A66-A9E0BB73F0E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloneExperimentGroupResponseBody) SetExperimentGroupId(v string) *CloneExperimentGroupResponseBody {
	s.ExperimentGroupId = &v
	return s
}

func (s *CloneExperimentGroupResponseBody) SetRequestId(v string) *CloneExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

type CloneExperimentGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloneExperimentGroupResponse) SetHeaders(v map[string]*string) *CloneExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloneExperimentGroupResponse) SetStatusCode(v int32) *CloneExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneExperimentGroupResponse) SetBody(v *CloneExperimentGroupResponseBody) *CloneExperimentGroupResponse {
	s.Body = v
	return s
}

type CloneFeatureConsistencyCheckJobConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneFeatureConsistencyCheckJobConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneFeatureConsistencyCheckJobConfigRequest) SetInstanceId(v string) *CloneFeatureConsistencyCheckJobConfigRequest {
	s.InstanceId = &v
	return s
}

type CloneFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// 4
	FeatureConsistencyCheckId *string `json:"FeatureConsistencyCheckId,omitempty" xml:"FeatureConsistencyCheckId,omitempty"`
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) SetFeatureConsistencyCheckId(v string) *CloneFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureConsistencyCheckId = &v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *CloneFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

type CloneFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneFeatureConsistencyCheckJobConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *CloneFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *CloneFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) SetBody(v *CloneFeatureConsistencyCheckJobConfigResponseBody) *CloneFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

type CloneLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	CloneExperimentGroup *bool `json:"CloneExperimentGroup,omitempty" xml:"CloneExperimentGroup,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *CloneLaboratoryRequest) SetCloneExperimentGroup(v bool) *CloneLaboratoryRequest {
	s.CloneExperimentGroup = &v
	return s
}

func (s *CloneLaboratoryRequest) SetEnvironment(v string) *CloneLaboratoryRequest {
	s.Environment = &v
	return s
}

func (s *CloneLaboratoryRequest) SetInstanceId(v string) *CloneLaboratoryRequest {
	s.InstanceId = &v
	return s
}

type CloneLaboratoryResponseBody struct {
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *CloneLaboratoryResponseBody) SetLaboratoryId(v string) *CloneLaboratoryResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *CloneLaboratoryResponseBody) SetRequestId(v string) *CloneLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

type CloneLaboratoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *CloneLaboratoryResponse) SetHeaders(v map[string]*string) *CloneLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *CloneLaboratoryResponse) SetStatusCode(v int32) *CloneLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneLaboratoryResponse) SetBody(v *CloneLaboratoryResponseBody) *CloneLaboratoryResponse {
	s.Body = v
	return s
}

type CloneTrafficControlTaskRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *CloneTrafficControlTaskRequest) SetInstanceId(v string) *CloneTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

type CloneTrafficControlTaskResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
}

func (s CloneTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloneTrafficControlTaskResponseBody) SetRequestId(v string) *CloneTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneTrafficControlTaskResponseBody) SetTrafficControlTaskId(v string) *CloneTrafficControlTaskResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

type CloneTrafficControlTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *CloneTrafficControlTaskResponse) SetHeaders(v map[string]*string) *CloneTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *CloneTrafficControlTaskResponse) SetStatusCode(v int32) *CloneTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneTrafficControlTaskResponse) SetBody(v *CloneTrafficControlTaskResponseBody) *CloneTrafficControlTaskResponse {
	s.Body = v
	return s
}

type CreateABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sum(click_cnt)
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 3
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateABMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateABMetricRequest) GoString() string {
	return s.String()
}

func (s *CreateABMetricRequest) SetDefinition(v string) *CreateABMetricRequest {
	s.Definition = &v
	return s
}

func (s *CreateABMetricRequest) SetDescription(v string) *CreateABMetricRequest {
	s.Description = &v
	return s
}

func (s *CreateABMetricRequest) SetInstanceId(v string) *CreateABMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateABMetricRequest) SetLeftMetricId(v string) *CreateABMetricRequest {
	s.LeftMetricId = &v
	return s
}

func (s *CreateABMetricRequest) SetName(v string) *CreateABMetricRequest {
	s.Name = &v
	return s
}

func (s *CreateABMetricRequest) SetOperator(v string) *CreateABMetricRequest {
	s.Operator = &v
	return s
}

func (s *CreateABMetricRequest) SetRealtime(v bool) *CreateABMetricRequest {
	s.Realtime = &v
	return s
}

func (s *CreateABMetricRequest) SetResultResourceId(v string) *CreateABMetricRequest {
	s.ResultResourceId = &v
	return s
}

func (s *CreateABMetricRequest) SetRightMetricId(v string) *CreateABMetricRequest {
	s.RightMetricId = &v
	return s
}

func (s *CreateABMetricRequest) SetSceneId(v string) *CreateABMetricRequest {
	s.SceneId = &v
	return s
}

func (s *CreateABMetricRequest) SetStatisticsCycle(v int32) *CreateABMetricRequest {
	s.StatisticsCycle = &v
	return s
}

func (s *CreateABMetricRequest) SetTableMetaId(v string) *CreateABMetricRequest {
	s.TableMetaId = &v
	return s
}

func (s *CreateABMetricRequest) SetType(v string) *CreateABMetricRequest {
	s.Type = &v
	return s
}

type CreateABMetricResponseBody struct {
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateABMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABMetricResponseBody) SetABMetricId(v string) *CreateABMetricResponseBody {
	s.ABMetricId = &v
	return s
}

func (s *CreateABMetricResponseBody) SetRequestId(v string) *CreateABMetricResponseBody {
	s.RequestId = &v
	return s
}

type CreateABMetricResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABMetricResponse) GoString() string {
	return s.String()
}

func (s *CreateABMetricResponse) SetHeaders(v map[string]*string) *CreateABMetricResponse {
	s.Headers = v
	return s
}

func (s *CreateABMetricResponse) SetStatusCode(v int32) *CreateABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABMetricResponse) SetBody(v *CreateABMetricResponseBody) *CreateABMetricResponse {
	s.Body = v
	return s
}

type CreateABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CreateABMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateABMetricGroupRequest) SetABMetricIds(v string) *CreateABMetricGroupRequest {
	s.ABMetricIds = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetDescription(v string) *CreateABMetricGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetInstanceId(v string) *CreateABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetName(v string) *CreateABMetricGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetRealtime(v bool) *CreateABMetricGroupRequest {
	s.Realtime = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetSceneId(v string) *CreateABMetricGroupRequest {
	s.SceneId = &v
	return s
}

type CreateABMetricGroupResponseBody struct {
	// example:
	//
	// 1
	ABMetricGroupId *string `json:"ABMetricGroupId,omitempty" xml:"ABMetricGroupId,omitempty"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateABMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABMetricGroupResponseBody) SetABMetricGroupId(v string) *CreateABMetricGroupResponseBody {
	s.ABMetricGroupId = &v
	return s
}

func (s *CreateABMetricGroupResponseBody) SetRequestId(v string) *CreateABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateABMetricGroupResponse) SetHeaders(v map[string]*string) *CreateABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateABMetricGroupResponse) SetStatusCode(v int32) *CreateABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABMetricGroupResponse) SetBody(v *CreateABMetricGroupResponseBody) *CreateABMetricGroupResponse {
	s.Body = v
	return s
}

type CreateCalculationJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2,3,4
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-03
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CreateCalculationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCalculationJobsRequest) GoString() string {
	return s.String()
}

func (s *CreateCalculationJobsRequest) SetABMetricIds(v string) *CreateCalculationJobsRequest {
	s.ABMetricIds = &v
	return s
}

func (s *CreateCalculationJobsRequest) SetEndDate(v string) *CreateCalculationJobsRequest {
	s.EndDate = &v
	return s
}

func (s *CreateCalculationJobsRequest) SetInstanceId(v string) *CreateCalculationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCalculationJobsRequest) SetStartDate(v string) *CreateCalculationJobsRequest {
	s.StartDate = &v
	return s
}

type CreateCalculationJobsResponseBody struct {
	CalculationJobIds []*string `json:"CalculationJobIds,omitempty" xml:"CalculationJobIds,omitempty" type:"Repeated"`
	// example:
	//
	// 8C27790E-CCA5-56BB-BA17-646295DEC0A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCalculationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCalculationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCalculationJobsResponseBody) SetCalculationJobIds(v []*string) *CreateCalculationJobsResponseBody {
	s.CalculationJobIds = v
	return s
}

func (s *CreateCalculationJobsResponseBody) SetRequestId(v string) *CreateCalculationJobsResponseBody {
	s.RequestId = &v
	return s
}

type CreateCalculationJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCalculationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCalculationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCalculationJobsResponse) GoString() string {
	return s.String()
}

func (s *CreateCalculationJobsResponse) SetHeaders(v map[string]*string) *CreateCalculationJobsResponse {
	s.Headers = v
	return s
}

func (s *CreateCalculationJobsResponse) SetStatusCode(v int32) *CreateCalculationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCalculationJobsResponse) SetBody(v *CreateCalculationJobsResponseBody) *CreateCalculationJobsResponse {
	s.Body = v
	return s
}

type CreateCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// os=android
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx人群
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// user1,user2,user3
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s CreateCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCrowdRequest) GoString() string {
	return s.String()
}

func (s *CreateCrowdRequest) SetDescription(v string) *CreateCrowdRequest {
	s.Description = &v
	return s
}

func (s *CreateCrowdRequest) SetInstanceId(v string) *CreateCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCrowdRequest) SetLabel(v string) *CreateCrowdRequest {
	s.Label = &v
	return s
}

func (s *CreateCrowdRequest) SetName(v string) *CreateCrowdRequest {
	s.Name = &v
	return s
}

func (s *CreateCrowdRequest) SetSource(v string) *CreateCrowdRequest {
	s.Source = &v
	return s
}

func (s *CreateCrowdRequest) SetUsers(v string) *CreateCrowdRequest {
	s.Users = &v
	return s
}

type CreateCrowdResponseBody struct {
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCrowdResponseBody) SetCrowdId(v string) *CreateCrowdResponseBody {
	s.CrowdId = &v
	return s
}

func (s *CreateCrowdResponseBody) SetRequestId(v string) *CreateCrowdResponseBody {
	s.RequestId = &v
	return s
}

type CreateCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCrowdResponse) GoString() string {
	return s.String()
}

func (s *CreateCrowdResponse) SetHeaders(v map[string]*string) *CreateCrowdResponse {
	s.Headers = v
	return s
}

func (s *CreateCrowdResponse) SetStatusCode(v int32) *CreateCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCrowdResponse) SetBody(v *CreateCrowdResponseBody) *CreateCrowdResponse {
	s.Body = v
	return s
}

type CreateEngineConfigRequest struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateEngineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateEngineConfigRequest) SetConfigValue(v string) *CreateEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *CreateEngineConfigRequest) SetDescription(v string) *CreateEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateEngineConfigRequest) SetEnvironment(v string) *CreateEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *CreateEngineConfigRequest) SetInstanceId(v string) *CreateEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEngineConfigRequest) SetName(v string) *CreateEngineConfigRequest {
	s.Name = &v
	return s
}

type CreateEngineConfigResponseBody struct {
	// example:
	//
	// 1
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEngineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEngineConfigResponseBody) SetEngineConfigId(v string) *CreateEngineConfigResponseBody {
	s.EngineConfigId = &v
	return s
}

func (s *CreateEngineConfigResponseBody) SetRequestId(v string) *CreateEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateEngineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEngineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateEngineConfigResponse) SetHeaders(v map[string]*string) *CreateEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateEngineConfigResponse) SetStatusCode(v int32) *CreateEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEngineConfigResponse) SetBody(v *CreateEngineConfigResponseBody) *CreateEngineConfigResponse {
	s.Body = v
	return s
}

type CreateExperimentRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// experiment_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) SetConfig(v string) *CreateExperimentRequest {
	s.Config = &v
	return s
}

func (s *CreateExperimentRequest) SetDebugCrowdId(v string) *CreateExperimentRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *CreateExperimentRequest) SetDebugUsers(v string) *CreateExperimentRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateExperimentRequest) SetDescription(v string) *CreateExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentRequest) SetExperimentGroupId(v string) *CreateExperimentRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *CreateExperimentRequest) SetFlowPercent(v int32) *CreateExperimentRequest {
	s.FlowPercent = &v
	return s
}

func (s *CreateExperimentRequest) SetInstanceId(v string) *CreateExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetType(v string) *CreateExperimentRequest {
	s.Type = &v
	return s
}

type CreateExperimentResponseBody struct {
	// example:
	//
	// 3
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AAA45F6-0798-5461-9360-81D133823CE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponseBody) SetExperimentId(v string) *CreateExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateExperimentResponseBody) SetRequestId(v string) *CreateExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponse) SetHeaders(v map[string]*string) *CreateExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentResponse) SetStatusCode(v int32) *CreateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentResponse) SetBody(v *CreateExperimentResponseBody) *CreateExperimentResponse {
	s.Body = v
	return s
}

type CreateExperimentGroupRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 1
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// example:
	//
	// gender=male
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// experiment_group_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedAA     *bool  `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	RandomFlow *int64 `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// example:
	//
	// 1,2,3
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
}

func (s CreateExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentGroupRequest) SetConfig(v string) *CreateExperimentGroupRequest {
	s.Config = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetCrowdId(v string) *CreateExperimentGroupRequest {
	s.CrowdId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetCrowdTargetType(v string) *CreateExperimentGroupRequest {
	s.CrowdTargetType = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDebugCrowdId(v string) *CreateExperimentGroupRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDebugUsers(v string) *CreateExperimentGroupRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDescription(v string) *CreateExperimentGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDistributionTimeDuration(v int32) *CreateExperimentGroupRequest {
	s.DistributionTimeDuration = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDistributionType(v string) *CreateExperimentGroupRequest {
	s.DistributionType = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetFilter(v string) *CreateExperimentGroupRequest {
	s.Filter = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetInstanceId(v string) *CreateExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetLayerId(v string) *CreateExperimentGroupRequest {
	s.LayerId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetName(v string) *CreateExperimentGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetNeedAA(v bool) *CreateExperimentGroupRequest {
	s.NeedAA = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetRandomFlow(v int64) *CreateExperimentGroupRequest {
	s.RandomFlow = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetReservedBuckets(v string) *CreateExperimentGroupRequest {
	s.ReservedBuckets = &v
	return s
}

type CreateExperimentGroupResponseBody struct {
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentGroupResponseBody) SetExperimentGroupId(v string) *CreateExperimentGroupResponseBody {
	s.ExperimentGroupId = &v
	return s
}

func (s *CreateExperimentGroupResponseBody) SetRequestId(v string) *CreateExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentGroupResponse) SetHeaders(v map[string]*string) *CreateExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentGroupResponse) SetStatusCode(v int32) *CreateExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentGroupResponse) SetBody(v *CreateExperimentGroupResponseBody) *CreateExperimentGroupResponse {
	s.Body = v
	return s
}

type CreateFeatureConsistencyCheckJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	SamplingDuration *int32 `json:"SamplingDuration,omitempty" xml:"SamplingDuration,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetEnvironment(v string) *CreateFeatureConsistencyCheckJobRequest {
	s.Environment = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetFeatureConsistencyCheckJobConfigId(v string) *CreateFeatureConsistencyCheckJobRequest {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetInstanceId(v string) *CreateFeatureConsistencyCheckJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetSamplingDuration(v int32) *CreateFeatureConsistencyCheckJobRequest {
	s.SamplingDuration = &v
	return s
}

type CreateFeatureConsistencyCheckJobResponseBody struct {
	// example:
	//
	// 4
	FeatureConsistencyCheckJobId *string `json:"FeatureConsistencyCheckJobId,omitempty" xml:"FeatureConsistencyCheckJobId,omitempty"`
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) SetFeatureConsistencyCheckJobId(v string) *CreateFeatureConsistencyCheckJobResponseBody {
	s.FeatureConsistencyCheckJobId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) SetRequestId(v string) *CreateFeatureConsistencyCheckJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateFeatureConsistencyCheckJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureConsistencyCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobResponse) SetHeaders(v map[string]*string) *CreateFeatureConsistencyCheckJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponse) SetStatusCode(v int32) *CreateFeatureConsistencyCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponse) SetBody(v *CreateFeatureConsistencyCheckJobResponseBody) *CreateFeatureConsistencyCheckJobResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
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

type CreateFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// 4
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) SetFeatureConsistencyCheckJobConfigId(v string) *CreateFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *CreateFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *CreateFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *CreateFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponse) SetBody(v *CreateFeatureConsistencyCheckJobConfigResponseBody) *CreateFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

type CreateInstanceResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateInstanceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceResourceRequest) SetCategory(v string) *CreateInstanceResourceRequest {
	s.Category = &v
	return s
}

func (s *CreateInstanceResourceRequest) SetGroup(v string) *CreateInstanceResourceRequest {
	s.Group = &v
	return s
}

func (s *CreateInstanceResourceRequest) SetType(v string) *CreateInstanceResourceRequest {
	s.Type = &v
	return s
}

func (s *CreateInstanceResourceRequest) SetUri(v string) *CreateInstanceResourceRequest {
	s.Uri = &v
	return s
}

type CreateInstanceResourceResponseBody struct {
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CreateInstanceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResourceResponseBody) SetRequestId(v string) *CreateInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResourceResponseBody) SetResourceId(v string) *CreateInstanceResourceResponseBody {
	s.ResourceId = &v
	return s
}

type CreateInstanceResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResourceResponse) SetHeaders(v map[string]*string) *CreateInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResourceResponse) SetStatusCode(v int32) *CreateInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResourceResponse) SetBody(v *CreateInstanceResourceResponseBody) *CreateInstanceResourceResponse {
	s.Body = v
	return s
}

type CreateLaboratoryRequest struct {
	// example:
	//
	// 24
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UidHash
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// example:
	//
	// 1,2,3,10-20
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// filter=xxx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *CreateLaboratoryRequest) SetBucketCount(v int32) *CreateLaboratoryRequest {
	s.BucketCount = &v
	return s
}

func (s *CreateLaboratoryRequest) SetBucketType(v string) *CreateLaboratoryRequest {
	s.BucketType = &v
	return s
}

func (s *CreateLaboratoryRequest) SetBuckets(v string) *CreateLaboratoryRequest {
	s.Buckets = &v
	return s
}

func (s *CreateLaboratoryRequest) SetDebugCrowdId(v string) *CreateLaboratoryRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *CreateLaboratoryRequest) SetDebugUsers(v string) *CreateLaboratoryRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateLaboratoryRequest) SetDescription(v string) *CreateLaboratoryRequest {
	s.Description = &v
	return s
}

func (s *CreateLaboratoryRequest) SetEnvironment(v string) *CreateLaboratoryRequest {
	s.Environment = &v
	return s
}

func (s *CreateLaboratoryRequest) SetFilter(v string) *CreateLaboratoryRequest {
	s.Filter = &v
	return s
}

func (s *CreateLaboratoryRequest) SetInstanceId(v string) *CreateLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLaboratoryRequest) SetName(v string) *CreateLaboratoryRequest {
	s.Name = &v
	return s
}

func (s *CreateLaboratoryRequest) SetSceneId(v string) *CreateLaboratoryRequest {
	s.SceneId = &v
	return s
}

func (s *CreateLaboratoryRequest) SetType(v string) *CreateLaboratoryRequest {
	s.Type = &v
	return s
}

type CreateLaboratoryResponseBody struct {
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 42391E6D-822C-58F8-9F7E-D991BB86D6AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLaboratoryResponseBody) SetLaboratoryId(v string) *CreateLaboratoryResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *CreateLaboratoryResponseBody) SetRequestId(v string) *CreateLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *CreateLaboratoryResponse) SetHeaders(v map[string]*string) *CreateLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *CreateLaboratoryResponse) SetStatusCode(v int32) *CreateLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLaboratoryResponse) SetBody(v *CreateLaboratoryResponseBody) *CreateLaboratoryResponse {
	s.Body = v
	return s
}

type CreateLayerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// layer1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerRequest) GoString() string {
	return s.String()
}

func (s *CreateLayerRequest) SetDescription(v string) *CreateLayerRequest {
	s.Description = &v
	return s
}

func (s *CreateLayerRequest) SetInstanceId(v string) *CreateLayerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLayerRequest) SetLaboratoryId(v string) *CreateLayerRequest {
	s.LaboratoryId = &v
	return s
}

func (s *CreateLayerRequest) SetName(v string) *CreateLayerRequest {
	s.Name = &v
	return s
}

type CreateLayerResponseBody struct {
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayerResponseBody) SetLayerId(v string) *CreateLayerResponseBody {
	s.LayerId = &v
	return s
}

func (s *CreateLayerResponseBody) SetRequestId(v string) *CreateLayerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerResponse) GoString() string {
	return s.String()
}

func (s *CreateLayerResponse) SetHeaders(v map[string]*string) *CreateLayerResponse {
	s.Headers = v
	return s
}

func (s *CreateLayerResponse) SetStatusCode(v int32) *CreateLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayerResponse) SetBody(v *CreateLayerResponseBody) *CreateLayerResponse {
	s.Body = v
	return s
}

type CreateParamRequest struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// house
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateParamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParamRequest) GoString() string {
	return s.String()
}

func (s *CreateParamRequest) SetEnvironment(v string) *CreateParamRequest {
	s.Environment = &v
	return s
}

func (s *CreateParamRequest) SetInstanceId(v string) *CreateParamRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateParamRequest) SetName(v string) *CreateParamRequest {
	s.Name = &v
	return s
}

func (s *CreateParamRequest) SetSceneId(v string) *CreateParamRequest {
	s.SceneId = &v
	return s
}

func (s *CreateParamRequest) SetValue(v string) *CreateParamRequest {
	s.Value = &v
	return s
}

type CreateParamResponseBody struct {
	// example:
	//
	// 4
	ParamId *int64 `json:"ParamId,omitempty" xml:"ParamId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateParamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParamResponseBody) SetParamId(v int64) *CreateParamResponseBody {
	s.ParamId = &v
	return s
}

func (s *CreateParamResponseBody) SetRequestId(v string) *CreateParamResponseBody {
	s.RequestId = &v
	return s
}

type CreateParamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateParamResponse) GoString() string {
	return s.String()
}

func (s *CreateParamResponse) SetHeaders(v map[string]*string) *CreateParamResponse {
	s.Headers = v
	return s
}

func (s *CreateParamResponse) SetStatusCode(v int32) *CreateParamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParamResponse) SetBody(v *CreateParamResponseBody) *CreateParamResponse {
	s.Body = v
	return s
}

type CreateResourceRuleRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricOperationType *string `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo      *string `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod    *string `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RuleComputingDefinition *string `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	// This parameter is required.
	RuleItems []*CreateResourceRuleRequestRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s CreateResourceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleRequest) SetDescription(v string) *CreateResourceRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceRuleRequest) SetInstanceId(v string) *CreateResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceRuleRequest) SetMetricOperationType(v string) *CreateResourceRuleRequest {
	s.MetricOperationType = &v
	return s
}

func (s *CreateResourceRuleRequest) SetMetricPullInfo(v string) *CreateResourceRuleRequest {
	s.MetricPullInfo = &v
	return s
}

func (s *CreateResourceRuleRequest) SetMetricPullPeriod(v string) *CreateResourceRuleRequest {
	s.MetricPullPeriod = &v
	return s
}

func (s *CreateResourceRuleRequest) SetName(v string) *CreateResourceRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceRuleRequest) SetRuleComputingDefinition(v string) *CreateResourceRuleRequest {
	s.RuleComputingDefinition = &v
	return s
}

func (s *CreateResourceRuleRequest) SetRuleItems(v []*CreateResourceRuleRequestRuleItems) *CreateResourceRuleRequest {
	s.RuleItems = v
	return s
}

type CreateResourceRuleRequestRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	MaxValue *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// This parameter is required.
	MinValue *float64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceRuleRequestRuleItems) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleRequestRuleItems) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleRequestRuleItems) SetDescription(v string) *CreateResourceRuleRequestRuleItems {
	s.Description = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetMaxValue(v float64) *CreateResourceRuleRequestRuleItems {
	s.MaxValue = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetMinValue(v float64) *CreateResourceRuleRequestRuleItems {
	s.MinValue = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetName(v string) *CreateResourceRuleRequestRuleItems {
	s.Name = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetValue(v float64) *CreateResourceRuleRequestRuleItems {
	s.Value = &v
	return s
}

type CreateResourceRuleResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleId *string `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
}

func (s CreateResourceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleResponseBody) SetRequestId(v string) *CreateResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceRuleResponseBody) SetResourceRuleId(v string) *CreateResourceRuleResponseBody {
	s.ResourceRuleId = &v
	return s
}

type CreateResourceRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleResponse) SetHeaders(v map[string]*string) *CreateResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceRuleResponse) SetStatusCode(v int32) *CreateResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceRuleResponse) SetBody(v *CreateResourceRuleResponseBody) *CreateResourceRuleResponse {
	s.Body = v
	return s
}

type CreateResourceRuleItemRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MaxValue *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// This parameter is required.
	MinValue *float64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceRuleItemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleItemRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleItemRequest) SetDescription(v string) *CreateResourceRuleItemRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetInstanceId(v string) *CreateResourceRuleItemRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetMaxValue(v float64) *CreateResourceRuleItemRequest {
	s.MaxValue = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetMinValue(v float64) *CreateResourceRuleItemRequest {
	s.MinValue = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetName(v string) *CreateResourceRuleItemRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetValue(v float64) *CreateResourceRuleItemRequest {
	s.Value = &v
	return s
}

type CreateResourceRuleItemResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleItemId *string `json:"ResourceRuleItemId,omitempty" xml:"ResourceRuleItemId,omitempty"`
}

func (s CreateResourceRuleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleItemResponseBody) SetRequestId(v string) *CreateResourceRuleItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceRuleItemResponseBody) SetResourceRuleItemId(v string) *CreateResourceRuleItemResponseBody {
	s.ResourceRuleItemId = &v
	return s
}

type CreateResourceRuleItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceRuleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceRuleItemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRuleItemResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleItemResponse) SetHeaders(v map[string]*string) *CreateResourceRuleItemResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceRuleItemResponse) SetStatusCode(v int32) *CreateResourceRuleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceRuleItemResponse) SetBody(v *CreateResourceRuleItemResponseBody) *CreateResourceRuleItemResponse {
	s.Body = v
	return s
}

type CreateSceneRequest struct {
	// example:
	//
	// This is a test.
	Description *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*CreateSceneRequestFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneRequest) SetDescription(v string) *CreateSceneRequest {
	s.Description = &v
	return s
}

func (s *CreateSceneRequest) SetFlows(v []*CreateSceneRequestFlows) *CreateSceneRequest {
	s.Flows = v
	return s
}

func (s *CreateSceneRequest) SetInstanceId(v string) *CreateSceneRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSceneRequest) SetName(v string) *CreateSceneRequest {
	s.Name = &v
	return s
}

type CreateSceneRequestFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s CreateSceneRequestFlows) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneRequestFlows) GoString() string {
	return s.String()
}

func (s *CreateSceneRequestFlows) SetFlowCode(v string) *CreateSceneRequestFlows {
	s.FlowCode = &v
	return s
}

func (s *CreateSceneRequestFlows) SetFlowName(v string) *CreateSceneRequestFlows {
	s.FlowName = &v
	return s
}

type CreateSceneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FCF741D8-9C30-578E-807F-B935487DB34A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CreateSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneResponseBody) SetRequestId(v string) *CreateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneResponseBody) SetSceneId(v string) *CreateSceneResponseBody {
	s.SceneId = &v
	return s
}

type CreateSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneResponse) SetHeaders(v map[string]*string) *CreateSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneResponse) SetStatusCode(v int32) *CreateSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSceneResponse) SetBody(v *CreateSceneResponseBody) *CreateSceneResponse {
	s.Body = v
	return s
}

type CreateSubCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user1,user2,user3
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s CreateSubCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubCrowdRequest) GoString() string {
	return s.String()
}

func (s *CreateSubCrowdRequest) SetInstanceId(v string) *CreateSubCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSubCrowdRequest) SetSource(v string) *CreateSubCrowdRequest {
	s.Source = &v
	return s
}

func (s *CreateSubCrowdRequest) SetUsers(v string) *CreateSubCrowdRequest {
	s.Users = &v
	return s
}

type CreateSubCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SubCrowdId *string `json:"SubCrowdId,omitempty" xml:"SubCrowdId,omitempty"`
}

func (s CreateSubCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubCrowdResponseBody) SetRequestId(v string) *CreateSubCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubCrowdResponseBody) SetSubCrowdId(v string) *CreateSubCrowdResponseBody {
	s.SubCrowdId = &v
	return s
}

type CreateSubCrowdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubCrowdResponse) GoString() string {
	return s.String()
}

func (s *CreateSubCrowdResponse) SetHeaders(v map[string]*string) *CreateSubCrowdResponse {
	s.Headers = v
	return s
}

func (s *CreateSubCrowdResponse) SetStatusCode(v int32) *CreateSubCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubCrowdResponse) SetBody(v *CreateSubCrowdResponseBody) *CreateSubCrowdResponse {
	s.Body = v
	return s
}

type CreateTableMetaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// this is a test table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Fields []*CreateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// reso-2s416t146ffjc3yefx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table_mysql
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreateTableMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTableMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateTableMetaRequest) SetDescription(v string) *CreateTableMetaRequest {
	s.Description = &v
	return s
}

func (s *CreateTableMetaRequest) SetFields(v []*CreateTableMetaRequestFields) *CreateTableMetaRequest {
	s.Fields = v
	return s
}

func (s *CreateTableMetaRequest) SetInstanceId(v string) *CreateTableMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTableMetaRequest) SetModule(v string) *CreateTableMetaRequest {
	s.Module = &v
	return s
}

func (s *CreateTableMetaRequest) SetName(v string) *CreateTableMetaRequest {
	s.Name = &v
	return s
}

func (s *CreateTableMetaRequest) SetResourceId(v string) *CreateTableMetaRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateTableMetaRequest) SetTableName(v string) *CreateTableMetaRequest {
	s.TableName = &v
	return s
}

type CreateTableMetaRequestFields struct {
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	// This parameter is required.
	IsPartitionField *string `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// this is gender of people
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gender
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTableMetaRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateTableMetaRequestFields) GoString() string {
	return s.String()
}

func (s *CreateTableMetaRequestFields) SetDataType(v string) *CreateTableMetaRequestFields {
	s.DataType = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetIsDimensionField(v bool) *CreateTableMetaRequestFields {
	s.IsDimensionField = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetIsPartitionField(v string) *CreateTableMetaRequestFields {
	s.IsPartitionField = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetMeaning(v string) *CreateTableMetaRequestFields {
	s.Meaning = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetName(v string) *CreateTableMetaRequestFields {
	s.Name = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetType(v string) *CreateTableMetaRequestFields {
	s.Type = &v
	return s
}

type CreateTableMetaResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
}

func (s CreateTableMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableMetaResponseBody) SetRequestId(v string) *CreateTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableMetaResponseBody) SetTableMetaId(v string) *CreateTableMetaResponseBody {
	s.TableMetaId = &v
	return s
}

type CreateTableMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTableMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTableMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateTableMetaResponse) SetHeaders(v map[string]*string) *CreateTableMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateTableMetaResponse) SetStatusCode(v int32) *CreateTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableMetaResponse) SetBody(v *CreateTableMetaResponseBody) *CreateTableMetaResponse {
	s.Body = v
	return s
}

type CreateTrafficControlTargetRequest struct {
	EndTime              *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ItemConditionArray   *string  `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress *string  `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType    *string  `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                 *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation *bool    `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName           *string  `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	StartTime            *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod         *string  `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status               *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue       *int64   `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	TrafficControlTaskId *string  `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	Value                *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTargetRequest) SetEndTime(v string) *CreateTrafficControlTargetRequest {
	s.EndTime = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetEvent(v string) *CreateTrafficControlTargetRequest {
	s.Event = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetItemConditionArray(v string) *CreateTrafficControlTargetRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetItemConditionExpress(v string) *CreateTrafficControlTargetRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetItemConditionType(v string) *CreateTrafficControlTargetRequest {
	s.ItemConditionType = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetName(v string) *CreateTrafficControlTargetRequest {
	s.Name = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetNewProductRegulation(v bool) *CreateTrafficControlTargetRequest {
	s.NewProductRegulation = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetRecallName(v string) *CreateTrafficControlTargetRequest {
	s.RecallName = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetStartTime(v string) *CreateTrafficControlTargetRequest {
	s.StartTime = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetStatisPeriod(v string) *CreateTrafficControlTargetRequest {
	s.StatisPeriod = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetStatus(v string) *CreateTrafficControlTargetRequest {
	s.Status = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetToleranceValue(v int64) *CreateTrafficControlTargetRequest {
	s.ToleranceValue = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetTrafficControlTaskId(v string) *CreateTrafficControlTargetRequest {
	s.TrafficControlTaskId = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetValue(v float32) *CreateTrafficControlTargetRequest {
	s.Value = &v
	return s
}

type CreateTrafficControlTargetResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTargetId *string `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
}

func (s CreateTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTargetResponseBody) SetRequestId(v string) *CreateTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficControlTargetResponseBody) SetTrafficControlTargetId(v string) *CreateTrafficControlTargetResponseBody {
	s.TrafficControlTargetId = &v
	return s
}

type CreateTrafficControlTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTargetResponse) SetHeaders(v map[string]*string) *CreateTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficControlTargetResponse) SetStatusCode(v int32) *CreateTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficControlTargetResponse) SetBody(v *CreateTrafficControlTargetResponseBody) *CreateTrafficControlTargetResponse {
	s.Body = v
	return s
}

type CreateTrafficControlTaskRequest struct {
	// example:
	//
	// 1
	BehaviorTableMetaId *string `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	// example:
	//
	// Global
	ControlGranularity *string `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	// example:
	//
	// Guaranteed
	ControlLogic *string `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	// example:
	//
	// Percent
	ControlType *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	// example:
	//
	// this is a test task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-03-26
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TimeRange
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{\\"field\\":\\"status\\",\\"option\\":\\"=\\",\\"value\\":\\"1\\"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// example:
	//
	// status=1
	ItemConditionExpress *string `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	// example:
	//
	// Array
	ItemConditionType *string `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	// example:
	//
	// 3
	ItemTableMetaId *string `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	// example:
	//
	// task-1
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds  *string `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	ProdExperimentIds *string `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// [{\\"field\\":\\"click\\",\\"option\\":\\"<=\\",\\"value\\":\\"30\\"}]
	StatisBehaviorConditionArray *string `json:"StatisBehaviorConditionArray,omitempty" xml:"StatisBehaviorConditionArray,omitempty"`
	// example:
	//
	// click=30
	StatisBehaviorConditionExpress *string `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	// example:
	//
	// Array
	StatisBehaviorConditionType *string                                                 `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets       []*CreateTrafficControlTaskRequestTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	// example:
	//
	// [{\\"field\\":\\"gender\\",\\"option\\":\\"=\\",\\"value\\":\\"male\\"}]
	UserConditionArray *string `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	// example:
	//
	// age<=30&&(3<=level<=6)&&gender=male
	UserConditionExpress *string `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	// example:
	//
	// Array
	UserConditionType *string `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	// example:
	//
	// 2
	UserTableMetaId *string `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s CreateTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskRequest) SetBehaviorTableMetaId(v string) *CreateTrafficControlTaskRequest {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetControlGranularity(v string) *CreateTrafficControlTaskRequest {
	s.ControlGranularity = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetControlLogic(v string) *CreateTrafficControlTaskRequest {
	s.ControlLogic = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetControlType(v string) *CreateTrafficControlTaskRequest {
	s.ControlType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetDescription(v string) *CreateTrafficControlTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetEndTime(v string) *CreateTrafficControlTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetExecutionTime(v string) *CreateTrafficControlTaskRequest {
	s.ExecutionTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetInstanceId(v string) *CreateTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemConditionArray(v string) *CreateTrafficControlTaskRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemConditionExpress(v string) *CreateTrafficControlTaskRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemConditionType(v string) *CreateTrafficControlTaskRequest {
	s.ItemConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemTableMetaId(v string) *CreateTrafficControlTaskRequest {
	s.ItemTableMetaId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetName(v string) *CreateTrafficControlTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetPreExperimentIds(v string) *CreateTrafficControlTaskRequest {
	s.PreExperimentIds = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetProdExperimentIds(v string) *CreateTrafficControlTaskRequest {
	s.ProdExperimentIds = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetSceneId(v string) *CreateTrafficControlTaskRequest {
	s.SceneId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetServiceId(v string) *CreateTrafficControlTaskRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStartTime(v string) *CreateTrafficControlTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStatisBehaviorConditionArray(v string) *CreateTrafficControlTaskRequest {
	s.StatisBehaviorConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStatisBehaviorConditionExpress(v string) *CreateTrafficControlTaskRequest {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStatisBehaviorConditionType(v string) *CreateTrafficControlTaskRequest {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetTrafficControlTargets(v []*CreateTrafficControlTaskRequestTrafficControlTargets) *CreateTrafficControlTaskRequest {
	s.TrafficControlTargets = v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserConditionArray(v string) *CreateTrafficControlTaskRequest {
	s.UserConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserConditionExpress(v string) *CreateTrafficControlTaskRequest {
	s.UserConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserConditionType(v string) *CreateTrafficControlTaskRequest {
	s.UserConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserTableMetaId(v string) *CreateTrafficControlTaskRequest {
	s.UserTableMetaId = &v
	return s
}

type CreateTrafficControlTaskRequestTrafficControlTargets struct {
	// example:
	//
	// 2024-04-25
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// click
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// [{\\"field\\":\\"status\\",\\"option\\":\\"=\\",\\"value\\":\\"1\\"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// example:
	//
	// status=1
	ItemConditionExpress *string `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	// example:
	//
	// Array
	ItemConditionType *string `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	// example:
	//
	// target_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	NewProductRegulation *bool `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	// example:
	//
	// recall_1
	RecallName *string `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Daily
	StatisPeriod *string `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	// example:
	//
	// Opened
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20
	ToleranceValue *int64 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	// example:
	//
	// 100
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrafficControlTaskRequestTrafficControlTargets) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTaskRequestTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetEndTime(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetEvent(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionArray(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionExpress(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionType(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetName(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetNewProductRegulation(v bool) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetRecallName(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetStartTime(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetStatisPeriod(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetStatus(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetToleranceValue(v int64) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetValue(v float32) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Value = &v
	return s
}

type CreateTrafficControlTaskResponseBody struct {
	// example:
	//
	// 42391E6D-822C-58F8-9F7E-D991BB86D6AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
}

func (s CreateTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskResponseBody) SetRequestId(v string) *CreateTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficControlTaskResponseBody) SetTrafficControlTaskId(v string) *CreateTrafficControlTaskResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

type CreateTrafficControlTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskResponse) SetHeaders(v map[string]*string) *CreateTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficControlTaskResponse) SetStatusCode(v int32) *CreateTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficControlTaskResponse) SetBody(v *CreateTrafficControlTaskResponseBody) *CreateTrafficControlTaskResponse {
	s.Body = v
	return s
}

type DebugResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfo map[string]interface{} `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
	RegionId   *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DebugResourceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleRequest) SetInstanceId(v string) *DebugResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugResourceRuleRequest) SetMetricInfo(v map[string]interface{}) *DebugResourceRuleRequest {
	s.MetricInfo = v
	return s
}

func (s *DebugResourceRuleRequest) SetRegionId(v string) *DebugResourceRuleRequest {
	s.RegionId = &v
	return s
}

type DebugResourceRuleShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfoShrink *string `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DebugResourceRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugResourceRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleShrinkRequest) SetInstanceId(v string) *DebugResourceRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugResourceRuleShrinkRequest) SetMetricInfoShrink(v string) *DebugResourceRuleShrinkRequest {
	s.MetricInfoShrink = &v
	return s
}

func (s *DebugResourceRuleShrinkRequest) SetRegionId(v string) *DebugResourceRuleShrinkRequest {
	s.RegionId = &v
	return s
}

type DebugResourceRuleResponseBody struct {
	CurrentValues map[string]interface{} `json:"CurrentValues,omitempty" xml:"CurrentValues,omitempty"`
	OutputValues  map[string]interface{} `json:"OutputValues,omitempty" xml:"OutputValues,omitempty"`
	RequestId     *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DebugResourceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleResponseBody) SetCurrentValues(v map[string]interface{}) *DebugResourceRuleResponseBody {
	s.CurrentValues = v
	return s
}

func (s *DebugResourceRuleResponseBody) SetOutputValues(v map[string]interface{}) *DebugResourceRuleResponseBody {
	s.OutputValues = v
	return s
}

func (s *DebugResourceRuleResponseBody) SetRequestId(v string) *DebugResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

type DebugResourceRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugResourceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleResponse) SetHeaders(v map[string]*string) *DebugResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *DebugResourceRuleResponse) SetStatusCode(v int32) *DebugResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugResourceRuleResponse) SetBody(v *DebugResourceRuleResponseBody) *DebugResourceRuleResponse {
	s.Body = v
	return s
}

type DeleteABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteABMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteABMetricRequest) GoString() string {
	return s.String()
}

func (s *DeleteABMetricRequest) SetInstanceId(v string) *DeleteABMetricRequest {
	s.InstanceId = &v
	return s
}

type DeleteABMetricResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteABMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABMetricResponseBody) SetRequestId(v string) *DeleteABMetricResponseBody {
	s.RequestId = &v
	return s
}

type DeleteABMetricResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteABMetricResponse) SetHeaders(v map[string]*string) *DeleteABMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteABMetricResponse) SetStatusCode(v int32) *DeleteABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABMetricResponse) SetBody(v *DeleteABMetricResponseBody) *DeleteABMetricResponse {
	s.Body = v
	return s
}

type DeleteABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteABMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteABMetricGroupRequest) SetInstanceId(v string) *DeleteABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

type DeleteABMetricGroupResponseBody struct {
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteABMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABMetricGroupResponseBody) SetRequestId(v string) *DeleteABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteABMetricGroupResponse) SetHeaders(v map[string]*string) *DeleteABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteABMetricGroupResponse) SetStatusCode(v int32) *DeleteABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABMetricGroupResponse) SetBody(v *DeleteABMetricGroupResponseBody) *DeleteABMetricGroupResponse {
	s.Body = v
	return s
}

type DeleteCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrowdRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrowdRequest) SetInstanceId(v string) *DeleteCrowdRequest {
	s.InstanceId = &v
	return s
}

type DeleteCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrowdResponseBody) SetRequestId(v string) *DeleteCrowdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrowdResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrowdResponse) SetHeaders(v map[string]*string) *DeleteCrowdResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrowdResponse) SetStatusCode(v int32) *DeleteCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrowdResponse) SetBody(v *DeleteCrowdResponseBody) *DeleteCrowdResponse {
	s.Body = v
	return s
}

type DeleteEngineConfigRequest struct {
	// example:
	//
	// pairec-cn-***test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteEngineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteEngineConfigRequest) SetInstanceId(v string) *DeleteEngineConfigRequest {
	s.InstanceId = &v
	return s
}

type DeleteEngineConfigResponseBody struct {
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEngineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEngineConfigResponseBody) SetRequestId(v string) *DeleteEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEngineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEngineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteEngineConfigResponse) SetHeaders(v map[string]*string) *DeleteEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteEngineConfigResponse) SetStatusCode(v int32) *DeleteEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEngineConfigResponse) SetBody(v *DeleteEngineConfigResponseBody) *DeleteEngineConfigResponse {
	s.Body = v
	return s
}

type DeleteExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRequest) SetInstanceId(v string) *DeleteExperimentRequest {
	s.InstanceId = &v
	return s
}

type DeleteExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2A734D87-2212-5C84-B63A-1AC87CA843D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponseBody) SetRequestId(v string) *DeleteExperimentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponse) SetHeaders(v map[string]*string) *DeleteExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentResponse) SetStatusCode(v int32) *DeleteExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentResponse) SetBody(v *DeleteExperimentResponseBody) *DeleteExperimentResponse {
	s.Body = v
	return s
}

type DeleteExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentGroupRequest) SetInstanceId(v string) *DeleteExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

type DeleteExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A009D9BE-C85E-57B2-AE05-BD78BB6EBF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentGroupResponseBody) SetRequestId(v string) *DeleteExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentGroupResponse) SetHeaders(v map[string]*string) *DeleteExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentGroupResponse) SetStatusCode(v int32) *DeleteExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentGroupResponse) SetBody(v *DeleteExperimentGroupResponseBody) *DeleteExperimentGroupResponse {
	s.Body = v
	return s
}

type DeleteInstanceResourceResponseBody struct {
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResourceResponseBody) SetRequestId(v string) *DeleteInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResourceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResourceResponse) SetStatusCode(v int32) *DeleteInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResourceResponse) SetBody(v *DeleteInstanceResourceResponseBody) *DeleteInstanceResourceResponse {
	s.Body = v
	return s
}

type DeleteLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLaboratoryRequest) SetInstanceId(v string) *DeleteLaboratoryRequest {
	s.InstanceId = &v
	return s
}

type DeleteLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1C0898E5-9220-5443-B2D9-445FF0688215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLaboratoryResponseBody) SetRequestId(v string) *DeleteLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLaboratoryResponse) SetHeaders(v map[string]*string) *DeleteLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteLaboratoryResponse) SetStatusCode(v int32) *DeleteLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLaboratoryResponse) SetBody(v *DeleteLaboratoryResponseBody) *DeleteLaboratoryResponse {
	s.Body = v
	return s
}

type DeleteLayerRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayerRequest) SetInstanceId(v string) *DeleteLayerRequest {
	s.InstanceId = &v
	return s
}

type DeleteLayerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F457D79-C4A2-5E8C-83E4-0D089456E2AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayerResponseBody) SetRequestId(v string) *DeleteLayerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayerResponse) SetHeaders(v map[string]*string) *DeleteLayerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayerResponse) SetStatusCode(v int32) *DeleteLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayerResponse) SetBody(v *DeleteLayerResponseBody) *DeleteLayerResponse {
	s.Body = v
	return s
}

type DeleteParamRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteParamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteParamRequest) GoString() string {
	return s.String()
}

func (s *DeleteParamRequest) SetInstanceId(v string) *DeleteParamRequest {
	s.InstanceId = &v
	return s
}

type DeleteParamResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F0AB6527-093F-5C44-B3BD-42C8C210C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteParamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParamResponseBody) SetRequestId(v string) *DeleteParamResponseBody {
	s.RequestId = &v
	return s
}

type DeleteParamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteParamResponse) GoString() string {
	return s.String()
}

func (s *DeleteParamResponse) SetHeaders(v map[string]*string) *DeleteParamResponse {
	s.Headers = v
	return s
}

func (s *DeleteParamResponse) SetStatusCode(v int32) *DeleteParamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParamResponse) SetBody(v *DeleteParamResponseBody) *DeleteParamResponse {
	s.Body = v
	return s
}

type DeleteResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteResourceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleRequest) SetInstanceId(v string) *DeleteResourceRuleRequest {
	s.InstanceId = &v
	return s
}

type DeleteResourceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleResponseBody) SetRequestId(v string) *DeleteResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleResponse) SetHeaders(v map[string]*string) *DeleteResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceRuleResponse) SetStatusCode(v int32) *DeleteResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceRuleResponse) SetBody(v *DeleteResourceRuleResponseBody) *DeleteResourceRuleResponse {
	s.Body = v
	return s
}

type DeleteResourceRuleItemRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteResourceRuleItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRuleItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleItemRequest) SetInstanceId(v string) *DeleteResourceRuleItemRequest {
	s.InstanceId = &v
	return s
}

type DeleteResourceRuleItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceRuleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRuleItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleItemResponseBody) SetRequestId(v string) *DeleteResourceRuleItemResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceRuleItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceRuleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceRuleItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRuleItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleItemResponse) SetHeaders(v map[string]*string) *DeleteResourceRuleItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceRuleItemResponse) SetStatusCode(v int32) *DeleteResourceRuleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceRuleItemResponse) SetBody(v *DeleteResourceRuleItemResponseBody) *DeleteResourceRuleItemResponse {
	s.Body = v
	return s
}

type DeleteSceneRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSceneRequest) SetInstanceId(v string) *DeleteSceneRequest {
	s.InstanceId = &v
	return s
}

type DeleteSceneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D75C43DC-3D3A-5CC8-9AAC-8C77306C433B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponseBody) SetRequestId(v string) *DeleteSceneResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponse) SetHeaders(v map[string]*string) *DeleteSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneResponse) SetStatusCode(v int32) *DeleteSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSceneResponse) SetBody(v *DeleteSceneResponseBody) *DeleteSceneResponse {
	s.Body = v
	return s
}

type DeleteSubCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteSubCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubCrowdRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubCrowdRequest) SetInstanceId(v string) *DeleteSubCrowdRequest {
	s.InstanceId = &v
	return s
}

type DeleteSubCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EE97D06A-2AA0-5AD9-B6CF-8A267924D691
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubCrowdResponseBody) SetRequestId(v string) *DeleteSubCrowdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubCrowdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubCrowdResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubCrowdResponse) SetHeaders(v map[string]*string) *DeleteSubCrowdResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubCrowdResponse) SetStatusCode(v int32) *DeleteSubCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubCrowdResponse) SetBody(v *DeleteSubCrowdResponseBody) *DeleteSubCrowdResponse {
	s.Body = v
	return s
}

type DeleteTableMetaRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteTableMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableMetaRequest) SetInstanceId(v string) *DeleteTableMetaRequest {
	s.InstanceId = &v
	return s
}

type DeleteTableMetaResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTableMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableMetaResponseBody) SetRequestId(v string) *DeleteTableMetaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTableMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTableMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableMetaResponse) SetHeaders(v map[string]*string) *DeleteTableMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableMetaResponse) SetStatusCode(v int32) *DeleteTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableMetaResponse) SetBody(v *DeleteTableMetaResponseBody) *DeleteTableMetaResponse {
	s.Body = v
	return s
}

type DeleteTrafficControlTargetRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTargetRequest) SetInstanceId(v string) *DeleteTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

type DeleteTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTargetResponseBody) SetRequestId(v string) *DeleteTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrafficControlTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTargetResponse) SetHeaders(v map[string]*string) *DeleteTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficControlTargetResponse) SetStatusCode(v int32) *DeleteTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficControlTargetResponse) SetBody(v *DeleteTrafficControlTargetResponseBody) *DeleteTrafficControlTargetResponse {
	s.Body = v
	return s
}

type DeleteTrafficControlTaskRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTaskRequest) SetInstanceId(v string) *DeleteTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

type DeleteTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTaskResponseBody) SetRequestId(v string) *DeleteTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrafficControlTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTaskResponse) SetHeaders(v map[string]*string) *DeleteTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficControlTaskResponse) SetStatusCode(v int32) *DeleteTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficControlTaskResponse) SetBody(v *DeleteTrafficControlTaskResponseBody) *DeleteTrafficControlTaskResponse {
	s.Body = v
	return s
}

type GenerateTrafficControlTaskCodeRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateTrafficControlTaskCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTrafficControlTaskCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskCodeRequest) SetEnvironment(v string) *GenerateTrafficControlTaskCodeRequest {
	s.Environment = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeRequest) SetInstanceId(v string) *GenerateTrafficControlTaskCodeRequest {
	s.InstanceId = &v
	return s
}

type GenerateTrafficControlTaskCodeResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	PreNeedConfig *bool   `json:"PreNeedConfig,omitempty" xml:"PreNeedConfig,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTrafficControlTaskCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTrafficControlTaskCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskCodeResponseBody) SetCode(v string) *GenerateTrafficControlTaskCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponseBody) SetPreNeedConfig(v bool) *GenerateTrafficControlTaskCodeResponseBody {
	s.PreNeedConfig = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponseBody) SetRequestId(v string) *GenerateTrafficControlTaskCodeResponseBody {
	s.RequestId = &v
	return s
}

type GenerateTrafficControlTaskCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTrafficControlTaskCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTrafficControlTaskCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTrafficControlTaskCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskCodeResponse) SetHeaders(v map[string]*string) *GenerateTrafficControlTaskCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponse) SetStatusCode(v int32) *GenerateTrafficControlTaskCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponse) SetBody(v *GenerateTrafficControlTaskCodeResponseBody) *GenerateTrafficControlTaskCodeResponse {
	s.Body = v
	return s
}

type GenerateTrafficControlTaskConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateTrafficControlTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTrafficControlTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskConfigRequest) SetInstanceId(v string) *GenerateTrafficControlTaskConfigRequest {
	s.InstanceId = &v
	return s
}

type GenerateTrafficControlTaskConfigResponseBody struct {
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTrafficControlTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTrafficControlTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskConfigResponseBody) SetConfig(v string) *GenerateTrafficControlTaskConfigResponseBody {
	s.Config = &v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponseBody) SetRequestId(v string) *GenerateTrafficControlTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

type GenerateTrafficControlTaskConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTrafficControlTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTrafficControlTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTrafficControlTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskConfigResponse) SetHeaders(v map[string]*string) *GenerateTrafficControlTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponse) SetStatusCode(v int32) *GenerateTrafficControlTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponse) SetBody(v *GenerateTrafficControlTaskConfigResponseBody) *GenerateTrafficControlTaskConfigResponse {
	s.Body = v
	return s
}

type GetABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetABMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s GetABMetricRequest) GoString() string {
	return s.String()
}

func (s *GetABMetricRequest) SetInstanceId(v string) *GetABMetricRequest {
	s.InstanceId = &v
	return s
}

type GetABMetricResponseBody struct {
	// example:
	//
	// sum(click_cnt)
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// false
	Realtime *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 3
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	// example:
	//
	// 2
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// home_feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetABMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetABMetricResponseBody) SetDefinition(v string) *GetABMetricResponseBody {
	s.Definition = &v
	return s
}

func (s *GetABMetricResponseBody) SetDescription(v string) *GetABMetricResponseBody {
	s.Description = &v
	return s
}

func (s *GetABMetricResponseBody) SetLeftMetricId(v string) *GetABMetricResponseBody {
	s.LeftMetricId = &v
	return s
}

func (s *GetABMetricResponseBody) SetName(v string) *GetABMetricResponseBody {
	s.Name = &v
	return s
}

func (s *GetABMetricResponseBody) SetOperator(v string) *GetABMetricResponseBody {
	s.Operator = &v
	return s
}

func (s *GetABMetricResponseBody) SetRealtime(v string) *GetABMetricResponseBody {
	s.Realtime = &v
	return s
}

func (s *GetABMetricResponseBody) SetRequestId(v string) *GetABMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetABMetricResponseBody) SetResultResourceId(v string) *GetABMetricResponseBody {
	s.ResultResourceId = &v
	return s
}

func (s *GetABMetricResponseBody) SetResultTableMetaId(v string) *GetABMetricResponseBody {
	s.ResultTableMetaId = &v
	return s
}

func (s *GetABMetricResponseBody) SetRightMetricId(v string) *GetABMetricResponseBody {
	s.RightMetricId = &v
	return s
}

func (s *GetABMetricResponseBody) SetSceneId(v string) *GetABMetricResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetABMetricResponseBody) SetSceneName(v string) *GetABMetricResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetABMetricResponseBody) SetStatisticsCycle(v int32) *GetABMetricResponseBody {
	s.StatisticsCycle = &v
	return s
}

func (s *GetABMetricResponseBody) SetTableMetaId(v string) *GetABMetricResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *GetABMetricResponseBody) SetType(v string) *GetABMetricResponseBody {
	s.Type = &v
	return s
}

type GetABMetricResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetABMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s GetABMetricResponse) GoString() string {
	return s.String()
}

func (s *GetABMetricResponse) SetHeaders(v map[string]*string) *GetABMetricResponse {
	s.Headers = v
	return s
}

func (s *GetABMetricResponse) SetStatusCode(v int32) *GetABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetABMetricResponse) SetBody(v *GetABMetricResponseBody) *GetABMetricResponse {
	s.Body = v
	return s
}

type GetABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetABMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *GetABMetricGroupRequest) SetInstanceId(v string) *GetABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

type GetABMetricGroupResponseBody struct {
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// example:
	//
	// pv,uv
	ABMetricNames *string `json:"ABMetricNames,omitempty" xml:"ABMetricNames,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2799614***
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetABMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetABMetricGroupResponseBody) SetABMetricIds(v string) *GetABMetricGroupResponseBody {
	s.ABMetricIds = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetABMetricNames(v string) *GetABMetricGroupResponseBody {
	s.ABMetricNames = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetDescription(v string) *GetABMetricGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetName(v string) *GetABMetricGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetOwner(v string) *GetABMetricGroupResponseBody {
	s.Owner = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetRealtime(v bool) *GetABMetricGroupResponseBody {
	s.Realtime = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetRequestId(v string) *GetABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetSceneId(v string) *GetABMetricGroupResponseBody {
	s.SceneId = &v
	return s
}

type GetABMetricGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetABMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *GetABMetricGroupResponse) SetHeaders(v map[string]*string) *GetABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *GetABMetricGroupResponse) SetStatusCode(v int32) *GetABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetABMetricGroupResponse) SetBody(v *GetABMetricGroupResponseBody) *GetABMetricGroupResponse {
	s.Body = v
	return s
}

type GetCalculationJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCalculationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCalculationJobRequest) GoString() string {
	return s.String()
}

func (s *GetCalculationJobRequest) SetInstanceId(v string) *GetCalculationJobRequest {
	s.InstanceId = &v
	return s
}

type GetCalculationJobResponseBody struct {
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	// example:
	//
	// pv
	ABMetricName *string `json:"ABMetricName,omitempty" xml:"ABMetricName,omitempty"`
	// example:
	//
	// 2021-12-15
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtRanTime *string   `json:"GmtRanTime,omitempty" xml:"GmtRanTime,omitempty"`
	JobMessage []*string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty" type:"Repeated"`
	// example:
	//
	// CronOffline
	JobSource *string `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCalculationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCalculationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetCalculationJobResponseBody) SetABMetricId(v string) *GetCalculationJobResponseBody {
	s.ABMetricId = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetABMetricName(v string) *GetCalculationJobResponseBody {
	s.ABMetricName = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetBizDate(v string) *GetCalculationJobResponseBody {
	s.BizDate = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetConfig(v string) *GetCalculationJobResponseBody {
	s.Config = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetGmtRanTime(v string) *GetCalculationJobResponseBody {
	s.GmtRanTime = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetJobMessage(v []*string) *GetCalculationJobResponseBody {
	s.JobMessage = v
	return s
}

func (s *GetCalculationJobResponseBody) SetJobSource(v string) *GetCalculationJobResponseBody {
	s.JobSource = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetRequestId(v string) *GetCalculationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetStatus(v string) *GetCalculationJobResponseBody {
	s.Status = &v
	return s
}

type GetCalculationJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCalculationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCalculationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCalculationJobResponse) GoString() string {
	return s.String()
}

func (s *GetCalculationJobResponse) SetHeaders(v map[string]*string) *GetCalculationJobResponse {
	s.Headers = v
	return s
}

func (s *GetCalculationJobResponse) SetStatusCode(v int32) *GetCalculationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCalculationJobResponse) SetBody(v *GetCalculationJobResponseBody) *GetCalculationJobResponse {
	s.Body = v
	return s
}

type GetEngineConfigRequest struct {
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEngineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *GetEngineConfigRequest) SetInstanceId(v string) *GetEngineConfigRequest {
	s.InstanceId = &v
	return s
}

type GetEngineConfigResponseBody struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// 2024-01-03T02:28:00.000Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-08-27T12:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2024-01-03 02:28:00
	GmtReleasedTime *string `json:"GmtReleasedTime,omitempty" xml:"GmtReleasedTime,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEngineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetEngineConfigResponseBody) SetConfigValue(v string) *GetEngineConfigResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetDescription(v string) *GetEngineConfigResponseBody {
	s.Description = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetEnvironment(v string) *GetEngineConfigResponseBody {
	s.Environment = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetGmtCreateTime(v string) *GetEngineConfigResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetGmtModifiedTime(v string) *GetEngineConfigResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetGmtReleasedTime(v string) *GetEngineConfigResponseBody {
	s.GmtReleasedTime = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetName(v string) *GetEngineConfigResponseBody {
	s.Name = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetRequestId(v string) *GetEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetStatus(v string) *GetEngineConfigResponseBody {
	s.Status = &v
	return s
}

type GetEngineConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEngineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *GetEngineConfigResponse) SetHeaders(v map[string]*string) *GetEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *GetEngineConfigResponse) SetStatusCode(v int32) *GetEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEngineConfigResponse) SetBody(v *GetEngineConfigResponseBody) *GetEngineConfigResponse {
	s.Body = v
	return s
}

type GetExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentRequest) SetInstanceId(v string) *GetExperimentRequest {
	s.InstanceId = &v
	return s
}

type GetExperimentResponseBody struct {
	// example:
	//
	// L1#EG1#E1
	AliasExperimentId *string `json:"AliasExperimentId,omitempty" xml:"AliasExperimentId,omitempty"`
	// example:
	//
	// 1,2,3
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// uid1,uid2,uid3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
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
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// experiment_test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBody) SetAliasExperimentId(v string) *GetExperimentResponseBody {
	s.AliasExperimentId = &v
	return s
}

func (s *GetExperimentResponseBody) SetBuckets(v string) *GetExperimentResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetExperimentResponseBody) SetConfig(v string) *GetExperimentResponseBody {
	s.Config = &v
	return s
}

func (s *GetExperimentResponseBody) SetDebugCrowdId(v string) *GetExperimentResponseBody {
	s.DebugCrowdId = &v
	return s
}

func (s *GetExperimentResponseBody) SetDebugUsers(v string) *GetExperimentResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetExperimentResponseBody) SetDescription(v string) *GetExperimentResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentResponseBody) SetExperimentGroupId(v string) *GetExperimentResponseBody {
	s.ExperimentGroupId = &v
	return s
}

func (s *GetExperimentResponseBody) SetFlowPercent(v int32) *GetExperimentResponseBody {
	s.FlowPercent = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtCreateTime(v string) *GetExperimentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtModifiedTime(v string) *GetExperimentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetLaboratoryId(v string) *GetExperimentResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *GetExperimentResponseBody) SetLayerId(v string) *GetExperimentResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetExperimentResponseBody) SetName(v string) *GetExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentResponseBody) SetRequestId(v string) *GetExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResponseBody) SetSceneId(v string) *GetExperimentResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetExperimentResponseBody) SetStatus(v string) *GetExperimentResponseBody {
	s.Status = &v
	return s
}

func (s *GetExperimentResponseBody) SetType(v string) *GetExperimentResponseBody {
	s.Type = &v
	return s
}

type GetExperimentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentResponse) SetHeaders(v map[string]*string) *GetExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentResponse) SetStatusCode(v int32) *GetExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentResponse) SetBody(v *GetExperimentResponseBody) *GetExperimentResponse {
	s.Body = v
	return s
}

type GetExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentGroupRequest) SetInstanceId(v string) *GetExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

type GetExperimentGroupResponseBody struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 4
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 5
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// example:
	//
	// gender=female
	Filter         *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	HoldingBuckets *string `json:"HoldingBuckets,omitempty" xml:"HoldingBuckets,omitempty"`
	// example:
	//
	// 4
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NeedAA *bool `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	// example:
	//
	// 1124512470******
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RandomFlow *int64  `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1,2,3,4
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentGroupResponseBody) SetConfig(v string) *GetExperimentGroupResponseBody {
	s.Config = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetCrowdId(v string) *GetExperimentGroupResponseBody {
	s.CrowdId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetCrowdTargetType(v string) *GetExperimentGroupResponseBody {
	s.CrowdTargetType = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDebugCrowdId(v string) *GetExperimentGroupResponseBody {
	s.DebugCrowdId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDebugUsers(v string) *GetExperimentGroupResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDescription(v string) *GetExperimentGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDistributionTimeDuration(v int32) *GetExperimentGroupResponseBody {
	s.DistributionTimeDuration = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDistributionType(v string) *GetExperimentGroupResponseBody {
	s.DistributionType = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetFilter(v string) *GetExperimentGroupResponseBody {
	s.Filter = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetHoldingBuckets(v string) *GetExperimentGroupResponseBody {
	s.HoldingBuckets = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetLaboratoryId(v string) *GetExperimentGroupResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetLayerId(v string) *GetExperimentGroupResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetName(v string) *GetExperimentGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetNeedAA(v bool) *GetExperimentGroupResponseBody {
	s.NeedAA = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetOwner(v string) *GetExperimentGroupResponseBody {
	s.Owner = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetRandomFlow(v int64) *GetExperimentGroupResponseBody {
	s.RandomFlow = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetRequestId(v string) *GetExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetReservedBuckets(v string) *GetExperimentGroupResponseBody {
	s.ReservedBuckets = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetSceneId(v string) *GetExperimentGroupResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetStatus(v string) *GetExperimentGroupResponseBody {
	s.Status = &v
	return s
}

type GetExperimentGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentGroupResponse) SetHeaders(v map[string]*string) *GetExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentGroupResponse) SetStatusCode(v int32) *GetExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentGroupResponse) SetBody(v *GetExperimentGroupResponseBody) *GetExperimentGroupResponse {
	s.Body = v
	return s
}

type GetFeatureConsistencyCheckJobRequest struct {
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetFeatureConsistencyCheckJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobRequest) SetInstanceId(v string) *GetFeatureConsistencyCheckJobRequest {
	s.InstanceId = &v
	return s
}

type GetFeatureConsistencyCheckJobResponseBody struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 5
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// example:
	//
	// feature_consistency_check_1
	FeatureConsistencyCheckJobConfigName *string `json:"FeatureConsistencyCheckJobConfigName,omitempty" xml:"FeatureConsistencyCheckJobConfigName,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtStartTime *string   `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Logs         []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFeatureConsistencyCheckJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetConfig(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.Config = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetFeatureConsistencyCheckJobConfigId(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetFeatureConsistencyCheckJobConfigName(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.FeatureConsistencyCheckJobConfigName = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetGmtEndTime(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.GmtEndTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetGmtStartTime(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.GmtStartTime = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetLogs(v []*string) *GetFeatureConsistencyCheckJobResponseBody {
	s.Logs = v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetRequestId(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponseBody) SetStatus(v string) *GetFeatureConsistencyCheckJobResponseBody {
	s.Status = &v
	return s
}

type GetFeatureConsistencyCheckJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureConsistencyCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureConsistencyCheckJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobResponse) SetHeaders(v map[string]*string) *GetFeatureConsistencyCheckJobResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponse) SetStatusCode(v int32) *GetFeatureConsistencyCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponse) SetBody(v *GetFeatureConsistencyCheckJobResponseBody) *GetFeatureConsistencyCheckJobResponse {
	s.Body = v
	return s
}

type GetFeatureConsistencyCheckJobConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetFeatureConsistencyCheckJobConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobConfigRequest) SetInstanceId(v string) *GetFeatureConsistencyCheckJobConfigRequest {
	s.InstanceId = &v
	return s
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return tea.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
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

type GetFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureConsistencyCheckJobConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *GetFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *GetFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) SetBody(v *GetFeatureConsistencyCheckJobConfigResponseBody) *GetFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// example:
	//
	// Subscription
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// airec_developers_public_cn
	CommodityCode *string                        `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Config        *GetInstanceResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 2022-12-14 00:00:00.0
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// 2022-10-13 17:34:52.0
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2022-11-05 09:02:30.0
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// pairec-test1
	InstanceId    *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatingTool *GetInstanceResponseBodyOperatingTool `json:"OperatingTool,omitempty" xml:"OperatingTool,omitempty" type:"Struct"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetChargeType(v string) *GetInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *GetInstanceResponseBody) SetCommodityCode(v string) *GetInstanceResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetConfig(v *GetInstanceResponseBodyConfig) *GetInstanceResponseBody {
	s.Config = v
	return s
}

func (s *GetInstanceResponseBody) SetExpiredTime(v string) *GetInstanceResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetOperatingTool(v *GetInstanceResponseBodyOperatingTool) *GetInstanceResponseBody {
	s.OperatingTool = v
	return s
}

func (s *GetInstanceResponseBody) SetRegionId(v string) *GetInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStatus(v string) *GetInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBody) SetType(v string) *GetInstanceResponseBody {
	s.Type = &v
	return s
}

type GetInstanceResponseBodyConfig struct {
	DataManagements []*GetInstanceResponseBodyConfigDataManagements `json:"DataManagements,omitempty" xml:"DataManagements,omitempty" type:"Repeated"`
	Engines         []*GetInstanceResponseBodyConfigEngines         `json:"Engines,omitempty" xml:"Engines,omitempty" type:"Repeated"`
	Monitors        []*GetInstanceResponseBodyConfigMonitors        `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfig) SetDataManagements(v []*GetInstanceResponseBodyConfigDataManagements) *GetInstanceResponseBodyConfig {
	s.DataManagements = v
	return s
}

func (s *GetInstanceResponseBodyConfig) SetEngines(v []*GetInstanceResponseBodyConfigEngines) *GetInstanceResponseBodyConfig {
	s.Engines = v
	return s
}

func (s *GetInstanceResponseBodyConfig) SetMonitors(v []*GetInstanceResponseBodyConfigMonitors) *GetInstanceResponseBodyConfig {
	s.Monitors = v
	return s
}

type GetInstanceResponseBodyConfigDataManagements struct {
	// example:
	//
	// storage
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBodyConfigDataManagements) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyConfigDataManagements) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfigDataManagements) SetComponentCode(v string) *GetInstanceResponseBodyConfigDataManagements {
	s.ComponentCode = &v
	return s
}

func (s *GetInstanceResponseBodyConfigDataManagements) SetMeta(v map[string]interface{}) *GetInstanceResponseBodyConfigDataManagements {
	s.Meta = v
	return s
}

func (s *GetInstanceResponseBodyConfigDataManagements) SetType(v string) *GetInstanceResponseBodyConfigDataManagements {
	s.Type = &v
	return s
}

type GetInstanceResponseBodyConfigEngines struct {
	// example:
	//
	// feature
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBodyConfigEngines) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyConfigEngines) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfigEngines) SetComponentCode(v string) *GetInstanceResponseBodyConfigEngines {
	s.ComponentCode = &v
	return s
}

func (s *GetInstanceResponseBodyConfigEngines) SetMeta(v map[string]interface{}) *GetInstanceResponseBodyConfigEngines {
	s.Meta = v
	return s
}

func (s *GetInstanceResponseBodyConfigEngines) SetType(v string) *GetInstanceResponseBodyConfigEngines {
	s.Type = &v
	return s
}

type GetInstanceResponseBodyConfigMonitors struct {
	// example:
	//
	// featuresets
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// Platform
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBodyConfigMonitors) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyConfigMonitors) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyConfigMonitors) SetComponentCode(v string) *GetInstanceResponseBodyConfigMonitors {
	s.ComponentCode = &v
	return s
}

func (s *GetInstanceResponseBodyConfigMonitors) SetMeta(v map[string]interface{}) *GetInstanceResponseBodyConfigMonitors {
	s.Meta = v
	return s
}

func (s *GetInstanceResponseBodyConfigMonitors) SetType(v string) *GetInstanceResponseBodyConfigMonitors {
	s.Type = &v
	return s
}

type GetInstanceResponseBodyOperatingTool struct {
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
}

func (s GetInstanceResponseBodyOperatingTool) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyOperatingTool) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyOperatingTool) SetIsEnable(v bool) *GetInstanceResponseBodyOperatingTool {
	s.IsEnable = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetInstanceResourceResponseBody struct {
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// D75C43DC-3D3A-5CC8-9AAC-8C77306C433B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetInstanceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceResponseBody) SetCategory(v string) *GetInstanceResourceResponseBody {
	s.Category = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetConfig(v string) *GetInstanceResourceResponseBody {
	s.Config = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetGmtCreateTime(v string) *GetInstanceResourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResourceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetGroup(v string) *GetInstanceResourceResponseBody {
	s.Group = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetRequestId(v string) *GetInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetResourceId(v string) *GetInstanceResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetType(v string) *GetInstanceResourceResponseBody {
	s.Type = &v
	return s
}

func (s *GetInstanceResourceResponseBody) SetUri(v string) *GetInstanceResourceResponseBody {
	s.Uri = &v
	return s
}

type GetInstanceResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceResponse) SetHeaders(v map[string]*string) *GetInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResourceResponse) SetStatusCode(v int32) *GetInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResourceResponse) SetBody(v *GetInstanceResourceResponseBody) *GetInstanceResourceResponse {
	s.Body = v
	return s
}

type GetInstanceResourceTableResponseBody struct {
	Fields []*GetInstanceResourceTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetInstanceResourceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResourceTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableResponseBody) SetFields(v []*GetInstanceResourceTableResponseBodyFields) *GetInstanceResourceTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetInstanceResourceTableResponseBody) SetRequestId(v string) *GetInstanceResourceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResourceTableResponseBody) SetTableName(v string) *GetInstanceResourceTableResponseBody {
	s.TableName = &v
	return s
}

type GetInstanceResourceTableResponseBodyFields struct {
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	IsPartitionField *bool `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	// example:
	//
	// ""
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResourceTableResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResourceTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableResponseBodyFields) SetIsDimensionField(v bool) *GetInstanceResourceTableResponseBodyFields {
	s.IsDimensionField = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetIsPartitionField(v bool) *GetInstanceResourceTableResponseBodyFields {
	s.IsPartitionField = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetMeaning(v string) *GetInstanceResourceTableResponseBodyFields {
	s.Meaning = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetName(v string) *GetInstanceResourceTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetType(v string) *GetInstanceResourceTableResponseBodyFields {
	s.Type = &v
	return s
}

type GetInstanceResourceTableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResourceTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResourceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResourceTableResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableResponse) SetHeaders(v map[string]*string) *GetInstanceResourceTableResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResourceTableResponse) SetStatusCode(v int32) *GetInstanceResourceTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResourceTableResponse) SetBody(v *GetInstanceResourceTableResponseBody) *GetInstanceResourceTableResponse {
	s.Body = v
	return s
}

type GetLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *GetLaboratoryRequest) SetInstanceId(v string) *GetLaboratoryRequest {
	s.InstanceId = &v
	return s
}

type GetLaboratoryResponseBody struct {
	// example:
	//
	// 100
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// example:
	//
	// Filter
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// example:
	//
	// 1,2,3,10-20
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// user1,user2,user3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// filter=xxx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1C0898E5-9220-5443-B2D9-445FF0688215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLaboratoryResponseBody) SetBucketCount(v int32) *GetLaboratoryResponseBody {
	s.BucketCount = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetBucketType(v string) *GetLaboratoryResponseBody {
	s.BucketType = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetBuckets(v string) *GetLaboratoryResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetCrowdId(v string) *GetLaboratoryResponseBody {
	s.CrowdId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetDebugCrowdId(v string) *GetLaboratoryResponseBody {
	s.DebugCrowdId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetDebugUsers(v string) *GetLaboratoryResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetDescription(v string) *GetLaboratoryResponseBody {
	s.Description = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetEnvironment(v string) *GetLaboratoryResponseBody {
	s.Environment = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetFilter(v string) *GetLaboratoryResponseBody {
	s.Filter = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetName(v string) *GetLaboratoryResponseBody {
	s.Name = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetRequestId(v string) *GetLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetSceneId(v string) *GetLaboratoryResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetStatus(v string) *GetLaboratoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetLaboratoryResponseBody) SetType(v string) *GetLaboratoryResponseBody {
	s.Type = &v
	return s
}

type GetLaboratoryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *GetLaboratoryResponse) SetHeaders(v map[string]*string) *GetLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *GetLaboratoryResponse) SetStatusCode(v int32) *GetLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLaboratoryResponse) SetBody(v *GetLaboratoryResponseBody) *GetLaboratoryResponse {
	s.Body = v
	return s
}

type GetLayerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLayerRequest) GoString() string {
	return s.String()
}

func (s *GetLayerRequest) SetInstanceId(v string) *GetLayerRequest {
	s.InstanceId = &v
	return s
}

type GetLayerResponseBody struct {
	// example:
	//
	// This is a test.
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// layer1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EE97D06A-2AA0-5AD9-B6CF-8A267924D691
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResidualFlow *int64  `json:"ResidualFlow,omitempty" xml:"ResidualFlow,omitempty"`
	// example:
	//
	// 4
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLayerResponseBody) GoString() string {
	return s.String()
}

func (s *GetLayerResponseBody) SetDescription(v string) *GetLayerResponseBody {
	s.Description = &v
	return s
}

func (s *GetLayerResponseBody) SetGmtCreateTime(v string) *GetLayerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLayerResponseBody) SetLaboratoryId(v string) *GetLayerResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *GetLayerResponseBody) SetName(v string) *GetLayerResponseBody {
	s.Name = &v
	return s
}

func (s *GetLayerResponseBody) SetRequestId(v string) *GetLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLayerResponseBody) SetResidualFlow(v int64) *GetLayerResponseBody {
	s.ResidualFlow = &v
	return s
}

func (s *GetLayerResponseBody) SetSceneId(v string) *GetLayerResponseBody {
	s.SceneId = &v
	return s
}

type GetLayerResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLayerResponse) GoString() string {
	return s.String()
}

func (s *GetLayerResponse) SetHeaders(v map[string]*string) *GetLayerResponse {
	s.Headers = v
	return s
}

func (s *GetLayerResponse) SetStatusCode(v int32) *GetLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayerResponse) SetBody(v *GetLayerResponseBody) *GetLayerResponse {
	s.Body = v
	return s
}

type GetResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetResourceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRuleRequest) SetInstanceId(v string) *GetResourceRuleRequest {
	s.InstanceId = &v
	return s
}

type GetResourceRuleResponseBody struct {
	Description             *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricOperationType     *string                                 `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo          *string                                 `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod        *string                                 `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	Name                    *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId               *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleId          *string                                 `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	RuleComputingDefinition *string                                 `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	RuleItems               []*GetResourceRuleResponseBodyRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s GetResourceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceRuleResponseBody) SetDescription(v string) *GetResourceRuleResponseBody {
	s.Description = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetMetricOperationType(v string) *GetResourceRuleResponseBody {
	s.MetricOperationType = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetMetricPullInfo(v string) *GetResourceRuleResponseBody {
	s.MetricPullInfo = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetMetricPullPeriod(v string) *GetResourceRuleResponseBody {
	s.MetricPullPeriod = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetName(v string) *GetResourceRuleResponseBody {
	s.Name = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetRequestId(v string) *GetResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetResourceRuleId(v string) *GetResourceRuleResponseBody {
	s.ResourceRuleId = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetRuleComputingDefinition(v string) *GetResourceRuleResponseBody {
	s.RuleComputingDefinition = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetRuleItems(v []*GetResourceRuleResponseBodyRuleItems) *GetResourceRuleResponseBody {
	s.RuleItems = v
	return s
}

type GetResourceRuleResponseBodyRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxValue    *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue    *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceRuleResponseBodyRuleItems) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRuleResponseBodyRuleItems) GoString() string {
	return s.String()
}

func (s *GetResourceRuleResponseBodyRuleItems) SetDescription(v string) *GetResourceRuleResponseBodyRuleItems {
	s.Description = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetMaxValue(v string) *GetResourceRuleResponseBodyRuleItems {
	s.MaxValue = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetMinValue(v string) *GetResourceRuleResponseBodyRuleItems {
	s.MinValue = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetName(v string) *GetResourceRuleResponseBodyRuleItems {
	s.Name = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetValue(v string) *GetResourceRuleResponseBodyRuleItems {
	s.Value = &v
	return s
}

type GetResourceRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *GetResourceRuleResponse) SetHeaders(v map[string]*string) *GetResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *GetResourceRuleResponse) SetStatusCode(v int32) *GetResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceRuleResponse) SetBody(v *GetResourceRuleResponseBody) *GetResourceRuleResponse {
	s.Body = v
	return s
}

type GetSceneRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneRequest) GoString() string {
	return s.String()
}

func (s *GetSceneRequest) SetInstanceId(v string) *GetSceneRequest {
	s.InstanceId = &v
	return s
}

type GetSceneResponseBody struct {
	// example:
	//
	// This is a test.
	Description *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*GetSceneResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B8987BF7-6028-5B17-80E0-251B7BD67BBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneResponseBody) SetDescription(v string) *GetSceneResponseBody {
	s.Description = &v
	return s
}

func (s *GetSceneResponseBody) SetFlows(v []*GetSceneResponseBodyFlows) *GetSceneResponseBody {
	s.Flows = v
	return s
}

func (s *GetSceneResponseBody) SetName(v string) *GetSceneResponseBody {
	s.Name = &v
	return s
}

func (s *GetSceneResponseBody) SetRequestId(v string) *GetSceneResponseBody {
	s.RequestId = &v
	return s
}

type GetSceneResponseBodyFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s GetSceneResponseBodyFlows) String() string {
	return tea.Prettify(s)
}

func (s GetSceneResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *GetSceneResponseBodyFlows) SetFlowCode(v string) *GetSceneResponseBodyFlows {
	s.FlowCode = &v
	return s
}

func (s *GetSceneResponseBodyFlows) SetFlowName(v string) *GetSceneResponseBodyFlows {
	s.FlowName = &v
	return s
}

type GetSceneResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneResponse) GoString() string {
	return s.String()
}

func (s *GetSceneResponse) SetHeaders(v map[string]*string) *GetSceneResponse {
	s.Headers = v
	return s
}

func (s *GetSceneResponse) SetStatusCode(v int32) *GetSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneResponse) SetBody(v *GetSceneResponseBody) *GetSceneResponse {
	s.Body = v
	return s
}

type GetSubCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSubCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubCrowdRequest) GoString() string {
	return s.String()
}

func (s *GetSubCrowdRequest) SetInstanceId(v string) *GetSubCrowdRequest {
	s.InstanceId = &v
	return s
}

type GetSubCrowdResponseBody struct {
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 3
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// user1,user2
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s GetSubCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubCrowdResponseBody) SetGmtCreateTime(v string) *GetSubCrowdResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetQuantity(v string) *GetSubCrowdResponseBody {
	s.Quantity = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetRequestId(v string) *GetSubCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetSource(v string) *GetSubCrowdResponseBody {
	s.Source = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetUsers(v string) *GetSubCrowdResponseBody {
	s.Users = &v
	return s
}

type GetSubCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubCrowdResponse) GoString() string {
	return s.String()
}

func (s *GetSubCrowdResponse) SetHeaders(v map[string]*string) *GetSubCrowdResponse {
	s.Headers = v
	return s
}

func (s *GetSubCrowdResponse) SetStatusCode(v int32) *GetSubCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubCrowdResponse) SetBody(v *GetSubCrowdResponseBody) *GetSubCrowdResponse {
	s.Body = v
	return s
}

type GetTableMetaRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTableMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableMetaRequest) GoString() string {
	return s.String()
}

func (s *GetTableMetaRequest) SetInstanceId(v string) *GetTableMetaRequest {
	s.InstanceId = &v
	return s
}

type GetTableMetaResponseBody struct {
	// example:
	//
	// false
	CanDelete *bool   `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// this is a test table
	Description *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields      []*GetTableMetaResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-15:24:33
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtImportedTime *string `json:"GmtImportedTime,omitempty" xml:"GmtImportedTime,omitempty"`
	// example:
	//
	// 2021-12-15:24:33
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// test_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 28C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// reso-wkgo***
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// table_mysql
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://dmc-xxx.com/dm/table/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTableMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponseBody) SetCanDelete(v bool) *GetTableMetaResponseBody {
	s.CanDelete = &v
	return s
}

func (s *GetTableMetaResponseBody) SetConfig(v string) *GetTableMetaResponseBody {
	s.Config = &v
	return s
}

func (s *GetTableMetaResponseBody) SetDescription(v string) *GetTableMetaResponseBody {
	s.Description = &v
	return s
}

func (s *GetTableMetaResponseBody) SetFields(v []*GetTableMetaResponseBodyFields) *GetTableMetaResponseBody {
	s.Fields = v
	return s
}

func (s *GetTableMetaResponseBody) SetGmtCreateTime(v string) *GetTableMetaResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTableMetaResponseBody) SetGmtImportedTime(v string) *GetTableMetaResponseBody {
	s.GmtImportedTime = &v
	return s
}

func (s *GetTableMetaResponseBody) SetGmtModifiedTime(v string) *GetTableMetaResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTableMetaResponseBody) SetModule(v string) *GetTableMetaResponseBody {
	s.Module = &v
	return s
}

func (s *GetTableMetaResponseBody) SetName(v string) *GetTableMetaResponseBody {
	s.Name = &v
	return s
}

func (s *GetTableMetaResponseBody) SetRequestId(v string) *GetTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetResourceId(v string) *GetTableMetaResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetTableMetaId(v string) *GetTableMetaResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetTableName(v string) *GetTableMetaResponseBody {
	s.TableName = &v
	return s
}

func (s *GetTableMetaResponseBody) SetType(v string) *GetTableMetaResponseBody {
	s.Type = &v
	return s
}

func (s *GetTableMetaResponseBody) SetUrl(v string) *GetTableMetaResponseBody {
	s.Url = &v
	return s
}

type GetTableMetaResponseBodyFields struct {
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	// example:
	//
	// the gender of people
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// example:
	//
	// gender
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTableMetaResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetTableMetaResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponseBodyFields) SetIsDimensionField(v bool) *GetTableMetaResponseBodyFields {
	s.IsDimensionField = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) SetMeaning(v string) *GetTableMetaResponseBodyFields {
	s.Meaning = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) SetName(v string) *GetTableMetaResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) SetType(v string) *GetTableMetaResponseBodyFields {
	s.Type = &v
	return s
}

type GetTableMetaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableMetaResponse) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponse) SetHeaders(v map[string]*string) *GetTableMetaResponse {
	s.Headers = v
	return s
}

func (s *GetTableMetaResponse) SetStatusCode(v int32) *GetTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableMetaResponse) SetBody(v *GetTableMetaResponseBody) *GetTableMetaResponse {
	s.Body = v
	return s
}

type GetTrafficControlTargetRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetRequest) SetInstanceId(v string) *GetTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

type GetTrafficControlTargetResponseBody struct {
	EndTime                *string                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                  *string                                        `json:"Event,omitempty" xml:"Event,omitempty"`
	GmtCreateTime          *string                                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	ItemConditionArray     *string                                        `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress   *string                                        `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType      *string                                        `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                   *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation   *bool                                          `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName             *string                                        `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	RequestId              *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SplitParts             *GetTrafficControlTargetResponseBodySplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	StartTime              *string                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod           *string                                        `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status                 *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue         *int64                                         `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	TrafficControlTargetId *string                                        `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	Value                  *float32                                       `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetResponseBody) SetEndTime(v string) *GetTrafficControlTargetResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetEvent(v string) *GetTrafficControlTargetResponseBody {
	s.Event = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetGmtCreateTime(v string) *GetTrafficControlTargetResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetItemConditionArray(v string) *GetTrafficControlTargetResponseBody {
	s.ItemConditionArray = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetItemConditionExpress(v string) *GetTrafficControlTargetResponseBody {
	s.ItemConditionExpress = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetItemConditionType(v string) *GetTrafficControlTargetResponseBody {
	s.ItemConditionType = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetName(v string) *GetTrafficControlTargetResponseBody {
	s.Name = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetNewProductRegulation(v bool) *GetTrafficControlTargetResponseBody {
	s.NewProductRegulation = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetRecallName(v string) *GetTrafficControlTargetResponseBody {
	s.RecallName = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetRequestId(v string) *GetTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetSplitParts(v *GetTrafficControlTargetResponseBodySplitParts) *GetTrafficControlTargetResponseBody {
	s.SplitParts = v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetStartTime(v string) *GetTrafficControlTargetResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetStatisPeriod(v string) *GetTrafficControlTargetResponseBody {
	s.StatisPeriod = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetStatus(v string) *GetTrafficControlTargetResponseBody {
	s.Status = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetToleranceValue(v int64) *GetTrafficControlTargetResponseBody {
	s.ToleranceValue = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetTrafficControlTargetId(v string) *GetTrafficControlTargetResponseBody {
	s.TrafficControlTargetId = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetValue(v float32) *GetTrafficControlTargetResponseBody {
	s.Value = &v
	return s
}

type GetTrafficControlTargetResponseBodySplitParts struct {
	SetPoints  []*int64 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints []*int64 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s GetTrafficControlTargetResponseBodySplitParts) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTargetResponseBodySplitParts) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetResponseBodySplitParts) SetSetPoints(v []*int64) *GetTrafficControlTargetResponseBodySplitParts {
	s.SetPoints = v
	return s
}

func (s *GetTrafficControlTargetResponseBodySplitParts) SetSetValues(v []*int64) *GetTrafficControlTargetResponseBodySplitParts {
	s.SetValues = v
	return s
}

func (s *GetTrafficControlTargetResponseBodySplitParts) SetTimePoints(v []*int64) *GetTrafficControlTargetResponseBodySplitParts {
	s.TimePoints = v
	return s
}

type GetTrafficControlTargetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetResponse) SetHeaders(v map[string]*string) *GetTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficControlTargetResponse) SetStatusCode(v int32) *GetTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficControlTargetResponse) SetBody(v *GetTrafficControlTargetResponseBody) *GetTrafficControlTargetResponse {
	s.Body = v
	return s
}

type GetTrafficControlTaskRequest struct {
	ControlTargetFilter *string `json:"ControlTargetFilter,omitempty" xml:"ControlTargetFilter,omitempty"`
	Environment         *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskRequest) SetControlTargetFilter(v string) *GetTrafficControlTaskRequest {
	s.ControlTargetFilter = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetEnvironment(v string) *GetTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetInstanceId(v string) *GetTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetRegionId(v string) *GetTrafficControlTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetVersion(v string) *GetTrafficControlTaskRequest {
	s.Version = &v
	return s
}

type GetTrafficControlTaskResponseBody struct {
	BehaviorTableMetaId            *string                                                   `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	ControlGranularity             *string                                                   `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	ControlLogic                   *string                                                   `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	ControlType                    *string                                                   `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	Description                    *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                        *string                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EverPublished                  *bool                                                     `json:"EverPublished,omitempty" xml:"EverPublished,omitempty"`
	ExecutionTime                  *string                                                   `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	GmtCreateTime                  *string                                                   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime                *string                                                   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemConditionArray             *string                                                   `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress           *string                                                   `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType              *string                                                   `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	ItemTableMetaId                *string                                                   `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	Name                           *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds               *string                                                   `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	PrepubStatus                   *string                                                   `json:"PrepubStatus,omitempty" xml:"PrepubStatus,omitempty"`
	ProdExperimentIds              *string                                                   `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	ProductStatus                  *string                                                   `json:"ProductStatus,omitempty" xml:"ProductStatus,omitempty"`
	RequestId                      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId                        *string                                                   `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName                      *string                                                   `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceId                      *string                                                   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	StartTime                      *string                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisBehaviorConditionArray   *string                                                   `json:"StatisBehaviorConditionArray,omitempty" xml:"StatisBehaviorConditionArray,omitempty"`
	StatisBehaviorConditionExpress *string                                                   `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	StatisBehaviorConditionType    *string                                                   `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets          []*GetTrafficControlTaskResponseBodyTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	TrafficControlTaskId           *string                                                   `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	UserConditionArray             *string                                                   `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	UserConditionExpress           *string                                                   `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	UserConditionType              *string                                                   `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	UserTableMetaId                *string                                                   `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s GetTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponseBody) SetBehaviorTableMetaId(v string) *GetTrafficControlTaskResponseBody {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetControlGranularity(v string) *GetTrafficControlTaskResponseBody {
	s.ControlGranularity = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetControlLogic(v string) *GetTrafficControlTaskResponseBody {
	s.ControlLogic = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetControlType(v string) *GetTrafficControlTaskResponseBody {
	s.ControlType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetDescription(v string) *GetTrafficControlTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetEndTime(v string) *GetTrafficControlTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetEverPublished(v bool) *GetTrafficControlTaskResponseBody {
	s.EverPublished = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetExecutionTime(v string) *GetTrafficControlTaskResponseBody {
	s.ExecutionTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetGmtCreateTime(v string) *GetTrafficControlTaskResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetGmtModifiedTime(v string) *GetTrafficControlTaskResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemConditionArray(v string) *GetTrafficControlTaskResponseBody {
	s.ItemConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemConditionExpress(v string) *GetTrafficControlTaskResponseBody {
	s.ItemConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemConditionType(v string) *GetTrafficControlTaskResponseBody {
	s.ItemConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemTableMetaId(v string) *GetTrafficControlTaskResponseBody {
	s.ItemTableMetaId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetName(v string) *GetTrafficControlTaskResponseBody {
	s.Name = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetPreExperimentIds(v string) *GetTrafficControlTaskResponseBody {
	s.PreExperimentIds = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetPrepubStatus(v string) *GetTrafficControlTaskResponseBody {
	s.PrepubStatus = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetProdExperimentIds(v string) *GetTrafficControlTaskResponseBody {
	s.ProdExperimentIds = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetProductStatus(v string) *GetTrafficControlTaskResponseBody {
	s.ProductStatus = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetRequestId(v string) *GetTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetSceneId(v string) *GetTrafficControlTaskResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetSceneName(v string) *GetTrafficControlTaskResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetServiceId(v string) *GetTrafficControlTaskResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStartTime(v string) *GetTrafficControlTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStatisBehaviorConditionArray(v string) *GetTrafficControlTaskResponseBody {
	s.StatisBehaviorConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStatisBehaviorConditionExpress(v string) *GetTrafficControlTaskResponseBody {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStatisBehaviorConditionType(v string) *GetTrafficControlTaskResponseBody {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetTrafficControlTargets(v []*GetTrafficControlTaskResponseBodyTrafficControlTargets) *GetTrafficControlTaskResponseBody {
	s.TrafficControlTargets = v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetTrafficControlTaskId(v string) *GetTrafficControlTaskResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserConditionArray(v string) *GetTrafficControlTaskResponseBody {
	s.UserConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserConditionExpress(v string) *GetTrafficControlTaskResponseBody {
	s.UserConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserConditionType(v string) *GetTrafficControlTaskResponseBody {
	s.UserConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserTableMetaId(v string) *GetTrafficControlTaskResponseBody {
	s.UserTableMetaId = &v
	return s
}

type GetTrafficControlTaskResponseBodyTrafficControlTargets struct {
	EndTime                *string                                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                  *string                                                           `json:"Event,omitempty" xml:"Event,omitempty"`
	GmtCreateTime          *string                                                           `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string                                                           `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemConditionArray     *string                                                           `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress   *string                                                           `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType      *string                                                           `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                   *string                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation   *bool                                                             `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName             *string                                                           `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	SplitParts             *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	StartTime              *string                                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod           *string                                                           `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status                 *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue         *int64                                                            `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	TrafficControlTargetId *string                                                           `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	Value                  *float32                                                          `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargets) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetEndTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetEvent(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetGmtCreateTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetGmtModifiedTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetItemConditionArray(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetItemConditionExpress(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetItemConditionType(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetName(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetNewProductRegulation(v bool) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetRecallName(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetSplitParts(v *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.SplitParts = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetStartTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetStatisPeriod(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetStatus(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetToleranceValue(v int64) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetTrafficControlTargetId(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.TrafficControlTargetId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetValue(v float32) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Value = &v
	return s
}

type GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts struct {
	SetPoints  []*int32 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints []*int32 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) SetSetPoints(v []*int32) *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	s.SetPoints = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) SetSetValues(v []*int64) *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	s.SetValues = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) SetTimePoints(v []*int32) *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	s.TimePoints = v
	return s
}

type GetTrafficControlTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponse) SetHeaders(v map[string]*string) *GetTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficControlTaskResponse) SetStatusCode(v int32) *GetTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficControlTaskResponse) SetBody(v *GetTrafficControlTaskResponseBody) *GetTrafficControlTaskResponse {
	s.Body = v
	return s
}

type GetTrafficControlTaskTrafficRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTrafficControlTaskTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskTrafficRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficRequest) SetEnvironment(v string) *GetTrafficControlTaskTrafficRequest {
	s.Environment = &v
	return s
}

func (s *GetTrafficControlTaskTrafficRequest) SetInstanceId(v string) *GetTrafficControlTaskTrafficRequest {
	s.InstanceId = &v
	return s
}

type GetTrafficControlTaskTrafficResponseBody struct {
	RequestId                     *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTaskTrafficInfo *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo `json:"TrafficControlTaskTrafficInfo,omitempty" xml:"TrafficControlTaskTrafficInfo,omitempty" type:"Struct"`
}

func (s GetTrafficControlTaskTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponseBody) SetRequestId(v string) *GetTrafficControlTaskTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBody) SetTrafficControlTaskTrafficInfo(v *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) *GetTrafficControlTaskTrafficResponseBody {
	s.TrafficControlTaskTrafficInfo = v
	return s
}

type GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo struct {
	TargetTraffics []*GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics `json:"TargetTraffics,omitempty" xml:"TargetTraffics,omitempty" type:"Repeated"`
	TaskTraffics   map[string]*TrafficControlTaskTrafficInfoTaskTrafficsValue                             `json:"TaskTraffics,omitempty" xml:"TaskTraffics,omitempty"`
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) SetTargetTraffics(v []*GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo {
	s.TargetTraffics = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) SetTaskTraffics(v map[string]*TrafficControlTaskTrafficInfoTaskTrafficsValue) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo {
	s.TaskTraffics = v
	return s
}

type GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics struct {
	Data                   map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	TrafficContorlTargetId *string                                                          `json:"TrafficContorlTargetId,omitempty" xml:"TrafficContorlTargetId,omitempty"`
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) SetData(v map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	s.Data = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) SetTrafficContorlTargetId(v string) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	s.TrafficContorlTargetId = &v
	return s
}

type GetTrafficControlTaskTrafficResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficControlTaskTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficControlTaskTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponse) SetHeaders(v map[string]*string) *GetTrafficControlTaskTrafficResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponse) SetStatusCode(v int32) *GetTrafficControlTaskTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficControlTaskTrafficResponse) SetBody(v *GetTrafficControlTaskTrafficResponseBody) *GetTrafficControlTaskTrafficResponse {
	s.Body = v
	return s
}

type ListABMetricGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SortBy  *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListABMetricGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsRequest) SetInstanceId(v string) *ListABMetricGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetOrder(v string) *ListABMetricGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetPageNumber(v int32) *ListABMetricGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetPageSize(v int32) *ListABMetricGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetRealtime(v bool) *ListABMetricGroupsRequest {
	s.Realtime = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetSceneId(v string) *ListABMetricGroupsRequest {
	s.SceneId = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetSortBy(v string) *ListABMetricGroupsRequest {
	s.SortBy = &v
	return s
}

type ListABMetricGroupsResponseBody struct {
	ABMetricGroups []*ListABMetricGroupsResponseBodyABMetricGroups `json:"ABMetricGroups,omitempty" xml:"ABMetricGroups,omitempty" type:"Repeated"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListABMetricGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsResponseBody) SetABMetricGroups(v []*ListABMetricGroupsResponseBodyABMetricGroups) *ListABMetricGroupsResponseBody {
	s.ABMetricGroups = v
	return s
}

func (s *ListABMetricGroupsResponseBody) SetRequestId(v string) *ListABMetricGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABMetricGroupsResponseBody) SetTotalCount(v int64) *ListABMetricGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListABMetricGroupsResponseBodyABMetricGroups struct {
	// example:
	//
	// 1
	ABMetricGroupId *string `json:"ABMetricGroupId,omitempty" xml:"ABMetricGroupId,omitempty"`
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// example:
	//
	// pv,uv
	ABMetricNames *string `json:"ABMetricNames,omitempty" xml:"ABMetricNames,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2799614***
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListABMetricGroupsResponseBodyABMetricGroups) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricGroupsResponseBodyABMetricGroups) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetABMetricGroupId(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.ABMetricGroupId = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetABMetricIds(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.ABMetricIds = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetABMetricNames(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.ABMetricNames = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetDescription(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Description = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetName(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Name = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetOwner(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Owner = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetRealtime(v bool) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Realtime = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetSceneId(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.SceneId = &v
	return s
}

type ListABMetricGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABMetricGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABMetricGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsResponse) SetHeaders(v map[string]*string) *ListABMetricGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListABMetricGroupsResponse) SetStatusCode(v int32) *ListABMetricGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABMetricGroupsResponse) SetBody(v *ListABMetricGroupsResponseBody) *ListABMetricGroupsResponse {
	s.Body = v
	return s
}

type ListABMetricsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 1
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListABMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListABMetricsRequest) SetInstanceId(v string) *ListABMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListABMetricsRequest) SetName(v string) *ListABMetricsRequest {
	s.Name = &v
	return s
}

func (s *ListABMetricsRequest) SetPageNumber(v int32) *ListABMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListABMetricsRequest) SetPageSize(v int32) *ListABMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListABMetricsRequest) SetRealtime(v bool) *ListABMetricsRequest {
	s.Realtime = &v
	return s
}

func (s *ListABMetricsRequest) SetSceneId(v string) *ListABMetricsRequest {
	s.SceneId = &v
	return s
}

func (s *ListABMetricsRequest) SetTableMetaId(v string) *ListABMetricsRequest {
	s.TableMetaId = &v
	return s
}

func (s *ListABMetricsRequest) SetType(v string) *ListABMetricsRequest {
	s.Type = &v
	return s
}

type ListABMetricsResponseBody struct {
	ABMetrics []*ListABMetricsResponseBodyABMetrics `json:"ABMetrics,omitempty" xml:"ABMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListABMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponseBody) SetABMetrics(v []*ListABMetricsResponseBodyABMetrics) *ListABMetricsResponseBody {
	s.ABMetrics = v
	return s
}

func (s *ListABMetricsResponseBody) SetRequestId(v string) *ListABMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABMetricsResponseBody) SetTotalCount(v int64) *ListABMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type ListABMetricsResponseBodyABMetrics struct {
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	// example:
	//
	// sum(click_cnt)
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// false
	Realtime *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 2
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	// example:
	//
	// 2
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// home_feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// example:
	//
	// 1
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListABMetricsResponseBodyABMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricsResponseBodyABMetrics) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponseBodyABMetrics) SetABMetricId(v string) *ListABMetricsResponseBodyABMetrics {
	s.ABMetricId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDefinition(v string) *ListABMetricsResponseBodyABMetrics {
	s.Definition = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDescription(v string) *ListABMetricsResponseBodyABMetrics {
	s.Description = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetLeftMetricId(v string) *ListABMetricsResponseBodyABMetrics {
	s.LeftMetricId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetName(v string) *ListABMetricsResponseBodyABMetrics {
	s.Name = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetOperator(v string) *ListABMetricsResponseBodyABMetrics {
	s.Operator = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetRealtime(v string) *ListABMetricsResponseBodyABMetrics {
	s.Realtime = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetResultResourceId(v string) *ListABMetricsResponseBodyABMetrics {
	s.ResultResourceId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetResultTableMetaId(v string) *ListABMetricsResponseBodyABMetrics {
	s.ResultTableMetaId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetRightMetricId(v string) *ListABMetricsResponseBodyABMetrics {
	s.RightMetricId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetSceneId(v string) *ListABMetricsResponseBodyABMetrics {
	s.SceneId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetSceneName(v string) *ListABMetricsResponseBodyABMetrics {
	s.SceneName = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetStatisticsCycle(v int32) *ListABMetricsResponseBodyABMetrics {
	s.StatisticsCycle = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetTableMetaId(v string) *ListABMetricsResponseBodyABMetrics {
	s.TableMetaId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetType(v string) *ListABMetricsResponseBodyABMetrics {
	s.Type = &v
	return s
}

type ListABMetricsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponse) SetHeaders(v map[string]*string) *ListABMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListABMetricsResponse) SetStatusCode(v int32) *ListABMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABMetricsResponse) SetBody(v *ListABMetricsResponseBody) *ListABMetricsResponse {
	s.Body = v
	return s
}

type ListCalculationJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCalculationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCalculationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsRequest) SetInstanceId(v string) *ListCalculationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCalculationJobsRequest) SetPageNumber(v int32) *ListCalculationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCalculationJobsRequest) SetPageSize(v int32) *ListCalculationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCalculationJobsRequest) SetSceneId(v string) *ListCalculationJobsRequest {
	s.SceneId = &v
	return s
}

func (s *ListCalculationJobsRequest) SetStatus(v string) *ListCalculationJobsRequest {
	s.Status = &v
	return s
}

type ListCalculationJobsResponseBody struct {
	CalculationJobs []*ListCalculationJobsResponseBodyCalculationJobs `json:"CalculationJobs,omitempty" xml:"CalculationJobs,omitempty" type:"Repeated"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCalculationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCalculationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsResponseBody) SetCalculationJobs(v []*ListCalculationJobsResponseBodyCalculationJobs) *ListCalculationJobsResponseBody {
	s.CalculationJobs = v
	return s
}

func (s *ListCalculationJobsResponseBody) SetRequestId(v string) *ListCalculationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalculationJobsResponseBody) SetTotalCount(v int64) *ListCalculationJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCalculationJobsResponseBodyCalculationJobs struct {
	// example:
	//
	// pv
	ABMetricName *string `json:"ABMetricName,omitempty" xml:"ABMetricName,omitempty"`
	// example:
	//
	// 2021-12-15
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2
	CalculationJobId *string `json:"CalculationJobId,omitempty" xml:"CalculationJobId,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtRanTime *string   `json:"GmtRanTime,omitempty" xml:"GmtRanTime,omitempty"`
	JobMessage []*string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty" type:"Repeated"`
	// example:
	//
	// CronOffline
	JobSource *string `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCalculationJobsResponseBodyCalculationJobs) String() string {
	return tea.Prettify(s)
}

func (s ListCalculationJobsResponseBodyCalculationJobs) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetABMetricName(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.ABMetricName = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetBizDate(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.BizDate = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetCalculationJobId(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.CalculationJobId = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetConfig(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.Config = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetGmtRanTime(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.GmtRanTime = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetJobMessage(v []*string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.JobMessage = v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetJobSource(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.JobSource = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetStatus(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.Status = &v
	return s
}

type ListCalculationJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalculationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalculationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCalculationJobsResponse) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsResponse) SetHeaders(v map[string]*string) *ListCalculationJobsResponse {
	s.Headers = v
	return s
}

func (s *ListCalculationJobsResponse) SetStatusCode(v int32) *ListCalculationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalculationJobsResponse) SetBody(v *ListCalculationJobsResponseBody) *ListCalculationJobsResponse {
	s.Body = v
	return s
}

type ListCrowdUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCrowdUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdUsersRequest) GoString() string {
	return s.String()
}

func (s *ListCrowdUsersRequest) SetInstanceId(v string) *ListCrowdUsersRequest {
	s.InstanceId = &v
	return s
}

type ListCrowdUsersResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F0AB6527-093F-5C44-B3BD-42C8C210C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListCrowdUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrowdUsersResponseBody) SetRequestId(v string) *ListCrowdUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrowdUsersResponseBody) SetTotalCount(v int64) *ListCrowdUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCrowdUsersResponseBody) SetUsers(v []*string) *ListCrowdUsersResponseBody {
	s.Users = v
	return s
}

type ListCrowdUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrowdUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrowdUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdUsersResponse) GoString() string {
	return s.String()
}

func (s *ListCrowdUsersResponse) SetHeaders(v map[string]*string) *ListCrowdUsersResponse {
	s.Headers = v
	return s
}

func (s *ListCrowdUsersResponse) SetStatusCode(v int32) *ListCrowdUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrowdUsersResponse) SetBody(v *ListCrowdUsersResponseBody) *ListCrowdUsersResponse {
	s.Body = v
	return s
}

type ListCrowdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCrowdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdsRequest) GoString() string {
	return s.String()
}

func (s *ListCrowdsRequest) SetInstanceId(v string) *ListCrowdsRequest {
	s.InstanceId = &v
	return s
}

type ListCrowdsResponseBody struct {
	Crowds []*ListCrowdsResponseBodyCrowds `json:"Crowds,omitempty" xml:"Crowds,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrowdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponseBody) SetCrowds(v []*ListCrowdsResponseBodyCrowds) *ListCrowdsResponseBody {
	s.Crowds = v
	return s
}

func (s *ListCrowdsResponseBody) SetRequestId(v string) *ListCrowdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrowdsResponseBody) SetTotalCount(v int64) *ListCrowdsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCrowdsResponseBodyCrowds struct {
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// os=android
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// crowd1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// user1,user2
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ListCrowdsResponseBodyCrowds) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdsResponseBodyCrowds) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponseBodyCrowds) SetCrowdId(v string) *ListCrowdsResponseBodyCrowds {
	s.CrowdId = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetDescription(v string) *ListCrowdsResponseBodyCrowds {
	s.Description = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetGmtCreateTime(v string) *ListCrowdsResponseBodyCrowds {
	s.GmtCreateTime = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetLabel(v string) *ListCrowdsResponseBodyCrowds {
	s.Label = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetName(v string) *ListCrowdsResponseBodyCrowds {
	s.Name = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetQuantity(v string) *ListCrowdsResponseBodyCrowds {
	s.Quantity = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetSource(v string) *ListCrowdsResponseBodyCrowds {
	s.Source = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetUsers(v string) *ListCrowdsResponseBodyCrowds {
	s.Users = &v
	return s
}

type ListCrowdsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrowdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrowdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdsResponse) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponse) SetHeaders(v map[string]*string) *ListCrowdsResponse {
	s.Headers = v
	return s
}

func (s *ListCrowdsResponse) SetStatusCode(v int32) *ListCrowdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrowdsResponse) SetBody(v *ListCrowdsResponseBody) *ListCrowdsResponse {
	s.Body = v
	return s
}

type ListEngineConfigsRequest struct {
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// latest
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListEngineConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEngineConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsRequest) SetEnvironment(v string) *ListEngineConfigsRequest {
	s.Environment = &v
	return s
}

func (s *ListEngineConfigsRequest) SetInstanceId(v string) *ListEngineConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListEngineConfigsRequest) SetName(v string) *ListEngineConfigsRequest {
	s.Name = &v
	return s
}

func (s *ListEngineConfigsRequest) SetPageNumber(v int32) *ListEngineConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEngineConfigsRequest) SetPageSize(v int32) *ListEngineConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEngineConfigsRequest) SetStatus(v string) *ListEngineConfigsRequest {
	s.Status = &v
	return s
}

func (s *ListEngineConfigsRequest) SetVersion(v string) *ListEngineConfigsRequest {
	s.Version = &v
	return s
}

type ListEngineConfigsResponseBody struct {
	EngineConfigs []*ListEngineConfigsResponseBodyEngineConfigs `json:"EngineConfigs,omitempty" xml:"EngineConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEngineConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEngineConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsResponseBody) SetEngineConfigs(v []*ListEngineConfigsResponseBodyEngineConfigs) *ListEngineConfigsResponseBody {
	s.EngineConfigs = v
	return s
}

func (s *ListEngineConfigsResponseBody) SetRequestId(v string) *ListEngineConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEngineConfigsResponseBody) SetTotalCount(v int64) *ListEngineConfigsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEngineConfigsResponseBodyEngineConfigs struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// 2023-08-07T01:43:42Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-08-27T12:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2023-08-29 12:00:00
	GmtReleasedTime *string `json:"GmtReleasedTime,omitempty" xml:"GmtReleasedTime,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20230509161300
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListEngineConfigsResponseBodyEngineConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListEngineConfigsResponseBodyEngineConfigs) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetConfigValue(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.ConfigValue = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetDescription(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Description = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetEngineConfigId(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.EngineConfigId = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetEnvironment(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Environment = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetGmtCreateTime(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetGmtModifiedTime(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetGmtReleasedTime(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.GmtReleasedTime = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetName(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Name = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetStatus(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Status = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetVersion(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Version = &v
	return s
}

type ListEngineConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEngineConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEngineConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEngineConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsResponse) SetHeaders(v map[string]*string) *ListEngineConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListEngineConfigsResponse) SetStatusCode(v int32) *ListEngineConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEngineConfigsResponse) SetBody(v *ListEngineConfigsResponseBody) *ListEngineConfigsResponse {
	s.Body = v
	return s
}

type ListExperimentGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 3
	LayerId  *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Online
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeRangeEnd   *string `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
}

func (s ListExperimentGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsRequest) SetInstanceId(v string) *ListExperimentGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetLayerId(v string) *ListExperimentGroupsRequest {
	s.LayerId = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetRegionId(v string) *ListExperimentGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetStatus(v string) *ListExperimentGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetTimeRangeEnd(v string) *ListExperimentGroupsRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetTimeRangeStart(v string) *ListExperimentGroupsRequest {
	s.TimeRangeStart = &v
	return s
}

type ListExperimentGroupsResponseBody struct {
	ExperimentGroups []*ListExperimentGroupsResponseBodyExperimentGroups `json:"ExperimentGroups,omitempty" xml:"ExperimentGroups,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsResponseBody) SetExperimentGroups(v []*ListExperimentGroupsResponseBodyExperimentGroups) *ListExperimentGroupsResponseBody {
	s.ExperimentGroups = v
	return s
}

func (s *ListExperimentGroupsResponseBody) SetRequestId(v string) *ListExperimentGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentGroupsResponseBody) SetTotalCount(v int64) *ListExperimentGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentGroupsResponseBodyExperimentGroups struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 4
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 5
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// example:
	//
	// gender=female
	Filter         *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	HoldingBuckets *string `json:"HoldingBuckets,omitempty" xml:"HoldingBuckets,omitempty"`
	// example:
	//
	// 4
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NeedAA *bool `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	// example:
	//
	// 1124512470******
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RandomFlow *int64  `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// example:
	//
	// 1,2,3,4
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExperimentGroupsResponseBodyExperimentGroups) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentGroupsResponseBodyExperimentGroups) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetConfig(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Config = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetCrowdId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.CrowdId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetCrowdTargetType(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.CrowdTargetType = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDebugCrowdId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DebugCrowdId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDebugUsers(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DebugUsers = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDescription(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Description = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDistributionTimeDuration(v int32) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DistributionTimeDuration = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDistributionType(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DistributionType = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetExperimentGroupId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetFilter(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Filter = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetHoldingBuckets(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.HoldingBuckets = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetLaboratoryId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.LaboratoryId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetLayerId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.LayerId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetName(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Name = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetNeedAA(v bool) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.NeedAA = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetOwner(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Owner = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetRandomFlow(v int64) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.RandomFlow = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetReservedBuckets(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.ReservedBuckets = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetSceneId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.SceneId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetStatus(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Status = &v
	return s
}

type ListExperimentGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsResponse) SetHeaders(v map[string]*string) *ListExperimentGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentGroupsResponse) SetStatusCode(v int32) *ListExperimentGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentGroupsResponse) SetBody(v *ListExperimentGroupsResponseBody) *ListExperimentGroupsResponse {
	s.Body = v
	return s
}

type ListExperimentsRequest struct {
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// experiment_test1
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentsRequest) SetExperimentGroupId(v string) *ListExperimentsRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListExperimentsRequest) SetInstanceId(v string) *ListExperimentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExperimentsRequest) SetQuery(v string) *ListExperimentsRequest {
	s.Query = &v
	return s
}

func (s *ListExperimentsRequest) SetStatus(v string) *ListExperimentsRequest {
	s.Status = &v
	return s
}

type ListExperimentsResponseBody struct {
	Experiments []*ListExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 68075085-1A7D-55C2-B51D-7AD9B02A6DD6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBody) SetExperiments(v []*ListExperimentsResponseBodyExperiments) *ListExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListExperimentsResponseBody) SetRequestId(v string) *ListExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentsResponseBody) SetTotalCount(v int64) *ListExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentsResponseBodyExperiments struct {
	// example:
	//
	// L1#EG1#E1
	AliasExperimentId *string `json:"AliasExperimentId,omitempty" xml:"AliasExperimentId,omitempty"`
	// example:
	//
	// 1,2,3
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// uid1,uid2,uid3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// example:
	//
	// 3
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
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
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// experiment_test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListExperimentsResponseBodyExperiments) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyExperiments) SetAliasExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.AliasExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetBuckets(v string) *ListExperimentsResponseBodyExperiments {
	s.Buckets = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetConfig(v string) *ListExperimentsResponseBodyExperiments {
	s.Config = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDebugCrowdId(v string) *ListExperimentsResponseBodyExperiments {
	s.DebugCrowdId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDebugUsers(v string) *ListExperimentsResponseBodyExperiments {
	s.DebugUsers = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDescription(v string) *ListExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentGroupId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetFlowPercent(v int32) *ListExperimentsResponseBodyExperiments {
	s.FlowPercent = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtCreateTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtCreateTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtModifiedTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetLaboratoryId(v string) *ListExperimentsResponseBodyExperiments {
	s.LaboratoryId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetLayerId(v string) *ListExperimentsResponseBodyExperiments {
	s.LayerId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetName(v string) *ListExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetSceneId(v string) *ListExperimentsResponseBodyExperiments {
	s.SceneId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetStatus(v string) *ListExperimentsResponseBodyExperiments {
	s.Status = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetType(v string) *ListExperimentsResponseBodyExperiments {
	s.Type = &v
	return s
}

type ListExperimentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponse) SetHeaders(v map[string]*string) *ListExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentsResponse) SetStatusCode(v int32) *ListExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentsResponse) SetBody(v *ListExperimentsResponseBody) *ListExperimentsResponse {
	s.Body = v
	return s
}

type ListFeatureConsistencyCheckJobConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListFeatureConsistencyCheckJobConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetOrder(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetPageNumber(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetPageSize(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetSortBy(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.SortBy = &v
	return s
}

type ListFeatureConsistencyCheckJobConfigsResponseBody struct {
	FeatureConsistencyCheckConfigs []*ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs `json:"FeatureConsistencyCheckConfigs,omitempty" xml:"FeatureConsistencyCheckConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// FCF741D8-9C30-578E-807F-B935487DB34A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBody) GoString() string {
	return s.String()
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

type ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs struct {
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
	// 3
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
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
	UseFeatureStore *string `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
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

func (s ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) GoString() string {
	return s.String()
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

type ListFeatureConsistencyCheckJobConfigsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) SetBody(v *ListFeatureConsistencyCheckJobConfigsResponseBody) *ListFeatureConsistencyCheckJobConfigsResponse {
	s.Body = v
	return s
}

type ListFeatureConsistencyCheckJobFeatureReportsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetLogItemId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.LogItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetLogRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetLogUserId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.LogUserId = &v
	return s
}

type ListFeatureConsistencyCheckJobFeatureReportsResponseBody struct {
	// example:
	//
	// https://********
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// example:
	//
	// oss://********
	OssPath              *string                                                                         `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	ReportsOfFeatureDiff []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff `json:"ReportsOfFeatureDiff,omitempty" xml:"ReportsOfFeatureDiff,omitempty" type:"Repeated"`
	// example:
	//
	// BBD41FBF-E75C-551A-92FA-CAD654AA006F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetDataPath(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.DataPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetOssPath(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.OssPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetReportsOfFeatureDiff(v []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.ReportsOfFeatureDiff = v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.RequestId = &v
	return s
}

type ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff struct {
	// example:
	//
	// gender
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// example:
	//
	// male
	OfflineValue *string `json:"OfflineValue,omitempty" xml:"OfflineValue,omitempty"`
	// example:
	//
	// male
	OnlineValue *string `json:"OnlineValue,omitempty" xml:"OnlineValue,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetFeatureName(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetLogItemId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.LogItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetLogRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.LogRequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetLogUserId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.LogUserId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetOfflineValue(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.OfflineValue = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetOnlineValue(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.OnlineValue = &v
	return s
}

type ListFeatureConsistencyCheckJobFeatureReportsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobFeatureReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobFeatureReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobFeatureReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) SetBody(v *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) *ListFeatureConsistencyCheckJobFeatureReportsResponse {
	s.Body = v
	return s
}

type ListFeatureConsistencyCheckJobScoreReportsRequest struct {
	ExcludeRequestIds []*string `json:"ExcludeRequestIds,omitempty" xml:"ExcludeRequestIds,omitempty" type:"Repeated"`
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) SetExcludeRequestIds(v []*string) *ListFeatureConsistencyCheckJobScoreReportsRequest {
	s.ExcludeRequestIds = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobScoreReportsRequest {
	s.InstanceId = &v
	return s
}

type ListFeatureConsistencyCheckJobScoreReportsShrinkRequest struct {
	ExcludeRequestIdsShrink *string `json:"ExcludeRequestIds,omitempty" xml:"ExcludeRequestIds,omitempty"`
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) SetExcludeRequestIdsShrink(v string) *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest {
	s.ExcludeRequestIdsShrink = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest {
	s.InstanceId = &v
	return s
}

type ListFeatureConsistencyCheckJobScoreReportsResponseBody struct {
	// example:
	//
	// http://*******
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// example:
	//
	// oss://********
	OssPath            *string                                                                     `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	ReportsOfScoreDiff []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff `json:"ReportsOfScoreDiff,omitempty" xml:"ReportsOfScoreDiff,omitempty" type:"Repeated"`
	// example:
	//
	// F0AB6527-093F-5C44-B3BD-42C8C210C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetDataPath(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.DataPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetOssPath(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.OssPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetReportsOfScoreDiff(v []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.ReportsOfScoreDiff = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.RequestId = &v
	return s
}

type ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff struct {
	// example:
	//
	// 4
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// example:
	//
	// 323
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// example:
	//
	// 3
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// example:
	//
	// 0.00000234
	ScoreDiff *string `json:"ScoreDiff,omitempty" xml:"ScoreDiff,omitempty"`
	// example:
	//
	// {}
	ScoreDiffDetail *string `json:"ScoreDiffDetail,omitempty" xml:"ScoreDiffDetail,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetLogItemId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.LogItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetLogRequestId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.LogRequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetLogUserId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.LogUserId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetScoreDiff(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.ScoreDiff = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetScoreDiffDetail(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.ScoreDiffDetail = &v
	return s
}

type ListFeatureConsistencyCheckJobScoreReportsResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobScoreReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobScoreReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobScoreReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) SetBody(v *ListFeatureConsistencyCheckJobScoreReportsResponseBody) *ListFeatureConsistencyCheckJobScoreReportsResponse {
	s.Body = v
	return s
}

type ListFeatureConsistencyCheckJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetOrder(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetPageNumber(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetPageSize(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetSortBy(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetStatus(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.Status = &v
	return s
}

type ListFeatureConsistencyCheckJobsResponseBody struct {
	FeatureConsistencyCheckJobs []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs `json:"FeatureConsistencyCheckJobs,omitempty" xml:"FeatureConsistencyCheckJobs,omitempty" type:"Repeated"`
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) SetFeatureConsistencyCheckJobs(v []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) *ListFeatureConsistencyCheckJobsResponseBody {
	s.FeatureConsistencyCheckJobs = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) SetTotalCount(v string) *ListFeatureConsistencyCheckJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 5
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// example:
	//
	// feature_consistency_check_1
	FeatureConsistencyCheckJobConfigName *string `json:"FeatureConsistencyCheckJobConfigName,omitempty" xml:"FeatureConsistencyCheckJobConfigName,omitempty"`
	// example:
	//
	// 4
	FeatureConsistencyCheckJobId *string `json:"FeatureConsistencyCheckJobId,omitempty" xml:"FeatureConsistencyCheckJobId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtStartTime *string   `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Logs         []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetConfig(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.Config = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetFeatureConsistencyCheckJobConfigId(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetFeatureConsistencyCheckJobConfigName(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.FeatureConsistencyCheckJobConfigName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetFeatureConsistencyCheckJobId(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.FeatureConsistencyCheckJobId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetGmtEndTime(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.GmtEndTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetGmtStartTime(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.GmtStartTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetLogs(v []*string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.Logs = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetStatus(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.Status = &v
	return s
}

type ListFeatureConsistencyCheckJobsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponse) SetBody(v *ListFeatureConsistencyCheckJobsResponseBody) *ListFeatureConsistencyCheckJobsResponse {
	s.Body = v
	return s
}

type ListInstanceResourcesRequest struct {
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesRequest) SetCategory(v string) *ListInstanceResourcesRequest {
	s.Category = &v
	return s
}

func (s *ListInstanceResourcesRequest) SetGroup(v string) *ListInstanceResourcesRequest {
	s.Group = &v
	return s
}

func (s *ListInstanceResourcesRequest) SetType(v string) *ListInstanceResourcesRequest {
	s.Type = &v
	return s
}

type ListInstanceResourcesResponseBody struct {
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*ListInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesResponseBody) SetRequestId(v string) *ListInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResourcesResponseBody) SetResources(v []*ListInstanceResourcesResponseBodyResources) *ListInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListInstanceResourcesResponseBody) SetTotalCount(v int64) *ListInstanceResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceResourcesResponseBodyResources struct {
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtCreateAt *string `json:"GmtCreateAt,omitempty" xml:"GmtCreateAt,omitempty"`
	// example:
	//
	// 2020-10-13 17:34:52
	GmtModifiedAt *string `json:"GmtModifiedAt,omitempty" xml:"GmtModifiedAt,omitempty"`
	// example:
	//
	// storage
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListInstanceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesResponseBodyResources) SetCategory(v string) *ListInstanceResourcesResponseBodyResources {
	s.Category = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetConfig(v string) *ListInstanceResourcesResponseBodyResources {
	s.Config = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetGmtCreateAt(v string) *ListInstanceResourcesResponseBodyResources {
	s.GmtCreateAt = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetGmtModifiedAt(v string) *ListInstanceResourcesResponseBodyResources {
	s.GmtModifiedAt = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetGroup(v string) *ListInstanceResourcesResponseBodyResources {
	s.Group = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetResourceId(v string) *ListInstanceResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetType(v string) *ListInstanceResourcesResponseBodyResources {
	s.Type = &v
	return s
}

func (s *ListInstanceResourcesResponseBodyResources) SetUri(v string) *ListInstanceResourcesResponseBodyResources {
	s.Uri = &v
	return s
}

type ListInstanceResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResourcesResponse) SetStatusCode(v int32) *ListInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResourcesResponse) SetBody(v *ListInstanceResourcesResponseBody) *ListInstanceResourcesResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Type
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetOrder(v string) *ListInstancesRequest {
	s.Order = &v
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

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetType(v string) *ListInstancesRequest {
	s.Type = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// example:
	//
	// Subscription
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// airec_developers_public_cn
	CommodityCode *string                                   `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Config        *ListInstancesResponseBodyInstancesConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 2022-12-14 00:00:00.0
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// 2022-10-13 17:34:52.0
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2022-11-05 09:02:30.0
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// pairec-test1
	InstanceId    *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatingTool *ListInstancesResponseBodyInstancesOperatingTool `json:"OperatingTool,omitempty" xml:"OperatingTool,omitempty" type:"Struct"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetChargeType(v string) *ListInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCommodityCode(v string) *ListInstancesResponseBodyInstances {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetConfig(v *ListInstancesResponseBodyInstancesConfig) *ListInstancesResponseBodyInstances {
	s.Config = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetExpiredTime(v string) *ListInstancesResponseBodyInstances {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOperatingTool(v *ListInstancesResponseBodyInstancesOperatingTool) *ListInstancesResponseBodyInstances {
	s.OperatingTool = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetType(v string) *ListInstancesResponseBodyInstances {
	s.Type = &v
	return s
}

type ListInstancesResponseBodyInstancesConfig struct {
	DataManagements []*ListInstancesResponseBodyInstancesConfigDataManagements `json:"DataManagements,omitempty" xml:"DataManagements,omitempty" type:"Repeated"`
	Engines         []*ListInstancesResponseBodyInstancesConfigEngines         `json:"Engines,omitempty" xml:"Engines,omitempty" type:"Repeated"`
	Monitors        []*ListInstancesResponseBodyInstancesConfigMonitors        `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstancesConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfig) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfig) SetDataManagements(v []*ListInstancesResponseBodyInstancesConfigDataManagements) *ListInstancesResponseBodyInstancesConfig {
	s.DataManagements = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfig) SetEngines(v []*ListInstancesResponseBodyInstancesConfigEngines) *ListInstancesResponseBodyInstancesConfig {
	s.Engines = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfig) SetMonitors(v []*ListInstancesResponseBodyInstancesConfigMonitors) *ListInstancesResponseBodyInstancesConfig {
	s.Monitors = v
	return s
}

type ListInstancesResponseBodyInstancesConfigDataManagements struct {
	// example:
	//
	// storage
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstancesConfigDataManagements) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfigDataManagements) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) SetComponentCode(v string) *ListInstancesResponseBodyInstancesConfigDataManagements {
	s.ComponentCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) SetMeta(v map[string]interface{}) *ListInstancesResponseBodyInstancesConfigDataManagements {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigDataManagements) SetType(v string) *ListInstancesResponseBodyInstancesConfigDataManagements {
	s.Type = &v
	return s
}

type ListInstancesResponseBodyInstancesConfigEngines struct {
	// example:
	//
	// feature
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstancesConfigEngines) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfigEngines) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) SetComponentCode(v string) *ListInstancesResponseBodyInstancesConfigEngines {
	s.ComponentCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) SetMeta(v map[string]interface{}) *ListInstancesResponseBodyInstancesConfigEngines {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigEngines) SetType(v string) *ListInstancesResponseBodyInstancesConfigEngines {
	s.Type = &v
	return s
}

type ListInstancesResponseBodyInstancesConfigMonitors struct {
	// example:
	//
	// featuresets
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// Platform
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstancesConfigMonitors) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesConfigMonitors) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) SetComponentCode(v string) *ListInstancesResponseBodyInstancesConfigMonitors {
	s.ComponentCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) SetMeta(v map[string]interface{}) *ListInstancesResponseBodyInstancesConfigMonitors {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyInstancesConfigMonitors) SetType(v string) *ListInstancesResponseBodyInstancesConfigMonitors {
	s.Type = &v
	return s
}

type ListInstancesResponseBodyInstancesOperatingTool struct {
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
}

func (s ListInstancesResponseBodyInstancesOperatingTool) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesOperatingTool) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesOperatingTool) SetIsEnable(v bool) *ListInstancesResponseBodyInstancesOperatingTool {
	s.IsEnable = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListLaboratoriesRequest struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLaboratoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLaboratoriesRequest) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesRequest) SetEnvironment(v string) *ListLaboratoriesRequest {
	s.Environment = &v
	return s
}

func (s *ListLaboratoriesRequest) SetInstanceId(v string) *ListLaboratoriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLaboratoriesRequest) SetSceneId(v string) *ListLaboratoriesRequest {
	s.SceneId = &v
	return s
}

func (s *ListLaboratoriesRequest) SetStatus(v string) *ListLaboratoriesRequest {
	s.Status = &v
	return s
}

type ListLaboratoriesResponseBody struct {
	Laboratories []*ListLaboratoriesResponseBodyLaboratories `json:"Laboratories,omitempty" xml:"Laboratories,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 1C0898E5-9220-5443-B2D9-445FF0688215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLaboratoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLaboratoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesResponseBody) SetLaboratories(v []*ListLaboratoriesResponseBodyLaboratories) *ListLaboratoriesResponseBody {
	s.Laboratories = v
	return s
}

func (s *ListLaboratoriesResponseBody) SetRequestId(v string) *ListLaboratoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLaboratoriesResponseBody) SetTotalCount(v int64) *ListLaboratoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListLaboratoriesResponseBodyLaboratories struct {
	// example:
	//
	// 100
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// example:
	//
	// Filter
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// example:
	//
	// 1,2,3,10-20
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// user1,user2,user3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// filter=xxx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLaboratoriesResponseBodyLaboratories) String() string {
	return tea.Prettify(s)
}

func (s ListLaboratoriesResponseBodyLaboratories) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetBucketCount(v int32) *ListLaboratoriesResponseBodyLaboratories {
	s.BucketCount = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetBucketType(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.BucketType = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetBuckets(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Buckets = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetCrowdId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.CrowdId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetDebugCrowdId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.DebugCrowdId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetDebugUsers(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.DebugUsers = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetDescription(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Description = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetEnvironment(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Environment = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetFilter(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Filter = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetLaboratoryId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.LaboratoryId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetName(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Name = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetSceneId(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.SceneId = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetStatus(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Status = &v
	return s
}

func (s *ListLaboratoriesResponseBodyLaboratories) SetType(v string) *ListLaboratoriesResponseBodyLaboratories {
	s.Type = &v
	return s
}

type ListLaboratoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLaboratoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLaboratoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLaboratoriesResponse) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesResponse) SetHeaders(v map[string]*string) *ListLaboratoriesResponse {
	s.Headers = v
	return s
}

func (s *ListLaboratoriesResponse) SetStatusCode(v int32) *ListLaboratoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLaboratoriesResponse) SetBody(v *ListLaboratoriesResponseBody) *ListLaboratoriesResponse {
	s.Body = v
	return s
}

type ListLayersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
}

func (s ListLayersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayersRequest) GoString() string {
	return s.String()
}

func (s *ListLayersRequest) SetInstanceId(v string) *ListLayersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLayersRequest) SetLaboratoryId(v string) *ListLayersRequest {
	s.LaboratoryId = &v
	return s
}

type ListLayersResponseBody struct {
	Layers []*ListLayersResponseBodyLayers `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 518C64F6-DFF7-11ED-85B0-00163E14B3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLayersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLayersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayersResponseBody) SetLayers(v []*ListLayersResponseBodyLayers) *ListLayersResponseBody {
	s.Layers = v
	return s
}

func (s *ListLayersResponseBody) SetRequestId(v string) *ListLayersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLayersResponseBody) SetTotalCount(v int64) *ListLayersResponseBody {
	s.TotalCount = &v
	return s
}

type ListLayersResponseBodyLayers struct {
	// example:
	//
	// This is a test.
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// layer1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResidualFlow *int64  `json:"ResidualFlow,omitempty" xml:"ResidualFlow,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListLayersResponseBodyLayers) String() string {
	return tea.Prettify(s)
}

func (s ListLayersResponseBodyLayers) GoString() string {
	return s.String()
}

func (s *ListLayersResponseBodyLayers) SetDescription(v string) *ListLayersResponseBodyLayers {
	s.Description = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetGmtCreateTime(v string) *ListLayersResponseBodyLayers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetLaboratoryId(v string) *ListLayersResponseBodyLayers {
	s.LaboratoryId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetLayerId(v string) *ListLayersResponseBodyLayers {
	s.LayerId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetName(v string) *ListLayersResponseBodyLayers {
	s.Name = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetResidualFlow(v int64) *ListLayersResponseBodyLayers {
	s.ResidualFlow = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetSceneId(v string) *ListLayersResponseBodyLayers {
	s.SceneId = &v
	return s
}

type ListLayersResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLayersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayersResponse) GoString() string {
	return s.String()
}

func (s *ListLayersResponse) SetHeaders(v map[string]*string) *ListLayersResponse {
	s.Headers = v
	return s
}

func (s *ListLayersResponse) SetStatusCode(v int32) *ListLayersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayersResponse) SetBody(v *ListLayersResponseBody) *ListLayersResponse {
	s.Body = v
	return s
}

type ListParamsRequest struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParamsRequest) GoString() string {
	return s.String()
}

func (s *ListParamsRequest) SetEnvironment(v string) *ListParamsRequest {
	s.Environment = &v
	return s
}

func (s *ListParamsRequest) SetInstanceId(v string) *ListParamsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListParamsRequest) SetName(v string) *ListParamsRequest {
	s.Name = &v
	return s
}

func (s *ListParamsRequest) SetPageNumber(v int32) *ListParamsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParamsRequest) SetPageSize(v int32) *ListParamsRequest {
	s.PageSize = &v
	return s
}

func (s *ListParamsRequest) SetSceneId(v string) *ListParamsRequest {
	s.SceneId = &v
	return s
}

type ListParamsResponseBody struct {
	Params []*ListParamsResponseBodyParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A2D07551-38DA-531E-9B22-877D1D86A579
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParamsResponseBody) SetParams(v []*ListParamsResponseBodyParams) *ListParamsResponseBody {
	s.Params = v
	return s
}

func (s *ListParamsResponseBody) SetRequestId(v string) *ListParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParamsResponseBody) SetTotalCount(v int64) *ListParamsResponseBody {
	s.TotalCount = &v
	return s
}

type ListParamsResponseBodyParams struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	ParamId *string `json:"ParamId,omitempty" xml:"ParamId,omitempty"`
	// example:
	//
	// house
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListParamsResponseBodyParams) String() string {
	return tea.Prettify(s)
}

func (s ListParamsResponseBodyParams) GoString() string {
	return s.String()
}

func (s *ListParamsResponseBodyParams) SetEnvironment(v string) *ListParamsResponseBodyParams {
	s.Environment = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetGmtModifiedTime(v string) *ListParamsResponseBodyParams {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetName(v string) *ListParamsResponseBodyParams {
	s.Name = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetParamId(v string) *ListParamsResponseBodyParams {
	s.ParamId = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetValue(v string) *ListParamsResponseBodyParams {
	s.Value = &v
	return s
}

type ListParamsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParamsResponse) GoString() string {
	return s.String()
}

func (s *ListParamsResponse) SetHeaders(v map[string]*string) *ListParamsResponse {
	s.Headers = v
	return s
}

func (s *ListParamsResponse) SetStatusCode(v int32) *ListParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParamsResponse) SetBody(v *ListParamsResponseBody) *ListParamsResponse {
	s.Body = v
	return s
}

type ListResourceRulesRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order            *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber       *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceRuleId   *string `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	ResourceRuleName *string `json:"ResourceRuleName,omitempty" xml:"ResourceRuleName,omitempty"`
	SortBy           *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListResourceRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRulesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceRulesRequest) SetAll(v bool) *ListResourceRulesRequest {
	s.All = &v
	return s
}

func (s *ListResourceRulesRequest) SetInstanceId(v string) *ListResourceRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceRulesRequest) SetOrder(v string) *ListResourceRulesRequest {
	s.Order = &v
	return s
}

func (s *ListResourceRulesRequest) SetPageNumber(v int64) *ListResourceRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceRulesRequest) SetPageSize(v int64) *ListResourceRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceRulesRequest) SetResourceRuleId(v string) *ListResourceRulesRequest {
	s.ResourceRuleId = &v
	return s
}

func (s *ListResourceRulesRequest) SetResourceRuleName(v string) *ListResourceRulesRequest {
	s.ResourceRuleName = &v
	return s
}

func (s *ListResourceRulesRequest) SetSortBy(v string) *ListResourceRulesRequest {
	s.SortBy = &v
	return s
}

type ListResourceRulesResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRules []*ListResourceRulesResponseBodyResourceRules `json:"ResourceRules,omitempty" xml:"ResourceRules,omitempty" type:"Repeated"`
	TotalCount    *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponseBody) SetRequestId(v string) *ListResourceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceRulesResponseBody) SetResourceRules(v []*ListResourceRulesResponseBodyResourceRules) *ListResourceRulesResponseBody {
	s.ResourceRules = v
	return s
}

func (s *ListResourceRulesResponseBody) SetTotalCount(v int64) *ListResourceRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceRulesResponseBodyResourceRules struct {
	Description             *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricOperationType     *string                                                `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo          *string                                                `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod        *string                                                `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	Name                    *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceRuleId          *string                                                `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	RuleComputingDefinition *string                                                `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	RuleItems               []*ListResourceRulesResponseBodyResourceRulesRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s ListResourceRulesResponseBodyResourceRules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRulesResponseBodyResourceRules) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponseBodyResourceRules) SetDescription(v string) *ListResourceRulesResponseBodyResourceRules {
	s.Description = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetMetricOperationType(v string) *ListResourceRulesResponseBodyResourceRules {
	s.MetricOperationType = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetMetricPullInfo(v string) *ListResourceRulesResponseBodyResourceRules {
	s.MetricPullInfo = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetMetricPullPeriod(v string) *ListResourceRulesResponseBodyResourceRules {
	s.MetricPullPeriod = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetName(v string) *ListResourceRulesResponseBodyResourceRules {
	s.Name = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetResourceRuleId(v string) *ListResourceRulesResponseBodyResourceRules {
	s.ResourceRuleId = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetRuleComputingDefinition(v string) *ListResourceRulesResponseBodyResourceRules {
	s.RuleComputingDefinition = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetRuleItems(v []*ListResourceRulesResponseBodyResourceRulesRuleItems) *ListResourceRulesResponseBodyResourceRules {
	s.RuleItems = v
	return s
}

type ListResourceRulesResponseBodyResourceRulesRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxValue    *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue    *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceRulesResponseBodyResourceRulesRuleItems) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRulesResponseBodyResourceRulesRuleItems) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetDescription(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.Description = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetMaxValue(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.MaxValue = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetMinValue(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.MinValue = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetName(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.Name = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetValue(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.Value = &v
	return s
}

type ListResourceRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRulesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponse) SetHeaders(v map[string]*string) *ListResourceRulesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceRulesResponse) SetStatusCode(v int32) *ListResourceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceRulesResponse) SetBody(v *ListResourceRulesResponseBody) *ListResourceRulesResponse {
	s.Body = v
	return s
}

type ListScenesRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenesRequest) GoString() string {
	return s.String()
}

func (s *ListScenesRequest) SetInstanceId(v string) *ListScenesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScenesRequest) SetName(v string) *ListScenesRequest {
	s.Name = &v
	return s
}

type ListScenesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B8987BF7-6028-5B17-80E0-251B7BD67BBA
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenes    []*ListScenesResponseBodyScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBody) SetRequestId(v string) *ListScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenesResponseBody) SetScenes(v []*ListScenesResponseBodyScenes) *ListScenesResponseBody {
	s.Scenes = v
	return s
}

func (s *ListScenesResponseBody) SetTotalCount(v int64) *ListScenesResponseBody {
	s.TotalCount = &v
	return s
}

type ListScenesResponseBodyScenes struct {
	// example:
	//
	// This is a test.
	Description *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*ListScenesResponseBodyScenesFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListScenesResponseBodyScenes) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseBodyScenes) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBodyScenes) SetDescription(v string) *ListScenesResponseBodyScenes {
	s.Description = &v
	return s
}

func (s *ListScenesResponseBodyScenes) SetFlows(v []*ListScenesResponseBodyScenesFlows) *ListScenesResponseBodyScenes {
	s.Flows = v
	return s
}

func (s *ListScenesResponseBodyScenes) SetName(v string) *ListScenesResponseBodyScenes {
	s.Name = &v
	return s
}

func (s *ListScenesResponseBodyScenes) SetSceneId(v string) *ListScenesResponseBodyScenes {
	s.SceneId = &v
	return s
}

type ListScenesResponseBodyScenesFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ListScenesResponseBodyScenesFlows) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseBodyScenesFlows) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBodyScenesFlows) SetFlowCode(v string) *ListScenesResponseBodyScenesFlows {
	s.FlowCode = &v
	return s
}

func (s *ListScenesResponseBodyScenesFlows) SetFlowName(v string) *ListScenesResponseBodyScenesFlows {
	s.FlowName = &v
	return s
}

type ListScenesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponse) GoString() string {
	return s.String()
}

func (s *ListScenesResponse) SetHeaders(v map[string]*string) *ListScenesResponse {
	s.Headers = v
	return s
}

func (s *ListScenesResponse) SetStatusCode(v int32) *ListScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScenesResponse) SetBody(v *ListScenesResponseBody) *ListScenesResponse {
	s.Body = v
	return s
}

type ListSubCrowdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSubCrowdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubCrowdsRequest) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsRequest) SetInstanceId(v string) *ListSubCrowdsRequest {
	s.InstanceId = &v
	return s
}

type ListSubCrowdsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C5AEB79E-FAA4-5DCE-8CD7-1CAF549ECC3E
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCrowds []*ListSubCrowdsResponseBodySubCrowds `json:"SubCrowds,omitempty" xml:"SubCrowds,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubCrowdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubCrowdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsResponseBody) SetRequestId(v string) *ListSubCrowdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubCrowdsResponseBody) SetSubCrowds(v []*ListSubCrowdsResponseBodySubCrowds) *ListSubCrowdsResponseBody {
	s.SubCrowds = v
	return s
}

func (s *ListSubCrowdsResponseBody) SetTotalCount(v int64) *ListSubCrowdsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSubCrowdsResponseBodySubCrowds struct {
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 3
	SubCrowdId *string `json:"SubCrowdId,omitempty" xml:"SubCrowdId,omitempty"`
	// example:
	//
	// user1,user2
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ListSubCrowdsResponseBodySubCrowds) String() string {
	return tea.Prettify(s)
}

func (s ListSubCrowdsResponseBodySubCrowds) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetGmtCreateTime(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetQuantity(v int32) *ListSubCrowdsResponseBodySubCrowds {
	s.Quantity = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetSource(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.Source = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetSubCrowdId(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.SubCrowdId = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetUsers(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.Users = &v
	return s
}

type ListSubCrowdsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubCrowdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubCrowdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubCrowdsResponse) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsResponse) SetHeaders(v map[string]*string) *ListSubCrowdsResponse {
	s.Headers = v
	return s
}

func (s *ListSubCrowdsResponse) SetStatusCode(v int32) *ListSubCrowdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubCrowdsResponse) SetBody(v *ListSubCrowdsResponseBody) *ListSubCrowdsResponse {
	s.Body = v
	return s
}

type ListTableMetasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTableMetasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasRequest) GoString() string {
	return s.String()
}

func (s *ListTableMetasRequest) SetInstanceId(v string) *ListTableMetasRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTableMetasRequest) SetModule(v string) *ListTableMetasRequest {
	s.Module = &v
	return s
}

func (s *ListTableMetasRequest) SetName(v string) *ListTableMetasRequest {
	s.Name = &v
	return s
}

func (s *ListTableMetasRequest) SetPageNumber(v int32) *ListTableMetasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTableMetasRequest) SetPageSize(v int32) *ListTableMetasRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableMetasRequest) SetType(v string) *ListTableMetasRequest {
	s.Type = &v
	return s
}

type ListTableMetasResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TableMetas []*ListTableMetasResponseBodyTableMetas `json:"TableMetas,omitempty" xml:"TableMetas,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTableMetasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponseBody) SetRequestId(v string) *ListTableMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableMetasResponseBody) SetTableMetas(v []*ListTableMetasResponseBodyTableMetas) *ListTableMetasResponseBody {
	s.TableMetas = v
	return s
}

func (s *ListTableMetasResponseBody) SetTotalCount(v int64) *ListTableMetasResponseBody {
	s.TotalCount = &v
	return s
}

type ListTableMetasResponseBodyTableMetas struct {
	// example:
	//
	// true
	CanDelete *bool   `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// this is a test table
	Description *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields      []*ListTableMetasResponseBodyTableMetasFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-12 12:24:33
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// imprecation
	GmtImportedTime *string `json:"GmtImportedTime,omitempty" xml:"GmtImportedTime,omitempty"`
	// example:
	//
	// 2021-12-12 12:24:33
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// test_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 3
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// table_mysql
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://dmc-xxx.com/dm/table/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListTableMetasResponseBodyTableMetas) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasResponseBodyTableMetas) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponseBodyTableMetas) SetCanDelete(v bool) *ListTableMetasResponseBodyTableMetas {
	s.CanDelete = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetConfig(v string) *ListTableMetasResponseBodyTableMetas {
	s.Config = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetDescription(v string) *ListTableMetasResponseBodyTableMetas {
	s.Description = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetFields(v []*ListTableMetasResponseBodyTableMetasFields) *ListTableMetasResponseBodyTableMetas {
	s.Fields = v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetGmtCreateTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetGmtImportedTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtImportedTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetGmtModifiedTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetModule(v string) *ListTableMetasResponseBodyTableMetas {
	s.Module = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetName(v string) *ListTableMetasResponseBodyTableMetas {
	s.Name = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetResourceId(v string) *ListTableMetasResponseBodyTableMetas {
	s.ResourceId = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetTableMetaId(v string) *ListTableMetasResponseBodyTableMetas {
	s.TableMetaId = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetTableName(v string) *ListTableMetasResponseBodyTableMetas {
	s.TableName = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetType(v string) *ListTableMetasResponseBodyTableMetas {
	s.Type = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetUrl(v string) *ListTableMetasResponseBodyTableMetas {
	s.Url = &v
	return s
}

type ListTableMetasResponseBodyTableMetasFields struct {
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	// example:
	//
	// the gender of people
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// example:
	//
	// gender
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTableMetasResponseBodyTableMetasFields) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasResponseBodyTableMetasFields) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetIsDimensionField(v bool) *ListTableMetasResponseBodyTableMetasFields {
	s.IsDimensionField = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetMeaning(v string) *ListTableMetasResponseBodyTableMetasFields {
	s.Meaning = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetName(v string) *ListTableMetasResponseBodyTableMetasFields {
	s.Name = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetType(v string) *ListTableMetasResponseBodyTableMetasFields {
	s.Type = &v
	return s
}

type ListTableMetasResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableMetasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasResponse) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponse) SetHeaders(v map[string]*string) *ListTableMetasResponse {
	s.Headers = v
	return s
}

func (s *ListTableMetasResponse) SetStatusCode(v int32) *ListTableMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableMetasResponse) SetBody(v *ListTableMetasResponseBody) *ListTableMetasResponse {
	s.Body = v
	return s
}

type ListTrafficControlTargetTrafficHistoryRequest struct {
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Environment       *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	ExperimentId      *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold         *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListTrafficControlTargetTrafficHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetEndTime(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetEnvironment(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.Environment = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetExperimentGroupId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetExperimentId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetInstanceId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetItemId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.ItemId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetStartTime(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetThreshold(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.Threshold = &v
	return s
}

type ListTrafficControlTargetTrafficHistoryResponseBody struct {
	RequestId                          *string                                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                         *string                                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficControlTaskTrafficHistories []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories `json:"TrafficControlTaskTrafficHistories,omitempty" xml:"TrafficControlTaskTrafficHistories,omitempty" type:"Repeated"`
}

func (s ListTrafficControlTargetTrafficHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) SetRequestId(v string) *ListTrafficControlTargetTrafficHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) SetTotalCount(v string) *ListTrafficControlTargetTrafficHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) SetTrafficControlTaskTrafficHistories(v []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) *ListTrafficControlTargetTrafficHistoryResponseBody {
	s.TrafficControlTaskTrafficHistories = v
	return s
}

type ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories struct {
	ExperimentId                   *string  `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ItemId                         *string  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	RecordTime                     *string  `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
	TrafficControlTargetAimTraffic *float64 `json:"TrafficControlTargetAimTraffic,omitempty" xml:"TrafficControlTargetAimTraffic,omitempty"`
	TrafficControlTargetTraffic    *float64 `json:"TrafficControlTargetTraffic,omitempty" xml:"TrafficControlTargetTraffic,omitempty"`
	TrafficControlTaskTraffic      *float64 `json:"TrafficControlTaskTraffic,omitempty" xml:"TrafficControlTaskTraffic,omitempty"`
}

func (s ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetExperimentId(v string) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.ExperimentId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetItemId(v string) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.ItemId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetRecordTime(v string) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.RecordTime = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetTrafficControlTargetAimTraffic(v float64) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.TrafficControlTargetAimTraffic = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetTrafficControlTargetTraffic(v float64) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.TrafficControlTargetTraffic = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetTrafficControlTaskTraffic(v float64) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.TrafficControlTaskTraffic = &v
	return s
}

type ListTrafficControlTargetTrafficHistoryResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficControlTargetTrafficHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficControlTargetTrafficHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) SetHeaders(v map[string]*string) *ListTrafficControlTargetTrafficHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) SetStatusCode(v int32) *ListTrafficControlTargetTrafficHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) SetBody(v *ListTrafficControlTargetTrafficHistoryResponseBody) *ListTrafficControlTargetTrafficHistoryResponse {
	s.Body = v
	return s
}

type ListTrafficControlTasksRequest struct {
	All                  *bool   `json:"All,omitempty" xml:"All,omitempty"`
	ControlTargetFilter  *string `json:"ControlTargetFilter,omitempty" xml:"ControlTargetFilter,omitempty"`
	Environment          *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId              *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SortBy               *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListTrafficControlTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksRequest) SetAll(v bool) *ListTrafficControlTasksRequest {
	s.All = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetControlTargetFilter(v string) *ListTrafficControlTasksRequest {
	s.ControlTargetFilter = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetEnvironment(v string) *ListTrafficControlTasksRequest {
	s.Environment = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetInstanceId(v string) *ListTrafficControlTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetName(v string) *ListTrafficControlTasksRequest {
	s.Name = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetOrder(v string) *ListTrafficControlTasksRequest {
	s.Order = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetPageNumber(v string) *ListTrafficControlTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetPageSize(v string) *ListTrafficControlTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetSceneId(v string) *ListTrafficControlTasksRequest {
	s.SceneId = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetSortBy(v string) *ListTrafficControlTasksRequest {
	s.SortBy = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetStatus(v string) *ListTrafficControlTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetTrafficControlTaskId(v string) *ListTrafficControlTasksRequest {
	s.TrafficControlTaskId = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetVersion(v string) *ListTrafficControlTasksRequest {
	s.Version = &v
	return s
}

type ListTrafficControlTasksResponseBody struct {
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *string                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficControlTasks []*ListTrafficControlTasksResponseBodyTrafficControlTasks `json:"TrafficControlTasks,omitempty" xml:"TrafficControlTasks,omitempty" type:"Repeated"`
}

func (s ListTrafficControlTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBody) SetRequestId(v string) *ListTrafficControlTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBody) SetTotalCount(v string) *ListTrafficControlTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficControlTasksResponseBody) SetTrafficControlTasks(v []*ListTrafficControlTasksResponseBodyTrafficControlTasks) *ListTrafficControlTasksResponseBody {
	s.TrafficControlTasks = v
	return s
}

type ListTrafficControlTasksResponseBodyTrafficControlTasks struct {
	BehaviorTableMetaId            *string                                                                        `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	ControlGranularity             *string                                                                        `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	ControlLogic                   *string                                                                        `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	ControlType                    *string                                                                        `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	Description                    *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                        *string                                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EverPublished                  *bool                                                                          `json:"EverPublished,omitempty" xml:"EverPublished,omitempty"`
	ExecutionTime                  *string                                                                        `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	GmtCreateTime                  *string                                                                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime                *string                                                                        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemConditionArray             *string                                                                        `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress           *string                                                                        `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType              *string                                                                        `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	ItemTableMetaId                *string                                                                        `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	Name                           *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds               *string                                                                        `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	PrepubStatus                   *string                                                                        `json:"PrepubStatus,omitempty" xml:"PrepubStatus,omitempty"`
	ProdExperimentIds              *string                                                                        `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	ProductStatus                  *string                                                                        `json:"ProductStatus,omitempty" xml:"ProductStatus,omitempty"`
	SceneId                        *string                                                                        `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName                      *string                                                                        `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceId                      *string                                                                        `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	StartTime                      *string                                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisBahaviorConditionExpress *string                                                                        `json:"StatisBahaviorConditionExpress,omitempty" xml:"StatisBahaviorConditionExpress,omitempty"`
	StatisBehaviorConditionArray   *string                                                                        `json:"StatisBehaviorConditionArray,omitempty" xml:"StatisBehaviorConditionArray,omitempty"`
	StatisBehaviorConditionType    *string                                                                        `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets          []*ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	TrafficControlTaskId           *string                                                                        `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	UserConditionArray             *string                                                                        `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	UserConditionExpress           *string                                                                        `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	UserConditionType              *string                                                                        `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	UserTableMetaId                *string                                                                        `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasks) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetBehaviorTableMetaId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetControlGranularity(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ControlGranularity = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetControlLogic(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ControlLogic = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetControlType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ControlType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetDescription(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.Description = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEndTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EndTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEverPublished(v bool) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EverPublished = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetExecutionTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ExecutionTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetGmtCreateTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetGmtModifiedTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemTableMetaId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemTableMetaId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.Name = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetPreExperimentIds(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.PreExperimentIds = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetPrepubStatus(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.PrepubStatus = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetProdExperimentIds(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ProdExperimentIds = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetProductStatus(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ProductStatus = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetSceneId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.SceneId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetSceneName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.SceneName = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetServiceId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ServiceId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStartTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StartTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBahaviorConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBahaviorConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBehaviorConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBehaviorConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBehaviorConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetTrafficControlTargets(v []*ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.TrafficControlTargets = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetTrafficControlTaskId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.TrafficControlTaskId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserTableMetaId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserTableMetaId = &v
	return s
}

type ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets struct {
	EndTime                *string                                                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                  *string                                                                                `json:"Event,omitempty" xml:"Event,omitempty"`
	GmtCreateTime          *string                                                                                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string                                                                                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemConditionArray     *string                                                                                `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress   *string                                                                                `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType      *string                                                                                `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                   *string                                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation   *bool                                                                                  `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName             *string                                                                                `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	SplitParts             *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	StartTime              *string                                                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod           *string                                                                                `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status                 *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue         *int64                                                                                 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	TrafficControlTargetId *string                                                                                `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	Value                  *float32                                                                               `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetEndTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetEvent(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetGmtCreateTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetGmtModifiedTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetItemConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetItemConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetItemConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetNewProductRegulation(v bool) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetRecallName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetSplitParts(v *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.SplitParts = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetStartTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetStatisPeriod(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetStatus(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetToleranceValue(v int64) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetTrafficControlTargetId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.TrafficControlTargetId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetValue(v float32) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Value = &v
	return s
}

type ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts struct {
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints []*int64 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) SetSetValues(v []*int64) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts {
	s.SetValues = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) SetTimePoints(v []*int64) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts {
	s.TimePoints = v
	return s
}

type ListTrafficControlTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficControlTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficControlTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficControlTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponse) SetHeaders(v map[string]*string) *ListTrafficControlTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficControlTasksResponse) SetStatusCode(v int32) *ListTrafficControlTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficControlTasksResponse) SetBody(v *ListTrafficControlTasksResponseBody) *ListTrafficControlTasksResponse {
	s.Body = v
	return s
}

type OfflineExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineExperimentRequest) GoString() string {
	return s.String()
}

func (s *OfflineExperimentRequest) SetInstanceId(v string) *OfflineExperimentRequest {
	s.InstanceId = &v
	return s
}

type OfflineExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 872951C9-7755-5FA1-AACD-7F9375A6D27A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineExperimentResponseBody) SetRequestId(v string) *OfflineExperimentResponseBody {
	s.RequestId = &v
	return s
}

type OfflineExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineExperimentResponse) GoString() string {
	return s.String()
}

func (s *OfflineExperimentResponse) SetHeaders(v map[string]*string) *OfflineExperimentResponse {
	s.Headers = v
	return s
}

func (s *OfflineExperimentResponse) SetStatusCode(v int32) *OfflineExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineExperimentResponse) SetBody(v *OfflineExperimentResponseBody) *OfflineExperimentResponse {
	s.Body = v
	return s
}

type OfflineExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *OfflineExperimentGroupRequest) SetInstanceId(v string) *OfflineExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

type OfflineExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 518C64F6-DFF7-11ED-85B0-00163E14B3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineExperimentGroupResponseBody) SetRequestId(v string) *OfflineExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

type OfflineExperimentGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *OfflineExperimentGroupResponse) SetHeaders(v map[string]*string) *OfflineExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *OfflineExperimentGroupResponse) SetStatusCode(v int32) *OfflineExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineExperimentGroupResponse) SetBody(v *OfflineExperimentGroupResponseBody) *OfflineExperimentGroupResponse {
	s.Body = v
	return s
}

type OfflineLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *OfflineLaboratoryRequest) SetInstanceId(v string) *OfflineLaboratoryRequest {
	s.InstanceId = &v
	return s
}

type OfflineLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineLaboratoryResponseBody) SetRequestId(v string) *OfflineLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

type OfflineLaboratoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *OfflineLaboratoryResponse) SetHeaders(v map[string]*string) *OfflineLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *OfflineLaboratoryResponse) SetStatusCode(v int32) *OfflineLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineLaboratoryResponse) SetBody(v *OfflineLaboratoryResponseBody) *OfflineLaboratoryResponse {
	s.Body = v
	return s
}

type OnlineExperimentRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s OnlineExperimentRequest) GoString() string {
	return s.String()
}

func (s *OnlineExperimentRequest) SetInstanceId(v string) *OnlineExperimentRequest {
	s.InstanceId = &v
	return s
}

type OnlineExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnlineExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineExperimentResponseBody) SetRequestId(v string) *OnlineExperimentResponseBody {
	s.RequestId = &v
	return s
}

type OnlineExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s OnlineExperimentResponse) GoString() string {
	return s.String()
}

func (s *OnlineExperimentResponse) SetHeaders(v map[string]*string) *OnlineExperimentResponse {
	s.Headers = v
	return s
}

func (s *OnlineExperimentResponse) SetStatusCode(v int32) *OnlineExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineExperimentResponse) SetBody(v *OnlineExperimentResponseBody) *OnlineExperimentResponse {
	s.Body = v
	return s
}

type OnlineExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s OnlineExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *OnlineExperimentGroupRequest) SetInstanceId(v string) *OnlineExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

type OnlineExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 47F761ED-BE4E-51A6-B678-78E1490DF313
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnlineExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineExperimentGroupResponseBody) SetRequestId(v string) *OnlineExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

type OnlineExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s OnlineExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *OnlineExperimentGroupResponse) SetHeaders(v map[string]*string) *OnlineExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *OnlineExperimentGroupResponse) SetStatusCode(v int32) *OnlineExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineExperimentGroupResponse) SetBody(v *OnlineExperimentGroupResponseBody) *OnlineExperimentGroupResponse {
	s.Body = v
	return s
}

type OnlineLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s OnlineLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *OnlineLaboratoryRequest) SetInstanceId(v string) *OnlineLaboratoryRequest {
	s.InstanceId = &v
	return s
}

type OnlineLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8C27790E-CCA5-56BB-BA17-646295DEC0A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnlineLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineLaboratoryResponseBody) SetRequestId(v string) *OnlineLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

type OnlineLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s OnlineLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *OnlineLaboratoryResponse) SetHeaders(v map[string]*string) *OnlineLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *OnlineLaboratoryResponse) SetStatusCode(v int32) *OnlineLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineLaboratoryResponse) SetBody(v *OnlineLaboratoryResponseBody) *OnlineLaboratoryResponse {
	s.Body = v
	return s
}

type PushAllExperimentRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PushAllExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s PushAllExperimentRequest) GoString() string {
	return s.String()
}

func (s *PushAllExperimentRequest) SetInstanceId(v string) *PushAllExperimentRequest {
	s.InstanceId = &v
	return s
}

type PushAllExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushAllExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushAllExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *PushAllExperimentResponseBody) SetRequestId(v string) *PushAllExperimentResponseBody {
	s.RequestId = &v
	return s
}

type PushAllExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushAllExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushAllExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s PushAllExperimentResponse) GoString() string {
	return s.String()
}

func (s *PushAllExperimentResponse) SetHeaders(v map[string]*string) *PushAllExperimentResponse {
	s.Headers = v
	return s
}

func (s *PushAllExperimentResponse) SetStatusCode(v int32) *PushAllExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *PushAllExperimentResponse) SetBody(v *PushAllExperimentResponseBody) *PushAllExperimentResponse {
	s.Body = v
	return s
}

type PushResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfo map[string]interface{} `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
}

func (s PushResourceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PushResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *PushResourceRuleRequest) SetInstanceId(v string) *PushResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *PushResourceRuleRequest) SetMetricInfo(v map[string]interface{}) *PushResourceRuleRequest {
	s.MetricInfo = v
	return s
}

type PushResourceRuleShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfoShrink *string `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
}

func (s PushResourceRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushResourceRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushResourceRuleShrinkRequest) SetInstanceId(v string) *PushResourceRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *PushResourceRuleShrinkRequest) SetMetricInfoShrink(v string) *PushResourceRuleShrinkRequest {
	s.MetricInfoShrink = &v
	return s
}

type PushResourceRuleResponseBody struct {
	Description             *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricOperationType     *string                                  `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo          *string                                  `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod        *string                                  `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	Name                    *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId               *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleId          *string                                  `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	RuleComputingDefinition *string                                  `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	RuleItems               []*PushResourceRuleResponseBodyRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s PushResourceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PushResourceRuleResponseBody) SetDescription(v string) *PushResourceRuleResponseBody {
	s.Description = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetMetricOperationType(v string) *PushResourceRuleResponseBody {
	s.MetricOperationType = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetMetricPullInfo(v string) *PushResourceRuleResponseBody {
	s.MetricPullInfo = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetMetricPullPeriod(v string) *PushResourceRuleResponseBody {
	s.MetricPullPeriod = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetName(v string) *PushResourceRuleResponseBody {
	s.Name = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetRequestId(v string) *PushResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetResourceRuleId(v string) *PushResourceRuleResponseBody {
	s.ResourceRuleId = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetRuleComputingDefinition(v string) *PushResourceRuleResponseBody {
	s.RuleComputingDefinition = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetRuleItems(v []*PushResourceRuleResponseBodyRuleItems) *PushResourceRuleResponseBody {
	s.RuleItems = v
	return s
}

type PushResourceRuleResponseBodyRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxValue    *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue    *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PushResourceRuleResponseBodyRuleItems) String() string {
	return tea.Prettify(s)
}

func (s PushResourceRuleResponseBodyRuleItems) GoString() string {
	return s.String()
}

func (s *PushResourceRuleResponseBodyRuleItems) SetDescription(v string) *PushResourceRuleResponseBodyRuleItems {
	s.Description = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetMaxValue(v string) *PushResourceRuleResponseBodyRuleItems {
	s.MaxValue = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetMinValue(v string) *PushResourceRuleResponseBodyRuleItems {
	s.MinValue = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetName(v string) *PushResourceRuleResponseBodyRuleItems {
	s.Name = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetValue(v string) *PushResourceRuleResponseBodyRuleItems {
	s.Value = &v
	return s
}

type PushResourceRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushResourceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PushResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *PushResourceRuleResponse) SetHeaders(v map[string]*string) *PushResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *PushResourceRuleResponse) SetStatusCode(v int32) *PushResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResourceRuleResponse) SetBody(v *PushResourceRuleResponseBody) *PushResourceRuleResponse {
	s.Body = v
	return s
}

type QueryTrafficControlTargetItemReportDetailRequest struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// This parameter is required.
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) SetDate(v string) *QueryTrafficControlTargetItemReportDetailRequest {
	s.Date = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) SetEnvironment(v string) *QueryTrafficControlTargetItemReportDetailRequest {
	s.Environment = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) SetInstanceId(v string) *QueryTrafficControlTargetItemReportDetailRequest {
	s.InstanceId = &v
	return s
}

type QueryTrafficControlTargetItemReportDetailResponseBody struct {
	RequestId                            *string                                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTargetItemReportDetail *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail `json:"TrafficControlTargetItemReportDetail,omitempty" xml:"TrafficControlTargetItemReportDetail,omitempty" type:"Struct"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) SetRequestId(v string) *QueryTrafficControlTargetItemReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) SetTrafficControlTargetItemReportDetail(v *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) *QueryTrafficControlTargetItemReportDetailResponseBody {
	s.TrafficControlTargetItemReportDetail = v
	return s
}

type QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail struct {
	ItemControlTailReportDetails []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails `json:"ItemControlTailReportDetails,omitempty" xml:"ItemControlTailReportDetails,omitempty" type:"Repeated"`
	ItemControlTopReportDetails  []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails  `json:"ItemControlTopReportDetails,omitempty" xml:"ItemControlTopReportDetails,omitempty" type:"Repeated"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) SetItemControlTailReportDetails(v []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail {
	s.ItemControlTailReportDetails = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) SetItemControlTopReportDetails(v []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail {
	s.ItemControlTopReportDetails = v
	return s
}

type QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails struct {
	Features       map[string]interface{} `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemId         *string                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	TargetProgress *string                `json:"TargetProgress,omitempty" xml:"TargetProgress,omitempty"`
	TargetTraffic  *int64                 `json:"TargetTraffic,omitempty" xml:"TargetTraffic,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetFeatures(v map[string]interface{}) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.Features = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetItemId(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.ItemId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetTargetProgress(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.TargetProgress = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetTargetTraffic(v int64) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.TargetTraffic = &v
	return s
}

type QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails struct {
	Features       map[string]interface{} `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemId         *string                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	TargetProgress *string                `json:"TargetProgress,omitempty" xml:"TargetProgress,omitempty"`
	TargetTraffic  *int64                 `json:"TargetTraffic,omitempty" xml:"TargetTraffic,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetFeatures(v map[string]interface{}) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.Features = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetItemId(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.ItemId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetTargetProgress(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.TargetProgress = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetTargetTraffic(v int64) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.TargetTraffic = &v
	return s
}

type QueryTrafficControlTargetItemReportDetailResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTrafficControlTargetItemReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) SetHeaders(v map[string]*string) *QueryTrafficControlTargetItemReportDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) SetStatusCode(v int32) *QueryTrafficControlTargetItemReportDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) SetBody(v *QueryTrafficControlTargetItemReportDetailResponseBody) *QueryTrafficControlTargetItemReportDetailResponse {
	s.Body = v
	return s
}

type ReleaseTrafficControlTaskRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *ReleaseTrafficControlTaskRequest) SetEnvironment(v string) *ReleaseTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *ReleaseTrafficControlTaskRequest) SetInstanceId(v string) *ReleaseTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

type ReleaseTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseTrafficControlTaskResponseBody) SetRequestId(v string) *ReleaseTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseTrafficControlTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *ReleaseTrafficControlTaskResponse) SetHeaders(v map[string]*string) *ReleaseTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *ReleaseTrafficControlTaskResponse) SetStatusCode(v int32) *ReleaseTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseTrafficControlTaskResponse) SetBody(v *ReleaseTrafficControlTaskResponseBody) *ReleaseTrafficControlTaskResponse {
	s.Body = v
	return s
}

type ReportABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3
	BaseExperimentId *string `json:"BaseExperimentId,omitempty" xml:"BaseExperimentId,omitempty"`
	// example:
	//
	// {"gender":"man"}
	DimensionFields *string `json:"DimensionFields,omitempty" xml:"DimensionFields,omitempty"`
	// example:
	//
	// 2021-07-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3,4,5
	ExperimentIds *string `json:"ExperimentIds,omitempty" xml:"ExperimentIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Offline
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 2021-07-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// Hour
	TimeStatisticsMethod *string `json:"TimeStatisticsMethod,omitempty" xml:"TimeStatisticsMethod,omitempty"`
}

func (s ReportABMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *ReportABMetricGroupRequest) SetBaseExperimentId(v string) *ReportABMetricGroupRequest {
	s.BaseExperimentId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetDimensionFields(v string) *ReportABMetricGroupRequest {
	s.DimensionFields = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetEndDate(v string) *ReportABMetricGroupRequest {
	s.EndDate = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetExperimentGroupId(v string) *ReportABMetricGroupRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetExperimentIds(v string) *ReportABMetricGroupRequest {
	s.ExperimentIds = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetInstanceId(v string) *ReportABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetReportType(v string) *ReportABMetricGroupRequest {
	s.ReportType = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetSceneId(v string) *ReportABMetricGroupRequest {
	s.SceneId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetStartDate(v string) *ReportABMetricGroupRequest {
	s.StartDate = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetTimeStatisticsMethod(v string) *ReportABMetricGroupRequest {
	s.TimeStatisticsMethod = &v
	return s
}

type ReportABMetricGroupResponseBody struct {
	ExperimentReport map[string]*ExperimentReportValue `json:"ExperimentReport,omitempty" xml:"ExperimentReport,omitempty"`
	GroupDimension   []*string                         `json:"GroupDimension,omitempty" xml:"GroupDimension,omitempty" type:"Repeated"`
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportABMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReportABMetricGroupResponseBody) SetExperimentReport(v map[string]*ExperimentReportValue) *ReportABMetricGroupResponseBody {
	s.ExperimentReport = v
	return s
}

func (s *ReportABMetricGroupResponseBody) SetGroupDimension(v []*string) *ReportABMetricGroupResponseBody {
	s.GroupDimension = v
	return s
}

func (s *ReportABMetricGroupResponseBody) SetRequestId(v string) *ReportABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type ReportABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportABMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *ReportABMetricGroupResponse) SetHeaders(v map[string]*string) *ReportABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *ReportABMetricGroupResponse) SetStatusCode(v int32) *ReportABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportABMetricGroupResponse) SetBody(v *ReportABMetricGroupResponseBody) *ReportABMetricGroupResponse {
	s.Body = v
	return s
}

type SplitTrafficControlTargetRequest struct {
	Environment *string  `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SetPoints   []*int64 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	SetValues   []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints  []*int64 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s SplitTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *SplitTrafficControlTargetRequest) SetEnvironment(v string) *SplitTrafficControlTargetRequest {
	s.Environment = &v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetInstanceId(v string) *SplitTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetSetPoints(v []*int64) *SplitTrafficControlTargetRequest {
	s.SetPoints = v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetSetValues(v []*int64) *SplitTrafficControlTargetRequest {
	s.SetValues = v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetTimePoints(v []*int64) *SplitTrafficControlTargetRequest {
	s.TimePoints = v
	return s
}

type SplitTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SplitTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SplitTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *SplitTrafficControlTargetResponseBody) SetRequestId(v string) *SplitTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

type SplitTrafficControlTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SplitTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s SplitTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *SplitTrafficControlTargetResponse) SetHeaders(v map[string]*string) *SplitTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *SplitTrafficControlTargetResponse) SetStatusCode(v int32) *SplitTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitTrafficControlTargetResponse) SetBody(v *SplitTrafficControlTargetResponseBody) *SplitTrafficControlTargetResponse {
	s.Body = v
	return s
}

type StartTrafficControlTargetRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTargetRequest) SetInstanceId(v string) *StartTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

type StartTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTargetResponseBody) SetRequestId(v string) *StartTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

type StartTrafficControlTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTargetResponse) SetHeaders(v map[string]*string) *StartTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *StartTrafficControlTargetResponse) SetStatusCode(v int32) *StartTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTrafficControlTargetResponse) SetBody(v *StartTrafficControlTargetResponseBody) *StartTrafficControlTargetResponse {
	s.Body = v
	return s
}

type StartTrafficControlTaskRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTaskRequest) SetEnvironment(v string) *StartTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *StartTrafficControlTaskRequest) SetInstanceId(v string) *StartTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

type StartTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTaskResponseBody) SetRequestId(v string) *StartTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartTrafficControlTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTaskResponse) SetHeaders(v map[string]*string) *StartTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTrafficControlTaskResponse) SetStatusCode(v int32) *StartTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTrafficControlTaskResponse) SetBody(v *StartTrafficControlTaskResponseBody) *StartTrafficControlTaskResponse {
	s.Body = v
	return s
}

type StopTrafficControlTargetRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTargetRequest) SetInstanceId(v string) *StopTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

type StopTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTargetResponseBody) SetRequestId(v string) *StopTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

type StopTrafficControlTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTargetResponse) SetHeaders(v map[string]*string) *StopTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *StopTrafficControlTargetResponse) SetStatusCode(v int32) *StopTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrafficControlTargetResponse) SetBody(v *StopTrafficControlTargetResponseBody) *StopTrafficControlTargetResponse {
	s.Body = v
	return s
}

type StopTrafficControlTaskRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTaskRequest) SetRegionId(v string) *StopTrafficControlTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopTrafficControlTaskRequest) SetEnvironment(v string) *StopTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *StopTrafficControlTaskRequest) SetInstanceId(v string) *StopTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

type StopTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTaskResponseBody) SetRequestId(v string) *StopTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopTrafficControlTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTaskResponse) SetHeaders(v map[string]*string) *StopTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTrafficControlTaskResponse) SetStatusCode(v int32) *StopTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrafficControlTaskResponse) SetBody(v *StopTrafficControlTaskResponseBody) *StopTrafficControlTaskResponse {
	s.Body = v
	return s
}

type SyncFeatureConsistencyCheckJobReplayLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{\\"Value\\":{\\"FloatFeature\\":0.1}}]
	ContextFeatures *string `json:"ContextFeatures,omitempty" xml:"ContextFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// week_day:1 | userid:3 | itemid:9001 | cate:cat1 | click_5_seq__cate:cat1
	GeneratedFeatures *string `json:"GeneratedFeatures,omitempty" xml:"GeneratedFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1693900981465
	LogRequestTime *int64 `json:"LogRequestTime,omitempty" xml:"LogRequestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// This parameter is required.
	RawFeatures *string `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video-feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s SyncFeatureConsistencyCheckJobReplayLogRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncFeatureConsistencyCheckJobReplayLogRequest) GoString() string {
	return s.String()
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetContextFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.ContextFeatures = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetFeatureConsistencyCheckJobConfigId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetGeneratedFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.GeneratedFeatures = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetInstanceId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogItemId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogItemId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogRequestId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogRequestId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogRequestTime(v int64) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogRequestTime = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetLogUserId(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.LogUserId = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetRawFeatures(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.RawFeatures = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogRequest) SetSceneName(v string) *SyncFeatureConsistencyCheckJobReplayLogRequest {
	s.SceneName = &v
	return s
}

type SyncFeatureConsistencyCheckJobReplayLogResponseBody struct {
	// example:
	//
	// C7D0B48F-0105-52B9-B60A-FA7606E2234D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponseBody) GoString() string {
	return s.String()
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponseBody) SetRequestId(v string) *SyncFeatureConsistencyCheckJobReplayLogResponseBody {
	s.RequestId = &v
	return s
}

type SyncFeatureConsistencyCheckJobReplayLogResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncFeatureConsistencyCheckJobReplayLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncFeatureConsistencyCheckJobReplayLogResponse) GoString() string {
	return s.String()
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) SetHeaders(v map[string]*string) *SyncFeatureConsistencyCheckJobReplayLogResponse {
	s.Headers = v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) SetStatusCode(v int32) *SyncFeatureConsistencyCheckJobReplayLogResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncFeatureConsistencyCheckJobReplayLogResponse) SetBody(v *SyncFeatureConsistencyCheckJobReplayLogResponseBody) *SyncFeatureConsistencyCheckJobReplayLogResponse {
	s.Body = v
	return s
}

type TerminateFeatureConsistencyCheckJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s TerminateFeatureConsistencyCheckJobRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateFeatureConsistencyCheckJobRequest) GoString() string {
	return s.String()
}

func (s *TerminateFeatureConsistencyCheckJobRequest) SetInstanceId(v string) *TerminateFeatureConsistencyCheckJobRequest {
	s.InstanceId = &v
	return s
}

type TerminateFeatureConsistencyCheckJobResponseBody struct {
	// example:
	//
	// A6C01890-54CA-5C49-BC91-AD85A98E4A98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateFeatureConsistencyCheckJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateFeatureConsistencyCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateFeatureConsistencyCheckJobResponseBody) SetRequestId(v string) *TerminateFeatureConsistencyCheckJobResponseBody {
	s.RequestId = &v
	return s
}

type TerminateFeatureConsistencyCheckJobResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateFeatureConsistencyCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateFeatureConsistencyCheckJobResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateFeatureConsistencyCheckJobResponse) GoString() string {
	return s.String()
}

func (s *TerminateFeatureConsistencyCheckJobResponse) SetHeaders(v map[string]*string) *TerminateFeatureConsistencyCheckJobResponse {
	s.Headers = v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobResponse) SetStatusCode(v int32) *TerminateFeatureConsistencyCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobResponse) SetBody(v *TerminateFeatureConsistencyCheckJobResponseBody) *TerminateFeatureConsistencyCheckJobResponse {
	s.Body = v
	return s
}

type UpdateABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sum(click_cnt)
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 3
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateABMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateABMetricRequest) GoString() string {
	return s.String()
}

func (s *UpdateABMetricRequest) SetDefinition(v string) *UpdateABMetricRequest {
	s.Definition = &v
	return s
}

func (s *UpdateABMetricRequest) SetDescription(v string) *UpdateABMetricRequest {
	s.Description = &v
	return s
}

func (s *UpdateABMetricRequest) SetInstanceId(v string) *UpdateABMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateABMetricRequest) SetLeftMetricId(v string) *UpdateABMetricRequest {
	s.LeftMetricId = &v
	return s
}

func (s *UpdateABMetricRequest) SetName(v string) *UpdateABMetricRequest {
	s.Name = &v
	return s
}

func (s *UpdateABMetricRequest) SetOperator(v string) *UpdateABMetricRequest {
	s.Operator = &v
	return s
}

func (s *UpdateABMetricRequest) SetRealtime(v bool) *UpdateABMetricRequest {
	s.Realtime = &v
	return s
}

func (s *UpdateABMetricRequest) SetResultResourceId(v string) *UpdateABMetricRequest {
	s.ResultResourceId = &v
	return s
}

func (s *UpdateABMetricRequest) SetRightMetricId(v string) *UpdateABMetricRequest {
	s.RightMetricId = &v
	return s
}

func (s *UpdateABMetricRequest) SetSceneId(v string) *UpdateABMetricRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateABMetricRequest) SetStatisticsCycle(v int32) *UpdateABMetricRequest {
	s.StatisticsCycle = &v
	return s
}

func (s *UpdateABMetricRequest) SetTableMetaId(v string) *UpdateABMetricRequest {
	s.TableMetaId = &v
	return s
}

func (s *UpdateABMetricRequest) SetType(v string) *UpdateABMetricRequest {
	s.Type = &v
	return s
}

type UpdateABMetricResponseBody struct {
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateABMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABMetricResponseBody) SetRequestId(v string) *UpdateABMetricResponseBody {
	s.RequestId = &v
	return s
}

type UpdateABMetricResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABMetricResponse) GoString() string {
	return s.String()
}

func (s *UpdateABMetricResponse) SetHeaders(v map[string]*string) *UpdateABMetricResponse {
	s.Headers = v
	return s
}

func (s *UpdateABMetricResponse) SetStatusCode(v int32) *UpdateABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABMetricResponse) SetBody(v *UpdateABMetricResponseBody) *UpdateABMetricResponse {
	s.Body = v
	return s
}

type UpdateABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdateABMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateABMetricGroupRequest) SetABMetricIds(v string) *UpdateABMetricGroupRequest {
	s.ABMetricIds = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetDescription(v string) *UpdateABMetricGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetInstanceId(v string) *UpdateABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetName(v string) *UpdateABMetricGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetRealtime(v bool) *UpdateABMetricGroupRequest {
	s.Realtime = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetSceneId(v string) *UpdateABMetricGroupRequest {
	s.SceneId = &v
	return s
}

type UpdateABMetricGroupResponseBody struct {
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateABMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABMetricGroupResponseBody) SetRequestId(v string) *UpdateABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateABMetricGroupResponse) SetHeaders(v map[string]*string) *UpdateABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateABMetricGroupResponse) SetStatusCode(v int32) *UpdateABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABMetricGroupResponse) SetBody(v *UpdateABMetricGroupResponseBody) *UpdateABMetricGroupResponse {
	s.Body = v
	return s
}

type UpdateCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx人群
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrowdRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrowdRequest) SetDescription(v string) *UpdateCrowdRequest {
	s.Description = &v
	return s
}

func (s *UpdateCrowdRequest) SetInstanceId(v string) *UpdateCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCrowdRequest) SetName(v string) *UpdateCrowdRequest {
	s.Name = &v
	return s
}

type UpdateCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8C27790E-CCA5-56BB-BA17-646295DEC0A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrowdResponseBody) SetRequestId(v string) *UpdateCrowdResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrowdResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrowdResponse) SetHeaders(v map[string]*string) *UpdateCrowdResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrowdResponse) SetStatusCode(v int32) *UpdateCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCrowdResponse) SetBody(v *UpdateCrowdResponseBody) *UpdateCrowdResponse {
	s.Body = v
	return s
}

type UpdateEngineConfigRequest struct {
	// example:
	//
	// {
	//
	// 	"ListenConf": {
	//
	// 		"HttpAddr": "",
	//
	// 		"HttpPort": 8000
	//
	// 	}
	//
	// }
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateEngineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateEngineConfigRequest) SetConfigValue(v string) *UpdateEngineConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetDescription(v string) *UpdateEngineConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetEnvironment(v string) *UpdateEngineConfigRequest {
	s.Environment = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetInstanceId(v string) *UpdateEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEngineConfigRequest) SetName(v string) *UpdateEngineConfigRequest {
	s.Name = &v
	return s
}

type UpdateEngineConfigResponseBody struct {
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEngineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEngineConfigResponseBody) SetRequestId(v string) *UpdateEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEngineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEngineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateEngineConfigResponse) SetHeaders(v map[string]*string) *UpdateEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateEngineConfigResponse) SetStatusCode(v int32) *UpdateEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEngineConfigResponse) SetBody(v *UpdateEngineConfigResponseBody) *UpdateEngineConfigResponse {
	s.Body = v
	return s
}

type UpdateExperimentRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// experiment_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRequest) SetConfig(v string) *UpdateExperimentRequest {
	s.Config = &v
	return s
}

func (s *UpdateExperimentRequest) SetDebugCrowdId(v string) *UpdateExperimentRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *UpdateExperimentRequest) SetDebugUsers(v string) *UpdateExperimentRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateExperimentRequest) SetDescription(v string) *UpdateExperimentRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentRequest) SetFlowPercent(v int32) *UpdateExperimentRequest {
	s.FlowPercent = &v
	return s
}

func (s *UpdateExperimentRequest) SetInstanceId(v string) *UpdateExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateExperimentRequest) SetName(v string) *UpdateExperimentRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentRequest) SetType(v string) *UpdateExperimentRequest {
	s.Type = &v
	return s
}

type UpdateExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A760D972-1475-58C0-BBB3-92B5FB08904F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponseBody) SetRequestId(v string) *UpdateExperimentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponse) SetHeaders(v map[string]*string) *UpdateExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentResponse) SetStatusCode(v int32) *UpdateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentResponse) SetBody(v *UpdateExperimentResponseBody) *UpdateExperimentResponse {
	s.Body = v
	return s
}

type UpdateExperimentGroupRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// user1,user2,user3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// example:
	//
	// gender=male
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NeedAA     *bool  `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	RandomFlow *int64 `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// example:
	//
	// 1,2,3
	ReservcedBuckets *string `json:"ReservcedBuckets,omitempty" xml:"ReservcedBuckets,omitempty"`
}

func (s UpdateExperimentGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentGroupRequest) SetConfig(v string) *UpdateExperimentGroupRequest {
	s.Config = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetCrowdId(v string) *UpdateExperimentGroupRequest {
	s.CrowdId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetCrowdTargetType(v string) *UpdateExperimentGroupRequest {
	s.CrowdTargetType = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDebugCrowdId(v string) *UpdateExperimentGroupRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDebugUsers(v string) *UpdateExperimentGroupRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDescription(v string) *UpdateExperimentGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDistributionTimeDuration(v int32) *UpdateExperimentGroupRequest {
	s.DistributionTimeDuration = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDistributionType(v string) *UpdateExperimentGroupRequest {
	s.DistributionType = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetFilter(v string) *UpdateExperimentGroupRequest {
	s.Filter = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetInstanceId(v string) *UpdateExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetLayerId(v string) *UpdateExperimentGroupRequest {
	s.LayerId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetName(v string) *UpdateExperimentGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetNeedAA(v bool) *UpdateExperimentGroupRequest {
	s.NeedAA = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetRandomFlow(v int64) *UpdateExperimentGroupRequest {
	s.RandomFlow = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetReservcedBuckets(v string) *UpdateExperimentGroupRequest {
	s.ReservcedBuckets = &v
	return s
}

type UpdateExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentGroupResponseBody) SetRequestId(v string) *UpdateExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentGroupResponse) SetHeaders(v map[string]*string) *UpdateExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentGroupResponse) SetStatusCode(v int32) *UpdateExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentGroupResponse) SetBody(v *UpdateExperimentGroupResponseBody) *UpdateExperimentGroupResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s UpdateFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
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

type UpdateFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *UpdateFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeatureConsistencyCheckJobConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *UpdateFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *UpdateFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponse) SetBody(v *UpdateFeatureConsistencyCheckJobConfigResponseBody) *UpdateFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

type UpdateInstanceResourceRequest struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateInstanceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResourceRequest) SetConfig(v string) *UpdateInstanceResourceRequest {
	s.Config = &v
	return s
}

func (s *UpdateInstanceResourceRequest) SetUri(v string) *UpdateInstanceResourceRequest {
	s.Uri = &v
	return s
}

type UpdateInstanceResourceResponseBody struct {
	// example:
	//
	// 3AAA45F6-0798-5461-9360-81D133823CE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResourceResponseBody) SetRequestId(v string) *UpdateInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResourceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResourceResponse) SetStatusCode(v int32) *UpdateInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResourceResponse) SetBody(v *UpdateInstanceResourceResponseBody) *UpdateInstanceResourceResponse {
	s.Body = v
	return s
}

type UpdateLaboratoryRequest struct {
	// example:
	//
	// 24
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Filter
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// example:
	//
	// 1,2,3,10-20
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// filter=xxx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// laboratory1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateLaboratoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateLaboratoryRequest) SetBucketCount(v int32) *UpdateLaboratoryRequest {
	s.BucketCount = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetBucketType(v string) *UpdateLaboratoryRequest {
	s.BucketType = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetBuckets(v string) *UpdateLaboratoryRequest {
	s.Buckets = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetDebugCrowdId(v string) *UpdateLaboratoryRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetDebugUsers(v string) *UpdateLaboratoryRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetDescription(v string) *UpdateLaboratoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetEnvironment(v string) *UpdateLaboratoryRequest {
	s.Environment = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetFilter(v string) *UpdateLaboratoryRequest {
	s.Filter = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetInstanceId(v string) *UpdateLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetName(v string) *UpdateLaboratoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateLaboratoryRequest) SetType(v string) *UpdateLaboratoryRequest {
	s.Type = &v
	return s
}

type UpdateLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLaboratoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLaboratoryResponseBody) SetRequestId(v string) *UpdateLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLaboratoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateLaboratoryResponse) SetHeaders(v map[string]*string) *UpdateLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateLaboratoryResponse) SetStatusCode(v int32) *UpdateLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLaboratoryResponse) SetBody(v *UpdateLaboratoryResponseBody) *UpdateLaboratoryResponse {
	s.Body = v
	return s
}

type UpdateLayerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// layer1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayerRequest) GoString() string {
	return s.String()
}

func (s *UpdateLayerRequest) SetDescription(v string) *UpdateLayerRequest {
	s.Description = &v
	return s
}

func (s *UpdateLayerRequest) SetInstanceId(v string) *UpdateLayerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLayerRequest) SetName(v string) *UpdateLayerRequest {
	s.Name = &v
	return s
}

type UpdateLayerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0EA9215E-EC21-53AB-B8D9-D3DEA90D040A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLayerResponseBody) SetRequestId(v string) *UpdateLayerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayerResponse) GoString() string {
	return s.String()
}

func (s *UpdateLayerResponse) SetHeaders(v map[string]*string) *UpdateLayerResponse {
	s.Headers = v
	return s
}

func (s *UpdateLayerResponse) SetStatusCode(v int32) *UpdateLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLayerResponse) SetBody(v *UpdateLayerResponseBody) *UpdateLayerResponse {
	s.Body = v
	return s
}

type UpdateParamRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// house
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateParamRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateParamRequest) GoString() string {
	return s.String()
}

func (s *UpdateParamRequest) SetInstanceId(v string) *UpdateParamRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateParamRequest) SetValue(v string) *UpdateParamRequest {
	s.Value = &v
	return s
}

type UpdateParamResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BBD41FBF-E75C-551A-92FA-CAD654AA006F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateParamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParamResponseBody) SetRequestId(v string) *UpdateParamResponseBody {
	s.RequestId = &v
	return s
}

type UpdateParamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateParamResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateParamResponse) GoString() string {
	return s.String()
}

func (s *UpdateParamResponse) SetHeaders(v map[string]*string) *UpdateParamResponse {
	s.Headers = v
	return s
}

func (s *UpdateParamResponse) SetStatusCode(v int32) *UpdateParamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateParamResponse) SetBody(v *UpdateParamResponseBody) *UpdateParamResponse {
	s.Body = v
	return s
}

type UpdateResourceRuleRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricOperationType *string `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo      *string `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod    *string `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RuleComputingDefinition *string `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
}

func (s UpdateResourceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleRequest) SetDescription(v string) *UpdateResourceRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetInstanceId(v string) *UpdateResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetMetricOperationType(v string) *UpdateResourceRuleRequest {
	s.MetricOperationType = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetMetricPullInfo(v string) *UpdateResourceRuleRequest {
	s.MetricPullInfo = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetMetricPullPeriod(v string) *UpdateResourceRuleRequest {
	s.MetricPullPeriod = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetName(v string) *UpdateResourceRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetRuleComputingDefinition(v string) *UpdateResourceRuleRequest {
	s.RuleComputingDefinition = &v
	return s
}

type UpdateResourceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleResponseBody) SetRequestId(v string) *UpdateResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleResponse) SetHeaders(v map[string]*string) *UpdateResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceRuleResponse) SetStatusCode(v int32) *UpdateResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceRuleResponse) SetBody(v *UpdateResourceRuleResponseBody) *UpdateResourceRuleResponse {
	s.Body = v
	return s
}

type UpdateResourceRuleItemRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxValue   *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue   *float64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// This parameter is required.
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateResourceRuleItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRuleItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleItemRequest) SetDescription(v string) *UpdateResourceRuleItemRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetInstanceId(v string) *UpdateResourceRuleItemRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetMaxValue(v float64) *UpdateResourceRuleItemRequest {
	s.MaxValue = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetMinValue(v float64) *UpdateResourceRuleItemRequest {
	s.MinValue = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetName(v string) *UpdateResourceRuleItemRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetValue(v float64) *UpdateResourceRuleItemRequest {
	s.Value = &v
	return s
}

type UpdateResourceRuleItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceRuleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRuleItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleItemResponseBody) SetRequestId(v string) *UpdateResourceRuleItemResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceRuleItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceRuleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceRuleItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRuleItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleItemResponse) SetHeaders(v map[string]*string) *UpdateResourceRuleItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceRuleItemResponse) SetStatusCode(v int32) *UpdateResourceRuleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceRuleItemResponse) SetBody(v *UpdateResourceRuleItemResponseBody) *UpdateResourceRuleItemResponse {
	s.Body = v
	return s
}

type UpdateSceneRequest struct {
	// example:
	//
	// This is a test.
	Description *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*UpdateSceneRequestFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSceneRequest) SetDescription(v string) *UpdateSceneRequest {
	s.Description = &v
	return s
}

func (s *UpdateSceneRequest) SetFlows(v []*UpdateSceneRequestFlows) *UpdateSceneRequest {
	s.Flows = v
	return s
}

func (s *UpdateSceneRequest) SetInstanceId(v string) *UpdateSceneRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSceneRequest) SetName(v string) *UpdateSceneRequest {
	s.Name = &v
	return s
}

type UpdateSceneRequestFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s UpdateSceneRequestFlows) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneRequestFlows) GoString() string {
	return s.String()
}

func (s *UpdateSceneRequestFlows) SetFlowCode(v string) *UpdateSceneRequestFlows {
	s.FlowCode = &v
	return s
}

func (s *UpdateSceneRequestFlows) SetFlowName(v string) *UpdateSceneRequestFlows {
	s.FlowName = &v
	return s
}

type UpdateSceneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FC17887E-3C82-5096-8AA6-F4C2E7417245
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponseBody) SetRequestId(v string) *UpdateSceneResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponse) SetHeaders(v map[string]*string) *UpdateSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSceneResponse) SetStatusCode(v int32) *UpdateSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSceneResponse) SetBody(v *UpdateSceneResponseBody) *UpdateSceneResponse {
	s.Body = v
	return s
}

type UpdateTableMetaRequest struct {
	// example:
	//
	// this is a test table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Fields []*UpdateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table_mysql
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s UpdateTableMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaRequest) SetDescription(v string) *UpdateTableMetaRequest {
	s.Description = &v
	return s
}

func (s *UpdateTableMetaRequest) SetFields(v []*UpdateTableMetaRequestFields) *UpdateTableMetaRequest {
	s.Fields = v
	return s
}

func (s *UpdateTableMetaRequest) SetInstanceId(v string) *UpdateTableMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTableMetaRequest) SetModule(v string) *UpdateTableMetaRequest {
	s.Module = &v
	return s
}

func (s *UpdateTableMetaRequest) SetName(v string) *UpdateTableMetaRequest {
	s.Name = &v
	return s
}

func (s *UpdateTableMetaRequest) SetResourceId(v string) *UpdateTableMetaRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateTableMetaRequest) SetTableName(v string) *UpdateTableMetaRequest {
	s.TableName = &v
	return s
}

type UpdateTableMetaRequestFields struct {
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	// This parameter is required.
	IsPartitionField *string `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	// This parameter is required.
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTableMetaRequestFields) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableMetaRequestFields) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaRequestFields) SetDataType(v string) *UpdateTableMetaRequestFields {
	s.DataType = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetIsDimensionField(v bool) *UpdateTableMetaRequestFields {
	s.IsDimensionField = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetIsPartitionField(v string) *UpdateTableMetaRequestFields {
	s.IsPartitionField = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetMeaning(v string) *UpdateTableMetaRequestFields {
	s.Meaning = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetName(v string) *UpdateTableMetaRequestFields {
	s.Name = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetType(v string) *UpdateTableMetaRequestFields {
	s.Type = &v
	return s
}

type UpdateTableMetaResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTableMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaResponseBody) SetRequestId(v string) *UpdateTableMetaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTableMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaResponse) SetHeaders(v map[string]*string) *UpdateTableMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableMetaResponse) SetStatusCode(v int32) *UpdateTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableMetaResponse) SetBody(v *UpdateTableMetaResponseBody) *UpdateTableMetaResponse {
	s.Body = v
	return s
}

type UpdateTrafficControlTargetRequest struct {
	EndTime              *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ItemConditionArray   *string  `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress *string  `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType    *string  `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                 *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation *bool    `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName           *string  `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	StartTime            *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod         *string  `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status               *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue       *int64   `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	Value                *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
	NewParam3            *string  `json:"new-param-3,omitempty" xml:"new-param-3,omitempty"`
}

func (s UpdateTrafficControlTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTargetRequest) SetEndTime(v string) *UpdateTrafficControlTargetRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetEvent(v string) *UpdateTrafficControlTargetRequest {
	s.Event = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetItemConditionArray(v string) *UpdateTrafficControlTargetRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetItemConditionExpress(v string) *UpdateTrafficControlTargetRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetItemConditionType(v string) *UpdateTrafficControlTargetRequest {
	s.ItemConditionType = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetName(v string) *UpdateTrafficControlTargetRequest {
	s.Name = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetNewProductRegulation(v bool) *UpdateTrafficControlTargetRequest {
	s.NewProductRegulation = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetRecallName(v string) *UpdateTrafficControlTargetRequest {
	s.RecallName = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetStartTime(v string) *UpdateTrafficControlTargetRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetStatisPeriod(v string) *UpdateTrafficControlTargetRequest {
	s.StatisPeriod = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetStatus(v string) *UpdateTrafficControlTargetRequest {
	s.Status = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetToleranceValue(v int64) *UpdateTrafficControlTargetRequest {
	s.ToleranceValue = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetValue(v float32) *UpdateTrafficControlTargetRequest {
	s.Value = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetNewParam3(v string) *UpdateTrafficControlTargetRequest {
	s.NewParam3 = &v
	return s
}

type UpdateTrafficControlTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficControlTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTargetResponseBody) SetRequestId(v string) *UpdateTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrafficControlTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficControlTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTargetResponse) SetHeaders(v map[string]*string) *UpdateTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficControlTargetResponse) SetStatusCode(v int32) *UpdateTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficControlTargetResponse) SetBody(v *UpdateTrafficControlTargetResponseBody) *UpdateTrafficControlTargetResponse {
	s.Body = v
	return s
}

type UpdateTrafficControlTaskRequest struct {
	BehaviorTableMetaId            *string                                                 `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	ControlGranularity             *string                                                 `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	ControlLogic                   *string                                                 `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	ControlType                    *string                                                 `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	Description                    *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                        *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutionTime                  *string                                                 `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	InstanceId                     *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemConditionArray             *string                                                 `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress           *string                                                 `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType              *string                                                 `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	ItemTableMetaId                *string                                                 `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	Name                           *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds               *string                                                 `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	ProdExperimentIds              *string                                                 `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	SceneId                        *string                                                 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ServiceId                      *string                                                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	StartTime                      *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisBaeaviorConditionArray   *string                                                 `json:"StatisBaeaviorConditionArray,omitempty" xml:"StatisBaeaviorConditionArray,omitempty"`
	StatisBehaviorConditionExpress *string                                                 `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	StatisBehaviorConditionType    *string                                                 `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets          []*UpdateTrafficControlTaskRequestTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	UserConditionArray             *string                                                 `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	UserConditionExpress           *string                                                 `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	UserConditionType              *string                                                 `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	UserTableMetaId                *string                                                 `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s UpdateTrafficControlTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskRequest) SetBehaviorTableMetaId(v string) *UpdateTrafficControlTaskRequest {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetControlGranularity(v string) *UpdateTrafficControlTaskRequest {
	s.ControlGranularity = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetControlLogic(v string) *UpdateTrafficControlTaskRequest {
	s.ControlLogic = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetControlType(v string) *UpdateTrafficControlTaskRequest {
	s.ControlType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetDescription(v string) *UpdateTrafficControlTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetEndTime(v string) *UpdateTrafficControlTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetExecutionTime(v string) *UpdateTrafficControlTaskRequest {
	s.ExecutionTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetInstanceId(v string) *UpdateTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemConditionArray(v string) *UpdateTrafficControlTaskRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemConditionExpress(v string) *UpdateTrafficControlTaskRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemConditionType(v string) *UpdateTrafficControlTaskRequest {
	s.ItemConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemTableMetaId(v string) *UpdateTrafficControlTaskRequest {
	s.ItemTableMetaId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetName(v string) *UpdateTrafficControlTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetPreExperimentIds(v string) *UpdateTrafficControlTaskRequest {
	s.PreExperimentIds = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetProdExperimentIds(v string) *UpdateTrafficControlTaskRequest {
	s.ProdExperimentIds = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetSceneId(v string) *UpdateTrafficControlTaskRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetServiceId(v string) *UpdateTrafficControlTaskRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStartTime(v string) *UpdateTrafficControlTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStatisBaeaviorConditionArray(v string) *UpdateTrafficControlTaskRequest {
	s.StatisBaeaviorConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStatisBehaviorConditionExpress(v string) *UpdateTrafficControlTaskRequest {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStatisBehaviorConditionType(v string) *UpdateTrafficControlTaskRequest {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetTrafficControlTargets(v []*UpdateTrafficControlTaskRequestTrafficControlTargets) *UpdateTrafficControlTaskRequest {
	s.TrafficControlTargets = v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserConditionArray(v string) *UpdateTrafficControlTaskRequest {
	s.UserConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserConditionExpress(v string) *UpdateTrafficControlTaskRequest {
	s.UserConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserConditionType(v string) *UpdateTrafficControlTaskRequest {
	s.UserConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserTableMetaId(v string) *UpdateTrafficControlTaskRequest {
	s.UserTableMetaId = &v
	return s
}

type UpdateTrafficControlTaskRequestTrafficControlTargets struct {
	EndTime              *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ItemConditionArray   *string  `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress *string  `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType    *string  `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                 *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation *bool    `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName           *string  `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	StartTime            *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod         *string  `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status               *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue       *int64   `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	Value                *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTrafficControlTaskRequestTrafficControlTargets) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskRequestTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetEndTime(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetEvent(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionArray(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionExpress(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionType(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetName(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetNewProductRegulation(v bool) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetRecallName(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetStartTime(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetStatisPeriod(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetStatus(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetToleranceValue(v int64) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetValue(v float32) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Value = &v
	return s
}

type UpdateTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficControlTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskResponseBody) SetRequestId(v string) *UpdateTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrafficControlTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficControlTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskResponse) SetHeaders(v map[string]*string) *UpdateTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficControlTaskResponse) SetStatusCode(v int32) *UpdateTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficControlTaskResponse) SetBody(v *UpdateTrafficControlTaskResponseBody) *UpdateTrafficControlTaskResponse {
	s.Body = v
	return s
}

type UpdateTrafficControlTaskTrafficRequest struct {
	Environment *string                                           `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Traffics    []*UpdateTrafficControlTaskTrafficRequestTraffics `json:"Traffics,omitempty" xml:"Traffics,omitempty" type:"Repeated"`
	NewParam3   *string                                           `json:"new-param-3,omitempty" xml:"new-param-3,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetEnvironment(v string) *UpdateTrafficControlTaskTrafficRequest {
	s.Environment = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetInstanceId(v string) *UpdateTrafficControlTaskTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetTraffics(v []*UpdateTrafficControlTaskTrafficRequestTraffics) *UpdateTrafficControlTaskTrafficRequest {
	s.Traffics = v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetNewParam3(v string) *UpdateTrafficControlTaskTrafficRequest {
	s.NewParam3 = &v
	return s
}

type UpdateTrafficControlTaskTrafficRequestTraffics struct {
	ItemOrExperimentId             *string  `json:"ItemOrExperimentId,omitempty" xml:"ItemOrExperimentId,omitempty"`
	RecordTime                     *string  `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
	TrafficControlTargetAimTraffic *float64 `json:"TrafficControlTargetAimTraffic,omitempty" xml:"TrafficControlTargetAimTraffic,omitempty"`
	TrafficControlTargetId         *string  `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	TrafficControlTargetTraffic    *int64   `json:"TrafficControlTargetTraffic,omitempty" xml:"TrafficControlTargetTraffic,omitempty"`
	TrafficControlTaskTraffic      *int64   `json:"TrafficControlTaskTraffic,omitempty" xml:"TrafficControlTaskTraffic,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficRequestTraffics) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficRequestTraffics) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetItemOrExperimentId(v string) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.ItemOrExperimentId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetRecordTime(v string) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.RecordTime = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTargetAimTraffic(v float64) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTargetAimTraffic = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTargetId(v string) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTargetId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTargetTraffic(v int64) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTargetTraffic = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTaskTraffic(v int64) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTaskTraffic = &v
	return s
}

type UpdateTrafficControlTaskTrafficResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficResponseBody) SetRequestId(v string) *UpdateTrafficControlTaskTrafficResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrafficControlTaskTrafficResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficControlTaskTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficResponse) SetHeaders(v map[string]*string) *UpdateTrafficControlTaskTrafficResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficControlTaskTrafficResponse) SetStatusCode(v int32) *UpdateTrafficControlTaskTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficResponse) SetBody(v *UpdateTrafficControlTaskTrafficResponseBody) *UpdateTrafficControlTaskTrafficResponse {
	s.Body = v
	return s
}

type UploadRecommendationDataRequest struct {
	RegionId *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Content  []*UploadRecommendationDataRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	DataType *string                                   `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s UploadRecommendationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRecommendationDataRequest) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataRequest) SetRegionId(v string) *UploadRecommendationDataRequest {
	s.RegionId = &v
	return s
}

func (s *UploadRecommendationDataRequest) SetContent(v []*UploadRecommendationDataRequestContent) *UploadRecommendationDataRequest {
	s.Content = v
	return s
}

func (s *UploadRecommendationDataRequest) SetDataType(v string) *UploadRecommendationDataRequest {
	s.DataType = &v
	return s
}

type UploadRecommendationDataRequestContent struct {
	Fields        *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s UploadRecommendationDataRequestContent) String() string {
	return tea.Prettify(s)
}

func (s UploadRecommendationDataRequestContent) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataRequestContent) SetFields(v string) *UploadRecommendationDataRequestContent {
	s.Fields = &v
	return s
}

func (s *UploadRecommendationDataRequestContent) SetOperationType(v string) *UploadRecommendationDataRequestContent {
	s.OperationType = &v
	return s
}

type UploadRecommendationDataResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadRecommendationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRecommendationDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataResponseBody) SetMessage(v string) *UploadRecommendationDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRecommendationDataResponseBody) SetRequestId(v string) *UploadRecommendationDataResponseBody {
	s.RequestId = &v
	return s
}

type UploadRecommendationDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRecommendationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRecommendationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadRecommendationDataResponse) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataResponse) SetHeaders(v map[string]*string) *UploadRecommendationDataResponse {
	s.Headers = v
	return s
}

func (s *UploadRecommendationDataResponse) SetStatusCode(v int32) *UploadRecommendationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRecommendationDataResponse) SetBody(v *UploadRecommendationDataResponseBody) *UploadRecommendationDataResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("pairecservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 应用/发布指定的推荐引擎配置
//
// @param request - ApplyEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyEngineConfigResponse
func (client *Client) ApplyEngineConfigWithOptions(EngineConfigId *string, request *ApplyEngineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyEngineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyEngineConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(EngineConfigId)) + "/action/apply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ApplyEngineConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ApplyEngineConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 应用/发布指定的推荐引擎配置
//
// @param request - ApplyEngineConfigRequest
//
// @return ApplyEngineConfigResponse
func (client *Client) ApplyEngineConfig(EngineConfigId *string, request *ApplyEngineConfigRequest) (_result *ApplyEngineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyEngineConfigResponse{}
	_body, _err := client.ApplyEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 特征一致性检查数据回流。
//
// @param request - BackflowFeatureConsistencyCheckJobDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackflowFeatureConsistencyCheckJobDataResponse
func (client *Client) BackflowFeatureConsistencyCheckJobDataWithOptions(request *BackflowFeatureConsistencyCheckJobDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BackflowFeatureConsistencyCheckJobDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureConsistencyCheckJobConfigId)) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemFeatures)) {
		body["ItemFeatures"] = request.ItemFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.LogItemId)) {
		body["LogItemId"] = request.LogItemId
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestTime)) {
		body["LogRequestTime"] = request.LogRequestTime
	}

	if !tea.BoolValue(util.IsUnset(request.LogUserId)) {
		body["LogUserId"] = request.LogUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		body["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.Scores)) {
		body["Scores"] = request.Scores
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.UserFeatures)) {
		body["UserFeatures"] = request.UserFeatures
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BackflowFeatureConsistencyCheckJobData"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs/action/backflowdata"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 特征一致性检查数据回流。
//
// @param request - BackflowFeatureConsistencyCheckJobDataRequest
//
// @return BackflowFeatureConsistencyCheckJobDataResponse
func (client *Client) BackflowFeatureConsistencyCheckJobData(request *BackflowFeatureConsistencyCheckJobDataRequest) (_result *BackflowFeatureConsistencyCheckJobDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
	_body, _err := client.BackflowFeatureConsistencyCheckJobDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检测实例下配置的资源的连接状态。
//
// @param request - CheckInstanceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceResourcesResponse
func (client *Client) CheckInstanceResourcesWithOptions(InstanceId *string, request *CheckInstanceResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckInstanceResources"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/action/checkresources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckInstanceResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckInstanceResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 检测实例下配置的资源的连接状态。
//
// @param request - CheckInstanceResourcesRequest
//
// @return CheckInstanceResourcesResponse
func (client *Client) CheckInstanceResources(InstanceId *string, request *CheckInstanceResourcesRequest) (_result *CheckInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceResourcesResponse{}
	_body, _err := client.CheckInstanceResourcesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验流量调控任务中的表达式
//
// @param request - CheckTrafficControlTaskExpressionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTrafficControlTaskExpressionResponse
func (client *Client) CheckTrafficControlTaskExpressionWithOptions(request *CheckTrafficControlTaskExpressionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckTrafficControlTaskExpressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TableMetaId)) {
		query["TableMetaId"] = request.TableMetaId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckTrafficControlTaskExpression"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/action/checkexpression"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckTrafficControlTaskExpressionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckTrafficControlTaskExpressionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 校验流量调控任务中的表达式
//
// @param request - CheckTrafficControlTaskExpressionRequest
//
// @return CheckTrafficControlTaskExpressionResponse
func (client *Client) CheckTrafficControlTaskExpression(request *CheckTrafficControlTaskExpressionRequest) (_result *CheckTrafficControlTaskExpressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckTrafficControlTaskExpressionResponse{}
	_body, _err := client.CheckTrafficControlTaskExpressionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆指定的推荐引擎配置
//
// @param request - CloneEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneEngineConfigResponse
func (client *Client) CloneEngineConfigWithOptions(EngineConfigId *string, request *CloneEngineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneEngineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigValue)) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneEngineConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(EngineConfigId)) + "/action/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloneEngineConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloneEngineConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 克隆指定的推荐引擎配置
//
// @param request - CloneEngineConfigRequest
//
// @return CloneEngineConfigResponse
func (client *Client) CloneEngineConfig(EngineConfigId *string, request *CloneEngineConfigRequest) (_result *CloneEngineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneEngineConfigResponse{}
	_body, _err := client.CloneEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆实验。
//
// @param request - CloneExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneExperimentResponse
func (client *Client) CloneExperimentWithOptions(ExperimentId *string, request *CloneExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/action/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloneExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloneExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 克隆实验。
//
// @param request - CloneExperimentRequest
//
// @return CloneExperimentResponse
func (client *Client) CloneExperiment(ExperimentId *string, request *CloneExperimentRequest) (_result *CloneExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneExperimentResponse{}
	_body, _err := client.CloneExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆实验组，并克隆实验组下的所有实验至新的实验组中。
//
// @param request - CloneExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneExperimentGroupResponse
func (client *Client) CloneExperimentGroupWithOptions(ExperimentGroupId *string, request *CloneExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		body["LayerId"] = request.LayerId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentGroupId)) + "/action/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloneExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloneExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 克隆实验组，并克隆实验组下的所有实验至新的实验组中。
//
// @param request - CloneExperimentGroupRequest
//
// @return CloneExperimentGroupResponse
func (client *Client) CloneExperimentGroup(ExperimentGroupId *string, request *CloneExperimentGroupRequest) (_result *CloneExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneExperimentGroupResponse{}
	_body, _err := client.CloneExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆特征一致性检查配置。
//
// @param request - CloneFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneFeatureConsistencyCheckJobConfigResponse
func (client *Client) CloneFeatureConsistencyCheckJobConfigWithOptions(SourceFeatureConsistencyCheckJobConfigId *string, request *CloneFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneFeatureConsistencyCheckJobConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(SourceFeatureConsistencyCheckJobConfigId)) + "/action/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 克隆特征一致性检查配置。
//
// @param request - CloneFeatureConsistencyCheckJobConfigRequest
//
// @return CloneFeatureConsistencyCheckJobConfigResponse
func (client *Client) CloneFeatureConsistencyCheckJobConfig(SourceFeatureConsistencyCheckJobConfigId *string, request *CloneFeatureConsistencyCheckJobConfigRequest) (_result *CloneFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CloneFeatureConsistencyCheckJobConfigWithOptions(SourceFeatureConsistencyCheckJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆实验室。
//
// @param request - CloneLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneLaboratoryResponse
func (client *Client) CloneLaboratoryWithOptions(LaboratoryId *string, request *CloneLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloneExperimentGroup)) {
		body["CloneExperimentGroup"] = request.CloneExperimentGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories/" + tea.StringValue(openapiutil.GetEncodeParam(LaboratoryId)) + "/action/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloneLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloneLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 克隆实验室。
//
// @param request - CloneLaboratoryRequest
//
// @return CloneLaboratoryResponse
func (client *Client) CloneLaboratory(LaboratoryId *string, request *CloneLaboratoryRequest) (_result *CloneLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneLaboratoryResponse{}
	_body, _err := client.CloneLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆流量调控任务
//
// @param request - CloneTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneTrafficControlTaskResponse
func (client *Client) CloneTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *CloneTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloneTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloneTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 克隆流量调控任务
//
// @param request - CloneTrafficControlTaskRequest
//
// @return CloneTrafficControlTaskResponse
func (client *Client) CloneTrafficControlTask(TrafficControlTaskId *string, request *CloneTrafficControlTaskRequest) (_result *CloneTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneTrafficControlTaskResponse{}
	_body, _err := client.CloneTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AB test实验指标
//
// @param request - CreateABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABMetricResponse
func (client *Client) CreateABMetricWithOptions(request *CreateABMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		body["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LeftMetricId)) {
		body["LeftMetricId"] = request.LeftMetricId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.Realtime)) {
		body["Realtime"] = request.Realtime
	}

	if !tea.BoolValue(util.IsUnset(request.ResultResourceId)) {
		body["ResultResourceId"] = request.ResultResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RightMetricId)) {
		body["RightMetricId"] = request.RightMetricId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.StatisticsCycle)) {
		body["StatisticsCycle"] = request.StatisticsCycle
	}

	if !tea.BoolValue(util.IsUnset(request.TableMetaId)) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABMetric"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetrics"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateABMetricResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateABMetricResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建AB test实验指标
//
// @param request - CreateABMetricRequest
//
// @return CreateABMetricResponse
func (client *Client) CreateABMetric(request *CreateABMetricRequest) (_result *CreateABMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABMetricResponse{}
	_body, _err := client.CreateABMetricWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建指标组
//
// @param request - CreateABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABMetricGroupResponse
func (client *Client) CreateABMetricGroupWithOptions(request *CreateABMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ABMetricIds)) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Realtime)) {
		body["Realtime"] = request.Realtime
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABMetricGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetricgroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateABMetricGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateABMetricGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建指标组
//
// @param request - CreateABMetricGroupRequest
//
// @return CreateABMetricGroupResponse
func (client *Client) CreateABMetricGroup(request *CreateABMetricGroupRequest) (_result *CreateABMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABMetricGroupResponse{}
	_body, _err := client.CreateABMetricGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AB指标的计算任务。
//
// @param request - CreateCalculationJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCalculationJobsResponse
func (client *Client) CreateCalculationJobsWithOptions(request *CreateCalculationJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCalculationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ABMetricIds)) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCalculationJobs"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/batch/calculationjobs/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateCalculationJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateCalculationJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建AB指标的计算任务。
//
// @param request - CreateCalculationJobsRequest
//
// @return CreateCalculationJobsResponse
func (client *Client) CreateCalculationJobs(request *CreateCalculationJobsRequest) (_result *CreateCalculationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCalculationJobsResponse{}
	_body, _err := client.CreateCalculationJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建人群。
//
// @param request - CreateCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCrowdResponse
func (client *Client) CreateCrowdWithOptions(request *CreateCrowdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCrowd"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateCrowdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateCrowdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建人群。
//
// @param request - CreateCrowdRequest
//
// @return CreateCrowdResponse
func (client *Client) CreateCrowd(request *CreateCrowdRequest) (_result *CreateCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCrowdResponse{}
	_body, _err := client.CreateCrowdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建引擎配置
//
// @param request - CreateEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEngineConfigResponse
func (client *Client) CreateEngineConfigWithOptions(request *CreateEngineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEngineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigValue)) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEngineConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEngineConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEngineConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建引擎配置
//
// @param request - CreateEngineConfigRequest
//
// @return CreateEngineConfigResponse
func (client *Client) CreateEngineConfig(request *CreateEngineConfigRequest) (_result *CreateEngineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEngineConfigResponse{}
	_body, _err := client.CreateEngineConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验。
//
// @param request - CreateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentResponse
func (client *Client) CreateExperimentWithOptions(request *CreateExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DebugCrowdId)) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentGroupId)) {
		body["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowPercent)) {
		body["FlowPercent"] = request.FlowPercent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建实验。
//
// @param request - CreateExperimentRequest
//
// @return CreateExperimentResponse
func (client *Client) CreateExperiment(request *CreateExperimentRequest) (_result *CreateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentResponse{}
	_body, _err := client.CreateExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验组。
//
// @param request - CreateExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentGroupResponse
func (client *Client) CreateExperimentGroupWithOptions(request *CreateExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdId)) {
		body["CrowdId"] = request.CrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdTargetType)) {
		body["CrowdTargetType"] = request.CrowdTargetType
	}

	if !tea.BoolValue(util.IsUnset(request.DebugCrowdId)) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionTimeDuration)) {
		body["DistributionTimeDuration"] = request.DistributionTimeDuration
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionType)) {
		body["DistributionType"] = request.DistributionType
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		body["LayerId"] = request.LayerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NeedAA)) {
		body["NeedAA"] = request.NeedAA
	}

	if !tea.BoolValue(util.IsUnset(request.RandomFlow)) {
		body["RandomFlow"] = request.RandomFlow
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedBuckets)) {
		body["ReservedBuckets"] = request.ReservedBuckets
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建实验组。
//
// @param request - CreateExperimentGroupRequest
//
// @return CreateExperimentGroupResponse
func (client *Client) CreateExperimentGroup(request *CreateExperimentGroupRequest) (_result *CreateExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentGroupResponse{}
	_body, _err := client.CreateExperimentGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建特征一致性检查任务。
//
// @param request - CreateFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureConsistencyCheckJobResponse
func (client *Client) CreateFeatureConsistencyCheckJobWithOptions(request *CreateFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureConsistencyCheckJobConfigId)) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingDuration)) {
		body["SamplingDuration"] = request.SamplingDuration
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFeatureConsistencyCheckJob"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateFeatureConsistencyCheckJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateFeatureConsistencyCheckJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建特征一致性检查任务。
//
// @param request - CreateFeatureConsistencyCheckJobRequest
//
// @return CreateFeatureConsistencyCheckJobResponse
func (client *Client) CreateFeatureConsistencyCheckJob(request *CreateFeatureConsistencyCheckJobRequest) (_result *CreateFeatureConsistencyCheckJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CreateFeatureConsistencyCheckJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建特征一致性检查配置。
//
// @param request - CreateFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureConsistencyCheckJobConfigResponse
func (client *Client) CreateFeatureConsistencyCheckJobConfigWithOptions(request *CreateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareFeature)) {
		body["CompareFeature"] = request.CompareFeature
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMountPath)) {
		body["DatasetMountPath"] = request.DatasetMountPath
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetType)) {
		body["DatasetType"] = request.DatasetType
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetUri)) {
		body["DatasetUri"] = request.DatasetUri
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRoute)) {
		body["DefaultRoute"] = request.DefaultRoute
	}

	if !tea.BoolValue(util.IsUnset(request.EasServiceName)) {
		body["EasServiceName"] = request.EasServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.EasyRecPackagePath)) {
		body["EasyRecPackagePath"] = request.EasyRecPackagePath
	}

	if !tea.BoolValue(util.IsUnset(request.EasyRecVersion)) {
		body["EasyRecVersion"] = request.EasyRecVersion
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureDisplayExclude)) {
		body["FeatureDisplayExclude"] = request.FeatureDisplayExclude
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureLandingResourceId)) {
		body["FeatureLandingResourceId"] = request.FeatureLandingResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.FeaturePriority)) {
		body["FeaturePriority"] = request.FeaturePriority
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreItemId)) {
		body["FeatureStoreItemId"] = request.FeatureStoreItemId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreModelId)) {
		body["FeatureStoreModelId"] = request.FeatureStoreModelId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreProjectId)) {
		body["FeatureStoreProjectId"] = request.FeatureStoreProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreProjectName)) {
		body["FeatureStoreProjectName"] = request.FeatureStoreProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreSeqFeatureView)) {
		body["FeatureStoreSeqFeatureView"] = request.FeatureStoreSeqFeatureView
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreUserId)) {
		body["FeatureStoreUserId"] = request.FeatureStoreUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FgJarVersion)) {
		body["FgJarVersion"] = request.FgJarVersion
	}

	if !tea.BoolValue(util.IsUnset(request.FgJsonFileName)) {
		body["FgJsonFileName"] = request.FgJsonFileName
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateZip)) {
		body["GenerateZip"] = request.GenerateZip
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemIdField)) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTable)) {
		body["ItemTable"] = request.ItemTable
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTablePartitionField)) {
		body["ItemTablePartitionField"] = request.ItemTablePartitionField
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTablePartitionFieldFormat)) {
		body["ItemTablePartitionFieldFormat"] = request.ItemTablePartitionFieldFormat
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssResourceId)) {
		body["OssResourceId"] = request.OssResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.PredictWorkerCount)) {
		body["PredictWorkerCount"] = request.PredictWorkerCount
	}

	if !tea.BoolValue(util.IsUnset(request.PredictWorkerCpu)) {
		body["PredictWorkerCpu"] = request.PredictWorkerCpu
	}

	if !tea.BoolValue(util.IsUnset(request.PredictWorkerMemory)) {
		body["PredictWorkerMemory"] = request.PredictWorkerMemory
	}

	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		body["SampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchId)) {
		body["SwitchId"] = request.SwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.UseFeatureStore)) {
		body["UseFeatureStore"] = request.UseFeatureStore
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdField)) {
		body["UserIdField"] = request.UserIdField
	}

	if !tea.BoolValue(util.IsUnset(request.UserTable)) {
		body["UserTable"] = request.UserTable
	}

	if !tea.BoolValue(util.IsUnset(request.UserTablePartitionField)) {
		body["UserTablePartitionField"] = request.UserTablePartitionField
	}

	if !tea.BoolValue(util.IsUnset(request.UserTablePartitionFieldFormat)) {
		body["UserTablePartitionFieldFormat"] = request.UserTablePartitionFieldFormat
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowName)) {
		body["WorkflowName"] = request.WorkflowName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFeatureConsistencyCheckJobConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobconfigs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建特征一致性检查配置。
//
// @param request - CreateFeatureConsistencyCheckJobConfigRequest
//
// @return CreateFeatureConsistencyCheckJobConfigResponse
func (client *Client) CreateFeatureConsistencyCheckJobConfig(request *CreateFeatureConsistencyCheckJobConfigRequest) (_result *CreateFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CreateFeatureConsistencyCheckJobConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为指定实例配置创建新的配置资源
//
// @param request - CreateInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResourceResponse
func (client *Client) CreateInstanceResourceWithOptions(InstanceId *string, request *CreateInstanceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		body["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceResource"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateInstanceResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateInstanceResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 为指定实例配置创建新的配置资源
//
// @param request - CreateInstanceResourceRequest
//
// @return CreateInstanceResourceResponse
func (client *Client) CreateInstanceResource(InstanceId *string, request *CreateInstanceResourceRequest) (_result *CreateInstanceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResourceResponse{}
	_body, _err := client.CreateInstanceResourceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验室
//
// @param request - CreateLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLaboratoryResponse
func (client *Client) CreateLaboratoryWithOptions(request *CreateLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketCount)) {
		body["BucketCount"] = request.BucketCount
	}

	if !tea.BoolValue(util.IsUnset(request.BucketType)) {
		body["BucketType"] = request.BucketType
	}

	if !tea.BoolValue(util.IsUnset(request.Buckets)) {
		body["Buckets"] = request.Buckets
	}

	if !tea.BoolValue(util.IsUnset(request.DebugCrowdId)) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建实验室
//
// @param request - CreateLaboratoryRequest
//
// @return CreateLaboratoryResponse
func (client *Client) CreateLaboratory(request *CreateLaboratoryRequest) (_result *CreateLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLaboratoryResponse{}
	_body, _err := client.CreateLaboratoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建层。
//
// @param request - CreateLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayerResponse
func (client *Client) CreateLayerWithOptions(request *CreateLayerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LaboratoryId)) {
		body["LaboratoryId"] = request.LaboratoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLayer"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/layers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLayerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLayerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建层。
//
// @param request - CreateLayerRequest
//
// @return CreateLayerResponse
func (client *Client) CreateLayer(request *CreateLayerRequest) (_result *CreateLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLayerResponse{}
	_body, _err := client.CreateLayerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建参数。
//
// @param request - CreateParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParamResponse
func (client *Client) CreateParamWithOptions(request *CreateParamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateParam"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/params"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateParamResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateParamResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建参数。
//
// @param request - CreateParamRequest
//
// @return CreateParamResponse
func (client *Client) CreateParam(request *CreateParamRequest) (_result *CreateParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateParamResponse{}
	_body, _err := client.CreateParamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源规则
//
// @param request - CreateResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceRuleResponse
func (client *Client) CreateResourceRuleWithOptions(request *CreateResourceRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricOperationType)) {
		body["MetricOperationType"] = request.MetricOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.MetricPullInfo)) {
		body["MetricPullInfo"] = request.MetricPullInfo
	}

	if !tea.BoolValue(util.IsUnset(request.MetricPullPeriod)) {
		body["MetricPullPeriod"] = request.MetricPullPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RuleComputingDefinition)) {
		body["RuleComputingDefinition"] = request.RuleComputingDefinition
	}

	if !tea.BoolValue(util.IsUnset(request.RuleItems)) {
		body["RuleItems"] = request.RuleItems
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceRule"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateResourceRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateResourceRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建资源规则
//
// @param request - CreateResourceRuleRequest
//
// @return CreateResourceRuleResponse
func (client *Client) CreateResourceRule(request *CreateResourceRuleRequest) (_result *CreateResourceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceRuleResponse{}
	_body, _err := client.CreateResourceRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源规则条目
//
// @param request - CreateResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceRuleItemResponse
func (client *Client) CreateResourceRuleItemWithOptions(ResourceRuleId *string, request *CreateResourceRuleItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceRuleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxValue)) {
		body["MaxValue"] = request.MaxValue
	}

	if !tea.BoolValue(util.IsUnset(request.MinValue)) {
		body["MinValue"] = request.MinValue
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceRuleItem"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId)) + "/items"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateResourceRuleItemResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateResourceRuleItemResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建资源规则条目
//
// @param request - CreateResourceRuleItemRequest
//
// @return CreateResourceRuleItemResponse
func (client *Client) CreateResourceRuleItem(ResourceRuleId *string, request *CreateResourceRuleItemRequest) (_result *CreateResourceRuleItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceRuleItemResponse{}
	_body, _err := client.CreateResourceRuleItemWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建场景
//
// @param request - CreateSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneResponse
func (client *Client) CreateSceneWithOptions(request *CreateSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Flows)) {
		body["Flows"] = request.Flows
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScene"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/scenes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSceneResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSceneResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建场景
//
// @param request - CreateSceneRequest
//
// @return CreateSceneResponse
func (client *Client) CreateScene(request *CreateSceneRequest) (_result *CreateSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSceneResponse{}
	_body, _err := client.CreateSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 在指定人群下创建子人群。
//
// @param request - CreateSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCrowdResponse
func (client *Client) CreateSubCrowdWithOptions(CrowdId *string, request *CreateSubCrowdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSubCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubCrowd"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId)) + "/subcrowds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSubCrowdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSubCrowdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 在指定人群下创建子人群。
//
// @param request - CreateSubCrowdRequest
//
// @return CreateSubCrowdResponse
func (client *Client) CreateSubCrowd(CrowdId *string, request *CreateSubCrowdRequest) (_result *CreateSubCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSubCrowdResponse{}
	_body, _err := client.CreateSubCrowdWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据表。
//
// @param request - CreateTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableMetaResponse
func (client *Client) CreateTableMetaWithOptions(request *CreateTableMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTableMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Module)) {
		body["Module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTableMeta"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tablemetas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTableMetaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTableMetaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建数据表。
//
// @param request - CreateTableMetaRequest
//
// @return CreateTableMetaResponse
func (client *Client) CreateTableMeta(request *CreateTableMetaRequest) (_result *CreateTableMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTableMetaResponse{}
	_body, _err := client.CreateTableMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流量调控目标
//
// @param request - CreateTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlTargetResponse
func (client *Client) CreateTrafficControlTargetWithOptions(request *CreateTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Event)) {
		body["Event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionArray)) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionExpress)) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionType)) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewProductRegulation)) {
		body["NewProductRegulation"] = request.NewProductRegulation
	}

	if !tea.BoolValue(util.IsUnset(request.RecallName)) {
		body["RecallName"] = request.RecallName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StatisPeriod)) {
		body["StatisPeriod"] = request.StatisPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ToleranceValue)) {
		body["ToleranceValue"] = request.ToleranceValue
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlTaskId)) {
		body["TrafficControlTaskId"] = request.TrafficControlTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建流量调控目标
//
// @param request - CreateTrafficControlTargetRequest
//
// @return CreateTrafficControlTargetResponse
func (client *Client) CreateTrafficControlTarget(request *CreateTrafficControlTargetRequest) (_result *CreateTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrafficControlTargetResponse{}
	_body, _err := client.CreateTrafficControlTargetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流量调控任务
//
// @param request - CreateTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlTaskResponse
func (client *Client) CreateTrafficControlTaskWithOptions(request *CreateTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BehaviorTableMetaId)) {
		body["BehaviorTableMetaId"] = request.BehaviorTableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.ControlGranularity)) {
		body["ControlGranularity"] = request.ControlGranularity
	}

	if !tea.BoolValue(util.IsUnset(request.ControlLogic)) {
		body["ControlLogic"] = request.ControlLogic
	}

	if !tea.BoolValue(util.IsUnset(request.ControlType)) {
		body["ControlType"] = request.ControlType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionTime)) {
		body["ExecutionTime"] = request.ExecutionTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionArray)) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionExpress)) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionType)) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTableMetaId)) {
		body["ItemTableMetaId"] = request.ItemTableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PreExperimentIds)) {
		body["PreExperimentIds"] = request.PreExperimentIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProdExperimentIds)) {
		body["ProdExperimentIds"] = request.ProdExperimentIds
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StatisBehaviorConditionArray)) {
		body["StatisBehaviorConditionArray"] = request.StatisBehaviorConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.StatisBehaviorConditionExpress)) {
		body["StatisBehaviorConditionExpress"] = request.StatisBehaviorConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.StatisBehaviorConditionType)) {
		body["StatisBehaviorConditionType"] = request.StatisBehaviorConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlTargets)) {
		body["TrafficControlTargets"] = request.TrafficControlTargets
	}

	if !tea.BoolValue(util.IsUnset(request.UserConditionArray)) {
		body["UserConditionArray"] = request.UserConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.UserConditionExpress)) {
		body["UserConditionExpress"] = request.UserConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.UserConditionType)) {
		body["UserConditionType"] = request.UserConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.UserTableMetaId)) {
		body["UserTableMetaId"] = request.UserTableMetaId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建流量调控任务
//
// @param request - CreateTrafficControlTaskRequest
//
// @return CreateTrafficControlTaskResponse
func (client *Client) CreateTrafficControlTask(request *CreateTrafficControlTaskRequest) (_result *CreateTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrafficControlTaskResponse{}
	_body, _err := client.CreateTrafficControlTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对指定资源规则中的计算公式进行调试
//
// @param tmpReq - DebugResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugResourceRuleResponse
func (client *Client) DebugResourceRuleWithOptions(ResourceRuleId *string, tmpReq *DebugResourceRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DebugResourceRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DebugResourceRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MetricInfo)) {
		request.MetricInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricInfo, tea.String("MetricInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricInfoShrink)) {
		query["MetricInfo"] = request.MetricInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugResourceRule"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId)) + "/action/debug"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DebugResourceRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DebugResourceRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 对指定资源规则中的计算公式进行调试
//
// @param request - DebugResourceRuleRequest
//
// @return DebugResourceRuleResponse
func (client *Client) DebugResourceRule(ResourceRuleId *string, request *DebugResourceRuleRequest) (_result *DebugResourceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DebugResourceRuleResponse{}
	_body, _err := client.DebugResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定AB实验指标。
//
// @param request - DeleteABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABMetricResponse
func (client *Client) DeleteABMetricWithOptions(ABMetricId *string, request *DeleteABMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABMetric"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetrics/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteABMetricResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteABMetricResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定AB实验指标。
//
// @param request - DeleteABMetricRequest
//
// @return DeleteABMetricResponse
func (client *Client) DeleteABMetric(ABMetricId *string, request *DeleteABMetricRequest) (_result *DeleteABMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABMetricResponse{}
	_body, _err := client.DeleteABMetricWithOptions(ABMetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AB实验指标组。
//
// @param request - DeleteABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABMetricGroupResponse
func (client *Client) DeleteABMetricGroupWithOptions(ABMetricGroupId *string, request *DeleteABMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABMetricGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricGroupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteABMetricGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteABMetricGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除AB实验指标组。
//
// @param request - DeleteABMetricGroupRequest
//
// @return DeleteABMetricGroupResponse
func (client *Client) DeleteABMetricGroup(ABMetricGroupId *string, request *DeleteABMetricGroupRequest) (_result *DeleteABMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABMetricGroupResponse{}
	_body, _err := client.DeleteABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定人群。
//
// @param request - DeleteCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowdWithOptions(CrowdId *string, request *DeleteCrowdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCrowd"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteCrowdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteCrowdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定人群。
//
// @param request - DeleteCrowdRequest
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowd(CrowdId *string, request *DeleteCrowdRequest) (_result *DeleteCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCrowdResponse{}
	_body, _err := client.DeleteCrowdWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定推荐引擎配置。
//
// @param request - DeleteEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEngineConfigResponse
func (client *Client) DeleteEngineConfigWithOptions(EngineConfigId *string, request *DeleteEngineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEngineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEngineConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(EngineConfigId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEngineConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEngineConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定推荐引擎配置。
//
// @param request - DeleteEngineConfigRequest
//
// @return DeleteEngineConfigResponse
func (client *Client) DeleteEngineConfig(EngineConfigId *string, request *DeleteEngineConfigRequest) (_result *DeleteEngineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEngineConfigResponse{}
	_body, _err := client.DeleteEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实验。
//
// @param request - DeleteExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithOptions(ExperimentId *string, request *DeleteExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除实验。
//
// @param request - DeleteExperimentRequest
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperiment(ExperimentId *string, request *DeleteExperimentRequest) (_result *DeleteExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定实验组。
//
// @param request - DeleteExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentGroupResponse
func (client *Client) DeleteExperimentGroupWithOptions(ExperimentGroupId *string, request *DeleteExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentGroupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定实验组。
//
// @param request - DeleteExperimentGroupRequest
//
// @return DeleteExperimentGroupResponse
func (client *Client) DeleteExperimentGroup(ExperimentGroupId *string, request *DeleteExperimentGroupRequest) (_result *DeleteExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentGroupResponse{}
	_body, _err := client.DeleteExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定实例下的指定配置资源。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResourceResponse
func (client *Client) DeleteInstanceResourceWithOptions(InstanceId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceResource"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteInstanceResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteInstanceResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定实例下的指定配置资源。
//
// @return DeleteInstanceResourceResponse
func (client *Client) DeleteInstanceResource(InstanceId *string, ResourceId *string) (_result *DeleteInstanceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResourceResponse{}
	_body, _err := client.DeleteInstanceResourceWithOptions(InstanceId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实验室。
//
// @param request - DeleteLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLaboratoryResponse
func (client *Client) DeleteLaboratoryWithOptions(LaboratoryId *string, request *DeleteLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories/" + tea.StringValue(openapiutil.GetEncodeParam(LaboratoryId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除实验室。
//
// @param request - DeleteLaboratoryRequest
//
// @return DeleteLaboratoryResponse
func (client *Client) DeleteLaboratory(LaboratoryId *string, request *DeleteLaboratoryRequest) (_result *DeleteLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLaboratoryResponse{}
	_body, _err := client.DeleteLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除层。
//
// @param request - DeleteLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayerWithOptions(LayerId *string, request *DeleteLayerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayer"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/layers/" + tea.StringValue(openapiutil.GetEncodeParam(LayerId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLayerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLayerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除层。
//
// @param request - DeleteLayerRequest
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayer(LayerId *string, request *DeleteLayerRequest) (_result *DeleteLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLayerResponse{}
	_body, _err := client.DeleteLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定参数。
//
// @param request - DeleteParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParamResponse
func (client *Client) DeleteParamWithOptions(ParamId *string, request *DeleteParamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteParam"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/params/" + tea.StringValue(openapiutil.GetEncodeParam(ParamId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteParamResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteParamResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定参数。
//
// @param request - DeleteParamRequest
//
// @return DeleteParamResponse
func (client *Client) DeleteParam(ParamId *string, request *DeleteParamRequest) (_result *DeleteParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteParamResponse{}
	_body, _err := client.DeleteParamWithOptions(ParamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源规则
//
// @param request - DeleteResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceRuleResponse
func (client *Client) DeleteResourceRuleWithOptions(ResourceRuleId *string, request *DeleteResourceRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceRule"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteResourceRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteResourceRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除资源规则
//
// @param request - DeleteResourceRuleRequest
//
// @return DeleteResourceRuleResponse
func (client *Client) DeleteResourceRule(ResourceRuleId *string, request *DeleteResourceRuleRequest) (_result *DeleteResourceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceRuleResponse{}
	_body, _err := client.DeleteResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源规则条目
//
// @param request - DeleteResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceRuleItemResponse
func (client *Client) DeleteResourceRuleItemWithOptions(ResourceRuleId *string, ResourceRuleItemId *string, request *DeleteResourceRuleItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceRuleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceRuleItem"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId)) + "/items/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleItemId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteResourceRuleItemResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteResourceRuleItemResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除资源规则条目
//
// @param request - DeleteResourceRuleItemRequest
//
// @return DeleteResourceRuleItemResponse
func (client *Client) DeleteResourceRuleItem(ResourceRuleId *string, ResourceRuleItemId *string, request *DeleteResourceRuleItemRequest) (_result *DeleteResourceRuleItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceRuleItemResponse{}
	_body, _err := client.DeleteResourceRuleItemWithOptions(ResourceRuleId, ResourceRuleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneResponse
func (client *Client) DeleteSceneWithOptions(SceneId *string, request *DeleteSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScene"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(SceneId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSceneResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSceneResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除场景
//
// @param request - DeleteSceneRequest
//
// @return DeleteSceneResponse
func (client *Client) DeleteScene(SceneId *string, request *DeleteSceneRequest) (_result *DeleteSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSceneResponse{}
	_body, _err := client.DeleteSceneWithOptions(SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定人群下的指定子人群。
//
// @param request - DeleteSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubCrowdResponse
func (client *Client) DeleteSubCrowdWithOptions(CrowdId *string, SubCrowdId *string, request *DeleteSubCrowdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSubCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubCrowd"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId)) + "/subcrowds/" + tea.StringValue(openapiutil.GetEncodeParam(SubCrowdId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSubCrowdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSubCrowdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定人群下的指定子人群。
//
// @param request - DeleteSubCrowdRequest
//
// @return DeleteSubCrowdResponse
func (client *Client) DeleteSubCrowd(CrowdId *string, SubCrowdId *string, request *DeleteSubCrowdRequest) (_result *DeleteSubCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSubCrowdResponse{}
	_body, _err := client.DeleteSubCrowdWithOptions(CrowdId, SubCrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据表。
//
// @param request - DeleteTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMetaWithOptions(TableMetaId *string, request *DeleteTableMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTableMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTableMeta"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tablemetas/" + tea.StringValue(openapiutil.GetEncodeParam(TableMetaId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTableMetaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTableMetaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除数据表。
//
// @param request - DeleteTableMetaRequest
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMeta(TableMetaId *string, request *DeleteTableMetaRequest) (_result *DeleteTableMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableMetaResponse{}
	_body, _err := client.DeleteTableMetaWithOptions(TableMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流量调控目标
//
// @param request - DeleteTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlTargetResponse
func (client *Client) DeleteTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *DeleteTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新流量调控目标
//
// @param request - DeleteTrafficControlTargetRequest
//
// @return DeleteTrafficControlTargetResponse
func (client *Client) DeleteTrafficControlTarget(TrafficControlTargetId *string, request *DeleteTrafficControlTargetRequest) (_result *DeleteTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrafficControlTargetResponse{}
	_body, _err := client.DeleteTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的流量调控任务
//
// @param request - DeleteTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlTaskResponse
func (client *Client) DeleteTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *DeleteTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定的流量调控任务
//
// @param request - DeleteTrafficControlTaskRequest
//
// @return DeleteTrafficControlTaskResponse
func (client *Client) DeleteTrafficControlTask(TrafficControlTaskId *string, request *DeleteTrafficControlTaskRequest) (_result *DeleteTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrafficControlTaskResponse{}
	_body, _err := client.DeleteTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 产生流量调控的相关代码
//
// @param request - GenerateTrafficControlTaskCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTrafficControlTaskCodeResponse
func (client *Client) GenerateTrafficControlTaskCodeWithOptions(TrafficControlTaskId *string, request *GenerateTrafficControlTaskCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateTrafficControlTaskCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateTrafficControlTaskCode"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/generatecode"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateTrafficControlTaskCodeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateTrafficControlTaskCodeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 产生流量调控的相关代码
//
// @param request - GenerateTrafficControlTaskCodeRequest
//
// @return GenerateTrafficControlTaskCodeResponse
func (client *Client) GenerateTrafficControlTaskCode(TrafficControlTaskId *string, request *GenerateTrafficControlTaskCodeRequest) (_result *GenerateTrafficControlTaskCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTrafficControlTaskCodeResponse{}
	_body, _err := client.GenerateTrafficControlTaskCodeWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 产生流量调控的相关召回配置
//
// @param request - GenerateTrafficControlTaskConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTrafficControlTaskConfigResponse
func (client *Client) GenerateTrafficControlTaskConfigWithOptions(TrafficControlTaskId *string, request *GenerateTrafficControlTaskConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateTrafficControlTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateTrafficControlTaskConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/generateconfig"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateTrafficControlTaskConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateTrafficControlTaskConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 产生流量调控的相关召回配置
//
// @param request - GenerateTrafficControlTaskConfigRequest
//
// @return GenerateTrafficControlTaskConfigResponse
func (client *Client) GenerateTrafficControlTaskConfig(TrafficControlTaskId *string, request *GenerateTrafficControlTaskConfigRequest) (_result *GenerateTrafficControlTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTrafficControlTaskConfigResponse{}
	_body, _err := client.GenerateTrafficControlTaskConfigWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AB Test实验指标详细信息。
//
// @param request - GetABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetABMetricResponse
func (client *Client) GetABMetricWithOptions(ABMetricId *string, request *GetABMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetABMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetABMetric"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetrics/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetABMetricResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetABMetricResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取AB Test实验指标详细信息。
//
// @param request - GetABMetricRequest
//
// @return GetABMetricResponse
func (client *Client) GetABMetric(ABMetricId *string, request *GetABMetricRequest) (_result *GetABMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetABMetricResponse{}
	_body, _err := client.GetABMetricWithOptions(ABMetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AB实验指标组详细信息。
//
// @param request - GetABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetABMetricGroupResponse
func (client *Client) GetABMetricGroupWithOptions(ABMetricGroupId *string, request *GetABMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetABMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetABMetricGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricGroupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetABMetricGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetABMetricGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取AB实验指标组详细信息。
//
// @param request - GetABMetricGroupRequest
//
// @return GetABMetricGroupResponse
func (client *Client) GetABMetricGroup(ABMetricGroupId *string, request *GetABMetricGroupRequest) (_result *GetABMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetABMetricGroupResponse{}
	_body, _err := client.GetABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定计算任务详细信息。
//
// @param request - GetCalculationJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCalculationJobResponse
func (client *Client) GetCalculationJobWithOptions(CalculationJobId *string, request *GetCalculationJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCalculationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCalculationJob"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/calculationjobs/" + tea.StringValue(openapiutil.GetEncodeParam(CalculationJobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetCalculationJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetCalculationJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定计算任务详细信息。
//
// @param request - GetCalculationJobRequest
//
// @return GetCalculationJobResponse
func (client *Client) GetCalculationJob(CalculationJobId *string, request *GetCalculationJobRequest) (_result *GetCalculationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCalculationJobResponse{}
	_body, _err := client.GetCalculationJobWithOptions(CalculationJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取引擎配置详细信息。
//
// @param request - GetEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEngineConfigResponse
func (client *Client) GetEngineConfigWithOptions(EngineConfigId *string, request *GetEngineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEngineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEngineConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(EngineConfigId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEngineConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEngineConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取引擎配置详细信息。
//
// @param request - GetEngineConfigRequest
//
// @return GetEngineConfigResponse
func (client *Client) GetEngineConfig(EngineConfigId *string, request *GetEngineConfigRequest) (_result *GetEngineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEngineConfigResponse{}
	_body, _err := client.GetEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验详细信息。
//
// @param request - GetExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithOptions(ExperimentId *string, request *GetExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实验详细信息。
//
// @param request - GetExperimentRequest
//
// @return GetExperimentResponse
func (client *Client) GetExperiment(ExperimentId *string, request *GetExperimentRequest) (_result *GetExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定实验组详细信息。
//
// @param request - GetExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentGroupResponse
func (client *Client) GetExperimentGroupWithOptions(ExperimentGroupId *string, request *GetExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentGroupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定实验组详细信息。
//
// @param request - GetExperimentGroupRequest
//
// @return GetExperimentGroupResponse
func (client *Client) GetExperimentGroup(ExperimentGroupId *string, request *GetExperimentGroupRequest) (_result *GetExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentGroupResponse{}
	_body, _err := client.GetExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务详细信息。
//
// @param request - GetFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConsistencyCheckJobResponse
func (client *Client) GetFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId *string, request *GetFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFeatureConsistencyCheckJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeatureConsistencyCheckJob"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureConsistencyCheckJobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetFeatureConsistencyCheckJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetFeatureConsistencyCheckJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取特征一致性检查任务详细信息。
//
// @param request - GetFeatureConsistencyCheckJobRequest
//
// @return GetFeatureConsistencyCheckJobResponse
func (client *Client) GetFeatureConsistencyCheckJob(FeatureConsistencyCheckJobId *string, request *GetFeatureConsistencyCheckJobRequest) (_result *GetFeatureConsistencyCheckJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureConsistencyCheckJobResponse{}
	_body, _err := client.GetFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征一致性检测配置详情。
//
// @param request - GetFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConsistencyCheckJobConfigResponse
func (client *Client) GetFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId *string, request *GetFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeatureConsistencyCheckJobConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureConsistencyCheckJobConfigId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取特征一致性检测配置详情。
//
// @param request - GetFeatureConsistencyCheckJobConfigRequest
//
// @return GetFeatureConsistencyCheckJobConfigResponse
func (client *Client) GetFeatureConsistencyCheckJobConfig(FeatureConsistencyCheckJobConfigId *string, request *GetFeatureConsistencyCheckJobConfigRequest) (_result *GetFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.GetFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定推荐全链路深度定制开发平台实例信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定推荐全链路深度定制开发平台实例信息。
//
// @return GetInstanceResponse
func (client *Client) GetInstance(InstanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定实例下指定资源的详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResourceResponse
func (client *Client) GetInstanceResourceWithOptions(InstanceId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceResource"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定实例下指定资源的详细信息。
//
// @return GetInstanceResourceResponse
func (client *Client) GetInstanceResource(InstanceId *string, ResourceId *string) (_result *GetInstanceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResourceResponse{}
	_body, _err := client.GetInstanceResourceWithOptions(InstanceId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据源下指定表的详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResourceTableResponse
func (client *Client) GetInstanceResourceTableWithOptions(InstanceId *string, ResourceId *string, TableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResourceTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceResourceTable"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(TableName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceResourceTableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceResourceTableResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取数据源下指定表的详细信息。
//
// @return GetInstanceResourceTableResponse
func (client *Client) GetInstanceResourceTable(InstanceId *string, ResourceId *string, TableName *string) (_result *GetInstanceResourceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResourceTableResponse{}
	_body, _err := client.GetInstanceResourceTableWithOptions(InstanceId, ResourceId, TableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验室详细信息。
//
// @param request - GetLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLaboratoryResponse
func (client *Client) GetLaboratoryWithOptions(LaboratoryId *string, request *GetLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories/" + tea.StringValue(openapiutil.GetEncodeParam(LaboratoryId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实验室详细信息。
//
// @param request - GetLaboratoryRequest
//
// @return GetLaboratoryResponse
func (client *Client) GetLaboratory(LaboratoryId *string, request *GetLaboratoryRequest) (_result *GetLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLaboratoryResponse{}
	_body, _err := client.GetLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取层详细信息。
//
// @param request - GetLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerResponse
func (client *Client) GetLayerWithOptions(LayerId *string, request *GetLayerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayer"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/layers/" + tea.StringValue(openapiutil.GetEncodeParam(LayerId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLayerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLayerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取层详细信息。
//
// @param request - GetLayerRequest
//
// @return GetLayerResponse
func (client *Client) GetLayer(LayerId *string, request *GetLayerRequest) (_result *GetLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLayerResponse{}
	_body, _err := client.GetLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源规则详细信息
//
// @param request - GetResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceRuleResponse
func (client *Client) GetResourceRuleWithOptions(ResourceRuleId *string, request *GetResourceRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceRule"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取资源规则详细信息
//
// @param request - GetResourceRuleRequest
//
// @return GetResourceRuleResponse
func (client *Client) GetResourceRule(ResourceRuleId *string, request *GetResourceRuleRequest) (_result *GetResourceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceRuleResponse{}
	_body, _err := client.GetResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取场景详细信息
//
// @param request - GetSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneResponse
func (client *Client) GetSceneWithOptions(SceneId *string, request *GetSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScene"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(SceneId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSceneResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSceneResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取场景详细信息
//
// @param request - GetSceneRequest
//
// @return GetSceneResponse
func (client *Client) GetScene(SceneId *string, request *GetSceneRequest) (_result *GetSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSceneResponse{}
	_body, _err := client.GetSceneWithOptions(SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定人群下的指定子人群的详细信息。
//
// @param request - GetSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubCrowdResponse
func (client *Client) GetSubCrowdWithOptions(CrowdId *string, SubCrowdId *string, request *GetSubCrowdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSubCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubCrowd"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId)) + "/subcrowds/" + tea.StringValue(openapiutil.GetEncodeParam(SubCrowdId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSubCrowdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSubCrowdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定人群下的指定子人群的详细信息。
//
// @param request - GetSubCrowdRequest
//
// @return GetSubCrowdResponse
func (client *Client) GetSubCrowd(CrowdId *string, SubCrowdId *string, request *GetSubCrowdRequest) (_result *GetSubCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSubCrowdResponse{}
	_body, _err := client.GetSubCrowdWithOptions(CrowdId, SubCrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据表详细信息。
//
// @param request - GetTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableMetaResponse
func (client *Client) GetTableMetaWithOptions(TableMetaId *string, request *GetTableMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableMeta"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tablemetas/" + tea.StringValue(openapiutil.GetEncodeParam(TableMetaId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTableMetaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTableMetaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取数据表详细信息。
//
// @param request - GetTableMetaRequest
//
// @return GetTableMetaResponse
func (client *Client) GetTableMeta(TableMetaId *string, request *GetTableMetaRequest) (_result *GetTableMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableMetaResponse{}
	_body, _err := client.GetTableMetaWithOptions(TableMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流量调控目标详情
//
// @param request - GetTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTargetResponse
func (client *Client) GetTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *GetTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取流量调控目标详情
//
// @param request - GetTrafficControlTargetRequest
//
// @return GetTrafficControlTargetResponse
func (client *Client) GetTrafficControlTarget(TrafficControlTargetId *string, request *GetTrafficControlTargetRequest) (_result *GetTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrafficControlTargetResponse{}
	_body, _err := client.GetTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流量调控任务详情
//
// @param request - GetTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTaskResponse
func (client *Client) GetTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *GetTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlTargetFilter)) {
		query["ControlTargetFilter"] = request.ControlTargetFilter
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取流量调控任务详情
//
// @param request - GetTrafficControlTaskRequest
//
// @return GetTrafficControlTaskResponse
func (client *Client) GetTrafficControlTask(TrafficControlTaskId *string, request *GetTrafficControlTaskRequest) (_result *GetTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrafficControlTaskResponse{}
	_body, _err := client.GetTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流量调控任务的流量详情
//
// @param request - GetTrafficControlTaskTrafficRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTaskTrafficResponse
func (client *Client) GetTrafficControlTaskTrafficWithOptions(TrafficControlTaskId *string, request *GetTrafficControlTaskTrafficRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrafficControlTaskTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrafficControlTaskTraffic"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/trafficinfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrafficControlTaskTrafficResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrafficControlTaskTrafficResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取流量调控任务的流量详情
//
// @param request - GetTrafficControlTaskTrafficRequest
//
// @return GetTrafficControlTaskTrafficResponse
func (client *Client) GetTrafficControlTaskTraffic(TrafficControlTaskId *string, request *GetTrafficControlTaskTrafficRequest) (_result *GetTrafficControlTaskTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrafficControlTaskTrafficResponse{}
	_body, _err := client.GetTrafficControlTaskTrafficWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AB Test实验指标组列表。
//
// @param request - ListABMetricGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABMetricGroupsResponse
func (client *Client) ListABMetricGroupsWithOptions(request *ListABMetricGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABMetricGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.Realtime)) {
		query["Realtime"] = request.Realtime
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListABMetricGroups"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetricgroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListABMetricGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListABMetricGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取AB Test实验指标组列表。
//
// @param request - ListABMetricGroupsRequest
//
// @return ListABMetricGroupsResponse
func (client *Client) ListABMetricGroups(request *ListABMetricGroupsRequest) (_result *ListABMetricGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABMetricGroupsResponse{}
	_body, _err := client.ListABMetricGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AB Test实验指标列表。
//
// @param request - ListABMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABMetricsResponse
func (client *Client) ListABMetricsWithOptions(request *ListABMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Realtime)) {
		query["Realtime"] = request.Realtime
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.TableMetaId)) {
		query["TableMetaId"] = request.TableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListABMetrics"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListABMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListABMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取AB Test实验指标列表。
//
// @param request - ListABMetricsRequest
//
// @return ListABMetricsResponse
func (client *Client) ListABMetrics(request *ListABMetricsRequest) (_result *ListABMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABMetricsResponse{}
	_body, _err := client.ListABMetricsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取计算任务列表。
//
// @param request - ListCalculationJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalculationJobsResponse
func (client *Client) ListCalculationJobsWithOptions(request *ListCalculationJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCalculationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCalculationJobs"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/calculationjobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListCalculationJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListCalculationJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取计算任务列表。
//
// @param request - ListCalculationJobsRequest
//
// @return ListCalculationJobsResponse
func (client *Client) ListCalculationJobs(request *ListCalculationJobsRequest) (_result *ListCalculationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCalculationJobsResponse{}
	_body, _err := client.ListCalculationJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取人群下的所有用户。
//
// @param request - ListCrowdUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrowdUsersResponse
func (client *Client) ListCrowdUsersWithOptions(CrowdId *string, request *ListCrowdUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCrowdUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCrowdUsers"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListCrowdUsersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListCrowdUsersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取人群下的所有用户。
//
// @param request - ListCrowdUsersRequest
//
// @return ListCrowdUsersResponse
func (client *Client) ListCrowdUsers(CrowdId *string, request *ListCrowdUsersRequest) (_result *ListCrowdUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCrowdUsersResponse{}
	_body, _err := client.ListCrowdUsersWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取人群列表。
//
// @param request - ListCrowdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrowdsResponse
func (client *Client) ListCrowdsWithOptions(request *ListCrowdsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCrowdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCrowds"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListCrowdsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListCrowdsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取人群列表。
//
// @param request - ListCrowdsRequest
//
// @return ListCrowdsResponse
func (client *Client) ListCrowds(request *ListCrowdsRequest) (_result *ListCrowdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCrowdsResponse{}
	_body, _err := client.ListCrowdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取引擎配置列表。
//
// @param request - ListEngineConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineConfigsResponse
func (client *Client) ListEngineConfigsWithOptions(request *ListEngineConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEngineConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEngineConfigs"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEngineConfigsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEngineConfigsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取引擎配置列表。
//
// @param request - ListEngineConfigsRequest
//
// @return ListEngineConfigsResponse
func (client *Client) ListEngineConfigs(request *ListEngineConfigsRequest) (_result *ListEngineConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEngineConfigsResponse{}
	_body, _err := client.ListEngineConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验组列表。
//
// @param request - ListExperimentGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentGroupsResponse
func (client *Client) ListExperimentGroupsWithOptions(request *ListExperimentGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExperimentGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		query["LayerId"] = request.LayerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRangeEnd)) {
		query["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRangeStart)) {
		query["TimeRangeStart"] = request.TimeRangeStart
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperimentGroups"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListExperimentGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListExperimentGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实验组列表。
//
// @param request - ListExperimentGroupsRequest
//
// @return ListExperimentGroupsResponse
func (client *Client) ListExperimentGroups(request *ListExperimentGroupsRequest) (_result *ListExperimentGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentGroupsResponse{}
	_body, _err := client.ListExperimentGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验列表。
//
// @param request - ListExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentsResponse
func (client *Client) ListExperimentsWithOptions(request *ListExperimentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExperimentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentGroupId)) {
		query["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperiments"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListExperimentsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListExperimentsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实验列表。
//
// @param request - ListExperimentsRequest
//
// @return ListExperimentsResponse
func (client *Client) ListExperiments(request *ListExperimentsRequest) (_result *ListExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentsResponse{}
	_body, _err := client.ListExperimentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查配置列表。
//
// @param request - ListFeatureConsistencyCheckJobConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobConfigsResponse
func (client *Client) ListFeatureConsistencyCheckJobConfigsWithOptions(request *ListFeatureConsistencyCheckJobConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureConsistencyCheckJobConfigs"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobconfigs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取特征一致性检查配置列表。
//
// @param request - ListFeatureConsistencyCheckJobConfigsRequest
//
// @return ListFeatureConsistencyCheckJobConfigsResponse
func (client *Client) ListFeatureConsistencyCheckJobConfigs(request *ListFeatureConsistencyCheckJobConfigsRequest) (_result *ListFeatureConsistencyCheckJobConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务的特征报表/比对结果。
//
// @param request - ListFeatureConsistencyCheckJobFeatureReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobFeatureReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobFeatureReportsWithOptions(FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobFeatureReportsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobFeatureReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogItemId)) {
		query["LogItemId"] = request.LogItemId
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.LogUserId)) {
		query["LogUserId"] = request.LogUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureConsistencyCheckJobFeatureReports"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureConsistencyCheckJobId)) + "/featurereports"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取特征一致性检查任务的特征报表/比对结果。
//
// @param request - ListFeatureConsistencyCheckJobFeatureReportsRequest
//
// @return ListFeatureConsistencyCheckJobFeatureReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobFeatureReports(FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobFeatureReportsRequest) (_result *ListFeatureConsistencyCheckJobFeatureReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobFeatureReportsWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务分数报表/比对结果。
//
// @param tmpReq - ListFeatureConsistencyCheckJobScoreReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobScoreReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobScoreReportsWithOptions(FeatureConsistencyCheckJobId *string, tmpReq *ListFeatureConsistencyCheckJobScoreReportsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobScoreReportsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureConsistencyCheckJobScoreReportsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExcludeRequestIds)) {
		request.ExcludeRequestIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeRequestIds, tea.String("ExcludeRequestIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeRequestIdsShrink)) {
		query["ExcludeRequestIds"] = request.ExcludeRequestIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureConsistencyCheckJobScoreReports"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureConsistencyCheckJobId)) + "/scorereports"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取特征一致性检查任务分数报表/比对结果。
//
// @param request - ListFeatureConsistencyCheckJobScoreReportsRequest
//
// @return ListFeatureConsistencyCheckJobScoreReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobScoreReports(FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobScoreReportsRequest) (_result *ListFeatureConsistencyCheckJobScoreReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobScoreReportsWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务列表。
//
// @param request - ListFeatureConsistencyCheckJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobsResponse
func (client *Client) ListFeatureConsistencyCheckJobsWithOptions(request *ListFeatureConsistencyCheckJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureConsistencyCheckJobs"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListFeatureConsistencyCheckJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListFeatureConsistencyCheckJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取特征一致性检查任务列表。
//
// @param request - ListFeatureConsistencyCheckJobsRequest
//
// @return ListFeatureConsistencyCheckJobsResponse
func (client *Client) ListFeatureConsistencyCheckJobs(request *ListFeatureConsistencyCheckJobsRequest) (_result *ListFeatureConsistencyCheckJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例下配置的资源列表。
//
// @param request - ListInstanceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResourcesResponse
func (client *Client) ListInstanceResourcesWithOptions(InstanceId *string, request *ListInstanceResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceResources"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstanceResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstanceResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例下配置的资源列表。
//
// @param request - ListInstanceResourcesRequest
//
// @return ListInstanceResourcesResponse
func (client *Client) ListInstanceResources(InstanceId *string, request *ListInstanceResourcesRequest) (_result *ListInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResourcesResponse{}
	_body, _err := client.ListInstanceResourcesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取推荐全链路深度定制开发平台实例信息列表。
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取推荐全链路深度定制开发平台实例信息列表。
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
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

// Summary:
//
// 获取实验室列表。
//
// @param request - ListLaboratoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLaboratoriesResponse
func (client *Client) ListLaboratoriesWithOptions(request *ListLaboratoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLaboratoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLaboratories"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLaboratoriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLaboratoriesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实验室列表。
//
// @param request - ListLaboratoriesRequest
//
// @return ListLaboratoriesResponse
func (client *Client) ListLaboratories(request *ListLaboratoriesRequest) (_result *ListLaboratoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLaboratoriesResponse{}
	_body, _err := client.ListLaboratoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取层列表。
//
// @param request - ListLayersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayersResponse
func (client *Client) ListLayersWithOptions(request *ListLayersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLayersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LaboratoryId)) {
		query["LaboratoryId"] = request.LaboratoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayers"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/layers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLayersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLayersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取层列表。
//
// @param request - ListLayersRequest
//
// @return ListLayersResponse
func (client *Client) ListLayers(request *ListLayersRequest) (_result *ListLayersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLayersResponse{}
	_body, _err := client.ListLayersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取参数列表。
//
// @param request - ListParamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParamsResponse
func (client *Client) ListParamsWithOptions(request *ListParamsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListParams"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/params"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListParamsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListParamsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取参数列表。
//
// @param request - ListParamsRequest
//
// @return ListParamsResponse
func (client *Client) ListParams(request *ListParamsRequest) (_result *ListParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListParamsResponse{}
	_body, _err := client.ListParamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - ListResourceRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRulesResponse
func (client *Client) ListResourceRulesWithOptions(request *ListResourceRulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceRuleId)) {
		query["ResourceRuleId"] = request.ResourceRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRuleName)) {
		query["ResourceRuleName"] = request.ResourceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceRules"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListResourceRulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListResourceRulesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取资源规则列表
//
// @param request - ListResourceRulesRequest
//
// @return ListResourceRulesResponse
func (client *Client) ListResourceRules(request *ListResourceRulesRequest) (_result *ListResourceRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceRulesResponse{}
	_body, _err := client.ListResourceRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取场景列表
//
// @param request - ListScenesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScenesResponse
func (client *Client) ListScenesWithOptions(request *ListScenesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScenes"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/scenes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListScenesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListScenesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取场景列表
//
// @param request - ListScenesRequest
//
// @return ListScenesResponse
func (client *Client) ListScenes(request *ListScenesRequest) (_result *ListScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScenesResponse{}
	_body, _err := client.ListScenesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取人群下的所有子人群。
//
// @param request - ListSubCrowdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubCrowdsResponse
func (client *Client) ListSubCrowdsWithOptions(CrowdId *string, request *ListSubCrowdsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSubCrowdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubCrowds"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId)) + "/subcrowds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSubCrowdsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSubCrowdsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取人群下的所有子人群。
//
// @param request - ListSubCrowdsRequest
//
// @return ListSubCrowdsResponse
func (client *Client) ListSubCrowds(CrowdId *string, request *ListSubCrowdsRequest) (_result *ListSubCrowdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubCrowdsResponse{}
	_body, _err := client.ListSubCrowdsWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据表列表。
//
// @param request - ListTableMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableMetasResponse
func (client *Client) ListTableMetasWithOptions(request *ListTableMetasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTableMetasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Module)) {
		query["Module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTableMetas"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tablemetas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTableMetasResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTableMetasResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取数据表列表。
//
// @param request - ListTableMetasRequest
//
// @return ListTableMetasResponse
func (client *Client) ListTableMetas(request *ListTableMetasRequest) (_result *ListTableMetasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTableMetasResponse{}
	_body, _err := client.ListTableMetasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流量调控任务流量变更的历史列表
//
// @param request - ListTrafficControlTargetTrafficHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficControlTargetTrafficHistoryResponse
func (client *Client) ListTrafficControlTargetTrafficHistoryWithOptions(TrafficControlTargetId *string, request *ListTrafficControlTargetTrafficHistoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrafficControlTargetTrafficHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentGroupId)) {
		query["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrafficControlTargetTrafficHistory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId)) + "/traffichistories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrafficControlTargetTrafficHistoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrafficControlTargetTrafficHistoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取流量调控任务流量变更的历史列表
//
// @param request - ListTrafficControlTargetTrafficHistoryRequest
//
// @return ListTrafficControlTargetTrafficHistoryResponse
func (client *Client) ListTrafficControlTargetTrafficHistory(TrafficControlTargetId *string, request *ListTrafficControlTargetTrafficHistoryRequest) (_result *ListTrafficControlTargetTrafficHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrafficControlTargetTrafficHistoryResponse{}
	_body, _err := client.ListTrafficControlTargetTrafficHistoryWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流量调控列表
//
// @param request - ListTrafficControlTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficControlTasksResponse
func (client *Client) ListTrafficControlTasksWithOptions(request *ListTrafficControlTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrafficControlTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ControlTargetFilter)) {
		query["ControlTargetFilter"] = request.ControlTargetFilter
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlTaskId)) {
		query["TrafficControlTaskId"] = request.TrafficControlTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrafficControlTasks"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrafficControlTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrafficControlTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取流量调控列表
//
// @param request - ListTrafficControlTasksRequest
//
// @return ListTrafficControlTasksResponse
func (client *Client) ListTrafficControlTasks(request *ListTrafficControlTasksRequest) (_result *ListTrafficControlTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrafficControlTasksResponse{}
	_body, _err := client.ListTrafficControlTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上线实验。
//
// @param request - OfflineExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineExperimentResponse
func (client *Client) OfflineExperimentWithOptions(ExperimentId *string, request *OfflineExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OfflineExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/action/offline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OfflineExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OfflineExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上线实验。
//
// @param request - OfflineExperimentRequest
//
// @return OfflineExperimentResponse
func (client *Client) OfflineExperiment(ExperimentId *string, request *OfflineExperimentRequest) (_result *OfflineExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineExperimentResponse{}
	_body, _err := client.OfflineExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下线实验组。
//
// @param request - OfflineExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineExperimentGroupResponse
func (client *Client) OfflineExperimentGroupWithOptions(ExperimentGroupId *string, request *OfflineExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OfflineExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentGroupId)) + "/action/offline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OfflineExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OfflineExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 下线实验组。
//
// @param request - OfflineExperimentGroupRequest
//
// @return OfflineExperimentGroupResponse
func (client *Client) OfflineExperimentGroup(ExperimentGroupId *string, request *OfflineExperimentGroupRequest) (_result *OfflineExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineExperimentGroupResponse{}
	_body, _err := client.OfflineExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下线实验室。
//
// @param request - OfflineLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineLaboratoryResponse
func (client *Client) OfflineLaboratoryWithOptions(LaboratoryId *string, request *OfflineLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OfflineLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories/" + tea.StringValue(openapiutil.GetEncodeParam(LaboratoryId)) + "/action/offline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OfflineLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OfflineLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 下线实验室。
//
// @param request - OfflineLaboratoryRequest
//
// @return OfflineLaboratoryResponse
func (client *Client) OfflineLaboratory(LaboratoryId *string, request *OfflineLaboratoryRequest) (_result *OfflineLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineLaboratoryResponse{}
	_body, _err := client.OfflineLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上线实验
//
// @param request - OnlineExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineExperimentResponse
func (client *Client) OnlineExperimentWithOptions(ExperimentId *string, request *OnlineExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OnlineExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OnlineExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/action/online"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnlineExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnlineExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上线实验
//
// @param request - OnlineExperimentRequest
//
// @return OnlineExperimentResponse
func (client *Client) OnlineExperiment(ExperimentId *string, request *OnlineExperimentRequest) (_result *OnlineExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineExperimentResponse{}
	_body, _err := client.OnlineExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上线实验组。
//
// @param request - OnlineExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineExperimentGroupResponse
func (client *Client) OnlineExperimentGroupWithOptions(ExperimentGroupId *string, request *OnlineExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OnlineExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OnlineExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentGroupId)) + "/action/online"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnlineExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnlineExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上线实验组。
//
// @param request - OnlineExperimentGroupRequest
//
// @return OnlineExperimentGroupResponse
func (client *Client) OnlineExperimentGroup(ExperimentGroupId *string, request *OnlineExperimentGroupRequest) (_result *OnlineExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineExperimentGroupResponse{}
	_body, _err := client.OnlineExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上线实验室。
//
// @param request - OnlineLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineLaboratoryResponse
func (client *Client) OnlineLaboratoryWithOptions(LaboratoryId *string, request *OnlineLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OnlineLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OnlineLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories/" + tea.StringValue(openapiutil.GetEncodeParam(LaboratoryId)) + "/action/online"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnlineLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnlineLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上线实验室。
//
// @param request - OnlineLaboratoryRequest
//
// @return OnlineLaboratoryResponse
func (client *Client) OnlineLaboratory(LaboratoryId *string, request *OnlineLaboratoryRequest) (_result *OnlineLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineLaboratoryResponse{}
	_body, _err := client.OnlineLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推全。
//
// @param request - PushAllExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushAllExperimentResponse
func (client *Client) PushAllExperimentWithOptions(ExperimentId *string, request *PushAllExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushAllExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushAllExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/action/pushall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PushAllExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PushAllExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 推全。
//
// @param request - PushAllExperimentRequest
//
// @return PushAllExperimentResponse
func (client *Client) PushAllExperiment(ExperimentId *string, request *PushAllExperimentRequest) (_result *PushAllExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushAllExperimentResponse{}
	_body, _err := client.PushAllExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送指标到指定资源规则
//
// @param tmpReq - PushResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResourceRuleResponse
func (client *Client) PushResourceRuleWithOptions(ResourceRuleId *string, tmpReq *PushResourceRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushResourceRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushResourceRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MetricInfo)) {
		request.MetricInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricInfo, tea.String("MetricInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricInfoShrink)) {
		query["MetricInfo"] = request.MetricInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushResourceRule"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId)) + "/action/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PushResourceRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PushResourceRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 推送指标到指定资源规则
//
// @param request - PushResourceRuleRequest
//
// @return PushResourceRuleResponse
func (client *Client) PushResourceRule(ResourceRuleId *string, request *PushResourceRuleRequest) (_result *PushResourceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushResourceRuleResponse{}
	_body, _err := client.PushResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流量调控目标的单品调控报表详情。
//
// @param request - QueryTrafficControlTargetItemReportDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTrafficControlTargetItemReportDetailResponse
func (client *Client) QueryTrafficControlTargetItemReportDetailWithOptions(TrafficControlTargetId *string, request *QueryTrafficControlTargetItemReportDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryTrafficControlTargetItemReportDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["Date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTrafficControlTargetItemReportDetail"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId)) + "/itemcontrolreportdetail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryTrafficControlTargetItemReportDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryTrafficControlTargetItemReportDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询流量调控目标的单品调控报表详情。
//
// @param request - QueryTrafficControlTargetItemReportDetailRequest
//
// @return QueryTrafficControlTargetItemReportDetailResponse
func (client *Client) QueryTrafficControlTargetItemReportDetail(TrafficControlTargetId *string, request *QueryTrafficControlTargetItemReportDetailRequest) (_result *QueryTrafficControlTargetItemReportDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTrafficControlTargetItemReportDetailResponse{}
	_body, _err := client.QueryTrafficControlTargetItemReportDetailWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布流量调控任务
//
// @param request - ReleaseTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseTrafficControlTaskResponse
func (client *Client) ReleaseTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *ReleaseTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 发布流量调控任务
//
// @param request - ReleaseTrafficControlTaskRequest
//
// @return ReleaseTrafficControlTaskResponse
func (client *Client) ReleaseTrafficControlTask(TrafficControlTaskId *string, request *ReleaseTrafficControlTaskRequest) (_result *ReleaseTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseTrafficControlTaskResponse{}
	_body, _err := client.ReleaseTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对指标组进行报表。
//
// @param request - ReportABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportABMetricGroupResponse
func (client *Client) ReportABMetricGroupWithOptions(ABMetricGroupId *string, request *ReportABMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReportABMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseExperimentId)) {
		body["BaseExperimentId"] = request.BaseExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionFields)) {
		body["DimensionFields"] = request.DimensionFields
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentGroupId)) {
		body["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentIds)) {
		body["ExperimentIds"] = request.ExperimentIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		body["ReportType"] = request.ReportType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStatisticsMethod)) {
		body["TimeStatisticsMethod"] = request.TimeStatisticsMethod
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportABMetricGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricGroupId)) + "/action/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReportABMetricGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReportABMetricGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 对指标组进行报表。
//
// @param request - ReportABMetricGroupRequest
//
// @return ReportABMetricGroupResponse
func (client *Client) ReportABMetricGroup(ABMetricGroupId *string, request *ReportABMetricGroupRequest) (_result *ReportABMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReportABMetricGroupResponse{}
	_body, _err := client.ReportABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拆分流量调控目标
//
// @param request - SplitTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitTrafficControlTargetResponse
func (client *Client) SplitTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *SplitTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SplitTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SetPoints)) {
		body["SetPoints"] = request.SetPoints
	}

	if !tea.BoolValue(util.IsUnset(request.SetValues)) {
		body["SetValues"] = request.SetValues
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoints)) {
		body["TimePoints"] = request.TimePoints
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SplitTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId)) + "/action/split"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SplitTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SplitTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 拆分流量调控目标
//
// @param request - SplitTrafficControlTargetRequest
//
// @return SplitTrafficControlTargetResponse
func (client *Client) SplitTrafficControlTarget(TrafficControlTargetId *string, request *SplitTrafficControlTargetRequest) (_result *SplitTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitTrafficControlTargetResponse{}
	_body, _err := client.SplitTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启流量调控目标
//
// @param request - StartTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTrafficControlTargetResponse
func (client *Client) StartTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *StartTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId)) + "/action/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 开启流量调控目标
//
// @param request - StartTrafficControlTargetRequest
//
// @return StartTrafficControlTargetResponse
func (client *Client) StartTrafficControlTarget(TrafficControlTargetId *string, request *StartTrafficControlTargetRequest) (_result *StartTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTrafficControlTargetResponse{}
	_body, _err := client.StartTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启流量调控任务
//
// @param request - StartTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTrafficControlTaskResponse
func (client *Client) StartTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *StartTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 开启流量调控任务
//
// @param request - StartTrafficControlTaskRequest
//
// @return StartTrafficControlTaskResponse
func (client *Client) StartTrafficControlTask(TrafficControlTaskId *string, request *StartTrafficControlTaskRequest) (_result *StartTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTrafficControlTaskResponse{}
	_body, _err := client.StartTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止流量调控目标
//
// @param request - StopTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlTargetResponse
func (client *Client) StopTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *StopTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId)) + "/action/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 停止流量调控目标
//
// @param request - StopTrafficControlTargetRequest
//
// @return StopTrafficControlTargetResponse
func (client *Client) StopTrafficControlTarget(TrafficControlTargetId *string, request *StopTrafficControlTargetRequest) (_result *StopTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrafficControlTargetResponse{}
	_body, _err := client.StopTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止流量调控任务
//
// @param request - StopTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlTaskResponse
func (client *Client) StopTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *StopTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 停止流量调控任务
//
// @param request - StopTrafficControlTaskRequest
//
// @return StopTrafficControlTaskResponse
func (client *Client) StopTrafficControlTask(TrafficControlTaskId *string, request *StopTrafficControlTaskRequest) (_result *StopTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrafficControlTaskResponse{}
	_body, _err := client.StopTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步特征一致性检测任务重放日志。
//
// @param request - SyncFeatureConsistencyCheckJobReplayLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncFeatureConsistencyCheckJobReplayLogResponse
func (client *Client) SyncFeatureConsistencyCheckJobReplayLogWithOptions(request *SyncFeatureConsistencyCheckJobReplayLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SyncFeatureConsistencyCheckJobReplayLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContextFeatures)) {
		body["ContextFeatures"] = request.ContextFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureConsistencyCheckJobConfigId)) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.GeneratedFeatures)) {
		body["GeneratedFeatures"] = request.GeneratedFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogItemId)) {
		body["LogItemId"] = request.LogItemId
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestTime)) {
		body["LogRequestTime"] = request.LogRequestTime
	}

	if !tea.BoolValue(util.IsUnset(request.LogUserId)) {
		body["LogUserId"] = request.LogUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RawFeatures)) {
		body["RawFeatures"] = request.RawFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		body["SceneName"] = request.SceneName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncFeatureConsistencyCheckJobReplayLog"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs/action/syncreplaylog"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 同步特征一致性检测任务重放日志。
//
// @param request - SyncFeatureConsistencyCheckJobReplayLogRequest
//
// @return SyncFeatureConsistencyCheckJobReplayLogResponse
func (client *Client) SyncFeatureConsistencyCheckJobReplayLog(request *SyncFeatureConsistencyCheckJobReplayLogRequest) (_result *SyncFeatureConsistencyCheckJobReplayLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
	_body, _err := client.SyncFeatureConsistencyCheckJobReplayLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消指定特征一致性检查正在运行中的任务。
//
// @param request - TerminateFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateFeatureConsistencyCheckJobResponse
func (client *Client) TerminateFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId *string, request *TerminateFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TerminateFeatureConsistencyCheckJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TerminateFeatureConsistencyCheckJob"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureConsistencyCheckJobId)) + "/action/terminate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TerminateFeatureConsistencyCheckJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TerminateFeatureConsistencyCheckJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 取消指定特征一致性检查正在运行中的任务。
//
// @param request - TerminateFeatureConsistencyCheckJobRequest
//
// @return TerminateFeatureConsistencyCheckJobResponse
func (client *Client) TerminateFeatureConsistencyCheckJob(FeatureConsistencyCheckJobId *string, request *TerminateFeatureConsistencyCheckJobRequest) (_result *TerminateFeatureConsistencyCheckJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TerminateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.TerminateFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AB Test实验指标。
//
// @param request - UpdateABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABMetricResponse
func (client *Client) UpdateABMetricWithOptions(ABMetricId *string, request *UpdateABMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		body["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LeftMetricId)) {
		body["LeftMetricId"] = request.LeftMetricId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.Realtime)) {
		body["Realtime"] = request.Realtime
	}

	if !tea.BoolValue(util.IsUnset(request.ResultResourceId)) {
		body["ResultResourceId"] = request.ResultResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RightMetricId)) {
		body["RightMetricId"] = request.RightMetricId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.StatisticsCycle)) {
		body["StatisticsCycle"] = request.StatisticsCycle
	}

	if !tea.BoolValue(util.IsUnset(request.TableMetaId)) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABMetric"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetrics/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateABMetricResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateABMetricResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新AB Test实验指标。
//
// @param request - UpdateABMetricRequest
//
// @return UpdateABMetricResponse
func (client *Client) UpdateABMetric(ABMetricId *string, request *UpdateABMetricRequest) (_result *UpdateABMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABMetricResponse{}
	_body, _err := client.UpdateABMetricWithOptions(ABMetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AB test实验指标组。
//
// @param request - UpdateABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABMetricGroupResponse
func (client *Client) UpdateABMetricGroupWithOptions(ABMetricGroupId *string, request *UpdateABMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ABMetricIds)) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Realtime)) {
		body["Realtime"] = request.Realtime
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABMetricGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/abmetricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ABMetricGroupId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateABMetricGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateABMetricGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新AB test实验指标组。
//
// @param request - UpdateABMetricGroupRequest
//
// @return UpdateABMetricGroupResponse
func (client *Client) UpdateABMetricGroup(ABMetricGroupId *string, request *UpdateABMetricGroupRequest) (_result *UpdateABMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABMetricGroupResponse{}
	_body, _err := client.UpdateABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定人群。
//
// @param request - UpdateCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrowdResponse
func (client *Client) UpdateCrowdWithOptions(CrowdId *string, request *UpdateCrowdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCrowd"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateCrowdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateCrowdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新指定人群。
//
// @param request - UpdateCrowdRequest
//
// @return UpdateCrowdResponse
func (client *Client) UpdateCrowd(CrowdId *string, request *UpdateCrowdRequest) (_result *UpdateCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCrowdResponse{}
	_body, _err := client.UpdateCrowdWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新引擎配置。
//
// @param request - UpdateEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEngineConfigResponse
func (client *Client) UpdateEngineConfigWithOptions(EngineConfigId *string, request *UpdateEngineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEngineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigValue)) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEngineConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/engineconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(EngineConfigId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateEngineConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateEngineConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新引擎配置。
//
// @param request - UpdateEngineConfigRequest
//
// @return UpdateEngineConfigResponse
func (client *Client) UpdateEngineConfig(EngineConfigId *string, request *UpdateEngineConfigRequest) (_result *UpdateEngineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEngineConfigResponse{}
	_body, _err := client.UpdateEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实验。
//
// @param request - UpdateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperimentWithOptions(ExperimentId *string, request *UpdateExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DebugCrowdId)) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FlowPercent)) {
		body["FlowPercent"] = request.FlowPercent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperiment"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateExperimentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateExperimentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新实验。
//
// @param request - UpdateExperimentRequest
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperiment(ExperimentId *string, request *UpdateExperimentRequest) (_result *UpdateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentResponse{}
	_body, _err := client.UpdateExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定实验组。
//
// @param request - UpdateExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentGroupResponse
func (client *Client) UpdateExperimentGroupWithOptions(ExperimentGroupId *string, request *UpdateExperimentGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdId)) {
		body["CrowdId"] = request.CrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdTargetType)) {
		body["CrowdTargetType"] = request.CrowdTargetType
	}

	if !tea.BoolValue(util.IsUnset(request.DebugCrowdId)) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionTimeDuration)) {
		body["DistributionTimeDuration"] = request.DistributionTimeDuration
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionType)) {
		body["DistributionType"] = request.DistributionType
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		body["LayerId"] = request.LayerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NeedAA)) {
		body["NeedAA"] = request.NeedAA
	}

	if !tea.BoolValue(util.IsUnset(request.RandomFlow)) {
		body["RandomFlow"] = request.RandomFlow
	}

	if !tea.BoolValue(util.IsUnset(request.ReservcedBuckets)) {
		body["ReservcedBuckets"] = request.ReservcedBuckets
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentGroup"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentgroups/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentGroupId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateExperimentGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateExperimentGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新指定实验组。
//
// @param request - UpdateExperimentGroupRequest
//
// @return UpdateExperimentGroupResponse
func (client *Client) UpdateExperimentGroup(ExperimentGroupId *string, request *UpdateExperimentGroupRequest) (_result *UpdateExperimentGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentGroupResponse{}
	_body, _err := client.UpdateExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新特征一致性检查配置信息。
//
// @param request - UpdateFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFeatureConsistencyCheckJobConfigResponse
func (client *Client) UpdateFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId *string, request *UpdateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareFeature)) {
		body["CompareFeature"] = request.CompareFeature
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMountPath)) {
		body["DatasetMountPath"] = request.DatasetMountPath
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetType)) {
		body["DatasetType"] = request.DatasetType
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetUri)) {
		body["DatasetUri"] = request.DatasetUri
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRoute)) {
		body["DefaultRoute"] = request.DefaultRoute
	}

	if !tea.BoolValue(util.IsUnset(request.EasServiceName)) {
		body["EasServiceName"] = request.EasServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.EasyRecPackagePath)) {
		body["EasyRecPackagePath"] = request.EasyRecPackagePath
	}

	if !tea.BoolValue(util.IsUnset(request.EasyRecVersion)) {
		body["EasyRecVersion"] = request.EasyRecVersion
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureDisplayExclude)) {
		body["FeatureDisplayExclude"] = request.FeatureDisplayExclude
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureLandingResourceId)) {
		body["FeatureLandingResourceId"] = request.FeatureLandingResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.FeaturePriority)) {
		body["FeaturePriority"] = request.FeaturePriority
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreItemId)) {
		body["FeatureStoreItemId"] = request.FeatureStoreItemId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreModelId)) {
		body["FeatureStoreModelId"] = request.FeatureStoreModelId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreProjectId)) {
		body["FeatureStoreProjectId"] = request.FeatureStoreProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreProjectName)) {
		body["FeatureStoreProjectName"] = request.FeatureStoreProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreSeqFeatureView)) {
		body["FeatureStoreSeqFeatureView"] = request.FeatureStoreSeqFeatureView
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureStoreUserId)) {
		body["FeatureStoreUserId"] = request.FeatureStoreUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FgJarVersion)) {
		body["FgJarVersion"] = request.FgJarVersion
	}

	if !tea.BoolValue(util.IsUnset(request.FgJsonFileName)) {
		body["FgJsonFileName"] = request.FgJsonFileName
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateZip)) {
		body["GenerateZip"] = request.GenerateZip
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsUseFeatureStore)) {
		body["IsUseFeatureStore"] = request.IsUseFeatureStore
	}

	if !tea.BoolValue(util.IsUnset(request.ItemIdField)) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTable)) {
		body["ItemTable"] = request.ItemTable
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTablePartitionField)) {
		body["ItemTablePartitionField"] = request.ItemTablePartitionField
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTablePartitionFieldFormat)) {
		body["ItemTablePartitionFieldFormat"] = request.ItemTablePartitionFieldFormat
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssResourceId)) {
		body["OssResourceId"] = request.OssResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.PredictWorkerCount)) {
		body["PredictWorkerCount"] = request.PredictWorkerCount
	}

	if !tea.BoolValue(util.IsUnset(request.PredictWorkerCpu)) {
		body["PredictWorkerCpu"] = request.PredictWorkerCpu
	}

	if !tea.BoolValue(util.IsUnset(request.PredictWorkerMemory)) {
		body["PredictWorkerMemory"] = request.PredictWorkerMemory
	}

	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		body["SampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchId)) {
		body["SwitchId"] = request.SwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdField)) {
		body["UserIdField"] = request.UserIdField
	}

	if !tea.BoolValue(util.IsUnset(request.UserTable)) {
		body["UserTable"] = request.UserTable
	}

	if !tea.BoolValue(util.IsUnset(request.UserTablePartitionField)) {
		body["UserTablePartitionField"] = request.UserTablePartitionField
	}

	if !tea.BoolValue(util.IsUnset(request.UserTablePartitionFieldFormat)) {
		body["UserTablePartitionFieldFormat"] = request.UserTablePartitionFieldFormat
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowName)) {
		body["WorkflowName"] = request.WorkflowName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFeatureConsistencyCheckJobConfig"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/featureconsistencycheck/jobconfigs/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureConsistencyCheckJobConfigId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新特征一致性检查配置信息。
//
// @param request - UpdateFeatureConsistencyCheckJobConfigRequest
//
// @return UpdateFeatureConsistencyCheckJobConfigResponse
func (client *Client) UpdateFeatureConsistencyCheckJobConfig(FeatureConsistencyCheckJobConfigId *string, request *UpdateFeatureConsistencyCheckJobConfigRequest) (_result *UpdateFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.UpdateFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定实例下指定资源的信息。
//
// @param request - UpdateInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResourceResponse
func (client *Client) UpdateInstanceResourceWithOptions(InstanceId *string, ResourceId *string, request *UpdateInstanceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceResource"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新指定实例下指定资源的信息。
//
// @param request - UpdateInstanceResourceRequest
//
// @return UpdateInstanceResourceResponse
func (client *Client) UpdateInstanceResource(InstanceId *string, ResourceId *string, request *UpdateInstanceResourceRequest) (_result *UpdateInstanceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResourceResponse{}
	_body, _err := client.UpdateInstanceResourceWithOptions(InstanceId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实验室。
//
// @param request - UpdateLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLaboratoryResponse
func (client *Client) UpdateLaboratoryWithOptions(LaboratoryId *string, request *UpdateLaboratoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLaboratoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketCount)) {
		body["BucketCount"] = request.BucketCount
	}

	if !tea.BoolValue(util.IsUnset(request.BucketType)) {
		body["BucketType"] = request.BucketType
	}

	if !tea.BoolValue(util.IsUnset(request.Buckets)) {
		body["Buckets"] = request.Buckets
	}

	if !tea.BoolValue(util.IsUnset(request.DebugCrowdId)) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLaboratory"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/laboratories/" + tea.StringValue(openapiutil.GetEncodeParam(LaboratoryId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLaboratoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLaboratoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新实验室。
//
// @param request - UpdateLaboratoryRequest
//
// @return UpdateLaboratoryResponse
func (client *Client) UpdateLaboratory(LaboratoryId *string, request *UpdateLaboratoryRequest) (_result *UpdateLaboratoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLaboratoryResponse{}
	_body, _err := client.UpdateLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新层。
//
// @param request - UpdateLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLayerResponse
func (client *Client) UpdateLayerWithOptions(LayerId *string, request *UpdateLayerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLayer"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/layers/" + tea.StringValue(openapiutil.GetEncodeParam(LayerId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLayerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLayerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新层。
//
// @param request - UpdateLayerRequest
//
// @return UpdateLayerResponse
func (client *Client) UpdateLayer(LayerId *string, request *UpdateLayerRequest) (_result *UpdateLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLayerResponse{}
	_body, _err := client.UpdateLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新参数。
//
// @param request - UpdateParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParamResponse
func (client *Client) UpdateParamWithOptions(ParamId *string, request *UpdateParamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateParam"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/params/" + tea.StringValue(openapiutil.GetEncodeParam(ParamId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateParamResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateParamResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新参数。
//
// @param request - UpdateParamRequest
//
// @return UpdateParamResponse
func (client *Client) UpdateParam(ParamId *string, request *UpdateParamRequest) (_result *UpdateParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateParamResponse{}
	_body, _err := client.UpdateParamWithOptions(ParamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - UpdateResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceRuleResponse
func (client *Client) UpdateResourceRuleWithOptions(ResourceRuleId *string, request *UpdateResourceRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricOperationType)) {
		body["MetricOperationType"] = request.MetricOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.MetricPullInfo)) {
		body["MetricPullInfo"] = request.MetricPullInfo
	}

	if !tea.BoolValue(util.IsUnset(request.MetricPullPeriod)) {
		body["MetricPullPeriod"] = request.MetricPullPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RuleComputingDefinition)) {
		body["RuleComputingDefinition"] = request.RuleComputingDefinition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceRule"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateResourceRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateResourceRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取资源规则列表
//
// @param request - UpdateResourceRuleRequest
//
// @return UpdateResourceRuleResponse
func (client *Client) UpdateResourceRule(ResourceRuleId *string, request *UpdateResourceRuleRequest) (_result *UpdateResourceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceRuleResponse{}
	_body, _err := client.UpdateResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源规则条目
//
// @param request - UpdateResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceRuleItemResponse
func (client *Client) UpdateResourceRuleItemWithOptions(ResourceRuleId *string, ResourceRuleItemId *string, request *UpdateResourceRuleItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceRuleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxValue)) {
		body["MaxValue"] = request.MaxValue
	}

	if !tea.BoolValue(util.IsUnset(request.MinValue)) {
		body["MinValue"] = request.MinValue
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceRuleItem"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resourcerules/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleId)) + "/items/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceRuleItemId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateResourceRuleItemResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateResourceRuleItemResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新资源规则条目
//
// @param request - UpdateResourceRuleItemRequest
//
// @return UpdateResourceRuleItemResponse
func (client *Client) UpdateResourceRuleItem(ResourceRuleId *string, ResourceRuleItemId *string, request *UpdateResourceRuleItemRequest) (_result *UpdateResourceRuleItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceRuleItemResponse{}
	_body, _err := client.UpdateResourceRuleItemWithOptions(ResourceRuleId, ResourceRuleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新场景
//
// @param request - UpdateSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSceneResponse
func (client *Client) UpdateSceneWithOptions(SceneId *string, request *UpdateSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Flows)) {
		body["Flows"] = request.Flows
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScene"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(SceneId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSceneResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSceneResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新场景
//
// @param request - UpdateSceneRequest
//
// @return UpdateSceneResponse
func (client *Client) UpdateScene(SceneId *string, request *UpdateSceneRequest) (_result *UpdateSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSceneResponse{}
	_body, _err := client.UpdateSceneWithOptions(SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据表详细信息。
//
// @param request - UpdateTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableMetaResponse
func (client *Client) UpdateTableMetaWithOptions(TableMetaId *string, request *UpdateTableMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTableMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Module)) {
		body["Module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTableMeta"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tablemetas/" + tea.StringValue(openapiutil.GetEncodeParam(TableMetaId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTableMetaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTableMetaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取数据表详细信息。
//
// @param request - UpdateTableMetaRequest
//
// @return UpdateTableMetaResponse
func (client *Client) UpdateTableMeta(TableMetaId *string, request *UpdateTableMetaRequest) (_result *UpdateTableMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTableMetaResponse{}
	_body, _err := client.UpdateTableMetaWithOptions(TableMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流量调控目标
//
// @param request - UpdateTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTargetResponse
func (client *Client) UpdateTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *UpdateTrafficControlTargetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTrafficControlTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewParam3)) {
		query["new-param-3"] = request.NewParam3
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Event)) {
		body["Event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionArray)) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionExpress)) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionType)) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewProductRegulation)) {
		body["NewProductRegulation"] = request.NewProductRegulation
	}

	if !tea.BoolValue(util.IsUnset(request.RecallName)) {
		body["RecallName"] = request.RecallName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StatisPeriod)) {
		body["StatisPeriod"] = request.StatisPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ToleranceValue)) {
		body["ToleranceValue"] = request.ToleranceValue
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrafficControlTarget"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltargets/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTargetId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTrafficControlTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTrafficControlTargetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新流量调控目标
//
// @param request - UpdateTrafficControlTargetRequest
//
// @return UpdateTrafficControlTargetResponse
func (client *Client) UpdateTrafficControlTarget(TrafficControlTargetId *string, request *UpdateTrafficControlTargetRequest) (_result *UpdateTrafficControlTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrafficControlTargetResponse{}
	_body, _err := client.UpdateTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流量调控任务
//
// @param request - UpdateTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTaskResponse
func (client *Client) UpdateTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *UpdateTrafficControlTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTrafficControlTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BehaviorTableMetaId)) {
		body["BehaviorTableMetaId"] = request.BehaviorTableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.ControlGranularity)) {
		body["ControlGranularity"] = request.ControlGranularity
	}

	if !tea.BoolValue(util.IsUnset(request.ControlLogic)) {
		body["ControlLogic"] = request.ControlLogic
	}

	if !tea.BoolValue(util.IsUnset(request.ControlType)) {
		body["ControlType"] = request.ControlType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionTime)) {
		body["ExecutionTime"] = request.ExecutionTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionArray)) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionExpress)) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.ItemConditionType)) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTableMetaId)) {
		body["ItemTableMetaId"] = request.ItemTableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PreExperimentIds)) {
		body["PreExperimentIds"] = request.PreExperimentIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProdExperimentIds)) {
		body["ProdExperimentIds"] = request.ProdExperimentIds
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StatisBaeaviorConditionArray)) {
		body["StatisBaeaviorConditionArray"] = request.StatisBaeaviorConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.StatisBehaviorConditionExpress)) {
		body["StatisBehaviorConditionExpress"] = request.StatisBehaviorConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.StatisBehaviorConditionType)) {
		body["StatisBehaviorConditionType"] = request.StatisBehaviorConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlTargets)) {
		body["TrafficControlTargets"] = request.TrafficControlTargets
	}

	if !tea.BoolValue(util.IsUnset(request.UserConditionArray)) {
		body["UserConditionArray"] = request.UserConditionArray
	}

	if !tea.BoolValue(util.IsUnset(request.UserConditionExpress)) {
		body["UserConditionExpress"] = request.UserConditionExpress
	}

	if !tea.BoolValue(util.IsUnset(request.UserConditionType)) {
		body["UserConditionType"] = request.UserConditionType
	}

	if !tea.BoolValue(util.IsUnset(request.UserTableMetaId)) {
		body["UserTableMetaId"] = request.UserTableMetaId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrafficControlTask"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTrafficControlTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTrafficControlTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新流量调控任务
//
// @param request - UpdateTrafficControlTaskRequest
//
// @return UpdateTrafficControlTaskResponse
func (client *Client) UpdateTrafficControlTask(TrafficControlTaskId *string, request *UpdateTrafficControlTaskRequest) (_result *UpdateTrafficControlTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrafficControlTaskResponse{}
	_body, _err := client.UpdateTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流量调控任务的流量参数
//
// @param request - UpdateTrafficControlTaskTrafficRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTaskTrafficResponse
func (client *Client) UpdateTrafficControlTaskTrafficWithOptions(TrafficControlTaskId *string, request *UpdateTrafficControlTaskTrafficRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTrafficControlTaskTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewParam3)) {
		query["new-param-3"] = request.NewParam3
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Traffics)) {
		body["Traffics"] = request.Traffics
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrafficControlTaskTraffic"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trafficcontroltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TrafficControlTaskId)) + "/action/traffic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTrafficControlTaskTrafficResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTrafficControlTaskTrafficResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新流量调控任务的流量参数
//
// @param request - UpdateTrafficControlTaskTrafficRequest
//
// @return UpdateTrafficControlTaskTrafficResponse
func (client *Client) UpdateTrafficControlTaskTraffic(TrafficControlTaskId *string, request *UpdateTrafficControlTaskTrafficRequest) (_result *UpdateTrafficControlTaskTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrafficControlTaskTrafficResponse{}
	_body, _err := client.UpdateTrafficControlTaskTrafficWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传数据
//
// @param request - UploadRecommendationDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRecommendationDataResponse
func (client *Client) UploadRecommendationDataWithOptions(request *UploadRecommendationDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadRecommendationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["DataType"] = request.DataType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadRecommendationData"),
		Version:     tea.String("2022-12-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/recommendationdata/action/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UploadRecommendationDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UploadRecommendationDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上传数据
//
// @param request - UploadRecommendationDataRequest
//
// @return UploadRecommendationDataResponse
func (client *Client) UploadRecommendationData(request *UploadRecommendationDataRequest) (_result *UploadRecommendationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadRecommendationDataResponse{}
	_body, _err := client.UploadRecommendationDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
