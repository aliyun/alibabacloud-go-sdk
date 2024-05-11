// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckLayerRequest struct {
	ParamNames *string `json:"ParamNames,omitempty" xml:"ParamNames,omitempty"`
}

func (s CheckLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckLayerRequest) GoString() string {
	return s.String()
}

func (s *CheckLayerRequest) SetParamNames(v string) *CheckLayerRequest {
	s.ParamNames = &v
	return s
}

type CheckLayerResponseBody struct {
	CheckResults []*CheckLayerResponseBodyCheckResults `json:"CheckResults,omitempty" xml:"CheckResults,omitempty" type:"Repeated"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckLayerResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLayerResponseBody) SetCheckResults(v []*CheckLayerResponseBodyCheckResults) *CheckLayerResponseBody {
	s.CheckResults = v
	return s
}

func (s *CheckLayerResponseBody) SetRequestId(v string) *CheckLayerResponseBody {
	s.RequestId = &v
	return s
}

type CheckLayerResponseBodyCheckResults struct {
	ExperimentId   *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	ParamName      *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s CheckLayerResponseBodyCheckResults) String() string {
	return tea.Prettify(s)
}

func (s CheckLayerResponseBodyCheckResults) GoString() string {
	return s.String()
}

func (s *CheckLayerResponseBodyCheckResults) SetExperimentId(v string) *CheckLayerResponseBodyCheckResults {
	s.ExperimentId = &v
	return s
}

func (s *CheckLayerResponseBodyCheckResults) SetExperimentName(v string) *CheckLayerResponseBodyCheckResults {
	s.ExperimentName = &v
	return s
}

func (s *CheckLayerResponseBodyCheckResults) SetParamName(v string) *CheckLayerResponseBodyCheckResults {
	s.ParamName = &v
	return s
}

type CheckLayerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckLayerResponse) GoString() string {
	return s.String()
}

func (s *CheckLayerResponse) SetHeaders(v map[string]*string) *CheckLayerResponse {
	s.Headers = v
	return s
}

func (s *CheckLayerResponse) SetStatusCode(v int32) *CheckLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckLayerResponse) SetBody(v *CheckLayerResponseBody) *CheckLayerResponse {
	s.Body = v
	return s
}

type CreateCrowdRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateCrowdRequest) SetLabel(v string) *CreateCrowdRequest {
	s.Label = &v
	return s
}

func (s *CreateCrowdRequest) SetName(v string) *CreateCrowdRequest {
	s.Name = &v
	return s
}

func (s *CreateCrowdRequest) SetUsers(v string) *CreateCrowdRequest {
	s.Users = &v
	return s
}

func (s *CreateCrowdRequest) SetWorkspaceId(v string) *CreateCrowdRequest {
	s.WorkspaceId = &v
	return s
}

type CreateCrowdResponseBody struct {
	CrowdId   *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
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

type CreateDomainRequest struct {
	// This parameter is required.
	BucketType  *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Condition   *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CrowdIds    *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers  *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Flow        *int64  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// This parameter is required.
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetBucketType(v string) *CreateDomainRequest {
	s.BucketType = &v
	return s
}

func (s *CreateDomainRequest) SetCondition(v string) *CreateDomainRequest {
	s.Condition = &v
	return s
}

func (s *CreateDomainRequest) SetCrowdIds(v string) *CreateDomainRequest {
	s.CrowdIds = &v
	return s
}

func (s *CreateDomainRequest) SetDebugUsers(v string) *CreateDomainRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateDomainRequest) SetDescription(v string) *CreateDomainRequest {
	s.Description = &v
	return s
}

func (s *CreateDomainRequest) SetFlow(v int64) *CreateDomainRequest {
	s.Flow = &v
	return s
}

func (s *CreateDomainRequest) SetLayerId(v string) *CreateDomainRequest {
	s.LayerId = &v
	return s
}

func (s *CreateDomainRequest) SetName(v string) *CreateDomainRequest {
	s.Name = &v
	return s
}

func (s *CreateDomainRequest) SetProjectId(v string) *CreateDomainRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDomainRequest) SetWorkspaceId(v string) *CreateDomainRequest {
	s.WorkspaceId = &v
	return s
}

type CreateDomainResponseBody struct {
	DomainId  *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetDomainId(v string) *CreateDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateExperimentRequest struct {
	// This parameter is required.
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Condition  *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// This parameter is required.
	CoreMetricId *string `json:"CoreMetricId,omitempty" xml:"CoreMetricId,omitempty"`
	CrowdIds     *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Flow    *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// This parameter is required.
	FocusMetricIds *string `json:"FocusMetricIds,omitempty" xml:"FocusMetricIds,omitempty"`
	// This parameter is required.
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) SetBucketType(v string) *CreateExperimentRequest {
	s.BucketType = &v
	return s
}

func (s *CreateExperimentRequest) SetCondition(v string) *CreateExperimentRequest {
	s.Condition = &v
	return s
}

func (s *CreateExperimentRequest) SetCoreMetricId(v string) *CreateExperimentRequest {
	s.CoreMetricId = &v
	return s
}

func (s *CreateExperimentRequest) SetCrowdIds(v string) *CreateExperimentRequest {
	s.CrowdIds = &v
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

func (s *CreateExperimentRequest) SetEndTime(v string) *CreateExperimentRequest {
	s.EndTime = &v
	return s
}

func (s *CreateExperimentRequest) SetFlow(v int32) *CreateExperimentRequest {
	s.Flow = &v
	return s
}

func (s *CreateExperimentRequest) SetFocusMetricIds(v string) *CreateExperimentRequest {
	s.FocusMetricIds = &v
	return s
}

func (s *CreateExperimentRequest) SetLayerId(v string) *CreateExperimentRequest {
	s.LayerId = &v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetStartTime(v string) *CreateExperimentRequest {
	s.StartTime = &v
	return s
}

func (s *CreateExperimentRequest) SetWorkspaceId(v string) *CreateExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type CreateExperimentResponseBody struct {
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type CreateExperimentVersionRequest struct {
	Config      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdIds    *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers  *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Flow         *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateExperimentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentVersionRequest) SetConfig(v string) *CreateExperimentVersionRequest {
	s.Config = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetCrowdIds(v string) *CreateExperimentVersionRequest {
	s.CrowdIds = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetDebugUsers(v string) *CreateExperimentVersionRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetDescription(v string) *CreateExperimentVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetExperimentId(v string) *CreateExperimentVersionRequest {
	s.ExperimentId = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetFlow(v int32) *CreateExperimentVersionRequest {
	s.Flow = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetName(v string) *CreateExperimentVersionRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentVersionRequest) SetType(v string) *CreateExperimentVersionRequest {
	s.Type = &v
	return s
}

type CreateExperimentVersionResponseBody struct {
	ExperimentVersionId *string `json:"ExperimentVersionId,omitempty" xml:"ExperimentVersionId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentVersionResponseBody) SetExperimentVersionId(v string) *CreateExperimentVersionResponseBody {
	s.ExperimentVersionId = &v
	return s
}

func (s *CreateExperimentVersionResponseBody) SetRequestId(v string) *CreateExperimentVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentVersionResponse) SetHeaders(v map[string]*string) *CreateExperimentVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentVersionResponse) SetStatusCode(v int32) *CreateExperimentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentVersionResponse) SetBody(v *CreateExperimentVersionResponseBody) *CreateExperimentVersionResponse {
	s.Body = v
	return s
}

type CreateFeatureRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureRequest) SetRegionId(v string) *CreateFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFeatureRequest) SetName(v string) *CreateFeatureRequest {
	s.Name = &v
	return s
}

func (s *CreateFeatureRequest) SetStatus(v string) *CreateFeatureRequest {
	s.Status = &v
	return s
}

type CreateFeatureResponseBody struct {
	FeatureId *string `json:"FeatureId,omitempty" xml:"FeatureId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureResponseBody) SetFeatureId(v string) *CreateFeatureResponseBody {
	s.FeatureId = &v
	return s
}

func (s *CreateFeatureResponseBody) SetRequestId(v string) *CreateFeatureResponseBody {
	s.RequestId = &v
	return s
}

type CreateFeatureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureResponse) SetHeaders(v map[string]*string) *CreateFeatureResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureResponse) SetStatusCode(v int32) *CreateFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureResponse) SetBody(v *CreateFeatureResponseBody) *CreateFeatureResponse {
	s.Body = v
	return s
}

type CreateLayerRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId    *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateLayerRequest) SetDomainId(v string) *CreateLayerRequest {
	s.DomainId = &v
	return s
}

func (s *CreateLayerRequest) SetName(v string) *CreateLayerRequest {
	s.Name = &v
	return s
}

func (s *CreateLayerRequest) SetProjectId(v string) *CreateLayerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateLayerRequest) SetWorkspaceId(v string) *CreateLayerRequest {
	s.WorkspaceId = &v
	return s
}

type CreateLayerResponseBody struct {
	LayerId   *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
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

type CreateMetricRequest struct {
	// This parameter is required.
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	MetricGroupId *string `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SourceTableMetaId *string `json:"SourceTableMetaId,omitempty" xml:"SourceTableMetaId,omitempty"`
}

func (s CreateMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricRequest) SetDefinition(v string) *CreateMetricRequest {
	s.Definition = &v
	return s
}

func (s *CreateMetricRequest) SetDescription(v string) *CreateMetricRequest {
	s.Description = &v
	return s
}

func (s *CreateMetricRequest) SetMetricGroupId(v string) *CreateMetricRequest {
	s.MetricGroupId = &v
	return s
}

func (s *CreateMetricRequest) SetName(v string) *CreateMetricRequest {
	s.Name = &v
	return s
}

func (s *CreateMetricRequest) SetSourceTableMetaId(v string) *CreateMetricRequest {
	s.SourceTableMetaId = &v
	return s
}

type CreateMetricResponseBody struct {
	MetricId  *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricResponseBody) SetMetricId(v string) *CreateMetricResponseBody {
	s.MetricId = &v
	return s
}

func (s *CreateMetricResponseBody) SetRequestId(v string) *CreateMetricResponseBody {
	s.RequestId = &v
	return s
}

type CreateMetricResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricResponse) SetHeaders(v map[string]*string) *CreateMetricResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricResponse) SetStatusCode(v int32) *CreateMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetricResponse) SetBody(v *CreateMetricResponseBody) *CreateMetricResponse {
	s.Body = v
	return s
}

type CreateMetricGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricGroupRequest) SetDescription(v string) *CreateMetricGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateMetricGroupRequest) SetName(v string) *CreateMetricGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateMetricGroupRequest) SetWorkspaceId(v string) *CreateMetricGroupRequest {
	s.WorkspaceId = &v
	return s
}

type CreateMetricGroupResponseBody struct {
	MetricGroupId *string `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricGroupResponseBody) SetMetricGroupId(v string) *CreateMetricGroupResponseBody {
	s.MetricGroupId = &v
	return s
}

func (s *CreateMetricGroupResponseBody) SetRequestId(v string) *CreateMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateMetricGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricGroupResponse) SetHeaders(v map[string]*string) *CreateMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricGroupResponse) SetStatusCode(v int32) *CreateMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetricGroupResponse) SetBody(v *CreateMetricGroupResponseBody) *CreateMetricGroupResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetWorkspaceId(v string) *CreateProjectRequest {
	s.WorkspaceId = &v
	return s
}

type CreateProjectResponseBody struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectId(v string) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateTableMetaRequest struct {
	// This parameter is required.
	DatasourceInfo *string `json:"DatasourceInfo,omitempty" xml:"DatasourceInfo,omitempty"`
	// This parameter is required.
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Fields []*CreateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTableMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTableMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateTableMetaRequest) SetDatasourceInfo(v string) *CreateTableMetaRequest {
	s.DatasourceInfo = &v
	return s
}

func (s *CreateTableMetaRequest) SetDatasourceType(v string) *CreateTableMetaRequest {
	s.DatasourceType = &v
	return s
}

func (s *CreateTableMetaRequest) SetDescription(v string) *CreateTableMetaRequest {
	s.Description = &v
	return s
}

func (s *CreateTableMetaRequest) SetFields(v []*CreateTableMetaRequestFields) *CreateTableMetaRequest {
	s.Fields = v
	return s
}

func (s *CreateTableMetaRequest) SetName(v string) *CreateTableMetaRequest {
	s.Name = &v
	return s
}

func (s *CreateTableMetaRequest) SetTableName(v string) *CreateTableMetaRequest {
	s.TableName = &v
	return s
}

func (s *CreateTableMetaRequest) SetWorkspaceId(v string) *CreateTableMetaRequest {
	s.WorkspaceId = &v
	return s
}

type CreateTableMetaRequestFields struct {
	// This parameter is required.
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTableMetaRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateTableMetaRequestFields) GoString() string {
	return s.String()
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

type DeleteCrowdResponseBody struct {
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

type DeleteDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteExperimentResponseBody struct {
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

type DeleteExperimentVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentVersionResponseBody) SetRequestId(v string) *DeleteExperimentVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentVersionResponse) SetHeaders(v map[string]*string) *DeleteExperimentVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentVersionResponse) SetStatusCode(v int32) *DeleteExperimentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentVersionResponse) SetBody(v *DeleteExperimentVersionResponseBody) *DeleteExperimentVersionResponse {
	s.Body = v
	return s
}

type DeleteFeatureResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureResponseBody) SetRequestId(v string) *DeleteFeatureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFeatureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureResponse) SetHeaders(v map[string]*string) *DeleteFeatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureResponse) SetStatusCode(v int32) *DeleteFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureResponse) SetBody(v *DeleteFeatureResponseBody) *DeleteFeatureResponse {
	s.Body = v
	return s
}

type DeleteLayerResponseBody struct {
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

type DeleteMetricResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricResponseBody) SetRequestId(v string) *DeleteMetricResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMetricResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricResponse) SetHeaders(v map[string]*string) *DeleteMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricResponse) SetStatusCode(v int32) *DeleteMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricResponse) SetBody(v *DeleteMetricResponseBody) *DeleteMetricResponse {
	s.Body = v
	return s
}

type DeleteMetricGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricGroupResponseBody) SetRequestId(v string) *DeleteMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMetricGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricGroupResponse) SetHeaders(v map[string]*string) *DeleteMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricGroupResponse) SetStatusCode(v int32) *DeleteMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricGroupResponse) SetBody(v *DeleteMetricGroupResponseBody) *DeleteMetricGroupResponse {
	s.Body = v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
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

type GetCrowdResponseBody struct {
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Label           *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Quantity        *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users           *string `json:"Users,omitempty" xml:"Users,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrowdResponseBody) SetCrowdId(v string) *GetCrowdResponseBody {
	s.CrowdId = &v
	return s
}

func (s *GetCrowdResponseBody) SetDescription(v string) *GetCrowdResponseBody {
	s.Description = &v
	return s
}

func (s *GetCrowdResponseBody) SetGmtCreateTime(v string) *GetCrowdResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetCrowdResponseBody) SetGmtModifiedTime(v string) *GetCrowdResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetCrowdResponseBody) SetLabel(v string) *GetCrowdResponseBody {
	s.Label = &v
	return s
}

func (s *GetCrowdResponseBody) SetName(v string) *GetCrowdResponseBody {
	s.Name = &v
	return s
}

func (s *GetCrowdResponseBody) SetQuantity(v string) *GetCrowdResponseBody {
	s.Quantity = &v
	return s
}

func (s *GetCrowdResponseBody) SetRequestId(v string) *GetCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrowdResponseBody) SetUsers(v string) *GetCrowdResponseBody {
	s.Users = &v
	return s
}

func (s *GetCrowdResponseBody) SetWorkspaceId(v string) *GetCrowdResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetCrowdResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdResponse) GoString() string {
	return s.String()
}

func (s *GetCrowdResponse) SetHeaders(v map[string]*string) *GetCrowdResponse {
	s.Headers = v
	return s
}

func (s *GetCrowdResponse) SetStatusCode(v int32) *GetCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrowdResponse) SetBody(v *GetCrowdResponseBody) *GetCrowdResponse {
	s.Body = v
	return s
}

type GetDomainRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) SetProjectId(v string) *GetDomainRequest {
	s.ProjectId = &v
	return s
}

type GetDomainResponseBody struct {
	BucketType      *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Condition       *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CrowdIds        *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers      *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId        *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	Flow            *int64  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsDefaultDomain *bool   `json:"IsDefaultDomain,omitempty" xml:"IsDefaultDomain,omitempty"`
	LayerId         *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	LayerName       *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) SetBucketType(v string) *GetDomainResponseBody {
	s.BucketType = &v
	return s
}

func (s *GetDomainResponseBody) SetCondition(v string) *GetDomainResponseBody {
	s.Condition = &v
	return s
}

func (s *GetDomainResponseBody) SetCrowdIds(v string) *GetDomainResponseBody {
	s.CrowdIds = &v
	return s
}

func (s *GetDomainResponseBody) SetDebugUsers(v string) *GetDomainResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetDomainResponseBody) SetDescription(v string) *GetDomainResponseBody {
	s.Description = &v
	return s
}

func (s *GetDomainResponseBody) SetDomainId(v string) *GetDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetDomainResponseBody) SetFlow(v int64) *GetDomainResponseBody {
	s.Flow = &v
	return s
}

func (s *GetDomainResponseBody) SetGmtCreateTime(v string) *GetDomainResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDomainResponseBody) SetGmtModifiedTime(v string) *GetDomainResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDomainResponseBody) SetIsDefaultDomain(v bool) *GetDomainResponseBody {
	s.IsDefaultDomain = &v
	return s
}

func (s *GetDomainResponseBody) SetLayerId(v string) *GetDomainResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetDomainResponseBody) SetLayerName(v string) *GetDomainResponseBody {
	s.LayerName = &v
	return s
}

func (s *GetDomainResponseBody) SetName(v string) *GetDomainResponseBody {
	s.Name = &v
	return s
}

func (s *GetDomainResponseBody) SetProjectId(v string) *GetDomainResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetDomainResponseBody) SetProjectName(v string) *GetDomainResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainResponseBody) SetWorkspaceId(v string) *GetDomainResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetStatusCode(v int32) *GetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainResponse) SetBody(v *GetDomainResponseBody) *GetDomainResponse {
	s.Body = v
	return s
}

type GetExperimentResponseBody struct {
	BucketType      *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Buckets         *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Condition       *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CoreMetricId    *string `json:"CoreMetricId,omitempty" xml:"CoreMetricId,omitempty"`
	CrowdIds        *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers      *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Flow            *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FocusMetricIds  *string `json:"FocusMetricIds,omitempty" xml:"FocusMetricIds,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	LayerId         *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	LayerName       *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner           *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBody) SetBucketType(v string) *GetExperimentResponseBody {
	s.BucketType = &v
	return s
}

func (s *GetExperimentResponseBody) SetBuckets(v string) *GetExperimentResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetExperimentResponseBody) SetCondition(v string) *GetExperimentResponseBody {
	s.Condition = &v
	return s
}

func (s *GetExperimentResponseBody) SetCoreMetricId(v string) *GetExperimentResponseBody {
	s.CoreMetricId = &v
	return s
}

func (s *GetExperimentResponseBody) SetCrowdIds(v string) *GetExperimentResponseBody {
	s.CrowdIds = &v
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

func (s *GetExperimentResponseBody) SetDomainName(v string) *GetExperimentResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetExperimentResponseBody) SetEndTime(v string) *GetExperimentResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetExperimentId(v string) *GetExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBody) SetFlow(v int32) *GetExperimentResponseBody {
	s.Flow = &v
	return s
}

func (s *GetExperimentResponseBody) SetFocusMetricIds(v string) *GetExperimentResponseBody {
	s.FocusMetricIds = &v
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

func (s *GetExperimentResponseBody) SetLayerId(v string) *GetExperimentResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetExperimentResponseBody) SetLayerName(v string) *GetExperimentResponseBody {
	s.LayerName = &v
	return s
}

func (s *GetExperimentResponseBody) SetName(v string) *GetExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentResponseBody) SetOwner(v string) *GetExperimentResponseBody {
	s.Owner = &v
	return s
}

func (s *GetExperimentResponseBody) SetProjectName(v string) *GetExperimentResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetExperimentResponseBody) SetRequestId(v string) *GetExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResponseBody) SetStartTime(v string) *GetExperimentResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetStatus(v string) *GetExperimentResponseBody {
	s.Status = &v
	return s
}

func (s *GetExperimentResponseBody) SetWorkspaceId(v string) *GetExperimentResponseBody {
	s.WorkspaceId = &v
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

type GetExperimentVersionResponseBody struct {
	Buckets             *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config              *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdIds            *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers          *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId        *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentVersionId *string `json:"ExperimentVersionId,omitempty" xml:"ExperimentVersionId,omitempty"`
	Flow                *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	GmtCreateTime       *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime     *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExperimentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentVersionResponseBody) SetBuckets(v string) *GetExperimentVersionResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetConfig(v string) *GetExperimentVersionResponseBody {
	s.Config = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetCrowdIds(v string) *GetExperimentVersionResponseBody {
	s.CrowdIds = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetDebugUsers(v string) *GetExperimentVersionResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetDescription(v string) *GetExperimentVersionResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetExperimentId(v string) *GetExperimentVersionResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetExperimentVersionId(v string) *GetExperimentVersionResponseBody {
	s.ExperimentVersionId = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetFlow(v int32) *GetExperimentVersionResponseBody {
	s.Flow = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetGmtCreateTime(v string) *GetExperimentVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetGmtModifiedTime(v string) *GetExperimentVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetName(v string) *GetExperimentVersionResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetRequestId(v string) *GetExperimentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentVersionResponseBody) SetType(v string) *GetExperimentVersionResponseBody {
	s.Type = &v
	return s
}

type GetExperimentVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentVersionResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentVersionResponse) SetHeaders(v map[string]*string) *GetExperimentVersionResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentVersionResponse) SetStatusCode(v int32) *GetExperimentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentVersionResponse) SetBody(v *GetExperimentVersionResponseBody) *GetExperimentVersionResponse {
	s.Body = v
	return s
}

type GetFeatureResponseBody struct {
	Condition             *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Config                *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DomainId              *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ExperimentId          *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName        *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	ExperimentOwner       *string `json:"ExperimentOwner,omitempty" xml:"ExperimentOwner,omitempty"`
	ExperimentVersionId   *string `json:"ExperimentVersionId,omitempty" xml:"ExperimentVersionId,omitempty"`
	ExperimentVersionName *string `json:"ExperimentVersionName,omitempty" xml:"ExperimentVersionName,omitempty"`
	FeatureId             *string `json:"FeatureId,omitempty" xml:"FeatureId,omitempty"`
	GmtCreateTime         *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime       *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName           *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ReleaseTime           *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureResponseBody) SetCondition(v string) *GetFeatureResponseBody {
	s.Condition = &v
	return s
}

func (s *GetFeatureResponseBody) SetConfig(v string) *GetFeatureResponseBody {
	s.Config = &v
	return s
}

func (s *GetFeatureResponseBody) SetDomainId(v string) *GetFeatureResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetFeatureResponseBody) SetDomainName(v string) *GetFeatureResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetFeatureResponseBody) SetExperimentId(v string) *GetFeatureResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetFeatureResponseBody) SetExperimentName(v string) *GetFeatureResponseBody {
	s.ExperimentName = &v
	return s
}

func (s *GetFeatureResponseBody) SetExperimentOwner(v string) *GetFeatureResponseBody {
	s.ExperimentOwner = &v
	return s
}

func (s *GetFeatureResponseBody) SetExperimentVersionId(v string) *GetFeatureResponseBody {
	s.ExperimentVersionId = &v
	return s
}

func (s *GetFeatureResponseBody) SetExperimentVersionName(v string) *GetFeatureResponseBody {
	s.ExperimentVersionName = &v
	return s
}

func (s *GetFeatureResponseBody) SetFeatureId(v string) *GetFeatureResponseBody {
	s.FeatureId = &v
	return s
}

func (s *GetFeatureResponseBody) SetGmtCreateTime(v string) *GetFeatureResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetFeatureResponseBody) SetGmtModifiedTime(v string) *GetFeatureResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetFeatureResponseBody) SetName(v string) *GetFeatureResponseBody {
	s.Name = &v
	return s
}

func (s *GetFeatureResponseBody) SetProjectId(v string) *GetFeatureResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetFeatureResponseBody) SetProjectName(v string) *GetFeatureResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetFeatureResponseBody) SetReleaseTime(v string) *GetFeatureResponseBody {
	s.ReleaseTime = &v
	return s
}

func (s *GetFeatureResponseBody) SetRequestId(v string) *GetFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureResponseBody) SetStatus(v string) *GetFeatureResponseBody {
	s.Status = &v
	return s
}

func (s *GetFeatureResponseBody) SetWorkspaceId(v string) *GetFeatureResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetFeatureResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureResponse) SetHeaders(v map[string]*string) *GetFeatureResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureResponse) SetStatusCode(v int32) *GetFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureResponse) SetBody(v *GetFeatureResponseBody) *GetFeatureResponse {
	s.Body = v
	return s
}

type GetLayerResponseBody struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId        *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsDefaultLayer  *bool   `json:"IsDefaultLayer,omitempty" xml:"IsDefaultLayer,omitempty"`
	LayerId         *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *GetLayerResponseBody) SetDomainId(v string) *GetLayerResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetLayerResponseBody) SetDomainName(v string) *GetLayerResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetLayerResponseBody) SetGmtCreateTime(v string) *GetLayerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLayerResponseBody) SetGmtModifiedTime(v string) *GetLayerResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetLayerResponseBody) SetIsDefaultLayer(v bool) *GetLayerResponseBody {
	s.IsDefaultLayer = &v
	return s
}

func (s *GetLayerResponseBody) SetLayerId(v string) *GetLayerResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetLayerResponseBody) SetName(v string) *GetLayerResponseBody {
	s.Name = &v
	return s
}

func (s *GetLayerResponseBody) SetProjectId(v string) *GetLayerResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetLayerResponseBody) SetProjectName(v string) *GetLayerResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetLayerResponseBody) SetRequestId(v string) *GetLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLayerResponseBody) SetWorkspaceId(v string) *GetLayerResponseBody {
	s.WorkspaceId = &v
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

type GetMetricResponseBody struct {
	Definition          *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime       *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime     *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MetricGroupId       *string `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	MetricId            *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceTableMetaId   *string `json:"SourceTableMetaId,omitempty" xml:"SourceTableMetaId,omitempty"`
	SourceTableMetaName *string `json:"SourceTableMetaName,omitempty" xml:"SourceTableMetaName,omitempty"`
}

func (s GetMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricResponseBody) SetDefinition(v string) *GetMetricResponseBody {
	s.Definition = &v
	return s
}

func (s *GetMetricResponseBody) SetDescription(v string) *GetMetricResponseBody {
	s.Description = &v
	return s
}

func (s *GetMetricResponseBody) SetGmtCreateTime(v string) *GetMetricResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetMetricResponseBody) SetGmtModifiedTime(v string) *GetMetricResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetMetricResponseBody) SetMetricGroupId(v string) *GetMetricResponseBody {
	s.MetricGroupId = &v
	return s
}

func (s *GetMetricResponseBody) SetMetricId(v string) *GetMetricResponseBody {
	s.MetricId = &v
	return s
}

func (s *GetMetricResponseBody) SetName(v string) *GetMetricResponseBody {
	s.Name = &v
	return s
}

func (s *GetMetricResponseBody) SetRequestId(v string) *GetMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetricResponseBody) SetSourceTableMetaId(v string) *GetMetricResponseBody {
	s.SourceTableMetaId = &v
	return s
}

func (s *GetMetricResponseBody) SetSourceTableMetaName(v string) *GetMetricResponseBody {
	s.SourceTableMetaName = &v
	return s
}

type GetMetricResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetricResponse) GoString() string {
	return s.String()
}

func (s *GetMetricResponse) SetHeaders(v map[string]*string) *GetMetricResponse {
	s.Headers = v
	return s
}

func (s *GetMetricResponse) SetStatusCode(v int32) *GetMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricResponse) SetBody(v *GetMetricResponseBody) *GetMetricResponse {
	s.Body = v
	return s
}

type GetMetricGroupResponseBody struct {
	Description     *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                              `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                              `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MetricGroupId   *string                              `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	Metrics         []*GetMetricGroupResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Name            *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId       *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkspaceId     *string                              `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricGroupResponseBody) SetDescription(v string) *GetMetricGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetMetricGroupResponseBody) SetGmtCreateTime(v string) *GetMetricGroupResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetMetricGroupResponseBody) SetGmtModifiedTime(v string) *GetMetricGroupResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetMetricGroupResponseBody) SetMetricGroupId(v string) *GetMetricGroupResponseBody {
	s.MetricGroupId = &v
	return s
}

func (s *GetMetricGroupResponseBody) SetMetrics(v []*GetMetricGroupResponseBodyMetrics) *GetMetricGroupResponseBody {
	s.Metrics = v
	return s
}

func (s *GetMetricGroupResponseBody) SetName(v string) *GetMetricGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetMetricGroupResponseBody) SetRequestId(v string) *GetMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetricGroupResponseBody) SetWorkspaceId(v string) *GetMetricGroupResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetMetricGroupResponseBodyMetrics struct {
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime           *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime         *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MetricId                *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedExperimentNumber *int64  `json:"RelatedExperimentNumber,omitempty" xml:"RelatedExperimentNumber,omitempty"`
	SourceTableMetaId       *string `json:"SourceTableMetaId,omitempty" xml:"SourceTableMetaId,omitempty"`
}

func (s GetMetricGroupResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetMetricGroupResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *GetMetricGroupResponseBodyMetrics) SetDefinition(v string) *GetMetricGroupResponseBodyMetrics {
	s.Definition = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetDescription(v string) *GetMetricGroupResponseBodyMetrics {
	s.Description = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetGmtCreateTime(v string) *GetMetricGroupResponseBodyMetrics {
	s.GmtCreateTime = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetGmtModifiedTime(v string) *GetMetricGroupResponseBodyMetrics {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetMetricId(v string) *GetMetricGroupResponseBodyMetrics {
	s.MetricId = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetName(v string) *GetMetricGroupResponseBodyMetrics {
	s.Name = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetRelatedExperimentNumber(v int64) *GetMetricGroupResponseBodyMetrics {
	s.RelatedExperimentNumber = &v
	return s
}

func (s *GetMetricGroupResponseBodyMetrics) SetSourceTableMetaId(v string) *GetMetricGroupResponseBodyMetrics {
	s.SourceTableMetaId = &v
	return s
}

type GetMetricGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMetricGroupResponse) SetHeaders(v map[string]*string) *GetMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMetricGroupResponse) SetStatusCode(v int32) *GetMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricGroupResponse) SetBody(v *GetMetricGroupResponseBody) *GetMetricGroupResponse {
	s.Body = v
	return s
}

type GetProjectResponseBody struct {
	DefaultDomainId *string `json:"DefaultDomainId,omitempty" xml:"DefaultDomainId,omitempty"`
	DefaultLayerId  *string `json:"DefaultLayerId,omitempty" xml:"DefaultLayerId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetDefaultDomainId(v string) *GetProjectResponseBody {
	s.DefaultDomainId = &v
	return s
}

func (s *GetProjectResponseBody) SetDefaultLayerId(v string) *GetProjectResponseBody {
	s.DefaultLayerId = &v
	return s
}

func (s *GetProjectResponseBody) SetDescription(v string) *GetProjectResponseBody {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBody) SetGmtCreateTime(v string) *GetProjectResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetGmtModifiedTime(v string) *GetProjectResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetProjectResponseBody) SetName(v string) *GetProjectResponseBody {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectId(v string) *GetProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetWorkspaceId(v string) *GetProjectResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetTableMetaResponseBody struct {
	DatasourceInfo  *string                           `json:"DatasourceInfo,omitempty" xml:"DatasourceInfo,omitempty"`
	DatasourceType  *string                           `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	Description     *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields          []*GetTableMetaResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	GmtCreateTime   *string                           `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                           `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TableMetaId     *string                           `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	TableName       *string                           `json:"TableName,omitempty" xml:"TableName,omitempty"`
	WorkspaceId     *string                           `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTableMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponseBody) SetDatasourceInfo(v string) *GetTableMetaResponseBody {
	s.DatasourceInfo = &v
	return s
}

func (s *GetTableMetaResponseBody) SetDatasourceType(v string) *GetTableMetaResponseBody {
	s.DatasourceType = &v
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

func (s *GetTableMetaResponseBody) SetGmtModifiedTime(v string) *GetTableMetaResponseBody {
	s.GmtModifiedTime = &v
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

func (s *GetTableMetaResponseBody) SetTableMetaId(v string) *GetTableMetaResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetTableName(v string) *GetTableMetaResponseBody {
	s.TableName = &v
	return s
}

func (s *GetTableMetaResponseBody) SetWorkspaceId(v string) *GetTableMetaResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetTableMetaResponseBodyFields struct {
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTableMetaResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetTableMetaResponseBodyFields) GoString() string {
	return s.String()
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

type ListCrowdsRequest struct {
	All         *bool   `json:"All,omitempty" xml:"All,omitempty"`
	CrowdId     *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdName   *string `json:"CrowdName,omitempty" xml:"CrowdName,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCrowdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdsRequest) GoString() string {
	return s.String()
}

func (s *ListCrowdsRequest) SetAll(v bool) *ListCrowdsRequest {
	s.All = &v
	return s
}

func (s *ListCrowdsRequest) SetCrowdId(v string) *ListCrowdsRequest {
	s.CrowdId = &v
	return s
}

func (s *ListCrowdsRequest) SetCrowdName(v string) *ListCrowdsRequest {
	s.CrowdName = &v
	return s
}

func (s *ListCrowdsRequest) SetOrder(v string) *ListCrowdsRequest {
	s.Order = &v
	return s
}

func (s *ListCrowdsRequest) SetPageNumber(v int64) *ListCrowdsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrowdsRequest) SetPageSize(v int64) *ListCrowdsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrowdsRequest) SetRegionId(v string) *ListCrowdsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCrowdsRequest) SetSortBy(v string) *ListCrowdsRequest {
	s.SortBy = &v
	return s
}

func (s *ListCrowdsRequest) SetWorkspaceId(v string) *ListCrowdsRequest {
	s.WorkspaceId = &v
	return s
}

type ListCrowdsResponseBody struct {
	Crowds     []*ListCrowdsResponseBodyCrowds `json:"Crowds,omitempty" xml:"Crowds,omitempty" type:"Repeated"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AliyunId      *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CrowdId       *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	Label         *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Quantity      *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Users         *string `json:"Users,omitempty" xml:"Users,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCrowdsResponseBodyCrowds) String() string {
	return tea.Prettify(s)
}

func (s ListCrowdsResponseBodyCrowds) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponseBodyCrowds) SetAliyunId(v string) *ListCrowdsResponseBodyCrowds {
	s.AliyunId = &v
	return s
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

func (s *ListCrowdsResponseBodyCrowds) SetGmtModifyTime(v string) *ListCrowdsResponseBodyCrowds {
	s.GmtModifyTime = &v
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

func (s *ListCrowdsResponseBodyCrowds) SetUsers(v string) *ListCrowdsResponseBodyCrowds {
	s.Users = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetWorkspaceId(v string) *ListCrowdsResponseBodyCrowds {
	s.WorkspaceId = &v
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

type ListDomainsRequest struct {
	All         *bool   `json:"All,omitempty" xml:"All,omitempty"`
	DomainId    *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LayerId     *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetAll(v bool) *ListDomainsRequest {
	s.All = &v
	return s
}

func (s *ListDomainsRequest) SetDomainId(v string) *ListDomainsRequest {
	s.DomainId = &v
	return s
}

func (s *ListDomainsRequest) SetDomainName(v string) *ListDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDomainsRequest) SetLayerId(v string) *ListDomainsRequest {
	s.LayerId = &v
	return s
}

func (s *ListDomainsRequest) SetOrder(v string) *ListDomainsRequest {
	s.Order = &v
	return s
}

func (s *ListDomainsRequest) SetPageNumber(v int64) *ListDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsRequest) SetPageSize(v int64) *ListDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainsRequest) SetProjectId(v string) *ListDomainsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDomainsRequest) SetRegionId(v string) *ListDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDomainsRequest) SetSortBy(v string) *ListDomainsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDomainsRequest) SetWorkspaceId(v string) *ListDomainsRequest {
	s.WorkspaceId = &v
	return s
}

type ListDomainsResponseBody struct {
	Domains    []*ListDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetDomains(v []*ListDomainsResponseBodyDomains) *ListDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsResponseBody) SetTotalCount(v string) *ListDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDomainsResponseBodyDomains struct {
	BucketType      *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Condition       *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CrowdIds        *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers      *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId        *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	Flow            *int64  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsDefaultDomain *bool   `json:"IsDefaultDomain,omitempty" xml:"IsDefaultDomain,omitempty"`
	LayerId         *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	LayerName       *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomains) SetBucketType(v string) *ListDomainsResponseBodyDomains {
	s.BucketType = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetCondition(v string) *ListDomainsResponseBodyDomains {
	s.Condition = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetCrowdIds(v string) *ListDomainsResponseBodyDomains {
	s.CrowdIds = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDebugUsers(v string) *ListDomainsResponseBodyDomains {
	s.DebugUsers = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDescription(v string) *ListDomainsResponseBodyDomains {
	s.Description = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomainId(v string) *ListDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetFlow(v int64) *ListDomainsResponseBodyDomains {
	s.Flow = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetGmtCreateTime(v string) *ListDomainsResponseBodyDomains {
	s.GmtCreateTime = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetGmtModifiedTime(v string) *ListDomainsResponseBodyDomains {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetIsDefaultDomain(v bool) *ListDomainsResponseBodyDomains {
	s.IsDefaultDomain = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetLayerId(v string) *ListDomainsResponseBodyDomains {
	s.LayerId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetLayerName(v string) *ListDomainsResponseBodyDomains {
	s.LayerName = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetName(v string) *ListDomainsResponseBodyDomains {
	s.Name = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetProjectId(v string) *ListDomainsResponseBodyDomains {
	s.ProjectId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetProjectName(v string) *ListDomainsResponseBodyDomains {
	s.ProjectName = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetWorkspaceId(v string) *ListDomainsResponseBodyDomains {
	s.WorkspaceId = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListExperimentVersionsRequest struct {
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListExperimentVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentVersionsRequest) SetAll(v bool) *ListExperimentVersionsRequest {
	s.All = &v
	return s
}

func (s *ListExperimentVersionsRequest) SetExperimentId(v string) *ListExperimentVersionsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentVersionsRequest) SetOrder(v string) *ListExperimentVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentVersionsRequest) SetPageNumber(v int64) *ListExperimentVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentVersionsRequest) SetPageSize(v int64) *ListExperimentVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentVersionsRequest) SetSortBy(v string) *ListExperimentVersionsRequest {
	s.SortBy = &v
	return s
}

type ListExperimentVersionsResponseBody struct {
	ExperimentVersions []*ListExperimentVersionsResponseBodyExperimentVersions `json:"ExperimentVersions,omitempty" xml:"ExperimentVersions,omitempty" type:"Repeated"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int64                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentVersionsResponseBody) SetExperimentVersions(v []*ListExperimentVersionsResponseBodyExperimentVersions) *ListExperimentVersionsResponseBody {
	s.ExperimentVersions = v
	return s
}

func (s *ListExperimentVersionsResponseBody) SetRequestId(v string) *ListExperimentVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentVersionsResponseBody) SetTotalCount(v int64) *ListExperimentVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentVersionsResponseBodyExperimentVersions struct {
	Buckets             *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config              *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdIds            *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers          *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId        *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentVersionId *string `json:"ExperimentVersionId,omitempty" xml:"ExperimentVersionId,omitempty"`
	Flow                *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	GmtCreateTime       *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime     *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListExperimentVersionsResponseBodyExperimentVersions) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentVersionsResponseBodyExperimentVersions) GoString() string {
	return s.String()
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetBuckets(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.Buckets = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetConfig(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.Config = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetCrowdIds(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.CrowdIds = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetDebugUsers(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.DebugUsers = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetDescription(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.Description = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetExperimentId(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetExperimentVersionId(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.ExperimentVersionId = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetFlow(v int32) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.Flow = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetGmtCreateTime(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.GmtCreateTime = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetGmtModifiedTime(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetName(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.Name = &v
	return s
}

func (s *ListExperimentVersionsResponseBodyExperimentVersions) SetType(v string) *ListExperimentVersionsResponseBodyExperimentVersions {
	s.Type = &v
	return s
}

type ListExperimentVersionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentVersionsResponse) SetHeaders(v map[string]*string) *ListExperimentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentVersionsResponse) SetStatusCode(v int32) *ListExperimentVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentVersionsResponse) SetBody(v *ListExperimentVersionsResponseBody) *ListExperimentVersionsResponse {
	s.Body = v
	return s
}

type ListExperimentsRequest struct {
	All            *bool   `json:"All,omitempty" xml:"All,omitempty"`
	ExperimentId   *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	LayerId        *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber     *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SortBy         *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentsRequest) SetAll(v bool) *ListExperimentsRequest {
	s.All = &v
	return s
}

func (s *ListExperimentsRequest) SetExperimentId(v string) *ListExperimentsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsRequest) SetExperimentName(v string) *ListExperimentsRequest {
	s.ExperimentName = &v
	return s
}

func (s *ListExperimentsRequest) SetLayerId(v string) *ListExperimentsRequest {
	s.LayerId = &v
	return s
}

func (s *ListExperimentsRequest) SetOrder(v string) *ListExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentsRequest) SetPageNumber(v int64) *ListExperimentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentsRequest) SetPageSize(v int64) *ListExperimentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentsRequest) SetProjectId(v string) *ListExperimentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListExperimentsRequest) SetSortBy(v string) *ListExperimentsRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentsRequest) SetStatus(v string) *ListExperimentsRequest {
	s.Status = &v
	return s
}

func (s *ListExperimentsRequest) SetWorkspaceId(v string) *ListExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

type ListExperimentsResponseBody struct {
	Experiments []*ListExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	BucketType      *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Buckets         *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Condition       *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CoreMetricId    *string `json:"CoreMetricId,omitempty" xml:"CoreMetricId,omitempty"`
	CrowdIds        *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers      *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Flow            *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FocusMetricIds  *string `json:"FocusMetricIds,omitempty" xml:"FocusMetricIds,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	LayerId         *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	LayerName       *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner           *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentsResponseBodyExperiments) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyExperiments) SetBucketType(v string) *ListExperimentsResponseBodyExperiments {
	s.BucketType = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetBuckets(v string) *ListExperimentsResponseBodyExperiments {
	s.Buckets = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetCondition(v string) *ListExperimentsResponseBodyExperiments {
	s.Condition = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetCoreMetricId(v string) *ListExperimentsResponseBodyExperiments {
	s.CoreMetricId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetCrowdIds(v string) *ListExperimentsResponseBodyExperiments {
	s.CrowdIds = &v
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

func (s *ListExperimentsResponseBodyExperiments) SetDomainName(v string) *ListExperimentsResponseBodyExperiments {
	s.DomainName = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetEndTime(v string) *ListExperimentsResponseBodyExperiments {
	s.EndTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetFlow(v int32) *ListExperimentsResponseBodyExperiments {
	s.Flow = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetFocusMetricIds(v string) *ListExperimentsResponseBodyExperiments {
	s.FocusMetricIds = &v
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

func (s *ListExperimentsResponseBodyExperiments) SetLayerId(v string) *ListExperimentsResponseBodyExperiments {
	s.LayerId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetLayerName(v string) *ListExperimentsResponseBodyExperiments {
	s.LayerName = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetName(v string) *ListExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetOwner(v string) *ListExperimentsResponseBodyExperiments {
	s.Owner = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetProjectName(v string) *ListExperimentsResponseBodyExperiments {
	s.ProjectName = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetStartTime(v string) *ListExperimentsResponseBodyExperiments {
	s.StartTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetStatus(v string) *ListExperimentsResponseBodyExperiments {
	s.Status = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetWorkspaceId(v string) *ListExperimentsResponseBodyExperiments {
	s.WorkspaceId = &v
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

type ListFeaturesRequest struct {
	All         *bool   `json:"All,omitempty" xml:"All,omitempty"`
	DomainId    *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	FeatureId   *string `json:"FeatureId,omitempty" xml:"FeatureId,omitempty"`
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListFeaturesRequest) SetAll(v bool) *ListFeaturesRequest {
	s.All = &v
	return s
}

func (s *ListFeaturesRequest) SetDomainId(v string) *ListFeaturesRequest {
	s.DomainId = &v
	return s
}

func (s *ListFeaturesRequest) SetFeatureId(v string) *ListFeaturesRequest {
	s.FeatureId = &v
	return s
}

func (s *ListFeaturesRequest) SetFeatureName(v string) *ListFeaturesRequest {
	s.FeatureName = &v
	return s
}

func (s *ListFeaturesRequest) SetOrder(v string) *ListFeaturesRequest {
	s.Order = &v
	return s
}

func (s *ListFeaturesRequest) SetPageNumber(v int64) *ListFeaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeaturesRequest) SetPageSize(v int64) *ListFeaturesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeaturesRequest) SetRegionId(v string) *ListFeaturesRequest {
	s.RegionId = &v
	return s
}

func (s *ListFeaturesRequest) SetSortBy(v string) *ListFeaturesRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeaturesRequest) SetStatus(v string) *ListFeaturesRequest {
	s.Status = &v
	return s
}

func (s *ListFeaturesRequest) SetWorkspaceId(v string) *ListFeaturesRequest {
	s.WorkspaceId = &v
	return s
}

type ListFeaturesResponseBody struct {
	Features   []*ListFeaturesResponseBodyFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponseBody) SetFeatures(v []*ListFeaturesResponseBodyFeatures) *ListFeaturesResponseBody {
	s.Features = v
	return s
}

func (s *ListFeaturesResponseBody) SetRequestId(v string) *ListFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeaturesResponseBody) SetTotalCount(v int64) *ListFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFeaturesResponseBodyFeatures struct {
	Config                *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DomainId              *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ExperimentId          *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ExperimentName        *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	ExperimentOwner       *string `json:"ExperimentOwner,omitempty" xml:"ExperimentOwner,omitempty"`
	ExperimentVersionId   *string `json:"ExperimentVersionId,omitempty" xml:"ExperimentVersionId,omitempty"`
	ExperimentVersionName *string `json:"ExperimentVersionName,omitempty" xml:"ExperimentVersionName,omitempty"`
	FeatureId             *string `json:"FeatureId,omitempty" xml:"FeatureId,omitempty"`
	Filter                *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	GmtCreateTime         *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime       *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName           *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ReleaseTime           *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListFeaturesResponseBodyFeatures) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesResponseBodyFeatures) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponseBodyFeatures) SetConfig(v string) *ListFeaturesResponseBodyFeatures {
	s.Config = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetDomainId(v string) *ListFeaturesResponseBodyFeatures {
	s.DomainId = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetDomainName(v string) *ListFeaturesResponseBodyFeatures {
	s.DomainName = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetExperimentId(v string) *ListFeaturesResponseBodyFeatures {
	s.ExperimentId = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetExperimentName(v string) *ListFeaturesResponseBodyFeatures {
	s.ExperimentName = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetExperimentOwner(v string) *ListFeaturesResponseBodyFeatures {
	s.ExperimentOwner = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetExperimentVersionId(v string) *ListFeaturesResponseBodyFeatures {
	s.ExperimentVersionId = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetExperimentVersionName(v string) *ListFeaturesResponseBodyFeatures {
	s.ExperimentVersionName = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetFeatureId(v string) *ListFeaturesResponseBodyFeatures {
	s.FeatureId = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetFilter(v string) *ListFeaturesResponseBodyFeatures {
	s.Filter = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetGmtCreateTime(v string) *ListFeaturesResponseBodyFeatures {
	s.GmtCreateTime = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetGmtModifiedTime(v string) *ListFeaturesResponseBodyFeatures {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetName(v string) *ListFeaturesResponseBodyFeatures {
	s.Name = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetProjectId(v string) *ListFeaturesResponseBodyFeatures {
	s.ProjectId = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetProjectName(v string) *ListFeaturesResponseBodyFeatures {
	s.ProjectName = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetReleaseTime(v string) *ListFeaturesResponseBodyFeatures {
	s.ReleaseTime = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetStatus(v string) *ListFeaturesResponseBodyFeatures {
	s.Status = &v
	return s
}

func (s *ListFeaturesResponseBodyFeatures) SetWorkspaceId(v string) *ListFeaturesResponseBodyFeatures {
	s.WorkspaceId = &v
	return s
}

type ListFeaturesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponse) SetHeaders(v map[string]*string) *ListFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListFeaturesResponse) SetStatusCode(v int32) *ListFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeaturesResponse) SetBody(v *ListFeaturesResponseBody) *ListFeaturesResponse {
	s.Body = v
	return s
}

type ListLayersRequest struct {
	All         *bool   `json:"All,omitempty" xml:"All,omitempty"`
	DomainId    *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	LayerId     *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	LayerName   *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListLayersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayersRequest) GoString() string {
	return s.String()
}

func (s *ListLayersRequest) SetAll(v bool) *ListLayersRequest {
	s.All = &v
	return s
}

func (s *ListLayersRequest) SetDomainId(v string) *ListLayersRequest {
	s.DomainId = &v
	return s
}

func (s *ListLayersRequest) SetLayerId(v string) *ListLayersRequest {
	s.LayerId = &v
	return s
}

func (s *ListLayersRequest) SetLayerName(v string) *ListLayersRequest {
	s.LayerName = &v
	return s
}

func (s *ListLayersRequest) SetOrder(v string) *ListLayersRequest {
	s.Order = &v
	return s
}

func (s *ListLayersRequest) SetPageNumber(v int64) *ListLayersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLayersRequest) SetPageSize(v string) *ListLayersRequest {
	s.PageSize = &v
	return s
}

func (s *ListLayersRequest) SetProjectId(v string) *ListLayersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListLayersRequest) SetSortBy(v string) *ListLayersRequest {
	s.SortBy = &v
	return s
}

func (s *ListLayersRequest) SetWorkspaceId(v string) *ListLayersRequest {
	s.WorkspaceId = &v
	return s
}

type ListLayersResponseBody struct {
	Layers     []*ListLayersResponseBodyLayers `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId       *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GmtCreateTime  *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime  *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	IsDefaultLayer *bool   `json:"IsDefaultLayer,omitempty" xml:"IsDefaultLayer,omitempty"`
	LayerId        *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *ListLayersResponseBodyLayers) SetDomainId(v string) *ListLayersResponseBodyLayers {
	s.DomainId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetDomainName(v string) *ListLayersResponseBodyLayers {
	s.DomainName = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetGmtCreateTime(v string) *ListLayersResponseBodyLayers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetGmtModifyTime(v string) *ListLayersResponseBodyLayers {
	s.GmtModifyTime = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetIsDefaultLayer(v bool) *ListLayersResponseBodyLayers {
	s.IsDefaultLayer = &v
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

func (s *ListLayersResponseBodyLayers) SetProjectId(v string) *ListLayersResponseBodyLayers {
	s.ProjectId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetProjectName(v string) *ListLayersResponseBodyLayers {
	s.ProjectName = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetWorkspaceId(v string) *ListLayersResponseBodyLayers {
	s.WorkspaceId = &v
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

type ListMetricGroupsRequest struct {
	All             *bool   `json:"All,omitempty" xml:"All,omitempty"`
	MetricGroupId   *string `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	MetricGroupName *string `json:"MetricGroupName,omitempty" xml:"MetricGroupName,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy          *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMetricGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMetricGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMetricGroupsRequest) SetAll(v bool) *ListMetricGroupsRequest {
	s.All = &v
	return s
}

func (s *ListMetricGroupsRequest) SetMetricGroupId(v string) *ListMetricGroupsRequest {
	s.MetricGroupId = &v
	return s
}

func (s *ListMetricGroupsRequest) SetMetricGroupName(v string) *ListMetricGroupsRequest {
	s.MetricGroupName = &v
	return s
}

func (s *ListMetricGroupsRequest) SetOrder(v string) *ListMetricGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListMetricGroupsRequest) SetPageNumber(v int64) *ListMetricGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetricGroupsRequest) SetPageSize(v int64) *ListMetricGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetricGroupsRequest) SetSortBy(v string) *ListMetricGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListMetricGroupsRequest) SetWorkspaceId(v string) *ListMetricGroupsRequest {
	s.WorkspaceId = &v
	return s
}

type ListMetricGroupsResponseBody struct {
	MetricGroups []*ListMetricGroupsResponseBodyMetricGroups `json:"MetricGroups,omitempty" xml:"MetricGroups,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetricGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMetricGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetricGroupsResponseBody) SetMetricGroups(v []*ListMetricGroupsResponseBodyMetricGroups) *ListMetricGroupsResponseBody {
	s.MetricGroups = v
	return s
}

func (s *ListMetricGroupsResponseBody) SetRequestId(v string) *ListMetricGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetricGroupsResponseBody) SetTotalCount(v int64) *ListMetricGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListMetricGroupsResponseBodyMetricGroups struct {
	Description     *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                                            `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MetricGroupId   *string                                            `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	Metrics         []*ListMetricGroupsResponseBodyMetricGroupsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Name            *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkspaceId     *string                                            `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMetricGroupsResponseBodyMetricGroups) String() string {
	return tea.Prettify(s)
}

func (s ListMetricGroupsResponseBodyMetricGroups) GoString() string {
	return s.String()
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetDescription(v string) *ListMetricGroupsResponseBodyMetricGroups {
	s.Description = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetGmtCreateTime(v string) *ListMetricGroupsResponseBodyMetricGroups {
	s.GmtCreateTime = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetGmtModifiedTime(v string) *ListMetricGroupsResponseBodyMetricGroups {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetMetricGroupId(v string) *ListMetricGroupsResponseBodyMetricGroups {
	s.MetricGroupId = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetMetrics(v []*ListMetricGroupsResponseBodyMetricGroupsMetrics) *ListMetricGroupsResponseBodyMetricGroups {
	s.Metrics = v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetName(v string) *ListMetricGroupsResponseBodyMetricGroups {
	s.Name = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroups) SetWorkspaceId(v string) *ListMetricGroupsResponseBodyMetricGroups {
	s.WorkspaceId = &v
	return s
}

type ListMetricGroupsResponseBodyMetricGroupsMetrics struct {
	Definition               *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime            *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime          *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MetricId                 *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedExperimentsNumber *int64  `json:"RelatedExperimentsNumber,omitempty" xml:"RelatedExperimentsNumber,omitempty"`
	SourceTableMetaId        *string `json:"SourceTableMetaId,omitempty" xml:"SourceTableMetaId,omitempty"`
}

func (s ListMetricGroupsResponseBodyMetricGroupsMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListMetricGroupsResponseBodyMetricGroupsMetrics) GoString() string {
	return s.String()
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetDefinition(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.Definition = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetDescription(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.Description = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetGmtCreateTime(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.GmtCreateTime = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetGmtModifiedTime(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetMetricId(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.MetricId = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetName(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.Name = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetRelatedExperimentsNumber(v int64) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.RelatedExperimentsNumber = &v
	return s
}

func (s *ListMetricGroupsResponseBodyMetricGroupsMetrics) SetSourceTableMetaId(v string) *ListMetricGroupsResponseBodyMetricGroupsMetrics {
	s.SourceTableMetaId = &v
	return s
}

type ListMetricGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetricGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetricGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMetricGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMetricGroupsResponse) SetHeaders(v map[string]*string) *ListMetricGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMetricGroupsResponse) SetStatusCode(v int32) *ListMetricGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetricGroupsResponse) SetBody(v *ListMetricGroupsResponseBody) *ListMetricGroupsResponse {
	s.Body = v
	return s
}

type ListMetricsRequest struct {
	All           *string `json:"All,omitempty" xml:"All,omitempty"`
	MetricGroupId *string `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	MetricId      *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	MetricName    *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListMetricsRequest) SetAll(v string) *ListMetricsRequest {
	s.All = &v
	return s
}

func (s *ListMetricsRequest) SetMetricGroupId(v string) *ListMetricsRequest {
	s.MetricGroupId = &v
	return s
}

func (s *ListMetricsRequest) SetMetricId(v string) *ListMetricsRequest {
	s.MetricId = &v
	return s
}

func (s *ListMetricsRequest) SetMetricName(v string) *ListMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *ListMetricsRequest) SetOrder(v string) *ListMetricsRequest {
	s.Order = &v
	return s
}

func (s *ListMetricsRequest) SetPageNumber(v string) *ListMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetricsRequest) SetPageSize(v string) *ListMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetricsRequest) SetSortBy(v string) *ListMetricsRequest {
	s.SortBy = &v
	return s
}

type ListMetricsResponseBody struct {
	Metrics    []*ListMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseBody) SetMetrics(v []*ListMetricsResponseBodyMetrics) *ListMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *ListMetricsResponseBody) SetRequestId(v string) *ListMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetricsResponseBody) SetTotalCount(v int64) *ListMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type ListMetricsResponseBodyMetrics struct {
	Definition          *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime       *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime     *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MetricGroupId       *string `json:"MetricGroupId,omitempty" xml:"MetricGroupId,omitempty"`
	MetricId            *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceTableMetaId   *string `json:"SourceTableMetaId,omitempty" xml:"SourceTableMetaId,omitempty"`
	SourceTableMetaName *string `json:"SourceTableMetaName,omitempty" xml:"SourceTableMetaName,omitempty"`
}

func (s ListMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseBodyMetrics) SetDefinition(v string) *ListMetricsResponseBodyMetrics {
	s.Definition = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetDescription(v string) *ListMetricsResponseBodyMetrics {
	s.Description = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetGmtCreateTime(v string) *ListMetricsResponseBodyMetrics {
	s.GmtCreateTime = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetGmtModifiedTime(v string) *ListMetricsResponseBodyMetrics {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetMetricGroupId(v string) *ListMetricsResponseBodyMetrics {
	s.MetricGroupId = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetMetricId(v string) *ListMetricsResponseBodyMetrics {
	s.MetricId = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetName(v string) *ListMetricsResponseBodyMetrics {
	s.Name = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetSourceTableMetaId(v string) *ListMetricsResponseBodyMetrics {
	s.SourceTableMetaId = &v
	return s
}

func (s *ListMetricsResponseBodyMetrics) SetSourceTableMetaName(v string) *ListMetricsResponseBodyMetrics {
	s.SourceTableMetaName = &v
	return s
}

type ListMetricsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListMetricsResponse) SetHeaders(v map[string]*string) *ListMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListMetricsResponse) SetStatusCode(v int32) *ListMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetricsResponse) SetBody(v *ListMetricsResponseBody) *ListMetricsResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	All         *bool   `json:"All,omitempty" xml:"All,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetAll(v bool) *ListProjectsRequest {
	s.All = &v
	return s
}

func (s *ListProjectsRequest) SetName(v string) *ListProjectsRequest {
	s.Name = &v
	return s
}

func (s *ListProjectsRequest) SetOrder(v string) *ListProjectsRequest {
	s.Order = &v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v string) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v string) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetProjectId(v string) *ListProjectsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsRequest) SetSortBy(v string) *ListProjectsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v string) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

type ListProjectsResponseBody struct {
	Projects   []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyProjects struct {
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	DefaultDomainId *string `json:"DefaultDomainId,omitempty" xml:"DefaultDomainId,omitempty"`
	DefaultLayerId  *string `json:"DefaultLayerId,omitempty" xml:"DefaultLayerId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetAliyunId(v string) *ListProjectsResponseBodyProjects {
	s.AliyunId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDefaultDomainId(v string) *ListProjectsResponseBodyProjects {
	s.DefaultDomainId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDefaultLayerId(v string) *ListProjectsResponseBodyProjects {
	s.DefaultLayerId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtCreateTime(v string) *ListProjectsResponseBodyProjects {
	s.GmtCreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtModifiedTime(v string) *ListProjectsResponseBodyProjects {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetName(v string) *ListProjectsResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectId(v string) *ListProjectsResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetWorkspaceId(v string) *ListProjectsResponseBodyProjects {
	s.WorkspaceId = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListTableMetasRequest struct {
	All            *bool   `json:"All,omitempty" xml:"All,omitempty"`
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber     *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy         *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TableMetaId    *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	TableMetaName  *string `json:"TableMetaName,omitempty" xml:"TableMetaName,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTableMetasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasRequest) GoString() string {
	return s.String()
}

func (s *ListTableMetasRequest) SetAll(v bool) *ListTableMetasRequest {
	s.All = &v
	return s
}

func (s *ListTableMetasRequest) SetDatasourceType(v string) *ListTableMetasRequest {
	s.DatasourceType = &v
	return s
}

func (s *ListTableMetasRequest) SetOrder(v string) *ListTableMetasRequest {
	s.Order = &v
	return s
}

func (s *ListTableMetasRequest) SetPageNumber(v string) *ListTableMetasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTableMetasRequest) SetPageSize(v string) *ListTableMetasRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableMetasRequest) SetSortBy(v string) *ListTableMetasRequest {
	s.SortBy = &v
	return s
}

func (s *ListTableMetasRequest) SetTableMetaId(v string) *ListTableMetasRequest {
	s.TableMetaId = &v
	return s
}

func (s *ListTableMetasRequest) SetTableMetaName(v string) *ListTableMetasRequest {
	s.TableMetaName = &v
	return s
}

func (s *ListTableMetasRequest) SetWorkspaceId(v string) *ListTableMetasRequest {
	s.WorkspaceId = &v
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
	DatasourceInfo  *string                                       `json:"DatasourceInfo,omitempty" xml:"DatasourceInfo,omitempty"`
	DatasourceType  *string                                       `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	Description     *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Fields          []*ListTableMetasResponseBodyTableMetasFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	GmtCreateTime   *string                                       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	TableMetaId     *string                                       `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	TableName       *string                                       `json:"TableName,omitempty" xml:"TableName,omitempty"`
	WorkspaceId     *string                                       `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *ListTableMetasResponseBodyTableMetas) SetDatasourceInfo(v string) *ListTableMetasResponseBodyTableMetas {
	s.DatasourceInfo = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetDatasourceType(v string) *ListTableMetasResponseBodyTableMetas {
	s.DatasourceType = &v
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

func (s *ListTableMetasResponseBodyTableMetas) SetGmtModifiedTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetName(v string) *ListTableMetasResponseBodyTableMetas {
	s.Name = &v
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

func (s *ListTableMetasResponseBodyTableMetas) SetWorkspaceId(v string) *ListTableMetasResponseBodyTableMetas {
	s.WorkspaceId = &v
	return s
}

type ListTableMetasResponseBodyTableMetasFields struct {
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTableMetasResponseBodyTableMetasFields) String() string {
	return tea.Prettify(s)
}

func (s ListTableMetasResponseBodyTableMetasFields) GoString() string {
	return s.String()
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

type PushAllExperimentVersionRequest struct {
	// This parameter is required.
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
}

func (s PushAllExperimentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PushAllExperimentVersionRequest) GoString() string {
	return s.String()
}

func (s *PushAllExperimentVersionRequest) SetFeatureName(v string) *PushAllExperimentVersionRequest {
	s.FeatureName = &v
	return s
}

type PushAllExperimentVersionResponseBody struct {
	FeatureId *string `json:"FeatureId,omitempty" xml:"FeatureId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushAllExperimentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushAllExperimentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PushAllExperimentVersionResponseBody) SetFeatureId(v string) *PushAllExperimentVersionResponseBody {
	s.FeatureId = &v
	return s
}

func (s *PushAllExperimentVersionResponseBody) SetRequestId(v string) *PushAllExperimentVersionResponseBody {
	s.RequestId = &v
	return s
}

type PushAllExperimentVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushAllExperimentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushAllExperimentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PushAllExperimentVersionResponse) GoString() string {
	return s.String()
}

func (s *PushAllExperimentVersionResponse) SetHeaders(v map[string]*string) *PushAllExperimentVersionResponse {
	s.Headers = v
	return s
}

func (s *PushAllExperimentVersionResponse) SetStatusCode(v int32) *PushAllExperimentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PushAllExperimentVersionResponse) SetBody(v *PushAllExperimentVersionResponseBody) *PushAllExperimentVersionResponse {
	s.Body = v
	return s
}

type StartExperimentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *StartExperimentResponseBody) SetRequestId(v string) *StartExperimentResponseBody {
	s.RequestId = &v
	return s
}

type StartExperimentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s StartExperimentResponse) GoString() string {
	return s.String()
}

func (s *StartExperimentResponse) SetHeaders(v map[string]*string) *StartExperimentResponse {
	s.Headers = v
	return s
}

func (s *StartExperimentResponse) SetStatusCode(v int32) *StartExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *StartExperimentResponse) SetBody(v *StartExperimentResponseBody) *StartExperimentResponse {
	s.Body = v
	return s
}

type StopExperimentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *StopExperimentResponseBody) SetRequestId(v string) *StopExperimentResponseBody {
	s.RequestId = &v
	return s
}

type StopExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentResponse) GoString() string {
	return s.String()
}

func (s *StopExperimentResponse) SetHeaders(v map[string]*string) *StopExperimentResponse {
	s.Headers = v
	return s
}

func (s *StopExperimentResponse) SetStatusCode(v int32) *StopExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopExperimentResponse) SetBody(v *StopExperimentResponseBody) *StopExperimentResponse {
	s.Body = v
	return s
}

type UpdateCrowdRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Users       *string `json:"Users,omitempty" xml:"Users,omitempty"`
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

func (s *UpdateCrowdRequest) SetLabel(v string) *UpdateCrowdRequest {
	s.Label = &v
	return s
}

func (s *UpdateCrowdRequest) SetName(v string) *UpdateCrowdRequest {
	s.Name = &v
	return s
}

func (s *UpdateCrowdRequest) SetUsers(v string) *UpdateCrowdRequest {
	s.Users = &v
	return s
}

type UpdateCrowdResponseBody struct {
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

type UpdateDomainRequest struct {
	// This parameter is required.
	BucketType  *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Condition   *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	CrowIds     *string `json:"CrowIds,omitempty" xml:"CrowIds,omitempty"`
	DebugUsers  *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Flow        *int64  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRequest) SetBucketType(v string) *UpdateDomainRequest {
	s.BucketType = &v
	return s
}

func (s *UpdateDomainRequest) SetCondition(v string) *UpdateDomainRequest {
	s.Condition = &v
	return s
}

func (s *UpdateDomainRequest) SetCrowIds(v string) *UpdateDomainRequest {
	s.CrowIds = &v
	return s
}

func (s *UpdateDomainRequest) SetDebugUsers(v string) *UpdateDomainRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateDomainRequest) SetDescription(v string) *UpdateDomainRequest {
	s.Description = &v
	return s
}

func (s *UpdateDomainRequest) SetFlow(v int64) *UpdateDomainRequest {
	s.Flow = &v
	return s
}

func (s *UpdateDomainRequest) SetName(v string) *UpdateDomainRequest {
	s.Name = &v
	return s
}

func (s *UpdateDomainRequest) SetProjectId(v string) *UpdateDomainRequest {
	s.ProjectId = &v
	return s
}

type UpdateDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponseBody) SetRequestId(v string) *UpdateDomainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponse) SetHeaders(v map[string]*string) *UpdateDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainResponse) SetStatusCode(v int32) *UpdateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainResponse) SetBody(v *UpdateDomainResponseBody) *UpdateDomainResponse {
	s.Body = v
	return s
}

type UpdateExperimentRequest struct {
	// This parameter is required.
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	Condition  *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// This parameter is required.
	CoreMetricId *string `json:"CoreMetricId,omitempty" xml:"CoreMetricId,omitempty"`
	CrowdIds     *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers   *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Flow    *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// This parameter is required.
	FocusMetricIds *string `json:"FocusMetricIds,omitempty" xml:"FocusMetricIds,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRequest) SetBucketType(v string) *UpdateExperimentRequest {
	s.BucketType = &v
	return s
}

func (s *UpdateExperimentRequest) SetCondition(v string) *UpdateExperimentRequest {
	s.Condition = &v
	return s
}

func (s *UpdateExperimentRequest) SetCoreMetricId(v string) *UpdateExperimentRequest {
	s.CoreMetricId = &v
	return s
}

func (s *UpdateExperimentRequest) SetCrowdIds(v string) *UpdateExperimentRequest {
	s.CrowdIds = &v
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

func (s *UpdateExperimentRequest) SetEndTime(v string) *UpdateExperimentRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateExperimentRequest) SetFlow(v int32) *UpdateExperimentRequest {
	s.Flow = &v
	return s
}

func (s *UpdateExperimentRequest) SetFocusMetricIds(v string) *UpdateExperimentRequest {
	s.FocusMetricIds = &v
	return s
}

func (s *UpdateExperimentRequest) SetName(v string) *UpdateExperimentRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentRequest) SetStartTime(v string) *UpdateExperimentRequest {
	s.StartTime = &v
	return s
}

type UpdateExperimentResponseBody struct {
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

type UpdateExperimentVersionRequest struct {
	Config      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CrowdIds    *string `json:"CrowdIds,omitempty" xml:"CrowdIds,omitempty"`
	DebugUsers  *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Flow        *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateExperimentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentVersionRequest) SetConfig(v string) *UpdateExperimentVersionRequest {
	s.Config = &v
	return s
}

func (s *UpdateExperimentVersionRequest) SetCrowdIds(v string) *UpdateExperimentVersionRequest {
	s.CrowdIds = &v
	return s
}

func (s *UpdateExperimentVersionRequest) SetDebugUsers(v string) *UpdateExperimentVersionRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateExperimentVersionRequest) SetDescription(v string) *UpdateExperimentVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentVersionRequest) SetFlow(v int32) *UpdateExperimentVersionRequest {
	s.Flow = &v
	return s
}

func (s *UpdateExperimentVersionRequest) SetName(v string) *UpdateExperimentVersionRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentVersionRequest) SetType(v string) *UpdateExperimentVersionRequest {
	s.Type = &v
	return s
}

type UpdateExperimentVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentVersionResponseBody) SetRequestId(v string) *UpdateExperimentVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExperimentVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentVersionResponse) SetHeaders(v map[string]*string) *UpdateExperimentVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentVersionResponse) SetStatusCode(v int32) *UpdateExperimentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentVersionResponse) SetBody(v *UpdateExperimentVersionResponseBody) *UpdateExperimentVersionResponse {
	s.Body = v
	return s
}

type UpdateFeatureRequest struct {
	// This parameter is required.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateFeatureRequest) SetStatus(v string) *UpdateFeatureRequest {
	s.Status = &v
	return s
}

type UpdateFeatureResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFeatureResponseBody) SetRequestId(v string) *UpdateFeatureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFeatureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateFeatureResponse) SetHeaders(v map[string]*string) *UpdateFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateFeatureResponse) SetStatusCode(v int32) *UpdateFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFeatureResponse) SetBody(v *UpdateFeatureResponseBody) *UpdateFeatureResponse {
	s.Body = v
	return s
}

type UpdateLayerRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId    *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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

func (s *UpdateLayerRequest) SetDomainId(v string) *UpdateLayerRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateLayerRequest) SetName(v string) *UpdateLayerRequest {
	s.Name = &v
	return s
}

func (s *UpdateLayerRequest) SetProjectId(v string) *UpdateLayerRequest {
	s.ProjectId = &v
	return s
}

type UpdateLayerResponseBody struct {
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

type UpdateMetricRequest struct {
	// This parameter is required.
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceTableMetaId *string `json:"SourceTableMetaId,omitempty" xml:"SourceTableMetaId,omitempty"`
}

func (s UpdateMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetricRequest) SetDefinition(v string) *UpdateMetricRequest {
	s.Definition = &v
	return s
}

func (s *UpdateMetricRequest) SetDescription(v string) *UpdateMetricRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetricRequest) SetName(v string) *UpdateMetricRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetricRequest) SetSourceTableMetaId(v string) *UpdateMetricRequest {
	s.SourceTableMetaId = &v
	return s
}

type UpdateMetricResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetricResponseBody) SetRequestId(v string) *UpdateMetricResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMetricResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetricResponse) SetHeaders(v map[string]*string) *UpdateMetricResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetricResponse) SetStatusCode(v int32) *UpdateMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetricResponse) SetBody(v *UpdateMetricResponseBody) *UpdateMetricResponse {
	s.Body = v
	return s
}

type UpdateMetricGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMetricGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetricGroupRequest) SetDescription(v string) *UpdateMetricGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetricGroupRequest) SetName(v string) *UpdateMetricGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetricGroupRequest) SetWorkspaceId(v string) *UpdateMetricGroupRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateMetricGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMetricGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetricGroupResponseBody) SetRequestId(v string) *UpdateMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMetricGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetricGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetricGroupResponse) SetHeaders(v map[string]*string) *UpdateMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetricGroupResponse) SetStatusCode(v int32) *UpdateMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetricGroupResponse) SetBody(v *UpdateMetricGroupResponseBody) *UpdateMetricGroupResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

type UpdateProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateTableMetaRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Fields []*UpdateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *UpdateTableMetaRequest) SetName(v string) *UpdateTableMetaRequest {
	s.Name = &v
	return s
}

type UpdateTableMetaRequestFields struct {
	// This parameter is required.
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTableMetaRequestFields) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableMetaRequestFields) GoString() string {
	return s.String()
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("paiabtest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 对层上的参数进行校验
//
// @param request - CheckLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckLayerResponse
func (client *Client) CheckLayerWithOptions(LayerId *string, request *CheckLayerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamNames)) {
		query["ParamNames"] = request.ParamNames
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckLayer"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/layers/" + tea.StringValue(openapiutil.GetEncodeParam(LayerId)) + "/action/check"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对层上的参数进行校验
//
// @param request - CheckLayerRequest
//
// @return CheckLayerResponse
func (client *Client) CheckLayer(LayerId *string, request *CheckLayerRequest) (_result *CheckLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckLayerResponse{}
	_body, _err := client.CheckLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建人群
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

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCrowd"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 创建人群
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
// 创建实验域
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketType)) {
		body["BucketType"] = request.BucketType
	}

	if !tea.BoolValue(util.IsUnset(request.Condition)) {
		body["Condition"] = request.Condition
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		body["CrowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Flow)) {
		body["Flow"] = request.Flow
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		body["LayerId"] = request.LayerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/domains"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验域
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验
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
	if !tea.BoolValue(util.IsUnset(request.BucketType)) {
		body["BucketType"] = request.BucketType
	}

	if !tea.BoolValue(util.IsUnset(request.Condition)) {
		body["Condition"] = request.Condition
	}

	if !tea.BoolValue(util.IsUnset(request.CoreMetricId)) {
		body["CoreMetricId"] = request.CoreMetricId
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		body["CrowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Flow)) {
		body["Flow"] = request.Flow
	}

	if !tea.BoolValue(util.IsUnset(request.FocusMetricIds)) {
		body["FocusMetricIds"] = request.FocusMetricIds
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		body["LayerId"] = request.LayerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperiment"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 创建实验
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
// 创建实验版本
//
// @param request - CreateExperimentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentVersionResponse
func (client *Client) CreateExperimentVersionWithOptions(request *CreateExperimentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		body["CrowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		body["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Flow)) {
		body["Flow"] = request.Flow
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
		Action:      tea.String("CreateExperimentVersion"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentversions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验版本
//
// @param request - CreateExperimentVersionRequest
//
// @return CreateExperimentVersionResponse
func (client *Client) CreateExperimentVersion(request *CreateExperimentVersionRequest) (_result *CreateExperimentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentVersionResponse{}
	_body, _err := client.CreateExperimentVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Feature
//
// @param request - CreateFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureResponse
func (client *Client) CreateFeatureWithOptions(request *CreateFeatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFeature"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/features"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Feature
//
// @param request - CreateFeatureRequest
//
// @return CreateFeatureResponse
func (client *Client) CreateFeature(request *CreateFeatureRequest) (_result *CreateFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureResponse{}
	_body, _err := client.CreateFeatureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验层
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

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		body["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLayer"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 创建实验层
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
// 创建指标
//
// @param request - CreateMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricResponse
func (client *Client) CreateMetricWithOptions(request *CreateMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMetricResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.MetricGroupId)) {
		body["MetricGroupId"] = request.MetricGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableMetaId)) {
		body["SourceTableMetaId"] = request.SourceTableMetaId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMetric"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metrics"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建指标
//
// @param request - CreateMetricRequest
//
// @return CreateMetricResponse
func (client *Client) CreateMetric(request *CreateMetricRequest) (_result *CreateMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMetricResponse{}
	_body, _err := client.CreateMetricWithOptions(request, headers, runtime)
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
// @param request - CreateMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricGroupResponse
func (client *Client) CreateMetricGroupWithOptions(request *CreateMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("CreateMetricGroup"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metricgroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建指标组
//
// @param request - CreateMetricGroupRequest
//
// @return CreateMetricGroupResponse
func (client *Client) CreateMetricGroup(request *CreateMetricGroupRequest) (_result *CreateMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMetricGroupResponse{}
	_body, _err := client.CreateMetricGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验项目
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据表
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
	if !tea.BoolValue(util.IsUnset(request.DatasourceInfo)) {
		body["DatasourceInfo"] = request.DatasourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.DatasourceType)) {
		body["DatasourceType"] = request.DatasourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTableMeta"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 创建数据表
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
// 删除指定的人群
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowdWithOptions(CrowdId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCrowdResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCrowd"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 删除指定的人群
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowd(CrowdId *string) (_result *DeleteCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCrowdResponse{}
	_body, _err := client.DeleteCrowdWithOptions(CrowdId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的实验域
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(DomainId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/domains/" + tea.StringValue(openapiutil.GetEncodeParam(DomainId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的实验域
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(DomainId *string) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(DomainId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的实验
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperiment"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 删除指定的实验
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperiment(ExperimentId *string) (_result *DeleteExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的实验版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentVersionResponse
func (client *Client) DeleteExperimentVersionWithOptions(ExperimentVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperimentVersion"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentversions/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentVersionId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的实验版本
//
// @return DeleteExperimentVersionResponse
func (client *Client) DeleteExperimentVersion(ExperimentVersionId *string) (_result *DeleteExperimentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentVersionResponse{}
	_body, _err := client.DeleteExperimentVersionWithOptions(ExperimentVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Feature
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureResponse
func (client *Client) DeleteFeatureWithOptions(FeatureId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFeatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFeature"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/features/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Feature
//
// @return DeleteFeatureResponse
func (client *Client) DeleteFeature(FeatureId *string) (_result *DeleteFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFeatureResponse{}
	_body, _err := client.DeleteFeatureWithOptions(FeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的实验层
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayerWithOptions(LayerId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLayerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayer"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 删除指定的实验层
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayer(LayerId *string) (_result *DeleteLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLayerResponse{}
	_body, _err := client.DeleteLayerWithOptions(LayerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定指标
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricResponse
func (client *Client) DeleteMetricWithOptions(MetricId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMetricResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMetric"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metrics/" + tea.StringValue(openapiutil.GetEncodeParam(MetricId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定指标
//
// @return DeleteMetricResponse
func (client *Client) DeleteMetric(MetricId *string) (_result *DeleteMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMetricResponse{}
	_body, _err := client.DeleteMetricWithOptions(MetricId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的指标组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricGroupResponse
func (client *Client) DeleteMetricGroupWithOptions(MetricGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMetricGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMetricGroup"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(MetricGroupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的指标组
//
// @return DeleteMetricGroupResponse
func (client *Client) DeleteMetricGroup(MetricGroupId *string) (_result *DeleteMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMetricGroupResponse{}
	_body, _err := client.DeleteMetricGroupWithOptions(MetricGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实验项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验项目
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(ProjectId *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMetaWithOptions(TableMetaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTableMetaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTableMeta"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 删除数据表
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMeta(TableMetaId *string) (_result *DeleteTableMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableMetaResponse{}
	_body, _err := client.DeleteTableMetaWithOptions(TableMetaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定人群详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrowdResponse
func (client *Client) GetCrowdWithOptions(CrowdId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCrowdResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCrowd"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/crowds/" + tea.StringValue(openapiutil.GetEncodeParam(CrowdId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定人群详情
//
// @return GetCrowdResponse
func (client *Client) GetCrowd(CrowdId *string) (_result *GetCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCrowdResponse{}
	_body, _err := client.GetCrowdWithOptions(CrowdId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定实验域详情
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(DomainId *string, request *GetDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomain"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/domains/" + tea.StringValue(openapiutil.GetEncodeParam(DomainId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定实验域详情
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
func (client *Client) GetDomain(DomainId *string, request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(DomainId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定实验的详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperiment"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取指定实验的详情
//
// @return GetExperimentResponse
func (client *Client) GetExperiment(ExperimentId *string) (_result *GetExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定实验版本的详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentVersionResponse
func (client *Client) GetExperimentVersionWithOptions(ExperimentVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentVersion"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentversions/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentVersionId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定实验版本的详情
//
// @return GetExperimentVersionResponse
func (client *Client) GetExperimentVersion(ExperimentVersionId *string) (_result *GetExperimentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentVersionResponse{}
	_body, _err := client.GetExperimentVersionWithOptions(ExperimentVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Feature详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureResponse
func (client *Client) GetFeatureWithOptions(FeatureId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFeatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeature"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/features/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Feature详情
//
// @return GetFeatureResponse
func (client *Client) GetFeature(FeatureId *string) (_result *GetFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureResponse{}
	_body, _err := client.GetFeatureWithOptions(FeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定的实验层详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerResponse
func (client *Client) GetLayerWithOptions(LayerId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLayerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayer"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取指定的实验层详情
//
// @return GetLayerResponse
func (client *Client) GetLayer(LayerId *string) (_result *GetLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLayerResponse{}
	_body, _err := client.GetLayerWithOptions(LayerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指标详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetricResponse
func (client *Client) GetMetricWithOptions(MetricId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMetricResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetric"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metrics/" + tea.StringValue(openapiutil.GetEncodeParam(MetricId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指标详情
//
// @return GetMetricResponse
func (client *Client) GetMetric(MetricId *string) (_result *GetMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMetricResponse{}
	_body, _err := client.GetMetricWithOptions(MetricId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指标组的详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetricGroupResponse
func (client *Client) GetMetricGroupWithOptions(MetricGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMetricGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetricGroup"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(MetricGroupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指标组的详细信息
//
// @return GetMetricGroupResponse
func (client *Client) GetMetricGroup(MetricGroupId *string) (_result *GetMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMetricGroupResponse{}
	_body, _err := client.GetMetricGroupWithOptions(MetricGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定的实验项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定的实验项目
//
// @return GetProjectResponse
func (client *Client) GetProject(ProjectId *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据表详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableMetaResponse
func (client *Client) GetTableMetaWithOptions(TableMetaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableMetaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableMeta"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取数据表详情
//
// @return GetTableMetaResponse
func (client *Client) GetTableMeta(TableMetaId *string) (_result *GetTableMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableMetaResponse{}
	_body, _err := client.GetTableMetaWithOptions(TableMetaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取人群列表
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
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdId)) {
		query["CrowdId"] = request.CrowdId
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdName)) {
		query["CrowdName"] = request.CrowdName
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCrowds"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取人群列表
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
// 获取实验域列表
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		query["LayerId"] = request.LayerId
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验域列表
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验版本列表
//
// @param request - ListExperimentVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentVersionsResponse
func (client *Client) ListExperimentVersionsWithOptions(request *ListExperimentVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExperimentVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
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
		Action:      tea.String("ListExperimentVersions"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentversions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验版本列表
//
// @param request - ListExperimentVersionsRequest
//
// @return ListExperimentVersionsResponse
func (client *Client) ListExperimentVersions(request *ListExperimentVersionsRequest) (_result *ListExperimentVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentVersionsResponse{}
	_body, _err := client.ListExperimentVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验列表
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
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentName)) {
		query["ExperimentName"] = request.ExperimentName
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		query["LayerId"] = request.LayerId
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
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
		Action:      tea.String("ListExperiments"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取实验列表
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
// 获取Faeture列表
//
// @param request - ListFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeaturesResponse
func (client *Client) ListFeaturesWithOptions(request *ListFeaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureId)) {
		query["FeatureId"] = request.FeatureId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureName)) {
		query["FeatureName"] = request.FeatureName
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Action:      tea.String("ListFeatures"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/features"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Faeture列表
//
// @param request - ListFeaturesRequest
//
// @return ListFeaturesResponse
func (client *Client) ListFeatures(request *ListFeaturesRequest) (_result *ListFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeaturesResponse{}
	_body, _err := client.ListFeaturesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验层列表
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
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.LayerId)) {
		query["LayerId"] = request.LayerId
	}

	if !tea.BoolValue(util.IsUnset(request.LayerName)) {
		query["LayerName"] = request.LayerName
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayers"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取实验层列表
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
// 获取指标组列表
//
// @param request - ListMetricGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetricGroupsResponse
func (client *Client) ListMetricGroupsWithOptions(request *ListMetricGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMetricGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.MetricGroupId)) {
		query["MetricGroupId"] = request.MetricGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricGroupName)) {
		query["MetricGroupName"] = request.MetricGroupName
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

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMetricGroups"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metricgroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMetricGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指标组列表
//
// @param request - ListMetricGroupsRequest
//
// @return ListMetricGroupsResponse
func (client *Client) ListMetricGroups(request *ListMetricGroupsRequest) (_result *ListMetricGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMetricGroupsResponse{}
	_body, _err := client.ListMetricGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指标列表
//
// @param request - ListMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetricsResponse
func (client *Client) ListMetricsWithOptions(request *ListMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.MetricGroupId)) {
		query["MetricGroupId"] = request.MetricGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricId)) {
		query["MetricId"] = request.MetricId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
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
		Action:      tea.String("ListMetrics"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指标列表
//
// @param request - ListMetricsRequest
//
// @return ListMetricsResponse
func (client *Client) ListMetrics(request *ListMetricsRequest) (_result *ListMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMetricsResponse{}
	_body, _err := client.ListMetricsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验项目列表
//
// @param request - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验项目列表
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据表列表
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
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.DatasourceType)) {
		query["DatasourceType"] = request.DatasourceType
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

	if !tea.BoolValue(util.IsUnset(request.TableMetaId)) {
		query["TableMetaId"] = request.TableMetaId
	}

	if !tea.BoolValue(util.IsUnset(request.TableMetaName)) {
		query["TableMetaName"] = request.TableMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTableMetas"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 获取数据表列表
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
// 对实验版本推全
//
// @param request - PushAllExperimentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushAllExperimentVersionResponse
func (client *Client) PushAllExperimentVersionWithOptions(ExperimentVersionId *string, request *PushAllExperimentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushAllExperimentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureName)) {
		body["FeatureName"] = request.FeatureName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushAllExperimentVersion"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentversions/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentVersionId)) + "/action/pushall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushAllExperimentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对实验版本推全
//
// @param request - PushAllExperimentVersionRequest
//
// @return PushAllExperimentVersionResponse
func (client *Client) PushAllExperimentVersion(ExperimentVersionId *string, request *PushAllExperimentVersionRequest) (_result *PushAllExperimentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushAllExperimentVersionResponse{}
	_body, _err := client.PushAllExperimentVersionWithOptions(ExperimentVersionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动实验
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartExperimentResponse
func (client *Client) StartExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartExperiment"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/action/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动实验
//
// @return StartExperimentResponse
func (client *Client) StartExperiment(ExperimentId *string) (_result *StartExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartExperimentResponse{}
	_body, _err := client.StartExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止实验
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopExperimentResponse
func (client *Client) StopExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopExperiment"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/action/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止实验
//
// @return StopExperimentResponse
func (client *Client) StopExperiment(ExperimentId *string) (_result *StopExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopExperimentResponse{}
	_body, _err := client.StopExperimentWithOptions(ExperimentId, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCrowd"),
		Version:     tea.String("2024-01-19"),
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
// 更新指定实验域
//
// @param request - UpdateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithOptions(DomainId *string, request *UpdateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketType)) {
		body["BucketType"] = request.BucketType
	}

	if !tea.BoolValue(util.IsUnset(request.Condition)) {
		body["Condition"] = request.Condition
	}

	if !tea.BoolValue(util.IsUnset(request.CrowIds)) {
		body["CrowIds"] = request.CrowIds
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Flow)) {
		body["Flow"] = request.Flow
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDomain"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/domains/" + tea.StringValue(openapiutil.GetEncodeParam(DomainId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定实验域
//
// @param request - UpdateDomainRequest
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomain(DomainId *string, request *UpdateDomainRequest) (_result *UpdateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDomainResponse{}
	_body, _err := client.UpdateDomainWithOptions(DomainId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定的实验
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
	if !tea.BoolValue(util.IsUnset(request.BucketType)) {
		body["BucketType"] = request.BucketType
	}

	if !tea.BoolValue(util.IsUnset(request.Condition)) {
		body["Condition"] = request.Condition
	}

	if !tea.BoolValue(util.IsUnset(request.CoreMetricId)) {
		body["CoreMetricId"] = request.CoreMetricId
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		body["CrowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Flow)) {
		body["Flow"] = request.Flow
	}

	if !tea.BoolValue(util.IsUnset(request.FocusMetricIds)) {
		body["FocusMetricIds"] = request.FocusMetricIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperiment"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 更新指定的实验
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
// 更新指定的实验版本
//
// @param request - UpdateExperimentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentVersionResponse
func (client *Client) UpdateExperimentVersionWithOptions(ExperimentVersionId *string, request *UpdateExperimentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		body["CrowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.DebugUsers)) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Flow)) {
		body["Flow"] = request.Flow
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
		Action:      tea.String("UpdateExperimentVersion"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentversions/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentVersionId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定的实验版本
//
// @param request - UpdateExperimentVersionRequest
//
// @return UpdateExperimentVersionResponse
func (client *Client) UpdateExperimentVersion(ExperimentVersionId *string, request *UpdateExperimentVersionRequest) (_result *UpdateExperimentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentVersionResponse{}
	_body, _err := client.UpdateExperimentVersionWithOptions(ExperimentVersionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Feature
//
// @param request - UpdateFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFeatureResponse
func (client *Client) UpdateFeatureWithOptions(FeatureId *string, request *UpdateFeatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFeature"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/features/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Feature
//
// @param request - UpdateFeatureRequest
//
// @return UpdateFeatureResponse
func (client *Client) UpdateFeature(FeatureId *string, request *UpdateFeatureRequest) (_result *UpdateFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFeatureResponse{}
	_body, _err := client.UpdateFeatureWithOptions(FeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定的实验层
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

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		body["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLayer"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 更新指定的实验层
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
// 更新指标
//
// @param request - UpdateMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetricResponse
func (client *Client) UpdateMetricWithOptions(MetricId *string, request *UpdateMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMetricResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTableMetaId)) {
		body["SourceTableMetaId"] = request.SourceTableMetaId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMetric"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metrics/" + tea.StringValue(openapiutil.GetEncodeParam(MetricId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指标
//
// @param request - UpdateMetricRequest
//
// @return UpdateMetricResponse
func (client *Client) UpdateMetric(MetricId *string, request *UpdateMetricRequest) (_result *UpdateMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMetricResponse{}
	_body, _err := client.UpdateMetricWithOptions(MetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定的指标组
//
// @param request - UpdateMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetricGroupResponse
func (client *Client) UpdateMetricGroupWithOptions(MetricGroupId *string, request *UpdateMetricGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMetricGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("UpdateMetricGroup"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/metricgroups/" + tea.StringValue(openapiutil.GetEncodeParam(MetricGroupId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定的指标组
//
// @param request - UpdateMetricGroupRequest
//
// @return UpdateMetricGroupResponse
func (client *Client) UpdateMetricGroup(MetricGroupId *string, request *UpdateMetricGroupRequest) (_result *UpdateMetricGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMetricGroupResponse{}
	_body, _err := client.UpdateMetricGroupWithOptions(MetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定的实验项目
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(ProjectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2024-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定的实验项目
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(ProjectId *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据表
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

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTableMeta"),
		Version:     tea.String("2024-01-19"),
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

// Summary:
//
// 更新数据表
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
