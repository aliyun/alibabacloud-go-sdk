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

type DeleteServiceInstancesRequest struct {
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceInstanceId []*string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty" type:"Repeated"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetClientToken(v string) *DeleteServiceInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetRegionId(v string) *DeleteServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest {
	s.ServiceInstanceId = v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	Outputs                   *string                                `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Status                    *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName              *string                                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UpdateTime                *string                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Progress                  *int64                                 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Parameters                *string                                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RequestId                 *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstanceId         *string                                `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	CreateTime                *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StatusDetail              *string                                `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Resources                 *string                                `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service                   *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	OperationStartTime        *string                                `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	OperationEndTime          *string                                `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OperatedServiceInstanceId *string                                `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	EnableInstanceOps         *bool                                  `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	IsOperated                *bool                                  `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	Status         *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	PublishTime    *string                                              `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Version        *string                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployMetadata *string                                              `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	DeployType     *string                                              `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId      *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl    *string                                              `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType    *string                                              `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName   *string                                              `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos   []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type ListServiceInstanceLogsRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults        *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequest) SetRegionId(v string) *ListServiceInstanceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetMaxResults(v string) *ListServiceInstanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetNextToken(v string) *ListServiceInstanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest {
	s.ServiceInstanceId = &v
	return s
}

type ListServiceInstanceLogsResponseBody struct {
	NextToken            *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults           *string                                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ServiceInstancesLogs []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs `json:"ServiceInstancesLogs,omitempty" xml:"ServiceInstancesLogs,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBody) SetNextToken(v string) *ListServiceInstanceLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetRequestId(v string) *ListServiceInstanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetMaxResults(v string) *ListServiceInstanceLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetServiceInstancesLogs(v []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs) *ListServiceInstanceLogsResponseBody {
	s.ServiceInstancesLogs = v
	return s
}

type ListServiceInstanceLogsResponseBodyServiceInstancesLogs struct {
	Timestamp         *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	Source            *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Phase             *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetTimestamp(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Timestamp = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetServiceInstanceId(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetSource(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Source = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetPhase(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Phase = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetContent(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Content = &v
	return s
}

type ListServiceInstanceLogsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceInstanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceInstanceLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse {
	s.Body = v
	return s
}

type CreateServiceInstanceRequest struct {
	RegionId          *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId         *string                                        `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion    *string                                        `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	Parameters        map[string]interface{}                         `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ClientToken       *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EnableInstanceOps *bool                                          `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	EnableAccountOps  *bool                                          `json:"EnableAccountOps,omitempty" xml:"EnableAccountOps,omitempty"`
	TemplateName      *string                                        `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	OperationMetadata *CreateServiceInstanceRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	RequestTags       []*CreateServiceInstanceRequestRequestTags     `json:"RequestTags,omitempty" xml:"RequestTags,omitempty" type:"Repeated"`
}

func (s CreateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceVersion(v string) *CreateServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableAccountOps(v bool) *CreateServiceInstanceRequest {
	s.EnableAccountOps = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetOperationMetadata(v *CreateServiceInstanceRequestOperationMetadata) *CreateServiceInstanceRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceRequest) SetRequestTags(v []*CreateServiceInstanceRequestRequestTags) *CreateServiceInstanceRequest {
	s.RequestTags = v
	return s
}

type CreateServiceInstanceRequestOperationMetadata struct {
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Resources         *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s CreateServiceInstanceRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestOperationMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.StartTime = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetEndTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetResources(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.Resources = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetServiceInstanceId(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.ServiceInstanceId = &v
	return s
}

type CreateServiceInstanceRequestRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceRequestRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestRequestTags) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestRequestTags) SetKey(v string) *CreateServiceInstanceRequestRequestTags {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestRequestTags) SetValue(v string) *CreateServiceInstanceRequestRequestTags {
	s.Value = &v
	return s
}

type CreateServiceInstanceShrinkRequest struct {
	RegionId          *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId         *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion    *string                                              `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ParametersShrink  *string                                              `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ClientToken       *string                                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EnableInstanceOps *bool                                                `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	EnableAccountOps  *bool                                                `json:"EnableAccountOps,omitempty" xml:"EnableAccountOps,omitempty"`
	TemplateName      *string                                              `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	OperationMetadata *CreateServiceInstanceShrinkRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	RequestTags       []*CreateServiceInstanceShrinkRequestRequestTags     `json:"RequestTags,omitempty" xml:"RequestTags,omitempty" type:"Repeated"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableAccountOps(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableAccountOps = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetOperationMetadata(v *CreateServiceInstanceShrinkRequestOperationMetadata) *CreateServiceInstanceShrinkRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRequestTags(v []*CreateServiceInstanceShrinkRequestRequestTags) *CreateServiceInstanceShrinkRequest {
	s.RequestTags = v
	return s
}

type CreateServiceInstanceShrinkRequestOperationMetadata struct {
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Resources         *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.StartTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetEndTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetResources(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.Resources = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetServiceInstanceId(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.ServiceInstanceId = &v
	return s
}

type CreateServiceInstanceShrinkRequestRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestRequestTags) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestRequestTags) SetKey(v string) *CreateServiceInstanceShrinkRequestRequestTags {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestRequestTags) SetValue(v string) *CreateServiceInstanceShrinkRequestRequestTags {
	s.Value = &v
	return s
}

type CreateServiceInstanceResponseBody struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

type CreateServiceInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

type ContinueDeployServiceInstanceRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Parameters        *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ContinueDeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceRequest) SetClientToken(v string) *ContinueDeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

type ContinueDeployServiceInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinueDeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBody) SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ContinueDeployServiceInstanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContinueDeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinueDeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponse) SetHeaders(v map[string]*string) *ContinueDeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetBody(v *ContinueDeployServiceInstanceResponseBody) *ContinueDeployServiceInstanceResponse {
	s.Body = v
	return s
}

type DisableOperationServiceInstanceRequest struct {
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableOperationServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableOperationServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DisableOperationServiceInstanceRequest) SetServiceInstanceId(v string) *DisableOperationServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *DisableOperationServiceInstanceRequest) SetRegionId(v string) *DisableOperationServiceInstanceRequest {
	s.RegionId = &v
	return s
}

type DisableOperationServiceInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableOperationServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableOperationServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableOperationServiceInstanceResponseBody) SetRequestId(v string) *DisableOperationServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DisableOperationServiceInstanceResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableOperationServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableOperationServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableOperationServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DisableOperationServiceInstanceResponse) SetHeaders(v map[string]*string) *DisableOperationServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DisableOperationServiceInstanceResponse) SetBody(v *DisableOperationServiceInstanceResponseBody) *DisableOperationServiceInstanceResponse {
	s.Body = v
	return s
}

type DeployServiceInstanceRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceRequest) SetClientToken(v string) *DeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetServiceInstanceId(v string) *DeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetRegionId(v string) *DeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

type DeployServiceInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponseBody) SetRequestId(v string) *DeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeployServiceInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponse) SetHeaders(v map[string]*string) *DeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployServiceInstanceResponse) SetBody(v *DeployServiceInstanceResponseBody) *DeployServiceInstanceResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	RegionId    *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults  *string                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Filter      []*ListServiceInstancesRequestFilter      `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	RequestTags []*ListServiceInstancesRequestRequestTags `json:"RequestTags,omitempty" xml:"RequestTags,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v string) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetRequestTags(v []*ListServiceInstancesRequestRequestTags) *ListServiceInstancesRequest {
	s.RequestTags = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

type ListServiceInstancesRequestRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestRequestTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestRequestTags) SetKey(v string) *ListServiceInstancesRequestRequestTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestRequestTags) SetValue(v string) *ListServiceInstancesRequestRequestTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	NextToken        *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults       *string                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int64) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v string) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	Status                    *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Outputs                   *string                                                  `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	UpdateTime                *string                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Parameters                *string                                                  `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ServiceInstanceId         *string                                                  `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	CreateTime                *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StatusDetail              *string                                                  `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Progress                  *int64                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Resources                 *string                                                  `json:"Resources,omitempty" xml:"Resources,omitempty"`
	TemplateName              *string                                                  `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	OperatedServiceInstanceId *string                                                  `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	Service                   *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	OperationStartTime        *string                                                  `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	OperationEndTime          *string                                                  `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	EnableInstanceOps         *bool                                                    `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOutputs(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Outputs = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResources(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Resources = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	Status       *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	PublishTime  *string                                                                `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Version      *string                                                                `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployType   *string                                                                `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId    *string                                                                `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl  *string                                                                `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType  *string                                                                `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName *string                                                                `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListPoliciesRequest struct {
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetMaxResults(v string) *ListPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesRequest) SetNextToken(v string) *ListPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesRequest) SetRegionId(v string) *ListPoliciesRequest {
	s.RegionId = &v
	return s
}

type ListPoliciesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// NextToken
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policies   []*ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	MaxResults *string                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) SetNextToken(v string) *ListPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetMaxResults(v string) *ListPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

type ListPoliciesResponseBodyPolicies struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) SetDescription(v string) *ListPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyName(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyType(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyDocument(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyDocument = &v
	return s
}

type ListPoliciesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) SetHeaders(v map[string]*string) *ListPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesResponse) SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse {
	s.Body = v
	return s
}

type ListInuseServicesRequest struct {
	RegionId   *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults *string                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Filter     []*ListInuseServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s ListInuseServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInuseServicesRequest) GoString() string {
	return s.String()
}

func (s *ListInuseServicesRequest) SetRegionId(v string) *ListInuseServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInuseServicesRequest) SetMaxResults(v string) *ListInuseServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInuseServicesRequest) SetNextToken(v string) *ListInuseServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInuseServicesRequest) SetFilter(v []*ListInuseServicesRequestFilter) *ListInuseServicesRequest {
	s.Filter = v
	return s
}

type ListInuseServicesRequestFilter struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInuseServicesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListInuseServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListInuseServicesRequestFilter) SetValue(v []*string) *ListInuseServicesRequestFilter {
	s.Value = v
	return s
}

func (s *ListInuseServicesRequestFilter) SetName(v string) *ListInuseServicesRequestFilter {
	s.Name = &v
	return s
}

type ListInuseServicesResponseBody struct {
	NextToken  *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults *string                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Services   []*ListInuseServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListInuseServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInuseServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInuseServicesResponseBody) SetNextToken(v string) *ListInuseServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInuseServicesResponseBody) SetRequestId(v string) *ListInuseServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInuseServicesResponseBody) SetTotalCount(v string) *ListInuseServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInuseServicesResponseBody) SetMaxResults(v string) *ListInuseServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInuseServicesResponseBody) SetServices(v []*ListInuseServicesResponseBodyServices) *ListInuseServicesResponseBody {
	s.Services = v
	return s
}

type ListInuseServicesResponseBodyServices struct {
	Status        *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	PublishTime   *string                                              `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Version       *string                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployType    *string                                              `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId     *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl   *string                                              `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType   *string                                              `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName  *string                                              `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos  []*ListInuseServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	CommodityCode *string                                              `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s ListInuseServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListInuseServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListInuseServicesResponseBodyServices) SetStatus(v string) *ListInuseServicesResponseBodyServices {
	s.Status = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetPublishTime(v string) *ListInuseServicesResponseBodyServices {
	s.PublishTime = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetVersion(v string) *ListInuseServicesResponseBodyServices {
	s.Version = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetDeployType(v string) *ListInuseServicesResponseBodyServices {
	s.DeployType = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetServiceId(v string) *ListInuseServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetSupplierUrl(v string) *ListInuseServicesResponseBodyServices {
	s.SupplierUrl = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetServiceType(v string) *ListInuseServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetSupplierName(v string) *ListInuseServicesResponseBodyServices {
	s.SupplierName = &v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetServiceInfos(v []*ListInuseServicesResponseBodyServicesServiceInfos) *ListInuseServicesResponseBodyServices {
	s.ServiceInfos = v
	return s
}

func (s *ListInuseServicesResponseBodyServices) SetCommodityCode(v string) *ListInuseServicesResponseBodyServices {
	s.CommodityCode = &v
	return s
}

type ListInuseServicesResponseBodyServicesServiceInfos struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListInuseServicesResponseBodyServicesServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListInuseServicesResponseBodyServicesServiceInfos) GoString() string {
	return s.String()
}

func (s *ListInuseServicesResponseBodyServicesServiceInfos) SetLocale(v string) *ListInuseServicesResponseBodyServicesServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListInuseServicesResponseBodyServicesServiceInfos) SetImage(v string) *ListInuseServicesResponseBodyServicesServiceInfos {
	s.Image = &v
	return s
}

func (s *ListInuseServicesResponseBodyServicesServiceInfos) SetName(v string) *ListInuseServicesResponseBodyServicesServiceInfos {
	s.Name = &v
	return s
}

func (s *ListInuseServicesResponseBodyServicesServiceInfos) SetShortDescription(v string) *ListInuseServicesResponseBodyServicesServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListInuseServicesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInuseServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInuseServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInuseServicesResponse) GoString() string {
	return s.String()
}

func (s *ListInuseServicesResponse) SetHeaders(v map[string]*string) *ListInuseServicesResponse {
	s.Headers = v
	return s
}

func (s *ListInuseServicesResponse) SetBody(v *ListInuseServicesResponseBody) *ListInuseServicesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type EnableOperationServiceInstanceRequest struct {
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableOperationServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableOperationServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *EnableOperationServiceInstanceRequest) SetServiceInstanceId(v string) *EnableOperationServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *EnableOperationServiceInstanceRequest) SetRegionId(v string) *EnableOperationServiceInstanceRequest {
	s.RegionId = &v
	return s
}

type EnableOperationServiceInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableOperationServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableOperationServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableOperationServiceInstanceResponseBody) SetRequestId(v string) *EnableOperationServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type EnableOperationServiceInstanceResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableOperationServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableOperationServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableOperationServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *EnableOperationServiceInstanceResponse) SetHeaders(v map[string]*string) *EnableOperationServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *EnableOperationServiceInstanceResponse) SetBody(v *EnableOperationServiceInstanceResponseBody) *EnableOperationServiceInstanceResponse {
	s.Body = v
	return s
}

type GetServiceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId      *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetRegionId(v string) *GetServiceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRequest) SetServiceId(v string) *GetServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceVersion(v string) *GetServiceRequest {
	s.ServiceVersion = &v
	return s
}

type GetServiceResponseBody struct {
	Status            *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	DeployMetadata    *string                               `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	PublishTime       *string                               `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	RequestId         *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version           *string                               `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployType        *string                               `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId         *string                               `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl       *string                               `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType       *string                               `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName      *string                               `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos      []*GetServiceResponseBodyServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	CommodityCode     *string                               `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	IsSupportOperated *bool                                 `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	PolicyNames       *string                               `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	Duration          *int64                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetStatus(v string) *GetServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployMetadata(v string) *GetServiceResponseBody {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetPublishTime(v string) *GetServiceResponseBody {
	s.PublishTime = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) SetVersion(v string) *GetServiceResponseBody {
	s.Version = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployType(v string) *GetServiceResponseBody {
	s.DeployType = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceId(v string) *GetServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierUrl(v string) *GetServiceResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceType(v string) *GetServiceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierName(v string) *GetServiceResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceInfos(v []*GetServiceResponseBodyServiceInfos) *GetServiceResponseBody {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceResponseBody) SetCommodityCode(v string) *GetServiceResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceResponseBody) SetIsSupportOperated(v bool) *GetServiceResponseBody {
	s.IsSupportOperated = &v
	return s
}

func (s *GetServiceResponseBody) SetPolicyNames(v string) *GetServiceResponseBody {
	s.PolicyNames = &v
	return s
}

func (s *GetServiceResponseBody) SetDuration(v int64) *GetServiceResponseBody {
	s.Duration = &v
	return s
}

type GetServiceResponseBodyServiceInfos struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceResponseBodyServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfos) SetLocale(v string) *GetServiceResponseBodyServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetImage(v string) *GetServiceResponseBodyServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetName(v string) *GetServiceResponseBodyServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetShortDescription(v string) *GetServiceResponseBodyServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteServiceInstances"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetServiceInstance"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceInstance(request *GetServiceInstanceRequest) (_result *GetServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.GetServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceInstanceLogsWithOptions(request *ListServiceInstanceLogsRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServiceInstanceLogs"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceInstanceLogs(request *ListServiceInstanceLogsRequest) (_result *ListServiceInstanceLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.ListServiceInstanceLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceInstance"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (_result *CreateServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CreateServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ContinueDeployServiceInstance"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContinueDeployServiceInstance(request *ContinueDeployServiceInstanceRequest) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.ContinueDeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableOperationServiceInstanceWithOptions(request *DisableOperationServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *DisableOperationServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableOperationServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableOperationServiceInstance"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableOperationServiceInstance(request *DisableOperationServiceInstanceRequest) (_result *DisableOperationServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableOperationServiceInstanceResponse{}
	_body, _err := client.DisableOperationServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployServiceInstanceWithOptions(request *DeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *DeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployServiceInstance"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployServiceInstance(request *DeployServiceInstanceRequest) (_result *DeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.DeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServiceInstances"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceInstances(request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPolicies"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInuseServicesWithOptions(request *ListInuseServicesRequest, runtime *util.RuntimeOptions) (_result *ListInuseServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInuseServicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInuseServices"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInuseServices(request *ListInuseServicesRequest) (_result *ListInuseServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInuseServicesResponse{}
	_body, _err := client.ListInuseServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableOperationServiceInstanceWithOptions(request *EnableOperationServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *EnableOperationServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableOperationServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableOperationServiceInstance"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableOperationServiceInstance(request *EnableOperationServiceInstanceRequest) (_result *EnableOperationServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableOperationServiceInstanceResponse{}
	_body, _err := client.EnableOperationServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetService"), tea.String("2021-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetService(request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
