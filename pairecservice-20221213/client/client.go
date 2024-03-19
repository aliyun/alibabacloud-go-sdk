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

type ExperimentReportValue struct {
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

type BackflowFeatureConsistencyCheckJobDataRequest struct {
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	InstanceId                         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemFeatures                       *string `json:"ItemFeatures,omitempty" xml:"ItemFeatures,omitempty"`
	LogItemId                          *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	LogRequestId                       *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	LogRequestTime                     *int64  `json:"LogRequestTime,omitempty" xml:"LogRequestTime,omitempty"`
	LogUserId                          *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	SceneName                          *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	Scores                             *string `json:"Scores,omitempty" xml:"Scores,omitempty"`
	UserFeatures                       *string `json:"UserFeatures,omitempty" xml:"UserFeatures,omitempty"`
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

func (s *BackflowFeatureConsistencyCheckJobDataRequest) SetUserFeatures(v string) *BackflowFeatureConsistencyCheckJobDataRequest {
	s.UserFeatures = &v
	return s
}

type BackflowFeatureConsistencyCheckJobDataResponseBody struct {
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
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri  *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri    *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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

type CloneExperimentRequest struct {
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
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
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
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LayerId     *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
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
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// Id of the request
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
	FeatureConsistencyCheckId *string `json:"FeatureConsistencyCheckId,omitempty" xml:"FeatureConsistencyCheckId,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CloneExperimentGroup *bool   `json:"CloneExperimentGroup,omitempty" xml:"CloneExperimentGroup,omitempty"`
	Environment          *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// Id of the request
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

type CreateABMetricRequest struct {
	Definition       *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LeftMetricId     *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator         *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Realtime         *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	RightMetricId    *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	SceneId          *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	StatisticsCycle  *int32  `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	TableMetaId      *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Realtime    *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	ABMetricGroupId *string `json:"ABMetricGroupId,omitempty" xml:"ABMetricGroupId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
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
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Users       *string `json:"Users,omitempty" xml:"Users,omitempty"`
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
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// Id of the request
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

type CreateExperimentRequest struct {
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DebugCrowdId      *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers        *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	FlowPercent       *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
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
	Config                   *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdId                  *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	DebugCrowdId             *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers               *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistributionTimeDuration *int32  `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	DistributionType         *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	Filter                   *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LayerId                  *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedAA                   *bool   `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	ReservedBuckets          *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
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

func (s *CreateExperimentGroupRequest) SetReservedBuckets(v string) *CreateExperimentGroupRequest {
	s.ReservedBuckets = &v
	return s
}

type CreateExperimentGroupResponseBody struct {
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// Id of the request
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
	Environment                        *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	InstanceId                         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SamplingDuration                   *int32  `json:"SamplingDuration,omitempty" xml:"SamplingDuration,omitempty"`
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
	FeatureConsistencyCheckJobId *string `json:"FeatureConsistencyCheckJobId,omitempty" xml:"FeatureConsistencyCheckJobId,omitempty"`
	RequestId                    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CompareFeature                *bool    `json:"CompareFeature,omitempty" xml:"CompareFeature,omitempty"`
	EasServiceName                *string  `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	EasyRecPackagePath            *string  `json:"EasyRecPackagePath,omitempty" xml:"EasyRecPackagePath,omitempty"`
	EasyRecVersion                *string  `json:"EasyRecVersion,omitempty" xml:"EasyRecVersion,omitempty"`
	FeatureDisplayExclude         *string  `json:"FeatureDisplayExclude,omitempty" xml:"FeatureDisplayExclude,omitempty"`
	FeatureLandingResourceId      *string  `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
	FeaturePriority               *string  `json:"FeaturePriority,omitempty" xml:"FeaturePriority,omitempty"`
	FeatureStoreItemId            *string  `json:"FeatureStoreItemId,omitempty" xml:"FeatureStoreItemId,omitempty"`
	FeatureStoreModelId           *string  `json:"FeatureStoreModelId,omitempty" xml:"FeatureStoreModelId,omitempty"`
	FeatureStoreProjectId         *string  `json:"FeatureStoreProjectId,omitempty" xml:"FeatureStoreProjectId,omitempty"`
	FeatureStoreProjectName       *string  `json:"FeatureStoreProjectName,omitempty" xml:"FeatureStoreProjectName,omitempty"`
	FeatureStoreSeqFeatureView    *string  `json:"FeatureStoreSeqFeatureView,omitempty" xml:"FeatureStoreSeqFeatureView,omitempty"`
	FeatureStoreUserId            *string  `json:"FeatureStoreUserId,omitempty" xml:"FeatureStoreUserId,omitempty"`
	FgJarVersion                  *string  `json:"FgJarVersion,omitempty" xml:"FgJarVersion,omitempty"`
	FgJsonFileName                *string  `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	GenerateZip                   *bool    `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	InstanceId                    *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemIdField                   *string  `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	ItemTable                     *string  `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	ItemTablePartitionField       *string  `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	ItemTablePartitionFieldFormat *string  `json:"ItemTablePartitionFieldFormat,omitempty" xml:"ItemTablePartitionFieldFormat,omitempty"`
	Name                          *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OssResourceId                 *string  `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	SampleRate                    *float64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SceneId                       *string  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ServiceId                     *string  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UseFeatureStore               *bool    `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
	UserIdField                   *string  `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	UserTable                     *string  `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	UserTablePartitionField       *string  `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	UserTablePartitionFieldFormat *string  `json:"UserTablePartitionFieldFormat,omitempty" xml:"UserTablePartitionFieldFormat,omitempty"`
	WorkflowName                  *string  `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
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

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetSampleRate(v float64) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.SampleRate = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetSceneId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.SceneId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetServiceId(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.ServiceId = &v
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

func (s *CreateFeatureConsistencyCheckJobConfigRequest) SetWorkflowName(v string) *CreateFeatureConsistencyCheckJobConfigRequest {
	s.WorkflowName = &v
	return s
}

type CreateFeatureConsistencyCheckJobConfigResponseBody struct {
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	RequestId                          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BucketCount  *int32  `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	BucketType   *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Buckets      *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment  *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId      *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// Id of the request
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
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// Id of the request
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
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	ParamId *int64 `json:"ParamId,omitempty" xml:"ParamId,omitempty"`
	// Id of the request
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

type CreateSceneRequest struct {
	Description *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*CreateSceneRequestFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	InstanceId  *string                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
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
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Users      *string `json:"Users,omitempty" xml:"Users,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields      []*CreateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	InstanceId  *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Module      *string                         `json:"Module,omitempty" xml:"Module,omitempty"`
	Name        *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceId  *string                         `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TableName   *string                         `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	DataType         *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	IsDimensionField *bool   `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	IsPartitionField *string `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	Meaning          *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DeleteABMetricRequest struct {
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

type DeleteExperimentRequest struct {
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

type DeleteSceneRequest struct {
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

type GetABMetricRequest struct {
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
	Definition        *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LeftMetricId      *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator          *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Realtime          *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultResourceId  *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	RightMetricId     *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	SceneId           *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName         *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	StatisticsCycle   *int32  `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	TableMetaId       *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ABMetricIds   *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	ABMetricNames *string `json:"ABMetricNames,omitempty" xml:"ABMetricNames,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner         *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Realtime      *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId       *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	ABMetricId   *string   `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	ABMetricName *string   `json:"ABMetricName,omitempty" xml:"ABMetricName,omitempty"`
	BizDate      *string   `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	Config       *string   `json:"Config,omitempty" xml:"Config,omitempty"`
	GmtRanTime   *string   `json:"GmtRanTime,omitempty" xml:"GmtRanTime,omitempty"`
	JobMessage   []*string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty" type:"Repeated"`
	JobSource    *string   `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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

type GetExperimentRequest struct {
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
	AliasExperimentId *string `json:"AliasExperimentId,omitempty" xml:"AliasExperimentId,omitempty"`
	Buckets           *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DebugCrowdId      *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers        *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	FlowPercent       *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GmtCreateTime     *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	LaboratoryId      *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	LayerId           *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Config                   *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdId                  *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	DebugCrowdId             *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers               *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistributionTimeDuration *int32  `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	DistributionType         *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	Filter                   *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	LaboratoryId             *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	LayerId                  *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedAA                   *bool   `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	Owner                    *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	SceneId         *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Config                               *string   `json:"Config,omitempty" xml:"Config,omitempty"`
	FeatureConsistencyCheckJobConfigId   *string   `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	FeatureConsistencyCheckJobConfigName *string   `json:"FeatureConsistencyCheckJobConfigName,omitempty" xml:"FeatureConsistencyCheckJobConfigName,omitempty"`
	GmtEndTime                           *string   `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	GmtStartTime                         *string   `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Logs                                 []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	RequestId                            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
	CompareFeature                *bool   `json:"CompareFeature,omitempty" xml:"CompareFeature,omitempty"`
	EasServiceName                *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	EasyRecPackagePath            *string `json:"EasyRecPackagePath,omitempty" xml:"EasyRecPackagePath,omitempty"`
	EasyRecVersion                *string `json:"EasyRecVersion,omitempty" xml:"EasyRecVersion,omitempty"`
	FeatureDisplayExclude         *string `json:"FeatureDisplayExclude,omitempty" xml:"FeatureDisplayExclude,omitempty"`
	FeatureLandingResourceId      *string `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
	FeatureLandingResourceUri     *string `json:"FeatureLandingResourceUri,omitempty" xml:"FeatureLandingResourceUri,omitempty"`
	FeaturePriority               *string `json:"FeaturePriority,omitempty" xml:"FeaturePriority,omitempty"`
	FeatureStoreItemId            *string `json:"FeatureStoreItemId,omitempty" xml:"FeatureStoreItemId,omitempty"`
	FeatureStoreModelId           *string `json:"FeatureStoreModelId,omitempty" xml:"FeatureStoreModelId,omitempty"`
	FeatureStoreProjectId         *string `json:"FeatureStoreProjectId,omitempty" xml:"FeatureStoreProjectId,omitempty"`
	FeatureStoreProjectName       *string `json:"FeatureStoreProjectName,omitempty" xml:"FeatureStoreProjectName,omitempty"`
	FeatureStoreSeqFeatureView    *string `json:"FeatureStoreSeqFeatureView,omitempty" xml:"FeatureStoreSeqFeatureView,omitempty"`
	FeatureStoreUserId            *string `json:"FeatureStoreUserId,omitempty" xml:"FeatureStoreUserId,omitempty"`
	FgJarVersion                  *string `json:"FgJarVersion,omitempty" xml:"FgJarVersion,omitempty"`
	FgJsonFileName                *string `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	GenerateZip                   *bool   `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	GmtCreateTime                 *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime               *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemIdField                   *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	ItemTable                     *string `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	ItemTablePartitionField       *string `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	ItemTablePartitionFieldFormat *string `json:"ItemTablePartitionFieldFormat,omitempty" xml:"ItemTablePartitionFieldFormat,omitempty"`
	LatestJobGmtSamplingEndTime   *string `json:"LatestJobGmtSamplingEndTime,omitempty" xml:"LatestJobGmtSamplingEndTime,omitempty"`
	LatestJobGmtSamplingStartTime *string `json:"LatestJobGmtSamplingStartTime,omitempty" xml:"LatestJobGmtSamplingStartTime,omitempty"`
	LatestJobId                   *string `json:"LatestJobId,omitempty" xml:"LatestJobId,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket                     *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssResourceId                 *string `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	RequestId                     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SampleRate                    *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SceneId                       *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName                     *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceId                     *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName                   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status                        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseFeatureStore               *bool   `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
	UserIdField                   *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	UserTable                     *string `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	UserTablePartitionField       *string `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	UserTablePartitionFieldFormat *string `json:"UserTablePartitionFieldFormat,omitempty" xml:"UserTablePartitionFieldFormat,omitempty"`
	WorkflowName                  *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
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

func (s *GetFeatureConsistencyCheckJobConfigResponseBody) SetWorkflowName(v string) *GetFeatureConsistencyCheckJobConfigResponseBody {
	s.WorkflowName = &v
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
	ChargeType      *string                        `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommodityCode   *string                        `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Config          *GetInstanceResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	ExpiredTime     *string                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreateTime   *string                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceId      *string                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId       *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *string                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string                        `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	Type          *string                `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	Type          *string                `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	Type          *string                `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Group           *string `json:"Group,omitempty" xml:"Group,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	Fields    []*GetInstanceResourceTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TableName *string                                       `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	IsDimensionField *bool   `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	IsPartitionField *bool   `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	Meaning          *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	BucketCount  *int32  `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	BucketType   *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Buckets      *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	CrowdId      *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment  *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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

type GetSceneRequest struct {
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
	Description *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*GetSceneResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	Name        *string                      `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
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
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
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
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	Quantity      *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Users     *string `json:"Users,omitempty" xml:"Users,omitempty"`
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
	CanDelete       *bool                             `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	Config          *string                           `json:"Config,omitempty" xml:"Config,omitempty"`
	Description     *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields          []*GetTableMetaResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	GmtCreateTime   *string                           `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtImportedTime *string                           `json:"GmtImportedTime,omitempty" xml:"GmtImportedTime,omitempty"`
	GmtModifiedTime *string                           `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Module          *string                           `json:"Module,omitempty" xml:"Module,omitempty"`
	Name            *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId      *string                           `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TableMetaId     *string                           `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	TableName       *string                           `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type            *string                           `json:"Type,omitempty" xml:"Type,omitempty"`
	Url             *string                           `json:"Url,omitempty" xml:"Url,omitempty"`
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
	IsDimensionField *bool   `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	Meaning          *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

type ListABMetricGroupsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Realtime   *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	ABMetricGroupId *string `json:"ABMetricGroupId,omitempty" xml:"ABMetricGroupId,omitempty"`
	ABMetricIds     *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	ABMetricNames   *string `json:"ABMetricNames,omitempty" xml:"ABMetricNames,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner           *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Realtime        *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	SceneId         *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Realtime    *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ABMetrics  []*ListABMetricsResponseBodyABMetrics `json:"ABMetrics,omitempty" xml:"ABMetrics,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	ABMetricId        *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	Definition        *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LeftMetricId      *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator          *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Realtime          *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	ResultResourceId  *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	RightMetricId     *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	SceneId           *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName         *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	StatisticsCycle   *int32  `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	TableMetaId       *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	ABMetricName     *string   `json:"ABMetricName,omitempty" xml:"ABMetricName,omitempty"`
	BizDate          *string   `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	CalculationJobId *string   `json:"CalculationJobId,omitempty" xml:"CalculationJobId,omitempty"`
	Config           *string   `json:"Config,omitempty" xml:"Config,omitempty"`
	GmtRanTime       *string   `json:"GmtRanTime,omitempty" xml:"GmtRanTime,omitempty"`
	JobMessage       []*string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty" type:"Repeated"`
	JobSource        *string   `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CrowdId       *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	Label         *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Quantity      *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Users         *string `json:"Users,omitempty" xml:"Users,omitempty"`
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

type ListExperimentGroupsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LayerId    *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListExperimentGroupsRequest) SetStatus(v string) *ListExperimentGroupsRequest {
	s.Status = &v
	return s
}

type ListExperimentGroupsResponseBody struct {
	ExperimentGroups []*ListExperimentGroupsResponseBodyExperimentGroups `json:"ExperimentGroups,omitempty" xml:"ExperimentGroups,omitempty" type:"Repeated"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Config                   *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdId                  *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	DebugCrowdId             *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers               *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistributionTimeDuration *int32  `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	DistributionType         *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	ExperimentGroupId        *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	Filter                   *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	LaboratoryId             *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	LayerId                  *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedAA                   *bool   `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	Owner                    *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ReservedBuckets          *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	SceneId                  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Query             *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AliasExperimentId *string `json:"AliasExperimentId,omitempty" xml:"AliasExperimentId,omitempty"`
	Buckets           *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DebugCrowdId      *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers        *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	ExperimentId      *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	FlowPercent       *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GmtCreateTime     *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	LaboratoryId      *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	LayerId           *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId           *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
	RequestId                      *string                                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                     *int64                                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CompareFeature                     *bool   `json:"CompareFeature,omitempty" xml:"CompareFeature,omitempty"`
	EasServiceName                     *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	EasyRecPackagePath                 *string `json:"EasyRecPackagePath,omitempty" xml:"EasyRecPackagePath,omitempty"`
	EasyRecVersion                     *string `json:"EasyRecVersion,omitempty" xml:"EasyRecVersion,omitempty"`
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	FeatureDisplayExclude              *string `json:"FeatureDisplayExclude,omitempty" xml:"FeatureDisplayExclude,omitempty"`
	FeatureLandingResourceId           *string `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
	FeatureLandingResourceUri          *string `json:"FeatureLandingResourceUri,omitempty" xml:"FeatureLandingResourceUri,omitempty"`
	FeaturePriority                    *string `json:"FeaturePriority,omitempty" xml:"FeaturePriority,omitempty"`
	FeatureStoreItemId                 *string `json:"FeatureStoreItemId,omitempty" xml:"FeatureStoreItemId,omitempty"`
	FeatureStoreModelId                *string `json:"FeatureStoreModelId,omitempty" xml:"FeatureStoreModelId,omitempty"`
	FeatureStoreProjectId              *string `json:"FeatureStoreProjectId,omitempty" xml:"FeatureStoreProjectId,omitempty"`
	FeatureStoreProjectName            *string `json:"FeatureStoreProjectName,omitempty" xml:"FeatureStoreProjectName,omitempty"`
	FeatureStoreSeqFeatureView         *string `json:"FeatureStoreSeqFeatureView,omitempty" xml:"FeatureStoreSeqFeatureView,omitempty"`
	FeatureStoreUserId                 *string `json:"FeatureStoreUserId,omitempty" xml:"FeatureStoreUserId,omitempty"`
	FgJarVersion                       *string `json:"FgJarVersion,omitempty" xml:"FgJarVersion,omitempty"`
	FgJsonFileName                     *string `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	GenerateZip                        *bool   `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	GmtCreateTime                      *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime                    *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemIdField                        *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	ItemTable                          *string `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	ItemTablePartitionField            *string `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	ItemTablePartitionFieldFormat      *string `json:"ItemTablePartitionFieldFormat,omitempty" xml:"ItemTablePartitionFieldFormat,omitempty"`
	LatestJobGmtSamplingEndTime        *string `json:"LatestJobGmtSamplingEndTime,omitempty" xml:"LatestJobGmtSamplingEndTime,omitempty"`
	LatestJobGmtSamplingStartTime      *string `json:"LatestJobGmtSamplingStartTime,omitempty" xml:"LatestJobGmtSamplingStartTime,omitempty"`
	LatestJobId                        *string `json:"LatestJobId,omitempty" xml:"LatestJobId,omitempty"`
	Name                               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket                          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssResourceId                      *string `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	SampleRate                         *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SceneId                            *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName                          *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceId                          *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName                        *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status                             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseFeatureStore                    *string `json:"UseFeatureStore,omitempty" xml:"UseFeatureStore,omitempty"`
	UserIdField                        *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	UserTable                          *string `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	UserTablePartitionField            *string `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	UserTablePartitionFieldFormat      *string `json:"UserTablePartitionFieldFormat,omitempty" xml:"UserTablePartitionFieldFormat,omitempty"`
	WorkflowName                       *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
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

func (s *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs) SetWorkflowName(v string) *ListFeatureConsistencyCheckJobConfigsResponseBodyFeatureConsistencyCheckConfigs {
	s.WorkflowName = &v
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LogItemId    *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	LogUserId    *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
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
	DataPath             *string                                                                         `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	OssPath              *string                                                                         `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	ReportsOfFeatureDiff []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff `json:"ReportsOfFeatureDiff,omitempty" xml:"ReportsOfFeatureDiff,omitempty" type:"Repeated"`
	RequestId            *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	FeatureName  *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	LogItemId    *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	LogUserId    *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	OfflineValue *string `json:"OfflineValue,omitempty" xml:"OfflineValue,omitempty"`
	OnlineValue  *string `json:"OnlineValue,omitempty" xml:"OnlineValue,omitempty"`
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
	InstanceId        *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	DataPath           *string                                                                     `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	OssPath            *string                                                                     `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	ReportsOfScoreDiff []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff `json:"ReportsOfScoreDiff,omitempty" xml:"ReportsOfScoreDiff,omitempty" type:"Repeated"`
	RequestId          *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	LogItemId       *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	LogRequestId    *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	LogUserId       *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	ScoreDiff       *string `json:"ScoreDiff,omitempty" xml:"ScoreDiff,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId                   *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                  *string                                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Config                               *string   `json:"Config,omitempty" xml:"Config,omitempty"`
	FeatureConsistencyCheckJobConfigId   *string   `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	FeatureConsistencyCheckJobConfigName *string   `json:"FeatureConsistencyCheckJobConfigName,omitempty" xml:"FeatureConsistencyCheckJobConfigName,omitempty"`
	FeatureConsistencyCheckJobId         *string   `json:"FeatureConsistencyCheckJobId,omitempty" xml:"FeatureConsistencyCheckJobId,omitempty"`
	GmtEndTime                           *string   `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	GmtStartTime                         *string   `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Logs                                 []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	Status                               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*ListInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Config        *string `json:"Config,omitempty" xml:"Config,omitempty"`
	GmtCreateAt   *string `json:"GmtCreateAt,omitempty" xml:"GmtCreateAt,omitempty"`
	GmtModifiedAt *string `json:"GmtModifiedAt,omitempty" xml:"GmtModifiedAt,omitempty"`
	Group         *string `json:"Group,omitempty" xml:"Group,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri           *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Instances  []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	ChargeType      *string                                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommodityCode   *string                                   `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Config          *ListInstancesResponseBodyInstancesConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	ExpiredTime     *string                                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreateTime   *string                                   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceId      *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	Type          *string                `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	Type          *string                `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ComponentCode *string                `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	Meta          map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
	Type          *string                `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	BucketCount  *int32  `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	BucketType   *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Buckets      *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	CrowdId      *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment  *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId      *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	LayerId      *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId      *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Environment     *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamId         *string `json:"ParamId,omitempty" xml:"ParamId,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type ListScenesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenes     []*ListScenesResponseBodyScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Description *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*ListScenesResponseBodyScenesFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	Name        *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId     *string                              `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
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
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCrowds  []*ListSubCrowdsResponseBodySubCrowds `json:"SubCrowds,omitempty" xml:"SubCrowds,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	Quantity      *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubCrowdId    *string `json:"SubCrowdId,omitempty" xml:"SubCrowdId,omitempty"`
	Users         *string `json:"Users,omitempty" xml:"Users,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Module     *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TableMetas []*ListTableMetasResponseBodyTableMetas `json:"TableMetas,omitempty" xml:"TableMetas,omitempty" type:"Repeated"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CanDelete       *bool                                         `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	Config          *string                                       `json:"Config,omitempty" xml:"Config,omitempty"`
	Description     *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields          []*ListTableMetasResponseBodyTableMetasFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	GmtCreateTime   *string                                       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtImportedTime *string                                       `json:"GmtImportedTime,omitempty" xml:"GmtImportedTime,omitempty"`
	GmtModifiedTime *string                                       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Module          *string                                       `json:"Module,omitempty" xml:"Module,omitempty"`
	Name            *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceId      *string                                       `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TableMetaId     *string                                       `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	TableName       *string                                       `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type            *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Url             *string                                       `json:"Url,omitempty" xml:"Url,omitempty"`
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
	IsDimensionField *bool   `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	Meaning          *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

type OfflineExperimentRequest struct {
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

type ReportABMetricGroupRequest struct {
	BaseExperimentId     *string `json:"BaseExperimentId,omitempty" xml:"BaseExperimentId,omitempty"`
	DimensionFields      *string `json:"DimensionFields,omitempty" xml:"DimensionFields,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ExperimentGroupId    *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	ExperimentIds        *string `json:"ExperimentIds,omitempty" xml:"ExperimentIds,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ReportType           *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	SceneId              *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
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
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type SyncFeatureConsistencyCheckJobReplayLogRequest struct {
	ContextFeatures                    *string `json:"ContextFeatures,omitempty" xml:"ContextFeatures,omitempty"`
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	GeneratedFeatures                  *string `json:"GeneratedFeatures,omitempty" xml:"GeneratedFeatures,omitempty"`
	InstanceId                         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LogItemId                          *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	LogRequestId                       *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	LogRequestTime                     *int64  `json:"LogRequestTime,omitempty" xml:"LogRequestTime,omitempty"`
	LogUserId                          *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	RawFeatures                        *string `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty"`
	SceneName                          *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
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
	Definition       *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LeftMetricId     *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator         *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Realtime         *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	RightMetricId    *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	SceneId          *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	StatisticsCycle  *int32  `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	TableMetaId      *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Realtime    *bool   `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type UpdateExperimentRequest struct {
	Config       *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowPercent  *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Config                   *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdId                  *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	DebugCrowdId             *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers               *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistributionTimeDuration *int32  `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	DistributionType         *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	Filter                   *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LayerId                  *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedAA                   *bool   `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	ReservcedBuckets         *string `json:"ReservcedBuckets,omitempty" xml:"ReservcedBuckets,omitempty"`
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

func (s *UpdateExperimentGroupRequest) SetReservcedBuckets(v string) *UpdateExperimentGroupRequest {
	s.ReservcedBuckets = &v
	return s
}

type UpdateExperimentGroupResponseBody struct {
	// Id of the request
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
	CompareFeature                *bool    `json:"CompareFeature,omitempty" xml:"CompareFeature,omitempty"`
	EasServiceName                *string  `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	EasyRecPackagePath            *string  `json:"EasyRecPackagePath,omitempty" xml:"EasyRecPackagePath,omitempty"`
	EasyRecVersion                *string  `json:"EasyRecVersion,omitempty" xml:"EasyRecVersion,omitempty"`
	FeatureDisplayExclude         *string  `json:"FeatureDisplayExclude,omitempty" xml:"FeatureDisplayExclude,omitempty"`
	FeatureLandingResourceId      *string  `json:"FeatureLandingResourceId,omitempty" xml:"FeatureLandingResourceId,omitempty"`
	FeaturePriority               *string  `json:"FeaturePriority,omitempty" xml:"FeaturePriority,omitempty"`
	FeatureStoreItemId            *string  `json:"FeatureStoreItemId,omitempty" xml:"FeatureStoreItemId,omitempty"`
	FeatureStoreModelId           *string  `json:"FeatureStoreModelId,omitempty" xml:"FeatureStoreModelId,omitempty"`
	FeatureStoreProjectId         *string  `json:"FeatureStoreProjectId,omitempty" xml:"FeatureStoreProjectId,omitempty"`
	FeatureStoreProjectName       *string  `json:"FeatureStoreProjectName,omitempty" xml:"FeatureStoreProjectName,omitempty"`
	FeatureStoreSeqFeatureView    *string  `json:"FeatureStoreSeqFeatureView,omitempty" xml:"FeatureStoreSeqFeatureView,omitempty"`
	FeatureStoreUserId            *string  `json:"FeatureStoreUserId,omitempty" xml:"FeatureStoreUserId,omitempty"`
	FgJarVersion                  *string  `json:"FgJarVersion,omitempty" xml:"FgJarVersion,omitempty"`
	FgJsonFileName                *string  `json:"FgJsonFileName,omitempty" xml:"FgJsonFileName,omitempty"`
	GenerateZip                   *bool    `json:"GenerateZip,omitempty" xml:"GenerateZip,omitempty"`
	InstanceId                    *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsUseFeatureStore             *bool    `json:"IsUseFeatureStore,omitempty" xml:"IsUseFeatureStore,omitempty"`
	ItemIdField                   *string  `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	ItemTable                     *string  `json:"ItemTable,omitempty" xml:"ItemTable,omitempty"`
	ItemTablePartitionField       *string  `json:"ItemTablePartitionField,omitempty" xml:"ItemTablePartitionField,omitempty"`
	ItemTablePartitionFieldFormat *string  `json:"ItemTablePartitionFieldFormat,omitempty" xml:"ItemTablePartitionFieldFormat,omitempty"`
	Name                          *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OssResourceId                 *string  `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	SampleRate                    *float64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SceneId                       *string  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ServiceId                     *string  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UserIdField                   *string  `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
	UserTable                     *string  `json:"UserTable,omitempty" xml:"UserTable,omitempty"`
	UserTablePartitionField       *string  `json:"UserTablePartitionField,omitempty" xml:"UserTablePartitionField,omitempty"`
	UserTablePartitionFieldFormat *string  `json:"UserTablePartitionFieldFormat,omitempty" xml:"UserTablePartitionFieldFormat,omitempty"`
	WorkflowName                  *string  `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
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

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetSampleRate(v float64) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.SampleRate = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetSceneId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetServiceId(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.ServiceId = &v
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

func (s *UpdateFeatureConsistencyCheckJobConfigRequest) SetWorkflowName(v string) *UpdateFeatureConsistencyCheckJobConfigRequest {
	s.WorkflowName = &v
	return s
}

type UpdateFeatureConsistencyCheckJobConfigResponseBody struct {
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
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Uri    *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	BucketCount  *int32  `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	BucketType   *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Buckets      *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment  *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type UpdateSceneRequest struct {
	Description *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*UpdateSceneRequestFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	InstanceId  *string                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
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
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
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
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields      []*UpdateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	InstanceId  *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Module      *string                         `json:"Module,omitempty" xml:"Module,omitempty"`
	Name        *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceId  *string                         `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TableName   *string                         `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	DataType         *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	IsDimensionField *bool   `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	IsPartitionField *string `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	Meaning          *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CheckInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CloneExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CloneExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CloneLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateCalculationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateFeatureConsistencyCheckJobConfigWithOptions(request *CreateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareFeature)) {
		body["CompareFeature"] = request.CompareFeature
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

	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		body["SampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
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

	if !tea.BoolValue(util.IsUnset(request.WorkflowName)) {
		body["WorkflowName"] = request.WorkflowName
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
	_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateSubCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteSubCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetCalculationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetInstanceResourceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetSubCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListABMetricGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListABMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListCalculationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListCrowdUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListCrowdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
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
	_result = &ListExperimentGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListFeatureConsistencyCheckJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListLaboratoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListLayersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListSubCrowdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListTableMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OfflineExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OfflineExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OfflineLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnlineExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnlineExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnlineLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &PushAllExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ReportABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &TerminateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId *string, request *UpdateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareFeature)) {
		body["CompareFeature"] = request.CompareFeature
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

	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		body["SampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
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

	if !tea.BoolValue(util.IsUnset(request.WorkflowName)) {
		body["WorkflowName"] = request.WorkflowName
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
	_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UploadRecommendationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
