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

type ContinueDeployServiceInstanceRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Parameters        *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
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

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ContinueDeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ContinueDeployServiceInstanceResponse) SetStatusCode(v int32) *ContinueDeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetBody(v *ContinueDeployServiceInstanceResponseBody) *ContinueDeployServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceInstanceRequest struct {
	ClientToken       *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContactGroup      *string                                        `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	DryRun            *bool                                          `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EnableInstanceOps *bool                                          `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	Name              *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMetadata *CreateServiceInstanceRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	Parameters        map[string]interface{}                         `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PayType           *int64                                         `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId          *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceId         *string                                        `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion    *string                                        `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationCode *string                                        `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	SpecificationName *string                                        `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	Tag               []*CreateServiceInstanceRequestTag             `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TemplateName      *string                                        `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TrialType         *string                                        `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetContactGroup(v string) *CreateServiceInstanceRequest {
	s.ContactGroup = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetDryRun(v bool) *CreateServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetName(v string) *CreateServiceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetOperationMetadata(v *CreateServiceInstanceRequestOperationMetadata) *CreateServiceInstanceRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetPayType(v int64) *CreateServiceInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
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

func (s *CreateServiceInstanceRequest) SetSpecificationCode(v string) *CreateServiceInstanceRequest {
	s.SpecificationCode = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationName(v string) *CreateServiceInstanceRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTrialType(v string) *CreateServiceInstanceRequest {
	s.TrialType = &v
	return s
}

type CreateServiceInstanceRequestOperationMetadata struct {
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Resources         *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestOperationMetadata) GoString() string {
	return s.String()
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

func (s *CreateServiceInstanceRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.StartTime = &v
	return s
}

type CreateServiceInstanceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceShrinkRequest struct {
	ClientToken       *string                                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContactGroup      *string                                              `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	DryRun            *bool                                                `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EnableInstanceOps *bool                                                `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	Name              *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMetadata *CreateServiceInstanceShrinkRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	ParametersShrink  *string                                              `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PayType           *int64                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId          *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceId         *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion    *string                                              `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationCode *string                                              `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	SpecificationName *string                                              `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	Tag               []*CreateServiceInstanceShrinkRequestTag             `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TemplateName      *string                                              `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TrialType         *string                                              `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetContactGroup(v string) *CreateServiceInstanceShrinkRequest {
	s.ContactGroup = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetDryRun(v bool) *CreateServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetName(v string) *CreateServiceInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetOperationMetadata(v *CreateServiceInstanceShrinkRequestOperationMetadata) *CreateServiceInstanceShrinkRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetPayType(v int64) *CreateServiceInstanceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
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

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationCode(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationCode = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTrialType(v string) *CreateServiceInstanceShrinkRequest {
	s.TrialType = &v
	return s
}

type CreateServiceInstanceShrinkRequestOperationMetadata struct {
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Resources         *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) GoString() string {
	return s.String()
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

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.StartTime = &v
	return s
}

type CreateServiceInstanceShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServiceInstanceResponse) SetStatusCode(v int32) *CreateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	CreateTime                *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableInstanceOps         *bool                                        `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	EndTime                   *string                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsOperated                *bool                                        `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	LicenseEndTime            *string                                      `json:"LicenseEndTime,omitempty" xml:"LicenseEndTime,omitempty"`
	Name                      *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkConfig             *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	OperatedServiceInstanceId *string                                      `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	OperationEndTime          *string                                      `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OperationStartTime        *string                                      `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	Outputs                   *string                                      `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Parameters                *string                                      `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PayType                   *string                                      `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Progress                  *int64                                       `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId                 *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources                 *string                                      `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service                   *GetServiceInstanceResponseBodyService       `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	ServiceInstanceId         *string                                      `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	ServiceType               *string                                      `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string                                      `json:"Source,omitempty" xml:"Source,omitempty"`
	Status                    *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDetail              *string                                      `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	SupplierUid               *int64                                       `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	Tags                      []*GetServiceInstanceResponseBodyTags        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName              *string                                      `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UpdateTime                *string                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                    *int64                                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseEndTime(v string) *GetServiceInstanceResponseBody {
	s.LicenseEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
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

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
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

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	EndpointId                   *string                                                                    `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PrivateVpcConnections        []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections        `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	PrivateZoneId                *string                                                                    `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	EndpointId    *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PrivateZoneId *string `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	DeployMetadata            *string                                              `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	DeployType                *string                                              `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	PublishTime               *string                                              `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	ServiceDocUrl             *string                                              `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	ServiceId                 *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos              []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	ServiceProductUrl         *string                                              `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	ServiceType               *string                                              `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status                    *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName              *string                                              `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl               *string                                              `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	UpgradableServiceVersions []*string                                            `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	UpgradeMetadata           *string                                              `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	Version                   *string                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName               *string                                              `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceVersions(v []*string) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceVersions = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradeMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.UpgradeMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
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

type GetServiceInstanceResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type ListServiceInstanceLogsRequest struct {
	MaxResults        *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequest) SetMaxResults(v string) *ListServiceInstanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetNextToken(v string) *ListServiceInstanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetRegionId(v string) *ListServiceInstanceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest {
	s.ServiceInstanceId = &v
	return s
}

type ListServiceInstanceLogsResponseBody struct {
	MaxResults           *string                                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstancesLogs []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs `json:"ServiceInstancesLogs,omitempty" xml:"ServiceInstancesLogs,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBody) SetMaxResults(v string) *ListServiceInstanceLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetNextToken(v string) *ListServiceInstanceLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetRequestId(v string) *ListServiceInstanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetServiceInstancesLogs(v []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs) *ListServiceInstanceLogsResponseBody {
	s.ServiceInstancesLogs = v
	return s
}

type ListServiceInstanceLogsResponseBodyServiceInstancesLogs struct {
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	LogType           *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	Source            *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Timestamp         *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetContent(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Content = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetLogType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.LogType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceId(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceType = &v
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

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetStatus(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetTimestamp(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Timestamp = &v
	return s
}

type ListServiceInstanceLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceInstanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListServiceInstanceLogsResponse) SetStatusCode(v int32) *ListServiceInstanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse {
	s.Body = v
	return s
}

type ListServiceInstanceResourcesRequest struct {
	ExpireTimeEnd     *string                                   `json:"ExpireTimeEnd,omitempty" xml:"ExpireTimeEnd,omitempty"`
	ExpireTimeStart   *string                                   `json:"ExpireTimeStart,omitempty" xml:"ExpireTimeStart,omitempty"`
	MaxResults        *string                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PayType           *string                                   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ResourceARN       []*string                                 `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ServiceInstanceId *string                                   `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	Tag               []*ListServiceInstanceResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequest) SetExpireTimeEnd(v string) *ListServiceInstanceResourcesRequest {
	s.ExpireTimeEnd = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetExpireTimeStart(v string) *ListServiceInstanceResourcesRequest {
	s.ExpireTimeStart = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetMaxResults(v string) *ListServiceInstanceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetNextToken(v string) *ListServiceInstanceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetPayType(v string) *ListServiceInstanceResourcesRequest {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetResourceARN(v []*string) *ListServiceInstanceResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetTag(v []*ListServiceInstanceResourcesRequestTag) *ListServiceInstanceResourcesRequest {
	s.Tag = v
	return s
}

type ListServiceInstanceResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstanceResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestTag) SetKey(v string) *ListServiceInstanceResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestTag) SetValue(v string) *ListServiceInstanceResourcesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstanceResourcesResponseBody struct {
	MaxResults *string                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*ListServiceInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBody) SetMaxResults(v string) *ListServiceInstanceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetNextToken(v string) *ListServiceInstanceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetRequestId(v string) *ListServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetResources(v []*ListServiceInstanceResourcesResponseBodyResources) *ListServiceInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

type ListServiceInstanceResourcesResponseBodyResources struct {
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType           *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RenewStatus       *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	RenewalPeriod     *int32  `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	ResourceARN       *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
}

func (s ListServiceInstanceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetCreateTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetExpireTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetPayType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductCode(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewStatus = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriod(v int32) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriod = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriodUnit(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetResourceARN(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ResourceARN = &v
	return s
}

type ListServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetStatusCode(v int32) *ListServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetBody(v *ListServiceInstanceResourcesResponseBody) *ListServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	Filter     []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	MaxResults *string                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag        []*ListServiceInstancesRequestTag    `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
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

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	MaxResults       *string                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	TotalCount       *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v string) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int64) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	CreateTime                *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableInstanceOps         *bool                                                    `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	EndTime                   *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MarketInstanceId          *string                                                  `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	Name                      *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OperatedServiceInstanceId *string                                                  `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	OperationEndTime          *string                                                  `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OperationStartTime        *string                                                  `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	Outputs                   *string                                                  `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Parameters                *string                                                  `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PayType                   *string                                                  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Progress                  *int64                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Resources                 *string                                                  `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service                   *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	ServiceInstanceId         *string                                                  `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	ServiceType               *string                                                  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string                                                  `json:"Source,omitempty" xml:"Source,omitempty"`
	Status                    *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDetail              *string                                                  `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Tags                      []*ListServiceInstancesResponseBodyServiceInstancesTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName              *string                                                  `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UpdateTime                *string                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetMarketInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.MarketInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOutputs(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Outputs = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
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

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	DeployType   *string                                                                `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	PublishTime  *string                                                                `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	ServiceId    *string                                                                `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	ServiceType  *string                                                                `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status       *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string                                                                `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl  *string                                                                `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	Version      *string                                                                `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName  *string                                                                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
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

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
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

func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueDeployServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroup)) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInstanceOps)) {
		query["EnableInstanceOps"] = request.EnableInstanceOps
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationCode)) {
		query["SpecificationCode"] = request.SpecificationCode
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceLogs"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListServiceInstanceResourcesWithOptions(request *ListServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTimeEnd)) {
		query["ExpireTimeEnd"] = request.ExpireTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTimeStart)) {
		query["ExpireTimeStart"] = request.ExpireTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceARN)) {
		query["ResourceARN"] = request.ResourceARN
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceInstanceResources(request *ListServiceInstanceResourcesRequest) (_result *ListServiceInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.ListServiceInstanceResourcesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
