// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateClusterCheckRequest struct {
	AddonName       *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ReportName      *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
}

func (s CreateClusterCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterCheckRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterCheckRequest) SetAddonName(v string) *CreateClusterCheckRequest {
	s.AddonName = &v
	return s
}

func (s *CreateClusterCheckRequest) SetClusterRegionId(v string) *CreateClusterCheckRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *CreateClusterCheckRequest) SetClusterUid(v string) *CreateClusterCheckRequest {
	s.ClusterUid = &v
	return s
}

func (s *CreateClusterCheckRequest) SetOwnerUid(v int64) *CreateClusterCheckRequest {
	s.OwnerUid = &v
	return s
}

func (s *CreateClusterCheckRequest) SetReportName(v string) *CreateClusterCheckRequest {
	s.ReportName = &v
	return s
}

type CreateClusterCheckResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReportUid *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterCheckResponseBody) SetCode(v string) *CreateClusterCheckResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClusterCheckResponseBody) SetIsSuccess(v bool) *CreateClusterCheckResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateClusterCheckResponseBody) SetReportUid(v string) *CreateClusterCheckResponseBody {
	s.ReportUid = &v
	return s
}

func (s *CreateClusterCheckResponseBody) SetRequestId(v string) *CreateClusterCheckResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterCheckResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterCheckResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterCheckResponse) SetHeaders(v map[string]*string) *CreateClusterCheckResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterCheckResponse) SetBody(v *CreateClusterCheckResponseBody) *CreateClusterCheckResponse {
	s.Body = v
	return s
}

type CreateClusterReportRequest struct {
	AddonName       *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ReportName      *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
}

func (s CreateClusterReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterReportRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterReportRequest) SetAddonName(v string) *CreateClusterReportRequest {
	s.AddonName = &v
	return s
}

func (s *CreateClusterReportRequest) SetClusterRegionId(v string) *CreateClusterReportRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *CreateClusterReportRequest) SetClusterUid(v string) *CreateClusterReportRequest {
	s.ClusterUid = &v
	return s
}

func (s *CreateClusterReportRequest) SetOwnerUid(v int64) *CreateClusterReportRequest {
	s.OwnerUid = &v
	return s
}

func (s *CreateClusterReportRequest) SetReportName(v string) *CreateClusterReportRequest {
	s.ReportName = &v
	return s
}

type CreateClusterReportResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReportUid *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterReportResponseBody) SetCode(v string) *CreateClusterReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClusterReportResponseBody) SetIsSuccess(v bool) *CreateClusterReportResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateClusterReportResponseBody) SetReportUid(v string) *CreateClusterReportResponseBody {
	s.ReportUid = &v
	return s
}

func (s *CreateClusterReportResponseBody) SetRequestId(v string) *CreateClusterReportResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterReportResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterReportResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterReportResponse) SetHeaders(v map[string]*string) *CreateClusterReportResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterReportResponse) SetBody(v *CreateClusterReportResponseBody) *CreateClusterReportResponse {
	s.Body = v
	return s
}

type CreateDiagnosisRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	Target          *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisRequest) SetClusterRegionId(v string) *CreateDiagnosisRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *CreateDiagnosisRequest) SetClusterUid(v string) *CreateDiagnosisRequest {
	s.ClusterUid = &v
	return s
}

func (s *CreateDiagnosisRequest) SetOwnerUid(v int64) *CreateDiagnosisRequest {
	s.OwnerUid = &v
	return s
}

func (s *CreateDiagnosisRequest) SetTarget(v string) *CreateDiagnosisRequest {
	s.Target = &v
	return s
}

func (s *CreateDiagnosisRequest) SetType(v string) *CreateDiagnosisRequest {
	s.Type = &v
	return s
}

type CreateDiagnosisResponseBody struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DiagnosisId *string `json:"DiagnosisId,omitempty" xml:"DiagnosisId,omitempty"`
	// bool of success
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisResponseBody) SetCode(v string) *CreateDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiagnosisResponseBody) SetDiagnosisId(v string) *CreateDiagnosisResponseBody {
	s.DiagnosisId = &v
	return s
}

func (s *CreateDiagnosisResponseBody) SetIsSuccess(v bool) *CreateDiagnosisResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateDiagnosisResponseBody) SetRequestId(v string) *CreateDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiagnosisResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisResponse) SetHeaders(v map[string]*string) *CreateDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosisResponse) SetBody(v *CreateDiagnosisResponseBody) *CreateDiagnosisResponse {
	s.Body = v
	return s
}

type CreateReportTaskRuleRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ReportName      *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	ScheduleRule    *string `json:"ScheduleRule,omitempty" xml:"ScheduleRule,omitempty"`
	TimeZoneId      *string `json:"TimeZoneId,omitempty" xml:"TimeZoneId,omitempty"`
}

func (s CreateReportTaskRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReportTaskRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateReportTaskRuleRequest) SetClusterRegionId(v string) *CreateReportTaskRuleRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *CreateReportTaskRuleRequest) SetClusterUid(v string) *CreateReportTaskRuleRequest {
	s.ClusterUid = &v
	return s
}

func (s *CreateReportTaskRuleRequest) SetOwnerUid(v int64) *CreateReportTaskRuleRequest {
	s.OwnerUid = &v
	return s
}

func (s *CreateReportTaskRuleRequest) SetReportName(v string) *CreateReportTaskRuleRequest {
	s.ReportName = &v
	return s
}

func (s *CreateReportTaskRuleRequest) SetScheduleRule(v string) *CreateReportTaskRuleRequest {
	s.ScheduleRule = &v
	return s
}

func (s *CreateReportTaskRuleRequest) SetTimeZoneId(v string) *CreateReportTaskRuleRequest {
	s.TimeZoneId = &v
	return s
}

type CreateReportTaskRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateReportTaskRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReportTaskRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportTaskRuleResponseBody) SetCode(v string) *CreateReportTaskRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateReportTaskRuleResponseBody) SetIsSuccess(v bool) *CreateReportTaskRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateReportTaskRuleResponseBody) SetRequestId(v string) *CreateReportTaskRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReportTaskRuleResponseBody) SetRuleId(v string) *CreateReportTaskRuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateReportTaskRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateReportTaskRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReportTaskRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReportTaskRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateReportTaskRuleResponse) SetHeaders(v map[string]*string) *CreateReportTaskRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateReportTaskRuleResponse) SetBody(v *CreateReportTaskRuleResponseBody) *CreateReportTaskRuleResponse {
	s.Body = v
	return s
}

type DeleteReportTaskRuleRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	RuleId          *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteReportTaskRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportTaskRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteReportTaskRuleRequest) SetClusterRegionId(v string) *DeleteReportTaskRuleRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *DeleteReportTaskRuleRequest) SetClusterUid(v string) *DeleteReportTaskRuleRequest {
	s.ClusterUid = &v
	return s
}

func (s *DeleteReportTaskRuleRequest) SetOwnerUid(v int64) *DeleteReportTaskRuleRequest {
	s.OwnerUid = &v
	return s
}

func (s *DeleteReportTaskRuleRequest) SetRuleId(v string) *DeleteReportTaskRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteReportTaskRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReportUid *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReportTaskRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportTaskRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReportTaskRuleResponseBody) SetCode(v string) *DeleteReportTaskRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteReportTaskRuleResponseBody) SetIsSuccess(v bool) *DeleteReportTaskRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteReportTaskRuleResponseBody) SetReportUid(v string) *DeleteReportTaskRuleResponseBody {
	s.ReportUid = &v
	return s
}

func (s *DeleteReportTaskRuleResponseBody) SetRequestId(v string) *DeleteReportTaskRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteReportTaskRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteReportTaskRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteReportTaskRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportTaskRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteReportTaskRuleResponse) SetHeaders(v map[string]*string) *DeleteReportTaskRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteReportTaskRuleResponse) SetBody(v *DeleteReportTaskRuleResponseBody) *DeleteReportTaskRuleResponse {
	s.Body = v
	return s
}

type GetClusterCheckItemRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	EnableFilter    *bool   `json:"EnableFilter,omitempty" xml:"EnableFilter,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ReportUid       *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
}

func (s GetClusterCheckItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckItemRequest) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemRequest) SetClusterRegionId(v string) *GetClusterCheckItemRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *GetClusterCheckItemRequest) SetClusterUid(v string) *GetClusterCheckItemRequest {
	s.ClusterUid = &v
	return s
}

func (s *GetClusterCheckItemRequest) SetEnableFilter(v bool) *GetClusterCheckItemRequest {
	s.EnableFilter = &v
	return s
}

func (s *GetClusterCheckItemRequest) SetLanguage(v string) *GetClusterCheckItemRequest {
	s.Language = &v
	return s
}

func (s *GetClusterCheckItemRequest) SetOwnerUid(v int64) *GetClusterCheckItemRequest {
	s.OwnerUid = &v
	return s
}

func (s *GetClusterCheckItemRequest) SetReportUid(v string) *GetClusterCheckItemRequest {
	s.ReportUid = &v
	return s
}

type GetClusterCheckItemResponseBody struct {
	CheckResult *GetClusterCheckItemResponseBodyCheckResult `json:"CheckResult,omitempty" xml:"CheckResult,omitempty" type:"Struct"`
	Code        *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess   *bool                                       `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterCheckItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemResponseBody) SetCheckResult(v *GetClusterCheckItemResponseBodyCheckResult) *GetClusterCheckItemResponseBody {
	s.CheckResult = v
	return s
}

func (s *GetClusterCheckItemResponseBody) SetCode(v string) *GetClusterCheckItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterCheckItemResponseBody) SetIsSuccess(v bool) *GetClusterCheckItemResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetClusterCheckItemResponseBody) SetRequestId(v string) *GetClusterCheckItemResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterCheckItemResponseBodyCheckResult struct {
	CheckId *string                                              `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Cluster []*GetClusterCheckItemResponseBodyCheckResultCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
	Nodes   []*GetClusterCheckItemResponseBodyCheckResultNodes   `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetClusterCheckItemResponseBodyCheckResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckItemResponseBodyCheckResult) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemResponseBodyCheckResult) SetCheckId(v string) *GetClusterCheckItemResponseBodyCheckResult {
	s.CheckId = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResult) SetCluster(v []*GetClusterCheckItemResponseBodyCheckResultCluster) *GetClusterCheckItemResponseBodyCheckResult {
	s.Cluster = v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResult) SetNodes(v []*GetClusterCheckItemResponseBodyCheckResultNodes) *GetClusterCheckItemResponseBodyCheckResult {
	s.Nodes = v
	return s
}

type GetClusterCheckItemResponseBodyCheckResultCluster struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	Fix     *string `json:"Fix,omitempty" xml:"Fix,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Level   *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Refer   *string `json:"Refer,omitempty" xml:"Refer,omitempty"`
	Value   *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetClusterCheckItemResponseBodyCheckResultCluster) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckItemResponseBodyCheckResultCluster) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetDesc(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Desc = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetDisplay(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Display = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetFix(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Fix = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetGroup(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Group = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetLevel(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Level = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetMessage(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Message = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetName(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Name = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetRefer(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Refer = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultCluster) SetValue(v string) *GetClusterCheckItemResponseBodyCheckResultCluster {
	s.Value = &v
	return s
}

type GetClusterCheckItemResponseBodyCheckResultNodes struct {
	Desc    *string                                `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Display *string                                `json:"Display,omitempty" xml:"Display,omitempty"`
	Fix     *string                                `json:"Fix,omitempty" xml:"Fix,omitempty"`
	Group   *string                                `json:"Group,omitempty" xml:"Group,omitempty"`
	Level   map[string]*CheckResultNodesLevelValue `json:"Level,omitempty" xml:"Level,omitempty"`
	Name    *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Refer   *string                                `json:"Refer,omitempty" xml:"Refer,omitempty"`
}

func (s GetClusterCheckItemResponseBodyCheckResultNodes) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckItemResponseBodyCheckResultNodes) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetDesc(v string) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Desc = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetDisplay(v string) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Display = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetFix(v string) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Fix = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetGroup(v string) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Group = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetLevel(v map[string]*CheckResultNodesLevelValue) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Level = v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetName(v string) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Name = &v
	return s
}

func (s *GetClusterCheckItemResponseBodyCheckResultNodes) SetRefer(v string) *GetClusterCheckItemResponseBodyCheckResultNodes {
	s.Refer = &v
	return s
}

type GetClusterCheckItemResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterCheckItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterCheckItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckItemResponse) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemResponse) SetHeaders(v map[string]*string) *GetClusterCheckItemResponse {
	s.Headers = v
	return s
}

func (s *GetClusterCheckItemResponse) SetBody(v *GetClusterCheckItemResponseBody) *GetClusterCheckItemResponse {
	s.Body = v
	return s
}

type GetClusterCheckResultRequest struct {
	ChecklistName   *string `json:"ChecklistName,omitempty" xml:"ChecklistName,omitempty"`
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ReportUid       *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
}

func (s GetClusterCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultRequest) SetChecklistName(v string) *GetClusterCheckResultRequest {
	s.ChecklistName = &v
	return s
}

func (s *GetClusterCheckResultRequest) SetClusterRegionId(v string) *GetClusterCheckResultRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *GetClusterCheckResultRequest) SetOwnerUid(v int64) *GetClusterCheckResultRequest {
	s.OwnerUid = &v
	return s
}

func (s *GetClusterCheckResultRequest) SetReportUid(v string) *GetClusterCheckResultRequest {
	s.ReportUid = &v
	return s
}

type GetClusterCheckResultResponseBody struct {
	CheckEntryResults []*GetClusterCheckResultResponseBodyCheckEntryResults `json:"CheckEntryResults,omitempty" xml:"CheckEntryResults,omitempty" type:"Repeated"`
	CheckStatus       *string                                               `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckSummary      *GetClusterCheckResultResponseBodyCheckSummary        `json:"CheckSummary,omitempty" xml:"CheckSummary,omitempty" type:"Struct"`
	Code              *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess         *bool                                                 `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Name              *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	ReportUid         *string                                               `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBody) SetCheckEntryResults(v []*GetClusterCheckResultResponseBodyCheckEntryResults) *GetClusterCheckResultResponseBody {
	s.CheckEntryResults = v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetCheckStatus(v string) *GetClusterCheckResultResponseBody {
	s.CheckStatus = &v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetCheckSummary(v *GetClusterCheckResultResponseBodyCheckSummary) *GetClusterCheckResultResponseBody {
	s.CheckSummary = v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetCode(v string) *GetClusterCheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetIsSuccess(v bool) *GetClusterCheckResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetName(v string) *GetClusterCheckResultResponseBody {
	s.Name = &v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetReportUid(v string) *GetClusterCheckResultResponseBody {
	s.ReportUid = &v
	return s
}

func (s *GetClusterCheckResultResponseBody) SetRequestId(v string) *GetClusterCheckResultResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResults struct {
	CheckSummary          *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary            `json:"CheckSummary,omitempty" xml:"CheckSummary,omitempty" type:"Struct"`
	EntryGroupName        *string                                                                    `json:"EntryGroupName,omitempty" xml:"EntryGroupName,omitempty"`
	EntryName             *string                                                                    `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	ErrorInstanceResult   []*GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult   `json:"ErrorInstanceResult,omitempty" xml:"ErrorInstanceResult,omitempty" type:"Repeated"`
	NormalInstanceResult  []*GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult  `json:"NormalInstanceResult,omitempty" xml:"NormalInstanceResult,omitempty" type:"Repeated"`
	UnknownInstanceResult []*GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult `json:"UnknownInstanceResult,omitempty" xml:"UnknownInstanceResult,omitempty" type:"Repeated"`
	WarnInstanceResult    []*GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult    `json:"WarnInstanceResult,omitempty" xml:"WarnInstanceResult,omitempty" type:"Repeated"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResults) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResults) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetCheckSummary(v *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.CheckSummary = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetEntryGroupName(v string) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.EntryGroupName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetEntryName(v string) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.EntryName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetErrorInstanceResult(v []*GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.ErrorInstanceResult = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetNormalInstanceResult(v []*GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.NormalInstanceResult = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetUnknownInstanceResult(v []*GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.UnknownInstanceResult = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResults) SetWarnInstanceResult(v []*GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) *GetClusterCheckResultResponseBodyCheckEntryResults {
	s.WarnInstanceResult = v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCount   *int32  `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	NormalCount  *int32  `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	Process      *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	UnknownCount *int32  `json:"UnknownCount,omitempty" xml:"UnknownCount,omitempty"`
	WarnCount    *int32  `json:"WarnCount,omitempty" xml:"WarnCount,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) SetCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary {
	s.Code = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) SetErrorCount(v int32) *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary {
	s.ErrorCount = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) SetNormalCount(v int32) *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary {
	s.NormalCount = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) SetProcess(v int32) *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary {
	s.Process = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) SetUnknownCount(v int32) *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary {
	s.UnknownCount = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary) SetWarnCount(v int32) *GetClusterCheckResultResponseBodyCheckEntryResultsCheckSummary {
	s.WarnCount = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult struct {
	CheckpointResults []*GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults `json:"CheckpointResults,omitempty" xml:"CheckpointResults,omitempty" type:"Repeated"`
	Extend            map[string]interface{}                                                                    `json:"Extend,omitempty" xml:"Extend,omitempty"`
	InstanceId        *string                                                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType      *string                                                                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) SetCheckpointResults(v []*GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult {
	s.CheckpointResults = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) SetExtend(v map[string]interface{}) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult {
	s.Extend = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) SetInstanceId(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult {
	s.InstanceId = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) SetInstanceName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult {
	s.InstanceName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult) SetInstanceType(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResult {
	s.InstanceType = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults struct {
	AdviseCode     *string `json:"AdviseCode,omitempty" xml:"AdviseCode,omitempty"`
	AffectCode     *string `json:"AffectCode,omitempty" xml:"AffectCode,omitempty"`
	CheckpointName *string `json:"CheckpointName,omitempty" xml:"CheckpointName,omitempty"`
	MessageCode    *string `json:"MessageCode,omitempty" xml:"MessageCode,omitempty"`
	MessageLevel   *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) SetAdviseCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults {
	s.AdviseCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) SetAffectCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults {
	s.AffectCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) SetCheckpointName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults {
	s.CheckpointName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) SetMessageCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults {
	s.MessageCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults) SetMessageLevel(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsErrorInstanceResultCheckpointResults {
	s.MessageLevel = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult struct {
	CheckpointResults []*GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults `json:"CheckpointResults,omitempty" xml:"CheckpointResults,omitempty" type:"Repeated"`
	Extend            map[string]interface{}                                                                     `json:"Extend,omitempty" xml:"Extend,omitempty"`
	InstanceId        *string                                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType      *string                                                                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) SetCheckpointResults(v []*GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult {
	s.CheckpointResults = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) SetExtend(v map[string]interface{}) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult {
	s.Extend = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) SetInstanceId(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult {
	s.InstanceId = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) SetInstanceName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult {
	s.InstanceName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult) SetInstanceType(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResult {
	s.InstanceType = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults struct {
	AdviseCode     *string `json:"AdviseCode,omitempty" xml:"AdviseCode,omitempty"`
	AffectCode     *string `json:"AffectCode,omitempty" xml:"AffectCode,omitempty"`
	CheckpointName *string `json:"CheckpointName,omitempty" xml:"CheckpointName,omitempty"`
	MessageCode    *string `json:"MessageCode,omitempty" xml:"MessageCode,omitempty"`
	MessageLevel   *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) SetAdviseCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults {
	s.AdviseCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) SetAffectCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults {
	s.AffectCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) SetCheckpointName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults {
	s.CheckpointName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) SetMessageCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults {
	s.MessageCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults) SetMessageLevel(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsNormalInstanceResultCheckpointResults {
	s.MessageLevel = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult struct {
	CheckpointResults []*GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults `json:"CheckpointResults,omitempty" xml:"CheckpointResults,omitempty" type:"Repeated"`
	Extend            map[string]interface{}                                                                      `json:"Extend,omitempty" xml:"Extend,omitempty"`
	InstanceId        *string                                                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                                                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType      *string                                                                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) SetCheckpointResults(v []*GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult {
	s.CheckpointResults = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) SetExtend(v map[string]interface{}) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult {
	s.Extend = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) SetInstanceId(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult {
	s.InstanceId = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) SetInstanceName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult {
	s.InstanceName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult) SetInstanceType(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResult {
	s.InstanceType = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults struct {
	AdviseCode     *string `json:"AdviseCode,omitempty" xml:"AdviseCode,omitempty"`
	AffectCode     *string `json:"AffectCode,omitempty" xml:"AffectCode,omitempty"`
	CheckpointName *string `json:"CheckpointName,omitempty" xml:"CheckpointName,omitempty"`
	MessageCode    *string `json:"MessageCode,omitempty" xml:"MessageCode,omitempty"`
	MessageLevel   *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) SetAdviseCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults {
	s.AdviseCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) SetAffectCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults {
	s.AffectCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) SetCheckpointName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults {
	s.CheckpointName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) SetMessageCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults {
	s.MessageCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults) SetMessageLevel(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsUnknownInstanceResultCheckpointResults {
	s.MessageLevel = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult struct {
	CheckpointResults []*GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults `json:"CheckpointResults,omitempty" xml:"CheckpointResults,omitempty" type:"Repeated"`
	Extend            map[string]interface{}                                                                   `json:"Extend,omitempty" xml:"Extend,omitempty"`
	InstanceId        *string                                                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType      *string                                                                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) SetCheckpointResults(v []*GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult {
	s.CheckpointResults = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) SetExtend(v map[string]interface{}) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult {
	s.Extend = v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) SetInstanceId(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult {
	s.InstanceId = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) SetInstanceName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult {
	s.InstanceName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult) SetInstanceType(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResult {
	s.InstanceType = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults struct {
	AdviseCode     *string `json:"AdviseCode,omitempty" xml:"AdviseCode,omitempty"`
	AffectCode     *string `json:"AffectCode,omitempty" xml:"AffectCode,omitempty"`
	CheckpointName *string `json:"CheckpointName,omitempty" xml:"CheckpointName,omitempty"`
	MessageCode    *string `json:"MessageCode,omitempty" xml:"MessageCode,omitempty"`
	MessageLevel   *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) SetAdviseCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults {
	s.AdviseCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) SetAffectCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults {
	s.AffectCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) SetCheckpointName(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults {
	s.CheckpointName = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) SetMessageCode(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults {
	s.MessageCode = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults) SetMessageLevel(v string) *GetClusterCheckResultResponseBodyCheckEntryResultsWarnInstanceResultCheckpointResults {
	s.MessageLevel = &v
	return s
}

type GetClusterCheckResultResponseBodyCheckSummary struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCount   *int32  `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	NormalCount  *int32  `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	Process      *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	UnknownCount *int32  `json:"UnknownCount,omitempty" xml:"UnknownCount,omitempty"`
	WarnCount    *int32  `json:"WarnCount,omitempty" xml:"WarnCount,omitempty"`
}

func (s GetClusterCheckResultResponseBodyCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponseBodyCheckSummary) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponseBodyCheckSummary) SetCode(v string) *GetClusterCheckResultResponseBodyCheckSummary {
	s.Code = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckSummary) SetErrorCount(v int32) *GetClusterCheckResultResponseBodyCheckSummary {
	s.ErrorCount = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckSummary) SetNormalCount(v int32) *GetClusterCheckResultResponseBodyCheckSummary {
	s.NormalCount = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckSummary) SetProcess(v int32) *GetClusterCheckResultResponseBodyCheckSummary {
	s.Process = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckSummary) SetUnknownCount(v int32) *GetClusterCheckResultResponseBodyCheckSummary {
	s.UnknownCount = &v
	return s
}

func (s *GetClusterCheckResultResponseBodyCheckSummary) SetWarnCount(v int32) *GetClusterCheckResultResponseBodyCheckSummary {
	s.WarnCount = &v
	return s
}

type GetClusterCheckResultResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterCheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResultResponse) SetHeaders(v map[string]*string) *GetClusterCheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetClusterCheckResultResponse) SetBody(v *GetClusterCheckResultResponseBody) *GetClusterCheckResultResponse {
	s.Body = v
	return s
}

type GetClusterReportSummaryRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ReportUid       *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
}

func (s GetClusterReportSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterReportSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetClusterReportSummaryRequest) SetClusterRegionId(v string) *GetClusterReportSummaryRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *GetClusterReportSummaryRequest) SetOwnerUid(v int64) *GetClusterReportSummaryRequest {
	s.OwnerUid = &v
	return s
}

func (s *GetClusterReportSummaryRequest) SetReportUid(v string) *GetClusterReportSummaryRequest {
	s.ReportUid = &v
	return s
}

type GetClusterReportSummaryResponseBody struct {
	AddonName        *string                                                `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	CheckEndTime     *int64                                                 `json:"CheckEndTime,omitempty" xml:"CheckEndTime,omitempty"`
	CheckStartTime   *int64                                                 `json:"CheckStartTime,omitempty" xml:"CheckStartTime,omitempty"`
	CheckStatus      *string                                                `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckSummary     *GetClusterReportSummaryResponseBodyCheckSummary       `json:"CheckSummary,omitempty" xml:"CheckSummary,omitempty" type:"Struct"`
	ChecklistResults []*GetClusterReportSummaryResponseBodyChecklistResults `json:"ChecklistResults,omitempty" xml:"ChecklistResults,omitempty" type:"Repeated"`
	ClusterUid       *string                                                `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	Code             *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime       *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IsSuccess        *bool                                                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Name             *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uid              *string                                                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetClusterReportSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterReportSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterReportSummaryResponseBody) SetAddonName(v string) *GetClusterReportSummaryResponseBody {
	s.AddonName = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetCheckEndTime(v int64) *GetClusterReportSummaryResponseBody {
	s.CheckEndTime = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetCheckStartTime(v int64) *GetClusterReportSummaryResponseBody {
	s.CheckStartTime = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetCheckStatus(v string) *GetClusterReportSummaryResponseBody {
	s.CheckStatus = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetCheckSummary(v *GetClusterReportSummaryResponseBodyCheckSummary) *GetClusterReportSummaryResponseBody {
	s.CheckSummary = v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetChecklistResults(v []*GetClusterReportSummaryResponseBodyChecklistResults) *GetClusterReportSummaryResponseBody {
	s.ChecklistResults = v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetClusterUid(v string) *GetClusterReportSummaryResponseBody {
	s.ClusterUid = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetCode(v string) *GetClusterReportSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetCreateTime(v string) *GetClusterReportSummaryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetIsSuccess(v bool) *GetClusterReportSummaryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetName(v string) *GetClusterReportSummaryResponseBody {
	s.Name = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetRequestId(v string) *GetClusterReportSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterReportSummaryResponseBody) SetUid(v string) *GetClusterReportSummaryResponseBody {
	s.Uid = &v
	return s
}

type GetClusterReportSummaryResponseBodyCheckSummary struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Process *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s GetClusterReportSummaryResponseBodyCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s GetClusterReportSummaryResponseBodyCheckSummary) GoString() string {
	return s.String()
}

func (s *GetClusterReportSummaryResponseBodyCheckSummary) SetCode(v string) *GetClusterReportSummaryResponseBodyCheckSummary {
	s.Code = &v
	return s
}

func (s *GetClusterReportSummaryResponseBodyCheckSummary) SetProcess(v int32) *GetClusterReportSummaryResponseBodyCheckSummary {
	s.Process = &v
	return s
}

type GetClusterReportSummaryResponseBodyChecklistResults struct {
	CheckStatus  *string                                                          `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckSummary *GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary `json:"CheckSummary,omitempty" xml:"CheckSummary,omitempty" type:"Struct"`
	Name         *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetClusterReportSummaryResponseBodyChecklistResults) String() string {
	return tea.Prettify(s)
}

func (s GetClusterReportSummaryResponseBodyChecklistResults) GoString() string {
	return s.String()
}

func (s *GetClusterReportSummaryResponseBodyChecklistResults) SetCheckStatus(v string) *GetClusterReportSummaryResponseBodyChecklistResults {
	s.CheckStatus = &v
	return s
}

func (s *GetClusterReportSummaryResponseBodyChecklistResults) SetCheckSummary(v *GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary) *GetClusterReportSummaryResponseBodyChecklistResults {
	s.CheckSummary = v
	return s
}

func (s *GetClusterReportSummaryResponseBodyChecklistResults) SetName(v string) *GetClusterReportSummaryResponseBodyChecklistResults {
	s.Name = &v
	return s
}

type GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Process *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary) GoString() string {
	return s.String()
}

func (s *GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary) SetCode(v string) *GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary {
	s.Code = &v
	return s
}

func (s *GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary) SetProcess(v int32) *GetClusterReportSummaryResponseBodyChecklistResultsCheckSummary {
	s.Process = &v
	return s
}

type GetClusterReportSummaryResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterReportSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterReportSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterReportSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetClusterReportSummaryResponse) SetHeaders(v map[string]*string) *GetClusterReportSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetClusterReportSummaryResponse) SetBody(v *GetClusterReportSummaryResponseBody) *GetClusterReportSummaryResponse {
	s.Body = v
	return s
}

type GetDiagnosisCheckItemRequest struct {
	DiagnosisId *string `json:"DiagnosisId,omitempty" xml:"DiagnosisId,omitempty"`
	OwnerUid    *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
}

func (s GetDiagnosisCheckItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisCheckItemRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisCheckItemRequest) SetDiagnosisId(v string) *GetDiagnosisCheckItemRequest {
	s.DiagnosisId = &v
	return s
}

func (s *GetDiagnosisCheckItemRequest) SetOwnerUid(v int64) *GetDiagnosisCheckItemRequest {
	s.OwnerUid = &v
	return s
}

type GetDiagnosisCheckItemResponseBody struct {
	CheckItems []*GetDiagnosisCheckItemResponseBodyCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
	Code       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess  *bool                                          `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiagnosisCheckItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisCheckItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosisCheckItemResponseBody) SetCheckItems(v []*GetDiagnosisCheckItemResponseBodyCheckItems) *GetDiagnosisCheckItemResponseBody {
	s.CheckItems = v
	return s
}

func (s *GetDiagnosisCheckItemResponseBody) SetCode(v string) *GetDiagnosisCheckItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBody) SetIsSuccess(v bool) *GetDiagnosisCheckItemResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBody) SetRequestId(v string) *GetDiagnosisCheckItemResponseBody {
	s.RequestId = &v
	return s
}

type GetDiagnosisCheckItemResponseBodyCheckItems struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Level   *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Refer   *string `json:"Refer,omitempty" xml:"Refer,omitempty"`
	Value   *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDiagnosisCheckItemResponseBodyCheckItems) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisCheckItemResponseBodyCheckItems) GoString() string {
	return s.String()
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetDesc(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Desc = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetGroup(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Group = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetLevel(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Level = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetMessage(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Message = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetName(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Name = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetRefer(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Refer = &v
	return s
}

func (s *GetDiagnosisCheckItemResponseBodyCheckItems) SetValue(v string) *GetDiagnosisCheckItemResponseBodyCheckItems {
	s.Value = &v
	return s
}

type GetDiagnosisCheckItemResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDiagnosisCheckItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiagnosisCheckItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisCheckItemResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisCheckItemResponse) SetHeaders(v map[string]*string) *GetDiagnosisCheckItemResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisCheckItemResponse) SetBody(v *GetDiagnosisCheckItemResponseBody) *GetDiagnosisCheckItemResponse {
	s.Body = v
	return s
}

type GetDiagnosisResultRequest struct {
	DiagnosisId *string `json:"DiagnosisId,omitempty" xml:"DiagnosisId,omitempty"`
	OwnerUid    *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) SetDiagnosisId(v string) *GetDiagnosisResultRequest {
	s.DiagnosisId = &v
	return s
}

func (s *GetDiagnosisResultRequest) SetOwnerUid(v int64) *GetDiagnosisResultRequest {
	s.OwnerUid = &v
	return s
}

type GetDiagnosisResultResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Diagnosis *GetDiagnosisResultResponseBodyDiagnosis `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty" type:"Struct"`
	// bool of success
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiagnosisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBody) SetCode(v string) *GetDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetDiagnosis(v *GetDiagnosisResultResponseBodyDiagnosis) *GetDiagnosisResultResponseBody {
	s.Diagnosis = v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetIsSuccess(v bool) *GetDiagnosisResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetRequestId(v string) *GetDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDiagnosisResultResponseBodyDiagnosis struct {
	ClusterUid  *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	Code        *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Created     *string `json:"Created,omitempty" xml:"Created,omitempty"`
	DiagnosisId *string `json:"DiagnosisId,omitempty" xml:"DiagnosisId,omitempty"`
	Finished    *string `json:"Finished,omitempty" xml:"Finished,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status      *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Target      *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDiagnosisResultResponseBodyDiagnosis) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBodyDiagnosis) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetClusterUid(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.ClusterUid = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetCode(v int64) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetCreated(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Created = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetDiagnosisId(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.DiagnosisId = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetFinished(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Finished = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetMessage(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Message = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetResult(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Result = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetStatus(v int64) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Status = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetTarget(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Target = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyDiagnosis) SetType(v string) *GetDiagnosisResultResponseBodyDiagnosis {
	s.Type = &v
	return s
}

type GetDiagnosisResultResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiagnosisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisResultResponse) SetBody(v *GetDiagnosisResultResponseBody) *GetDiagnosisResultResponse {
	s.Body = v
	return s
}

type ListClusterReportSummaryRequest struct {
	AddonName       *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	PageNo          *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReportName      *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
}

func (s ListClusterReportSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryRequest) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryRequest) SetAddonName(v string) *ListClusterReportSummaryRequest {
	s.AddonName = &v
	return s
}

func (s *ListClusterReportSummaryRequest) SetClusterRegionId(v string) *ListClusterReportSummaryRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *ListClusterReportSummaryRequest) SetClusterUid(v string) *ListClusterReportSummaryRequest {
	s.ClusterUid = &v
	return s
}

func (s *ListClusterReportSummaryRequest) SetOwnerUid(v int64) *ListClusterReportSummaryRequest {
	s.OwnerUid = &v
	return s
}

func (s *ListClusterReportSummaryRequest) SetPageNo(v string) *ListClusterReportSummaryRequest {
	s.PageNo = &v
	return s
}

func (s *ListClusterReportSummaryRequest) SetPageSize(v string) *ListClusterReportSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterReportSummaryRequest) SetReportName(v string) *ListClusterReportSummaryRequest {
	s.ReportName = &v
	return s
}

type ListClusterReportSummaryResponseBody struct {
	Code            *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess       *bool                                                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo          *int32                                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReportSummaries []*ListClusterReportSummaryResponseBodyReportSummaries `json:"ReportSummaries,omitempty" xml:"ReportSummaries,omitempty" type:"Repeated"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterReportSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryResponseBody) SetCode(v string) *ListClusterReportSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterReportSummaryResponseBody) SetIsSuccess(v bool) *ListClusterReportSummaryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListClusterReportSummaryResponseBody) SetPageNo(v int32) *ListClusterReportSummaryResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListClusterReportSummaryResponseBody) SetPageSize(v int32) *ListClusterReportSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterReportSummaryResponseBody) SetReportSummaries(v []*ListClusterReportSummaryResponseBodyReportSummaries) *ListClusterReportSummaryResponseBody {
	s.ReportSummaries = v
	return s
}

func (s *ListClusterReportSummaryResponseBody) SetRequestId(v string) *ListClusterReportSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterReportSummaryResponseBody) SetTotalCount(v int32) *ListClusterReportSummaryResponseBody {
	s.TotalCount = &v
	return s
}

type ListClusterReportSummaryResponseBodyReportSummaries struct {
	CheckEndTime     *int64                                                                 `json:"CheckEndTime,omitempty" xml:"CheckEndTime,omitempty"`
	CheckStartTime   *int64                                                                 `json:"CheckStartTime,omitempty" xml:"CheckStartTime,omitempty"`
	CheckStatus      *string                                                                `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckSummary     *ListClusterReportSummaryResponseBodyReportSummariesCheckSummary       `json:"CheckSummary,omitempty" xml:"CheckSummary,omitempty" type:"Struct"`
	ChecklistResults []*ListClusterReportSummaryResponseBodyReportSummariesChecklistResults `json:"ChecklistResults,omitempty" xml:"ChecklistResults,omitempty" type:"Repeated"`
	ClusterUid       *string                                                                `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	CreateTime       *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name             *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Uid              *string                                                                `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListClusterReportSummaryResponseBodyReportSummaries) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryResponseBodyReportSummaries) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetCheckEndTime(v int64) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.CheckEndTime = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetCheckStartTime(v int64) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.CheckStartTime = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetCheckStatus(v string) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.CheckStatus = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetCheckSummary(v *ListClusterReportSummaryResponseBodyReportSummariesCheckSummary) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.CheckSummary = v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetChecklistResults(v []*ListClusterReportSummaryResponseBodyReportSummariesChecklistResults) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.ChecklistResults = v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetClusterUid(v string) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.ClusterUid = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetCreateTime(v string) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.CreateTime = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetName(v string) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.Name = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummaries) SetUid(v string) *ListClusterReportSummaryResponseBodyReportSummaries {
	s.Uid = &v
	return s
}

type ListClusterReportSummaryResponseBodyReportSummariesCheckSummary struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Process *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s ListClusterReportSummaryResponseBodyReportSummariesCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryResponseBodyReportSummariesCheckSummary) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesCheckSummary) SetCode(v string) *ListClusterReportSummaryResponseBodyReportSummariesCheckSummary {
	s.Code = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesCheckSummary) SetProcess(v int32) *ListClusterReportSummaryResponseBodyReportSummariesCheckSummary {
	s.Process = &v
	return s
}

type ListClusterReportSummaryResponseBodyReportSummariesChecklistResults struct {
	CheckStatus  *string                                                                          `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckSummary *ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary `json:"CheckSummary,omitempty" xml:"CheckSummary,omitempty" type:"Struct"`
	Name         *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListClusterReportSummaryResponseBodyReportSummariesChecklistResults) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryResponseBodyReportSummariesChecklistResults) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesChecklistResults) SetCheckStatus(v string) *ListClusterReportSummaryResponseBodyReportSummariesChecklistResults {
	s.CheckStatus = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesChecklistResults) SetCheckSummary(v *ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary) *ListClusterReportSummaryResponseBodyReportSummariesChecklistResults {
	s.CheckSummary = v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesChecklistResults) SetName(v string) *ListClusterReportSummaryResponseBodyReportSummariesChecklistResults {
	s.Name = &v
	return s
}

type ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Process *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary) SetCode(v string) *ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary {
	s.Code = &v
	return s
}

func (s *ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary) SetProcess(v int32) *ListClusterReportSummaryResponseBodyReportSummariesChecklistResultsCheckSummary {
	s.Process = &v
	return s
}

type ListClusterReportSummaryResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterReportSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterReportSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterReportSummaryResponse) GoString() string {
	return s.String()
}

func (s *ListClusterReportSummaryResponse) SetHeaders(v map[string]*string) *ListClusterReportSummaryResponse {
	s.Headers = v
	return s
}

func (s *ListClusterReportSummaryResponse) SetBody(v *ListClusterReportSummaryResponseBody) *ListClusterReportSummaryResponse {
	s.Body = v
	return s
}

type ListDiagnosisResultRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	OwnerUid        *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	PageNo          *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Since           *int64  `json:"Since,omitempty" xml:"Since,omitempty"`
}

func (s ListDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResultRequest) SetClusterRegionId(v string) *ListDiagnosisResultRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *ListDiagnosisResultRequest) SetClusterUid(v string) *ListDiagnosisResultRequest {
	s.ClusterUid = &v
	return s
}

func (s *ListDiagnosisResultRequest) SetOwnerUid(v int64) *ListDiagnosisResultRequest {
	s.OwnerUid = &v
	return s
}

func (s *ListDiagnosisResultRequest) SetPageNo(v int64) *ListDiagnosisResultRequest {
	s.PageNo = &v
	return s
}

func (s *ListDiagnosisResultRequest) SetPageSize(v int64) *ListDiagnosisResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListDiagnosisResultRequest) SetSince(v int64) *ListDiagnosisResultRequest {
	s.Since = &v
	return s
}

type ListDiagnosisResultResponseBody struct {
	Code             *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DiagnosisResults []*ListDiagnosisResultResponseBodyDiagnosisResults `json:"DiagnosisResults,omitempty" xml:"DiagnosisResults,omitempty" type:"Repeated"`
	IsSuccess        *bool                                              `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo           *int64                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int64                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDiagnosisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResultResponseBody) SetCode(v string) *ListDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResultResponseBody) SetDiagnosisResults(v []*ListDiagnosisResultResponseBodyDiagnosisResults) *ListDiagnosisResultResponseBody {
	s.DiagnosisResults = v
	return s
}

func (s *ListDiagnosisResultResponseBody) SetIsSuccess(v bool) *ListDiagnosisResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListDiagnosisResultResponseBody) SetPageNo(v int64) *ListDiagnosisResultResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListDiagnosisResultResponseBody) SetPageSize(v int64) *ListDiagnosisResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDiagnosisResultResponseBody) SetRequestId(v string) *ListDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosisResultResponseBody) SetTotalCount(v int32) *ListDiagnosisResultResponseBody {
	s.TotalCount = &v
	return s
}

type ListDiagnosisResultResponseBodyDiagnosisResults struct {
	Code        *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Created     *int64  `json:"Created,omitempty" xml:"Created,omitempty"`
	DiagnosisId *string `json:"DiagnosisId,omitempty" xml:"DiagnosisId,omitempty"`
	Finished    *int64  `json:"Finished,omitempty" xml:"Finished,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status      *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Target      *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDiagnosisResultResponseBodyDiagnosisResults) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResultResponseBodyDiagnosisResults) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetCode(v int64) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetCreated(v int64) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Created = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetDiagnosisId(v string) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.DiagnosisId = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetFinished(v int64) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Finished = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetMessage(v string) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Message = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetStatus(v int64) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Status = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetTarget(v string) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Target = &v
	return s
}

func (s *ListDiagnosisResultResponseBodyDiagnosisResults) SetType(v string) *ListDiagnosisResultResponseBodyDiagnosisResults {
	s.Type = &v
	return s
}

type ListDiagnosisResultResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiagnosisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResultResponse) SetHeaders(v map[string]*string) *ListDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosisResultResponse) SetBody(v *ListDiagnosisResultResponseBody) *ListDiagnosisResultResponse {
	s.Body = v
	return s
}

type ListReportTaskRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	ReportName      *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
}

func (s ListReportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskRequest) GoString() string {
	return s.String()
}

func (s *ListReportTaskRequest) SetClusterRegionId(v string) *ListReportTaskRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *ListReportTaskRequest) SetClusterUid(v string) *ListReportTaskRequest {
	s.ClusterUid = &v
	return s
}

func (s *ListReportTaskRequest) SetReportName(v string) *ListReportTaskRequest {
	s.ReportName = &v
	return s
}

type ListReportTaskResponseBody struct {
	Code        *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	PageNo      *int64                                   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int64                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReportTasks []*ListReportTaskResponseBodyReportTasks `json:"ReportTasks,omitempty" xml:"ReportTasks,omitempty" type:"Repeated"`
	RequestId   map[string]interface{}                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListReportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportTaskResponseBody) SetCode(v string) *ListReportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListReportTaskResponseBody) SetPageNo(v int64) *ListReportTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListReportTaskResponseBody) SetPageSize(v int64) *ListReportTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListReportTaskResponseBody) SetReportTasks(v []*ListReportTaskResponseBodyReportTasks) *ListReportTaskResponseBody {
	s.ReportTasks = v
	return s
}

func (s *ListReportTaskResponseBody) SetRequestId(v map[string]interface{}) *ListReportTaskResponseBody {
	s.RequestId = v
	return s
}

func (s *ListReportTaskResponseBody) SetSuccess(v bool) *ListReportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListReportTaskResponseBody) SetTotalCount(v int64) *ListReportTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListReportTaskResponseBodyReportTasks struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReportUid *string `json:"ReportUid,omitempty" xml:"ReportUid,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListReportTaskResponseBodyReportTasks) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskResponseBodyReportTasks) GoString() string {
	return s.String()
}

func (s *ListReportTaskResponseBodyReportTasks) SetEndTime(v int64) *ListReportTaskResponseBodyReportTasks {
	s.EndTime = &v
	return s
}

func (s *ListReportTaskResponseBodyReportTasks) SetReportUid(v string) *ListReportTaskResponseBodyReportTasks {
	s.ReportUid = &v
	return s
}

func (s *ListReportTaskResponseBodyReportTasks) SetResult(v string) *ListReportTaskResponseBodyReportTasks {
	s.Result = &v
	return s
}

func (s *ListReportTaskResponseBodyReportTasks) SetStartTime(v int64) *ListReportTaskResponseBodyReportTasks {
	s.StartTime = &v
	return s
}

func (s *ListReportTaskResponseBodyReportTasks) SetStatus(v string) *ListReportTaskResponseBodyReportTasks {
	s.Status = &v
	return s
}

func (s *ListReportTaskResponseBodyReportTasks) SetTaskId(v string) *ListReportTaskResponseBodyReportTasks {
	s.TaskId = &v
	return s
}

type ListReportTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListReportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListReportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskResponse) GoString() string {
	return s.String()
}

func (s *ListReportTaskResponse) SetHeaders(v map[string]*string) *ListReportTaskResponse {
	s.Headers = v
	return s
}

func (s *ListReportTaskResponse) SetBody(v *ListReportTaskResponseBody) *ListReportTaskResponse {
	s.Body = v
	return s
}

type ListReportTaskRuleRequest struct {
	ClusterRegionId *string `json:"ClusterRegionId,omitempty" xml:"ClusterRegionId,omitempty"`
	ClusterUid      *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
}

func (s ListReportTaskRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskRuleRequest) GoString() string {
	return s.String()
}

func (s *ListReportTaskRuleRequest) SetClusterRegionId(v string) *ListReportTaskRuleRequest {
	s.ClusterRegionId = &v
	return s
}

func (s *ListReportTaskRuleRequest) SetClusterUid(v string) *ListReportTaskRuleRequest {
	s.ClusterUid = &v
	return s
}

type ListReportTaskRuleResponseBody struct {
	Code            *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess       *bool                                            `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReportTaskRules []*ListReportTaskRuleResponseBodyReportTaskRules `json:"ReportTaskRules,omitempty" xml:"ReportTaskRules,omitempty" type:"Repeated"`
	RequestId       map[string]interface{}                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListReportTaskRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportTaskRuleResponseBody) SetCode(v string) *ListReportTaskRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListReportTaskRuleResponseBody) SetIsSuccess(v bool) *ListReportTaskRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListReportTaskRuleResponseBody) SetReportTaskRules(v []*ListReportTaskRuleResponseBodyReportTaskRules) *ListReportTaskRuleResponseBody {
	s.ReportTaskRules = v
	return s
}

func (s *ListReportTaskRuleResponseBody) SetRequestId(v map[string]interface{}) *ListReportTaskRuleResponseBody {
	s.RequestId = v
	return s
}

type ListReportTaskRuleResponseBodyReportTaskRules struct {
	ClusterUid   *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	ReportName   *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ScheduleRule *string `json:"ScheduleRule,omitempty" xml:"ScheduleRule,omitempty"`
	TimeZoneId   *string `json:"TimeZoneId,omitempty" xml:"TimeZoneId,omitempty"`
}

func (s ListReportTaskRuleResponseBodyReportTaskRules) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskRuleResponseBodyReportTaskRules) GoString() string {
	return s.String()
}

func (s *ListReportTaskRuleResponseBodyReportTaskRules) SetClusterUid(v string) *ListReportTaskRuleResponseBodyReportTaskRules {
	s.ClusterUid = &v
	return s
}

func (s *ListReportTaskRuleResponseBodyReportTaskRules) SetReportName(v string) *ListReportTaskRuleResponseBodyReportTaskRules {
	s.ReportName = &v
	return s
}

func (s *ListReportTaskRuleResponseBodyReportTaskRules) SetRuleId(v string) *ListReportTaskRuleResponseBodyReportTaskRules {
	s.RuleId = &v
	return s
}

func (s *ListReportTaskRuleResponseBodyReportTaskRules) SetScheduleRule(v string) *ListReportTaskRuleResponseBodyReportTaskRules {
	s.ScheduleRule = &v
	return s
}

func (s *ListReportTaskRuleResponseBodyReportTaskRules) SetTimeZoneId(v string) *ListReportTaskRuleResponseBodyReportTaskRules {
	s.TimeZoneId = &v
	return s
}

type ListReportTaskRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListReportTaskRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListReportTaskRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReportTaskRuleResponse) GoString() string {
	return s.String()
}

func (s *ListReportTaskRuleResponse) SetHeaders(v map[string]*string) *ListReportTaskRuleResponse {
	s.Headers = v
	return s
}

func (s *ListReportTaskRuleResponse) SetBody(v *ListReportTaskRuleResponseBody) *ListReportTaskRuleResponse {
	s.Body = v
	return s
}

type CheckResultNodesLevelValue struct {
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Targets []*CheckResultNodesLevelValueTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s CheckResultNodesLevelValue) String() string {
	return tea.Prettify(s)
}

func (s CheckResultNodesLevelValue) GoString() string {
	return s.String()
}

func (s *CheckResultNodesLevelValue) SetMessage(v string) *CheckResultNodesLevelValue {
	s.Message = &v
	return s
}

func (s *CheckResultNodesLevelValue) SetTargets(v []*CheckResultNodesLevelValueTargets) *CheckResultNodesLevelValue {
	s.Targets = v
	return s
}

type CheckResultNodesLevelValueTargets struct {
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CheckResultNodesLevelValueTargets) String() string {
	return tea.Prettify(s)
}

func (s CheckResultNodesLevelValueTargets) GoString() string {
	return s.String()
}

func (s *CheckResultNodesLevelValueTargets) SetTarget(v string) *CheckResultNodesLevelValueTargets {
	s.Target = &v
	return s
}

func (s *CheckResultNodesLevelValueTargets) SetValue(v string) *CheckResultNodesLevelValueTargets {
	s.Value = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cis"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateClusterCheckWithOptions(request *CreateClusterCheckRequest, runtime *util.RuntimeOptions) (_result *CreateClusterCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonName)) {
		query["AddonName"] = request.AddonName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportName)) {
		query["ReportName"] = request.ReportName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClusterCheck"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterCheck(request *CreateClusterCheckRequest) (_result *CreateClusterCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterCheckResponse{}
	_body, _err := client.CreateClusterCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterReportWithOptions(request *CreateClusterReportRequest, runtime *util.RuntimeOptions) (_result *CreateClusterReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonName)) {
		query["AddonName"] = request.AddonName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportName)) {
		query["ReportName"] = request.ReportName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClusterReport"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterReport(request *CreateClusterReportRequest) (_result *CreateClusterReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterReportResponse{}
	_body, _err := client.CreateClusterReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDiagnosisWithOptions(request *CreateDiagnosisRequest, runtime *util.RuntimeOptions) (_result *CreateDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiagnosis"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiagnosis(request *CreateDiagnosisRequest) (_result *CreateDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiagnosisResponse{}
	_body, _err := client.CreateDiagnosisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateReportTaskRuleWithOptions(request *CreateReportTaskRuleRequest, runtime *util.RuntimeOptions) (_result *CreateReportTaskRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportName)) {
		query["ReportName"] = request.ReportName
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleRule)) {
		query["ScheduleRule"] = request.ScheduleRule
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZoneId)) {
		query["TimeZoneId"] = request.TimeZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReportTaskRule"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReportTaskRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReportTaskRule(request *CreateReportTaskRuleRequest) (_result *CreateReportTaskRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReportTaskRuleResponse{}
	_body, _err := client.CreateReportTaskRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteReportTaskRuleWithOptions(request *DeleteReportTaskRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteReportTaskRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteReportTaskRule"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteReportTaskRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteReportTaskRule(request *DeleteReportTaskRuleRequest) (_result *DeleteReportTaskRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteReportTaskRuleResponse{}
	_body, _err := client.DeleteReportTaskRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterCheckItemWithOptions(request *GetClusterCheckItemRequest, runtime *util.RuntimeOptions) (_result *GetClusterCheckItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterCheckItem"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterCheckItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterCheckItem(request *GetClusterCheckItemRequest) (_result *GetClusterCheckItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterCheckItemResponse{}
	_body, _err := client.GetClusterCheckItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterCheckResultWithOptions(request *GetClusterCheckResultRequest, runtime *util.RuntimeOptions) (_result *GetClusterCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChecklistName)) {
		query["ChecklistName"] = request.ChecklistName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportUid)) {
		query["ReportUid"] = request.ReportUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterCheckResult"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterCheckResult(request *GetClusterCheckResultRequest) (_result *GetClusterCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterCheckResultResponse{}
	_body, _err := client.GetClusterCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterReportSummaryWithOptions(request *GetClusterReportSummaryRequest, runtime *util.RuntimeOptions) (_result *GetClusterReportSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportUid)) {
		query["ReportUid"] = request.ReportUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterReportSummary"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterReportSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterReportSummary(request *GetClusterReportSummaryRequest) (_result *GetClusterReportSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterReportSummaryResponse{}
	_body, _err := client.GetClusterReportSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiagnosisCheckItemWithOptions(request *GetDiagnosisCheckItemRequest, runtime *util.RuntimeOptions) (_result *GetDiagnosisCheckItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisCheckItem"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosisCheckItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiagnosisCheckItem(request *GetDiagnosisCheckItemRequest) (_result *GetDiagnosisCheckItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiagnosisCheckItemResponse{}
	_body, _err := client.GetDiagnosisCheckItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiagnosisResultWithOptions(request *GetDiagnosisResultRequest, runtime *util.RuntimeOptions) (_result *GetDiagnosisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisResult"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiagnosisResult(request *GetDiagnosisResultRequest) (_result *GetDiagnosisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.GetDiagnosisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterReportSummaryWithOptions(request *ListClusterReportSummaryRequest, runtime *util.RuntimeOptions) (_result *ListClusterReportSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonName)) {
		query["AddonName"] = request.AddonName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReportName)) {
		query["ReportName"] = request.ReportName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterReportSummary"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterReportSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterReportSummary(request *ListClusterReportSummaryRequest) (_result *ListClusterReportSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterReportSummaryResponse{}
	_body, _err := client.ListClusterReportSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiagnosisResultWithOptions(request *ListDiagnosisResultRequest, runtime *util.RuntimeOptions) (_result *ListDiagnosisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnosisResult"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiagnosisResult(request *ListDiagnosisResultRequest) (_result *ListDiagnosisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiagnosisResultResponse{}
	_body, _err := client.ListDiagnosisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListReportTaskWithOptions(request *ListReportTaskRequest, runtime *util.RuntimeOptions) (_result *ListReportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportName)) {
		query["ReportName"] = request.ReportName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReportTask"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListReportTask(request *ListReportTaskRequest) (_result *ListReportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListReportTaskResponse{}
	_body, _err := client.ListReportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListReportTaskRuleWithOptions(request *ListReportTaskRuleRequest, runtime *util.RuntimeOptions) (_result *ListReportTaskRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterRegionId)) {
		query["ClusterRegionId"] = request.ClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUid)) {
		query["ClusterUid"] = request.ClusterUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReportTaskRule"),
		Version:     tea.String("2018-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReportTaskRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListReportTaskRule(request *ListReportTaskRuleRequest) (_result *ListReportTaskRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListReportTaskRuleResponse{}
	_body, _err := client.ListReportTaskRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
