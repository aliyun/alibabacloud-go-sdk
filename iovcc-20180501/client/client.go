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

type AddUploadedFunctionFileInfoRequest struct {
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddUploadedFunctionFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUploadedFunctionFileInfoRequest) GoString() string {
	return s.String()
}

func (s *AddUploadedFunctionFileInfoRequest) SetFileName(v string) *AddUploadedFunctionFileInfoRequest {
	s.FileName = &v
	return s
}

func (s *AddUploadedFunctionFileInfoRequest) SetObjectKey(v string) *AddUploadedFunctionFileInfoRequest {
	s.ObjectKey = &v
	return s
}

func (s *AddUploadedFunctionFileInfoRequest) SetProjectId(v string) *AddUploadedFunctionFileInfoRequest {
	s.ProjectId = &v
	return s
}

type AddUploadedFunctionFileInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUploadedFunctionFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUploadedFunctionFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddUploadedFunctionFileInfoResponseBody) SetRequestId(v string) *AddUploadedFunctionFileInfoResponseBody {
	s.RequestId = &v
	return s
}

type AddUploadedFunctionFileInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUploadedFunctionFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUploadedFunctionFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUploadedFunctionFileInfoResponse) GoString() string {
	return s.String()
}

func (s *AddUploadedFunctionFileInfoResponse) SetHeaders(v map[string]*string) *AddUploadedFunctionFileInfoResponse {
	s.Headers = v
	return s
}

func (s *AddUploadedFunctionFileInfoResponse) SetStatusCode(v int32) *AddUploadedFunctionFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUploadedFunctionFileInfoResponse) SetBody(v *AddUploadedFunctionFileInfoResponseBody) *AddUploadedFunctionFileInfoResponse {
	s.Body = v
	return s
}

type AddVersionBlackDevicesRequest struct {
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s AddVersionBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionBlackDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddVersionBlackDevicesRequest) SetDeviceIdType(v string) *AddVersionBlackDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetDeviceIds(v string) *AddVersionBlackDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetProjectId(v string) *AddVersionBlackDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetVersionId(v string) *AddVersionBlackDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetVersionType(v string) *AddVersionBlackDevicesRequest {
	s.VersionType = &v
	return s
}

type AddVersionBlackDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVersionBlackDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVersionBlackDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *AddVersionBlackDevicesResponseBody) SetRequestId(v string) *AddVersionBlackDevicesResponseBody {
	s.RequestId = &v
	return s
}

type AddVersionBlackDevicesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddVersionBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVersionBlackDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVersionBlackDevicesResponse) GoString() string {
	return s.String()
}

func (s *AddVersionBlackDevicesResponse) SetHeaders(v map[string]*string) *AddVersionBlackDevicesResponse {
	s.Headers = v
	return s
}

func (s *AddVersionBlackDevicesResponse) SetStatusCode(v int32) *AddVersionBlackDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVersionBlackDevicesResponse) SetBody(v *AddVersionBlackDevicesResponseBody) *AddVersionBlackDevicesResponse {
	s.Body = v
	return s
}

type AddVersionGroupDevicesRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceIdType  *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceIds     *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddVersionGroupDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionGroupDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddVersionGroupDevicesRequest) SetDeviceGroupId(v string) *AddVersionGroupDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddVersionGroupDevicesRequest) SetDeviceIdType(v string) *AddVersionGroupDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *AddVersionGroupDevicesRequest) SetDeviceIds(v string) *AddVersionGroupDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *AddVersionGroupDevicesRequest) SetProjectId(v string) *AddVersionGroupDevicesRequest {
	s.ProjectId = &v
	return s
}

type AddVersionGroupDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVersionGroupDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVersionGroupDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *AddVersionGroupDevicesResponseBody) SetRequestId(v string) *AddVersionGroupDevicesResponseBody {
	s.RequestId = &v
	return s
}

type AddVersionGroupDevicesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddVersionGroupDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVersionGroupDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVersionGroupDevicesResponse) GoString() string {
	return s.String()
}

func (s *AddVersionGroupDevicesResponse) SetHeaders(v map[string]*string) *AddVersionGroupDevicesResponse {
	s.Headers = v
	return s
}

func (s *AddVersionGroupDevicesResponse) SetStatusCode(v int32) *AddVersionGroupDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVersionGroupDevicesResponse) SetBody(v *AddVersionGroupDevicesResponseBody) *AddVersionGroupDevicesResponse {
	s.Body = v
	return s
}

type AddVersionWhiteDevicesRequest struct {
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s AddVersionWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesRequest) SetDeviceIdType(v string) *AddVersionWhiteDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetDeviceIds(v string) *AddVersionWhiteDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetProjectId(v string) *AddVersionWhiteDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetVersionId(v string) *AddVersionWhiteDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetVersionType(v string) *AddVersionWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

type AddVersionWhiteDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVersionWhiteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesResponseBody) SetRequestId(v string) *AddVersionWhiteDevicesResponseBody {
	s.RequestId = &v
	return s
}

type AddVersionWhiteDevicesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddVersionWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVersionWhiteDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesResponse) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesResponse) SetHeaders(v map[string]*string) *AddVersionWhiteDevicesResponse {
	s.Headers = v
	return s
}

func (s *AddVersionWhiteDevicesResponse) SetStatusCode(v int32) *AddVersionWhiteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVersionWhiteDevicesResponse) SetBody(v *AddVersionWhiteDevicesResponseBody) *AddVersionWhiteDevicesResponse {
	s.Body = v
	return s
}

type AddVersionWhiteDevicesByDeviceGroupsRequest struct {
	GroupIds    *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s AddVersionWhiteDevicesByDeviceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesByDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetGroupIds(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetProjectId(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetVersionId(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.VersionId = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetVersionType(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.VersionType = &v
	return s
}

type AddVersionWhiteDevicesByDeviceGroupsResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVersionWhiteDevicesByDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesByDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponseBody) SetData(v string) *AddVersionWhiteDevicesByDeviceGroupsResponseBody {
	s.Data = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponseBody) SetRequestId(v string) *AddVersionWhiteDevicesByDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type AddVersionWhiteDevicesByDeviceGroupsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddVersionWhiteDevicesByDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVersionWhiteDevicesByDeviceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesByDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponse) SetHeaders(v map[string]*string) *AddVersionWhiteDevicesByDeviceGroupsResponse {
	s.Headers = v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponse) SetStatusCode(v int32) *AddVersionWhiteDevicesByDeviceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponse) SetBody(v *AddVersionWhiteDevicesByDeviceGroupsResponseBody) *AddVersionWhiteDevicesByDeviceGroupsResponse {
	s.Body = v
	return s
}

type ConnectAssistDeviceRequest struct {
	AllowCommandExtension *bool   `json:"AllowCommandExtension,omitempty" xml:"AllowCommandExtension,omitempty"`
	DeviceId              *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	HardwareId            *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SerialNumber          *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	UUID                  *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VIN                   *string `json:"VIN,omitempty" xml:"VIN,omitempty"`
}

func (s ConnectAssistDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConnectAssistDeviceRequest) GoString() string {
	return s.String()
}

func (s *ConnectAssistDeviceRequest) SetAllowCommandExtension(v bool) *ConnectAssistDeviceRequest {
	s.AllowCommandExtension = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetDeviceId(v string) *ConnectAssistDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetHardwareId(v string) *ConnectAssistDeviceRequest {
	s.HardwareId = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetProjectId(v string) *ConnectAssistDeviceRequest {
	s.ProjectId = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetSerialNumber(v string) *ConnectAssistDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetUUID(v string) *ConnectAssistDeviceRequest {
	s.UUID = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetVIN(v string) *ConnectAssistDeviceRequest {
	s.VIN = &v
	return s
}

type ConnectAssistDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConnectAssistDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConnectAssistDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectAssistDeviceResponseBody) SetRequestId(v string) *ConnectAssistDeviceResponseBody {
	s.RequestId = &v
	return s
}

type ConnectAssistDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConnectAssistDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConnectAssistDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConnectAssistDeviceResponse) GoString() string {
	return s.String()
}

func (s *ConnectAssistDeviceResponse) SetHeaders(v map[string]*string) *ConnectAssistDeviceResponse {
	s.Headers = v
	return s
}

func (s *ConnectAssistDeviceResponse) SetStatusCode(v int32) *ConnectAssistDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectAssistDeviceResponse) SetBody(v *ConnectAssistDeviceResponseBody) *ConnectAssistDeviceResponse {
	s.Body = v
	return s
}

type CountActivatedOrNewRegistrationDeviceRequest struct {
	DeviceBrand                  *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId                *string `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	DeviceCountStatType          *string `json:"DeviceCountStatType,omitempty" xml:"DeviceCountStatType,omitempty"`
	DeviceModel                  *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId                *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceType                   *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EndTime                      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsQueryNewRegistrationDevice *string `json:"IsQueryNewRegistrationDevice,omitempty" xml:"IsQueryNewRegistrationDevice,omitempty"`
	IsQueryYearlyActivate        *string `json:"IsQueryYearlyActivate,omitempty" xml:"IsQueryYearlyActivate,omitempty"`
	ProjectId                    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartTime                    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CountActivatedOrNewRegistrationDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CountActivatedOrNewRegistrationDeviceRequest) GoString() string {
	return s.String()
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetDeviceBrand(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.DeviceBrand = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetDeviceBrandId(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetDeviceCountStatType(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.DeviceCountStatType = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetDeviceModel(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.DeviceModel = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetDeviceModelId(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetDeviceType(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetEndTime(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.EndTime = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetIsQueryNewRegistrationDevice(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.IsQueryNewRegistrationDevice = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetIsQueryYearlyActivate(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.IsQueryYearlyActivate = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetProjectId(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.ProjectId = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceRequest) SetStartTime(v string) *CountActivatedOrNewRegistrationDeviceRequest {
	s.StartTime = &v
	return s
}

type CountActivatedOrNewRegistrationDeviceResponseBody struct {
	RequestId  *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics *CountActivatedOrNewRegistrationDeviceResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
}

func (s CountActivatedOrNewRegistrationDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountActivatedOrNewRegistrationDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CountActivatedOrNewRegistrationDeviceResponseBody) SetRequestId(v string) *CountActivatedOrNewRegistrationDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceResponseBody) SetStatistics(v *CountActivatedOrNewRegistrationDeviceResponseBodyStatistics) *CountActivatedOrNewRegistrationDeviceResponseBody {
	s.Statistics = v
	return s
}

type CountActivatedOrNewRegistrationDeviceResponseBodyStatistics struct {
	Categories []*string                                                            `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Series     []*CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s CountActivatedOrNewRegistrationDeviceResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s CountActivatedOrNewRegistrationDeviceResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *CountActivatedOrNewRegistrationDeviceResponseBodyStatistics) SetCategories(v []*string) *CountActivatedOrNewRegistrationDeviceResponseBodyStatistics {
	s.Categories = v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceResponseBodyStatistics) SetSeries(v []*CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries) *CountActivatedOrNewRegistrationDeviceResponseBodyStatistics {
	s.Series = v
	return s
}

type CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Name *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries) String() string {
	return tea.Prettify(s)
}

func (s CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries) GoString() string {
	return s.String()
}

func (s *CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries) SetData(v []*string) *CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries {
	s.Data = v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries) SetName(v string) *CountActivatedOrNewRegistrationDeviceResponseBodyStatisticsSeries {
	s.Name = &v
	return s
}

type CountActivatedOrNewRegistrationDeviceResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountActivatedOrNewRegistrationDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountActivatedOrNewRegistrationDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CountActivatedOrNewRegistrationDeviceResponse) GoString() string {
	return s.String()
}

func (s *CountActivatedOrNewRegistrationDeviceResponse) SetHeaders(v map[string]*string) *CountActivatedOrNewRegistrationDeviceResponse {
	s.Headers = v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceResponse) SetStatusCode(v int32) *CountActivatedOrNewRegistrationDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CountActivatedOrNewRegistrationDeviceResponse) SetBody(v *CountActivatedOrNewRegistrationDeviceResponseBody) *CountActivatedOrNewRegistrationDeviceResponse {
	s.Body = v
	return s
}

type CountDeviceBrandsRequest struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CountDeviceBrandsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceBrandsRequest) GoString() string {
	return s.String()
}

func (s *CountDeviceBrandsRequest) SetDeviceBrand(v string) *CountDeviceBrandsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *CountDeviceBrandsRequest) SetDeviceBrandId(v int64) *CountDeviceBrandsRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *CountDeviceBrandsRequest) SetProjectId(v string) *CountDeviceBrandsRequest {
	s.ProjectId = &v
	return s
}

type CountDeviceBrandsResponseBody struct {
	BrandCount *int32  `json:"BrandCount,omitempty" xml:"BrandCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CountDeviceBrandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceBrandsResponseBody) GoString() string {
	return s.String()
}

func (s *CountDeviceBrandsResponseBody) SetBrandCount(v int32) *CountDeviceBrandsResponseBody {
	s.BrandCount = &v
	return s
}

func (s *CountDeviceBrandsResponseBody) SetRequestId(v string) *CountDeviceBrandsResponseBody {
	s.RequestId = &v
	return s
}

type CountDeviceBrandsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountDeviceBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountDeviceBrandsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceBrandsResponse) GoString() string {
	return s.String()
}

func (s *CountDeviceBrandsResponse) SetHeaders(v map[string]*string) *CountDeviceBrandsResponse {
	s.Headers = v
	return s
}

func (s *CountDeviceBrandsResponse) SetStatusCode(v int32) *CountDeviceBrandsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountDeviceBrandsResponse) SetBody(v *CountDeviceBrandsResponseBody) *CountDeviceBrandsResponse {
	s.Body = v
	return s
}

type CountDeviceModelsRequest struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CountDeviceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceModelsRequest) GoString() string {
	return s.String()
}

func (s *CountDeviceModelsRequest) SetDeviceBrand(v string) *CountDeviceModelsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *CountDeviceModelsRequest) SetDeviceModel(v string) *CountDeviceModelsRequest {
	s.DeviceModel = &v
	return s
}

func (s *CountDeviceModelsRequest) SetDeviceModelId(v int32) *CountDeviceModelsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CountDeviceModelsRequest) SetProjectId(v string) *CountDeviceModelsRequest {
	s.ProjectId = &v
	return s
}

type CountDeviceModelsResponseBody struct {
	ModelCount *int32  `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CountDeviceModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *CountDeviceModelsResponseBody) SetModelCount(v int32) *CountDeviceModelsResponseBody {
	s.ModelCount = &v
	return s
}

func (s *CountDeviceModelsResponseBody) SetRequestId(v string) *CountDeviceModelsResponseBody {
	s.RequestId = &v
	return s
}

type CountDeviceModelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountDeviceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountDeviceModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceModelsResponse) GoString() string {
	return s.String()
}

func (s *CountDeviceModelsResponse) SetHeaders(v map[string]*string) *CountDeviceModelsResponse {
	s.Headers = v
	return s
}

func (s *CountDeviceModelsResponse) SetStatusCode(v int32) *CountDeviceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountDeviceModelsResponse) SetBody(v *CountDeviceModelsResponseBody) *CountDeviceModelsResponse {
	s.Body = v
	return s
}

type CountDevicesRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CountDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s CountDevicesRequest) GoString() string {
	return s.String()
}

func (s *CountDevicesRequest) SetDeviceModel(v string) *CountDevicesRequest {
	s.DeviceModel = &v
	return s
}

func (s *CountDevicesRequest) SetDeviceModelId(v int32) *CountDevicesRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CountDevicesRequest) SetProjectId(v string) *CountDevicesRequest {
	s.ProjectId = &v
	return s
}

type CountDevicesResponseBody struct {
	DeviceCount *int32  `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CountDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *CountDevicesResponseBody) SetDeviceCount(v int32) *CountDevicesResponseBody {
	s.DeviceCount = &v
	return s
}

func (s *CountDevicesResponseBody) SetRequestId(v string) *CountDevicesResponseBody {
	s.RequestId = &v
	return s
}

type CountDevicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s CountDevicesResponse) GoString() string {
	return s.String()
}

func (s *CountDevicesResponse) SetHeaders(v map[string]*string) *CountDevicesResponse {
	s.Headers = v
	return s
}

func (s *CountDevicesResponse) SetStatusCode(v int32) *CountDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *CountDevicesResponse) SetBody(v *CountDevicesResponseBody) *CountDevicesResponse {
	s.Body = v
	return s
}

type CountYunIdInfoResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	YunIdInfo []*CountYunIdInfoResponseBodyYunIdInfo `json:"YunIdInfo,omitempty" xml:"YunIdInfo,omitempty" type:"Repeated"`
}

func (s CountYunIdInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountYunIdInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CountYunIdInfoResponseBody) SetRequestId(v string) *CountYunIdInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountYunIdInfoResponseBody) SetYunIdInfo(v []*CountYunIdInfoResponseBodyYunIdInfo) *CountYunIdInfoResponseBody {
	s.YunIdInfo = v
	return s
}

type CountYunIdInfoResponseBodyYunIdInfo struct {
	TotalBrandCount       *int64 `json:"TotalBrandCount,omitempty" xml:"TotalBrandCount,omitempty"`
	TotalDeviceCount      *int64 `json:"TotalDeviceCount,omitempty" xml:"TotalDeviceCount,omitempty"`
	TotalDeviceModelCount *int64 `json:"TotalDeviceModelCount,omitempty" xml:"TotalDeviceModelCount,omitempty"`
}

func (s CountYunIdInfoResponseBodyYunIdInfo) String() string {
	return tea.Prettify(s)
}

func (s CountYunIdInfoResponseBodyYunIdInfo) GoString() string {
	return s.String()
}

func (s *CountYunIdInfoResponseBodyYunIdInfo) SetTotalBrandCount(v int64) *CountYunIdInfoResponseBodyYunIdInfo {
	s.TotalBrandCount = &v
	return s
}

func (s *CountYunIdInfoResponseBodyYunIdInfo) SetTotalDeviceCount(v int64) *CountYunIdInfoResponseBodyYunIdInfo {
	s.TotalDeviceCount = &v
	return s
}

func (s *CountYunIdInfoResponseBodyYunIdInfo) SetTotalDeviceModelCount(v int64) *CountYunIdInfoResponseBodyYunIdInfo {
	s.TotalDeviceModelCount = &v
	return s
}

type CountYunIdInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountYunIdInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountYunIdInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CountYunIdInfoResponse) GoString() string {
	return s.String()
}

func (s *CountYunIdInfoResponse) SetHeaders(v map[string]*string) *CountYunIdInfoResponse {
	s.Headers = v
	return s
}

func (s *CountYunIdInfoResponse) SetStatusCode(v int32) *CountYunIdInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CountYunIdInfoResponse) SetBody(v *CountYunIdInfoResponseBody) *CountYunIdInfoResponse {
	s.Body = v
	return s
}

type CreateAppVersionRequest struct {
	ApkMd5            *string `json:"ApkMd5,omitempty" xml:"ApkMd5,omitempty"`
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion        *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BlackVersionList  *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	DeviceAdapterList *string `json:"DeviceAdapterList,omitempty" xml:"DeviceAdapterList,omitempty"`
	InstallType       *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	IsAllowNewInstall *string `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	IsForceUpgrade    *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsNeedRestart     *string `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	IsSilentUpgrade   *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	PackageUrl        *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReleaseNote       *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RestartAppParam   *string `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	RestartAppType    *string `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	RestartType       *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	WhiteVersionList  *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
}

func (s CreateAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppVersionRequest) SetApkMd5(v string) *CreateAppVersionRequest {
	s.ApkMd5 = &v
	return s
}

func (s *CreateAppVersionRequest) SetAppId(v string) *CreateAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppVersionRequest) SetAppVersion(v string) *CreateAppVersionRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppVersionRequest) SetBlackVersionList(v string) *CreateAppVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *CreateAppVersionRequest) SetDeviceAdapterList(v string) *CreateAppVersionRequest {
	s.DeviceAdapterList = &v
	return s
}

func (s *CreateAppVersionRequest) SetInstallType(v string) *CreateAppVersionRequest {
	s.InstallType = &v
	return s
}

func (s *CreateAppVersionRequest) SetIsAllowNewInstall(v string) *CreateAppVersionRequest {
	s.IsAllowNewInstall = &v
	return s
}

func (s *CreateAppVersionRequest) SetIsForceUpgrade(v string) *CreateAppVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *CreateAppVersionRequest) SetIsNeedRestart(v string) *CreateAppVersionRequest {
	s.IsNeedRestart = &v
	return s
}

func (s *CreateAppVersionRequest) SetIsSilentUpgrade(v string) *CreateAppVersionRequest {
	s.IsSilentUpgrade = &v
	return s
}

func (s *CreateAppVersionRequest) SetPackageUrl(v string) *CreateAppVersionRequest {
	s.PackageUrl = &v
	return s
}

func (s *CreateAppVersionRequest) SetProjectId(v string) *CreateAppVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAppVersionRequest) SetReleaseNote(v string) *CreateAppVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *CreateAppVersionRequest) SetRemark(v string) *CreateAppVersionRequest {
	s.Remark = &v
	return s
}

func (s *CreateAppVersionRequest) SetRestartAppParam(v string) *CreateAppVersionRequest {
	s.RestartAppParam = &v
	return s
}

func (s *CreateAppVersionRequest) SetRestartAppType(v string) *CreateAppVersionRequest {
	s.RestartAppType = &v
	return s
}

func (s *CreateAppVersionRequest) SetRestartType(v string) *CreateAppVersionRequest {
	s.RestartType = &v
	return s
}

func (s *CreateAppVersionRequest) SetVersionCode(v string) *CreateAppVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *CreateAppVersionRequest) SetWhiteVersionList(v string) *CreateAppVersionRequest {
	s.WhiteVersionList = &v
	return s
}

type CreateAppVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppVersionResponseBody) SetRequestId(v string) *CreateAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetVersionId(v string) *CreateAppVersionResponseBody {
	s.VersionId = &v
	return s
}

type CreateAppVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppVersionResponse) SetHeaders(v map[string]*string) *CreateAppVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppVersionResponse) SetStatusCode(v int32) *CreateAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppVersionResponse) SetBody(v *CreateAppVersionResponseBody) *CreateAppVersionResponse {
	s.Body = v
	return s
}

type CreateCustomizedFilterRequest struct {
	BlackWhiteType   *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueCompareType *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	ValueType        *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType      *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedFilterRequest) SetBlackWhiteType(v string) *CreateCustomizedFilterRequest {
	s.BlackWhiteType = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetName(v string) *CreateCustomizedFilterRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetProjectId(v string) *CreateCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetValue(v string) *CreateCustomizedFilterRequest {
	s.Value = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetValueCompareType(v string) *CreateCustomizedFilterRequest {
	s.ValueCompareType = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetValueType(v string) *CreateCustomizedFilterRequest {
	s.ValueType = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetVersionId(v string) *CreateCustomizedFilterRequest {
	s.VersionId = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetVersionType(v string) *CreateCustomizedFilterRequest {
	s.VersionType = &v
	return s
}

type CreateCustomizedFilterResponseBody struct {
	CustomizedFilterId *string `json:"CustomizedFilterId,omitempty" xml:"CustomizedFilterId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomizedFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedFilterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedFilterResponseBody) SetCustomizedFilterId(v string) *CreateCustomizedFilterResponseBody {
	s.CustomizedFilterId = &v
	return s
}

func (s *CreateCustomizedFilterResponseBody) SetRequestId(v string) *CreateCustomizedFilterResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomizedFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomizedFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedFilterResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedFilterResponse) SetHeaders(v map[string]*string) *CreateCustomizedFilterResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedFilterResponse) SetStatusCode(v int32) *CreateCustomizedFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedFilterResponse) SetBody(v *CreateCustomizedFilterResponseBody) *CreateCustomizedFilterResponse {
	s.Body = v
	return s
}

type CreateCustomizedPropertyRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateCustomizedPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedPropertyRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedPropertyRequest) SetName(v string) *CreateCustomizedPropertyRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetProjectId(v string) *CreateCustomizedPropertyRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetValue(v string) *CreateCustomizedPropertyRequest {
	s.Value = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetVersionId(v string) *CreateCustomizedPropertyRequest {
	s.VersionId = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetVersionType(v string) *CreateCustomizedPropertyRequest {
	s.VersionType = &v
	return s
}

type CreateCustomizedPropertyResponseBody struct {
	CustomizedPropertyId *string `json:"CustomizedPropertyId,omitempty" xml:"CustomizedPropertyId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomizedPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedPropertyResponseBody) SetCustomizedPropertyId(v string) *CreateCustomizedPropertyResponseBody {
	s.CustomizedPropertyId = &v
	return s
}

func (s *CreateCustomizedPropertyResponseBody) SetRequestId(v string) *CreateCustomizedPropertyResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomizedPropertyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomizedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomizedPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedPropertyResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedPropertyResponse) SetHeaders(v map[string]*string) *CreateCustomizedPropertyResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedPropertyResponse) SetStatusCode(v int32) *CreateCustomizedPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedPropertyResponse) SetBody(v *CreateCustomizedPropertyResponseBody) *CreateCustomizedPropertyResponse {
	s.Body = v
	return s
}

type CreateDeviceBrandRequest struct {
	BrandName   *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Manufacture *string `json:"Manufacture,omitempty" xml:"Manufacture,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateDeviceBrandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceBrandRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceBrandRequest) SetBrandName(v string) *CreateDeviceBrandRequest {
	s.BrandName = &v
	return s
}

func (s *CreateDeviceBrandRequest) SetDescription(v string) *CreateDeviceBrandRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceBrandRequest) SetManufacture(v string) *CreateDeviceBrandRequest {
	s.Manufacture = &v
	return s
}

func (s *CreateDeviceBrandRequest) SetProjectId(v string) *CreateDeviceBrandRequest {
	s.ProjectId = &v
	return s
}

type CreateDeviceBrandResponseBody struct {
	BrandId   *int64  `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeviceBrandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceBrandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceBrandResponseBody) SetBrandId(v int64) *CreateDeviceBrandResponseBody {
	s.BrandId = &v
	return s
}

func (s *CreateDeviceBrandResponseBody) SetRequestId(v string) *CreateDeviceBrandResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeviceBrandResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeviceBrandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceBrandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceBrandResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceBrandResponse) SetHeaders(v map[string]*string) *CreateDeviceBrandResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceBrandResponse) SetStatusCode(v int32) *CreateDeviceBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceBrandResponse) SetBody(v *CreateDeviceBrandResponseBody) *CreateDeviceBrandResponse {
	s.Body = v
	return s
}

type CreateDeviceModelRequest struct {
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	CanCreateDeviceId *string `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	InitUsageType     *string `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	ModelName         *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
}

func (s CreateDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceModelRequest) SetBrandName(v string) *CreateDeviceModelRequest {
	s.BrandName = &v
	return s
}

func (s *CreateDeviceModelRequest) SetCanCreateDeviceId(v string) *CreateDeviceModelRequest {
	s.CanCreateDeviceId = &v
	return s
}

func (s *CreateDeviceModelRequest) SetDescription(v string) *CreateDeviceModelRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceModelRequest) SetDeviceName(v string) *CreateDeviceModelRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateDeviceModelRequest) SetDeviceType(v string) *CreateDeviceModelRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateDeviceModelRequest) SetHardwareType(v string) *CreateDeviceModelRequest {
	s.HardwareType = &v
	return s
}

func (s *CreateDeviceModelRequest) SetInitUsageType(v string) *CreateDeviceModelRequest {
	s.InitUsageType = &v
	return s
}

func (s *CreateDeviceModelRequest) SetModelName(v string) *CreateDeviceModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateDeviceModelRequest) SetObjectKey(v string) *CreateDeviceModelRequest {
	s.ObjectKey = &v
	return s
}

func (s *CreateDeviceModelRequest) SetOsPlatform(v string) *CreateDeviceModelRequest {
	s.OsPlatform = &v
	return s
}

func (s *CreateDeviceModelRequest) SetProjectId(v string) *CreateDeviceModelRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeviceModelRequest) SetSecurityChip(v string) *CreateDeviceModelRequest {
	s.SecurityChip = &v
	return s
}

type CreateDeviceModelResponseBody struct {
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeviceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceModelResponseBody) SetDeviceModelId(v int64) *CreateDeviceModelResponseBody {
	s.DeviceModelId = &v
	return s
}

func (s *CreateDeviceModelResponseBody) SetRequestId(v string) *CreateDeviceModelResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeviceModelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceModelResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceModelResponse) SetHeaders(v map[string]*string) *CreateDeviceModelResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceModelResponse) SetStatusCode(v int32) *CreateDeviceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceModelResponse) SetBody(v *CreateDeviceModelResponseBody) *CreateDeviceModelResponse {
	s.Body = v
	return s
}

type CreateNamespaceRequest struct {
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetAuthType(v string) *CreateNamespaceRequest {
	s.AuthType = &v
	return s
}

func (s *CreateNamespaceRequest) SetDesc(v string) *CreateNamespaceRequest {
	s.Desc = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) SetProjectId(v string) *CreateNamespaceRequest {
	s.ProjectId = &v
	return s
}

type CreateNamespaceResponseBody struct {
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetNamespace(v string) *CreateNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type CreateOsVersionRequest struct {
	BlackVersionList            *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	DeviceModelId               *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	EnableMobileDownload        *string `json:"EnableMobileDownload,omitempty" xml:"EnableMobileDownload,omitempty"`
	IsForceNightUpgrade         *string `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	IsForceUpgrade              *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsMilestone                 *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	MaxClientVersion            *string `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	MinClientVersion            *string `json:"MinClientVersion,omitempty" xml:"MinClientVersion,omitempty"`
	MobileDownloadMaxSize       *string `json:"MobileDownloadMaxSize,omitempty" xml:"MobileDownloadMaxSize,omitempty"`
	NightUpgradeDownloadType    *string `json:"NightUpgradeDownloadType,omitempty" xml:"NightUpgradeDownloadType,omitempty"`
	NightUpgradeIsAllowedCancel *string `json:"NightUpgradeIsAllowedCancel,omitempty" xml:"NightUpgradeIsAllowedCancel,omitempty"`
	NightUpgradeIsShowTip       *string `json:"NightUpgradeIsShowTip,omitempty" xml:"NightUpgradeIsShowTip,omitempty"`
	ProjectId                   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReleaseNote                 *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark                      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RomList                     *string `json:"RomList,omitempty" xml:"RomList,omitempty"`
	SystemVersion               *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	WhiteVersionList            *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
}

func (s CreateOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateOsVersionRequest) SetBlackVersionList(v string) *CreateOsVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *CreateOsVersionRequest) SetDeviceModelId(v string) *CreateOsVersionRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CreateOsVersionRequest) SetEnableMobileDownload(v string) *CreateOsVersionRequest {
	s.EnableMobileDownload = &v
	return s
}

func (s *CreateOsVersionRequest) SetIsForceNightUpgrade(v string) *CreateOsVersionRequest {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *CreateOsVersionRequest) SetIsForceUpgrade(v string) *CreateOsVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *CreateOsVersionRequest) SetIsMilestone(v string) *CreateOsVersionRequest {
	s.IsMilestone = &v
	return s
}

func (s *CreateOsVersionRequest) SetMaxClientVersion(v string) *CreateOsVersionRequest {
	s.MaxClientVersion = &v
	return s
}

func (s *CreateOsVersionRequest) SetMinClientVersion(v string) *CreateOsVersionRequest {
	s.MinClientVersion = &v
	return s
}

func (s *CreateOsVersionRequest) SetMobileDownloadMaxSize(v string) *CreateOsVersionRequest {
	s.MobileDownloadMaxSize = &v
	return s
}

func (s *CreateOsVersionRequest) SetNightUpgradeDownloadType(v string) *CreateOsVersionRequest {
	s.NightUpgradeDownloadType = &v
	return s
}

func (s *CreateOsVersionRequest) SetNightUpgradeIsAllowedCancel(v string) *CreateOsVersionRequest {
	s.NightUpgradeIsAllowedCancel = &v
	return s
}

func (s *CreateOsVersionRequest) SetNightUpgradeIsShowTip(v string) *CreateOsVersionRequest {
	s.NightUpgradeIsShowTip = &v
	return s
}

func (s *CreateOsVersionRequest) SetProjectId(v string) *CreateOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateOsVersionRequest) SetReleaseNote(v string) *CreateOsVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *CreateOsVersionRequest) SetRemark(v string) *CreateOsVersionRequest {
	s.Remark = &v
	return s
}

func (s *CreateOsVersionRequest) SetRomList(v string) *CreateOsVersionRequest {
	s.RomList = &v
	return s
}

func (s *CreateOsVersionRequest) SetSystemVersion(v string) *CreateOsVersionRequest {
	s.SystemVersion = &v
	return s
}

func (s *CreateOsVersionRequest) SetWhiteVersionList(v string) *CreateOsVersionRequest {
	s.WhiteVersionList = &v
	return s
}

type CreateOsVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOsVersionResponseBody) SetRequestId(v string) *CreateOsVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOsVersionResponseBody) SetVersionId(v string) *CreateOsVersionResponseBody {
	s.VersionId = &v
	return s
}

type CreateOsVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateOsVersionResponse) SetHeaders(v map[string]*string) *CreateOsVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateOsVersionResponse) SetStatusCode(v int32) *CreateOsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOsVersionResponse) SetBody(v *CreateOsVersionResponseBody) *CreateOsVersionResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	ProjectDesc *string `json:"ProjectDesc,omitempty" xml:"ProjectDesc,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetProjectDesc(v string) *CreateProjectRequest {
	s.ProjectDesc = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

type CreateProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateProjectAppRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPkgName *string `json:"AppPkgName,omitempty" xml:"AppPkgName,omitempty"`
	OsType     *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateProjectAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectAppRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectAppRequest) SetAppName(v string) *CreateProjectAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateProjectAppRequest) SetAppPkgName(v string) *CreateProjectAppRequest {
	s.AppPkgName = &v
	return s
}

func (s *CreateProjectAppRequest) SetOsType(v int32) *CreateProjectAppRequest {
	s.OsType = &v
	return s
}

func (s *CreateProjectAppRequest) SetProjectId(v string) *CreateProjectAppRequest {
	s.ProjectId = &v
	return s
}

type CreateProjectAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectAppResponseBody) SetRequestId(v string) *CreateProjectAppResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectAppResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectAppResponse) SetHeaders(v map[string]*string) *CreateProjectAppResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectAppResponse) SetStatusCode(v int32) *CreateProjectAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectAppResponse) SetBody(v *CreateProjectAppResponseBody) *CreateProjectAppResponse {
	s.Body = v
	return s
}

type CreateRpcServiceRequest struct {
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	InvokeType    *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	MethodName    *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	Params        *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionCode   *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateRpcServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRpcServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateRpcServiceRequest) SetAppKey(v string) *CreateRpcServiceRequest {
	s.AppKey = &v
	return s
}

func (s *CreateRpcServiceRequest) SetGroupName(v string) *CreateRpcServiceRequest {
	s.GroupName = &v
	return s
}

func (s *CreateRpcServiceRequest) SetInterfaceName(v string) *CreateRpcServiceRequest {
	s.InterfaceName = &v
	return s
}

func (s *CreateRpcServiceRequest) SetInvokeType(v string) *CreateRpcServiceRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateRpcServiceRequest) SetMethodName(v string) *CreateRpcServiceRequest {
	s.MethodName = &v
	return s
}

func (s *CreateRpcServiceRequest) SetParams(v string) *CreateRpcServiceRequest {
	s.Params = &v
	return s
}

func (s *CreateRpcServiceRequest) SetProjectId(v string) *CreateRpcServiceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateRpcServiceRequest) SetVersionCode(v string) *CreateRpcServiceRequest {
	s.VersionCode = &v
	return s
}

type CreateRpcServiceResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRpcServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRpcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRpcServiceResponseBody) SetId(v int64) *CreateRpcServiceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateRpcServiceResponseBody) SetRequestId(v string) *CreateRpcServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateRpcServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRpcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRpcServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRpcServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateRpcServiceResponse) SetHeaders(v map[string]*string) *CreateRpcServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateRpcServiceResponse) SetStatusCode(v int32) *CreateRpcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRpcServiceResponse) SetBody(v *CreateRpcServiceResponseBody) *CreateRpcServiceResponse {
	s.Body = v
	return s
}

type CreateSchemaSubscribeRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubscribeList *string `json:"SubscribeList,omitempty" xml:"SubscribeList,omitempty"`
}

func (s CreateSchemaSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemaSubscribeRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemaSubscribeRequest) SetDeviceModel(v string) *CreateSchemaSubscribeRequest {
	s.DeviceModel = &v
	return s
}

func (s *CreateSchemaSubscribeRequest) SetProjectId(v string) *CreateSchemaSubscribeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateSchemaSubscribeRequest) SetSchemaVersion(v string) *CreateSchemaSubscribeRequest {
	s.SchemaVersion = &v
	return s
}

func (s *CreateSchemaSubscribeRequest) SetSubscribeList(v string) *CreateSchemaSubscribeRequest {
	s.SubscribeList = &v
	return s
}

type CreateSchemaSubscribeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchemaSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemaSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchemaSubscribeResponseBody) SetRequestId(v string) *CreateSchemaSubscribeResponseBody {
	s.RequestId = &v
	return s
}

type CreateSchemaSubscribeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSchemaSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSchemaSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemaSubscribeResponse) GoString() string {
	return s.String()
}

func (s *CreateSchemaSubscribeResponse) SetHeaders(v map[string]*string) *CreateSchemaSubscribeResponse {
	s.Headers = v
	return s
}

func (s *CreateSchemaSubscribeResponse) SetStatusCode(v int32) *CreateSchemaSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchemaSubscribeResponse) SetBody(v *CreateSchemaSubscribeResponseBody) *CreateSchemaSubscribeResponse {
	s.Body = v
	return s
}

type CreateShadowSchemaRequest struct {
	AuthType      *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Schema        *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s CreateShadowSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShadowSchemaRequest) GoString() string {
	return s.String()
}

func (s *CreateShadowSchemaRequest) SetAuthType(v string) *CreateShadowSchemaRequest {
	s.AuthType = &v
	return s
}

func (s *CreateShadowSchemaRequest) SetDeviceModelId(v string) *CreateShadowSchemaRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CreateShadowSchemaRequest) SetNamespace(v string) *CreateShadowSchemaRequest {
	s.Namespace = &v
	return s
}

func (s *CreateShadowSchemaRequest) SetProjectId(v string) *CreateShadowSchemaRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateShadowSchemaRequest) SetSchema(v string) *CreateShadowSchemaRequest {
	s.Schema = &v
	return s
}

type CreateShadowSchemaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateShadowSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateShadowSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShadowSchemaResponseBody) SetRequestId(v string) *CreateShadowSchemaResponseBody {
	s.RequestId = &v
	return s
}

type CreateShadowSchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateShadowSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShadowSchemaResponse) GoString() string {
	return s.String()
}

func (s *CreateShadowSchemaResponse) SetHeaders(v map[string]*string) *CreateShadowSchemaResponse {
	s.Headers = v
	return s
}

func (s *CreateShadowSchemaResponse) SetStatusCode(v int32) *CreateShadowSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShadowSchemaResponse) SetBody(v *CreateShadowSchemaResponseBody) *CreateShadowSchemaResponse {
	s.Body = v
	return s
}

type CreateTriggerRequest struct {
	FileIds        *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	FunctionIds    *string `json:"FunctionIds,omitempty" xml:"FunctionIds,omitempty"`
	InvocationMode *int32  `json:"InvocationMode,omitempty" xml:"InvocationMode,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Production     *int32  `json:"Production,omitempty" xml:"Production,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Sandbox        *int32  `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetFileIds(v string) *CreateTriggerRequest {
	s.FileIds = &v
	return s
}

func (s *CreateTriggerRequest) SetFunctionIds(v string) *CreateTriggerRequest {
	s.FunctionIds = &v
	return s
}

func (s *CreateTriggerRequest) SetInvocationMode(v int32) *CreateTriggerRequest {
	s.InvocationMode = &v
	return s
}

func (s *CreateTriggerRequest) SetNamespace(v string) *CreateTriggerRequest {
	s.Namespace = &v
	return s
}

func (s *CreateTriggerRequest) SetProduction(v int32) *CreateTriggerRequest {
	s.Production = &v
	return s
}

func (s *CreateTriggerRequest) SetProjectId(v string) *CreateTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateTriggerRequest) SetSandbox(v int32) *CreateTriggerRequest {
	s.Sandbox = &v
	return s
}

func (s *CreateTriggerRequest) SetSource(v string) *CreateTriggerRequest {
	s.Source = &v
	return s
}

type CreateTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) SetRequestId(v string) *CreateTriggerResponseBody {
	s.RequestId = &v
	return s
}

type CreateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetStatusCode(v int32) *CreateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse {
	s.Body = v
	return s
}

type CreateUpstreamAppKeyRelationRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	PAppKey   *string `json:"PAppKey,omitempty" xml:"PAppKey,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateUpstreamAppKeyRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationRequest) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationRequest) SetAppKey(v string) *CreateUpstreamAppKeyRelationRequest {
	s.AppKey = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationRequest) SetPAppKey(v string) *CreateUpstreamAppKeyRelationRequest {
	s.PAppKey = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationRequest) SetProjectId(v string) *CreateUpstreamAppKeyRelationRequest {
	s.ProjectId = &v
	return s
}

type CreateUpstreamAppKeyRelationResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUpstreamAppKeyRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationResponseBody) SetId(v int64) *CreateUpstreamAppKeyRelationResponseBody {
	s.Id = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationResponseBody) SetRequestId(v string) *CreateUpstreamAppKeyRelationResponseBody {
	s.RequestId = &v
	return s
}

type CreateUpstreamAppKeyRelationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUpstreamAppKeyRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUpstreamAppKeyRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationResponse) SetHeaders(v map[string]*string) *CreateUpstreamAppKeyRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateUpstreamAppKeyRelationResponse) SetStatusCode(v int32) *CreateUpstreamAppKeyRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationResponse) SetBody(v *CreateUpstreamAppKeyRelationResponseBody) *CreateUpstreamAppKeyRelationResponse {
	s.Body = v
	return s
}

type CreateUpstreamAppKeyRelationsRequest struct {
	AppKeys     *string `json:"AppKeys,omitempty" xml:"AppKeys,omitempty"`
	AppServerId *string `json:"AppServerId,omitempty" xml:"AppServerId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateUpstreamAppKeyRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationsRequest) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationsRequest) SetAppKeys(v string) *CreateUpstreamAppKeyRelationsRequest {
	s.AppKeys = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationsRequest) SetAppServerId(v string) *CreateUpstreamAppKeyRelationsRequest {
	s.AppServerId = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationsRequest) SetProjectId(v string) *CreateUpstreamAppKeyRelationsRequest {
	s.ProjectId = &v
	return s
}

type CreateUpstreamAppKeyRelationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUpstreamAppKeyRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationsResponseBody) SetRequestId(v string) *CreateUpstreamAppKeyRelationsResponseBody {
	s.RequestId = &v
	return s
}

type CreateUpstreamAppKeyRelationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUpstreamAppKeyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUpstreamAppKeyRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationsResponse) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationsResponse) SetHeaders(v map[string]*string) *CreateUpstreamAppKeyRelationsResponse {
	s.Headers = v
	return s
}

func (s *CreateUpstreamAppKeyRelationsResponse) SetStatusCode(v int32) *CreateUpstreamAppKeyRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationsResponse) SetBody(v *CreateUpstreamAppKeyRelationsResponseBody) *CreateUpstreamAppKeyRelationsResponse {
	s.Body = v
	return s
}

type CreateUpstreamAppServerRequest struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateUpstreamAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppServerRequest) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppServerRequest) SetName(v string) *CreateUpstreamAppServerRequest {
	s.Name = &v
	return s
}

func (s *CreateUpstreamAppServerRequest) SetProjectId(v string) *CreateUpstreamAppServerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateUpstreamAppServerRequest) SetTags(v string) *CreateUpstreamAppServerRequest {
	s.Tags = &v
	return s
}

type CreateUpstreamAppServerResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUpstreamAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppServerResponseBody) SetId(v int64) *CreateUpstreamAppServerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateUpstreamAppServerResponseBody) SetRequestId(v string) *CreateUpstreamAppServerResponseBody {
	s.RequestId = &v
	return s
}

type CreateUpstreamAppServerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUpstreamAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUpstreamAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppServerResponse) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppServerResponse) SetHeaders(v map[string]*string) *CreateUpstreamAppServerResponse {
	s.Headers = v
	return s
}

func (s *CreateUpstreamAppServerResponse) SetStatusCode(v int32) *CreateUpstreamAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUpstreamAppServerResponse) SetBody(v *CreateUpstreamAppServerResponseBody) *CreateUpstreamAppServerResponse {
	s.Body = v
	return s
}

type CreateVersionDeviceGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxCount    *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionDeviceGroupRequest) SetDescription(v string) *CreateVersionDeviceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateVersionDeviceGroupRequest) SetMaxCount(v string) *CreateVersionDeviceGroupRequest {
	s.MaxCount = &v
	return s
}

func (s *CreateVersionDeviceGroupRequest) SetName(v string) *CreateVersionDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateVersionDeviceGroupRequest) SetProjectId(v string) *CreateVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

type CreateVersionDeviceGroupResponseBody struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVersionDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionDeviceGroupResponseBody) SetDeviceGroupId(v string) *CreateVersionDeviceGroupResponseBody {
	s.DeviceGroupId = &v
	return s
}

func (s *CreateVersionDeviceGroupResponseBody) SetRequestId(v string) *CreateVersionDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateVersionDeviceGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVersionDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateVersionDeviceGroupResponse) SetHeaders(v map[string]*string) *CreateVersionDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateVersionDeviceGroupResponse) SetStatusCode(v int32) *CreateVersionDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVersionDeviceGroupResponse) SetBody(v *CreateVersionDeviceGroupResponseBody) *CreateVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type CreateVersionPrepublishRequest struct {
	BarrierCount      *string `json:"BarrierCount,omitempty" xml:"BarrierCount,omitempty"`
	IsTotalPrepublish *string `json:"IsTotalPrepublish,omitempty" xml:"IsTotalPrepublish,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId         *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType       *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateVersionPrepublishRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionPrepublishRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionPrepublishRequest) SetBarrierCount(v string) *CreateVersionPrepublishRequest {
	s.BarrierCount = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetIsTotalPrepublish(v string) *CreateVersionPrepublishRequest {
	s.IsTotalPrepublish = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetName(v string) *CreateVersionPrepublishRequest {
	s.Name = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetProjectId(v string) *CreateVersionPrepublishRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetVersionId(v string) *CreateVersionPrepublishRequest {
	s.VersionId = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetVersionType(v string) *CreateVersionPrepublishRequest {
	s.VersionType = &v
	return s
}

type CreateVersionPrepublishResponseBody struct {
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVersionPrepublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionPrepublishResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionPrepublishResponseBody) SetPrepublishId(v string) *CreateVersionPrepublishResponseBody {
	s.PrepublishId = &v
	return s
}

func (s *CreateVersionPrepublishResponseBody) SetRequestId(v string) *CreateVersionPrepublishResponseBody {
	s.RequestId = &v
	return s
}

type CreateVersionPrepublishResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVersionPrepublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVersionPrepublishResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionPrepublishResponse) GoString() string {
	return s.String()
}

func (s *CreateVersionPrepublishResponse) SetHeaders(v map[string]*string) *CreateVersionPrepublishResponse {
	s.Headers = v
	return s
}

func (s *CreateVersionPrepublishResponse) SetStatusCode(v int32) *CreateVersionPrepublishResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVersionPrepublishResponse) SetBody(v *CreateVersionPrepublishResponseBody) *CreateVersionPrepublishResponse {
	s.Body = v
	return s
}

type CreateVersionTestRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType   *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateVersionTestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionTestRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionTestRequest) SetDescription(v string) *CreateVersionTestRequest {
	s.Description = &v
	return s
}

func (s *CreateVersionTestRequest) SetDeviceGroupId(v string) *CreateVersionTestRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *CreateVersionTestRequest) SetName(v string) *CreateVersionTestRequest {
	s.Name = &v
	return s
}

func (s *CreateVersionTestRequest) SetProjectId(v string) *CreateVersionTestRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateVersionTestRequest) SetVersionId(v string) *CreateVersionTestRequest {
	s.VersionId = &v
	return s
}

func (s *CreateVersionTestRequest) SetVersionType(v string) *CreateVersionTestRequest {
	s.VersionType = &v
	return s
}

type CreateVersionTestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TestId    *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s CreateVersionTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionTestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionTestResponseBody) SetRequestId(v string) *CreateVersionTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVersionTestResponseBody) SetTestId(v string) *CreateVersionTestResponseBody {
	s.TestId = &v
	return s
}

type CreateVersionTestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVersionTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVersionTestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionTestResponse) GoString() string {
	return s.String()
}

func (s *CreateVersionTestResponse) SetHeaders(v map[string]*string) *CreateVersionTestResponse {
	s.Headers = v
	return s
}

func (s *CreateVersionTestResponse) SetStatusCode(v int32) *CreateVersionTestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVersionTestResponse) SetBody(v *CreateVersionTestResponseBody) *CreateVersionTestResponse {
	s.Body = v
	return s
}

type DelayPublishOsVersionRequest struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DownTime        *int64  `json:"DownTime,omitempty" xml:"DownTime,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PrepubTime      *int64  `json:"PrepubTime,omitempty" xml:"PrepubTime,omitempty"`
	PrepublishCount *string `json:"PrepublishCount,omitempty" xml:"PrepublishCount,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PublishTime     *int64  `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	SendMessage     *string `json:"SendMessage,omitempty" xml:"SendMessage,omitempty"`
	VersionId       *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DelayPublishOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DelayPublishOsVersionRequest) GoString() string {
	return s.String()
}

func (s *DelayPublishOsVersionRequest) SetDescription(v string) *DelayPublishOsVersionRequest {
	s.Description = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetDownTime(v int64) *DelayPublishOsVersionRequest {
	s.DownTime = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetEmail(v string) *DelayPublishOsVersionRequest {
	s.Email = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetPrepubTime(v int64) *DelayPublishOsVersionRequest {
	s.PrepubTime = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetPrepublishCount(v string) *DelayPublishOsVersionRequest {
	s.PrepublishCount = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetProjectId(v string) *DelayPublishOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetPublishTime(v int64) *DelayPublishOsVersionRequest {
	s.PublishTime = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetSendMessage(v string) *DelayPublishOsVersionRequest {
	s.SendMessage = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetVersionId(v string) *DelayPublishOsVersionRequest {
	s.VersionId = &v
	return s
}

type DelayPublishOsVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DelayPublishOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelayPublishOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DelayPublishOsVersionResponseBody) SetRequestId(v string) *DelayPublishOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type DelayPublishOsVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DelayPublishOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DelayPublishOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DelayPublishOsVersionResponse) GoString() string {
	return s.String()
}

func (s *DelayPublishOsVersionResponse) SetHeaders(v map[string]*string) *DelayPublishOsVersionResponse {
	s.Headers = v
	return s
}

func (s *DelayPublishOsVersionResponse) SetStatusCode(v int32) *DelayPublishOsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DelayPublishOsVersionResponse) SetBody(v *DelayPublishOsVersionResponseBody) *DelayPublishOsVersionResponse {
	s.Body = v
	return s
}

type DeleteAllCustomizedFiltersRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteAllCustomizedFiltersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllCustomizedFiltersRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllCustomizedFiltersRequest) SetProjectId(v string) *DeleteAllCustomizedFiltersRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteAllCustomizedFiltersRequest) SetVersionId(v string) *DeleteAllCustomizedFiltersRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteAllCustomizedFiltersRequest) SetVersionType(v string) *DeleteAllCustomizedFiltersRequest {
	s.VersionType = &v
	return s
}

type DeleteAllCustomizedFiltersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAllCustomizedFiltersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllCustomizedFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllCustomizedFiltersResponseBody) SetRequestId(v string) *DeleteAllCustomizedFiltersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAllCustomizedFiltersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAllCustomizedFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAllCustomizedFiltersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllCustomizedFiltersResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllCustomizedFiltersResponse) SetHeaders(v map[string]*string) *DeleteAllCustomizedFiltersResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllCustomizedFiltersResponse) SetStatusCode(v int32) *DeleteAllCustomizedFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllCustomizedFiltersResponse) SetBody(v *DeleteAllCustomizedFiltersResponseBody) *DeleteAllCustomizedFiltersResponse {
	s.Body = v
	return s
}

type DeleteAllVersionGroupDevicesRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteAllVersionGroupDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllVersionGroupDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllVersionGroupDevicesRequest) SetDeviceGroupId(v string) *DeleteAllVersionGroupDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteAllVersionGroupDevicesRequest) SetProjectId(v string) *DeleteAllVersionGroupDevicesRequest {
	s.ProjectId = &v
	return s
}

type DeleteAllVersionGroupDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAllVersionGroupDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllVersionGroupDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllVersionGroupDevicesResponseBody) SetRequestId(v string) *DeleteAllVersionGroupDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAllVersionGroupDevicesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAllVersionGroupDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAllVersionGroupDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllVersionGroupDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllVersionGroupDevicesResponse) SetHeaders(v map[string]*string) *DeleteAllVersionGroupDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllVersionGroupDevicesResponse) SetStatusCode(v int32) *DeleteAllVersionGroupDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllVersionGroupDevicesResponse) SetBody(v *DeleteAllVersionGroupDevicesResponseBody) *DeleteAllVersionGroupDevicesResponse {
	s.Body = v
	return s
}

type DeleteCustomizedFilterRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedFilterRequest) SetId(v string) *DeleteCustomizedFilterRequest {
	s.Id = &v
	return s
}

func (s *DeleteCustomizedFilterRequest) SetProjectId(v string) *DeleteCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

type DeleteCustomizedFilterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomizedFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedFilterResponseBody) SetRequestId(v string) *DeleteCustomizedFilterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomizedFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomizedFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedFilterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedFilterResponse) SetHeaders(v map[string]*string) *DeleteCustomizedFilterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizedFilterResponse) SetStatusCode(v int32) *DeleteCustomizedFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizedFilterResponse) SetBody(v *DeleteCustomizedFilterResponseBody) *DeleteCustomizedFilterResponse {
	s.Body = v
	return s
}

type DeleteCustomizedPropertyRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteCustomizedPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedPropertyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedPropertyRequest) SetId(v string) *DeleteCustomizedPropertyRequest {
	s.Id = &v
	return s
}

func (s *DeleteCustomizedPropertyRequest) SetProjectId(v string) *DeleteCustomizedPropertyRequest {
	s.ProjectId = &v
	return s
}

type DeleteCustomizedPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomizedPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedPropertyResponseBody) SetRequestId(v string) *DeleteCustomizedPropertyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomizedPropertyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomizedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomizedPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedPropertyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedPropertyResponse) SetHeaders(v map[string]*string) *DeleteCustomizedPropertyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizedPropertyResponse) SetStatusCode(v int32) *DeleteCustomizedPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizedPropertyResponse) SetBody(v *DeleteCustomizedPropertyResponseBody) *DeleteCustomizedPropertyResponse {
	s.Body = v
	return s
}

type DeleteDeviceRequest struct {
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) SetDeviceId(v string) *DeleteDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *DeleteDeviceRequest) SetProjectId(v string) *DeleteDeviceRequest {
	s.ProjectId = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDeviceResponse) SetStatusCode(v int32) *DeleteDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceResponse) SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse {
	s.Body = v
	return s
}

type DeleteFunctionFileRequest struct {
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteFunctionFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionFileRequest) SetFileName(v string) *DeleteFunctionFileRequest {
	s.FileName = &v
	return s
}

func (s *DeleteFunctionFileRequest) SetFileType(v int32) *DeleteFunctionFileRequest {
	s.FileType = &v
	return s
}

func (s *DeleteFunctionFileRequest) SetProjectId(v string) *DeleteFunctionFileRequest {
	s.ProjectId = &v
	return s
}

type DeleteFunctionFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFunctionFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionFileResponseBody) SetRequestId(v string) *DeleteFunctionFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFunctionFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFunctionFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionFileResponse) SetHeaders(v map[string]*string) *DeleteFunctionFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionFileResponse) SetStatusCode(v int32) *DeleteFunctionFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionFileResponse) SetBody(v *DeleteFunctionFileResponseBody) *DeleteFunctionFileResponse {
	s.Body = v
	return s
}

type DeleteNamespaceRequest struct {
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequest) SetProjectId(v string) *DeleteNamespaceRequest {
	s.ProjectId = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

type DeleteOpenAccountRequest struct {
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteOpenAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOpenAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpenAccountRequest) SetIdentityId(v string) *DeleteOpenAccountRequest {
	s.IdentityId = &v
	return s
}

func (s *DeleteOpenAccountRequest) SetProjectId(v string) *DeleteOpenAccountRequest {
	s.ProjectId = &v
	return s
}

type DeleteOpenAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOpenAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOpenAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpenAccountResponseBody) SetRequestId(v string) *DeleteOpenAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOpenAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteOpenAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOpenAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOpenAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpenAccountResponse) SetHeaders(v map[string]*string) *DeleteOpenAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpenAccountResponse) SetStatusCode(v int32) *DeleteOpenAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpenAccountResponse) SetBody(v *DeleteOpenAccountResponseBody) *DeleteOpenAccountResponse {
	s.Body = v
	return s
}

type DeleteProjectAppRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectAppRequest) SetAppId(v string) *DeleteProjectAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteProjectAppRequest) SetProjectId(v string) *DeleteProjectAppRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectAppResponseBody) SetRequestId(v string) *DeleteProjectAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectAppResponse) SetHeaders(v map[string]*string) *DeleteProjectAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectAppResponse) SetStatusCode(v int32) *DeleteProjectAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectAppResponse) SetBody(v *DeleteProjectAppResponseBody) *DeleteProjectAppResponse {
	s.Body = v
	return s
}

type DeleteRpcServiceRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteRpcServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRpcServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRpcServiceRequest) SetId(v string) *DeleteRpcServiceRequest {
	s.Id = &v
	return s
}

func (s *DeleteRpcServiceRequest) SetProjectId(v string) *DeleteRpcServiceRequest {
	s.ProjectId = &v
	return s
}

type DeleteRpcServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRpcServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRpcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRpcServiceResponseBody) SetRequestId(v string) *DeleteRpcServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRpcServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRpcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRpcServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRpcServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRpcServiceResponse) SetHeaders(v map[string]*string) *DeleteRpcServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRpcServiceResponse) SetStatusCode(v int32) *DeleteRpcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRpcServiceResponse) SetBody(v *DeleteRpcServiceResponseBody) *DeleteRpcServiceResponse {
	s.Body = v
	return s
}

type DeleteSchemaSubscribeRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteSchemaSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemaSubscribeRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchemaSubscribeRequest) SetId(v string) *DeleteSchemaSubscribeRequest {
	s.Id = &v
	return s
}

func (s *DeleteSchemaSubscribeRequest) SetProjectId(v string) *DeleteSchemaSubscribeRequest {
	s.ProjectId = &v
	return s
}

type DeleteSchemaSubscribeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchemaSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemaSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchemaSubscribeResponseBody) SetRequestId(v string) *DeleteSchemaSubscribeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSchemaSubscribeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSchemaSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSchemaSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchemaSubscribeResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchemaSubscribeResponse) SetHeaders(v map[string]*string) *DeleteSchemaSubscribeResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchemaSubscribeResponse) SetStatusCode(v int32) *DeleteSchemaSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchemaSubscribeResponse) SetBody(v *DeleteSchemaSubscribeResponseBody) *DeleteSchemaSubscribeResponse {
	s.Body = v
	return s
}

type DeleteShadowSchemaRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteShadowSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteShadowSchemaRequest) GoString() string {
	return s.String()
}

func (s *DeleteShadowSchemaRequest) SetId(v string) *DeleteShadowSchemaRequest {
	s.Id = &v
	return s
}

func (s *DeleteShadowSchemaRequest) SetProjectId(v string) *DeleteShadowSchemaRequest {
	s.ProjectId = &v
	return s
}

type DeleteShadowSchemaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteShadowSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteShadowSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteShadowSchemaResponseBody) SetRequestId(v string) *DeleteShadowSchemaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteShadowSchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteShadowSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteShadowSchemaResponse) GoString() string {
	return s.String()
}

func (s *DeleteShadowSchemaResponse) SetHeaders(v map[string]*string) *DeleteShadowSchemaResponse {
	s.Headers = v
	return s
}

func (s *DeleteShadowSchemaResponse) SetStatusCode(v int32) *DeleteShadowSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteShadowSchemaResponse) SetBody(v *DeleteShadowSchemaResponseBody) *DeleteShadowSchemaResponse {
	s.Body = v
	return s
}

type DeleteTriggerRequest struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteTriggerRequest) SetId(v int64) *DeleteTriggerRequest {
	s.Id = &v
	return s
}

func (s *DeleteTriggerRequest) SetProjectId(v string) *DeleteTriggerRequest {
	s.ProjectId = &v
	return s
}

type DeleteTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponseBody) SetRequestId(v string) *DeleteTriggerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponse) SetHeaders(v map[string]*string) *DeleteTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteTriggerResponse) SetStatusCode(v int32) *DeleteTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTriggerResponse) SetBody(v *DeleteTriggerResponseBody) *DeleteTriggerResponse {
	s.Body = v
	return s
}

type DeleteUpstreamAppKeyRelationRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteUpstreamAppKeyRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppKeyRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppKeyRelationRequest) SetId(v string) *DeleteUpstreamAppKeyRelationRequest {
	s.Id = &v
	return s
}

func (s *DeleteUpstreamAppKeyRelationRequest) SetProjectId(v string) *DeleteUpstreamAppKeyRelationRequest {
	s.ProjectId = &v
	return s
}

type DeleteUpstreamAppKeyRelationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUpstreamAppKeyRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppKeyRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppKeyRelationResponseBody) SetRequestId(v string) *DeleteUpstreamAppKeyRelationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUpstreamAppKeyRelationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUpstreamAppKeyRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUpstreamAppKeyRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppKeyRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppKeyRelationResponse) SetHeaders(v map[string]*string) *DeleteUpstreamAppKeyRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteUpstreamAppKeyRelationResponse) SetStatusCode(v int32) *DeleteUpstreamAppKeyRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUpstreamAppKeyRelationResponse) SetBody(v *DeleteUpstreamAppKeyRelationResponseBody) *DeleteUpstreamAppKeyRelationResponse {
	s.Body = v
	return s
}

type DeleteUpstreamAppServerRequest struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteUpstreamAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppServerRequest) SetId(v int64) *DeleteUpstreamAppServerRequest {
	s.Id = &v
	return s
}

func (s *DeleteUpstreamAppServerRequest) SetProjectId(v string) *DeleteUpstreamAppServerRequest {
	s.ProjectId = &v
	return s
}

type DeleteUpstreamAppServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUpstreamAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppServerResponseBody) SetRequestId(v string) *DeleteUpstreamAppServerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUpstreamAppServerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUpstreamAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUpstreamAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppServerResponse) SetHeaders(v map[string]*string) *DeleteUpstreamAppServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteUpstreamAppServerResponse) SetStatusCode(v int32) *DeleteUpstreamAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUpstreamAppServerResponse) SetBody(v *DeleteUpstreamAppServerResponseBody) *DeleteUpstreamAppServerResponse {
	s.Body = v
	return s
}

type DeleteVersionAllBlackDevicesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteVersionAllBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionAllBlackDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionAllBlackDevicesRequest) SetProjectId(v string) *DeleteVersionAllBlackDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionAllBlackDevicesRequest) SetVersionId(v string) *DeleteVersionAllBlackDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionAllBlackDevicesRequest) SetVersionType(v string) *DeleteVersionAllBlackDevicesRequest {
	s.VersionType = &v
	return s
}

type DeleteVersionAllBlackDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionAllBlackDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionAllBlackDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionAllBlackDevicesResponseBody) SetRequestId(v string) *DeleteVersionAllBlackDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionAllBlackDevicesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionAllBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionAllBlackDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionAllBlackDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionAllBlackDevicesResponse) SetHeaders(v map[string]*string) *DeleteVersionAllBlackDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionAllBlackDevicesResponse) SetStatusCode(v int32) *DeleteVersionAllBlackDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionAllBlackDevicesResponse) SetBody(v *DeleteVersionAllBlackDevicesResponseBody) *DeleteVersionAllBlackDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionAllWhiteDevicesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteVersionAllWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionAllWhiteDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionAllWhiteDevicesRequest) SetProjectId(v string) *DeleteVersionAllWhiteDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionAllWhiteDevicesRequest) SetVersionId(v string) *DeleteVersionAllWhiteDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionAllWhiteDevicesRequest) SetVersionType(v string) *DeleteVersionAllWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

type DeleteVersionAllWhiteDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionAllWhiteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionAllWhiteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionAllWhiteDevicesResponseBody) SetRequestId(v string) *DeleteVersionAllWhiteDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionAllWhiteDevicesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionAllWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionAllWhiteDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionAllWhiteDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionAllWhiteDevicesResponse) SetHeaders(v map[string]*string) *DeleteVersionAllWhiteDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionAllWhiteDevicesResponse) SetStatusCode(v int32) *DeleteVersionAllWhiteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionAllWhiteDevicesResponse) SetBody(v *DeleteVersionAllWhiteDevicesResponseBody) *DeleteVersionAllWhiteDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionBlackDevicesRequest struct {
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteVersionBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesRequest) SetDeviceIdType(v string) *DeleteVersionBlackDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetDeviceIds(v string) *DeleteVersionBlackDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetProjectId(v string) *DeleteVersionBlackDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetVersionId(v string) *DeleteVersionBlackDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetVersionType(v string) *DeleteVersionBlackDevicesRequest {
	s.VersionType = &v
	return s
}

type DeleteVersionBlackDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionBlackDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesResponseBody) SetRequestId(v string) *DeleteVersionBlackDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionBlackDevicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionBlackDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesResponse) SetHeaders(v map[string]*string) *DeleteVersionBlackDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionBlackDevicesResponse) SetStatusCode(v int32) *DeleteVersionBlackDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionBlackDevicesResponse) SetBody(v *DeleteVersionBlackDevicesResponseBody) *DeleteVersionBlackDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionBlackDevicesByIdRequest struct {
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteVersionBlackDevicesByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesByIdRequest) SetIds(v string) *DeleteVersionBlackDevicesByIdRequest {
	s.Ids = &v
	return s
}

func (s *DeleteVersionBlackDevicesByIdRequest) SetProjectId(v string) *DeleteVersionBlackDevicesByIdRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionBlackDevicesByIdRequest) SetVersionId(v string) *DeleteVersionBlackDevicesByIdRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionBlackDevicesByIdRequest) SetVersionType(v string) *DeleteVersionBlackDevicesByIdRequest {
	s.VersionType = &v
	return s
}

type DeleteVersionBlackDevicesByIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionBlackDevicesByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesByIdResponseBody) SetRequestId(v string) *DeleteVersionBlackDevicesByIdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionBlackDevicesByIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionBlackDevicesByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionBlackDevicesByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesByIdResponse) SetHeaders(v map[string]*string) *DeleteVersionBlackDevicesByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionBlackDevicesByIdResponse) SetStatusCode(v int32) *DeleteVersionBlackDevicesByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionBlackDevicesByIdResponse) SetBody(v *DeleteVersionBlackDevicesByIdResponseBody) *DeleteVersionBlackDevicesByIdResponse {
	s.Body = v
	return s
}

type DeleteVersionDeviceGroupRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionDeviceGroupRequest) SetId(v string) *DeleteVersionDeviceGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteVersionDeviceGroupRequest) SetProjectId(v string) *DeleteVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

type DeleteVersionDeviceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionDeviceGroupResponseBody) SetRequestId(v string) *DeleteVersionDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionDeviceGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionDeviceGroupResponse) SetHeaders(v map[string]*string) *DeleteVersionDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionDeviceGroupResponse) SetStatusCode(v int32) *DeleteVersionDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionDeviceGroupResponse) SetBody(v *DeleteVersionDeviceGroupResponseBody) *DeleteVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type DeleteVersionGroupDeviceRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceIdType  *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceIds     *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteVersionGroupDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceRequest) SetDeviceGroupId(v string) *DeleteVersionGroupDeviceRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteVersionGroupDeviceRequest) SetDeviceIdType(v string) *DeleteVersionGroupDeviceRequest {
	s.DeviceIdType = &v
	return s
}

func (s *DeleteVersionGroupDeviceRequest) SetDeviceIds(v string) *DeleteVersionGroupDeviceRequest {
	s.DeviceIds = &v
	return s
}

func (s *DeleteVersionGroupDeviceRequest) SetProjectId(v string) *DeleteVersionGroupDeviceRequest {
	s.ProjectId = &v
	return s
}

type DeleteVersionGroupDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionGroupDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceResponseBody) SetRequestId(v string) *DeleteVersionGroupDeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionGroupDeviceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionGroupDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionGroupDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceResponse) SetHeaders(v map[string]*string) *DeleteVersionGroupDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionGroupDeviceResponse) SetStatusCode(v int32) *DeleteVersionGroupDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionGroupDeviceResponse) SetBody(v *DeleteVersionGroupDeviceResponseBody) *DeleteVersionGroupDeviceResponse {
	s.Body = v
	return s
}

type DeleteVersionGroupDeviceByIdRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Ids           *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteVersionGroupDeviceByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceByIdRequest) SetDeviceGroupId(v string) *DeleteVersionGroupDeviceByIdRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteVersionGroupDeviceByIdRequest) SetIds(v string) *DeleteVersionGroupDeviceByIdRequest {
	s.Ids = &v
	return s
}

func (s *DeleteVersionGroupDeviceByIdRequest) SetProjectId(v string) *DeleteVersionGroupDeviceByIdRequest {
	s.ProjectId = &v
	return s
}

type DeleteVersionGroupDeviceByIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionGroupDeviceByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceByIdResponseBody) SetRequestId(v string) *DeleteVersionGroupDeviceByIdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionGroupDeviceByIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionGroupDeviceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionGroupDeviceByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceByIdResponse) SetHeaders(v map[string]*string) *DeleteVersionGroupDeviceByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionGroupDeviceByIdResponse) SetStatusCode(v int32) *DeleteVersionGroupDeviceByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionGroupDeviceByIdResponse) SetBody(v *DeleteVersionGroupDeviceByIdResponseBody) *DeleteVersionGroupDeviceByIdResponse {
	s.Body = v
	return s
}

type DeleteVersionTestRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteVersionTestRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionTestRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionTestRequest) SetId(v string) *DeleteVersionTestRequest {
	s.Id = &v
	return s
}

func (s *DeleteVersionTestRequest) SetProjectId(v string) *DeleteVersionTestRequest {
	s.ProjectId = &v
	return s
}

type DeleteVersionTestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionTestResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionTestResponseBody) SetRequestId(v string) *DeleteVersionTestResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionTestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionTestResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionTestResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionTestResponse) SetHeaders(v map[string]*string) *DeleteVersionTestResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionTestResponse) SetStatusCode(v int32) *DeleteVersionTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionTestResponse) SetBody(v *DeleteVersionTestResponseBody) *DeleteVersionTestResponse {
	s.Body = v
	return s
}

type DeleteVersionWhiteDevicesRequest struct {
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteVersionWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesRequest) SetDeviceIdType(v string) *DeleteVersionWhiteDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetDeviceIds(v string) *DeleteVersionWhiteDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetProjectId(v string) *DeleteVersionWhiteDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetVersionId(v string) *DeleteVersionWhiteDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetVersionType(v string) *DeleteVersionWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

type DeleteVersionWhiteDevicesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionWhiteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesResponseBody) SetRequestId(v string) *DeleteVersionWhiteDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionWhiteDevicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionWhiteDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesResponse) SetHeaders(v map[string]*string) *DeleteVersionWhiteDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionWhiteDevicesResponse) SetStatusCode(v int32) *DeleteVersionWhiteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionWhiteDevicesResponse) SetBody(v *DeleteVersionWhiteDevicesResponseBody) *DeleteVersionWhiteDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionWhiteDevicesByIdRequest struct {
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DeleteVersionWhiteDevicesByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesByIdRequest) SetIds(v string) *DeleteVersionWhiteDevicesByIdRequest {
	s.Ids = &v
	return s
}

func (s *DeleteVersionWhiteDevicesByIdRequest) SetProjectId(v string) *DeleteVersionWhiteDevicesByIdRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionWhiteDevicesByIdRequest) SetVersionId(v string) *DeleteVersionWhiteDevicesByIdRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionWhiteDevicesByIdRequest) SetVersionType(v string) *DeleteVersionWhiteDevicesByIdRequest {
	s.VersionType = &v
	return s
}

type DeleteVersionWhiteDevicesByIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionWhiteDevicesByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesByIdResponseBody) SetRequestId(v string) *DeleteVersionWhiteDevicesByIdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVersionWhiteDevicesByIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVersionWhiteDevicesByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVersionWhiteDevicesByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesByIdResponse) SetHeaders(v map[string]*string) *DeleteVersionWhiteDevicesByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionWhiteDevicesByIdResponse) SetStatusCode(v int32) *DeleteVersionWhiteDevicesByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionWhiteDevicesByIdResponse) SetBody(v *DeleteVersionWhiteDevicesByIdResponseBody) *DeleteVersionWhiteDevicesByIdResponse {
	s.Body = v
	return s
}

type DeployFunctionFileRequest struct {
	DeployEnv *int32  `json:"DeployEnv,omitempty" xml:"DeployEnv,omitempty"`
	FileId    *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeployFunctionFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionFileRequest) GoString() string {
	return s.String()
}

func (s *DeployFunctionFileRequest) SetDeployEnv(v int32) *DeployFunctionFileRequest {
	s.DeployEnv = &v
	return s
}

func (s *DeployFunctionFileRequest) SetFileId(v string) *DeployFunctionFileRequest {
	s.FileId = &v
	return s
}

func (s *DeployFunctionFileRequest) SetProjectId(v string) *DeployFunctionFileRequest {
	s.ProjectId = &v
	return s
}

type DeployFunctionFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployFunctionFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeployFunctionFileResponseBody) SetRequestId(v string) *DeployFunctionFileResponseBody {
	s.RequestId = &v
	return s
}

type DeployFunctionFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployFunctionFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployFunctionFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionFileResponse) GoString() string {
	return s.String()
}

func (s *DeployFunctionFileResponse) SetHeaders(v map[string]*string) *DeployFunctionFileResponse {
	s.Headers = v
	return s
}

func (s *DeployFunctionFileResponse) SetStatusCode(v int32) *DeployFunctionFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployFunctionFileResponse) SetBody(v *DeployFunctionFileResponseBody) *DeployFunctionFileResponse {
	s.Body = v
	return s
}

type DescribeApiGatewayAppSecurityRequest struct {
	GatewayAppId *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeApiGatewayAppSecurityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityRequest) SetGatewayAppId(v string) *DescribeApiGatewayAppSecurityRequest {
	s.GatewayAppId = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityRequest) SetProjectId(v string) *DescribeApiGatewayAppSecurityRequest {
	s.ProjectId = &v
	return s
}

type DescribeApiGatewayAppSecurityResponseBody struct {
	ApiGatewayAppSecurity *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity `json:"ApiGatewayAppSecurity,omitempty" xml:"ApiGatewayAppSecurity,omitempty" type:"Struct"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiGatewayAppSecurityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityResponseBody) SetApiGatewayAppSecurity(v *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) *DescribeApiGatewayAppSecurityResponseBody {
	s.ApiGatewayAppSecurity = v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBody) SetRequestId(v string) *DescribeApiGatewayAppSecurityResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity struct {
	GatewayAppId     *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	GatewayAppKey    *string `json:"GatewayAppKey,omitempty" xml:"GatewayAppKey,omitempty"`
	GatewayAppSecret *string `json:"GatewayAppSecret,omitempty" xml:"GatewayAppSecret,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGatewayAppId(v string) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GatewayAppId = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGatewayAppKey(v string) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GatewayAppKey = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGatewayAppSecret(v string) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GatewayAppSecret = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGmtCreate(v int64) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GmtCreate = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGmtModified(v int64) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GmtModified = &v
	return s
}

type DescribeApiGatewayAppSecurityResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiGatewayAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiGatewayAppSecurityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityResponse) SetHeaders(v map[string]*string) *DescribeApiGatewayAppSecurityResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponse) SetStatusCode(v int32) *DescribeApiGatewayAppSecurityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponse) SetBody(v *DescribeApiGatewayAppSecurityResponseBody) *DescribeApiGatewayAppSecurityResponse {
	s.Body = v
	return s
}

type DescribeAppVersionRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppVersionRequest) SetProjectId(v string) *DescribeAppVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAppVersionRequest) SetVersionId(v string) *DescribeAppVersionRequest {
	s.VersionId = &v
	return s
}

type DescribeAppVersionResponseBody struct {
	AppVersion *DescribeAppVersionResponseBodyAppVersion `json:"AppVersion,omitempty" xml:"AppVersion,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppVersionResponseBody) SetAppVersion(v *DescribeAppVersionResponseBodyAppVersion) *DescribeAppVersionResponseBody {
	s.AppVersion = v
	return s
}

func (s *DescribeAppVersionResponseBody) SetRequestId(v string) *DescribeAppVersionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppVersionResponseBodyAppVersion struct {
	Adapters          []*DescribeAppVersionResponseBodyAppVersionAdapters `json:"Adapters,omitempty" xml:"Adapters,omitempty" type:"Repeated"`
	ApkMd5            *string                                             `json:"ApkMd5,omitempty" xml:"ApkMd5,omitempty"`
	AppId             *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName           *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion        *string                                             `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BlackVersionList  *string                                             `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	DownloadUrl       *string                                             `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	GmtCreate         *string                                             `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify         *string                                             `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Id                *int64                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallType       *string                                             `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	IsAllowNewInstall *string                                             `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	IsForceUpgrade    *string                                             `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsNeedRestart     *string                                             `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	IsSilentUpgrade   *string                                             `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	Md5               *string                                             `json:"Md5,omitempty" xml:"Md5,omitempty"`
	OriginalUrl       *string                                             `json:"OriginalUrl,omitempty" xml:"OriginalUrl,omitempty"`
	PackageName       *string                                             `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ReleaseNote       *string                                             `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark            *string                                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RestartAppParam   *string                                             `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	RestartAppType    *string                                             `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	RestartType       *string                                             `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	Size              *string                                             `json:"Size,omitempty" xml:"Size,omitempty"`
	Status            *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName        *string                                             `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	VersionCode       *int64                                              `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	WhiteVersionList  *string                                             `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
}

func (s DescribeAppVersionResponseBodyAppVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppVersionResponseBodyAppVersion) GoString() string {
	return s.String()
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAdapters(v []*DescribeAppVersionResponseBodyAppVersionAdapters) *DescribeAppVersionResponseBodyAppVersion {
	s.Adapters = v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetApkMd5(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.ApkMd5 = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAppId(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.AppId = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAppName(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.AppName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAppVersion(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetBlackVersionList(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.BlackVersionList = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetDownloadUrl(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetGmtCreate(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetGmtModify(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.GmtModify = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetId(v int64) *DescribeAppVersionResponseBodyAppVersion {
	s.Id = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetInstallType(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.InstallType = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsAllowNewInstall(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsAllowNewInstall = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsForceUpgrade(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsForceUpgrade = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsNeedRestart(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsNeedRestart = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsSilentUpgrade(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsSilentUpgrade = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetMd5(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Md5 = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetOriginalUrl(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.OriginalUrl = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetPackageName(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.PackageName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetReleaseNote(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRemark(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Remark = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRestartAppParam(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.RestartAppParam = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRestartAppType(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.RestartAppType = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRestartType(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.RestartType = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetSize(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Size = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetStatus(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Status = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetStatusName(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.StatusName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetVersionCode(v int64) *DescribeAppVersionResponseBodyAppVersion {
	s.VersionCode = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetWhiteVersionList(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.WhiteVersionList = &v
	return s
}

type DescribeAppVersionResponseBodyAppVersionAdapters struct {
	DeviceModelId   *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModelName *string `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxOsVersion    *string `json:"MaxOsVersion,omitempty" xml:"MaxOsVersion,omitempty"`
	MinOsVersion    *string `json:"MinOsVersion,omitempty" xml:"MinOsVersion,omitempty"`
	VersionId       *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeAppVersionResponseBodyAppVersionAdapters) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppVersionResponseBodyAppVersionAdapters) GoString() string {
	return s.String()
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetDeviceModelId(v string) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetDeviceModelName(v string) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.DeviceModelName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetId(v int64) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.Id = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetMaxOsVersion(v string) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.MaxOsVersion = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetMinOsVersion(v string) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.MinOsVersion = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetVersionId(v int64) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.VersionId = &v
	return s
}

type DescribeAppVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppVersionResponse) SetHeaders(v map[string]*string) *DescribeAppVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppVersionResponse) SetStatusCode(v int32) *DescribeAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppVersionResponse) SetBody(v *DescribeAppVersionResponseBody) *DescribeAppVersionResponse {
	s.Body = v
	return s
}

type DescribeAssistRTMPServerAddressRequest struct {
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeAssistRTMPServerAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistRTMPServerAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssistRTMPServerAddressRequest) SetDeviceId(v string) *DescribeAssistRTMPServerAddressRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressRequest) SetProjectId(v string) *DescribeAssistRTMPServerAddressRequest {
	s.ProjectId = &v
	return s
}

type DescribeAssistRTMPServerAddressResponseBody struct {
	OTP        *string `json:"OTP,omitempty" xml:"OTP,omitempty"`
	RTMPServer *string `json:"RTMPServer,omitempty" xml:"RTMPServer,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssistRTMPServerAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistRTMPServerAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssistRTMPServerAddressResponseBody) SetOTP(v string) *DescribeAssistRTMPServerAddressResponseBody {
	s.OTP = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressResponseBody) SetRTMPServer(v string) *DescribeAssistRTMPServerAddressResponseBody {
	s.RTMPServer = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressResponseBody) SetRequestId(v string) *DescribeAssistRTMPServerAddressResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAssistRTMPServerAddressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssistRTMPServerAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssistRTMPServerAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistRTMPServerAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssistRTMPServerAddressResponse) SetHeaders(v map[string]*string) *DescribeAssistRTMPServerAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssistRTMPServerAddressResponse) SetStatusCode(v int32) *DescribeAssistRTMPServerAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressResponse) SetBody(v *DescribeAssistRTMPServerAddressResponseBody) *DescribeAssistRTMPServerAddressResponse {
	s.Body = v
	return s
}

type DescribeAssistReportRequest struct {
	AssistId  *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeAssistReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssistReportRequest) SetAssistId(v string) *DescribeAssistReportRequest {
	s.AssistId = &v
	return s
}

func (s *DescribeAssistReportRequest) SetProjectId(v string) *DescribeAssistReportRequest {
	s.ProjectId = &v
	return s
}

type DescribeAssistReportResponseBody struct {
	AssistDescription *string `json:"AssistDescription,omitempty" xml:"AssistDescription,omitempty"`
	AssistId          *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	AssistReason      *string `json:"AssistReason,omitempty" xml:"AssistReason,omitempty"`
	AssistResult      *string `json:"AssistResult,omitempty" xml:"AssistResult,omitempty"`
	AssistTag         *string `json:"AssistTag,omitempty" xml:"AssistTag,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssistReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssistReportResponseBody) SetAssistDescription(v string) *DescribeAssistReportResponseBody {
	s.AssistDescription = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistId(v string) *DescribeAssistReportResponseBody {
	s.AssistId = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistReason(v string) *DescribeAssistReportResponseBody {
	s.AssistReason = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistResult(v string) *DescribeAssistReportResponseBody {
	s.AssistResult = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistTag(v string) *DescribeAssistReportResponseBody {
	s.AssistTag = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetRequestId(v string) *DescribeAssistReportResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAssistReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssistReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssistReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssistReportResponse) SetHeaders(v map[string]*string) *DescribeAssistReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssistReportResponse) SetStatusCode(v int32) *DescribeAssistReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssistReportResponse) SetBody(v *DescribeAssistReportResponseBody) *DescribeAssistReportResponse {
	s.Body = v
	return s
}

type DescribeAssistWSServerAddressRequest struct {
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeAssistWSServerAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistWSServerAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssistWSServerAddressRequest) SetDeviceId(v string) *DescribeAssistWSServerAddressRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeAssistWSServerAddressRequest) SetProjectId(v string) *DescribeAssistWSServerAddressRequest {
	s.ProjectId = &v
	return s
}

type DescribeAssistWSServerAddressResponseBody struct {
	OTP       *string `json:"OTP,omitempty" xml:"OTP,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WsServer  *string `json:"WsServer,omitempty" xml:"WsServer,omitempty"`
}

func (s DescribeAssistWSServerAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistWSServerAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssistWSServerAddressResponseBody) SetOTP(v string) *DescribeAssistWSServerAddressResponseBody {
	s.OTP = &v
	return s
}

func (s *DescribeAssistWSServerAddressResponseBody) SetRequestId(v string) *DescribeAssistWSServerAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssistWSServerAddressResponseBody) SetWsServer(v string) *DescribeAssistWSServerAddressResponseBody {
	s.WsServer = &v
	return s
}

type DescribeAssistWSServerAddressResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssistWSServerAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssistWSServerAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistWSServerAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssistWSServerAddressResponse) SetHeaders(v map[string]*string) *DescribeAssistWSServerAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssistWSServerAddressResponse) SetStatusCode(v int32) *DescribeAssistWSServerAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssistWSServerAddressResponse) SetBody(v *DescribeAssistWSServerAddressResponseBody) *DescribeAssistWSServerAddressResponse {
	s.Body = v
	return s
}

type DescribeCustomizedFilterRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedFilterRequest) SetId(v string) *DescribeCustomizedFilterRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizedFilterRequest) SetProjectId(v string) *DescribeCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

type DescribeCustomizedFilterResponseBody struct {
	CustomizedFilter *DescribeCustomizedFilterResponseBodyCustomizedFilter `json:"CustomizedFilter,omitempty" xml:"CustomizedFilter,omitempty" type:"Struct"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomizedFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizedFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedFilterResponseBody) SetCustomizedFilter(v *DescribeCustomizedFilterResponseBodyCustomizedFilter) *DescribeCustomizedFilterResponseBody {
	s.CustomizedFilter = v
	return s
}

func (s *DescribeCustomizedFilterResponseBody) SetRequestId(v string) *DescribeCustomizedFilterResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCustomizedFilterResponseBodyCustomizedFilter struct {
	BlackWhiteType     *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueCompareType   *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	ValueType          *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	VersionId          *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DescribeCustomizedFilterResponseBodyCustomizedFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizedFilterResponseBodyCustomizedFilter) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetBlackWhiteType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.BlackWhiteType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtCreate(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtCreateTimestamp(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtModify(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtModify = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtModifyTimestamp(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetId(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.Id = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetName(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.Name = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetValue(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.Value = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetValueCompareType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.ValueCompareType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetValueType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.ValueType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetVersionId(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.VersionId = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetVersionType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.VersionType = &v
	return s
}

type DescribeCustomizedFilterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomizedFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizedFilterResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedFilterResponse) SetHeaders(v map[string]*string) *DescribeCustomizedFilterResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizedFilterResponse) SetStatusCode(v int32) *DescribeCustomizedFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizedFilterResponse) SetBody(v *DescribeCustomizedFilterResponseBody) *DescribeCustomizedFilterResponse {
	s.Body = v
	return s
}

type DescribeDefaultSchemaRequest struct {
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeDefaultSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefaultSchemaRequest) SetDeviceModelId(v string) *DescribeDefaultSchemaRequest {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDefaultSchemaRequest) SetProjectId(v string) *DescribeDefaultSchemaRequest {
	s.ProjectId = &v
	return s
}

type DescribeDefaultSchemaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schema    *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeDefaultSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefaultSchemaResponseBody) SetRequestId(v string) *DescribeDefaultSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefaultSchemaResponseBody) SetSchema(v string) *DescribeDefaultSchemaResponseBody {
	s.Schema = &v
	return s
}

type DescribeDefaultSchemaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDefaultSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefaultSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefaultSchemaResponse) SetHeaders(v map[string]*string) *DescribeDefaultSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefaultSchemaResponse) SetStatusCode(v int32) *DescribeDefaultSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefaultSchemaResponse) SetBody(v *DescribeDefaultSchemaResponseBody) *DescribeDefaultSchemaResponse {
	s.Body = v
	return s
}

type DescribeDeviceBrandRequest struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeDeviceBrandRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandRequest) SetDeviceBrand(v string) *DescribeDeviceBrandRequest {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetDeviceBrandId(v int64) *DescribeDeviceBrandRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetLength(v string) *DescribeDeviceBrandRequest {
	s.Length = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetProjectId(v string) *DescribeDeviceBrandRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetStart(v string) *DescribeDeviceBrandRequest {
	s.Start = &v
	return s
}

type DescribeDeviceBrandResponseBody struct {
	DeviceBrand *DescribeDeviceBrandResponseBodyDeviceBrand `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceBrandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandResponseBody) SetDeviceBrand(v *DescribeDeviceBrandResponseBodyDeviceBrand) *DescribeDeviceBrandResponseBody {
	s.DeviceBrand = v
	return s
}

func (s *DescribeDeviceBrandResponseBody) SetRequestId(v string) *DescribeDeviceBrandResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceBrandResponseBodyDeviceBrand struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Manufacture   *string `json:"Manufacture,omitempty" xml:"Manufacture,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeDeviceBrandResponseBodyDeviceBrand) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandResponseBodyDeviceBrand) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetDescription(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.Description = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetDeviceBrand(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetDeviceBrandId(v int64) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.DeviceBrandId = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetManufacture(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.Manufacture = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetProjectId(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.ProjectId = &v
	return s
}

type DescribeDeviceBrandResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceBrandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceBrandResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandResponse) SetHeaders(v map[string]*string) *DescribeDeviceBrandResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceBrandResponse) SetStatusCode(v int32) *DescribeDeviceBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceBrandResponse) SetBody(v *DescribeDeviceBrandResponseBody) *DescribeDeviceBrandResponse {
	s.Body = v
	return s
}

type DescribeDeviceIdByOuterInfoRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	QueryValue *string `json:"QueryValue,omitempty" xml:"QueryValue,omitempty"`
}

func (s DescribeDeviceIdByOuterInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceIdByOuterInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceIdByOuterInfoRequest) SetProjectId(v string) *DescribeDeviceIdByOuterInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceIdByOuterInfoRequest) SetQueryType(v string) *DescribeDeviceIdByOuterInfoRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeDeviceIdByOuterInfoRequest) SetQueryValue(v string) *DescribeDeviceIdByOuterInfoRequest {
	s.QueryValue = &v
	return s
}

type DescribeDeviceIdByOuterInfoResponseBody struct {
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceIdByOuterInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceIdByOuterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceIdByOuterInfoResponseBody) SetDeviceId(v string) *DescribeDeviceIdByOuterInfoResponseBody {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceIdByOuterInfoResponseBody) SetRequestId(v string) *DescribeDeviceIdByOuterInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceIdByOuterInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceIdByOuterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceIdByOuterInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceIdByOuterInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceIdByOuterInfoResponse) SetHeaders(v map[string]*string) *DescribeDeviceIdByOuterInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceIdByOuterInfoResponse) SetStatusCode(v int32) *DescribeDeviceIdByOuterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceIdByOuterInfoResponse) SetBody(v *DescribeDeviceIdByOuterInfoResponseBody) *DescribeDeviceIdByOuterInfoResponse {
	s.Body = v
	return s
}

type DescribeDeviceInfoRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) SetDeviceId(v string) *DescribeDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetDeviceToken(v string) *DescribeDeviceInfoRequest {
	s.DeviceToken = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetProjectId(v string) *DescribeDeviceInfoRequest {
	s.ProjectId = &v
	return s
}

type DescribeDeviceInfoResponseBody struct {
	DeviceInfo *DescribeDeviceInfoResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBody) SetDeviceInfo(v *DescribeDeviceInfoResponseBodyDeviceInfo) *DescribeDeviceInfoResponseBody {
	s.DeviceInfo = v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetRequestId(v string) *DescribeDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceInfoResponseBodyDeviceInfo struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareId    *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	MacAddress    *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SoftwareId    *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsageType     *int32  `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	UsageTypeDesc *string `json:"UsageTypeDesc,omitempty" xml:"UsageTypeDesc,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Vin           *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceBrand(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceModel(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceModelId(v int64) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceType(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetHardwareId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.HardwareId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetMacAddress(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.MacAddress = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetName(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Name = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetProjectId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetRegion(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Region = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetSerialNumber(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetSoftwareId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.SoftwareId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetStatus(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Status = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetUsageType(v int32) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.UsageType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetUsageTypeDesc(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.UsageTypeDesc = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetUuid(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Uuid = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetVin(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Vin = &v
	return s
}

type DescribeDeviceInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceInfoResponse) SetStatusCode(v int32) *DescribeDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse {
	s.Body = v
	return s
}

type DescribeDeviceModelRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceModelRequest) SetDeviceModel(v string) *DescribeDeviceModelRequest {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceModelRequest) SetDeviceModelId(v int32) *DescribeDeviceModelRequest {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceModelRequest) SetProjectId(v string) *DescribeDeviceModelRequest {
	s.ProjectId = &v
	return s
}

type DescribeDeviceModelResponseBody struct {
	DeviceModel *DescribeDeviceModelResponseBodyDeviceModel `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceModelResponseBody) SetDeviceModel(v *DescribeDeviceModelResponseBodyDeviceModel) *DescribeDeviceModelResponseBody {
	s.DeviceModel = v
	return s
}

func (s *DescribeDeviceModelResponseBody) SetRequestId(v string) *DescribeDeviceModelResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceModelResponseBodyDeviceModel struct {
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceLogoUrl     *string `json:"DeviceLogoUrl,omitempty" xml:"DeviceLogoUrl,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
}

func (s DescribeDeviceModelResponseBodyDeviceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceModelResponseBodyDeviceModel) GoString() string {
	return s.String()
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetCanCreateDeviceId(v int32) *DescribeDeviceModelResponseBodyDeviceModel {
	s.CanCreateDeviceId = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDescription(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.Description = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceBrand(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceLogoUrl(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceLogoUrl = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceModel(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceModelId(v int64) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceName(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceName = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceType(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceType = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetHardwareType(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.HardwareType = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetInitUsageType(v int32) *DescribeDeviceModelResponseBodyDeviceModel {
	s.InitUsageType = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetInitUsageTypeDesc(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetObjectKey(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.ObjectKey = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetOsPlatform(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.OsPlatform = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetProjectId(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetSecurityChip(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.SecurityChip = &v
	return s
}

type DescribeDeviceModelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceModelResponse) SetHeaders(v map[string]*string) *DescribeDeviceModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceModelResponse) SetStatusCode(v int32) *DescribeDeviceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceModelResponse) SetBody(v *DescribeDeviceModelResponseBody) *DescribeDeviceModelResponse {
	s.Body = v
	return s
}

type DescribeDeviceOnlineInfoRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDeviceOnlineInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceOnlineInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceOnlineInfoRequest) SetProjectId(v string) *DescribeDeviceOnlineInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceOnlineInfoRequest) SetType(v string) *DescribeDeviceOnlineInfoRequest {
	s.Type = &v
	return s
}

func (s *DescribeDeviceOnlineInfoRequest) SetValue(v string) *DescribeDeviceOnlineInfoRequest {
	s.Value = &v
	return s
}

type DescribeDeviceOnlineInfoResponseBody struct {
	Devices   []*DescribeDeviceOnlineInfoResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceOnlineInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceOnlineInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceOnlineInfoResponseBody) SetDevices(v []*DescribeDeviceOnlineInfoResponseBodyDevices) *DescribeDeviceOnlineInfoResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBody) SetRequestId(v string) *DescribeDeviceOnlineInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceOnlineInfoResponseBodyDevices struct {
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IasId         *string `json:"IasId,omitempty" xml:"IasId,omitempty"`
	LoginTime     *int64  `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	Online        *int32  `json:"Online,omitempty" xml:"Online,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	Terminal      *string `json:"Terminal,omitempty" xml:"Terminal,omitempty"`
}

func (s DescribeDeviceOnlineInfoResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceOnlineInfoResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetClientVersion(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.ClientVersion = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetDeviceId(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetIasId(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.IasId = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetLoginTime(v int64) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.LoginTime = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetOnline(v int32) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.Online = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetProjectId(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetSystemVersion(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.SystemVersion = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetTerminal(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.Terminal = &v
	return s
}

type DescribeDeviceOnlineInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceOnlineInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceOnlineInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceOnlineInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceOnlineInfoResponse) SetHeaders(v map[string]*string) *DescribeDeviceOnlineInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceOnlineInfoResponse) SetStatusCode(v int32) *DescribeDeviceOnlineInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponse) SetBody(v *DescribeDeviceOnlineInfoResponseBody) *DescribeDeviceOnlineInfoResponse {
	s.Body = v
	return s
}

type DescribeDeviceShadowRequest struct {
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceModel    *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ViewSubscribed *bool   `json:"ViewSubscribed,omitempty" xml:"ViewSubscribed,omitempty"`
}

func (s DescribeDeviceShadowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowRequest) SetDeviceId(v string) *DescribeDeviceShadowRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceShadowRequest) SetDeviceModel(v string) *DescribeDeviceShadowRequest {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceShadowRequest) SetPath(v string) *DescribeDeviceShadowRequest {
	s.Path = &v
	return s
}

func (s *DescribeDeviceShadowRequest) SetProjectId(v string) *DescribeDeviceShadowRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceShadowRequest) SetViewSubscribed(v bool) *DescribeDeviceShadowRequest {
	s.ViewSubscribed = &v
	return s
}

type DescribeDeviceShadowResponseBody struct {
	DeviceShadow *DescribeDeviceShadowResponseBodyDeviceShadow `json:"DeviceShadow,omitempty" xml:"DeviceShadow,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceShadowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowResponseBody) SetDeviceShadow(v *DescribeDeviceShadowResponseBodyDeviceShadow) *DescribeDeviceShadowResponseBody {
	s.DeviceShadow = v
	return s
}

func (s *DescribeDeviceShadowResponseBody) SetRequestId(v string) *DescribeDeviceShadowResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceShadowResponseBodyDeviceShadow struct {
	DeviceInfo   *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	DeviceShadow *string `json:"DeviceShadow,omitempty" xml:"DeviceShadow,omitempty"`
}

func (s DescribeDeviceShadowResponseBodyDeviceShadow) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowResponseBodyDeviceShadow) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowResponseBodyDeviceShadow) SetDeviceInfo(v string) *DescribeDeviceShadowResponseBodyDeviceShadow {
	s.DeviceInfo = &v
	return s
}

func (s *DescribeDeviceShadowResponseBodyDeviceShadow) SetDeviceShadow(v string) *DescribeDeviceShadowResponseBodyDeviceShadow {
	s.DeviceShadow = &v
	return s
}

type DescribeDeviceShadowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceShadowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceShadowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowResponse) SetHeaders(v map[string]*string) *DescribeDeviceShadowResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceShadowResponse) SetStatusCode(v int32) *DescribeDeviceShadowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceShadowResponse) SetBody(v *DescribeDeviceShadowResponseBody) *DescribeDeviceShadowResponse {
	s.Body = v
	return s
}

type DescribeDeviceValiditySchemaRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s DescribeDeviceValiditySchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceValiditySchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceValiditySchemaRequest) SetDeviceModel(v string) *DescribeDeviceValiditySchemaRequest {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceValiditySchemaRequest) SetProjectId(v string) *DescribeDeviceValiditySchemaRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceValiditySchemaRequest) SetSchemaVersion(v string) *DescribeDeviceValiditySchemaRequest {
	s.SchemaVersion = &v
	return s
}

type DescribeDeviceValiditySchemaResponseBody struct {
	ItemList  []*DescribeDeviceValiditySchemaResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceValiditySchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceValiditySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceValiditySchemaResponseBody) SetItemList(v []*DescribeDeviceValiditySchemaResponseBodyItemList) *DescribeDeviceValiditySchemaResponseBody {
	s.ItemList = v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBody) SetRequestId(v string) *DescribeDeviceValiditySchemaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeviceValiditySchemaResponseBodyItemList struct {
	Description      *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EnumListStr      *string  `json:"EnumListStr,omitempty" xml:"EnumListStr,omitempty"`
	ExclusiveMaximum *bool    `json:"ExclusiveMaximum,omitempty" xml:"ExclusiveMaximum,omitempty"`
	ExclusiveMinimum *bool    `json:"ExclusiveMinimum,omitempty" xml:"ExclusiveMinimum,omitempty"`
	ItemType         *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	MaxLength        *int32   `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	Maximum          *float32 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	MinLength        *int32   `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	Minimum          *float32 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	Path             *string  `json:"Path,omitempty" xml:"Path,omitempty"`
	Required         *string  `json:"Required,omitempty" xml:"Required,omitempty"`
	Type             *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeviceValiditySchemaResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceValiditySchemaResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetDescription(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Description = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetEnumListStr(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.EnumListStr = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetExclusiveMaximum(v bool) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.ExclusiveMaximum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetExclusiveMinimum(v bool) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.ExclusiveMinimum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetItemType(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.ItemType = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMaxLength(v int32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.MaxLength = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMaximum(v float32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Maximum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMinLength(v int32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.MinLength = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMinimum(v float32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Minimum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetPath(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Path = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetRequired(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Required = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetType(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Type = &v
	return s
}

type DescribeDeviceValiditySchemaResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceValiditySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceValiditySchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceValiditySchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceValiditySchemaResponse) SetHeaders(v map[string]*string) *DescribeDeviceValiditySchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceValiditySchemaResponse) SetStatusCode(v int32) *DescribeDeviceValiditySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponse) SetBody(v *DescribeDeviceValiditySchemaResponseBody) *DescribeDeviceValiditySchemaResponse {
	s.Body = v
	return s
}

type DescribeMessageRequest struct {
	MessageId *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMessageRequest) SetMessageId(v int64) *DescribeMessageRequest {
	s.MessageId = &v
	return s
}

func (s *DescribeMessageRequest) SetProjectId(v string) *DescribeMessageRequest {
	s.ProjectId = &v
	return s
}

type DescribeMessageResponseBody struct {
	Message   *DescribeMessageResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMessageResponseBody) SetMessage(v *DescribeMessageResponseBodyMessage) *DescribeMessageResponseBody {
	s.Message = v
	return s
}

func (s *DescribeMessageResponseBody) SetRequestId(v string) *DescribeMessageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMessageResponseBodyMessage struct {
	AckCnt         *int32  `json:"AckCnt,omitempty" xml:"AckCnt,omitempty"`
	Action         *string `json:"Action,omitempty" xml:"Action,omitempty"`
	AppKey         *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Audit          *int32  `json:"Audit,omitempty" xml:"Audit,omitempty"`
	AuditMsg       *string `json:"AuditMsg,omitempty" xml:"AuditMsg,omitempty"`
	Desc           *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	ExipiredTime   *int64  `json:"ExipiredTime,omitempty" xml:"ExipiredTime,omitempty"`
	GmtCreateTime  *int64  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Parameter      *string `json:"Parameter,omitempty" xml:"Parameter,omitempty"`
	PredictSendCnt *int32  `json:"PredictSendCnt,omitempty" xml:"PredictSendCnt,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SendStatus     *int32  `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri            *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeMessageResponseBodyMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *DescribeMessageResponseBodyMessage) SetAckCnt(v int32) *DescribeMessageResponseBodyMessage {
	s.AckCnt = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAction(v string) *DescribeMessageResponseBodyMessage {
	s.Action = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAppKey(v string) *DescribeMessageResponseBodyMessage {
	s.AppKey = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAppName(v string) *DescribeMessageResponseBodyMessage {
	s.AppName = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAudit(v int32) *DescribeMessageResponseBodyMessage {
	s.Audit = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAuditMsg(v string) *DescribeMessageResponseBodyMessage {
	s.AuditMsg = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetDesc(v string) *DescribeMessageResponseBodyMessage {
	s.Desc = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetExipiredTime(v int64) *DescribeMessageResponseBodyMessage {
	s.ExipiredTime = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetGmtCreateTime(v int64) *DescribeMessageResponseBodyMessage {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetId(v int64) *DescribeMessageResponseBodyMessage {
	s.Id = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetParameter(v string) *DescribeMessageResponseBodyMessage {
	s.Parameter = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetPredictSendCnt(v int32) *DescribeMessageResponseBodyMessage {
	s.PredictSendCnt = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetProjectId(v string) *DescribeMessageResponseBodyMessage {
	s.ProjectId = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetSendStatus(v int32) *DescribeMessageResponseBodyMessage {
	s.SendStatus = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetTitle(v string) *DescribeMessageResponseBodyMessage {
	s.Title = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetType(v int32) *DescribeMessageResponseBodyMessage {
	s.Type = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetUri(v string) *DescribeMessageResponseBodyMessage {
	s.Uri = &v
	return s
}

type DescribeMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMessageResponse) SetHeaders(v map[string]*string) *DescribeMessageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMessageResponse) SetStatusCode(v int32) *DescribeMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMessageResponse) SetBody(v *DescribeMessageResponseBody) *DescribeMessageResponse {
	s.Body = v
	return s
}

type DescribeOpenAccountRequest struct {
	IdToken    *string `json:"IdToken,omitempty" xml:"IdToken,omitempty"`
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp        *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	OpenId     *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeOpenAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenAccountRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenAccountRequest) SetIdToken(v string) *DescribeOpenAccountRequest {
	s.IdToken = &v
	return s
}

func (s *DescribeOpenAccountRequest) SetIdentityId(v string) *DescribeOpenAccountRequest {
	s.IdentityId = &v
	return s
}

func (s *DescribeOpenAccountRequest) SetIdp(v string) *DescribeOpenAccountRequest {
	s.Idp = &v
	return s
}

func (s *DescribeOpenAccountRequest) SetOpenId(v string) *DescribeOpenAccountRequest {
	s.OpenId = &v
	return s
}

func (s *DescribeOpenAccountRequest) SetProjectId(v string) *DescribeOpenAccountRequest {
	s.ProjectId = &v
	return s
}

type DescribeOpenAccountResponseBody struct {
	OpenAccount *DescribeOpenAccountResponseBodyOpenAccount `json:"OpenAccount,omitempty" xml:"OpenAccount,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenAccountResponseBody) SetOpenAccount(v *DescribeOpenAccountResponseBodyOpenAccount) *DescribeOpenAccountResponseBody {
	s.OpenAccount = v
	return s
}

func (s *DescribeOpenAccountResponseBody) SetRequestId(v string) *DescribeOpenAccountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOpenAccountResponseBodyOpenAccount struct {
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CreateAccessKey *string `json:"CreateAccessKey,omitempty" xml:"CreateAccessKey,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	IdentityId      *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp             *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	LoginId         *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenId          *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeOpenAccountResponseBodyOpenAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenAccountResponseBodyOpenAccount) GoString() string {
	return s.String()
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetAliyunId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.AliyunId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetCreateAccessKey(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.CreateAccessKey = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetDisplayName(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.DisplayName = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetIdentityId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.IdentityId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetIdp(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Idp = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetLoginId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.LoginId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetMobile(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Mobile = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetOpenId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.OpenId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetRegion(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Region = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetStatus(v int32) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Status = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetType(v int32) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Type = &v
	return s
}

type DescribeOpenAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOpenAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOpenAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenAccountResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenAccountResponse) SetHeaders(v map[string]*string) *DescribeOpenAccountResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenAccountResponse) SetStatusCode(v int32) *DescribeOpenAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenAccountResponse) SetBody(v *DescribeOpenAccountResponseBody) *DescribeOpenAccountResponse {
	s.Body = v
	return s
}

type DescribeOsVersionRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionRequest) SetProjectId(v string) *DescribeOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeOsVersionRequest) SetVersionId(v int64) *DescribeOsVersionRequest {
	s.VersionId = &v
	return s
}

type DescribeOsVersionResponseBody struct {
	OsVersion *DescribeOsVersionResponseBodyOsVersion `json:"OsVersion,omitempty" xml:"OsVersion,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBody) SetOsVersion(v *DescribeOsVersionResponseBodyOsVersion) *DescribeOsVersionResponseBody {
	s.OsVersion = v
	return s
}

func (s *DescribeOsVersionResponseBody) SetRequestId(v string) *DescribeOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOsVersionResponseBodyOsVersion struct {
	BlackVersionList      *string                                                   `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	DeviceModelId         *string                                                   `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModelName       *string                                                   `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	EnableMobileDownload  *string                                                   `json:"EnableMobileDownload,omitempty" xml:"EnableMobileDownload,omitempty"`
	GmtCreate             *string                                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify             *string                                                   `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Id                    *int64                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	IsForceNightUpgrade   *string                                                   `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	IsForceUpgrade        *string                                                   `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsMilestone           *string                                                   `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	MaxClientVersion      *string                                                   `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	MinClientVersion      *string                                                   `json:"MinClientVersion,omitempty" xml:"MinClientVersion,omitempty"`
	MobileDownloadMaxSize *string                                                   `json:"MobileDownloadMaxSize,omitempty" xml:"MobileDownloadMaxSize,omitempty"`
	NightUpgradeOption    *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption `json:"NightUpgradeOption,omitempty" xml:"NightUpgradeOption,omitempty" type:"Struct"`
	ReleaseNote           *string                                                   `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark                *string                                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RomList               []*DescribeOsVersionResponseBodyOsVersionRomList          `json:"RomList,omitempty" xml:"RomList,omitempty" type:"Repeated"`
	Status                *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName            *string                                                   `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	SystemVersion         *string                                                   `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	WhiteVersionList      *string                                                   `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
}

func (s DescribeOsVersionResponseBodyOsVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBodyOsVersion) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetBlackVersionList(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.BlackVersionList = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetDeviceModelId(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetDeviceModelName(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.DeviceModelName = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetEnableMobileDownload(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.EnableMobileDownload = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetGmtCreate(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetGmtModify(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.GmtModify = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetId(v int64) *DescribeOsVersionResponseBodyOsVersion {
	s.Id = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetIsForceNightUpgrade(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetIsForceUpgrade(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.IsForceUpgrade = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetIsMilestone(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.IsMilestone = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetMaxClientVersion(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.MaxClientVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetMinClientVersion(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.MinClientVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetMobileDownloadMaxSize(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.MobileDownloadMaxSize = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetNightUpgradeOption(v *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) *DescribeOsVersionResponseBodyOsVersion {
	s.NightUpgradeOption = v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetReleaseNote(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetRemark(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.Remark = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetRomList(v []*DescribeOsVersionResponseBodyOsVersionRomList) *DescribeOsVersionResponseBodyOsVersion {
	s.RomList = v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetStatus(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.Status = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetStatusName(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.StatusName = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetSystemVersion(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.SystemVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetWhiteVersionList(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.WhiteVersionList = &v
	return s
}

type DescribeOsVersionResponseBodyOsVersionNightUpgradeOption struct {
	DownloadType    *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	IsAllowedCancel *string `json:"IsAllowedCancel,omitempty" xml:"IsAllowedCancel,omitempty"`
	IsShowTip       *string `json:"IsShowTip,omitempty" xml:"IsShowTip,omitempty"`
}

func (s DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) SetDownloadType(v string) *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption {
	s.DownloadType = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) SetIsAllowedCancel(v string) *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption {
	s.IsAllowedCancel = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) SetIsShowTip(v string) *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption {
	s.IsShowTip = &v
	return s
}

type DescribeOsVersionResponseBodyOsVersionRomList struct {
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify   *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	OriginalUrl *string `json:"OriginalUrl,omitempty" xml:"OriginalUrl,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	SplitNum    *string `json:"SplitNum,omitempty" xml:"SplitNum,omitempty"`
	VersionId   *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeOsVersionResponseBodyOsVersionRomList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBodyOsVersionRomList) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetBaseVersion(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.BaseVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetDownloadUrl(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetGmtCreate(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetGmtModify(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.GmtModify = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetId(v int64) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.Id = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetMd5(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.Md5 = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetOriginalUrl(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.OriginalUrl = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetSize(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.Size = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetSplitNum(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.SplitNum = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetVersionId(v int64) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.VersionId = &v
	return s
}

type DescribeOsVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponse) SetHeaders(v map[string]*string) *DescribeOsVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeOsVersionResponse) SetStatusCode(v int32) *DescribeOsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOsVersionResponse) SetBody(v *DescribeOsVersionResponseBody) *DescribeOsVersionResponse {
	s.Body = v
	return s
}

type DescribeProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectRequest) SetProjectId(v string) *DescribeProjectRequest {
	s.ProjectId = &v
	return s
}

type DescribeProjectResponseBody struct {
	Project   *DescribeProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBody) SetProject(v *DescribeProjectResponseBodyProject) *DescribeProjectResponseBody {
	s.Project = v
	return s
}

func (s *DescribeProjectResponseBody) SetRequestId(v string) *DescribeProjectResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProjectResponseBodyProject struct {
	Creator                      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate                    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified                  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId                    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status                       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId                       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VehicleCommunicationProtocol *string `json:"VehicleCommunicationProtocol,omitempty" xml:"VehicleCommunicationProtocol,omitempty"`
}

func (s DescribeProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBodyProject) SetCreator(v string) *DescribeProjectResponseBodyProject {
	s.Creator = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetDescription(v string) *DescribeProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetGmtCreate(v int64) *DescribeProjectResponseBodyProject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetGmtModified(v int64) *DescribeProjectResponseBodyProject {
	s.GmtModified = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetId(v int64) *DescribeProjectResponseBodyProject {
	s.Id = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetName(v string) *DescribeProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetProjectId(v string) *DescribeProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetStatus(v int32) *DescribeProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetUserId(v string) *DescribeProjectResponseBodyProject {
	s.UserId = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetVehicleCommunicationProtocol(v string) *DescribeProjectResponseBodyProject {
	s.VehicleCommunicationProtocol = &v
	return s
}

type DescribeProjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponse) SetHeaders(v map[string]*string) *DescribeProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectResponse) SetStatusCode(v int32) *DescribeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectResponse) SetBody(v *DescribeProjectResponseBody) *DescribeProjectResponse {
	s.Body = v
	return s
}

type DescribeProjectAppSecurityRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeProjectAppSecurityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectAppSecurityRequest) SetAppId(v string) *DescribeProjectAppSecurityRequest {
	s.AppId = &v
	return s
}

func (s *DescribeProjectAppSecurityRequest) SetProjectId(v string) *DescribeProjectAppSecurityRequest {
	s.ProjectId = &v
	return s
}

type DescribeProjectAppSecurityResponseBody struct {
	ProjectAppSecurity *DescribeProjectAppSecurityResponseBodyProjectAppSecurity `json:"ProjectAppSecurity,omitempty" xml:"ProjectAppSecurity,omitempty" type:"Struct"`
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProjectAppSecurityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAppSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectAppSecurityResponseBody) SetProjectAppSecurity(v *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) *DescribeProjectAppSecurityResponseBody {
	s.ProjectAppSecurity = v
	return s
}

func (s *DescribeProjectAppSecurityResponseBody) SetRequestId(v string) *DescribeProjectAppSecurityResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProjectAppSecurityResponseBodyProjectAppSecurity struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppSecret   *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeProjectAppSecurityResponseBodyProjectAppSecurity) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAppSecurityResponseBodyProjectAppSecurity) GoString() string {
	return s.String()
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetAppId(v string) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.AppId = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetAppKey(v string) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.AppKey = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetAppSecret(v string) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.AppSecret = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetGmtCreate(v int64) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetGmtModified(v int64) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.GmtModified = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetId(v int64) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.Id = &v
	return s
}

type DescribeProjectAppSecurityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectAppSecurityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAppSecurityResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectAppSecurityResponse) SetHeaders(v map[string]*string) *DescribeProjectAppSecurityResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectAppSecurityResponse) SetStatusCode(v int32) *DescribeProjectAppSecurityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectAppSecurityResponse) SetBody(v *DescribeProjectAppSecurityResponseBody) *DescribeProjectAppSecurityResponse {
	s.Body = v
	return s
}

type DescribeShadowSchemaRequest struct {
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	IsSimple    *bool   `json:"IsSimple,omitempty" xml:"IsSimple,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeShadowSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeShadowSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeShadowSchemaRequest) SetDeviceModel(v string) *DescribeShadowSchemaRequest {
	s.DeviceModel = &v
	return s
}

func (s *DescribeShadowSchemaRequest) SetIsSimple(v bool) *DescribeShadowSchemaRequest {
	s.IsSimple = &v
	return s
}

func (s *DescribeShadowSchemaRequest) SetProjectId(v string) *DescribeShadowSchemaRequest {
	s.ProjectId = &v
	return s
}

type DescribeShadowSchemaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schema    *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeShadowSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeShadowSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShadowSchemaResponseBody) SetRequestId(v string) *DescribeShadowSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShadowSchemaResponseBody) SetSchema(v string) *DescribeShadowSchemaResponseBody {
	s.Schema = &v
	return s
}

type DescribeShadowSchemaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeShadowSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeShadowSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeShadowSchemaResponse) SetHeaders(v map[string]*string) *DescribeShadowSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeShadowSchemaResponse) SetStatusCode(v int32) *DescribeShadowSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShadowSchemaResponse) SetBody(v *DescribeShadowSchemaResponseBody) *DescribeShadowSchemaResponse {
	s.Body = v
	return s
}

type DescribeVersionDeviceGroupRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupRequest) SetId(v string) *DescribeVersionDeviceGroupRequest {
	s.Id = &v
	return s
}

func (s *DescribeVersionDeviceGroupRequest) SetProjectId(v string) *DescribeVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

type DescribeVersionDeviceGroupResponseBody struct {
	DeviceGroup *DescribeVersionDeviceGroupResponseBodyDeviceGroup `json:"DeviceGroup,omitempty" xml:"DeviceGroup,omitempty" type:"Struct"`
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVersionDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupResponseBody) SetDeviceGroup(v *DescribeVersionDeviceGroupResponseBodyDeviceGroup) *DescribeVersionDeviceGroupResponseBody {
	s.DeviceGroup = v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBody) SetRequestId(v string) *DescribeVersionDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVersionDeviceGroupResponseBodyDeviceGroup struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify   *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxCount    *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeVersionDeviceGroupResponseBodyDeviceGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupResponseBodyDeviceGroup) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetDescription(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.Description = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetGmtCreate(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetGmtModify(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.GmtModify = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetId(v int64) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.Id = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetMaxCount(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.MaxCount = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetName(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.Name = &v
	return s
}

type DescribeVersionDeviceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVersionDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupResponse) SetHeaders(v map[string]*string) *DescribeVersionDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeVersionDeviceGroupResponse) SetStatusCode(v int32) *DescribeVersionDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponse) SetBody(v *DescribeVersionDeviceGroupResponseBody) *DescribeVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type DiagnosisVersionRequest struct {
	DiagnoseStyle *string `json:"DiagnoseStyle,omitempty" xml:"DiagnoseStyle,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IdType        *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId    *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType   *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DiagnosisVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DiagnosisVersionRequest) GoString() string {
	return s.String()
}

func (s *DiagnosisVersionRequest) SetDiagnoseStyle(v string) *DiagnosisVersionRequest {
	s.DiagnoseStyle = &v
	return s
}

func (s *DiagnosisVersionRequest) SetEndTime(v string) *DiagnosisVersionRequest {
	s.EndTime = &v
	return s
}

func (s *DiagnosisVersionRequest) SetIdType(v string) *DiagnosisVersionRequest {
	s.IdType = &v
	return s
}

func (s *DiagnosisVersionRequest) SetOriginalId(v string) *DiagnosisVersionRequest {
	s.OriginalId = &v
	return s
}

func (s *DiagnosisVersionRequest) SetProjectId(v string) *DiagnosisVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *DiagnosisVersionRequest) SetStartTime(v string) *DiagnosisVersionRequest {
	s.StartTime = &v
	return s
}

func (s *DiagnosisVersionRequest) SetVersionId(v string) *DiagnosisVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DiagnosisVersionRequest) SetVersionType(v string) *DiagnosisVersionRequest {
	s.VersionType = &v
	return s
}

type DiagnosisVersionResponseBody struct {
	DiagnosisResult *string `json:"DiagnosisResult,omitempty" xml:"DiagnosisResult,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DiagnosisVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DiagnosisVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnosisVersionResponseBody) SetDiagnosisResult(v string) *DiagnosisVersionResponseBody {
	s.DiagnosisResult = &v
	return s
}

func (s *DiagnosisVersionResponseBody) SetRequestId(v string) *DiagnosisVersionResponseBody {
	s.RequestId = &v
	return s
}

type DiagnosisVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DiagnosisVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DiagnosisVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DiagnosisVersionResponse) GoString() string {
	return s.String()
}

func (s *DiagnosisVersionResponse) SetHeaders(v map[string]*string) *DiagnosisVersionResponse {
	s.Headers = v
	return s
}

func (s *DiagnosisVersionResponse) SetStatusCode(v int32) *DiagnosisVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DiagnosisVersionResponse) SetBody(v *DiagnosisVersionResponseBody) *DiagnosisVersionResponse {
	s.Body = v
	return s
}

type FindAppVersionsRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s FindAppVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *FindAppVersionsRequest) SetAppId(v string) *FindAppVersionsRequest {
	s.AppId = &v
	return s
}

func (s *FindAppVersionsRequest) SetDeviceModelId(v string) *FindAppVersionsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *FindAppVersionsRequest) SetPageIndex(v int32) *FindAppVersionsRequest {
	s.PageIndex = &v
	return s
}

func (s *FindAppVersionsRequest) SetPageSize(v int32) *FindAppVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *FindAppVersionsRequest) SetProjectId(v string) *FindAppVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *FindAppVersionsRequest) SetRemark(v string) *FindAppVersionsRequest {
	s.Remark = &v
	return s
}

func (s *FindAppVersionsRequest) SetStatus(v string) *FindAppVersionsRequest {
	s.Status = &v
	return s
}

func (s *FindAppVersionsRequest) SetVersionId(v string) *FindAppVersionsRequest {
	s.VersionId = &v
	return s
}

type FindAppVersionsResponseBody struct {
	AppVersionList *FindAppVersionsResponseBodyAppVersionList `json:"AppVersionList,omitempty" xml:"AppVersionList,omitempty" type:"Struct"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindAppVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *FindAppVersionsResponseBody) SetAppVersionList(v *FindAppVersionsResponseBodyAppVersionList) *FindAppVersionsResponseBody {
	s.AppVersionList = v
	return s
}

func (s *FindAppVersionsResponseBody) SetRequestId(v string) *FindAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

type FindAppVersionsResponseBodyAppVersionList struct {
	Items      []*FindAppVersionsResponseBodyAppVersionListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindAppVersionsResponseBodyAppVersionList) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsResponseBodyAppVersionList) GoString() string {
	return s.String()
}

func (s *FindAppVersionsResponseBodyAppVersionList) SetItems(v []*FindAppVersionsResponseBodyAppVersionListItems) *FindAppVersionsResponseBodyAppVersionList {
	s.Items = v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionList) SetTotalCount(v int32) *FindAppVersionsResponseBodyAppVersionList {
	s.TotalCount = &v
	return s
}

type FindAppVersionsResponseBodyAppVersionListItems struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPackageName     *string `json:"AppPackageName,omitempty" xml:"AppPackageName,omitempty"`
	AppVersion         *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallType        *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	IsAllowNewInstall  *string `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	IsForceUpgrade     *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsNeedRestart      *string `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	IsSilentUpgrade    *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	RestartAppParam    *string `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	RestartAppType     *string `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	RestartType        *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName         *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	VersionCode        *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s FindAppVersionsResponseBodyAppVersionListItems) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsResponseBodyAppVersionListItems) GoString() string {
	return s.String()
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppId(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppId = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppName(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppName = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppPackageName(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppPackageName = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppVersion(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppVersion = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetGmtCreate(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetGmtCreateTimestamp(v int64) *FindAppVersionsResponseBodyAppVersionListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetGmtModify(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.GmtModify = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetGmtModifyTimestamp(v int64) *FindAppVersionsResponseBodyAppVersionListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetId(v int64) *FindAppVersionsResponseBodyAppVersionListItems {
	s.Id = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetInstallType(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.InstallType = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsAllowNewInstall(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsAllowNewInstall = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsForceUpgrade(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsForceUpgrade = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsNeedRestart(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsNeedRestart = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsSilentUpgrade(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsSilentUpgrade = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetRestartAppParam(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.RestartAppParam = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetRestartAppType(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.RestartAppType = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetRestartType(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.RestartType = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetStatus(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.Status = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetStatusName(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.StatusName = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetVersionCode(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.VersionCode = &v
	return s
}

type FindAppVersionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindAppVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *FindAppVersionsResponse) SetHeaders(v map[string]*string) *FindAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *FindAppVersionsResponse) SetStatusCode(v int32) *FindAppVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *FindAppVersionsResponse) SetBody(v *FindAppVersionsResponseBody) *FindAppVersionsResponse {
	s.Body = v
	return s
}

type FindCustomizedFiltersRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindCustomizedFiltersRequest) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersRequest) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersRequest) SetName(v string) *FindCustomizedFiltersRequest {
	s.Name = &v
	return s
}

func (s *FindCustomizedFiltersRequest) SetPageIndex(v int32) *FindCustomizedFiltersRequest {
	s.PageIndex = &v
	return s
}

func (s *FindCustomizedFiltersRequest) SetPageSize(v int32) *FindCustomizedFiltersRequest {
	s.PageSize = &v
	return s
}

func (s *FindCustomizedFiltersRequest) SetProjectId(v string) *FindCustomizedFiltersRequest {
	s.ProjectId = &v
	return s
}

func (s *FindCustomizedFiltersRequest) SetVersionId(v string) *FindCustomizedFiltersRequest {
	s.VersionId = &v
	return s
}

func (s *FindCustomizedFiltersRequest) SetVersionType(v string) *FindCustomizedFiltersRequest {
	s.VersionType = &v
	return s
}

type FindCustomizedFiltersResponseBody struct {
	CustomizedFilterList *FindCustomizedFiltersResponseBodyCustomizedFilterList `json:"CustomizedFilterList,omitempty" xml:"CustomizedFilterList,omitempty" type:"Struct"`
	RequestId            *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindCustomizedFiltersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersResponseBody) SetCustomizedFilterList(v *FindCustomizedFiltersResponseBodyCustomizedFilterList) *FindCustomizedFiltersResponseBody {
	s.CustomizedFilterList = v
	return s
}

func (s *FindCustomizedFiltersResponseBody) SetRequestId(v string) *FindCustomizedFiltersResponseBody {
	s.RequestId = &v
	return s
}

type FindCustomizedFiltersResponseBodyCustomizedFilterList struct {
	Items      []*FindCustomizedFiltersResponseBodyCustomizedFilterListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindCustomizedFiltersResponseBodyCustomizedFilterList) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersResponseBodyCustomizedFilterList) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterList) SetItems(v []*FindCustomizedFiltersResponseBodyCustomizedFilterListItems) *FindCustomizedFiltersResponseBodyCustomizedFilterList {
	s.Items = v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterList) SetTotalCount(v int32) *FindCustomizedFiltersResponseBodyCustomizedFilterList {
	s.TotalCount = &v
	return s
}

type FindCustomizedFiltersResponseBodyCustomizedFilterListItems struct {
	BlackWhiteType     *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueCompareType   *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
}

func (s FindCustomizedFiltersResponseBodyCustomizedFilterListItems) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersResponseBodyCustomizedFilterListItems) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetBlackWhiteType(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.BlackWhiteType = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtCreate(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtCreateTimestamp(v int64) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtModify(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtModify = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtModifyTimestamp(v int64) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetId(v int64) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.Id = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetName(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.Name = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetValue(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.Value = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetValueCompareType(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.ValueCompareType = &v
	return s
}

type FindCustomizedFiltersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindCustomizedFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindCustomizedFiltersResponse) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersResponse) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersResponse) SetHeaders(v map[string]*string) *FindCustomizedFiltersResponse {
	s.Headers = v
	return s
}

func (s *FindCustomizedFiltersResponse) SetStatusCode(v int32) *FindCustomizedFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *FindCustomizedFiltersResponse) SetBody(v *FindCustomizedFiltersResponseBody) *FindCustomizedFiltersResponse {
	s.Body = v
	return s
}

type FindCustomizedPropertiesRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindCustomizedPropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesRequest) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesRequest) SetName(v string) *FindCustomizedPropertiesRequest {
	s.Name = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetPageIndex(v int32) *FindCustomizedPropertiesRequest {
	s.PageIndex = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetPageSize(v int32) *FindCustomizedPropertiesRequest {
	s.PageSize = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetProjectId(v string) *FindCustomizedPropertiesRequest {
	s.ProjectId = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetVersionId(v string) *FindCustomizedPropertiesRequest {
	s.VersionId = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetVersionType(v string) *FindCustomizedPropertiesRequest {
	s.VersionType = &v
	return s
}

type FindCustomizedPropertiesResponseBody struct {
	CustomizedPropertyList *FindCustomizedPropertiesResponseBodyCustomizedPropertyList `json:"CustomizedPropertyList,omitempty" xml:"CustomizedPropertyList,omitempty" type:"Struct"`
	RequestId              *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindCustomizedPropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesResponseBody) SetCustomizedPropertyList(v *FindCustomizedPropertiesResponseBodyCustomizedPropertyList) *FindCustomizedPropertiesResponseBody {
	s.CustomizedPropertyList = v
	return s
}

func (s *FindCustomizedPropertiesResponseBody) SetRequestId(v string) *FindCustomizedPropertiesResponseBody {
	s.RequestId = &v
	return s
}

type FindCustomizedPropertiesResponseBodyCustomizedPropertyList struct {
	Items      []*FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindCustomizedPropertiesResponseBodyCustomizedPropertyList) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesResponseBodyCustomizedPropertyList) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyList) SetItems(v []*FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) *FindCustomizedPropertiesResponseBodyCustomizedPropertyList {
	s.Items = v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyList) SetTotalCount(v int32) *FindCustomizedPropertiesResponseBodyCustomizedPropertyList {
	s.TotalCount = &v
	return s
}

type FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems struct {
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetGmtCreate(v string) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetGmtCreateTimestamp(v int64) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetId(v int64) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.Id = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetName(v string) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.Name = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetValue(v string) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.Value = &v
	return s
}

type FindCustomizedPropertiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindCustomizedPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindCustomizedPropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesResponse) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesResponse) SetHeaders(v map[string]*string) *FindCustomizedPropertiesResponse {
	s.Headers = v
	return s
}

func (s *FindCustomizedPropertiesResponse) SetStatusCode(v int32) *FindCustomizedPropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindCustomizedPropertiesResponse) SetBody(v *FindCustomizedPropertiesResponseBody) *FindCustomizedPropertiesResponse {
	s.Body = v
	return s
}

type FindOsVersionsRequest struct {
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	IsMilestone   *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s FindOsVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsRequest) GoString() string {
	return s.String()
}

func (s *FindOsVersionsRequest) SetDeviceModelId(v string) *FindOsVersionsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *FindOsVersionsRequest) SetIsMilestone(v string) *FindOsVersionsRequest {
	s.IsMilestone = &v
	return s
}

func (s *FindOsVersionsRequest) SetPageIndex(v int32) *FindOsVersionsRequest {
	s.PageIndex = &v
	return s
}

func (s *FindOsVersionsRequest) SetPageSize(v int32) *FindOsVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *FindOsVersionsRequest) SetProjectId(v string) *FindOsVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *FindOsVersionsRequest) SetRemark(v string) *FindOsVersionsRequest {
	s.Remark = &v
	return s
}

func (s *FindOsVersionsRequest) SetStatus(v string) *FindOsVersionsRequest {
	s.Status = &v
	return s
}

func (s *FindOsVersionsRequest) SetSystemVersion(v string) *FindOsVersionsRequest {
	s.SystemVersion = &v
	return s
}

func (s *FindOsVersionsRequest) SetVersionId(v string) *FindOsVersionsRequest {
	s.VersionId = &v
	return s
}

type FindOsVersionsResponseBody struct {
	OsVersionList *FindOsVersionsResponseBodyOsVersionList `json:"OsVersionList,omitempty" xml:"OsVersionList,omitempty" type:"Struct"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindOsVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *FindOsVersionsResponseBody) SetOsVersionList(v *FindOsVersionsResponseBodyOsVersionList) *FindOsVersionsResponseBody {
	s.OsVersionList = v
	return s
}

func (s *FindOsVersionsResponseBody) SetRequestId(v string) *FindOsVersionsResponseBody {
	s.RequestId = &v
	return s
}

type FindOsVersionsResponseBodyOsVersionList struct {
	Items      []*FindOsVersionsResponseBodyOsVersionListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindOsVersionsResponseBodyOsVersionList) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsResponseBodyOsVersionList) GoString() string {
	return s.String()
}

func (s *FindOsVersionsResponseBodyOsVersionList) SetItems(v []*FindOsVersionsResponseBodyOsVersionListItems) *FindOsVersionsResponseBodyOsVersionList {
	s.Items = v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionList) SetTotalCount(v int32) *FindOsVersionsResponseBodyOsVersionList {
	s.TotalCount = &v
	return s
}

type FindOsVersionsResponseBodyOsVersionListItems struct {
	DeviceModelId       *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModelName     *string `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp  *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify           *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp  *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsForceNightUpgrade *string `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	IsForceReboot       *string `json:"IsForceReboot,omitempty" xml:"IsForceReboot,omitempty"`
	IsForceUpgrade      *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsMilestone         *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	IsSilentUpgrade     *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	Remark              *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName          *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	SystemVersion       *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
}

func (s FindOsVersionsResponseBodyOsVersionListItems) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsResponseBodyOsVersionListItems) GoString() string {
	return s.String()
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetDeviceModelId(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.DeviceModelId = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetDeviceModelName(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.DeviceModelName = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtCreate(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtCreateTimestamp(v int64) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtModify(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtModify = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtModifyTimestamp(v int64) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetId(v int64) *FindOsVersionsResponseBodyOsVersionListItems {
	s.Id = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsForceNightUpgrade(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsForceReboot(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsForceReboot = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsForceUpgrade(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsForceUpgrade = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsMilestone(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsMilestone = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsSilentUpgrade(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsSilentUpgrade = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetRemark(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.Remark = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetStatus(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.Status = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetStatusName(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.StatusName = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetSystemVersion(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.SystemVersion = &v
	return s
}

type FindOsVersionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindOsVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindOsVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsResponse) GoString() string {
	return s.String()
}

func (s *FindOsVersionsResponse) SetHeaders(v map[string]*string) *FindOsVersionsResponse {
	s.Headers = v
	return s
}

func (s *FindOsVersionsResponse) SetStatusCode(v int32) *FindOsVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *FindOsVersionsResponse) SetBody(v *FindOsVersionsResponseBody) *FindOsVersionsResponse {
	s.Body = v
	return s
}

type FindPrepublishPassedDevicesRequest struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s FindPrepublishPassedDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesRequest) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesRequest) SetDeviceId(v string) *FindPrepublishPassedDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *FindPrepublishPassedDevicesRequest) SetPageIndex(v int32) *FindPrepublishPassedDevicesRequest {
	s.PageIndex = &v
	return s
}

func (s *FindPrepublishPassedDevicesRequest) SetPageSize(v int32) *FindPrepublishPassedDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *FindPrepublishPassedDevicesRequest) SetPrepublishId(v string) *FindPrepublishPassedDevicesRequest {
	s.PrepublishId = &v
	return s
}

func (s *FindPrepublishPassedDevicesRequest) SetProjectId(v string) *FindPrepublishPassedDevicesRequest {
	s.ProjectId = &v
	return s
}

type FindPrepublishPassedDevicesResponseBody struct {
	DeviceList *FindPrepublishPassedDevicesResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Struct"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindPrepublishPassedDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesResponseBody) SetDeviceList(v *FindPrepublishPassedDevicesResponseBodyDeviceList) *FindPrepublishPassedDevicesResponseBody {
	s.DeviceList = v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBody) SetRequestId(v string) *FindPrepublishPassedDevicesResponseBody {
	s.RequestId = &v
	return s
}

type FindPrepublishPassedDevicesResponseBodyDeviceList struct {
	Items      []*FindPrepublishPassedDevicesResponseBodyDeviceListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindPrepublishPassedDevicesResponseBodyDeviceList) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceList) SetItems(v []*FindPrepublishPassedDevicesResponseBodyDeviceListItems) *FindPrepublishPassedDevicesResponseBodyDeviceList {
	s.Items = v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceList) SetTotalCount(v int32) *FindPrepublishPassedDevicesResponseBodyDeviceList {
	s.TotalCount = &v
	return s
}

type FindPrepublishPassedDevicesResponseBodyDeviceListItems struct {
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
}

func (s FindPrepublishPassedDevicesResponseBodyDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesResponseBodyDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceListItems) SetDeviceId(v string) *FindPrepublishPassedDevicesResponseBodyDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceListItems) SetGmtCreate(v string) *FindPrepublishPassedDevicesResponseBodyDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceListItems) SetGmtCreateTimestamp(v int64) *FindPrepublishPassedDevicesResponseBodyDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

type FindPrepublishPassedDevicesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindPrepublishPassedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindPrepublishPassedDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesResponse) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesResponse) SetHeaders(v map[string]*string) *FindPrepublishPassedDevicesResponse {
	s.Headers = v
	return s
}

func (s *FindPrepublishPassedDevicesResponse) SetStatusCode(v int32) *FindPrepublishPassedDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindPrepublishPassedDevicesResponse) SetBody(v *FindPrepublishPassedDevicesResponseBody) *FindPrepublishPassedDevicesResponse {
	s.Body = v
	return s
}

type FindPrepublishesByParentIdRequest struct {
	ParentId  *int32  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s FindPrepublishesByParentIdRequest) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdRequest) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdRequest) SetParentId(v int32) *FindPrepublishesByParentIdRequest {
	s.ParentId = &v
	return s
}

func (s *FindPrepublishesByParentIdRequest) SetProjectId(v string) *FindPrepublishesByParentIdRequest {
	s.ProjectId = &v
	return s
}

type FindPrepublishesByParentIdResponseBody struct {
	PrepublishList *FindPrepublishesByParentIdResponseBodyPrepublishList `json:"PrepublishList,omitempty" xml:"PrepublishList,omitempty" type:"Struct"`
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindPrepublishesByParentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdResponseBody) SetPrepublishList(v *FindPrepublishesByParentIdResponseBodyPrepublishList) *FindPrepublishesByParentIdResponseBody {
	s.PrepublishList = v
	return s
}

func (s *FindPrepublishesByParentIdResponseBody) SetRequestId(v string) *FindPrepublishesByParentIdResponseBody {
	s.RequestId = &v
	return s
}

type FindPrepublishesByParentIdResponseBodyPrepublishList struct {
	Items      []*FindPrepublishesByParentIdResponseBodyPrepublishListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindPrepublishesByParentIdResponseBodyPrepublishList) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdResponseBodyPrepublishList) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishList) SetItems(v []*FindPrepublishesByParentIdResponseBodyPrepublishListItems) *FindPrepublishesByParentIdResponseBodyPrepublishList {
	s.Items = v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishList) SetTotalCount(v int32) *FindPrepublishesByParentIdResponseBodyPrepublishList {
	s.TotalCount = &v
	return s
}

type FindPrepublishesByParentIdResponseBodyPrepublishListItems struct {
	BarrierCount       *string `json:"BarrierCount,omitempty" xml:"BarrierCount,omitempty"`
	DeviceModelId      *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsActive           *string `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	IsTotalPrepublish  *string `json:"IsTotalPrepublish,omitempty" xml:"IsTotalPrepublish,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId           *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindPrepublishesByParentIdResponseBodyPrepublishListItems) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdResponseBodyPrepublishListItems) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetBarrierCount(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.BarrierCount = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetDeviceModelId(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.DeviceModelId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtCreate(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtCreateTimestamp(v int64) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtModify(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtModify = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtModifyTimestamp(v int64) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetId(v int64) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.Id = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetIsActive(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.IsActive = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetIsTotalPrepublish(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.IsTotalPrepublish = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetName(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.Name = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetParentId(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.ParentId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetVersionId(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.VersionId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetVersionType(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.VersionType = &v
	return s
}

type FindPrepublishesByParentIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindPrepublishesByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindPrepublishesByParentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdResponse) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdResponse) SetHeaders(v map[string]*string) *FindPrepublishesByParentIdResponse {
	s.Headers = v
	return s
}

func (s *FindPrepublishesByParentIdResponse) SetStatusCode(v int32) *FindPrepublishesByParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *FindPrepublishesByParentIdResponse) SetBody(v *FindPrepublishesByParentIdResponseBody) *FindPrepublishesByParentIdResponse {
	s.Body = v
	return s
}

type FindPrepublishesByVersionIdRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *int32  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindPrepublishesByVersionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByVersionIdRequest) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByVersionIdRequest) SetProjectId(v string) *FindPrepublishesByVersionIdRequest {
	s.ProjectId = &v
	return s
}

func (s *FindPrepublishesByVersionIdRequest) SetVersionId(v int32) *FindPrepublishesByVersionIdRequest {
	s.VersionId = &v
	return s
}

func (s *FindPrepublishesByVersionIdRequest) SetVersionType(v string) *FindPrepublishesByVersionIdRequest {
	s.VersionType = &v
	return s
}

type FindPrepublishesByVersionIdResponseBody struct {
	PrepublishList []*FindPrepublishesByVersionIdResponseBodyPrepublishList `json:"PrepublishList,omitempty" xml:"PrepublishList,omitempty" type:"Repeated"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindPrepublishesByVersionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByVersionIdResponseBody) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByVersionIdResponseBody) SetPrepublishList(v []*FindPrepublishesByVersionIdResponseBodyPrepublishList) *FindPrepublishesByVersionIdResponseBody {
	s.PrepublishList = v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBody) SetRequestId(v string) *FindPrepublishesByVersionIdResponseBody {
	s.RequestId = &v
	return s
}

type FindPrepublishesByVersionIdResponseBodyPrepublishList struct {
	BarrierCount       *string `json:"BarrierCount,omitempty" xml:"BarrierCount,omitempty"`
	DeviceModelId      *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModelName    *string `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsActive           *string `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	IsTotalPrepublish  *string `json:"IsTotalPrepublish,omitempty" xml:"IsTotalPrepublish,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId           *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	PassedCount        *string `json:"PassedCount,omitempty" xml:"PassedCount,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindPrepublishesByVersionIdResponseBodyPrepublishList) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByVersionIdResponseBodyPrepublishList) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetBarrierCount(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.BarrierCount = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetDeviceModelId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.DeviceModelId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetDeviceModelName(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.DeviceModelName = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtCreate(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtCreate = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtCreateTimestamp(v int64) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtModify(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtModify = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtModifyTimestamp(v int64) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.Id = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetIsActive(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.IsActive = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetIsTotalPrepublish(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.IsTotalPrepublish = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetName(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.Name = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetParentId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.ParentId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetPassedCount(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.PassedCount = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetVersionId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.VersionId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetVersionType(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.VersionType = &v
	return s
}

type FindPrepublishesByVersionIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindPrepublishesByVersionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindPrepublishesByVersionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByVersionIdResponse) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByVersionIdResponse) SetHeaders(v map[string]*string) *FindPrepublishesByVersionIdResponse {
	s.Headers = v
	return s
}

func (s *FindPrepublishesByVersionIdResponse) SetStatusCode(v int32) *FindPrepublishesByVersionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponse) SetBody(v *FindPrepublishesByVersionIdResponseBody) *FindPrepublishesByVersionIdResponse {
	s.Body = v
	return s
}

type FindVersionBlackDevicesRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OriginalId  *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesRequest) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesRequest) SetDeviceId(v string) *FindVersionBlackDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *FindVersionBlackDevicesRequest) SetOriginalId(v string) *FindVersionBlackDevicesRequest {
	s.OriginalId = &v
	return s
}

func (s *FindVersionBlackDevicesRequest) SetPageIndex(v int32) *FindVersionBlackDevicesRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionBlackDevicesRequest) SetPageSize(v int32) *FindVersionBlackDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionBlackDevicesRequest) SetProjectId(v string) *FindVersionBlackDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *FindVersionBlackDevicesRequest) SetVersionId(v string) *FindVersionBlackDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *FindVersionBlackDevicesRequest) SetVersionType(v string) *FindVersionBlackDevicesRequest {
	s.VersionType = &v
	return s
}

type FindVersionBlackDevicesResponseBody struct {
	DeviceList *FindVersionBlackDevicesResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindVersionBlackDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesResponseBody) SetDeviceList(v *FindVersionBlackDevicesResponseBodyDeviceList) *FindVersionBlackDevicesResponseBody {
	s.DeviceList = v
	return s
}

func (s *FindVersionBlackDevicesResponseBody) SetRequestId(v string) *FindVersionBlackDevicesResponseBody {
	s.RequestId = &v
	return s
}

type FindVersionBlackDevicesResponseBodyDeviceList struct {
	Items      []*FindVersionBlackDevicesResponseBodyDeviceListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionBlackDevicesResponseBodyDeviceList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesResponseBodyDeviceList) SetItems(v []*FindVersionBlackDevicesResponseBodyDeviceListItems) *FindVersionBlackDevicesResponseBodyDeviceList {
	s.Items = v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceList) SetTotalCount(v int32) *FindVersionBlackDevicesResponseBodyDeviceList {
	s.TotalCount = &v
	return s
}

type FindVersionBlackDevicesResponseBodyDeviceListItems struct {
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId         *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
}

func (s FindVersionBlackDevicesResponseBodyDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesResponseBodyDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetDeviceId(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetGmtCreate(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetGmtCreateTimestamp(v int64) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetId(v int64) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.Id = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetIdType(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.IdType = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetOriginalId(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.OriginalId = &v
	return s
}

type FindVersionBlackDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionBlackDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesResponse) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesResponse) SetHeaders(v map[string]*string) *FindVersionBlackDevicesResponse {
	s.Headers = v
	return s
}

func (s *FindVersionBlackDevicesResponse) SetStatusCode(v int32) *FindVersionBlackDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionBlackDevicesResponse) SetBody(v *FindVersionBlackDevicesResponseBody) *FindVersionBlackDevicesResponse {
	s.Body = v
	return s
}

type FindVersionDeviceGroupsRequest struct {
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginalId *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s FindVersionDeviceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsRequest) SetDeviceId(v string) *FindVersionDeviceGroupsRequest {
	s.DeviceId = &v
	return s
}

func (s *FindVersionDeviceGroupsRequest) SetName(v string) *FindVersionDeviceGroupsRequest {
	s.Name = &v
	return s
}

func (s *FindVersionDeviceGroupsRequest) SetOriginalId(v string) *FindVersionDeviceGroupsRequest {
	s.OriginalId = &v
	return s
}

func (s *FindVersionDeviceGroupsRequest) SetPageIndex(v int32) *FindVersionDeviceGroupsRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionDeviceGroupsRequest) SetPageSize(v int32) *FindVersionDeviceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionDeviceGroupsRequest) SetProjectId(v string) *FindVersionDeviceGroupsRequest {
	s.ProjectId = &v
	return s
}

type FindVersionDeviceGroupsResponseBody struct {
	DeviceGroupList *FindVersionDeviceGroupsResponseBodyDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Struct"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindVersionDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsResponseBody) SetDeviceGroupList(v *FindVersionDeviceGroupsResponseBodyDeviceGroupList) *FindVersionDeviceGroupsResponseBody {
	s.DeviceGroupList = v
	return s
}

func (s *FindVersionDeviceGroupsResponseBody) SetRequestId(v string) *FindVersionDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type FindVersionDeviceGroupsResponseBodyDeviceGroupList struct {
	Items      []*FindVersionDeviceGroupsResponseBodyDeviceGroupListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionDeviceGroupsResponseBodyDeviceGroupList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsResponseBodyDeviceGroupList) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupList) SetItems(v []*FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) *FindVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Items = v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupList) SetTotalCount(v int32) *FindVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.TotalCount = &v
	return s
}

type FindVersionDeviceGroupsResponseBodyDeviceGroupListItems struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxCount           *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetDescription(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.Description = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetGmtCreate(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetGmtCreateTimestamp(v int64) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetGmtModify(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.GmtModify = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetGmtModifyTimestamp(v int64) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetId(v int64) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.Id = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetMaxCount(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.MaxCount = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetName(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.Name = &v
	return s
}

type FindVersionDeviceGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionDeviceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsResponse) SetHeaders(v map[string]*string) *FindVersionDeviceGroupsResponse {
	s.Headers = v
	return s
}

func (s *FindVersionDeviceGroupsResponse) SetStatusCode(v int32) *FindVersionDeviceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionDeviceGroupsResponse) SetBody(v *FindVersionDeviceGroupsResponseBody) *FindVersionDeviceGroupsResponse {
	s.Body = v
	return s
}

type FindVersionGroupDevicesRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OriginalId    *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s FindVersionGroupDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesRequest) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesRequest) SetDeviceGroupId(v string) *FindVersionGroupDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *FindVersionGroupDevicesRequest) SetDeviceId(v string) *FindVersionGroupDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *FindVersionGroupDevicesRequest) SetOriginalId(v string) *FindVersionGroupDevicesRequest {
	s.OriginalId = &v
	return s
}

func (s *FindVersionGroupDevicesRequest) SetPageIndex(v int32) *FindVersionGroupDevicesRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionGroupDevicesRequest) SetPageSize(v int32) *FindVersionGroupDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionGroupDevicesRequest) SetProjectId(v string) *FindVersionGroupDevicesRequest {
	s.ProjectId = &v
	return s
}

type FindVersionGroupDevicesResponseBody struct {
	GroupDeviceList *FindVersionGroupDevicesResponseBodyGroupDeviceList `json:"GroupDeviceList,omitempty" xml:"GroupDeviceList,omitempty" type:"Struct"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindVersionGroupDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesResponseBody) SetGroupDeviceList(v *FindVersionGroupDevicesResponseBodyGroupDeviceList) *FindVersionGroupDevicesResponseBody {
	s.GroupDeviceList = v
	return s
}

func (s *FindVersionGroupDevicesResponseBody) SetRequestId(v string) *FindVersionGroupDevicesResponseBody {
	s.RequestId = &v
	return s
}

type FindVersionGroupDevicesResponseBodyGroupDeviceList struct {
	Items      []*FindVersionGroupDevicesResponseBodyGroupDeviceListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionGroupDevicesResponseBodyGroupDeviceList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesResponseBodyGroupDeviceList) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceList) SetItems(v []*FindVersionGroupDevicesResponseBodyGroupDeviceListItems) *FindVersionGroupDevicesResponseBodyGroupDeviceList {
	s.Items = v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceList) SetTotalCount(v int32) *FindVersionGroupDevicesResponseBodyGroupDeviceList {
	s.TotalCount = &v
	return s
}

type FindVersionGroupDevicesResponseBodyGroupDeviceListItems struct {
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId         *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
}

func (s FindVersionGroupDevicesResponseBodyGroupDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesResponseBodyGroupDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetDeviceId(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetGmtCreate(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetGmtCreateTimestamp(v int64) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetId(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.Id = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetIdType(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.IdType = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetOriginalId(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.OriginalId = &v
	return s
}

type FindVersionGroupDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionGroupDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionGroupDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesResponse) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesResponse) SetHeaders(v map[string]*string) *FindVersionGroupDevicesResponse {
	s.Headers = v
	return s
}

func (s *FindVersionGroupDevicesResponse) SetStatusCode(v int32) *FindVersionGroupDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionGroupDevicesResponse) SetBody(v *FindVersionGroupDevicesResponseBody) *FindVersionGroupDevicesResponse {
	s.Body = v
	return s
}

type FindVersionMessageSendRecordsRequest struct {
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionMessageSendRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsRequest) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsRequest) SetMessageType(v string) *FindVersionMessageSendRecordsRequest {
	s.MessageType = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetPageIndex(v int32) *FindVersionMessageSendRecordsRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetPageSize(v int32) *FindVersionMessageSendRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetProjectId(v string) *FindVersionMessageSendRecordsRequest {
	s.ProjectId = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetVersionId(v string) *FindVersionMessageSendRecordsRequest {
	s.VersionId = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetVersionType(v string) *FindVersionMessageSendRecordsRequest {
	s.VersionType = &v
	return s
}

type FindVersionMessageSendRecordsResponseBody struct {
	MessageSendRecordList *FindVersionMessageSendRecordsResponseBodyMessageSendRecordList `json:"MessageSendRecordList,omitempty" xml:"MessageSendRecordList,omitempty" type:"Struct"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindVersionMessageSendRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsResponseBody) SetMessageSendRecordList(v *FindVersionMessageSendRecordsResponseBodyMessageSendRecordList) *FindVersionMessageSendRecordsResponseBody {
	s.MessageSendRecordList = v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBody) SetRequestId(v string) *FindVersionMessageSendRecordsResponseBody {
	s.RequestId = &v
	return s
}

type FindVersionMessageSendRecordsResponseBodyMessageSendRecordList struct {
	Items      []*FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionMessageSendRecordsResponseBodyMessageSendRecordList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsResponseBodyMessageSendRecordList) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordList) SetItems(v []*FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordList {
	s.Items = v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordList) SetTotalCount(v int32) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordList {
	s.TotalCount = &v
	return s
}

type FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems struct {
	FailedCount        *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MessageType        *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Result             *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultDesc         *string `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty"`
	SkippedCount       *string `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	SucceededCount     *string `json:"SucceededCount,omitempty" xml:"SucceededCount,omitempty"`
	TargetId           *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetFailedCount(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.FailedCount = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetGmtCreate(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetGmtCreateTimestamp(v int64) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetId(v int64) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.Id = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetMessageType(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.MessageType = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetResult(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.Result = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetResultDesc(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.ResultDesc = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetSkippedCount(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.SkippedCount = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetSucceededCount(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.SucceededCount = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetTargetId(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.TargetId = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetVersionId(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.VersionId = &v
	return s
}

type FindVersionMessageSendRecordsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionMessageSendRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionMessageSendRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsResponse) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsResponse) SetHeaders(v map[string]*string) *FindVersionMessageSendRecordsResponse {
	s.Headers = v
	return s
}

func (s *FindVersionMessageSendRecordsResponse) SetStatusCode(v int32) *FindVersionMessageSendRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponse) SetBody(v *FindVersionMessageSendRecordsResponseBody) *FindVersionMessageSendRecordsResponse {
	s.Body = v
	return s
}

type FindVersionMessagesRequest struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	MessageType  *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SendRecordId *string `json:"SendRecordId,omitempty" xml:"SendRecordId,omitempty"`
	TestId       *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesRequest) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesRequest) SetDeviceId(v string) *FindVersionMessagesRequest {
	s.DeviceId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetMessageType(v string) *FindVersionMessagesRequest {
	s.MessageType = &v
	return s
}

func (s *FindVersionMessagesRequest) SetPageIndex(v int32) *FindVersionMessagesRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionMessagesRequest) SetPageSize(v int32) *FindVersionMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionMessagesRequest) SetProjectId(v string) *FindVersionMessagesRequest {
	s.ProjectId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetSendRecordId(v string) *FindVersionMessagesRequest {
	s.SendRecordId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetTestId(v string) *FindVersionMessagesRequest {
	s.TestId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetVersionId(v string) *FindVersionMessagesRequest {
	s.VersionId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetVersionType(v string) *FindVersionMessagesRequest {
	s.VersionType = &v
	return s
}

type FindVersionMessagesResponseBody struct {
	MessageList *FindVersionMessagesResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindVersionMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesResponseBody) SetMessageList(v *FindVersionMessagesResponseBodyMessageList) *FindVersionMessagesResponseBody {
	s.MessageList = v
	return s
}

func (s *FindVersionMessagesResponseBody) SetRequestId(v string) *FindVersionMessagesResponseBody {
	s.RequestId = &v
	return s
}

type FindVersionMessagesResponseBodyMessageList struct {
	Items      []*FindVersionMessagesResponseBodyMessageListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionMessagesResponseBodyMessageList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesResponseBodyMessageList) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesResponseBodyMessageList) SetItems(v []*FindVersionMessagesResponseBodyMessageListItems) *FindVersionMessagesResponseBodyMessageList {
	s.Items = v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageList) SetTotalCount(v int32) *FindVersionMessagesResponseBodyMessageList {
	s.TotalCount = &v
	return s
}

type FindVersionMessagesResponseBodyMessageListItems struct {
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MessageId          *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDesc         *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	TestId             *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s FindVersionMessagesResponseBodyMessageListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesResponseBodyMessageListItems) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetDeviceId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtCreate(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtCreateTimestamp(v int64) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtModify(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtModify = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtModifyTimestamp(v int64) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetId(v int64) *FindVersionMessagesResponseBodyMessageListItems {
	s.Id = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetMessageId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.MessageId = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetStatus(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.Status = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetStatusDesc(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.StatusDesc = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetTestId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.TestId = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetVersionId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.VersionId = &v
	return s
}

type FindVersionMessagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesResponse) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesResponse) SetHeaders(v map[string]*string) *FindVersionMessagesResponse {
	s.Headers = v
	return s
}

func (s *FindVersionMessagesResponse) SetStatusCode(v int32) *FindVersionMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionMessagesResponse) SetBody(v *FindVersionMessagesResponseBody) *FindVersionMessagesResponse {
	s.Body = v
	return s
}

type FindVersionTestsRequest struct {
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionTestsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsRequest) GoString() string {
	return s.String()
}

func (s *FindVersionTestsRequest) SetPageIndex(v int32) *FindVersionTestsRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionTestsRequest) SetPageSize(v int32) *FindVersionTestsRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionTestsRequest) SetProjectId(v string) *FindVersionTestsRequest {
	s.ProjectId = &v
	return s
}

func (s *FindVersionTestsRequest) SetVersionId(v string) *FindVersionTestsRequest {
	s.VersionId = &v
	return s
}

func (s *FindVersionTestsRequest) SetVersionType(v string) *FindVersionTestsRequest {
	s.VersionType = &v
	return s
}

type FindVersionTestsResponseBody struct {
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionTestList *FindVersionTestsResponseBodyVersionTestList `json:"VersionTestList,omitempty" xml:"VersionTestList,omitempty" type:"Struct"`
}

func (s FindVersionTestsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionTestsResponseBody) SetRequestId(v string) *FindVersionTestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindVersionTestsResponseBody) SetVersionTestList(v *FindVersionTestsResponseBodyVersionTestList) *FindVersionTestsResponseBody {
	s.VersionTestList = v
	return s
}

type FindVersionTestsResponseBodyVersionTestList struct {
	Items      []*FindVersionTestsResponseBodyVersionTestListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionTestsResponseBodyVersionTestList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsResponseBodyVersionTestList) GoString() string {
	return s.String()
}

func (s *FindVersionTestsResponseBodyVersionTestList) SetItems(v []*FindVersionTestsResponseBodyVersionTestListItems) *FindVersionTestsResponseBodyVersionTestList {
	s.Items = v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestList) SetTotalCount(v int32) *FindVersionTestsResponseBodyVersionTestList {
	s.TotalCount = &v
	return s
}

type FindVersionTestsResponseBodyVersionTestListItems struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceGroupId      *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceGroupName    *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	FailedCount        *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkippedCount       *string `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	SucceededCount     *string `json:"SucceededCount,omitempty" xml:"SucceededCount,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionTestsResponseBodyVersionTestListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsResponseBodyVersionTestListItems) GoString() string {
	return s.String()
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetDescription(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.Description = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetDeviceGroupId(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.DeviceGroupId = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetDeviceGroupName(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.DeviceGroupName = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetFailedCount(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.FailedCount = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtCreate(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtCreateTimestamp(v int64) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtModify(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtModify = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtModifyTimestamp(v int64) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetId(v int64) *FindVersionTestsResponseBodyVersionTestListItems {
	s.Id = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetName(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.Name = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetSkippedCount(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.SkippedCount = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetSucceededCount(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.SucceededCount = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetVersionId(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.VersionId = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetVersionType(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.VersionType = &v
	return s
}

type FindVersionTestsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionTestsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionTestsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsResponse) GoString() string {
	return s.String()
}

func (s *FindVersionTestsResponse) SetHeaders(v map[string]*string) *FindVersionTestsResponse {
	s.Headers = v
	return s
}

func (s *FindVersionTestsResponse) SetStatusCode(v int32) *FindVersionTestsResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionTestsResponse) SetBody(v *FindVersionTestsResponseBody) *FindVersionTestsResponse {
	s.Body = v
	return s
}

type FindVersionWhiteDevicesRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OriginalId  *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesRequest) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesRequest) SetDeviceId(v string) *FindVersionWhiteDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *FindVersionWhiteDevicesRequest) SetOriginalId(v string) *FindVersionWhiteDevicesRequest {
	s.OriginalId = &v
	return s
}

func (s *FindVersionWhiteDevicesRequest) SetPageIndex(v int32) *FindVersionWhiteDevicesRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionWhiteDevicesRequest) SetPageSize(v int32) *FindVersionWhiteDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *FindVersionWhiteDevicesRequest) SetProjectId(v string) *FindVersionWhiteDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *FindVersionWhiteDevicesRequest) SetVersionId(v string) *FindVersionWhiteDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *FindVersionWhiteDevicesRequest) SetVersionType(v string) *FindVersionWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

type FindVersionWhiteDevicesResponseBody struct {
	DeviceList *FindVersionWhiteDevicesResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindVersionWhiteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesResponseBody) SetDeviceList(v *FindVersionWhiteDevicesResponseBodyDeviceList) *FindVersionWhiteDevicesResponseBody {
	s.DeviceList = v
	return s
}

func (s *FindVersionWhiteDevicesResponseBody) SetRequestId(v string) *FindVersionWhiteDevicesResponseBody {
	s.RequestId = &v
	return s
}

type FindVersionWhiteDevicesResponseBodyDeviceList struct {
	Items      []*FindVersionWhiteDevicesResponseBodyDeviceListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCount *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindVersionWhiteDevicesResponseBodyDeviceList) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceList) SetItems(v []*FindVersionWhiteDevicesResponseBodyDeviceListItems) *FindVersionWhiteDevicesResponseBodyDeviceList {
	s.Items = v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceList) SetTotalCount(v int32) *FindVersionWhiteDevicesResponseBodyDeviceList {
	s.TotalCount = &v
	return s
}

type FindVersionWhiteDevicesResponseBodyDeviceListItems struct {
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId         *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
}

func (s FindVersionWhiteDevicesResponseBodyDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesResponseBodyDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetDeviceId(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetGmtCreate(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetGmtCreateTimestamp(v int64) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetId(v int64) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.Id = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetIdType(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.IdType = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetOriginalId(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.OriginalId = &v
	return s
}

type FindVersionWhiteDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FindVersionWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindVersionWhiteDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesResponse) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesResponse) SetHeaders(v map[string]*string) *FindVersionWhiteDevicesResponse {
	s.Headers = v
	return s
}

func (s *FindVersionWhiteDevicesResponse) SetStatusCode(v int32) *FindVersionWhiteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindVersionWhiteDevicesResponse) SetBody(v *FindVersionWhiteDevicesResponseBody) *FindVersionWhiteDevicesResponse {
	s.Body = v
	return s
}

type GenerateFunctionFileUploadMetaRequest struct {
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GenerateFunctionFileUploadMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaRequest) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaRequest) SetFileName(v string) *GenerateFunctionFileUploadMetaRequest {
	s.FileName = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaRequest) SetProjectId(v string) *GenerateFunctionFileUploadMetaRequest {
	s.ProjectId = &v
	return s
}

type GenerateFunctionFileUploadMetaResponseBody struct {
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadMeta *GenerateFunctionFileUploadMetaResponseBodyUploadMeta `json:"UploadMeta,omitempty" xml:"UploadMeta,omitempty" type:"Struct"`
}

func (s GenerateFunctionFileUploadMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaResponseBody) SetRequestId(v string) *GenerateFunctionFileUploadMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBody) SetUploadMeta(v *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) *GenerateFunctionFileUploadMetaResponseBody {
	s.UploadMeta = v
	return s
}

type GenerateFunctionFileUploadMetaResponseBodyUploadMeta struct {
	ObjectKey        *string                                                               `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	PostObjectPolicy *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy `json:"PostObjectPolicy,omitempty" xml:"PostObjectPolicy,omitempty" type:"Struct"`
	SecurityToken    *string                                                               `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMeta) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMeta) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) SetObjectKey(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMeta {
	s.ObjectKey = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) SetPostObjectPolicy(v *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) *GenerateFunctionFileUploadMetaResponseBodyUploadMeta {
	s.PostObjectPolicy = v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) SetSecurityToken(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMeta {
	s.SecurityToken = &v
	return s
}

type GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetAccessId(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.AccessId = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetExpire(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.Expire = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetHost(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.Host = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetPolicy(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.Policy = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetSignature(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.Signature = &v
	return s
}

type GenerateFunctionFileUploadMetaResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateFunctionFileUploadMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateFunctionFileUploadMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaResponse) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaResponse) SetHeaders(v map[string]*string) *GenerateFunctionFileUploadMetaResponse {
	s.Headers = v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponse) SetStatusCode(v int32) *GenerateFunctionFileUploadMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponse) SetBody(v *GenerateFunctionFileUploadMetaResponseBody) *GenerateFunctionFileUploadMetaResponse {
	s.Body = v
	return s
}

type GenerateOssPostPolicyRequest struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GenerateOssPostPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyRequest) SetAccessId(v string) *GenerateOssPostPolicyRequest {
	s.AccessId = &v
	return s
}

func (s *GenerateOssPostPolicyRequest) SetAccessKey(v string) *GenerateOssPostPolicyRequest {
	s.AccessKey = &v
	return s
}

func (s *GenerateOssPostPolicyRequest) SetExt(v string) *GenerateOssPostPolicyRequest {
	s.Ext = &v
	return s
}

func (s *GenerateOssPostPolicyRequest) SetProjectId(v string) *GenerateOssPostPolicyRequest {
	s.ProjectId = &v
	return s
}

type GenerateOssPostPolicyResponseBody struct {
	OssPostPolicy *GenerateOssPostPolicyResponseBodyOssPostPolicy `json:"OssPostPolicy,omitempty" xml:"OssPostPolicy,omitempty" type:"Struct"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateOssPostPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyResponseBody) SetOssPostPolicy(v *GenerateOssPostPolicyResponseBodyOssPostPolicy) *GenerateOssPostPolicyResponseBody {
	s.OssPostPolicy = v
	return s
}

func (s *GenerateOssPostPolicyResponseBody) SetRequestId(v string) *GenerateOssPostPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateOssPostPolicyResponseBodyOssPostPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateOssPostPolicyResponseBodyOssPostPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyResponseBodyOssPostPolicy) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetAccessId(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.AccessId = &v
	return s
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetExpire(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.Expire = &v
	return s
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetHost(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.Host = &v
	return s
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetPolicy(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.Policy = &v
	return s
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetSignature(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.Signature = &v
	return s
}

type GenerateOssPostPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateOssPostPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateOssPostPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyResponse) SetHeaders(v map[string]*string) *GenerateOssPostPolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateOssPostPolicyResponse) SetStatusCode(v int32) *GenerateOssPostPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOssPostPolicyResponse) SetBody(v *GenerateOssPostPolicyResponseBody) *GenerateOssPostPolicyResponse {
	s.Body = v
	return s
}

type GenerateOssUploadMetaRequest struct {
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GenerateOssUploadMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaRequest) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaRequest) SetExt(v string) *GenerateOssUploadMetaRequest {
	s.Ext = &v
	return s
}

func (s *GenerateOssUploadMetaRequest) SetProjectId(v string) *GenerateOssUploadMetaRequest {
	s.ProjectId = &v
	return s
}

type GenerateOssUploadMetaResponseBody struct {
	OssUploadMeta *GenerateOssUploadMetaResponseBodyOssUploadMeta `json:"OssUploadMeta,omitempty" xml:"OssUploadMeta,omitempty" type:"Struct"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateOssUploadMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaResponseBody) SetOssUploadMeta(v *GenerateOssUploadMetaResponseBodyOssUploadMeta) *GenerateOssUploadMetaResponseBody {
	s.OssUploadMeta = v
	return s
}

func (s *GenerateOssUploadMetaResponseBody) SetRequestId(v string) *GenerateOssUploadMetaResponseBody {
	s.RequestId = &v
	return s
}

type GenerateOssUploadMetaResponseBodyOssUploadMeta struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Bucket          *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	ObjectKey       *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GenerateOssUploadMetaResponseBodyOssUploadMeta) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaResponseBodyOssUploadMeta) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetAccessKeyId(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.AccessKeyId = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetAccessKeySecret(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.AccessKeySecret = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetBucket(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.Bucket = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetObjectKey(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.ObjectKey = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetSecurityToken(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.SecurityToken = &v
	return s
}

type GenerateOssUploadMetaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateOssUploadMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateOssUploadMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaResponse) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaResponse) SetHeaders(v map[string]*string) *GenerateOssUploadMetaResponse {
	s.Headers = v
	return s
}

func (s *GenerateOssUploadMetaResponse) SetStatusCode(v int32) *GenerateOssUploadMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOssUploadMetaResponse) SetBody(v *GenerateOssUploadMetaResponseBody) *GenerateOssUploadMetaResponse {
	s.Body = v
	return s
}

type GenerateSdkDownloadInfoRequest struct {
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CertFileObjectKey *string `json:"CertFileObjectKey,omitempty" xml:"CertFileObjectKey,omitempty"`
	OsType            *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PkgName           *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Sdks              *string `json:"Sdks,omitempty" xml:"Sdks,omitempty"`
}

func (s GenerateSdkDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateSdkDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GenerateSdkDownloadInfoRequest) SetAppId(v string) *GenerateSdkDownloadInfoRequest {
	s.AppId = &v
	return s
}

func (s *GenerateSdkDownloadInfoRequest) SetCertFileObjectKey(v string) *GenerateSdkDownloadInfoRequest {
	s.CertFileObjectKey = &v
	return s
}

func (s *GenerateSdkDownloadInfoRequest) SetOsType(v int32) *GenerateSdkDownloadInfoRequest {
	s.OsType = &v
	return s
}

func (s *GenerateSdkDownloadInfoRequest) SetPkgName(v string) *GenerateSdkDownloadInfoRequest {
	s.PkgName = &v
	return s
}

func (s *GenerateSdkDownloadInfoRequest) SetProjectId(v string) *GenerateSdkDownloadInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateSdkDownloadInfoRequest) SetSdks(v string) *GenerateSdkDownloadInfoRequest {
	s.Sdks = &v
	return s
}

type GenerateSdkDownloadInfoResponseBody struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkDownloadInfo *GenerateSdkDownloadInfoResponseBodySdkDownloadInfo `json:"SdkDownloadInfo,omitempty" xml:"SdkDownloadInfo,omitempty" type:"Struct"`
}

func (s GenerateSdkDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateSdkDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSdkDownloadInfoResponseBody) SetRequestId(v string) *GenerateSdkDownloadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateSdkDownloadInfoResponseBody) SetSdkDownloadInfo(v *GenerateSdkDownloadInfoResponseBodySdkDownloadInfo) *GenerateSdkDownloadInfoResponseBody {
	s.SdkDownloadInfo = v
	return s
}

type GenerateSdkDownloadInfoResponseBodySdkDownloadInfo struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateSdkDownloadInfoResponseBodySdkDownloadInfo) String() string {
	return tea.Prettify(s)
}

func (s GenerateSdkDownloadInfoResponseBodySdkDownloadInfo) GoString() string {
	return s.String()
}

func (s *GenerateSdkDownloadInfoResponseBodySdkDownloadInfo) SetUrl(v string) *GenerateSdkDownloadInfoResponseBodySdkDownloadInfo {
	s.Url = &v
	return s
}

type GenerateSdkDownloadInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateSdkDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateSdkDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateSdkDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GenerateSdkDownloadInfoResponse) SetHeaders(v map[string]*string) *GenerateSdkDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GenerateSdkDownloadInfoResponse) SetStatusCode(v int32) *GenerateSdkDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSdkDownloadInfoResponse) SetBody(v *GenerateSdkDownloadInfoResponseBody) *GenerateSdkDownloadInfoResponse {
	s.Body = v
	return s
}

type GenerateSysAppDownloadInfoRequest struct {
	CertFileObjectKey *string `json:"CertFileObjectKey,omitempty" xml:"CertFileObjectKey,omitempty"`
	OsType            *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PkgName           *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	Plugins           *string `json:"Plugins,omitempty" xml:"Plugins,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SignMode          *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
}

func (s GenerateSysAppDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateSysAppDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GenerateSysAppDownloadInfoRequest) SetCertFileObjectKey(v string) *GenerateSysAppDownloadInfoRequest {
	s.CertFileObjectKey = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetOsType(v int32) *GenerateSysAppDownloadInfoRequest {
	s.OsType = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetPkgName(v string) *GenerateSysAppDownloadInfoRequest {
	s.PkgName = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetPlugins(v string) *GenerateSysAppDownloadInfoRequest {
	s.Plugins = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetProjectId(v string) *GenerateSysAppDownloadInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetSignMode(v string) *GenerateSysAppDownloadInfoRequest {
	s.SignMode = &v
	return s
}

type GenerateSysAppDownloadInfoResponseBody struct {
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SysAppDownloadInfo *GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo `json:"SysAppDownloadInfo,omitempty" xml:"SysAppDownloadInfo,omitempty" type:"Struct"`
}

func (s GenerateSysAppDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateSysAppDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSysAppDownloadInfoResponseBody) SetRequestId(v string) *GenerateSysAppDownloadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateSysAppDownloadInfoResponseBody) SetSysAppDownloadInfo(v *GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo) *GenerateSysAppDownloadInfoResponseBody {
	s.SysAppDownloadInfo = v
	return s
}

type GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo) String() string {
	return tea.Prettify(s)
}

func (s GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo) GoString() string {
	return s.String()
}

func (s *GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo) SetUrl(v string) *GenerateSysAppDownloadInfoResponseBodySysAppDownloadInfo {
	s.Url = &v
	return s
}

type GenerateSysAppDownloadInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateSysAppDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateSysAppDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateSysAppDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GenerateSysAppDownloadInfoResponse) SetHeaders(v map[string]*string) *GenerateSysAppDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GenerateSysAppDownloadInfoResponse) SetStatusCode(v int32) *GenerateSysAppDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSysAppDownloadInfoResponse) SetBody(v *GenerateSysAppDownloadInfoResponseBody) *GenerateSysAppDownloadInfoResponse {
	s.Body = v
	return s
}

type GetDeviceAppUpdateFunnelEventsRequest struct {
	IdType            *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId        *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PackageName       *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TargetVersionCode *string `json:"TargetVersionCode,omitempty" xml:"TargetVersionCode,omitempty"`
}

func (s GetDeviceAppUpdateFunnelEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetIdType(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.IdType = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetOriginalId(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.OriginalId = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetPackageName(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.PackageName = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetProjectId(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetTargetVersionCode(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.TargetVersionCode = &v
	return s
}

type GetDeviceAppUpdateFunnelEventsResponseBody struct {
	EventList []*GetDeviceAppUpdateFunnelEventsResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceAppUpdateFunnelEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBody) SetEventList(v []*GetDeviceAppUpdateFunnelEventsResponseBodyEventList) *GetDeviceAppUpdateFunnelEventsResponseBody {
	s.EventList = v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBody) SetRequestId(v string) *GetDeviceAppUpdateFunnelEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceAppUpdateFunnelEventsResponseBodyEventList struct {
	DeviceId          *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Event             *string `json:"Event,omitempty" xml:"Event,omitempty"`
	PackageName       *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ReportTime        *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	ReportTimestamp   *int64  `json:"ReportTimestamp,omitempty" xml:"ReportTimestamp,omitempty"`
	TargetVersionCode *string `json:"TargetVersionCode,omitempty" xml:"TargetVersionCode,omitempty"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceAppUpdateFunnelEventsResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetDeviceId(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetEvent(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.Event = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetPackageName(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.PackageName = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetReportTime(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.ReportTime = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetReportTimestamp(v int64) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.ReportTimestamp = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetTargetVersionCode(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.TargetVersionCode = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetTenantId(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.TenantId = &v
	return s
}

type GetDeviceAppUpdateFunnelEventsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceAppUpdateFunnelEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceAppUpdateFunnelEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceAppUpdateFunnelEventsResponse) SetHeaders(v map[string]*string) *GetDeviceAppUpdateFunnelEventsResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponse) SetStatusCode(v int32) *GetDeviceAppUpdateFunnelEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponse) SetBody(v *GetDeviceAppUpdateFunnelEventsResponseBody) *GetDeviceAppUpdateFunnelEventsResponse {
	s.Body = v
	return s
}

type GetDeviceSystemUpdateFunnelEventsRequest struct {
	IdType        *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId    *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s GetDeviceSystemUpdateFunnelEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSystemUpdateFunnelEventsRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceSystemUpdateFunnelEventsRequest) SetIdType(v string) *GetDeviceSystemUpdateFunnelEventsRequest {
	s.IdType = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsRequest) SetOriginalId(v string) *GetDeviceSystemUpdateFunnelEventsRequest {
	s.OriginalId = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsRequest) SetProjectId(v string) *GetDeviceSystemUpdateFunnelEventsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsRequest) SetTargetVersion(v string) *GetDeviceSystemUpdateFunnelEventsRequest {
	s.TargetVersion = &v
	return s
}

type GetDeviceSystemUpdateFunnelEventsResponseBody struct {
	EventList []*GetDeviceSystemUpdateFunnelEventsResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceSystemUpdateFunnelEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSystemUpdateFunnelEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBody) SetEventList(v []*GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) *GetDeviceSystemUpdateFunnelEventsResponseBody {
	s.EventList = v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBody) SetRequestId(v string) *GetDeviceSystemUpdateFunnelEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceSystemUpdateFunnelEventsResponseBodyEventList struct {
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Event           *string `json:"Event,omitempty" xml:"Event,omitempty"`
	ReportTime      *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	ReportTimestamp *int64  `json:"ReportTimestamp,omitempty" xml:"ReportTimestamp,omitempty"`
	TargetVersion   *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	TenantId        *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetDeviceId(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetEvent(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.Event = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetReportTime(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.ReportTime = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetReportTimestamp(v int64) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.ReportTimestamp = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetTargetVersion(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.TargetVersion = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetTenantId(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.TenantId = &v
	return s
}

type GetDeviceSystemUpdateFunnelEventsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceSystemUpdateFunnelEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceSystemUpdateFunnelEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSystemUpdateFunnelEventsResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceSystemUpdateFunnelEventsResponse) SetHeaders(v map[string]*string) *GetDeviceSystemUpdateFunnelEventsResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponse) SetStatusCode(v int32) *GetDeviceSystemUpdateFunnelEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponse) SetBody(v *GetDeviceSystemUpdateFunnelEventsResponseBody) *GetDeviceSystemUpdateFunnelEventsResponse {
	s.Body = v
	return s
}

type GetNamespaceDataRequest struct {
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountType  *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AuthType     *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetNamespaceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceDataRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceDataRequest) SetAccountId(v string) *GetNamespaceDataRequest {
	s.AccountId = &v
	return s
}

func (s *GetNamespaceDataRequest) SetAccountType(v string) *GetNamespaceDataRequest {
	s.AccountType = &v
	return s
}

func (s *GetNamespaceDataRequest) SetAuthType(v string) *GetNamespaceDataRequest {
	s.AuthType = &v
	return s
}

func (s *GetNamespaceDataRequest) SetDeviceId(v string) *GetNamespaceDataRequest {
	s.DeviceId = &v
	return s
}

func (s *GetNamespaceDataRequest) SetDeviceIdType(v string) *GetNamespaceDataRequest {
	s.DeviceIdType = &v
	return s
}

func (s *GetNamespaceDataRequest) SetNamespace(v string) *GetNamespaceDataRequest {
	s.Namespace = &v
	return s
}

func (s *GetNamespaceDataRequest) SetProjectId(v string) *GetNamespaceDataRequest {
	s.ProjectId = &v
	return s
}

type GetNamespaceDataResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNamespaceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceDataResponseBody) SetData(v string) *GetNamespaceDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetNamespaceDataResponseBody) SetRequestId(v string) *GetNamespaceDataResponseBody {
	s.RequestId = &v
	return s
}

type GetNamespaceDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNamespaceDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNamespaceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceDataResponse) GoString() string {
	return s.String()
}

func (s *GetNamespaceDataResponse) SetHeaders(v map[string]*string) *GetNamespaceDataResponse {
	s.Headers = v
	return s
}

func (s *GetNamespaceDataResponse) SetStatusCode(v int32) *GetNamespaceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNamespaceDataResponse) SetBody(v *GetNamespaceDataResponseBody) *GetNamespaceDataResponse {
	s.Body = v
	return s
}

type GetNamespaceStatisticsDataRequest struct {
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetNamespaceStatisticsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceStatisticsDataRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceStatisticsDataRequest) SetDimensionType(v string) *GetNamespaceStatisticsDataRequest {
	s.DimensionType = &v
	return s
}

func (s *GetNamespaceStatisticsDataRequest) SetEndTime(v string) *GetNamespaceStatisticsDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetNamespaceStatisticsDataRequest) SetNamespace(v string) *GetNamespaceStatisticsDataRequest {
	s.Namespace = &v
	return s
}

func (s *GetNamespaceStatisticsDataRequest) SetProjectId(v string) *GetNamespaceStatisticsDataRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNamespaceStatisticsDataRequest) SetStartTime(v string) *GetNamespaceStatisticsDataRequest {
	s.StartTime = &v
	return s
}

type GetNamespaceStatisticsDataResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics *GetNamespaceStatisticsDataResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
}

func (s GetNamespaceStatisticsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceStatisticsDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceStatisticsDataResponseBody) SetRequestId(v string) *GetNamespaceStatisticsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNamespaceStatisticsDataResponseBody) SetStatistics(v *GetNamespaceStatisticsDataResponseBodyStatistics) *GetNamespaceStatisticsDataResponseBody {
	s.Statistics = v
	return s
}

type GetNamespaceStatisticsDataResponseBodyStatistics struct {
	Categories []*int64                                                  `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Series     []*GetNamespaceStatisticsDataResponseBodyStatisticsSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s GetNamespaceStatisticsDataResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceStatisticsDataResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *GetNamespaceStatisticsDataResponseBodyStatistics) SetCategories(v []*int64) *GetNamespaceStatisticsDataResponseBodyStatistics {
	s.Categories = v
	return s
}

func (s *GetNamespaceStatisticsDataResponseBodyStatistics) SetSeries(v []*GetNamespaceStatisticsDataResponseBodyStatisticsSeries) *GetNamespaceStatisticsDataResponseBodyStatistics {
	s.Series = v
	return s
}

type GetNamespaceStatisticsDataResponseBodyStatisticsSeries struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Name *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetNamespaceStatisticsDataResponseBodyStatisticsSeries) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceStatisticsDataResponseBodyStatisticsSeries) GoString() string {
	return s.String()
}

func (s *GetNamespaceStatisticsDataResponseBodyStatisticsSeries) SetData(v []*int64) *GetNamespaceStatisticsDataResponseBodyStatisticsSeries {
	s.Data = v
	return s
}

func (s *GetNamespaceStatisticsDataResponseBodyStatisticsSeries) SetName(v string) *GetNamespaceStatisticsDataResponseBodyStatisticsSeries {
	s.Name = &v
	return s
}

type GetNamespaceStatisticsDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNamespaceStatisticsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNamespaceStatisticsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceStatisticsDataResponse) GoString() string {
	return s.String()
}

func (s *GetNamespaceStatisticsDataResponse) SetHeaders(v map[string]*string) *GetNamespaceStatisticsDataResponse {
	s.Headers = v
	return s
}

func (s *GetNamespaceStatisticsDataResponse) SetStatusCode(v int32) *GetNamespaceStatisticsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNamespaceStatisticsDataResponse) SetBody(v *GetNamespaceStatisticsDataResponseBody) *GetNamespaceStatisticsDataResponse {
	s.Body = v
	return s
}

type GetOssUploadMetaRequest struct {
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetOssUploadMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadMetaRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadMetaRequest) SetExt(v string) *GetOssUploadMetaRequest {
	s.Ext = &v
	return s
}

func (s *GetOssUploadMetaRequest) SetProjectId(v string) *GetOssUploadMetaRequest {
	s.ProjectId = &v
	return s
}

type GetOssUploadMetaResponseBody struct {
	OssUploadMeta *GetOssUploadMetaResponseBodyOssUploadMeta `json:"OssUploadMeta,omitempty" xml:"OssUploadMeta,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssUploadMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadMetaResponseBody) SetOssUploadMeta(v *GetOssUploadMetaResponseBodyOssUploadMeta) *GetOssUploadMetaResponseBody {
	s.OssUploadMeta = v
	return s
}

func (s *GetOssUploadMetaResponseBody) SetRequestId(v string) *GetOssUploadMetaResponseBody {
	s.RequestId = &v
	return s
}

type GetOssUploadMetaResponseBodyOssUploadMeta struct {
	AccessKey     *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	ObjectKey     *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOssUploadMetaResponseBodyOssUploadMeta) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadMetaResponseBodyOssUploadMeta) GoString() string {
	return s.String()
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetAccessKey(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.AccessKey = &v
	return s
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetHost(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.Host = &v
	return s
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetObjectKey(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.ObjectKey = &v
	return s
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetPolicy(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.Policy = &v
	return s
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetSecurityToken(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.SecurityToken = &v
	return s
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetSignature(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.Signature = &v
	return s
}

type GetOssUploadMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOssUploadMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssUploadMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadMetaResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadMetaResponse) SetHeaders(v map[string]*string) *GetOssUploadMetaResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadMetaResponse) SetStatusCode(v int32) *GetOssUploadMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadMetaResponse) SetBody(v *GetOssUploadMetaResponseBody) *GetOssUploadMetaResponse {
	s.Body = v
	return s
}

type InvokeFunctionRequest struct {
	Env          *int32  `json:"Env,omitempty" xml:"Env,omitempty"`
	FileId       *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Parameters   *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s InvokeFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionRequest) GoString() string {
	return s.String()
}

func (s *InvokeFunctionRequest) SetEnv(v int32) *InvokeFunctionRequest {
	s.Env = &v
	return s
}

func (s *InvokeFunctionRequest) SetFileId(v int64) *InvokeFunctionRequest {
	s.FileId = &v
	return s
}

func (s *InvokeFunctionRequest) SetFunctionName(v string) *InvokeFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *InvokeFunctionRequest) SetParameters(v string) *InvokeFunctionRequest {
	s.Parameters = &v
	return s
}

func (s *InvokeFunctionRequest) SetProjectId(v string) *InvokeFunctionRequest {
	s.ProjectId = &v
	return s
}

type InvokeFunctionResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InvokeFunctionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InvokeFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponseBody) SetRequestId(v string) *InvokeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeFunctionResponseBody) SetResult(v *InvokeFunctionResponseBodyResult) *InvokeFunctionResponseBody {
	s.Result = v
	return s
}

type InvokeFunctionResponseBodyResult struct {
	BackEndRequestId *string `json:"BackEndRequestId,omitempty" xml:"BackEndRequestId,omitempty"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s InvokeFunctionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponseBodyResult) SetBackEndRequestId(v string) *InvokeFunctionResponseBodyResult {
	s.BackEndRequestId = &v
	return s
}

func (s *InvokeFunctionResponseBodyResult) SetOutput(v string) *InvokeFunctionResponseBodyResult {
	s.Output = &v
	return s
}

type InvokeFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionResponse) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponse) SetHeaders(v map[string]*string) *InvokeFunctionResponse {
	s.Headers = v
	return s
}

func (s *InvokeFunctionResponse) SetStatusCode(v int32) *InvokeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeFunctionResponse) SetBody(v *InvokeFunctionResponseBody) *InvokeFunctionResponse {
	s.Body = v
	return s
}

type ListApiGatewayAppsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListApiGatewayAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApiGatewayAppsRequest) GoString() string {
	return s.String()
}

func (s *ListApiGatewayAppsRequest) SetProjectId(v string) *ListApiGatewayAppsRequest {
	s.ProjectId = &v
	return s
}

type ListApiGatewayAppsResponseBody struct {
	ApiGatewayApps []*ListApiGatewayAppsResponseBodyApiGatewayApps `json:"ApiGatewayApps,omitempty" xml:"ApiGatewayApps,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApiGatewayAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApiGatewayAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiGatewayAppsResponseBody) SetApiGatewayApps(v []*ListApiGatewayAppsResponseBodyApiGatewayApps) *ListApiGatewayAppsResponseBody {
	s.ApiGatewayApps = v
	return s
}

func (s *ListApiGatewayAppsResponseBody) SetRequestId(v string) *ListApiGatewayAppsResponseBody {
	s.RequestId = &v
	return s
}

type ListApiGatewayAppsResponseBodyApiGatewayApps struct {
	GatewayAppId     *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	GatewayAppKey    *string `json:"GatewayAppKey,omitempty" xml:"GatewayAppKey,omitempty"`
	GatewayAppSecret *string `json:"GatewayAppSecret,omitempty" xml:"GatewayAppSecret,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApiGatewayAppsResponseBodyApiGatewayApps) String() string {
	return tea.Prettify(s)
}

func (s ListApiGatewayAppsResponseBodyApiGatewayApps) GoString() string {
	return s.String()
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGatewayAppId(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GatewayAppId = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGatewayAppKey(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GatewayAppKey = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGatewayAppSecret(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GatewayAppSecret = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGmtCreate(v int64) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GmtCreate = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGmtModified(v int64) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GmtModified = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetId(v int64) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.Id = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetProjectId(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.ProjectId = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetStatus(v int32) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.Status = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetUserId(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.UserId = &v
	return s
}

type ListApiGatewayAppsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApiGatewayAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApiGatewayAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApiGatewayAppsResponse) GoString() string {
	return s.String()
}

func (s *ListApiGatewayAppsResponse) SetHeaders(v map[string]*string) *ListApiGatewayAppsResponse {
	s.Headers = v
	return s
}

func (s *ListApiGatewayAppsResponse) SetStatusCode(v int32) *ListApiGatewayAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiGatewayAppsResponse) SetBody(v *ListApiGatewayAppsResponseBody) *ListApiGatewayAppsResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetProjectId(v string) *ListAppsRequest {
	s.ProjectId = &v
	return s
}

type ListAppsResponseBody struct {
	Apps      []*ListAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetApps(v []*ListAppsResponseBodyApps) *ListAppsResponseBody {
	s.Apps = v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

type ListAppsResponseBodyApps struct {
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPackage *string `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	OsType     *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s ListAppsResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyApps) SetAppKey(v string) *ListAppsResponseBodyApps {
	s.AppKey = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppName(v string) *ListAppsResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppPackage(v string) *ListAppsResponseBodyApps {
	s.AppPackage = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetOsType(v int32) *ListAppsResponseBodyApps {
	s.OsType = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListAssistActionDetailsRequest struct {
	ActionTimestamp *string `json:"ActionTimestamp,omitempty" xml:"ActionTimestamp,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAssistActionDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistActionDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListAssistActionDetailsRequest) SetActionTimestamp(v string) *ListAssistActionDetailsRequest {
	s.ActionTimestamp = &v
	return s
}

func (s *ListAssistActionDetailsRequest) SetProjectId(v string) *ListAssistActionDetailsRequest {
	s.ProjectId = &v
	return s
}

type ListAssistActionDetailsResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ListAssistActionDetailsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ListAssistActionDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistActionDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistActionDetailsResponseBody) SetRequestId(v string) *ListAssistActionDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssistActionDetailsResponseBody) SetResults(v []*ListAssistActionDetailsResponseBodyResults) *ListAssistActionDetailsResponseBody {
	s.Results = v
	return s
}

type ListAssistActionDetailsResponseBodyResults struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ID        *string `json:"ID,omitempty" xml:"ID,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListAssistActionDetailsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListAssistActionDetailsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListAssistActionDetailsResponseBodyResults) SetAction(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Action = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetCreatedAt(v int64) *ListAssistActionDetailsResponseBodyResults {
	s.CreatedAt = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetData(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Data = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetDeviceId(v string) *ListAssistActionDetailsResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetID(v string) *ListAssistActionDetailsResponseBodyResults {
	s.ID = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetTimestamp(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Timestamp = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetType(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Type = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetUpdatedAt(v int64) *ListAssistActionDetailsResponseBodyResults {
	s.UpdatedAt = &v
	return s
}

type ListAssistActionDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAssistActionDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAssistActionDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistActionDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListAssistActionDetailsResponse) SetHeaders(v map[string]*string) *ListAssistActionDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListAssistActionDetailsResponse) SetStatusCode(v int32) *ListAssistActionDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistActionDetailsResponse) SetBody(v *ListAssistActionDetailsResponseBody) *ListAssistActionDetailsResponse {
	s.Body = v
	return s
}

type ListAssistDevicesRequest struct {
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PerPage   *int32  `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAssistDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesRequest) SetCondition(v string) *ListAssistDevicesRequest {
	s.Condition = &v
	return s
}

func (s *ListAssistDevicesRequest) SetPageIndex(v int32) *ListAssistDevicesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAssistDevicesRequest) SetPerPage(v int32) *ListAssistDevicesRequest {
	s.PerPage = &v
	return s
}

func (s *ListAssistDevicesRequest) SetProjectId(v string) *ListAssistDevicesRequest {
	s.ProjectId = &v
	return s
}

type ListAssistDevicesResponseBody struct {
	Devices    []*ListAssistDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageIndex  *int32                                  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PerPage    *int32                                  `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAssistDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesResponseBody) SetDevices(v []*ListAssistDevicesResponseBodyDevices) *ListAssistDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *ListAssistDevicesResponseBody) SetPageIndex(v int32) *ListAssistDevicesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetPerPage(v int32) *ListAssistDevicesResponseBody {
	s.PerPage = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetRequestId(v string) *ListAssistDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetTotalCount(v int32) *ListAssistDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAssistDevicesResponseBodyDevices struct {
	AccessTime   *int64  `json:"AccessTime,omitempty" xml:"AccessTime,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	HardwareId   *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	UUID         *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VIN          *string `json:"VIN,omitempty" xml:"VIN,omitempty"`
}

func (s ListAssistDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesResponseBodyDevices) SetAccessTime(v int64) *ListAssistDevicesResponseBodyDevices {
	s.AccessTime = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetDeviceId(v string) *ListAssistDevicesResponseBodyDevices {
	s.DeviceId = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetDeviceName(v string) *ListAssistDevicesResponseBodyDevices {
	s.DeviceName = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetHardwareId(v string) *ListAssistDevicesResponseBodyDevices {
	s.HardwareId = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetSerialNumber(v string) *ListAssistDevicesResponseBodyDevices {
	s.SerialNumber = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetUUID(v string) *ListAssistDevicesResponseBodyDevices {
	s.UUID = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetVIN(v string) *ListAssistDevicesResponseBodyDevices {
	s.VIN = &v
	return s
}

type ListAssistDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAssistDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAssistDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesResponse) SetHeaders(v map[string]*string) *ListAssistDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListAssistDevicesResponse) SetStatusCode(v int32) *ListAssistDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistDevicesResponse) SetBody(v *ListAssistDevicesResponseBody) *ListAssistDevicesResponse {
	s.Body = v
	return s
}

type ListAssistHistoriesRequest struct {
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PerPage   *int32  `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAssistHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesRequest) SetCondition(v string) *ListAssistHistoriesRequest {
	s.Condition = &v
	return s
}

func (s *ListAssistHistoriesRequest) SetPageIndex(v int32) *ListAssistHistoriesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAssistHistoriesRequest) SetPerPage(v int32) *ListAssistHistoriesRequest {
	s.PerPage = &v
	return s
}

func (s *ListAssistHistoriesRequest) SetProjectId(v string) *ListAssistHistoriesRequest {
	s.ProjectId = &v
	return s
}

type ListAssistHistoriesResponseBody struct {
	Histories  []*ListAssistHistoriesResponseBodyHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Repeated"`
	PageIndex  *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PerPage    *int32                                      `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAssistHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesResponseBody) SetHistories(v []*ListAssistHistoriesResponseBodyHistories) *ListAssistHistoriesResponseBody {
	s.Histories = v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetPageIndex(v int32) *ListAssistHistoriesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetPerPage(v int32) *ListAssistHistoriesResponseBody {
	s.PerPage = &v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetRequestId(v string) *ListAssistHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetTotalCount(v int32) *ListAssistHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAssistHistoriesResponseBodyHistories struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HardwareId   *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	ID           *string `json:"ID,omitempty" xml:"ID,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UID          *string `json:"UID,omitempty" xml:"UID,omitempty"`
	UNAME        *string `json:"UNAME,omitempty" xml:"UNAME,omitempty"`
	UUID         *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VIN          *string `json:"VIN,omitempty" xml:"VIN,omitempty"`
}

func (s ListAssistHistoriesResponseBodyHistories) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesResponseBodyHistories) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesResponseBodyHistories) SetDeviceId(v string) *ListAssistHistoriesResponseBodyHistories {
	s.DeviceId = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetDeviceName(v string) *ListAssistHistoriesResponseBodyHistories {
	s.DeviceName = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetEndTime(v int64) *ListAssistHistoriesResponseBodyHistories {
	s.EndTime = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetHardwareId(v string) *ListAssistHistoriesResponseBodyHistories {
	s.HardwareId = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetID(v string) *ListAssistHistoriesResponseBodyHistories {
	s.ID = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetSerialNumber(v string) *ListAssistHistoriesResponseBodyHistories {
	s.SerialNumber = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetStartTime(v int64) *ListAssistHistoriesResponseBodyHistories {
	s.StartTime = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetUID(v string) *ListAssistHistoriesResponseBodyHistories {
	s.UID = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetUNAME(v string) *ListAssistHistoriesResponseBodyHistories {
	s.UNAME = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetUUID(v string) *ListAssistHistoriesResponseBodyHistories {
	s.UUID = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetVIN(v string) *ListAssistHistoriesResponseBodyHistories {
	s.VIN = &v
	return s
}

type ListAssistHistoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAssistHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAssistHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesResponse) SetHeaders(v map[string]*string) *ListAssistHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAssistHistoriesResponse) SetStatusCode(v int32) *ListAssistHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistHistoriesResponse) SetBody(v *ListAssistHistoriesResponseBody) *ListAssistHistoriesResponse {
	s.Body = v
	return s
}

type ListAssistHistoryDetailsRequest struct {
	AssistId  *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAssistHistoryDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoryDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListAssistHistoryDetailsRequest) SetAssistId(v string) *ListAssistHistoryDetailsRequest {
	s.AssistId = &v
	return s
}

func (s *ListAssistHistoryDetailsRequest) SetProjectId(v string) *ListAssistHistoryDetailsRequest {
	s.ProjectId = &v
	return s
}

type ListAssistHistoryDetailsResponseBody struct {
	Actions   []*ListAssistHistoryDetailsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssistHistoryDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoryDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistHistoryDetailsResponseBody) SetActions(v []*ListAssistHistoryDetailsResponseBodyActions) *ListAssistHistoryDetailsResponseBody {
	s.Actions = v
	return s
}

func (s *ListAssistHistoryDetailsResponseBody) SetRequestId(v string) *ListAssistHistoryDetailsResponseBody {
	s.RequestId = &v
	return s
}

type ListAssistHistoryDetailsResponseBodyActions struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	AssistId  *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ID        *string `json:"ID,omitempty" xml:"ID,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListAssistHistoryDetailsResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoryDetailsResponseBodyActions) GoString() string {
	return s.String()
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetAction(v string) *ListAssistHistoryDetailsResponseBodyActions {
	s.Action = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetAssistId(v string) *ListAssistHistoryDetailsResponseBodyActions {
	s.AssistId = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetCreatedAt(v int64) *ListAssistHistoryDetailsResponseBodyActions {
	s.CreatedAt = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetID(v string) *ListAssistHistoryDetailsResponseBodyActions {
	s.ID = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetTimestamp(v string) *ListAssistHistoryDetailsResponseBodyActions {
	s.Timestamp = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetUpdatedAt(v int64) *ListAssistHistoryDetailsResponseBodyActions {
	s.UpdatedAt = &v
	return s
}

type ListAssistHistoryDetailsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAssistHistoryDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAssistHistoryDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoryDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListAssistHistoryDetailsResponse) SetHeaders(v map[string]*string) *ListAssistHistoryDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListAssistHistoryDetailsResponse) SetStatusCode(v int32) *ListAssistHistoryDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistHistoryDetailsResponse) SetBody(v *ListAssistHistoryDetailsResponseBody) *ListAssistHistoryDetailsResponse {
	s.Body = v
	return s
}

type ListClientPluginVersionsRequest struct {
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PkgName *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
}

func (s ListClientPluginVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListClientPluginVersionsRequest) SetOsType(v string) *ListClientPluginVersionsRequest {
	s.OsType = &v
	return s
}

func (s *ListClientPluginVersionsRequest) SetPkgName(v string) *ListClientPluginVersionsRequest {
	s.PkgName = &v
	return s
}

type ListClientPluginVersionsResponseBody struct {
	ClientPluginVersions []*ListClientPluginVersionsResponseBodyClientPluginVersions `json:"ClientPluginVersions,omitempty" xml:"ClientPluginVersions,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientPluginVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientPluginVersionsResponseBody) SetClientPluginVersions(v []*ListClientPluginVersionsResponseBodyClientPluginVersions) *ListClientPluginVersionsResponseBody {
	s.ClientPluginVersions = v
	return s
}

func (s *ListClientPluginVersionsResponseBody) SetRequestId(v string) *ListClientPluginVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListClientPluginVersionsResponseBodyClientPluginVersions struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PkgName     *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode *int64  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListClientPluginVersionsResponseBodyClientPluginVersions) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginVersionsResponseBodyClientPluginVersions) GoString() string {
	return s.String()
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetDownloadUrl(v string) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.DownloadUrl = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetId(v int64) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.Id = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetPkgName(v string) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.PkgName = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetSize(v int64) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.Size = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetVersion(v string) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.Version = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetVersionCode(v int64) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.VersionCode = &v
	return s
}

type ListClientPluginVersionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClientPluginVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClientPluginVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListClientPluginVersionsResponse) SetHeaders(v map[string]*string) *ListClientPluginVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListClientPluginVersionsResponse) SetStatusCode(v int32) *ListClientPluginVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientPluginVersionsResponse) SetBody(v *ListClientPluginVersionsResponseBody) *ListClientPluginVersionsResponse {
	s.Body = v
	return s
}

type ListClientPluginsRequest struct {
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s ListClientPluginsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListClientPluginsRequest) SetOsType(v string) *ListClientPluginsRequest {
	s.OsType = &v
	return s
}

type ListClientPluginsResponseBody struct {
	ClientPlugins []*ListClientPluginsResponseBodyClientPlugins `json:"ClientPlugins,omitempty" xml:"ClientPlugins,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientPluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientPluginsResponseBody) SetClientPlugins(v []*ListClientPluginsResponseBodyClientPlugins) *ListClientPluginsResponseBody {
	s.ClientPlugins = v
	return s
}

func (s *ListClientPluginsResponseBody) SetRequestId(v string) *ListClientPluginsResponseBody {
	s.RequestId = &v
	return s
}

type ListClientPluginsResponseBodyClientPlugins struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PkgName *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
}

func (s ListClientPluginsResponseBodyClientPlugins) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginsResponseBodyClientPlugins) GoString() string {
	return s.String()
}

func (s *ListClientPluginsResponseBodyClientPlugins) SetName(v string) *ListClientPluginsResponseBodyClientPlugins {
	s.Name = &v
	return s
}

func (s *ListClientPluginsResponseBodyClientPlugins) SetPkgName(v string) *ListClientPluginsResponseBodyClientPlugins {
	s.PkgName = &v
	return s
}

type ListClientPluginsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClientPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClientPluginsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListClientPluginsResponse) SetHeaders(v map[string]*string) *ListClientPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListClientPluginsResponse) SetStatusCode(v int32) *ListClientPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientPluginsResponse) SetBody(v *ListClientPluginsResponseBody) *ListClientPluginsResponse {
	s.Body = v
	return s
}

type ListClientSdksRequest struct {
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s ListClientSdksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClientSdksRequest) GoString() string {
	return s.String()
}

func (s *ListClientSdksRequest) SetOsType(v string) *ListClientSdksRequest {
	s.OsType = &v
	return s
}

type ListClientSdksResponseBody struct {
	ClientSdks []*ListClientSdksResponseBodyClientSdks `json:"ClientSdks,omitempty" xml:"ClientSdks,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientSdksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientSdksResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientSdksResponseBody) SetClientSdks(v []*ListClientSdksResponseBodyClientSdks) *ListClientSdksResponseBody {
	s.ClientSdks = v
	return s
}

func (s *ListClientSdksResponseBody) SetRequestId(v string) *ListClientSdksResponseBody {
	s.RequestId = &v
	return s
}

type ListClientSdksResponseBodyClientSdks struct {
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OsType      *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PkgName     *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	PkgType     *int32  `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
}

func (s ListClientSdksResponseBodyClientSdks) String() string {
	return tea.Prettify(s)
}

func (s ListClientSdksResponseBodyClientSdks) GoString() string {
	return s.String()
}

func (s *ListClientSdksResponseBodyClientSdks) SetGmtCreate(v int64) *ListClientSdksResponseBodyClientSdks {
	s.GmtCreate = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetGmtModified(v int64) *ListClientSdksResponseBodyClientSdks {
	s.GmtModified = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetId(v int64) *ListClientSdksResponseBodyClientSdks {
	s.Id = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetName(v string) *ListClientSdksResponseBodyClientSdks {
	s.Name = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetOsType(v int32) *ListClientSdksResponseBodyClientSdks {
	s.OsType = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetPkgName(v string) *ListClientSdksResponseBodyClientSdks {
	s.PkgName = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetPkgType(v int32) *ListClientSdksResponseBodyClientSdks {
	s.PkgType = &v
	return s
}

type ListClientSdksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClientSdksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClientSdksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClientSdksResponse) GoString() string {
	return s.String()
}

func (s *ListClientSdksResponse) SetHeaders(v map[string]*string) *ListClientSdksResponse {
	s.Headers = v
	return s
}

func (s *ListClientSdksResponse) SetStatusCode(v int32) *ListClientSdksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientSdksResponse) SetBody(v *ListClientSdksResponseBody) *ListClientSdksResponse {
	s.Body = v
	return s
}

type ListConnectLogsRequest struct {
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListConnectLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectLogsRequest) SetDeviceId(v string) *ListConnectLogsRequest {
	s.DeviceId = &v
	return s
}

func (s *ListConnectLogsRequest) SetEndTime(v int32) *ListConnectLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListConnectLogsRequest) SetPageIndex(v int32) *ListConnectLogsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListConnectLogsRequest) SetPageSize(v int32) *ListConnectLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectLogsRequest) SetProjectId(v string) *ListConnectLogsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListConnectLogsRequest) SetStartTime(v int32) *ListConnectLogsRequest {
	s.StartTime = &v
	return s
}

type ListConnectLogsResponseBody struct {
	Logs      *ListConnectLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConnectLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBody) SetLogs(v *ListConnectLogsResponseBodyLogs) *ListConnectLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListConnectLogsResponseBody) SetRequestId(v string) *ListConnectLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListConnectLogsResponseBodyLogs struct {
	List       []*ListConnectLogsResponseBodyLogsList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListConnectLogsResponseBodyLogsPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListConnectLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBodyLogs) SetList(v []*ListConnectLogsResponseBodyLogsList) *ListConnectLogsResponseBodyLogs {
	s.List = v
	return s
}

func (s *ListConnectLogsResponseBodyLogs) SetPagination(v *ListConnectLogsResponseBodyLogsPagination) *ListConnectLogsResponseBodyLogs {
	s.Pagination = v
	return s
}

type ListConnectLogsResponseBodyLogsList struct {
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NetWorking    *string `json:"NetWorking,omitempty" xml:"NetWorking,omitempty"`
	Sid           *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	Terminal      *string `json:"Terminal,omitempty" xml:"Terminal,omitempty"`
	Time          *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListConnectLogsResponseBodyLogsList) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBodyLogsList) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBodyLogsList) SetDeviceId(v string) *ListConnectLogsResponseBodyLogsList {
	s.DeviceId = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetIp(v string) *ListConnectLogsResponseBodyLogsList {
	s.Ip = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetNetWorking(v string) *ListConnectLogsResponseBodyLogsList {
	s.NetWorking = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetSid(v string) *ListConnectLogsResponseBodyLogsList {
	s.Sid = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetStatus(v string) *ListConnectLogsResponseBodyLogsList {
	s.Status = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetSystemVersion(v string) *ListConnectLogsResponseBodyLogsList {
	s.SystemVersion = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetTerminal(v string) *ListConnectLogsResponseBodyLogsList {
	s.Terminal = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetTime(v int64) *ListConnectLogsResponseBodyLogsList {
	s.Time = &v
	return s
}

type ListConnectLogsResponseBodyLogsPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListConnectLogsResponseBodyLogsPagination) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBodyLogsPagination) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBodyLogsPagination) SetPageIndex(v int32) *ListConnectLogsResponseBodyLogsPagination {
	s.PageIndex = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsPagination) SetPageSize(v int32) *ListConnectLogsResponseBodyLogsPagination {
	s.PageSize = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsPagination) SetTotalCount(v int32) *ListConnectLogsResponseBodyLogsPagination {
	s.TotalCount = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsPagination) SetTotalPageCount(v int32) *ListConnectLogsResponseBodyLogsPagination {
	s.TotalPageCount = &v
	return s
}

type ListConnectLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConnectLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponse) SetHeaders(v map[string]*string) *ListConnectLogsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectLogsResponse) SetStatusCode(v int32) *ListConnectLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectLogsResponse) SetBody(v *ListConnectLogsResponseBody) *ListConnectLogsResponse {
	s.Body = v
	return s
}

type ListDeployedFunctionsRequest struct {
	FileId    *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDeployedFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsRequest) SetFileId(v int64) *ListDeployedFunctionsRequest {
	s.FileId = &v
	return s
}

func (s *ListDeployedFunctionsRequest) SetProjectId(v string) *ListDeployedFunctionsRequest {
	s.ProjectId = &v
	return s
}

type ListDeployedFunctionsResponseBody struct {
	Functions []*ListDeployedFunctionsResponseBodyFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeployedFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsResponseBody) SetFunctions(v []*ListDeployedFunctionsResponseBodyFunctions) *ListDeployedFunctionsResponseBody {
	s.Functions = v
	return s
}

func (s *ListDeployedFunctionsResponseBody) SetRequestId(v string) *ListDeployedFunctionsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeployedFunctionsResponseBodyFunctions struct {
	FileId      *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDeployedFunctionsResponseBodyFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsResponseBodyFunctions) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetFileId(v int64) *ListDeployedFunctionsResponseBodyFunctions {
	s.FileId = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetGmtCreate(v int64) *ListDeployedFunctionsResponseBodyFunctions {
	s.GmtCreate = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetGmtModified(v int64) *ListDeployedFunctionsResponseBodyFunctions {
	s.GmtModified = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetId(v int64) *ListDeployedFunctionsResponseBodyFunctions {
	s.Id = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetName(v string) *ListDeployedFunctionsResponseBodyFunctions {
	s.Name = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetProjectId(v string) *ListDeployedFunctionsResponseBodyFunctions {
	s.ProjectId = &v
	return s
}

type ListDeployedFunctionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeployedFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployedFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsResponse) SetHeaders(v map[string]*string) *ListDeployedFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListDeployedFunctionsResponse) SetStatusCode(v int32) *ListDeployedFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeployedFunctionsResponse) SetBody(v *ListDeployedFunctionsResponseBody) *ListDeployedFunctionsResponse {
	s.Body = v
	return s
}

type ListDeviceBrandsRequest struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListDeviceBrandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBrandsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBrandsRequest) SetDeviceBrand(v string) *ListDeviceBrandsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetDeviceBrandId(v int64) *ListDeviceBrandsRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetLength(v string) *ListDeviceBrandsRequest {
	s.Length = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetProjectId(v string) *ListDeviceBrandsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetStart(v string) *ListDeviceBrandsRequest {
	s.Start = &v
	return s
}

type ListDeviceBrandsResponseBody struct {
	DeviceBrands []*ListDeviceBrandsResponseBodyDeviceBrands `json:"DeviceBrands,omitempty" xml:"DeviceBrands,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceBrandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBrandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceBrandsResponseBody) SetDeviceBrands(v []*ListDeviceBrandsResponseBodyDeviceBrands) *ListDeviceBrandsResponseBody {
	s.DeviceBrands = v
	return s
}

func (s *ListDeviceBrandsResponseBody) SetRequestId(v string) *ListDeviceBrandsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceBrandsResponseBodyDeviceBrands struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Manufacture   *string `json:"Manufacture,omitempty" xml:"Manufacture,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDeviceBrandsResponseBodyDeviceBrands) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBrandsResponseBodyDeviceBrands) GoString() string {
	return s.String()
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetDescription(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.Description = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetDeviceBrand(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.DeviceBrand = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetDeviceBrandId(v int64) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.DeviceBrandId = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetManufacture(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.Manufacture = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetProjectId(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.ProjectId = &v
	return s
}

type ListDeviceBrandsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceBrandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBrandsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceBrandsResponse) SetHeaders(v map[string]*string) *ListDeviceBrandsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceBrandsResponse) SetStatusCode(v int32) *ListDeviceBrandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceBrandsResponse) SetBody(v *ListDeviceBrandsResponseBody) *ListDeviceBrandsResponse {
	s.Body = v
	return s
}

type ListDeviceModelsRequest struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListDeviceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceModelsRequest) SetDeviceBrand(v string) *ListDeviceModelsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *ListDeviceModelsRequest) SetDeviceBrandId(v int64) *ListDeviceModelsRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *ListDeviceModelsRequest) SetDeviceModel(v string) *ListDeviceModelsRequest {
	s.DeviceModel = &v
	return s
}

func (s *ListDeviceModelsRequest) SetDeviceModelId(v int32) *ListDeviceModelsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *ListDeviceModelsRequest) SetLength(v string) *ListDeviceModelsRequest {
	s.Length = &v
	return s
}

func (s *ListDeviceModelsRequest) SetProjectId(v string) *ListDeviceModelsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceModelsRequest) SetStart(v string) *ListDeviceModelsRequest {
	s.Start = &v
	return s
}

type ListDeviceModelsResponseBody struct {
	DeviceModels []*ListDeviceModelsResponseBodyDeviceModels `json:"DeviceModels,omitempty" xml:"DeviceModels,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceModelsResponseBody) SetDeviceModels(v []*ListDeviceModelsResponseBodyDeviceModels) *ListDeviceModelsResponseBody {
	s.DeviceModels = v
	return s
}

func (s *ListDeviceModelsResponseBody) SetRequestId(v string) *ListDeviceModelsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceModelsResponseBodyDeviceModels struct {
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceLogoUrl     *string `json:"DeviceLogoUrl,omitempty" xml:"DeviceLogoUrl,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
}

func (s ListDeviceModelsResponseBodyDeviceModels) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelsResponseBodyDeviceModels) GoString() string {
	return s.String()
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetCanCreateDeviceId(v int32) *ListDeviceModelsResponseBodyDeviceModels {
	s.CanCreateDeviceId = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDescription(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.Description = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceBrand(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceBrand = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceLogoUrl(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceLogoUrl = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceModel(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceModel = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceModelId(v int64) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceModelId = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceName(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceType(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceType = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetHardwareType(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.HardwareType = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetInitUsageType(v int32) *ListDeviceModelsResponseBodyDeviceModels {
	s.InitUsageType = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetInitUsageTypeDesc(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetObjectKey(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.ObjectKey = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetOsPlatform(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.OsPlatform = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetProjectId(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetSecurityChip(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.SecurityChip = &v
	return s
}

type ListDeviceModelsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceModelsResponse) SetHeaders(v map[string]*string) *ListDeviceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceModelsResponse) SetStatusCode(v int32) *ListDeviceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceModelsResponse) SetBody(v *ListDeviceModelsResponseBody) *ListDeviceModelsResponse {
	s.Body = v
	return s
}

type ListDeviceTypesRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDeviceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceTypesRequest) SetProjectId(v string) *ListDeviceTypesRequest {
	s.ProjectId = &v
	return s
}

type ListDeviceTypesResponseBody struct {
	DeviceTypes []*ListDeviceTypesResponseBodyDeviceTypes `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceTypesResponseBody) SetDeviceTypes(v []*ListDeviceTypesResponseBodyDeviceTypes) *ListDeviceTypesResponseBody {
	s.DeviceTypes = v
	return s
}

func (s *ListDeviceTypesResponseBody) SetRequestId(v string) *ListDeviceTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceTypesResponseBodyDeviceTypes struct {
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDeviceTypesResponseBodyDeviceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceTypesResponseBodyDeviceTypes) GoString() string {
	return s.String()
}

func (s *ListDeviceTypesResponseBodyDeviceTypes) SetDeviceType(v string) *ListDeviceTypesResponseBodyDeviceTypes {
	s.DeviceType = &v
	return s
}

func (s *ListDeviceTypesResponseBodyDeviceTypes) SetName(v string) *ListDeviceTypesResponseBodyDeviceTypes {
	s.Name = &v
	return s
}

type ListDeviceTypesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceTypesResponse) SetHeaders(v map[string]*string) *ListDeviceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceTypesResponse) SetStatusCode(v int32) *ListDeviceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceTypesResponse) SetBody(v *ListDeviceTypesResponseBody) *ListDeviceTypesResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetDeviceModel(v string) *ListDevicesRequest {
	s.DeviceModel = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceModelId(v int32) *ListDevicesRequest {
	s.DeviceModelId = &v
	return s
}

func (s *ListDevicesRequest) SetLength(v string) *ListDevicesRequest {
	s.Length = &v
	return s
}

func (s *ListDevicesRequest) SetProjectId(v string) *ListDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDevicesRequest) SetStart(v string) *ListDevicesRequest {
	s.Start = &v
	return s
}

type ListDevicesResponseBody struct {
	Devices   []*ListDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetDevices(v []*ListDevicesResponseBodyDevices) *ListDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

type ListDevicesResponseBodyDevices struct {
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareId    *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	MacAddress    *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SoftwareId    *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsageType     *int32  `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	UsageTypeDesc *string `json:"UsageTypeDesc,omitempty" xml:"UsageTypeDesc,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Vin           *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
}

func (s ListDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDevices) SetDeviceBrand(v string) *ListDevicesResponseBodyDevices {
	s.DeviceBrand = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceId(v string) *ListDevicesResponseBodyDevices {
	s.DeviceId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceModel(v string) *ListDevicesResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceModelId(v int64) *ListDevicesResponseBodyDevices {
	s.DeviceModelId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceType(v string) *ListDevicesResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetHardwareId(v string) *ListDevicesResponseBodyDevices {
	s.HardwareId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetMacAddress(v string) *ListDevicesResponseBodyDevices {
	s.MacAddress = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetName(v string) *ListDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetProjectId(v string) *ListDevicesResponseBodyDevices {
	s.ProjectId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetRegion(v string) *ListDevicesResponseBodyDevices {
	s.Region = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSerialNumber(v string) *ListDevicesResponseBodyDevices {
	s.SerialNumber = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSoftwareId(v string) *ListDevicesResponseBodyDevices {
	s.SoftwareId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetStatus(v string) *ListDevicesResponseBodyDevices {
	s.Status = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetUsageType(v int32) *ListDevicesResponseBodyDevices {
	s.UsageType = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetUsageTypeDesc(v string) *ListDevicesResponseBodyDevices {
	s.UsageTypeDesc = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetUuid(v string) *ListDevicesResponseBodyDevices {
	s.Uuid = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetVin(v string) *ListDevicesResponseBodyDevices {
	s.Vin = &v
	return s
}

type ListDevicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) SetHeaders(v map[string]*string) *ListDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicesResponse) SetStatusCode(v int32) *ListDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

type ListFunctionExecuteLogRequest struct {
	Env          *int32  `json:"Env,omitempty" xml:"Env,omitempty"`
	FileId       *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListFunctionExecuteLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogRequest) SetEnv(v int32) *ListFunctionExecuteLogRequest {
	s.Env = &v
	return s
}

func (s *ListFunctionExecuteLogRequest) SetFileId(v int64) *ListFunctionExecuteLogRequest {
	s.FileId = &v
	return s
}

func (s *ListFunctionExecuteLogRequest) SetFunctionName(v string) *ListFunctionExecuteLogRequest {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionExecuteLogRequest) SetPageIndex(v int32) *ListFunctionExecuteLogRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionExecuteLogRequest) SetPageSize(v int32) *ListFunctionExecuteLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionExecuteLogRequest) SetProjectId(v string) *ListFunctionExecuteLogRequest {
	s.ProjectId = &v
	return s
}

type ListFunctionExecuteLogResponseBody struct {
	LogList   *ListFunctionExecuteLogResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionExecuteLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBody) SetLogList(v *ListFunctionExecuteLogResponseBodyLogList) *ListFunctionExecuteLogResponseBody {
	s.LogList = v
	return s
}

func (s *ListFunctionExecuteLogResponseBody) SetRequestId(v string) *ListFunctionExecuteLogResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionExecuteLogResponseBodyLogList struct {
	Logs       []*ListFunctionExecuteLogResponseBodyLogListLogs     `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	Pagination *ListFunctionExecuteLogResponseBodyLogListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListFunctionExecuteLogResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBodyLogList) SetLogs(v []*ListFunctionExecuteLogResponseBodyLogListLogs) *ListFunctionExecuteLogResponseBodyLogList {
	s.Logs = v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogList) SetPagination(v *ListFunctionExecuteLogResponseBodyLogListPagination) *ListFunctionExecuteLogResponseBodyLogList {
	s.Pagination = v
	return s
}

type ListFunctionExecuteLogResponseBodyLogListLogs struct {
	BackEndRequestId *string `json:"BackEndRequestId,omitempty" xml:"BackEndRequestId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListFunctionExecuteLogResponseBodyLogListLogs) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBodyLogListLogs) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBodyLogListLogs) SetBackEndRequestId(v string) *ListFunctionExecuteLogResponseBodyLogListLogs {
	s.BackEndRequestId = &v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogListLogs) SetMessage(v string) *ListFunctionExecuteLogResponseBodyLogListLogs {
	s.Message = &v
	return s
}

type ListFunctionExecuteLogResponseBodyLogListPagination struct {
	HasNextPage *bool  `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	PageIndex   *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFunctionExecuteLogResponseBodyLogListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBodyLogListPagination) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBodyLogListPagination) SetHasNextPage(v bool) *ListFunctionExecuteLogResponseBodyLogListPagination {
	s.HasNextPage = &v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogListPagination) SetPageIndex(v int32) *ListFunctionExecuteLogResponseBodyLogListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogListPagination) SetPageSize(v int32) *ListFunctionExecuteLogResponseBodyLogListPagination {
	s.PageSize = &v
	return s
}

type ListFunctionExecuteLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionExecuteLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionExecuteLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponse) SetHeaders(v map[string]*string) *ListFunctionExecuteLogResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionExecuteLogResponse) SetStatusCode(v int32) *ListFunctionExecuteLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionExecuteLogResponse) SetBody(v *ListFunctionExecuteLogResponseBody) *ListFunctionExecuteLogResponse {
	s.Body = v
	return s
}

type ListFunctionFilesRequest struct {
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListFunctionFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesRequest) SetFileType(v int32) *ListFunctionFilesRequest {
	s.FileType = &v
	return s
}

func (s *ListFunctionFilesRequest) SetPageIndex(v int32) *ListFunctionFilesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionFilesRequest) SetPageSize(v int32) *ListFunctionFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionFilesRequest) SetProjectId(v string) *ListFunctionFilesRequest {
	s.ProjectId = &v
	return s
}

type ListFunctionFilesResponseBody struct {
	FileList  *ListFunctionFilesResponseBodyFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBody) SetFileList(v *ListFunctionFilesResponseBodyFileList) *ListFunctionFilesResponseBody {
	s.FileList = v
	return s
}

func (s *ListFunctionFilesResponseBody) SetRequestId(v string) *ListFunctionFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionFilesResponseBodyFileList struct {
	Files      []*ListFunctionFilesResponseBodyFileListFiles    `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	Pagination *ListFunctionFilesResponseBodyFileListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListFunctionFilesResponseBodyFileList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBodyFileList) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBodyFileList) SetFiles(v []*ListFunctionFilesResponseBodyFileListFiles) *ListFunctionFilesResponseBodyFileList {
	s.Files = v
	return s
}

func (s *ListFunctionFilesResponseBodyFileList) SetPagination(v *ListFunctionFilesResponseBodyFileListPagination) *ListFunctionFilesResponseBodyFileList {
	s.Pagination = v
	return s
}

type ListFunctionFilesResponseBodyFileListFiles struct {
	ContentId              *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate              *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified            *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductionDeployStatus *int32  `json:"ProductionDeployStatus,omitempty" xml:"ProductionDeployStatus,omitempty"`
	ProductionDeployTime   *int64  `json:"ProductionDeployTime,omitempty" xml:"ProductionDeployTime,omitempty"`
	SandboxDeployStatus    *int32  `json:"SandboxDeployStatus,omitempty" xml:"SandboxDeployStatus,omitempty"`
	SandboxDeployTime      *int64  `json:"SandboxDeployTime,omitempty" xml:"SandboxDeployTime,omitempty"`
	Status                 *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionFilesResponseBodyFileListFiles) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBodyFileListFiles) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetContentId(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.ContentId = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetDescription(v string) *ListFunctionFilesResponseBodyFileListFiles {
	s.Description = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetGmtCreate(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.GmtCreate = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetGmtModified(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.GmtModified = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetId(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.Id = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetName(v string) *ListFunctionFilesResponseBodyFileListFiles {
	s.Name = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetProductionDeployStatus(v int32) *ListFunctionFilesResponseBodyFileListFiles {
	s.ProductionDeployStatus = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetProductionDeployTime(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.ProductionDeployTime = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetSandboxDeployStatus(v int32) *ListFunctionFilesResponseBodyFileListFiles {
	s.SandboxDeployStatus = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetSandboxDeployTime(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.SandboxDeployTime = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetStatus(v int32) *ListFunctionFilesResponseBodyFileListFiles {
	s.Status = &v
	return s
}

type ListFunctionFilesResponseBodyFileListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListFunctionFilesResponseBodyFileListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBodyFileListPagination) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBodyFileListPagination) SetPageIndex(v int32) *ListFunctionFilesResponseBodyFileListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListPagination) SetPageSize(v int32) *ListFunctionFilesResponseBodyFileListPagination {
	s.PageSize = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListPagination) SetTotalCount(v int32) *ListFunctionFilesResponseBodyFileListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListPagination) SetTotalPageCount(v int32) *ListFunctionFilesResponseBodyFileListPagination {
	s.TotalPageCount = &v
	return s
}

type ListFunctionFilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponse) SetHeaders(v map[string]*string) *ListFunctionFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionFilesResponse) SetStatusCode(v int32) *ListFunctionFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionFilesResponse) SetBody(v *ListFunctionFilesResponseBody) *ListFunctionFilesResponse {
	s.Body = v
	return s
}

type ListFunctionFilesByProjectIdRequest struct {
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListFunctionFilesByProjectIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdRequest) SetPageIndex(v int32) *ListFunctionFilesByProjectIdRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionFilesByProjectIdRequest) SetPageSize(v int32) *ListFunctionFilesByProjectIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionFilesByProjectIdRequest) SetProjectId(v string) *ListFunctionFilesByProjectIdRequest {
	s.ProjectId = &v
	return s
}

type ListFunctionFilesByProjectIdResponseBody struct {
	Files     []*ListFunctionFilesByProjectIdResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionFilesByProjectIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdResponseBody) SetFiles(v []*ListFunctionFilesByProjectIdResponseBodyFiles) *ListFunctionFilesByProjectIdResponseBody {
	s.Files = v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBody) SetRequestId(v string) *ListFunctionFilesByProjectIdResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionFilesByProjectIdResponseBodyFiles struct {
	ContentId   *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionFilesByProjectIdResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetContentId(v int64) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.ContentId = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetGmtCreate(v int64) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.GmtCreate = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetGmtModified(v int64) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.GmtModified = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetId(v int64) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Id = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetName(v string) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetStatus(v int32) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Status = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetType(v int32) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Type = &v
	return s
}

type ListFunctionFilesByProjectIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionFilesByProjectIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionFilesByProjectIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdResponse) SetHeaders(v map[string]*string) *ListFunctionFilesByProjectIdResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionFilesByProjectIdResponse) SetStatusCode(v int32) *ListFunctionFilesByProjectIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponse) SetBody(v *ListFunctionFilesByProjectIdResponseBody) *ListFunctionFilesByProjectIdResponse {
	s.Body = v
	return s
}

type ListMessageAcksRequest struct {
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	MessageId *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListMessageAcksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksRequest) GoString() string {
	return s.String()
}

func (s *ListMessageAcksRequest) SetDeviceId(v string) *ListMessageAcksRequest {
	s.DeviceId = &v
	return s
}

func (s *ListMessageAcksRequest) SetMessageId(v int64) *ListMessageAcksRequest {
	s.MessageId = &v
	return s
}

func (s *ListMessageAcksRequest) SetPageIndex(v int32) *ListMessageAcksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMessageAcksRequest) SetPageSize(v int32) *ListMessageAcksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageAcksRequest) SetProjectId(v string) *ListMessageAcksRequest {
	s.ProjectId = &v
	return s
}

type ListMessageAcksResponseBody struct {
	MessageAcks *ListMessageAcksResponseBodyMessageAcks `json:"MessageAcks,omitempty" xml:"MessageAcks,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMessageAcksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBody) SetMessageAcks(v *ListMessageAcksResponseBodyMessageAcks) *ListMessageAcksResponseBody {
	s.MessageAcks = v
	return s
}

func (s *ListMessageAcksResponseBody) SetRequestId(v string) *ListMessageAcksResponseBody {
	s.RequestId = &v
	return s
}

type ListMessageAcksResponseBodyMessageAcks struct {
	List       []*ListMessageAcksResponseBodyMessageAcksList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListMessageAcksResponseBodyMessageAcksPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListMessageAcksResponseBodyMessageAcks) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBodyMessageAcks) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBodyMessageAcks) SetList(v []*ListMessageAcksResponseBodyMessageAcksList) *ListMessageAcksResponseBodyMessageAcks {
	s.List = v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcks) SetPagination(v *ListMessageAcksResponseBodyMessageAcksPagination) *ListMessageAcksResponseBodyMessageAcks {
	s.Pagination = v
	return s
}

type ListMessageAcksResponseBodyMessageAcksList struct {
	AckTime  *int64  `json:"AckTime,omitempty" xml:"AckTime,omitempty"`
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Mid      *int64  `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s ListMessageAcksResponseBodyMessageAcksList) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBodyMessageAcksList) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBodyMessageAcksList) SetAckTime(v int64) *ListMessageAcksResponseBodyMessageAcksList {
	s.AckTime = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksList) SetDeviceId(v string) *ListMessageAcksResponseBodyMessageAcksList {
	s.DeviceId = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksList) SetMid(v int64) *ListMessageAcksResponseBodyMessageAcksList {
	s.Mid = &v
	return s
}

type ListMessageAcksResponseBodyMessageAcksPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListMessageAcksResponseBodyMessageAcksPagination) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBodyMessageAcksPagination) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBodyMessageAcksPagination) SetPageIndex(v int32) *ListMessageAcksResponseBodyMessageAcksPagination {
	s.PageIndex = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksPagination) SetPageSize(v int32) *ListMessageAcksResponseBodyMessageAcksPagination {
	s.PageSize = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksPagination) SetTotalCount(v int32) *ListMessageAcksResponseBodyMessageAcksPagination {
	s.TotalCount = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksPagination) SetTotalPageCount(v int32) *ListMessageAcksResponseBodyMessageAcksPagination {
	s.TotalPageCount = &v
	return s
}

type ListMessageAcksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMessageAcksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMessageAcksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponse) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponse) SetHeaders(v map[string]*string) *ListMessageAcksResponse {
	s.Headers = v
	return s
}

func (s *ListMessageAcksResponse) SetStatusCode(v int32) *ListMessageAcksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageAcksResponse) SetBody(v *ListMessageAcksResponseBody) *ListMessageAcksResponse {
	s.Body = v
	return s
}

type ListMessageReceiversRequest struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListMessageReceiversRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversRequest) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversRequest) SetMessageId(v string) *ListMessageReceiversRequest {
	s.MessageId = &v
	return s
}

func (s *ListMessageReceiversRequest) SetPageIndex(v int32) *ListMessageReceiversRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMessageReceiversRequest) SetPageSize(v int32) *ListMessageReceiversRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageReceiversRequest) SetProjectId(v string) *ListMessageReceiversRequest {
	s.ProjectId = &v
	return s
}

type ListMessageReceiversResponseBody struct {
	MessageReceivers *ListMessageReceiversResponseBodyMessageReceivers `json:"MessageReceivers,omitempty" xml:"MessageReceivers,omitempty" type:"Struct"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMessageReceiversResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponseBody) SetMessageReceivers(v *ListMessageReceiversResponseBodyMessageReceivers) *ListMessageReceiversResponseBody {
	s.MessageReceivers = v
	return s
}

func (s *ListMessageReceiversResponseBody) SetRequestId(v string) *ListMessageReceiversResponseBody {
	s.RequestId = &v
	return s
}

type ListMessageReceiversResponseBodyMessageReceivers struct {
	List       []*ListMessageReceiversResponseBodyMessageReceiversList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListMessageReceiversResponseBodyMessageReceiversPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListMessageReceiversResponseBodyMessageReceivers) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponseBodyMessageReceivers) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponseBodyMessageReceivers) SetList(v []*ListMessageReceiversResponseBodyMessageReceiversList) *ListMessageReceiversResponseBodyMessageReceivers {
	s.List = v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceivers) SetPagination(v *ListMessageReceiversResponseBodyMessageReceiversPagination) *ListMessageReceiversResponseBodyMessageReceivers {
	s.Pagination = v
	return s
}

type ListMessageReceiversResponseBodyMessageReceiversList struct {
	Mid   *int64  `json:"Mid,omitempty" xml:"Mid,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListMessageReceiversResponseBodyMessageReceiversList) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponseBodyMessageReceiversList) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponseBodyMessageReceiversList) SetMid(v int64) *ListMessageReceiversResponseBodyMessageReceiversList {
	s.Mid = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversList) SetType(v string) *ListMessageReceiversResponseBodyMessageReceiversList {
	s.Type = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversList) SetValue(v string) *ListMessageReceiversResponseBodyMessageReceiversList {
	s.Value = &v
	return s
}

type ListMessageReceiversResponseBodyMessageReceiversPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListMessageReceiversResponseBodyMessageReceiversPagination) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponseBodyMessageReceiversPagination) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponseBodyMessageReceiversPagination) SetPageIndex(v int32) *ListMessageReceiversResponseBodyMessageReceiversPagination {
	s.PageIndex = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversPagination) SetPageSize(v int32) *ListMessageReceiversResponseBodyMessageReceiversPagination {
	s.PageSize = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversPagination) SetTotalCount(v int32) *ListMessageReceiversResponseBodyMessageReceiversPagination {
	s.TotalCount = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversPagination) SetTotalPageCount(v int32) *ListMessageReceiversResponseBodyMessageReceiversPagination {
	s.TotalPageCount = &v
	return s
}

type ListMessageReceiversResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMessageReceiversResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMessageReceiversResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponse) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponse) SetHeaders(v map[string]*string) *ListMessageReceiversResponse {
	s.Headers = v
	return s
}

func (s *ListMessageReceiversResponse) SetStatusCode(v int32) *ListMessageReceiversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageReceiversResponse) SetBody(v *ListMessageReceiversResponseBody) *ListMessageReceiversResponse {
	s.Body = v
	return s
}

type ListNamespacesRequest struct {
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) SetAuthType(v string) *ListNamespacesRequest {
	s.AuthType = &v
	return s
}

func (s *ListNamespacesRequest) SetProjectId(v string) *ListNamespacesRequest {
	s.ProjectId = &v
	return s
}

type ListNamespacesResponseBody struct {
	Namespaces []*ListNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) SetNamespaces(v []*ListNamespacesResponseBodyNamespaces) *ListNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

type ListNamespacesResponseBodyNamespaces struct {
	AuthType    *int32  `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListNamespacesResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyNamespaces) SetAuthType(v int32) *ListNamespacesResponseBodyNamespaces {
	s.AuthType = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetDescription(v string) *ListNamespacesResponseBodyNamespaces {
	s.Description = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetGmtCreate(v int64) *ListNamespacesResponseBodyNamespaces {
	s.GmtCreate = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetGmtModified(v int64) *ListNamespacesResponseBodyNamespaces {
	s.GmtModified = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetName(v string) *ListNamespacesResponseBodyNamespaces {
	s.Name = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetNamespace(v string) *ListNamespacesResponseBodyNamespaces {
	s.Namespace = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetProjectId(v string) *ListNamespacesResponseBodyNamespaces {
	s.ProjectId = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetUserId(v string) *ListNamespacesResponseBodyNamespaces {
	s.UserId = &v
	return s
}

type ListNamespacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponse) SetHeaders(v map[string]*string) *ListNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacesResponse) SetStatusCode(v int32) *ListNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

type ListOfflineMessagesRequest struct {
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListOfflineMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesRequest) SetPageIndex(v int32) *ListOfflineMessagesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListOfflineMessagesRequest) SetPageSize(v int32) *ListOfflineMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOfflineMessagesRequest) SetProjectId(v string) *ListOfflineMessagesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListOfflineMessagesRequest) SetType(v string) *ListOfflineMessagesRequest {
	s.Type = &v
	return s
}

func (s *ListOfflineMessagesRequest) SetValue(v string) *ListOfflineMessagesRequest {
	s.Value = &v
	return s
}

type ListOfflineMessagesResponseBody struct {
	OfflineMessages *ListOfflineMessagesResponseBodyOfflineMessages `json:"OfflineMessages,omitempty" xml:"OfflineMessages,omitempty" type:"Struct"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOfflineMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponseBody) SetOfflineMessages(v *ListOfflineMessagesResponseBodyOfflineMessages) *ListOfflineMessagesResponseBody {
	s.OfflineMessages = v
	return s
}

func (s *ListOfflineMessagesResponseBody) SetRequestId(v string) *ListOfflineMessagesResponseBody {
	s.RequestId = &v
	return s
}

type ListOfflineMessagesResponseBodyOfflineMessages struct {
	List       []*ListOfflineMessagesResponseBodyOfflineMessagesList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListOfflineMessagesResponseBodyOfflineMessagesPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListOfflineMessagesResponseBodyOfflineMessages) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponseBodyOfflineMessages) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponseBodyOfflineMessages) SetList(v []*ListOfflineMessagesResponseBodyOfflineMessagesList) *ListOfflineMessagesResponseBodyOfflineMessages {
	s.List = v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessages) SetPagination(v *ListOfflineMessagesResponseBodyOfflineMessagesPagination) *ListOfflineMessagesResponseBodyOfflineMessages {
	s.Pagination = v
	return s
}

type ListOfflineMessagesResponseBodyOfflineMessagesList struct {
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate   *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Mid         *int64 `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s ListOfflineMessagesResponseBodyOfflineMessagesList) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponseBodyOfflineMessagesList) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesList) SetExpiredTime(v int64) *ListOfflineMessagesResponseBodyOfflineMessagesList {
	s.ExpiredTime = &v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesList) SetGmtCreate(v int64) *ListOfflineMessagesResponseBodyOfflineMessagesList {
	s.GmtCreate = &v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesList) SetMid(v int64) *ListOfflineMessagesResponseBodyOfflineMessagesList {
	s.Mid = &v
	return s
}

type ListOfflineMessagesResponseBodyOfflineMessagesPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListOfflineMessagesResponseBodyOfflineMessagesPagination) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponseBodyOfflineMessagesPagination) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesPagination) SetPageIndex(v int32) *ListOfflineMessagesResponseBodyOfflineMessagesPagination {
	s.PageIndex = &v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesPagination) SetPageSize(v int32) *ListOfflineMessagesResponseBodyOfflineMessagesPagination {
	s.PageSize = &v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesPagination) SetTotalCount(v int32) *ListOfflineMessagesResponseBodyOfflineMessagesPagination {
	s.TotalCount = &v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesPagination) SetTotalPageCount(v int32) *ListOfflineMessagesResponseBodyOfflineMessagesPagination {
	s.TotalPageCount = &v
	return s
}

type ListOfflineMessagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOfflineMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfflineMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponse) SetHeaders(v map[string]*string) *ListOfflineMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListOfflineMessagesResponse) SetStatusCode(v int32) *ListOfflineMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOfflineMessagesResponse) SetBody(v *ListOfflineMessagesResponseBody) *ListOfflineMessagesResponse {
	s.Body = v
	return s
}

type ListOpenAccountLinksRequest struct {
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp        *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	OpenId     *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListOpenAccountLinksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksRequest) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksRequest) SetIdentityId(v string) *ListOpenAccountLinksRequest {
	s.IdentityId = &v
	return s
}

func (s *ListOpenAccountLinksRequest) SetIdp(v string) *ListOpenAccountLinksRequest {
	s.Idp = &v
	return s
}

func (s *ListOpenAccountLinksRequest) SetOpenId(v string) *ListOpenAccountLinksRequest {
	s.OpenId = &v
	return s
}

func (s *ListOpenAccountLinksRequest) SetProjectId(v string) *ListOpenAccountLinksRequest {
	s.ProjectId = &v
	return s
}

type ListOpenAccountLinksResponseBody struct {
	OpenAccounts []*ListOpenAccountLinksResponseBodyOpenAccounts `json:"OpenAccounts,omitempty" xml:"OpenAccounts,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOpenAccountLinksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksResponseBody) SetOpenAccounts(v []*ListOpenAccountLinksResponseBodyOpenAccounts) *ListOpenAccountLinksResponseBody {
	s.OpenAccounts = v
	return s
}

func (s *ListOpenAccountLinksResponseBody) SetRequestId(v string) *ListOpenAccountLinksResponseBody {
	s.RequestId = &v
	return s
}

type ListOpenAccountLinksResponseBodyOpenAccounts struct {
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CreateAccessKey *string `json:"CreateAccessKey,omitempty" xml:"CreateAccessKey,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	IdentityId      *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp             *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	LoginId         *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenId          *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOpenAccountLinksResponseBodyOpenAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksResponseBodyOpenAccounts) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetAliyunId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.AliyunId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetCreateAccessKey(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.CreateAccessKey = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetDisplayName(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetIdentityId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.IdentityId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetIdp(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Idp = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetLoginId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.LoginId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetMobile(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Mobile = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetOpenId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.OpenId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetRegion(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Region = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetStatus(v int32) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Status = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetType(v int32) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Type = &v
	return s
}

type ListOpenAccountLinksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOpenAccountLinksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenAccountLinksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksResponse) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksResponse) SetHeaders(v map[string]*string) *ListOpenAccountLinksResponse {
	s.Headers = v
	return s
}

func (s *ListOpenAccountLinksResponse) SetStatusCode(v int32) *ListOpenAccountLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenAccountLinksResponse) SetBody(v *ListOpenAccountLinksResponseBody) *ListOpenAccountLinksResponse {
	s.Body = v
	return s
}

type ListOpenAccountsRequest struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Length      *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Start       *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListOpenAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsRequest) SetDisplayName(v string) *ListOpenAccountsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListOpenAccountsRequest) SetEmail(v string) *ListOpenAccountsRequest {
	s.Email = &v
	return s
}

func (s *ListOpenAccountsRequest) SetLength(v int32) *ListOpenAccountsRequest {
	s.Length = &v
	return s
}

func (s *ListOpenAccountsRequest) SetMobile(v string) *ListOpenAccountsRequest {
	s.Mobile = &v
	return s
}

func (s *ListOpenAccountsRequest) SetProjectId(v string) *ListOpenAccountsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListOpenAccountsRequest) SetStart(v int32) *ListOpenAccountsRequest {
	s.Start = &v
	return s
}

type ListOpenAccountsResponseBody struct {
	OpenAccounts []*ListOpenAccountsResponseBodyOpenAccounts `json:"OpenAccounts,omitempty" xml:"OpenAccounts,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOpenAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsResponseBody) SetOpenAccounts(v []*ListOpenAccountsResponseBodyOpenAccounts) *ListOpenAccountsResponseBody {
	s.OpenAccounts = v
	return s
}

func (s *ListOpenAccountsResponseBody) SetRequestId(v string) *ListOpenAccountsResponseBody {
	s.RequestId = &v
	return s
}

type ListOpenAccountsResponseBodyOpenAccounts struct {
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CreateAccessKey *string `json:"CreateAccessKey,omitempty" xml:"CreateAccessKey,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	IdentityId      *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp             *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	LoginId         *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenId          *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOpenAccountsResponseBodyOpenAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsResponseBodyOpenAccounts) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetAliyunId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.AliyunId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetCreateAccessKey(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.CreateAccessKey = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetDisplayName(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetIdentityId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.IdentityId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetIdp(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Idp = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetLoginId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.LoginId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetMobile(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Mobile = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetOpenId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.OpenId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetRegion(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Region = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetStatus(v int32) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Status = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetType(v int32) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Type = &v
	return s
}

type ListOpenAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOpenAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsResponse) SetHeaders(v map[string]*string) *ListOpenAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListOpenAccountsResponse) SetStatusCode(v int32) *ListOpenAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenAccountsResponse) SetBody(v *ListOpenAccountsResponseBody) *ListOpenAccountsResponse {
	s.Body = v
	return s
}

type ListPreChecksResponseBody struct {
	PreChecks []*ListPreChecksResponseBodyPreChecks `json:"PreChecks,omitempty" xml:"PreChecks,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPreChecksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPreChecksResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreChecksResponseBody) SetPreChecks(v []*ListPreChecksResponseBodyPreChecks) *ListPreChecksResponseBody {
	s.PreChecks = v
	return s
}

func (s *ListPreChecksResponseBody) SetRequestId(v string) *ListPreChecksResponseBody {
	s.RequestId = &v
	return s
}

type ListPreChecksResponseBodyPreChecks struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Link  *string `json:"Link,omitempty" xml:"Link,omitempty"`
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListPreChecksResponseBodyPreChecks) String() string {
	return tea.Prettify(s)
}

func (s ListPreChecksResponseBodyPreChecks) GoString() string {
	return s.String()
}

func (s *ListPreChecksResponseBodyPreChecks) SetKey(v string) *ListPreChecksResponseBodyPreChecks {
	s.Key = &v
	return s
}

func (s *ListPreChecksResponseBodyPreChecks) SetLink(v string) *ListPreChecksResponseBodyPreChecks {
	s.Link = &v
	return s
}

func (s *ListPreChecksResponseBodyPreChecks) SetPrice(v string) *ListPreChecksResponseBodyPreChecks {
	s.Price = &v
	return s
}

func (s *ListPreChecksResponseBodyPreChecks) SetState(v string) *ListPreChecksResponseBodyPreChecks {
	s.State = &v
	return s
}

type ListPreChecksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPreChecksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPreChecksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPreChecksResponse) GoString() string {
	return s.String()
}

func (s *ListPreChecksResponse) SetHeaders(v map[string]*string) *ListPreChecksResponse {
	s.Headers = v
	return s
}

func (s *ListPreChecksResponse) SetStatusCode(v int32) *ListPreChecksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPreChecksResponse) SetBody(v *ListPreChecksResponseBody) *ListPreChecksResponse {
	s.Body = v
	return s
}

type ListProjectAppsRequest struct {
	Keywords  *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectAppsRequest) SetKeywords(v string) *ListProjectAppsRequest {
	s.Keywords = &v
	return s
}

func (s *ListProjectAppsRequest) SetPageIndex(v int32) *ListProjectAppsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListProjectAppsRequest) SetPageSize(v int32) *ListProjectAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectAppsRequest) SetProjectId(v string) *ListProjectAppsRequest {
	s.ProjectId = &v
	return s
}

type ListProjectAppsResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListProjectAppsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListProjectAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectAppsResponseBody) SetRequestId(v string) *ListProjectAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectAppsResponseBody) SetResult(v *ListProjectAppsResponseBodyResult) *ListProjectAppsResponseBody {
	s.Result = v
	return s
}

type ListProjectAppsResponseBodyResult struct {
	ProjectApps []*ListProjectAppsResponseBodyResultProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	TotalCount  *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int32                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListProjectAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListProjectAppsResponseBodyResult) SetProjectApps(v []*ListProjectAppsResponseBodyResultProjectApps) *ListProjectAppsResponseBodyResult {
	s.ProjectApps = v
	return s
}

func (s *ListProjectAppsResponseBodyResult) SetTotalCount(v int32) *ListProjectAppsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListProjectAppsResponseBodyResult) SetTotalPage(v int32) *ListProjectAppsResponseBodyResult {
	s.TotalPage = &v
	return s
}

type ListProjectAppsResponseBodyResultProjectApps struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPkgName  *string `json:"AppPkgName,omitempty" xml:"AppPkgName,omitempty"`
	AppSecret   *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OsType      *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListProjectAppsResponseBodyResultProjectApps) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsResponseBodyResultProjectApps) GoString() string {
	return s.String()
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppId(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppId = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppKey(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppKey = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppName(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppName = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppPkgName(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppPkgName = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppSecret(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppSecret = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetGmtCreate(v int64) *ListProjectAppsResponseBodyResultProjectApps {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetGmtModified(v int64) *ListProjectAppsResponseBodyResultProjectApps {
	s.GmtModified = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetId(v int64) *ListProjectAppsResponseBodyResultProjectApps {
	s.Id = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetOsType(v int32) *ListProjectAppsResponseBodyResultProjectApps {
	s.OsType = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetProjectId(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.ProjectId = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetStatus(v int32) *ListProjectAppsResponseBodyResultProjectApps {
	s.Status = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetUserId(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.UserId = &v
	return s
}

type ListProjectAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectAppsResponse) SetHeaders(v map[string]*string) *ListProjectAppsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectAppsResponse) SetStatusCode(v int32) *ListProjectAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectAppsResponse) SetBody(v *ListProjectAppsResponseBody) *ListProjectAppsResponse {
	s.Body = v
	return s
}

type ListProjectsResponseBody struct {
	Projects  []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ListProjectsResponseBodyProjects struct {
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetCreator(v string) *ListProjectsResponseBodyProjects {
	s.Creator = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtCreate(v int64) *ListProjectsResponseBodyProjects {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtModified(v int64) *ListProjectsResponseBodyProjects {
	s.GmtModified = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetId(v int64) *ListProjectsResponseBodyProjects {
	s.Id = &v
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

func (s *ListProjectsResponseBodyProjects) SetStatus(v int32) *ListProjectsResponseBodyProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetUserId(v string) *ListProjectsResponseBodyProjects {
	s.UserId = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListRpcServicesRequest struct {
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListRpcServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesRequest) GoString() string {
	return s.String()
}

func (s *ListRpcServicesRequest) SetPageIndex(v int32) *ListRpcServicesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListRpcServicesRequest) SetPageSize(v int32) *ListRpcServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRpcServicesRequest) SetProjectId(v string) *ListRpcServicesRequest {
	s.ProjectId = &v
	return s
}

type ListRpcServicesResponseBody struct {
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RpcServices *ListRpcServicesResponseBodyRpcServices `json:"RpcServices,omitempty" xml:"RpcServices,omitempty" type:"Struct"`
}

func (s ListRpcServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponseBody) SetRequestId(v string) *ListRpcServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRpcServicesResponseBody) SetRpcServices(v *ListRpcServicesResponseBodyRpcServices) *ListRpcServicesResponseBody {
	s.RpcServices = v
	return s
}

type ListRpcServicesResponseBodyRpcServices struct {
	List       []*ListRpcServicesResponseBodyRpcServicesList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListRpcServicesResponseBodyRpcServicesPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListRpcServicesResponseBodyRpcServices) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponseBodyRpcServices) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponseBodyRpcServices) SetList(v []*ListRpcServicesResponseBodyRpcServicesList) *ListRpcServicesResponseBodyRpcServices {
	s.List = v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServices) SetPagination(v *ListRpcServicesResponseBodyRpcServicesPagination) *ListRpcServicesResponseBodyRpcServices {
	s.Pagination = v
	return s
}

type ListRpcServicesResponseBodyRpcServicesList struct {
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	IsDelete      *string `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	MethodName    *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	Params        *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VersionCode   *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListRpcServicesResponseBodyRpcServicesList) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponseBodyRpcServicesList) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetAppKey(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.AppKey = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetGmtCreate(v int64) *ListRpcServicesResponseBodyRpcServicesList {
	s.GmtCreate = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetGmtModified(v int64) *ListRpcServicesResponseBodyRpcServicesList {
	s.GmtModified = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetGroupName(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.GroupName = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetId(v int64) *ListRpcServicesResponseBodyRpcServicesList {
	s.Id = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetInterfaceName(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.InterfaceName = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetIsDelete(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.IsDelete = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetMethodName(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.MethodName = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetParams(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.Params = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetType(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.Type = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetVersionCode(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.VersionCode = &v
	return s
}

type ListRpcServicesResponseBodyRpcServicesPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListRpcServicesResponseBodyRpcServicesPagination) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponseBodyRpcServicesPagination) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponseBodyRpcServicesPagination) SetPageIndex(v int32) *ListRpcServicesResponseBodyRpcServicesPagination {
	s.PageIndex = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesPagination) SetPageSize(v int32) *ListRpcServicesResponseBodyRpcServicesPagination {
	s.PageSize = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesPagination) SetTotalCount(v int32) *ListRpcServicesResponseBodyRpcServicesPagination {
	s.TotalCount = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesPagination) SetTotalPageCount(v int32) *ListRpcServicesResponseBodyRpcServicesPagination {
	s.TotalPageCount = &v
	return s
}

type ListRpcServicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRpcServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRpcServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponse) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponse) SetHeaders(v map[string]*string) *ListRpcServicesResponse {
	s.Headers = v
	return s
}

func (s *ListRpcServicesResponse) SetStatusCode(v int32) *ListRpcServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRpcServicesResponse) SetBody(v *ListRpcServicesResponseBody) *ListRpcServicesResponse {
	s.Body = v
	return s
}

type ListSchemaSubscribesRequest struct {
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListSchemaSubscribesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesRequest) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesRequest) SetDeviceModel(v string) *ListSchemaSubscribesRequest {
	s.DeviceModel = &v
	return s
}

func (s *ListSchemaSubscribesRequest) SetPageIndex(v int32) *ListSchemaSubscribesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListSchemaSubscribesRequest) SetPageSize(v int32) *ListSchemaSubscribesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchemaSubscribesRequest) SetProjectId(v string) *ListSchemaSubscribesRequest {
	s.ProjectId = &v
	return s
}

type ListSchemaSubscribesResponseBody struct {
	PageList  []*ListSchemaSubscribesResponseBodyPageList `json:"PageList,omitempty" xml:"PageList,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSchemaSubscribesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBody) SetPageList(v []*ListSchemaSubscribesResponseBodyPageList) *ListSchemaSubscribesResponseBody {
	s.PageList = v
	return s
}

func (s *ListSchemaSubscribesResponseBody) SetRequestId(v string) *ListSchemaSubscribesResponseBody {
	s.RequestId = &v
	return s
}

type ListSchemaSubscribesResponseBodyPageList struct {
	List       []*ListSchemaSubscribesResponseBodyPageListList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListSchemaSubscribesResponseBodyPageListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListSchemaSubscribesResponseBodyPageList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBodyPageList) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBodyPageList) SetList(v []*ListSchemaSubscribesResponseBodyPageListList) *ListSchemaSubscribesResponseBodyPageList {
	s.List = v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageList) SetPagination(v *ListSchemaSubscribesResponseBodyPageListPagination) *ListSchemaSubscribesResponseBodyPageList {
	s.Pagination = v
	return s
}

type ListSchemaSubscribesResponseBodyPageListList struct {
	DeviceModel    *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId  *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ValiditySchema *string `json:"ValiditySchema,omitempty" xml:"ValiditySchema,omitempty"`
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSchemaSubscribesResponseBodyPageListList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBodyPageListList) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetDeviceModel(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.DeviceModel = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetDeviceModelId(v int64) *ListSchemaSubscribesResponseBodyPageListList {
	s.DeviceModelId = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetGmtCreate(v int64) *ListSchemaSubscribesResponseBodyPageListList {
	s.GmtCreate = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetGmtModified(v int64) *ListSchemaSubscribesResponseBodyPageListList {
	s.GmtModified = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetId(v int64) *ListSchemaSubscribesResponseBodyPageListList {
	s.Id = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetNamespace(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.Namespace = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetProjectId(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.ProjectId = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetValiditySchema(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.ValiditySchema = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetVersion(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.Version = &v
	return s
}

type ListSchemaSubscribesResponseBodyPageListPagination struct {
	HasNextPage    *bool  `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SimpleSign     *bool  `json:"SimpleSign,omitempty" xml:"SimpleSign,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListSchemaSubscribesResponseBodyPageListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBodyPageListPagination) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetHasNextPage(v bool) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.HasNextPage = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetPageIndex(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetPageSize(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.PageSize = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetSimpleSign(v bool) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.SimpleSign = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetTotalCount(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetTotalPageCount(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.TotalPageCount = &v
	return s
}

type ListSchemaSubscribesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSchemaSubscribesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSchemaSubscribesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponse) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponse) SetHeaders(v map[string]*string) *ListSchemaSubscribesResponse {
	s.Headers = v
	return s
}

func (s *ListSchemaSubscribesResponse) SetStatusCode(v int32) *ListSchemaSubscribesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchemaSubscribesResponse) SetBody(v *ListSchemaSubscribesResponseBody) *ListSchemaSubscribesResponse {
	s.Body = v
	return s
}

type ListServicesResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceList []*string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServiceList(v []*string) *ListServicesResponseBody {
	s.ServiceList = v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListShadowSchemaDeviceModelsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListShadowSchemaDeviceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemaDeviceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListShadowSchemaDeviceModelsRequest) SetProjectId(v string) *ListShadowSchemaDeviceModelsRequest {
	s.ProjectId = &v
	return s
}

type ListShadowSchemaDeviceModelsResponseBody struct {
	ModelList []*ListShadowSchemaDeviceModelsResponseBodyModelList `json:"ModelList,omitempty" xml:"ModelList,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListShadowSchemaDeviceModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemaDeviceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListShadowSchemaDeviceModelsResponseBody) SetModelList(v []*ListShadowSchemaDeviceModelsResponseBodyModelList) *ListShadowSchemaDeviceModelsResponseBody {
	s.ModelList = v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBody) SetRequestId(v string) *ListShadowSchemaDeviceModelsResponseBody {
	s.RequestId = &v
	return s
}

type ListShadowSchemaDeviceModelsResponseBodyModelList struct {
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SecurityChip      *int32  `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
}

func (s ListShadowSchemaDeviceModelsResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemaDeviceModelsResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetCanCreateDeviceId(v int32) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.CanCreateDeviceId = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDescription(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.Description = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceBrand(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceBrand = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceModel(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceModel = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceModelId(v int64) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceModelId = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceType(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceType = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetHardwareType(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.HardwareType = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetInitUsageType(v int32) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.InitUsageType = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetInitUsageTypeDesc(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetOsPlatform(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.OsPlatform = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetProjectId(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.ProjectId = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetSecurityChip(v int32) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.SecurityChip = &v
	return s
}

type ListShadowSchemaDeviceModelsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListShadowSchemaDeviceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListShadowSchemaDeviceModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemaDeviceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListShadowSchemaDeviceModelsResponse) SetHeaders(v map[string]*string) *ListShadowSchemaDeviceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponse) SetStatusCode(v int32) *ListShadowSchemaDeviceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponse) SetBody(v *ListShadowSchemaDeviceModelsResponseBody) *ListShadowSchemaDeviceModelsResponse {
	s.Body = v
	return s
}

type ListShadowSchemasRequest struct {
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	QueryValue *string `json:"QueryValue,omitempty" xml:"QueryValue,omitempty"`
}

func (s ListShadowSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasRequest) SetPageIndex(v int32) *ListShadowSchemasRequest {
	s.PageIndex = &v
	return s
}

func (s *ListShadowSchemasRequest) SetPageSize(v int32) *ListShadowSchemasRequest {
	s.PageSize = &v
	return s
}

func (s *ListShadowSchemasRequest) SetProjectId(v string) *ListShadowSchemasRequest {
	s.ProjectId = &v
	return s
}

func (s *ListShadowSchemasRequest) SetQueryType(v string) *ListShadowSchemasRequest {
	s.QueryType = &v
	return s
}

func (s *ListShadowSchemasRequest) SetQueryValue(v string) *ListShadowSchemasRequest {
	s.QueryValue = &v
	return s
}

type ListShadowSchemasResponseBody struct {
	PageList  *ListShadowSchemasResponseBodyPageList `json:"PageList,omitempty" xml:"PageList,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListShadowSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBody) SetPageList(v *ListShadowSchemasResponseBodyPageList) *ListShadowSchemasResponseBody {
	s.PageList = v
	return s
}

func (s *ListShadowSchemasResponseBody) SetRequestId(v string) *ListShadowSchemasResponseBody {
	s.RequestId = &v
	return s
}

type ListShadowSchemasResponseBodyPageList struct {
	List       []*ListShadowSchemasResponseBodyPageListList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListShadowSchemasResponseBodyPageListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListShadowSchemasResponseBodyPageList) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBodyPageList) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBodyPageList) SetList(v []*ListShadowSchemasResponseBodyPageListList) *ListShadowSchemasResponseBodyPageList {
	s.List = v
	return s
}

func (s *ListShadowSchemasResponseBodyPageList) SetPagination(v *ListShadowSchemasResponseBodyPageListPagination) *ListShadowSchemasResponseBodyPageList {
	s.Pagination = v
	return s
}

type ListShadowSchemasResponseBodyPageListList struct {
	AuthType      *int32  `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AuthTypeDesc  *string `json:"AuthTypeDesc,omitempty" xml:"AuthTypeDesc,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModuleSchema  *string `json:"ModuleSchema,omitempty" xml:"ModuleSchema,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListShadowSchemasResponseBodyPageListList) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBodyPageListList) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBodyPageListList) SetAuthType(v int32) *ListShadowSchemasResponseBodyPageListList {
	s.AuthType = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetAuthTypeDesc(v string) *ListShadowSchemasResponseBodyPageListList {
	s.AuthTypeDesc = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetDeviceModel(v string) *ListShadowSchemasResponseBodyPageListList {
	s.DeviceModel = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetDeviceModelId(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.DeviceModelId = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetGmtCreate(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.GmtCreate = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetGmtModified(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.GmtModified = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetId(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.Id = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetModuleSchema(v string) *ListShadowSchemasResponseBodyPageListList {
	s.ModuleSchema = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetNamespace(v string) *ListShadowSchemasResponseBodyPageListList {
	s.Namespace = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetProjectId(v string) *ListShadowSchemasResponseBodyPageListList {
	s.ProjectId = &v
	return s
}

type ListShadowSchemasResponseBodyPageListPagination struct {
	HasNextPage    *bool  `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SimpleSign     *bool  `json:"SimpleSign,omitempty" xml:"SimpleSign,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListShadowSchemasResponseBodyPageListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBodyPageListPagination) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetHasNextPage(v bool) *ListShadowSchemasResponseBodyPageListPagination {
	s.HasNextPage = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetPageIndex(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetPageSize(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.PageSize = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetSimpleSign(v bool) *ListShadowSchemasResponseBodyPageListPagination {
	s.SimpleSign = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetTotalCount(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetTotalPageCount(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.TotalPageCount = &v
	return s
}

type ListShadowSchemasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListShadowSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListShadowSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponse) SetHeaders(v map[string]*string) *ListShadowSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListShadowSchemasResponse) SetStatusCode(v int32) *ListShadowSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShadowSchemasResponse) SetBody(v *ListShadowSchemasResponseBody) *ListShadowSchemasResponse {
	s.Body = v
	return s
}

type ListSupportFeaturesResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportFeatures []*ListSupportFeaturesResponseBodySupportFeatures `json:"SupportFeatures,omitempty" xml:"SupportFeatures,omitempty" type:"Repeated"`
}

func (s ListSupportFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSupportFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportFeaturesResponseBody) SetRequestId(v string) *ListSupportFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportFeaturesResponseBody) SetSupportFeatures(v []*ListSupportFeaturesResponseBodySupportFeatures) *ListSupportFeaturesResponseBody {
	s.SupportFeatures = v
	return s
}

type ListSupportFeaturesResponseBodySupportFeatures struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListSupportFeaturesResponseBodySupportFeatures) String() string {
	return tea.Prettify(s)
}

func (s ListSupportFeaturesResponseBodySupportFeatures) GoString() string {
	return s.String()
}

func (s *ListSupportFeaturesResponseBodySupportFeatures) SetName(v string) *ListSupportFeaturesResponseBodySupportFeatures {
	s.Name = &v
	return s
}

type ListSupportFeaturesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSupportFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSupportFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSupportFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListSupportFeaturesResponse) SetHeaders(v map[string]*string) *ListSupportFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListSupportFeaturesResponse) SetStatusCode(v int32) *ListSupportFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportFeaturesResponse) SetBody(v *ListSupportFeaturesResponseBody) *ListSupportFeaturesResponse {
	s.Body = v
	return s
}

type ListTriggersRequest struct {
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) SetNamespace(v string) *ListTriggersRequest {
	s.Namespace = &v
	return s
}

func (s *ListTriggersRequest) SetPageIndex(v int32) *ListTriggersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTriggersRequest) SetPageSize(v int32) *ListTriggersRequest {
	s.PageSize = &v
	return s
}

func (s *ListTriggersRequest) SetProjectId(v string) *ListTriggersRequest {
	s.ProjectId = &v
	return s
}

type ListTriggersResponseBody struct {
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TriggerList *ListTriggersResponseBodyTriggerList `json:"TriggerList,omitempty" xml:"TriggerList,omitempty" type:"Struct"`
}

func (s ListTriggersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBody) SetRequestId(v string) *ListTriggersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTriggersResponseBody) SetTriggerList(v *ListTriggersResponseBodyTriggerList) *ListTriggersResponseBody {
	s.TriggerList = v
	return s
}

type ListTriggersResponseBodyTriggerList struct {
	Pagination *ListTriggersResponseBodyTriggerListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Triggers   []*ListTriggersResponseBodyTriggerListTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListTriggersResponseBodyTriggerList) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggerList) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggerList) SetPagination(v *ListTriggersResponseBodyTriggerListPagination) *ListTriggersResponseBodyTriggerList {
	s.Pagination = v
	return s
}

func (s *ListTriggersResponseBodyTriggerList) SetTriggers(v []*ListTriggersResponseBodyTriggerListTriggers) *ListTriggersResponseBodyTriggerList {
	s.Triggers = v
	return s
}

type ListTriggersResponseBodyTriggerListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListTriggersResponseBodyTriggerListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggerListPagination) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggerListPagination) SetPageIndex(v int32) *ListTriggersResponseBodyTriggerListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListPagination) SetPageSize(v int32) *ListTriggersResponseBodyTriggerListPagination {
	s.PageSize = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListPagination) SetTotalCount(v int32) *ListTriggersResponseBodyTriggerListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListPagination) SetTotalPageCount(v int32) *ListTriggersResponseBodyTriggerListPagination {
	s.TotalPageCount = &v
	return s
}

type ListTriggersResponseBodyTriggerListTriggers struct {
	ChainedFunctionIds *string                                                 `json:"ChainedFunctionIds,omitempty" xml:"ChainedFunctionIds,omitempty"`
	Functions          []*ListTriggersResponseBodyTriggerListTriggersFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	GmtCreate          *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *int64                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	InvocationMode     *int32                                                  `json:"InvocationMode,omitempty" xml:"InvocationMode,omitempty"`
	Namespace          *string                                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Production         *int32                                                  `json:"Production,omitempty" xml:"Production,omitempty"`
	Sandbox            *int32                                                  `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Source             *string                                                 `json:"Source,omitempty" xml:"Source,omitempty"`
	Status             *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *int32                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTriggersResponseBodyTriggerListTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggerListTriggers) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetChainedFunctionIds(v string) *ListTriggersResponseBodyTriggerListTriggers {
	s.ChainedFunctionIds = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetFunctions(v []*ListTriggersResponseBodyTriggerListTriggersFunctions) *ListTriggersResponseBodyTriggerListTriggers {
	s.Functions = v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetGmtCreate(v int64) *ListTriggersResponseBodyTriggerListTriggers {
	s.GmtCreate = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetGmtModified(v int64) *ListTriggersResponseBodyTriggerListTriggers {
	s.GmtModified = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetId(v int64) *ListTriggersResponseBodyTriggerListTriggers {
	s.Id = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetInvocationMode(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.InvocationMode = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetNamespace(v string) *ListTriggersResponseBodyTriggerListTriggers {
	s.Namespace = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetProduction(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Production = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetSandbox(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Sandbox = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetSource(v string) *ListTriggersResponseBodyTriggerListTriggers {
	s.Source = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetStatus(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Status = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetType(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Type = &v
	return s
}

type ListTriggersResponseBodyTriggerListTriggersFunctions struct {
	FileId      *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTriggersResponseBodyTriggerListTriggersFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggerListTriggersFunctions) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetFileId(v int64) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.FileId = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetFileName(v string) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.FileName = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetGmtCreate(v int64) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.GmtCreate = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetGmtModified(v int64) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.GmtModified = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetId(v int64) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.Id = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetName(v string) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.Name = &v
	return s
}

type ListTriggersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListTriggersResponse) SetHeaders(v map[string]*string) *ListTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListTriggersResponse) SetStatusCode(v int32) *ListTriggersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTriggersResponse) SetBody(v *ListTriggersResponseBody) *ListTriggersResponse {
	s.Body = v
	return s
}

type ListUpstreamAppKeyRelationsRequest struct {
	AppServerId *int64  `json:"AppServerId,omitempty" xml:"AppServerId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListUpstreamAppKeyRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsRequest) SetAppServerId(v int64) *ListUpstreamAppKeyRelationsRequest {
	s.AppServerId = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsRequest) SetPageIndex(v int32) *ListUpstreamAppKeyRelationsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsRequest) SetPageSize(v int32) *ListUpstreamAppKeyRelationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsRequest) SetProjectId(v string) *ListUpstreamAppKeyRelationsRequest {
	s.ProjectId = &v
	return s
}

type ListUpstreamAppKeyRelationsResponseBody struct {
	RelationList *ListUpstreamAppKeyRelationsResponseBodyRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Struct"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUpstreamAppKeyRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBody) SetRelationList(v *ListUpstreamAppKeyRelationsResponseBodyRelationList) *ListUpstreamAppKeyRelationsResponseBody {
	s.RelationList = v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBody) SetRequestId(v string) *ListUpstreamAppKeyRelationsResponseBody {
	s.RequestId = &v
	return s
}

type ListUpstreamAppKeyRelationsResponseBodyRelationList struct {
	List       []*ListUpstreamAppKeyRelationsResponseBodyRelationListList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationList) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationList) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationList) SetList(v []*ListUpstreamAppKeyRelationsResponseBodyRelationListList) *ListUpstreamAppKeyRelationsResponseBodyRelationList {
	s.List = v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationList) SetPagination(v *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) *ListUpstreamAppKeyRelationsResponseBodyRelationList {
	s.Pagination = v
	return s
}

type ListUpstreamAppKeyRelationsResponseBodyRelationListList struct {
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPackage *string `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PAppKey    *string `json:"PAppKey,omitempty" xml:"PAppKey,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationListList) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationListList) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetAppKey(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.AppKey = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetAppName(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.AppName = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetAppPackage(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.AppPackage = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetGmtCreate(v int64) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.GmtCreate = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetId(v int64) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.Id = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetPAppKey(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.PAppKey = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetProjectId(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.ProjectId = &v
	return s
}

type ListUpstreamAppKeyRelationsResponseBodyRelationListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) SetPageIndex(v int32) *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) SetPageSize(v int32) *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) SetTotalCount(v int32) *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) SetTotalPageCount(v int32) *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination {
	s.TotalPageCount = &v
	return s
}

type ListUpstreamAppKeyRelationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUpstreamAppKeyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUpstreamAppKeyRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponse) SetHeaders(v map[string]*string) *ListUpstreamAppKeyRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponse) SetStatusCode(v int32) *ListUpstreamAppKeyRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponse) SetBody(v *ListUpstreamAppKeyRelationsResponseBody) *ListUpstreamAppKeyRelationsResponse {
	s.Body = v
	return s
}

type ListUpstreamAppServersRequest struct {
	PageIndex *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListUpstreamAppServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersRequest) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersRequest) SetPageIndex(v string) *ListUpstreamAppServersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListUpstreamAppServersRequest) SetPageSize(v string) *ListUpstreamAppServersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamAppServersRequest) SetProjectId(v string) *ListUpstreamAppServersRequest {
	s.ProjectId = &v
	return s
}

type ListUpstreamAppServersResponseBody struct {
	AppServers *ListUpstreamAppServersResponseBodyAppServers `json:"AppServers,omitempty" xml:"AppServers,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUpstreamAppServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBody) SetAppServers(v *ListUpstreamAppServersResponseBodyAppServers) *ListUpstreamAppServersResponseBody {
	s.AppServers = v
	return s
}

func (s *ListUpstreamAppServersResponseBody) SetRequestId(v string) *ListUpstreamAppServersResponseBody {
	s.RequestId = &v
	return s
}

type ListUpstreamAppServersResponseBodyAppServers struct {
	List       []*ListUpstreamAppServersResponseBodyAppServersList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Pagination *ListUpstreamAppServersResponseBodyAppServersPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s ListUpstreamAppServersResponseBodyAppServers) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBodyAppServers) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBodyAppServers) SetList(v []*ListUpstreamAppServersResponseBodyAppServersList) *ListUpstreamAppServersResponseBodyAppServers {
	s.List = v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServers) SetPagination(v *ListUpstreamAppServersResponseBodyAppServersPagination) *ListUpstreamAppServersResponseBodyAppServers {
	s.Pagination = v
	return s
}

type ListUpstreamAppServersResponseBodyAppServersList struct {
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PAppKey       *string `json:"PAppKey,omitempty" xml:"PAppKey,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueueNameList *string `json:"QueueNameList,omitempty" xml:"QueueNameList,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListUpstreamAppServersResponseBodyAppServersList) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBodyAppServersList) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetGmtCreate(v int64) *ListUpstreamAppServersResponseBodyAppServersList {
	s.GmtCreate = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetGmtModified(v int64) *ListUpstreamAppServersResponseBodyAppServersList {
	s.GmtModified = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetId(v int64) *ListUpstreamAppServersResponseBodyAppServersList {
	s.Id = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetName(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.Name = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetPAppKey(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.PAppKey = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetProjectId(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetQueueNameList(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.QueueNameList = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetTags(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.Tags = &v
	return s
}

type ListUpstreamAppServersResponseBodyAppServersPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListUpstreamAppServersResponseBodyAppServersPagination) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBodyAppServersPagination) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBodyAppServersPagination) SetPageIndex(v int32) *ListUpstreamAppServersResponseBodyAppServersPagination {
	s.PageIndex = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersPagination) SetPageSize(v int32) *ListUpstreamAppServersResponseBodyAppServersPagination {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersPagination) SetTotalCount(v int32) *ListUpstreamAppServersResponseBodyAppServersPagination {
	s.TotalCount = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersPagination) SetTotalPageCount(v int32) *ListUpstreamAppServersResponseBodyAppServersPagination {
	s.TotalPageCount = &v
	return s
}

type ListUpstreamAppServersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUpstreamAppServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUpstreamAppServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponse) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponse) SetHeaders(v map[string]*string) *ListUpstreamAppServersResponse {
	s.Headers = v
	return s
}

func (s *ListUpstreamAppServersResponse) SetStatusCode(v int32) *ListUpstreamAppServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpstreamAppServersResponse) SetBody(v *ListUpstreamAppServersResponseBody) *ListUpstreamAppServersResponse {
	s.Body = v
	return s
}

type ListVersionDeviceGroupsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListVersionDeviceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVersionDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionDeviceGroupsRequest) SetProjectId(v string) *ListVersionDeviceGroupsRequest {
	s.ProjectId = &v
	return s
}

type ListVersionDeviceGroupsResponseBody struct {
	DeviceGroupList []*ListVersionDeviceGroupsResponseBodyDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Repeated"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVersionDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVersionDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionDeviceGroupsResponseBody) SetDeviceGroupList(v []*ListVersionDeviceGroupsResponseBodyDeviceGroupList) *ListVersionDeviceGroupsResponseBody {
	s.DeviceGroupList = v
	return s
}

func (s *ListVersionDeviceGroupsResponseBody) SetRequestId(v string) *ListVersionDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListVersionDeviceGroupsResponseBodyDeviceGroupList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify   *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListVersionDeviceGroupsResponseBodyDeviceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListVersionDeviceGroupsResponseBodyDeviceGroupList) GoString() string {
	return s.String()
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetDescription(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Description = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetGmtCreate(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.GmtCreate = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetGmtModify(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.GmtModify = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetId(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Id = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetName(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Name = &v
	return s
}

type ListVersionDeviceGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVersionDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVersionDeviceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVersionDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListVersionDeviceGroupsResponse) SetHeaders(v map[string]*string) *ListVersionDeviceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListVersionDeviceGroupsResponse) SetStatusCode(v int32) *ListVersionDeviceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVersionDeviceGroupsResponse) SetBody(v *ListVersionDeviceGroupsResponseBody) *ListVersionDeviceGroupsResponse {
	s.Body = v
	return s
}

type PublishAppVersionRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SendMessage *bool   `json:"SendMessage,omitempty" xml:"SendMessage,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s PublishAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishAppVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishAppVersionRequest) SetProjectId(v string) *PublishAppVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *PublishAppVersionRequest) SetSendMessage(v bool) *PublishAppVersionRequest {
	s.SendMessage = &v
	return s
}

func (s *PublishAppVersionRequest) SetVersionId(v string) *PublishAppVersionRequest {
	s.VersionId = &v
	return s
}

type PublishAppVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAppVersionResponseBody) SetRequestId(v string) *PublishAppVersionResponseBody {
	s.RequestId = &v
	return s
}

type PublishAppVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishAppVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishAppVersionResponse) SetHeaders(v map[string]*string) *PublishAppVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishAppVersionResponse) SetStatusCode(v int32) *PublishAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAppVersionResponse) SetBody(v *PublishAppVersionResponseBody) *PublishAppVersionResponse {
	s.Body = v
	return s
}

type PublishOsVersionRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SendMessage *bool   `json:"SendMessage,omitempty" xml:"SendMessage,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s PublishOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishOsVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishOsVersionRequest) SetProjectId(v string) *PublishOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *PublishOsVersionRequest) SetSendMessage(v bool) *PublishOsVersionRequest {
	s.SendMessage = &v
	return s
}

func (s *PublishOsVersionRequest) SetVersionId(v string) *PublishOsVersionRequest {
	s.VersionId = &v
	return s
}

type PublishOsVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishOsVersionResponseBody) SetRequestId(v string) *PublishOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type PublishOsVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishOsVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishOsVersionResponse) SetHeaders(v map[string]*string) *PublishOsVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishOsVersionResponse) SetStatusCode(v int32) *PublishOsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishOsVersionResponse) SetBody(v *PublishOsVersionResponseBody) *PublishOsVersionResponse {
	s.Body = v
	return s
}

type PushMessageRequest struct {
	Act            *string `json:"Act,omitempty" xml:"Act,omitempty"`
	AppKey         *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppPackage     *string `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	CustomContent  *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	Desc           *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	ExpiredTime    *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PkgContent     *string `json:"PkgContent,omitempty" xml:"PkgContent,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReceiverType   *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ReceiverValues *string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri            *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s PushMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMessageRequest) GoString() string {
	return s.String()
}

func (s *PushMessageRequest) SetAct(v string) *PushMessageRequest {
	s.Act = &v
	return s
}

func (s *PushMessageRequest) SetAppKey(v string) *PushMessageRequest {
	s.AppKey = &v
	return s
}

func (s *PushMessageRequest) SetAppPackage(v string) *PushMessageRequest {
	s.AppPackage = &v
	return s
}

func (s *PushMessageRequest) SetCustomContent(v string) *PushMessageRequest {
	s.CustomContent = &v
	return s
}

func (s *PushMessageRequest) SetDesc(v string) *PushMessageRequest {
	s.Desc = &v
	return s
}

func (s *PushMessageRequest) SetExpiredTime(v int64) *PushMessageRequest {
	s.ExpiredTime = &v
	return s
}

func (s *PushMessageRequest) SetPkgContent(v string) *PushMessageRequest {
	s.PkgContent = &v
	return s
}

func (s *PushMessageRequest) SetProjectId(v string) *PushMessageRequest {
	s.ProjectId = &v
	return s
}

func (s *PushMessageRequest) SetReceiverType(v string) *PushMessageRequest {
	s.ReceiverType = &v
	return s
}

func (s *PushMessageRequest) SetReceiverValues(v string) *PushMessageRequest {
	s.ReceiverValues = &v
	return s
}

func (s *PushMessageRequest) SetTitle(v string) *PushMessageRequest {
	s.Title = &v
	return s
}

func (s *PushMessageRequest) SetType(v int32) *PushMessageRequest {
	s.Type = &v
	return s
}

func (s *PushMessageRequest) SetUri(v string) *PushMessageRequest {
	s.Uri = &v
	return s
}

type PushMessageResponseBody struct {
	Mid       *int64  `json:"Mid,omitempty" xml:"Mid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushMessageResponseBody) SetMid(v int64) *PushMessageResponseBody {
	s.Mid = &v
	return s
}

func (s *PushMessageResponseBody) SetRequestId(v string) *PushMessageResponseBody {
	s.RequestId = &v
	return s
}

type PushMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMessageResponse) GoString() string {
	return s.String()
}

func (s *PushMessageResponse) SetHeaders(v map[string]*string) *PushMessageResponse {
	s.Headers = v
	return s
}

func (s *PushMessageResponse) SetStatusCode(v int32) *PushMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMessageResponse) SetBody(v *PushMessageResponseBody) *PushMessageResponse {
	s.Body = v
	return s
}

type PushVersionMessageRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s PushVersionMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushVersionMessageRequest) GoString() string {
	return s.String()
}

func (s *PushVersionMessageRequest) SetProjectId(v string) *PushVersionMessageRequest {
	s.ProjectId = &v
	return s
}

func (s *PushVersionMessageRequest) SetVersionId(v string) *PushVersionMessageRequest {
	s.VersionId = &v
	return s
}

func (s *PushVersionMessageRequest) SetVersionType(v string) *PushVersionMessageRequest {
	s.VersionType = &v
	return s
}

type PushVersionMessageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushVersionMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushVersionMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushVersionMessageResponseBody) SetRequestId(v string) *PushVersionMessageResponseBody {
	s.RequestId = &v
	return s
}

type PushVersionMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushVersionMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushVersionMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushVersionMessageResponse) GoString() string {
	return s.String()
}

func (s *PushVersionMessageResponse) SetHeaders(v map[string]*string) *PushVersionMessageResponse {
	s.Headers = v
	return s
}

func (s *PushVersionMessageResponse) SetStatusCode(v int32) *PushVersionMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PushVersionMessageResponse) SetBody(v *PushVersionMessageResponseBody) *PushVersionMessageResponse {
	s.Body = v
	return s
}

type QueryPrepublishPassedDeviceCountRequest struct {
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryPrepublishPassedDeviceCountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepublishPassedDeviceCountRequest) GoString() string {
	return s.String()
}

func (s *QueryPrepublishPassedDeviceCountRequest) SetPrepublishId(v string) *QueryPrepublishPassedDeviceCountRequest {
	s.PrepublishId = &v
	return s
}

func (s *QueryPrepublishPassedDeviceCountRequest) SetProjectId(v string) *QueryPrepublishPassedDeviceCountRequest {
	s.ProjectId = &v
	return s
}

type QueryPrepublishPassedDeviceCountResponseBody struct {
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPrepublishPassedDeviceCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepublishPassedDeviceCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPrepublishPassedDeviceCountResponseBody) SetCount(v int32) *QueryPrepublishPassedDeviceCountResponseBody {
	s.Count = &v
	return s
}

func (s *QueryPrepublishPassedDeviceCountResponseBody) SetRequestId(v string) *QueryPrepublishPassedDeviceCountResponseBody {
	s.RequestId = &v
	return s
}

type QueryPrepublishPassedDeviceCountResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPrepublishPassedDeviceCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPrepublishPassedDeviceCountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepublishPassedDeviceCountResponse) GoString() string {
	return s.String()
}

func (s *QueryPrepublishPassedDeviceCountResponse) SetHeaders(v map[string]*string) *QueryPrepublishPassedDeviceCountResponse {
	s.Headers = v
	return s
}

func (s *QueryPrepublishPassedDeviceCountResponse) SetStatusCode(v int32) *QueryPrepublishPassedDeviceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPrepublishPassedDeviceCountResponse) SetBody(v *QueryPrepublishPassedDeviceCountResponseBody) *QueryPrepublishPassedDeviceCountResponse {
	s.Body = v
	return s
}

type SubmitAssistReportRequest struct {
	AssistDescription *string `json:"AssistDescription,omitempty" xml:"AssistDescription,omitempty"`
	AssistId          *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	AssistReason      *string `json:"AssistReason,omitempty" xml:"AssistReason,omitempty"`
	AssistResult      *string `json:"AssistResult,omitempty" xml:"AssistResult,omitempty"`
	AssistTag         *string `json:"AssistTag,omitempty" xml:"AssistTag,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SubmitAssistReportRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAssistReportRequest) GoString() string {
	return s.String()
}

func (s *SubmitAssistReportRequest) SetAssistDescription(v string) *SubmitAssistReportRequest {
	s.AssistDescription = &v
	return s
}

func (s *SubmitAssistReportRequest) SetAssistId(v string) *SubmitAssistReportRequest {
	s.AssistId = &v
	return s
}

func (s *SubmitAssistReportRequest) SetAssistReason(v string) *SubmitAssistReportRequest {
	s.AssistReason = &v
	return s
}

func (s *SubmitAssistReportRequest) SetAssistResult(v string) *SubmitAssistReportRequest {
	s.AssistResult = &v
	return s
}

func (s *SubmitAssistReportRequest) SetAssistTag(v string) *SubmitAssistReportRequest {
	s.AssistTag = &v
	return s
}

func (s *SubmitAssistReportRequest) SetDeviceModel(v string) *SubmitAssistReportRequest {
	s.DeviceModel = &v
	return s
}

func (s *SubmitAssistReportRequest) SetProjectId(v string) *SubmitAssistReportRequest {
	s.ProjectId = &v
	return s
}

type SubmitAssistReportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAssistReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAssistReportResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAssistReportResponseBody) SetRequestId(v string) *SubmitAssistReportResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAssistReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitAssistReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAssistReportResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAssistReportResponse) GoString() string {
	return s.String()
}

func (s *SubmitAssistReportResponse) SetHeaders(v map[string]*string) *SubmitAssistReportResponse {
	s.Headers = v
	return s
}

func (s *SubmitAssistReportResponse) SetStatusCode(v int32) *SubmitAssistReportResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAssistReportResponse) SetBody(v *SubmitAssistReportResponseBody) *SubmitAssistReportResponse {
	s.Body = v
	return s
}

type UpdateApiGatewayAppStatusRequest struct {
	GatewayAppId *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateApiGatewayAppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiGatewayAppStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiGatewayAppStatusRequest) SetGatewayAppId(v string) *UpdateApiGatewayAppStatusRequest {
	s.GatewayAppId = &v
	return s
}

func (s *UpdateApiGatewayAppStatusRequest) SetProjectId(v string) *UpdateApiGatewayAppStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateApiGatewayAppStatusRequest) SetStatus(v int32) *UpdateApiGatewayAppStatusRequest {
	s.Status = &v
	return s
}

type UpdateApiGatewayAppStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApiGatewayAppStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiGatewayAppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiGatewayAppStatusResponseBody) SetRequestId(v string) *UpdateApiGatewayAppStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApiGatewayAppStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApiGatewayAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApiGatewayAppStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiGatewayAppStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiGatewayAppStatusResponse) SetHeaders(v map[string]*string) *UpdateApiGatewayAppStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiGatewayAppStatusResponse) SetStatusCode(v int32) *UpdateApiGatewayAppStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiGatewayAppStatusResponse) SetBody(v *UpdateApiGatewayAppStatusResponseBody) *UpdateApiGatewayAppStatusResponse {
	s.Body = v
	return s
}

type UpdateAppBlackWhiteVersionsRequest struct {
	BlackAppVersions *string `json:"BlackAppVersions,omitempty" xml:"BlackAppVersions,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	WhiteAppVersions *string `json:"WhiteAppVersions,omitempty" xml:"WhiteAppVersions,omitempty"`
}

func (s UpdateAppBlackWhiteVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppBlackWhiteVersionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppBlackWhiteVersionsRequest) SetBlackAppVersions(v string) *UpdateAppBlackWhiteVersionsRequest {
	s.BlackAppVersions = &v
	return s
}

func (s *UpdateAppBlackWhiteVersionsRequest) SetProjectId(v string) *UpdateAppBlackWhiteVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAppBlackWhiteVersionsRequest) SetVersionId(v string) *UpdateAppBlackWhiteVersionsRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateAppBlackWhiteVersionsRequest) SetWhiteAppVersions(v string) *UpdateAppBlackWhiteVersionsRequest {
	s.WhiteAppVersions = &v
	return s
}

type UpdateAppBlackWhiteVersionsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppBlackWhiteVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppBlackWhiteVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppBlackWhiteVersionsResponseBody) SetRequestId(v string) *UpdateAppBlackWhiteVersionsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppBlackWhiteVersionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppBlackWhiteVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppBlackWhiteVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppBlackWhiteVersionsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppBlackWhiteVersionsResponse) SetHeaders(v map[string]*string) *UpdateAppBlackWhiteVersionsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppBlackWhiteVersionsResponse) SetStatusCode(v int32) *UpdateAppBlackWhiteVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppBlackWhiteVersionsResponse) SetBody(v *UpdateAppBlackWhiteVersionsResponseBody) *UpdateAppBlackWhiteVersionsResponse {
	s.Body = v
	return s
}

type UpdateAppVersionRequest struct {
	ApkMd5            *string `json:"ApkMd5,omitempty" xml:"ApkMd5,omitempty"`
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion        *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BlackVersionList  *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	DeviceAdapterList *string `json:"DeviceAdapterList,omitempty" xml:"DeviceAdapterList,omitempty"`
	InstallType       *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	IsAllowNewInstall *string `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	IsForceUpgrade    *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsNeedRestart     *string `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	IsSilentUpgrade   *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	PackageUrl        *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReleaseNote       *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RestartAppParam   *string `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	RestartAppType    *string `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	RestartType       *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionId         *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	WhiteVersionList  *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
}

func (s UpdateAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionRequest) SetApkMd5(v string) *UpdateAppVersionRequest {
	s.ApkMd5 = &v
	return s
}

func (s *UpdateAppVersionRequest) SetAppId(v string) *UpdateAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppVersionRequest) SetAppVersion(v string) *UpdateAppVersionRequest {
	s.AppVersion = &v
	return s
}

func (s *UpdateAppVersionRequest) SetBlackVersionList(v string) *UpdateAppVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *UpdateAppVersionRequest) SetDeviceAdapterList(v string) *UpdateAppVersionRequest {
	s.DeviceAdapterList = &v
	return s
}

func (s *UpdateAppVersionRequest) SetInstallType(v string) *UpdateAppVersionRequest {
	s.InstallType = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsAllowNewInstall(v string) *UpdateAppVersionRequest {
	s.IsAllowNewInstall = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsForceUpgrade(v string) *UpdateAppVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsNeedRestart(v string) *UpdateAppVersionRequest {
	s.IsNeedRestart = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsSilentUpgrade(v string) *UpdateAppVersionRequest {
	s.IsSilentUpgrade = &v
	return s
}

func (s *UpdateAppVersionRequest) SetPackageUrl(v string) *UpdateAppVersionRequest {
	s.PackageUrl = &v
	return s
}

func (s *UpdateAppVersionRequest) SetProjectId(v string) *UpdateAppVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAppVersionRequest) SetReleaseNote(v string) *UpdateAppVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRemark(v string) *UpdateAppVersionRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRestartAppParam(v string) *UpdateAppVersionRequest {
	s.RestartAppParam = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRestartAppType(v string) *UpdateAppVersionRequest {
	s.RestartAppType = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRestartType(v string) *UpdateAppVersionRequest {
	s.RestartType = &v
	return s
}

func (s *UpdateAppVersionRequest) SetVersionCode(v string) *UpdateAppVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *UpdateAppVersionRequest) SetVersionId(v string) *UpdateAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateAppVersionRequest) SetWhiteVersionList(v string) *UpdateAppVersionRequest {
	s.WhiteVersionList = &v
	return s
}

type UpdateAppVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionResponseBody) SetRequestId(v string) *UpdateAppVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionResponse) SetHeaders(v map[string]*string) *UpdateAppVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppVersionResponse) SetStatusCode(v int32) *UpdateAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppVersionResponse) SetBody(v *UpdateAppVersionResponseBody) *UpdateAppVersionResponse {
	s.Body = v
	return s
}

type UpdateAppVersionReleaseNoteRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateAppVersionReleaseNoteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionReleaseNoteRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionReleaseNoteRequest) SetProjectId(v string) *UpdateAppVersionReleaseNoteRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAppVersionReleaseNoteRequest) SetReleaseNote(v string) *UpdateAppVersionReleaseNoteRequest {
	s.ReleaseNote = &v
	return s
}

func (s *UpdateAppVersionReleaseNoteRequest) SetVersionId(v string) *UpdateAppVersionReleaseNoteRequest {
	s.VersionId = &v
	return s
}

type UpdateAppVersionReleaseNoteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppVersionReleaseNoteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionReleaseNoteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionReleaseNoteResponseBody) SetRequestId(v string) *UpdateAppVersionReleaseNoteResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppVersionReleaseNoteResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppVersionReleaseNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppVersionReleaseNoteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionReleaseNoteResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionReleaseNoteResponse) SetHeaders(v map[string]*string) *UpdateAppVersionReleaseNoteResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppVersionReleaseNoteResponse) SetStatusCode(v int32) *UpdateAppVersionReleaseNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppVersionReleaseNoteResponse) SetBody(v *UpdateAppVersionReleaseNoteResponseBody) *UpdateAppVersionReleaseNoteResponse {
	s.Body = v
	return s
}

type UpdateAppVersionRemarkRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateAppVersionRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionRemarkRequest) SetProjectId(v string) *UpdateAppVersionRemarkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAppVersionRemarkRequest) SetRemark(v string) *UpdateAppVersionRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAppVersionRemarkRequest) SetVersionId(v string) *UpdateAppVersionRemarkRequest {
	s.VersionId = &v
	return s
}

type UpdateAppVersionRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppVersionRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionRemarkResponseBody) SetRequestId(v string) *UpdateAppVersionRemarkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppVersionRemarkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppVersionRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppVersionRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionRemarkResponse) SetHeaders(v map[string]*string) *UpdateAppVersionRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppVersionRemarkResponse) SetStatusCode(v int32) *UpdateAppVersionRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppVersionRemarkResponse) SetBody(v *UpdateAppVersionRemarkResponseBody) *UpdateAppVersionRemarkResponse {
	s.Body = v
	return s
}

type UpdateAppVersionStatusRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAppVersionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionStatusRequest) SetId(v string) *UpdateAppVersionStatusRequest {
	s.Id = &v
	return s
}

func (s *UpdateAppVersionStatusRequest) SetProjectId(v string) *UpdateAppVersionStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAppVersionStatusRequest) SetStatus(v string) *UpdateAppVersionStatusRequest {
	s.Status = &v
	return s
}

type UpdateAppVersionStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppVersionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionStatusResponseBody) SetRequestId(v string) *UpdateAppVersionStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppVersionStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppVersionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionStatusResponse) SetHeaders(v map[string]*string) *UpdateAppVersionStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppVersionStatusResponse) SetStatusCode(v int32) *UpdateAppVersionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppVersionStatusResponse) SetBody(v *UpdateAppVersionStatusResponseBody) *UpdateAppVersionStatusResponse {
	s.Body = v
	return s
}

type UpdateCustomizedFilterRequest struct {
	BlackWhiteType   *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueCompareType *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	ValueType        *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomizedFilterRequest) SetBlackWhiteType(v string) *UpdateCustomizedFilterRequest {
	s.BlackWhiteType = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetId(v int64) *UpdateCustomizedFilterRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetName(v string) *UpdateCustomizedFilterRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetProjectId(v string) *UpdateCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetValue(v string) *UpdateCustomizedFilterRequest {
	s.Value = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetValueCompareType(v string) *UpdateCustomizedFilterRequest {
	s.ValueCompareType = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetValueType(v string) *UpdateCustomizedFilterRequest {
	s.ValueType = &v
	return s
}

type UpdateCustomizedFilterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomizedFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomizedFilterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomizedFilterResponseBody) SetRequestId(v string) *UpdateCustomizedFilterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomizedFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomizedFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomizedFilterResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomizedFilterResponse) SetHeaders(v map[string]*string) *UpdateCustomizedFilterResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomizedFilterResponse) SetStatusCode(v int32) *UpdateCustomizedFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomizedFilterResponse) SetBody(v *UpdateCustomizedFilterResponseBody) *UpdateCustomizedFilterResponse {
	s.Body = v
	return s
}

type UpdateDeviceModelRequest struct {
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	CanCreateDeviceId *string `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InitUsageType     *string `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	ModelName         *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
}

func (s UpdateDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceModelRequest) SetBrandName(v string) *UpdateDeviceModelRequest {
	s.BrandName = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetCanCreateDeviceId(v string) *UpdateDeviceModelRequest {
	s.CanCreateDeviceId = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetDescription(v string) *UpdateDeviceModelRequest {
	s.Description = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetDeviceName(v string) *UpdateDeviceModelRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetDeviceType(v string) *UpdateDeviceModelRequest {
	s.DeviceType = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetHardwareType(v string) *UpdateDeviceModelRequest {
	s.HardwareType = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetId(v string) *UpdateDeviceModelRequest {
	s.Id = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetInitUsageType(v string) *UpdateDeviceModelRequest {
	s.InitUsageType = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetModelName(v string) *UpdateDeviceModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetObjectKey(v string) *UpdateDeviceModelRequest {
	s.ObjectKey = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetOsPlatform(v string) *UpdateDeviceModelRequest {
	s.OsPlatform = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetProjectId(v string) *UpdateDeviceModelRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetSecurityChip(v string) *UpdateDeviceModelRequest {
	s.SecurityChip = &v
	return s
}

type UpdateDeviceModelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceModelResponseBody) SetRequestId(v string) *UpdateDeviceModelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeviceModelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceModelResponse) SetHeaders(v map[string]*string) *UpdateDeviceModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceModelResponse) SetStatusCode(v int32) *UpdateDeviceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceModelResponse) SetBody(v *UpdateDeviceModelResponseBody) *UpdateDeviceModelResponse {
	s.Body = v
	return s
}

type UpdateNamespaceDataRequest struct {
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountType  *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AuthType     *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NewData      *string `json:"NewData,omitempty" xml:"NewData,omitempty"`
	OldData      *string `json:"OldData,omitempty" xml:"OldData,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateNamespaceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceDataRequest) SetAccountId(v string) *UpdateNamespaceDataRequest {
	s.AccountId = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetAccountType(v string) *UpdateNamespaceDataRequest {
	s.AccountType = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetAuthType(v string) *UpdateNamespaceDataRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetDeviceId(v string) *UpdateNamespaceDataRequest {
	s.DeviceId = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetDeviceIdType(v string) *UpdateNamespaceDataRequest {
	s.DeviceIdType = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetNamespace(v string) *UpdateNamespaceDataRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetNewData(v string) *UpdateNamespaceDataRequest {
	s.NewData = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetOldData(v string) *UpdateNamespaceDataRequest {
	s.OldData = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetPath(v string) *UpdateNamespaceDataRequest {
	s.Path = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetProjectId(v string) *UpdateNamespaceDataRequest {
	s.ProjectId = &v
	return s
}

type UpdateNamespaceDataResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNamespaceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceDataResponseBody) SetRequestId(v string) *UpdateNamespaceDataResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNamespaceDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNamespaceDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNamespaceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceDataResponse) SetHeaders(v map[string]*string) *UpdateNamespaceDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceDataResponse) SetStatusCode(v int32) *UpdateNamespaceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceDataResponse) SetBody(v *UpdateNamespaceDataResponseBody) *UpdateNamespaceDataResponse {
	s.Body = v
	return s
}

type UpdateOsBlackWhiteVersionsRequest struct {
	BlackVersions *string `json:"BlackVersions,omitempty" xml:"BlackVersions,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	WhiteVersions *string `json:"WhiteVersions,omitempty" xml:"WhiteVersions,omitempty"`
}

func (s UpdateOsBlackWhiteVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsBlackWhiteVersionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsBlackWhiteVersionsRequest) SetBlackVersions(v string) *UpdateOsBlackWhiteVersionsRequest {
	s.BlackVersions = &v
	return s
}

func (s *UpdateOsBlackWhiteVersionsRequest) SetProjectId(v string) *UpdateOsBlackWhiteVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsBlackWhiteVersionsRequest) SetVersionId(v string) *UpdateOsBlackWhiteVersionsRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateOsBlackWhiteVersionsRequest) SetWhiteVersions(v string) *UpdateOsBlackWhiteVersionsRequest {
	s.WhiteVersions = &v
	return s
}

type UpdateOsBlackWhiteVersionsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOsBlackWhiteVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsBlackWhiteVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOsBlackWhiteVersionsResponseBody) SetRequestId(v string) *UpdateOsBlackWhiteVersionsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOsBlackWhiteVersionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOsBlackWhiteVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOsBlackWhiteVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsBlackWhiteVersionsResponse) GoString() string {
	return s.String()
}

func (s *UpdateOsBlackWhiteVersionsResponse) SetHeaders(v map[string]*string) *UpdateOsBlackWhiteVersionsResponse {
	s.Headers = v
	return s
}

func (s *UpdateOsBlackWhiteVersionsResponse) SetStatusCode(v int32) *UpdateOsBlackWhiteVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOsBlackWhiteVersionsResponse) SetBody(v *UpdateOsBlackWhiteVersionsResponseBody) *UpdateOsBlackWhiteVersionsResponse {
	s.Body = v
	return s
}

type UpdateOsVersionRequest struct {
	BlackVersionList            *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	DeviceModelId               *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	EnableMobileDownload        *string `json:"EnableMobileDownload,omitempty" xml:"EnableMobileDownload,omitempty"`
	Id                          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsForceNightUpgrade         *string `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	IsForceUpgrade              *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsMilestone                 *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	MaxClientVersion            *string `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	MinClientVersion            *string `json:"MinClientVersion,omitempty" xml:"MinClientVersion,omitempty"`
	MobileDownloadMaxSize       *string `json:"MobileDownloadMaxSize,omitempty" xml:"MobileDownloadMaxSize,omitempty"`
	NightUpgradeDownloadType    *string `json:"NightUpgradeDownloadType,omitempty" xml:"NightUpgradeDownloadType,omitempty"`
	NightUpgradeIsAllowedCancel *string `json:"NightUpgradeIsAllowedCancel,omitempty" xml:"NightUpgradeIsAllowedCancel,omitempty"`
	NightUpgradeIsShowTip       *string `json:"NightUpgradeIsShowTip,omitempty" xml:"NightUpgradeIsShowTip,omitempty"`
	ProjectId                   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReleaseNote                 *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark                      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RomList                     *string `json:"RomList,omitempty" xml:"RomList,omitempty"`
	SystemVersion               *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	WhiteVersionList            *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
}

func (s UpdateOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionRequest) SetBlackVersionList(v string) *UpdateOsVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *UpdateOsVersionRequest) SetDeviceModelId(v string) *UpdateOsVersionRequest {
	s.DeviceModelId = &v
	return s
}

func (s *UpdateOsVersionRequest) SetEnableMobileDownload(v string) *UpdateOsVersionRequest {
	s.EnableMobileDownload = &v
	return s
}

func (s *UpdateOsVersionRequest) SetId(v string) *UpdateOsVersionRequest {
	s.Id = &v
	return s
}

func (s *UpdateOsVersionRequest) SetIsForceNightUpgrade(v string) *UpdateOsVersionRequest {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *UpdateOsVersionRequest) SetIsForceUpgrade(v string) *UpdateOsVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *UpdateOsVersionRequest) SetIsMilestone(v string) *UpdateOsVersionRequest {
	s.IsMilestone = &v
	return s
}

func (s *UpdateOsVersionRequest) SetMaxClientVersion(v string) *UpdateOsVersionRequest {
	s.MaxClientVersion = &v
	return s
}

func (s *UpdateOsVersionRequest) SetMinClientVersion(v string) *UpdateOsVersionRequest {
	s.MinClientVersion = &v
	return s
}

func (s *UpdateOsVersionRequest) SetMobileDownloadMaxSize(v string) *UpdateOsVersionRequest {
	s.MobileDownloadMaxSize = &v
	return s
}

func (s *UpdateOsVersionRequest) SetNightUpgradeDownloadType(v string) *UpdateOsVersionRequest {
	s.NightUpgradeDownloadType = &v
	return s
}

func (s *UpdateOsVersionRequest) SetNightUpgradeIsAllowedCancel(v string) *UpdateOsVersionRequest {
	s.NightUpgradeIsAllowedCancel = &v
	return s
}

func (s *UpdateOsVersionRequest) SetNightUpgradeIsShowTip(v string) *UpdateOsVersionRequest {
	s.NightUpgradeIsShowTip = &v
	return s
}

func (s *UpdateOsVersionRequest) SetProjectId(v string) *UpdateOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsVersionRequest) SetReleaseNote(v string) *UpdateOsVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *UpdateOsVersionRequest) SetRemark(v string) *UpdateOsVersionRequest {
	s.Remark = &v
	return s
}

func (s *UpdateOsVersionRequest) SetRomList(v string) *UpdateOsVersionRequest {
	s.RomList = &v
	return s
}

func (s *UpdateOsVersionRequest) SetSystemVersion(v string) *UpdateOsVersionRequest {
	s.SystemVersion = &v
	return s
}

func (s *UpdateOsVersionRequest) SetWhiteVersionList(v string) *UpdateOsVersionRequest {
	s.WhiteVersionList = &v
	return s
}

type UpdateOsVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionResponseBody) SetRequestId(v string) *UpdateOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOsVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionResponse) SetHeaders(v map[string]*string) *UpdateOsVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOsVersionResponse) SetStatusCode(v int32) *UpdateOsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOsVersionResponse) SetBody(v *UpdateOsVersionResponseBody) *UpdateOsVersionResponse {
	s.Body = v
	return s
}

type UpdateOsVersionReleaseNoteRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateOsVersionReleaseNoteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionReleaseNoteRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionReleaseNoteRequest) SetProjectId(v string) *UpdateOsVersionReleaseNoteRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsVersionReleaseNoteRequest) SetReleaseNote(v string) *UpdateOsVersionReleaseNoteRequest {
	s.ReleaseNote = &v
	return s
}

func (s *UpdateOsVersionReleaseNoteRequest) SetVersionId(v string) *UpdateOsVersionReleaseNoteRequest {
	s.VersionId = &v
	return s
}

type UpdateOsVersionReleaseNoteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOsVersionReleaseNoteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionReleaseNoteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionReleaseNoteResponseBody) SetRequestId(v string) *UpdateOsVersionReleaseNoteResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOsVersionReleaseNoteResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOsVersionReleaseNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOsVersionReleaseNoteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionReleaseNoteResponse) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionReleaseNoteResponse) SetHeaders(v map[string]*string) *UpdateOsVersionReleaseNoteResponse {
	s.Headers = v
	return s
}

func (s *UpdateOsVersionReleaseNoteResponse) SetStatusCode(v int32) *UpdateOsVersionReleaseNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOsVersionReleaseNoteResponse) SetBody(v *UpdateOsVersionReleaseNoteResponseBody) *UpdateOsVersionReleaseNoteResponse {
	s.Body = v
	return s
}

type UpdateOsVersionRemarkRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateOsVersionRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionRemarkRequest) SetProjectId(v string) *UpdateOsVersionRemarkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsVersionRemarkRequest) SetRemark(v string) *UpdateOsVersionRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateOsVersionRemarkRequest) SetVersionId(v string) *UpdateOsVersionRemarkRequest {
	s.VersionId = &v
	return s
}

type UpdateOsVersionRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOsVersionRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionRemarkResponseBody) SetRequestId(v string) *UpdateOsVersionRemarkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOsVersionRemarkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOsVersionRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOsVersionRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionRemarkResponse) SetHeaders(v map[string]*string) *UpdateOsVersionRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateOsVersionRemarkResponse) SetStatusCode(v int32) *UpdateOsVersionRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOsVersionRemarkResponse) SetBody(v *UpdateOsVersionRemarkResponseBody) *UpdateOsVersionRemarkResponse {
	s.Body = v
	return s
}

type UpdateOsVersionStatusRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateOsVersionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionStatusRequest) SetId(v string) *UpdateOsVersionStatusRequest {
	s.Id = &v
	return s
}

func (s *UpdateOsVersionStatusRequest) SetProjectId(v string) *UpdateOsVersionStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsVersionStatusRequest) SetStatus(v string) *UpdateOsVersionStatusRequest {
	s.Status = &v
	return s
}

type UpdateOsVersionStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOsVersionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionStatusResponseBody) SetRequestId(v string) *UpdateOsVersionStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOsVersionStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOsVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOsVersionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionStatusResponse) SetHeaders(v map[string]*string) *UpdateOsVersionStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateOsVersionStatusResponse) SetStatusCode(v int32) *UpdateOsVersionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOsVersionStatusResponse) SetBody(v *UpdateOsVersionStatusResponseBody) *UpdateOsVersionStatusResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateSchemaSubscribeRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubscribeList *string `json:"SubscribeList,omitempty" xml:"SubscribeList,omitempty"`
}

func (s UpdateSchemaSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemaSubscribeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSchemaSubscribeRequest) SetDeviceModel(v string) *UpdateSchemaSubscribeRequest {
	s.DeviceModel = &v
	return s
}

func (s *UpdateSchemaSubscribeRequest) SetProjectId(v string) *UpdateSchemaSubscribeRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateSchemaSubscribeRequest) SetSchemaVersion(v string) *UpdateSchemaSubscribeRequest {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateSchemaSubscribeRequest) SetSubscribeList(v string) *UpdateSchemaSubscribeRequest {
	s.SubscribeList = &v
	return s
}

type UpdateSchemaSubscribeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSchemaSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemaSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSchemaSubscribeResponseBody) SetRequestId(v string) *UpdateSchemaSubscribeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSchemaSubscribeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSchemaSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSchemaSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSchemaSubscribeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSchemaSubscribeResponse) SetHeaders(v map[string]*string) *UpdateSchemaSubscribeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSchemaSubscribeResponse) SetStatusCode(v int32) *UpdateSchemaSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSchemaSubscribeResponse) SetBody(v *UpdateSchemaSubscribeResponseBody) *UpdateSchemaSubscribeResponse {
	s.Body = v
	return s
}

type UpdateShadowSchemaRequest struct {
	AuthType      *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Schema        *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s UpdateShadowSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateShadowSchemaRequest) GoString() string {
	return s.String()
}

func (s *UpdateShadowSchemaRequest) SetAuthType(v string) *UpdateShadowSchemaRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateShadowSchemaRequest) SetDeviceModelId(v string) *UpdateShadowSchemaRequest {
	s.DeviceModelId = &v
	return s
}

func (s *UpdateShadowSchemaRequest) SetId(v string) *UpdateShadowSchemaRequest {
	s.Id = &v
	return s
}

func (s *UpdateShadowSchemaRequest) SetNamespace(v string) *UpdateShadowSchemaRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateShadowSchemaRequest) SetProjectId(v string) *UpdateShadowSchemaRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateShadowSchemaRequest) SetSchema(v string) *UpdateShadowSchemaRequest {
	s.Schema = &v
	return s
}

type UpdateShadowSchemaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateShadowSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateShadowSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateShadowSchemaResponseBody) SetRequestId(v string) *UpdateShadowSchemaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateShadowSchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateShadowSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateShadowSchemaResponse) GoString() string {
	return s.String()
}

func (s *UpdateShadowSchemaResponse) SetHeaders(v map[string]*string) *UpdateShadowSchemaResponse {
	s.Headers = v
	return s
}

func (s *UpdateShadowSchemaResponse) SetStatusCode(v int32) *UpdateShadowSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateShadowSchemaResponse) SetBody(v *UpdateShadowSchemaResponseBody) *UpdateShadowSchemaResponse {
	s.Body = v
	return s
}

type UpdateTriggerRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Production *int32  `json:"Production,omitempty" xml:"Production,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Sandbox    *int32  `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) SetId(v int64) *UpdateTriggerRequest {
	s.Id = &v
	return s
}

func (s *UpdateTriggerRequest) SetProduction(v int32) *UpdateTriggerRequest {
	s.Production = &v
	return s
}

func (s *UpdateTriggerRequest) SetProjectId(v string) *UpdateTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTriggerRequest) SetSandbox(v int32) *UpdateTriggerRequest {
	s.Sandbox = &v
	return s
}

type UpdateTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponseBody) SetRequestId(v string) *UpdateTriggerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponse) SetHeaders(v map[string]*string) *UpdateTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTriggerResponse) SetStatusCode(v int32) *UpdateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTriggerResponse) SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse {
	s.Body = v
	return s
}

type UpdateUpstreamAppServerRequest struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateUpstreamAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUpstreamAppServerRequest) GoString() string {
	return s.String()
}

func (s *UpdateUpstreamAppServerRequest) SetId(v int64) *UpdateUpstreamAppServerRequest {
	s.Id = &v
	return s
}

func (s *UpdateUpstreamAppServerRequest) SetName(v string) *UpdateUpstreamAppServerRequest {
	s.Name = &v
	return s
}

func (s *UpdateUpstreamAppServerRequest) SetProjectId(v string) *UpdateUpstreamAppServerRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateUpstreamAppServerRequest) SetTags(v string) *UpdateUpstreamAppServerRequest {
	s.Tags = &v
	return s
}

type UpdateUpstreamAppServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUpstreamAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUpstreamAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUpstreamAppServerResponseBody) SetRequestId(v string) *UpdateUpstreamAppServerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUpstreamAppServerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUpstreamAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUpstreamAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUpstreamAppServerResponse) GoString() string {
	return s.String()
}

func (s *UpdateUpstreamAppServerResponse) SetHeaders(v map[string]*string) *UpdateUpstreamAppServerResponse {
	s.Headers = v
	return s
}

func (s *UpdateUpstreamAppServerResponse) SetStatusCode(v int32) *UpdateUpstreamAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUpstreamAppServerResponse) SetBody(v *UpdateUpstreamAppServerResponseBody) *UpdateUpstreamAppServerResponse {
	s.Body = v
	return s
}

type UpdateVersionDeviceGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateVersionDeviceGroupRequest) SetDescription(v string) *UpdateVersionDeviceGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateVersionDeviceGroupRequest) SetId(v string) *UpdateVersionDeviceGroupRequest {
	s.Id = &v
	return s
}

func (s *UpdateVersionDeviceGroupRequest) SetName(v string) *UpdateVersionDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateVersionDeviceGroupRequest) SetProjectId(v string) *UpdateVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

type UpdateVersionDeviceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVersionDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVersionDeviceGroupResponseBody) SetRequestId(v string) *UpdateVersionDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVersionDeviceGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVersionDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateVersionDeviceGroupResponse) SetHeaders(v map[string]*string) *UpdateVersionDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateVersionDeviceGroupResponse) SetStatusCode(v int32) *UpdateVersionDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVersionDeviceGroupResponse) SetBody(v *UpdateVersionDeviceGroupResponseBody) *UpdateVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type UpdateVersionPrepublishActiveStatusRequest struct {
	IsActive     *string `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateVersionPrepublishActiveStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionPrepublishActiveStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateVersionPrepublishActiveStatusRequest) SetIsActive(v string) *UpdateVersionPrepublishActiveStatusRequest {
	s.IsActive = &v
	return s
}

func (s *UpdateVersionPrepublishActiveStatusRequest) SetPrepublishId(v string) *UpdateVersionPrepublishActiveStatusRequest {
	s.PrepublishId = &v
	return s
}

func (s *UpdateVersionPrepublishActiveStatusRequest) SetProjectId(v string) *UpdateVersionPrepublishActiveStatusRequest {
	s.ProjectId = &v
	return s
}

type UpdateVersionPrepublishActiveStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVersionPrepublishActiveStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionPrepublishActiveStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVersionPrepublishActiveStatusResponseBody) SetRequestId(v string) *UpdateVersionPrepublishActiveStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVersionPrepublishActiveStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVersionPrepublishActiveStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVersionPrepublishActiveStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionPrepublishActiveStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateVersionPrepublishActiveStatusResponse) SetHeaders(v map[string]*string) *UpdateVersionPrepublishActiveStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateVersionPrepublishActiveStatusResponse) SetStatusCode(v int32) *UpdateVersionPrepublishActiveStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVersionPrepublishActiveStatusResponse) SetBody(v *UpdateVersionPrepublishActiveStatusResponseBody) *UpdateVersionPrepublishActiveStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iovcc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddUploadedFunctionFileInfoWithOptions(request *AddUploadedFunctionFileInfoRequest, runtime *util.RuntimeOptions) (_result *AddUploadedFunctionFileInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKey)) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUploadedFunctionFileInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUploadedFunctionFileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUploadedFunctionFileInfo(request *AddUploadedFunctionFileInfoRequest) (_result *AddUploadedFunctionFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUploadedFunctionFileInfoResponse{}
	_body, _err := client.AddUploadedFunctionFileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVersionBlackDevicesWithOptions(request *AddVersionBlackDevicesRequest, runtime *util.RuntimeOptions) (_result *AddVersionBlackDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVersionBlackDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVersionBlackDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVersionBlackDevices(request *AddVersionBlackDevicesRequest) (_result *AddVersionBlackDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVersionBlackDevicesResponse{}
	_body, _err := client.AddVersionBlackDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVersionGroupDevicesWithOptions(request *AddVersionGroupDevicesRequest, runtime *util.RuntimeOptions) (_result *AddVersionGroupDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVersionGroupDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVersionGroupDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVersionGroupDevices(request *AddVersionGroupDevicesRequest) (_result *AddVersionGroupDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVersionGroupDevicesResponse{}
	_body, _err := client.AddVersionGroupDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVersionWhiteDevicesWithOptions(request *AddVersionWhiteDevicesRequest, runtime *util.RuntimeOptions) (_result *AddVersionWhiteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVersionWhiteDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVersionWhiteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVersionWhiteDevices(request *AddVersionWhiteDevicesRequest) (_result *AddVersionWhiteDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVersionWhiteDevicesResponse{}
	_body, _err := client.AddVersionWhiteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVersionWhiteDevicesByDeviceGroupsWithOptions(request *AddVersionWhiteDevicesByDeviceGroupsRequest, runtime *util.RuntimeOptions) (_result *AddVersionWhiteDevicesByDeviceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVersionWhiteDevicesByDeviceGroups"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVersionWhiteDevicesByDeviceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVersionWhiteDevicesByDeviceGroups(request *AddVersionWhiteDevicesByDeviceGroupsRequest) (_result *AddVersionWhiteDevicesByDeviceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVersionWhiteDevicesByDeviceGroupsResponse{}
	_body, _err := client.AddVersionWhiteDevicesByDeviceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConnectAssistDeviceWithOptions(request *ConnectAssistDeviceRequest, runtime *util.RuntimeOptions) (_result *ConnectAssistDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowCommandExtension)) {
		body["AllowCommandExtension"] = request.AllowCommandExtension
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareId)) {
		body["HardwareId"] = request.HardwareId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UUID)) {
		body["UUID"] = request.UUID
	}

	if !tea.BoolValue(util.IsUnset(request.VIN)) {
		body["VIN"] = request.VIN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConnectAssistDevice"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConnectAssistDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConnectAssistDevice(request *ConnectAssistDeviceRequest) (_result *ConnectAssistDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConnectAssistDeviceResponse{}
	_body, _err := client.ConnectAssistDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountActivatedOrNewRegistrationDeviceWithOptions(request *CountActivatedOrNewRegistrationDeviceRequest, runtime *util.RuntimeOptions) (_result *CountActivatedOrNewRegistrationDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceBrand)) {
		query["DeviceBrand"] = request.DeviceBrand
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceBrandId)) {
		query["DeviceBrandId"] = request.DeviceBrandId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCountStatType)) {
		query["DeviceCountStatType"] = request.DeviceCountStatType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		query["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsQueryNewRegistrationDevice)) {
		query["IsQueryNewRegistrationDevice"] = request.IsQueryNewRegistrationDevice
	}

	if !tea.BoolValue(util.IsUnset(request.IsQueryYearlyActivate)) {
		query["IsQueryYearlyActivate"] = request.IsQueryYearlyActivate
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountActivatedOrNewRegistrationDevice"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountActivatedOrNewRegistrationDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountActivatedOrNewRegistrationDevice(request *CountActivatedOrNewRegistrationDeviceRequest) (_result *CountActivatedOrNewRegistrationDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountActivatedOrNewRegistrationDeviceResponse{}
	_body, _err := client.CountActivatedOrNewRegistrationDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountDeviceBrandsWithOptions(request *CountDeviceBrandsRequest, runtime *util.RuntimeOptions) (_result *CountDeviceBrandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountDeviceBrands"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountDeviceBrandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountDeviceBrands(request *CountDeviceBrandsRequest) (_result *CountDeviceBrandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountDeviceBrandsResponse{}
	_body, _err := client.CountDeviceBrandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountDeviceModelsWithOptions(request *CountDeviceModelsRequest, runtime *util.RuntimeOptions) (_result *CountDeviceModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountDeviceModels"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountDeviceModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountDeviceModels(request *CountDeviceModelsRequest) (_result *CountDeviceModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountDeviceModelsResponse{}
	_body, _err := client.CountDeviceModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountDevicesWithOptions(request *CountDevicesRequest, runtime *util.RuntimeOptions) (_result *CountDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountDevices(request *CountDevicesRequest) (_result *CountDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountDevicesResponse{}
	_body, _err := client.CountDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountYunIdInfoWithOptions(runtime *util.RuntimeOptions) (_result *CountYunIdInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CountYunIdInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountYunIdInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountYunIdInfo() (_result *CountYunIdInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountYunIdInfoResponse{}
	_body, _err := client.CountYunIdInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppVersionWithOptions(request *CreateAppVersionRequest, runtime *util.RuntimeOptions) (_result *CreateAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApkMd5)) {
		query["ApkMd5"] = request.ApkMd5
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BlackVersionList)) {
		query["BlackVersionList"] = request.BlackVersionList
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceAdapterList)) {
		query["DeviceAdapterList"] = request.DeviceAdapterList
	}

	if !tea.BoolValue(util.IsUnset(request.InstallType)) {
		query["InstallType"] = request.InstallType
	}

	if !tea.BoolValue(util.IsUnset(request.IsAllowNewInstall)) {
		query["IsAllowNewInstall"] = request.IsAllowNewInstall
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceUpgrade)) {
		query["IsForceUpgrade"] = request.IsForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.IsNeedRestart)) {
		query["IsNeedRestart"] = request.IsNeedRestart
	}

	if !tea.BoolValue(util.IsUnset(request.IsSilentUpgrade)) {
		query["IsSilentUpgrade"] = request.IsSilentUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.PackageUrl)) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RestartAppParam)) {
		query["RestartAppParam"] = request.RestartAppParam
	}

	if !tea.BoolValue(util.IsUnset(request.RestartAppType)) {
		query["RestartAppType"] = request.RestartAppType
	}

	if !tea.BoolValue(util.IsUnset(request.RestartType)) {
		query["RestartType"] = request.RestartType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionCode)) {
		query["VersionCode"] = request.VersionCode
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteVersionList)) {
		query["WhiteVersionList"] = request.WhiteVersionList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppVersion(request *CreateAppVersionRequest) (_result *CreateAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppVersionResponse{}
	_body, _err := client.CreateAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomizedFilterWithOptions(request *CreateCustomizedFilterRequest, runtime *util.RuntimeOptions) (_result *CreateCustomizedFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackWhiteType)) {
		query["BlackWhiteType"] = request.BlackWhiteType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.ValueCompareType)) {
		query["ValueCompareType"] = request.ValueCompareType
	}

	if !tea.BoolValue(util.IsUnset(request.ValueType)) {
		query["ValueType"] = request.ValueType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomizedFilter"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomizedFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomizedFilter(request *CreateCustomizedFilterRequest) (_result *CreateCustomizedFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomizedFilterResponse{}
	_body, _err := client.CreateCustomizedFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomizedPropertyWithOptions(request *CreateCustomizedPropertyRequest, runtime *util.RuntimeOptions) (_result *CreateCustomizedPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomizedProperty"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomizedPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomizedProperty(request *CreateCustomizedPropertyRequest) (_result *CreateCustomizedPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomizedPropertyResponse{}
	_body, _err := client.CreateCustomizedPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceBrandWithOptions(request *CreateDeviceBrandRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceBrandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandName)) {
		query["BrandName"] = request.BrandName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Manufacture)) {
		query["Manufacture"] = request.Manufacture
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeviceBrand"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceBrandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceBrand(request *CreateDeviceBrandRequest) (_result *CreateDeviceBrandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceBrandResponse{}
	_body, _err := client.CreateDeviceBrandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceModelWithOptions(request *CreateDeviceModelRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandName)) {
		query["BrandName"] = request.BrandName
	}

	if !tea.BoolValue(util.IsUnset(request.CanCreateDeviceId)) {
		query["CanCreateDeviceId"] = request.CanCreateDeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareType)) {
		query["HardwareType"] = request.HardwareType
	}

	if !tea.BoolValue(util.IsUnset(request.InitUsageType)) {
		query["InitUsageType"] = request.InitUsageType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKey)) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.OsPlatform)) {
		query["OsPlatform"] = request.OsPlatform
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityChip)) {
		query["SecurityChip"] = request.SecurityChip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeviceModel"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceModel(request *CreateDeviceModelRequest) (_result *CreateDeviceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceModelResponse{}
	_body, _err := client.CreateDeviceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		query["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOsVersionWithOptions(request *CreateOsVersionRequest, runtime *util.RuntimeOptions) (_result *CreateOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackVersionList)) {
		query["BlackVersionList"] = request.BlackVersionList
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMobileDownload)) {
		query["EnableMobileDownload"] = request.EnableMobileDownload
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceNightUpgrade)) {
		query["IsForceNightUpgrade"] = request.IsForceNightUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceUpgrade)) {
		query["IsForceUpgrade"] = request.IsForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.IsMilestone)) {
		query["IsMilestone"] = request.IsMilestone
	}

	if !tea.BoolValue(util.IsUnset(request.MaxClientVersion)) {
		query["MaxClientVersion"] = request.MaxClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MinClientVersion)) {
		query["MinClientVersion"] = request.MinClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MobileDownloadMaxSize)) {
		query["MobileDownloadMaxSize"] = request.MobileDownloadMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.NightUpgradeDownloadType)) {
		query["NightUpgradeDownloadType"] = request.NightUpgradeDownloadType
	}

	if !tea.BoolValue(util.IsUnset(request.NightUpgradeIsAllowedCancel)) {
		query["NightUpgradeIsAllowedCancel"] = request.NightUpgradeIsAllowedCancel
	}

	if !tea.BoolValue(util.IsUnset(request.NightUpgradeIsShowTip)) {
		query["NightUpgradeIsShowTip"] = request.NightUpgradeIsShowTip
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RomList)) {
		query["RomList"] = request.RomList
	}

	if !tea.BoolValue(util.IsUnset(request.SystemVersion)) {
		query["SystemVersion"] = request.SystemVersion
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteVersionList)) {
		query["WhiteVersionList"] = request.WhiteVersionList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOsVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOsVersion(request *CreateOsVersionRequest) (_result *CreateOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOsVersionResponse{}
	_body, _err := client.CreateOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectDesc)) {
		query["ProjectDesc"] = request.ProjectDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectAppWithOptions(request *CreateProjectAppRequest, runtime *util.RuntimeOptions) (_result *CreateProjectAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppPkgName)) {
		query["AppPkgName"] = request.AppPkgName
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProjectApp"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProjectApp(request *CreateProjectAppRequest) (_result *CreateProjectAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectAppResponse{}
	_body, _err := client.CreateProjectAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRpcServiceWithOptions(request *CreateRpcServiceRequest, runtime *util.RuntimeOptions) (_result *CreateRpcServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InterfaceName)) {
		query["InterfaceName"] = request.InterfaceName
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeType)) {
		query["InvokeType"] = request.InvokeType
	}

	if !tea.BoolValue(util.IsUnset(request.MethodName)) {
		query["MethodName"] = request.MethodName
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionCode)) {
		query["VersionCode"] = request.VersionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRpcService"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRpcServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRpcService(request *CreateRpcServiceRequest) (_result *CreateRpcServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRpcServiceResponse{}
	_body, _err := client.CreateRpcServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSchemaSubscribeWithOptions(request *CreateSchemaSubscribeRequest, runtime *util.RuntimeOptions) (_result *CreateSchemaSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		query["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeList)) {
		query["SubscribeList"] = request.SubscribeList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchemaSubscribe"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSchemaSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchemaSubscribe(request *CreateSchemaSubscribeRequest) (_result *CreateSchemaSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchemaSubscribeResponse{}
	_body, _err := client.CreateSchemaSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateShadowSchemaWithOptions(request *CreateShadowSchemaRequest, runtime *util.RuntimeOptions) (_result *CreateShadowSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		query["Schema"] = request.Schema
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateShadowSchema"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateShadowSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateShadowSchema(request *CreateShadowSchemaRequest) (_result *CreateShadowSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateShadowSchemaResponse{}
	_body, _err := client.CreateShadowSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTriggerWithOptions(request *CreateTriggerRequest, runtime *util.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileIds)) {
		query["FileIds"] = request.FileIds
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionIds)) {
		query["FunctionIds"] = request.FunctionIds
	}

	if !tea.BoolValue(util.IsUnset(request.InvocationMode)) {
		query["InvocationMode"] = request.InvocationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Production)) {
		query["Production"] = request.Production
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Sandbox)) {
		query["Sandbox"] = request.Sandbox
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrigger"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrigger(request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUpstreamAppKeyRelationWithOptions(request *CreateUpstreamAppKeyRelationRequest, runtime *util.RuntimeOptions) (_result *CreateUpstreamAppKeyRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.PAppKey)) {
		query["PAppKey"] = request.PAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUpstreamAppKeyRelation"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUpstreamAppKeyRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUpstreamAppKeyRelation(request *CreateUpstreamAppKeyRelationRequest) (_result *CreateUpstreamAppKeyRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUpstreamAppKeyRelationResponse{}
	_body, _err := client.CreateUpstreamAppKeyRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUpstreamAppKeyRelationsWithOptions(request *CreateUpstreamAppKeyRelationsRequest, runtime *util.RuntimeOptions) (_result *CreateUpstreamAppKeyRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKeys)) {
		query["AppKeys"] = request.AppKeys
	}

	if !tea.BoolValue(util.IsUnset(request.AppServerId)) {
		query["AppServerId"] = request.AppServerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUpstreamAppKeyRelations"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUpstreamAppKeyRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUpstreamAppKeyRelations(request *CreateUpstreamAppKeyRelationsRequest) (_result *CreateUpstreamAppKeyRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUpstreamAppKeyRelationsResponse{}
	_body, _err := client.CreateUpstreamAppKeyRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUpstreamAppServerWithOptions(request *CreateUpstreamAppServerRequest, runtime *util.RuntimeOptions) (_result *CreateUpstreamAppServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUpstreamAppServer"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUpstreamAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUpstreamAppServer(request *CreateUpstreamAppServerRequest) (_result *CreateUpstreamAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUpstreamAppServerResponse{}
	_body, _err := client.CreateUpstreamAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVersionDeviceGroupWithOptions(request *CreateVersionDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateVersionDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCount)) {
		query["MaxCount"] = request.MaxCount
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVersionDeviceGroup"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVersionDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVersionDeviceGroup(request *CreateVersionDeviceGroupRequest) (_result *CreateVersionDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVersionDeviceGroupResponse{}
	_body, _err := client.CreateVersionDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVersionPrepublishWithOptions(request *CreateVersionPrepublishRequest, runtime *util.RuntimeOptions) (_result *CreateVersionPrepublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BarrierCount)) {
		query["BarrierCount"] = request.BarrierCount
	}

	if !tea.BoolValue(util.IsUnset(request.IsTotalPrepublish)) {
		query["IsTotalPrepublish"] = request.IsTotalPrepublish
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVersionPrepublish"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVersionPrepublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVersionPrepublish(request *CreateVersionPrepublishRequest) (_result *CreateVersionPrepublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVersionPrepublishResponse{}
	_body, _err := client.CreateVersionPrepublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVersionTestWithOptions(request *CreateVersionTestRequest, runtime *util.RuntimeOptions) (_result *CreateVersionTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVersionTest"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVersionTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVersionTest(request *CreateVersionTestRequest) (_result *CreateVersionTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVersionTestResponse{}
	_body, _err := client.CreateVersionTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DelayPublishOsVersionWithOptions(request *DelayPublishOsVersionRequest, runtime *util.RuntimeOptions) (_result *DelayPublishOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DownTime)) {
		query["DownTime"] = request.DownTime
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PrepubTime)) {
		query["PrepubTime"] = request.PrepubTime
	}

	if !tea.BoolValue(util.IsUnset(request.PrepublishCount)) {
		query["PrepublishCount"] = request.PrepublishCount
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.PublishTime)) {
		query["PublishTime"] = request.PublishTime
	}

	if !tea.BoolValue(util.IsUnset(request.SendMessage)) {
		query["SendMessage"] = request.SendMessage
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DelayPublishOsVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelayPublishOsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DelayPublishOsVersion(request *DelayPublishOsVersionRequest) (_result *DelayPublishOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelayPublishOsVersionResponse{}
	_body, _err := client.DelayPublishOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAllCustomizedFiltersWithOptions(request *DeleteAllCustomizedFiltersRequest, runtime *util.RuntimeOptions) (_result *DeleteAllCustomizedFiltersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAllCustomizedFilters"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAllCustomizedFiltersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAllCustomizedFilters(request *DeleteAllCustomizedFiltersRequest) (_result *DeleteAllCustomizedFiltersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAllCustomizedFiltersResponse{}
	_body, _err := client.DeleteAllCustomizedFiltersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAllVersionGroupDevicesWithOptions(request *DeleteAllVersionGroupDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteAllVersionGroupDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAllVersionGroupDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAllVersionGroupDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAllVersionGroupDevices(request *DeleteAllVersionGroupDevicesRequest) (_result *DeleteAllVersionGroupDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAllVersionGroupDevicesResponse{}
	_body, _err := client.DeleteAllVersionGroupDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomizedFilterWithOptions(request *DeleteCustomizedFilterRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomizedFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomizedFilter"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomizedFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomizedFilter(request *DeleteCustomizedFilterRequest) (_result *DeleteCustomizedFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomizedFilterResponse{}
	_body, _err := client.DeleteCustomizedFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomizedPropertyWithOptions(request *DeleteCustomizedPropertyRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomizedPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomizedProperty"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomizedPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomizedProperty(request *DeleteCustomizedPropertyRequest) (_result *DeleteCustomizedPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomizedPropertyResponse{}
	_body, _err := client.DeleteCustomizedPropertyWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevice"),
		Version:     tea.String("2018-05-01"),
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

func (client *Client) DeleteFunctionFileWithOptions(request *DeleteFunctionFileRequest, runtime *util.RuntimeOptions) (_result *DeleteFunctionFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionFile"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunctionFile(request *DeleteFunctionFileRequest) (_result *DeleteFunctionFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFunctionFileResponse{}
	_body, _err := client.DeleteFunctionFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOpenAccountWithOptions(request *DeleteOpenAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteOpenAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOpenAccount"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOpenAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOpenAccount(request *DeleteOpenAccountRequest) (_result *DeleteOpenAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOpenAccountResponse{}
	_body, _err := client.DeleteOpenAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectAppWithOptions(request *DeleteProjectAppRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProjectApp"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProjectApp(request *DeleteProjectAppRequest) (_result *DeleteProjectAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectAppResponse{}
	_body, _err := client.DeleteProjectAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRpcServiceWithOptions(request *DeleteRpcServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteRpcServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRpcService"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRpcServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRpcService(request *DeleteRpcServiceRequest) (_result *DeleteRpcServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRpcServiceResponse{}
	_body, _err := client.DeleteRpcServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSchemaSubscribeWithOptions(request *DeleteSchemaSubscribeRequest, runtime *util.RuntimeOptions) (_result *DeleteSchemaSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchemaSubscribe"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSchemaSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSchemaSubscribe(request *DeleteSchemaSubscribeRequest) (_result *DeleteSchemaSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSchemaSubscribeResponse{}
	_body, _err := client.DeleteSchemaSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteShadowSchemaWithOptions(request *DeleteShadowSchemaRequest, runtime *util.RuntimeOptions) (_result *DeleteShadowSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteShadowSchema"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteShadowSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteShadowSchema(request *DeleteShadowSchemaRequest) (_result *DeleteShadowSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteShadowSchemaResponse{}
	_body, _err := client.DeleteShadowSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTriggerWithOptions(request *DeleteTriggerRequest, runtime *util.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrigger"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrigger(request *DeleteTriggerRequest) (_result *DeleteTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DeleteTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUpstreamAppKeyRelationWithOptions(request *DeleteUpstreamAppKeyRelationRequest, runtime *util.RuntimeOptions) (_result *DeleteUpstreamAppKeyRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUpstreamAppKeyRelation"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUpstreamAppKeyRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUpstreamAppKeyRelation(request *DeleteUpstreamAppKeyRelationRequest) (_result *DeleteUpstreamAppKeyRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUpstreamAppKeyRelationResponse{}
	_body, _err := client.DeleteUpstreamAppKeyRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUpstreamAppServerWithOptions(request *DeleteUpstreamAppServerRequest, runtime *util.RuntimeOptions) (_result *DeleteUpstreamAppServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUpstreamAppServer"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUpstreamAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUpstreamAppServer(request *DeleteUpstreamAppServerRequest) (_result *DeleteUpstreamAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUpstreamAppServerResponse{}
	_body, _err := client.DeleteUpstreamAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionAllBlackDevicesWithOptions(request *DeleteVersionAllBlackDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionAllBlackDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionAllBlackDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionAllBlackDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionAllBlackDevices(request *DeleteVersionAllBlackDevicesRequest) (_result *DeleteVersionAllBlackDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionAllBlackDevicesResponse{}
	_body, _err := client.DeleteVersionAllBlackDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionAllWhiteDevicesWithOptions(request *DeleteVersionAllWhiteDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionAllWhiteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionAllWhiteDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionAllWhiteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionAllWhiteDevices(request *DeleteVersionAllWhiteDevicesRequest) (_result *DeleteVersionAllWhiteDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionAllWhiteDevicesResponse{}
	_body, _err := client.DeleteVersionAllWhiteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionBlackDevicesWithOptions(request *DeleteVersionBlackDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionBlackDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionBlackDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionBlackDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionBlackDevices(request *DeleteVersionBlackDevicesRequest) (_result *DeleteVersionBlackDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionBlackDevicesResponse{}
	_body, _err := client.DeleteVersionBlackDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionBlackDevicesByIdWithOptions(request *DeleteVersionBlackDevicesByIdRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionBlackDevicesByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionBlackDevicesById"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionBlackDevicesByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionBlackDevicesById(request *DeleteVersionBlackDevicesByIdRequest) (_result *DeleteVersionBlackDevicesByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionBlackDevicesByIdResponse{}
	_body, _err := client.DeleteVersionBlackDevicesByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionDeviceGroupWithOptions(request *DeleteVersionDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionDeviceGroup"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionDeviceGroup(request *DeleteVersionDeviceGroupRequest) (_result *DeleteVersionDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionDeviceGroupResponse{}
	_body, _err := client.DeleteVersionDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionGroupDeviceWithOptions(request *DeleteVersionGroupDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionGroupDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionGroupDevice"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionGroupDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionGroupDevice(request *DeleteVersionGroupDeviceRequest) (_result *DeleteVersionGroupDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionGroupDeviceResponse{}
	_body, _err := client.DeleteVersionGroupDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionGroupDeviceByIdWithOptions(request *DeleteVersionGroupDeviceByIdRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionGroupDeviceByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionGroupDeviceById"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionGroupDeviceByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionGroupDeviceById(request *DeleteVersionGroupDeviceByIdRequest) (_result *DeleteVersionGroupDeviceByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionGroupDeviceByIdResponse{}
	_body, _err := client.DeleteVersionGroupDeviceByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionTestWithOptions(request *DeleteVersionTestRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionTest"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionTest(request *DeleteVersionTestRequest) (_result *DeleteVersionTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionTestResponse{}
	_body, _err := client.DeleteVersionTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionWhiteDevicesWithOptions(request *DeleteVersionWhiteDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionWhiteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionWhiteDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionWhiteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionWhiteDevices(request *DeleteVersionWhiteDevicesRequest) (_result *DeleteVersionWhiteDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionWhiteDevicesResponse{}
	_body, _err := client.DeleteVersionWhiteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVersionWhiteDevicesByIdWithOptions(request *DeleteVersionWhiteDevicesByIdRequest, runtime *util.RuntimeOptions) (_result *DeleteVersionWhiteDevicesByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVersionWhiteDevicesById"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVersionWhiteDevicesByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVersionWhiteDevicesById(request *DeleteVersionWhiteDevicesByIdRequest) (_result *DeleteVersionWhiteDevicesByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVersionWhiteDevicesByIdResponse{}
	_body, _err := client.DeleteVersionWhiteDevicesByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployFunctionFileWithOptions(request *DeployFunctionFileRequest, runtime *util.RuntimeOptions) (_result *DeployFunctionFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployEnv)) {
		query["DeployEnv"] = request.DeployEnv
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployFunctionFile"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployFunctionFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployFunctionFile(request *DeployFunctionFileRequest) (_result *DeployFunctionFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployFunctionFileResponse{}
	_body, _err := client.DeployFunctionFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiGatewayAppSecurityWithOptions(request *DescribeApiGatewayAppSecurityRequest, runtime *util.RuntimeOptions) (_result *DescribeApiGatewayAppSecurityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayAppId)) {
		query["GatewayAppId"] = request.GatewayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiGatewayAppSecurity"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiGatewayAppSecurityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiGatewayAppSecurity(request *DescribeApiGatewayAppSecurityRequest) (_result *DescribeApiGatewayAppSecurityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiGatewayAppSecurityResponse{}
	_body, _err := client.DescribeApiGatewayAppSecurityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppVersionWithOptions(request *DescribeAppVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppVersion(request *DescribeAppVersionRequest) (_result *DescribeAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppVersionResponse{}
	_body, _err := client.DescribeAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssistRTMPServerAddressWithOptions(request *DescribeAssistRTMPServerAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeAssistRTMPServerAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssistRTMPServerAddress"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssistRTMPServerAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssistRTMPServerAddress(request *DescribeAssistRTMPServerAddressRequest) (_result *DescribeAssistRTMPServerAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssistRTMPServerAddressResponse{}
	_body, _err := client.DescribeAssistRTMPServerAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssistReportWithOptions(request *DescribeAssistReportRequest, runtime *util.RuntimeOptions) (_result *DescribeAssistReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssistReport"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssistReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssistReport(request *DescribeAssistReportRequest) (_result *DescribeAssistReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssistReportResponse{}
	_body, _err := client.DescribeAssistReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssistWSServerAddressWithOptions(request *DescribeAssistWSServerAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeAssistWSServerAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssistWSServerAddress"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssistWSServerAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssistWSServerAddress(request *DescribeAssistWSServerAddressRequest) (_result *DescribeAssistWSServerAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssistWSServerAddressResponse{}
	_body, _err := client.DescribeAssistWSServerAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomizedFilterWithOptions(request *DescribeCustomizedFilterRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizedFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizedFilter"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizedFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomizedFilter(request *DescribeCustomizedFilterRequest) (_result *DescribeCustomizedFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizedFilterResponse{}
	_body, _err := client.DescribeCustomizedFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefaultSchemaWithOptions(request *DescribeDefaultSchemaRequest, runtime *util.RuntimeOptions) (_result *DescribeDefaultSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefaultSchema"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefaultSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefaultSchema(request *DescribeDefaultSchemaRequest) (_result *DescribeDefaultSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefaultSchemaResponse{}
	_body, _err := client.DescribeDefaultSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceBrandWithOptions(request *DescribeDeviceBrandRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceBrandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceBrand"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceBrandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceBrand(request *DescribeDeviceBrandRequest) (_result *DescribeDeviceBrandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceBrandResponse{}
	_body, _err := client.DescribeDeviceBrandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceIdByOuterInfoWithOptions(request *DescribeDeviceIdByOuterInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceIdByOuterInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceIdByOuterInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceIdByOuterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceIdByOuterInfo(request *DescribeDeviceIdByOuterInfoRequest) (_result *DescribeDeviceIdByOuterInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceIdByOuterInfoResponse{}
	_body, _err := client.DescribeDeviceIdByOuterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceInfoWithOptions(request *DescribeDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (_result *DescribeDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DescribeDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceModelWithOptions(request *DescribeDeviceModelRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceModel"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceModel(request *DescribeDeviceModelRequest) (_result *DescribeDeviceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceModelResponse{}
	_body, _err := client.DescribeDeviceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceOnlineInfoWithOptions(request *DescribeDeviceOnlineInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceOnlineInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceOnlineInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceOnlineInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceOnlineInfo(request *DescribeDeviceOnlineInfoRequest) (_result *DescribeDeviceOnlineInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceOnlineInfoResponse{}
	_body, _err := client.DescribeDeviceOnlineInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceShadowWithOptions(request *DescribeDeviceShadowRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceShadowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		query["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewSubscribed)) {
		query["ViewSubscribed"] = request.ViewSubscribed
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceShadow"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceShadowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceShadow(request *DescribeDeviceShadowRequest) (_result *DescribeDeviceShadowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceShadowResponse{}
	_body, _err := client.DescribeDeviceShadowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceValiditySchemaWithOptions(request *DescribeDeviceValiditySchemaRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceValiditySchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		query["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceValiditySchema"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceValiditySchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceValiditySchema(request *DescribeDeviceValiditySchemaRequest) (_result *DescribeDeviceValiditySchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceValiditySchemaResponse{}
	_body, _err := client.DescribeDeviceValiditySchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMessageWithOptions(request *DescribeMessageRequest, runtime *util.RuntimeOptions) (_result *DescribeMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMessage"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMessage(request *DescribeMessageRequest) (_result *DescribeMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMessageResponse{}
	_body, _err := client.DescribeMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpenAccountWithOptions(request *DescribeOpenAccountRequest, runtime *util.RuntimeOptions) (_result *DescribeOpenAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpenAccount"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpenAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpenAccount(request *DescribeOpenAccountRequest) (_result *DescribeOpenAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpenAccountResponse{}
	_body, _err := client.DescribeOpenAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOsVersionWithOptions(request *DescribeOsVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOsVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOsVersion(request *DescribeOsVersionRequest) (_result *DescribeOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOsVersionResponse{}
	_body, _err := client.DescribeOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectWithOptions(request *DescribeProjectRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProject"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProject(request *DescribeProjectRequest) (_result *DescribeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectResponse{}
	_body, _err := client.DescribeProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectAppSecurityWithOptions(request *DescribeProjectAppSecurityRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectAppSecurityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjectAppSecurity"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectAppSecurityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjectAppSecurity(request *DescribeProjectAppSecurityRequest) (_result *DescribeProjectAppSecurityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectAppSecurityResponse{}
	_body, _err := client.DescribeProjectAppSecurityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeShadowSchemaWithOptions(request *DescribeShadowSchemaRequest, runtime *util.RuntimeOptions) (_result *DescribeShadowSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		query["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.IsSimple)) {
		query["IsSimple"] = request.IsSimple
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeShadowSchema"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeShadowSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeShadowSchema(request *DescribeShadowSchemaRequest) (_result *DescribeShadowSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeShadowSchemaResponse{}
	_body, _err := client.DescribeShadowSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVersionDeviceGroupWithOptions(request *DescribeVersionDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeVersionDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVersionDeviceGroup"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVersionDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVersionDeviceGroup(request *DescribeVersionDeviceGroupRequest) (_result *DescribeVersionDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVersionDeviceGroupResponse{}
	_body, _err := client.DescribeVersionDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DiagnosisVersionWithOptions(request *DiagnosisVersionRequest, runtime *util.RuntimeOptions) (_result *DiagnosisVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagnoseStyle)) {
		query["DiagnoseStyle"] = request.DiagnoseStyle
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IdType)) {
		query["IdType"] = request.IdType
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DiagnosisVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DiagnosisVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DiagnosisVersion(request *DiagnosisVersionRequest) (_result *DiagnosisVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DiagnosisVersionResponse{}
	_body, _err := client.DiagnosisVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindAppVersionsWithOptions(request *FindAppVersionsRequest, runtime *util.RuntimeOptions) (_result *FindAppVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindAppVersions"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindAppVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindAppVersions(request *FindAppVersionsRequest) (_result *FindAppVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindAppVersionsResponse{}
	_body, _err := client.FindAppVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindCustomizedFiltersWithOptions(request *FindCustomizedFiltersRequest, runtime *util.RuntimeOptions) (_result *FindCustomizedFiltersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindCustomizedFilters"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindCustomizedFiltersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindCustomizedFilters(request *FindCustomizedFiltersRequest) (_result *FindCustomizedFiltersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindCustomizedFiltersResponse{}
	_body, _err := client.FindCustomizedFiltersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindCustomizedPropertiesWithOptions(request *FindCustomizedPropertiesRequest, runtime *util.RuntimeOptions) (_result *FindCustomizedPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindCustomizedProperties"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindCustomizedPropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindCustomizedProperties(request *FindCustomizedPropertiesRequest) (_result *FindCustomizedPropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindCustomizedPropertiesResponse{}
	_body, _err := client.FindCustomizedPropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindOsVersionsWithOptions(request *FindOsVersionsRequest, runtime *util.RuntimeOptions) (_result *FindOsVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.IsMilestone)) {
		query["IsMilestone"] = request.IsMilestone
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SystemVersion)) {
		query["SystemVersion"] = request.SystemVersion
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindOsVersions"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindOsVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindOsVersions(request *FindOsVersionsRequest) (_result *FindOsVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindOsVersionsResponse{}
	_body, _err := client.FindOsVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindPrepublishPassedDevicesWithOptions(request *FindPrepublishPassedDevicesRequest, runtime *util.RuntimeOptions) (_result *FindPrepublishPassedDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrepublishId)) {
		query["PrepublishId"] = request.PrepublishId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindPrepublishPassedDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindPrepublishPassedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindPrepublishPassedDevices(request *FindPrepublishPassedDevicesRequest) (_result *FindPrepublishPassedDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindPrepublishPassedDevicesResponse{}
	_body, _err := client.FindPrepublishPassedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindPrepublishesByParentIdWithOptions(request *FindPrepublishesByParentIdRequest, runtime *util.RuntimeOptions) (_result *FindPrepublishesByParentIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindPrepublishesByParentId"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindPrepublishesByParentIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindPrepublishesByParentId(request *FindPrepublishesByParentIdRequest) (_result *FindPrepublishesByParentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindPrepublishesByParentIdResponse{}
	_body, _err := client.FindPrepublishesByParentIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindPrepublishesByVersionIdWithOptions(request *FindPrepublishesByVersionIdRequest, runtime *util.RuntimeOptions) (_result *FindPrepublishesByVersionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindPrepublishesByVersionId"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindPrepublishesByVersionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindPrepublishesByVersionId(request *FindPrepublishesByVersionIdRequest) (_result *FindPrepublishesByVersionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindPrepublishesByVersionIdResponse{}
	_body, _err := client.FindPrepublishesByVersionIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionBlackDevicesWithOptions(request *FindVersionBlackDevicesRequest, runtime *util.RuntimeOptions) (_result *FindVersionBlackDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionBlackDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionBlackDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionBlackDevices(request *FindVersionBlackDevicesRequest) (_result *FindVersionBlackDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionBlackDevicesResponse{}
	_body, _err := client.FindVersionBlackDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionDeviceGroupsWithOptions(request *FindVersionDeviceGroupsRequest, runtime *util.RuntimeOptions) (_result *FindVersionDeviceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionDeviceGroups"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionDeviceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionDeviceGroups(request *FindVersionDeviceGroupsRequest) (_result *FindVersionDeviceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionDeviceGroupsResponse{}
	_body, _err := client.FindVersionDeviceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionGroupDevicesWithOptions(request *FindVersionGroupDevicesRequest, runtime *util.RuntimeOptions) (_result *FindVersionGroupDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionGroupDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionGroupDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionGroupDevices(request *FindVersionGroupDevicesRequest) (_result *FindVersionGroupDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionGroupDevicesResponse{}
	_body, _err := client.FindVersionGroupDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionMessageSendRecordsWithOptions(request *FindVersionMessageSendRecordsRequest, runtime *util.RuntimeOptions) (_result *FindVersionMessageSendRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionMessageSendRecords"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionMessageSendRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionMessageSendRecords(request *FindVersionMessageSendRecordsRequest) (_result *FindVersionMessageSendRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionMessageSendRecordsResponse{}
	_body, _err := client.FindVersionMessageSendRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionMessagesWithOptions(request *FindVersionMessagesRequest, runtime *util.RuntimeOptions) (_result *FindVersionMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SendRecordId)) {
		query["SendRecordId"] = request.SendRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.TestId)) {
		query["TestId"] = request.TestId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionMessages"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionMessages(request *FindVersionMessagesRequest) (_result *FindVersionMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionMessagesResponse{}
	_body, _err := client.FindVersionMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionTestsWithOptions(request *FindVersionTestsRequest, runtime *util.RuntimeOptions) (_result *FindVersionTestsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionTests"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionTestsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionTests(request *FindVersionTestsRequest) (_result *FindVersionTestsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionTestsResponse{}
	_body, _err := client.FindVersionTestsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindVersionWhiteDevicesWithOptions(request *FindVersionWhiteDevicesRequest, runtime *util.RuntimeOptions) (_result *FindVersionWhiteDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindVersionWhiteDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindVersionWhiteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindVersionWhiteDevices(request *FindVersionWhiteDevicesRequest) (_result *FindVersionWhiteDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindVersionWhiteDevicesResponse{}
	_body, _err := client.FindVersionWhiteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateFunctionFileUploadMetaWithOptions(request *GenerateFunctionFileUploadMetaRequest, runtime *util.RuntimeOptions) (_result *GenerateFunctionFileUploadMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateFunctionFileUploadMeta"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateFunctionFileUploadMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateFunctionFileUploadMeta(request *GenerateFunctionFileUploadMetaRequest) (_result *GenerateFunctionFileUploadMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateFunctionFileUploadMetaResponse{}
	_body, _err := client.GenerateFunctionFileUploadMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateOssPostPolicyWithOptions(request *GenerateOssPostPolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateOssPostPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessId)) {
		query["AccessId"] = request.AccessId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		query["Ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateOssPostPolicy"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateOssPostPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateOssPostPolicy(request *GenerateOssPostPolicyRequest) (_result *GenerateOssPostPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateOssPostPolicyResponse{}
	_body, _err := client.GenerateOssPostPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateOssUploadMetaWithOptions(request *GenerateOssUploadMetaRequest, runtime *util.RuntimeOptions) (_result *GenerateOssUploadMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		query["Ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateOssUploadMeta"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateOssUploadMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateOssUploadMeta(request *GenerateOssUploadMetaRequest) (_result *GenerateOssUploadMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateOssUploadMetaResponse{}
	_body, _err := client.GenerateOssUploadMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateSdkDownloadInfoWithOptions(request *GenerateSdkDownloadInfoRequest, runtime *util.RuntimeOptions) (_result *GenerateSdkDownloadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CertFileObjectKey)) {
		query["CertFileObjectKey"] = request.CertFileObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.PkgName)) {
		query["PkgName"] = request.PkgName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Sdks)) {
		query["Sdks"] = request.Sdks
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateSdkDownloadInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateSdkDownloadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateSdkDownloadInfo(request *GenerateSdkDownloadInfoRequest) (_result *GenerateSdkDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateSdkDownloadInfoResponse{}
	_body, _err := client.GenerateSdkDownloadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateSysAppDownloadInfoWithOptions(request *GenerateSysAppDownloadInfoRequest, runtime *util.RuntimeOptions) (_result *GenerateSysAppDownloadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertFileObjectKey)) {
		query["CertFileObjectKey"] = request.CertFileObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.PkgName)) {
		query["PkgName"] = request.PkgName
	}

	if !tea.BoolValue(util.IsUnset(request.Plugins)) {
		query["Plugins"] = request.Plugins
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SignMode)) {
		query["SignMode"] = request.SignMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateSysAppDownloadInfo"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateSysAppDownloadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateSysAppDownloadInfo(request *GenerateSysAppDownloadInfoRequest) (_result *GenerateSysAppDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateSysAppDownloadInfoResponse{}
	_body, _err := client.GenerateSysAppDownloadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceAppUpdateFunnelEventsWithOptions(request *GetDeviceAppUpdateFunnelEventsRequest, runtime *util.RuntimeOptions) (_result *GetDeviceAppUpdateFunnelEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdType)) {
		query["IdType"] = request.IdType
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersionCode)) {
		query["TargetVersionCode"] = request.TargetVersionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceAppUpdateFunnelEvents"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceAppUpdateFunnelEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceAppUpdateFunnelEvents(request *GetDeviceAppUpdateFunnelEventsRequest) (_result *GetDeviceAppUpdateFunnelEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceAppUpdateFunnelEventsResponse{}
	_body, _err := client.GetDeviceAppUpdateFunnelEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceSystemUpdateFunnelEventsWithOptions(request *GetDeviceSystemUpdateFunnelEventsRequest, runtime *util.RuntimeOptions) (_result *GetDeviceSystemUpdateFunnelEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdType)) {
		query["IdType"] = request.IdType
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalId)) {
		query["OriginalId"] = request.OriginalId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		query["TargetVersion"] = request.TargetVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceSystemUpdateFunnelEvents"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceSystemUpdateFunnelEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceSystemUpdateFunnelEvents(request *GetDeviceSystemUpdateFunnelEventsRequest) (_result *GetDeviceSystemUpdateFunnelEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceSystemUpdateFunnelEventsResponse{}
	_body, _err := client.GetDeviceSystemUpdateFunnelEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNamespaceDataWithOptions(request *GetNamespaceDataRequest, runtime *util.RuntimeOptions) (_result *GetNamespaceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNamespaceData"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNamespaceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNamespaceData(request *GetNamespaceDataRequest) (_result *GetNamespaceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNamespaceDataResponse{}
	_body, _err := client.GetNamespaceDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNamespaceStatisticsDataWithOptions(request *GetNamespaceStatisticsDataRequest, runtime *util.RuntimeOptions) (_result *GetNamespaceStatisticsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DimensionType)) {
		query["DimensionType"] = request.DimensionType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNamespaceStatisticsData"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNamespaceStatisticsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNamespaceStatisticsData(request *GetNamespaceStatisticsDataRequest) (_result *GetNamespaceStatisticsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNamespaceStatisticsDataResponse{}
	_body, _err := client.GetNamespaceStatisticsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssUploadMetaWithOptions(request *GetOssUploadMetaRequest, runtime *util.RuntimeOptions) (_result *GetOssUploadMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		query["Ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssUploadMeta"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssUploadMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssUploadMeta(request *GetOssUploadMetaRequest) (_result *GetOssUploadMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssUploadMetaResponse{}
	_body, _err := client.GetOssUploadMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeFunctionWithOptions(request *InvokeFunctionRequest, runtime *util.RuntimeOptions) (_result *InvokeFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeFunction"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeFunction(request *InvokeFunctionRequest) (_result *InvokeFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeFunctionResponse{}
	_body, _err := client.InvokeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApiGatewayAppsWithOptions(request *ListApiGatewayAppsRequest, runtime *util.RuntimeOptions) (_result *ListApiGatewayAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApiGatewayApps"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApiGatewayAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApiGatewayApps(request *ListApiGatewayAppsRequest) (_result *ListApiGatewayAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApiGatewayAppsResponse{}
	_body, _err := client.ListApiGatewayAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAssistActionDetailsWithOptions(request *ListAssistActionDetailsRequest, runtime *util.RuntimeOptions) (_result *ListAssistActionDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAssistActionDetails"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistActionDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAssistActionDetails(request *ListAssistActionDetailsRequest) (_result *ListAssistActionDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAssistActionDetailsResponse{}
	_body, _err := client.ListAssistActionDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAssistDevicesWithOptions(request *ListAssistDevicesRequest, runtime *util.RuntimeOptions) (_result *ListAssistDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAssistDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAssistDevices(request *ListAssistDevicesRequest) (_result *ListAssistDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAssistDevicesResponse{}
	_body, _err := client.ListAssistDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAssistHistoriesWithOptions(request *ListAssistHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListAssistHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAssistHistories"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAssistHistories(request *ListAssistHistoriesRequest) (_result *ListAssistHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAssistHistoriesResponse{}
	_body, _err := client.ListAssistHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAssistHistoryDetailsWithOptions(request *ListAssistHistoryDetailsRequest, runtime *util.RuntimeOptions) (_result *ListAssistHistoryDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAssistHistoryDetails"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistHistoryDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAssistHistoryDetails(request *ListAssistHistoryDetailsRequest) (_result *ListAssistHistoryDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAssistHistoryDetailsResponse{}
	_body, _err := client.ListAssistHistoryDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClientPluginVersionsWithOptions(request *ListClientPluginVersionsRequest, runtime *util.RuntimeOptions) (_result *ListClientPluginVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.PkgName)) {
		query["PkgName"] = request.PkgName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClientPluginVersions"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClientPluginVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClientPluginVersions(request *ListClientPluginVersionsRequest) (_result *ListClientPluginVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClientPluginVersionsResponse{}
	_body, _err := client.ListClientPluginVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClientPluginsWithOptions(request *ListClientPluginsRequest, runtime *util.RuntimeOptions) (_result *ListClientPluginsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClientPlugins"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClientPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClientPlugins(request *ListClientPluginsRequest) (_result *ListClientPluginsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClientPluginsResponse{}
	_body, _err := client.ListClientPluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClientSdksWithOptions(request *ListClientSdksRequest, runtime *util.RuntimeOptions) (_result *ListClientSdksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClientSdks"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClientSdksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClientSdks(request *ListClientSdksRequest) (_result *ListClientSdksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClientSdksResponse{}
	_body, _err := client.ListClientSdksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectLogsWithOptions(request *ListConnectLogsRequest, runtime *util.RuntimeOptions) (_result *ListConnectLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnectLogs"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectLogs(request *ListConnectLogsRequest) (_result *ListConnectLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectLogsResponse{}
	_body, _err := client.ListConnectLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployedFunctionsWithOptions(request *ListDeployedFunctionsRequest, runtime *util.RuntimeOptions) (_result *ListDeployedFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployedFunctions"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeployedFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployedFunctions(request *ListDeployedFunctionsRequest) (_result *ListDeployedFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployedFunctionsResponse{}
	_body, _err := client.ListDeployedFunctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceBrandsWithOptions(request *ListDeviceBrandsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceBrandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceBrands"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceBrandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceBrands(request *ListDeviceBrandsRequest) (_result *ListDeviceBrandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceBrandsResponse{}
	_body, _err := client.ListDeviceBrandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceModelsWithOptions(request *ListDeviceModelsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceModels"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceModels(request *ListDeviceModelsRequest) (_result *ListDeviceModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceModelsResponse{}
	_body, _err := client.ListDeviceModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceTypesWithOptions(request *ListDeviceTypesRequest, runtime *util.RuntimeOptions) (_result *ListDeviceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceTypes"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceTypes(request *ListDeviceTypesRequest) (_result *ListDeviceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceTypesResponse{}
	_body, _err := client.ListDeviceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionExecuteLogWithOptions(request *ListFunctionExecuteLogRequest, runtime *util.RuntimeOptions) (_result *ListFunctionExecuteLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionExecuteLog"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionExecuteLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionExecuteLog(request *ListFunctionExecuteLogRequest) (_result *ListFunctionExecuteLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionExecuteLogResponse{}
	_body, _err := client.ListFunctionExecuteLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionFilesWithOptions(request *ListFunctionFilesRequest, runtime *util.RuntimeOptions) (_result *ListFunctionFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionFiles"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionFiles(request *ListFunctionFilesRequest) (_result *ListFunctionFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionFilesResponse{}
	_body, _err := client.ListFunctionFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionFilesByProjectIdWithOptions(request *ListFunctionFilesByProjectIdRequest, runtime *util.RuntimeOptions) (_result *ListFunctionFilesByProjectIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionFilesByProjectId"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionFilesByProjectIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionFilesByProjectId(request *ListFunctionFilesByProjectIdRequest) (_result *ListFunctionFilesByProjectIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionFilesByProjectIdResponse{}
	_body, _err := client.ListFunctionFilesByProjectIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMessageAcksWithOptions(request *ListMessageAcksRequest, runtime *util.RuntimeOptions) (_result *ListMessageAcksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessageAcks"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessageAcksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMessageAcks(request *ListMessageAcksRequest) (_result *ListMessageAcksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMessageAcksResponse{}
	_body, _err := client.ListMessageAcksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMessageReceiversWithOptions(request *ListMessageReceiversRequest, runtime *util.RuntimeOptions) (_result *ListMessageReceiversResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessageReceivers"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessageReceiversResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMessageReceivers(request *ListMessageReceiversRequest) (_result *ListMessageReceiversResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMessageReceiversResponse{}
	_body, _err := client.ListMessageReceiversWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *util.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespaces"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfflineMessagesWithOptions(request *ListOfflineMessagesRequest, runtime *util.RuntimeOptions) (_result *ListOfflineMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOfflineMessages"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOfflineMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfflineMessages(request *ListOfflineMessagesRequest) (_result *ListOfflineMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfflineMessagesResponse{}
	_body, _err := client.ListOfflineMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenAccountLinksWithOptions(request *ListOpenAccountLinksRequest, runtime *util.RuntimeOptions) (_result *ListOpenAccountLinksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpenAccountLinks"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpenAccountLinksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenAccountLinks(request *ListOpenAccountLinksRequest) (_result *ListOpenAccountLinksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenAccountLinksResponse{}
	_body, _err := client.ListOpenAccountLinksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenAccountsWithOptions(request *ListOpenAccountsRequest, runtime *util.RuntimeOptions) (_result *ListOpenAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpenAccounts"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpenAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenAccounts(request *ListOpenAccountsRequest) (_result *ListOpenAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenAccountsResponse{}
	_body, _err := client.ListOpenAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPreChecksWithOptions(runtime *util.RuntimeOptions) (_result *ListPreChecksResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListPreChecks"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPreChecksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPreChecks() (_result *ListPreChecksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPreChecksResponse{}
	_body, _err := client.ListPreChecksWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectAppsWithOptions(request *ListProjectAppsRequest, runtime *util.RuntimeOptions) (_result *ListProjectAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectApps"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectApps(request *ListProjectAppsRequest) (_result *ListProjectAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectAppsResponse{}
	_body, _err := client.ListProjectAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) ListProjects() (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRpcServicesWithOptions(request *ListRpcServicesRequest, runtime *util.RuntimeOptions) (_result *ListRpcServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRpcServices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRpcServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRpcServices(request *ListRpcServicesRequest) (_result *ListRpcServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRpcServicesResponse{}
	_body, _err := client.ListRpcServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSchemaSubscribesWithOptions(request *ListSchemaSubscribesRequest, runtime *util.RuntimeOptions) (_result *ListSchemaSubscribesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSchemaSubscribes"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSchemaSubscribesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSchemaSubscribes(request *ListSchemaSubscribesRequest) (_result *ListSchemaSubscribesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSchemaSubscribesResponse{}
	_body, _err := client.ListSchemaSubscribesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices() (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListShadowSchemaDeviceModelsWithOptions(request *ListShadowSchemaDeviceModelsRequest, runtime *util.RuntimeOptions) (_result *ListShadowSchemaDeviceModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListShadowSchemaDeviceModels"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListShadowSchemaDeviceModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListShadowSchemaDeviceModels(request *ListShadowSchemaDeviceModelsRequest) (_result *ListShadowSchemaDeviceModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListShadowSchemaDeviceModelsResponse{}
	_body, _err := client.ListShadowSchemaDeviceModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListShadowSchemasWithOptions(request *ListShadowSchemasRequest, runtime *util.RuntimeOptions) (_result *ListShadowSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryValue)) {
		query["QueryValue"] = request.QueryValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListShadowSchemas"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListShadowSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListShadowSchemas(request *ListShadowSchemasRequest) (_result *ListShadowSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListShadowSchemasResponse{}
	_body, _err := client.ListShadowSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSupportFeaturesWithOptions(runtime *util.RuntimeOptions) (_result *ListSupportFeaturesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListSupportFeatures"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSupportFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSupportFeatures() (_result *ListSupportFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSupportFeaturesResponse{}
	_body, _err := client.ListSupportFeaturesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTriggersWithOptions(request *ListTriggersRequest, runtime *util.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTriggers"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTriggers(request *ListTriggersRequest) (_result *ListTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTriggersResponse{}
	_body, _err := client.ListTriggersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUpstreamAppKeyRelationsWithOptions(request *ListUpstreamAppKeyRelationsRequest, runtime *util.RuntimeOptions) (_result *ListUpstreamAppKeyRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppServerId)) {
		query["AppServerId"] = request.AppServerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUpstreamAppKeyRelations"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUpstreamAppKeyRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUpstreamAppKeyRelations(request *ListUpstreamAppKeyRelationsRequest) (_result *ListUpstreamAppKeyRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUpstreamAppKeyRelationsResponse{}
	_body, _err := client.ListUpstreamAppKeyRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUpstreamAppServersWithOptions(request *ListUpstreamAppServersRequest, runtime *util.RuntimeOptions) (_result *ListUpstreamAppServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUpstreamAppServers"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUpstreamAppServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUpstreamAppServers(request *ListUpstreamAppServersRequest) (_result *ListUpstreamAppServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUpstreamAppServersResponse{}
	_body, _err := client.ListUpstreamAppServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVersionDeviceGroupsWithOptions(request *ListVersionDeviceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListVersionDeviceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVersionDeviceGroups"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVersionDeviceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVersionDeviceGroups(request *ListVersionDeviceGroupsRequest) (_result *ListVersionDeviceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVersionDeviceGroupsResponse{}
	_body, _err := client.ListVersionDeviceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishAppVersionWithOptions(request *PublishAppVersionRequest, runtime *util.RuntimeOptions) (_result *PublishAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SendMessage)) {
		query["SendMessage"] = request.SendMessage
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishAppVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishAppVersion(request *PublishAppVersionRequest) (_result *PublishAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishAppVersionResponse{}
	_body, _err := client.PublishAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishOsVersionWithOptions(request *PublishOsVersionRequest, runtime *util.RuntimeOptions) (_result *PublishOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SendMessage)) {
		query["SendMessage"] = request.SendMessage
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishOsVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishOsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishOsVersion(request *PublishOsVersionRequest) (_result *PublishOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishOsVersionResponse{}
	_body, _err := client.PublishOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushMessageWithOptions(request *PushMessageRequest, runtime *util.RuntimeOptions) (_result *PushMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Act)) {
		query["Act"] = request.Act
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppPackage)) {
		query["AppPackage"] = request.AppPackage
	}

	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		query["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		query["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.PkgContent)) {
		query["PkgContent"] = request.PkgContent
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverType)) {
		query["ReceiverType"] = request.ReceiverType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverValues)) {
		query["ReceiverValues"] = request.ReceiverValues
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		query["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMessage"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushMessage(request *PushMessageRequest) (_result *PushMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMessageResponse{}
	_body, _err := client.PushMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushVersionMessageWithOptions(request *PushVersionMessageRequest, runtime *util.RuntimeOptions) (_result *PushVersionMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		query["VersionType"] = request.VersionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushVersionMessage"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushVersionMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushVersionMessage(request *PushVersionMessageRequest) (_result *PushVersionMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushVersionMessageResponse{}
	_body, _err := client.PushVersionMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPrepublishPassedDeviceCountWithOptions(request *QueryPrepublishPassedDeviceCountRequest, runtime *util.RuntimeOptions) (_result *QueryPrepublishPassedDeviceCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrepublishId)) {
		query["PrepublishId"] = request.PrepublishId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPrepublishPassedDeviceCount"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPrepublishPassedDeviceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPrepublishPassedDeviceCount(request *QueryPrepublishPassedDeviceCountRequest) (_result *QueryPrepublishPassedDeviceCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPrepublishPassedDeviceCountResponse{}
	_body, _err := client.QueryPrepublishPassedDeviceCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAssistReportWithOptions(request *SubmitAssistReportRequest, runtime *util.RuntimeOptions) (_result *SubmitAssistReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistDescription)) {
		body["AssistDescription"] = request.AssistDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AssistId)) {
		body["AssistId"] = request.AssistId
	}

	if !tea.BoolValue(util.IsUnset(request.AssistReason)) {
		body["AssistReason"] = request.AssistReason
	}

	if !tea.BoolValue(util.IsUnset(request.AssistResult)) {
		body["AssistResult"] = request.AssistResult
	}

	if !tea.BoolValue(util.IsUnset(request.AssistTag)) {
		body["AssistTag"] = request.AssistTag
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		body["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAssistReport"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAssistReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAssistReport(request *SubmitAssistReportRequest) (_result *SubmitAssistReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAssistReportResponse{}
	_body, _err := client.SubmitAssistReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApiGatewayAppStatusWithOptions(request *UpdateApiGatewayAppStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateApiGatewayAppStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayAppId)) {
		query["GatewayAppId"] = request.GatewayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApiGatewayAppStatus"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApiGatewayAppStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApiGatewayAppStatus(request *UpdateApiGatewayAppStatusRequest) (_result *UpdateApiGatewayAppStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApiGatewayAppStatusResponse{}
	_body, _err := client.UpdateApiGatewayAppStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppBlackWhiteVersionsWithOptions(request *UpdateAppBlackWhiteVersionsRequest, runtime *util.RuntimeOptions) (_result *UpdateAppBlackWhiteVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackAppVersions)) {
		query["BlackAppVersions"] = request.BlackAppVersions
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteAppVersions)) {
		query["WhiteAppVersions"] = request.WhiteAppVersions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppBlackWhiteVersions"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppBlackWhiteVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppBlackWhiteVersions(request *UpdateAppBlackWhiteVersionsRequest) (_result *UpdateAppBlackWhiteVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppBlackWhiteVersionsResponse{}
	_body, _err := client.UpdateAppBlackWhiteVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppVersionWithOptions(request *UpdateAppVersionRequest, runtime *util.RuntimeOptions) (_result *UpdateAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApkMd5)) {
		query["ApkMd5"] = request.ApkMd5
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BlackVersionList)) {
		query["BlackVersionList"] = request.BlackVersionList
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceAdapterList)) {
		query["DeviceAdapterList"] = request.DeviceAdapterList
	}

	if !tea.BoolValue(util.IsUnset(request.InstallType)) {
		query["InstallType"] = request.InstallType
	}

	if !tea.BoolValue(util.IsUnset(request.IsAllowNewInstall)) {
		query["IsAllowNewInstall"] = request.IsAllowNewInstall
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceUpgrade)) {
		query["IsForceUpgrade"] = request.IsForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.IsNeedRestart)) {
		query["IsNeedRestart"] = request.IsNeedRestart
	}

	if !tea.BoolValue(util.IsUnset(request.IsSilentUpgrade)) {
		query["IsSilentUpgrade"] = request.IsSilentUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.PackageUrl)) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RestartAppParam)) {
		query["RestartAppParam"] = request.RestartAppParam
	}

	if !tea.BoolValue(util.IsUnset(request.RestartAppType)) {
		query["RestartAppType"] = request.RestartAppType
	}

	if !tea.BoolValue(util.IsUnset(request.RestartType)) {
		query["RestartType"] = request.RestartType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionCode)) {
		query["VersionCode"] = request.VersionCode
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteVersionList)) {
		query["WhiteVersionList"] = request.WhiteVersionList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppVersion(request *UpdateAppVersionRequest) (_result *UpdateAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppVersionResponse{}
	_body, _err := client.UpdateAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppVersionReleaseNoteWithOptions(request *UpdateAppVersionReleaseNoteRequest, runtime *util.RuntimeOptions) (_result *UpdateAppVersionReleaseNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppVersionReleaseNote"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppVersionReleaseNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppVersionReleaseNote(request *UpdateAppVersionReleaseNoteRequest) (_result *UpdateAppVersionReleaseNoteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppVersionReleaseNoteResponse{}
	_body, _err := client.UpdateAppVersionReleaseNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppVersionRemarkWithOptions(request *UpdateAppVersionRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateAppVersionRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppVersionRemark"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppVersionRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppVersionRemark(request *UpdateAppVersionRemarkRequest) (_result *UpdateAppVersionRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppVersionRemarkResponse{}
	_body, _err := client.UpdateAppVersionRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppVersionStatusWithOptions(request *UpdateAppVersionStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAppVersionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppVersionStatus"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppVersionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppVersionStatus(request *UpdateAppVersionStatusRequest) (_result *UpdateAppVersionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppVersionStatusResponse{}
	_body, _err := client.UpdateAppVersionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomizedFilterWithOptions(request *UpdateCustomizedFilterRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomizedFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackWhiteType)) {
		query["BlackWhiteType"] = request.BlackWhiteType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.ValueCompareType)) {
		query["ValueCompareType"] = request.ValueCompareType
	}

	if !tea.BoolValue(util.IsUnset(request.ValueType)) {
		query["ValueType"] = request.ValueType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomizedFilter"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomizedFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomizedFilter(request *UpdateCustomizedFilterRequest) (_result *UpdateCustomizedFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomizedFilterResponse{}
	_body, _err := client.UpdateCustomizedFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceModelWithOptions(request *UpdateDeviceModelRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandName)) {
		query["BrandName"] = request.BrandName
	}

	if !tea.BoolValue(util.IsUnset(request.CanCreateDeviceId)) {
		query["CanCreateDeviceId"] = request.CanCreateDeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareType)) {
		query["HardwareType"] = request.HardwareType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InitUsageType)) {
		query["InitUsageType"] = request.InitUsageType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKey)) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.OsPlatform)) {
		query["OsPlatform"] = request.OsPlatform
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityChip)) {
		query["SecurityChip"] = request.SecurityChip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeviceModel"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeviceModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceModel(request *UpdateDeviceModelRequest) (_result *UpdateDeviceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceModelResponse{}
	_body, _err := client.UpdateDeviceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNamespaceDataWithOptions(request *UpdateNamespaceDataRequest, runtime *util.RuntimeOptions) (_result *UpdateNamespaceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIdType)) {
		query["DeviceIdType"] = request.DeviceIdType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NewData)) {
		query["NewData"] = request.NewData
	}

	if !tea.BoolValue(util.IsUnset(request.OldData)) {
		query["OldData"] = request.OldData
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNamespaceData"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNamespaceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNamespaceData(request *UpdateNamespaceDataRequest) (_result *UpdateNamespaceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNamespaceDataResponse{}
	_body, _err := client.UpdateNamespaceDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOsBlackWhiteVersionsWithOptions(request *UpdateOsBlackWhiteVersionsRequest, runtime *util.RuntimeOptions) (_result *UpdateOsBlackWhiteVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackVersions)) {
		query["BlackVersions"] = request.BlackVersions
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteVersions)) {
		query["WhiteVersions"] = request.WhiteVersions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOsBlackWhiteVersions"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOsBlackWhiteVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOsBlackWhiteVersions(request *UpdateOsBlackWhiteVersionsRequest) (_result *UpdateOsBlackWhiteVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOsBlackWhiteVersionsResponse{}
	_body, _err := client.UpdateOsBlackWhiteVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOsVersionWithOptions(request *UpdateOsVersionRequest, runtime *util.RuntimeOptions) (_result *UpdateOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackVersionList)) {
		query["BlackVersionList"] = request.BlackVersionList
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMobileDownload)) {
		query["EnableMobileDownload"] = request.EnableMobileDownload
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceNightUpgrade)) {
		query["IsForceNightUpgrade"] = request.IsForceNightUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceUpgrade)) {
		query["IsForceUpgrade"] = request.IsForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.IsMilestone)) {
		query["IsMilestone"] = request.IsMilestone
	}

	if !tea.BoolValue(util.IsUnset(request.MaxClientVersion)) {
		query["MaxClientVersion"] = request.MaxClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MinClientVersion)) {
		query["MinClientVersion"] = request.MinClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MobileDownloadMaxSize)) {
		query["MobileDownloadMaxSize"] = request.MobileDownloadMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.NightUpgradeDownloadType)) {
		query["NightUpgradeDownloadType"] = request.NightUpgradeDownloadType
	}

	if !tea.BoolValue(util.IsUnset(request.NightUpgradeIsAllowedCancel)) {
		query["NightUpgradeIsAllowedCancel"] = request.NightUpgradeIsAllowedCancel
	}

	if !tea.BoolValue(util.IsUnset(request.NightUpgradeIsShowTip)) {
		query["NightUpgradeIsShowTip"] = request.NightUpgradeIsShowTip
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RomList)) {
		query["RomList"] = request.RomList
	}

	if !tea.BoolValue(util.IsUnset(request.SystemVersion)) {
		query["SystemVersion"] = request.SystemVersion
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteVersionList)) {
		query["WhiteVersionList"] = request.WhiteVersionList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOsVersion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOsVersion(request *UpdateOsVersionRequest) (_result *UpdateOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOsVersionResponse{}
	_body, _err := client.UpdateOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOsVersionReleaseNoteWithOptions(request *UpdateOsVersionReleaseNoteRequest, runtime *util.RuntimeOptions) (_result *UpdateOsVersionReleaseNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseNote)) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOsVersionReleaseNote"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOsVersionReleaseNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOsVersionReleaseNote(request *UpdateOsVersionReleaseNoteRequest) (_result *UpdateOsVersionReleaseNoteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOsVersionReleaseNoteResponse{}
	_body, _err := client.UpdateOsVersionReleaseNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOsVersionRemarkWithOptions(request *UpdateOsVersionRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateOsVersionRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOsVersionRemark"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOsVersionRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOsVersionRemark(request *UpdateOsVersionRemarkRequest) (_result *UpdateOsVersionRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOsVersionRemarkResponse{}
	_body, _err := client.UpdateOsVersionRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOsVersionStatusWithOptions(request *UpdateOsVersionStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateOsVersionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOsVersionStatus"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOsVersionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOsVersionStatus(request *UpdateOsVersionStatusRequest) (_result *UpdateOsVersionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOsVersionStatusResponse{}
	_body, _err := client.UpdateOsVersionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSchemaSubscribeWithOptions(request *UpdateSchemaSubscribeRequest, runtime *util.RuntimeOptions) (_result *UpdateSchemaSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceModel)) {
		query["DeviceModel"] = request.DeviceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeList)) {
		query["SubscribeList"] = request.SubscribeList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSchemaSubscribe"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSchemaSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSchemaSubscribe(request *UpdateSchemaSubscribeRequest) (_result *UpdateSchemaSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSchemaSubscribeResponse{}
	_body, _err := client.UpdateSchemaSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateShadowSchemaWithOptions(request *UpdateShadowSchemaRequest, runtime *util.RuntimeOptions) (_result *UpdateShadowSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelId)) {
		query["DeviceModelId"] = request.DeviceModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		query["Schema"] = request.Schema
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateShadowSchema"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateShadowSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateShadowSchema(request *UpdateShadowSchemaRequest) (_result *UpdateShadowSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateShadowSchemaResponse{}
	_body, _err := client.UpdateShadowSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTriggerWithOptions(request *UpdateTriggerRequest, runtime *util.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Production)) {
		query["Production"] = request.Production
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Sandbox)) {
		query["Sandbox"] = request.Sandbox
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrigger"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrigger(request *UpdateTriggerRequest) (_result *UpdateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.UpdateTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUpstreamAppServerWithOptions(request *UpdateUpstreamAppServerRequest, runtime *util.RuntimeOptions) (_result *UpdateUpstreamAppServerResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUpstreamAppServer"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUpstreamAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUpstreamAppServer(request *UpdateUpstreamAppServerRequest) (_result *UpdateUpstreamAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUpstreamAppServerResponse{}
	_body, _err := client.UpdateUpstreamAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVersionDeviceGroupWithOptions(request *UpdateVersionDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateVersionDeviceGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVersionDeviceGroup"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVersionDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVersionDeviceGroup(request *UpdateVersionDeviceGroupRequest) (_result *UpdateVersionDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVersionDeviceGroupResponse{}
	_body, _err := client.UpdateVersionDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVersionPrepublishActiveStatusWithOptions(request *UpdateVersionPrepublishActiveStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateVersionPrepublishActiveStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsActive)) {
		query["IsActive"] = request.IsActive
	}

	if !tea.BoolValue(util.IsUnset(request.PrepublishId)) {
		query["PrepublishId"] = request.PrepublishId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVersionPrepublishActiveStatus"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVersionPrepublishActiveStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVersionPrepublishActiveStatus(request *UpdateVersionPrepublishActiveStatusRequest) (_result *UpdateVersionPrepublishActiveStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVersionPrepublishActiveStatusResponse{}
	_body, _err := client.UpdateVersionPrepublishActiveStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
