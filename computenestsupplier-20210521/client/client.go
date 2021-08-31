// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelServiceRegistrationRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CancelServiceRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceRegistrationRequest) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationRequest) SetRegionId(v string) *CancelServiceRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *CancelServiceRegistrationRequest) SetRegistrationId(v string) *CancelServiceRegistrationRequest {
	s.RegistrationId = &v
	return s
}

func (s *CancelServiceRegistrationRequest) SetClientToken(v string) *CancelServiceRegistrationRequest {
	s.ClientToken = &v
	return s
}

type CancelServiceRegistrationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelServiceRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationResponseBody) SetRequestId(v string) *CancelServiceRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type CancelServiceRegistrationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelServiceRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelServiceRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceRegistrationResponse) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationResponse) SetHeaders(v map[string]*string) *CancelServiceRegistrationResponse {
	s.Headers = v
	return s
}

func (s *CancelServiceRegistrationResponse) SetBody(v *CancelServiceRegistrationResponseBody) *CancelServiceRegistrationResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	RegionId       *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupplierName   *string                            `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl    *string                            `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ClientToken    *string                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ServiceId      *string                            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	DeployType     *string                            `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DeployMetadata *string                            `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	ServiceType    *string                            `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	ServiceInfo    []*CreateServiceRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetRegionId(v string) *CreateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceRequest) SetSupplierName(v string) *CreateServiceRequest {
	s.SupplierName = &v
	return s
}

func (s *CreateServiceRequest) SetSupplierUrl(v string) *CreateServiceRequest {
	s.SupplierUrl = &v
	return s
}

func (s *CreateServiceRequest) SetClientToken(v string) *CreateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceRequest) SetServiceId(v string) *CreateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceRequest) SetDeployType(v string) *CreateServiceRequest {
	s.DeployType = &v
	return s
}

func (s *CreateServiceRequest) SetDeployMetadata(v string) *CreateServiceRequest {
	s.DeployMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetServiceType(v string) *CreateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateServiceRequest) SetServiceInfo(v []*CreateServiceRequestServiceInfo) *CreateServiceRequest {
	s.ServiceInfo = v
	return s
}

type CreateServiceRequestServiceInfo struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateServiceRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfo) SetLocale(v string) *CreateServiceRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetShortDescription(v string) *CreateServiceRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetImage(v string) *CreateServiceRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetName(v string) *CreateServiceRequestServiceInfo {
	s.Name = &v
	return s
}

type CreateServiceResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetVersion(v string) *CreateServiceResponseBody {
	s.Version = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

type CreateServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId      *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetRegionId(v string) *DeleteServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceId(v string) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceVersion(v string) *DeleteServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

type DeleteServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
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
	Status         *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	DeployMetadata *string                               `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	PublishTime    *string                               `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version        *string                               `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployType     *string                               `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId      *string                               `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl    *string                               `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType    *string                               `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName   *string                               `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos   []*GetServiceResponseBodyServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	CommodityCode  *string                               `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
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
	Status            *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime        *string                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	RequestId         *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstanceId *string                                `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	CreateTime        *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId            *int64                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Service           *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	Parameters        *string                                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	StatusDetail      *string                                `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Progress          *int64                                 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	TemplateName      *string                                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
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

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	Status         *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	PublishTime    *string                                              `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Version        *string                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployType     *string                                              `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId      *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl    *string                                              `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType    *string                                              `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName   *string                                              `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos   []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	DeployMetadata *string                                              `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
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

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
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

type GetSupplierInformationRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupplierInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationRequest) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationRequest) SetRegionId(v string) *GetSupplierInformationRequest {
	s.RegionId = &v
	return s
}

type GetSupplierInformationResponseBody struct {
	// Id of the request
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl  *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
}

func (s GetSupplierInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBody) SetRequestId(v string) *GetSupplierInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierName(v string) *GetSupplierInformationResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierUrl(v string) *GetSupplierInformationResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierDesc(v string) *GetSupplierInformationResponseBody {
	s.SupplierDesc = &v
	return s
}

type GetSupplierInformationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSupplierInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSupplierInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationResponse) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponse) SetHeaders(v map[string]*string) *GetSupplierInformationResponse {
	s.Headers = v
	return s
}

func (s *GetSupplierInformationResponse) SetBody(v *GetSupplierInformationResponseBody) *GetSupplierInformationResponse {
	s.Body = v
	return s
}

type LaunchServiceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId      *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s LaunchServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s LaunchServiceRequest) GoString() string {
	return s.String()
}

func (s *LaunchServiceRequest) SetRegionId(v string) *LaunchServiceRequest {
	s.RegionId = &v
	return s
}

func (s *LaunchServiceRequest) SetServiceId(v string) *LaunchServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *LaunchServiceRequest) SetServiceVersion(v string) *LaunchServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *LaunchServiceRequest) SetClientToken(v string) *LaunchServiceRequest {
	s.ClientToken = &v
	return s
}

type LaunchServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s LaunchServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LaunchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchServiceResponseBody) SetRequestId(v string) *LaunchServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchServiceResponseBody) SetServiceId(v string) *LaunchServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *LaunchServiceResponseBody) SetVersion(v string) *LaunchServiceResponseBody {
	s.Version = &v
	return s
}

func (s *LaunchServiceResponseBody) SetStatus(v string) *LaunchServiceResponseBody {
	s.Status = &v
	return s
}

type LaunchServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LaunchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LaunchServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s LaunchServiceResponse) GoString() string {
	return s.String()
}

func (s *LaunchServiceResponse) SetHeaders(v map[string]*string) *LaunchServiceResponse {
	s.Headers = v
	return s
}

func (s *LaunchServiceResponse) SetBody(v *LaunchServiceResponseBody) *LaunchServiceResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	RegionId   *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults *string                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Filter     []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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

type ListServiceInstancesResponseBody struct {
	NextToken        *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
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
	Status            *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime        *string                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ServiceInstanceId *string                                                  `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	CreateTime        *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId            *int64                                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Service           *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	Parameters        *string                                                  `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	StatusDetail      *string                                                  `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Progress          *int64                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	TemplateName      *string                                                  `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
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

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUserId(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.UserId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
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

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
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

type ListServiceRegistrationsRequest struct {
	RegionId   *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults *string                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Filter     []*ListServiceRegistrationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s ListServiceRegistrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsRequest) SetRegionId(v string) *ListServiceRegistrationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetMaxResults(v string) *ListServiceRegistrationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetNextToken(v string) *ListServiceRegistrationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetFilter(v []*ListServiceRegistrationsRequestFilter) *ListServiceRegistrationsRequest {
	s.Filter = v
	return s
}

type ListServiceRegistrationsRequestFilter struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListServiceRegistrationsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsRequestFilter) SetValue(v []*string) *ListServiceRegistrationsRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceRegistrationsRequestFilter) SetName(v string) *ListServiceRegistrationsRequestFilter {
	s.Name = &v
	return s
}

type ListServiceRegistrationsResponseBody struct {
	NextToken            *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *string                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults           *int32                                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ServiceRegistrations []*ListServiceRegistrationsResponseBodyServiceRegistrations `json:"ServiceRegistrations,omitempty" xml:"ServiceRegistrations,omitempty" type:"Repeated"`
}

func (s ListServiceRegistrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponseBody) SetNextToken(v string) *ListServiceRegistrationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetRequestId(v string) *ListServiceRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetTotalCount(v string) *ListServiceRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetMaxResults(v int32) *ListServiceRegistrationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetServiceRegistrations(v []*ListServiceRegistrationsResponseBodyServiceRegistrations) *ListServiceRegistrationsResponseBody {
	s.ServiceRegistrations = v
	return s
}

type ListServiceRegistrationsResponseBodyServiceRegistrations struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	FinishTime     *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Comment        *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ServiceId      *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SubmitTime     *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListServiceRegistrationsResponseBodyServiceRegistrations) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsResponseBodyServiceRegistrations) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetStatus(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.Status = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetRegistrationId(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetFinishTime(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.FinishTime = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetComment(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.Comment = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetServiceId(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.ServiceId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetSubmitTime(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.SubmitTime = &v
	return s
}

type ListServiceRegistrationsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceRegistrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponse) SetHeaders(v map[string]*string) *ListServiceRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceRegistrationsResponse) SetBody(v *ListServiceRegistrationsResponseBody) *ListServiceRegistrationsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	RegionId    *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults  *string                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	AllVersions *bool                        `json:"AllVersions,omitempty" xml:"AllVersions,omitempty"`
	Filter      []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetMaxResults(v string) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetAllVersions(v bool) *ListServicesRequest {
	s.AllVersions = &v
	return s
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

type ListServicesRequestFilter struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListServicesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServicesRequestFilter) SetValue(v []*string) *ListServicesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServicesRequestFilter) SetName(v string) *ListServicesRequestFilter {
	s.Name = &v
	return s
}

type ListServicesResponseBody struct {
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Services   []*ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetNextToken(v string) *ListServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v string) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServicesResponseBody) SetMaxResults(v int32) *ListServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

type ListServicesResponseBodyServices struct {
	Status         *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	DefaultVersion *bool                                           `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	PublishTime    *string                                         `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Version        *string                                         `json:"Version,omitempty" xml:"Version,omitempty"`
	DeployType     *string                                         `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	ServiceId      *string                                         `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	SupplierUrl    *string                                         `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	ServiceType    *string                                         `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupplierName   *string                                         `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	ServiceInfos   []*ListServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	CommodityCode  *string                                         `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetStatus(v string) *ListServicesResponseBodyServices {
	s.Status = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDefaultVersion(v bool) *ListServicesResponseBodyServices {
	s.DefaultVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetPublishTime(v string) *ListServicesResponseBodyServices {
	s.PublishTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersion(v string) *ListServicesResponseBodyServices {
	s.Version = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDeployType(v string) *ListServicesResponseBodyServices {
	s.DeployType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceId(v string) *ListServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierUrl(v string) *ListServicesResponseBodyServices {
	s.SupplierUrl = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceType(v string) *ListServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierName(v string) *ListServicesResponseBodyServices {
	s.SupplierName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceInfos(v []*ListServicesResponseBodyServicesServiceInfos) *ListServicesResponseBodyServices {
	s.ServiceInfos = v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodityCode(v string) *ListServicesResponseBodyServices {
	s.CommodityCode = &v
	return s
}

type ListServicesResponseBodyServicesServiceInfos struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServicesResponseBodyServicesServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetLocale(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetImage(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetName(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetShortDescription(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServicesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type RegisterServiceRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RegisterServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceRequest) GoString() string {
	return s.String()
}

func (s *RegisterServiceRequest) SetRegionId(v string) *RegisterServiceRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterServiceRequest) SetServiceId(v string) *RegisterServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *RegisterServiceRequest) SetClientToken(v string) *RegisterServiceRequest {
	s.ClientToken = &v
	return s
}

type RegisterServiceResponseBody struct {
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponseBody) SetRegistrationId(v string) *RegisterServiceResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *RegisterServiceResponseBody) SetRequestId(v string) *RegisterServiceResponseBody {
	s.RequestId = &v
	return s
}

type RegisterServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceResponse) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponse) SetHeaders(v map[string]*string) *RegisterServiceResponse {
	s.Headers = v
	return s
}

func (s *RegisterServiceResponse) SetBody(v *RegisterServiceResponseBody) *RegisterServiceResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	RegionId       *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupplierName   *string                            `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl    *string                            `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	DeployType     *string                            `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DeployMetadata *string                            `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	ClientToken    *string                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ServiceId      *string                            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceType    *string                            `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	ServiceInfo    []*UpdateServiceRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetRegionId(v string) *UpdateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceRequest) SetSupplierName(v string) *UpdateServiceRequest {
	s.SupplierName = &v
	return s
}

func (s *UpdateServiceRequest) SetSupplierUrl(v string) *UpdateServiceRequest {
	s.SupplierUrl = &v
	return s
}

func (s *UpdateServiceRequest) SetDeployType(v string) *UpdateServiceRequest {
	s.DeployType = &v
	return s
}

func (s *UpdateServiceRequest) SetDeployMetadata(v string) *UpdateServiceRequest {
	s.DeployMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetClientToken(v string) *UpdateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceId(v string) *UpdateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceType(v string) *UpdateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceInfo(v []*UpdateServiceRequestServiceInfo) *UpdateServiceRequest {
	s.ServiceInfo = v
	return s
}

type UpdateServiceRequestServiceInfo struct {
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateServiceRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfo) SetLocale(v string) *UpdateServiceRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetShortDescription(v string) *UpdateServiceRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetImage(v string) *UpdateServiceRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetName(v string) *UpdateServiceRequestServiceInfo {
	s.Name = &v
	return s
}

type UpdateServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type WithdrawServiceRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId      *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s WithdrawServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawServiceRequest) GoString() string {
	return s.String()
}

func (s *WithdrawServiceRequest) SetRegionId(v string) *WithdrawServiceRequest {
	s.RegionId = &v
	return s
}

func (s *WithdrawServiceRequest) SetServiceId(v string) *WithdrawServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *WithdrawServiceRequest) SetServiceVersion(v string) *WithdrawServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *WithdrawServiceRequest) SetClientToken(v string) *WithdrawServiceRequest {
	s.ClientToken = &v
	return s
}

type WithdrawServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WithdrawServiceResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawServiceResponseBody) SetRequestId(v string) *WithdrawServiceResponseBody {
	s.RequestId = &v
	return s
}

type WithdrawServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WithdrawServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WithdrawServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s WithdrawServiceResponse) GoString() string {
	return s.String()
}

func (s *WithdrawServiceResponse) SetHeaders(v map[string]*string) *WithdrawServiceResponse {
	s.Headers = v
	return s
}

func (s *WithdrawServiceResponse) SetBody(v *WithdrawServiceResponseBody) *WithdrawServiceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenestsupplier"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelServiceRegistrationWithOptions(request *CancelServiceRegistrationRequest, runtime *util.RuntimeOptions) (_result *CancelServiceRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelServiceRegistrationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelServiceRegistration"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelServiceRegistration(request *CancelServiceRegistrationRequest) (_result *CancelServiceRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelServiceRegistrationResponse{}
	_body, _err := client.CancelServiceRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("GetService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetServiceInstance"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetSupplierInformationWithOptions(request *GetSupplierInformationRequest, runtime *util.RuntimeOptions) (_result *GetSupplierInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSupplierInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSupplierInformation"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSupplierInformation(request *GetSupplierInformationRequest) (_result *GetSupplierInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSupplierInformationResponse{}
	_body, _err := client.GetSupplierInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LaunchServiceWithOptions(request *LaunchServiceRequest, runtime *util.RuntimeOptions) (_result *LaunchServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LaunchServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LaunchService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LaunchService(request *LaunchServiceRequest) (_result *LaunchServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LaunchServiceResponse{}
	_body, _err := client.LaunchServiceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListServiceInstances"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListServiceRegistrationsWithOptions(request *ListServiceRegistrationsRequest, runtime *util.RuntimeOptions) (_result *ListServiceRegistrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServiceRegistrationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServiceRegistrations"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceRegistrations(request *ListServiceRegistrationsRequest) (_result *ListServiceRegistrationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceRegistrationsResponse{}
	_body, _err := client.ListServiceRegistrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServices"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterServiceWithOptions(request *RegisterServiceRequest, runtime *util.RuntimeOptions) (_result *RegisterServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterService(request *RegisterServiceRequest) (_result *RegisterServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterServiceResponse{}
	_body, _err := client.RegisterServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceWithOptions(request *UpdateServiceRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateService(request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WithdrawServiceWithOptions(request *WithdrawServiceRequest, runtime *util.RuntimeOptions) (_result *WithdrawServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WithdrawServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WithdrawService"), tea.String("2021-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WithdrawService(request *WithdrawServiceRequest) (_result *WithdrawServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WithdrawServiceResponse{}
	_body, _err := client.WithdrawServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
