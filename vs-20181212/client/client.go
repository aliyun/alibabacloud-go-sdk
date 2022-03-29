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

type AddRegisteredDeviceRequest struct {
	Dsn          *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	SecretKey    *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s AddRegisteredDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRegisteredDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddRegisteredDeviceRequest) SetDsn(v string) *AddRegisteredDeviceRequest {
	s.Dsn = &v
	return s
}

func (s *AddRegisteredDeviceRequest) SetOwnerId(v int64) *AddRegisteredDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *AddRegisteredDeviceRequest) SetRegisterCode(v string) *AddRegisteredDeviceRequest {
	s.RegisterCode = &v
	return s
}

func (s *AddRegisteredDeviceRequest) SetSecretKey(v string) *AddRegisteredDeviceRequest {
	s.SecretKey = &v
	return s
}

func (s *AddRegisteredDeviceRequest) SetVendor(v string) *AddRegisteredDeviceRequest {
	s.Vendor = &v
	return s
}

type AddRegisteredDeviceResponseBody struct {
	Dsn          *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretKey    *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s AddRegisteredDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRegisteredDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddRegisteredDeviceResponseBody) SetDsn(v string) *AddRegisteredDeviceResponseBody {
	s.Dsn = &v
	return s
}

func (s *AddRegisteredDeviceResponseBody) SetId(v string) *AddRegisteredDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *AddRegisteredDeviceResponseBody) SetRegisterCode(v string) *AddRegisteredDeviceResponseBody {
	s.RegisterCode = &v
	return s
}

func (s *AddRegisteredDeviceResponseBody) SetRequestId(v string) *AddRegisteredDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRegisteredDeviceResponseBody) SetSecretKey(v string) *AddRegisteredDeviceResponseBody {
	s.SecretKey = &v
	return s
}

func (s *AddRegisteredDeviceResponseBody) SetVendor(v string) *AddRegisteredDeviceResponseBody {
	s.Vendor = &v
	return s
}

type AddRegisteredDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRegisteredDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRegisteredDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRegisteredDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddRegisteredDeviceResponse) SetHeaders(v map[string]*string) *AddRegisteredDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddRegisteredDeviceResponse) SetBody(v *AddRegisteredDeviceResponseBody) *AddRegisteredDeviceResponse {
	s.Body = v
	return s
}

type AddRegisteredVendorRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s AddRegisteredVendorRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRegisteredVendorRequest) GoString() string {
	return s.String()
}

func (s *AddRegisteredVendorRequest) SetDescription(v string) *AddRegisteredVendorRequest {
	s.Description = &v
	return s
}

func (s *AddRegisteredVendorRequest) SetName(v string) *AddRegisteredVendorRequest {
	s.Name = &v
	return s
}

func (s *AddRegisteredVendorRequest) SetOwnerId(v int64) *AddRegisteredVendorRequest {
	s.OwnerId = &v
	return s
}

type AddRegisteredVendorResponseBody struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Vendor      *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s AddRegisteredVendorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRegisteredVendorResponseBody) GoString() string {
	return s.String()
}

func (s *AddRegisteredVendorResponseBody) SetDescription(v string) *AddRegisteredVendorResponseBody {
	s.Description = &v
	return s
}

func (s *AddRegisteredVendorResponseBody) SetName(v string) *AddRegisteredVendorResponseBody {
	s.Name = &v
	return s
}

func (s *AddRegisteredVendorResponseBody) SetRequestId(v string) *AddRegisteredVendorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRegisteredVendorResponseBody) SetVendor(v string) *AddRegisteredVendorResponseBody {
	s.Vendor = &v
	return s
}

type AddRegisteredVendorResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRegisteredVendorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRegisteredVendorResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRegisteredVendorResponse) GoString() string {
	return s.String()
}

func (s *AddRegisteredVendorResponse) SetHeaders(v map[string]*string) *AddRegisteredVendorResponse {
	s.Headers = v
	return s
}

func (s *AddRegisteredVendorResponse) SetBody(v *AddRegisteredVendorResponseBody) *AddRegisteredVendorResponse {
	s.Body = v
	return s
}

type AddRenderingDeviceInternetPortsRequest struct {
	ISP          *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s AddRenderingDeviceInternetPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRenderingDeviceInternetPortsRequest) GoString() string {
	return s.String()
}

func (s *AddRenderingDeviceInternetPortsRequest) SetISP(v string) *AddRenderingDeviceInternetPortsRequest {
	s.ISP = &v
	return s
}

func (s *AddRenderingDeviceInternetPortsRequest) SetInstanceIds(v string) *AddRenderingDeviceInternetPortsRequest {
	s.InstanceIds = &v
	return s
}

func (s *AddRenderingDeviceInternetPortsRequest) SetInternalPort(v string) *AddRenderingDeviceInternetPortsRequest {
	s.InternalPort = &v
	return s
}

func (s *AddRenderingDeviceInternetPortsRequest) SetIpProtocol(v string) *AddRenderingDeviceInternetPortsRequest {
	s.IpProtocol = &v
	return s
}

func (s *AddRenderingDeviceInternetPortsRequest) SetOwnerId(v int64) *AddRenderingDeviceInternetPortsRequest {
	s.OwnerId = &v
	return s
}

type AddRenderingDeviceInternetPortsResponseBody struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRenderingDeviceInternetPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRenderingDeviceInternetPortsResponseBody) GoString() string {
	return s.String()
}

func (s *AddRenderingDeviceInternetPortsResponseBody) SetInstanceIds(v []*string) *AddRenderingDeviceInternetPortsResponseBody {
	s.InstanceIds = v
	return s
}

func (s *AddRenderingDeviceInternetPortsResponseBody) SetRequestId(v string) *AddRenderingDeviceInternetPortsResponseBody {
	s.RequestId = &v
	return s
}

type AddRenderingDeviceInternetPortsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRenderingDeviceInternetPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRenderingDeviceInternetPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRenderingDeviceInternetPortsResponse) GoString() string {
	return s.String()
}

func (s *AddRenderingDeviceInternetPortsResponse) SetHeaders(v map[string]*string) *AddRenderingDeviceInternetPortsResponse {
	s.Headers = v
	return s
}

func (s *AddRenderingDeviceInternetPortsResponse) SetBody(v *AddRenderingDeviceInternetPortsResponseBody) *AddRenderingDeviceInternetPortsResponse {
	s.Body = v
	return s
}

type AddVsPullStreamInfoConfigRequest struct {
	Always     *string `json:"Always,omitempty" xml:"Always,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s AddVsPullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *AddVsPullStreamInfoConfigRequest) SetAlways(v string) *AddVsPullStreamInfoConfigRequest {
	s.Always = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetAppName(v string) *AddVsPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetDomainName(v string) *AddVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetEndTime(v string) *AddVsPullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *AddVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetSourceUrl(v string) *AddVsPullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetStartTime(v string) *AddVsPullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *AddVsPullStreamInfoConfigRequest) SetStreamName(v string) *AddVsPullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

type AddVsPullStreamInfoConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVsPullStreamInfoConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *AddVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

type AddVsPullStreamInfoConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVsPullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *AddVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *AddVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *AddVsPullStreamInfoConfigResponse) SetBody(v *AddVsPullStreamInfoConfigResponseBody) *AddVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

type BatchBindDirectoriesRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchBindDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBindDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesRequest) SetDeviceId(v string) *BatchBindDirectoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchBindDirectoriesRequest) SetDirectoryId(v string) *BatchBindDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *BatchBindDirectoriesRequest) SetOwnerId(v int64) *BatchBindDirectoriesRequest {
	s.OwnerId = &v
	return s
}

type BatchBindDirectoriesResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindDirectoriesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBindDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesResponseBody) SetRequestId(v string) *BatchBindDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindDirectoriesResponseBody) SetResults(v []*BatchBindDirectoriesResponseBodyResults) *BatchBindDirectoriesResponseBody {
	s.Results = v
	return s
}

type BatchBindDirectoriesResponseBodyResults struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Error       *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s BatchBindDirectoriesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchBindDirectoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesResponseBodyResults) SetDeviceId(v string) *BatchBindDirectoriesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchBindDirectoriesResponseBodyResults) SetDirectoryId(v string) *BatchBindDirectoriesResponseBodyResults {
	s.DirectoryId = &v
	return s
}

func (s *BatchBindDirectoriesResponseBodyResults) SetError(v string) *BatchBindDirectoriesResponseBodyResults {
	s.Error = &v
	return s
}

type BatchBindDirectoriesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchBindDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchBindDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBindDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesResponse) SetHeaders(v map[string]*string) *BatchBindDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindDirectoriesResponse) SetBody(v *BatchBindDirectoriesResponseBody) *BatchBindDirectoriesResponse {
	s.Body = v
	return s
}

type BatchBindParentPlatformDevicesRequest struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BatchBindParentPlatformDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBindParentPlatformDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesRequest) SetDeviceId(v string) *BatchBindParentPlatformDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesRequest) SetOwnerId(v int64) *BatchBindParentPlatformDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesRequest) SetParentPlatformId(v string) *BatchBindParentPlatformDevicesRequest {
	s.ParentPlatformId = &v
	return s
}

type BatchBindParentPlatformDevicesResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindParentPlatformDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindParentPlatformDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBindParentPlatformDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesResponseBody) SetRequestId(v string) *BatchBindParentPlatformDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBody) SetResults(v []*BatchBindParentPlatformDevicesResponseBodyResults) *BatchBindParentPlatformDevicesResponseBody {
	s.Results = v
	return s
}

type BatchBindParentPlatformDevicesResponseBodyResults struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Error            *string `json:"Error,omitempty" xml:"Error,omitempty"`
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BatchBindParentPlatformDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchBindParentPlatformDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) SetDeviceId(v string) *BatchBindParentPlatformDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) SetError(v string) *BatchBindParentPlatformDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) SetParentPlatformId(v string) *BatchBindParentPlatformDevicesResponseBodyResults {
	s.ParentPlatformId = &v
	return s
}

type BatchBindParentPlatformDevicesResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchBindParentPlatformDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchBindParentPlatformDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBindParentPlatformDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesResponse) SetHeaders(v map[string]*string) *BatchBindParentPlatformDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindParentPlatformDevicesResponse) SetBody(v *BatchBindParentPlatformDevicesResponseBody) *BatchBindParentPlatformDevicesResponse {
	s.Body = v
	return s
}

type BatchBindPurchasedDevicesRequest struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s BatchBindPurchasedDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBindPurchasedDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesRequest) SetDeviceId(v string) *BatchBindPurchasedDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) SetGroupId(v string) *BatchBindPurchasedDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) SetOwnerId(v int64) *BatchBindPurchasedDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindPurchasedDevicesRequest) SetRegion(v string) *BatchBindPurchasedDevicesRequest {
	s.Region = &v
	return s
}

type BatchBindPurchasedDevicesResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindPurchasedDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindPurchasedDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBindPurchasedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesResponseBody) SetRequestId(v string) *BatchBindPurchasedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBody) SetResults(v []*BatchBindPurchasedDevicesResponseBodyResults) *BatchBindPurchasedDevicesResponseBody {
	s.Results = v
	return s
}

type BatchBindPurchasedDevicesResponseBodyResults struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Error    *string `json:"Error,omitempty" xml:"Error,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s BatchBindPurchasedDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchBindPurchasedDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetDeviceId(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetError(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetGroupId(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.GroupId = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetRegion(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.Region = &v
	return s
}

type BatchBindPurchasedDevicesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchBindPurchasedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchBindPurchasedDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBindPurchasedDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesResponse) SetHeaders(v map[string]*string) *BatchBindPurchasedDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindPurchasedDevicesResponse) SetBody(v *BatchBindPurchasedDevicesResponseBody) *BatchBindPurchasedDevicesResponse {
	s.Body = v
	return s
}

type BatchBindTemplateRequest struct {
	ApplyAll     *bool   `json:"ApplyAll,omitempty" xml:"ApplyAll,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Replace      *bool   `json:"Replace,omitempty" xml:"Replace,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s BatchBindTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplateRequest) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateRequest) SetApplyAll(v bool) *BatchBindTemplateRequest {
	s.ApplyAll = &v
	return s
}

func (s *BatchBindTemplateRequest) SetInstanceId(v string) *BatchBindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplateRequest) SetInstanceType(v string) *BatchBindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplateRequest) SetOwnerId(v int64) *BatchBindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindTemplateRequest) SetReplace(v bool) *BatchBindTemplateRequest {
	s.Replace = &v
	return s
}

func (s *BatchBindTemplateRequest) SetTemplateId(v string) *BatchBindTemplateRequest {
	s.TemplateId = &v
	return s
}

type BatchBindTemplateResponseBody struct {
	Bindings  []*BatchBindTemplateResponseBodyBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchBindTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateResponseBody) SetBindings(v []*BatchBindTemplateResponseBodyBindings) *BatchBindTemplateResponseBody {
	s.Bindings = v
	return s
}

func (s *BatchBindTemplateResponseBody) SetRequestId(v string) *BatchBindTemplateResponseBody {
	s.RequestId = &v
	return s
}

type BatchBindTemplateResponseBodyBindings struct {
	Error        *string `json:"Error,omitempty" xml:"Error,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s BatchBindTemplateResponseBodyBindings) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplateResponseBodyBindings) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateResponseBodyBindings) SetError(v string) *BatchBindTemplateResponseBodyBindings {
	s.Error = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) SetInstanceId(v string) *BatchBindTemplateResponseBodyBindings {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) SetInstanceType(v string) *BatchBindTemplateResponseBodyBindings {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) SetTemplateId(v string) *BatchBindTemplateResponseBodyBindings {
	s.TemplateId = &v
	return s
}

type BatchBindTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchBindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchBindTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplateResponse) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateResponse) SetHeaders(v map[string]*string) *BatchBindTemplateResponse {
	s.Headers = v
	return s
}

func (s *BatchBindTemplateResponse) SetBody(v *BatchBindTemplateResponseBody) *BatchBindTemplateResponse {
	s.Body = v
	return s
}

type BatchBindTemplatesRequest struct {
	ApplyAll     *bool   `json:"ApplyAll,omitempty" xml:"ApplyAll,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Replace      *bool   `json:"Replace,omitempty" xml:"Replace,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchBindTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplatesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesRequest) SetApplyAll(v bool) *BatchBindTemplatesRequest {
	s.ApplyAll = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetInstanceId(v string) *BatchBindTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetInstanceType(v string) *BatchBindTemplatesRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetOwnerId(v int64) *BatchBindTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetReplace(v bool) *BatchBindTemplatesRequest {
	s.Replace = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetTemplateId(v string) *BatchBindTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetTemplateType(v string) *BatchBindTemplatesRequest {
	s.TemplateType = &v
	return s
}

type BatchBindTemplatesResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindTemplatesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesResponseBody) SetRequestId(v string) *BatchBindTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindTemplatesResponseBody) SetResults(v []*BatchBindTemplatesResponseBodyResults) *BatchBindTemplatesResponseBody {
	s.Results = v
	return s
}

type BatchBindTemplatesResponseBodyResults struct {
	Error        *string `json:"Error,omitempty" xml:"Error,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s BatchBindTemplatesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplatesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesResponseBodyResults) SetError(v string) *BatchBindTemplatesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) SetInstanceId(v string) *BatchBindTemplatesResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) SetInstanceType(v string) *BatchBindTemplatesResponseBodyResults {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) SetTemplateId(v string) *BatchBindTemplatesResponseBodyResults {
	s.TemplateId = &v
	return s
}

type BatchBindTemplatesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchBindTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchBindTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBindTemplatesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesResponse) SetHeaders(v map[string]*string) *BatchBindTemplatesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindTemplatesResponse) SetBody(v *BatchBindTemplatesResponseBody) *BatchBindTemplatesResponse {
	s.Body = v
	return s
}

type BatchDeleteDevicesRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchDeleteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesRequest) SetId(v string) *BatchDeleteDevicesRequest {
	s.Id = &v
	return s
}

func (s *BatchDeleteDevicesRequest) SetOwnerId(v int64) *BatchDeleteDevicesRequest {
	s.OwnerId = &v
	return s
}

type BatchDeleteDevicesResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchDeleteDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchDeleteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesResponseBody) SetRequestId(v string) *BatchDeleteDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDevicesResponseBody) SetResults(v []*BatchDeleteDevicesResponseBodyResults) *BatchDeleteDevicesResponseBody {
	s.Results = v
	return s
}

type BatchDeleteDevicesResponseBodyResults struct {
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s BatchDeleteDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesResponseBodyResults) SetError(v string) *BatchDeleteDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchDeleteDevicesResponseBodyResults) SetId(v string) *BatchDeleteDevicesResponseBodyResults {
	s.Id = &v
	return s
}

type BatchDeleteDevicesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesResponse) SetHeaders(v map[string]*string) *BatchDeleteDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDevicesResponse) SetBody(v *BatchDeleteDevicesResponseBody) *BatchDeleteDevicesResponse {
	s.Body = v
	return s
}

type BatchDeleteVsDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchDeleteVsDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteVsDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteVsDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteVsDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteVsDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsRequest) SetOwnerId(v int64) *BatchDeleteVsDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

type BatchDeleteVsDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteVsDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteVsDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteVsDomainConfigsResponseBody) SetRequestId(v string) *BatchDeleteVsDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteVsDomainConfigsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteVsDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteVsDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteVsDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteVsDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchDeleteVsDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteVsDomainConfigsResponse) SetBody(v *BatchDeleteVsDomainConfigsResponseBody) *BatchDeleteVsDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchForbidVsStreamRequest struct {
	Channel             *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LiveStreamType      *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	Oneshot             *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResumeTime          *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
}

func (s BatchForbidVsStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchForbidVsStreamRequest) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamRequest) SetChannel(v string) *BatchForbidVsStreamRequest {
	s.Channel = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetControlStreamAction(v string) *BatchForbidVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetDomainName(v string) *BatchForbidVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetLiveStreamType(v string) *BatchForbidVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetOneshot(v string) *BatchForbidVsStreamRequest {
	s.Oneshot = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetOwnerId(v int64) *BatchForbidVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetResumeTime(v string) *BatchForbidVsStreamRequest {
	s.ResumeTime = &v
	return s
}

type BatchForbidVsStreamResponseBody struct {
	ForbidResult *BatchForbidVsStreamResponseBodyForbidResult `json:"ForbidResult,omitempty" xml:"ForbidResult,omitempty" type:"Struct"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchForbidVsStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchForbidVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBody) SetForbidResult(v *BatchForbidVsStreamResponseBodyForbidResult) *BatchForbidVsStreamResponseBody {
	s.ForbidResult = v
	return s
}

func (s *BatchForbidVsStreamResponseBody) SetRequestId(v string) *BatchForbidVsStreamResponseBody {
	s.RequestId = &v
	return s
}

type BatchForbidVsStreamResponseBodyForbidResult struct {
	ForbidResultInfo []*BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo `json:"ForbidResultInfo,omitempty" xml:"ForbidResultInfo,omitempty" type:"Repeated"`
}

func (s BatchForbidVsStreamResponseBodyForbidResult) String() string {
	return tea.Prettify(s)
}

func (s BatchForbidVsStreamResponseBodyForbidResult) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBodyForbidResult) SetForbidResultInfo(v []*BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) *BatchForbidVsStreamResponseBodyForbidResult {
	s.ForbidResultInfo = v
	return s
}

type BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo struct {
	Channels *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	Count    *int32                                                               `json:"Count,omitempty" xml:"Count,omitempty"`
	Detail   *string                                                              `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Result   *string                                                              `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetChannels(v *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Channels = v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetCount(v int32) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Count = &v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetDetail(v string) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Detail = &v
	return s
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo) SetResult(v string) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfo {
	s.Result = &v
	return s
}

type BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) String() string {
	return tea.Prettify(s)
}

func (s BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels) SetChannel(v []*string) *BatchForbidVsStreamResponseBodyForbidResultForbidResultInfoChannels {
	s.Channel = v
	return s
}

type BatchForbidVsStreamResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchForbidVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchForbidVsStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchForbidVsStreamResponse) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponse) SetHeaders(v map[string]*string) *BatchForbidVsStreamResponse {
	s.Headers = v
	return s
}

func (s *BatchForbidVsStreamResponse) SetBody(v *BatchForbidVsStreamResponseBody) *BatchForbidVsStreamResponse {
	s.Body = v
	return s
}

type BatchResumeVsStreamRequest struct {
	Channel             *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LiveStreamType      *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchResumeVsStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchResumeVsStreamRequest) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamRequest) SetChannel(v string) *BatchResumeVsStreamRequest {
	s.Channel = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetControlStreamAction(v string) *BatchResumeVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetDomainName(v string) *BatchResumeVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetLiveStreamType(v string) *BatchResumeVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetOwnerId(v int64) *BatchResumeVsStreamRequest {
	s.OwnerId = &v
	return s
}

type BatchResumeVsStreamResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResumeResult *BatchResumeVsStreamResponseBodyResumeResult `json:"ResumeResult,omitempty" xml:"ResumeResult,omitempty" type:"Struct"`
}

func (s BatchResumeVsStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchResumeVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBody) SetRequestId(v string) *BatchResumeVsStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchResumeVsStreamResponseBody) SetResumeResult(v *BatchResumeVsStreamResponseBodyResumeResult) *BatchResumeVsStreamResponseBody {
	s.ResumeResult = v
	return s
}

type BatchResumeVsStreamResponseBodyResumeResult struct {
	ResumeResultInfo []*BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo `json:"ResumeResultInfo,omitempty" xml:"ResumeResultInfo,omitempty" type:"Repeated"`
}

func (s BatchResumeVsStreamResponseBodyResumeResult) String() string {
	return tea.Prettify(s)
}

func (s BatchResumeVsStreamResponseBodyResumeResult) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBodyResumeResult) SetResumeResultInfo(v []*BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) *BatchResumeVsStreamResponseBodyResumeResult {
	s.ResumeResultInfo = v
	return s
}

type BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo struct {
	Channels *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	Count    *int32                                                               `json:"Count,omitempty" xml:"Count,omitempty"`
	Detail   *string                                                              `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Result   *string                                                              `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetChannels(v *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Channels = v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetCount(v int32) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Count = &v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetDetail(v string) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Detail = &v
	return s
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo) SetResult(v string) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfo {
	s.Result = &v
	return s
}

type BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels struct {
	Channel []*string `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) String() string {
	return tea.Prettify(s)
}

func (s BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels) SetChannel(v []*string) *BatchResumeVsStreamResponseBodyResumeResultResumeResultInfoChannels {
	s.Channel = v
	return s
}

type BatchResumeVsStreamResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchResumeVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchResumeVsStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchResumeVsStreamResponse) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponse) SetHeaders(v map[string]*string) *BatchResumeVsStreamResponse {
	s.Headers = v
	return s
}

func (s *BatchResumeVsStreamResponse) SetBody(v *BatchResumeVsStreamResponseBody) *BatchResumeVsStreamResponse {
	s.Body = v
	return s
}

type BatchSetVsDomainConfigsRequest struct {
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Functions   *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchSetVsDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetVsDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetVsDomainConfigsRequest) SetDomainNames(v string) *BatchSetVsDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetVsDomainConfigsRequest) SetFunctions(v string) *BatchSetVsDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetVsDomainConfigsRequest) SetOwnerId(v int64) *BatchSetVsDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

type BatchSetVsDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetVsDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetVsDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetVsDomainConfigsResponseBody) SetRequestId(v string) *BatchSetVsDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetVsDomainConfigsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetVsDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetVsDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetVsDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetVsDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetVsDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetVsDomainConfigsResponse) SetBody(v *BatchSetVsDomainConfigsResponseBody) *BatchSetVsDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchStartDevicesRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchStartDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesRequest) SetId(v string) *BatchStartDevicesRequest {
	s.Id = &v
	return s
}

func (s *BatchStartDevicesRequest) SetOwnerId(v int64) *BatchStartDevicesRequest {
	s.OwnerId = &v
	return s
}

type BatchStartDevicesResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStartDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStartDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponseBody) SetRequestId(v string) *BatchStartDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartDevicesResponseBody) SetResults(v []*BatchStartDevicesResponseBodyResults) *BatchStartDevicesResponseBody {
	s.Results = v
	return s
}

type BatchStartDevicesResponseBodyResults struct {
	Id      *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Streams []*BatchStartDevicesResponseBodyResultsStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
}

func (s BatchStartDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponseBodyResults) SetId(v string) *BatchStartDevicesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResults) SetStreams(v []*BatchStartDevicesResponseBodyResultsStreams) *BatchStartDevicesResponseBodyResults {
	s.Streams = v
	return s
}

type BatchStartDevicesResponseBodyResultsStreams struct {
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStartDevicesResponseBodyResultsStreams) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDevicesResponseBodyResultsStreams) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponseBodyResultsStreams) SetError(v string) *BatchStartDevicesResponseBodyResultsStreams {
	s.Error = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResultsStreams) SetId(v string) *BatchStartDevicesResponseBodyResultsStreams {
	s.Id = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResultsStreams) SetName(v string) *BatchStartDevicesResponseBodyResultsStreams {
	s.Name = &v
	return s
}

type BatchStartDevicesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStartDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponse) SetHeaders(v map[string]*string) *BatchStartDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchStartDevicesResponse) SetBody(v *BatchStartDevicesResponseBody) *BatchStartDevicesResponse {
	s.Body = v
	return s
}

type BatchStartStreamsRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchStartStreamsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartStreamsRequest) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsRequest) SetId(v string) *BatchStartStreamsRequest {
	s.Id = &v
	return s
}

func (s *BatchStartStreamsRequest) SetOwnerId(v int64) *BatchStartStreamsRequest {
	s.OwnerId = &v
	return s
}

type BatchStartStreamsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStartStreamsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStartStreamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsResponseBody) SetRequestId(v string) *BatchStartStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartStreamsResponseBody) SetResults(v []*BatchStartStreamsResponseBodyResults) *BatchStartStreamsResponseBody {
	s.Results = v
	return s
}

type BatchStartStreamsResponseBodyResults struct {
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStartStreamsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchStartStreamsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsResponseBodyResults) SetError(v string) *BatchStartStreamsResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchStartStreamsResponseBodyResults) SetId(v string) *BatchStartStreamsResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStartStreamsResponseBodyResults) SetName(v string) *BatchStartStreamsResponseBodyResults {
	s.Name = &v
	return s
}

type BatchStartStreamsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStartStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartStreamsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartStreamsResponse) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsResponse) SetHeaders(v map[string]*string) *BatchStartStreamsResponse {
	s.Headers = v
	return s
}

func (s *BatchStartStreamsResponse) SetBody(v *BatchStartStreamsResponseBody) *BatchStartStreamsResponse {
	s.Body = v
	return s
}

type BatchStopDevicesRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s BatchStopDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesRequest) SetId(v string) *BatchStopDevicesRequest {
	s.Id = &v
	return s
}

func (s *BatchStopDevicesRequest) SetOwnerId(v int64) *BatchStopDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopDevicesRequest) SetStartTime(v string) *BatchStopDevicesRequest {
	s.StartTime = &v
	return s
}

type BatchStopDevicesResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStopDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStopDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponseBody) SetRequestId(v string) *BatchStopDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopDevicesResponseBody) SetResults(v []*BatchStopDevicesResponseBodyResults) *BatchStopDevicesResponseBody {
	s.Results = v
	return s
}

type BatchStopDevicesResponseBodyResults struct {
	Id      *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Streams []*BatchStopDevicesResponseBodyResultsStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
}

func (s BatchStopDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponseBodyResults) SetId(v string) *BatchStopDevicesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResults) SetStreams(v []*BatchStopDevicesResponseBodyResultsStreams) *BatchStopDevicesResponseBodyResults {
	s.Streams = v
	return s
}

type BatchStopDevicesResponseBodyResultsStreams struct {
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStopDevicesResponseBodyResultsStreams) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDevicesResponseBodyResultsStreams) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponseBodyResultsStreams) SetError(v string) *BatchStopDevicesResponseBodyResultsStreams {
	s.Error = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResultsStreams) SetId(v string) *BatchStopDevicesResponseBodyResultsStreams {
	s.Id = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResultsStreams) SetName(v string) *BatchStopDevicesResponseBodyResultsStreams {
	s.Name = &v
	return s
}

type BatchStopDevicesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStopDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponse) SetHeaders(v map[string]*string) *BatchStopDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchStopDevicesResponse) SetBody(v *BatchStopDevicesResponseBody) *BatchStopDevicesResponse {
	s.Body = v
	return s
}

type BatchStopStreamsRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s BatchStopStreamsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopStreamsRequest) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsRequest) SetId(v string) *BatchStopStreamsRequest {
	s.Id = &v
	return s
}

func (s *BatchStopStreamsRequest) SetOwnerId(v int64) *BatchStopStreamsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopStreamsRequest) SetStartTime(v string) *BatchStopStreamsRequest {
	s.StartTime = &v
	return s
}

type BatchStopStreamsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStopStreamsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStopStreamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsResponseBody) SetRequestId(v string) *BatchStopStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopStreamsResponseBody) SetResults(v []*BatchStopStreamsResponseBodyResults) *BatchStopStreamsResponseBody {
	s.Results = v
	return s
}

type BatchStopStreamsResponseBodyResults struct {
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStopStreamsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchStopStreamsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsResponseBodyResults) SetError(v string) *BatchStopStreamsResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchStopStreamsResponseBodyResults) SetId(v string) *BatchStopStreamsResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStopStreamsResponseBodyResults) SetName(v string) *BatchStopStreamsResponseBodyResults {
	s.Name = &v
	return s
}

type BatchStopStreamsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStopStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopStreamsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopStreamsResponse) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsResponse) SetHeaders(v map[string]*string) *BatchStopStreamsResponse {
	s.Headers = v
	return s
}

func (s *BatchStopStreamsResponse) SetBody(v *BatchStopStreamsResponseBody) *BatchStopStreamsResponse {
	s.Body = v
	return s
}

type BatchUnbindDirectoriesRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchUnbindDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesRequest) SetDeviceId(v string) *BatchUnbindDirectoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindDirectoriesRequest) SetDirectoryId(v string) *BatchUnbindDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *BatchUnbindDirectoriesRequest) SetOwnerId(v int64) *BatchUnbindDirectoriesRequest {
	s.OwnerId = &v
	return s
}

type BatchUnbindDirectoriesResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindDirectoriesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesResponseBody) SetRequestId(v string) *BatchUnbindDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBody) SetResults(v []*BatchUnbindDirectoriesResponseBodyResults) *BatchUnbindDirectoriesResponseBody {
	s.Results = v
	return s
}

type BatchUnbindDirectoriesResponseBodyResults struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Error       *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s BatchUnbindDirectoriesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindDirectoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesResponseBodyResults) SetDeviceId(v string) *BatchUnbindDirectoriesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBodyResults) SetDirectoryId(v string) *BatchUnbindDirectoriesResponseBodyResults {
	s.DirectoryId = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBodyResults) SetError(v string) *BatchUnbindDirectoriesResponseBodyResults {
	s.Error = &v
	return s
}

type BatchUnbindDirectoriesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUnbindDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUnbindDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesResponse) SetHeaders(v map[string]*string) *BatchUnbindDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindDirectoriesResponse) SetBody(v *BatchUnbindDirectoriesResponseBody) *BatchUnbindDirectoriesResponse {
	s.Body = v
	return s
}

type BatchUnbindParentPlatformDevicesRequest struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BatchUnbindParentPlatformDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesRequest) SetDeviceId(v string) *BatchUnbindParentPlatformDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesRequest) SetOwnerId(v int64) *BatchUnbindParentPlatformDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesRequest) SetParentPlatformId(v string) *BatchUnbindParentPlatformDevicesRequest {
	s.ParentPlatformId = &v
	return s
}

type BatchUnbindParentPlatformDevicesResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindParentPlatformDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindParentPlatformDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) SetRequestId(v string) *BatchUnbindParentPlatformDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) SetResults(v []*BatchUnbindParentPlatformDevicesResponseBodyResults) *BatchUnbindParentPlatformDevicesResponseBody {
	s.Results = v
	return s
}

type BatchUnbindParentPlatformDevicesResponseBodyResults struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Error            *string `json:"Error,omitempty" xml:"Error,omitempty"`
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BatchUnbindParentPlatformDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) SetDeviceId(v string) *BatchUnbindParentPlatformDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) SetError(v string) *BatchUnbindParentPlatformDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) SetParentPlatformId(v string) *BatchUnbindParentPlatformDevicesResponseBodyResults {
	s.ParentPlatformId = &v
	return s
}

type BatchUnbindParentPlatformDevicesResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUnbindParentPlatformDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUnbindParentPlatformDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesResponse) SetHeaders(v map[string]*string) *BatchUnbindParentPlatformDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponse) SetBody(v *BatchUnbindParentPlatformDevicesResponseBody) *BatchUnbindParentPlatformDevicesResponse {
	s.Body = v
	return s
}

type BatchUnbindPurchasedDevicesRequest struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchUnbindPurchasedDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesRequest) SetDeviceId(v string) *BatchUnbindPurchasedDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesRequest) SetOwnerId(v int64) *BatchUnbindPurchasedDevicesRequest {
	s.OwnerId = &v
	return s
}

type BatchUnbindPurchasedDevicesResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindPurchasedDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindPurchasedDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesResponseBody) SetRequestId(v string) *BatchUnbindPurchasedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponseBody) SetResults(v []*BatchUnbindPurchasedDevicesResponseBodyResults) *BatchUnbindPurchasedDevicesResponseBody {
	s.Results = v
	return s
}

type BatchUnbindPurchasedDevicesResponseBodyResults struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Error    *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s BatchUnbindPurchasedDevicesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) SetDeviceId(v string) *BatchUnbindPurchasedDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) SetError(v string) *BatchUnbindPurchasedDevicesResponseBodyResults {
	s.Error = &v
	return s
}

type BatchUnbindPurchasedDevicesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUnbindPurchasedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUnbindPurchasedDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesResponse) SetHeaders(v map[string]*string) *BatchUnbindPurchasedDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponse) SetBody(v *BatchUnbindPurchasedDevicesResponseBody) *BatchUnbindPurchasedDevicesResponse {
	s.Body = v
	return s
}

type BatchUnbindTemplateRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchUnbindTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplateRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateRequest) SetInstanceId(v string) *BatchUnbindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetInstanceType(v string) *BatchUnbindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetOwnerId(v int64) *BatchUnbindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetTemplateId(v string) *BatchUnbindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetTemplateType(v string) *BatchUnbindTemplateRequest {
	s.TemplateType = &v
	return s
}

type BatchUnbindTemplateResponseBody struct {
	Bindings  []*BatchUnbindTemplateResponseBodyBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUnbindTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateResponseBody) SetBindings(v []*BatchUnbindTemplateResponseBodyBindings) *BatchUnbindTemplateResponseBody {
	s.Bindings = v
	return s
}

func (s *BatchUnbindTemplateResponseBody) SetRequestId(v string) *BatchUnbindTemplateResponseBody {
	s.RequestId = &v
	return s
}

type BatchUnbindTemplateResponseBodyBindings struct {
	Error        *string `json:"Error,omitempty" xml:"Error,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s BatchUnbindTemplateResponseBodyBindings) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplateResponseBodyBindings) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetError(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.Error = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetInstanceId(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetInstanceType(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetTemplateId(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.TemplateId = &v
	return s
}

type BatchUnbindTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUnbindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUnbindTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplateResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateResponse) SetHeaders(v map[string]*string) *BatchUnbindTemplateResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindTemplateResponse) SetBody(v *BatchUnbindTemplateResponseBody) *BatchUnbindTemplateResponse {
	s.Body = v
	return s
}

type BatchUnbindTemplatesRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchUnbindTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplatesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesRequest) SetInstanceId(v string) *BatchUnbindTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetInstanceType(v string) *BatchUnbindTemplatesRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetOwnerId(v int64) *BatchUnbindTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetTemplateId(v string) *BatchUnbindTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetTemplateType(v string) *BatchUnbindTemplatesRequest {
	s.TemplateType = &v
	return s
}

type BatchUnbindTemplatesResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindTemplatesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesResponseBody) SetRequestId(v string) *BatchUnbindTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBody) SetResults(v []*BatchUnbindTemplatesResponseBodyResults) *BatchUnbindTemplatesResponseBody {
	s.Results = v
	return s
}

type BatchUnbindTemplatesResponseBodyResults struct {
	Error        *string `json:"Error,omitempty" xml:"Error,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchUnbindTemplatesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplatesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetError(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetInstanceId(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetInstanceType(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetTemplateId(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetTemplateType(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.TemplateType = &v
	return s
}

type BatchUnbindTemplatesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUnbindTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUnbindTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUnbindTemplatesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesResponse) SetHeaders(v map[string]*string) *BatchUnbindTemplatesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindTemplatesResponse) SetBody(v *BatchUnbindTemplatesResponseBody) *BatchUnbindTemplatesResponse {
	s.Body = v
	return s
}

type BindDirectoryRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BindDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDirectoryRequest) GoString() string {
	return s.String()
}

func (s *BindDirectoryRequest) SetDeviceId(v string) *BindDirectoryRequest {
	s.DeviceId = &v
	return s
}

func (s *BindDirectoryRequest) SetDirectoryId(v string) *BindDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *BindDirectoryRequest) SetOwnerId(v int64) *BindDirectoryRequest {
	s.OwnerId = &v
	return s
}

type BindDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *BindDirectoryResponseBody) SetRequestId(v string) *BindDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type BindDirectoryResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDirectoryResponse) GoString() string {
	return s.String()
}

func (s *BindDirectoryResponse) SetHeaders(v map[string]*string) *BindDirectoryResponse {
	s.Headers = v
	return s
}

func (s *BindDirectoryResponse) SetBody(v *BindDirectoryResponseBody) *BindDirectoryResponse {
	s.Body = v
	return s
}

type BindParentPlatformDeviceRequest struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BindParentPlatformDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindParentPlatformDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindParentPlatformDeviceRequest) SetDeviceId(v string) *BindParentPlatformDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *BindParentPlatformDeviceRequest) SetOwnerId(v int64) *BindParentPlatformDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *BindParentPlatformDeviceRequest) SetParentPlatformId(v string) *BindParentPlatformDeviceRequest {
	s.ParentPlatformId = &v
	return s
}

type BindParentPlatformDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindParentPlatformDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindParentPlatformDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindParentPlatformDeviceResponseBody) SetRequestId(v string) *BindParentPlatformDeviceResponseBody {
	s.RequestId = &v
	return s
}

type BindParentPlatformDeviceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindParentPlatformDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindParentPlatformDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindParentPlatformDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindParentPlatformDeviceResponse) SetHeaders(v map[string]*string) *BindParentPlatformDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindParentPlatformDeviceResponse) SetBody(v *BindParentPlatformDeviceResponseBody) *BindParentPlatformDeviceResponse {
	s.Body = v
	return s
}

type BindPurchasedDeviceRequest struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s BindPurchasedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindPurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindPurchasedDeviceRequest) SetDeviceId(v string) *BindPurchasedDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *BindPurchasedDeviceRequest) SetGroupId(v string) *BindPurchasedDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *BindPurchasedDeviceRequest) SetOwnerId(v int64) *BindPurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *BindPurchasedDeviceRequest) SetRegion(v string) *BindPurchasedDeviceRequest {
	s.Region = &v
	return s
}

type BindPurchasedDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindPurchasedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindPurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindPurchasedDeviceResponseBody) SetRequestId(v string) *BindPurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

type BindPurchasedDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindPurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindPurchasedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindPurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindPurchasedDeviceResponse) SetHeaders(v map[string]*string) *BindPurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindPurchasedDeviceResponse) SetBody(v *BindPurchasedDeviceResponseBody) *BindPurchasedDeviceResponse {
	s.Body = v
	return s
}

type BindTemplateRequest struct {
	ApplyAll     *bool   `json:"ApplyAll,omitempty" xml:"ApplyAll,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Replace      *bool   `json:"Replace,omitempty" xml:"Replace,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BindTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s BindTemplateRequest) GoString() string {
	return s.String()
}

func (s *BindTemplateRequest) SetApplyAll(v bool) *BindTemplateRequest {
	s.ApplyAll = &v
	return s
}

func (s *BindTemplateRequest) SetInstanceId(v string) *BindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *BindTemplateRequest) SetInstanceType(v string) *BindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *BindTemplateRequest) SetOwnerId(v int64) *BindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *BindTemplateRequest) SetReplace(v bool) *BindTemplateRequest {
	s.Replace = &v
	return s
}

func (s *BindTemplateRequest) SetTemplateId(v string) *BindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *BindTemplateRequest) SetTemplateType(v string) *BindTemplateRequest {
	s.TemplateType = &v
	return s
}

type BindTemplateResponseBody struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s BindTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BindTemplateResponseBody) SetInstanceId(v string) *BindTemplateResponseBody {
	s.InstanceId = &v
	return s
}

func (s *BindTemplateResponseBody) SetInstanceType(v string) *BindTemplateResponseBody {
	s.InstanceType = &v
	return s
}

func (s *BindTemplateResponseBody) SetRequestId(v string) *BindTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindTemplateResponseBody) SetTemplateId(v string) *BindTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type BindTemplateResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s BindTemplateResponse) GoString() string {
	return s.String()
}

func (s *BindTemplateResponse) SetHeaders(v map[string]*string) *BindTemplateResponse {
	s.Headers = v
	return s
}

func (s *BindTemplateResponse) SetBody(v *BindTemplateResponseBody) *BindTemplateResponse {
	s.Body = v
	return s
}

type CaptureDeviceSnapshotRequest struct {
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Mode           *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SnapshotConfig *string `json:"SnapshotConfig,omitempty" xml:"SnapshotConfig,omitempty"`
	StreamId       *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s CaptureDeviceSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CaptureDeviceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CaptureDeviceSnapshotRequest) SetDeviceId(v string) *CaptureDeviceSnapshotRequest {
	s.DeviceId = &v
	return s
}

func (s *CaptureDeviceSnapshotRequest) SetMode(v string) *CaptureDeviceSnapshotRequest {
	s.Mode = &v
	return s
}

func (s *CaptureDeviceSnapshotRequest) SetOwnerId(v int64) *CaptureDeviceSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *CaptureDeviceSnapshotRequest) SetSnapshotConfig(v string) *CaptureDeviceSnapshotRequest {
	s.SnapshotConfig = &v
	return s
}

func (s *CaptureDeviceSnapshotRequest) SetStreamId(v string) *CaptureDeviceSnapshotRequest {
	s.StreamId = &v
	return s
}

type CaptureDeviceSnapshotResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CaptureDeviceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CaptureDeviceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CaptureDeviceSnapshotResponseBody) SetId(v string) *CaptureDeviceSnapshotResponseBody {
	s.Id = &v
	return s
}

func (s *CaptureDeviceSnapshotResponseBody) SetRequestId(v string) *CaptureDeviceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type CaptureDeviceSnapshotResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CaptureDeviceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CaptureDeviceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CaptureDeviceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CaptureDeviceSnapshotResponse) SetHeaders(v map[string]*string) *CaptureDeviceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CaptureDeviceSnapshotResponse) SetBody(v *CaptureDeviceSnapshotResponseBody) *CaptureDeviceSnapshotResponse {
	s.Body = v
	return s
}

type ContinuousAdjustRequest struct {
	Focus   *string `json:"Focus,omitempty" xml:"Focus,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Iris    *string `json:"Iris,omitempty" xml:"Iris,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ContinuousAdjustRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinuousAdjustRequest) GoString() string {
	return s.String()
}

func (s *ContinuousAdjustRequest) SetFocus(v string) *ContinuousAdjustRequest {
	s.Focus = &v
	return s
}

func (s *ContinuousAdjustRequest) SetId(v string) *ContinuousAdjustRequest {
	s.Id = &v
	return s
}

func (s *ContinuousAdjustRequest) SetIris(v string) *ContinuousAdjustRequest {
	s.Iris = &v
	return s
}

func (s *ContinuousAdjustRequest) SetOwnerId(v int64) *ContinuousAdjustRequest {
	s.OwnerId = &v
	return s
}

type ContinuousAdjustResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuousAdjustResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinuousAdjustResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuousAdjustResponseBody) SetId(v string) *ContinuousAdjustResponseBody {
	s.Id = &v
	return s
}

func (s *ContinuousAdjustResponseBody) SetRequestId(v string) *ContinuousAdjustResponseBody {
	s.RequestId = &v
	return s
}

type ContinuousAdjustResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContinuousAdjustResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinuousAdjustResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinuousAdjustResponse) GoString() string {
	return s.String()
}

func (s *ContinuousAdjustResponse) SetHeaders(v map[string]*string) *ContinuousAdjustResponse {
	s.Headers = v
	return s
}

func (s *ContinuousAdjustResponse) SetBody(v *ContinuousAdjustResponseBody) *ContinuousAdjustResponse {
	s.Body = v
	return s
}

type ContinuousMoveRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Pan     *string `json:"Pan,omitempty" xml:"Pan,omitempty"`
	Tilt    *string `json:"Tilt,omitempty" xml:"Tilt,omitempty"`
	Zoom    *string `json:"Zoom,omitempty" xml:"Zoom,omitempty"`
}

func (s ContinuousMoveRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinuousMoveRequest) GoString() string {
	return s.String()
}

func (s *ContinuousMoveRequest) SetId(v string) *ContinuousMoveRequest {
	s.Id = &v
	return s
}

func (s *ContinuousMoveRequest) SetOwnerId(v int64) *ContinuousMoveRequest {
	s.OwnerId = &v
	return s
}

func (s *ContinuousMoveRequest) SetPan(v string) *ContinuousMoveRequest {
	s.Pan = &v
	return s
}

func (s *ContinuousMoveRequest) SetTilt(v string) *ContinuousMoveRequest {
	s.Tilt = &v
	return s
}

func (s *ContinuousMoveRequest) SetZoom(v string) *ContinuousMoveRequest {
	s.Zoom = &v
	return s
}

type ContinuousMoveResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuousMoveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinuousMoveResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuousMoveResponseBody) SetId(v string) *ContinuousMoveResponseBody {
	s.Id = &v
	return s
}

func (s *ContinuousMoveResponseBody) SetRequestId(v string) *ContinuousMoveResponseBody {
	s.RequestId = &v
	return s
}

type ContinuousMoveResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContinuousMoveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinuousMoveResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinuousMoveResponse) GoString() string {
	return s.String()
}

func (s *ContinuousMoveResponse) SetHeaders(v map[string]*string) *ContinuousMoveResponse {
	s.Headers = v
	return s
}

func (s *ContinuousMoveResponse) SetBody(v *ContinuousMoveResponseBody) *ContinuousMoveResponse {
	s.Body = v
	return s
}

type CreateAIConfigRequest struct {
	CaptureInterval *int32  `json:"CaptureInterval,omitempty" xml:"CaptureInterval,omitempty"`
	Configs         *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Features        *string `json:"Features,omitempty" xml:"Features,omitempty"`
	InstanceId      *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAIConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAIConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAIConfigRequest) SetCaptureInterval(v int32) *CreateAIConfigRequest {
	s.CaptureInterval = &v
	return s
}

func (s *CreateAIConfigRequest) SetConfigs(v string) *CreateAIConfigRequest {
	s.Configs = &v
	return s
}

func (s *CreateAIConfigRequest) SetDescription(v string) *CreateAIConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateAIConfigRequest) SetEndTime(v int64) *CreateAIConfigRequest {
	s.EndTime = &v
	return s
}

func (s *CreateAIConfigRequest) SetFeatures(v string) *CreateAIConfigRequest {
	s.Features = &v
	return s
}

func (s *CreateAIConfigRequest) SetInstanceId(v int64) *CreateAIConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAIConfigRequest) SetInstanceType(v string) *CreateAIConfigRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateAIConfigRequest) SetOwnerId(v int64) *CreateAIConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAIConfigRequest) SetStartTime(v int64) *CreateAIConfigRequest {
	s.StartTime = &v
	return s
}

func (s *CreateAIConfigRequest) SetStatus(v string) *CreateAIConfigRequest {
	s.Status = &v
	return s
}

type CreateAIConfigResponseBody struct {
	ConfigId  *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAIConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAIConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIConfigResponseBody) SetConfigId(v string) *CreateAIConfigResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateAIConfigResponseBody) SetRequestId(v string) *CreateAIConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateAIConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAIConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAIConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAIConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAIConfigResponse) SetHeaders(v map[string]*string) *CreateAIConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAIConfigResponse) SetBody(v *CreateAIConfigResponseBody) *CreateAIConfigResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InternalPorts   *string `json:"InternalPorts,omitempty" xml:"InternalPorts,omitempty"`
	MaintainTime    *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetInternalPorts(v string) *CreateClusterRequest {
	s.InternalPorts = &v
	return s
}

func (s *CreateClusterRequest) SetMaintainTime(v string) *CreateClusterRequest {
	s.MaintainTime = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetOwnerId(v int64) *CreateClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateDeviceRequest struct {
	AlarmMethod *string `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoPos     *bool   `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	AutoStart   *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Dsn         *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	GbId        *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Latitude    *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude   *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port        *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	PosInterval *int64  `json:"PosInterval,omitempty" xml:"PosInterval,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Username    *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Vendor      *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceRequest) SetAlarmMethod(v string) *CreateDeviceRequest {
	s.AlarmMethod = &v
	return s
}

func (s *CreateDeviceRequest) SetAutoPos(v bool) *CreateDeviceRequest {
	s.AutoPos = &v
	return s
}

func (s *CreateDeviceRequest) SetAutoStart(v bool) *CreateDeviceRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateDeviceRequest) SetDescription(v string) *CreateDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceRequest) SetDirectoryId(v string) *CreateDeviceRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDeviceRequest) SetDsn(v string) *CreateDeviceRequest {
	s.Dsn = &v
	return s
}

func (s *CreateDeviceRequest) SetGbId(v string) *CreateDeviceRequest {
	s.GbId = &v
	return s
}

func (s *CreateDeviceRequest) SetGroupId(v string) *CreateDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDeviceRequest) SetIp(v string) *CreateDeviceRequest {
	s.Ip = &v
	return s
}

func (s *CreateDeviceRequest) SetLatitude(v string) *CreateDeviceRequest {
	s.Latitude = &v
	return s
}

func (s *CreateDeviceRequest) SetLongitude(v string) *CreateDeviceRequest {
	s.Longitude = &v
	return s
}

func (s *CreateDeviceRequest) SetName(v string) *CreateDeviceRequest {
	s.Name = &v
	return s
}

func (s *CreateDeviceRequest) SetOwnerId(v int64) *CreateDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDeviceRequest) SetParams(v string) *CreateDeviceRequest {
	s.Params = &v
	return s
}

func (s *CreateDeviceRequest) SetParentId(v string) *CreateDeviceRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDeviceRequest) SetPassword(v string) *CreateDeviceRequest {
	s.Password = &v
	return s
}

func (s *CreateDeviceRequest) SetPort(v int64) *CreateDeviceRequest {
	s.Port = &v
	return s
}

func (s *CreateDeviceRequest) SetPosInterval(v int64) *CreateDeviceRequest {
	s.PosInterval = &v
	return s
}

func (s *CreateDeviceRequest) SetType(v string) *CreateDeviceRequest {
	s.Type = &v
	return s
}

func (s *CreateDeviceRequest) SetUrl(v string) *CreateDeviceRequest {
	s.Url = &v
	return s
}

func (s *CreateDeviceRequest) SetUsername(v string) *CreateDeviceRequest {
	s.Username = &v
	return s
}

func (s *CreateDeviceRequest) SetVendor(v string) *CreateDeviceRequest {
	s.Vendor = &v
	return s
}

type CreateDeviceResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponseBody) SetId(v string) *CreateDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDeviceResponseBody) SetRequestId(v string) *CreateDeviceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponse) SetHeaders(v map[string]*string) *CreateDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceResponse) SetBody(v *CreateDeviceResponseBody) *CreateDeviceResponse {
	s.Body = v
	return s
}

type CreateDeviceAlarmRequest struct {
	Alarm      *int32  `json:"Alarm,omitempty" xml:"Alarm,omitempty"`
	ChannelId  *int32  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Expire     *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ObjectType *int32  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubAlarm   *int32  `json:"SubAlarm,omitempty" xml:"SubAlarm,omitempty"`
}

func (s CreateDeviceAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceAlarmRequest) SetAlarm(v int32) *CreateDeviceAlarmRequest {
	s.Alarm = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetChannelId(v int32) *CreateDeviceAlarmRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetEndTime(v int64) *CreateDeviceAlarmRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetExpire(v int64) *CreateDeviceAlarmRequest {
	s.Expire = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetId(v string) *CreateDeviceAlarmRequest {
	s.Id = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetObjectType(v int32) *CreateDeviceAlarmRequest {
	s.ObjectType = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetOwnerId(v int64) *CreateDeviceAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetStartTime(v int64) *CreateDeviceAlarmRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetSubAlarm(v int32) *CreateDeviceAlarmRequest {
	s.SubAlarm = &v
	return s
}

type CreateDeviceAlarmResponseBody struct {
	AlarmDelay *int64  `json:"AlarmDelay,omitempty" xml:"AlarmDelay,omitempty"`
	AlarmId    *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	Expire     *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDeviceAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceAlarmResponseBody) SetAlarmDelay(v int64) *CreateDeviceAlarmResponseBody {
	s.AlarmDelay = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetAlarmId(v string) *CreateDeviceAlarmResponseBody {
	s.AlarmId = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetExpire(v int64) *CreateDeviceAlarmResponseBody {
	s.Expire = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetRequestId(v string) *CreateDeviceAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceAlarmResponseBody) SetUrl(v string) *CreateDeviceAlarmResponseBody {
	s.Url = &v
	return s
}

type CreateDeviceAlarmResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceAlarmResponse) SetHeaders(v map[string]*string) *CreateDeviceAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceAlarmResponse) SetBody(v *CreateDeviceAlarmResponseBody) *CreateDeviceAlarmResponse {
	s.Body = v
	return s
}

type CreateDirectoryRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequest) SetDescription(v string) *CreateDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateDirectoryRequest) SetGroupId(v string) *CreateDirectoryRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDirectoryRequest) SetName(v string) *CreateDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateDirectoryRequest) SetOwnerId(v int64) *CreateDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDirectoryRequest) SetParentId(v string) *CreateDirectoryRequest {
	s.ParentId = &v
	return s
}

type CreateDirectoryResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBody) SetId(v string) *CreateDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetRequestId(v string) *CreateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateDirectoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponse) SetHeaders(v map[string]*string) *CreateDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateDirectoryResponse) SetBody(v *CreateDirectoryResponseBody) *CreateDirectoryResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	App         *string `json:"App,omitempty" xml:"App,omitempty"`
	Callback    *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InProtocol  *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	LazyPull    *bool   `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayDomain  *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	PushDomain  *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetApp(v string) *CreateGroupRequest {
	s.App = &v
	return s
}

func (s *CreateGroupRequest) SetCallback(v string) *CreateGroupRequest {
	s.Callback = &v
	return s
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetInProtocol(v string) *CreateGroupRequest {
	s.InProtocol = &v
	return s
}

func (s *CreateGroupRequest) SetLazyPull(v bool) *CreateGroupRequest {
	s.LazyPull = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetOutProtocol(v string) *CreateGroupRequest {
	s.OutProtocol = &v
	return s
}

func (s *CreateGroupRequest) SetOwnerId(v int64) *CreateGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGroupRequest) SetPlayDomain(v string) *CreateGroupRequest {
	s.PlayDomain = &v
	return s
}

func (s *CreateGroupRequest) SetPushDomain(v string) *CreateGroupRequest {
	s.PushDomain = &v
	return s
}

func (s *CreateGroupRequest) SetRegion(v string) *CreateGroupRequest {
	s.Region = &v
	return s
}

type CreateGroupResponseBody struct {
	GbId      *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GbIp      *string `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	GbPort    *int64  `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGbId(v string) *CreateGroupResponseBody {
	s.GbId = &v
	return s
}

func (s *CreateGroupResponseBody) SetGbIp(v string) *CreateGroupResponseBody {
	s.GbIp = &v
	return s
}

func (s *CreateGroupResponseBody) SetGbPort(v int64) *CreateGroupResponseBody {
	s.GbPort = &v
	return s
}

func (s *CreateGroupResponseBody) SetId(v string) *CreateGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateParentPlatformRequest struct {
	AutoStart      *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ClientAuth     *bool   `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port           *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol       *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateParentPlatformRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *CreateParentPlatformRequest) SetAutoStart(v bool) *CreateParentPlatformRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateParentPlatformRequest) SetClientAuth(v bool) *CreateParentPlatformRequest {
	s.ClientAuth = &v
	return s
}

func (s *CreateParentPlatformRequest) SetClientPassword(v string) *CreateParentPlatformRequest {
	s.ClientPassword = &v
	return s
}

func (s *CreateParentPlatformRequest) SetClientUsername(v string) *CreateParentPlatformRequest {
	s.ClientUsername = &v
	return s
}

func (s *CreateParentPlatformRequest) SetDescription(v string) *CreateParentPlatformRequest {
	s.Description = &v
	return s
}

func (s *CreateParentPlatformRequest) SetGbId(v string) *CreateParentPlatformRequest {
	s.GbId = &v
	return s
}

func (s *CreateParentPlatformRequest) SetIp(v string) *CreateParentPlatformRequest {
	s.Ip = &v
	return s
}

func (s *CreateParentPlatformRequest) SetName(v string) *CreateParentPlatformRequest {
	s.Name = &v
	return s
}

func (s *CreateParentPlatformRequest) SetOwnerId(v int64) *CreateParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateParentPlatformRequest) SetPort(v int64) *CreateParentPlatformRequest {
	s.Port = &v
	return s
}

func (s *CreateParentPlatformRequest) SetProtocol(v string) *CreateParentPlatformRequest {
	s.Protocol = &v
	return s
}

type CreateParentPlatformResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParentPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParentPlatformResponseBody) SetId(v string) *CreateParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *CreateParentPlatformResponseBody) SetRequestId(v string) *CreateParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

type CreateParentPlatformResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateParentPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *CreateParentPlatformResponse) SetHeaders(v map[string]*string) *CreateParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *CreateParentPlatformResponse) SetBody(v *CreateParentPlatformResponseBody) *CreateParentPlatformResponse {
	s.Body = v
	return s
}

type CreateRenderingDeviceRequest struct {
	AutoRenew          *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod    *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Count              *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EdgeNodeName       *string `json:"EdgeNodeName,omitempty" xml:"EdgeNodeName,omitempty"`
	ISP                *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password           *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SecurityGroupId    *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Specification      *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s CreateRenderingDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRenderingDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingDeviceRequest) SetAutoRenew(v bool) *CreateRenderingDeviceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetAutoRenewPeriod(v int32) *CreateRenderingDeviceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetClusterId(v string) *CreateRenderingDeviceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetCount(v int32) *CreateRenderingDeviceRequest {
	s.Count = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetDescription(v string) *CreateRenderingDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetEdgeNodeName(v string) *CreateRenderingDeviceRequest {
	s.EdgeNodeName = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetISP(v string) *CreateRenderingDeviceRequest {
	s.ISP = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetImageId(v string) *CreateRenderingDeviceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetInstanceChargeType(v string) *CreateRenderingDeviceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetInstanceName(v string) *CreateRenderingDeviceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetOwnerId(v int64) *CreateRenderingDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetPassword(v string) *CreateRenderingDeviceRequest {
	s.Password = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetPeriod(v int32) *CreateRenderingDeviceRequest {
	s.Period = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetPeriodUnit(v string) *CreateRenderingDeviceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetSecurityGroupId(v string) *CreateRenderingDeviceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateRenderingDeviceRequest) SetSpecification(v string) *CreateRenderingDeviceRequest {
	s.Specification = &v
	return s
}

type CreateRenderingDeviceResponseBody struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRenderingDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRenderingDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRenderingDeviceResponseBody) SetInstanceIds(v []*string) *CreateRenderingDeviceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateRenderingDeviceResponseBody) SetRequestId(v string) *CreateRenderingDeviceResponseBody {
	s.RequestId = &v
	return s
}

type CreateRenderingDeviceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRenderingDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRenderingDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRenderingDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateRenderingDeviceResponse) SetHeaders(v map[string]*string) *CreateRenderingDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateRenderingDeviceResponse) SetBody(v *CreateRenderingDeviceResponseBody) *CreateRenderingDeviceResponse {
	s.Body = v
	return s
}

type CreateStreamSnapshotRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateStreamSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotRequest) SetId(v string) *CreateStreamSnapshotRequest {
	s.Id = &v
	return s
}

func (s *CreateStreamSnapshotRequest) SetLocation(v string) *CreateStreamSnapshotRequest {
	s.Location = &v
	return s
}

func (s *CreateStreamSnapshotRequest) SetOwnerId(v int64) *CreateStreamSnapshotRequest {
	s.OwnerId = &v
	return s
}

type CreateStreamSnapshotResponseBody struct {
	Format      *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Height      *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Width       *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateStreamSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotResponseBody) SetFormat(v string) *CreateStreamSnapshotResponseBody {
	s.Format = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetHeight(v int64) *CreateStreamSnapshotResponseBody {
	s.Height = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetId(v string) *CreateStreamSnapshotResponseBody {
	s.Id = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetOssBucket(v string) *CreateStreamSnapshotResponseBody {
	s.OssBucket = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetOssEndpoint(v string) *CreateStreamSnapshotResponseBody {
	s.OssEndpoint = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetOssObject(v string) *CreateStreamSnapshotResponseBody {
	s.OssObject = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetRequestId(v string) *CreateStreamSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetTimestamp(v int64) *CreateStreamSnapshotResponseBody {
	s.Timestamp = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetUrl(v string) *CreateStreamSnapshotResponseBody {
	s.Url = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetWidth(v int64) *CreateStreamSnapshotResponseBody {
	s.Width = &v
	return s
}

type CreateStreamSnapshotResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStreamSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStreamSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotResponse) SetHeaders(v map[string]*string) *CreateStreamSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamSnapshotResponse) SetBody(v *CreateStreamSnapshotResponseBody) *CreateStreamSnapshotResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	Callback         *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileFormat       *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	Flv              *string `json:"Flv,omitempty" xml:"Flv,omitempty"`
	HlsM3u8          *string `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	HlsTs            *string `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	Interval         *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JpgOnDemand      *string `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	JpgOverwrite     *string `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	JpgSequence      *string `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	Mp4              *string `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket        *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint      *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssFilePrefix    *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Retention        *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	TransConfigsJSON *string `json:"TransConfigsJSON,omitempty" xml:"TransConfigsJSON,omitempty"`
	Trigger          *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetCallback(v string) *CreateTemplateRequest {
	s.Callback = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetFileFormat(v string) *CreateTemplateRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateTemplateRequest) SetFlv(v string) *CreateTemplateRequest {
	s.Flv = &v
	return s
}

func (s *CreateTemplateRequest) SetHlsM3u8(v string) *CreateTemplateRequest {
	s.HlsM3u8 = &v
	return s
}

func (s *CreateTemplateRequest) SetHlsTs(v string) *CreateTemplateRequest {
	s.HlsTs = &v
	return s
}

func (s *CreateTemplateRequest) SetInterval(v int64) *CreateTemplateRequest {
	s.Interval = &v
	return s
}

func (s *CreateTemplateRequest) SetJpgOnDemand(v string) *CreateTemplateRequest {
	s.JpgOnDemand = &v
	return s
}

func (s *CreateTemplateRequest) SetJpgOverwrite(v string) *CreateTemplateRequest {
	s.JpgOverwrite = &v
	return s
}

func (s *CreateTemplateRequest) SetJpgSequence(v string) *CreateTemplateRequest {
	s.JpgSequence = &v
	return s
}

func (s *CreateTemplateRequest) SetMp4(v string) *CreateTemplateRequest {
	s.Mp4 = &v
	return s
}

func (s *CreateTemplateRequest) SetName(v string) *CreateTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplateRequest) SetOssBucket(v string) *CreateTemplateRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateTemplateRequest) SetOssEndpoint(v string) *CreateTemplateRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateTemplateRequest) SetOssFilePrefix(v string) *CreateTemplateRequest {
	s.OssFilePrefix = &v
	return s
}

func (s *CreateTemplateRequest) SetOwnerId(v int64) *CreateTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTemplateRequest) SetRegion(v string) *CreateTemplateRequest {
	s.Region = &v
	return s
}

func (s *CreateTemplateRequest) SetRetention(v int64) *CreateTemplateRequest {
	s.Retention = &v
	return s
}

func (s *CreateTemplateRequest) SetTransConfigsJSON(v string) *CreateTemplateRequest {
	s.TransConfigsJSON = &v
	return s
}

func (s *CreateTemplateRequest) SetTrigger(v string) *CreateTemplateRequest {
	s.Trigger = &v
	return s
}

func (s *CreateTemplateRequest) SetType(v string) *CreateTemplateRequest {
	s.Type = &v
	return s
}

type CreateTemplateResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetId(v string) *CreateTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteAIConfigRequest struct {
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteAIConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAIConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIConfigRequest) SetConfigId(v string) *DeleteAIConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteAIConfigRequest) SetOwnerId(v int64) *DeleteAIConfigRequest {
	s.OwnerId = &v
	return s
}

type DeleteAIConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAIConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIConfigResponseBody) SetRequestId(v string) *DeleteAIConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAIConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAIConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAIConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAIConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIConfigResponse) SetHeaders(v map[string]*string) *DeleteAIConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIConfigResponse) SetBody(v *DeleteAIConfigResponseBody) *DeleteAIConfigResponse {
	s.Body = v
	return s
}

type DeleteBucketRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteBucketRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketRequest) SetBucketName(v string) *DeleteBucketRequest {
	s.BucketName = &v
	return s
}

func (s *DeleteBucketRequest) SetOwnerId(v int64) *DeleteBucketRequest {
	s.OwnerId = &v
	return s
}

type DeleteBucketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBucketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBucketResponseBody) SetRequestId(v string) *DeleteBucketResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBucketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBucketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketResponse) SetHeaders(v map[string]*string) *DeleteBucketResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketResponse) SetBody(v *DeleteBucketResponseBody) *DeleteBucketResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetOwnerId(v int64) *DeleteClusterRequest {
	s.OwnerId = &v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteDeviceRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) SetId(v string) *DeleteDeviceRequest {
	s.Id = &v
	return s
}

func (s *DeleteDeviceRequest) SetOwnerId(v int64) *DeleteDeviceRequest {
	s.OwnerId = &v
	return s
}

type DeleteDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponseBody) SetRequestId(v string) *DeleteDeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponse) SetHeaders(v map[string]*string) *DeleteDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceResponse) SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse {
	s.Body = v
	return s
}

type DeleteDirectoryRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryRequest) SetId(v string) *DeleteDirectoryRequest {
	s.Id = &v
	return s
}

func (s *DeleteDirectoryRequest) SetOwnerId(v int64) *DeleteDirectoryRequest {
	s.OwnerId = &v
	return s
}

type DeleteDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponseBody) SetRequestId(v string) *DeleteDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDirectoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponse) SetHeaders(v map[string]*string) *DeleteDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoryResponse) SetBody(v *DeleteDirectoryResponseBody) *DeleteDirectoryResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetId(v string) *DeleteGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteGroupRequest) SetOwnerId(v int64) *DeleteGroupRequest {
	s.OwnerId = &v
	return s
}

type DeleteGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteParentPlatformRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteParentPlatformRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *DeleteParentPlatformRequest) SetId(v string) *DeleteParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *DeleteParentPlatformRequest) SetOwnerId(v int64) *DeleteParentPlatformRequest {
	s.OwnerId = &v
	return s
}

type DeleteParentPlatformResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParentPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParentPlatformResponseBody) SetRequestId(v string) *DeleteParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

type DeleteParentPlatformResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteParentPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *DeleteParentPlatformResponse) SetHeaders(v map[string]*string) *DeleteParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *DeleteParentPlatformResponse) SetBody(v *DeleteParentPlatformResponseBody) *DeleteParentPlatformResponse {
	s.Body = v
	return s
}

type DeletePresetRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
}

func (s DeletePresetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePresetRequest) GoString() string {
	return s.String()
}

func (s *DeletePresetRequest) SetId(v string) *DeletePresetRequest {
	s.Id = &v
	return s
}

func (s *DeletePresetRequest) SetOwnerId(v int64) *DeletePresetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePresetRequest) SetPresetId(v string) *DeletePresetRequest {
	s.PresetId = &v
	return s
}

type DeletePresetResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePresetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePresetResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePresetResponseBody) SetId(v string) *DeletePresetResponseBody {
	s.Id = &v
	return s
}

func (s *DeletePresetResponseBody) SetRequestId(v string) *DeletePresetResponseBody {
	s.RequestId = &v
	return s
}

type DeletePresetResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePresetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePresetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePresetResponse) GoString() string {
	return s.String()
}

func (s *DeletePresetResponse) SetHeaders(v map[string]*string) *DeletePresetResponse {
	s.Headers = v
	return s
}

func (s *DeletePresetResponse) SetBody(v *DeletePresetResponseBody) *DeletePresetResponse {
	s.Body = v
	return s
}

type DeletePurchasedDeviceRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeletePurchasedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeletePurchasedDeviceRequest) SetId(v string) *DeletePurchasedDeviceRequest {
	s.Id = &v
	return s
}

func (s *DeletePurchasedDeviceRequest) SetOwnerId(v int64) *DeletePurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

type DeletePurchasedDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePurchasedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePurchasedDeviceResponseBody) SetRequestId(v string) *DeletePurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeletePurchasedDeviceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePurchasedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeletePurchasedDeviceResponse) SetHeaders(v map[string]*string) *DeletePurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeletePurchasedDeviceResponse) SetBody(v *DeletePurchasedDeviceResponseBody) *DeletePurchasedDeviceResponse {
	s.Body = v
	return s
}

type DeleteRenderingDeviceInternetPortsRequest struct {
	ISP          *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteRenderingDeviceInternetPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDeviceInternetPortsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDeviceInternetPortsRequest) SetISP(v string) *DeleteRenderingDeviceInternetPortsRequest {
	s.ISP = &v
	return s
}

func (s *DeleteRenderingDeviceInternetPortsRequest) SetInstanceIds(v string) *DeleteRenderingDeviceInternetPortsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DeleteRenderingDeviceInternetPortsRequest) SetInternalPort(v string) *DeleteRenderingDeviceInternetPortsRequest {
	s.InternalPort = &v
	return s
}

func (s *DeleteRenderingDeviceInternetPortsRequest) SetIpProtocol(v string) *DeleteRenderingDeviceInternetPortsRequest {
	s.IpProtocol = &v
	return s
}

func (s *DeleteRenderingDeviceInternetPortsRequest) SetOwnerId(v int64) *DeleteRenderingDeviceInternetPortsRequest {
	s.OwnerId = &v
	return s
}

type DeleteRenderingDeviceInternetPortsResponseBody struct {
	InstanceIds *DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	RequestId   *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRenderingDeviceInternetPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDeviceInternetPortsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDeviceInternetPortsResponseBody) SetInstanceIds(v *DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds) *DeleteRenderingDeviceInternetPortsResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DeleteRenderingDeviceInternetPortsResponseBody) SetRequestId(v string) *DeleteRenderingDeviceInternetPortsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds struct {
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
}

func (s DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds) SetInstanceIds(v []*string) *DeleteRenderingDeviceInternetPortsResponseBodyInstanceIds {
	s.InstanceIds = v
	return s
}

type DeleteRenderingDeviceInternetPortsResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRenderingDeviceInternetPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRenderingDeviceInternetPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDeviceInternetPortsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDeviceInternetPortsResponse) SetHeaders(v map[string]*string) *DeleteRenderingDeviceInternetPortsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRenderingDeviceInternetPortsResponse) SetBody(v *DeleteRenderingDeviceInternetPortsResponseBody) *DeleteRenderingDeviceInternetPortsResponse {
	s.Body = v
	return s
}

type DeleteRenderingDevicesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteRenderingDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDevicesRequest) SetInstanceIds(v string) *DeleteRenderingDevicesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DeleteRenderingDevicesRequest) SetOwnerId(v int64) *DeleteRenderingDevicesRequest {
	s.OwnerId = &v
	return s
}

type DeleteRenderingDevicesResponseBody struct {
	FailedIds []*string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRenderingDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDevicesResponseBody) SetFailedIds(v []*string) *DeleteRenderingDevicesResponseBody {
	s.FailedIds = v
	return s
}

func (s *DeleteRenderingDevicesResponseBody) SetRequestId(v string) *DeleteRenderingDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRenderingDevicesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRenderingDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRenderingDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRenderingDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRenderingDevicesResponse) SetHeaders(v map[string]*string) *DeleteRenderingDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRenderingDevicesResponse) SetBody(v *DeleteRenderingDevicesResponseBody) *DeleteRenderingDevicesResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetId(v string) *DeleteTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteTemplateRequest) SetOwnerId(v int64) *DeleteTemplateRequest {
	s.OwnerId = &v
	return s
}

type DeleteTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteVsPullStreamInfoConfigRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteVsPullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetAppName(v string) *DeleteVsPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetDomainName(v string) *DeleteVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *DeleteVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetStreamName(v string) *DeleteVsPullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

type DeleteVsPullStreamInfoConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVsPullStreamInfoConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *DeleteVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVsPullStreamInfoConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVsPullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *DeleteVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVsPullStreamInfoConfigResponse) SetBody(v *DeleteVsPullStreamInfoConfigResponseBody) *DeleteVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

type DeleteVsStreamsNotifyUrlConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteVsStreamsNotifyUrlConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVsStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DeleteVsStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *DeleteVsStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

type DeleteVsStreamsNotifyUrlConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVsStreamsNotifyUrlConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVsStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVsStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *DeleteVsStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVsStreamsNotifyUrlConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVsStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVsStreamsNotifyUrlConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVsStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *DeleteVsStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) SetBody(v *DeleteVsStreamsNotifyUrlConfigResponseBody) *DeleteVsStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

type DescribeAIConfigRequest struct {
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAIConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigRequest) SetConfigId(v string) *DescribeAIConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeAIConfigRequest) SetOwnerId(v int64) *DescribeAIConfigRequest {
	s.OwnerId = &v
	return s
}

type DescribeAIConfigResponseBody struct {
	AIConfig  *DescribeAIConfigResponseBodyAIConfig `json:"AIConfig,omitempty" xml:"AIConfig,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigResponseBody) SetAIConfig(v *DescribeAIConfigResponseBodyAIConfig) *DescribeAIConfigResponseBody {
	s.AIConfig = v
	return s
}

func (s *DescribeAIConfigResponseBody) SetRequestId(v string) *DescribeAIConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAIConfigResponseBodyAIConfig struct {
	CaptureInterval *int32  `json:"CaptureInterval,omitempty" xml:"CaptureInterval,omitempty"`
	Configs         *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Features        *string `json:"Features,omitempty" xml:"Features,omitempty"`
	InstanceId      *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAIConfigResponseBodyAIConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigResponseBodyAIConfig) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetCaptureInterval(v int32) *DescribeAIConfigResponseBodyAIConfig {
	s.CaptureInterval = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetConfigs(v string) *DescribeAIConfigResponseBodyAIConfig {
	s.Configs = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetDescription(v string) *DescribeAIConfigResponseBodyAIConfig {
	s.Description = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetEndTime(v int64) *DescribeAIConfigResponseBodyAIConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetFeatures(v string) *DescribeAIConfigResponseBodyAIConfig {
	s.Features = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetInstanceId(v int64) *DescribeAIConfigResponseBodyAIConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetInstanceType(v string) *DescribeAIConfigResponseBodyAIConfig {
	s.InstanceType = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetStartTime(v int64) *DescribeAIConfigResponseBodyAIConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeAIConfigResponseBodyAIConfig) SetStatus(v string) *DescribeAIConfigResponseBodyAIConfig {
	s.Status = &v
	return s
}

type DescribeAIConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAIConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAIConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigResponse) SetHeaders(v map[string]*string) *DescribeAIConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIConfigResponse) SetBody(v *DescribeAIConfigResponseBody) *DescribeAIConfigResponse {
	s.Body = v
	return s
}

type DescribeAIConfigListRequest struct {
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo   *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAIConfigListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigListRequest) SetOwnerId(v int64) *DescribeAIConfigListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIConfigListRequest) SetPageNo(v int64) *DescribeAIConfigListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAIConfigListRequest) SetPageSize(v int64) *DescribeAIConfigListRequest {
	s.PageSize = &v
	return s
}

type DescribeAIConfigListResponseBody struct {
	AIConfigList *DescribeAIConfigListResponseBodyAIConfigList `json:"AIConfigList,omitempty" xml:"AIConfigList,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIConfigListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigListResponseBody) SetAIConfigList(v *DescribeAIConfigListResponseBodyAIConfigList) *DescribeAIConfigListResponseBody {
	s.AIConfigList = v
	return s
}

func (s *DescribeAIConfigListResponseBody) SetRequestId(v string) *DescribeAIConfigListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAIConfigListResponseBodyAIConfigList struct {
	AIConfigList []*DescribeAIConfigListResponseBodyAIConfigListAIConfigList `json:"AIConfigList,omitempty" xml:"AIConfigList,omitempty" type:"Repeated"`
}

func (s DescribeAIConfigListResponseBodyAIConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigListResponseBodyAIConfigList) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigListResponseBodyAIConfigList) SetAIConfigList(v []*DescribeAIConfigListResponseBodyAIConfigListAIConfigList) *DescribeAIConfigListResponseBodyAIConfigList {
	s.AIConfigList = v
	return s
}

type DescribeAIConfigListResponseBodyAIConfigListAIConfigList struct {
	CaptureInterval *int32  `json:"CaptureInterval,omitempty" xml:"CaptureInterval,omitempty"`
	ConfigId        *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Configs         *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Features        *string `json:"Features,omitempty" xml:"Features,omitempty"`
	InstanceId      *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAIConfigListResponseBodyAIConfigListAIConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigListResponseBodyAIConfigListAIConfigList) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetCaptureInterval(v int32) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.CaptureInterval = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetConfigId(v string) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.ConfigId = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetConfigs(v string) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.Configs = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetDescription(v string) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.Description = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetEndTime(v int64) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.EndTime = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetFeatures(v string) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.Features = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetInstanceId(v int64) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetInstanceType(v string) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.InstanceType = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetStartTime(v int64) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.StartTime = &v
	return s
}

func (s *DescribeAIConfigListResponseBodyAIConfigListAIConfigList) SetStatus(v string) *DescribeAIConfigListResponseBodyAIConfigListAIConfigList {
	s.Status = &v
	return s
}

type DescribeAIConfigListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAIConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAIConfigListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIConfigListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIConfigListResponse) SetHeaders(v map[string]*string) *DescribeAIConfigListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIConfigListResponse) SetBody(v *DescribeAIConfigListResponseBody) *DescribeAIConfigListResponse {
	s.Body = v
	return s
}

type DescribeAIEventListRequest struct {
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Feature      *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	InstanceId   *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Triggered    *bool   `json:"Triggered,omitempty" xml:"Triggered,omitempty"`
}

func (s DescribeAIEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIEventListRequest) SetEndTime(v int64) *DescribeAIEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAIEventListRequest) SetFeature(v string) *DescribeAIEventListRequest {
	s.Feature = &v
	return s
}

func (s *DescribeAIEventListRequest) SetInstanceId(v int64) *DescribeAIEventListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAIEventListRequest) SetInstanceType(v string) *DescribeAIEventListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeAIEventListRequest) SetOwnerId(v int64) *DescribeAIEventListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIEventListRequest) SetStartTime(v int64) *DescribeAIEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAIEventListRequest) SetTriggered(v bool) *DescribeAIEventListRequest {
	s.Triggered = &v
	return s
}

type DescribeAIEventListResponseBody struct {
	AIEventList *string `json:"AIEventList,omitempty" xml:"AIEventList,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIEventListResponseBody) SetAIEventList(v string) *DescribeAIEventListResponseBody {
	s.AIEventList = &v
	return s
}

func (s *DescribeAIEventListResponseBody) SetRequestId(v string) *DescribeAIEventListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAIEventListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAIEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAIEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAIEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIEventListResponse) SetHeaders(v map[string]*string) *DescribeAIEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIEventListResponse) SetBody(v *DescribeAIEventListResponseBody) *DescribeAIEventListResponse {
	s.Body = v
	return s
}

type DescribeAccountStatRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAccountStatRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountStatRequest) SetId(v string) *DescribeAccountStatRequest {
	s.Id = &v
	return s
}

func (s *DescribeAccountStatRequest) SetOwnerId(v int64) *DescribeAccountStatRequest {
	s.OwnerId = &v
	return s
}

type DescribeAccountStatResponseBody struct {
	GroupLimit    *int64  `json:"GroupLimit,omitempty" xml:"GroupLimit,omitempty"`
	GroupNum      *int64  `json:"GroupNum,omitempty" xml:"GroupNum,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateLimit *int64  `json:"TemplateLimit,omitempty" xml:"TemplateLimit,omitempty"`
	TemplateNum   *int64  `json:"TemplateNum,omitempty" xml:"TemplateNum,omitempty"`
}

func (s DescribeAccountStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountStatResponseBody) SetGroupLimit(v int64) *DescribeAccountStatResponseBody {
	s.GroupLimit = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetGroupNum(v int64) *DescribeAccountStatResponseBody {
	s.GroupNum = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetId(v string) *DescribeAccountStatResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetRequestId(v string) *DescribeAccountStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetTemplateLimit(v int64) *DescribeAccountStatResponseBody {
	s.TemplateLimit = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetTemplateNum(v int64) *DescribeAccountStatResponseBody {
	s.TemplateNum = &v
	return s
}

type DescribeAccountStatResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountStatResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountStatResponse) SetHeaders(v map[string]*string) *DescribeAccountStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountStatResponse) SetBody(v *DescribeAccountStatResponseBody) *DescribeAccountStatResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterRequest) SetOwnerId(v int64) *DescribeClusterRequest {
	s.OwnerId = &v
	return s
}

type DescribeClusterResponseBody struct {
	ClusterId     *string                                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description   *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	InternalPorts []*DescribeClusterResponseBodyInternalPorts `json:"InternalPorts,omitempty" xml:"InternalPorts,omitempty" type:"Repeated"`
	MaintainTime  *string                                     `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Name          *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterId(v string) *DescribeClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetDescription(v string) *DescribeClusterResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeClusterResponseBody) SetInternalPorts(v []*DescribeClusterResponseBodyInternalPorts) *DescribeClusterResponseBody {
	s.InternalPorts = v
	return s
}

func (s *DescribeClusterResponseBody) SetMaintainTime(v string) *DescribeClusterResponseBody {
	s.MaintainTime = &v
	return s
}

func (s *DescribeClusterResponseBody) SetName(v string) *DescribeClusterResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetStatus(v string) *DescribeClusterResponseBody {
	s.Status = &v
	return s
}

type DescribeClusterResponseBodyInternalPorts struct {
	IpProtocol *string   `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Platform   *string   `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Port       []*string `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyInternalPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyInternalPorts) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyInternalPorts) SetIpProtocol(v string) *DescribeClusterResponseBodyInternalPorts {
	s.IpProtocol = &v
	return s
}

func (s *DescribeClusterResponseBodyInternalPorts) SetPlatform(v string) *DescribeClusterResponseBodyInternalPorts {
	s.Platform = &v
	return s
}

func (s *DescribeClusterResponseBodyInternalPorts) SetPort(v []*string) *DescribeClusterResponseBodyInternalPorts {
	s.Port = v
	return s
}

type DescribeClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeClusterDevicesRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EdgeNodeName  *string `json:"EdgeNodeName,omitempty" xml:"EdgeNodeName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo        *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platform      *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s DescribeClusterDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesRequest) SetClusterId(v string) *DescribeClusterDevicesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetDescription(v string) *DescribeClusterDevicesRequest {
	s.Description = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetEdgeNodeName(v string) *DescribeClusterDevicesRequest {
	s.EdgeNodeName = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetOwnerId(v int64) *DescribeClusterDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetPageNo(v int64) *DescribeClusterDevicesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetPageSize(v int64) *DescribeClusterDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetPlatform(v string) *DescribeClusterDevicesRequest {
	s.Platform = &v
	return s
}

func (s *DescribeClusterDevicesRequest) SetSpecification(v string) *DescribeClusterDevicesRequest {
	s.Specification = &v
	return s
}

type DescribeClusterDevicesResponseBody struct {
	Devices   []*DescribeClusterDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeClusterDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesResponseBody) SetDevices(v []*DescribeClusterDevicesResponseBodyDevices) *DescribeClusterDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeClusterDevicesResponseBody) SetRequestId(v string) *DescribeClusterDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterDevicesResponseBody) SetTotal(v int64) *DescribeClusterDevicesResponseBody {
	s.Total = &v
	return s
}

type DescribeClusterDevicesResponseBodyDevices struct {
	AutoRenew          *bool                                                `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod    *int32                                               `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	Description        *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	EdgeNodeName       *string                                              `json:"EdgeNodeName,omitempty" xml:"EdgeNodeName,omitempty"`
	ImageId            *string                                              `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string                                              `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceId         *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                              `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IpInfos            []*DescribeClusterDevicesResponseBodyDevicesIpInfos  `json:"IpInfos,omitempty" xml:"IpInfos,omitempty" type:"Repeated"`
	MacAddress         *string                                              `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	Period             *int32                                               `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string                                              `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PlatformType       *string                                              `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	PodInfos           []*DescribeClusterDevicesResponseBodyDevicesPodInfos `json:"PodInfos,omitempty" xml:"PodInfos,omitempty" type:"Repeated"`
	Server             *string                                              `json:"Server,omitempty" xml:"Server,omitempty"`
	Status             *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetAutoRenew(v bool) *DescribeClusterDevicesResponseBodyDevices {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetAutoRenewPeriod(v int32) *DescribeClusterDevicesResponseBodyDevices {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetDescription(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.Description = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetEdgeNodeName(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.EdgeNodeName = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetImageId(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetInstanceChargeType(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetInstanceId(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetInstanceName(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetIpInfos(v []*DescribeClusterDevicesResponseBodyDevicesIpInfos) *DescribeClusterDevicesResponseBodyDevices {
	s.IpInfos = v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetMacAddress(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.MacAddress = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetPeriod(v int32) *DescribeClusterDevicesResponseBodyDevices {
	s.Period = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetPeriodUnit(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetPlatformType(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.PlatformType = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetPodInfos(v []*DescribeClusterDevicesResponseBodyDevicesPodInfos) *DescribeClusterDevicesResponseBodyDevices {
	s.PodInfos = v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetServer(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.Server = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevices) SetStatus(v string) *DescribeClusterDevicesResponseBodyDevices {
	s.Status = &v
	return s
}

type DescribeClusterDevicesResponseBodyDevicesIpInfos struct {
	ExternalIp   *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	ISP          *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InternalIp   *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NatType      *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
}

func (s DescribeClusterDevicesResponseBodyDevicesIpInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesResponseBodyDevicesIpInfos) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetExternalIp(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.ExternalIp = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetExternalPort(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.ExternalPort = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetISP(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.ISP = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetInternalIp(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.InternalIp = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetInternalPort(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.InternalPort = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetIpProtocol(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.IpProtocol = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesIpInfos) SetNatType(v string) *DescribeClusterDevicesResponseBodyDevicesIpInfos {
	s.NatType = &v
	return s
}

type DescribeClusterDevicesResponseBodyDevicesPodInfos struct {
	Network []*DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Repeated"`
	PodId   *string                                                     `json:"PodId,omitempty" xml:"PodId,omitempty"`
	Status  *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterDevicesResponseBodyDevicesPodInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesResponseBodyDevicesPodInfos) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfos) SetNetwork(v []*DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) *DescribeClusterDevicesResponseBodyDevicesPodInfos {
	s.Network = v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfos) SetPodId(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfos {
	s.PodId = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfos) SetStatus(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfos {
	s.Status = &v
	return s
}

type DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork struct {
	ContainerPorts *string `json:"ContainerPorts,omitempty" xml:"ContainerPorts,omitempty"`
	ExternalIp     *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalIsp    *string `json:"ExternalIsp,omitempty" xml:"ExternalIsp,omitempty"`
	ExternalPorts  *string `json:"ExternalPorts,omitempty" xml:"ExternalPorts,omitempty"`
	OutgoingIp     *string `json:"OutgoingIp,omitempty" xml:"OutgoingIp,omitempty"`
	OutgoingIsp    *string `json:"OutgoingIsp,omitempty" xml:"OutgoingIsp,omitempty"`
}

func (s DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) SetContainerPorts(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork {
	s.ContainerPorts = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) SetExternalIp(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork {
	s.ExternalIp = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) SetExternalIsp(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork {
	s.ExternalIsp = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) SetExternalPorts(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork {
	s.ExternalPorts = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) SetOutgoingIp(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork {
	s.OutgoingIp = &v
	return s
}

func (s *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork) SetOutgoingIsp(v string) *DescribeClusterDevicesResponseBodyDevicesPodInfosNetwork {
	s.OutgoingIsp = &v
	return s
}

type DescribeClusterDevicesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDevicesResponse) SetHeaders(v map[string]*string) *DescribeClusterDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterDevicesResponse) SetBody(v *DescribeClusterDevicesResponseBody) *DescribeClusterDevicesResponse {
	s.Body = v
	return s
}

type DescribeClustersRequest struct {
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo   *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetOwnerId(v int64) *DescribeClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClustersRequest) SetPageNo(v int64) *DescribeClustersRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeClustersRequest) SetPageSize(v int64) *DescribeClustersRequest {
	s.PageSize = &v
	return s
}

type DescribeClustersResponseBody struct {
	Clusters  []*DescribeClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) SetClusters(v []*DescribeClustersResponseBodyClusters) *DescribeClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersResponseBody) SetRequestId(v string) *DescribeClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetTotal(v int64) *DescribeClustersResponseBody {
	s.Total = &v
	return s
}

type DescribeClustersResponseBodyClusters struct {
	ClusterId     *string                                              `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description   *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	InternalPorts []*DescribeClustersResponseBodyClustersInternalPorts `json:"InternalPorts,omitempty" xml:"InternalPorts,omitempty" type:"Repeated"`
	MaintainTime  *string                                              `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Name          *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Status        *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClusters) SetClusterId(v string) *DescribeClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetDescription(v string) *DescribeClustersResponseBodyClusters {
	s.Description = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetInternalPorts(v []*DescribeClustersResponseBodyClustersInternalPorts) *DescribeClustersResponseBodyClusters {
	s.InternalPorts = v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetMaintainTime(v string) *DescribeClustersResponseBodyClusters {
	s.MaintainTime = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetName(v string) *DescribeClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetStatus(v string) *DescribeClustersResponseBodyClusters {
	s.Status = &v
	return s
}

type DescribeClustersResponseBodyClustersInternalPorts struct {
	IpProtocol *string   `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Platform   *string   `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Port       []*string `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
}

func (s DescribeClustersResponseBodyClustersInternalPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersInternalPorts) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersInternalPorts) SetIpProtocol(v string) *DescribeClustersResponseBodyClustersInternalPorts {
	s.IpProtocol = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersInternalPorts) SetPlatform(v string) *DescribeClustersResponseBodyClustersInternalPorts {
	s.Platform = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersInternalPorts) SetPort(v []*string) *DescribeClustersResponseBodyClustersInternalPorts {
	s.Port = v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetBody(v *DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
	return s
}

type DescribeContainerInstanceIdRequest struct {
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PodIndex *int32  `json:"PodIndex,omitempty" xml:"PodIndex,omitempty"`
}

func (s DescribeContainerInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstanceIdRequest) SetNodeName(v string) *DescribeContainerInstanceIdRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeContainerInstanceIdRequest) SetOwnerId(v int64) *DescribeContainerInstanceIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerInstanceIdRequest) SetPodIndex(v int32) *DescribeContainerInstanceIdRequest {
	s.PodIndex = &v
	return s
}

type DescribeContainerInstanceIdResponseBody struct {
	InstanceDetail *DescribeContainerInstanceIdResponseBodyInstanceDetail `json:"InstanceDetail,omitempty" xml:"InstanceDetail,omitempty" type:"Struct"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstanceIdResponseBody) SetInstanceDetail(v *DescribeContainerInstanceIdResponseBodyInstanceDetail) *DescribeContainerInstanceIdResponseBody {
	s.InstanceDetail = v
	return s
}

func (s *DescribeContainerInstanceIdResponseBody) SetRequestId(v string) *DescribeContainerInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerInstanceIdResponseBodyInstanceDetail struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeContainerInstanceIdResponseBodyInstanceDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerInstanceIdResponseBodyInstanceDetail) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstanceIdResponseBodyInstanceDetail) SetInstanceId(v string) *DescribeContainerInstanceIdResponseBodyInstanceDetail {
	s.InstanceId = &v
	return s
}

type DescribeContainerInstanceIdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstanceIdResponse) SetHeaders(v map[string]*string) *DescribeContainerInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerInstanceIdResponse) SetBody(v *DescribeContainerInstanceIdResponseBody) *DescribeContainerInstanceIdResponse {
	s.Body = v
	return s
}

type DescribeDeviceRequest struct {
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IncludeDirectory *bool   `json:"IncludeDirectory,omitempty" xml:"IncludeDirectory,omitempty"`
	IncludeStats     *bool   `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceRequest) SetId(v string) *DescribeDeviceRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceRequest) SetIncludeDirectory(v bool) *DescribeDeviceRequest {
	s.IncludeDirectory = &v
	return s
}

func (s *DescribeDeviceRequest) SetIncludeStats(v bool) *DescribeDeviceRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeDeviceRequest) SetOwnerId(v int64) *DescribeDeviceRequest {
	s.OwnerId = &v
	return s
}

type DescribeDeviceResponseBody struct {
	AlarmMethod     *string                              `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoPos         *bool                                `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	AutoStart       *bool                                `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ChannelSyncTime *string                              `json:"ChannelSyncTime,omitempty" xml:"ChannelSyncTime,omitempty"`
	CreatedTime     *string                              `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description     *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory       *DescribeDeviceResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	DirectoryId     *string                              `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Dsn             *string                              `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	Enabled         *bool                                `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GbId            *string                              `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GroupId         *string                              `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id              *string                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip              *string                              `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Latitude        *string                              `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude       *string                              `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name            *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Params          *string                              `json:"Params,omitempty" xml:"Params,omitempty"`
	ParentId        *string                              `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Password        *string                              `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int64                               `json:"Port,omitempty" xml:"Port,omitempty"`
	PosInterval     *int64                               `json:"PosInterval,omitempty" xml:"PosInterval,omitempty"`
	Protocol        *string                              `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegisteredTime  *string                              `json:"RegisteredTime,omitempty" xml:"RegisteredTime,omitempty"`
	RequestId       *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stats           *DescribeDeviceResponseBodyStats     `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	Status          *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Url             *string                              `json:"Url,omitempty" xml:"Url,omitempty"`
	Username        *string                              `json:"Username,omitempty" xml:"Username,omitempty"`
	Vendor          *string                              `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBody) SetAlarmMethod(v string) *DescribeDeviceResponseBody {
	s.AlarmMethod = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetAutoPos(v bool) *DescribeDeviceResponseBody {
	s.AutoPos = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetAutoStart(v bool) *DescribeDeviceResponseBody {
	s.AutoStart = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetChannelSyncTime(v string) *DescribeDeviceResponseBody {
	s.ChannelSyncTime = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetCreatedTime(v string) *DescribeDeviceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDescription(v string) *DescribeDeviceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDirectory(v *DescribeDeviceResponseBodyDirectory) *DescribeDeviceResponseBody {
	s.Directory = v
	return s
}

func (s *DescribeDeviceResponseBody) SetDirectoryId(v string) *DescribeDeviceResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDsn(v string) *DescribeDeviceResponseBody {
	s.Dsn = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetEnabled(v bool) *DescribeDeviceResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetGbId(v string) *DescribeDeviceResponseBody {
	s.GbId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetGroupId(v string) *DescribeDeviceResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetId(v string) *DescribeDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetIp(v string) *DescribeDeviceResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetLatitude(v string) *DescribeDeviceResponseBody {
	s.Latitude = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetLongitude(v string) *DescribeDeviceResponseBody {
	s.Longitude = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetName(v string) *DescribeDeviceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetParams(v string) *DescribeDeviceResponseBody {
	s.Params = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetParentId(v string) *DescribeDeviceResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetPassword(v string) *DescribeDeviceResponseBody {
	s.Password = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetPort(v int64) *DescribeDeviceResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetPosInterval(v int64) *DescribeDeviceResponseBody {
	s.PosInterval = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetProtocol(v string) *DescribeDeviceResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetRegisteredTime(v string) *DescribeDeviceResponseBody {
	s.RegisteredTime = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetRequestId(v string) *DescribeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetStats(v *DescribeDeviceResponseBodyStats) *DescribeDeviceResponseBody {
	s.Stats = v
	return s
}

func (s *DescribeDeviceResponseBody) SetStatus(v string) *DescribeDeviceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetType(v string) *DescribeDeviceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetUrl(v string) *DescribeDeviceResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetUsername(v string) *DescribeDeviceResponseBody {
	s.Username = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetVendor(v string) *DescribeDeviceResponseBody {
	s.Vendor = &v
	return s
}

type DescribeDeviceResponseBodyDirectory struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDeviceResponseBodyDirectory) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBodyDirectory) SetCreatedTime(v string) *DescribeDeviceResponseBodyDirectory {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetDescription(v string) *DescribeDeviceResponseBodyDirectory {
	s.Description = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetGroupId(v string) *DescribeDeviceResponseBodyDirectory {
	s.GroupId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetId(v string) *DescribeDeviceResponseBodyDirectory {
	s.Id = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetName(v string) *DescribeDeviceResponseBodyDirectory {
	s.Name = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetParentId(v string) *DescribeDeviceResponseBodyDirectory {
	s.ParentId = &v
	return s
}

type DescribeDeviceResponseBodyStats struct {
	ChannelNum *int64 `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	FailedNum  *int64 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	OfflineNum *int64 `json:"OfflineNum,omitempty" xml:"OfflineNum,omitempty"`
	OnlineNum  *int64 `json:"OnlineNum,omitempty" xml:"OnlineNum,omitempty"`
	StreamNum  *int64 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
}

func (s DescribeDeviceResponseBodyStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceResponseBodyStats) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBodyStats) SetChannelNum(v int64) *DescribeDeviceResponseBodyStats {
	s.ChannelNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetFailedNum(v int64) *DescribeDeviceResponseBodyStats {
	s.FailedNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetOfflineNum(v int64) *DescribeDeviceResponseBodyStats {
	s.OfflineNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetOnlineNum(v int64) *DescribeDeviceResponseBodyStats {
	s.OnlineNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetStreamNum(v int64) *DescribeDeviceResponseBodyStats {
	s.StreamNum = &v
	return s
}

type DescribeDeviceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponse) SetHeaders(v map[string]*string) *DescribeDeviceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceResponse) SetBody(v *DescribeDeviceResponseBody) *DescribeDeviceResponse {
	s.Body = v
	return s
}

type DescribeDeviceChannelsRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum  *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDeviceChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsRequest) SetId(v string) *DescribeDeviceChannelsRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) SetOwnerId(v int64) *DescribeDeviceChannelsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) SetPageNum(v int64) *DescribeDeviceChannelsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDeviceChannelsRequest) SetPageSize(v int64) *DescribeDeviceChannelsRequest {
	s.PageSize = &v
	return s
}

type DescribeDeviceChannelsResponseBody struct {
	Channels   []*DescribeDeviceChannelsResponseBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	PageCount  *int64                                        `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsResponseBody) SetChannels(v []*DescribeDeviceChannelsResponseBodyChannels) *DescribeDeviceChannelsResponseBody {
	s.Channels = v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetPageCount(v int64) *DescribeDeviceChannelsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetPageNum(v int64) *DescribeDeviceChannelsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetPageSize(v int64) *DescribeDeviceChannelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetRequestId(v string) *DescribeDeviceChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBody) SetTotalCount(v int64) *DescribeDeviceChannelsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeviceChannelsResponseBodyChannels struct {
	ChannelId    *int64  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	GbId         *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	StreamId     *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	StreamStatus *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s DescribeDeviceChannelsResponseBodyChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceChannelsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetChannelId(v int64) *DescribeDeviceChannelsResponseBodyChannels {
	s.ChannelId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetDeviceId(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetDeviceStatus(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.DeviceStatus = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetGbId(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.GbId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetName(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.Name = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetParams(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.Params = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetStreamId(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.StreamId = &v
	return s
}

func (s *DescribeDeviceChannelsResponseBodyChannels) SetStreamStatus(v string) *DescribeDeviceChannelsResponseBodyChannels {
	s.StreamStatus = &v
	return s
}

type DescribeDeviceChannelsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsResponse) SetHeaders(v map[string]*string) *DescribeDeviceChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceChannelsResponse) SetBody(v *DescribeDeviceChannelsResponseBody) *DescribeDeviceChannelsResponse {
	s.Body = v
	return s
}

type DescribeDeviceGatewayRequest struct {
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Expire   *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDeviceGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceGatewayRequest) SetClientIp(v string) *DescribeDeviceGatewayRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) SetExpire(v int64) *DescribeDeviceGatewayRequest {
	s.Expire = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) SetId(v string) *DescribeDeviceGatewayRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) SetOwnerId(v int64) *DescribeDeviceGatewayRequest {
	s.OwnerId = &v
	return s
}

type DescribeDeviceGatewayResponseBody struct {
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port      *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol  *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeDeviceGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceGatewayResponseBody) SetHost(v string) *DescribeDeviceGatewayResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetPort(v int64) *DescribeDeviceGatewayResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetProtocol(v string) *DescribeDeviceGatewayResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetRequestId(v string) *DescribeDeviceGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetToken(v string) *DescribeDeviceGatewayResponseBody {
	s.Token = &v
	return s
}

type DescribeDeviceGatewayResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceGatewayResponse) SetHeaders(v map[string]*string) *DescribeDeviceGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceGatewayResponse) SetBody(v *DescribeDeviceGatewayResponseBody) *DescribeDeviceGatewayResponse {
	s.Body = v
	return s
}

type DescribeDeviceURLRequest struct {
	Auth        *bool   `json:"Auth,omitempty" xml:"Auth,omitempty"`
	Expire      *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Stream      *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeviceURLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceURLRequest) SetAuth(v bool) *DescribeDeviceURLRequest {
	s.Auth = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetExpire(v int64) *DescribeDeviceURLRequest {
	s.Expire = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetId(v string) *DescribeDeviceURLRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetMode(v string) *DescribeDeviceURLRequest {
	s.Mode = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetOutProtocol(v string) *DescribeDeviceURLRequest {
	s.OutProtocol = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetOwnerId(v int64) *DescribeDeviceURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetStream(v string) *DescribeDeviceURLRequest {
	s.Stream = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetType(v string) *DescribeDeviceURLRequest {
	s.Type = &v
	return s
}

type DescribeDeviceURLResponseBody struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDeviceURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceURLResponseBody) SetExpireTime(v int64) *DescribeDeviceURLResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDeviceURLResponseBody) SetRequestId(v string) *DescribeDeviceURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceURLResponseBody) SetUrl(v string) *DescribeDeviceURLResponseBody {
	s.Url = &v
	return s
}

type DescribeDeviceURLResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceURLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceURLResponse) SetHeaders(v map[string]*string) *DescribeDeviceURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceURLResponse) SetBody(v *DescribeDeviceURLResponseBody) *DescribeDeviceURLResponse {
	s.Body = v
	return s
}

type DescribeDevicesRequest struct {
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Dsn              *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IncludeDirectory *bool   `json:"IncludeDirectory,omitempty" xml:"IncludeDirectory,omitempty"`
	IncludeStats     *bool   `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentId         *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	SortBy           *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection    *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor           *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDevicesRequest) SetDirectoryId(v string) *DescribeDevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDevicesRequest) SetDsn(v string) *DescribeDevicesRequest {
	s.Dsn = &v
	return s
}

func (s *DescribeDevicesRequest) SetGbId(v string) *DescribeDevicesRequest {
	s.GbId = &v
	return s
}

func (s *DescribeDevicesRequest) SetGroupId(v string) *DescribeDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDevicesRequest) SetId(v string) *DescribeDevicesRequest {
	s.Id = &v
	return s
}

func (s *DescribeDevicesRequest) SetIncludeDirectory(v bool) *DescribeDevicesRequest {
	s.IncludeDirectory = &v
	return s
}

func (s *DescribeDevicesRequest) SetIncludeStats(v bool) *DescribeDevicesRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeDevicesRequest) SetName(v string) *DescribeDevicesRequest {
	s.Name = &v
	return s
}

func (s *DescribeDevicesRequest) SetOwnerId(v int64) *DescribeDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageNum(v int64) *DescribeDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageSize(v int64) *DescribeDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesRequest) SetParentId(v string) *DescribeDevicesRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDevicesRequest) SetSortBy(v string) *DescribeDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDevicesRequest) SetSortDirection(v string) *DescribeDevicesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeDevicesRequest) SetStatus(v string) *DescribeDevicesRequest {
	s.Status = &v
	return s
}

func (s *DescribeDevicesRequest) SetType(v string) *DescribeDevicesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDevicesRequest) SetVendor(v string) *DescribeDevicesRequest {
	s.Vendor = &v
	return s
}

type DescribeDevicesResponseBody struct {
	Devices    []*DescribeDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageCount  *int64                                `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBody) SetDevices(v []*DescribeDevicesResponseBodyDevices) *DescribeDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeDevicesResponseBody) SetPageCount(v int64) *DescribeDevicesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetPageNum(v int64) *DescribeDevicesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetPageSize(v int64) *DescribeDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetRequestId(v string) *DescribeDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetTotalCount(v int64) *DescribeDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDevicesResponseBodyDevices struct {
	AlarmMethod     *string                                      `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoDirectory   *bool                                        `json:"AutoDirectory,omitempty" xml:"AutoDirectory,omitempty"`
	AutoPos         *bool                                        `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	AutoStart       *bool                                        `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ChannelSyncTime *string                                      `json:"ChannelSyncTime,omitempty" xml:"ChannelSyncTime,omitempty"`
	CreatedTime     *string                                      `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description     *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory       *DescribeDevicesResponseBodyDevicesDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	DirectoryId     *string                                      `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Dsn             *string                                      `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	Enabled         *bool                                        `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GbId            *string                                      `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GroupId         *string                                      `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id              *string                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip              *string                                      `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Latitude        *string                                      `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude       *string                                      `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name            *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Params          *string                                      `json:"Params,omitempty" xml:"Params,omitempty"`
	ParentId        *string                                      `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Password        *string                                      `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int64                                       `json:"Port,omitempty" xml:"Port,omitempty"`
	PosInterval     *int64                                       `json:"PosInterval,omitempty" xml:"PosInterval,omitempty"`
	Protocol        *string                                      `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegisteredTime  *string                                      `json:"RegisteredTime,omitempty" xml:"RegisteredTime,omitempty"`
	Stats           *DescribeDevicesResponseBodyDevicesStats     `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	Status          *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	Url             *string                                      `json:"Url,omitempty" xml:"Url,omitempty"`
	Username        *string                                      `json:"Username,omitempty" xml:"Username,omitempty"`
	Vendor          *string                                      `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevices) SetAlarmMethod(v string) *DescribeDevicesResponseBodyDevices {
	s.AlarmMethod = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetAutoDirectory(v bool) *DescribeDevicesResponseBodyDevices {
	s.AutoDirectory = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetAutoPos(v bool) *DescribeDevicesResponseBodyDevices {
	s.AutoPos = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetAutoStart(v bool) *DescribeDevicesResponseBodyDevices {
	s.AutoStart = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetChannelSyncTime(v string) *DescribeDevicesResponseBodyDevices {
	s.ChannelSyncTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetCreatedTime(v string) *DescribeDevicesResponseBodyDevices {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDescription(v string) *DescribeDevicesResponseBodyDevices {
	s.Description = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDirectory(v *DescribeDevicesResponseBodyDevicesDirectory) *DescribeDevicesResponseBodyDevices {
	s.Directory = v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDirectoryId(v string) *DescribeDevicesResponseBodyDevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDsn(v string) *DescribeDevicesResponseBodyDevices {
	s.Dsn = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetEnabled(v bool) *DescribeDevicesResponseBodyDevices {
	s.Enabled = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetGbId(v string) *DescribeDevicesResponseBodyDevices {
	s.GbId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetGroupId(v string) *DescribeDevicesResponseBodyDevices {
	s.GroupId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetId(v string) *DescribeDevicesResponseBodyDevices {
	s.Id = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetIp(v string) *DescribeDevicesResponseBodyDevices {
	s.Ip = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetLatitude(v string) *DescribeDevicesResponseBodyDevices {
	s.Latitude = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetLongitude(v string) *DescribeDevicesResponseBodyDevices {
	s.Longitude = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetName(v string) *DescribeDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetParams(v string) *DescribeDevicesResponseBodyDevices {
	s.Params = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetParentId(v string) *DescribeDevicesResponseBodyDevices {
	s.ParentId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetPassword(v string) *DescribeDevicesResponseBodyDevices {
	s.Password = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetPort(v int64) *DescribeDevicesResponseBodyDevices {
	s.Port = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetPosInterval(v int64) *DescribeDevicesResponseBodyDevices {
	s.PosInterval = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetProtocol(v string) *DescribeDevicesResponseBodyDevices {
	s.Protocol = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetRegisteredTime(v string) *DescribeDevicesResponseBodyDevices {
	s.RegisteredTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetStats(v *DescribeDevicesResponseBodyDevicesStats) *DescribeDevicesResponseBodyDevices {
	s.Stats = v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetStatus(v string) *DescribeDevicesResponseBodyDevices {
	s.Status = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetType(v string) *DescribeDevicesResponseBodyDevices {
	s.Type = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetUrl(v string) *DescribeDevicesResponseBodyDevices {
	s.Url = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetUsername(v string) *DescribeDevicesResponseBodyDevices {
	s.Username = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetVendor(v string) *DescribeDevicesResponseBodyDevices {
	s.Vendor = &v
	return s
}

type DescribeDevicesResponseBodyDevicesDirectory struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDevicesResponseBodyDevicesDirectory) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevicesDirectory) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetCreatedTime(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetDescription(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.Description = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetGroupId(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.GroupId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetId(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.Id = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetName(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.Name = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetParentId(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.ParentId = &v
	return s
}

type DescribeDevicesResponseBodyDevicesStats struct {
	ChannelNum *int64 `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	FailedNum  *int64 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	OfflineNum *int64 `json:"OfflineNum,omitempty" xml:"OfflineNum,omitempty"`
	OnlineNum  *int64 `json:"OnlineNum,omitempty" xml:"OnlineNum,omitempty"`
	StreamNum  *int64 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
}

func (s DescribeDevicesResponseBodyDevicesStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevicesStats) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetChannelNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.ChannelNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetFailedNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.FailedNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetOfflineNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.OfflineNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetOnlineNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.OnlineNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetStreamNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.StreamNum = &v
	return s
}

type DescribeDevicesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponse) SetHeaders(v map[string]*string) *DescribeDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDevicesResponse) SetBody(v *DescribeDevicesResponseBody) *DescribeDevicesResponse {
	s.Body = v
	return s
}

type DescribeDirectoriesRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NoPagination  *bool   `json:"NoPagination,omitempty" xml:"NoPagination,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentId      *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetGroupId(v string) *DescribeDirectoriesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNoPagination(v bool) *DescribeDirectoriesRequest {
	s.NoPagination = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetOwnerId(v int64) *DescribeDirectoriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetPageNum(v int64) *DescribeDirectoriesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetPageSize(v int64) *DescribeDirectoriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetParentId(v string) *DescribeDirectoriesRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetSortBy(v string) *DescribeDirectoriesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetSortDirection(v string) *DescribeDirectoriesRequest {
	s.SortDirection = &v
	return s
}

type DescribeDirectoriesResponseBody struct {
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	PageCount   *int64                                        `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum     *int64                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int64                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetPageCount(v int64) *DescribeDirectoriesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetPageNum(v int64) *DescribeDirectoriesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetPageSize(v int64) *DescribeDirectoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetTotalCount(v int64) *DescribeDirectoriesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCreatedTime(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDescription(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Description = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.GroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Id = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetParentId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.ParentId = &v
	return s
}

type DescribeDirectoriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

type DescribeDirectoryRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoryRequest) SetId(v string) *DescribeDirectoryRequest {
	s.Id = &v
	return s
}

func (s *DescribeDirectoryRequest) SetOwnerId(v int64) *DescribeDirectoryRequest {
	s.OwnerId = &v
	return s
}

type DescribeDirectoryResponseBody struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoryResponseBody) SetCreatedTime(v string) *DescribeDirectoryResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetDescription(v string) *DescribeDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetGroupId(v string) *DescribeDirectoryResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetId(v string) *DescribeDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetName(v string) *DescribeDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetParentId(v string) *DescribeDirectoryResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetRequestId(v string) *DescribeDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDirectoryResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoryResponse) SetHeaders(v map[string]*string) *DescribeDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoryResponse) SetBody(v *DescribeDirectoryResponseBody) *DescribeDirectoryResponse {
	s.Body = v
	return s
}

type DescribeExternalStreamURLRequest struct {
	Kind    *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	TxId    *string `json:"TxId,omitempty" xml:"TxId,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeExternalStreamURLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalStreamURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeExternalStreamURLRequest) SetKind(v string) *DescribeExternalStreamURLRequest {
	s.Kind = &v
	return s
}

func (s *DescribeExternalStreamURLRequest) SetOwnerId(v int64) *DescribeExternalStreamURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExternalStreamURLRequest) SetProfile(v string) *DescribeExternalStreamURLRequest {
	s.Profile = &v
	return s
}

func (s *DescribeExternalStreamURLRequest) SetTxId(v string) *DescribeExternalStreamURLRequest {
	s.TxId = &v
	return s
}

func (s *DescribeExternalStreamURLRequest) SetUrl(v string) *DescribeExternalStreamURLRequest {
	s.Url = &v
	return s
}

type DescribeExternalStreamURLResponseBody struct {
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	Port        *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Profile     *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TxId        *string `json:"TxId,omitempty" xml:"TxId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeExternalStreamURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalStreamURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExternalStreamURLResponseBody) SetIp(v string) *DescribeExternalStreamURLResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeExternalStreamURLResponseBody) SetOutProtocol(v string) *DescribeExternalStreamURLResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *DescribeExternalStreamURLResponseBody) SetPort(v int64) *DescribeExternalStreamURLResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeExternalStreamURLResponseBody) SetProfile(v string) *DescribeExternalStreamURLResponseBody {
	s.Profile = &v
	return s
}

func (s *DescribeExternalStreamURLResponseBody) SetRequestId(v string) *DescribeExternalStreamURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExternalStreamURLResponseBody) SetTxId(v string) *DescribeExternalStreamURLResponseBody {
	s.TxId = &v
	return s
}

func (s *DescribeExternalStreamURLResponseBody) SetUrl(v string) *DescribeExternalStreamURLResponseBody {
	s.Url = &v
	return s
}

type DescribeExternalStreamURLResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExternalStreamURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExternalStreamURLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalStreamURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeExternalStreamURLResponse) SetHeaders(v map[string]*string) *DescribeExternalStreamURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeExternalStreamURLResponse) SetBody(v *DescribeExternalStreamURLResponseBody) *DescribeExternalStreamURLResponse {
	s.Body = v
	return s
}

type DescribeGroupRequest struct {
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IncludeStats *bool   `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupRequest) SetId(v string) *DescribeGroupRequest {
	s.Id = &v
	return s
}

func (s *DescribeGroupRequest) SetIncludeStats(v bool) *DescribeGroupRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeGroupRequest) SetOwnerId(v int64) *DescribeGroupRequest {
	s.OwnerId = &v
	return s
}

type DescribeGroupResponseBody struct {
	AliasId     *string                         `json:"AliasId,omitempty" xml:"AliasId,omitempty"`
	App         *string                         `json:"App,omitempty" xml:"App,omitempty"`
	Callback    *string                         `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CreatedTime *string                         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool                           `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GbId        *string                         `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GbIp        *string                         `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	GbPort      *int64                          `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	GbTcpPorts  []*string                       `json:"GbTcpPorts,omitempty" xml:"GbTcpPorts,omitempty" type:"Repeated"`
	GbUdpPorts  []*string                       `json:"GbUdpPorts,omitempty" xml:"GbUdpPorts,omitempty" type:"Repeated"`
	Id          *string                         `json:"Id,omitempty" xml:"Id,omitempty"`
	InProtocol  *string                         `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	LazyPull    *bool                           `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	Name        *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OutProtocol *string                         `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	PlayDomain  *string                         `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	PushDomain  *string                         `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	Region      *string                         `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stats       *DescribeGroupResponseBodyStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	Status      *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponseBody) SetAliasId(v string) *DescribeGroupResponseBody {
	s.AliasId = &v
	return s
}

func (s *DescribeGroupResponseBody) SetApp(v string) *DescribeGroupResponseBody {
	s.App = &v
	return s
}

func (s *DescribeGroupResponseBody) SetCallback(v string) *DescribeGroupResponseBody {
	s.Callback = &v
	return s
}

func (s *DescribeGroupResponseBody) SetCreatedTime(v string) *DescribeGroupResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeGroupResponseBody) SetDescription(v string) *DescribeGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeGroupResponseBody) SetEnabled(v bool) *DescribeGroupResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbId(v string) *DescribeGroupResponseBody {
	s.GbId = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbIp(v string) *DescribeGroupResponseBody {
	s.GbIp = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbPort(v int64) *DescribeGroupResponseBody {
	s.GbPort = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbTcpPorts(v []*string) *DescribeGroupResponseBody {
	s.GbTcpPorts = v
	return s
}

func (s *DescribeGroupResponseBody) SetGbUdpPorts(v []*string) *DescribeGroupResponseBody {
	s.GbUdpPorts = v
	return s
}

func (s *DescribeGroupResponseBody) SetId(v string) *DescribeGroupResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeGroupResponseBody) SetInProtocol(v string) *DescribeGroupResponseBody {
	s.InProtocol = &v
	return s
}

func (s *DescribeGroupResponseBody) SetLazyPull(v bool) *DescribeGroupResponseBody {
	s.LazyPull = &v
	return s
}

func (s *DescribeGroupResponseBody) SetName(v string) *DescribeGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGroupResponseBody) SetOutProtocol(v string) *DescribeGroupResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *DescribeGroupResponseBody) SetPlayDomain(v string) *DescribeGroupResponseBody {
	s.PlayDomain = &v
	return s
}

func (s *DescribeGroupResponseBody) SetPushDomain(v string) *DescribeGroupResponseBody {
	s.PushDomain = &v
	return s
}

func (s *DescribeGroupResponseBody) SetRegion(v string) *DescribeGroupResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeGroupResponseBody) SetRequestId(v string) *DescribeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupResponseBody) SetStats(v *DescribeGroupResponseBodyStats) *DescribeGroupResponseBody {
	s.Stats = v
	return s
}

func (s *DescribeGroupResponseBody) SetStatus(v string) *DescribeGroupResponseBody {
	s.Status = &v
	return s
}

type DescribeGroupResponseBodyStats struct {
	DeviceNum   *int64 `json:"DeviceNum,omitempty" xml:"DeviceNum,omitempty"`
	IedNum      *int64 `json:"IedNum,omitempty" xml:"IedNum,omitempty"`
	IpcNum      *int64 `json:"IpcNum,omitempty" xml:"IpcNum,omitempty"`
	PlatformNum *int64 `json:"PlatformNum,omitempty" xml:"PlatformNum,omitempty"`
}

func (s DescribeGroupResponseBodyStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupResponseBodyStats) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponseBodyStats) SetDeviceNum(v int64) *DescribeGroupResponseBodyStats {
	s.DeviceNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) SetIedNum(v int64) *DescribeGroupResponseBodyStats {
	s.IedNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) SetIpcNum(v int64) *DescribeGroupResponseBodyStats {
	s.IpcNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) SetPlatformNum(v int64) *DescribeGroupResponseBodyStats {
	s.PlatformNum = &v
	return s
}

type DescribeGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponse) SetHeaders(v map[string]*string) *DescribeGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupResponse) SetBody(v *DescribeGroupResponseBody) *DescribeGroupResponse {
	s.Body = v
	return s
}

type DescribeGroupsRequest struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InProtocol    *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	IncludeStats  *bool   `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupsRequest) SetId(v string) *DescribeGroupsRequest {
	s.Id = &v
	return s
}

func (s *DescribeGroupsRequest) SetInProtocol(v string) *DescribeGroupsRequest {
	s.InProtocol = &v
	return s
}

func (s *DescribeGroupsRequest) SetIncludeStats(v bool) *DescribeGroupsRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeGroupsRequest) SetName(v string) *DescribeGroupsRequest {
	s.Name = &v
	return s
}

func (s *DescribeGroupsRequest) SetOwnerId(v int64) *DescribeGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGroupsRequest) SetPageNum(v int64) *DescribeGroupsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeGroupsRequest) SetPageSize(v int64) *DescribeGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupsRequest) SetRegion(v string) *DescribeGroupsRequest {
	s.Region = &v
	return s
}

func (s *DescribeGroupsRequest) SetSortBy(v string) *DescribeGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeGroupsRequest) SetSortDirection(v string) *DescribeGroupsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeGroupsRequest) SetStatus(v string) *DescribeGroupsRequest {
	s.Status = &v
	return s
}

type DescribeGroupsResponseBody struct {
	Groups     []*DescribeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	PageCount  *int64                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBody) SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupsResponseBody) SetPageCount(v int64) *DescribeGroupsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetPageNum(v int64) *DescribeGroupsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetPageSize(v int64) *DescribeGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetRequestId(v string) *DescribeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetTotalCount(v int64) *DescribeGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGroupsResponseBodyGroups struct {
	AliasId     *string                                `json:"AliasId,omitempty" xml:"AliasId,omitempty"`
	App         *string                                `json:"App,omitempty" xml:"App,omitempty"`
	Callback    *string                                `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CreatedTime *string                                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool                                  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GbId        *string                                `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GbIp        *string                                `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	GbPort      *int64                                 `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	GbTcpPorts  []*string                              `json:"GbTcpPorts,omitempty" xml:"GbTcpPorts,omitempty" type:"Repeated"`
	GbUdpPorts  []*string                              `json:"GbUdpPorts,omitempty" xml:"GbUdpPorts,omitempty" type:"Repeated"`
	Id          *string                                `json:"Id,omitempty" xml:"Id,omitempty"`
	InProtocol  *string                                `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	LazyPull    *bool                                  `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	Name        *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	OutProtocol *string                                `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	PlayDomain  *string                                `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	PushDomain  *string                                `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	Region      *string                                `json:"Region,omitempty" xml:"Region,omitempty"`
	Stats       *DescribeGroupsResponseBodyGroupsStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	Status      *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroups) SetAliasId(v string) *DescribeGroupsResponseBodyGroups {
	s.AliasId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetApp(v string) *DescribeGroupsResponseBodyGroups {
	s.App = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetCallback(v string) *DescribeGroupsResponseBodyGroups {
	s.Callback = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetCreatedTime(v string) *DescribeGroupsResponseBodyGroups {
	s.CreatedTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetDescription(v string) *DescribeGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetEnabled(v bool) *DescribeGroupsResponseBodyGroups {
	s.Enabled = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbId(v string) *DescribeGroupsResponseBodyGroups {
	s.GbId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbIp(v string) *DescribeGroupsResponseBodyGroups {
	s.GbIp = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbPort(v int64) *DescribeGroupsResponseBodyGroups {
	s.GbPort = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbTcpPorts(v []*string) *DescribeGroupsResponseBodyGroups {
	s.GbTcpPorts = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbUdpPorts(v []*string) *DescribeGroupsResponseBodyGroups {
	s.GbUdpPorts = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetId(v string) *DescribeGroupsResponseBodyGroups {
	s.Id = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetInProtocol(v string) *DescribeGroupsResponseBodyGroups {
	s.InProtocol = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetLazyPull(v bool) *DescribeGroupsResponseBodyGroups {
	s.LazyPull = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetName(v string) *DescribeGroupsResponseBodyGroups {
	s.Name = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetOutProtocol(v string) *DescribeGroupsResponseBodyGroups {
	s.OutProtocol = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetPlayDomain(v string) *DescribeGroupsResponseBodyGroups {
	s.PlayDomain = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetPushDomain(v string) *DescribeGroupsResponseBodyGroups {
	s.PushDomain = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetRegion(v string) *DescribeGroupsResponseBodyGroups {
	s.Region = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetStats(v *DescribeGroupsResponseBodyGroupsStats) *DescribeGroupsResponseBodyGroups {
	s.Stats = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetStatus(v string) *DescribeGroupsResponseBodyGroups {
	s.Status = &v
	return s
}

type DescribeGroupsResponseBodyGroupsStats struct {
	DeviceNum   *int64 `json:"DeviceNum,omitempty" xml:"DeviceNum,omitempty"`
	IedNum      *int64 `json:"IedNum,omitempty" xml:"IedNum,omitempty"`
	IpcNum      *int64 `json:"IpcNum,omitempty" xml:"IpcNum,omitempty"`
	PlatformNum *int64 `json:"PlatformNum,omitempty" xml:"PlatformNum,omitempty"`
}

func (s DescribeGroupsResponseBodyGroupsStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroupsStats) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetDeviceNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.DeviceNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetIedNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.IedNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetIpcNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.IpcNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetPlatformNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.PlatformNum = &v
	return s
}

type DescribeGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponse) SetHeaders(v map[string]*string) *DescribeGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupsResponse) SetBody(v *DescribeGroupsResponseBody) *DescribeGroupsResponse {
	s.Body = v
	return s
}

type DescribeNodeDevicesInfoRequest struct {
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeNodeDevicesInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeDevicesInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeDevicesInfoRequest) SetNodeName(v string) *DescribeNodeDevicesInfoRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeDevicesInfoRequest) SetOwnerId(v int64) *DescribeNodeDevicesInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeNodeDevicesInfoResponseBody struct {
	NodeDevices []*DescribeNodeDevicesInfoResponseBodyNodeDevices `json:"NodeDevices,omitempty" xml:"NodeDevices,omitempty" type:"Repeated"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeDevicesInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeDevicesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeDevicesInfoResponseBody) SetNodeDevices(v []*DescribeNodeDevicesInfoResponseBodyNodeDevices) *DescribeNodeDevicesInfoResponseBody {
	s.NodeDevices = v
	return s
}

func (s *DescribeNodeDevicesInfoResponseBody) SetRequestId(v string) *DescribeNodeDevicesInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNodeDevicesInfoResponseBodyNodeDevices struct {
	DeviceInfos []*DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	NodeName    *string                                                      `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s DescribeNodeDevicesInfoResponseBodyNodeDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeDevicesInfoResponseBodyNodeDevices) GoString() string {
	return s.String()
}

func (s *DescribeNodeDevicesInfoResponseBodyNodeDevices) SetDeviceInfos(v []*DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) *DescribeNodeDevicesInfoResponseBodyNodeDevices {
	s.DeviceInfos = v
	return s
}

func (s *DescribeNodeDevicesInfoResponseBodyNodeDevices) SetNodeName(v string) *DescribeNodeDevicesInfoResponseBodyNodeDevices {
	s.NodeName = &v
	return s
}

type DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos struct {
	IP         *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Server     *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) GoString() string {
	return s.String()
}

func (s *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) SetIP(v string) *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos {
	s.IP = &v
	return s
}

func (s *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) SetInstanceId(v string) *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) SetName(v string) *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos {
	s.Name = &v
	return s
}

func (s *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos) SetServer(v string) *DescribeNodeDevicesInfoResponseBodyNodeDevicesDeviceInfos {
	s.Server = &v
	return s
}

type DescribeNodeDevicesInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNodeDevicesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeDevicesInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeDevicesInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeDevicesInfoResponse) SetHeaders(v map[string]*string) *DescribeNodeDevicesInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeDevicesInfoResponse) SetBody(v *DescribeNodeDevicesInfoResponseBody) *DescribeNodeDevicesInfoResponse {
	s.Body = v
	return s
}

type DescribeParentPlatformRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeParentPlatformRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformRequest) SetId(v string) *DescribeParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformRequest) SetOwnerId(v int64) *DescribeParentPlatformRequest {
	s.OwnerId = &v
	return s
}

type DescribeParentPlatformResponseBody struct {
	AutoStart      *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ClientAuth     *bool   `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	ClientGbId     *string `json:"ClientGbId,omitempty" xml:"ClientGbId,omitempty"`
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	ClientPort     *int64  `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	CreatedTime    *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port           *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol       *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeParentPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformResponseBody) SetAutoStart(v bool) *DescribeParentPlatformResponseBody {
	s.AutoStart = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientAuth(v bool) *DescribeParentPlatformResponseBody {
	s.ClientAuth = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientGbId(v string) *DescribeParentPlatformResponseBody {
	s.ClientGbId = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientIp(v string) *DescribeParentPlatformResponseBody {
	s.ClientIp = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientPassword(v string) *DescribeParentPlatformResponseBody {
	s.ClientPassword = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientPort(v int64) *DescribeParentPlatformResponseBody {
	s.ClientPort = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientUsername(v string) *DescribeParentPlatformResponseBody {
	s.ClientUsername = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetCreatedTime(v string) *DescribeParentPlatformResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetDescription(v string) *DescribeParentPlatformResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetGbId(v string) *DescribeParentPlatformResponseBody {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetId(v string) *DescribeParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetIp(v string) *DescribeParentPlatformResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetName(v string) *DescribeParentPlatformResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetPort(v int64) *DescribeParentPlatformResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetProtocol(v string) *DescribeParentPlatformResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetRequestId(v string) *DescribeParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetStatus(v string) *DescribeParentPlatformResponseBody {
	s.Status = &v
	return s
}

type DescribeParentPlatformResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParentPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformResponse) SetHeaders(v map[string]*string) *DescribeParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentPlatformResponse) SetBody(v *DescribeParentPlatformResponseBody) *DescribeParentPlatformResponse {
	s.Body = v
	return s
}

type DescribeParentPlatformDevicesRequest struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s DescribeParentPlatformDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesRequest) SetId(v string) *DescribeParentPlatformDevicesRequest {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetOwnerId(v int64) *DescribeParentPlatformDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetPageNum(v int64) *DescribeParentPlatformDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetPageSize(v int64) *DescribeParentPlatformDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetSortBy(v string) *DescribeParentPlatformDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeParentPlatformDevicesRequest) SetSortDirection(v string) *DescribeParentPlatformDevicesRequest {
	s.SortDirection = &v
	return s
}

type DescribeParentPlatformDevicesResponseBody struct {
	Devices    []*DescribeParentPlatformDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageCount  *int64                                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParentPlatformDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesResponseBody) SetDevices(v []*DescribeParentPlatformDevicesResponseBodyDevices) *DescribeParentPlatformDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetPageCount(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetPageNum(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetPageSize(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetRequestId(v string) *DescribeParentPlatformDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetTotalCount(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeParentPlatformDevicesResponseBodyDevices struct {
	GbId     *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeParentPlatformDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetGbId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetGroupId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.GroupId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetName(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetParentId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.ParentId = &v
	return s
}

type DescribeParentPlatformDevicesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParentPlatformDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParentPlatformDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesResponse) SetHeaders(v map[string]*string) *DescribeParentPlatformDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentPlatformDevicesResponse) SetBody(v *DescribeParentPlatformDevicesResponseBody) *DescribeParentPlatformDevicesResponse {
	s.Body = v
	return s
}

type DescribeParentPlatformsRequest struct {
	GbId          *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeParentPlatformsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformsRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsRequest) SetGbId(v string) *DescribeParentPlatformsRequest {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetOwnerId(v int64) *DescribeParentPlatformsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetPageNum(v int64) *DescribeParentPlatformsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetPageSize(v int64) *DescribeParentPlatformsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetSortBy(v string) *DescribeParentPlatformsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetSortDirection(v string) *DescribeParentPlatformsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeParentPlatformsRequest) SetStatus(v string) *DescribeParentPlatformsRequest {
	s.Status = &v
	return s
}

type DescribeParentPlatformsResponseBody struct {
	PageCount  *int64                                          `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platforms  []*DescribeParentPlatformsResponseBodyPlatforms `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParentPlatformsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsResponseBody) SetPageCount(v int64) *DescribeParentPlatformsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetPageNum(v int64) *DescribeParentPlatformsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetPageSize(v int64) *DescribeParentPlatformsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetPlatforms(v []*DescribeParentPlatformsResponseBodyPlatforms) *DescribeParentPlatformsResponseBody {
	s.Platforms = v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetRequestId(v string) *DescribeParentPlatformsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetTotalCount(v int64) *DescribeParentPlatformsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeParentPlatformsResponseBodyPlatforms struct {
	AutoStart      *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ClientAuth     *bool   `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	ClientGbId     *string `json:"ClientGbId,omitempty" xml:"ClientGbId,omitempty"`
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	ClientPort     *int64  `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	CreatedTime    *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port           *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol       *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeParentPlatformsResponseBodyPlatforms) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformsResponseBodyPlatforms) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetAutoStart(v bool) *DescribeParentPlatformsResponseBodyPlatforms {
	s.AutoStart = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientAuth(v bool) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientAuth = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientGbId(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientGbId = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientIp(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientIp = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientPassword(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientPassword = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientPort(v int64) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientPort = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientUsername(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientUsername = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetCreatedTime(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.CreatedTime = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetDescription(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Description = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetGbId(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetId(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetIp(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Ip = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetName(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Name = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetPort(v int64) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Port = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetProtocol(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Protocol = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetStatus(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Status = &v
	return s
}

type DescribeParentPlatformsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParentPlatformsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParentPlatformsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentPlatformsResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsResponse) SetHeaders(v map[string]*string) *DescribeParentPlatformsResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentPlatformsResponse) SetBody(v *DescribeParentPlatformsResponseBody) *DescribeParentPlatformsResponse {
	s.Body = v
	return s
}

type DescribePresetsRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribePresetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePresetsRequest) GoString() string {
	return s.String()
}

func (s *DescribePresetsRequest) SetId(v string) *DescribePresetsRequest {
	s.Id = &v
	return s
}

func (s *DescribePresetsRequest) SetOwnerId(v int64) *DescribePresetsRequest {
	s.OwnerId = &v
	return s
}

type DescribePresetsResponseBody struct {
	Id        *string                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Presets   []*DescribePresetsResponseBodyPresets `json:"Presets,omitempty" xml:"Presets,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePresetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePresetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePresetsResponseBody) SetId(v string) *DescribePresetsResponseBody {
	s.Id = &v
	return s
}

func (s *DescribePresetsResponseBody) SetPresets(v []*DescribePresetsResponseBodyPresets) *DescribePresetsResponseBody {
	s.Presets = v
	return s
}

func (s *DescribePresetsResponseBody) SetRequestId(v string) *DescribePresetsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePresetsResponseBodyPresets struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePresetsResponseBodyPresets) String() string {
	return tea.Prettify(s)
}

func (s DescribePresetsResponseBodyPresets) GoString() string {
	return s.String()
}

func (s *DescribePresetsResponseBodyPresets) SetId(v string) *DescribePresetsResponseBodyPresets {
	s.Id = &v
	return s
}

func (s *DescribePresetsResponseBodyPresets) SetName(v string) *DescribePresetsResponseBodyPresets {
	s.Name = &v
	return s
}

type DescribePresetsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePresetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePresetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePresetsResponse) GoString() string {
	return s.String()
}

func (s *DescribePresetsResponse) SetHeaders(v map[string]*string) *DescribePresetsResponse {
	s.Headers = v
	return s
}

func (s *DescribePresetsResponse) SetBody(v *DescribePresetsResponseBody) *DescribePresetsResponse {
	s.Body = v
	return s
}

type DescribePurchasedDeviceRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribePurchasedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDeviceRequest) SetId(v string) *DescribePurchasedDeviceRequest {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDeviceRequest) SetOwnerId(v int64) *DescribePurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

type DescribePurchasedDeviceResponseBody struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubType      *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribePurchasedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDeviceResponseBody) SetCreatedTime(v string) *DescribePurchasedDeviceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetDescription(v string) *DescribePurchasedDeviceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetGroupId(v string) *DescribePurchasedDeviceResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetGroupName(v string) *DescribePurchasedDeviceResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetId(v string) *DescribePurchasedDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetName(v string) *DescribePurchasedDeviceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetOrderId(v string) *DescribePurchasedDeviceResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetRegion(v string) *DescribePurchasedDeviceResponseBody {
	s.Region = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetRegisterCode(v string) *DescribePurchasedDeviceResponseBody {
	s.RegisterCode = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetRequestId(v string) *DescribePurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetSubType(v string) *DescribePurchasedDeviceResponseBody {
	s.SubType = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetType(v string) *DescribePurchasedDeviceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetVendor(v string) *DescribePurchasedDeviceResponseBody {
	s.Vendor = &v
	return s
}

type DescribePurchasedDeviceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePurchasedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDeviceResponse) SetHeaders(v map[string]*string) *DescribePurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedDeviceResponse) SetBody(v *DescribePurchasedDeviceResponseBody) *DescribePurchasedDeviceResponse {
	s.Body = v
	return s
}

type DescribePurchasedDevicesRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor        *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribePurchasedDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesRequest) SetGroupId(v string) *DescribePurchasedDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetId(v string) *DescribePurchasedDevicesRequest {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetName(v string) *DescribePurchasedDevicesRequest {
	s.Name = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetOwnerId(v int64) *DescribePurchasedDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetPageNum(v int64) *DescribePurchasedDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetPageSize(v int64) *DescribePurchasedDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetSortBy(v string) *DescribePurchasedDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetSortDirection(v string) *DescribePurchasedDevicesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetSubType(v string) *DescribePurchasedDevicesRequest {
	s.SubType = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetType(v string) *DescribePurchasedDevicesRequest {
	s.Type = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetVendor(v string) *DescribePurchasedDevicesRequest {
	s.Vendor = &v
	return s
}

type DescribePurchasedDevicesResponseBody struct {
	Devices    []*DescribePurchasedDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageCount  *int64                                         `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                         `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurchasedDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesResponseBody) SetDevices(v []*DescribePurchasedDevicesResponseBodyDevices) *DescribePurchasedDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetPageCount(v int64) *DescribePurchasedDevicesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetPageNum(v int64) *DescribePurchasedDevicesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetPageSize(v int64) *DescribePurchasedDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetRequestId(v string) *DescribePurchasedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetTotalCount(v int64) *DescribePurchasedDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePurchasedDevicesResponseBodyDevices struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	SubType      *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribePurchasedDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetCreatedTime(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetDescription(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Description = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetGroupId(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetGroupName(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetId(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetName(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetOrderId(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.OrderId = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetRegion(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Region = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetRegisterCode(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.RegisterCode = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetSubType(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.SubType = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetType(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Type = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetVendor(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Vendor = &v
	return s
}

type DescribePurchasedDevicesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePurchasedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePurchasedDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesResponse) SetHeaders(v map[string]*string) *DescribePurchasedDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedDevicesResponse) SetBody(v *DescribePurchasedDevicesResponseBody) *DescribePurchasedDevicesResponse {
	s.Body = v
	return s
}

type DescribeRecordsRequest struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrivateBucket *bool   `json:"PrivateBucket,omitempty" xml:"PrivateBucket,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamId      *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordsRequest) SetEndTime(v string) *DescribeRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordsRequest) SetOwnerId(v int64) *DescribeRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordsRequest) SetPageNum(v int64) *DescribeRecordsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordsRequest) SetPageSize(v int64) *DescribeRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordsRequest) SetPrivateBucket(v bool) *DescribeRecordsRequest {
	s.PrivateBucket = &v
	return s
}

func (s *DescribeRecordsRequest) SetSortBy(v string) *DescribeRecordsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeRecordsRequest) SetSortDirection(v string) *DescribeRecordsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeRecordsRequest) SetStartTime(v string) *DescribeRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordsRequest) SetStreamId(v string) *DescribeRecordsRequest {
	s.StreamId = &v
	return s
}

func (s *DescribeRecordsRequest) SetType(v string) *DescribeRecordsRequest {
	s.Type = &v
	return s
}

type DescribeRecordsResponseBody struct {
	NextStartTime *string                               `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	PageCount     *int64                                `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum       *int64                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records       []*DescribeRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId     *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponseBody) SetNextStartTime(v string) *DescribeRecordsResponseBody {
	s.NextStartTime = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetPageCount(v int64) *DescribeRecordsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetPageNum(v int64) *DescribeRecordsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetPageSize(v int64) *DescribeRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetRecords(v []*DescribeRecordsResponseBodyRecords) *DescribeRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeRecordsResponseBody) SetRequestId(v string) *DescribeRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetTotalCount(v int64) *DescribeRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRecordsResponseBodyRecords struct {
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileFormat  *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	Height      *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamId    *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Width       *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponseBodyRecords) SetEndTime(v string) *DescribeRecordsResponseBodyRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetFileFormat(v string) *DescribeRecordsResponseBodyRecords {
	s.FileFormat = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetHeight(v int64) *DescribeRecordsResponseBodyRecords {
	s.Height = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetId(v string) *DescribeRecordsResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetOssBucket(v string) *DescribeRecordsResponseBodyRecords {
	s.OssBucket = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetOssEndpoint(v string) *DescribeRecordsResponseBodyRecords {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetOssObject(v string) *DescribeRecordsResponseBodyRecords {
	s.OssObject = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetStartTime(v string) *DescribeRecordsResponseBodyRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetStreamId(v string) *DescribeRecordsResponseBodyRecords {
	s.StreamId = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetTemplateId(v string) *DescribeRecordsResponseBodyRecords {
	s.TemplateId = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetType(v string) *DescribeRecordsResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetUrl(v string) *DescribeRecordsResponseBodyRecords {
	s.Url = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetWidth(v int64) *DescribeRecordsResponseBodyRecords {
	s.Width = &v
	return s
}

type DescribeRecordsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponse) SetHeaders(v map[string]*string) *DescribeRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordsResponse) SetBody(v *DescribeRecordsResponseBody) *DescribeRecordsResponse {
	s.Body = v
	return s
}

type DescribeRenderingDevicesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRenderingDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesRequest) SetInstanceIds(v string) *DescribeRenderingDevicesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeRenderingDevicesRequest) SetOwnerId(v int64) *DescribeRenderingDevicesRequest {
	s.OwnerId = &v
	return s
}

type DescribeRenderingDevicesResponseBody struct {
	Devices   []*DescribeRenderingDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRenderingDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesResponseBody) SetDevices(v []*DescribeRenderingDevicesResponseBodyDevices) *DescribeRenderingDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeRenderingDevicesResponseBody) SetRequestId(v string) *DescribeRenderingDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBody) SetTotal(v int64) *DescribeRenderingDevicesResponseBody {
	s.Total = &v
	return s
}

type DescribeRenderingDevicesResponseBodyDevices struct {
	AutoRenew          *bool                                                  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod    *int32                                                 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClusterId          *string                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description        *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EdgeNodeName       *string                                                `json:"EdgeNodeName,omitempty" xml:"EdgeNodeName,omitempty"`
	ImageId            *string                                                `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string                                                `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceId         *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IpInfos            []*DescribeRenderingDevicesResponseBodyDevicesIpInfos  `json:"IpInfos,omitempty" xml:"IpInfos,omitempty" type:"Repeated"`
	MacAddress         *string                                                `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	Period             *int32                                                 `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string                                                `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PlatformType       *string                                                `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	PodInfos           []*DescribeRenderingDevicesResponseBodyDevicesPodInfos `json:"PodInfos,omitempty" xml:"PodInfos,omitempty" type:"Repeated"`
	ServerName         *string                                                `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	Specification      *string                                                `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status             *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRenderingDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetAutoRenew(v bool) *DescribeRenderingDevicesResponseBodyDevices {
	s.AutoRenew = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetAutoRenewPeriod(v int32) *DescribeRenderingDevicesResponseBodyDevices {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetClusterId(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.ClusterId = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetDescription(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.Description = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetEdgeNodeName(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.EdgeNodeName = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetImageId(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.ImageId = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetInstanceChargeType(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetInstanceId(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.InstanceId = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetInstanceName(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.InstanceName = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetIpInfos(v []*DescribeRenderingDevicesResponseBodyDevicesIpInfos) *DescribeRenderingDevicesResponseBodyDevices {
	s.IpInfos = v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetMacAddress(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.MacAddress = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetPeriod(v int32) *DescribeRenderingDevicesResponseBodyDevices {
	s.Period = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetPeriodUnit(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetPlatformType(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.PlatformType = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetPodInfos(v []*DescribeRenderingDevicesResponseBodyDevicesPodInfos) *DescribeRenderingDevicesResponseBodyDevices {
	s.PodInfos = v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetServerName(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.ServerName = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetSpecification(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.Specification = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevices) SetStatus(v string) *DescribeRenderingDevicesResponseBodyDevices {
	s.Status = &v
	return s
}

type DescribeRenderingDevicesResponseBodyDevicesIpInfos struct {
	ExternalIp   *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	ISP          *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InternalIp   *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NatType      *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
}

func (s DescribeRenderingDevicesResponseBodyDevicesIpInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesResponseBodyDevicesIpInfos) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetExternalIp(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.ExternalIp = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetExternalPort(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.ExternalPort = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetISP(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.ISP = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetInternalIp(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.InternalIp = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetInternalPort(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.InternalPort = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetIpProtocol(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.IpProtocol = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesIpInfos) SetNatType(v string) *DescribeRenderingDevicesResponseBodyDevicesIpInfos {
	s.NatType = &v
	return s
}

type DescribeRenderingDevicesResponseBodyDevicesPodInfos struct {
	Network []*DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Repeated"`
	PodId   *string                                                       `json:"PodId,omitempty" xml:"PodId,omitempty"`
	Status  *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRenderingDevicesResponseBodyDevicesPodInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesResponseBodyDevicesPodInfos) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfos) SetNetwork(v []*DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) *DescribeRenderingDevicesResponseBodyDevicesPodInfos {
	s.Network = v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfos) SetPodId(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfos {
	s.PodId = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfos) SetStatus(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfos {
	s.Status = &v
	return s
}

type DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork struct {
	ContainerPorts *string `json:"ContainerPorts,omitempty" xml:"ContainerPorts,omitempty"`
	ExternalIp     *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalIsp    *string `json:"ExternalIsp,omitempty" xml:"ExternalIsp,omitempty"`
	ExternalPorts  *string `json:"ExternalPorts,omitempty" xml:"ExternalPorts,omitempty"`
	OutgoingIp     *string `json:"OutgoingIp,omitempty" xml:"OutgoingIp,omitempty"`
	OutgoingIsp    *string `json:"OutgoingIsp,omitempty" xml:"OutgoingIsp,omitempty"`
}

func (s DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) SetContainerPorts(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork {
	s.ContainerPorts = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) SetExternalIp(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork {
	s.ExternalIp = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) SetExternalIsp(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork {
	s.ExternalIsp = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) SetExternalPorts(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork {
	s.ExternalPorts = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) SetOutgoingIp(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork {
	s.OutgoingIp = &v
	return s
}

func (s *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork) SetOutgoingIsp(v string) *DescribeRenderingDevicesResponseBodyDevicesPodInfosNetwork {
	s.OutgoingIsp = &v
	return s
}

type DescribeRenderingDevicesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRenderingDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRenderingDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenderingDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenderingDevicesResponse) SetHeaders(v map[string]*string) *DescribeRenderingDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenderingDevicesResponse) SetBody(v *DescribeRenderingDevicesResponseBody) *DescribeRenderingDevicesResponse {
	s.Body = v
	return s
}

type DescribeStreamRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamRequest) SetId(v string) *DescribeStreamRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamRequest) SetOwnerId(v int64) *DescribeStreamRequest {
	s.OwnerId = &v
	return s
}

type DescribeStreamResponseBody struct {
	App         *string `json:"App,omitempty" xml:"App,omitempty"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Height      *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlayDomain  *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	PushDomain  *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Width       *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamResponseBody) SetApp(v string) *DescribeStreamResponseBody {
	s.App = &v
	return s
}

func (s *DescribeStreamResponseBody) SetCreatedTime(v string) *DescribeStreamResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStreamResponseBody) SetDeviceId(v string) *DescribeStreamResponseBody {
	s.DeviceId = &v
	return s
}

func (s *DescribeStreamResponseBody) SetEnabled(v bool) *DescribeStreamResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeStreamResponseBody) SetGroupId(v string) *DescribeStreamResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeStreamResponseBody) SetHeight(v int32) *DescribeStreamResponseBody {
	s.Height = &v
	return s
}

func (s *DescribeStreamResponseBody) SetId(v string) *DescribeStreamResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeStreamResponseBody) SetName(v string) *DescribeStreamResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeStreamResponseBody) SetPlayDomain(v string) *DescribeStreamResponseBody {
	s.PlayDomain = &v
	return s
}

func (s *DescribeStreamResponseBody) SetProtocol(v string) *DescribeStreamResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeStreamResponseBody) SetPushDomain(v string) *DescribeStreamResponseBody {
	s.PushDomain = &v
	return s
}

func (s *DescribeStreamResponseBody) SetRequestId(v string) *DescribeStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamResponseBody) SetStatus(v string) *DescribeStreamResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStreamResponseBody) SetWidth(v int32) *DescribeStreamResponseBody {
	s.Width = &v
	return s
}

type DescribeStreamResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamResponse) SetHeaders(v map[string]*string) *DescribeStreamResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamResponse) SetBody(v *DescribeStreamResponseBody) *DescribeStreamResponse {
	s.Body = v
	return s
}

type DescribeStreamURLRequest struct {
	Auth        *bool   `json:"Auth,omitempty" xml:"Auth,omitempty"`
	AuthKey     *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Expire      *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Transcode   *string `json:"Transcode,omitempty" xml:"Transcode,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeStreamURLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamURLRequest) SetAuth(v bool) *DescribeStreamURLRequest {
	s.Auth = &v
	return s
}

func (s *DescribeStreamURLRequest) SetAuthKey(v string) *DescribeStreamURLRequest {
	s.AuthKey = &v
	return s
}

func (s *DescribeStreamURLRequest) SetEndTime(v int64) *DescribeStreamURLRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamURLRequest) SetExpire(v int64) *DescribeStreamURLRequest {
	s.Expire = &v
	return s
}

func (s *DescribeStreamURLRequest) SetId(v string) *DescribeStreamURLRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamURLRequest) SetOutProtocol(v string) *DescribeStreamURLRequest {
	s.OutProtocol = &v
	return s
}

func (s *DescribeStreamURLRequest) SetOwnerId(v int64) *DescribeStreamURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamURLRequest) SetStartTime(v int64) *DescribeStreamURLRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeStreamURLRequest) SetTranscode(v string) *DescribeStreamURLRequest {
	s.Transcode = &v
	return s
}

func (s *DescribeStreamURLRequest) SetType(v string) *DescribeStreamURLRequest {
	s.Type = &v
	return s
}

type DescribeStreamURLResponseBody struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeStreamURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamURLResponseBody) SetExpireTime(v int64) *DescribeStreamURLResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeStreamURLResponseBody) SetRequestId(v string) *DescribeStreamURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamURLResponseBody) SetUrl(v string) *DescribeStreamURLResponseBody {
	s.Url = &v
	return s
}

type DescribeStreamURLResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamURLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamURLResponse) SetHeaders(v map[string]*string) *DescribeStreamURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamURLResponse) SetBody(v *DescribeStreamURLResponseBody) *DescribeStreamURLResponse {
	s.Body = v
	return s
}

type DescribeStreamVodListRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStreamVodListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamVodListRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListRequest) SetEndTime(v int64) *DescribeStreamVodListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamVodListRequest) SetId(v string) *DescribeStreamVodListRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamVodListRequest) SetOwnerId(v int64) *DescribeStreamVodListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamVodListRequest) SetStartTime(v int64) *DescribeStreamVodListRequest {
	s.StartTime = &v
	return s
}

type DescribeStreamVodListResponseBody struct {
	Records   []*DescribeStreamVodListResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStreamVodListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamVodListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListResponseBody) SetRecords(v []*DescribeStreamVodListResponseBodyRecords) *DescribeStreamVodListResponseBody {
	s.Records = v
	return s
}

func (s *DescribeStreamVodListResponseBody) SetRequestId(v string) *DescribeStreamVodListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeStreamVodListResponseBodyRecords struct {
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStreamVodListResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamVodListResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListResponseBodyRecords) SetEndTime(v int64) *DescribeStreamVodListResponseBodyRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamVodListResponseBodyRecords) SetStartTime(v int64) *DescribeStreamVodListResponseBodyRecords {
	s.StartTime = &v
	return s
}

type DescribeStreamVodListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamVodListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamVodListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamVodListResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListResponse) SetHeaders(v map[string]*string) *DescribeStreamVodListResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamVodListResponse) SetBody(v *DescribeStreamVodListResponseBody) *DescribeStreamVodListResponse {
	s.Body = v
	return s
}

type DescribeStreamsRequest struct {
	App           *string `json:"App,omitempty" xml:"App,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentId      *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s DescribeStreamsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamsRequest) SetApp(v string) *DescribeStreamsRequest {
	s.App = &v
	return s
}

func (s *DescribeStreamsRequest) SetDeviceId(v string) *DescribeStreamsRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeStreamsRequest) SetDomain(v string) *DescribeStreamsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeStreamsRequest) SetGroupId(v string) *DescribeStreamsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeStreamsRequest) SetId(v string) *DescribeStreamsRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamsRequest) SetName(v string) *DescribeStreamsRequest {
	s.Name = &v
	return s
}

func (s *DescribeStreamsRequest) SetOwnerId(v int64) *DescribeStreamsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamsRequest) SetPageNum(v int64) *DescribeStreamsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeStreamsRequest) SetPageSize(v int64) *DescribeStreamsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamsRequest) SetParentId(v string) *DescribeStreamsRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeStreamsRequest) SetSortBy(v string) *DescribeStreamsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeStreamsRequest) SetSortDirection(v string) *DescribeStreamsRequest {
	s.SortDirection = &v
	return s
}

type DescribeStreamsResponseBody struct {
	PageCount  *int64                                `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Streams    []*DescribeStreamsResponseBodyStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStreamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBody) SetPageCount(v int64) *DescribeStreamsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetPageNum(v int64) *DescribeStreamsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetPageSize(v int64) *DescribeStreamsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetRequestId(v string) *DescribeStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetStreams(v []*DescribeStreamsResponseBodyStreams) *DescribeStreamsResponseBody {
	s.Streams = v
	return s
}

func (s *DescribeStreamsResponseBody) SetTotalCount(v int64) *DescribeStreamsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeStreamsResponseBodyStreams struct {
	App         *string `json:"App,omitempty" xml:"App,omitempty"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Height      *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlayDomain  *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	PushDomain  *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Width       *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeStreamsResponseBodyStreams) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponseBodyStreams) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBodyStreams) SetApp(v string) *DescribeStreamsResponseBodyStreams {
	s.App = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetCreatedTime(v string) *DescribeStreamsResponseBodyStreams {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetDeviceId(v string) *DescribeStreamsResponseBodyStreams {
	s.DeviceId = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetEnabled(v bool) *DescribeStreamsResponseBodyStreams {
	s.Enabled = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetGroupId(v string) *DescribeStreamsResponseBodyStreams {
	s.GroupId = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetHeight(v int32) *DescribeStreamsResponseBodyStreams {
	s.Height = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetId(v string) *DescribeStreamsResponseBodyStreams {
	s.Id = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetName(v string) *DescribeStreamsResponseBodyStreams {
	s.Name = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetPlayDomain(v string) *DescribeStreamsResponseBodyStreams {
	s.PlayDomain = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetProtocol(v string) *DescribeStreamsResponseBodyStreams {
	s.Protocol = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetPushDomain(v string) *DescribeStreamsResponseBodyStreams {
	s.PushDomain = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetStatus(v string) *DescribeStreamsResponseBodyStreams {
	s.Status = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreams) SetWidth(v int32) *DescribeStreamsResponseBodyStreams {
	s.Width = &v
	return s
}

type DescribeStreamsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponse) SetHeaders(v map[string]*string) *DescribeStreamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamsResponse) SetBody(v *DescribeStreamsResponseBody) *DescribeStreamsResponse {
	s.Body = v
	return s
}

type DescribeTemplateRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateRequest) SetId(v string) *DescribeTemplateRequest {
	s.Id = &v
	return s
}

func (s *DescribeTemplateRequest) SetOwnerId(v int64) *DescribeTemplateRequest {
	s.OwnerId = &v
	return s
}

type DescribeTemplateResponseBody struct {
	Callback      *string                                     `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CreatedTime   *string                                     `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	FileFormat    *string                                     `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	Flv           *string                                     `json:"Flv,omitempty" xml:"Flv,omitempty"`
	HlsM3u8       *string                                     `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	HlsTs         *string                                     `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	Id            *string                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval      *int64                                      `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JpgOnDemand   *string                                     `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	JpgOverwrite  *string                                     `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	JpgSequence   *string                                     `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	Mp4           *string                                     `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	Name          *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket     *string                                     `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint   *string                                     `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssFilePrefix *string                                     `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	Region        *string                                     `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Retention     *int64                                      `json:"Retention,omitempty" xml:"Retention,omitempty"`
	TransConfigs  []*DescribeTemplateResponseBodyTransConfigs `json:"TransConfigs,omitempty" xml:"TransConfigs,omitempty" type:"Repeated"`
	Trigger       *string                                     `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	Type          *string                                     `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResponseBody) SetCallback(v string) *DescribeTemplateResponseBody {
	s.Callback = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetCreatedTime(v string) *DescribeTemplateResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetDescription(v string) *DescribeTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetFileFormat(v string) *DescribeTemplateResponseBody {
	s.FileFormat = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetFlv(v string) *DescribeTemplateResponseBody {
	s.Flv = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetHlsM3u8(v string) *DescribeTemplateResponseBody {
	s.HlsM3u8 = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetHlsTs(v string) *DescribeTemplateResponseBody {
	s.HlsTs = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetId(v string) *DescribeTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetInterval(v int64) *DescribeTemplateResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetJpgOnDemand(v string) *DescribeTemplateResponseBody {
	s.JpgOnDemand = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetJpgOverwrite(v string) *DescribeTemplateResponseBody {
	s.JpgOverwrite = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetJpgSequence(v string) *DescribeTemplateResponseBody {
	s.JpgSequence = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetMp4(v string) *DescribeTemplateResponseBody {
	s.Mp4 = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetName(v string) *DescribeTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetOssBucket(v string) *DescribeTemplateResponseBody {
	s.OssBucket = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetOssEndpoint(v string) *DescribeTemplateResponseBody {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetOssFilePrefix(v string) *DescribeTemplateResponseBody {
	s.OssFilePrefix = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetRegion(v string) *DescribeTemplateResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetRequestId(v string) *DescribeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetRetention(v int64) *DescribeTemplateResponseBody {
	s.Retention = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetTransConfigs(v []*DescribeTemplateResponseBodyTransConfigs) *DescribeTemplateResponseBody {
	s.TransConfigs = v
	return s
}

func (s *DescribeTemplateResponseBody) SetTrigger(v string) *DescribeTemplateResponseBody {
	s.Trigger = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetType(v string) *DescribeTemplateResponseBody {
	s.Type = &v
	return s
}

type DescribeTemplateResponseBodyTransConfigs struct {
	Fps          *int64  `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop          *int64  `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height       *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VideoBitrate *int64  `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoCodec   *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	Width        *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeTemplateResponseBodyTransConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResponseBodyTransConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetFps(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Fps = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetGop(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Gop = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetHeight(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Height = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetId(v string) *DescribeTemplateResponseBodyTransConfigs {
	s.Id = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetName(v string) *DescribeTemplateResponseBodyTransConfigs {
	s.Name = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetVideoBitrate(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetVideoCodec(v string) *DescribeTemplateResponseBodyTransConfigs {
	s.VideoCodec = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetWidth(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Width = &v
	return s
}

type DescribeTemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResponse) SetHeaders(v map[string]*string) *DescribeTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResponse) SetBody(v *DescribeTemplateResponseBody) *DescribeTemplateResponse {
	s.Body = v
	return s
}

type DescribeTemplatesRequest struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum       *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) SetId(v string) *DescribeTemplatesRequest {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesRequest) SetInstanceId(v string) *DescribeTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplatesRequest) SetOwnerId(v int64) *DescribeTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageNum(v int64) *DescribeTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageSize(v int64) *DescribeTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesRequest) SetSortBy(v string) *DescribeTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeTemplatesRequest) SetSortDirection(v string) *DescribeTemplatesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeTemplatesRequest) SetType(v string) *DescribeTemplatesRequest {
	s.Type = &v
	return s
}

type DescribeTemplatesResponseBody struct {
	PageCount  *int64                                    `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum    *int64                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates  []*DescribeTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) SetPageCount(v int64) *DescribeTemplatesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageNum(v int64) *DescribeTemplatesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageSize(v int64) *DescribeTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetRequestId(v string) *DescribeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTotalCount(v int64) *DescribeTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTemplatesResponseBodyTemplates struct {
	Callback      *string                                               `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CreatedTime   *string                                               `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	FileFormat    *string                                               `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	Flv           *string                                               `json:"Flv,omitempty" xml:"Flv,omitempty"`
	HlsM3u8       *string                                               `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	HlsTs         *string                                               `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	Id            *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval      *int64                                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JpgOnDemand   *string                                               `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	JpgOverwrite  *string                                               `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	JpgSequence   *string                                               `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	Mp4           *string                                               `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	Name          *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket     *string                                               `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint   *string                                               `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssFilePrefix *string                                               `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	Region        *string                                               `json:"Region,omitempty" xml:"Region,omitempty"`
	Retention     *int64                                                `json:"Retention,omitempty" xml:"Retention,omitempty"`
	TransConfigs  []*DescribeTemplatesResponseBodyTemplatesTransConfigs `json:"TransConfigs,omitempty" xml:"TransConfigs,omitempty" type:"Repeated"`
	Trigger       *string                                               `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	Type          *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCallback(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Callback = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCreatedTime(v string) *DescribeTemplatesResponseBodyTemplates {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetFileFormat(v string) *DescribeTemplatesResponseBodyTemplates {
	s.FileFormat = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetFlv(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Flv = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetHlsM3u8(v string) *DescribeTemplatesResponseBodyTemplates {
	s.HlsM3u8 = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetHlsTs(v string) *DescribeTemplatesResponseBodyTemplates {
	s.HlsTs = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetInterval(v int64) *DescribeTemplatesResponseBodyTemplates {
	s.Interval = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetJpgOnDemand(v string) *DescribeTemplatesResponseBodyTemplates {
	s.JpgOnDemand = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetJpgOverwrite(v string) *DescribeTemplatesResponseBodyTemplates {
	s.JpgOverwrite = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetJpgSequence(v string) *DescribeTemplatesResponseBodyTemplates {
	s.JpgSequence = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetMp4(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Mp4 = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetName(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetOssBucket(v string) *DescribeTemplatesResponseBodyTemplates {
	s.OssBucket = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetOssEndpoint(v string) *DescribeTemplatesResponseBodyTemplates {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetOssFilePrefix(v string) *DescribeTemplatesResponseBodyTemplates {
	s.OssFilePrefix = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetRegion(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Region = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetRetention(v int64) *DescribeTemplatesResponseBodyTemplates {
	s.Retention = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTransConfigs(v []*DescribeTemplatesResponseBodyTemplatesTransConfigs) *DescribeTemplatesResponseBodyTemplates {
	s.TransConfigs = v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTrigger(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Trigger = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetType(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

type DescribeTemplatesResponseBodyTemplatesTransConfigs struct {
	Fps          *int64  `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop          *int64  `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height       *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VideoBitrate *int64  `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoCodec   *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	Width        *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Id           *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DescribeTemplatesResponseBodyTemplatesTransConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplatesTransConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetFps(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Fps = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetGop(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Gop = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetHeight(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Height = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetName(v string) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetVideoBitrate(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetVideoCodec(v string) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.VideoCodec = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetWidth(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Width = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetId(v string) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Id = &v
	return s
}

type DescribeTemplatesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponse) SetHeaders(v map[string]*string) *DescribeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplatesResponse) SetBody(v *DescribeTemplatesResponseBody) *DescribeTemplatesResponse {
	s.Body = v
	return s
}

type DescribeUserDevicesRequest struct {
	EnsInstanceIds *string `json:"EnsInstanceIds,omitempty" xml:"EnsInstanceIds,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ServerName     *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeUserDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDevicesRequest) SetEnsInstanceIds(v string) *DescribeUserDevicesRequest {
	s.EnsInstanceIds = &v
	return s
}

func (s *DescribeUserDevicesRequest) SetOwnerId(v int64) *DescribeUserDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDevicesRequest) SetServerName(v string) *DescribeUserDevicesRequest {
	s.ServerName = &v
	return s
}

type DescribeUserDevicesResponseBody struct {
	List      []*DescribeUserDevicesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDevicesResponseBody) SetList(v []*DescribeUserDevicesResponseBodyList) *DescribeUserDevicesResponseBody {
	s.List = v
	return s
}

func (s *DescribeUserDevicesResponseBody) SetRequestId(v string) *DescribeUserDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserDevicesResponseBodyList struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EdgeNodeName  *string `json:"EdgeNodeName,omitempty" xml:"EdgeNodeName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MacAddress    *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	MatrixId      *string `json:"MatrixId,omitempty" xml:"MatrixId,omitempty"`
	Server        *string `json:"Server,omitempty" xml:"Server,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserDevicesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDevicesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeUserDevicesResponseBodyList) SetAliUid(v string) *DescribeUserDevicesResponseBodyList {
	s.AliUid = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetEdgeNodeName(v string) *DescribeUserDevicesResponseBodyList {
	s.EdgeNodeName = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetInstanceId(v string) *DescribeUserDevicesResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetMacAddress(v string) *DescribeUserDevicesResponseBodyList {
	s.MacAddress = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetMatrixId(v string) *DescribeUserDevicesResponseBodyList {
	s.MatrixId = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetServer(v string) *DescribeUserDevicesResponseBodyList {
	s.Server = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetSpecification(v string) *DescribeUserDevicesResponseBodyList {
	s.Specification = &v
	return s
}

func (s *DescribeUserDevicesResponseBodyList) SetStatus(v string) *DescribeUserDevicesResponseBodyList {
	s.Status = &v
	return s
}

type DescribeUserDevicesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDevicesResponse) SetHeaders(v map[string]*string) *DescribeUserDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDevicesResponse) SetBody(v *DescribeUserDevicesResponseBody) *DescribeUserDevicesResponse {
	s.Body = v
	return s
}

type DescribeVodStreamURLRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TxId    *string `json:"TxId,omitempty" xml:"TxId,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeVodStreamURLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStreamURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodStreamURLRequest) SetOwnerId(v int64) *DescribeVodStreamURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodStreamURLRequest) SetTxId(v string) *DescribeVodStreamURLRequest {
	s.TxId = &v
	return s
}

func (s *DescribeVodStreamURLRequest) SetUrl(v string) *DescribeVodStreamURLRequest {
	s.Url = &v
	return s
}

type DescribeVodStreamURLResponseBody struct {
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	Port        *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TxId        *string `json:"TxId,omitempty" xml:"TxId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeVodStreamURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStreamURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodStreamURLResponseBody) SetOutProtocol(v string) *DescribeVodStreamURLResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetPort(v int64) *DescribeVodStreamURLResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetRequestId(v string) *DescribeVodStreamURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetTxId(v string) *DescribeVodStreamURLResponseBody {
	s.TxId = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetUrl(v string) *DescribeVodStreamURLResponseBody {
	s.Url = &v
	return s
}

type DescribeVodStreamURLResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodStreamURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodStreamURLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStreamURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodStreamURLResponse) SetHeaders(v map[string]*string) *DescribeVodStreamURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodStreamURLResponse) SetBody(v *DescribeVodStreamURLResponseBody) *DescribeVodStreamURLResponse {
	s.Body = v
	return s
}

type DescribeVsCertificateDetailRequest struct {
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateDetailRequest) SetCertName(v string) *DescribeVsCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeVsCertificateDetailRequest) SetOwnerId(v int64) *DescribeVsCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsCertificateDetailResponseBody struct {
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateDetailResponseBody) SetCert(v string) *DescribeVsCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetCertId(v int64) *DescribeVsCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetCertName(v string) *DescribeVsCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetKey(v string) *DescribeVsCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetRequestId(v string) *DescribeVsCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsCertificateDetailResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeVsCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsCertificateDetailResponse) SetBody(v *DescribeVsCertificateDetailResponseBody) *DescribeVsCertificateDetailResponse {
	s.Body = v
	return s
}

type DescribeVsCertificateListRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListRequest) SetDomainName(v string) *DescribeVsCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsCertificateListRequest) SetOwnerId(v int64) *DescribeVsCertificateListRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsCertificateListResponseBody struct {
	CertificateListModel *DescribeVsCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	RequestId            *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponseBody) SetCertificateListModel(v *DescribeVsCertificateListResponseBodyCertificateListModel) *DescribeVsCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeVsCertificateListResponseBody) SetRequestId(v string) *DescribeVsCertificateListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsCertificateListResponseBodyCertificateListModel struct {
	CertList []*DescribeVsCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	Count    *int32                                                               `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeVsCertificateListResponseBodyCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) SetCertList(v []*DescribeVsCertificateListResponseBodyCertificateListModelCertList) *DescribeVsCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeVsCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

type DescribeVsCertificateListResponseBodyCertificateListModelCertList struct {
	CertId      *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	LastTime    *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeVsCertificateListResponseBodyCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetCertId(v int64) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.CertId = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetCertName(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.CertName = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetCommon(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.Common = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetFingerprint(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.Fingerprint = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetIssuer(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.Issuer = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetLastTime(v int64) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.LastTime = &v
	return s
}

type DescribeVsCertificateListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponse) SetHeaders(v map[string]*string) *DescribeVsCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsCertificateListResponse) SetBody(v *DescribeVsCertificateListResponseBody) *DescribeVsCertificateListResponse {
	s.Body = v
	return s
}

type DescribeVsDevicesDataRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDevicesDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDevicesDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataRequest) SetEndTime(v string) *DescribeVsDevicesDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) SetGroupId(v string) *DescribeVsDevicesDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) SetOwnerId(v int64) *DescribeVsDevicesDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) SetStartTime(v string) *DescribeVsDevicesDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDevicesDataResponseBody struct {
	DevicesDataPerInterval *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval `json:"DevicesDataPerInterval,omitempty" xml:"DevicesDataPerInterval,omitempty" type:"Struct"`
	RequestId              *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDevicesDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDevicesDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponseBody) SetDevicesDataPerInterval(v *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) *DescribeVsDevicesDataResponseBody {
	s.DevicesDataPerInterval = v
	return s
}

func (s *DescribeVsDevicesDataResponseBody) SetRequestId(v string) *DescribeVsDevicesDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsDevicesDataResponseBodyDevicesDataPerInterval struct {
	DataModule []*DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) SetDataModule(v []*DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule struct {
	DevicesDataValue *string `json:"DevicesDataValue,omitempty" xml:"DevicesDataValue,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) SetDevicesDataValue(v string) *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule {
	s.DevicesDataValue = &v
	return s
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsDevicesDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDevicesDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDevicesDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDevicesDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponse) SetHeaders(v map[string]*string) *DescribeVsDevicesDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDevicesDataResponse) SetBody(v *DescribeVsDevicesDataResponseBody) *DescribeVsDevicesDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataRequest) SetDomainName(v string) *DescribeVsDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetEndTime(v string) *DescribeVsDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetInterval(v string) *DescribeVsDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetIspNameEn(v string) *DescribeVsDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeVsDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetOwnerId(v int64) *DescribeVsDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetStartTime(v string) *DescribeVsDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeVsDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeVsDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetDomainName(v string) *DescribeVsDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetEndTime(v string) *DescribeVsDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetRequestId(v string) *DescribeVsDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetStartTime(v string) *DescribeVsDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	BpsValue  *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsDomainBpsDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainBpsDataResponse) SetBody(v *DescribeVsDomainBpsDataResponseBody) *DescribeVsDomainBpsDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainCertificateInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsDomainCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoRequest) SetDomainName(v string) *DescribeVsDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeVsDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeVsDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainCertificateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeVsDomainCertificateInfoResponseBodyCertInfos) *DescribeVsDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeVsDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeVsDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertDomainName          *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpireTime          *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertLife                *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg                 *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SSLPub                  *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

type DescribeVsDomainCertificateInfoResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeVsDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponse) SetBody(v *DescribeVsDomainCertificateInfoResponseBody) *DescribeVsDomainCertificateInfoResponse {
	s.Body = v
	return s
}

type DescribeVsDomainConfigsRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsRequest) SetDomainName(v string) *DescribeVsDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainConfigsRequest) SetFunctionNames(v string) *DescribeVsDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeVsDomainConfigsRequest) SetOwnerId(v int64) *DescribeVsDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsDomainConfigsResponseBody struct {
	DomainConfigs []*DescribeVsDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponseBody) SetDomainConfigs(v []*DescribeVsDomainConfigsResponseBodyDomainConfigs) *DescribeVsDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeVsDomainConfigsResponseBody) SetRequestId(v string) *DescribeVsDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsDomainConfigsResponseBodyDomainConfigs struct {
	ConfigId     *string                                                         `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs []*DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	FunctionName *string                                                         `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Status       *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetConfigId(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetStatus(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

type DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

type DescribeVsDomainConfigsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeVsDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainConfigsResponse) SetBody(v *DescribeVsDomainConfigsResponseBody) *DescribeVsDomainConfigsResponse {
	s.Body = v
	return s
}

type DescribeVsDomainDetailRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailRequest) SetDomainName(v string) *DescribeVsDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainDetailRequest) SetOwnerId(v int64) *DescribeVsDomainDetailRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsDomainDetailResponseBody struct {
	DomainConfig *DescribeVsDomainDetailResponseBodyDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailResponseBody) SetDomainConfig(v *DescribeVsDomainDetailResponseBodyDomainConfig) *DescribeVsDomainDetailResponseBody {
	s.DomainConfig = v
	return s
}

func (s *DescribeVsDomainDetailResponseBody) SetRequestId(v string) *DescribeVsDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsDomainDetailResponseBodyDomainConfig struct {
	Cname        *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType   *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	GmtCreated   *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SSLProtocol  *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeVsDomainDetailResponseBodyDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainDetailResponseBodyDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetCname(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Cname = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDescription(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Description = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDomainName(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDomainStatus(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDomainType(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.DomainType = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetGmtCreated(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetGmtModified(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.GmtModified = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetRegion(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Region = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetSSLProtocol(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetScope(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Scope = &v
	return s
}

type DescribeVsDomainDetailResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeVsDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainDetailResponse) SetBody(v *DescribeVsDomainDetailResponseBody) *DescribeVsDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeVsDomainOnlineUserNumRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QueryTime  *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
}

func (s DescribeVsDomainOnlineUserNumRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumRequest) SetDomainName(v string) *DescribeVsDomainOnlineUserNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainOnlineUserNumRequest) SetOwnerId(v int64) *DescribeVsDomainOnlineUserNumRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainOnlineUserNumRequest) SetQueryTime(v string) *DescribeVsDomainOnlineUserNumRequest {
	s.QueryTime = &v
	return s
}

type DescribeVsDomainOnlineUserNumResponseBody struct {
	OnlineUserInfo *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo `json:"OnlineUserInfo,omitempty" xml:"OnlineUserInfo,omitempty" type:"Struct"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StreamCount    *int32                                                   `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	UserCount      *int32                                                   `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeVsDomainOnlineUserNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumResponseBody) SetOnlineUserInfo(v *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo) *DescribeVsDomainOnlineUserNumResponseBody {
	s.OnlineUserInfo = v
	return s
}

func (s *DescribeVsDomainOnlineUserNumResponseBody) SetRequestId(v string) *DescribeVsDomainOnlineUserNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainOnlineUserNumResponseBody) SetStreamCount(v int32) *DescribeVsDomainOnlineUserNumResponseBody {
	s.StreamCount = &v
	return s
}

func (s *DescribeVsDomainOnlineUserNumResponseBody) SetUserCount(v int32) *DescribeVsDomainOnlineUserNumResponseBody {
	s.UserCount = &v
	return s
}

type DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo struct {
	LiveStreamOnlineUserNumInfo []*DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo `json:"LiveStreamOnlineUserNumInfo,omitempty" xml:"LiveStreamOnlineUserNumInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo) SetLiveStreamOnlineUserNumInfo(v []*DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfo {
	s.LiveStreamOnlineUserNumInfo = v
	return s
}

type DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo struct {
	Infos      *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos `json:"Infos,omitempty" xml:"Infos,omitempty" type:"Struct"`
	StreamName *string                                                                                  `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) SetInfos(v *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.Infos = v
	return s
}

func (s *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) SetStreamName(v string) *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.StreamName = &v
	return s
}

type DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos struct {
	Info []*DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) SetInfo(v []*DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos {
	s.Info = v
	return s
}

type DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo struct {
	TranscodeTemplate *string `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty"`
	UserNumber        *int64  `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) SetTranscodeTemplate(v string) *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	s.TranscodeTemplate = &v
	return s
}

func (s *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) SetUserNumber(v int64) *DescribeVsDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	s.UserNumber = &v
	return s
}

type DescribeVsDomainOnlineUserNumResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainOnlineUserNumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainOnlineUserNumResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainOnlineUserNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainOnlineUserNumResponse) SetHeaders(v map[string]*string) *DescribeVsDomainOnlineUserNumResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainOnlineUserNumResponse) SetBody(v *DescribeVsDomainOnlineUserNumResponseBody) *DescribeVsDomainOnlineUserNumResponse {
	s.Body = v
	return s
}

type DescribeVsDomainPvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataRequest) SetDomainName(v string) *DescribeVsDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) SetEndTime(v string) *DescribeVsDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) SetOwnerId(v int64) *DescribeVsDomainPvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) SetStartTime(v string) *DescribeVsDomainPvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainPvDataResponseBody struct {
	DataInterval   *string                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvDataInterval *DescribeVsDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponseBody) SetDataInterval(v string) *DescribeVsDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetDomainName(v string) *DescribeVsDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetEndTime(v string) *DescribeVsDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetPvDataInterval(v *DescribeVsDomainPvDataResponseBodyPvDataInterval) *DescribeVsDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetRequestId(v string) *DescribeVsDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetStartTime(v string) *DescribeVsDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeVsDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainPvDataResponseBodyPvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeVsDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

type DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeVsDomainPvDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainPvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainPvDataResponse) SetBody(v *DescribeVsDomainPvDataResponseBody) *DescribeVsDomainPvDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainPvUvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvUvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataRequest) SetDomainName(v string) *DescribeVsDomainPvUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) SetEndTime(v string) *DescribeVsDomainPvUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) SetOwnerId(v int64) *DescribeVsDomainPvUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) SetStartTime(v string) *DescribeVsDomainPvUvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainPvUvDataResponseBody struct {
	DataInterval  *string                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName    *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvUvDataInfos *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos `json:"PvUvDataInfos,omitempty" xml:"PvUvDataInfos,omitempty" type:"Struct"`
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime     *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvUvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetDataInterval(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetDomainName(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetEndTime(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetPvUvDataInfos(v *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) *DescribeVsDomainPvUvDataResponseBody {
	s.PvUvDataInfos = v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetRequestId(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetStartTime(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos struct {
	PvUvDataInfo []*DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo `json:"PvUvDataInfo,omitempty" xml:"PvUvDataInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) SetPvUvDataInfo(v []*DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos {
	s.PvUvDataInfo = v
	return s
}

type DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo struct {
	PV        *string `json:"PV,omitempty" xml:"PV,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UV        *string `json:"UV,omitempty" xml:"UV,omitempty"`
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetPV(v string) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.PV = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetTimeStamp(v string) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetUV(v string) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.UV = &v
	return s
}

type DescribeVsDomainPvUvDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainPvUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainPvUvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainPvUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainPvUvDataResponse) SetBody(v *DescribeVsDomainPvUvDataResponseBody) *DescribeVsDomainPvUvDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainRecordDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainRecordDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRecordDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataRequest) SetDomainName(v string) *DescribeVsDomainRecordDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetEndTime(v string) *DescribeVsDomainRecordDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetOwnerId(v int64) *DescribeVsDomainRecordDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetRegion(v string) *DescribeVsDomainRecordDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetStartTime(v string) *DescribeVsDomainRecordDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainRecordDataResponseBody struct {
	RecordDataPerInterval *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval `json:"RecordDataPerInterval,omitempty" xml:"RecordDataPerInterval,omitempty" type:"Struct"`
	RequestId             *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainRecordDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponseBody) SetRecordDataPerInterval(v *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) *DescribeVsDomainRecordDataResponseBody {
	s.RecordDataPerInterval = v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBody) SetRequestId(v string) *DescribeVsDomainRecordDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval struct {
	DataModule []*DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) SetDataModule(v []*DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule struct {
	RecordValue      *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	StreamCountValue *string `json:"StreamCountValue,omitempty" xml:"StreamCountValue,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) SetRecordValue(v string) *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	s.RecordValue = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) SetStreamCountValue(v string) *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	s.StreamCountValue = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsDomainRecordDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainRecordDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainRecordDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainRecordDataResponse) SetBody(v *DescribeVsDomainRecordDataResponseBody) *DescribeVsDomainRecordDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainRegionDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainRegionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataRequest) SetDomainName(v string) *DescribeVsDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) SetEndTime(v string) *DescribeVsDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) SetOwnerId(v int64) *DescribeVsDomainRegionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) SetStartTime(v string) *DescribeVsDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainRegionDataResponseBody struct {
	DataInterval *string                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeVsDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeVsDomainRegionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeVsDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetDomainName(v string) *DescribeVsDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetEndTime(v string) *DescribeVsDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetRequestId(v string) *DescribeVsDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetStartTime(v string) *DescribeVsDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetValue(v *DescribeVsDomainRegionDataResponseBodyValue) *DescribeVsDomainRegionDataResponseBody {
	s.Value = v
	return s
}

type DescribeVsDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeVsDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainRegionDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) *DescribeVsDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

type DescribeVsDomainRegionDataResponseBodyValueRegionProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionEname     *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeVsDomainRegionDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainRegionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainRegionDataResponse) SetBody(v *DescribeVsDomainRegionDataResponseBody) *DescribeVsDomainRegionDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainReqBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataRequest) SetDomainName(v string) *DescribeVsDomainReqBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetEndTime(v string) *DescribeVsDomainReqBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetInterval(v string) *DescribeVsDomainReqBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetIspNameEn(v string) *DescribeVsDomainReqBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetLocationNameEn(v string) *DescribeVsDomainReqBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetOwnerId(v int64) *DescribeVsDomainReqBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetStartTime(v string) *DescribeVsDomainReqBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainReqBpsDataResponseBody struct {
	DataInterval          *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName            *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime               *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReqBpsDataPerInterval *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval `json:"ReqBpsDataPerInterval,omitempty" xml:"ReqBpsDataPerInterval,omitempty" type:"Struct"`
	RequestId             *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime             *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetDataInterval(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetDomainName(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetEndTime(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetReqBpsDataPerInterval(v *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) *DescribeVsDomainReqBpsDataResponseBody {
	s.ReqBpsDataPerInterval = v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetRequestId(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetStartTime(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval struct {
	DataModule []*DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) SetDataModule(v []*DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule struct {
	ReqBpsValue *string `json:"ReqBpsValue,omitempty" xml:"ReqBpsValue,omitempty"`
	TimeStamp   *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) SetReqBpsValue(v string) *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule {
	s.ReqBpsValue = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsDomainReqBpsDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainReqBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainReqBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainReqBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponse) SetBody(v *DescribeVsDomainReqBpsDataResponseBody) *DescribeVsDomainReqBpsDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainReqTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetDomainName(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetEndTime(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetInterval(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetIspNameEn(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetLocationNameEn(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetOwnerId(v int64) *DescribeVsDomainReqTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetStartTime(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainReqTrafficDataResponseBody struct {
	DataInterval              *string                                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                *string                                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                   *string                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReqTrafficDataPerInterval *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval `json:"ReqTrafficDataPerInterval,omitempty" xml:"ReqTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                 *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                 *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetDataInterval(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetDomainName(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetEndTime(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetReqTrafficDataPerInterval(v *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) *DescribeVsDomainReqTrafficDataResponseBody {
	s.ReqTrafficDataPerInterval = v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetRequestId(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetStartTime(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval struct {
	DataModule []*DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) SetDataModule(v []*DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule struct {
	ReqTrafficValue *string `json:"ReqTrafficValue,omitempty" xml:"ReqTrafficValue,omitempty"`
	TimeStamp       *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) SetReqTrafficValue(v string) *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule {
	s.ReqTrafficValue = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsDomainReqTrafficDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainReqTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainReqTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainReqTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponse) SetBody(v *DescribeVsDomainReqTrafficDataResponseBody) *DescribeVsDomainReqTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainSnapshotDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainSnapshotDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataRequest) SetDomainName(v string) *DescribeVsDomainSnapshotDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) SetEndTime(v string) *DescribeVsDomainSnapshotDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) SetOwnerId(v int64) *DescribeVsDomainSnapshotDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) SetStartTime(v string) *DescribeVsDomainSnapshotDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainSnapshotDataResponseBody struct {
	RequestId               *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotDataPerInterval *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval `json:"SnapshotDataPerInterval,omitempty" xml:"SnapshotDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVsDomainSnapshotDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponseBody) SetRequestId(v string) *DescribeVsDomainSnapshotDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBody) SetSnapshotDataPerInterval(v *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) *DescribeVsDomainSnapshotDataResponseBody {
	s.SnapshotDataPerInterval = v
	return s
}

type DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval struct {
	DataModule []*DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) SetDataModule(v []*DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule struct {
	SnapshotValue *string `json:"SnapshotValue,omitempty" xml:"SnapshotValue,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) SetSnapshotValue(v string) *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule {
	s.SnapshotValue = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsDomainSnapshotDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainSnapshotDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainSnapshotDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainSnapshotDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponse) SetBody(v *DescribeVsDomainSnapshotDataResponseBody) *DescribeVsDomainSnapshotDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataRequest) SetDomainName(v string) *DescribeVsDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetEndTime(v string) *DescribeVsDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetInterval(v string) *DescribeVsDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeVsDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeVsDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeVsDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetStartTime(v string) *DescribeVsDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainTrafficDataResponseBody struct {
	DataInterval           *string                                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVsDomainTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeVsDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

type DescribeVsDomainTrafficDataResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainTrafficDataResponse) SetBody(v *DescribeVsDomainTrafficDataResponseBody) *DescribeVsDomainTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeVsDomainUvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainUvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataRequest) SetDomainName(v string) *DescribeVsDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) SetEndTime(v string) *DescribeVsDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) SetOwnerId(v int64) *DescribeVsDomainUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) SetStartTime(v string) *DescribeVsDomainUvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsDomainUvDataResponseBody struct {
	DataInterval   *string                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UvDataInterval *DescribeVsDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeVsDomainUvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponseBody) SetDataInterval(v string) *DescribeVsDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetDomainName(v string) *DescribeVsDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetEndTime(v string) *DescribeVsDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetRequestId(v string) *DescribeVsDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetStartTime(v string) *DescribeVsDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetUvDataInterval(v *DescribeVsDomainUvDataResponseBodyUvDataInterval) *DescribeVsDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

type DescribeVsDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainUvDataResponseBodyUvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeVsDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

type DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeVsDomainUvDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsDomainUvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainUvDataResponse) SetBody(v *DescribeVsDomainUvDataResponseBody) *DescribeVsDomainUvDataResponse {
	s.Body = v
	return s
}

type DescribeVsPullStreamConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsPullStreamConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamConfigRequest) SetDomainName(v string) *DescribeVsPullStreamConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsPullStreamConfigRequest) SetOwnerId(v int64) *DescribeVsPullStreamConfigRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsPullStreamConfigResponseBody struct {
	LiveAppRecordList *DescribeVsPullStreamConfigResponseBodyLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" type:"Struct"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsPullStreamConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamConfigResponseBody) SetLiveAppRecordList(v *DescribeVsPullStreamConfigResponseBodyLiveAppRecordList) *DescribeVsPullStreamConfigResponseBody {
	s.LiveAppRecordList = v
	return s
}

func (s *DescribeVsPullStreamConfigResponseBody) SetRequestId(v string) *DescribeVsPullStreamConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsPullStreamConfigResponseBodyLiveAppRecordList struct {
	LiveAppRecord []*DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" type:"Repeated"`
}

func (s DescribeVsPullStreamConfigResponseBodyLiveAppRecordList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamConfigResponseBodyLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordList) SetLiveAppRecord(v []*DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

type DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetSourceUrl(v string) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.SourceUrl = &v
	return s
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeVsPullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

type DescribeVsPullStreamConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsPullStreamConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsPullStreamConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamConfigResponse) SetHeaders(v map[string]*string) *DescribeVsPullStreamConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsPullStreamConfigResponse) SetBody(v *DescribeVsPullStreamConfigResponseBody) *DescribeVsPullStreamConfigResponse {
	s.Body = v
	return s
}

type DescribeVsPullStreamInfoConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigRequest) SetDomainName(v string) *DescribeVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *DescribeVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsPullStreamInfoConfigResponseBody struct {
	LiveAppRecordList *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" type:"Struct"`
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) SetLiveAppRecordList(v *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) *DescribeVsPullStreamInfoConfigResponseBody {
	s.LiveAppRecordList = v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *DescribeVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList struct {
	LiveAppRecord []*DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" type:"Repeated"`
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) SetLiveAppRecord(v []*DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

type DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetSourceUrl(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.SourceUrl = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

type DescribeVsPullStreamInfoConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsPullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *DescribeVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponse) SetBody(v *DescribeVsPullStreamInfoConfigResponseBody) *DescribeVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

type DescribeVsStorageTrafficUsageDataRequest struct {
	Bucket    *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SplitBy   *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsStorageTrafficUsageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageTrafficUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageTrafficUsageDataRequest) SetBucket(v string) *DescribeVsStorageTrafficUsageDataRequest {
	s.Bucket = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataRequest) SetEndTime(v string) *DescribeVsStorageTrafficUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataRequest) SetInterval(v string) *DescribeVsStorageTrafficUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataRequest) SetOwnerId(v int64) *DescribeVsStorageTrafficUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataRequest) SetSplitBy(v string) *DescribeVsStorageTrafficUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataRequest) SetStartTime(v string) *DescribeVsStorageTrafficUsageDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsStorageTrafficUsageDataResponseBody struct {
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficUsage *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage `json:"TrafficUsage,omitempty" xml:"TrafficUsage,omitempty" type:"Struct"`
}

func (s DescribeVsStorageTrafficUsageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageTrafficUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageTrafficUsageDataResponseBody) SetRequestId(v string) *DescribeVsStorageTrafficUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBody) SetTrafficUsage(v *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage) *DescribeVsStorageTrafficUsageDataResponseBody {
	s.TrafficUsage = v
	return s
}

type DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage struct {
	TrafficUsageDataModule []*DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule `json:"TrafficUsageDataModule,omitempty" xml:"TrafficUsageDataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage) SetTrafficUsageDataModule(v []*DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsage {
	s.TrafficUsageDataModule = v
	return s
}

type DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule struct {
	Bucket                   *string  `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	LanBandwidthInDataValue  *float32 `json:"LanBandwidthInDataValue,omitempty" xml:"LanBandwidthInDataValue,omitempty"`
	LanBandwidthOutDataValue *float32 `json:"LanBandwidthOutDataValue,omitempty" xml:"LanBandwidthOutDataValue,omitempty"`
	LanTrafficInDataValue    *int64   `json:"LanTrafficInDataValue,omitempty" xml:"LanTrafficInDataValue,omitempty"`
	LanTrafficOutDataValue   *int64   `json:"LanTrafficOutDataValue,omitempty" xml:"LanTrafficOutDataValue,omitempty"`
	TimeStamp                *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	WanBandwidthInDataValue  *float32 `json:"WanBandwidthInDataValue,omitempty" xml:"WanBandwidthInDataValue,omitempty"`
	WanBandwidthOutDataValue *float32 `json:"WanBandwidthOutDataValue,omitempty" xml:"WanBandwidthOutDataValue,omitempty"`
	WanTrafficInDataValue    *int64   `json:"WanTrafficInDataValue,omitempty" xml:"WanTrafficInDataValue,omitempty"`
	WanTrafficOutDataValue   *int64   `json:"WanTrafficOutDataValue,omitempty" xml:"WanTrafficOutDataValue,omitempty"`
}

func (s DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetBucket(v string) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.Bucket = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetLanBandwidthInDataValue(v float32) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.LanBandwidthInDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetLanBandwidthOutDataValue(v float32) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.LanBandwidthOutDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetLanTrafficInDataValue(v int64) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.LanTrafficInDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetLanTrafficOutDataValue(v int64) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.LanTrafficOutDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetTimeStamp(v string) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetWanBandwidthInDataValue(v float32) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.WanBandwidthInDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetWanBandwidthOutDataValue(v float32) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.WanBandwidthOutDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetWanTrafficInDataValue(v int64) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.WanTrafficInDataValue = &v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule) SetWanTrafficOutDataValue(v int64) *DescribeVsStorageTrafficUsageDataResponseBodyTrafficUsageTrafficUsageDataModule {
	s.WanTrafficOutDataValue = &v
	return s
}

type DescribeVsStorageTrafficUsageDataResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsStorageTrafficUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsStorageTrafficUsageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageTrafficUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageTrafficUsageDataResponse) SetHeaders(v map[string]*string) *DescribeVsStorageTrafficUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStorageTrafficUsageDataResponse) SetBody(v *DescribeVsStorageTrafficUsageDataResponseBody) *DescribeVsStorageTrafficUsageDataResponse {
	s.Body = v
	return s
}

type DescribeVsStorageUsageDataRequest struct {
	Bucket    *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SplitBy   *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsStorageUsageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageUsageDataRequest) SetBucket(v string) *DescribeVsStorageUsageDataRequest {
	s.Bucket = &v
	return s
}

func (s *DescribeVsStorageUsageDataRequest) SetEndTime(v string) *DescribeVsStorageUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsStorageUsageDataRequest) SetInterval(v string) *DescribeVsStorageUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsStorageUsageDataRequest) SetOwnerId(v int64) *DescribeVsStorageUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStorageUsageDataRequest) SetSplitBy(v string) *DescribeVsStorageUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeVsStorageUsageDataRequest) SetStartTime(v string) *DescribeVsStorageUsageDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsStorageUsageDataResponseBody struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageUsage *DescribeVsStorageUsageDataResponseBodyStorageUsage `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty" type:"Struct"`
}

func (s DescribeVsStorageUsageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageUsageDataResponseBody) SetRequestId(v string) *DescribeVsStorageUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStorageUsageDataResponseBody) SetStorageUsage(v *DescribeVsStorageUsageDataResponseBodyStorageUsage) *DescribeVsStorageUsageDataResponseBody {
	s.StorageUsage = v
	return s
}

type DescribeVsStorageUsageDataResponseBodyStorageUsage struct {
	StorageUsageDataModule []*DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule `json:"StorageUsageDataModule,omitempty" xml:"StorageUsageDataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsStorageUsageDataResponseBodyStorageUsage) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageUsageDataResponseBodyStorageUsage) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageUsageDataResponseBodyStorageUsage) SetStorageUsageDataModule(v []*DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule) *DescribeVsStorageUsageDataResponseBodyStorageUsage {
	s.StorageUsageDataModule = v
	return s
}

type DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule struct {
	Bucket           *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	StorageDataValue *int32  `json:"StorageDataValue,omitempty" xml:"StorageDataValue,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule) SetBucket(v string) *DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule {
	s.Bucket = &v
	return s
}

func (s *DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule) SetStorageDataValue(v int32) *DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule {
	s.StorageDataValue = &v
	return s
}

func (s *DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule) SetTimeStamp(v string) *DescribeVsStorageUsageDataResponseBodyStorageUsageStorageUsageDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeVsStorageUsageDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsStorageUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsStorageUsageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStorageUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStorageUsageDataResponse) SetHeaders(v map[string]*string) *DescribeVsStorageUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStorageUsageDataResponse) SetBody(v *DescribeVsStorageUsageDataResponseBody) *DescribeVsStorageUsageDataResponse {
	s.Body = v
	return s
}

type DescribeVsStreamsNotifyUrlConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DescribeVsStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *DescribeVsStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

type DescribeVsStreamsNotifyUrlConfigResponseBody struct {
	LiveStreamsNotifyConfig *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig `json:"LiveStreamsNotifyConfig,omitempty" xml:"LiveStreamsNotifyConfig,omitempty" type:"Struct"`
	RequestId               *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) SetLiveStreamsNotifyConfig(v *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) *DescribeVsStreamsNotifyUrlConfigResponseBody {
	s.LiveStreamsNotifyConfig = v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *DescribeVsStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig struct {
	AuthKey    *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthType   *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	NotifyUrl  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetAuthKey(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.AuthKey = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetAuthType(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.AuthType = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetDomainName(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetNotifyUrl(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.NotifyUrl = &v
	return s
}

type DescribeVsStreamsNotifyUrlConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsStreamsNotifyUrlConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *DescribeVsStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) SetBody(v *DescribeVsStreamsNotifyUrlConfigResponseBody) *DescribeVsStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

type DescribeVsStreamsOnlineListRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum    *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DescribeVsStreamsOnlineListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsOnlineListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListRequest) SetAppName(v string) *DescribeVsStreamsOnlineListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetDomainName(v string) *DescribeVsStreamsOnlineListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetEndTime(v string) *DescribeVsStreamsOnlineListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetOrderBy(v string) *DescribeVsStreamsOnlineListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetOwnerId(v int64) *DescribeVsStreamsOnlineListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetPageNum(v int32) *DescribeVsStreamsOnlineListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetPageSize(v int32) *DescribeVsStreamsOnlineListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetQueryType(v string) *DescribeVsStreamsOnlineListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetStartTime(v string) *DescribeVsStreamsOnlineListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetStreamName(v string) *DescribeVsStreamsOnlineListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetStreamType(v string) *DescribeVsStreamsOnlineListRequest {
	s.StreamType = &v
	return s
}

type DescribeVsStreamsOnlineListResponseBody struct {
	OnlineInfo *DescribeVsStreamsOnlineListResponseBodyOnlineInfo `json:"OnlineInfo,omitempty" xml:"OnlineInfo,omitempty" type:"Struct"`
	PageNum    *int32                                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum   *int32                                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage  *int32                                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeVsStreamsOnlineListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetOnlineInfo(v *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) *DescribeVsStreamsOnlineListResponseBody {
	s.OnlineInfo = v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetPageNum(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetPageSize(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetRequestId(v string) *DescribeVsStreamsOnlineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetTotalNum(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetTotalPage(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeVsStreamsOnlineListResponseBodyOnlineInfo struct {
	LiveStreamOnlineInfo []*DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo `json:"LiveStreamOnlineInfo,omitempty" xml:"LiveStreamOnlineInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) SetLiveStreamOnlineInfo(v []*DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) *DescribeVsStreamsOnlineListResponseBodyOnlineInfo {
	s.LiveStreamOnlineInfo = v
	return s
}

type DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty"`
	PublishTime   *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	PublishType   *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	PublishUrl    *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	TranscodeId   *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty"`
	Transcoded    *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty"`
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetAppName(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetDomainName(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishDomain(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishTime(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishType(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishUrl(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetStreamName(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetTranscodeId(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.TranscodeId = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetTranscoded(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.Transcoded = &v
	return s
}

type DescribeVsStreamsOnlineListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsStreamsOnlineListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsStreamsOnlineListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponse) SetHeaders(v map[string]*string) *DescribeVsStreamsOnlineListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStreamsOnlineListResponse) SetBody(v *DescribeVsStreamsOnlineListResponseBody) *DescribeVsStreamsOnlineListResponse {
	s.Body = v
	return s
}

type DescribeVsStreamsPublishListRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DescribeVsStreamsPublishListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsPublishListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListRequest) SetAppName(v string) *DescribeVsStreamsPublishListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetDomainName(v string) *DescribeVsStreamsPublishListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetEndTime(v string) *DescribeVsStreamsPublishListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetOrderBy(v string) *DescribeVsStreamsPublishListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetOwnerId(v int64) *DescribeVsStreamsPublishListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetPageNumber(v int32) *DescribeVsStreamsPublishListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetPageSize(v int32) *DescribeVsStreamsPublishListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetQueryType(v string) *DescribeVsStreamsPublishListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetStartTime(v string) *DescribeVsStreamsPublishListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetStreamName(v string) *DescribeVsStreamsPublishListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetStreamType(v string) *DescribeVsStreamsPublishListRequest {
	s.StreamType = &v
	return s
}

type DescribeVsStreamsPublishListResponseBody struct {
	PageNum     *int32                                               `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublishInfo *DescribeVsStreamsPublishListResponseBodyPublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" type:"Struct"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum    *int32                                               `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage   *int32                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeVsStreamsPublishListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponseBody) SetPageNum(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetPageSize(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetPublishInfo(v *DescribeVsStreamsPublishListResponseBodyPublishInfo) *DescribeVsStreamsPublishListResponseBody {
	s.PublishInfo = v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetRequestId(v string) *DescribeVsStreamsPublishListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetTotalNum(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetTotalPage(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeVsStreamsPublishListResponseBodyPublishInfo struct {
	LiveStreamPublishInfo []*DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo `json:"LiveStreamPublishInfo,omitempty" xml:"LiveStreamPublishInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfo) SetLiveStreamPublishInfo(v []*DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) *DescribeVsStreamsPublishListResponseBodyPublishInfo {
	s.LiveStreamPublishInfo = v
	return s
}

type DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClientAddr    *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EdgeNodeAddr  *string `json:"EdgeNodeAddr,omitempty" xml:"EdgeNodeAddr,omitempty"`
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty"`
	PublishTime   *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	PublishType   *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	PublishUrl    *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
	StopTime      *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamUrl     *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	TranscodeId   *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty"`
	Transcoded    *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty"`
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetAppName(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetClientAddr(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.ClientAddr = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetDomainName(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetEdgeNodeAddr(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.EdgeNodeAddr = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishDomain(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishTime(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishType(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishUrl(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStopTime(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StopTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStreamName(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStreamUrl(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetTranscodeId(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.TranscodeId = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetTranscoded(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.Transcoded = &v
	return s
}

type DescribeVsStreamsPublishListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsStreamsPublishListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsStreamsPublishListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponse) SetHeaders(v map[string]*string) *DescribeVsStreamsPublishListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStreamsPublishListResponse) SetBody(v *DescribeVsStreamsPublishListResponseBody) *DescribeVsStreamsPublishListResponse {
	s.Body = v
	return s
}

type DescribeVsTopDomainsByFlowRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsTopDomainsByFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowRequest) SetEndTime(v string) *DescribeVsTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) SetLimit(v int64) *DescribeVsTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeVsTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) SetStartTime(v string) *DescribeVsTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

type DescribeVsTopDomainsByFlowResponseBody struct {
	DomainCount       *int64                                            `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainOnlineCount *int64                                            `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	EndTime           *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime         *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains        *DescribeVsTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeVsTopDomainsByFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeVsTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeVsTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeVsTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeVsTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeVsTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeVsTopDomainsByFlowResponseBodyTopDomains) *DescribeVsTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

type DescribeVsTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeVsTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

type DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MaxBps         *int64  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	Rank           *int64  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64  `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

type DescribeVsTopDomainsByFlowResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsTopDomainsByFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeVsTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponse) SetBody(v *DescribeVsTopDomainsByFlowResponseBody) *DescribeVsTopDomainsByFlowResponse {
	s.Body = v
	return s
}

type DescribeVsUpPeakPublishStreamDataRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetDomainName(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetDomainSwitch(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetEndTime(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetOwnerId(v int64) *DescribeVsUpPeakPublishStreamDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetStartTime(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVsUpPeakPublishStreamDataResponseBody struct {
	DescribeVsUpPeakPublishStreamDatas *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas `json:"DescribeVsUpPeakPublishStreamDatas,omitempty" xml:"DescribeVsUpPeakPublishStreamDatas,omitempty" type:"Struct"`
	RequestId                          *string                                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) SetDescribeVsUpPeakPublishStreamDatas(v *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) *DescribeVsUpPeakPublishStreamDataResponseBody {
	s.DescribeVsUpPeakPublishStreamDatas = v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBody) SetRequestId(v string) *DescribeVsUpPeakPublishStreamDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas struct {
	DescribeVsUpPeakPublishStreamData []*DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData `json:"DescribeVsUpPeakPublishStreamData,omitempty" xml:"DescribeVsUpPeakPublishStreamData,omitempty" type:"Repeated"`
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas) SetDescribeVsUpPeakPublishStreamData(v []*DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatas {
	s.DescribeVsUpPeakPublishStreamData = v
	return s
}

type DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData struct {
	BandWidth        *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	PeakTime         *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	PublishStreamNum *int32  `json:"PublishStreamNum,omitempty" xml:"PublishStreamNum,omitempty"`
	QueryTime        *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	StatName         *string `json:"StatName,omitempty" xml:"StatName,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetBandWidth(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.BandWidth = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetPeakTime(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.PeakTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetPublishStreamNum(v int32) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.PublishStreamNum = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetQueryTime(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.QueryTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData) SetStatName(v string) *DescribeVsUpPeakPublishStreamDataResponseBodyDescribeVsUpPeakPublishStreamDatasDescribeVsUpPeakPublishStreamData {
	s.StatName = &v
	return s
}

type DescribeVsUpPeakPublishStreamDataResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsUpPeakPublishStreamDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsUpPeakPublishStreamDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) SetHeaders(v map[string]*string) *DescribeVsUpPeakPublishStreamDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) SetBody(v *DescribeVsUpPeakPublishStreamDataResponseBody) *DescribeVsUpPeakPublishStreamDataResponse {
	s.Body = v
	return s
}

type DescribeVsUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVsUserResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageRequest) SetOwnerId(v int64) *DescribeVsUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsUserResourcePackageRequest) SetSecurityToken(v string) *DescribeVsUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

type DescribeVsUserResourcePackageResponseBody struct {
	RequestId            *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePackageInfos *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeVsUserResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponseBody) SetRequestId(v string) *DescribeVsUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) *DescribeVsUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

type DescribeVsUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

type DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CurrCapacity  *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InitCapacity  *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

type DescribeVsUserResourcePackageResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVsUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVsUserResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeVsUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsUserResourcePackageResponse) SetBody(v *DescribeVsUserResourcePackageResponseBody) *DescribeVsUserResourcePackageResponse {
	s.Body = v
	return s
}

type ForbidVsStreamRequest struct {
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LiveStreamType      *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	Oneshot             *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResumeTime          *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
	StreamName          *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ForbidVsStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ForbidVsStreamRequest) GoString() string {
	return s.String()
}

func (s *ForbidVsStreamRequest) SetAppName(v string) *ForbidVsStreamRequest {
	s.AppName = &v
	return s
}

func (s *ForbidVsStreamRequest) SetControlStreamAction(v string) *ForbidVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *ForbidVsStreamRequest) SetDomainName(v string) *ForbidVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ForbidVsStreamRequest) SetLiveStreamType(v string) *ForbidVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ForbidVsStreamRequest) SetOneshot(v string) *ForbidVsStreamRequest {
	s.Oneshot = &v
	return s
}

func (s *ForbidVsStreamRequest) SetOwnerId(v int64) *ForbidVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ForbidVsStreamRequest) SetResumeTime(v string) *ForbidVsStreamRequest {
	s.ResumeTime = &v
	return s
}

func (s *ForbidVsStreamRequest) SetStreamName(v string) *ForbidVsStreamRequest {
	s.StreamName = &v
	return s
}

type ForbidVsStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ForbidVsStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForbidVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ForbidVsStreamResponseBody) SetRequestId(v string) *ForbidVsStreamResponseBody {
	s.RequestId = &v
	return s
}

type ForbidVsStreamResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ForbidVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ForbidVsStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ForbidVsStreamResponse) GoString() string {
	return s.String()
}

func (s *ForbidVsStreamResponse) SetHeaders(v map[string]*string) *ForbidVsStreamResponse {
	s.Headers = v
	return s
}

func (s *ForbidVsStreamResponse) SetBody(v *ForbidVsStreamResponseBody) *ForbidVsStreamResponse {
	s.Body = v
	return s
}

type GetBucketInfoRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetBucketInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBucketInfoRequest) SetBucketName(v string) *GetBucketInfoRequest {
	s.BucketName = &v
	return s
}

func (s *GetBucketInfoRequest) SetOwnerId(v int64) *GetBucketInfoRequest {
	s.OwnerId = &v
	return s
}

type GetBucketInfoResponseBody struct {
	BucketInfo *GetBucketInfoResponseBodyBucketInfo `json:"BucketInfo,omitempty" xml:"BucketInfo,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBucketInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBody) SetBucketInfo(v *GetBucketInfoResponseBodyBucketInfo) *GetBucketInfoResponseBody {
	s.BucketInfo = v
	return s
}

func (s *GetBucketInfoResponseBody) SetRequestId(v string) *GetBucketInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetBucketInfoResponseBodyBucketInfo struct {
	BucketAcl          *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	BucketName         *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	DispatcherType     *string `json:"DispatcherType,omitempty" xml:"DispatcherType,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ModifyTime         *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StorageClass       *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s GetBucketInfoResponseBodyBucketInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponseBodyBucketInfo) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetBucketAcl(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.BucketAcl = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetBucketName(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.BucketName = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetComment(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.Comment = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetCreateTime(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.CreateTime = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetDataRedundancyType(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.DataRedundancyType = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetDispatcherType(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.DispatcherType = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetEndpoint(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.Endpoint = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetModifyTime(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetResourceType(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.ResourceType = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetStorageClass(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.StorageClass = &v
	return s
}

type GetBucketInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBucketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponse) SetHeaders(v map[string]*string) *GetBucketInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBucketInfoResponse) SetBody(v *GetBucketInfoResponseBody) *GetBucketInfoResponse {
	s.Body = v
	return s
}

type GetObjectTotalRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetObjectTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetObjectTotalRequest) GoString() string {
	return s.String()
}

func (s *GetObjectTotalRequest) SetBucketName(v string) *GetObjectTotalRequest {
	s.BucketName = &v
	return s
}

func (s *GetObjectTotalRequest) SetOwnerId(v int64) *GetObjectTotalRequest {
	s.OwnerId = &v
	return s
}

type GetObjectTotalResponseBody struct {
	ObjectCount *int64  `json:"ObjectCount,omitempty" xml:"ObjectCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetObjectTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetObjectTotalResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectTotalResponseBody) SetObjectCount(v int64) *GetObjectTotalResponseBody {
	s.ObjectCount = &v
	return s
}

func (s *GetObjectTotalResponseBody) SetRequestId(v string) *GetObjectTotalResponseBody {
	s.RequestId = &v
	return s
}

type GetObjectTotalResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetObjectTotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetObjectTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetObjectTotalResponse) GoString() string {
	return s.String()
}

func (s *GetObjectTotalResponse) SetHeaders(v map[string]*string) *GetObjectTotalResponse {
	s.Headers = v
	return s
}

func (s *GetObjectTotalResponse) SetBody(v *GetObjectTotalResponseBody) *GetObjectTotalResponse {
	s.Body = v
	return s
}

type GotoPresetRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
}

func (s GotoPresetRequest) String() string {
	return tea.Prettify(s)
}

func (s GotoPresetRequest) GoString() string {
	return s.String()
}

func (s *GotoPresetRequest) SetId(v string) *GotoPresetRequest {
	s.Id = &v
	return s
}

func (s *GotoPresetRequest) SetOwnerId(v int64) *GotoPresetRequest {
	s.OwnerId = &v
	return s
}

func (s *GotoPresetRequest) SetPresetId(v string) *GotoPresetRequest {
	s.PresetId = &v
	return s
}

type GotoPresetResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GotoPresetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GotoPresetResponseBody) GoString() string {
	return s.String()
}

func (s *GotoPresetResponseBody) SetId(v string) *GotoPresetResponseBody {
	s.Id = &v
	return s
}

func (s *GotoPresetResponseBody) SetRequestId(v string) *GotoPresetResponseBody {
	s.RequestId = &v
	return s
}

type GotoPresetResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GotoPresetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GotoPresetResponse) String() string {
	return tea.Prettify(s)
}

func (s GotoPresetResponse) GoString() string {
	return s.String()
}

func (s *GotoPresetResponse) SetHeaders(v map[string]*string) *GotoPresetResponse {
	s.Headers = v
	return s
}

func (s *GotoPresetResponse) SetBody(v *GotoPresetResponseBody) *GotoPresetResponse {
	s.Body = v
	return s
}

type ListBucketsRequest struct {
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Marker     *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Prefix     *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListBucketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListBucketsRequest) SetKeyword(v string) *ListBucketsRequest {
	s.Keyword = &v
	return s
}

func (s *ListBucketsRequest) SetMarker(v string) *ListBucketsRequest {
	s.Marker = &v
	return s
}

func (s *ListBucketsRequest) SetOwnerId(v int64) *ListBucketsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListBucketsRequest) SetPageNumber(v int32) *ListBucketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBucketsRequest) SetPageSize(v int32) *ListBucketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBucketsRequest) SetPrefix(v string) *ListBucketsRequest {
	s.Prefix = &v
	return s
}

type ListBucketsResponseBody struct {
	BucketInfos []*ListBucketsResponseBodyBucketInfos `json:"BucketInfos,omitempty" xml:"BucketInfos,omitempty" type:"Repeated"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBucketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBody) SetBucketInfos(v []*ListBucketsResponseBodyBucketInfos) *ListBucketsResponseBody {
	s.BucketInfos = v
	return s
}

func (s *ListBucketsResponseBody) SetRequestId(v string) *ListBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBucketsResponseBody) SetTotalCount(v int64) *ListBucketsResponseBody {
	s.TotalCount = &v
	return s
}

type ListBucketsResponseBodyBucketInfos struct {
	BucketAcl          *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	BucketName         *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	DispatcherType     *string `json:"DispatcherType,omitempty" xml:"DispatcherType,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ModifyTime         *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StorageClass       *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s ListBucketsResponseBodyBucketInfos) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsResponseBodyBucketInfos) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBodyBucketInfos) SetBucketAcl(v string) *ListBucketsResponseBodyBucketInfos {
	s.BucketAcl = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetBucketName(v string) *ListBucketsResponseBodyBucketInfos {
	s.BucketName = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetComment(v string) *ListBucketsResponseBodyBucketInfos {
	s.Comment = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetCreateTime(v string) *ListBucketsResponseBodyBucketInfos {
	s.CreateTime = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetDataRedundancyType(v string) *ListBucketsResponseBodyBucketInfos {
	s.DataRedundancyType = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetDispatcherType(v string) *ListBucketsResponseBodyBucketInfos {
	s.DispatcherType = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetEndpoint(v string) *ListBucketsResponseBodyBucketInfos {
	s.Endpoint = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetModifyTime(v string) *ListBucketsResponseBodyBucketInfos {
	s.ModifyTime = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetResourceType(v string) *ListBucketsResponseBodyBucketInfos {
	s.ResourceType = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetStorageClass(v string) *ListBucketsResponseBodyBucketInfos {
	s.StorageClass = &v
	return s
}

type ListBucketsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBucketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListBucketsResponse) SetHeaders(v map[string]*string) *ListBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListBucketsResponse) SetBody(v *ListBucketsResponseBody) *ListBucketsResponse {
	s.Body = v
	return s
}

type ListObjectsRequest struct {
	BucketName        *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	Delimiter         *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	EncodingType      *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	Marker            *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys           *int32  `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Prefix            *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	StartAfter        *string `json:"StartAfter,omitempty" xml:"StartAfter,omitempty"`
}

func (s ListObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsRequest) GoString() string {
	return s.String()
}

func (s *ListObjectsRequest) SetBucketName(v string) *ListObjectsRequest {
	s.BucketName = &v
	return s
}

func (s *ListObjectsRequest) SetContinuationToken(v string) *ListObjectsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsRequest) SetDelimiter(v string) *ListObjectsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsRequest) SetEncodingType(v string) *ListObjectsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsRequest) SetMarker(v string) *ListObjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListObjectsRequest) SetMaxKeys(v int32) *ListObjectsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsRequest) SetOwnerId(v int64) *ListObjectsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListObjectsRequest) SetPrefix(v string) *ListObjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListObjectsRequest) SetStartAfter(v string) *ListObjectsRequest {
	s.StartAfter = &v
	return s
}

type ListObjectsResponseBody struct {
	BucketName            *string                            `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CommonPrefixes        []*string                          `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	Contents              []*ListObjectsResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	ContinuationToken     *string                            `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	Delimiter             *string                            `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	EncodingType          *string                            `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	IsTruncated           *bool                              `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	KeyCount              *int32                             `json:"KeyCount,omitempty" xml:"KeyCount,omitempty"`
	Marker                *string                            `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys               *int32                             `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	NextContinuationToken *string                            `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	NextMarker            *string                            `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Prefix                *string                            `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	RequestId             *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBody) SetBucketName(v string) *ListObjectsResponseBody {
	s.BucketName = &v
	return s
}

func (s *ListObjectsResponseBody) SetCommonPrefixes(v []*string) *ListObjectsResponseBody {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectsResponseBody) SetContents(v []*ListObjectsResponseBodyContents) *ListObjectsResponseBody {
	s.Contents = v
	return s
}

func (s *ListObjectsResponseBody) SetContinuationToken(v string) *ListObjectsResponseBody {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsResponseBody) SetDelimiter(v string) *ListObjectsResponseBody {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsResponseBody) SetEncodingType(v string) *ListObjectsResponseBody {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsResponseBody) SetIsTruncated(v bool) *ListObjectsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectsResponseBody) SetKeyCount(v int32) *ListObjectsResponseBody {
	s.KeyCount = &v
	return s
}

func (s *ListObjectsResponseBody) SetMarker(v string) *ListObjectsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListObjectsResponseBody) SetMaxKeys(v int32) *ListObjectsResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsResponseBody) SetNextContinuationToken(v string) *ListObjectsResponseBody {
	s.NextContinuationToken = &v
	return s
}

func (s *ListObjectsResponseBody) SetNextMarker(v string) *ListObjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListObjectsResponseBody) SetPrefix(v string) *ListObjectsResponseBody {
	s.Prefix = &v
	return s
}

func (s *ListObjectsResponseBody) SetRequestId(v string) *ListObjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListObjectsResponseBodyContents struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s ListObjectsResponseBodyContents) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsResponseBodyContents) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBodyContents) SetETag(v string) *ListObjectsResponseBodyContents {
	s.ETag = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetKey(v string) *ListObjectsResponseBodyContents {
	s.Key = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetLastModified(v string) *ListObjectsResponseBodyContents {
	s.LastModified = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetSize(v int64) *ListObjectsResponseBodyContents {
	s.Size = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetStorageClass(v string) *ListObjectsResponseBodyContents {
	s.StorageClass = &v
	return s
}

type ListObjectsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsResponse) GoString() string {
	return s.String()
}

func (s *ListObjectsResponse) SetHeaders(v map[string]*string) *ListObjectsResponse {
	s.Headers = v
	return s
}

func (s *ListObjectsResponse) SetBody(v *ListObjectsResponseBody) *ListObjectsResponse {
	s.Body = v
	return s
}

type ModifyDeviceRequest struct {
	AlarmMethod *string `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoPos     *bool   `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	AutoStart   *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GbId        *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Latitude    *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude   *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port        *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	PosInterval *int64  `json:"PosInterval,omitempty" xml:"PosInterval,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Username    *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Vendor      *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ModifyDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceRequest) SetAlarmMethod(v string) *ModifyDeviceRequest {
	s.AlarmMethod = &v
	return s
}

func (s *ModifyDeviceRequest) SetAutoPos(v bool) *ModifyDeviceRequest {
	s.AutoPos = &v
	return s
}

func (s *ModifyDeviceRequest) SetAutoStart(v bool) *ModifyDeviceRequest {
	s.AutoStart = &v
	return s
}

func (s *ModifyDeviceRequest) SetDescription(v string) *ModifyDeviceRequest {
	s.Description = &v
	return s
}

func (s *ModifyDeviceRequest) SetDirectoryId(v string) *ModifyDeviceRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyDeviceRequest) SetGbId(v string) *ModifyDeviceRequest {
	s.GbId = &v
	return s
}

func (s *ModifyDeviceRequest) SetGroupId(v string) *ModifyDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyDeviceRequest) SetId(v string) *ModifyDeviceRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceRequest) SetIp(v string) *ModifyDeviceRequest {
	s.Ip = &v
	return s
}

func (s *ModifyDeviceRequest) SetLatitude(v string) *ModifyDeviceRequest {
	s.Latitude = &v
	return s
}

func (s *ModifyDeviceRequest) SetLongitude(v string) *ModifyDeviceRequest {
	s.Longitude = &v
	return s
}

func (s *ModifyDeviceRequest) SetName(v string) *ModifyDeviceRequest {
	s.Name = &v
	return s
}

func (s *ModifyDeviceRequest) SetOwnerId(v int64) *ModifyDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceRequest) SetParams(v string) *ModifyDeviceRequest {
	s.Params = &v
	return s
}

func (s *ModifyDeviceRequest) SetParentId(v string) *ModifyDeviceRequest {
	s.ParentId = &v
	return s
}

func (s *ModifyDeviceRequest) SetPassword(v string) *ModifyDeviceRequest {
	s.Password = &v
	return s
}

func (s *ModifyDeviceRequest) SetPort(v int64) *ModifyDeviceRequest {
	s.Port = &v
	return s
}

func (s *ModifyDeviceRequest) SetPosInterval(v int64) *ModifyDeviceRequest {
	s.PosInterval = &v
	return s
}

func (s *ModifyDeviceRequest) SetType(v string) *ModifyDeviceRequest {
	s.Type = &v
	return s
}

func (s *ModifyDeviceRequest) SetUrl(v string) *ModifyDeviceRequest {
	s.Url = &v
	return s
}

func (s *ModifyDeviceRequest) SetUsername(v string) *ModifyDeviceRequest {
	s.Username = &v
	return s
}

func (s *ModifyDeviceRequest) SetVendor(v string) *ModifyDeviceRequest {
	s.Vendor = &v
	return s
}

type ModifyDeviceResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceResponseBody) SetId(v string) *ModifyDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyDeviceResponseBody) SetRequestId(v string) *ModifyDeviceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceResponse) SetHeaders(v map[string]*string) *ModifyDeviceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceResponse) SetBody(v *ModifyDeviceResponseBody) *ModifyDeviceResponse {
	s.Body = v
	return s
}

type ModifyDeviceAlarmRequest struct {
	AlarmId   *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	ChannelId *int32  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDeviceAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceAlarmRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAlarmRequest) SetAlarmId(v string) *ModifyDeviceAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetChannelId(v int32) *ModifyDeviceAlarmRequest {
	s.ChannelId = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetId(v string) *ModifyDeviceAlarmRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetOwnerId(v int64) *ModifyDeviceAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetStatus(v int32) *ModifyDeviceAlarmRequest {
	s.Status = &v
	return s
}

type ModifyDeviceAlarmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAlarmResponseBody) SetRequestId(v string) *ModifyDeviceAlarmResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDeviceAlarmResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDeviceAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeviceAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceAlarmResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAlarmResponse) SetHeaders(v map[string]*string) *ModifyDeviceAlarmResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceAlarmResponse) SetBody(v *ModifyDeviceAlarmResponseBody) *ModifyDeviceAlarmResponse {
	s.Body = v
	return s
}

type ModifyDeviceCaptureRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Image   *int32  `json:"Image,omitempty" xml:"Image,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Video   *int32  `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s ModifyDeviceCaptureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceCaptureRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceCaptureRequest) SetId(v string) *ModifyDeviceCaptureRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) SetImage(v int32) *ModifyDeviceCaptureRequest {
	s.Image = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) SetOwnerId(v int64) *ModifyDeviceCaptureRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) SetVideo(v int32) *ModifyDeviceCaptureRequest {
	s.Video = &v
	return s
}

type ModifyDeviceCaptureResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceCaptureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceCaptureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceCaptureResponseBody) SetRequestId(v string) *ModifyDeviceCaptureResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDeviceCaptureResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDeviceCaptureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeviceCaptureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceCaptureResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceCaptureResponse) SetHeaders(v map[string]*string) *ModifyDeviceCaptureResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceCaptureResponse) SetBody(v *ModifyDeviceCaptureResponseBody) *ModifyDeviceCaptureResponse {
	s.Body = v
	return s
}

type ModifyDeviceChannelsRequest struct {
	Channels     *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	Dsn          *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyDeviceChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceChannelsRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceChannelsRequest) SetChannels(v string) *ModifyDeviceChannelsRequest {
	s.Channels = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetDeviceStatus(v string) *ModifyDeviceChannelsRequest {
	s.DeviceStatus = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetDsn(v string) *ModifyDeviceChannelsRequest {
	s.Dsn = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetId(v string) *ModifyDeviceChannelsRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetOwnerId(v int64) *ModifyDeviceChannelsRequest {
	s.OwnerId = &v
	return s
}

type ModifyDeviceChannelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceChannelsResponseBody) SetRequestId(v string) *ModifyDeviceChannelsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDeviceChannelsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDeviceChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeviceChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceChannelsResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceChannelsResponse) SetHeaders(v map[string]*string) *ModifyDeviceChannelsResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceChannelsResponse) SetBody(v *ModifyDeviceChannelsResponseBody) *ModifyDeviceChannelsResponse {
	s.Body = v
	return s
}

type ModifyDirectoryRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyDirectoryRequest) SetDescription(v string) *ModifyDirectoryRequest {
	s.Description = &v
	return s
}

func (s *ModifyDirectoryRequest) SetId(v string) *ModifyDirectoryRequest {
	s.Id = &v
	return s
}

func (s *ModifyDirectoryRequest) SetName(v string) *ModifyDirectoryRequest {
	s.Name = &v
	return s
}

func (s *ModifyDirectoryRequest) SetOwnerId(v int64) *ModifyDirectoryRequest {
	s.OwnerId = &v
	return s
}

type ModifyDirectoryResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDirectoryResponseBody) SetId(v string) *ModifyDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyDirectoryResponseBody) SetRequestId(v string) *ModifyDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDirectoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyDirectoryResponse) SetHeaders(v map[string]*string) *ModifyDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyDirectoryResponse) SetBody(v *ModifyDirectoryResponseBody) *ModifyDirectoryResponse {
	s.Body = v
	return s
}

type ModifyGroupRequest struct {
	Callback    *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InProtocol  *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	LazyPull    *bool   `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayDomain  *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	PushDomain  *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupRequest) SetCallback(v string) *ModifyGroupRequest {
	s.Callback = &v
	return s
}

func (s *ModifyGroupRequest) SetDescription(v string) *ModifyGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyGroupRequest) SetEnabled(v bool) *ModifyGroupRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyGroupRequest) SetId(v string) *ModifyGroupRequest {
	s.Id = &v
	return s
}

func (s *ModifyGroupRequest) SetInProtocol(v string) *ModifyGroupRequest {
	s.InProtocol = &v
	return s
}

func (s *ModifyGroupRequest) SetLazyPull(v bool) *ModifyGroupRequest {
	s.LazyPull = &v
	return s
}

func (s *ModifyGroupRequest) SetName(v string) *ModifyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyGroupRequest) SetOutProtocol(v string) *ModifyGroupRequest {
	s.OutProtocol = &v
	return s
}

func (s *ModifyGroupRequest) SetOwnerId(v int64) *ModifyGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGroupRequest) SetPlayDomain(v string) *ModifyGroupRequest {
	s.PlayDomain = &v
	return s
}

func (s *ModifyGroupRequest) SetPushDomain(v string) *ModifyGroupRequest {
	s.PushDomain = &v
	return s
}

func (s *ModifyGroupRequest) SetRegion(v string) *ModifyGroupRequest {
	s.Region = &v
	return s
}

type ModifyGroupResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponseBody) SetId(v string) *ModifyGroupResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyGroupResponseBody) SetRequestId(v string) *ModifyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponse) SetHeaders(v map[string]*string) *ModifyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupResponse) SetBody(v *ModifyGroupResponseBody) *ModifyGroupResponse {
	s.Body = v
	return s
}

type ModifyParentPlatformRequest struct {
	AutoStart      *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ClientAuth     *bool   `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port           *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyParentPlatformRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *ModifyParentPlatformRequest) SetAutoStart(v bool) *ModifyParentPlatformRequest {
	s.AutoStart = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetClientAuth(v bool) *ModifyParentPlatformRequest {
	s.ClientAuth = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetClientPassword(v string) *ModifyParentPlatformRequest {
	s.ClientPassword = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetClientUsername(v string) *ModifyParentPlatformRequest {
	s.ClientUsername = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetDescription(v string) *ModifyParentPlatformRequest {
	s.Description = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetGbId(v string) *ModifyParentPlatformRequest {
	s.GbId = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetId(v string) *ModifyParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetIp(v string) *ModifyParentPlatformRequest {
	s.Ip = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetName(v string) *ModifyParentPlatformRequest {
	s.Name = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetOwnerId(v int64) *ModifyParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetPort(v int64) *ModifyParentPlatformRequest {
	s.Port = &v
	return s
}

type ModifyParentPlatformResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParentPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParentPlatformResponseBody) SetId(v string) *ModifyParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyParentPlatformResponseBody) SetRequestId(v string) *ModifyParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

type ModifyParentPlatformResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParentPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *ModifyParentPlatformResponse) SetHeaders(v map[string]*string) *ModifyParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *ModifyParentPlatformResponse) SetBody(v *ModifyParentPlatformResponseBody) *ModifyParentPlatformResponse {
	s.Body = v
	return s
}

type ModifyPurchasedDeviceRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	SubType      *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ModifyPurchasedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *ModifyPurchasedDeviceRequest) SetDescription(v string) *ModifyPurchasedDeviceRequest {
	s.Description = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetId(v string) *ModifyPurchasedDeviceRequest {
	s.Id = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetName(v string) *ModifyPurchasedDeviceRequest {
	s.Name = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetOrderId(v string) *ModifyPurchasedDeviceRequest {
	s.OrderId = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetOwnerId(v int64) *ModifyPurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetRegisterCode(v string) *ModifyPurchasedDeviceRequest {
	s.RegisterCode = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetSubType(v string) *ModifyPurchasedDeviceRequest {
	s.SubType = &v
	return s
}

func (s *ModifyPurchasedDeviceRequest) SetVendor(v string) *ModifyPurchasedDeviceRequest {
	s.Vendor = &v
	return s
}

type ModifyPurchasedDeviceResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPurchasedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPurchasedDeviceResponseBody) SetId(v string) *ModifyPurchasedDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyPurchasedDeviceResponseBody) SetRequestId(v string) *ModifyPurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPurchasedDeviceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPurchasedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *ModifyPurchasedDeviceResponse) SetHeaders(v map[string]*string) *ModifyPurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *ModifyPurchasedDeviceResponse) SetBody(v *ModifyPurchasedDeviceResponseBody) *ModifyPurchasedDeviceResponse {
	s.Body = v
	return s
}

type ModifyTemplateRequest struct {
	Callback         *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileFormat       *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	Flv              *string `json:"Flv,omitempty" xml:"Flv,omitempty"`
	HlsM3u8          *string `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	HlsTs            *string `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval         *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JpgOnDemand      *string `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	JpgOverwrite     *string `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	JpgSequence      *string `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	Mp4              *string `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket        *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint      *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssFilePrefix    *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Retention        *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	TransConfigsJSON *string `json:"TransConfigsJSON,omitempty" xml:"TransConfigsJSON,omitempty"`
	Trigger          *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s ModifyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequest) SetCallback(v string) *ModifyTemplateRequest {
	s.Callback = &v
	return s
}

func (s *ModifyTemplateRequest) SetDescription(v string) *ModifyTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyTemplateRequest) SetFileFormat(v string) *ModifyTemplateRequest {
	s.FileFormat = &v
	return s
}

func (s *ModifyTemplateRequest) SetFlv(v string) *ModifyTemplateRequest {
	s.Flv = &v
	return s
}

func (s *ModifyTemplateRequest) SetHlsM3u8(v string) *ModifyTemplateRequest {
	s.HlsM3u8 = &v
	return s
}

func (s *ModifyTemplateRequest) SetHlsTs(v string) *ModifyTemplateRequest {
	s.HlsTs = &v
	return s
}

func (s *ModifyTemplateRequest) SetId(v string) *ModifyTemplateRequest {
	s.Id = &v
	return s
}

func (s *ModifyTemplateRequest) SetInterval(v int64) *ModifyTemplateRequest {
	s.Interval = &v
	return s
}

func (s *ModifyTemplateRequest) SetJpgOnDemand(v string) *ModifyTemplateRequest {
	s.JpgOnDemand = &v
	return s
}

func (s *ModifyTemplateRequest) SetJpgOverwrite(v string) *ModifyTemplateRequest {
	s.JpgOverwrite = &v
	return s
}

func (s *ModifyTemplateRequest) SetJpgSequence(v string) *ModifyTemplateRequest {
	s.JpgSequence = &v
	return s
}

func (s *ModifyTemplateRequest) SetMp4(v string) *ModifyTemplateRequest {
	s.Mp4 = &v
	return s
}

func (s *ModifyTemplateRequest) SetName(v string) *ModifyTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyTemplateRequest) SetOssBucket(v string) *ModifyTemplateRequest {
	s.OssBucket = &v
	return s
}

func (s *ModifyTemplateRequest) SetOssEndpoint(v string) *ModifyTemplateRequest {
	s.OssEndpoint = &v
	return s
}

func (s *ModifyTemplateRequest) SetOssFilePrefix(v string) *ModifyTemplateRequest {
	s.OssFilePrefix = &v
	return s
}

func (s *ModifyTemplateRequest) SetOwnerId(v int64) *ModifyTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTemplateRequest) SetRegion(v string) *ModifyTemplateRequest {
	s.Region = &v
	return s
}

func (s *ModifyTemplateRequest) SetRetention(v int64) *ModifyTemplateRequest {
	s.Retention = &v
	return s
}

func (s *ModifyTemplateRequest) SetTransConfigsJSON(v string) *ModifyTemplateRequest {
	s.TransConfigsJSON = &v
	return s
}

func (s *ModifyTemplateRequest) SetTrigger(v string) *ModifyTemplateRequest {
	s.Trigger = &v
	return s
}

type ModifyTemplateResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponseBody) SetId(v string) *ModifyTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyTemplateResponseBody) SetRequestId(v string) *ModifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponse) SetHeaders(v map[string]*string) *ModifyTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateResponse) SetBody(v *ModifyTemplateResponseBody) *ModifyTemplateResponse {
	s.Body = v
	return s
}

type OpenVsServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenVsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVsServiceResponseBody) SetOrderId(v string) *OpenVsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenVsServiceResponseBody) SetRequestId(v string) *OpenVsServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenVsServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenVsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenVsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenVsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenVsServiceResponse) SetHeaders(v map[string]*string) *OpenVsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenVsServiceResponse) SetBody(v *OpenVsServiceResponseBody) *OpenVsServiceResponse {
	s.Body = v
	return s
}

type OperateRenderingDevicesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PodId       *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s OperateRenderingDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateRenderingDevicesRequest) GoString() string {
	return s.String()
}

func (s *OperateRenderingDevicesRequest) SetInstanceIds(v string) *OperateRenderingDevicesRequest {
	s.InstanceIds = &v
	return s
}

func (s *OperateRenderingDevicesRequest) SetOperation(v string) *OperateRenderingDevicesRequest {
	s.Operation = &v
	return s
}

func (s *OperateRenderingDevicesRequest) SetOwnerId(v int64) *OperateRenderingDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateRenderingDevicesRequest) SetPodId(v string) *OperateRenderingDevicesRequest {
	s.PodId = &v
	return s
}

type OperateRenderingDevicesResponseBody struct {
	FailedIds []*string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateRenderingDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateRenderingDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *OperateRenderingDevicesResponseBody) SetFailedIds(v []*string) *OperateRenderingDevicesResponseBody {
	s.FailedIds = v
	return s
}

func (s *OperateRenderingDevicesResponseBody) SetRequestId(v string) *OperateRenderingDevicesResponseBody {
	s.RequestId = &v
	return s
}

type OperateRenderingDevicesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateRenderingDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateRenderingDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateRenderingDevicesResponse) GoString() string {
	return s.String()
}

func (s *OperateRenderingDevicesResponse) SetHeaders(v map[string]*string) *OperateRenderingDevicesResponse {
	s.Headers = v
	return s
}

func (s *OperateRenderingDevicesResponse) SetBody(v *OperateRenderingDevicesResponseBody) *OperateRenderingDevicesResponse {
	s.Body = v
	return s
}

type PrepareUploadRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	ClientIp   *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PrepareUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s PrepareUploadRequest) GoString() string {
	return s.String()
}

func (s *PrepareUploadRequest) SetBucketName(v string) *PrepareUploadRequest {
	s.BucketName = &v
	return s
}

func (s *PrepareUploadRequest) SetClientIp(v string) *PrepareUploadRequest {
	s.ClientIp = &v
	return s
}

func (s *PrepareUploadRequest) SetOwnerId(v int64) *PrepareUploadRequest {
	s.OwnerId = &v
	return s
}

type PrepareUploadResponseBody struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Endpoint   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PrepareUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PrepareUploadResponseBody) GoString() string {
	return s.String()
}

func (s *PrepareUploadResponseBody) SetBucketName(v string) *PrepareUploadResponseBody {
	s.BucketName = &v
	return s
}

func (s *PrepareUploadResponseBody) SetEndpoint(v string) *PrepareUploadResponseBody {
	s.Endpoint = &v
	return s
}

func (s *PrepareUploadResponseBody) SetRequestId(v string) *PrepareUploadResponseBody {
	s.RequestId = &v
	return s
}

type PrepareUploadResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PrepareUploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PrepareUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s PrepareUploadResponse) GoString() string {
	return s.String()
}

func (s *PrepareUploadResponse) SetHeaders(v map[string]*string) *PrepareUploadResponse {
	s.Headers = v
	return s
}

func (s *PrepareUploadResponse) SetBody(v *PrepareUploadResponseBody) *PrepareUploadResponse {
	s.Body = v
	return s
}

type PutBucketRequest struct {
	BucketAcl          *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	BucketName         *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	DispatcherType     *string `json:"DispatcherType,omitempty" xml:"DispatcherType,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StorageClass       *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s PutBucketRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequest) SetBucketAcl(v string) *PutBucketRequest {
	s.BucketAcl = &v
	return s
}

func (s *PutBucketRequest) SetBucketName(v string) *PutBucketRequest {
	s.BucketName = &v
	return s
}

func (s *PutBucketRequest) SetComment(v string) *PutBucketRequest {
	s.Comment = &v
	return s
}

func (s *PutBucketRequest) SetDataRedundancyType(v string) *PutBucketRequest {
	s.DataRedundancyType = &v
	return s
}

func (s *PutBucketRequest) SetDispatcherType(v string) *PutBucketRequest {
	s.DispatcherType = &v
	return s
}

func (s *PutBucketRequest) SetEndpoint(v string) *PutBucketRequest {
	s.Endpoint = &v
	return s
}

func (s *PutBucketRequest) SetOwnerId(v int64) *PutBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *PutBucketRequest) SetResourceType(v string) *PutBucketRequest {
	s.ResourceType = &v
	return s
}

func (s *PutBucketRequest) SetStorageClass(v string) *PutBucketRequest {
	s.StorageClass = &v
	return s
}

type PutBucketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutBucketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutBucketResponseBody) GoString() string {
	return s.String()
}

func (s *PutBucketResponseBody) SetRequestId(v string) *PutBucketResponseBody {
	s.RequestId = &v
	return s
}

type PutBucketResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutBucketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketResponse) GoString() string {
	return s.String()
}

func (s *PutBucketResponse) SetHeaders(v map[string]*string) *PutBucketResponse {
	s.Headers = v
	return s
}

func (s *PutBucketResponse) SetBody(v *PutBucketResponseBody) *PutBucketResponse {
	s.Body = v
	return s
}

type ResetRenderingDevicesRequest struct {
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PodId       *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s ResetRenderingDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetRenderingDevicesRequest) GoString() string {
	return s.String()
}

func (s *ResetRenderingDevicesRequest) SetImageId(v string) *ResetRenderingDevicesRequest {
	s.ImageId = &v
	return s
}

func (s *ResetRenderingDevicesRequest) SetInstanceIds(v string) *ResetRenderingDevicesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ResetRenderingDevicesRequest) SetOwnerId(v int64) *ResetRenderingDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetRenderingDevicesRequest) SetPodId(v string) *ResetRenderingDevicesRequest {
	s.PodId = &v
	return s
}

type ResetRenderingDevicesResponseBody struct {
	FailedIds []*string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetRenderingDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetRenderingDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetRenderingDevicesResponseBody) SetFailedIds(v []*string) *ResetRenderingDevicesResponseBody {
	s.FailedIds = v
	return s
}

func (s *ResetRenderingDevicesResponseBody) SetRequestId(v string) *ResetRenderingDevicesResponseBody {
	s.RequestId = &v
	return s
}

type ResetRenderingDevicesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetRenderingDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetRenderingDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetRenderingDevicesResponse) GoString() string {
	return s.String()
}

func (s *ResetRenderingDevicesResponse) SetHeaders(v map[string]*string) *ResetRenderingDevicesResponse {
	s.Headers = v
	return s
}

func (s *ResetRenderingDevicesResponse) SetBody(v *ResetRenderingDevicesResponseBody) *ResetRenderingDevicesResponse {
	s.Body = v
	return s
}

type ResumeVsStreamRequest struct {
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LiveStreamType      *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StreamName          *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ResumeVsStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeVsStreamRequest) GoString() string {
	return s.String()
}

func (s *ResumeVsStreamRequest) SetAppName(v string) *ResumeVsStreamRequest {
	s.AppName = &v
	return s
}

func (s *ResumeVsStreamRequest) SetControlStreamAction(v string) *ResumeVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *ResumeVsStreamRequest) SetDomainName(v string) *ResumeVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ResumeVsStreamRequest) SetLiveStreamType(v string) *ResumeVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ResumeVsStreamRequest) SetOwnerId(v int64) *ResumeVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeVsStreamRequest) SetStreamName(v string) *ResumeVsStreamRequest {
	s.StreamName = &v
	return s
}

type ResumeVsStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeVsStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeVsStreamResponseBody) SetRequestId(v string) *ResumeVsStreamResponseBody {
	s.RequestId = &v
	return s
}

type ResumeVsStreamResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeVsStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeVsStreamResponse) GoString() string {
	return s.String()
}

func (s *ResumeVsStreamResponse) SetHeaders(v map[string]*string) *ResumeVsStreamResponse {
	s.Headers = v
	return s
}

func (s *ResumeVsStreamResponse) SetBody(v *ResumeVsStreamResponseBody) *ResumeVsStreamResponse {
	s.Body = v
	return s
}

type SetPresetRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
}

func (s SetPresetRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPresetRequest) GoString() string {
	return s.String()
}

func (s *SetPresetRequest) SetId(v string) *SetPresetRequest {
	s.Id = &v
	return s
}

func (s *SetPresetRequest) SetOwnerId(v int64) *SetPresetRequest {
	s.OwnerId = &v
	return s
}

func (s *SetPresetRequest) SetPresetId(v string) *SetPresetRequest {
	s.PresetId = &v
	return s
}

type SetPresetResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPresetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPresetResponseBody) GoString() string {
	return s.String()
}

func (s *SetPresetResponseBody) SetId(v string) *SetPresetResponseBody {
	s.Id = &v
	return s
}

func (s *SetPresetResponseBody) SetRequestId(v string) *SetPresetResponseBody {
	s.RequestId = &v
	return s
}

type SetPresetResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetPresetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetPresetResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPresetResponse) GoString() string {
	return s.String()
}

func (s *SetPresetResponse) SetHeaders(v map[string]*string) *SetPresetResponse {
	s.Headers = v
	return s
}

func (s *SetPresetResponse) SetBody(v *SetPresetResponseBody) *SetPresetResponse {
	s.Body = v
	return s
}

type SetVsDomainCertificateRequest struct {
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType    *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ForceSet    *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SSLPri      *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub      *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
}

func (s SetVsDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetVsDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetVsDomainCertificateRequest) SetCertName(v string) *SetVsDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetCertType(v string) *SetVsDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetDomainName(v string) *SetVsDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetForceSet(v string) *SetVsDomainCertificateRequest {
	s.ForceSet = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetOwnerId(v int64) *SetVsDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetRegion(v string) *SetVsDomainCertificateRequest {
	s.Region = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetSSLPri(v string) *SetVsDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetSSLProtocol(v string) *SetVsDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetSSLPub(v string) *SetVsDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

type SetVsDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVsDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetVsDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetVsDomainCertificateResponseBody) SetRequestId(v string) *SetVsDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetVsDomainCertificateResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetVsDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetVsDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetVsDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetVsDomainCertificateResponse) SetHeaders(v map[string]*string) *SetVsDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetVsDomainCertificateResponse) SetBody(v *SetVsDomainCertificateResponseBody) *SetVsDomainCertificateResponse {
	s.Body = v
	return s
}

type SetVsStreamsNotifyUrlConfigRequest struct {
	AuthKey    *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthType   *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	NotifyUrl  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetVsStreamsNotifyUrlConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetVsStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetAuthKey(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.AuthKey = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetAuthType(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.AuthType = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetDomainName(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetNotifyUrl(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *SetVsStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

type SetVsStreamsNotifyUrlConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVsStreamsNotifyUrlConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetVsStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetVsStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *SetVsStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetVsStreamsNotifyUrlConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetVsStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetVsStreamsNotifyUrlConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetVsStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *SetVsStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *SetVsStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigResponse) SetBody(v *SetVsStreamsNotifyUrlConfigResponseBody) *SetVsStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

type StartDeviceRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StartDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDeviceRequest) GoString() string {
	return s.String()
}

func (s *StartDeviceRequest) SetId(v string) *StartDeviceRequest {
	s.Id = &v
	return s
}

func (s *StartDeviceRequest) SetOwnerId(v int64) *StartDeviceRequest {
	s.OwnerId = &v
	return s
}

type StartDeviceResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDeviceResponseBody) SetId(v string) *StartDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *StartDeviceResponseBody) SetRequestId(v string) *StartDeviceResponseBody {
	s.RequestId = &v
	return s
}

type StartDeviceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDeviceResponse) GoString() string {
	return s.String()
}

func (s *StartDeviceResponse) SetHeaders(v map[string]*string) *StartDeviceResponse {
	s.Headers = v
	return s
}

func (s *StartDeviceResponse) SetBody(v *StartDeviceResponseBody) *StartDeviceResponse {
	s.Body = v
	return s
}

type StartParentPlatformRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StartParentPlatformRequest) String() string {
	return tea.Prettify(s)
}

func (s StartParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *StartParentPlatformRequest) SetId(v string) *StartParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *StartParentPlatformRequest) SetOwnerId(v int64) *StartParentPlatformRequest {
	s.OwnerId = &v
	return s
}

type StartParentPlatformResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartParentPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *StartParentPlatformResponseBody) SetId(v string) *StartParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *StartParentPlatformResponseBody) SetRequestId(v string) *StartParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

type StartParentPlatformResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartParentPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s StartParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *StartParentPlatformResponse) SetHeaders(v map[string]*string) *StartParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *StartParentPlatformResponse) SetBody(v *StartParentPlatformResponseBody) *StartParentPlatformResponse {
	s.Body = v
	return s
}

type StartRecordStreamRequest struct {
	App        *string `json:"App,omitempty" xml:"App,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
}

func (s StartRecordStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRecordStreamRequest) GoString() string {
	return s.String()
}

func (s *StartRecordStreamRequest) SetApp(v string) *StartRecordStreamRequest {
	s.App = &v
	return s
}

func (s *StartRecordStreamRequest) SetId(v string) *StartRecordStreamRequest {
	s.Id = &v
	return s
}

func (s *StartRecordStreamRequest) SetName(v string) *StartRecordStreamRequest {
	s.Name = &v
	return s
}

func (s *StartRecordStreamRequest) SetOwnerId(v int64) *StartRecordStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRecordStreamRequest) SetPlayDomain(v string) *StartRecordStreamRequest {
	s.PlayDomain = &v
	return s
}

type StartRecordStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRecordStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordStreamResponseBody) SetRequestId(v string) *StartRecordStreamResponseBody {
	s.RequestId = &v
	return s
}

type StartRecordStreamResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartRecordStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartRecordStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRecordStreamResponse) GoString() string {
	return s.String()
}

func (s *StartRecordStreamResponse) SetHeaders(v map[string]*string) *StartRecordStreamResponse {
	s.Headers = v
	return s
}

func (s *StartRecordStreamResponse) SetBody(v *StartRecordStreamResponseBody) *StartRecordStreamResponse {
	s.Body = v
	return s
}

type StartStreamRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StartStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s StartStreamRequest) GoString() string {
	return s.String()
}

func (s *StartStreamRequest) SetEndTime(v int64) *StartStreamRequest {
	s.EndTime = &v
	return s
}

func (s *StartStreamRequest) SetId(v string) *StartStreamRequest {
	s.Id = &v
	return s
}

func (s *StartStreamRequest) SetOwnerId(v int64) *StartStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartStreamRequest) SetStartTime(v int64) *StartStreamRequest {
	s.StartTime = &v
	return s
}

type StartStreamResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamResponseBody) SetId(v string) *StartStreamResponseBody {
	s.Id = &v
	return s
}

func (s *StartStreamResponseBody) SetName(v string) *StartStreamResponseBody {
	s.Name = &v
	return s
}

func (s *StartStreamResponseBody) SetRequestId(v string) *StartStreamResponseBody {
	s.RequestId = &v
	return s
}

type StartStreamResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s StartStreamResponse) GoString() string {
	return s.String()
}

func (s *StartStreamResponse) SetHeaders(v map[string]*string) *StartStreamResponse {
	s.Headers = v
	return s
}

func (s *StartStreamResponse) SetBody(v *StartStreamResponseBody) *StartStreamResponse {
	s.Body = v
	return s
}

type StartTransferStreamRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Transcode *string `json:"Transcode,omitempty" xml:"Transcode,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartTransferStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTransferStreamRequest) GoString() string {
	return s.String()
}

func (s *StartTransferStreamRequest) SetId(v string) *StartTransferStreamRequest {
	s.Id = &v
	return s
}

func (s *StartTransferStreamRequest) SetOwnerId(v int64) *StartTransferStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTransferStreamRequest) SetTranscode(v string) *StartTransferStreamRequest {
	s.Transcode = &v
	return s
}

func (s *StartTransferStreamRequest) SetUrl(v string) *StartTransferStreamRequest {
	s.Url = &v
	return s
}

type StartTransferStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTransferStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTransferStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartTransferStreamResponseBody) SetRequestId(v string) *StartTransferStreamResponseBody {
	s.RequestId = &v
	return s
}

type StartTransferStreamResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartTransferStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTransferStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTransferStreamResponse) GoString() string {
	return s.String()
}

func (s *StartTransferStreamResponse) SetHeaders(v map[string]*string) *StartTransferStreamResponse {
	s.Headers = v
	return s
}

func (s *StartTransferStreamResponse) SetBody(v *StartTransferStreamResponseBody) *StartTransferStreamResponse {
	s.Body = v
	return s
}

type StopAdjustRequest struct {
	Focus   *bool   `json:"Focus,omitempty" xml:"Focus,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Iris    *bool   `json:"Iris,omitempty" xml:"Iris,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StopAdjustRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAdjustRequest) GoString() string {
	return s.String()
}

func (s *StopAdjustRequest) SetFocus(v bool) *StopAdjustRequest {
	s.Focus = &v
	return s
}

func (s *StopAdjustRequest) SetId(v string) *StopAdjustRequest {
	s.Id = &v
	return s
}

func (s *StopAdjustRequest) SetIris(v bool) *StopAdjustRequest {
	s.Iris = &v
	return s
}

func (s *StopAdjustRequest) SetOwnerId(v int64) *StopAdjustRequest {
	s.OwnerId = &v
	return s
}

type StopAdjustResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAdjustResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAdjustResponseBody) GoString() string {
	return s.String()
}

func (s *StopAdjustResponseBody) SetId(v string) *StopAdjustResponseBody {
	s.Id = &v
	return s
}

func (s *StopAdjustResponseBody) SetRequestId(v string) *StopAdjustResponseBody {
	s.RequestId = &v
	return s
}

type StopAdjustResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopAdjustResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAdjustResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAdjustResponse) GoString() string {
	return s.String()
}

func (s *StopAdjustResponse) SetHeaders(v map[string]*string) *StopAdjustResponse {
	s.Headers = v
	return s
}

func (s *StopAdjustResponse) SetBody(v *StopAdjustResponseBody) *StopAdjustResponse {
	s.Body = v
	return s
}

type StopDeviceRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StopDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDeviceRequest) GoString() string {
	return s.String()
}

func (s *StopDeviceRequest) SetId(v string) *StopDeviceRequest {
	s.Id = &v
	return s
}

func (s *StopDeviceRequest) SetOwnerId(v int64) *StopDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDeviceRequest) SetStartTime(v string) *StopDeviceRequest {
	s.StartTime = &v
	return s
}

type StopDeviceResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDeviceResponseBody) SetId(v string) *StopDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *StopDeviceResponseBody) SetRequestId(v string) *StopDeviceResponseBody {
	s.RequestId = &v
	return s
}

type StopDeviceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDeviceResponse) GoString() string {
	return s.String()
}

func (s *StopDeviceResponse) SetHeaders(v map[string]*string) *StopDeviceResponse {
	s.Headers = v
	return s
}

func (s *StopDeviceResponse) SetBody(v *StopDeviceResponseBody) *StopDeviceResponse {
	s.Body = v
	return s
}

type StopMoveRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Pan     *bool   `json:"Pan,omitempty" xml:"Pan,omitempty"`
	Tilt    *bool   `json:"Tilt,omitempty" xml:"Tilt,omitempty"`
	Zoom    *bool   `json:"Zoom,omitempty" xml:"Zoom,omitempty"`
}

func (s StopMoveRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMoveRequest) GoString() string {
	return s.String()
}

func (s *StopMoveRequest) SetId(v string) *StopMoveRequest {
	s.Id = &v
	return s
}

func (s *StopMoveRequest) SetOwnerId(v int64) *StopMoveRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMoveRequest) SetPan(v bool) *StopMoveRequest {
	s.Pan = &v
	return s
}

func (s *StopMoveRequest) SetTilt(v bool) *StopMoveRequest {
	s.Tilt = &v
	return s
}

func (s *StopMoveRequest) SetZoom(v bool) *StopMoveRequest {
	s.Zoom = &v
	return s
}

type StopMoveResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMoveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMoveResponseBody) GoString() string {
	return s.String()
}

func (s *StopMoveResponseBody) SetId(v string) *StopMoveResponseBody {
	s.Id = &v
	return s
}

func (s *StopMoveResponseBody) SetRequestId(v string) *StopMoveResponseBody {
	s.RequestId = &v
	return s
}

type StopMoveResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopMoveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMoveResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMoveResponse) GoString() string {
	return s.String()
}

func (s *StopMoveResponse) SetHeaders(v map[string]*string) *StopMoveResponse {
	s.Headers = v
	return s
}

func (s *StopMoveResponse) SetBody(v *StopMoveResponseBody) *StopMoveResponse {
	s.Body = v
	return s
}

type StopParentPlatformRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StopParentPlatformRequest) String() string {
	return tea.Prettify(s)
}

func (s StopParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *StopParentPlatformRequest) SetId(v string) *StopParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *StopParentPlatformRequest) SetOwnerId(v int64) *StopParentPlatformRequest {
	s.OwnerId = &v
	return s
}

type StopParentPlatformResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopParentPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *StopParentPlatformResponseBody) SetId(v string) *StopParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *StopParentPlatformResponseBody) SetRequestId(v string) *StopParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

type StopParentPlatformResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopParentPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s StopParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *StopParentPlatformResponse) SetHeaders(v map[string]*string) *StopParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *StopParentPlatformResponse) SetBody(v *StopParentPlatformResponseBody) *StopParentPlatformResponse {
	s.Body = v
	return s
}

type StopRecordStreamRequest struct {
	App        *string `json:"App,omitempty" xml:"App,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
}

func (s StopRecordStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s StopRecordStreamRequest) GoString() string {
	return s.String()
}

func (s *StopRecordStreamRequest) SetApp(v string) *StopRecordStreamRequest {
	s.App = &v
	return s
}

func (s *StopRecordStreamRequest) SetId(v string) *StopRecordStreamRequest {
	s.Id = &v
	return s
}

func (s *StopRecordStreamRequest) SetName(v string) *StopRecordStreamRequest {
	s.Name = &v
	return s
}

func (s *StopRecordStreamRequest) SetOwnerId(v int64) *StopRecordStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRecordStreamRequest) SetPlayDomain(v string) *StopRecordStreamRequest {
	s.PlayDomain = &v
	return s
}

type StopRecordStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopRecordStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordStreamResponseBody) SetRequestId(v string) *StopRecordStreamResponseBody {
	s.RequestId = &v
	return s
}

type StopRecordStreamResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopRecordStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopRecordStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s StopRecordStreamResponse) GoString() string {
	return s.String()
}

func (s *StopRecordStreamResponse) SetHeaders(v map[string]*string) *StopRecordStreamResponse {
	s.Headers = v
	return s
}

func (s *StopRecordStreamResponse) SetBody(v *StopRecordStreamResponseBody) *StopRecordStreamResponse {
	s.Body = v
	return s
}

type StopStreamRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StopStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStreamRequest) GoString() string {
	return s.String()
}

func (s *StopStreamRequest) SetId(v string) *StopStreamRequest {
	s.Id = &v
	return s
}

func (s *StopStreamRequest) SetName(v string) *StopStreamRequest {
	s.Name = &v
	return s
}

func (s *StopStreamRequest) SetOwnerId(v int64) *StopStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopStreamRequest) SetStartTime(v string) *StopStreamRequest {
	s.StartTime = &v
	return s
}

type StopStreamResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamResponseBody) SetId(v string) *StopStreamResponseBody {
	s.Id = &v
	return s
}

func (s *StopStreamResponseBody) SetRequestId(v string) *StopStreamResponseBody {
	s.RequestId = &v
	return s
}

type StopStreamResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStreamResponse) GoString() string {
	return s.String()
}

func (s *StopStreamResponse) SetHeaders(v map[string]*string) *StopStreamResponse {
	s.Headers = v
	return s
}

func (s *StopStreamResponse) SetBody(v *StopStreamResponseBody) *StopStreamResponse {
	s.Body = v
	return s
}

type StopTransferStreamRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Transcode *string `json:"Transcode,omitempty" xml:"Transcode,omitempty"`
}

func (s StopTransferStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTransferStreamRequest) GoString() string {
	return s.String()
}

func (s *StopTransferStreamRequest) SetId(v string) *StopTransferStreamRequest {
	s.Id = &v
	return s
}

func (s *StopTransferStreamRequest) SetOwnerId(v int64) *StopTransferStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopTransferStreamRequest) SetTranscode(v string) *StopTransferStreamRequest {
	s.Transcode = &v
	return s
}

type StopTransferStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTransferStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTransferStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopTransferStreamResponseBody) SetRequestId(v string) *StopTransferStreamResponseBody {
	s.RequestId = &v
	return s
}

type StopTransferStreamResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopTransferStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopTransferStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTransferStreamResponse) GoString() string {
	return s.String()
}

func (s *StopTransferStreamResponse) SetHeaders(v map[string]*string) *StopTransferStreamResponse {
	s.Headers = v
	return s
}

func (s *StopTransferStreamResponse) SetBody(v *StopTransferStreamResponseBody) *StopTransferStreamResponse {
	s.Body = v
	return s
}

type SyncCatalogsRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SyncCatalogsRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCatalogsRequest) GoString() string {
	return s.String()
}

func (s *SyncCatalogsRequest) SetId(v string) *SyncCatalogsRequest {
	s.Id = &v
	return s
}

func (s *SyncCatalogsRequest) SetOwnerId(v int64) *SyncCatalogsRequest {
	s.OwnerId = &v
	return s
}

type SyncCatalogsResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncCatalogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *SyncCatalogsResponseBody) SetId(v string) *SyncCatalogsResponseBody {
	s.Id = &v
	return s
}

func (s *SyncCatalogsResponseBody) SetRequestId(v string) *SyncCatalogsResponseBody {
	s.RequestId = &v
	return s
}

type SyncCatalogsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncCatalogsResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncCatalogsResponse) GoString() string {
	return s.String()
}

func (s *SyncCatalogsResponse) SetHeaders(v map[string]*string) *SyncCatalogsResponse {
	s.Headers = v
	return s
}

func (s *SyncCatalogsResponse) SetBody(v *SyncCatalogsResponseBody) *SyncCatalogsResponse {
	s.Body = v
	return s
}

type UnbindDirectoryRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnbindDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UnbindDirectoryRequest) SetDeviceId(v string) *UnbindDirectoryRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindDirectoryRequest) SetDirectoryId(v string) *UnbindDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UnbindDirectoryRequest) SetOwnerId(v int64) *UnbindDirectoryRequest {
	s.OwnerId = &v
	return s
}

type UnbindDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDirectoryResponseBody) SetRequestId(v string) *UnbindDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type UnbindDirectoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UnbindDirectoryResponse) SetHeaders(v map[string]*string) *UnbindDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UnbindDirectoryResponse) SetBody(v *UnbindDirectoryResponseBody) *UnbindDirectoryResponse {
	s.Body = v
	return s
}

type UnbindParentPlatformDeviceRequest struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s UnbindParentPlatformDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindParentPlatformDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindParentPlatformDeviceRequest) SetDeviceId(v string) *UnbindParentPlatformDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindParentPlatformDeviceRequest) SetOwnerId(v int64) *UnbindParentPlatformDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindParentPlatformDeviceRequest) SetParentPlatformId(v string) *UnbindParentPlatformDeviceRequest {
	s.ParentPlatformId = &v
	return s
}

type UnbindParentPlatformDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindParentPlatformDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindParentPlatformDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindParentPlatformDeviceResponseBody) SetRequestId(v string) *UnbindParentPlatformDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnbindParentPlatformDeviceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindParentPlatformDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindParentPlatformDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindParentPlatformDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindParentPlatformDeviceResponse) SetHeaders(v map[string]*string) *UnbindParentPlatformDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindParentPlatformDeviceResponse) SetBody(v *UnbindParentPlatformDeviceResponseBody) *UnbindParentPlatformDeviceResponse {
	s.Body = v
	return s
}

type UnbindPurchasedDeviceRequest struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnbindPurchasedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindPurchasedDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindPurchasedDeviceRequest) SetDeviceId(v string) *UnbindPurchasedDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindPurchasedDeviceRequest) SetOwnerId(v int64) *UnbindPurchasedDeviceRequest {
	s.OwnerId = &v
	return s
}

type UnbindPurchasedDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindPurchasedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindPurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPurchasedDeviceResponseBody) SetRequestId(v string) *UnbindPurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnbindPurchasedDeviceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindPurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindPurchasedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindPurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindPurchasedDeviceResponse) SetHeaders(v map[string]*string) *UnbindPurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindPurchasedDeviceResponse) SetBody(v *UnbindPurchasedDeviceResponseBody) *UnbindPurchasedDeviceResponse {
	s.Body = v
	return s
}

type UnbindTemplateRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s UnbindTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindTemplateRequest) GoString() string {
	return s.String()
}

func (s *UnbindTemplateRequest) SetInstanceId(v string) *UnbindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindTemplateRequest) SetInstanceType(v string) *UnbindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *UnbindTemplateRequest) SetOwnerId(v int64) *UnbindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindTemplateRequest) SetTemplateId(v string) *UnbindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UnbindTemplateRequest) SetTemplateType(v string) *UnbindTemplateRequest {
	s.TemplateType = &v
	return s
}

type UnbindTemplateResponseBody struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s UnbindTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindTemplateResponseBody) SetInstanceId(v string) *UnbindTemplateResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetInstanceType(v string) *UnbindTemplateResponseBody {
	s.InstanceType = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetRequestId(v string) *UnbindTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetTemplateId(v string) *UnbindTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetTemplateType(v string) *UnbindTemplateResponseBody {
	s.TemplateType = &v
	return s
}

type UnbindTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindTemplateResponse) GoString() string {
	return s.String()
}

func (s *UnbindTemplateResponse) SetHeaders(v map[string]*string) *UnbindTemplateResponse {
	s.Headers = v
	return s
}

func (s *UnbindTemplateResponse) SetBody(v *UnbindTemplateResponseBody) *UnbindTemplateResponse {
	s.Body = v
	return s
}

type UnlockDeviceRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnlockDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockDeviceRequest) SetId(v string) *UnlockDeviceRequest {
	s.Id = &v
	return s
}

func (s *UnlockDeviceRequest) SetOwnerId(v int64) *UnlockDeviceRequest {
	s.OwnerId = &v
	return s
}

type UnlockDeviceResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockDeviceResponseBody) SetId(v string) *UnlockDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *UnlockDeviceResponseBody) SetRequestId(v string) *UnlockDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnlockDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnlockDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockDeviceResponse) SetHeaders(v map[string]*string) *UnlockDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockDeviceResponse) SetBody(v *UnlockDeviceResponseBody) *UnlockDeviceResponse {
	s.Body = v
	return s
}

type UpdateAIConfigRequest struct {
	CaptureInterval *int32  `json:"CaptureInterval,omitempty" xml:"CaptureInterval,omitempty"`
	ConfigId        *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Configs         *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Features        *string `json:"Features,omitempty" xml:"Features,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAIConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAIConfigRequest) SetCaptureInterval(v int32) *UpdateAIConfigRequest {
	s.CaptureInterval = &v
	return s
}

func (s *UpdateAIConfigRequest) SetConfigId(v string) *UpdateAIConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateAIConfigRequest) SetConfigs(v string) *UpdateAIConfigRequest {
	s.Configs = &v
	return s
}

func (s *UpdateAIConfigRequest) SetDescription(v string) *UpdateAIConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateAIConfigRequest) SetEndTime(v int64) *UpdateAIConfigRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateAIConfigRequest) SetFeatures(v string) *UpdateAIConfigRequest {
	s.Features = &v
	return s
}

func (s *UpdateAIConfigRequest) SetOwnerId(v int64) *UpdateAIConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAIConfigRequest) SetStartTime(v int64) *UpdateAIConfigRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateAIConfigRequest) SetStatus(v string) *UpdateAIConfigRequest {
	s.Status = &v
	return s
}

type UpdateAIConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAIConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAIConfigResponseBody) SetRequestId(v string) *UpdateAIConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAIConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAIConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAIConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAIConfigResponse) SetHeaders(v map[string]*string) *UpdateAIConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAIConfigResponse) SetBody(v *UpdateAIConfigResponseBody) *UpdateAIConfigResponse {
	s.Body = v
	return s
}

type UpdateBucketInfoRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Comment    *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UpdateBucketInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBucketInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateBucketInfoRequest) SetBucketName(v string) *UpdateBucketInfoRequest {
	s.BucketName = &v
	return s
}

func (s *UpdateBucketInfoRequest) SetComment(v string) *UpdateBucketInfoRequest {
	s.Comment = &v
	return s
}

func (s *UpdateBucketInfoRequest) SetOwnerId(v int64) *UpdateBucketInfoRequest {
	s.OwnerId = &v
	return s
}

type UpdateBucketInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBucketInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBucketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBucketInfoResponseBody) SetRequestId(v string) *UpdateBucketInfoResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBucketInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateBucketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBucketInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBucketInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateBucketInfoResponse) SetHeaders(v map[string]*string) *UpdateBucketInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateBucketInfoResponse) SetBody(v *UpdateBucketInfoResponseBody) *UpdateBucketInfoResponse {
	s.Body = v
	return s
}

type UpdateClusterRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveTime   *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	InternalPorts   *string `json:"InternalPorts,omitempty" xml:"InternalPorts,omitempty"`
	MaintainTime    *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s UpdateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) SetClusterId(v string) *UpdateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterRequest) SetDescription(v string) *UpdateClusterRequest {
	s.Description = &v
	return s
}

func (s *UpdateClusterRequest) SetEffectiveTime(v string) *UpdateClusterRequest {
	s.EffectiveTime = &v
	return s
}

func (s *UpdateClusterRequest) SetInternalPorts(v string) *UpdateClusterRequest {
	s.InternalPorts = &v
	return s
}

func (s *UpdateClusterRequest) SetMaintainTime(v string) *UpdateClusterRequest {
	s.MaintainTime = &v
	return s
}

func (s *UpdateClusterRequest) SetName(v string) *UpdateClusterRequest {
	s.Name = &v
	return s
}

func (s *UpdateClusterRequest) SetOwnerId(v int64) *UpdateClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateClusterRequest) SetSecurityGroupId(v string) *UpdateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

type UpdateClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponse) SetHeaders(v map[string]*string) *UpdateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterResponse) SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse {
	s.Body = v
	return s
}

type UpdateRenderingDeviceSpecRequest struct {
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod *int64  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveTime   *bool   `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	InstanceIds     *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Specification   *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s UpdateRenderingDeviceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRenderingDeviceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingDeviceSpecRequest) SetAutoRenew(v bool) *UpdateRenderingDeviceSpecRequest {
	s.AutoRenew = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetAutoRenewPeriod(v int64) *UpdateRenderingDeviceSpecRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetDescription(v string) *UpdateRenderingDeviceSpecRequest {
	s.Description = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetEffectiveTime(v bool) *UpdateRenderingDeviceSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetInstanceIds(v string) *UpdateRenderingDeviceSpecRequest {
	s.InstanceIds = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetOwnerId(v int64) *UpdateRenderingDeviceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetPeriod(v int32) *UpdateRenderingDeviceSpecRequest {
	s.Period = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetPeriodUnit(v string) *UpdateRenderingDeviceSpecRequest {
	s.PeriodUnit = &v
	return s
}

func (s *UpdateRenderingDeviceSpecRequest) SetSpecification(v string) *UpdateRenderingDeviceSpecRequest {
	s.Specification = &v
	return s
}

type UpdateRenderingDeviceSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRenderingDeviceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRenderingDeviceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRenderingDeviceSpecResponseBody) SetRequestId(v string) *UpdateRenderingDeviceSpecResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRenderingDeviceSpecResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRenderingDeviceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRenderingDeviceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRenderingDeviceSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateRenderingDeviceSpecResponse) SetHeaders(v map[string]*string) *UpdateRenderingDeviceSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateRenderingDeviceSpecResponse) SetBody(v *UpdateRenderingDeviceSpecResponseBody) *UpdateRenderingDeviceSpecResponse {
	s.Body = v
	return s
}

type UpdateVsPullStreamInfoConfigRequest struct {
	Always     *string `json:"Always,omitempty" xml:"Always,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s UpdateVsPullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetAlways(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.Always = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetAppName(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetDomainName(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetEndTime(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *UpdateVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetSourceUrl(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetStartTime(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigRequest) SetStreamName(v string) *UpdateVsPullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

type UpdateVsPullStreamInfoConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVsPullStreamInfoConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *UpdateVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVsPullStreamInfoConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVsPullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *UpdateVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVsPullStreamInfoConfigResponse) SetBody(v *UpdateVsPullStreamInfoConfigResponseBody) *UpdateVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

type UpgradeRenderingDevicesHostOSRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RomName     *string `json:"RomName,omitempty" xml:"RomName,omitempty"`
	RomPath     *string `json:"RomPath,omitempty" xml:"RomPath,omitempty"`
	RomVersion  *string `json:"RomVersion,omitempty" xml:"RomVersion,omitempty"`
}

func (s UpgradeRenderingDevicesHostOSRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeRenderingDevicesHostOSRequest) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingDevicesHostOSRequest) SetInstanceIds(v string) *UpgradeRenderingDevicesHostOSRequest {
	s.InstanceIds = &v
	return s
}

func (s *UpgradeRenderingDevicesHostOSRequest) SetOwnerId(v int64) *UpgradeRenderingDevicesHostOSRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeRenderingDevicesHostOSRequest) SetRomName(v string) *UpgradeRenderingDevicesHostOSRequest {
	s.RomName = &v
	return s
}

func (s *UpgradeRenderingDevicesHostOSRequest) SetRomPath(v string) *UpgradeRenderingDevicesHostOSRequest {
	s.RomPath = &v
	return s
}

func (s *UpgradeRenderingDevicesHostOSRequest) SetRomVersion(v string) *UpgradeRenderingDevicesHostOSRequest {
	s.RomVersion = &v
	return s
}

type UpgradeRenderingDevicesHostOSResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeRenderingDevicesHostOSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeRenderingDevicesHostOSResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingDevicesHostOSResponseBody) SetRequestId(v string) *UpgradeRenderingDevicesHostOSResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeRenderingDevicesHostOSResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeRenderingDevicesHostOSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeRenderingDevicesHostOSResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeRenderingDevicesHostOSResponse) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingDevicesHostOSResponse) SetHeaders(v map[string]*string) *UpgradeRenderingDevicesHostOSResponse {
	s.Headers = v
	return s
}

func (s *UpgradeRenderingDevicesHostOSResponse) SetBody(v *UpgradeRenderingDevicesHostOSResponseBody) *UpgradeRenderingDevicesHostOSResponse {
	s.Body = v
	return s
}

type UpgradeRenderingDevicesImageRequest struct {
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UpgradeRenderingDevicesImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeRenderingDevicesImageRequest) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingDevicesImageRequest) SetImageId(v string) *UpgradeRenderingDevicesImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeRenderingDevicesImageRequest) SetInstanceIds(v string) *UpgradeRenderingDevicesImageRequest {
	s.InstanceIds = &v
	return s
}

func (s *UpgradeRenderingDevicesImageRequest) SetOwnerId(v int64) *UpgradeRenderingDevicesImageRequest {
	s.OwnerId = &v
	return s
}

type UpgradeRenderingDevicesImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeRenderingDevicesImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeRenderingDevicesImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingDevicesImageResponseBody) SetRequestId(v string) *UpgradeRenderingDevicesImageResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeRenderingDevicesImageResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeRenderingDevicesImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeRenderingDevicesImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeRenderingDevicesImageResponse) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingDevicesImageResponse) SetHeaders(v map[string]*string) *UpgradeRenderingDevicesImageResponse {
	s.Headers = v
	return s
}

func (s *UpgradeRenderingDevicesImageResponse) SetBody(v *UpgradeRenderingDevicesImageResponseBody) *UpgradeRenderingDevicesImageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("vs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddRegisteredDeviceWithOptions(request *AddRegisteredDeviceRequest, runtime *util.RuntimeOptions) (_result *AddRegisteredDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dsn)) {
		query["Dsn"] = request.Dsn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterCode)) {
		query["RegisterCode"] = request.RegisterCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		query["SecretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRegisteredDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRegisteredDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRegisteredDevice(request *AddRegisteredDeviceRequest) (_result *AddRegisteredDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRegisteredDeviceResponse{}
	_body, _err := client.AddRegisteredDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRegisteredVendorWithOptions(request *AddRegisteredVendorRequest, runtime *util.RuntimeOptions) (_result *AddRegisteredVendorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRegisteredVendor"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRegisteredVendorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRegisteredVendor(request *AddRegisteredVendorRequest) (_result *AddRegisteredVendorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRegisteredVendorResponse{}
	_body, _err := client.AddRegisteredVendorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRenderingDeviceInternetPortsWithOptions(request *AddRenderingDeviceInternetPortsRequest, runtime *util.RuntimeOptions) (_result *AddRenderingDeviceInternetPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InternalPort)) {
		query["InternalPort"] = request.InternalPort
	}

	if !tea.BoolValue(util.IsUnset(request.IpProtocol)) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRenderingDeviceInternetPorts"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRenderingDeviceInternetPortsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRenderingDeviceInternetPorts(request *AddRenderingDeviceInternetPortsRequest) (_result *AddRenderingDeviceInternetPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRenderingDeviceInternetPortsResponse{}
	_body, _err := client.AddRenderingDeviceInternetPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVsPullStreamInfoConfigWithOptions(request *AddVsPullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *AddVsPullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Always)) {
		query["Always"] = request.Always
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		query["SourceUrl"] = request.SourceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVsPullStreamInfoConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVsPullStreamInfoConfig(request *AddVsPullStreamInfoConfigRequest) (_result *AddVsPullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVsPullStreamInfoConfigResponse{}
	_body, _err := client.AddVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchBindDirectoriesWithOptions(request *BatchBindDirectoriesRequest, runtime *util.RuntimeOptions) (_result *BatchBindDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchBindDirectories"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBindDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchBindDirectories(request *BatchBindDirectoriesRequest) (_result *BatchBindDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchBindDirectoriesResponse{}
	_body, _err := client.BatchBindDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchBindParentPlatformDevicesWithOptions(request *BatchBindParentPlatformDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchBindParentPlatformDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentPlatformId)) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchBindParentPlatformDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBindParentPlatformDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchBindParentPlatformDevices(request *BatchBindParentPlatformDevicesRequest) (_result *BatchBindParentPlatformDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchBindParentPlatformDevicesResponse{}
	_body, _err := client.BatchBindParentPlatformDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchBindPurchasedDevicesWithOptions(request *BatchBindPurchasedDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchBindPurchasedDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchBindPurchasedDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBindPurchasedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchBindPurchasedDevices(request *BatchBindPurchasedDevicesRequest) (_result *BatchBindPurchasedDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchBindPurchasedDevicesResponse{}
	_body, _err := client.BatchBindPurchasedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchBindTemplateWithOptions(request *BatchBindTemplateRequest, runtime *util.RuntimeOptions) (_result *BatchBindTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyAll)) {
		query["ApplyAll"] = request.ApplyAll
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Replace)) {
		query["Replace"] = request.Replace
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchBindTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchBindTemplate(request *BatchBindTemplateRequest) (_result *BatchBindTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchBindTemplateResponse{}
	_body, _err := client.BatchBindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchBindTemplatesWithOptions(request *BatchBindTemplatesRequest, runtime *util.RuntimeOptions) (_result *BatchBindTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyAll)) {
		query["ApplyAll"] = request.ApplyAll
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Replace)) {
		query["Replace"] = request.Replace
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchBindTemplates"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBindTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchBindTemplates(request *BatchBindTemplatesRequest) (_result *BatchBindTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchBindTemplatesResponse{}
	_body, _err := client.BatchBindTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteDevicesWithOptions(request *BatchDeleteDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteDevices(request *BatchDeleteDevicesRequest) (_result *BatchDeleteDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteDevicesResponse{}
	_body, _err := client.BatchDeleteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteVsDomainConfigsWithOptions(request *BatchDeleteVsDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteVsDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteVsDomainConfigs"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteVsDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteVsDomainConfigs(request *BatchDeleteVsDomainConfigsRequest) (_result *BatchDeleteVsDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteVsDomainConfigsResponse{}
	_body, _err := client.BatchDeleteVsDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchForbidVsStreamWithOptions(request *BatchForbidVsStreamRequest, runtime *util.RuntimeOptions) (_result *BatchForbidVsStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ControlStreamAction)) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.LiveStreamType)) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !tea.BoolValue(util.IsUnset(request.Oneshot)) {
		query["Oneshot"] = request.Oneshot
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeTime)) {
		query["ResumeTime"] = request.ResumeTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchForbidVsStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchForbidVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchForbidVsStream(request *BatchForbidVsStreamRequest) (_result *BatchForbidVsStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchForbidVsStreamResponse{}
	_body, _err := client.BatchForbidVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchResumeVsStreamWithOptions(request *BatchResumeVsStreamRequest, runtime *util.RuntimeOptions) (_result *BatchResumeVsStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ControlStreamAction)) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.LiveStreamType)) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchResumeVsStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchResumeVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchResumeVsStream(request *BatchResumeVsStreamRequest) (_result *BatchResumeVsStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchResumeVsStreamResponse{}
	_body, _err := client.BatchResumeVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetVsDomainConfigsWithOptions(request *BatchSetVsDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchSetVsDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.Functions)) {
		query["Functions"] = request.Functions
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSetVsDomainConfigs"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSetVsDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetVsDomainConfigs(request *BatchSetVsDomainConfigsRequest) (_result *BatchSetVsDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetVsDomainConfigsResponse{}
	_body, _err := client.BatchSetVsDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartDevicesWithOptions(request *BatchStartDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchStartDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStartDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStartDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartDevices(request *BatchStartDevicesRequest) (_result *BatchStartDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStartDevicesResponse{}
	_body, _err := client.BatchStartDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartStreamsWithOptions(request *BatchStartStreamsRequest, runtime *util.RuntimeOptions) (_result *BatchStartStreamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStartStreams"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStartStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartStreams(request *BatchStartStreamsRequest) (_result *BatchStartStreamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStartStreamsResponse{}
	_body, _err := client.BatchStartStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopDevicesWithOptions(request *BatchStopDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchStopDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStopDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStopDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopDevices(request *BatchStopDevicesRequest) (_result *BatchStopDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopDevicesResponse{}
	_body, _err := client.BatchStopDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopStreamsWithOptions(request *BatchStopStreamsRequest, runtime *util.RuntimeOptions) (_result *BatchStopStreamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStopStreams"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStopStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopStreams(request *BatchStopStreamsRequest) (_result *BatchStopStreamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopStreamsResponse{}
	_body, _err := client.BatchStopStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUnbindDirectoriesWithOptions(request *BatchUnbindDirectoriesRequest, runtime *util.RuntimeOptions) (_result *BatchUnbindDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUnbindDirectories"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUnbindDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUnbindDirectories(request *BatchUnbindDirectoriesRequest) (_result *BatchUnbindDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUnbindDirectoriesResponse{}
	_body, _err := client.BatchUnbindDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUnbindParentPlatformDevicesWithOptions(request *BatchUnbindParentPlatformDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchUnbindParentPlatformDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentPlatformId)) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUnbindParentPlatformDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUnbindParentPlatformDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUnbindParentPlatformDevices(request *BatchUnbindParentPlatformDevicesRequest) (_result *BatchUnbindParentPlatformDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUnbindParentPlatformDevicesResponse{}
	_body, _err := client.BatchUnbindParentPlatformDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUnbindPurchasedDevicesWithOptions(request *BatchUnbindPurchasedDevicesRequest, runtime *util.RuntimeOptions) (_result *BatchUnbindPurchasedDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUnbindPurchasedDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUnbindPurchasedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUnbindPurchasedDevices(request *BatchUnbindPurchasedDevicesRequest) (_result *BatchUnbindPurchasedDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUnbindPurchasedDevicesResponse{}
	_body, _err := client.BatchUnbindPurchasedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUnbindTemplateWithOptions(request *BatchUnbindTemplateRequest, runtime *util.RuntimeOptions) (_result *BatchUnbindTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUnbindTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUnbindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUnbindTemplate(request *BatchUnbindTemplateRequest) (_result *BatchUnbindTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUnbindTemplateResponse{}
	_body, _err := client.BatchUnbindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUnbindTemplatesWithOptions(request *BatchUnbindTemplatesRequest, runtime *util.RuntimeOptions) (_result *BatchUnbindTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUnbindTemplates"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUnbindTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUnbindTemplates(request *BatchUnbindTemplatesRequest) (_result *BatchUnbindTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUnbindTemplatesResponse{}
	_body, _err := client.BatchUnbindTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindDirectoryWithOptions(request *BindDirectoryRequest, runtime *util.RuntimeOptions) (_result *BindDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindDirectory"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindDirectory(request *BindDirectoryRequest) (_result *BindDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindDirectoryResponse{}
	_body, _err := client.BindDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindParentPlatformDeviceWithOptions(request *BindParentPlatformDeviceRequest, runtime *util.RuntimeOptions) (_result *BindParentPlatformDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentPlatformId)) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindParentPlatformDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindParentPlatformDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindParentPlatformDevice(request *BindParentPlatformDeviceRequest) (_result *BindParentPlatformDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindParentPlatformDeviceResponse{}
	_body, _err := client.BindParentPlatformDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindPurchasedDeviceWithOptions(request *BindPurchasedDeviceRequest, runtime *util.RuntimeOptions) (_result *BindPurchasedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindPurchasedDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindPurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindPurchasedDevice(request *BindPurchasedDeviceRequest) (_result *BindPurchasedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindPurchasedDeviceResponse{}
	_body, _err := client.BindPurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindTemplateWithOptions(request *BindTemplateRequest, runtime *util.RuntimeOptions) (_result *BindTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyAll)) {
		query["ApplyAll"] = request.ApplyAll
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Replace)) {
		query["Replace"] = request.Replace
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindTemplate(request *BindTemplateRequest) (_result *BindTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindTemplateResponse{}
	_body, _err := client.BindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CaptureDeviceSnapshotWithOptions(request *CaptureDeviceSnapshotRequest, runtime *util.RuntimeOptions) (_result *CaptureDeviceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotConfig)) {
		query["SnapshotConfig"] = request.SnapshotConfig
	}

	if !tea.BoolValue(util.IsUnset(request.StreamId)) {
		query["StreamId"] = request.StreamId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CaptureDeviceSnapshot"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CaptureDeviceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CaptureDeviceSnapshot(request *CaptureDeviceSnapshotRequest) (_result *CaptureDeviceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CaptureDeviceSnapshotResponse{}
	_body, _err := client.CaptureDeviceSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContinuousAdjustWithOptions(request *ContinuousAdjustRequest, runtime *util.RuntimeOptions) (_result *ContinuousAdjustResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Focus)) {
		query["Focus"] = request.Focus
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Iris)) {
		query["Iris"] = request.Iris
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinuousAdjust"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinuousAdjustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContinuousAdjust(request *ContinuousAdjustRequest) (_result *ContinuousAdjustResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinuousAdjustResponse{}
	_body, _err := client.ContinuousAdjustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContinuousMoveWithOptions(request *ContinuousMoveRequest, runtime *util.RuntimeOptions) (_result *ContinuousMoveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Pan)) {
		query["Pan"] = request.Pan
	}

	if !tea.BoolValue(util.IsUnset(request.Tilt)) {
		query["Tilt"] = request.Tilt
	}

	if !tea.BoolValue(util.IsUnset(request.Zoom)) {
		query["Zoom"] = request.Zoom
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinuousMove"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinuousMoveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContinuousMove(request *ContinuousMoveRequest) (_result *ContinuousMoveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinuousMoveResponse{}
	_body, _err := client.ContinuousMoveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAIConfigWithOptions(request *CreateAIConfigRequest, runtime *util.RuntimeOptions) (_result *CreateAIConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaptureInterval)) {
		query["CaptureInterval"] = request.CaptureInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Configs)) {
		query["Configs"] = request.Configs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Features)) {
		query["Features"] = request.Features
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAIConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAIConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAIConfig(request *CreateAIConfigRequest) (_result *CreateAIConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAIConfigResponse{}
	_body, _err := client.CreateAIConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InternalPorts)) {
		query["InternalPorts"] = request.InternalPorts
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainTime)) {
		query["MaintainTime"] = request.MaintainTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceWithOptions(request *CreateDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmMethod)) {
		query["AlarmMethod"] = request.AlarmMethod
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPos)) {
		query["AutoPos"] = request.AutoPos
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStart)) {
		query["AutoStart"] = request.AutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Dsn)) {
		query["Dsn"] = request.Dsn
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["Latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["Longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.PosInterval)) {
		query["PosInterval"] = request.PosInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevice(request *CreateDeviceRequest) (_result *CreateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CreateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceAlarmWithOptions(request *CreateDeviceAlarmRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alarm)) {
		query["Alarm"] = request.Alarm
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Expire)) {
		query["Expire"] = request.Expire
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubAlarm)) {
		query["SubAlarm"] = request.SubAlarm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeviceAlarm"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceAlarm(request *CreateDeviceAlarmRequest) (_result *CreateDeviceAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceAlarmResponse{}
	_body, _err := client.CreateDeviceAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDirectoryWithOptions(request *CreateDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDirectory"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDirectory(request *CreateDirectoryRequest) (_result *CreateDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CreateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.App)) {
		query["App"] = request.App
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InProtocol)) {
		query["InProtocol"] = request.InProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LazyPull)) {
		query["LazyPull"] = request.LazyPull
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OutProtocol)) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayDomain)) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !tea.BoolValue(util.IsUnset(request.PushDomain)) {
		query["PushDomain"] = request.PushDomain
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateParentPlatformWithOptions(request *CreateParentPlatformRequest, runtime *util.RuntimeOptions) (_result *CreateParentPlatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoStart)) {
		query["AutoStart"] = request.AutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.ClientAuth)) {
		query["ClientAuth"] = request.ClientAuth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientPassword)) {
		query["ClientPassword"] = request.ClientPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUsername)) {
		query["ClientUsername"] = request.ClientUsername
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateParentPlatform"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateParentPlatform(request *CreateParentPlatformRequest) (_result *CreateParentPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateParentPlatformResponse{}
	_body, _err := client.CreateParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRenderingDeviceWithOptions(request *CreateRenderingDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateRenderingDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EdgeNodeName)) {
		query["EdgeNodeName"] = request.EdgeNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceChargeType)) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRenderingDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRenderingDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRenderingDevice(request *CreateRenderingDeviceRequest) (_result *CreateRenderingDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRenderingDeviceResponse{}
	_body, _err := client.CreateRenderingDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStreamSnapshotWithOptions(request *CreateStreamSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateStreamSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		query["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStreamSnapshot"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStreamSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStreamSnapshot(request *CreateStreamSnapshotRequest) (_result *CreateStreamSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStreamSnapshotResponse{}
	_body, _err := client.CreateStreamSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileFormat)) {
		query["FileFormat"] = request.FileFormat
	}

	if !tea.BoolValue(util.IsUnset(request.Flv)) {
		query["Flv"] = request.Flv
	}

	if !tea.BoolValue(util.IsUnset(request.HlsM3u8)) {
		query["HlsM3u8"] = request.HlsM3u8
	}

	if !tea.BoolValue(util.IsUnset(request.HlsTs)) {
		query["HlsTs"] = request.HlsTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.JpgOnDemand)) {
		query["JpgOnDemand"] = request.JpgOnDemand
	}

	if !tea.BoolValue(util.IsUnset(request.JpgOverwrite)) {
		query["JpgOverwrite"] = request.JpgOverwrite
	}

	if !tea.BoolValue(util.IsUnset(request.JpgSequence)) {
		query["JpgSequence"] = request.JpgSequence
	}

	if !tea.BoolValue(util.IsUnset(request.Mp4)) {
		query["Mp4"] = request.Mp4
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.OssFilePrefix)) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Retention)) {
		query["Retention"] = request.Retention
	}

	if !tea.BoolValue(util.IsUnset(request.TransConfigsJSON)) {
		query["TransConfigsJSON"] = request.TransConfigsJSON
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		query["Trigger"] = request.Trigger
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAIConfigWithOptions(request *DeleteAIConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteAIConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAIConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAIConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAIConfig(request *DeleteAIConfigRequest) (_result *DeleteAIConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAIConfigResponse{}
	_body, _err := client.DeleteAIConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketWithOptions(request *DeleteBucketRequest, runtime *util.RuntimeOptions) (_result *DeleteBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucket"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucket(request *DeleteBucketRequest) (_result *DeleteBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBucketResponse{}
	_body, _err := client.DeleteBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (_result *DeleteDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DeleteDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectory"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectory(request *DeleteDirectoryRequest) (_result *DeleteDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.DeleteDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteParentPlatformWithOptions(request *DeleteParentPlatformRequest, runtime *util.RuntimeOptions) (_result *DeleteParentPlatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteParentPlatform"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteParentPlatform(request *DeleteParentPlatformRequest) (_result *DeleteParentPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteParentPlatformResponse{}
	_body, _err := client.DeleteParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePresetWithOptions(request *DeletePresetRequest, runtime *util.RuntimeOptions) (_result *DeletePresetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PresetId)) {
		query["PresetId"] = request.PresetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePreset"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePresetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePreset(request *DeletePresetRequest) (_result *DeletePresetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePresetResponse{}
	_body, _err := client.DeletePresetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePurchasedDeviceWithOptions(request *DeletePurchasedDeviceRequest, runtime *util.RuntimeOptions) (_result *DeletePurchasedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePurchasedDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePurchasedDevice(request *DeletePurchasedDeviceRequest) (_result *DeletePurchasedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePurchasedDeviceResponse{}
	_body, _err := client.DeletePurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRenderingDeviceInternetPortsWithOptions(request *DeleteRenderingDeviceInternetPortsRequest, runtime *util.RuntimeOptions) (_result *DeleteRenderingDeviceInternetPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InternalPort)) {
		query["InternalPort"] = request.InternalPort
	}

	if !tea.BoolValue(util.IsUnset(request.IpProtocol)) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRenderingDeviceInternetPorts"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRenderingDeviceInternetPortsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRenderingDeviceInternetPorts(request *DeleteRenderingDeviceInternetPortsRequest) (_result *DeleteRenderingDeviceInternetPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRenderingDeviceInternetPortsResponse{}
	_body, _err := client.DeleteRenderingDeviceInternetPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRenderingDevicesWithOptions(request *DeleteRenderingDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteRenderingDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRenderingDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRenderingDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRenderingDevices(request *DeleteRenderingDevicesRequest) (_result *DeleteRenderingDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRenderingDevicesResponse{}
	_body, _err := client.DeleteRenderingDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVsPullStreamInfoConfigWithOptions(request *DeleteVsPullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteVsPullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVsPullStreamInfoConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVsPullStreamInfoConfig(request *DeleteVsPullStreamInfoConfigRequest) (_result *DeleteVsPullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVsPullStreamInfoConfigResponse{}
	_body, _err := client.DeleteVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVsStreamsNotifyUrlConfigWithOptions(request *DeleteVsStreamsNotifyUrlConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteVsStreamsNotifyUrlConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVsStreamsNotifyUrlConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVsStreamsNotifyUrlConfig(request *DeleteVsStreamsNotifyUrlConfigRequest) (_result *DeleteVsStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DeleteVsStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAIConfigWithOptions(request *DescribeAIConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAIConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAIConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAIConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAIConfig(request *DescribeAIConfigRequest) (_result *DescribeAIConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAIConfigResponse{}
	_body, _err := client.DescribeAIConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAIConfigListWithOptions(request *DescribeAIConfigListRequest, runtime *util.RuntimeOptions) (_result *DescribeAIConfigListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAIConfigList"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAIConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAIConfigList(request *DescribeAIConfigListRequest) (_result *DescribeAIConfigListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAIConfigListResponse{}
	_body, _err := client.DescribeAIConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAIEventListWithOptions(request *DescribeAIEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeAIEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAIEventList"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAIEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAIEventList(request *DescribeAIEventListRequest) (_result *DescribeAIEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAIEventListResponse{}
	_body, _err := client.DescribeAIEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountStatWithOptions(request *DescribeAccountStatRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountStat"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountStat(request *DescribeAccountStatRequest) (_result *DescribeAccountStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountStatResponse{}
	_body, _err := client.DescribeAccountStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCluster"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterDevicesWithOptions(request *DescribeClusterDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EdgeNodeName)) {
		query["EdgeNodeName"] = request.EdgeNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDevices(request *DescribeClusterDevicesRequest) (_result *DescribeClusterDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterDevicesResponse{}
	_body, _err := client.DescribeClusterDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusters"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerInstanceIdWithOptions(request *DescribeContainerInstanceIdRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PodIndex)) {
		query["PodIndex"] = request.PodIndex
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerInstanceId"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerInstanceId(request *DescribeContainerInstanceIdRequest) (_result *DescribeContainerInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerInstanceIdResponse{}
	_body, _err := client.DescribeContainerInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceWithOptions(request *DescribeDeviceRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDirectory)) {
		query["IncludeDirectory"] = request.IncludeDirectory
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeStats)) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDevice(request *DescribeDeviceRequest) (_result *DescribeDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceResponse{}
	_body, _err := client.DescribeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceChannelsWithOptions(request *DescribeDeviceChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceChannels"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceChannels(request *DescribeDeviceChannelsRequest) (_result *DescribeDeviceChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceChannelsResponse{}
	_body, _err := client.DescribeDeviceChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceGatewayWithOptions(request *DescribeDeviceGatewayRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.Expire)) {
		query["Expire"] = request.Expire
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceGateway"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceGateway(request *DescribeDeviceGatewayRequest) (_result *DescribeDeviceGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceGatewayResponse{}
	_body, _err := client.DescribeDeviceGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceURLWithOptions(request *DescribeDeviceURLRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Auth)) {
		query["Auth"] = request.Auth
	}

	if !tea.BoolValue(util.IsUnset(request.Expire)) {
		query["Expire"] = request.Expire
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OutProtocol)) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		query["Stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceURL"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceURL(request *DescribeDeviceURLRequest) (_result *DescribeDeviceURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceURLResponse{}
	_body, _err := client.DescribeDeviceURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDevicesWithOptions(request *DescribeDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Dsn)) {
		query["Dsn"] = request.Dsn
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDirectory)) {
		query["IncludeDirectory"] = request.IncludeDirectory
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeStats)) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDevices(request *DescribeDevicesRequest) (_result *DescribeDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.DescribeDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NoPagination)) {
		query["NoPagination"] = request.NoPagination
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirectories"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoryWithOptions(request *DescribeDirectoryRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirectory"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectory(request *DescribeDirectoryRequest) (_result *DescribeDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoryResponse{}
	_body, _err := client.DescribeDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExternalStreamURLWithOptions(request *DescribeExternalStreamURLRequest, runtime *util.RuntimeOptions) (_result *DescribeExternalStreamURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		query["Kind"] = request.Kind
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Profile)) {
		query["Profile"] = request.Profile
	}

	if !tea.BoolValue(util.IsUnset(request.TxId)) {
		query["TxId"] = request.TxId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExternalStreamURL"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExternalStreamURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExternalStreamURL(request *DescribeExternalStreamURLRequest) (_result *DescribeExternalStreamURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExternalStreamURLResponse{}
	_body, _err := client.DescribeExternalStreamURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupWithOptions(request *DescribeGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeStats)) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroup"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroup(request *DescribeGroupRequest) (_result *DescribeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupResponse{}
	_body, _err := client.DescribeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupsWithOptions(request *DescribeGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InProtocol)) {
		query["InProtocol"] = request.InProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeStats)) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroups"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroups(request *DescribeGroupsRequest) (_result *DescribeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.DescribeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeDevicesInfoWithOptions(request *DescribeNodeDevicesInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeDevicesInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeDevicesInfo"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeDevicesInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeDevicesInfo(request *DescribeNodeDevicesInfoRequest) (_result *DescribeNodeDevicesInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeDevicesInfoResponse{}
	_body, _err := client.DescribeNodeDevicesInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParentPlatformWithOptions(request *DescribeParentPlatformRequest, runtime *util.RuntimeOptions) (_result *DescribeParentPlatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParentPlatform"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParentPlatform(request *DescribeParentPlatformRequest) (_result *DescribeParentPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParentPlatformResponse{}
	_body, _err := client.DescribeParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParentPlatformDevicesWithOptions(request *DescribeParentPlatformDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeParentPlatformDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParentPlatformDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParentPlatformDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParentPlatformDevices(request *DescribeParentPlatformDevicesRequest) (_result *DescribeParentPlatformDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParentPlatformDevicesResponse{}
	_body, _err := client.DescribeParentPlatformDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParentPlatformsWithOptions(request *DescribeParentPlatformsRequest, runtime *util.RuntimeOptions) (_result *DescribeParentPlatformsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParentPlatforms"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParentPlatformsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParentPlatforms(request *DescribeParentPlatformsRequest) (_result *DescribeParentPlatformsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParentPlatformsResponse{}
	_body, _err := client.DescribeParentPlatformsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePresetsWithOptions(request *DescribePresetsRequest, runtime *util.RuntimeOptions) (_result *DescribePresetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePresets"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePresetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePresets(request *DescribePresetsRequest) (_result *DescribePresetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePresetsResponse{}
	_body, _err := client.DescribePresetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedDeviceWithOptions(request *DescribePurchasedDeviceRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedDevice(request *DescribePurchasedDeviceRequest) (_result *DescribePurchasedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedDeviceResponse{}
	_body, _err := client.DescribePurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedDevicesWithOptions(request *DescribePurchasedDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedDevices(request *DescribePurchasedDevicesRequest) (_result *DescribePurchasedDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedDevicesResponse{}
	_body, _err := client.DescribePurchasedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordsWithOptions(request *DescribeRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateBucket)) {
		query["PrivateBucket"] = request.PrivateBucket
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StreamId)) {
		query["StreamId"] = request.StreamId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecords"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecords(request *DescribeRecordsRequest) (_result *DescribeRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordsResponse{}
	_body, _err := client.DescribeRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRenderingDevicesWithOptions(request *DescribeRenderingDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeRenderingDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRenderingDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRenderingDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRenderingDevices(request *DescribeRenderingDevicesRequest) (_result *DescribeRenderingDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRenderingDevicesResponse{}
	_body, _err := client.DescribeRenderingDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamWithOptions(request *DescribeStreamRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStream(request *DescribeStreamRequest) (_result *DescribeStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamResponse{}
	_body, _err := client.DescribeStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamURLWithOptions(request *DescribeStreamURLRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Auth)) {
		query["Auth"] = request.Auth
	}

	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Expire)) {
		query["Expire"] = request.Expire
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OutProtocol)) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Transcode)) {
		query["Transcode"] = request.Transcode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStreamURL"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStreamURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStreamURL(request *DescribeStreamURLRequest) (_result *DescribeStreamURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamURLResponse{}
	_body, _err := client.DescribeStreamURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamVodListWithOptions(request *DescribeStreamVodListRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamVodListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStreamVodList"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStreamVodListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStreamVodList(request *DescribeStreamVodListRequest) (_result *DescribeStreamVodListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamVodListResponse{}
	_body, _err := client.DescribeStreamVodListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamsWithOptions(request *DescribeStreamsRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.App)) {
		query["App"] = request.App
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStreams"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStreams(request *DescribeStreamsRequest) (_result *DescribeStreamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamsResponse{}
	_body, _err := client.DescribeStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateWithOptions(request *DescribeTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplate(request *DescribeTemplateRequest) (_result *DescribeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateResponse{}
	_body, _err := client.DescribeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplates"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserDevicesWithOptions(request *DescribeUserDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeUserDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnsInstanceIds)) {
		query["EnsInstanceIds"] = request.EnsInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerName)) {
		query["ServerName"] = request.ServerName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserDevices(request *DescribeUserDevicesRequest) (_result *DescribeUserDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserDevicesResponse{}
	_body, _err := client.DescribeUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodStreamURLWithOptions(request *DescribeVodStreamURLRequest, runtime *util.RuntimeOptions) (_result *DescribeVodStreamURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TxId)) {
		query["TxId"] = request.TxId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVodStreamURL"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVodStreamURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodStreamURL(request *DescribeVodStreamURLRequest) (_result *DescribeVodStreamURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodStreamURLResponse{}
	_body, _err := client.DescribeVodStreamURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsCertificateDetailWithOptions(request *DescribeVsCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVsCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsCertificateDetail"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsCertificateDetail(request *DescribeVsCertificateDetailRequest) (_result *DescribeVsCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsCertificateDetailResponse{}
	_body, _err := client.DescribeVsCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsCertificateListWithOptions(request *DescribeVsCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeVsCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsCertificateList"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsCertificateList(request *DescribeVsCertificateListRequest) (_result *DescribeVsCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsCertificateListResponse{}
	_body, _err := client.DescribeVsCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDevicesDataWithOptions(request *DescribeVsDevicesDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDevicesDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDevicesData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDevicesDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDevicesData(request *DescribeVsDevicesDataRequest) (_result *DescribeVsDevicesDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDevicesDataResponse{}
	_body, _err := client.DescribeVsDevicesDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainBpsDataWithOptions(request *DescribeVsDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainBpsData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainBpsData(request *DescribeVsDomainBpsDataRequest) (_result *DescribeVsDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainBpsDataResponse{}
	_body, _err := client.DescribeVsDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainCertificateInfoWithOptions(request *DescribeVsDomainCertificateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainCertificateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainCertificateInfo"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainCertificateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainCertificateInfo(request *DescribeVsDomainCertificateInfoRequest) (_result *DescribeVsDomainCertificateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainCertificateInfoResponse{}
	_body, _err := client.DescribeVsDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainConfigsWithOptions(request *DescribeVsDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainConfigs"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainConfigs(request *DescribeVsDomainConfigsRequest) (_result *DescribeVsDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainConfigsResponse{}
	_body, _err := client.DescribeVsDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainDetailWithOptions(request *DescribeVsDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainDetail"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainDetail(request *DescribeVsDomainDetailRequest) (_result *DescribeVsDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainDetailResponse{}
	_body, _err := client.DescribeVsDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainOnlineUserNumWithOptions(request *DescribeVsDomainOnlineUserNumRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainOnlineUserNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryTime)) {
		query["QueryTime"] = request.QueryTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainOnlineUserNum"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainOnlineUserNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainOnlineUserNum(request *DescribeVsDomainOnlineUserNumRequest) (_result *DescribeVsDomainOnlineUserNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainOnlineUserNumResponse{}
	_body, _err := client.DescribeVsDomainOnlineUserNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainPvDataWithOptions(request *DescribeVsDomainPvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainPvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainPvData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainPvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainPvData(request *DescribeVsDomainPvDataRequest) (_result *DescribeVsDomainPvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainPvDataResponse{}
	_body, _err := client.DescribeVsDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainPvUvDataWithOptions(request *DescribeVsDomainPvUvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainPvUvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainPvUvData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainPvUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainPvUvData(request *DescribeVsDomainPvUvDataRequest) (_result *DescribeVsDomainPvUvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainPvUvDataResponse{}
	_body, _err := client.DescribeVsDomainPvUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainRecordDataWithOptions(request *DescribeVsDomainRecordDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainRecordDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainRecordData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainRecordDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainRecordData(request *DescribeVsDomainRecordDataRequest) (_result *DescribeVsDomainRecordDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainRecordDataResponse{}
	_body, _err := client.DescribeVsDomainRecordDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainRegionDataWithOptions(request *DescribeVsDomainRegionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainRegionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainRegionData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainRegionData(request *DescribeVsDomainRegionDataRequest) (_result *DescribeVsDomainRegionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainRegionDataResponse{}
	_body, _err := client.DescribeVsDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainReqBpsDataWithOptions(request *DescribeVsDomainReqBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainReqBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainReqBpsData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainReqBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainReqBpsData(request *DescribeVsDomainReqBpsDataRequest) (_result *DescribeVsDomainReqBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainReqBpsDataResponse{}
	_body, _err := client.DescribeVsDomainReqBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainReqTrafficDataWithOptions(request *DescribeVsDomainReqTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainReqTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainReqTrafficData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainReqTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainReqTrafficData(request *DescribeVsDomainReqTrafficDataRequest) (_result *DescribeVsDomainReqTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainReqTrafficDataResponse{}
	_body, _err := client.DescribeVsDomainReqTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainSnapshotDataWithOptions(request *DescribeVsDomainSnapshotDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainSnapshotDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainSnapshotData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainSnapshotDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainSnapshotData(request *DescribeVsDomainSnapshotDataRequest) (_result *DescribeVsDomainSnapshotDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainSnapshotDataResponse{}
	_body, _err := client.DescribeVsDomainSnapshotDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainTrafficDataWithOptions(request *DescribeVsDomainTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainTrafficData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainTrafficData(request *DescribeVsDomainTrafficDataRequest) (_result *DescribeVsDomainTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainTrafficDataResponse{}
	_body, _err := client.DescribeVsDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsDomainUvDataWithOptions(request *DescribeVsDomainUvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsDomainUvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsDomainUvData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsDomainUvData(request *DescribeVsDomainUvDataRequest) (_result *DescribeVsDomainUvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsDomainUvDataResponse{}
	_body, _err := client.DescribeVsDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsPullStreamConfigWithOptions(request *DescribeVsPullStreamConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeVsPullStreamConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsPullStreamConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsPullStreamConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsPullStreamConfig(request *DescribeVsPullStreamConfigRequest) (_result *DescribeVsPullStreamConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsPullStreamConfigResponse{}
	_body, _err := client.DescribeVsPullStreamConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsPullStreamInfoConfigWithOptions(request *DescribeVsPullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeVsPullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsPullStreamInfoConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsPullStreamInfoConfig(request *DescribeVsPullStreamInfoConfigRequest) (_result *DescribeVsPullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsPullStreamInfoConfigResponse{}
	_body, _err := client.DescribeVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsStorageTrafficUsageDataWithOptions(request *DescribeVsStorageTrafficUsageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsStorageTrafficUsageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SplitBy)) {
		query["SplitBy"] = request.SplitBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsStorageTrafficUsageData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsStorageTrafficUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsStorageTrafficUsageData(request *DescribeVsStorageTrafficUsageDataRequest) (_result *DescribeVsStorageTrafficUsageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsStorageTrafficUsageDataResponse{}
	_body, _err := client.DescribeVsStorageTrafficUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsStorageUsageDataWithOptions(request *DescribeVsStorageUsageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsStorageUsageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SplitBy)) {
		query["SplitBy"] = request.SplitBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsStorageUsageData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsStorageUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsStorageUsageData(request *DescribeVsStorageUsageDataRequest) (_result *DescribeVsStorageUsageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsStorageUsageDataResponse{}
	_body, _err := client.DescribeVsStorageUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsStreamsNotifyUrlConfigWithOptions(request *DescribeVsStreamsNotifyUrlConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeVsStreamsNotifyUrlConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsStreamsNotifyUrlConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsStreamsNotifyUrlConfig(request *DescribeVsStreamsNotifyUrlConfigRequest) (_result *DescribeVsStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DescribeVsStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsStreamsOnlineListWithOptions(request *DescribeVsStreamsOnlineListRequest, runtime *util.RuntimeOptions) (_result *DescribeVsStreamsOnlineListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsStreamsOnlineList"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsStreamsOnlineListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsStreamsOnlineList(request *DescribeVsStreamsOnlineListRequest) (_result *DescribeVsStreamsOnlineListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsStreamsOnlineListResponse{}
	_body, _err := client.DescribeVsStreamsOnlineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsStreamsPublishListWithOptions(request *DescribeVsStreamsPublishListRequest, runtime *util.RuntimeOptions) (_result *DescribeVsStreamsPublishListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsStreamsPublishList"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsStreamsPublishListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsStreamsPublishList(request *DescribeVsStreamsPublishListRequest) (_result *DescribeVsStreamsPublishListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsStreamsPublishListResponse{}
	_body, _err := client.DescribeVsStreamsPublishListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsTopDomainsByFlowWithOptions(request *DescribeVsTopDomainsByFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeVsTopDomainsByFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsTopDomainsByFlow"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsTopDomainsByFlow(request *DescribeVsTopDomainsByFlowRequest) (_result *DescribeVsTopDomainsByFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsTopDomainsByFlowResponse{}
	_body, _err := client.DescribeVsTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsUpPeakPublishStreamDataWithOptions(request *DescribeVsUpPeakPublishStreamDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVsUpPeakPublishStreamDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainSwitch)) {
		query["DomainSwitch"] = request.DomainSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsUpPeakPublishStreamData"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsUpPeakPublishStreamDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsUpPeakPublishStreamData(request *DescribeVsUpPeakPublishStreamDataRequest) (_result *DescribeVsUpPeakPublishStreamDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsUpPeakPublishStreamDataResponse{}
	_body, _err := client.DescribeVsUpPeakPublishStreamDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVsUserResourcePackageWithOptions(request *DescribeVsUserResourcePackageRequest, runtime *util.RuntimeOptions) (_result *DescribeVsUserResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsUserResourcePackage"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVsUserResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVsUserResourcePackage(request *DescribeVsUserResourcePackageRequest) (_result *DescribeVsUserResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVsUserResourcePackageResponse{}
	_body, _err := client.DescribeVsUserResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ForbidVsStreamWithOptions(request *ForbidVsStreamRequest, runtime *util.RuntimeOptions) (_result *ForbidVsStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ControlStreamAction)) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.LiveStreamType)) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !tea.BoolValue(util.IsUnset(request.Oneshot)) {
		query["Oneshot"] = request.Oneshot
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeTime)) {
		query["ResumeTime"] = request.ResumeTime
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ForbidVsStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ForbidVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ForbidVsStream(request *ForbidVsStreamRequest) (_result *ForbidVsStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForbidVsStreamResponse{}
	_body, _err := client.ForbidVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketInfoWithOptions(request *GetBucketInfoRequest, runtime *util.RuntimeOptions) (_result *GetBucketInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketInfo"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBucketInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketInfo(request *GetBucketInfoRequest) (_result *GetBucketInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBucketInfoResponse{}
	_body, _err := client.GetBucketInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetObjectTotalWithOptions(request *GetObjectTotalRequest, runtime *util.RuntimeOptions) (_result *GetObjectTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetObjectTotal"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetObjectTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetObjectTotal(request *GetObjectTotalRequest) (_result *GetObjectTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetObjectTotalResponse{}
	_body, _err := client.GetObjectTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GotoPresetWithOptions(request *GotoPresetRequest, runtime *util.RuntimeOptions) (_result *GotoPresetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PresetId)) {
		query["PresetId"] = request.PresetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GotoPreset"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GotoPresetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GotoPreset(request *GotoPresetRequest) (_result *GotoPresetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GotoPresetResponse{}
	_body, _err := client.GotoPresetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBucketsWithOptions(request *ListBucketsRequest, runtime *util.RuntimeOptions) (_result *ListBucketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBuckets"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBucketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBuckets(request *ListBucketsRequest) (_result *ListBucketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBucketsResponse{}
	_body, _err := client.ListBucketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListObjectsWithOptions(request *ListObjectsRequest, runtime *util.RuntimeOptions) (_result *ListObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.ContinuationToken)) {
		query["ContinuationToken"] = request.ContinuationToken
	}

	if !tea.BoolValue(util.IsUnset(request.Delimiter)) {
		query["Delimiter"] = request.Delimiter
	}

	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["EncodingType"] = request.EncodingType
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["MaxKeys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartAfter)) {
		query["StartAfter"] = request.StartAfter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListObjects"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListObjects(request *ListObjectsRequest) (_result *ListObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListObjectsResponse{}
	_body, _err := client.ListObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceWithOptions(request *ModifyDeviceRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmMethod)) {
		query["AlarmMethod"] = request.AlarmMethod
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPos)) {
		query["AutoPos"] = request.AutoPos
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStart)) {
		query["AutoStart"] = request.AutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["Latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["Longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.PosInterval)) {
		query["PosInterval"] = request.PosInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDevice(request *ModifyDeviceRequest) (_result *ModifyDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceResponse{}
	_body, _err := client.ModifyDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceAlarmWithOptions(request *ModifyDeviceAlarmRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmId)) {
		query["AlarmId"] = request.AlarmId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeviceAlarm"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceAlarm(request *ModifyDeviceAlarmRequest) (_result *ModifyDeviceAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceAlarmResponse{}
	_body, _err := client.ModifyDeviceAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceCaptureWithOptions(request *ModifyDeviceCaptureRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceCaptureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Video)) {
		query["Video"] = request.Video
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeviceCapture"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceCaptureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceCapture(request *ModifyDeviceCaptureRequest) (_result *ModifyDeviceCaptureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceCaptureResponse{}
	_body, _err := client.ModifyDeviceCaptureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceChannelsWithOptions(request *ModifyDeviceChannelsRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channels)) {
		query["Channels"] = request.Channels
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceStatus)) {
		query["DeviceStatus"] = request.DeviceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Dsn)) {
		query["Dsn"] = request.Dsn
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeviceChannels"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceChannels(request *ModifyDeviceChannelsRequest) (_result *ModifyDeviceChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceChannelsResponse{}
	_body, _err := client.ModifyDeviceChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDirectoryWithOptions(request *ModifyDirectoryRequest, runtime *util.RuntimeOptions) (_result *ModifyDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDirectory"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDirectory(request *ModifyDirectoryRequest) (_result *ModifyDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDirectoryResponse{}
	_body, _err := client.ModifyDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGroupWithOptions(request *ModifyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InProtocol)) {
		query["InProtocol"] = request.InProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LazyPull)) {
		query["LazyPull"] = request.LazyPull
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OutProtocol)) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayDomain)) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !tea.BoolValue(util.IsUnset(request.PushDomain)) {
		query["PushDomain"] = request.PushDomain
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGroup"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGroup(request *ModifyGroupRequest) (_result *ModifyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGroupResponse{}
	_body, _err := client.ModifyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyParentPlatformWithOptions(request *ModifyParentPlatformRequest, runtime *util.RuntimeOptions) (_result *ModifyParentPlatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoStart)) {
		query["AutoStart"] = request.AutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.ClientAuth)) {
		query["ClientAuth"] = request.ClientAuth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientPassword)) {
		query["ClientPassword"] = request.ClientPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUsername)) {
		query["ClientUsername"] = request.ClientUsername
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyParentPlatform"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyParentPlatform(request *ModifyParentPlatformRequest) (_result *ModifyParentPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParentPlatformResponse{}
	_body, _err := client.ModifyParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPurchasedDeviceWithOptions(request *ModifyPurchasedDeviceRequest, runtime *util.RuntimeOptions) (_result *ModifyPurchasedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterCode)) {
		query["RegisterCode"] = request.RegisterCode
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPurchasedDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPurchasedDevice(request *ModifyPurchasedDeviceRequest) (_result *ModifyPurchasedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPurchasedDeviceResponse{}
	_body, _err := client.ModifyPurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTemplateWithOptions(request *ModifyTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileFormat)) {
		query["FileFormat"] = request.FileFormat
	}

	if !tea.BoolValue(util.IsUnset(request.Flv)) {
		query["Flv"] = request.Flv
	}

	if !tea.BoolValue(util.IsUnset(request.HlsM3u8)) {
		query["HlsM3u8"] = request.HlsM3u8
	}

	if !tea.BoolValue(util.IsUnset(request.HlsTs)) {
		query["HlsTs"] = request.HlsTs
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.JpgOnDemand)) {
		query["JpgOnDemand"] = request.JpgOnDemand
	}

	if !tea.BoolValue(util.IsUnset(request.JpgOverwrite)) {
		query["JpgOverwrite"] = request.JpgOverwrite
	}

	if !tea.BoolValue(util.IsUnset(request.JpgSequence)) {
		query["JpgSequence"] = request.JpgSequence
	}

	if !tea.BoolValue(util.IsUnset(request.Mp4)) {
		query["Mp4"] = request.Mp4
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.OssFilePrefix)) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Retention)) {
		query["Retention"] = request.Retention
	}

	if !tea.BoolValue(util.IsUnset(request.TransConfigsJSON)) {
		query["TransConfigsJSON"] = request.TransConfigsJSON
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		query["Trigger"] = request.Trigger
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTemplate(request *ModifyTemplateRequest) (_result *ModifyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.ModifyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenVsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenVsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenVsService"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenVsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenVsService() (_result *OpenVsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenVsServiceResponse{}
	_body, _err := client.OpenVsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateRenderingDevicesWithOptions(request *OperateRenderingDevicesRequest, runtime *util.RuntimeOptions) (_result *OperateRenderingDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PodId)) {
		query["PodId"] = request.PodId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateRenderingDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateRenderingDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateRenderingDevices(request *OperateRenderingDevicesRequest) (_result *OperateRenderingDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateRenderingDevicesResponse{}
	_body, _err := client.OperateRenderingDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PrepareUploadWithOptions(request *PrepareUploadRequest, runtime *util.RuntimeOptions) (_result *PrepareUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PrepareUpload"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PrepareUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PrepareUpload(request *PrepareUploadRequest) (_result *PrepareUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PrepareUploadResponse{}
	_body, _err := client.PrepareUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketWithOptions(request *PutBucketRequest, runtime *util.RuntimeOptions) (_result *PutBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketAcl)) {
		query["BucketAcl"] = request.BucketAcl
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.DataRedundancyType)) {
		query["DataRedundancyType"] = request.DataRedundancyType
	}

	if !tea.BoolValue(util.IsUnset(request.DispatcherType)) {
		query["DispatcherType"] = request.DispatcherType
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageClass)) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucket"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucket(request *PutBucketRequest) (_result *PutBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutBucketResponse{}
	_body, _err := client.PutBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetRenderingDevicesWithOptions(request *ResetRenderingDevicesRequest, runtime *util.RuntimeOptions) (_result *ResetRenderingDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PodId)) {
		query["PodId"] = request.PodId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetRenderingDevices"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetRenderingDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetRenderingDevices(request *ResetRenderingDevicesRequest) (_result *ResetRenderingDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetRenderingDevicesResponse{}
	_body, _err := client.ResetRenderingDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeVsStreamWithOptions(request *ResumeVsStreamRequest, runtime *util.RuntimeOptions) (_result *ResumeVsStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ControlStreamAction)) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.LiveStreamType)) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeVsStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeVsStream(request *ResumeVsStreamRequest) (_result *ResumeVsStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeVsStreamResponse{}
	_body, _err := client.ResumeVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetPresetWithOptions(request *SetPresetRequest, runtime *util.RuntimeOptions) (_result *SetPresetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PresetId)) {
		query["PresetId"] = request.PresetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPreset"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPresetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPreset(request *SetPresetRequest) (_result *SetPresetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPresetResponse{}
	_body, _err := client.SetPresetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetVsDomainCertificateWithOptions(request *SetVsDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *SetVsDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ForceSet)) {
		query["ForceSet"] = request.ForceSet
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPri)) {
		query["SSLPri"] = request.SSLPri
	}

	if !tea.BoolValue(util.IsUnset(request.SSLProtocol)) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPub)) {
		query["SSLPub"] = request.SSLPub
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetVsDomainCertificate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetVsDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetVsDomainCertificate(request *SetVsDomainCertificateRequest) (_result *SetVsDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetVsDomainCertificateResponse{}
	_body, _err := client.SetVsDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetVsStreamsNotifyUrlConfigWithOptions(request *SetVsStreamsNotifyUrlConfigRequest, runtime *util.RuntimeOptions) (_result *SetVsStreamsNotifyUrlConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthKey)) {
		query["AuthKey"] = request.AuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyUrl)) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetVsStreamsNotifyUrlConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetVsStreamsNotifyUrlConfig(request *SetVsStreamsNotifyUrlConfigRequest) (_result *SetVsStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.SetVsStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDeviceWithOptions(request *StartDeviceRequest, runtime *util.RuntimeOptions) (_result *StartDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDevice(request *StartDeviceRequest) (_result *StartDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDeviceResponse{}
	_body, _err := client.StartDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartParentPlatformWithOptions(request *StartParentPlatformRequest, runtime *util.RuntimeOptions) (_result *StartParentPlatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartParentPlatform"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartParentPlatform(request *StartParentPlatformRequest) (_result *StartParentPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartParentPlatformResponse{}
	_body, _err := client.StartParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRecordStreamWithOptions(request *StartRecordStreamRequest, runtime *util.RuntimeOptions) (_result *StartRecordStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.App)) {
		query["App"] = request.App
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayDomain)) {
		query["PlayDomain"] = request.PlayDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRecordStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRecordStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRecordStream(request *StartRecordStreamRequest) (_result *StartRecordStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRecordStreamResponse{}
	_body, _err := client.StartRecordStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartStreamWithOptions(request *StartStreamRequest, runtime *util.RuntimeOptions) (_result *StartStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartStream(request *StartStreamRequest) (_result *StartStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartStreamResponse{}
	_body, _err := client.StartStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTransferStreamWithOptions(request *StartTransferStreamRequest, runtime *util.RuntimeOptions) (_result *StartTransferStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Transcode)) {
		query["Transcode"] = request.Transcode
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTransferStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTransferStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTransferStream(request *StartTransferStreamRequest) (_result *StartTransferStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTransferStreamResponse{}
	_body, _err := client.StartTransferStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAdjustWithOptions(request *StopAdjustRequest, runtime *util.RuntimeOptions) (_result *StopAdjustResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Focus)) {
		query["Focus"] = request.Focus
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Iris)) {
		query["Iris"] = request.Iris
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAdjust"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAdjustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAdjust(request *StopAdjustRequest) (_result *StopAdjustResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAdjustResponse{}
	_body, _err := client.StopAdjustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDeviceWithOptions(request *StopDeviceRequest, runtime *util.RuntimeOptions) (_result *StopDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDevice(request *StopDeviceRequest) (_result *StopDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDeviceResponse{}
	_body, _err := client.StopDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMoveWithOptions(request *StopMoveRequest, runtime *util.RuntimeOptions) (_result *StopMoveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Pan)) {
		query["Pan"] = request.Pan
	}

	if !tea.BoolValue(util.IsUnset(request.Tilt)) {
		query["Tilt"] = request.Tilt
	}

	if !tea.BoolValue(util.IsUnset(request.Zoom)) {
		query["Zoom"] = request.Zoom
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMove"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMoveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMove(request *StopMoveRequest) (_result *StopMoveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopMoveResponse{}
	_body, _err := client.StopMoveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopParentPlatformWithOptions(request *StopParentPlatformRequest, runtime *util.RuntimeOptions) (_result *StopParentPlatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopParentPlatform"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopParentPlatform(request *StopParentPlatformRequest) (_result *StopParentPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopParentPlatformResponse{}
	_body, _err := client.StopParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopRecordStreamWithOptions(request *StopRecordStreamRequest, runtime *util.RuntimeOptions) (_result *StopRecordStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.App)) {
		query["App"] = request.App
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayDomain)) {
		query["PlayDomain"] = request.PlayDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopRecordStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopRecordStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopRecordStream(request *StopRecordStreamRequest) (_result *StopRecordStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopRecordStreamResponse{}
	_body, _err := client.StopRecordStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStreamWithOptions(request *StopStreamRequest, runtime *util.RuntimeOptions) (_result *StopStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStream(request *StopStreamRequest) (_result *StopStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopStreamResponse{}
	_body, _err := client.StopStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTransferStreamWithOptions(request *StopTransferStreamRequest, runtime *util.RuntimeOptions) (_result *StopTransferStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Transcode)) {
		query["Transcode"] = request.Transcode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopTransferStream"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTransferStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTransferStream(request *StopTransferStreamRequest) (_result *StopTransferStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopTransferStreamResponse{}
	_body, _err := client.StopTransferStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncCatalogsWithOptions(request *SyncCatalogsRequest, runtime *util.RuntimeOptions) (_result *SyncCatalogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncCatalogs"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncCatalogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncCatalogs(request *SyncCatalogsRequest) (_result *SyncCatalogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncCatalogsResponse{}
	_body, _err := client.SyncCatalogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDirectoryWithOptions(request *UnbindDirectoryRequest, runtime *util.RuntimeOptions) (_result *UnbindDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDirectory"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDirectory(request *UnbindDirectoryRequest) (_result *UnbindDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDirectoryResponse{}
	_body, _err := client.UnbindDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindParentPlatformDeviceWithOptions(request *UnbindParentPlatformDeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindParentPlatformDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentPlatformId)) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindParentPlatformDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindParentPlatformDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindParentPlatformDevice(request *UnbindParentPlatformDeviceRequest) (_result *UnbindParentPlatformDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindParentPlatformDeviceResponse{}
	_body, _err := client.UnbindParentPlatformDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindPurchasedDeviceWithOptions(request *UnbindPurchasedDeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindPurchasedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindPurchasedDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindPurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindPurchasedDevice(request *UnbindPurchasedDeviceRequest) (_result *UnbindPurchasedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindPurchasedDeviceResponse{}
	_body, _err := client.UnbindPurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindTemplateWithOptions(request *UnbindTemplateRequest, runtime *util.RuntimeOptions) (_result *UnbindTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindTemplate"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindTemplate(request *UnbindTemplateRequest) (_result *UnbindTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindTemplateResponse{}
	_body, _err := client.UnbindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockDeviceWithOptions(request *UnlockDeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockDevice"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockDevice(request *UnlockDeviceRequest) (_result *UnlockDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockDeviceResponse{}
	_body, _err := client.UnlockDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAIConfigWithOptions(request *UpdateAIConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateAIConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaptureInterval)) {
		query["CaptureInterval"] = request.CaptureInterval
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.Configs)) {
		query["Configs"] = request.Configs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Features)) {
		query["Features"] = request.Features
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAIConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAIConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAIConfig(request *UpdateAIConfigRequest) (_result *UpdateAIConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAIConfigResponse{}
	_body, _err := client.UpdateAIConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBucketInfoWithOptions(request *UpdateBucketInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateBucketInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBucketInfo"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBucketInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBucketInfo(request *UpdateBucketInfoRequest) (_result *UpdateBucketInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBucketInfoResponse{}
	_body, _err := client.UpdateBucketInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterWithOptions(request *UpdateClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.InternalPorts)) {
		query["InternalPorts"] = request.InternalPorts
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainTime)) {
		query["MaintainTime"] = request.MaintainTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCluster"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRenderingDeviceSpecWithOptions(request *UpdateRenderingDeviceSpecRequest, runtime *util.RuntimeOptions) (_result *UpdateRenderingDeviceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRenderingDeviceSpec"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRenderingDeviceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRenderingDeviceSpec(request *UpdateRenderingDeviceSpecRequest) (_result *UpdateRenderingDeviceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRenderingDeviceSpecResponse{}
	_body, _err := client.UpdateRenderingDeviceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVsPullStreamInfoConfigWithOptions(request *UpdateVsPullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateVsPullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Always)) {
		query["Always"] = request.Always
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		query["SourceUrl"] = request.SourceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		query["StreamName"] = request.StreamName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVsPullStreamInfoConfig"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVsPullStreamInfoConfig(request *UpdateVsPullStreamInfoConfigRequest) (_result *UpdateVsPullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVsPullStreamInfoConfigResponse{}
	_body, _err := client.UpdateVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeRenderingDevicesHostOSWithOptions(request *UpgradeRenderingDevicesHostOSRequest, runtime *util.RuntimeOptions) (_result *UpgradeRenderingDevicesHostOSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RomName)) {
		query["RomName"] = request.RomName
	}

	if !tea.BoolValue(util.IsUnset(request.RomPath)) {
		query["RomPath"] = request.RomPath
	}

	if !tea.BoolValue(util.IsUnset(request.RomVersion)) {
		query["RomVersion"] = request.RomVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeRenderingDevicesHostOS"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeRenderingDevicesHostOSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeRenderingDevicesHostOS(request *UpgradeRenderingDevicesHostOSRequest) (_result *UpgradeRenderingDevicesHostOSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeRenderingDevicesHostOSResponse{}
	_body, _err := client.UpgradeRenderingDevicesHostOSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeRenderingDevicesImageWithOptions(request *UpgradeRenderingDevicesImageRequest, runtime *util.RuntimeOptions) (_result *UpgradeRenderingDevicesImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeRenderingDevicesImage"),
		Version:     tea.String("2018-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeRenderingDevicesImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeRenderingDevicesImage(request *UpgradeRenderingDevicesImageRequest) (_result *UpgradeRenderingDevicesImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeRenderingDevicesImageResponse{}
	_body, _err := client.UpgradeRenderingDevicesImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
