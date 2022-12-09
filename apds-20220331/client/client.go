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

type CreateFileJobRequest struct {
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SurveyJobId  *int64  `json:"surveyJobId,omitempty" xml:"surveyJobId,omitempty"`
	RegionId     *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateFileJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFileJobRequest) SetResourceType(v string) *CreateFileJobRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateFileJobRequest) SetSurveyJobId(v int64) *CreateFileJobRequest {
	s.SurveyJobId = &v
	return s
}

func (s *CreateFileJobRequest) SetRegionId(v string) *CreateFileJobRequest {
	s.RegionId = &v
	return s
}

type CreateFileJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s CreateFileJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileJobResponseBody) SetCode(v string) *CreateFileJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFileJobResponseBody) SetData(v interface{}) *CreateFileJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateFileJobResponseBody) SetSuccess(v bool) *CreateFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFileJobResponseBody) SetError(v string) *CreateFileJobResponseBody {
	s.Error = &v
	return s
}

type CreateFileJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFileJobResponse) SetHeaders(v map[string]*string) *CreateFileJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFileJobResponse) SetStatusCode(v int32) *CreateFileJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileJobResponse) SetBody(v *CreateFileJobResponseBody) *CreateFileJobResponse {
	s.Body = v
	return s
}

type CreateMigrationGroupRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Extra       *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Id          *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateMigrationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrationGroupRequest) SetDescription(v string) *CreateMigrationGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateMigrationGroupRequest) SetExtra(v string) *CreateMigrationGroupRequest {
	s.Extra = &v
	return s
}

func (s *CreateMigrationGroupRequest) SetId(v int32) *CreateMigrationGroupRequest {
	s.Id = &v
	return s
}

func (s *CreateMigrationGroupRequest) SetName(v string) *CreateMigrationGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateMigrationGroupRequest) SetRegionId(v string) *CreateMigrationGroupRequest {
	s.RegionId = &v
	return s
}

type CreateMigrationGroupResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMigrationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrationGroupResponseBody) SetCode(v string) *CreateMigrationGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMigrationGroupResponseBody) SetData(v interface{}) *CreateMigrationGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateMigrationGroupResponseBody) SetError(v string) *CreateMigrationGroupResponseBody {
	s.Error = &v
	return s
}

func (s *CreateMigrationGroupResponseBody) SetSuccess(v bool) *CreateMigrationGroupResponseBody {
	s.Success = &v
	return s
}

type CreateMigrationGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMigrationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMigrationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMigrationGroupResponse) SetHeaders(v map[string]*string) *CreateMigrationGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMigrationGroupResponse) SetStatusCode(v int32) *CreateMigrationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMigrationGroupResponse) SetBody(v *CreateMigrationGroupResponseBody) *CreateMigrationGroupResponse {
	s.Body = v
	return s
}

type CreateMigrationJobRequest struct {
	MigrationJobList []*CreateMigrationJobRequestMigrationJobList `json:"migrationJobList,omitempty" xml:"migrationJobList,omitempty" type:"Repeated"`
	Type             *string                                      `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobRequest) SetMigrationJobList(v []*CreateMigrationJobRequestMigrationJobList) *CreateMigrationJobRequest {
	s.MigrationJobList = v
	return s
}

func (s *CreateMigrationJobRequest) SetType(v string) *CreateMigrationJobRequest {
	s.Type = &v
	return s
}

type CreateMigrationJobRequestMigrationJobList struct {
	Destination       *string `json:"destination,omitempty" xml:"destination,omitempty"`
	DestinationIp     *string `json:"destinationIp,omitempty" xml:"destinationIp,omitempty"`
	DestinationRegion *string `json:"destinationRegion,omitempty" xml:"destinationRegion,omitempty"`
	JobGmtCreate      *string `json:"jobGmtCreate,omitempty" xml:"jobGmtCreate,omitempty"`
	JobGmtModified    *string `json:"jobGmtModified,omitempty" xml:"jobGmtModified,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	OriginalPercent   *string `json:"originalPercent,omitempty" xml:"originalPercent,omitempty"`
	OriginalProgress  *string `json:"originalProgress,omitempty" xml:"originalProgress,omitempty"`
	OriginalStatus    *string `json:"originalStatus,omitempty" xml:"originalStatus,omitempty"`
	OutSideId         *string `json:"outSideId,omitempty" xml:"outSideId,omitempty"`
	Properties        *string `json:"properties,omitempty" xml:"properties,omitempty"`
	Source            *string `json:"source,omitempty" xml:"source,omitempty"`
	SourceIp          *string `json:"sourceIp,omitempty" xml:"sourceIp,omitempty"`
}

func (s CreateMigrationJobRequestMigrationJobList) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobRequestMigrationJobList) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobRequestMigrationJobList) SetDestination(v string) *CreateMigrationJobRequestMigrationJobList {
	s.Destination = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetDestinationIp(v string) *CreateMigrationJobRequestMigrationJobList {
	s.DestinationIp = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetDestinationRegion(v string) *CreateMigrationJobRequestMigrationJobList {
	s.DestinationRegion = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetJobGmtCreate(v string) *CreateMigrationJobRequestMigrationJobList {
	s.JobGmtCreate = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetJobGmtModified(v string) *CreateMigrationJobRequestMigrationJobList {
	s.JobGmtModified = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetName(v string) *CreateMigrationJobRequestMigrationJobList {
	s.Name = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetOriginalPercent(v string) *CreateMigrationJobRequestMigrationJobList {
	s.OriginalPercent = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetOriginalProgress(v string) *CreateMigrationJobRequestMigrationJobList {
	s.OriginalProgress = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetOriginalStatus(v string) *CreateMigrationJobRequestMigrationJobList {
	s.OriginalStatus = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetOutSideId(v string) *CreateMigrationJobRequestMigrationJobList {
	s.OutSideId = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetProperties(v string) *CreateMigrationJobRequestMigrationJobList {
	s.Properties = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetSource(v string) *CreateMigrationJobRequestMigrationJobList {
	s.Source = &v
	return s
}

func (s *CreateMigrationJobRequestMigrationJobList) SetSourceIp(v string) *CreateMigrationJobRequestMigrationJobList {
	s.SourceIp = &v
	return s
}

type CreateMigrationJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponseBody) SetCode(v string) *CreateMigrationJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetData(v interface{}) *CreateMigrationJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateMigrationJobResponseBody) SetError(v string) *CreateMigrationJobResponseBody {
	s.Error = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetSuccess(v bool) *CreateMigrationJobResponseBody {
	s.Success = &v
	return s
}

type CreateMigrationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponse) SetHeaders(v map[string]*string) *CreateMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMigrationJobResponse) SetStatusCode(v int32) *CreateMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMigrationJobResponse) SetBody(v *CreateMigrationJobResponseBody) *CreateMigrationJobResponse {
	s.Body = v
	return s
}

type CreatePayOrderCallbackRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreatePayOrderCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePayOrderCallbackRequest) GoString() string {
	return s.String()
}

func (s *CreatePayOrderCallbackRequest) SetData(v string) *CreatePayOrderCallbackRequest {
	s.Data = &v
	return s
}

type CreatePayOrderCallbackResponseBody struct {
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObj interface{} `json:"ResultObj,omitempty" xml:"ResultObj,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *bool       `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreatePayOrderCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePayOrderCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePayOrderCallbackResponseBody) SetCode(v string) *CreatePayOrderCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePayOrderCallbackResponseBody) SetMessage(v string) *CreatePayOrderCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePayOrderCallbackResponseBody) SetRequestId(v string) *CreatePayOrderCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePayOrderCallbackResponseBody) SetResultObj(v interface{}) *CreatePayOrderCallbackResponseBody {
	s.ResultObj = v
	return s
}

func (s *CreatePayOrderCallbackResponseBody) SetSuccess(v bool) *CreatePayOrderCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePayOrderCallbackResponseBody) SetSynchro(v bool) *CreatePayOrderCallbackResponseBody {
	s.Synchro = &v
	return s
}

type CreatePayOrderCallbackResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePayOrderCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePayOrderCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePayOrderCallbackResponse) GoString() string {
	return s.String()
}

func (s *CreatePayOrderCallbackResponse) SetHeaders(v map[string]*string) *CreatePayOrderCallbackResponse {
	s.Headers = v
	return s
}

func (s *CreatePayOrderCallbackResponse) SetStatusCode(v int32) *CreatePayOrderCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePayOrderCallbackResponse) SetBody(v *CreatePayOrderCallbackResponseBody) *CreatePayOrderCallbackResponse {
	s.Body = v
	return s
}

type CreateRefundRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundRequest) GoString() string {
	return s.String()
}

func (s *CreateRefundRequest) SetData(v string) *CreateRefundRequest {
	s.Data = &v
	return s
}

type CreateRefundResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateRefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRefundResponseBody) SetCode(v string) *CreateRefundResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRefundResponseBody) SetData(v string) *CreateRefundResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRefundResponseBody) SetMessage(v string) *CreateRefundResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRefundResponseBody) SetRequestId(v string) *CreateRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRefundResponseBody) SetSuccess(v bool) *CreateRefundResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRefundResponseBody) SetSynchro(v bool) *CreateRefundResponseBody {
	s.Synchro = &v
	return s
}

type CreateRefundResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRefundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundResponse) GoString() string {
	return s.String()
}

func (s *CreateRefundResponse) SetHeaders(v map[string]*string) *CreateRefundResponse {
	s.Headers = v
	return s
}

func (s *CreateRefundResponse) SetStatusCode(v int32) *CreateRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRefundResponse) SetBody(v *CreateRefundResponseBody) *CreateRefundResponse {
	s.Body = v
	return s
}

type CreateSurveyJobRequest struct {
	Ak               *string   `json:"ak,omitempty" xml:"ak,omitempty"`
	Channel          *string   `json:"channel,omitempty" xml:"channel,omitempty"`
	CloudType        *string   `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	Name             *string   `json:"name,omitempty" xml:"name,omitempty"`
	Region           []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
	ResourceTypeList []*string `json:"resourceTypeList,omitempty" xml:"resourceTypeList,omitempty" type:"Repeated"`
	Sk               *string   `json:"sk,omitempty" xml:"sk,omitempty"`
	TenantId         *string   `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Zone             []*string `json:"zone,omitempty" xml:"zone,omitempty" type:"Repeated"`
	RegionId         *string   `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateSurveyJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSurveyJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSurveyJobRequest) SetAk(v string) *CreateSurveyJobRequest {
	s.Ak = &v
	return s
}

func (s *CreateSurveyJobRequest) SetChannel(v string) *CreateSurveyJobRequest {
	s.Channel = &v
	return s
}

func (s *CreateSurveyJobRequest) SetCloudType(v string) *CreateSurveyJobRequest {
	s.CloudType = &v
	return s
}

func (s *CreateSurveyJobRequest) SetName(v string) *CreateSurveyJobRequest {
	s.Name = &v
	return s
}

func (s *CreateSurveyJobRequest) SetRegion(v []*string) *CreateSurveyJobRequest {
	s.Region = v
	return s
}

func (s *CreateSurveyJobRequest) SetResourceTypeList(v []*string) *CreateSurveyJobRequest {
	s.ResourceTypeList = v
	return s
}

func (s *CreateSurveyJobRequest) SetSk(v string) *CreateSurveyJobRequest {
	s.Sk = &v
	return s
}

func (s *CreateSurveyJobRequest) SetTenantId(v string) *CreateSurveyJobRequest {
	s.TenantId = &v
	return s
}

func (s *CreateSurveyJobRequest) SetZone(v []*string) *CreateSurveyJobRequest {
	s.Zone = v
	return s
}

func (s *CreateSurveyJobRequest) SetRegionId(v string) *CreateSurveyJobRequest {
	s.RegionId = &v
	return s
}

type CreateSurveyJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s CreateSurveyJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSurveyJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSurveyJobResponseBody) SetCode(v string) *CreateSurveyJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSurveyJobResponseBody) SetData(v interface{}) *CreateSurveyJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateSurveyJobResponseBody) SetSuccess(v bool) *CreateSurveyJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSurveyJobResponseBody) SetError(v string) *CreateSurveyJobResponseBody {
	s.Error = &v
	return s
}

type CreateSurveyJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSurveyJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSurveyJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSurveyJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSurveyJobResponse) SetHeaders(v map[string]*string) *CreateSurveyJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSurveyJobResponse) SetStatusCode(v int32) *CreateSurveyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSurveyJobResponse) SetBody(v *CreateSurveyJobResponseBody) *CreateSurveyJobResponse {
	s.Body = v
	return s
}

type CreateSurveyJobOfflineRequest struct {
	Channel    *string `json:"channel,omitempty" xml:"channel,omitempty"`
	CloudType  *string `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	FileName   *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateSurveyJobOfflineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSurveyJobOfflineRequest) GoString() string {
	return s.String()
}

func (s *CreateSurveyJobOfflineRequest) SetChannel(v string) *CreateSurveyJobOfflineRequest {
	s.Channel = &v
	return s
}

func (s *CreateSurveyJobOfflineRequest) SetCloudType(v string) *CreateSurveyJobOfflineRequest {
	s.CloudType = &v
	return s
}

func (s *CreateSurveyJobOfflineRequest) SetFileName(v string) *CreateSurveyJobOfflineRequest {
	s.FileName = &v
	return s
}

func (s *CreateSurveyJobOfflineRequest) SetName(v string) *CreateSurveyJobOfflineRequest {
	s.Name = &v
	return s
}

func (s *CreateSurveyJobOfflineRequest) SetObjectName(v string) *CreateSurveyJobOfflineRequest {
	s.ObjectName = &v
	return s
}

func (s *CreateSurveyJobOfflineRequest) SetRegionId(v string) *CreateSurveyJobOfflineRequest {
	s.RegionId = &v
	return s
}

type CreateSurveyJobOfflineResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s CreateSurveyJobOfflineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSurveyJobOfflineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSurveyJobOfflineResponseBody) SetCode(v string) *CreateSurveyJobOfflineResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSurveyJobOfflineResponseBody) SetData(v interface{}) *CreateSurveyJobOfflineResponseBody {
	s.Data = v
	return s
}

func (s *CreateSurveyJobOfflineResponseBody) SetSuccess(v bool) *CreateSurveyJobOfflineResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSurveyJobOfflineResponseBody) SetError(v string) *CreateSurveyJobOfflineResponseBody {
	s.Error = &v
	return s
}

type CreateSurveyJobOfflineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSurveyJobOfflineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSurveyJobOfflineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSurveyJobOfflineResponse) GoString() string {
	return s.String()
}

func (s *CreateSurveyJobOfflineResponse) SetHeaders(v map[string]*string) *CreateSurveyJobOfflineResponse {
	s.Headers = v
	return s
}

func (s *CreateSurveyJobOfflineResponse) SetStatusCode(v int32) *CreateSurveyJobOfflineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSurveyJobOfflineResponse) SetBody(v *CreateSurveyJobOfflineResponseBody) *CreateSurveyJobOfflineResponse {
	s.Body = v
	return s
}

type DeleteMigrationJobRequest struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobRequest) SetId(v int64) *DeleteMigrationJobRequest {
	s.Id = &v
	return s
}

type DeleteMigrationJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DeleteMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponseBody) SetCode(v string) *DeleteMigrationJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetData(v interface{}) *DeleteMigrationJobResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetSuccess(v bool) *DeleteMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetError(v string) *DeleteMigrationJobResponseBody {
	s.Error = &v
	return s
}

type DeleteMigrationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponse) SetHeaders(v map[string]*string) *DeleteMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteMigrationJobResponse) SetStatusCode(v int32) *DeleteMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMigrationJobResponse) SetBody(v *DeleteMigrationJobResponseBody) *DeleteMigrationJobResponse {
	s.Body = v
	return s
}

type DeleteOssFileRequest struct {
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteOssFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOssFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteOssFileRequest) SetObjectName(v string) *DeleteOssFileRequest {
	s.ObjectName = &v
	return s
}

func (s *DeleteOssFileRequest) SetRegionId(v string) *DeleteOssFileRequest {
	s.RegionId = &v
	return s
}

type DeleteOssFileResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DeleteOssFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOssFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOssFileResponseBody) SetCode(v string) *DeleteOssFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOssFileResponseBody) SetData(v interface{}) *DeleteOssFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteOssFileResponseBody) SetSuccess(v bool) *DeleteOssFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOssFileResponseBody) SetError(v string) *DeleteOssFileResponseBody {
	s.Error = &v
	return s
}

type DeleteOssFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteOssFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOssFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOssFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteOssFileResponse) SetHeaders(v map[string]*string) *DeleteOssFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteOssFileResponse) SetStatusCode(v int32) *DeleteOssFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOssFileResponse) SetBody(v *DeleteOssFileResponseBody) *DeleteOssFileResponse {
	s.Body = v
	return s
}

type DeleteSurveyJobRequest struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteSurveyJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSurveyJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSurveyJobRequest) SetId(v int64) *DeleteSurveyJobRequest {
	s.Id = &v
	return s
}

type DeleteSurveyJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DeleteSurveyJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSurveyJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSurveyJobResponseBody) SetCode(v string) *DeleteSurveyJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSurveyJobResponseBody) SetData(v interface{}) *DeleteSurveyJobResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSurveyJobResponseBody) SetSuccess(v bool) *DeleteSurveyJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSurveyJobResponseBody) SetError(v string) *DeleteSurveyJobResponseBody {
	s.Error = &v
	return s
}

type DeleteSurveyJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSurveyJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSurveyJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSurveyJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSurveyJobResponse) SetHeaders(v map[string]*string) *DeleteSurveyJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSurveyJobResponse) SetStatusCode(v int32) *DeleteSurveyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSurveyJobResponse) SetBody(v *DeleteSurveyJobResponseBody) *DeleteSurveyJobResponse {
	s.Body = v
	return s
}

type DeleteSurveyResourcesRequest struct {
	Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
}

func (s DeleteSurveyResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSurveyResourcesRequest) GoString() string {
	return s.String()
}

func (s *DeleteSurveyResourcesRequest) SetIds(v string) *DeleteSurveyResourcesRequest {
	s.Ids = &v
	return s
}

type DeleteSurveyResourcesResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSurveyResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSurveyResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSurveyResourcesResponseBody) SetCode(v string) *DeleteSurveyResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSurveyResourcesResponseBody) SetData(v interface{}) *DeleteSurveyResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSurveyResourcesResponseBody) SetError(v string) *DeleteSurveyResourcesResponseBody {
	s.Error = &v
	return s
}

func (s *DeleteSurveyResourcesResponseBody) SetSuccess(v bool) *DeleteSurveyResourcesResponseBody {
	s.Success = &v
	return s
}

type DeleteSurveyResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSurveyResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSurveyResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSurveyResourcesResponse) GoString() string {
	return s.String()
}

func (s *DeleteSurveyResourcesResponse) SetHeaders(v map[string]*string) *DeleteSurveyResourcesResponse {
	s.Headers = v
	return s
}

func (s *DeleteSurveyResourcesResponse) SetStatusCode(v int32) *DeleteSurveyResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSurveyResourcesResponse) SetBody(v *DeleteSurveyResourcesResponseBody) *DeleteSurveyResourcesResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobConfigResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DescribeMigrationJobConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobConfigResponseBody) SetCode(v string) *DescribeMigrationJobConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMigrationJobConfigResponseBody) SetData(v interface{}) *DescribeMigrationJobConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMigrationJobConfigResponseBody) SetSuccess(v bool) *DescribeMigrationJobConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobConfigResponseBody) SetError(v string) *DescribeMigrationJobConfigResponseBody {
	s.Error = &v
	return s
}

type DescribeMigrationJobConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMigrationJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMigrationJobConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobConfigResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobConfigResponse) SetStatusCode(v int32) *DescribeMigrationJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobConfigResponse) SetBody(v *DescribeMigrationJobConfigResponseBody) *DescribeMigrationJobConfigResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobCountRequest struct {
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	Source   *string   `json:"source,omitempty" xml:"source,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
	TypeList []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobCountRequest) SetName(v string) *DescribeMigrationJobCountRequest {
	s.Name = &v
	return s
}

func (s *DescribeMigrationJobCountRequest) SetSource(v string) *DescribeMigrationJobCountRequest {
	s.Source = &v
	return s
}

func (s *DescribeMigrationJobCountRequest) SetType(v string) *DescribeMigrationJobCountRequest {
	s.Type = &v
	return s
}

func (s *DescribeMigrationJobCountRequest) SetTypeList(v []*string) *DescribeMigrationJobCountRequest {
	s.TypeList = v
	return s
}

type DescribeMigrationJobCountResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DescribeMigrationJobCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobCountResponseBody) SetCode(v string) *DescribeMigrationJobCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMigrationJobCountResponseBody) SetData(v interface{}) *DescribeMigrationJobCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMigrationJobCountResponseBody) SetSuccess(v bool) *DescribeMigrationJobCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobCountResponseBody) SetError(v string) *DescribeMigrationJobCountResponseBody {
	s.Error = &v
	return s
}

type DescribeMigrationJobCountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMigrationJobCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMigrationJobCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobCountResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobCountResponse) SetStatusCode(v int32) *DescribeMigrationJobCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobCountResponse) SetBody(v *DescribeMigrationJobCountResponseBody) *DescribeMigrationJobCountResponse {
	s.Body = v
	return s
}

type DescribeOssStsRequest struct {
	Ak        *string `json:"ak,omitempty" xml:"ak,omitempty"`
	CloudType *string `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Sk        *string `json:"sk,omitempty" xml:"sk,omitempty"`
	TenantId  *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DescribeOssStsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssStsRequest) SetAk(v string) *DescribeOssStsRequest {
	s.Ak = &v
	return s
}

func (s *DescribeOssStsRequest) SetCloudType(v string) *DescribeOssStsRequest {
	s.CloudType = &v
	return s
}

func (s *DescribeOssStsRequest) SetRegion(v string) *DescribeOssStsRequest {
	s.Region = &v
	return s
}

func (s *DescribeOssStsRequest) SetRegionId(v string) *DescribeOssStsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOssStsRequest) SetSk(v string) *DescribeOssStsRequest {
	s.Sk = &v
	return s
}

func (s *DescribeOssStsRequest) SetTenantId(v string) *DescribeOssStsRequest {
	s.TenantId = &v
	return s
}

type DescribeOssStsResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DescribeOssStsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssStsResponseBody) SetCode(v string) *DescribeOssStsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOssStsResponseBody) SetData(v interface{}) *DescribeOssStsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOssStsResponseBody) SetSuccess(v bool) *DescribeOssStsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeOssStsResponseBody) SetError(v string) *DescribeOssStsResponseBody {
	s.Error = &v
	return s
}

type DescribeOssStsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOssStsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssStsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssStsResponse) SetHeaders(v map[string]*string) *DescribeOssStsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssStsResponse) SetStatusCode(v int32) *DescribeOssStsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssStsResponse) SetBody(v *DescribeOssStsResponseBody) *DescribeOssStsResponse {
	s.Body = v
	return s
}

type DescribeSummaryByStatusRequest struct {
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeSummaryByStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryByStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSummaryByStatusRequest) SetRegionId(v string) *DescribeSummaryByStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeSummaryByStatusResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSummaryByStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryByStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSummaryByStatusResponseBody) SetCode(v string) *DescribeSummaryByStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSummaryByStatusResponseBody) SetData(v interface{}) *DescribeSummaryByStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSummaryByStatusResponseBody) SetError(v string) *DescribeSummaryByStatusResponseBody {
	s.Error = &v
	return s
}

func (s *DescribeSummaryByStatusResponseBody) SetSuccess(v bool) *DescribeSummaryByStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeSummaryByStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSummaryByStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSummaryByStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryByStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSummaryByStatusResponse) SetHeaders(v map[string]*string) *DescribeSummaryByStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSummaryByStatusResponse) SetStatusCode(v int32) *DescribeSummaryByStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSummaryByStatusResponse) SetBody(v *DescribeSummaryByStatusResponseBody) *DescribeSummaryByStatusResponse {
	s.Body = v
	return s
}

type DescribeSummaryByStatusAndGroupRequest struct {
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeSummaryByStatusAndGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryByStatusAndGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeSummaryByStatusAndGroupRequest) SetRegionId(v string) *DescribeSummaryByStatusAndGroupRequest {
	s.RegionId = &v
	return s
}

type DescribeSummaryByStatusAndGroupResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSummaryByStatusAndGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryByStatusAndGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSummaryByStatusAndGroupResponseBody) SetCode(v string) *DescribeSummaryByStatusAndGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSummaryByStatusAndGroupResponseBody) SetData(v interface{}) *DescribeSummaryByStatusAndGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSummaryByStatusAndGroupResponseBody) SetError(v string) *DescribeSummaryByStatusAndGroupResponseBody {
	s.Error = &v
	return s
}

func (s *DescribeSummaryByStatusAndGroupResponseBody) SetSuccess(v bool) *DescribeSummaryByStatusAndGroupResponseBody {
	s.Success = &v
	return s
}

type DescribeSummaryByStatusAndGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSummaryByStatusAndGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSummaryByStatusAndGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryByStatusAndGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeSummaryByStatusAndGroupResponse) SetHeaders(v map[string]*string) *DescribeSummaryByStatusAndGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeSummaryByStatusAndGroupResponse) SetStatusCode(v int32) *DescribeSummaryByStatusAndGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSummaryByStatusAndGroupResponse) SetBody(v *DescribeSummaryByStatusAndGroupResponseBody) *DescribeSummaryByStatusAndGroupResponse {
	s.Body = v
	return s
}

type DescribeSurveyJobRequest struct {
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeSurveyJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeSurveyJobRequest) SetId(v int64) *DescribeSurveyJobRequest {
	s.Id = &v
	return s
}

func (s *DescribeSurveyJobRequest) SetRegionId(v string) *DescribeSurveyJobRequest {
	s.RegionId = &v
	return s
}

type DescribeSurveyJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DescribeSurveyJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSurveyJobResponseBody) SetCode(v string) *DescribeSurveyJobResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSurveyJobResponseBody) SetData(v interface{}) *DescribeSurveyJobResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSurveyJobResponseBody) SetSuccess(v bool) *DescribeSurveyJobResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSurveyJobResponseBody) SetError(v string) *DescribeSurveyJobResponseBody {
	s.Error = &v
	return s
}

type DescribeSurveyJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSurveyJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSurveyJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeSurveyJobResponse) SetHeaders(v map[string]*string) *DescribeSurveyJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeSurveyJobResponse) SetStatusCode(v int32) *DescribeSurveyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSurveyJobResponse) SetBody(v *DescribeSurveyJobResponseBody) *DescribeSurveyJobResponse {
	s.Body = v
	return s
}

type DescribeSurveyJobCountRequest struct {
	ChannelList   []*string `json:"channelList,omitempty" xml:"channelList,omitempty" type:"Repeated"`
	CloudTypeList []*string `json:"cloudTypeList,omitempty" xml:"cloudTypeList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	RegionId      *string   `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeSurveyJobCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyJobCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSurveyJobCountRequest) SetChannelList(v []*string) *DescribeSurveyJobCountRequest {
	s.ChannelList = v
	return s
}

func (s *DescribeSurveyJobCountRequest) SetCloudTypeList(v []*string) *DescribeSurveyJobCountRequest {
	s.CloudTypeList = v
	return s
}

func (s *DescribeSurveyJobCountRequest) SetName(v string) *DescribeSurveyJobCountRequest {
	s.Name = &v
	return s
}

func (s *DescribeSurveyJobCountRequest) SetRegionId(v string) *DescribeSurveyJobCountRequest {
	s.RegionId = &v
	return s
}

type DescribeSurveyJobCountResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DescribeSurveyJobCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyJobCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSurveyJobCountResponseBody) SetCode(v string) *DescribeSurveyJobCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSurveyJobCountResponseBody) SetData(v interface{}) *DescribeSurveyJobCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSurveyJobCountResponseBody) SetSuccess(v bool) *DescribeSurveyJobCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSurveyJobCountResponseBody) SetError(v string) *DescribeSurveyJobCountResponseBody {
	s.Error = &v
	return s
}

type DescribeSurveyJobCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSurveyJobCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSurveyJobCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyJobCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSurveyJobCountResponse) SetHeaders(v map[string]*string) *DescribeSurveyJobCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSurveyJobCountResponse) SetStatusCode(v int32) *DescribeSurveyJobCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSurveyJobCountResponse) SetBody(v *DescribeSurveyJobCountResponseBody) *DescribeSurveyJobCountResponse {
	s.Body = v
	return s
}

type DescribeSurveyResourceTagResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSurveyResourceTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyResourceTagResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSurveyResourceTagResponseBody) SetCode(v string) *DescribeSurveyResourceTagResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSurveyResourceTagResponseBody) SetData(v interface{}) *DescribeSurveyResourceTagResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSurveyResourceTagResponseBody) SetError(v string) *DescribeSurveyResourceTagResponseBody {
	s.Error = &v
	return s
}

func (s *DescribeSurveyResourceTagResponseBody) SetSuccess(v bool) *DescribeSurveyResourceTagResponseBody {
	s.Success = &v
	return s
}

type DescribeSurveyResourceTagResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSurveyResourceTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSurveyResourceTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyResourceTagResponse) GoString() string {
	return s.String()
}

func (s *DescribeSurveyResourceTagResponse) SetHeaders(v map[string]*string) *DescribeSurveyResourceTagResponse {
	s.Headers = v
	return s
}

func (s *DescribeSurveyResourceTagResponse) SetStatusCode(v int32) *DescribeSurveyResourceTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSurveyResourceTagResponse) SetBody(v *DescribeSurveyResourceTagResponseBody) *DescribeSurveyResourceTagResponse {
	s.Body = v
	return s
}

type DescribeSurveyTemplateRequest struct {
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DescribeSurveyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeSurveyTemplateRequest) SetResourceType(v string) *DescribeSurveyTemplateRequest {
	s.ResourceType = &v
	return s
}

type DescribeSurveyTemplateResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s DescribeSurveyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSurveyTemplateResponseBody) SetCode(v string) *DescribeSurveyTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSurveyTemplateResponseBody) SetData(v interface{}) *DescribeSurveyTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSurveyTemplateResponseBody) SetSuccess(v bool) *DescribeSurveyTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSurveyTemplateResponseBody) SetError(v string) *DescribeSurveyTemplateResponseBody {
	s.Error = &v
	return s
}

type DescribeSurveyTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSurveyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSurveyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSurveyTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeSurveyTemplateResponse) SetHeaders(v map[string]*string) *DescribeSurveyTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeSurveyTemplateResponse) SetStatusCode(v int32) *DescribeSurveyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSurveyTemplateResponse) SetBody(v *DescribeSurveyTemplateResponseBody) *DescribeSurveyTemplateResponse {
	s.Body = v
	return s
}

type ListMigrationJobsRequest struct {
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	PageNum  *int32    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SortCol  *string   `json:"sortCol,omitempty" xml:"sortCol,omitempty"`
	SortType *string   `json:"sortType,omitempty" xml:"sortType,omitempty"`
	Source   *string   `json:"source,omitempty" xml:"source,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
	TypeList []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
}

func (s ListMigrationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMigrationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationJobsRequest) SetName(v string) *ListMigrationJobsRequest {
	s.Name = &v
	return s
}

func (s *ListMigrationJobsRequest) SetPageNum(v int32) *ListMigrationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMigrationJobsRequest) SetPageSize(v int32) *ListMigrationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMigrationJobsRequest) SetSortCol(v string) *ListMigrationJobsRequest {
	s.SortCol = &v
	return s
}

func (s *ListMigrationJobsRequest) SetSortType(v string) *ListMigrationJobsRequest {
	s.SortType = &v
	return s
}

func (s *ListMigrationJobsRequest) SetSource(v string) *ListMigrationJobsRequest {
	s.Source = &v
	return s
}

func (s *ListMigrationJobsRequest) SetType(v string) *ListMigrationJobsRequest {
	s.Type = &v
	return s
}

func (s *ListMigrationJobsRequest) SetTypeList(v []*string) *ListMigrationJobsRequest {
	s.TypeList = v
	return s
}

type ListMigrationJobsResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s ListMigrationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationJobsResponseBody) SetCode(v string) *ListMigrationJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMigrationJobsResponseBody) SetData(v interface{}) *ListMigrationJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListMigrationJobsResponseBody) SetSuccess(v bool) *ListMigrationJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMigrationJobsResponseBody) SetError(v string) *ListMigrationJobsResponseBody {
	s.Error = &v
	return s
}

type ListMigrationJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMigrationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMigrationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMigrationJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMigrationJobsResponse) SetHeaders(v map[string]*string) *ListMigrationJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMigrationJobsResponse) SetStatusCode(v int32) *ListMigrationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMigrationJobsResponse) SetBody(v *ListMigrationJobsResponseBody) *ListMigrationJobsResponse {
	s.Body = v
	return s
}

type ListSurveyJobDownLoadJobsRequest struct {
	PageNum  *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SortCol  *string `json:"sortCol,omitempty" xml:"sortCol,omitempty"`
	SortType *string `json:"sortType,omitempty" xml:"sortType,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListSurveyJobDownLoadJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyJobDownLoadJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyJobDownLoadJobsRequest) SetPageNum(v int32) *ListSurveyJobDownLoadJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsRequest) SetPageSize(v int32) *ListSurveyJobDownLoadJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsRequest) SetSortCol(v string) *ListSurveyJobDownLoadJobsRequest {
	s.SortCol = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsRequest) SetSortType(v string) *ListSurveyJobDownLoadJobsRequest {
	s.SortType = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsRequest) SetRegionId(v string) *ListSurveyJobDownLoadJobsRequest {
	s.RegionId = &v
	return s
}

type ListSurveyJobDownLoadJobsResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s ListSurveyJobDownLoadJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyJobDownLoadJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSurveyJobDownLoadJobsResponseBody) SetCode(v string) *ListSurveyJobDownLoadJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsResponseBody) SetData(v interface{}) *ListSurveyJobDownLoadJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListSurveyJobDownLoadJobsResponseBody) SetSuccess(v bool) *ListSurveyJobDownLoadJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsResponseBody) SetError(v string) *ListSurveyJobDownLoadJobsResponseBody {
	s.Error = &v
	return s
}

type ListSurveyJobDownLoadJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSurveyJobDownLoadJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSurveyJobDownLoadJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyJobDownLoadJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSurveyJobDownLoadJobsResponse) SetHeaders(v map[string]*string) *ListSurveyJobDownLoadJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSurveyJobDownLoadJobsResponse) SetStatusCode(v int32) *ListSurveyJobDownLoadJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSurveyJobDownLoadJobsResponse) SetBody(v *ListSurveyJobDownLoadJobsResponseBody) *ListSurveyJobDownLoadJobsResponse {
	s.Body = v
	return s
}

type ListSurveyJobsRequest struct {
	ChannelList   []*string `json:"channelList,omitempty" xml:"channelList,omitempty" type:"Repeated"`
	CloudTypeList []*string `json:"cloudTypeList,omitempty" xml:"cloudTypeList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	RegionId      *string   `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListSurveyJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyJobsRequest) SetChannelList(v []*string) *ListSurveyJobsRequest {
	s.ChannelList = v
	return s
}

func (s *ListSurveyJobsRequest) SetCloudTypeList(v []*string) *ListSurveyJobsRequest {
	s.CloudTypeList = v
	return s
}

func (s *ListSurveyJobsRequest) SetName(v string) *ListSurveyJobsRequest {
	s.Name = &v
	return s
}

func (s *ListSurveyJobsRequest) SetRegionId(v string) *ListSurveyJobsRequest {
	s.RegionId = &v
	return s
}

type ListSurveyJobsResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s ListSurveyJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSurveyJobsResponseBody) SetCode(v string) *ListSurveyJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSurveyJobsResponseBody) SetData(v interface{}) *ListSurveyJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListSurveyJobsResponseBody) SetSuccess(v bool) *ListSurveyJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSurveyJobsResponseBody) SetError(v string) *ListSurveyJobsResponseBody {
	s.Error = &v
	return s
}

type ListSurveyJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSurveyJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSurveyJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSurveyJobsResponse) SetHeaders(v map[string]*string) *ListSurveyJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSurveyJobsResponse) SetStatusCode(v int32) *ListSurveyJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSurveyJobsResponse) SetBody(v *ListSurveyJobsResponseBody) *ListSurveyJobsResponse {
	s.Body = v
	return s
}

type ListSurveyResourceByMigrationGroupsRequest struct {
	Body *ListSurveyResourceByMigrationGroupsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ListSurveyResourceByMigrationGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceByMigrationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceByMigrationGroupsRequest) SetBody(v *ListSurveyResourceByMigrationGroupsRequestBody) *ListSurveyResourceByMigrationGroupsRequest {
	s.Body = v
	return s
}

type ListSurveyResourceByMigrationGroupsRequestBody struct {
	Ids []*int32 `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
}

func (s ListSurveyResourceByMigrationGroupsRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceByMigrationGroupsRequestBody) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceByMigrationGroupsRequestBody) SetIds(v []*int32) *ListSurveyResourceByMigrationGroupsRequestBody {
	s.Ids = v
	return s
}

type ListSurveyResourceByMigrationGroupsShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSurveyResourceByMigrationGroupsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceByMigrationGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceByMigrationGroupsShrinkRequest) SetBodyShrink(v string) *ListSurveyResourceByMigrationGroupsShrinkRequest {
	s.BodyShrink = &v
	return s
}

type ListSurveyResourceByMigrationGroupsResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSurveyResourceByMigrationGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceByMigrationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceByMigrationGroupsResponseBody) SetCode(v string) *ListSurveyResourceByMigrationGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSurveyResourceByMigrationGroupsResponseBody) SetData(v interface{}) *ListSurveyResourceByMigrationGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListSurveyResourceByMigrationGroupsResponseBody) SetError(v string) *ListSurveyResourceByMigrationGroupsResponseBody {
	s.Error = &v
	return s
}

func (s *ListSurveyResourceByMigrationGroupsResponseBody) SetSuccess(v bool) *ListSurveyResourceByMigrationGroupsResponseBody {
	s.Success = &v
	return s
}

type ListSurveyResourceByMigrationGroupsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSurveyResourceByMigrationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSurveyResourceByMigrationGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceByMigrationGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceByMigrationGroupsResponse) SetHeaders(v map[string]*string) *ListSurveyResourceByMigrationGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSurveyResourceByMigrationGroupsResponse) SetStatusCode(v int32) *ListSurveyResourceByMigrationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSurveyResourceByMigrationGroupsResponse) SetBody(v *ListSurveyResourceByMigrationGroupsResponseBody) *ListSurveyResourceByMigrationGroupsResponse {
	s.Body = v
	return s
}

type ListSurveyResourceConnectionsRequest struct {
	Ids      []*int32 `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
	RegionId *string  `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListSurveyResourceConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceConnectionsRequest) SetIds(v []*int32) *ListSurveyResourceConnectionsRequest {
	s.Ids = v
	return s
}

func (s *ListSurveyResourceConnectionsRequest) SetRegionId(v string) *ListSurveyResourceConnectionsRequest {
	s.RegionId = &v
	return s
}

type ListSurveyResourceConnectionsResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Error   *string     `json:"Error,omitempty" xml:"Error,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSurveyResourceConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceConnectionsResponseBody) SetCode(v string) *ListSurveyResourceConnectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSurveyResourceConnectionsResponseBody) SetData(v interface{}) *ListSurveyResourceConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListSurveyResourceConnectionsResponseBody) SetError(v string) *ListSurveyResourceConnectionsResponseBody {
	s.Error = &v
	return s
}

func (s *ListSurveyResourceConnectionsResponseBody) SetSuccess(v bool) *ListSurveyResourceConnectionsResponseBody {
	s.Success = &v
	return s
}

type ListSurveyResourceConnectionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSurveyResourceConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSurveyResourceConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceConnectionsResponse) SetHeaders(v map[string]*string) *ListSurveyResourceConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListSurveyResourceConnectionsResponse) SetStatusCode(v int32) *ListSurveyResourceConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSurveyResourceConnectionsResponse) SetBody(v *ListSurveyResourceConnectionsResponseBody) *ListSurveyResourceConnectionsResponse {
	s.Body = v
	return s
}

type ListSurveyResourceTypesRequest struct {
	Ak        *string `json:"ak,omitempty" xml:"ak,omitempty"`
	CloudType *string `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Sk        *string `json:"sk,omitempty" xml:"sk,omitempty"`
}

func (s ListSurveyResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceTypesRequest) SetAk(v string) *ListSurveyResourceTypesRequest {
	s.Ak = &v
	return s
}

func (s *ListSurveyResourceTypesRequest) SetCloudType(v string) *ListSurveyResourceTypesRequest {
	s.CloudType = &v
	return s
}

func (s *ListSurveyResourceTypesRequest) SetRegion(v string) *ListSurveyResourceTypesRequest {
	s.Region = &v
	return s
}

func (s *ListSurveyResourceTypesRequest) SetRegionId(v string) *ListSurveyResourceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSurveyResourceTypesRequest) SetSk(v string) *ListSurveyResourceTypesRequest {
	s.Sk = &v
	return s
}

type ListSurveyResourceTypesResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s ListSurveyResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceTypesResponseBody) SetCode(v string) *ListSurveyResourceTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSurveyResourceTypesResponseBody) SetData(v interface{}) *ListSurveyResourceTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListSurveyResourceTypesResponseBody) SetSuccess(v bool) *ListSurveyResourceTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSurveyResourceTypesResponseBody) SetError(v string) *ListSurveyResourceTypesResponseBody {
	s.Error = &v
	return s
}

type ListSurveyResourceTypesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSurveyResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSurveyResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListSurveyResourceTypesResponse) SetHeaders(v map[string]*string) *ListSurveyResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListSurveyResourceTypesResponse) SetStatusCode(v int32) *ListSurveyResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSurveyResourceTypesResponse) SetBody(v *ListSurveyResourceTypesResponseBody) *ListSurveyResourceTypesResponse {
	s.Body = v
	return s
}

type ListSurveyResourcesDetailRequest struct {
	Ip           *string `json:"ip,omitempty" xml:"ip,omitempty"`
	JobId        *int64  `json:"jobId,omitempty" xml:"jobId,omitempty"`
	ProjectId    *int64  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubProjectId *int64  `json:"subProjectId,omitempty" xml:"subProjectId,omitempty"`
	RegionId     *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListSurveyResourcesDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourcesDetailRequest) GoString() string {
	return s.String()
}

func (s *ListSurveyResourcesDetailRequest) SetIp(v string) *ListSurveyResourcesDetailRequest {
	s.Ip = &v
	return s
}

func (s *ListSurveyResourcesDetailRequest) SetJobId(v int64) *ListSurveyResourcesDetailRequest {
	s.JobId = &v
	return s
}

func (s *ListSurveyResourcesDetailRequest) SetProjectId(v int64) *ListSurveyResourcesDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSurveyResourcesDetailRequest) SetResourceType(v string) *ListSurveyResourcesDetailRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSurveyResourcesDetailRequest) SetSubProjectId(v int64) *ListSurveyResourcesDetailRequest {
	s.SubProjectId = &v
	return s
}

func (s *ListSurveyResourcesDetailRequest) SetRegionId(v string) *ListSurveyResourcesDetailRequest {
	s.RegionId = &v
	return s
}

type ListSurveyResourcesDetailResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s ListSurveyResourcesDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourcesDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListSurveyResourcesDetailResponseBody) SetCode(v string) *ListSurveyResourcesDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListSurveyResourcesDetailResponseBody) SetData(v interface{}) *ListSurveyResourcesDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListSurveyResourcesDetailResponseBody) SetSuccess(v bool) *ListSurveyResourcesDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListSurveyResourcesDetailResponseBody) SetError(v string) *ListSurveyResourcesDetailResponseBody {
	s.Error = &v
	return s
}

type ListSurveyResourcesDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSurveyResourcesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSurveyResourcesDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSurveyResourcesDetailResponse) GoString() string {
	return s.String()
}

func (s *ListSurveyResourcesDetailResponse) SetHeaders(v map[string]*string) *ListSurveyResourcesDetailResponse {
	s.Headers = v
	return s
}

func (s *ListSurveyResourcesDetailResponse) SetStatusCode(v int32) *ListSurveyResourcesDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSurveyResourcesDetailResponse) SetBody(v *ListSurveyResourcesDetailResponseBody) *ListSurveyResourcesDetailResponse {
	s.Body = v
	return s
}

type RecoverMigrationJobRequest struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s RecoverMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *RecoverMigrationJobRequest) SetId(v int64) *RecoverMigrationJobRequest {
	s.Id = &v
	return s
}

type RecoverMigrationJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s RecoverMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverMigrationJobResponseBody) SetCode(v string) *RecoverMigrationJobResponseBody {
	s.Code = &v
	return s
}

func (s *RecoverMigrationJobResponseBody) SetData(v interface{}) *RecoverMigrationJobResponseBody {
	s.Data = v
	return s
}

func (s *RecoverMigrationJobResponseBody) SetSuccess(v bool) *RecoverMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *RecoverMigrationJobResponseBody) SetError(v string) *RecoverMigrationJobResponseBody {
	s.Error = &v
	return s
}

type RecoverMigrationJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoverMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoverMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *RecoverMigrationJobResponse) SetHeaders(v map[string]*string) *RecoverMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *RecoverMigrationJobResponse) SetStatusCode(v int32) *RecoverMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverMigrationJobResponse) SetBody(v *RecoverMigrationJobResponseBody) *RecoverMigrationJobResponse {
	s.Body = v
	return s
}

type StopSyncMigrationJobRequest struct {
	JobType  *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopSyncMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSyncMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StopSyncMigrationJobRequest) SetJobType(v string) *StopSyncMigrationJobRequest {
	s.JobType = &v
	return s
}

func (s *StopSyncMigrationJobRequest) SetRegionId(v string) *StopSyncMigrationJobRequest {
	s.RegionId = &v
	return s
}

type StopSyncMigrationJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s StopSyncMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSyncMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopSyncMigrationJobResponseBody) SetCode(v string) *StopSyncMigrationJobResponseBody {
	s.Code = &v
	return s
}

func (s *StopSyncMigrationJobResponseBody) SetData(v interface{}) *StopSyncMigrationJobResponseBody {
	s.Data = v
	return s
}

func (s *StopSyncMigrationJobResponseBody) SetSuccess(v bool) *StopSyncMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopSyncMigrationJobResponseBody) SetError(v string) *StopSyncMigrationJobResponseBody {
	s.Error = &v
	return s
}

type StopSyncMigrationJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopSyncMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopSyncMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopSyncMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *StopSyncMigrationJobResponse) SetHeaders(v map[string]*string) *StopSyncMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *StopSyncMigrationJobResponse) SetStatusCode(v int32) *StopSyncMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSyncMigrationJobResponse) SetBody(v *StopSyncMigrationJobResponseBody) *StopSyncMigrationJobResponse {
	s.Body = v
	return s
}

type SyncMigrationJobRequest struct {
	JobType       *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	RegionId      *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Regions       *string `json:"regions,omitempty" xml:"regions,omitempty"`
}

func (s SyncMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *SyncMigrationJobRequest) SetJobType(v string) *SyncMigrationJobRequest {
	s.JobType = &v
	return s
}

func (s *SyncMigrationJobRequest) SetOperationType(v string) *SyncMigrationJobRequest {
	s.OperationType = &v
	return s
}

func (s *SyncMigrationJobRequest) SetRegionId(v string) *SyncMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *SyncMigrationJobRequest) SetRegions(v string) *SyncMigrationJobRequest {
	s.Regions = &v
	return s
}

type SyncMigrationJobResponseBody struct {
	Code    *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Error   *string     `json:"error,omitempty" xml:"error,omitempty"`
}

func (s SyncMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMigrationJobResponseBody) SetCode(v string) *SyncMigrationJobResponseBody {
	s.Code = &v
	return s
}

func (s *SyncMigrationJobResponseBody) SetData(v interface{}) *SyncMigrationJobResponseBody {
	s.Data = v
	return s
}

func (s *SyncMigrationJobResponseBody) SetSuccess(v bool) *SyncMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SyncMigrationJobResponseBody) SetError(v string) *SyncMigrationJobResponseBody {
	s.Error = &v
	return s
}

type SyncMigrationJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *SyncMigrationJobResponse) SetHeaders(v map[string]*string) *SyncMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *SyncMigrationJobResponse) SetStatusCode(v int32) *SyncMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMigrationJobResponse) SetBody(v *SyncMigrationJobResponseBody) *SyncMigrationJobResponse {
	s.Body = v
	return s
}

type UpdatePushAppHeaders struct {
	CommonHeaders  map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	TenantProxyUid *string            `json:"tenant-proxy-uid,omitempty" xml:"tenant-proxy-uid,omitempty"`
}

func (s UpdatePushAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePushAppHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePushAppHeaders) SetCommonHeaders(v map[string]*string) *UpdatePushAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePushAppHeaders) SetTenantProxyUid(v string) *UpdatePushAppHeaders {
	s.TenantProxyUid = &v
	return s
}

type UpdatePushAppRequest struct {
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s UpdatePushAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePushAppRequest) GoString() string {
	return s.String()
}

func (s *UpdatePushAppRequest) SetAppId(v int64) *UpdatePushAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdatePushAppRequest) SetDownloadLink(v string) *UpdatePushAppRequest {
	s.DownloadLink = &v
	return s
}

type UpdatePushAppResponseBody struct {
	IsDebugEnable *bool       `json:"IsDebugEnable,omitempty" xml:"IsDebugEnable,omitempty"`
	RequestId     *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObj     interface{} `json:"ResultObj,omitempty" xml:"ResultObj,omitempty"`
	Success       *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	TipMsg        *string     `json:"TipMsg,omitempty" xml:"TipMsg,omitempty"`
}

func (s UpdatePushAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePushAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePushAppResponseBody) SetIsDebugEnable(v bool) *UpdatePushAppResponseBody {
	s.IsDebugEnable = &v
	return s
}

func (s *UpdatePushAppResponseBody) SetRequestId(v string) *UpdatePushAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePushAppResponseBody) SetResultObj(v interface{}) *UpdatePushAppResponseBody {
	s.ResultObj = v
	return s
}

func (s *UpdatePushAppResponseBody) SetSuccess(v bool) *UpdatePushAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePushAppResponseBody) SetTipMsg(v string) *UpdatePushAppResponseBody {
	s.TipMsg = &v
	return s
}

type UpdatePushAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePushAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePushAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePushAppResponse) GoString() string {
	return s.String()
}

func (s *UpdatePushAppResponse) SetHeaders(v map[string]*string) *UpdatePushAppResponse {
	s.Headers = v
	return s
}

func (s *UpdatePushAppResponse) SetStatusCode(v int32) *UpdatePushAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePushAppResponse) SetBody(v *UpdatePushAppResponseBody) *UpdatePushAppResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("apds"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateFileJob(request *CreateFileJobRequest) (_result *CreateFileJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFileJobResponse{}
	_body, _err := client.CreateFileJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileJobWithOptions(request *CreateFileJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFileJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SurveyJobId)) {
		body["surveyJobId"] = request.SurveyJobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/file-job/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMigrationGroup(request *CreateMigrationGroupRequest) (_result *CreateMigrationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMigrationGroupResponse{}
	_body, _err := client.CreateMigrationGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMigrationGroupWithOptions(request *CreateMigrationGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMigrationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMigrationGroup"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-group/save-migration-group"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMigrationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMigrationJob(request *CreateMigrationJobRequest) (_result *CreateMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.CreateMigrationJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMigrationJobWithOptions(request *CreateMigrationJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobList)) {
		body["migrationJobList"] = request.MigrationJobList
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMigrationJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/create-migration-jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePayOrderCallback(request *CreatePayOrderCallbackRequest) (_result *CreatePayOrderCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePayOrderCallbackResponse{}
	_body, _err := client.CreatePayOrderCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePayOrderCallbackWithOptions(request *CreatePayOrderCallbackRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePayOrderCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePayOrderCallback"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sys/user/pop/api/v1/payOrderCallback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePayOrderCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRefund(request *CreateRefundRequest) (_result *CreateRefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRefundResponse{}
	_body, _err := client.CreateRefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRefundWithOptions(request *CreateRefundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRefund"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sys/user/pop/api/v1/refund"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSurveyJob(request *CreateSurveyJobRequest) (_result *CreateSurveyJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSurveyJobResponse{}
	_body, _err := client.CreateSurveyJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSurveyJobWithOptions(request *CreateSurveyJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSurveyJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ak)) {
		body["ak"] = request.Ak
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		body["cloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeList)) {
		body["resourceTypeList"] = request.ResourceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.Sk)) {
		body["sk"] = request.Sk
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		body["zone"] = request.Zone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSurveyJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/add-survey-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSurveyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSurveyJobOffline(request *CreateSurveyJobOfflineRequest) (_result *CreateSurveyJobOfflineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSurveyJobOfflineResponse{}
	_body, _err := client.CreateSurveyJobOfflineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSurveyJobOfflineWithOptions(request *CreateSurveyJobOfflineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSurveyJobOfflineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		body["cloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectName)) {
		body["objectName"] = request.ObjectName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSurveyJobOffline"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/add-import-survey-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSurveyJobOfflineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMigrationJob(request *DeleteMigrationJobRequest) (_result *DeleteMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.DeleteMigrationJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMigrationJobWithOptions(request *DeleteMigrationJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMigrationJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/remove-migration-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOssFile(request *DeleteOssFileRequest) (_result *DeleteOssFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteOssFileResponse{}
	_body, _err := client.DeleteOssFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOssFileWithOptions(request *DeleteOssFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteOssFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectName)) {
		query["objectName"] = request.ObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOssFile"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/file-job/delete-file"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOssFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSurveyJob(request *DeleteSurveyJobRequest) (_result *DeleteSurveyJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSurveyJobResponse{}
	_body, _err := client.DeleteSurveyJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSurveyJobWithOptions(request *DeleteSurveyJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSurveyJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSurveyJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/delete-survey-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSurveyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSurveyResources(request *DeleteSurveyResourcesRequest) (_result *DeleteSurveyResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSurveyResourcesResponse{}
	_body, _err := client.DeleteSurveyResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSurveyResourcesWithOptions(request *DeleteSurveyResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSurveyResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSurveyResources"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/confirm-resource/destroy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSurveyResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMigrationJobConfig() (_result *DescribeMigrationJobConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeMigrationJobConfigResponse{}
	_body, _err := client.DescribeMigrationJobConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMigrationJobConfigWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMigrationJobConfig"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/describe-migration-job-config"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMigrationJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMigrationJobCount(request *DescribeMigrationJobCountRequest) (_result *DescribeMigrationJobCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeMigrationJobCountResponse{}
	_body, _err := client.DescribeMigrationJobCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMigrationJobCountWithOptions(request *DescribeMigrationJobCountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.TypeList)) {
		body["typeList"] = request.TypeList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMigrationJobCount"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/count-migration-jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMigrationJobCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssSts(request *DescribeOssStsRequest) (_result *DescribeOssStsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeOssStsResponse{}
	_body, _err := client.DescribeOssStsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssStsWithOptions(request *DescribeOssStsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeOssStsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ak)) {
		query["ak"] = request.Ak
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		query["cloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Sk)) {
		query["sk"] = request.Sk
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssSts"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/file-job/sts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssStsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSummaryByStatus(request *DescribeSummaryByStatusRequest) (_result *DescribeSummaryByStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSummaryByStatusResponse{}
	_body, _err := client.DescribeSummaryByStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSummaryByStatusWithOptions(request *DescribeSummaryByStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSummaryByStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSummaryByStatus"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/summary/summary-by-status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSummaryByStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSummaryByStatusAndGroup(request *DescribeSummaryByStatusAndGroupRequest) (_result *DescribeSummaryByStatusAndGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSummaryByStatusAndGroupResponse{}
	_body, _err := client.DescribeSummaryByStatusAndGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSummaryByStatusAndGroupWithOptions(request *DescribeSummaryByStatusAndGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSummaryByStatusAndGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSummaryByStatusAndGroup"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/summary/summary-by-status-and-region"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSummaryByStatusAndGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSurveyJob(request *DescribeSurveyJobRequest) (_result *DescribeSurveyJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSurveyJobResponse{}
	_body, _err := client.DescribeSurveyJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSurveyJobWithOptions(request *DescribeSurveyJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSurveyJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSurveyJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/query-survey-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSurveyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSurveyJobCount(request *DescribeSurveyJobCountRequest) (_result *DescribeSurveyJobCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSurveyJobCountResponse{}
	_body, _err := client.DescribeSurveyJobCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSurveyJobCountWithOptions(request *DescribeSurveyJobCountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSurveyJobCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelList)) {
		body["channelList"] = request.ChannelList
	}

	if !tea.BoolValue(util.IsUnset(request.CloudTypeList)) {
		body["cloudTypeList"] = request.CloudTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSurveyJobCount"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/count-survey-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSurveyJobCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSurveyResourceTag() (_result *DescribeSurveyResourceTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSurveyResourceTagResponse{}
	_body, _err := client.DescribeSurveyResourceTagWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSurveyResourceTagWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSurveyResourceTagResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSurveyResourceTag"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/confirm-resource/get-resource-tag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSurveyResourceTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSurveyTemplate(request *DescribeSurveyTemplateRequest) (_result *DescribeSurveyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSurveyTemplateResponse{}
	_body, _err := client.DescribeSurveyTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSurveyTemplateWithOptions(request *DescribeSurveyTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSurveyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSurveyTemplate"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/survey-template/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSurveyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMigrationJobs(request *ListMigrationJobsRequest) (_result *ListMigrationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMigrationJobsResponse{}
	_body, _err := client.ListMigrationJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMigrationJobsWithOptions(request *ListMigrationJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMigrationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortCol)) {
		body["sortCol"] = request.SortCol
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		body["sortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.TypeList)) {
		body["typeList"] = request.TypeList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMigrationJobs"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/describe-migration-jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMigrationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSurveyJobDownLoadJobs(request *ListSurveyJobDownLoadJobsRequest) (_result *ListSurveyJobDownLoadJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSurveyJobDownLoadJobsResponse{}
	_body, _err := client.ListSurveyJobDownLoadJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSurveyJobDownLoadJobsWithOptions(request *ListSurveyJobDownLoadJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSurveyJobDownLoadJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortCol)) {
		body["sortCol"] = request.SortCol
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		body["sortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSurveyJobDownLoadJobs"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/file-job/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSurveyJobDownLoadJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSurveyJobs(request *ListSurveyJobsRequest) (_result *ListSurveyJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSurveyJobsResponse{}
	_body, _err := client.ListSurveyJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSurveyJobsWithOptions(request *ListSurveyJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSurveyJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelList)) {
		body["channelList"] = request.ChannelList
	}

	if !tea.BoolValue(util.IsUnset(request.CloudTypeList)) {
		body["cloudTypeList"] = request.CloudTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSurveyJobs"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/query-survey-jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSurveyJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSurveyResourceByMigrationGroups(request *ListSurveyResourceByMigrationGroupsRequest) (_result *ListSurveyResourceByMigrationGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSurveyResourceByMigrationGroupsResponse{}
	_body, _err := client.ListSurveyResourceByMigrationGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSurveyResourceByMigrationGroupsWithOptions(tmpReq *ListSurveyResourceByMigrationGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSurveyResourceByMigrationGroupsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSurveyResourceByMigrationGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSurveyResourceByMigrationGroups"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-group/get-survey-resource"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSurveyResourceByMigrationGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSurveyResourceConnections(request *ListSurveyResourceConnectionsRequest) (_result *ListSurveyResourceConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSurveyResourceConnectionsResponse{}
	_body, _err := client.ListSurveyResourceConnectionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSurveyResourceConnectionsWithOptions(request *ListSurveyResourceConnectionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSurveyResourceConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSurveyResourceConnections"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/resource-connects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSurveyResourceConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSurveyResourceTypes(request *ListSurveyResourceTypesRequest) (_result *ListSurveyResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSurveyResourceTypesResponse{}
	_body, _err := client.ListSurveyResourceTypesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSurveyResourceTypesWithOptions(request *ListSurveyResourceTypesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSurveyResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ak)) {
		query["ak"] = request.Ak
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		query["cloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Sk)) {
		query["sk"] = request.Sk
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSurveyResourceTypes"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/winback/query-resource-type"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSurveyResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSurveyResourcesDetail(request *ListSurveyResourcesDetailRequest) (_result *ListSurveyResourcesDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSurveyResourcesDetailResponse{}
	_body, _err := client.ListSurveyResourcesDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSurveyResourcesDetailWithOptions(request *ListSurveyResourcesDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSurveyResourcesDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["jobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SubProjectId)) {
		body["subProjectId"] = request.SubProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSurveyResourcesDetail"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/survey-detail/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSurveyResourcesDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverMigrationJob(request *RecoverMigrationJobRequest) (_result *RecoverMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecoverMigrationJobResponse{}
	_body, _err := client.RecoverMigrationJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverMigrationJobWithOptions(request *RecoverMigrationJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecoverMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverMigrationJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/recover-migration-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopSyncMigrationJob(request *StopSyncMigrationJobRequest) (_result *StopSyncMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopSyncMigrationJobResponse{}
	_body, _err := client.StopSyncMigrationJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopSyncMigrationJobWithOptions(request *StopSyncMigrationJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopSyncMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["jobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopSyncMigrationJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/unsync-migration-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopSyncMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncMigrationJob(request *SyncMigrationJobRequest) (_result *SyncMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncMigrationJobResponse{}
	_body, _err := client.SyncMigrationJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncMigrationJobWithOptions(request *SyncMigrationJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SyncMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["jobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["operationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Regions)) {
		query["regions"] = request.Regions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncMigrationJob"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/okss-services/migration-job/sync-migration-job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePushApp(request *UpdatePushAppRequest) (_result *UpdatePushAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePushAppHeaders{}
	_result = &UpdatePushAppResponse{}
	_body, _err := client.UpdatePushAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePushAppWithOptions(request *UpdatePushAppRequest, headers *UpdatePushAppHeaders, runtime *util.RuntimeOptions) (_result *UpdatePushAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadLink)) {
		body["DownloadLink"] = request.DownloadLink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.TenantProxyUid)) {
		realHeaders["tenant-proxy-uid"] = util.ToJSONString(headers.TenantProxyUid)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePushApp"),
		Version:     tea.String("2022-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/abm/app/manager/api/v1/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePushAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
