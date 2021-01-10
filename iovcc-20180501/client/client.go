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

type AddUploadedFunctionFileInfoRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s AddUploadedFunctionFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUploadedFunctionFileInfoRequest) GoString() string {
	return s.String()
}

func (s *AddUploadedFunctionFileInfoRequest) SetProjectId(v string) *AddUploadedFunctionFileInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *AddUploadedFunctionFileInfoRequest) SetObjectKey(v string) *AddUploadedFunctionFileInfoRequest {
	s.ObjectKey = &v
	return s
}

func (s *AddUploadedFunctionFileInfoRequest) SetFileName(v string) *AddUploadedFunctionFileInfoRequest {
	s.FileName = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUploadedFunctionFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddUploadedFunctionFileInfoResponse) SetBody(v *AddUploadedFunctionFileInfoResponseBody) *AddUploadedFunctionFileInfoResponse {
	s.Body = v
	return s
}

type AddVersionBlackDevicesRequest struct {
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s AddVersionBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionBlackDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddVersionBlackDevicesRequest) SetDeviceIds(v string) *AddVersionBlackDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetProjectId(v string) *AddVersionBlackDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetVersionType(v string) *AddVersionBlackDevicesRequest {
	s.VersionType = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetDeviceIdType(v string) *AddVersionBlackDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *AddVersionBlackDevicesRequest) SetVersionId(v string) *AddVersionBlackDevicesRequest {
	s.VersionId = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVersionBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddVersionBlackDevicesResponse) SetBody(v *AddVersionBlackDevicesResponseBody) *AddVersionBlackDevicesResponse {
	s.Body = v
	return s
}

type AddVersionGroupDevicesRequest struct {
	DeviceIdType  *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceIds     *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s AddVersionGroupDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionGroupDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddVersionGroupDevicesRequest) SetDeviceIdType(v string) *AddVersionGroupDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *AddVersionGroupDevicesRequest) SetProjectId(v string) *AddVersionGroupDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionGroupDevicesRequest) SetDeviceIds(v string) *AddVersionGroupDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *AddVersionGroupDevicesRequest) SetDeviceGroupId(v string) *AddVersionGroupDevicesRequest {
	s.DeviceGroupId = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVersionGroupDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddVersionGroupDevicesResponse) SetBody(v *AddVersionGroupDevicesResponseBody) *AddVersionGroupDevicesResponse {
	s.Body = v
	return s
}

type AddVersionWhiteDevicesRequest struct {
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s AddVersionWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesRequest) SetDeviceIds(v string) *AddVersionWhiteDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetProjectId(v string) *AddVersionWhiteDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetVersionType(v string) *AddVersionWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetDeviceIdType(v string) *AddVersionWhiteDevicesRequest {
	s.DeviceIdType = &v
	return s
}

func (s *AddVersionWhiteDevicesRequest) SetVersionId(v string) *AddVersionWhiteDevicesRequest {
	s.VersionId = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVersionWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddVersionWhiteDevicesResponse) SetBody(v *AddVersionWhiteDevicesResponseBody) *AddVersionWhiteDevicesResponse {
	s.Body = v
	return s
}

type AddVersionWhiteDevicesByDeviceGroupsRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	GroupIds    *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s AddVersionWhiteDevicesByDeviceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesByDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetProjectId(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetVersionType(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.VersionType = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetGroupIds(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsRequest) SetVersionId(v string) *AddVersionWhiteDevicesByDeviceGroupsRequest {
	s.VersionId = &v
	return s
}

type AddVersionWhiteDevicesByDeviceGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s AddVersionWhiteDevicesByDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVersionWhiteDevicesByDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponseBody) SetRequestId(v string) *AddVersionWhiteDevicesByDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVersionWhiteDevicesByDeviceGroupsResponseBody) SetData(v string) *AddVersionWhiteDevicesByDeviceGroupsResponseBody {
	s.Data = &v
	return s
}

type AddVersionWhiteDevicesByDeviceGroupsResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVersionWhiteDevicesByDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddVersionWhiteDevicesByDeviceGroupsResponse) SetBody(v *AddVersionWhiteDevicesByDeviceGroupsResponseBody) *AddVersionWhiteDevicesByDeviceGroupsResponse {
	s.Body = v
	return s
}

type ConnectAssistDeviceRequest struct {
	HardwareId            *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	AllowCommandExtension *bool   `json:"AllowCommandExtension,omitempty" xml:"AllowCommandExtension,omitempty"`
	DeviceId              *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	SerialNumber          *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	VIN                   *string `json:"VIN,omitempty" xml:"VIN,omitempty"`
	UUID                  *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ConnectAssistDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConnectAssistDeviceRequest) GoString() string {
	return s.String()
}

func (s *ConnectAssistDeviceRequest) SetHardwareId(v string) *ConnectAssistDeviceRequest {
	s.HardwareId = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetAllowCommandExtension(v bool) *ConnectAssistDeviceRequest {
	s.AllowCommandExtension = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetDeviceId(v string) *ConnectAssistDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetSerialNumber(v string) *ConnectAssistDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetVIN(v string) *ConnectAssistDeviceRequest {
	s.VIN = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetUUID(v string) *ConnectAssistDeviceRequest {
	s.UUID = &v
	return s
}

func (s *ConnectAssistDeviceRequest) SetProjectId(v string) *ConnectAssistDeviceRequest {
	s.ProjectId = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConnectAssistDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConnectAssistDeviceResponse) SetBody(v *ConnectAssistDeviceResponseBody) *ConnectAssistDeviceResponse {
	s.Body = v
	return s
}

type CountDeviceBrandsRequest struct {
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CountDeviceBrandsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceBrandsRequest) GoString() string {
	return s.String()
}

func (s *CountDeviceBrandsRequest) SetDeviceBrandId(v int64) *CountDeviceBrandsRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *CountDeviceBrandsRequest) SetDeviceBrand(v string) *CountDeviceBrandsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *CountDeviceBrandsRequest) SetProjectId(v string) *CountDeviceBrandsRequest {
	s.ProjectId = &v
	return s
}

type CountDeviceBrandsResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BrandCount *int32  `json:"BrandCount,omitempty" xml:"BrandCount,omitempty"`
}

func (s CountDeviceBrandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceBrandsResponseBody) GoString() string {
	return s.String()
}

func (s *CountDeviceBrandsResponseBody) SetRequestId(v string) *CountDeviceBrandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountDeviceBrandsResponseBody) SetBrandCount(v int32) *CountDeviceBrandsResponseBody {
	s.BrandCount = &v
	return s
}

type CountDeviceBrandsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountDeviceBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CountDeviceBrandsResponse) SetBody(v *CountDeviceBrandsResponseBody) *CountDeviceBrandsResponse {
	s.Body = v
	return s
}

type CountDeviceModelsRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s CountDeviceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceModelsRequest) GoString() string {
	return s.String()
}

func (s *CountDeviceModelsRequest) SetProjectId(v string) *CountDeviceModelsRequest {
	s.ProjectId = &v
	return s
}

func (s *CountDeviceModelsRequest) SetDeviceModelId(v int32) *CountDeviceModelsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CountDeviceModelsRequest) SetDeviceModel(v string) *CountDeviceModelsRequest {
	s.DeviceModel = &v
	return s
}

func (s *CountDeviceModelsRequest) SetDeviceBrand(v string) *CountDeviceModelsRequest {
	s.DeviceBrand = &v
	return s
}

type CountDeviceModelsResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ModelCount *int32  `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
}

func (s CountDeviceModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountDeviceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *CountDeviceModelsResponseBody) SetRequestId(v string) *CountDeviceModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountDeviceModelsResponseBody) SetModelCount(v int32) *CountDeviceModelsResponseBody {
	s.ModelCount = &v
	return s
}

type CountDeviceModelsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountDeviceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CountDeviceModelsResponse) SetBody(v *CountDeviceModelsResponseBody) *CountDeviceModelsResponse {
	s.Body = v
	return s
}

type CountDevicesRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
}

func (s CountDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s CountDevicesRequest) GoString() string {
	return s.String()
}

func (s *CountDevicesRequest) SetProjectId(v string) *CountDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *CountDevicesRequest) SetDeviceModelId(v int32) *CountDevicesRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CountDevicesRequest) SetDeviceModel(v string) *CountDevicesRequest {
	s.DeviceModel = &v
	return s
}

type CountDevicesResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceCount *int32  `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
}

func (s CountDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *CountDevicesResponseBody) SetRequestId(v string) *CountDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountDevicesResponseBody) SetDeviceCount(v int32) *CountDevicesResponseBody {
	s.DeviceCount = &v
	return s
}

type CountDevicesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CountDevicesResponse) SetBody(v *CountDevicesResponseBody) *CountDevicesResponse {
	s.Body = v
	return s
}

type CountProjectsResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectCount *int32  `json:"ProjectCount,omitempty" xml:"ProjectCount,omitempty"`
}

func (s CountProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *CountProjectsResponseBody) SetRequestId(v string) *CountProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountProjectsResponseBody) SetProjectCount(v int32) *CountProjectsResponseBody {
	s.ProjectCount = &v
	return s
}

type CountProjectsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountProjectsResponse) GoString() string {
	return s.String()
}

func (s *CountProjectsResponse) SetHeaders(v map[string]*string) *CountProjectsResponse {
	s.Headers = v
	return s
}

func (s *CountProjectsResponse) SetBody(v *CountProjectsResponseBody) *CountProjectsResponse {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountYunIdInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CountYunIdInfoResponse) SetBody(v *CountYunIdInfoResponseBody) *CountYunIdInfoResponse {
	s.Body = v
	return s
}

type CreateAppVersionRequest struct {
	IsForceUpgrade    *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsAllowNewInstall *string `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion        *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	InstallType       *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ReleaseNote       *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	IsSilentUpgrade   *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	PackageUrl        *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	IsNeedRestart     *string `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	BlackVersionList  *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	WhiteVersionList  *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
	RestartType       *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	RestartAppType    *string `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	RestartAppParam   *string `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	DeviceAdapterList *string `json:"DeviceAdapterList,omitempty" xml:"DeviceAdapterList,omitempty"`
	ApkMd5            *string `json:"ApkMd5,omitempty" xml:"ApkMd5,omitempty"`
}

func (s CreateAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppVersionRequest) SetIsForceUpgrade(v string) *CreateAppVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *CreateAppVersionRequest) SetIsAllowNewInstall(v string) *CreateAppVersionRequest {
	s.IsAllowNewInstall = &v
	return s
}

func (s *CreateAppVersionRequest) SetProjectId(v string) *CreateAppVersionRequest {
	s.ProjectId = &v
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

func (s *CreateAppVersionRequest) SetVersionCode(v string) *CreateAppVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *CreateAppVersionRequest) SetInstallType(v string) *CreateAppVersionRequest {
	s.InstallType = &v
	return s
}

func (s *CreateAppVersionRequest) SetRemark(v string) *CreateAppVersionRequest {
	s.Remark = &v
	return s
}

func (s *CreateAppVersionRequest) SetReleaseNote(v string) *CreateAppVersionRequest {
	s.ReleaseNote = &v
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

func (s *CreateAppVersionRequest) SetIsNeedRestart(v string) *CreateAppVersionRequest {
	s.IsNeedRestart = &v
	return s
}

func (s *CreateAppVersionRequest) SetBlackVersionList(v string) *CreateAppVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *CreateAppVersionRequest) SetWhiteVersionList(v string) *CreateAppVersionRequest {
	s.WhiteVersionList = &v
	return s
}

func (s *CreateAppVersionRequest) SetRestartType(v string) *CreateAppVersionRequest {
	s.RestartType = &v
	return s
}

func (s *CreateAppVersionRequest) SetRestartAppType(v string) *CreateAppVersionRequest {
	s.RestartAppType = &v
	return s
}

func (s *CreateAppVersionRequest) SetRestartAppParam(v string) *CreateAppVersionRequest {
	s.RestartAppParam = &v
	return s
}

func (s *CreateAppVersionRequest) SetDeviceAdapterList(v string) *CreateAppVersionRequest {
	s.DeviceAdapterList = &v
	return s
}

func (s *CreateAppVersionRequest) SetApkMd5(v string) *CreateAppVersionRequest {
	s.ApkMd5 = &v
	return s
}

type CreateAppVersionResponseBody struct {
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppVersionResponseBody) SetVersionId(v string) *CreateAppVersionResponseBody {
	s.VersionId = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetRequestId(v string) *CreateAppVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAppVersionResponse) SetBody(v *CreateAppVersionResponseBody) *CreateAppVersionResponse {
	s.Body = v
	return s
}

type CreateCustomizedFilterRequest struct {
	VersionType      *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	BlackWhiteType   *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ValueCompareType *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ValueType        *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedFilterRequest) SetVersionType(v string) *CreateCustomizedFilterRequest {
	s.VersionType = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetBlackWhiteType(v string) *CreateCustomizedFilterRequest {
	s.BlackWhiteType = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetValue(v string) *CreateCustomizedFilterRequest {
	s.Value = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetProjectId(v string) *CreateCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetValueCompareType(v string) *CreateCustomizedFilterRequest {
	s.ValueCompareType = &v
	return s
}

func (s *CreateCustomizedFilterRequest) SetName(v string) *CreateCustomizedFilterRequest {
	s.Name = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCustomizedFilterResponse) SetBody(v *CreateCustomizedFilterResponseBody) *CreateCustomizedFilterResponse {
	s.Body = v
	return s
}

type CreateCustomizedPropertyRequest struct {
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateCustomizedPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedPropertyRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedPropertyRequest) SetVersionId(v string) *CreateCustomizedPropertyRequest {
	s.VersionId = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetProjectId(v string) *CreateCustomizedPropertyRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetName(v string) *CreateCustomizedPropertyRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetValue(v string) *CreateCustomizedPropertyRequest {
	s.Value = &v
	return s
}

func (s *CreateCustomizedPropertyRequest) SetVersionType(v string) *CreateCustomizedPropertyRequest {
	s.VersionType = &v
	return s
}

type CreateCustomizedPropertyResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CustomizedPropertyId *string `json:"CustomizedPropertyId,omitempty" xml:"CustomizedPropertyId,omitempty"`
}

func (s CreateCustomizedPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedPropertyResponseBody) SetRequestId(v string) *CreateCustomizedPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomizedPropertyResponseBody) SetCustomizedPropertyId(v string) *CreateCustomizedPropertyResponseBody {
	s.CustomizedPropertyId = &v
	return s
}

type CreateCustomizedPropertyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomizedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCustomizedPropertyResponse) SetBody(v *CreateCustomizedPropertyResponseBody) *CreateCustomizedPropertyResponse {
	s.Body = v
	return s
}

type CreateDeviceRequest struct {
	ModelName  *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	HardwareId *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
}

func (s CreateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceRequest) SetModelName(v string) *CreateDeviceRequest {
	s.ModelName = &v
	return s
}

func (s *CreateDeviceRequest) SetProjectId(v string) *CreateDeviceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeviceRequest) SetHardwareId(v string) *CreateDeviceRequest {
	s.HardwareId = &v
	return s
}

type CreateDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceId  *int64  `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s CreateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponseBody) SetRequestId(v string) *CreateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceResponseBody) SetDeviceId(v int64) *CreateDeviceResponseBody {
	s.DeviceId = &v
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

type CreateDeviceBrandRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	BrandName   *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	Manufacture *string `json:"Manufacture,omitempty" xml:"Manufacture,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateDeviceBrandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceBrandRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceBrandRequest) SetProjectId(v string) *CreateDeviceBrandRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeviceBrandRequest) SetBrandName(v string) *CreateDeviceBrandRequest {
	s.BrandName = &v
	return s
}

func (s *CreateDeviceBrandRequest) SetManufacture(v string) *CreateDeviceBrandRequest {
	s.Manufacture = &v
	return s
}

func (s *CreateDeviceBrandRequest) SetDescription(v string) *CreateDeviceBrandRequest {
	s.Description = &v
	return s
}

type CreateDeviceBrandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BrandId   *int64  `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
}

func (s CreateDeviceBrandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceBrandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceBrandResponseBody) SetRequestId(v string) *CreateDeviceBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceBrandResponseBody) SetBrandId(v int64) *CreateDeviceBrandResponseBody {
	s.BrandId = &v
	return s
}

type CreateDeviceBrandResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceBrandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDeviceBrandResponse) SetBody(v *CreateDeviceBrandResponseBody) *CreateDeviceBrandResponse {
	s.Body = v
	return s
}

type CreateDeviceModelRequest struct {
	InitUsageType     *string `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	CanCreateDeviceId *string `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	ModelName         *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s CreateDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceModelRequest) SetInitUsageType(v string) *CreateDeviceModelRequest {
	s.InitUsageType = &v
	return s
}

func (s *CreateDeviceModelRequest) SetCanCreateDeviceId(v string) *CreateDeviceModelRequest {
	s.CanCreateDeviceId = &v
	return s
}

func (s *CreateDeviceModelRequest) SetModelName(v string) *CreateDeviceModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateDeviceModelRequest) SetHardwareType(v string) *CreateDeviceModelRequest {
	s.HardwareType = &v
	return s
}

func (s *CreateDeviceModelRequest) SetBrandName(v string) *CreateDeviceModelRequest {
	s.BrandName = &v
	return s
}

func (s *CreateDeviceModelRequest) SetDescription(v string) *CreateDeviceModelRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceModelRequest) SetDeviceType(v string) *CreateDeviceModelRequest {
	s.DeviceType = &v
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

func (s *CreateDeviceModelRequest) SetOsPlatform(v string) *CreateDeviceModelRequest {
	s.OsPlatform = &v
	return s
}

func (s *CreateDeviceModelRequest) SetObjectKey(v string) *CreateDeviceModelRequest {
	s.ObjectKey = &v
	return s
}

func (s *CreateDeviceModelRequest) SetDeviceName(v string) *CreateDeviceModelRequest {
	s.DeviceName = &v
	return s
}

type CreateDeviceModelResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
}

func (s CreateDeviceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceModelResponseBody) SetRequestId(v string) *CreateDeviceModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceModelResponseBody) SetDeviceModelId(v int64) *CreateDeviceModelResponseBody {
	s.DeviceModelId = &v
	return s
}

type CreateDeviceModelResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDeviceModelResponse) SetBody(v *CreateDeviceModelResponseBody) *CreateDeviceModelResponse {
	s.Body = v
	return s
}

type CreateMqttRootTopicRequest struct {
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	RootTopic   *string `json:"RootTopic,omitempty" xml:"RootTopic,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateMqttRootTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMqttRootTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateMqttRootTopicRequest) SetAppKey(v string) *CreateMqttRootTopicRequest {
	s.AppKey = &v
	return s
}

func (s *CreateMqttRootTopicRequest) SetRootTopic(v string) *CreateMqttRootTopicRequest {
	s.RootTopic = &v
	return s
}

func (s *CreateMqttRootTopicRequest) SetProjectId(v string) *CreateMqttRootTopicRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateMqttRootTopicRequest) SetClientToken(v string) *CreateMqttRootTopicRequest {
	s.ClientToken = &v
	return s
}

type CreateMqttRootTopicResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s CreateMqttRootTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMqttRootTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMqttRootTopicResponseBody) SetRequestId(v string) *CreateMqttRootTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMqttRootTopicResponseBody) SetQueueName(v string) *CreateMqttRootTopicResponseBody {
	s.QueueName = &v
	return s
}

type CreateMqttRootTopicResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMqttRootTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMqttRootTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMqttRootTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateMqttRootTopicResponse) SetHeaders(v map[string]*string) *CreateMqttRootTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateMqttRootTopicResponse) SetBody(v *CreateMqttRootTopicResponseBody) *CreateMqttRootTopicResponse {
	s.Body = v
	return s
}

type CreateNamespaceRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetProjectId(v string) *CreateNamespaceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateNamespaceRequest) SetAuthType(v string) *CreateNamespaceRequest {
	s.AuthType = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) SetDesc(v string) *CreateNamespaceRequest {
	s.Desc = &v
	return s
}

type CreateNamespaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetNamespace(v string) *CreateNamespaceResponseBody {
	s.Namespace = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type CreateOsVersionRequest struct {
	IsForceNightUpgrade         *string `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	MaxClientVersion            *string `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	ProjectId                   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModelId               *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	SystemVersion               *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	ReleaseNote                 *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark                      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	BlackVersionList            *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	IsMilestone                 *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	MinClientVersion            *string `json:"MinClientVersion,omitempty" xml:"MinClientVersion,omitempty"`
	WhiteVersionList            *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
	IsForceUpgrade              *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	NightUpgradeDownloadType    *string `json:"NightUpgradeDownloadType,omitempty" xml:"NightUpgradeDownloadType,omitempty"`
	NightUpgradeIsShowTip       *string `json:"NightUpgradeIsShowTip,omitempty" xml:"NightUpgradeIsShowTip,omitempty"`
	NightUpgradeIsAllowedCancel *string `json:"NightUpgradeIsAllowedCancel,omitempty" xml:"NightUpgradeIsAllowedCancel,omitempty"`
	RomList                     *string `json:"RomList,omitempty" xml:"RomList,omitempty"`
	EnableMobileDownload        *string `json:"EnableMobileDownload,omitempty" xml:"EnableMobileDownload,omitempty"`
	MobileDownloadMaxSize       *string `json:"MobileDownloadMaxSize,omitempty" xml:"MobileDownloadMaxSize,omitempty"`
}

func (s CreateOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateOsVersionRequest) SetIsForceNightUpgrade(v string) *CreateOsVersionRequest {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *CreateOsVersionRequest) SetMaxClientVersion(v string) *CreateOsVersionRequest {
	s.MaxClientVersion = &v
	return s
}

func (s *CreateOsVersionRequest) SetProjectId(v string) *CreateOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateOsVersionRequest) SetDeviceModelId(v string) *CreateOsVersionRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CreateOsVersionRequest) SetSystemVersion(v string) *CreateOsVersionRequest {
	s.SystemVersion = &v
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

func (s *CreateOsVersionRequest) SetBlackVersionList(v string) *CreateOsVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *CreateOsVersionRequest) SetIsMilestone(v string) *CreateOsVersionRequest {
	s.IsMilestone = &v
	return s
}

func (s *CreateOsVersionRequest) SetMinClientVersion(v string) *CreateOsVersionRequest {
	s.MinClientVersion = &v
	return s
}

func (s *CreateOsVersionRequest) SetWhiteVersionList(v string) *CreateOsVersionRequest {
	s.WhiteVersionList = &v
	return s
}

func (s *CreateOsVersionRequest) SetIsForceUpgrade(v string) *CreateOsVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *CreateOsVersionRequest) SetNightUpgradeDownloadType(v string) *CreateOsVersionRequest {
	s.NightUpgradeDownloadType = &v
	return s
}

func (s *CreateOsVersionRequest) SetNightUpgradeIsShowTip(v string) *CreateOsVersionRequest {
	s.NightUpgradeIsShowTip = &v
	return s
}

func (s *CreateOsVersionRequest) SetNightUpgradeIsAllowedCancel(v string) *CreateOsVersionRequest {
	s.NightUpgradeIsAllowedCancel = &v
	return s
}

func (s *CreateOsVersionRequest) SetRomList(v string) *CreateOsVersionRequest {
	s.RomList = &v
	return s
}

func (s *CreateOsVersionRequest) SetEnableMobileDownload(v string) *CreateOsVersionRequest {
	s.EnableMobileDownload = &v
	return s
}

func (s *CreateOsVersionRequest) SetMobileDownloadMaxSize(v string) *CreateOsVersionRequest {
	s.MobileDownloadMaxSize = &v
	return s
}

type CreateOsVersionResponseBody struct {
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOsVersionResponseBody) SetVersionId(v string) *CreateOsVersionResponseBody {
	s.VersionId = &v
	return s
}

func (s *CreateOsVersionResponseBody) SetRequestId(v string) *CreateOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateOsVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateOsVersionResponse) SetBody(v *CreateOsVersionResponseBody) *CreateOsVersionResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectDesc *string `json:"ProjectDesc,omitempty" xml:"ProjectDesc,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetProjectDesc(v string) *CreateProjectRequest {
	s.ProjectDesc = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateProjectAppRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPkgName *string `json:"AppPkgName,omitempty" xml:"AppPkgName,omitempty"`
	OsType     *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s CreateProjectAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectAppRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectAppRequest) SetProjectId(v string) *CreateProjectAppRequest {
	s.ProjectId = &v
	return s
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectAppResponse) SetBody(v *CreateProjectAppResponseBody) *CreateProjectAppResponse {
	s.Body = v
	return s
}

type CreateRpcServiceRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	InvokeType    *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	Params        *string `json:"Params,omitempty" xml:"Params,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MethodName    *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	VersionCode   *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateRpcServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRpcServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateRpcServiceRequest) SetProjectId(v string) *CreateRpcServiceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateRpcServiceRequest) SetAppKey(v string) *CreateRpcServiceRequest {
	s.AppKey = &v
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

func (s *CreateRpcServiceRequest) SetParams(v string) *CreateRpcServiceRequest {
	s.Params = &v
	return s
}

func (s *CreateRpcServiceRequest) SetGroupName(v string) *CreateRpcServiceRequest {
	s.GroupName = &v
	return s
}

func (s *CreateRpcServiceRequest) SetMethodName(v string) *CreateRpcServiceRequest {
	s.MethodName = &v
	return s
}

func (s *CreateRpcServiceRequest) SetVersionCode(v string) *CreateRpcServiceRequest {
	s.VersionCode = &v
	return s
}

type CreateRpcServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateRpcServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRpcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRpcServiceResponseBody) SetRequestId(v string) *CreateRpcServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRpcServiceResponseBody) SetId(v int64) *CreateRpcServiceResponseBody {
	s.Id = &v
	return s
}

type CreateRpcServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRpcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRpcServiceResponse) SetBody(v *CreateRpcServiceResponseBody) *CreateRpcServiceResponse {
	s.Body = v
	return s
}

type CreateSchemaSubscribeRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SubscribeList *string `json:"SubscribeList,omitempty" xml:"SubscribeList,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
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

func (s *CreateSchemaSubscribeRequest) SetSubscribeList(v string) *CreateSchemaSubscribeRequest {
	s.SubscribeList = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSchemaSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSchemaSubscribeResponse) SetBody(v *CreateSchemaSubscribeResponseBody) *CreateSchemaSubscribeResponse {
	s.Body = v
	return s
}

type CreateShadowSchemaRequest struct {
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	AuthType      *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
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

func (s *CreateShadowSchemaRequest) SetDeviceModelId(v string) *CreateShadowSchemaRequest {
	s.DeviceModelId = &v
	return s
}

func (s *CreateShadowSchemaRequest) SetAuthType(v string) *CreateShadowSchemaRequest {
	s.AuthType = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateShadowSchemaResponse) SetBody(v *CreateShadowSchemaResponseBody) *CreateShadowSchemaResponse {
	s.Body = v
	return s
}

type CreateTriggerRequest struct {
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	FileIds        *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	FunctionIds    *string `json:"FunctionIds,omitempty" xml:"FunctionIds,omitempty"`
	InvocationMode *int32  `json:"InvocationMode,omitempty" xml:"InvocationMode,omitempty"`
	Sandbox        *int32  `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Production     *int32  `json:"Production,omitempty" xml:"Production,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetProjectId(v string) *CreateTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateTriggerRequest) SetNamespace(v string) *CreateTriggerRequest {
	s.Namespace = &v
	return s
}

func (s *CreateTriggerRequest) SetSource(v string) *CreateTriggerRequest {
	s.Source = &v
	return s
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

func (s *CreateTriggerRequest) SetSandbox(v int32) *CreateTriggerRequest {
	s.Sandbox = &v
	return s
}

func (s *CreateTriggerRequest) SetProduction(v int32) *CreateTriggerRequest {
	s.Production = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateUpstreamAppKeyRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppKeyRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppKeyRelationResponseBody) SetRequestId(v string) *CreateUpstreamAppKeyRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUpstreamAppKeyRelationResponseBody) SetId(v int64) *CreateUpstreamAppKeyRelationResponseBody {
	s.Id = &v
	return s
}

type CreateUpstreamAppKeyRelationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUpstreamAppKeyRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUpstreamAppKeyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateUpstreamAppKeyRelationsResponse) SetBody(v *CreateUpstreamAppKeyRelationsResponseBody) *CreateUpstreamAppKeyRelationsResponse {
	s.Body = v
	return s
}

type CreateUpstreamAppServerRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateUpstreamAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppServerRequest) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppServerRequest) SetProjectId(v string) *CreateUpstreamAppServerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateUpstreamAppServerRequest) SetName(v string) *CreateUpstreamAppServerRequest {
	s.Name = &v
	return s
}

func (s *CreateUpstreamAppServerRequest) SetTags(v string) *CreateUpstreamAppServerRequest {
	s.Tags = &v
	return s
}

type CreateUpstreamAppServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateUpstreamAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUpstreamAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUpstreamAppServerResponseBody) SetRequestId(v string) *CreateUpstreamAppServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUpstreamAppServerResponseBody) SetId(v int64) *CreateUpstreamAppServerResponseBody {
	s.Id = &v
	return s
}

type CreateUpstreamAppServerResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUpstreamAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateUpstreamAppServerResponse) SetBody(v *CreateUpstreamAppServerResponseBody) *CreateUpstreamAppServerResponse {
	s.Body = v
	return s
}

type CreateVersionDeviceGroupRequest struct {
	MaxCount    *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionDeviceGroupRequest) SetMaxCount(v string) *CreateVersionDeviceGroupRequest {
	s.MaxCount = &v
	return s
}

func (s *CreateVersionDeviceGroupRequest) SetProjectId(v string) *CreateVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateVersionDeviceGroupRequest) SetName(v string) *CreateVersionDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateVersionDeviceGroupRequest) SetDescription(v string) *CreateVersionDeviceGroupRequest {
	s.Description = &v
	return s
}

type CreateVersionDeviceGroupResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s CreateVersionDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionDeviceGroupResponseBody) SetRequestId(v string) *CreateVersionDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVersionDeviceGroupResponseBody) SetDeviceGroupId(v string) *CreateVersionDeviceGroupResponseBody {
	s.DeviceGroupId = &v
	return s
}

type CreateVersionDeviceGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVersionDeviceGroupResponse) SetBody(v *CreateVersionDeviceGroupResponseBody) *CreateVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type CreateVersionPrepublishRequest struct {
	IsTotalPrepublish *string `json:"IsTotalPrepublish,omitempty" xml:"IsTotalPrepublish,omitempty"`
	VersionId         *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType       *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	BarrierCount      *string `json:"BarrierCount,omitempty" xml:"BarrierCount,omitempty"`
}

func (s CreateVersionPrepublishRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionPrepublishRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionPrepublishRequest) SetIsTotalPrepublish(v string) *CreateVersionPrepublishRequest {
	s.IsTotalPrepublish = &v
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

func (s *CreateVersionPrepublishRequest) SetName(v string) *CreateVersionPrepublishRequest {
	s.Name = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetProjectId(v string) *CreateVersionPrepublishRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateVersionPrepublishRequest) SetBarrierCount(v string) *CreateVersionPrepublishRequest {
	s.BarrierCount = &v
	return s
}

type CreateVersionPrepublishResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
}

func (s CreateVersionPrepublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionPrepublishResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionPrepublishResponseBody) SetRequestId(v string) *CreateVersionPrepublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVersionPrepublishResponseBody) SetPrepublishId(v string) *CreateVersionPrepublishResponseBody {
	s.PrepublishId = &v
	return s
}

type CreateVersionPrepublishResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVersionPrepublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVersionPrepublishResponse) SetBody(v *CreateVersionPrepublishResponseBody) *CreateVersionPrepublishResponse {
	s.Body = v
	return s
}

type CreateVersionTestRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType   *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateVersionTestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionTestRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionTestRequest) SetDeviceGroupId(v string) *CreateVersionTestRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *CreateVersionTestRequest) SetDescription(v string) *CreateVersionTestRequest {
	s.Description = &v
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

func (s *CreateVersionTestRequest) SetName(v string) *CreateVersionTestRequest {
	s.Name = &v
	return s
}

func (s *CreateVersionTestRequest) SetProjectId(v string) *CreateVersionTestRequest {
	s.ProjectId = &v
	return s
}

type CreateVersionTestResponseBody struct {
	TestId    *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVersionTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionTestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionTestResponseBody) SetTestId(v string) *CreateVersionTestResponseBody {
	s.TestId = &v
	return s
}

func (s *CreateVersionTestResponseBody) SetRequestId(v string) *CreateVersionTestResponseBody {
	s.RequestId = &v
	return s
}

type CreateVersionTestResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVersionTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVersionTestResponse) SetBody(v *CreateVersionTestResponseBody) *CreateVersionTestResponse {
	s.Body = v
	return s
}

type DelayPublishOsVersionRequest struct {
	VersionId       *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PrepubTime      *int64  `json:"PrepubTime,omitempty" xml:"PrepubTime,omitempty"`
	PublishTime     *int64  `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	DownTime        *int64  `json:"DownTime,omitempty" xml:"DownTime,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SendMessage     *string `json:"SendMessage,omitempty" xml:"SendMessage,omitempty"`
	PrepublishCount *string `json:"PrepublishCount,omitempty" xml:"PrepublishCount,omitempty"`
}

func (s DelayPublishOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DelayPublishOsVersionRequest) GoString() string {
	return s.String()
}

func (s *DelayPublishOsVersionRequest) SetVersionId(v string) *DelayPublishOsVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetProjectId(v string) *DelayPublishOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetPrepubTime(v int64) *DelayPublishOsVersionRequest {
	s.PrepubTime = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetPublishTime(v int64) *DelayPublishOsVersionRequest {
	s.PublishTime = &v
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

func (s *DelayPublishOsVersionRequest) SetDescription(v string) *DelayPublishOsVersionRequest {
	s.Description = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetSendMessage(v string) *DelayPublishOsVersionRequest {
	s.SendMessage = &v
	return s
}

func (s *DelayPublishOsVersionRequest) SetPrepublishCount(v string) *DelayPublishOsVersionRequest {
	s.PrepublishCount = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DelayPublishOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAllCustomizedFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAllCustomizedFiltersResponse) SetBody(v *DeleteAllCustomizedFiltersResponseBody) *DeleteAllCustomizedFiltersResponse {
	s.Body = v
	return s
}

type DeleteAllVersionGroupDevicesRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s DeleteAllVersionGroupDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllVersionGroupDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllVersionGroupDevicesRequest) SetProjectId(v string) *DeleteAllVersionGroupDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteAllVersionGroupDevicesRequest) SetDeviceGroupId(v string) *DeleteAllVersionGroupDevicesRequest {
	s.DeviceGroupId = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAllVersionGroupDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAllVersionGroupDevicesResponse) SetBody(v *DeleteAllVersionGroupDevicesResponseBody) *DeleteAllVersionGroupDevicesResponse {
	s.Body = v
	return s
}

type DeleteCustomizedFilterRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedFilterRequest) SetProjectId(v string) *DeleteCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteCustomizedFilterRequest) SetId(v string) *DeleteCustomizedFilterRequest {
	s.Id = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteCustomizedFilterResponse) SetBody(v *DeleteCustomizedFilterResponseBody) *DeleteCustomizedFilterResponse {
	s.Body = v
	return s
}

type DeleteCustomizedPropertyRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCustomizedPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizedPropertyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedPropertyRequest) SetProjectId(v string) *DeleteCustomizedPropertyRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteCustomizedPropertyRequest) SetId(v string) *DeleteCustomizedPropertyRequest {
	s.Id = &v
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCustomizedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteFunctionFileRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s DeleteFunctionFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionFileRequest) SetProjectId(v string) *DeleteFunctionFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFunctionFileRequest) SetFileName(v string) *DeleteFunctionFileRequest {
	s.FileName = &v
	return s
}

func (s *DeleteFunctionFileRequest) SetFileType(v int32) *DeleteFunctionFileRequest {
	s.FileType = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFunctionFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFunctionFileResponse) SetBody(v *DeleteFunctionFileResponseBody) *DeleteFunctionFileResponse {
	s.Body = v
	return s
}

type DeleteMqttRootTopicRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	RootTopic *string `json:"RootTopic,omitempty" xml:"RootTopic,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteMqttRootTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqttRootTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteMqttRootTopicRequest) SetAppKey(v string) *DeleteMqttRootTopicRequest {
	s.AppKey = &v
	return s
}

func (s *DeleteMqttRootTopicRequest) SetRootTopic(v string) *DeleteMqttRootTopicRequest {
	s.RootTopic = &v
	return s
}

func (s *DeleteMqttRootTopicRequest) SetProjectId(v string) *DeleteMqttRootTopicRequest {
	s.ProjectId = &v
	return s
}

type DeleteMqttRootTopicResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMqttRootTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqttRootTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMqttRootTopicResponseBody) SetRequestId(v string) *DeleteMqttRootTopicResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMqttRootTopicResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMqttRootTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMqttRootTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqttRootTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqttRootTopicResponse) SetHeaders(v map[string]*string) *DeleteMqttRootTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqttRootTopicResponse) SetBody(v *DeleteMqttRootTopicResponseBody) *DeleteMqttRootTopicResponse {
	s.Body = v
	return s
}

type DeleteNamespaceRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetProjectId(v string) *DeleteNamespaceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOpenAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteOpenAccountResponse) SetBody(v *DeleteOpenAccountResponseBody) *DeleteOpenAccountResponse {
	s.Body = v
	return s
}

type DeleteProjectAppRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteProjectAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectAppRequest) SetProjectId(v string) *DeleteProjectAppRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectAppRequest) SetAppId(v string) *DeleteProjectAppRequest {
	s.AppId = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRpcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSchemaSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteShadowSchemaResponse) SetBody(v *DeleteShadowSchemaResponseBody) *DeleteShadowSchemaResponse {
	s.Body = v
	return s
}

type DeleteTriggerRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteTriggerRequest) SetProjectId(v string) *DeleteTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteTriggerRequest) SetId(v int64) *DeleteTriggerRequest {
	s.Id = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUpstreamAppKeyRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteUpstreamAppKeyRelationResponse) SetBody(v *DeleteUpstreamAppKeyRelationResponseBody) *DeleteUpstreamAppKeyRelationResponse {
	s.Body = v
	return s
}

type DeleteUpstreamAppServerRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteUpstreamAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUpstreamAppServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteUpstreamAppServerRequest) SetProjectId(v string) *DeleteUpstreamAppServerRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteUpstreamAppServerRequest) SetId(v int64) *DeleteUpstreamAppServerRequest {
	s.Id = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUpstreamAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteUpstreamAppServerResponse) SetBody(v *DeleteUpstreamAppServerResponseBody) *DeleteUpstreamAppServerResponse {
	s.Body = v
	return s
}

type DeleteVersionAllBlackDevicesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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

func (s *DeleteVersionAllBlackDevicesRequest) SetVersionType(v string) *DeleteVersionAllBlackDevicesRequest {
	s.VersionType = &v
	return s
}

func (s *DeleteVersionAllBlackDevicesRequest) SetVersionId(v string) *DeleteVersionAllBlackDevicesRequest {
	s.VersionId = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionAllBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionAllBlackDevicesResponse) SetBody(v *DeleteVersionAllBlackDevicesResponseBody) *DeleteVersionAllBlackDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionAllWhiteDevicesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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

func (s *DeleteVersionAllWhiteDevicesRequest) SetVersionType(v string) *DeleteVersionAllWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

func (s *DeleteVersionAllWhiteDevicesRequest) SetVersionId(v string) *DeleteVersionAllWhiteDevicesRequest {
	s.VersionId = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionAllWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionAllWhiteDevicesResponse) SetBody(v *DeleteVersionAllWhiteDevicesResponseBody) *DeleteVersionAllWhiteDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionBlackDevicesRequest struct {
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
}

func (s DeleteVersionBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionBlackDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionBlackDevicesRequest) SetDeviceIds(v string) *DeleteVersionBlackDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetProjectId(v string) *DeleteVersionBlackDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetVersionType(v string) *DeleteVersionBlackDevicesRequest {
	s.VersionType = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetVersionId(v string) *DeleteVersionBlackDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionBlackDevicesRequest) SetDeviceIdType(v string) *DeleteVersionBlackDevicesRequest {
	s.DeviceIdType = &v
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionBlackDevicesResponse) SetBody(v *DeleteVersionBlackDevicesResponseBody) *DeleteVersionBlackDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionBlackDevicesByIdRequest struct {
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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

func (s *DeleteVersionBlackDevicesByIdRequest) SetVersionType(v string) *DeleteVersionBlackDevicesByIdRequest {
	s.VersionType = &v
	return s
}

func (s *DeleteVersionBlackDevicesByIdRequest) SetVersionId(v string) *DeleteVersionBlackDevicesByIdRequest {
	s.VersionId = &v
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionBlackDevicesByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionBlackDevicesByIdResponse) SetBody(v *DeleteVersionBlackDevicesByIdResponseBody) *DeleteVersionBlackDevicesByIdResponse {
	s.Body = v
	return s
}

type DeleteVersionDeviceGroupRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionDeviceGroupRequest) SetProjectId(v string) *DeleteVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionDeviceGroupRequest) SetId(v string) *DeleteVersionDeviceGroupRequest {
	s.Id = &v
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionDeviceGroupResponse) SetBody(v *DeleteVersionDeviceGroupResponseBody) *DeleteVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type DeleteVersionGroupDeviceRequest struct {
	DeviceIds     *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceIdType  *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
}

func (s DeleteVersionGroupDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceRequest) SetDeviceIds(v string) *DeleteVersionGroupDeviceRequest {
	s.DeviceIds = &v
	return s
}

func (s *DeleteVersionGroupDeviceRequest) SetProjectId(v string) *DeleteVersionGroupDeviceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionGroupDeviceRequest) SetDeviceGroupId(v string) *DeleteVersionGroupDeviceRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteVersionGroupDeviceRequest) SetDeviceIdType(v string) *DeleteVersionGroupDeviceRequest {
	s.DeviceIdType = &v
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionGroupDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionGroupDeviceResponse) SetBody(v *DeleteVersionGroupDeviceResponseBody) *DeleteVersionGroupDeviceResponse {
	s.Body = v
	return s
}

type DeleteVersionGroupDeviceByIdRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Ids           *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DeleteVersionGroupDeviceByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionGroupDeviceByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionGroupDeviceByIdRequest) SetProjectId(v string) *DeleteVersionGroupDeviceByIdRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionGroupDeviceByIdRequest) SetDeviceGroupId(v string) *DeleteVersionGroupDeviceByIdRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteVersionGroupDeviceByIdRequest) SetIds(v string) *DeleteVersionGroupDeviceByIdRequest {
	s.Ids = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionGroupDeviceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionGroupDeviceByIdResponse) SetBody(v *DeleteVersionGroupDeviceByIdResponseBody) *DeleteVersionGroupDeviceByIdResponse {
	s.Body = v
	return s
}

type DeleteVersionTestRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteVersionTestRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionTestRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionTestRequest) SetProjectId(v string) *DeleteVersionTestRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionTestRequest) SetId(v string) *DeleteVersionTestRequest {
	s.Id = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionTestResponse) SetBody(v *DeleteVersionTestResponseBody) *DeleteVersionTestResponse {
	s.Body = v
	return s
}

type DeleteVersionWhiteDevicesRequest struct {
	DeviceIds    *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
}

func (s DeleteVersionWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVersionWhiteDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteVersionWhiteDevicesRequest) SetDeviceIds(v string) *DeleteVersionWhiteDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetProjectId(v string) *DeleteVersionWhiteDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetVersionType(v string) *DeleteVersionWhiteDevicesRequest {
	s.VersionType = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetVersionId(v string) *DeleteVersionWhiteDevicesRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteVersionWhiteDevicesRequest) SetDeviceIdType(v string) *DeleteVersionWhiteDevicesRequest {
	s.DeviceIdType = &v
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionWhiteDevicesResponse) SetBody(v *DeleteVersionWhiteDevicesResponseBody) *DeleteVersionWhiteDevicesResponse {
	s.Body = v
	return s
}

type DeleteVersionWhiteDevicesByIdRequest struct {
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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

func (s *DeleteVersionWhiteDevicesByIdRequest) SetVersionType(v string) *DeleteVersionWhiteDevicesByIdRequest {
	s.VersionType = &v
	return s
}

func (s *DeleteVersionWhiteDevicesByIdRequest) SetVersionId(v string) *DeleteVersionWhiteDevicesByIdRequest {
	s.VersionId = &v
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVersionWhiteDevicesByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVersionWhiteDevicesByIdResponse) SetBody(v *DeleteVersionWhiteDevicesByIdResponseBody) *DeleteVersionWhiteDevicesByIdResponse {
	s.Body = v
	return s
}

type DeployFunctionFileRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileId    *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	DeployEnv *int32  `json:"DeployEnv,omitempty" xml:"DeployEnv,omitempty"`
}

func (s DeployFunctionFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionFileRequest) GoString() string {
	return s.String()
}

func (s *DeployFunctionFileRequest) SetProjectId(v string) *DeployFunctionFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeployFunctionFileRequest) SetFileId(v string) *DeployFunctionFileRequest {
	s.FileId = &v
	return s
}

func (s *DeployFunctionFileRequest) SetDeployEnv(v int32) *DeployFunctionFileRequest {
	s.DeployEnv = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployFunctionFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeployFunctionFileResponse) SetBody(v *DeployFunctionFileResponseBody) *DeployFunctionFileResponse {
	s.Body = v
	return s
}

type DescribeApiGatewayAppSecurityRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GatewayAppId *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
}

func (s DescribeApiGatewayAppSecurityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityRequest) SetProjectId(v string) *DescribeApiGatewayAppSecurityRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityRequest) SetGatewayAppId(v string) *DescribeApiGatewayAppSecurityRequest {
	s.GatewayAppId = &v
	return s
}

type DescribeApiGatewayAppSecurityResponseBody struct {
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApiGatewayAppSecurity *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity `json:"ApiGatewayAppSecurity,omitempty" xml:"ApiGatewayAppSecurity,omitempty" type:"Struct"`
}

func (s DescribeApiGatewayAppSecurityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityResponseBody) SetRequestId(v string) *DescribeApiGatewayAppSecurityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBody) SetApiGatewayAppSecurity(v *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) *DescribeApiGatewayAppSecurityResponseBody {
	s.ApiGatewayAppSecurity = v
	return s
}

type DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity struct {
	GatewayAppKey    *string `json:"GatewayAppKey,omitempty" xml:"GatewayAppKey,omitempty"`
	GatewayAppSecret *string `json:"GatewayAppSecret,omitempty" xml:"GatewayAppSecret,omitempty"`
	GatewayAppId     *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) GoString() string {
	return s.String()
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGatewayAppKey(v string) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GatewayAppKey = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGatewayAppSecret(v string) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GatewayAppSecret = &v
	return s
}

func (s *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity) SetGatewayAppId(v string) *DescribeApiGatewayAppSecurityResponseBodyApiGatewayAppSecurity {
	s.GatewayAppId = &v
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiGatewayAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Status            *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	IsAllowNewInstall *string                                             `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	ReleaseNote       *string                                             `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	PackageName       *string                                             `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	Remark            *string                                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	StatusName        *string                                             `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	ApkMd5            *string                                             `json:"ApkMd5,omitempty" xml:"ApkMd5,omitempty"`
	RestartAppParam   *string                                             `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	WhiteVersionList  *string                                             `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
	AppName           *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppId             *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RestartAppType    *string                                             `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	VersionCode       *int64                                              `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	BlackVersionList  *string                                             `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	GmtModify         *string                                             `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	DownloadUrl       *string                                             `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Adapters          []*DescribeAppVersionResponseBodyAppVersionAdapters `json:"Adapters,omitempty" xml:"Adapters,omitempty" type:"Repeated"`
	IsSilentUpgrade   *string                                             `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	InstallType       *string                                             `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	IsNeedRestart     *string                                             `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	Size              *string                                             `json:"Size,omitempty" xml:"Size,omitempty"`
	RestartType       *string                                             `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	GmtCreate         *string                                             `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Md5               *string                                             `json:"Md5,omitempty" xml:"Md5,omitempty"`
	AppVersion        *string                                             `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	IsForceUpgrade    *string                                             `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	Id                *int64                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	OriginalUrl       *string                                             `json:"OriginalUrl,omitempty" xml:"OriginalUrl,omitempty"`
}

func (s DescribeAppVersionResponseBodyAppVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppVersionResponseBodyAppVersion) GoString() string {
	return s.String()
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetStatus(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Status = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsAllowNewInstall(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsAllowNewInstall = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetReleaseNote(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetPackageName(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.PackageName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRemark(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Remark = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetStatusName(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.StatusName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetApkMd5(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.ApkMd5 = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRestartAppParam(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.RestartAppParam = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetWhiteVersionList(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.WhiteVersionList = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAppName(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.AppName = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAppId(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.AppId = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRestartAppType(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.RestartAppType = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetVersionCode(v int64) *DescribeAppVersionResponseBodyAppVersion {
	s.VersionCode = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetBlackVersionList(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.BlackVersionList = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetGmtModify(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.GmtModify = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetDownloadUrl(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAdapters(v []*DescribeAppVersionResponseBodyAppVersionAdapters) *DescribeAppVersionResponseBodyAppVersion {
	s.Adapters = v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsSilentUpgrade(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsSilentUpgrade = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetInstallType(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.InstallType = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsNeedRestart(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsNeedRestart = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetSize(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Size = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetRestartType(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.RestartType = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetGmtCreate(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetMd5(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.Md5 = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetAppVersion(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetIsForceUpgrade(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.IsForceUpgrade = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetId(v int64) *DescribeAppVersionResponseBodyAppVersion {
	s.Id = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersion) SetOriginalUrl(v string) *DescribeAppVersionResponseBodyAppVersion {
	s.OriginalUrl = &v
	return s
}

type DescribeAppVersionResponseBodyAppVersionAdapters struct {
	DeviceModelId   *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	MaxOsVersion    *string `json:"MaxOsVersion,omitempty" xml:"MaxOsVersion,omitempty"`
	MinOsVersion    *string `json:"MinOsVersion,omitempty" xml:"MinOsVersion,omitempty"`
	VersionId       *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	DeviceModelName *string `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
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

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetId(v int64) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.Id = &v
	return s
}

func (s *DescribeAppVersionResponseBodyAppVersionAdapters) SetDeviceModelName(v string) *DescribeAppVersionResponseBodyAppVersionAdapters {
	s.DeviceModelName = &v
	return s
}

type DescribeAppVersionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAppVersionResponse) SetBody(v *DescribeAppVersionResponseBody) *DescribeAppVersionResponse {
	s.Body = v
	return s
}

type DescribeAssistReportRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AssistId  *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
}

func (s DescribeAssistReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssistReportRequest) SetProjectId(v string) *DescribeAssistReportRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAssistReportRequest) SetAssistId(v string) *DescribeAssistReportRequest {
	s.AssistId = &v
	return s
}

type DescribeAssistReportResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AssistResult      *string `json:"AssistResult,omitempty" xml:"AssistResult,omitempty"`
	AssistReason      *string `json:"AssistReason,omitempty" xml:"AssistReason,omitempty"`
	AssistId          *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	AssistDescription *string `json:"AssistDescription,omitempty" xml:"AssistDescription,omitempty"`
	AssistTag         *string `json:"AssistTag,omitempty" xml:"AssistTag,omitempty"`
}

func (s DescribeAssistReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssistReportResponseBody) SetRequestId(v string) *DescribeAssistReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistResult(v string) *DescribeAssistReportResponseBody {
	s.AssistResult = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistReason(v string) *DescribeAssistReportResponseBody {
	s.AssistReason = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistId(v string) *DescribeAssistReportResponseBody {
	s.AssistId = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistDescription(v string) *DescribeAssistReportResponseBody {
	s.AssistDescription = &v
	return s
}

func (s *DescribeAssistReportResponseBody) SetAssistTag(v string) *DescribeAssistReportResponseBody {
	s.AssistTag = &v
	return s
}

type DescribeAssistReportResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAssistReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAssistReportResponse) SetBody(v *DescribeAssistReportResponseBody) *DescribeAssistReportResponse {
	s.Body = v
	return s
}

type DescribeAssistRTMPServerAddressRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s DescribeAssistRTMPServerAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistRTMPServerAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssistRTMPServerAddressRequest) SetProjectId(v string) *DescribeAssistRTMPServerAddressRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressRequest) SetDeviceId(v string) *DescribeAssistRTMPServerAddressRequest {
	s.DeviceId = &v
	return s
}

type DescribeAssistRTMPServerAddressResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RTMPServer *string `json:"RTMPServer,omitempty" xml:"RTMPServer,omitempty"`
	OTP        *string `json:"OTP,omitempty" xml:"OTP,omitempty"`
}

func (s DescribeAssistRTMPServerAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistRTMPServerAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssistRTMPServerAddressResponseBody) SetRequestId(v string) *DescribeAssistRTMPServerAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressResponseBody) SetRTMPServer(v string) *DescribeAssistRTMPServerAddressResponseBody {
	s.RTMPServer = &v
	return s
}

func (s *DescribeAssistRTMPServerAddressResponseBody) SetOTP(v string) *DescribeAssistRTMPServerAddressResponseBody {
	s.OTP = &v
	return s
}

type DescribeAssistRTMPServerAddressResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAssistRTMPServerAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAssistRTMPServerAddressResponse) SetBody(v *DescribeAssistRTMPServerAddressResponseBody) *DescribeAssistRTMPServerAddressResponse {
	s.Body = v
	return s
}

type DescribeAssistWSServerAddressRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s DescribeAssistWSServerAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistWSServerAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssistWSServerAddressRequest) SetProjectId(v string) *DescribeAssistWSServerAddressRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAssistWSServerAddressRequest) SetDeviceId(v string) *DescribeAssistWSServerAddressRequest {
	s.DeviceId = &v
	return s
}

type DescribeAssistWSServerAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WsServer  *string `json:"WsServer,omitempty" xml:"WsServer,omitempty"`
	OTP       *string `json:"OTP,omitempty" xml:"OTP,omitempty"`
}

func (s DescribeAssistWSServerAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssistWSServerAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssistWSServerAddressResponseBody) SetRequestId(v string) *DescribeAssistWSServerAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssistWSServerAddressResponseBody) SetWsServer(v string) *DescribeAssistWSServerAddressResponseBody {
	s.WsServer = &v
	return s
}

func (s *DescribeAssistWSServerAddressResponseBody) SetOTP(v string) *DescribeAssistWSServerAddressResponseBody {
	s.OTP = &v
	return s
}

type DescribeAssistWSServerAddressResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAssistWSServerAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAssistWSServerAddressResponse) SetBody(v *DescribeAssistWSServerAddressResponseBody) *DescribeAssistWSServerAddressResponse {
	s.Body = v
	return s
}

type DescribeCustomizedFilterRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCustomizedFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizedFilterRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedFilterRequest) SetProjectId(v string) *DescribeCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeCustomizedFilterRequest) SetId(v string) *DescribeCustomizedFilterRequest {
	s.Id = &v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ValueCompareType   *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	VersionId          *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueType          *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	BlackWhiteType     *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCustomizedFilterResponseBodyCustomizedFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizedFilterResponseBodyCustomizedFilter) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtCreateTimestamp(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtModify(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtModify = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetValueCompareType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.ValueCompareType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetVersionId(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.VersionId = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtModifyTimestamp(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetValue(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.Value = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetValueType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.ValueType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetGmtCreate(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetBlackWhiteType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.BlackWhiteType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetName(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.Name = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetVersionType(v string) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.VersionType = &v
	return s
}

func (s *DescribeCustomizedFilterResponseBodyCustomizedFilter) SetId(v int64) *DescribeCustomizedFilterResponseBodyCustomizedFilter {
	s.Id = &v
	return s
}

type DescribeCustomizedFilterResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCustomizedFilterResponse) SetBody(v *DescribeCustomizedFilterResponseBody) *DescribeCustomizedFilterResponse {
	s.Body = v
	return s
}

type DescribeDefaultSchemaRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
}

func (s DescribeDefaultSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefaultSchemaRequest) SetProjectId(v string) *DescribeDefaultSchemaRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDefaultSchemaRequest) SetDeviceModelId(v string) *DescribeDefaultSchemaRequest {
	s.DeviceModelId = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDefaultSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDefaultSchemaResponse) SetBody(v *DescribeDefaultSchemaResponseBody) *DescribeDefaultSchemaResponse {
	s.Body = v
	return s
}

type DescribeDeviceRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s DescribeDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceRequest) SetProjectId(v string) *DescribeDeviceRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceRequest) SetDeviceId(v string) *DescribeDeviceRequest {
	s.DeviceId = &v
	return s
}

type DescribeDeviceResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceInfo *DescribeDeviceResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s DescribeDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBody) SetRequestId(v string) *DescribeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDeviceInfo(v *DescribeDeviceResponseBodyDeviceInfo) *DescribeDeviceResponseBody {
	s.DeviceInfo = v
	return s
}

type DescribeDeviceResponseBodyDeviceInfo struct {
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	MacAddress    *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OsPlatform    *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	UsageType     *int32  `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	Vin           *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
	UsageTypeDesc *string `json:"UsageTypeDesc,omitempty" xml:"UsageTypeDesc,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	HardwareId    *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Attributes    *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	SoftwareId    *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceProduct *string `json:"DeviceProduct,omitempty" xml:"DeviceProduct,omitempty"`
}

func (s DescribeDeviceResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetSerialNumber(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetStatus(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.Status = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceModelId(v int64) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetMacAddress(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.MacAddress = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceId(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceType(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceType = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetProjectId(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetOsPlatform(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.OsPlatform = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceModel(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetUsageType(v int32) *DescribeDeviceResponseBodyDeviceInfo {
	s.UsageType = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetVin(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.Vin = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetUsageTypeDesc(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.UsageTypeDesc = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetUuid(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.Uuid = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetHardwareId(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.HardwareId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceBrandId(v int64) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceBrandId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetRegion(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.Region = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetAttributes(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.Attributes = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetSoftwareId(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.SoftwareId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetName(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.Name = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceBrand(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeDeviceResponseBodyDeviceInfo) SetDeviceProduct(v string) *DescribeDeviceResponseBodyDeviceInfo {
	s.DeviceProduct = &v
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

type DescribeDeviceBrandRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s DescribeDeviceBrandRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandRequest) SetProjectId(v string) *DescribeDeviceBrandRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetDeviceBrandId(v int64) *DescribeDeviceBrandRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetDeviceBrand(v string) *DescribeDeviceBrandRequest {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetStart(v string) *DescribeDeviceBrandRequest {
	s.Start = &v
	return s
}

func (s *DescribeDeviceBrandRequest) SetLength(v string) *DescribeDeviceBrandRequest {
	s.Length = &v
	return s
}

type DescribeDeviceBrandResponseBody struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceBrand *DescribeDeviceBrandResponseBodyDeviceBrand `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty" type:"Struct"`
}

func (s DescribeDeviceBrandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandResponseBody) SetRequestId(v string) *DescribeDeviceBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceBrandResponseBody) SetDeviceBrand(v *DescribeDeviceBrandResponseBodyDeviceBrand) *DescribeDeviceBrandResponseBody {
	s.DeviceBrand = v
	return s
}

type DescribeDeviceBrandResponseBodyDeviceBrand struct {
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Manufacture   *string `json:"Manufacture,omitempty" xml:"Manufacture,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s DescribeDeviceBrandResponseBodyDeviceBrand) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceBrandResponseBodyDeviceBrand) GoString() string {
	return s.String()
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetDeviceBrandId(v int64) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.DeviceBrandId = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetDescription(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.Description = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetProjectId(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetManufacture(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.Manufacture = &v
	return s
}

func (s *DescribeDeviceBrandResponseBodyDeviceBrand) SetDeviceBrand(v string) *DescribeDeviceBrandResponseBodyDeviceBrand {
	s.DeviceBrand = &v
	return s
}

type DescribeDeviceBrandResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceBrandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s DescribeDeviceIdByOuterInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceIdByOuterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceIdByOuterInfoResponseBody) SetRequestId(v string) *DescribeDeviceIdByOuterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceIdByOuterInfoResponseBody) SetDeviceId(v string) *DescribeDeviceIdByOuterInfoResponseBody {
	s.DeviceId = &v
	return s
}

type DescribeDeviceIdByOuterInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceIdByOuterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceInfo *DescribeDeviceInfoResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s DescribeDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBody) SetRequestId(v string) *DescribeDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetDeviceInfo(v *DescribeDeviceInfoResponseBodyDeviceInfo) *DescribeDeviceInfoResponseBody {
	s.DeviceInfo = v
	return s
}

type DescribeDeviceInfoResponseBodyDeviceInfo struct {
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	MacAddress    *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	UsageType     *int32  `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	Vin           *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
	UsageTypeDesc *string `json:"UsageTypeDesc,omitempty" xml:"UsageTypeDesc,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	HardwareId    *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SoftwareId    *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetSerialNumber(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetStatus(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Status = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceModelId(v int64) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetMacAddress(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.MacAddress = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceType(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetProjectId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceModel(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetUsageType(v int32) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.UsageType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetVin(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Vin = &v
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

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetHardwareId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.HardwareId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetRegion(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Region = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetSoftwareId(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.SoftwareId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetName(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.Name = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfo) SetDeviceBrand(v string) *DescribeDeviceInfoResponseBodyDeviceInfo {
	s.DeviceBrand = &v
	return s
}

type DescribeDeviceInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeviceInfoResponse) SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse {
	s.Body = v
	return s
}

type DescribeDeviceModelRequest struct {
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceModelRequest) SetDeviceModelId(v int32) *DescribeDeviceModelRequest {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceModelRequest) SetDeviceModel(v string) *DescribeDeviceModelRequest {
	s.DeviceModel = &v
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
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
	DeviceLogoUrl     *string `json:"DeviceLogoUrl,omitempty" xml:"DeviceLogoUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s DescribeDeviceModelResponseBodyDeviceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceModelResponseBodyDeviceModel) GoString() string {
	return s.String()
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceModelId(v int64) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetHardwareType(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.HardwareType = &v
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

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetCanCreateDeviceId(v int32) *DescribeDeviceModelResponseBodyDeviceModel {
	s.CanCreateDeviceId = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetProjectId(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.ProjectId = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetOsPlatform(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.OsPlatform = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceModel(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceModel = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetSecurityChip(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.SecurityChip = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceLogoUrl(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceLogoUrl = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDescription(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.Description = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetObjectKey(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.ObjectKey = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetInitUsageTypeDesc(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetInitUsageType(v int32) *DescribeDeviceModelResponseBodyDeviceModel {
	s.InitUsageType = &v
	return s
}

func (s *DescribeDeviceModelResponseBodyDeviceModel) SetDeviceBrand(v string) *DescribeDeviceModelResponseBodyDeviceModel {
	s.DeviceBrand = &v
	return s
}

type DescribeDeviceModelResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Devices   []*DescribeDeviceOnlineInfoResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
}

func (s DescribeDeviceOnlineInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceOnlineInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceOnlineInfoResponseBody) SetRequestId(v string) *DescribeDeviceOnlineInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBody) SetDevices(v []*DescribeDeviceOnlineInfoResponseBodyDevices) *DescribeDeviceOnlineInfoResponseBody {
	s.Devices = v
	return s
}

type DescribeDeviceOnlineInfoResponseBodyDevices struct {
	LoginTime     *int64  `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Online        *int32  `json:"Online,omitempty" xml:"Online,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IasId         *string `json:"IasId,omitempty" xml:"IasId,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	Terminal      *string `json:"Terminal,omitempty" xml:"Terminal,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s DescribeDeviceOnlineInfoResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceOnlineInfoResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetLoginTime(v int64) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.LoginTime = &v
	return s
}

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetDeviceId(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.DeviceId = &v
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

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetIasId(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.IasId = &v
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

func (s *DescribeDeviceOnlineInfoResponseBodyDevices) SetClientVersion(v string) *DescribeDeviceOnlineInfoResponseBodyDevices {
	s.ClientVersion = &v
	return s
}

type DescribeDeviceOnlineInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceOnlineInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeviceOnlineInfoResponse) SetBody(v *DescribeDeviceOnlineInfoResponseBody) *DescribeDeviceOnlineInfoResponse {
	s.Body = v
	return s
}

type DescribeDeviceShadowRequest struct {
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceModel    *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ViewSubscribed *bool   `json:"ViewSubscribed,omitempty" xml:"ViewSubscribed,omitempty"`
}

func (s DescribeDeviceShadowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowRequest) SetProjectId(v string) *DescribeDeviceShadowRequest {
	s.ProjectId = &v
	return s
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

func (s *DescribeDeviceShadowRequest) SetViewSubscribed(v bool) *DescribeDeviceShadowRequest {
	s.ViewSubscribed = &v
	return s
}

type DescribeDeviceShadowResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceShadow *DescribeDeviceShadowResponseBodyDeviceShadow `json:"DeviceShadow,omitempty" xml:"DeviceShadow,omitempty" type:"Struct"`
}

func (s DescribeDeviceShadowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowResponseBody) SetRequestId(v string) *DescribeDeviceShadowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceShadowResponseBody) SetDeviceShadow(v *DescribeDeviceShadowResponseBodyDeviceShadow) *DescribeDeviceShadowResponseBody {
	s.DeviceShadow = v
	return s
}

type DescribeDeviceShadowResponseBodyDeviceShadow struct {
	DeviceShadow *string `json:"DeviceShadow,omitempty" xml:"DeviceShadow,omitempty"`
	DeviceInfo   *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s DescribeDeviceShadowResponseBodyDeviceShadow) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceShadowResponseBodyDeviceShadow) GoString() string {
	return s.String()
}

func (s *DescribeDeviceShadowResponseBodyDeviceShadow) SetDeviceShadow(v string) *DescribeDeviceShadowResponseBodyDeviceShadow {
	s.DeviceShadow = &v
	return s
}

func (s *DescribeDeviceShadowResponseBodyDeviceShadow) SetDeviceInfo(v string) *DescribeDeviceShadowResponseBodyDeviceShadow {
	s.DeviceInfo = &v
	return s
}

type DescribeDeviceShadowResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceShadowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeviceShadowResponse) SetBody(v *DescribeDeviceShadowResponseBody) *DescribeDeviceShadowResponse {
	s.Body = v
	return s
}

type DescribeDeviceValiditySchemaRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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

func (s *DescribeDeviceValiditySchemaRequest) SetSchemaVersion(v string) *DescribeDeviceValiditySchemaRequest {
	s.SchemaVersion = &v
	return s
}

func (s *DescribeDeviceValiditySchemaRequest) SetProjectId(v string) *DescribeDeviceValiditySchemaRequest {
	s.ProjectId = &v
	return s
}

type DescribeDeviceValiditySchemaResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ItemList  []*DescribeDeviceValiditySchemaResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s DescribeDeviceValiditySchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceValiditySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceValiditySchemaResponseBody) SetRequestId(v string) *DescribeDeviceValiditySchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBody) SetItemList(v []*DescribeDeviceValiditySchemaResponseBodyItemList) *DescribeDeviceValiditySchemaResponseBody {
	s.ItemList = v
	return s
}

type DescribeDeviceValiditySchemaResponseBodyItemList struct {
	Minimum          *float32 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	Type             *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Maximum          *float32 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	ItemType         *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	EnumListStr      *string  `json:"EnumListStr,omitempty" xml:"EnumListStr,omitempty"`
	ExclusiveMinimum *bool    `json:"ExclusiveMinimum,omitempty" xml:"ExclusiveMinimum,omitempty"`
	MaxLength        *int32   `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	Required         *string  `json:"Required,omitempty" xml:"Required,omitempty"`
	Description      *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	ExclusiveMaximum *bool    `json:"ExclusiveMaximum,omitempty" xml:"ExclusiveMaximum,omitempty"`
	Path             *string  `json:"Path,omitempty" xml:"Path,omitempty"`
	MinLength        *int32   `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
}

func (s DescribeDeviceValiditySchemaResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceValiditySchemaResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMinimum(v float32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Minimum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetType(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Type = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMaximum(v float32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Maximum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetItemType(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.ItemType = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetEnumListStr(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.EnumListStr = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetExclusiveMinimum(v bool) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.ExclusiveMinimum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMaxLength(v int32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.MaxLength = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetRequired(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Required = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetDescription(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Description = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetExclusiveMaximum(v bool) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.ExclusiveMaximum = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetPath(v string) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.Path = &v
	return s
}

func (s *DescribeDeviceValiditySchemaResponseBodyItemList) SetMinLength(v int32) *DescribeDeviceValiditySchemaResponseBodyItemList {
	s.MinLength = &v
	return s
}

type DescribeDeviceValiditySchemaResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceValiditySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeviceValiditySchemaResponse) SetBody(v *DescribeDeviceValiditySchemaResponseBody) *DescribeDeviceValiditySchemaResponse {
	s.Body = v
	return s
}

type DescribeMessageRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	MessageId *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s DescribeMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMessageRequest) SetProjectId(v string) *DescribeMessageRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeMessageRequest) SetMessageId(v int64) *DescribeMessageRequest {
	s.MessageId = &v
	return s
}

type DescribeMessageResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *DescribeMessageResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
}

func (s DescribeMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMessageResponseBody) SetRequestId(v string) *DescribeMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMessageResponseBody) SetMessage(v *DescribeMessageResponseBodyMessage) *DescribeMessageResponseBody {
	s.Message = v
	return s
}

type DescribeMessageResponseBodyMessage struct {
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Action         *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PredictSendCnt *int32  `json:"PredictSendCnt,omitempty" xml:"PredictSendCnt,omitempty"`
	Uri            *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	Desc           *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	AuditMsg       *string `json:"AuditMsg,omitempty" xml:"AuditMsg,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppKey         *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	GmtCreateTime  *int64  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	ExipiredTime   *int64  `json:"ExipiredTime,omitempty" xml:"ExipiredTime,omitempty"`
	AckCnt         *int32  `json:"AckCnt,omitempty" xml:"AckCnt,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Parameter      *string `json:"Parameter,omitempty" xml:"Parameter,omitempty"`
	Audit          *int32  `json:"Audit,omitempty" xml:"Audit,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SendStatus     *int32  `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
}

func (s DescribeMessageResponseBodyMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeMessageResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *DescribeMessageResponseBodyMessage) SetType(v int32) *DescribeMessageResponseBodyMessage {
	s.Type = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAction(v string) *DescribeMessageResponseBodyMessage {
	s.Action = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetProjectId(v string) *DescribeMessageResponseBodyMessage {
	s.ProjectId = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetPredictSendCnt(v int32) *DescribeMessageResponseBodyMessage {
	s.PredictSendCnt = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetUri(v string) *DescribeMessageResponseBodyMessage {
	s.Uri = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetDesc(v string) *DescribeMessageResponseBodyMessage {
	s.Desc = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAuditMsg(v string) *DescribeMessageResponseBodyMessage {
	s.AuditMsg = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAppName(v string) *DescribeMessageResponseBodyMessage {
	s.AppName = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAppKey(v string) *DescribeMessageResponseBodyMessage {
	s.AppKey = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetGmtCreateTime(v int64) *DescribeMessageResponseBodyMessage {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetExipiredTime(v int64) *DescribeMessageResponseBodyMessage {
	s.ExipiredTime = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAckCnt(v int32) *DescribeMessageResponseBodyMessage {
	s.AckCnt = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetTitle(v string) *DescribeMessageResponseBodyMessage {
	s.Title = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetParameter(v string) *DescribeMessageResponseBodyMessage {
	s.Parameter = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetAudit(v int32) *DescribeMessageResponseBodyMessage {
	s.Audit = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetId(v int64) *DescribeMessageResponseBodyMessage {
	s.Id = &v
	return s
}

func (s *DescribeMessageResponseBodyMessage) SetSendStatus(v int32) *DescribeMessageResponseBodyMessage {
	s.SendStatus = &v
	return s
}

type DescribeMessageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeMessageResponse) SetBody(v *DescribeMessageResponseBody) *DescribeMessageResponse {
	s.Body = v
	return s
}

type DescribeMqttClientStatusRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ClientId  *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeMqttClientStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttClientStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeMqttClientStatusRequest) SetAppKey(v string) *DescribeMqttClientStatusRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeMqttClientStatusRequest) SetClientId(v string) *DescribeMqttClientStatusRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeMqttClientStatusRequest) SetProjectId(v string) *DescribeMqttClientStatusRequest {
	s.ProjectId = &v
	return s
}

type DescribeMqttClientStatusResponseBody struct {
	ClientStatus *DescribeMqttClientStatusResponseBodyClientStatus `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty" type:"Struct"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMqttClientStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttClientStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMqttClientStatusResponseBody) SetClientStatus(v *DescribeMqttClientStatusResponseBodyClientStatus) *DescribeMqttClientStatusResponseBody {
	s.ClientStatus = v
	return s
}

func (s *DescribeMqttClientStatusResponseBody) SetRequestId(v string) *DescribeMqttClientStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMqttClientStatusResponseBodyClientStatus struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	CleanSession *string `json:"CleanSession,omitempty" xml:"CleanSession,omitempty"`
	LastUpdate   *int64  `json:"LastUpdate,omitempty" xml:"LastUpdate,omitempty"`
}

func (s DescribeMqttClientStatusResponseBodyClientStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttClientStatusResponseBodyClientStatus) GoString() string {
	return s.String()
}

func (s *DescribeMqttClientStatusResponseBodyClientStatus) SetStatus(v int32) *DescribeMqttClientStatusResponseBodyClientStatus {
	s.Status = &v
	return s
}

func (s *DescribeMqttClientStatusResponseBodyClientStatus) SetCleanSession(v string) *DescribeMqttClientStatusResponseBodyClientStatus {
	s.CleanSession = &v
	return s
}

func (s *DescribeMqttClientStatusResponseBodyClientStatus) SetLastUpdate(v int64) *DescribeMqttClientStatusResponseBodyClientStatus {
	s.LastUpdate = &v
	return s
}

type DescribeMqttClientStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMqttClientStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMqttClientStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttClientStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeMqttClientStatusResponse) SetHeaders(v map[string]*string) *DescribeMqttClientStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeMqttClientStatusResponse) SetBody(v *DescribeMqttClientStatusResponseBody) *DescribeMqttClientStatusResponse {
	s.Body = v
	return s
}

type DescribeMqttMessageRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Mid       *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeMqttMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttMessageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMqttMessageRequest) SetAppKey(v string) *DescribeMqttMessageRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeMqttMessageRequest) SetMid(v string) *DescribeMqttMessageRequest {
	s.Mid = &v
	return s
}

func (s *DescribeMqttMessageRequest) SetProjectId(v string) *DescribeMqttMessageRequest {
	s.ProjectId = &v
	return s
}

type DescribeMqttMessageResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *DescribeMqttMessageResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
}

func (s DescribeMqttMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMqttMessageResponseBody) SetRequestId(v string) *DescribeMqttMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMqttMessageResponseBody) SetMessage(v *DescribeMqttMessageResponseBodyMessage) *DescribeMqttMessageResponseBody {
	s.Message = v
	return s
}

type DescribeMqttMessageResponseBodyMessage struct {
	Time    *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	AppKey  *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Mid     *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	Topic   *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	QoS     *int32  `json:"QoS,omitempty" xml:"QoS,omitempty"`
}

func (s DescribeMqttMessageResponseBodyMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttMessageResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *DescribeMqttMessageResponseBodyMessage) SetTime(v int64) *DescribeMqttMessageResponseBodyMessage {
	s.Time = &v
	return s
}

func (s *DescribeMqttMessageResponseBodyMessage) SetAppKey(v string) *DescribeMqttMessageResponseBodyMessage {
	s.AppKey = &v
	return s
}

func (s *DescribeMqttMessageResponseBodyMessage) SetMid(v string) *DescribeMqttMessageResponseBodyMessage {
	s.Mid = &v
	return s
}

func (s *DescribeMqttMessageResponseBodyMessage) SetTopic(v string) *DescribeMqttMessageResponseBodyMessage {
	s.Topic = &v
	return s
}

func (s *DescribeMqttMessageResponseBodyMessage) SetPayload(v string) *DescribeMqttMessageResponseBodyMessage {
	s.Payload = &v
	return s
}

func (s *DescribeMqttMessageResponseBodyMessage) SetQoS(v int32) *DescribeMqttMessageResponseBodyMessage {
	s.QoS = &v
	return s
}

type DescribeMqttMessageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMqttMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMqttMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttMessageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMqttMessageResponse) SetHeaders(v map[string]*string) *DescribeMqttMessageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMqttMessageResponse) SetBody(v *DescribeMqttMessageResponseBody) *DescribeMqttMessageResponse {
	s.Body = v
	return s
}

type DescribeMqttTopicSubscriptionRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeMqttTopicSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttTopicSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DescribeMqttTopicSubscriptionRequest) SetAppKey(v string) *DescribeMqttTopicSubscriptionRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeMqttTopicSubscriptionRequest) SetTopic(v string) *DescribeMqttTopicSubscriptionRequest {
	s.Topic = &v
	return s
}

func (s *DescribeMqttTopicSubscriptionRequest) SetProjectId(v string) *DescribeMqttTopicSubscriptionRequest {
	s.ProjectId = &v
	return s
}

type DescribeMqttTopicSubscriptionResponseBody struct {
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subscription *DescribeMqttTopicSubscriptionResponseBodySubscription `json:"Subscription,omitempty" xml:"Subscription,omitempty" type:"Struct"`
}

func (s DescribeMqttTopicSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttTopicSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMqttTopicSubscriptionResponseBody) SetRequestId(v string) *DescribeMqttTopicSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMqttTopicSubscriptionResponseBody) SetSubscription(v *DescribeMqttTopicSubscriptionResponseBodySubscription) *DescribeMqttTopicSubscriptionResponseBody {
	s.Subscription = v
	return s
}

type DescribeMqttTopicSubscriptionResponseBodySubscription struct {
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeMqttTopicSubscriptionResponseBodySubscription) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttTopicSubscriptionResponseBodySubscription) GoString() string {
	return s.String()
}

func (s *DescribeMqttTopicSubscriptionResponseBodySubscription) SetTopic(v string) *DescribeMqttTopicSubscriptionResponseBodySubscription {
	s.Topic = &v
	return s
}

func (s *DescribeMqttTopicSubscriptionResponseBodySubscription) SetCount(v int32) *DescribeMqttTopicSubscriptionResponseBodySubscription {
	s.Count = &v
	return s
}

type DescribeMqttTopicSubscriptionResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMqttTopicSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMqttTopicSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttTopicSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DescribeMqttTopicSubscriptionResponse) SetHeaders(v map[string]*string) *DescribeMqttTopicSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DescribeMqttTopicSubscriptionResponse) SetBody(v *DescribeMqttTopicSubscriptionResponseBody) *DescribeMqttTopicSubscriptionResponse {
	s.Body = v
	return s
}

type DescribeOpenAccountRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp        *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	IdToken    *string `json:"IdToken,omitempty" xml:"IdToken,omitempty"`
	OpenId     *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
}

func (s DescribeOpenAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenAccountRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenAccountRequest) SetProjectId(v string) *DescribeOpenAccountRequest {
	s.ProjectId = &v
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

func (s *DescribeOpenAccountRequest) SetIdToken(v string) *DescribeOpenAccountRequest {
	s.IdToken = &v
	return s
}

func (s *DescribeOpenAccountRequest) SetOpenId(v string) *DescribeOpenAccountRequest {
	s.OpenId = &v
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
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	CreateAccessKey *string `json:"CreateAccessKey,omitempty" xml:"CreateAccessKey,omitempty"`
	OpenId          *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IdentityId      *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	LoginId         *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	Idp             *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s DescribeOpenAccountResponseBodyOpenAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenAccountResponseBodyOpenAccount) GoString() string {
	return s.String()
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetStatus(v int32) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Status = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetType(v int32) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Type = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetDisplayName(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.DisplayName = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetCreateAccessKey(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.CreateAccessKey = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetOpenId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.OpenId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetMobile(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Mobile = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetRegion(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Region = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetIdentityId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.IdentityId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetLoginId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.LoginId = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetIdp(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.Idp = &v
	return s
}

func (s *DescribeOpenAccountResponseBodyOpenAccount) SetAliyunId(v string) *DescribeOpenAccountResponseBodyOpenAccount {
	s.AliyunId = &v
	return s
}

type DescribeOpenAccountResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOpenAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OsVersion *DescribeOsVersionResponseBodyOsVersion `json:"OsVersion,omitempty" xml:"OsVersion,omitempty" type:"Struct"`
}

func (s DescribeOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBody) SetRequestId(v string) *DescribeOsVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOsVersionResponseBody) SetOsVersion(v *DescribeOsVersionResponseBodyOsVersion) *DescribeOsVersionResponseBody {
	s.OsVersion = v
	return s
}

type DescribeOsVersionResponseBodyOsVersion struct {
	Status                *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceModelId         *string                                                   `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	BlackVersionList      *string                                                   `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	IsMilestone           *string                                                   `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	GmtModify             *string                                                   `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ReleaseNote           *string                                                   `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark                *string                                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SystemVersion         *string                                                   `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	StatusName            *string                                                   `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	DeviceModelName       *string                                                   `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	WhiteVersionList      *string                                                   `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
	MaxClientVersion      *string                                                   `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	RomList               []*DescribeOsVersionResponseBodyOsVersionRomList          `json:"RomList,omitempty" xml:"RomList,omitempty" type:"Repeated"`
	MinClientVersion      *string                                                   `json:"MinClientVersion,omitempty" xml:"MinClientVersion,omitempty"`
	NightUpgradeOption    *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption `json:"NightUpgradeOption,omitempty" xml:"NightUpgradeOption,omitempty" type:"Struct"`
	GmtCreate             *string                                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IsForceNightUpgrade   *string                                                   `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	MobileDownloadMaxSize *string                                                   `json:"MobileDownloadMaxSize,omitempty" xml:"MobileDownloadMaxSize,omitempty"`
	EnableMobileDownload  *string                                                   `json:"EnableMobileDownload,omitempty" xml:"EnableMobileDownload,omitempty"`
	IsForceUpgrade        *string                                                   `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	Id                    *int64                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeOsVersionResponseBodyOsVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBodyOsVersion) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetStatus(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.Status = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetDeviceModelId(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.DeviceModelId = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetBlackVersionList(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.BlackVersionList = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetIsMilestone(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.IsMilestone = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetGmtModify(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.GmtModify = &v
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

func (s *DescribeOsVersionResponseBodyOsVersion) SetSystemVersion(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.SystemVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetStatusName(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.StatusName = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetDeviceModelName(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.DeviceModelName = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetWhiteVersionList(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.WhiteVersionList = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetMaxClientVersion(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.MaxClientVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetRomList(v []*DescribeOsVersionResponseBodyOsVersionRomList) *DescribeOsVersionResponseBodyOsVersion {
	s.RomList = v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetMinClientVersion(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.MinClientVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetNightUpgradeOption(v *DescribeOsVersionResponseBodyOsVersionNightUpgradeOption) *DescribeOsVersionResponseBodyOsVersion {
	s.NightUpgradeOption = v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetGmtCreate(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetIsForceNightUpgrade(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetMobileDownloadMaxSize(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.MobileDownloadMaxSize = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetEnableMobileDownload(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.EnableMobileDownload = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetIsForceUpgrade(v string) *DescribeOsVersionResponseBodyOsVersion {
	s.IsForceUpgrade = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersion) SetId(v int64) *DescribeOsVersionResponseBodyOsVersion {
	s.Id = &v
	return s
}

type DescribeOsVersionResponseBodyOsVersionRomList struct {
	GmtModify   *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	SplitNum    *string `json:"SplitNum,omitempty" xml:"SplitNum,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Size        *string `json:"Size,omitempty" xml:"Size,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId   *int64  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OriginalUrl *string `json:"OriginalUrl,omitempty" xml:"OriginalUrl,omitempty"`
}

func (s DescribeOsVersionResponseBodyOsVersionRomList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOsVersionResponseBodyOsVersionRomList) GoString() string {
	return s.String()
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetGmtModify(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.GmtModify = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetSplitNum(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.SplitNum = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetDownloadUrl(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetSize(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.Size = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetGmtCreate(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetVersionId(v int64) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.VersionId = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetMd5(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.Md5 = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetBaseVersion(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.BaseVersion = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetId(v int64) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.Id = &v
	return s
}

func (s *DescribeOsVersionResponseBodyOsVersionRomList) SetOriginalUrl(v string) *DescribeOsVersionResponseBodyOsVersionRomList {
	s.OriginalUrl = &v
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

type DescribeOsVersionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
}

func (s DescribeProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBodyProject) SetStatus(v int32) *DescribeProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetDescription(v string) *DescribeProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetUserId(v string) *DescribeProjectResponseBodyProject {
	s.UserId = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetProjectId(v string) *DescribeProjectResponseBodyProject {
	s.ProjectId = &v
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

func (s *DescribeProjectResponseBodyProject) SetName(v string) *DescribeProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetId(v int64) *DescribeProjectResponseBodyProject {
	s.Id = &v
	return s
}

func (s *DescribeProjectResponseBodyProject) SetCreator(v string) *DescribeProjectResponseBodyProject {
	s.Creator = &v
	return s
}

type DescribeProjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeProjectResponse) SetBody(v *DescribeProjectResponseBody) *DescribeProjectResponse {
	s.Body = v
	return s
}

type DescribeProjectAppSecurityRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeProjectAppSecurityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectAppSecurityRequest) SetProjectId(v string) *DescribeProjectAppSecurityRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeProjectAppSecurityRequest) SetAppId(v string) *DescribeProjectAppSecurityRequest {
	s.AppId = &v
	return s
}

type DescribeProjectAppSecurityResponseBody struct {
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectAppSecurity *DescribeProjectAppSecurityResponseBodyProjectAppSecurity `json:"ProjectAppSecurity,omitempty" xml:"ProjectAppSecurity,omitempty" type:"Struct"`
}

func (s DescribeProjectAppSecurityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAppSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectAppSecurityResponseBody) SetRequestId(v string) *DescribeProjectAppSecurityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBody) SetProjectAppSecurity(v *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) *DescribeProjectAppSecurityResponseBody {
	s.ProjectAppSecurity = v
	return s
}

type DescribeProjectAppSecurityResponseBodyProjectAppSecurity struct {
	AppSecret   *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetAppSecret(v string) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.AppSecret = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetAppKey(v string) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.AppKey = &v
	return s
}

func (s *DescribeProjectAppSecurityResponseBodyProjectAppSecurity) SetAppId(v string) *DescribeProjectAppSecurityResponseBodyProjectAppSecurity {
	s.AppId = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProjectAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeShadowSchemaResponse) SetBody(v *DescribeShadowSchemaResponseBody) *DescribeShadowSchemaResponse {
	s.Body = v
	return s
}

type DescribeVersionDeviceGroupRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupRequest) SetProjectId(v string) *DescribeVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeVersionDeviceGroupRequest) SetId(v string) *DescribeVersionDeviceGroupRequest {
	s.Id = &v
	return s
}

type DescribeVersionDeviceGroupResponseBody struct {
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceGroup *DescribeVersionDeviceGroupResponseBodyDeviceGroup `json:"DeviceGroup,omitempty" xml:"DeviceGroup,omitempty" type:"Struct"`
}

func (s DescribeVersionDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupResponseBody) SetRequestId(v string) *DescribeVersionDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBody) SetDeviceGroup(v *DescribeVersionDeviceGroupResponseBodyDeviceGroup) *DescribeVersionDeviceGroupResponseBody {
	s.DeviceGroup = v
	return s
}

type DescribeVersionDeviceGroupResponseBodyDeviceGroup struct {
	GmtModify   *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxCount    *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s DescribeVersionDeviceGroupResponseBodyDeviceGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionDeviceGroupResponseBodyDeviceGroup) GoString() string {
	return s.String()
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetGmtModify(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.GmtModify = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetDescription(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.Description = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetGmtCreate(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVersionDeviceGroupResponseBodyDeviceGroup) SetName(v string) *DescribeVersionDeviceGroupResponseBodyDeviceGroup {
	s.Name = &v
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

type DescribeVersionDeviceGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVersionDeviceGroupResponse) SetBody(v *DescribeVersionDeviceGroupResponseBody) *DescribeVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type DiagnosisVersionRequest struct {
	OriginalId    *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionType   *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	IdType        *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	DiagnoseStyle *string `json:"DiagnoseStyle,omitempty" xml:"DiagnoseStyle,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DiagnosisVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DiagnosisVersionRequest) GoString() string {
	return s.String()
}

func (s *DiagnosisVersionRequest) SetOriginalId(v string) *DiagnosisVersionRequest {
	s.OriginalId = &v
	return s
}

func (s *DiagnosisVersionRequest) SetProjectId(v string) *DiagnosisVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *DiagnosisVersionRequest) SetVersionType(v string) *DiagnosisVersionRequest {
	s.VersionType = &v
	return s
}

func (s *DiagnosisVersionRequest) SetVersionId(v string) *DiagnosisVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DiagnosisVersionRequest) SetIdType(v string) *DiagnosisVersionRequest {
	s.IdType = &v
	return s
}

func (s *DiagnosisVersionRequest) SetDiagnoseStyle(v string) *DiagnosisVersionRequest {
	s.DiagnoseStyle = &v
	return s
}

func (s *DiagnosisVersionRequest) SetStartTime(v string) *DiagnosisVersionRequest {
	s.StartTime = &v
	return s
}

func (s *DiagnosisVersionRequest) SetEndTime(v string) *DiagnosisVersionRequest {
	s.EndTime = &v
	return s
}

type DiagnosisVersionResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DiagnosisResult *string `json:"DiagnosisResult,omitempty" xml:"DiagnosisResult,omitempty"`
}

func (s DiagnosisVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DiagnosisVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnosisVersionResponseBody) SetRequestId(v string) *DiagnosisVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnosisVersionResponseBody) SetDiagnosisResult(v string) *DiagnosisVersionResponseBody {
	s.DiagnosisResult = &v
	return s
}

type DiagnosisVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DiagnosisVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DiagnosisVersionResponse) SetBody(v *DiagnosisVersionResponseBody) *DiagnosisVersionResponse {
	s.Body = v
	return s
}

type FindAppVersionsRequest struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s FindAppVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *FindAppVersionsRequest) SetStatus(v string) *FindAppVersionsRequest {
	s.Status = &v
	return s
}

func (s *FindAppVersionsRequest) SetProjectId(v string) *FindAppVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *FindAppVersionsRequest) SetPageSize(v int32) *FindAppVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *FindAppVersionsRequest) SetVersionId(v string) *FindAppVersionsRequest {
	s.VersionId = &v
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

func (s *FindAppVersionsRequest) SetAppId(v string) *FindAppVersionsRequest {
	s.AppId = &v
	return s
}

func (s *FindAppVersionsRequest) SetRemark(v string) *FindAppVersionsRequest {
	s.Remark = &v
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
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	IsAllowNewInstall  *string `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	StatusName         *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	RestartAppParam    *string `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	IsSilentUpgrade    *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	AppPackageName     *string `json:"AppPackageName,omitempty" xml:"AppPackageName,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InstallType        *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	IsNeedRestart      *string `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	RestartAppType     *string `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RestartType        *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	AppVersion         *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	VersionCode        *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	IsForceUpgrade     *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindAppVersionsResponseBodyAppVersionListItems) String() string {
	return tea.Prettify(s)
}

func (s FindAppVersionsResponseBodyAppVersionListItems) GoString() string {
	return s.String()
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetStatus(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.Status = &v
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

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsAllowNewInstall(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsAllowNewInstall = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetStatusName(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.StatusName = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetRestartAppParam(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.RestartAppParam = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsSilentUpgrade(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsSilentUpgrade = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppPackageName(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppPackageName = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetGmtModifyTimestamp(v int64) *FindAppVersionsResponseBodyAppVersionListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppName(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppName = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetInstallType(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.InstallType = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsNeedRestart(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsNeedRestart = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetRestartAppType(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.RestartAppType = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppId(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppId = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetRestartType(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.RestartType = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetGmtCreate(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetAppVersion(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.AppVersion = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetVersionCode(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.VersionCode = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetIsForceUpgrade(v string) *FindAppVersionsResponseBodyAppVersionListItems {
	s.IsForceUpgrade = &v
	return s
}

func (s *FindAppVersionsResponseBodyAppVersionListItems) SetId(v int64) *FindAppVersionsResponseBodyAppVersionListItems {
	s.Id = &v
	return s
}

type FindAppVersionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindAppVersionsResponse) SetBody(v *FindAppVersionsResponseBody) *FindAppVersionsResponse {
	s.Body = v
	return s
}

type FindCustomizedFiltersRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindCustomizedFiltersRequest) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersRequest) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersRequest) SetProjectId(v string) *FindCustomizedFiltersRequest {
	s.ProjectId = &v
	return s
}

func (s *FindCustomizedFiltersRequest) SetVersionId(v string) *FindCustomizedFiltersRequest {
	s.VersionId = &v
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

func (s *FindCustomizedFiltersRequest) SetName(v string) *FindCustomizedFiltersRequest {
	s.Name = &v
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
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ValueCompareType   *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	BlackWhiteType     *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindCustomizedFiltersResponseBodyCustomizedFilterListItems) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedFiltersResponseBodyCustomizedFilterListItems) GoString() string {
	return s.String()
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtModifyTimestamp(v int64) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtCreateTimestamp(v int64) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetValue(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.Value = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtModify(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtModify = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetValueCompareType(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.ValueCompareType = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetGmtCreate(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetBlackWhiteType(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.BlackWhiteType = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetName(v string) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.Name = &v
	return s
}

func (s *FindCustomizedFiltersResponseBodyCustomizedFilterListItems) SetId(v int64) *FindCustomizedFiltersResponseBodyCustomizedFilterListItems {
	s.Id = &v
	return s
}

type FindCustomizedFiltersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindCustomizedFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindCustomizedFiltersResponse) SetBody(v *FindCustomizedFiltersResponseBody) *FindCustomizedFiltersResponse {
	s.Body = v
	return s
}

type FindCustomizedPropertiesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindCustomizedPropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesRequest) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesRequest) SetProjectId(v string) *FindCustomizedPropertiesRequest {
	s.ProjectId = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetVersionId(v string) *FindCustomizedPropertiesRequest {
	s.VersionId = &v
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

func (s *FindCustomizedPropertiesRequest) SetName(v string) *FindCustomizedPropertiesRequest {
	s.Name = &v
	return s
}

func (s *FindCustomizedPropertiesRequest) SetVersionType(v string) *FindCustomizedPropertiesRequest {
	s.VersionType = &v
	return s
}

type FindCustomizedPropertiesResponseBody struct {
	RequestId              *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CustomizedPropertyList *FindCustomizedPropertiesResponseBodyCustomizedPropertyList `json:"CustomizedPropertyList,omitempty" xml:"CustomizedPropertyList,omitempty" type:"Struct"`
}

func (s FindCustomizedPropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesResponseBody) SetRequestId(v string) *FindCustomizedPropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBody) SetCustomizedPropertyList(v *FindCustomizedPropertiesResponseBodyCustomizedPropertyList) *FindCustomizedPropertiesResponseBody {
	s.CustomizedPropertyList = v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) String() string {
	return tea.Prettify(s)
}

func (s FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) GoString() string {
	return s.String()
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetGmtCreateTimestamp(v int64) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetValue(v string) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.Value = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetGmtCreate(v string) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetName(v string) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.Name = &v
	return s
}

func (s *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems) SetId(v int64) *FindCustomizedPropertiesResponseBodyCustomizedPropertyListItems {
	s.Id = &v
	return s
}

type FindCustomizedPropertiesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindCustomizedPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindCustomizedPropertiesResponse) SetBody(v *FindCustomizedPropertiesResponseBody) *FindCustomizedPropertiesResponse {
	s.Body = v
	return s
}

type FindOsVersionsRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	IsMilestone   *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s FindOsVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsRequest) GoString() string {
	return s.String()
}

func (s *FindOsVersionsRequest) SetProjectId(v string) *FindOsVersionsRequest {
	s.ProjectId = &v
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

func (s *FindOsVersionsRequest) SetVersionId(v string) *FindOsVersionsRequest {
	s.VersionId = &v
	return s
}

func (s *FindOsVersionsRequest) SetDeviceModelId(v string) *FindOsVersionsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *FindOsVersionsRequest) SetSystemVersion(v string) *FindOsVersionsRequest {
	s.SystemVersion = &v
	return s
}

func (s *FindOsVersionsRequest) SetStatus(v string) *FindOsVersionsRequest {
	s.Status = &v
	return s
}

func (s *FindOsVersionsRequest) SetIsMilestone(v string) *FindOsVersionsRequest {
	s.IsMilestone = &v
	return s
}

func (s *FindOsVersionsRequest) SetRemark(v string) *FindOsVersionsRequest {
	s.Remark = &v
	return s
}

type FindOsVersionsResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OsVersionList *FindOsVersionsResponseBodyOsVersionList `json:"OsVersionList,omitempty" xml:"OsVersionList,omitempty" type:"Struct"`
}

func (s FindOsVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *FindOsVersionsResponseBody) SetRequestId(v string) *FindOsVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindOsVersionsResponseBody) SetOsVersionList(v *FindOsVersionsResponseBodyOsVersionList) *FindOsVersionsResponseBody {
	s.OsVersionList = v
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
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreateTimestamp  *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	DeviceModelId       *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	GmtModify           *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	IsMilestone         *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	Remark              *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SystemVersion       *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	StatusName          *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	IsForceReboot       *string `json:"IsForceReboot,omitempty" xml:"IsForceReboot,omitempty"`
	DeviceModelName     *string `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	IsSilentUpgrade     *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	GmtModifyTimestamp  *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	IsForceNightUpgrade *string `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IsForceUpgrade      *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindOsVersionsResponseBodyOsVersionListItems) String() string {
	return tea.Prettify(s)
}

func (s FindOsVersionsResponseBodyOsVersionListItems) GoString() string {
	return s.String()
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetStatus(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.Status = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtCreateTimestamp(v int64) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetDeviceModelId(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.DeviceModelId = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtModify(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtModify = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsMilestone(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsMilestone = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetRemark(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.Remark = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetSystemVersion(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.SystemVersion = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetStatusName(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.StatusName = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsForceReboot(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsForceReboot = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetDeviceModelName(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.DeviceModelName = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsSilentUpgrade(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsSilentUpgrade = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtModifyTimestamp(v int64) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsForceNightUpgrade(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetGmtCreate(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetIsForceUpgrade(v string) *FindOsVersionsResponseBodyOsVersionListItems {
	s.IsForceUpgrade = &v
	return s
}

func (s *FindOsVersionsResponseBodyOsVersionListItems) SetId(v int64) *FindOsVersionsResponseBodyOsVersionListItems {
	s.Id = &v
	return s
}

type FindOsVersionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindOsVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindOsVersionsResponse) SetBody(v *FindOsVersionsResponseBody) *FindOsVersionsResponse {
	s.Body = v
	return s
}

type FindPrepublishesByParentIdRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ParentId  *int32  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s FindPrepublishesByParentIdRequest) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdRequest) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdRequest) SetProjectId(v string) *FindPrepublishesByParentIdRequest {
	s.ProjectId = &v
	return s
}

func (s *FindPrepublishesByParentIdRequest) SetParentId(v int32) *FindPrepublishesByParentIdRequest {
	s.ParentId = &v
	return s
}

type FindPrepublishesByParentIdResponseBody struct {
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrepublishList *FindPrepublishesByParentIdResponseBodyPrepublishList `json:"PrepublishList,omitempty" xml:"PrepublishList,omitempty" type:"Struct"`
}

func (s FindPrepublishesByParentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdResponseBody) SetRequestId(v string) *FindPrepublishesByParentIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBody) SetPrepublishList(v *FindPrepublishesByParentIdResponseBodyPrepublishList) *FindPrepublishesByParentIdResponseBody {
	s.PrepublishList = v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	DeviceModelId      *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	IsActive           *string `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	BarrierCount       *string `json:"BarrierCount,omitempty" xml:"BarrierCount,omitempty"`
	IsTotalPrepublish  *string `json:"IsTotalPrepublish,omitempty" xml:"IsTotalPrepublish,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	ParentId           *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindPrepublishesByParentIdResponseBodyPrepublishListItems) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByParentIdResponseBodyPrepublishListItems) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtCreateTimestamp(v int64) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetDeviceModelId(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.DeviceModelId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtModify(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtModify = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetIsActive(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.IsActive = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetVersionId(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.VersionId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetBarrierCount(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.BarrierCount = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetIsTotalPrepublish(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.IsTotalPrepublish = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtModifyTimestamp(v int64) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetParentId(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.ParentId = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetGmtCreate(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetName(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.Name = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetId(v int64) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.Id = &v
	return s
}

func (s *FindPrepublishesByParentIdResponseBodyPrepublishListItems) SetVersionType(v string) *FindPrepublishesByParentIdResponseBodyPrepublishListItems {
	s.VersionType = &v
	return s
}

type FindPrepublishesByParentIdResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindPrepublishesByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrepublishList []*FindPrepublishesByVersionIdResponseBodyPrepublishList `json:"PrepublishList,omitempty" xml:"PrepublishList,omitempty" type:"Repeated"`
}

func (s FindPrepublishesByVersionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByVersionIdResponseBody) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByVersionIdResponseBody) SetRequestId(v string) *FindPrepublishesByVersionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBody) SetPrepublishList(v []*FindPrepublishesByVersionIdResponseBodyPrepublishList) *FindPrepublishesByVersionIdResponseBody {
	s.PrepublishList = v
	return s
}

type FindPrepublishesByVersionIdResponseBodyPrepublishList struct {
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	DeviceModelId      *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	IsActive           *string `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	BarrierCount       *string `json:"BarrierCount,omitempty" xml:"BarrierCount,omitempty"`
	DeviceModelName    *string `json:"DeviceModelName,omitempty" xml:"DeviceModelName,omitempty"`
	IsTotalPrepublish  *string `json:"IsTotalPrepublish,omitempty" xml:"IsTotalPrepublish,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	ParentId           *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	PassedCount        *string `json:"PassedCount,omitempty" xml:"PassedCount,omitempty"`
}

func (s FindPrepublishesByVersionIdResponseBodyPrepublishList) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishesByVersionIdResponseBodyPrepublishList) GoString() string {
	return s.String()
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtCreateTimestamp(v int64) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetDeviceModelId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.DeviceModelId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtModify(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtModify = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetIsActive(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.IsActive = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetVersionId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.VersionId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetBarrierCount(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.BarrierCount = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetDeviceModelName(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.DeviceModelName = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetIsTotalPrepublish(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.IsTotalPrepublish = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtModifyTimestamp(v int64) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetParentId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.ParentId = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetGmtCreate(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.GmtCreate = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetName(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.Name = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetId(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.Id = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetVersionType(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.VersionType = &v
	return s
}

func (s *FindPrepublishesByVersionIdResponseBodyPrepublishList) SetPassedCount(v string) *FindPrepublishesByVersionIdResponseBodyPrepublishList {
	s.PassedCount = &v
	return s
}

type FindPrepublishesByVersionIdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindPrepublishesByVersionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindPrepublishesByVersionIdResponse) SetBody(v *FindPrepublishesByVersionIdResponseBody) *FindPrepublishesByVersionIdResponse {
	s.Body = v
	return s
}

type FindPrepublishPassedDevicesRequest struct {
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s FindPrepublishPassedDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesRequest) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesRequest) SetPrepublishId(v string) *FindPrepublishPassedDevicesRequest {
	s.PrepublishId = &v
	return s
}

func (s *FindPrepublishPassedDevicesRequest) SetProjectId(v string) *FindPrepublishPassedDevicesRequest {
	s.ProjectId = &v
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

func (s *FindPrepublishPassedDevicesRequest) SetDeviceId(v string) *FindPrepublishPassedDevicesRequest {
	s.DeviceId = &v
	return s
}

type FindPrepublishPassedDevicesResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceList *FindPrepublishPassedDevicesResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Struct"`
}

func (s FindPrepublishPassedDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesResponseBody) SetRequestId(v string) *FindPrepublishPassedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBody) SetDeviceList(v *FindPrepublishPassedDevicesResponseBodyDeviceList) *FindPrepublishPassedDevicesResponseBody {
	s.DeviceList = v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s FindPrepublishPassedDevicesResponseBodyDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindPrepublishPassedDevicesResponseBodyDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceListItems) SetGmtCreateTimestamp(v int64) *FindPrepublishPassedDevicesResponseBodyDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceListItems) SetDeviceId(v string) *FindPrepublishPassedDevicesResponseBodyDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindPrepublishPassedDevicesResponseBodyDeviceListItems) SetGmtCreate(v string) *FindPrepublishPassedDevicesResponseBodyDeviceListItems {
	s.GmtCreate = &v
	return s
}

type FindPrepublishPassedDevicesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindPrepublishPassedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindPrepublishPassedDevicesResponse) SetBody(v *FindPrepublishPassedDevicesResponseBody) *FindPrepublishPassedDevicesResponse {
	s.Body = v
	return s
}

type FindVersionBlackDevicesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OriginalId  *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s FindVersionBlackDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesRequest) GoString() string {
	return s.String()
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

type FindVersionBlackDevicesResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceList *FindVersionBlackDevicesResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Struct"`
}

func (s FindVersionBlackDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesResponseBody) SetRequestId(v string) *FindVersionBlackDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBody) SetDeviceList(v *FindVersionBlackDevicesResponseBodyDeviceList) *FindVersionBlackDevicesResponseBody {
	s.DeviceList = v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	OriginalId         *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindVersionBlackDevicesResponseBodyDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionBlackDevicesResponseBodyDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetGmtCreateTimestamp(v int64) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetOriginalId(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.OriginalId = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetDeviceId(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetIdType(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.IdType = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetGmtCreate(v string) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionBlackDevicesResponseBodyDeviceListItems) SetId(v int64) *FindVersionBlackDevicesResponseBodyDeviceListItems {
	s.Id = &v
	return s
}

type FindVersionBlackDevicesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionBlackDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionBlackDevicesResponse) SetBody(v *FindVersionBlackDevicesResponseBody) *FindVersionBlackDevicesResponse {
	s.Body = v
	return s
}

type FindVersionDeviceGroupsRequest struct {
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OriginalId *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *FindVersionDeviceGroupsRequest) SetName(v string) *FindVersionDeviceGroupsRequest {
	s.Name = &v
	return s
}

func (s *FindVersionDeviceGroupsRequest) SetProjectId(v string) *FindVersionDeviceGroupsRequest {
	s.ProjectId = &v
	return s
}

type FindVersionDeviceGroupsResponseBody struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceGroupList *FindVersionDeviceGroupsResponseBodyDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Struct"`
}

func (s FindVersionDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsResponseBody) SetRequestId(v string) *FindVersionDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBody) SetDeviceGroupList(v *FindVersionDeviceGroupsResponseBodyDeviceGroupList) *FindVersionDeviceGroupsResponseBody {
	s.DeviceGroupList = v
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
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxCount           *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) GoString() string {
	return s.String()
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetGmtModifyTimestamp(v int64) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.GmtModifyTimestamp = &v
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

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetDescription(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.Description = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetGmtCreate(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems) SetName(v string) *FindVersionDeviceGroupsResponseBodyDeviceGroupListItems {
	s.Name = &v
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

type FindVersionDeviceGroupsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionDeviceGroupsResponse) SetBody(v *FindVersionDeviceGroupsResponseBody) *FindVersionDeviceGroupsResponse {
	s.Body = v
	return s
}

type FindVersionGroupDevicesRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OriginalId    *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s FindVersionGroupDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesRequest) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesRequest) SetProjectId(v string) *FindVersionGroupDevicesRequest {
	s.ProjectId = &v
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

func (s *FindVersionGroupDevicesRequest) SetDeviceGroupId(v string) *FindVersionGroupDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *FindVersionGroupDevicesRequest) SetDeviceId(v string) *FindVersionGroupDevicesRequest {
	s.DeviceId = &v
	return s
}

type FindVersionGroupDevicesResponseBody struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupDeviceList *FindVersionGroupDevicesResponseBodyGroupDeviceList `json:"GroupDeviceList,omitempty" xml:"GroupDeviceList,omitempty" type:"Struct"`
}

func (s FindVersionGroupDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesResponseBody) SetRequestId(v string) *FindVersionGroupDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBody) SetGroupDeviceList(v *FindVersionGroupDevicesResponseBodyGroupDeviceList) *FindVersionGroupDevicesResponseBody {
	s.GroupDeviceList = v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	OriginalId         *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindVersionGroupDevicesResponseBodyGroupDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionGroupDevicesResponseBodyGroupDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetGmtCreateTimestamp(v int64) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetOriginalId(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.OriginalId = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetDeviceId(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetIdType(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.IdType = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetGmtCreate(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionGroupDevicesResponseBodyGroupDeviceListItems) SetId(v string) *FindVersionGroupDevicesResponseBodyGroupDeviceListItems {
	s.Id = &v
	return s
}

type FindVersionGroupDevicesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionGroupDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionGroupDevicesResponse) SetBody(v *FindVersionGroupDevicesResponseBody) *FindVersionGroupDevicesResponse {
	s.Body = v
	return s
}

type FindVersionMessagesRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MessageType  *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	TestId       *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	SendRecordId *string `json:"SendRecordId,omitempty" xml:"SendRecordId,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesRequest) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesRequest) SetProjectId(v string) *FindVersionMessagesRequest {
	s.ProjectId = &v
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

func (s *FindVersionMessagesRequest) SetMessageType(v string) *FindVersionMessagesRequest {
	s.MessageType = &v
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

func (s *FindVersionMessagesRequest) SetSendRecordId(v string) *FindVersionMessagesRequest {
	s.SendRecordId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetDeviceId(v string) *FindVersionMessagesRequest {
	s.DeviceId = &v
	return s
}

func (s *FindVersionMessagesRequest) SetVersionType(v string) *FindVersionMessagesRequest {
	s.VersionType = &v
	return s
}

type FindVersionMessagesResponseBody struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MessageList *FindVersionMessagesResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Struct"`
}

func (s FindVersionMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesResponseBody) SetRequestId(v string) *FindVersionMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindVersionMessagesResponseBody) SetMessageList(v *FindVersionMessagesResponseBodyMessageList) *FindVersionMessagesResponseBody {
	s.MessageList = v
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
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	MessageId          *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	StatusDesc         *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	TestId             *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindVersionMessagesResponseBodyMessageListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessagesResponseBodyMessageListItems) GoString() string {
	return s.String()
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtModifyTimestamp(v int64) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtCreateTimestamp(v int64) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetStatus(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.Status = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtModify(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtModify = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetMessageId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.MessageId = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetDeviceId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetGmtCreate(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionMessagesResponseBodyMessageListItems) SetVersionId(v string) *FindVersionMessagesResponseBodyMessageListItems {
	s.VersionId = &v
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

func (s *FindVersionMessagesResponseBodyMessageListItems) SetId(v int64) *FindVersionMessagesResponseBodyMessageListItems {
	s.Id = &v
	return s
}

type FindVersionMessagesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionMessagesResponse) SetBody(v *FindVersionMessagesResponseBody) *FindVersionMessagesResponse {
	s.Body = v
	return s
}

type FindVersionMessageSendRecordsRequest struct {
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s FindVersionMessageSendRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsRequest) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsRequest) SetVersionType(v string) *FindVersionMessageSendRecordsRequest {
	s.VersionType = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetProjectId(v string) *FindVersionMessageSendRecordsRequest {
	s.ProjectId = &v
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

func (s *FindVersionMessageSendRecordsRequest) SetMessageType(v string) *FindVersionMessageSendRecordsRequest {
	s.MessageType = &v
	return s
}

func (s *FindVersionMessageSendRecordsRequest) SetVersionId(v string) *FindVersionMessageSendRecordsRequest {
	s.VersionId = &v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	MessageType        *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	FailedCount        *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	SkippedCount       *string `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	Result             *string `json:"Result,omitempty" xml:"Result,omitempty"`
	SucceededCount     *string `json:"SucceededCount,omitempty" xml:"SucceededCount,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ResultDesc         *string `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty"`
	TargetId           *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) GoString() string {
	return s.String()
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetGmtCreateTimestamp(v int64) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetMessageType(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.MessageType = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetFailedCount(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.FailedCount = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetSkippedCount(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.SkippedCount = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetResult(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.Result = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetSucceededCount(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.SucceededCount = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetGmtCreate(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetVersionId(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.VersionId = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetResultDesc(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.ResultDesc = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetTargetId(v string) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.TargetId = &v
	return s
}

func (s *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems) SetId(v int64) *FindVersionMessageSendRecordsResponseBodyMessageSendRecordListItems {
	s.Id = &v
	return s
}

type FindVersionMessageSendRecordsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionMessageSendRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionMessageSendRecordsResponse) SetBody(v *FindVersionMessageSendRecordsResponseBody) *FindVersionMessageSendRecordsResponse {
	s.Body = v
	return s
}

type FindVersionTestsRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionTestsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsRequest) GoString() string {
	return s.String()
}

func (s *FindVersionTestsRequest) SetProjectId(v string) *FindVersionTestsRequest {
	s.ProjectId = &v
	return s
}

func (s *FindVersionTestsRequest) SetPageIndex(v int32) *FindVersionTestsRequest {
	s.PageIndex = &v
	return s
}

func (s *FindVersionTestsRequest) SetPageSize(v int32) *FindVersionTestsRequest {
	s.PageSize = &v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	GmtModify          *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	GmtModifyTimestamp *int64  `json:"GmtModifyTimestamp,omitempty" xml:"GmtModifyTimestamp,omitempty"`
	FailedCount        *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	DeviceGroupId      *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	SkippedCount       *string `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SucceededCount     *string `json:"SucceededCount,omitempty" xml:"SucceededCount,omitempty"`
	DeviceGroupName    *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	VersionType        *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s FindVersionTestsResponseBodyVersionTestListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionTestsResponseBodyVersionTestListItems) GoString() string {
	return s.String()
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtCreateTimestamp(v int64) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtModify(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtModify = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetVersionId(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.VersionId = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtModifyTimestamp(v int64) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtModifyTimestamp = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetFailedCount(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.FailedCount = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetDeviceGroupId(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.DeviceGroupId = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetSkippedCount(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.SkippedCount = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetDescription(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.Description = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetSucceededCount(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.SucceededCount = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetDeviceGroupName(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.DeviceGroupName = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetGmtCreate(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetName(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.Name = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetId(v int64) *FindVersionTestsResponseBodyVersionTestListItems {
	s.Id = &v
	return s
}

func (s *FindVersionTestsResponseBodyVersionTestListItems) SetVersionType(v string) *FindVersionTestsResponseBodyVersionTestListItems {
	s.VersionType = &v
	return s
}

type FindVersionTestsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionTestsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionTestsResponse) SetBody(v *FindVersionTestsResponseBody) *FindVersionTestsResponse {
	s.Body = v
	return s
}

type FindVersionWhiteDevicesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OriginalId  *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s FindVersionWhiteDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesRequest) GoString() string {
	return s.String()
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

type FindVersionWhiteDevicesResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceList *FindVersionWhiteDevicesResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Struct"`
}

func (s FindVersionWhiteDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesResponseBody) SetRequestId(v string) *FindVersionWhiteDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBody) SetDeviceList(v *FindVersionWhiteDevicesResponseBodyDeviceList) *FindVersionWhiteDevicesResponseBody {
	s.DeviceList = v
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
	GmtCreateTimestamp *int64  `json:"GmtCreateTimestamp,omitempty" xml:"GmtCreateTimestamp,omitempty"`
	OriginalId         *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	DeviceId           *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FindVersionWhiteDevicesResponseBodyDeviceListItems) String() string {
	return tea.Prettify(s)
}

func (s FindVersionWhiteDevicesResponseBodyDeviceListItems) GoString() string {
	return s.String()
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetGmtCreateTimestamp(v int64) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.GmtCreateTimestamp = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetOriginalId(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.OriginalId = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetDeviceId(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.DeviceId = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetIdType(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.IdType = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetGmtCreate(v string) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.GmtCreate = &v
	return s
}

func (s *FindVersionWhiteDevicesResponseBodyDeviceListItems) SetId(v int64) *FindVersionWhiteDevicesResponseBodyDeviceListItems {
	s.Id = &v
	return s
}

type FindVersionWhiteDevicesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindVersionWhiteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FindVersionWhiteDevicesResponse) SetBody(v *FindVersionWhiteDevicesResponseBody) *FindVersionWhiteDevicesResponse {
	s.Body = v
	return s
}

type GenerateAssistFileUploadUrlRequest struct {
	Filename  *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s GenerateAssistFileUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAssistFileUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateAssistFileUploadUrlRequest) SetFilename(v string) *GenerateAssistFileUploadUrlRequest {
	s.Filename = &v
	return s
}

func (s *GenerateAssistFileUploadUrlRequest) SetProjectId(v string) *GenerateAssistFileUploadUrlRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateAssistFileUploadUrlRequest) SetDeviceId(v string) *GenerateAssistFileUploadUrlRequest {
	s.DeviceId = &v
	return s
}

type GenerateAssistFileUploadUrlResponseBody struct {
	FileKey   *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAssistFileUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAssistFileUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAssistFileUploadUrlResponseBody) SetFileKey(v string) *GenerateAssistFileUploadUrlResponseBody {
	s.FileKey = &v
	return s
}

func (s *GenerateAssistFileUploadUrlResponseBody) SetUploadUrl(v string) *GenerateAssistFileUploadUrlResponseBody {
	s.UploadUrl = &v
	return s
}

func (s *GenerateAssistFileUploadUrlResponseBody) SetRequestId(v string) *GenerateAssistFileUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAssistFileUploadUrlResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateAssistFileUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateAssistFileUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAssistFileUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateAssistFileUploadUrlResponse) SetHeaders(v map[string]*string) *GenerateAssistFileUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateAssistFileUploadUrlResponse) SetBody(v *GenerateAssistFileUploadUrlResponseBody) *GenerateAssistFileUploadUrlResponse {
	s.Body = v
	return s
}

type GenerateFunctionFileUploadMetaRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GenerateFunctionFileUploadMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaRequest) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaRequest) SetProjectId(v string) *GenerateFunctionFileUploadMetaRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaRequest) SetFileName(v string) *GenerateFunctionFileUploadMetaRequest {
	s.FileName = &v
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
	PostObjectPolicy *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy `json:"PostObjectPolicy,omitempty" xml:"PostObjectPolicy,omitempty" type:"Struct"`
	SecurityToken    *string                                                               `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ObjectKey        *string                                                               `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMeta) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMeta) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) SetPostObjectPolicy(v *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) *GenerateFunctionFileUploadMetaResponseBodyUploadMeta {
	s.PostObjectPolicy = v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) SetSecurityToken(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMeta {
	s.SecurityToken = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMeta) SetObjectKey(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMeta {
	s.ObjectKey = &v
	return s
}

type GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy struct {
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) GoString() string {
	return s.String()
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetSignature(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.Signature = &v
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

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetExpire(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.Expire = &v
	return s
}

func (s *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy) SetAccessId(v string) *GenerateFunctionFileUploadMetaResponseBodyUploadMetaPostObjectPolicy {
	s.AccessId = &v
	return s
}

type GenerateFunctionFileUploadMetaResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateFunctionFileUploadMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateFunctionFileUploadMetaResponse) SetBody(v *GenerateFunctionFileUploadMetaResponseBody) *GenerateFunctionFileUploadMetaResponse {
	s.Body = v
	return s
}

type GenerateOssPostPolicyRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
}

func (s GenerateOssPostPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyRequest) SetProjectId(v string) *GenerateOssPostPolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateOssPostPolicyRequest) SetExt(v string) *GenerateOssPostPolicyRequest {
	s.Ext = &v
	return s
}

func (s *GenerateOssPostPolicyRequest) SetAccessId(v string) *GenerateOssPostPolicyRequest {
	s.AccessId = &v
	return s
}

func (s *GenerateOssPostPolicyRequest) SetAccessKey(v string) *GenerateOssPostPolicyRequest {
	s.AccessKey = &v
	return s
}

type GenerateOssPostPolicyResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OssPostPolicy *GenerateOssPostPolicyResponseBodyOssPostPolicy `json:"OssPostPolicy,omitempty" xml:"OssPostPolicy,omitempty" type:"Struct"`
}

func (s GenerateOssPostPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyResponseBody) SetRequestId(v string) *GenerateOssPostPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOssPostPolicyResponseBody) SetOssPostPolicy(v *GenerateOssPostPolicyResponseBodyOssPostPolicy) *GenerateOssPostPolicyResponseBody {
	s.OssPostPolicy = v
	return s
}

type GenerateOssPostPolicyResponseBodyOssPostPolicy struct {
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
}

func (s GenerateOssPostPolicyResponseBodyOssPostPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssPostPolicyResponseBodyOssPostPolicy) GoString() string {
	return s.String()
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetSignature(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.Signature = &v
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

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetExpire(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.Expire = &v
	return s
}

func (s *GenerateOssPostPolicyResponseBodyOssPostPolicy) SetAccessId(v string) *GenerateOssPostPolicyResponseBodyOssPostPolicy {
	s.AccessId = &v
	return s
}

type GenerateOssPostPolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateOssPostPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateOssPostPolicyResponse) SetBody(v *GenerateOssPostPolicyResponseBody) *GenerateOssPostPolicyResponse {
	s.Body = v
	return s
}

type GenerateOssUploadMetaRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
}

func (s GenerateOssUploadMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaRequest) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaRequest) SetProjectId(v string) *GenerateOssUploadMetaRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateOssUploadMetaRequest) SetExt(v string) *GenerateOssUploadMetaRequest {
	s.Ext = &v
	return s
}

type GenerateOssUploadMetaResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OssUploadMeta *GenerateOssUploadMetaResponseBodyOssUploadMeta `json:"OssUploadMeta,omitempty" xml:"OssUploadMeta,omitempty" type:"Struct"`
}

func (s GenerateOssUploadMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaResponseBody) SetRequestId(v string) *GenerateOssUploadMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBody) SetOssUploadMeta(v *GenerateOssUploadMetaResponseBodyOssUploadMeta) *GenerateOssUploadMetaResponseBody {
	s.OssUploadMeta = v
	return s
}

type GenerateOssUploadMetaResponseBodyOssUploadMeta struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ObjectKey       *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Bucket          *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s GenerateOssUploadMetaResponseBodyOssUploadMeta) String() string {
	return tea.Prettify(s)
}

func (s GenerateOssUploadMetaResponseBodyOssUploadMeta) GoString() string {
	return s.String()
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetSecurityToken(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.SecurityToken = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetObjectKey(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.ObjectKey = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetAccessKeySecret(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.AccessKeySecret = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetAccessKeyId(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.AccessKeyId = &v
	return s
}

func (s *GenerateOssUploadMetaResponseBodyOssUploadMeta) SetBucket(v string) *GenerateOssUploadMetaResponseBodyOssUploadMeta {
	s.Bucket = &v
	return s
}

type GenerateOssUploadMetaResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateOssUploadMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateOssUploadMetaResponse) SetBody(v *GenerateOssUploadMetaResponseBody) *GenerateOssUploadMetaResponse {
	s.Body = v
	return s
}

type GenerateSdkDownloadInfoRequest struct {
	Sdks              *string `json:"Sdks,omitempty" xml:"Sdks,omitempty"`
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OsType            *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PkgName           *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	CertFileObjectKey *string `json:"CertFileObjectKey,omitempty" xml:"CertFileObjectKey,omitempty"`
}

func (s GenerateSdkDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateSdkDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GenerateSdkDownloadInfoRequest) SetSdks(v string) *GenerateSdkDownloadInfoRequest {
	s.Sdks = &v
	return s
}

func (s *GenerateSdkDownloadInfoRequest) SetAppId(v string) *GenerateSdkDownloadInfoRequest {
	s.AppId = &v
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

func (s *GenerateSdkDownloadInfoRequest) SetCertFileObjectKey(v string) *GenerateSdkDownloadInfoRequest {
	s.CertFileObjectKey = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateSdkDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateSdkDownloadInfoResponse) SetBody(v *GenerateSdkDownloadInfoResponseBody) *GenerateSdkDownloadInfoResponse {
	s.Body = v
	return s
}

type GenerateSysAppDownloadInfoRequest struct {
	Plugins           *string `json:"Plugins,omitempty" xml:"Plugins,omitempty"`
	SignMode          *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
	OsType            *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PkgName           *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	CertFileObjectKey *string `json:"CertFileObjectKey,omitempty" xml:"CertFileObjectKey,omitempty"`
}

func (s GenerateSysAppDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateSysAppDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GenerateSysAppDownloadInfoRequest) SetPlugins(v string) *GenerateSysAppDownloadInfoRequest {
	s.Plugins = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetSignMode(v string) *GenerateSysAppDownloadInfoRequest {
	s.SignMode = &v
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

func (s *GenerateSysAppDownloadInfoRequest) SetProjectId(v string) *GenerateSysAppDownloadInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateSysAppDownloadInfoRequest) SetCertFileObjectKey(v string) *GenerateSysAppDownloadInfoRequest {
	s.CertFileObjectKey = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateSysAppDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateSysAppDownloadInfoResponse) SetBody(v *GenerateSysAppDownloadInfoResponseBody) *GenerateSysAppDownloadInfoResponse {
	s.Body = v
	return s
}

type GetDeviceAppUpdateFunnelEventsRequest struct {
	PackageName       *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TargetVersionCode *string `json:"TargetVersionCode,omitempty" xml:"TargetVersionCode,omitempty"`
	IdType            *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OriginalId        *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
}

func (s GetDeviceAppUpdateFunnelEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsRequest) GoString() string {
	return s.String()
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

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetIdType(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.IdType = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsRequest) SetOriginalId(v string) *GetDeviceAppUpdateFunnelEventsRequest {
	s.OriginalId = &v
	return s
}

type GetDeviceAppUpdateFunnelEventsResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventList []*GetDeviceAppUpdateFunnelEventsResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
}

func (s GetDeviceAppUpdateFunnelEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBody) SetRequestId(v string) *GetDeviceAppUpdateFunnelEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBody) SetEventList(v []*GetDeviceAppUpdateFunnelEventsResponseBodyEventList) *GetDeviceAppUpdateFunnelEventsResponseBody {
	s.EventList = v
	return s
}

type GetDeviceAppUpdateFunnelEventsResponseBodyEventList struct {
	PackageName       *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	DeviceId          *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	TargetVersionCode *string `json:"TargetVersionCode,omitempty" xml:"TargetVersionCode,omitempty"`
	Event             *string `json:"Event,omitempty" xml:"Event,omitempty"`
	ReportTimestamp   *int64  `json:"ReportTimestamp,omitempty" xml:"ReportTimestamp,omitempty"`
	ReportTime        *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceAppUpdateFunnelEventsResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceAppUpdateFunnelEventsResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetPackageName(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.PackageName = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetDeviceId(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetTargetVersionCode(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.TargetVersionCode = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetEvent(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.Event = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetReportTimestamp(v int64) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.ReportTimestamp = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetReportTime(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.ReportTime = &v
	return s
}

func (s *GetDeviceAppUpdateFunnelEventsResponseBodyEventList) SetTenantId(v string) *GetDeviceAppUpdateFunnelEventsResponseBodyEventList {
	s.TenantId = &v
	return s
}

type GetDeviceAppUpdateFunnelEventsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceAppUpdateFunnelEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceAppUpdateFunnelEventsResponse) SetBody(v *GetDeviceAppUpdateFunnelEventsResponseBody) *GetDeviceAppUpdateFunnelEventsResponse {
	s.Body = v
	return s
}

type GetDeviceSystemUpdateFunnelEventsRequest struct {
	OriginalId    *string `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	IdType        *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s GetDeviceSystemUpdateFunnelEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSystemUpdateFunnelEventsRequest) GoString() string {
	return s.String()
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

func (s *GetDeviceSystemUpdateFunnelEventsRequest) SetIdType(v string) *GetDeviceSystemUpdateFunnelEventsRequest {
	s.IdType = &v
	return s
}

type GetDeviceSystemUpdateFunnelEventsResponseBody struct {
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventList []*GetDeviceSystemUpdateFunnelEventsResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
}

func (s GetDeviceSystemUpdateFunnelEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceSystemUpdateFunnelEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBody) SetRequestId(v string) *GetDeviceSystemUpdateFunnelEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBody) SetEventList(v []*GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) *GetDeviceSystemUpdateFunnelEventsResponseBody {
	s.EventList = v
	return s
}

type GetDeviceSystemUpdateFunnelEventsResponseBodyEventList struct {
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	TargetVersion   *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	Event           *string `json:"Event,omitempty" xml:"Event,omitempty"`
	ReportTimestamp *int64  `json:"ReportTimestamp,omitempty" xml:"ReportTimestamp,omitempty"`
	ReportTime      *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
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

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetTargetVersion(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.TargetVersion = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetEvent(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.Event = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetReportTimestamp(v int64) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.ReportTimestamp = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetReportTime(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.ReportTime = &v
	return s
}

func (s *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList) SetTenantId(v string) *GetDeviceSystemUpdateFunnelEventsResponseBodyEventList {
	s.TenantId = &v
	return s
}

type GetDeviceSystemUpdateFunnelEventsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceSystemUpdateFunnelEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceSystemUpdateFunnelEventsResponse) SetBody(v *GetDeviceSystemUpdateFunnelEventsResponseBody) *GetDeviceSystemUpdateFunnelEventsResponse {
	s.Body = v
	return s
}

type GetNamespaceDataRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	AuthType     *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	AccountType  *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetNamespaceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceDataRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceDataRequest) SetProjectId(v string) *GetNamespaceDataRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNamespaceDataRequest) SetNamespace(v string) *GetNamespaceDataRequest {
	s.Namespace = &v
	return s
}

func (s *GetNamespaceDataRequest) SetAuthType(v string) *GetNamespaceDataRequest {
	s.AuthType = &v
	return s
}

func (s *GetNamespaceDataRequest) SetDeviceIdType(v string) *GetNamespaceDataRequest {
	s.DeviceIdType = &v
	return s
}

func (s *GetNamespaceDataRequest) SetDeviceId(v string) *GetNamespaceDataRequest {
	s.DeviceId = &v
	return s
}

func (s *GetNamespaceDataRequest) SetAccountType(v string) *GetNamespaceDataRequest {
	s.AccountType = &v
	return s
}

func (s *GetNamespaceDataRequest) SetAccountId(v string) *GetNamespaceDataRequest {
	s.AccountId = &v
	return s
}

type GetNamespaceDataResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetNamespaceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceDataResponseBody) SetRequestId(v string) *GetNamespaceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNamespaceDataResponseBody) SetData(v string) *GetNamespaceDataResponseBody {
	s.Data = &v
	return s
}

type GetNamespaceDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNamespaceDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetNamespaceDataResponse) SetBody(v *GetNamespaceDataResponseBody) *GetNamespaceDataResponse {
	s.Body = v
	return s
}

type GetOssUploadMetaRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
}

func (s GetOssUploadMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadMetaRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadMetaRequest) SetProjectId(v string) *GetOssUploadMetaRequest {
	s.ProjectId = &v
	return s
}

func (s *GetOssUploadMetaRequest) SetExt(v string) *GetOssUploadMetaRequest {
	s.Ext = &v
	return s
}

type GetOssUploadMetaResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OssUploadMeta *GetOssUploadMetaResponseBodyOssUploadMeta `json:"OssUploadMeta,omitempty" xml:"OssUploadMeta,omitempty" type:"Struct"`
}

func (s GetOssUploadMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadMetaResponseBody) SetRequestId(v string) *GetOssUploadMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadMetaResponseBody) SetOssUploadMeta(v *GetOssUploadMetaResponseBodyOssUploadMeta) *GetOssUploadMetaResponseBody {
	s.OssUploadMeta = v
	return s
}

type GetOssUploadMetaResponseBodyOssUploadMeta struct {
	AccessKey     *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ObjectKey     *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
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

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetSignature(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.Signature = &v
	return s
}

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetHost(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.Host = &v
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

func (s *GetOssUploadMetaResponseBodyOssUploadMeta) SetObjectKey(v string) *GetOssUploadMetaResponseBodyOssUploadMeta {
	s.ObjectKey = &v
	return s
}

type GetOssUploadMetaResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssUploadMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOssUploadMetaResponse) SetBody(v *GetOssUploadMetaResponseBody) *GetOssUploadMetaResponse {
	s.Body = v
	return s
}

type InvokeFunctionRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileId       *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Env          *int32  `json:"Env,omitempty" xml:"Env,omitempty"`
	Parameters   *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s InvokeFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionRequest) GoString() string {
	return s.String()
}

func (s *InvokeFunctionRequest) SetProjectId(v string) *InvokeFunctionRequest {
	s.ProjectId = &v
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

func (s *InvokeFunctionRequest) SetEnv(v int32) *InvokeFunctionRequest {
	s.Env = &v
	return s
}

func (s *InvokeFunctionRequest) SetParameters(v string) *InvokeFunctionRequest {
	s.Parameters = &v
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
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty"`
	BackEndRequestId *string `json:"BackEndRequestId,omitempty" xml:"BackEndRequestId,omitempty"`
}

func (s InvokeFunctionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponseBodyResult) SetOutput(v string) *InvokeFunctionResponseBodyResult {
	s.Output = &v
	return s
}

func (s *InvokeFunctionResponseBodyResult) SetBackEndRequestId(v string) *InvokeFunctionResponseBodyResult {
	s.BackEndRequestId = &v
	return s
}

type InvokeFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InvokeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GatewayAppKey    *string `json:"GatewayAppKey,omitempty" xml:"GatewayAppKey,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	GatewayAppSecret *string `json:"GatewayAppSecret,omitempty" xml:"GatewayAppSecret,omitempty"`
	GatewayAppId     *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListApiGatewayAppsResponseBodyApiGatewayApps) String() string {
	return tea.Prettify(s)
}

func (s ListApiGatewayAppsResponseBodyApiGatewayApps) GoString() string {
	return s.String()
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGatewayAppKey(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GatewayAppKey = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetStatus(v int32) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.Status = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGatewayAppSecret(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GatewayAppSecret = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetGatewayAppId(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.GatewayAppId = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetUserId(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.UserId = &v
	return s
}

func (s *ListApiGatewayAppsResponseBodyApiGatewayApps) SetProjectId(v string) *ListApiGatewayAppsResponseBodyApiGatewayApps {
	s.ProjectId = &v
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

type ListApiGatewayAppsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApiGatewayAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Apps      []*ListAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetApps(v []*ListAppsResponseBodyApps) *ListAppsResponseBody {
	s.Apps = v
	return s
}

type ListAppsResponseBodyApps struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	OsType     *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	AppPackage *string `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
}

func (s ListAppsResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyApps) SetAppName(v string) *ListAppsResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppKey(v string) *ListAppsResponseBodyApps {
	s.AppKey = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetOsType(v int32) *ListAppsResponseBodyApps {
	s.OsType = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppPackage(v string) *ListAppsResponseBodyApps {
	s.AppPackage = &v
	return s
}

type ListAppsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ID        *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s ListAssistActionDetailsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListAssistActionDetailsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListAssistActionDetailsResponseBodyResults) SetType(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Type = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetAction(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Action = &v
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

func (s *ListAssistActionDetailsResponseBodyResults) SetCreatedAt(v int64) *ListAssistActionDetailsResponseBodyResults {
	s.CreatedAt = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetUpdatedAt(v int64) *ListAssistActionDetailsResponseBodyResults {
	s.UpdatedAt = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetTimestamp(v string) *ListAssistActionDetailsResponseBodyResults {
	s.Timestamp = &v
	return s
}

func (s *ListAssistActionDetailsResponseBodyResults) SetID(v string) *ListAssistActionDetailsResponseBodyResults {
	s.ID = &v
	return s
}

type ListAssistActionDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAssistActionDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAssistActionDetailsResponse) SetBody(v *ListAssistActionDetailsResponseBody) *ListAssistActionDetailsResponse {
	s.Body = v
	return s
}

type ListAssistDevicesRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	PerPage   *int32  `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
}

func (s ListAssistDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesRequest) SetProjectId(v string) *ListAssistDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListAssistDevicesRequest) SetPageIndex(v int32) *ListAssistDevicesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAssistDevicesRequest) SetCondition(v string) *ListAssistDevicesRequest {
	s.Condition = &v
	return s
}

func (s *ListAssistDevicesRequest) SetPerPage(v int32) *ListAssistDevicesRequest {
	s.PerPage = &v
	return s
}

type ListAssistDevicesResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PerPage    *int32                                  `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	PageIndex  *int32                                  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	Devices    []*ListAssistDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
}

func (s ListAssistDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesResponseBody) SetTotalCount(v int32) *ListAssistDevicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetRequestId(v string) *ListAssistDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetPerPage(v int32) *ListAssistDevicesResponseBody {
	s.PerPage = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetPageIndex(v int32) *ListAssistDevicesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListAssistDevicesResponseBody) SetDevices(v []*ListAssistDevicesResponseBodyDevices) *ListAssistDevicesResponseBody {
	s.Devices = v
	return s
}

type ListAssistDevicesResponseBodyDevices struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	HardwareId   *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	AccessTime   *int64  `json:"AccessTime,omitempty" xml:"AccessTime,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UUID         *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VIN          *string `json:"VIN,omitempty" xml:"VIN,omitempty"`
}

func (s ListAssistDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListAssistDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListAssistDevicesResponseBodyDevices) SetSerialNumber(v string) *ListAssistDevicesResponseBodyDevices {
	s.SerialNumber = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetHardwareId(v string) *ListAssistDevicesResponseBodyDevices {
	s.HardwareId = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetDeviceName(v string) *ListAssistDevicesResponseBodyDevices {
	s.DeviceName = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetAccessTime(v int64) *ListAssistDevicesResponseBodyDevices {
	s.AccessTime = &v
	return s
}

func (s *ListAssistDevicesResponseBodyDevices) SetDeviceId(v string) *ListAssistDevicesResponseBodyDevices {
	s.DeviceId = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAssistDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAssistDevicesResponse) SetBody(v *ListAssistDevicesResponseBody) *ListAssistDevicesResponse {
	s.Body = v
	return s
}

type ListAssistHistoriesRequest struct {
	PerPage   *int32  `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAssistHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesRequest) SetPerPage(v int32) *ListAssistHistoriesRequest {
	s.PerPage = &v
	return s
}

func (s *ListAssistHistoriesRequest) SetPageIndex(v int32) *ListAssistHistoriesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAssistHistoriesRequest) SetCondition(v string) *ListAssistHistoriesRequest {
	s.Condition = &v
	return s
}

func (s *ListAssistHistoriesRequest) SetProjectId(v string) *ListAssistHistoriesRequest {
	s.ProjectId = &v
	return s
}

type ListAssistHistoriesResponseBody struct {
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PerPage    *int32                                      `json:"PerPage,omitempty" xml:"PerPage,omitempty"`
	Histories  []*ListAssistHistoriesResponseBodyHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Repeated"`
	PageIndex  *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s ListAssistHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesResponseBody) SetTotalCount(v int32) *ListAssistHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetRequestId(v string) *ListAssistHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetPerPage(v int32) *ListAssistHistoriesResponseBody {
	s.PerPage = &v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetHistories(v []*ListAssistHistoriesResponseBodyHistories) *ListAssistHistoriesResponseBody {
	s.Histories = v
	return s
}

func (s *ListAssistHistoriesResponseBody) SetPageIndex(v int32) *ListAssistHistoriesResponseBody {
	s.PageIndex = &v
	return s
}

type ListAssistHistoriesResponseBodyHistories struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UNAME        *string `json:"UNAME,omitempty" xml:"UNAME,omitempty"`
	HardwareId   *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	UUID         *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	VIN          *string `json:"VIN,omitempty" xml:"VIN,omitempty"`
	UID          *string `json:"UID,omitempty" xml:"UID,omitempty"`
	ID           *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s ListAssistHistoriesResponseBodyHistories) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoriesResponseBodyHistories) GoString() string {
	return s.String()
}

func (s *ListAssistHistoriesResponseBodyHistories) SetSerialNumber(v string) *ListAssistHistoriesResponseBodyHistories {
	s.SerialNumber = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetEndTime(v int64) *ListAssistHistoriesResponseBodyHistories {
	s.EndTime = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetStartTime(v int64) *ListAssistHistoriesResponseBodyHistories {
	s.StartTime = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetUNAME(v string) *ListAssistHistoriesResponseBodyHistories {
	s.UNAME = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetHardwareId(v string) *ListAssistHistoriesResponseBodyHistories {
	s.HardwareId = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetDeviceName(v string) *ListAssistHistoriesResponseBodyHistories {
	s.DeviceName = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetUUID(v string) *ListAssistHistoriesResponseBodyHistories {
	s.UUID = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetDeviceId(v string) *ListAssistHistoriesResponseBodyHistories {
	s.DeviceId = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetVIN(v string) *ListAssistHistoriesResponseBodyHistories {
	s.VIN = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetUID(v string) *ListAssistHistoriesResponseBodyHistories {
	s.UID = &v
	return s
}

func (s *ListAssistHistoriesResponseBodyHistories) SetID(v string) *ListAssistHistoriesResponseBodyHistories {
	s.ID = &v
	return s
}

type ListAssistHistoriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAssistHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAssistHistoriesResponse) SetBody(v *ListAssistHistoriesResponseBody) *ListAssistHistoriesResponse {
	s.Body = v
	return s
}

type ListAssistHistoryDetailsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AssistId  *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
}

func (s ListAssistHistoryDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoryDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListAssistHistoryDetailsRequest) SetProjectId(v string) *ListAssistHistoryDetailsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListAssistHistoryDetailsRequest) SetAssistId(v string) *ListAssistHistoryDetailsRequest {
	s.AssistId = &v
	return s
}

type ListAssistHistoryDetailsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Actions   []*ListAssistHistoryDetailsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
}

func (s ListAssistHistoryDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistHistoryDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistHistoryDetailsResponseBody) SetRequestId(v string) *ListAssistHistoryDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBody) SetActions(v []*ListAssistHistoryDetailsResponseBodyActions) *ListAssistHistoryDetailsResponseBody {
	s.Actions = v
	return s
}

type ListAssistHistoryDetailsResponseBodyActions struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	AssistId  *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	ID        *string `json:"ID,omitempty" xml:"ID,omitempty"`
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

func (s *ListAssistHistoryDetailsResponseBodyActions) SetCreatedAt(v int64) *ListAssistHistoryDetailsResponseBodyActions {
	s.CreatedAt = &v
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

func (s *ListAssistHistoryDetailsResponseBodyActions) SetAssistId(v string) *ListAssistHistoryDetailsResponseBodyActions {
	s.AssistId = &v
	return s
}

func (s *ListAssistHistoryDetailsResponseBodyActions) SetID(v string) *ListAssistHistoryDetailsResponseBodyActions {
	s.ID = &v
	return s
}

type ListAssistHistoryDetailsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAssistHistoryDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAssistHistoryDetailsResponse) SetBody(v *ListAssistHistoryDetailsResponseBody) *ListAssistHistoryDetailsResponse {
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
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClientPlugins []*ListClientPluginsResponseBodyClientPlugins `json:"ClientPlugins,omitempty" xml:"ClientPlugins,omitempty" type:"Repeated"`
}

func (s ListClientPluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientPluginsResponseBody) SetRequestId(v string) *ListClientPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientPluginsResponseBody) SetClientPlugins(v []*ListClientPluginsResponseBodyClientPlugins) *ListClientPluginsResponseBody {
	s.ClientPlugins = v
	return s
}

type ListClientPluginsResponseBodyClientPlugins struct {
	PkgName *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListClientPluginsResponseBodyClientPlugins) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginsResponseBodyClientPlugins) GoString() string {
	return s.String()
}

func (s *ListClientPluginsResponseBodyClientPlugins) SetPkgName(v string) *ListClientPluginsResponseBodyClientPlugins {
	s.PkgName = &v
	return s
}

func (s *ListClientPluginsResponseBodyClientPlugins) SetName(v string) *ListClientPluginsResponseBodyClientPlugins {
	s.Name = &v
	return s
}

type ListClientPluginsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClientPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListClientPluginsResponse) SetBody(v *ListClientPluginsResponseBody) *ListClientPluginsResponse {
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
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	PkgName     *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	VersionCode *int64  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListClientPluginVersionsResponseBodyClientPluginVersions) String() string {
	return tea.Prettify(s)
}

func (s ListClientPluginVersionsResponseBodyClientPluginVersions) GoString() string {
	return s.String()
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetVersion(v string) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.Version = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetDownloadUrl(v string) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.DownloadUrl = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetSize(v int64) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.Size = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetPkgName(v string) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.PkgName = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetVersionCode(v int64) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.VersionCode = &v
	return s
}

func (s *ListClientPluginVersionsResponseBodyClientPluginVersions) SetId(v int64) *ListClientPluginVersionsResponseBodyClientPluginVersions {
	s.Id = &v
	return s
}

type ListClientPluginVersionsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClientPluginVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListClientPluginVersionsResponse) SetBody(v *ListClientPluginVersionsResponseBody) *ListClientPluginVersionsResponse {
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
	OsType      *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PkgName     *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	PkgType     *int32  `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListClientSdksResponseBodyClientSdks) String() string {
	return tea.Prettify(s)
}

func (s ListClientSdksResponseBodyClientSdks) GoString() string {
	return s.String()
}

func (s *ListClientSdksResponseBodyClientSdks) SetOsType(v int32) *ListClientSdksResponseBodyClientSdks {
	s.OsType = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetGmtCreate(v int64) *ListClientSdksResponseBodyClientSdks {
	s.GmtCreate = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetGmtModified(v int64) *ListClientSdksResponseBodyClientSdks {
	s.GmtModified = &v
	return s
}

func (s *ListClientSdksResponseBodyClientSdks) SetName(v string) *ListClientSdksResponseBodyClientSdks {
	s.Name = &v
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

func (s *ListClientSdksResponseBodyClientSdks) SetId(v int64) *ListClientSdksResponseBodyClientSdks {
	s.Id = &v
	return s
}

type ListClientSdksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClientSdksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListClientSdksResponse) SetBody(v *ListClientSdksResponseBody) *ListClientSdksResponse {
	s.Body = v
	return s
}

type ListConnectLogsRequest struct {
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s ListConnectLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectLogsRequest) SetPageSize(v int32) *ListConnectLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectLogsRequest) SetProjectId(v string) *ListConnectLogsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListConnectLogsRequest) SetDeviceId(v string) *ListConnectLogsRequest {
	s.DeviceId = &v
	return s
}

func (s *ListConnectLogsRequest) SetStartTime(v int32) *ListConnectLogsRequest {
	s.StartTime = &v
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

type ListConnectLogsResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Logs      *ListConnectLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
}

func (s ListConnectLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBody) SetRequestId(v string) *ListConnectLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectLogsResponseBody) SetLogs(v *ListConnectLogsResponseBodyLogs) *ListConnectLogsResponseBody {
	s.Logs = v
	return s
}

type ListConnectLogsResponseBodyLogs struct {
	Pagination *ListConnectLogsResponseBodyLogsPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListConnectLogsResponseBodyLogsList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListConnectLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBodyLogs) SetPagination(v *ListConnectLogsResponseBodyLogsPagination) *ListConnectLogsResponseBodyLogs {
	s.Pagination = v
	return s
}

func (s *ListConnectLogsResponseBodyLogs) SetList(v []*ListConnectLogsResponseBodyLogsList) *ListConnectLogsResponseBodyLogs {
	s.List = v
	return s
}

type ListConnectLogsResponseBodyLogsPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListConnectLogsResponseBodyLogsPagination) SetTotalPageCount(v int32) *ListConnectLogsResponseBodyLogsPagination {
	s.TotalPageCount = &v
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

type ListConnectLogsResponseBodyLogsList struct {
	Sid           *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Time          *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NetWorking    *string `json:"NetWorking,omitempty" xml:"NetWorking,omitempty"`
	Terminal      *string `json:"Terminal,omitempty" xml:"Terminal,omitempty"`
}

func (s ListConnectLogsResponseBodyLogsList) String() string {
	return tea.Prettify(s)
}

func (s ListConnectLogsResponseBodyLogsList) GoString() string {
	return s.String()
}

func (s *ListConnectLogsResponseBodyLogsList) SetSid(v string) *ListConnectLogsResponseBodyLogsList {
	s.Sid = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetStatus(v string) *ListConnectLogsResponseBodyLogsList {
	s.Status = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetTime(v int64) *ListConnectLogsResponseBodyLogsList {
	s.Time = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetDeviceId(v string) *ListConnectLogsResponseBodyLogsList {
	s.DeviceId = &v
	return s
}

func (s *ListConnectLogsResponseBodyLogsList) SetSystemVersion(v string) *ListConnectLogsResponseBodyLogsList {
	s.SystemVersion = &v
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

func (s *ListConnectLogsResponseBodyLogsList) SetTerminal(v string) *ListConnectLogsResponseBodyLogsList {
	s.Terminal = &v
	return s
}

type ListConnectLogsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListConnectLogsResponse) SetBody(v *ListConnectLogsResponseBody) *ListConnectLogsResponse {
	s.Body = v
	return s
}

type ListDeployedFunctionsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileId    *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s ListDeployedFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsRequest) SetProjectId(v string) *ListDeployedFunctionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeployedFunctionsRequest) SetFileId(v int64) *ListDeployedFunctionsRequest {
	s.FileId = &v
	return s
}

type ListDeployedFunctionsResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Functions []*ListDeployedFunctionsResponseBodyFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
}

func (s ListDeployedFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsResponseBody) SetRequestId(v string) *ListDeployedFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployedFunctionsResponseBody) SetFunctions(v []*ListDeployedFunctionsResponseBodyFunctions) *ListDeployedFunctionsResponseBody {
	s.Functions = v
	return s
}

type ListDeployedFunctionsResponseBodyFunctions struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	FileId      *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s ListDeployedFunctionsResponseBodyFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedFunctionsResponseBodyFunctions) GoString() string {
	return s.String()
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetProjectId(v string) *ListDeployedFunctionsResponseBodyFunctions {
	s.ProjectId = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetGmtCreate(v int64) *ListDeployedFunctionsResponseBodyFunctions {
	s.GmtCreate = &v
	return s
}

func (s *ListDeployedFunctionsResponseBodyFunctions) SetName(v string) *ListDeployedFunctionsResponseBodyFunctions {
	s.Name = &v
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

func (s *ListDeployedFunctionsResponseBodyFunctions) SetFileId(v int64) *ListDeployedFunctionsResponseBodyFunctions {
	s.FileId = &v
	return s
}

type ListDeployedFunctionsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeployedFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDeployedFunctionsResponse) SetBody(v *ListDeployedFunctionsResponseBody) *ListDeployedFunctionsResponse {
	s.Body = v
	return s
}

type ListDeviceBrandsRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s ListDeviceBrandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBrandsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBrandsRequest) SetProjectId(v string) *ListDeviceBrandsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetDeviceBrandId(v int64) *ListDeviceBrandsRequest {
	s.DeviceBrandId = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetDeviceBrand(v string) *ListDeviceBrandsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetStart(v string) *ListDeviceBrandsRequest {
	s.Start = &v
	return s
}

func (s *ListDeviceBrandsRequest) SetLength(v string) *ListDeviceBrandsRequest {
	s.Length = &v
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
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Manufacture   *string `json:"Manufacture,omitempty" xml:"Manufacture,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s ListDeviceBrandsResponseBodyDeviceBrands) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceBrandsResponseBodyDeviceBrands) GoString() string {
	return s.String()
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetDeviceBrandId(v int64) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.DeviceBrandId = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetDescription(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.Description = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetProjectId(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetManufacture(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.Manufacture = &v
	return s
}

func (s *ListDeviceBrandsResponseBodyDeviceBrands) SetDeviceBrand(v string) *ListDeviceBrandsResponseBodyDeviceBrands {
	s.DeviceBrand = &v
	return s
}

type ListDeviceBrandsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDeviceBrandsResponse) SetBody(v *ListDeviceBrandsResponseBody) *ListDeviceBrandsResponse {
	s.Body = v
	return s
}

type ListDeviceModelRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceModelRequest) SetProjectId(v string) *ListDeviceModelRequest {
	s.ProjectId = &v
	return s
}

type ListDeviceModelResponseBody struct {
	ModelList []*ListDeviceModelResponseBodyModelList `json:"ModelList,omitempty" xml:"ModelList,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceModelResponseBody) SetModelList(v []*ListDeviceModelResponseBodyModelList) *ListDeviceModelResponseBody {
	s.ModelList = v
	return s
}

func (s *ListDeviceModelResponseBody) SetRequestId(v string) *ListDeviceModelResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceModelResponseBodyModelList struct {
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s ListDeviceModelResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *ListDeviceModelResponseBodyModelList) SetDeviceModelId(v int64) *ListDeviceModelResponseBodyModelList {
	s.DeviceModelId = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetHardwareType(v string) *ListDeviceModelResponseBodyModelList {
	s.HardwareType = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetDeviceType(v string) *ListDeviceModelResponseBodyModelList {
	s.DeviceType = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetCanCreateDeviceId(v int32) *ListDeviceModelResponseBodyModelList {
	s.CanCreateDeviceId = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetProjectId(v string) *ListDeviceModelResponseBodyModelList {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetOsPlatform(v string) *ListDeviceModelResponseBodyModelList {
	s.OsPlatform = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetDeviceModel(v string) *ListDeviceModelResponseBodyModelList {
	s.DeviceModel = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetSecurityChip(v string) *ListDeviceModelResponseBodyModelList {
	s.SecurityChip = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetDescription(v string) *ListDeviceModelResponseBodyModelList {
	s.Description = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetInitUsageTypeDesc(v string) *ListDeviceModelResponseBodyModelList {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetInitUsageType(v int32) *ListDeviceModelResponseBodyModelList {
	s.InitUsageType = &v
	return s
}

func (s *ListDeviceModelResponseBodyModelList) SetDeviceBrand(v string) *ListDeviceModelResponseBodyModelList {
	s.DeviceBrand = &v
	return s
}

type ListDeviceModelResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceModelResponse) SetHeaders(v map[string]*string) *ListDeviceModelResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceModelResponse) SetBody(v *ListDeviceModelResponseBody) *ListDeviceModelResponse {
	s.Body = v
	return s
}

type ListDeviceModelsRequest struct {
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceBrandId *int64  `json:"DeviceBrandId,omitempty" xml:"DeviceBrandId,omitempty"`
}

func (s ListDeviceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceModelsRequest) SetDeviceModelId(v int32) *ListDeviceModelsRequest {
	s.DeviceModelId = &v
	return s
}

func (s *ListDeviceModelsRequest) SetLength(v string) *ListDeviceModelsRequest {
	s.Length = &v
	return s
}

func (s *ListDeviceModelsRequest) SetDeviceModel(v string) *ListDeviceModelsRequest {
	s.DeviceModel = &v
	return s
}

func (s *ListDeviceModelsRequest) SetDeviceBrand(v string) *ListDeviceModelsRequest {
	s.DeviceBrand = &v
	return s
}

func (s *ListDeviceModelsRequest) SetStart(v string) *ListDeviceModelsRequest {
	s.Start = &v
	return s
}

func (s *ListDeviceModelsRequest) SetProjectId(v string) *ListDeviceModelsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceModelsRequest) SetDeviceBrandId(v int64) *ListDeviceModelsRequest {
	s.DeviceBrandId = &v
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
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
	DeviceLogoUrl     *string `json:"DeviceLogoUrl,omitempty" xml:"DeviceLogoUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s ListDeviceModelsResponseBodyDeviceModels) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceModelsResponseBodyDeviceModels) GoString() string {
	return s.String()
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceModelId(v int64) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceModelId = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetHardwareType(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.HardwareType = &v
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

func (s *ListDeviceModelsResponseBodyDeviceModels) SetCanCreateDeviceId(v int32) *ListDeviceModelsResponseBodyDeviceModels {
	s.CanCreateDeviceId = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetProjectId(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.ProjectId = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetOsPlatform(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.OsPlatform = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceModel(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceModel = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetSecurityChip(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.SecurityChip = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceLogoUrl(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceLogoUrl = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDescription(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.Description = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetObjectKey(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.ObjectKey = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetInitUsageTypeDesc(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetInitUsageType(v int32) *ListDeviceModelsResponseBodyDeviceModels {
	s.InitUsageType = &v
	return s
}

func (s *ListDeviceModelsResponseBodyDeviceModels) SetDeviceBrand(v string) *ListDeviceModelsResponseBodyDeviceModels {
	s.DeviceBrand = &v
	return s
}

type ListDeviceModelsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDeviceModelsResponse) SetBody(v *ListDeviceModelsResponseBody) *ListDeviceModelsResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModelId *int32  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Length        *string `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetProjectId(v string) *ListDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceModelId(v int32) *ListDevicesRequest {
	s.DeviceModelId = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceModel(v string) *ListDevicesRequest {
	s.DeviceModel = &v
	return s
}

func (s *ListDevicesRequest) SetStart(v string) *ListDevicesRequest {
	s.Start = &v
	return s
}

func (s *ListDevicesRequest) SetLength(v string) *ListDevicesRequest {
	s.Length = &v
	return s
}

type ListDevicesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Devices   []*ListDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetDevices(v []*ListDevicesResponseBodyDevices) *ListDevicesResponseBody {
	s.Devices = v
	return s
}

type ListDevicesResponseBodyDevices struct {
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	MacAddress    *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	UsageType     *int32  `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	Vin           *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
	UsageTypeDesc *string `json:"UsageTypeDesc,omitempty" xml:"UsageTypeDesc,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	HardwareId    *string `json:"HardwareId,omitempty" xml:"HardwareId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SoftwareId    *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s ListDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDevices) SetSerialNumber(v string) *ListDevicesResponseBodyDevices {
	s.SerialNumber = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetStatus(v string) *ListDevicesResponseBodyDevices {
	s.Status = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceModelId(v int64) *ListDevicesResponseBodyDevices {
	s.DeviceModelId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetMacAddress(v string) *ListDevicesResponseBodyDevices {
	s.MacAddress = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceId(v string) *ListDevicesResponseBodyDevices {
	s.DeviceId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceType(v string) *ListDevicesResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetProjectId(v string) *ListDevicesResponseBodyDevices {
	s.ProjectId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceModel(v string) *ListDevicesResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetUsageType(v int32) *ListDevicesResponseBodyDevices {
	s.UsageType = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetVin(v string) *ListDevicesResponseBodyDevices {
	s.Vin = &v
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

func (s *ListDevicesResponseBodyDevices) SetHardwareId(v string) *ListDevicesResponseBodyDevices {
	s.HardwareId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetRegion(v string) *ListDevicesResponseBodyDevices {
	s.Region = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSoftwareId(v string) *ListDevicesResponseBodyDevices {
	s.SoftwareId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetName(v string) *ListDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceBrand(v string) *ListDevicesResponseBodyDevices {
	s.DeviceBrand = &v
	return s
}

type ListDevicesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
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
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceTypes []*ListDeviceTypesResponseBodyDeviceTypes `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
}

func (s ListDeviceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceTypesResponseBody) SetRequestId(v string) *ListDeviceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceTypesResponseBody) SetDeviceTypes(v []*ListDeviceTypesResponseBodyDeviceTypes) *ListDeviceTypesResponseBody {
	s.DeviceTypes = v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDeviceTypesResponse) SetBody(v *ListDeviceTypesResponseBody) *ListDeviceTypesResponse {
	s.Body = v
	return s
}

type ListFunctionExecuteLogRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileId       *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Env          *int32  `json:"Env,omitempty" xml:"Env,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFunctionExecuteLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogRequest) SetProjectId(v string) *ListFunctionExecuteLogRequest {
	s.ProjectId = &v
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

func (s *ListFunctionExecuteLogRequest) SetEnv(v int32) *ListFunctionExecuteLogRequest {
	s.Env = &v
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
	Pagination *ListFunctionExecuteLogResponseBodyLogListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Logs       []*ListFunctionExecuteLogResponseBodyLogListLogs     `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s ListFunctionExecuteLogResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBodyLogList) SetPagination(v *ListFunctionExecuteLogResponseBodyLogListPagination) *ListFunctionExecuteLogResponseBodyLogList {
	s.Pagination = v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogList) SetLogs(v []*ListFunctionExecuteLogResponseBodyLogListLogs) *ListFunctionExecuteLogResponseBodyLogList {
	s.Logs = v
	return s
}

type ListFunctionExecuteLogResponseBodyLogListPagination struct {
	PageIndex   *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HasNextPage *bool  `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
}

func (s ListFunctionExecuteLogResponseBodyLogListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBodyLogListPagination) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBodyLogListPagination) SetPageIndex(v int32) *ListFunctionExecuteLogResponseBodyLogListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogListPagination) SetPageSize(v int32) *ListFunctionExecuteLogResponseBodyLogListPagination {
	s.PageSize = &v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogListPagination) SetHasNextPage(v bool) *ListFunctionExecuteLogResponseBodyLogListPagination {
	s.HasNextPage = &v
	return s
}

type ListFunctionExecuteLogResponseBodyLogListLogs struct {
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	BackEndRequestId *string `json:"BackEndRequestId,omitempty" xml:"BackEndRequestId,omitempty"`
}

func (s ListFunctionExecuteLogResponseBodyLogListLogs) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionExecuteLogResponseBodyLogListLogs) GoString() string {
	return s.String()
}

func (s *ListFunctionExecuteLogResponseBodyLogListLogs) SetMessage(v string) *ListFunctionExecuteLogResponseBodyLogListLogs {
	s.Message = &v
	return s
}

func (s *ListFunctionExecuteLogResponseBodyLogListLogs) SetBackEndRequestId(v string) *ListFunctionExecuteLogResponseBodyLogListLogs {
	s.BackEndRequestId = &v
	return s
}

type ListFunctionExecuteLogResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionExecuteLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListFunctionExecuteLogResponse) SetBody(v *ListFunctionExecuteLogResponseBody) *ListFunctionExecuteLogResponse {
	s.Body = v
	return s
}

type ListFunctionFilesRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFunctionFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesRequest) SetProjectId(v string) *ListFunctionFilesRequest {
	s.ProjectId = &v
	return s
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

type ListFunctionFilesResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FileList  *ListFunctionFilesResponseBodyFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Struct"`
}

func (s ListFunctionFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBody) SetRequestId(v string) *ListFunctionFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionFilesResponseBody) SetFileList(v *ListFunctionFilesResponseBodyFileList) *ListFunctionFilesResponseBody {
	s.FileList = v
	return s
}

type ListFunctionFilesResponseBodyFileList struct {
	Pagination *ListFunctionFilesResponseBodyFileListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Files      []*ListFunctionFilesResponseBodyFileListFiles    `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s ListFunctionFilesResponseBodyFileList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBodyFileList) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBodyFileList) SetPagination(v *ListFunctionFilesResponseBodyFileListPagination) *ListFunctionFilesResponseBodyFileList {
	s.Pagination = v
	return s
}

func (s *ListFunctionFilesResponseBodyFileList) SetFiles(v []*ListFunctionFilesResponseBodyFileListFiles) *ListFunctionFilesResponseBodyFileList {
	s.Files = v
	return s
}

type ListFunctionFilesResponseBodyFileListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListFunctionFilesResponseBodyFileListPagination) SetTotalPageCount(v int32) *ListFunctionFilesResponseBodyFileListPagination {
	s.TotalPageCount = &v
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

type ListFunctionFilesResponseBodyFileListFiles struct {
	Status                 *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ProductionDeployTime   *int64  `json:"ProductionDeployTime,omitempty" xml:"ProductionDeployTime,omitempty"`
	ProductionDeployStatus *int32  `json:"ProductionDeployStatus,omitempty" xml:"ProductionDeployStatus,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SandboxDeployTime      *int64  `json:"SandboxDeployTime,omitempty" xml:"SandboxDeployTime,omitempty"`
	GmtCreate              *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SandboxDeployStatus    *int32  `json:"SandboxDeployStatus,omitempty" xml:"SandboxDeployStatus,omitempty"`
	GmtModified            *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ContentId              *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	Id                     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFunctionFilesResponseBodyFileListFiles) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesResponseBodyFileListFiles) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetStatus(v int32) *ListFunctionFilesResponseBodyFileListFiles {
	s.Status = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetProductionDeployTime(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.ProductionDeployTime = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetProductionDeployStatus(v int32) *ListFunctionFilesResponseBodyFileListFiles {
	s.ProductionDeployStatus = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetDescription(v string) *ListFunctionFilesResponseBodyFileListFiles {
	s.Description = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetSandboxDeployTime(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.SandboxDeployTime = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetGmtCreate(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.GmtCreate = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetSandboxDeployStatus(v int32) *ListFunctionFilesResponseBodyFileListFiles {
	s.SandboxDeployStatus = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetGmtModified(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.GmtModified = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetName(v string) *ListFunctionFilesResponseBodyFileListFiles {
	s.Name = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetContentId(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.ContentId = &v
	return s
}

func (s *ListFunctionFilesResponseBodyFileListFiles) SetId(v int64) *ListFunctionFilesResponseBodyFileListFiles {
	s.Id = &v
	return s
}

type ListFunctionFilesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListFunctionFilesResponse) SetBody(v *ListFunctionFilesResponseBody) *ListFunctionFilesResponse {
	s.Body = v
	return s
}

type ListFunctionFilesByProjectIdRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFunctionFilesByProjectIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdRequest) SetProjectId(v string) *ListFunctionFilesByProjectIdRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionFilesByProjectIdRequest) SetPageIndex(v int32) *ListFunctionFilesByProjectIdRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFunctionFilesByProjectIdRequest) SetPageSize(v int32) *ListFunctionFilesByProjectIdRequest {
	s.PageSize = &v
	return s
}

type ListFunctionFilesByProjectIdResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Files     []*ListFunctionFilesByProjectIdResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s ListFunctionFilesByProjectIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdResponseBody) SetRequestId(v string) *ListFunctionFilesByProjectIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBody) SetFiles(v []*ListFunctionFilesByProjectIdResponseBodyFiles) *ListFunctionFilesByProjectIdResponseBody {
	s.Files = v
	return s
}

type ListFunctionFilesByProjectIdResponseBodyFiles struct {
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ContentId   *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
}

func (s ListFunctionFilesByProjectIdResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionFilesByProjectIdResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetType(v int32) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Type = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetStatus(v int32) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Status = &v
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

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetName(v string) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetId(v int64) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.Id = &v
	return s
}

func (s *ListFunctionFilesByProjectIdResponseBodyFiles) SetContentId(v int64) *ListFunctionFilesByProjectIdResponseBodyFiles {
	s.ContentId = &v
	return s
}

type ListFunctionFilesByProjectIdResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionFilesByProjectIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListFunctionFilesByProjectIdResponse) SetBody(v *ListFunctionFilesByProjectIdResponseBody) *ListFunctionFilesByProjectIdResponse {
	s.Body = v
	return s
}

type ListMessageAcksRequest struct {
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	MessageId *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s ListMessageAcksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksRequest) GoString() string {
	return s.String()
}

func (s *ListMessageAcksRequest) SetPageSize(v int32) *ListMessageAcksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageAcksRequest) SetProjectId(v string) *ListMessageAcksRequest {
	s.ProjectId = &v
	return s
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

type ListMessageAcksResponseBody struct {
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MessageAcks *ListMessageAcksResponseBodyMessageAcks `json:"MessageAcks,omitempty" xml:"MessageAcks,omitempty" type:"Struct"`
}

func (s ListMessageAcksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBody) SetRequestId(v string) *ListMessageAcksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageAcksResponseBody) SetMessageAcks(v *ListMessageAcksResponseBodyMessageAcks) *ListMessageAcksResponseBody {
	s.MessageAcks = v
	return s
}

type ListMessageAcksResponseBodyMessageAcks struct {
	Pagination *ListMessageAcksResponseBodyMessageAcksPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListMessageAcksResponseBodyMessageAcksList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListMessageAcksResponseBodyMessageAcks) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBodyMessageAcks) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBodyMessageAcks) SetPagination(v *ListMessageAcksResponseBodyMessageAcksPagination) *ListMessageAcksResponseBodyMessageAcks {
	s.Pagination = v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcks) SetList(v []*ListMessageAcksResponseBodyMessageAcksList) *ListMessageAcksResponseBodyMessageAcks {
	s.List = v
	return s
}

type ListMessageAcksResponseBodyMessageAcksPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListMessageAcksResponseBodyMessageAcksPagination) SetTotalPageCount(v int32) *ListMessageAcksResponseBodyMessageAcksPagination {
	s.TotalPageCount = &v
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

type ListMessageAcksResponseBodyMessageAcksList struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	AckTime  *int64  `json:"AckTime,omitempty" xml:"AckTime,omitempty"`
	Mid      *int64  `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s ListMessageAcksResponseBodyMessageAcksList) String() string {
	return tea.Prettify(s)
}

func (s ListMessageAcksResponseBodyMessageAcksList) GoString() string {
	return s.String()
}

func (s *ListMessageAcksResponseBodyMessageAcksList) SetDeviceId(v string) *ListMessageAcksResponseBodyMessageAcksList {
	s.DeviceId = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksList) SetAckTime(v int64) *ListMessageAcksResponseBodyMessageAcksList {
	s.AckTime = &v
	return s
}

func (s *ListMessageAcksResponseBodyMessageAcksList) SetMid(v int64) *ListMessageAcksResponseBodyMessageAcksList {
	s.Mid = &v
	return s
}

type ListMessageAcksResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMessageAcksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListMessageAcksResponse) SetBody(v *ListMessageAcksResponseBody) *ListMessageAcksResponse {
	s.Body = v
	return s
}

type ListMessageReceiversRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMessageReceiversRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversRequest) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversRequest) SetProjectId(v string) *ListMessageReceiversRequest {
	s.ProjectId = &v
	return s
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
	Pagination *ListMessageReceiversResponseBodyMessageReceiversPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListMessageReceiversResponseBodyMessageReceiversList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListMessageReceiversResponseBodyMessageReceivers) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponseBodyMessageReceivers) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponseBodyMessageReceivers) SetPagination(v *ListMessageReceiversResponseBodyMessageReceiversPagination) *ListMessageReceiversResponseBodyMessageReceivers {
	s.Pagination = v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceivers) SetList(v []*ListMessageReceiversResponseBodyMessageReceiversList) *ListMessageReceiversResponseBodyMessageReceivers {
	s.List = v
	return s
}

type ListMessageReceiversResponseBodyMessageReceiversPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListMessageReceiversResponseBodyMessageReceiversPagination) SetTotalPageCount(v int32) *ListMessageReceiversResponseBodyMessageReceiversPagination {
	s.TotalPageCount = &v
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

type ListMessageReceiversResponseBodyMessageReceiversList struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Mid   *int64  `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s ListMessageReceiversResponseBodyMessageReceiversList) String() string {
	return tea.Prettify(s)
}

func (s ListMessageReceiversResponseBodyMessageReceiversList) GoString() string {
	return s.String()
}

func (s *ListMessageReceiversResponseBodyMessageReceiversList) SetType(v string) *ListMessageReceiversResponseBodyMessageReceiversList {
	s.Type = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversList) SetValue(v string) *ListMessageReceiversResponseBodyMessageReceiversList {
	s.Value = &v
	return s
}

func (s *ListMessageReceiversResponseBodyMessageReceiversList) SetMid(v int64) *ListMessageReceiversResponseBodyMessageReceiversList {
	s.Mid = &v
	return s
}

type ListMessageReceiversResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMessageReceiversResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListMessageReceiversResponse) SetBody(v *ListMessageReceiversResponseBody) *ListMessageReceiversResponse {
	s.Body = v
	return s
}

type ListMqttClientSubscriptionsRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ClientId  *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListMqttClientSubscriptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqttClientSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListMqttClientSubscriptionsRequest) SetAppKey(v string) *ListMqttClientSubscriptionsRequest {
	s.AppKey = &v
	return s
}

func (s *ListMqttClientSubscriptionsRequest) SetClientId(v string) *ListMqttClientSubscriptionsRequest {
	s.ClientId = &v
	return s
}

func (s *ListMqttClientSubscriptionsRequest) SetProjectId(v string) *ListMqttClientSubscriptionsRequest {
	s.ProjectId = &v
	return s
}

type ListMqttClientSubscriptionsResponseBody struct {
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClientSubscriptions *ListMqttClientSubscriptionsResponseBodyClientSubscriptions `json:"ClientSubscriptions,omitempty" xml:"ClientSubscriptions,omitempty" type:"Struct"`
}

func (s ListMqttClientSubscriptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqttClientSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqttClientSubscriptionsResponseBody) SetRequestId(v string) *ListMqttClientSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqttClientSubscriptionsResponseBody) SetClientSubscriptions(v *ListMqttClientSubscriptionsResponseBodyClientSubscriptions) *ListMqttClientSubscriptionsResponseBody {
	s.ClientSubscriptions = v
	return s
}

type ListMqttClientSubscriptionsResponseBodyClientSubscriptions struct {
	Pagination *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListMqttClientSubscriptionsResponseBodyClientSubscriptions) String() string {
	return tea.Prettify(s)
}

func (s ListMqttClientSubscriptionsResponseBodyClientSubscriptions) GoString() string {
	return s.String()
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptions) SetPagination(v *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) *ListMqttClientSubscriptionsResponseBodyClientSubscriptions {
	s.Pagination = v
	return s
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptions) SetList(v []*ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList) *ListMqttClientSubscriptionsResponseBodyClientSubscriptions {
	s.List = v
	return s
}

type ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) String() string {
	return tea.Prettify(s)
}

func (s ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) GoString() string {
	return s.String()
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) SetPageIndex(v int32) *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination {
	s.PageIndex = &v
	return s
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) SetTotalPageCount(v int32) *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination {
	s.TotalPageCount = &v
	return s
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) SetPageSize(v int32) *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination {
	s.PageSize = &v
	return s
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination) SetTotalCount(v int32) *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsPagination {
	s.TotalCount = &v
	return s
}

type ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList struct {
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	QoS   *int32  `json:"QoS,omitempty" xml:"QoS,omitempty"`
}

func (s ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList) String() string {
	return tea.Prettify(s)
}

func (s ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList) GoString() string {
	return s.String()
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList) SetTopic(v string) *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList {
	s.Topic = &v
	return s
}

func (s *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList) SetQoS(v int32) *ListMqttClientSubscriptionsResponseBodyClientSubscriptionsList {
	s.QoS = &v
	return s
}

type ListMqttClientSubscriptionsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMqttClientSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqttClientSubscriptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqttClientSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListMqttClientSubscriptionsResponse) SetHeaders(v map[string]*string) *ListMqttClientSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListMqttClientSubscriptionsResponse) SetBody(v *ListMqttClientSubscriptionsResponseBody) *ListMqttClientSubscriptionsResponse {
	s.Body = v
	return s
}

type ListMqttMessageLogsRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Mid       *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	ClientId  *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMqttMessageLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqttMessageLogsRequest) GoString() string {
	return s.String()
}

func (s *ListMqttMessageLogsRequest) SetAppKey(v string) *ListMqttMessageLogsRequest {
	s.AppKey = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetProjectId(v string) *ListMqttMessageLogsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetTopic(v string) *ListMqttMessageLogsRequest {
	s.Topic = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetMid(v string) *ListMqttMessageLogsRequest {
	s.Mid = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetClientId(v string) *ListMqttMessageLogsRequest {
	s.ClientId = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetStartTime(v int32) *ListMqttMessageLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetEndTime(v int32) *ListMqttMessageLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMqttMessageLogsRequest) SetType(v string) *ListMqttMessageLogsRequest {
	s.Type = &v
	return s
}

type ListMqttMessageLogsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Traces    *ListMqttMessageLogsResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Struct"`
}

func (s ListMqttMessageLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqttMessageLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqttMessageLogsResponseBody) SetRequestId(v string) *ListMqttMessageLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqttMessageLogsResponseBody) SetTraces(v *ListMqttMessageLogsResponseBodyTraces) *ListMqttMessageLogsResponseBody {
	s.Traces = v
	return s
}

type ListMqttMessageLogsResponseBodyTraces struct {
	Pagination *ListMqttMessageLogsResponseBodyTracesPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListMqttMessageLogsResponseBodyTracesList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListMqttMessageLogsResponseBodyTraces) String() string {
	return tea.Prettify(s)
}

func (s ListMqttMessageLogsResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *ListMqttMessageLogsResponseBodyTraces) SetPagination(v *ListMqttMessageLogsResponseBodyTracesPagination) *ListMqttMessageLogsResponseBodyTraces {
	s.Pagination = v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTraces) SetList(v []*ListMqttMessageLogsResponseBodyTracesList) *ListMqttMessageLogsResponseBodyTraces {
	s.List = v
	return s
}

type ListMqttMessageLogsResponseBodyTracesPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMqttMessageLogsResponseBodyTracesPagination) String() string {
	return tea.Prettify(s)
}

func (s ListMqttMessageLogsResponseBodyTracesPagination) GoString() string {
	return s.String()
}

func (s *ListMqttMessageLogsResponseBodyTracesPagination) SetPageIndex(v int32) *ListMqttMessageLogsResponseBodyTracesPagination {
	s.PageIndex = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesPagination) SetTotalPageCount(v int32) *ListMqttMessageLogsResponseBodyTracesPagination {
	s.TotalPageCount = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesPagination) SetPageSize(v int32) *ListMqttMessageLogsResponseBodyTracesPagination {
	s.PageSize = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesPagination) SetTotalCount(v int32) *ListMqttMessageLogsResponseBodyTracesPagination {
	s.TotalCount = &v
	return s
}

type ListMqttMessageLogsResponseBodyTracesList struct {
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Time      *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Mid       *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	ClientMid *string `json:"ClientMid,omitempty" xml:"ClientMid,omitempty"`
	ClientId  *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s ListMqttMessageLogsResponseBodyTracesList) String() string {
	return tea.Prettify(s)
}

func (s ListMqttMessageLogsResponseBodyTracesList) GoString() string {
	return s.String()
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetType(v string) *ListMqttMessageLogsResponseBodyTracesList {
	s.Type = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetTime(v int64) *ListMqttMessageLogsResponseBodyTracesList {
	s.Time = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetAction(v string) *ListMqttMessageLogsResponseBodyTracesList {
	s.Action = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetTopic(v string) *ListMqttMessageLogsResponseBodyTracesList {
	s.Topic = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetMid(v string) *ListMqttMessageLogsResponseBodyTracesList {
	s.Mid = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetClientMid(v string) *ListMqttMessageLogsResponseBodyTracesList {
	s.ClientMid = &v
	return s
}

func (s *ListMqttMessageLogsResponseBodyTracesList) SetClientId(v string) *ListMqttMessageLogsResponseBodyTracesList {
	s.ClientId = &v
	return s
}

type ListMqttMessageLogsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMqttMessageLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqttMessageLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqttMessageLogsResponse) GoString() string {
	return s.String()
}

func (s *ListMqttMessageLogsResponse) SetHeaders(v map[string]*string) *ListMqttMessageLogsResponse {
	s.Headers = v
	return s
}

func (s *ListMqttMessageLogsResponse) SetBody(v *ListMqttMessageLogsResponseBody) *ListMqttMessageLogsResponse {
	s.Body = v
	return s
}

type ListMqttRootTopicsRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListMqttRootTopicsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqttRootTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListMqttRootTopicsRequest) SetAppKey(v string) *ListMqttRootTopicsRequest {
	s.AppKey = &v
	return s
}

func (s *ListMqttRootTopicsRequest) SetProjectId(v string) *ListMqttRootTopicsRequest {
	s.ProjectId = &v
	return s
}

type ListMqttRootTopicsResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootTopics *ListMqttRootTopicsResponseBodyRootTopics `json:"RootTopics,omitempty" xml:"RootTopics,omitempty" type:"Struct"`
}

func (s ListMqttRootTopicsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqttRootTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqttRootTopicsResponseBody) SetRequestId(v string) *ListMqttRootTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqttRootTopicsResponseBody) SetRootTopics(v *ListMqttRootTopicsResponseBodyRootTopics) *ListMqttRootTopicsResponseBody {
	s.RootTopics = v
	return s
}

type ListMqttRootTopicsResponseBodyRootTopics struct {
	Pagination *ListMqttRootTopicsResponseBodyRootTopicsPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListMqttRootTopicsResponseBodyRootTopicsList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListMqttRootTopicsResponseBodyRootTopics) String() string {
	return tea.Prettify(s)
}

func (s ListMqttRootTopicsResponseBodyRootTopics) GoString() string {
	return s.String()
}

func (s *ListMqttRootTopicsResponseBodyRootTopics) SetPagination(v *ListMqttRootTopicsResponseBodyRootTopicsPagination) *ListMqttRootTopicsResponseBodyRootTopics {
	s.Pagination = v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopics) SetList(v []*ListMqttRootTopicsResponseBodyRootTopicsList) *ListMqttRootTopicsResponseBodyRootTopics {
	s.List = v
	return s
}

type ListMqttRootTopicsResponseBodyRootTopicsPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMqttRootTopicsResponseBodyRootTopicsPagination) String() string {
	return tea.Prettify(s)
}

func (s ListMqttRootTopicsResponseBodyRootTopicsPagination) GoString() string {
	return s.String()
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsPagination) SetPageIndex(v int32) *ListMqttRootTopicsResponseBodyRootTopicsPagination {
	s.PageIndex = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsPagination) SetTotalPageCount(v int32) *ListMqttRootTopicsResponseBodyRootTopicsPagination {
	s.TotalPageCount = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsPagination) SetPageSize(v int32) *ListMqttRootTopicsResponseBodyRootTopicsPagination {
	s.PageSize = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsPagination) SetTotalCount(v int32) *ListMqttRootTopicsResponseBodyRootTopicsPagination {
	s.TotalCount = &v
	return s
}

type ListMqttRootTopicsResponseBodyRootTopicsList struct {
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	QueueName  *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	CreateTime *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RootTopic  *string `json:"RootTopic,omitempty" xml:"RootTopic,omitempty"`
}

func (s ListMqttRootTopicsResponseBodyRootTopicsList) String() string {
	return tea.Prettify(s)
}

func (s ListMqttRootTopicsResponseBodyRootTopicsList) GoString() string {
	return s.String()
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsList) SetType(v string) *ListMqttRootTopicsResponseBodyRootTopicsList {
	s.Type = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsList) SetAppKey(v string) *ListMqttRootTopicsResponseBodyRootTopicsList {
	s.AppKey = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsList) SetQueueName(v string) *ListMqttRootTopicsResponseBodyRootTopicsList {
	s.QueueName = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsList) SetCreateTime(v int32) *ListMqttRootTopicsResponseBodyRootTopicsList {
	s.CreateTime = &v
	return s
}

func (s *ListMqttRootTopicsResponseBodyRootTopicsList) SetRootTopic(v string) *ListMqttRootTopicsResponseBodyRootTopicsList {
	s.RootTopic = &v
	return s
}

type ListMqttRootTopicsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMqttRootTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqttRootTopicsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqttRootTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListMqttRootTopicsResponse) SetHeaders(v map[string]*string) *ListMqttRootTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListMqttRootTopicsResponse) SetBody(v *ListMqttRootTopicsResponseBody) *ListMqttRootTopicsResponse {
	s.Body = v
	return s
}

type ListNamespacesRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) SetProjectId(v string) *ListNamespacesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNamespacesRequest) SetAuthType(v string) *ListNamespacesRequest {
	s.AuthType = &v
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
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *ListNamespacesResponseBodyNamespaces) SetUserId(v string) *ListNamespacesResponseBodyNamespaces {
	s.UserId = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetProjectId(v string) *ListNamespacesResponseBodyNamespaces {
	s.ProjectId = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetGmtCreate(v int64) *ListNamespacesResponseBodyNamespaces {
	s.GmtCreate = &v
	return s
}

func (s *ListNamespacesResponseBodyNamespaces) SetNamespace(v string) *ListNamespacesResponseBodyNamespaces {
	s.Namespace = &v
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

type ListNamespacesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

type ListOfflineMessagesRequest struct {
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s ListOfflineMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesRequest) GoString() string {
	return s.String()
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

func (s *ListOfflineMessagesRequest) SetPageIndex(v int32) *ListOfflineMessagesRequest {
	s.PageIndex = &v
	return s
}

type ListOfflineMessagesResponseBody struct {
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OfflineMessages *ListOfflineMessagesResponseBodyOfflineMessages `json:"OfflineMessages,omitempty" xml:"OfflineMessages,omitempty" type:"Struct"`
}

func (s ListOfflineMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponseBody) SetRequestId(v string) *ListOfflineMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfflineMessagesResponseBody) SetOfflineMessages(v *ListOfflineMessagesResponseBodyOfflineMessages) *ListOfflineMessagesResponseBody {
	s.OfflineMessages = v
	return s
}

type ListOfflineMessagesResponseBodyOfflineMessages struct {
	Pagination *ListOfflineMessagesResponseBodyOfflineMessagesPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListOfflineMessagesResponseBodyOfflineMessagesList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListOfflineMessagesResponseBodyOfflineMessages) String() string {
	return tea.Prettify(s)
}

func (s ListOfflineMessagesResponseBodyOfflineMessages) GoString() string {
	return s.String()
}

func (s *ListOfflineMessagesResponseBodyOfflineMessages) SetPagination(v *ListOfflineMessagesResponseBodyOfflineMessagesPagination) *ListOfflineMessagesResponseBodyOfflineMessages {
	s.Pagination = v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessages) SetList(v []*ListOfflineMessagesResponseBodyOfflineMessagesList) *ListOfflineMessagesResponseBodyOfflineMessages {
	s.List = v
	return s
}

type ListOfflineMessagesResponseBodyOfflineMessagesPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListOfflineMessagesResponseBodyOfflineMessagesPagination) SetTotalPageCount(v int32) *ListOfflineMessagesResponseBodyOfflineMessagesPagination {
	s.TotalPageCount = &v
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

type ListOfflineMessagesResponseBodyOfflineMessagesList struct {
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Mid         *int64 `json:"Mid,omitempty" xml:"Mid,omitempty"`
	GmtCreate   *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
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

func (s *ListOfflineMessagesResponseBodyOfflineMessagesList) SetMid(v int64) *ListOfflineMessagesResponseBodyOfflineMessagesList {
	s.Mid = &v
	return s
}

func (s *ListOfflineMessagesResponseBodyOfflineMessagesList) SetGmtCreate(v int64) *ListOfflineMessagesResponseBodyOfflineMessagesList {
	s.GmtCreate = &v
	return s
}

type ListOfflineMessagesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOfflineMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListOfflineMessagesResponse) SetBody(v *ListOfflineMessagesResponseBody) *ListOfflineMessagesResponse {
	s.Body = v
	return s
}

type ListOpenAccountLinksRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	Idp        *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	OpenId     *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
}

func (s ListOpenAccountLinksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksRequest) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksRequest) SetProjectId(v string) *ListOpenAccountLinksRequest {
	s.ProjectId = &v
	return s
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

type ListOpenAccountLinksResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OpenAccounts []*ListOpenAccountLinksResponseBodyOpenAccounts `json:"OpenAccounts,omitempty" xml:"OpenAccounts,omitempty" type:"Repeated"`
}

func (s ListOpenAccountLinksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksResponseBody) SetRequestId(v string) *ListOpenAccountLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBody) SetOpenAccounts(v []*ListOpenAccountLinksResponseBodyOpenAccounts) *ListOpenAccountLinksResponseBody {
	s.OpenAccounts = v
	return s
}

type ListOpenAccountLinksResponseBodyOpenAccounts struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	CreateAccessKey *string `json:"CreateAccessKey,omitempty" xml:"CreateAccessKey,omitempty"`
	OpenId          *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IdentityId      *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	LoginId         *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	Idp             *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s ListOpenAccountLinksResponseBodyOpenAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountLinksResponseBodyOpenAccounts) GoString() string {
	return s.String()
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetStatus(v int32) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Status = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetType(v int32) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Type = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetDisplayName(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetCreateAccessKey(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.CreateAccessKey = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetOpenId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.OpenId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetMobile(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Mobile = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetRegion(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Region = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetIdentityId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.IdentityId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetLoginId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.LoginId = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetIdp(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.Idp = &v
	return s
}

func (s *ListOpenAccountLinksResponseBodyOpenAccounts) SetAliyunId(v string) *ListOpenAccountLinksResponseBodyOpenAccounts {
	s.AliyunId = &v
	return s
}

type ListOpenAccountLinksResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOpenAccountLinksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListOpenAccountLinksResponse) SetBody(v *ListOpenAccountLinksResponseBody) *ListOpenAccountLinksResponse {
	s.Body = v
	return s
}

type ListOpenAccountsRequest struct {
	Length      *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Start       *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListOpenAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsRequest) SetLength(v int32) *ListOpenAccountsRequest {
	s.Length = &v
	return s
}

func (s *ListOpenAccountsRequest) SetStart(v int32) *ListOpenAccountsRequest {
	s.Start = &v
	return s
}

func (s *ListOpenAccountsRequest) SetMobile(v string) *ListOpenAccountsRequest {
	s.Mobile = &v
	return s
}

func (s *ListOpenAccountsRequest) SetEmail(v string) *ListOpenAccountsRequest {
	s.Email = &v
	return s
}

func (s *ListOpenAccountsRequest) SetDisplayName(v string) *ListOpenAccountsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListOpenAccountsRequest) SetProjectId(v string) *ListOpenAccountsRequest {
	s.ProjectId = &v
	return s
}

type ListOpenAccountsResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OpenAccounts []*ListOpenAccountsResponseBodyOpenAccounts `json:"OpenAccounts,omitempty" xml:"OpenAccounts,omitempty" type:"Repeated"`
}

func (s ListOpenAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsResponseBody) SetRequestId(v string) *ListOpenAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenAccountsResponseBody) SetOpenAccounts(v []*ListOpenAccountsResponseBodyOpenAccounts) *ListOpenAccountsResponseBody {
	s.OpenAccounts = v
	return s
}

type ListOpenAccountsResponseBodyOpenAccounts struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	CreateAccessKey *string `json:"CreateAccessKey,omitempty" xml:"CreateAccessKey,omitempty"`
	OpenId          *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IdentityId      *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	LoginId         *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	Idp             *string `json:"Idp,omitempty" xml:"Idp,omitempty"`
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s ListOpenAccountsResponseBodyOpenAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListOpenAccountsResponseBodyOpenAccounts) GoString() string {
	return s.String()
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetStatus(v int32) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Status = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetType(v int32) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Type = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetDisplayName(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetCreateAccessKey(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.CreateAccessKey = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetOpenId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.OpenId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetMobile(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Mobile = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetRegion(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Region = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetIdentityId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.IdentityId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetLoginId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.LoginId = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetIdp(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.Idp = &v
	return s
}

func (s *ListOpenAccountsResponseBodyOpenAccounts) SetAliyunId(v string) *ListOpenAccountsResponseBodyOpenAccounts {
	s.AliyunId = &v
	return s
}

type ListOpenAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOpenAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListOpenAccountsResponse) SetBody(v *ListOpenAccountsResponseBody) *ListOpenAccountsResponse {
	s.Body = v
	return s
}

type ListPreChecksResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreChecks []*ListPreChecksResponseBodyPreChecks `json:"PreChecks,omitempty" xml:"PreChecks,omitempty" type:"Repeated"`
}

func (s ListPreChecksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPreChecksResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreChecksResponseBody) SetRequestId(v string) *ListPreChecksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPreChecksResponseBody) SetPreChecks(v []*ListPreChecksResponseBodyPreChecks) *ListPreChecksResponseBody {
	s.PreChecks = v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPreChecksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListPreChecksResponse) SetBody(v *ListPreChecksResponseBody) *ListPreChecksResponse {
	s.Body = v
	return s
}

type ListProjectAppsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keywords  *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
}

func (s ListProjectAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectAppsRequest) SetProjectId(v string) *ListProjectAppsRequest {
	s.ProjectId = &v
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

func (s *ListProjectAppsRequest) SetKeywords(v string) *ListProjectAppsRequest {
	s.Keywords = &v
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
	TotalPage   *int32                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalCount  *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListProjectAppsResponseBodyResult) SetTotalPage(v int32) *ListProjectAppsResponseBodyResult {
	s.TotalPage = &v
	return s
}

func (s *ListProjectAppsResponseBodyResult) SetTotalCount(v int32) *ListProjectAppsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListProjectAppsResponseBodyResultProjectApps struct {
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	AppPkgName  *string `json:"AppPkgName,omitempty" xml:"AppPkgName,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppSecret   *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OsType      *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListProjectAppsResponseBodyResultProjectApps) String() string {
	return tea.Prettify(s)
}

func (s ListProjectAppsResponseBodyResultProjectApps) GoString() string {
	return s.String()
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetStatus(v int32) *ListProjectAppsResponseBodyResultProjectApps {
	s.Status = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetProjectId(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.ProjectId = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetUserId(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.UserId = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetGmtModified(v int64) *ListProjectAppsResponseBodyResultProjectApps {
	s.GmtModified = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppPkgName(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppPkgName = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppName(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppName = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppSecret(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppSecret = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppKey(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppKey = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetAppId(v string) *ListProjectAppsResponseBodyResultProjectApps {
	s.AppId = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetOsType(v int32) *ListProjectAppsResponseBodyResultProjectApps {
	s.OsType = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetGmtCreate(v int64) *ListProjectAppsResponseBodyResultProjectApps {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectAppsResponseBodyResultProjectApps) SetId(v int64) *ListProjectAppsResponseBodyResultProjectApps {
	s.Id = &v
	return s
}

type ListProjectAppsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProjectAppsResponse) SetBody(v *ListProjectAppsResponseBody) *ListProjectAppsResponse {
	s.Body = v
	return s
}

type ListProjectsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Projects  []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

type ListProjectsResponseBodyProjects struct {
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetStatus(v int32) *ListProjectsResponseBodyProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetUserId(v string) *ListProjectsResponseBodyProjects {
	s.UserId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectId(v string) *ListProjectsResponseBodyProjects {
	s.ProjectId = &v
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

func (s *ListProjectsResponseBodyProjects) SetName(v string) *ListProjectsResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetId(v int64) *ListProjectsResponseBodyProjects {
	s.Id = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreator(v string) *ListProjectsResponseBodyProjects {
	s.Creator = &v
	return s
}

type ListProjectsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListRpcServicesRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRpcServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesRequest) GoString() string {
	return s.String()
}

func (s *ListRpcServicesRequest) SetProjectId(v string) *ListRpcServicesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListRpcServicesRequest) SetPageIndex(v int32) *ListRpcServicesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListRpcServicesRequest) SetPageSize(v int32) *ListRpcServicesRequest {
	s.PageSize = &v
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
	Pagination *ListRpcServicesResponseBodyRpcServicesPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListRpcServicesResponseBodyRpcServicesList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListRpcServicesResponseBodyRpcServices) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponseBodyRpcServices) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponseBodyRpcServices) SetPagination(v *ListRpcServicesResponseBodyRpcServicesPagination) *ListRpcServicesResponseBodyRpcServices {
	s.Pagination = v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServices) SetList(v []*ListRpcServicesResponseBodyRpcServicesList) *ListRpcServicesResponseBodyRpcServices {
	s.List = v
	return s
}

type ListRpcServicesResponseBodyRpcServicesPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListRpcServicesResponseBodyRpcServicesPagination) SetTotalPageCount(v int32) *ListRpcServicesResponseBodyRpcServicesPagination {
	s.TotalPageCount = &v
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

type ListRpcServicesResponseBodyRpcServicesList struct {
	MethodName    *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	Params        *string `json:"Params,omitempty" xml:"Params,omitempty"`
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IsDelete      *string `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	VersionCode   *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRpcServicesResponseBodyRpcServicesList) String() string {
	return tea.Prettify(s)
}

func (s ListRpcServicesResponseBodyRpcServicesList) GoString() string {
	return s.String()
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetMethodName(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.MethodName = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetType(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.Type = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetInterfaceName(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.InterfaceName = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetParams(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.Params = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetAppKey(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.AppKey = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetGroupName(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.GroupName = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetGmtCreate(v int64) *ListRpcServicesResponseBodyRpcServicesList {
	s.GmtCreate = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetIsDelete(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.IsDelete = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetVersionCode(v string) *ListRpcServicesResponseBodyRpcServicesList {
	s.VersionCode = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetGmtModified(v int64) *ListRpcServicesResponseBodyRpcServicesList {
	s.GmtModified = &v
	return s
}

func (s *ListRpcServicesResponseBodyRpcServicesList) SetId(v int64) *ListRpcServicesResponseBodyRpcServicesList {
	s.Id = &v
	return s
}

type ListRpcServicesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRpcServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRpcServicesResponse) SetBody(v *ListRpcServicesResponseBody) *ListRpcServicesResponse {
	s.Body = v
	return s
}

type ListSchemaSubscribesRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSchemaSubscribesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesRequest) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesRequest) SetProjectId(v string) *ListSchemaSubscribesRequest {
	s.ProjectId = &v
	return s
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

type ListSchemaSubscribesResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageList  []*ListSchemaSubscribesResponseBodyPageList `json:"PageList,omitempty" xml:"PageList,omitempty" type:"Repeated"`
}

func (s ListSchemaSubscribesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBody) SetRequestId(v string) *ListSchemaSubscribesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemaSubscribesResponseBody) SetPageList(v []*ListSchemaSubscribesResponseBodyPageList) *ListSchemaSubscribesResponseBody {
	s.PageList = v
	return s
}

type ListSchemaSubscribesResponseBodyPageList struct {
	Pagination *ListSchemaSubscribesResponseBodyPageListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListSchemaSubscribesResponseBodyPageListList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListSchemaSubscribesResponseBodyPageList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBodyPageList) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBodyPageList) SetPagination(v *ListSchemaSubscribesResponseBodyPageListPagination) *ListSchemaSubscribesResponseBodyPageList {
	s.Pagination = v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageList) SetList(v []*ListSchemaSubscribesResponseBodyPageListList) *ListSchemaSubscribesResponseBodyPageList {
	s.List = v
	return s
}

type ListSchemaSubscribesResponseBodyPageListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SimpleSign     *bool  `json:"SimpleSign,omitempty" xml:"SimpleSign,omitempty"`
	HasNextPage    *bool  `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
}

func (s ListSchemaSubscribesResponseBodyPageListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBodyPageListPagination) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetPageIndex(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetTotalPageCount(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.TotalPageCount = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetPageSize(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.PageSize = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetTotalCount(v int32) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetSimpleSign(v bool) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.SimpleSign = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListPagination) SetHasNextPage(v bool) *ListSchemaSubscribesResponseBodyPageListPagination {
	s.HasNextPage = &v
	return s
}

type ListSchemaSubscribesResponseBodyPageListList struct {
	DeviceModelId  *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ValiditySchema *string `json:"ValiditySchema,omitempty" xml:"ValiditySchema,omitempty"`
	DeviceModel    *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListSchemaSubscribesResponseBodyPageListList) String() string {
	return tea.Prettify(s)
}

func (s ListSchemaSubscribesResponseBodyPageListList) GoString() string {
	return s.String()
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetDeviceModelId(v int64) *ListSchemaSubscribesResponseBodyPageListList {
	s.DeviceModelId = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetVersion(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.Version = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetProjectId(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.ProjectId = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetGmtCreate(v int64) *ListSchemaSubscribesResponseBodyPageListList {
	s.GmtCreate = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetNamespace(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.Namespace = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetValiditySchema(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.ValiditySchema = &v
	return s
}

func (s *ListSchemaSubscribesResponseBodyPageListList) SetDeviceModel(v string) *ListSchemaSubscribesResponseBodyPageListList {
	s.DeviceModel = &v
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

type ListSchemaSubscribesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSchemaSubscribesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeviceModelId     *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	CanCreateDeviceId *int32  `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SecurityChip      *int32  `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InitUsageTypeDesc *string `json:"InitUsageTypeDesc,omitempty" xml:"InitUsageTypeDesc,omitempty"`
	InitUsageType     *int32  `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	DeviceBrand       *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
}

func (s ListShadowSchemaDeviceModelsResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemaDeviceModelsResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceModelId(v int64) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceModelId = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetHardwareType(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.HardwareType = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceType(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceType = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetCanCreateDeviceId(v int32) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.CanCreateDeviceId = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetProjectId(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.ProjectId = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetOsPlatform(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.OsPlatform = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceModel(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceModel = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetSecurityChip(v int32) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.SecurityChip = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDescription(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.Description = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetInitUsageTypeDesc(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.InitUsageTypeDesc = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetInitUsageType(v int32) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.InitUsageType = &v
	return s
}

func (s *ListShadowSchemaDeviceModelsResponseBodyModelList) SetDeviceBrand(v string) *ListShadowSchemaDeviceModelsResponseBodyModelList {
	s.DeviceBrand = &v
	return s
}

type ListShadowSchemaDeviceModelsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListShadowSchemaDeviceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListShadowSchemaDeviceModelsResponse) SetBody(v *ListShadowSchemaDeviceModelsResponseBody) *ListShadowSchemaDeviceModelsResponse {
	s.Body = v
	return s
}

type ListShadowSchemasRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	QueryValue *string `json:"QueryValue,omitempty" xml:"QueryValue,omitempty"`
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListShadowSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasRequest) GoString() string {
	return s.String()
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

func (s *ListShadowSchemasRequest) SetPageIndex(v int32) *ListShadowSchemasRequest {
	s.PageIndex = &v
	return s
}

func (s *ListShadowSchemasRequest) SetPageSize(v int32) *ListShadowSchemasRequest {
	s.PageSize = &v
	return s
}

type ListShadowSchemasResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageList  *ListShadowSchemasResponseBodyPageList `json:"PageList,omitempty" xml:"PageList,omitempty" type:"Struct"`
}

func (s ListShadowSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBody) SetRequestId(v string) *ListShadowSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListShadowSchemasResponseBody) SetPageList(v *ListShadowSchemasResponseBodyPageList) *ListShadowSchemasResponseBody {
	s.PageList = v
	return s
}

type ListShadowSchemasResponseBodyPageList struct {
	Pagination *ListShadowSchemasResponseBodyPageListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListShadowSchemasResponseBodyPageListList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListShadowSchemasResponseBodyPageList) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBodyPageList) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBodyPageList) SetPagination(v *ListShadowSchemasResponseBodyPageListPagination) *ListShadowSchemasResponseBodyPageList {
	s.Pagination = v
	return s
}

func (s *ListShadowSchemasResponseBodyPageList) SetList(v []*ListShadowSchemasResponseBodyPageListList) *ListShadowSchemasResponseBodyPageList {
	s.List = v
	return s
}

type ListShadowSchemasResponseBodyPageListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SimpleSign     *bool  `json:"SimpleSign,omitempty" xml:"SimpleSign,omitempty"`
	HasNextPage    *bool  `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
}

func (s ListShadowSchemasResponseBodyPageListPagination) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBodyPageListPagination) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetPageIndex(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.PageIndex = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetTotalPageCount(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.TotalPageCount = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetPageSize(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.PageSize = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetTotalCount(v int32) *ListShadowSchemasResponseBodyPageListPagination {
	s.TotalCount = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetSimpleSign(v bool) *ListShadowSchemasResponseBodyPageListPagination {
	s.SimpleSign = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListPagination) SetHasNextPage(v bool) *ListShadowSchemasResponseBodyPageListPagination {
	s.HasNextPage = &v
	return s
}

type ListShadowSchemasResponseBodyPageListList struct {
	AuthTypeDesc  *string `json:"AuthTypeDesc,omitempty" xml:"AuthTypeDesc,omitempty"`
	DeviceModelId *int64  `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	AuthType      *int32  `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ModuleSchema  *string `json:"ModuleSchema,omitempty" xml:"ModuleSchema,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListShadowSchemasResponseBodyPageListList) String() string {
	return tea.Prettify(s)
}

func (s ListShadowSchemasResponseBodyPageListList) GoString() string {
	return s.String()
}

func (s *ListShadowSchemasResponseBodyPageListList) SetAuthTypeDesc(v string) *ListShadowSchemasResponseBodyPageListList {
	s.AuthTypeDesc = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetDeviceModelId(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.DeviceModelId = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetAuthType(v int32) *ListShadowSchemasResponseBodyPageListList {
	s.AuthType = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetProjectId(v string) *ListShadowSchemasResponseBodyPageListList {
	s.ProjectId = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetGmtCreate(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.GmtCreate = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetNamespace(v string) *ListShadowSchemasResponseBodyPageListList {
	s.Namespace = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetDeviceModel(v string) *ListShadowSchemasResponseBodyPageListList {
	s.DeviceModel = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetGmtModified(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.GmtModified = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetModuleSchema(v string) *ListShadowSchemasResponseBodyPageListList {
	s.ModuleSchema = &v
	return s
}

func (s *ListShadowSchemasResponseBodyPageListList) SetId(v int64) *ListShadowSchemasResponseBodyPageListList {
	s.Id = &v
	return s
}

type ListShadowSchemasResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListShadowSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSupportFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSupportFeaturesResponse) SetBody(v *ListSupportFeaturesResponseBody) *ListSupportFeaturesResponse {
	s.Body = v
	return s
}

type ListTriggersRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) SetProjectId(v string) *ListTriggersRequest {
	s.ProjectId = &v
	return s
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
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListTriggersResponseBodyTriggerListPagination) SetTotalPageCount(v int32) *ListTriggersResponseBodyTriggerListPagination {
	s.TotalPageCount = &v
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

type ListTriggersResponseBodyTriggerListTriggers struct {
	Status             *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *int32                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	Production         *int32                                                  `json:"Production,omitempty" xml:"Production,omitempty"`
	Functions          []*ListTriggersResponseBodyTriggerListTriggersFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	Sandbox            *int32                                                  `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Namespace          *string                                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	GmtModified        *int64                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Source             *string                                                 `json:"Source,omitempty" xml:"Source,omitempty"`
	ChainedFunctionIds *string                                                 `json:"ChainedFunctionIds,omitempty" xml:"ChainedFunctionIds,omitempty"`
	GmtCreate          *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InvocationMode     *int32                                                  `json:"InvocationMode,omitempty" xml:"InvocationMode,omitempty"`
	Id                 *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTriggersResponseBodyTriggerListTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggerListTriggers) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetStatus(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Status = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetType(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Type = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetProduction(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Production = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetFunctions(v []*ListTriggersResponseBodyTriggerListTriggersFunctions) *ListTriggersResponseBodyTriggerListTriggers {
	s.Functions = v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetSandbox(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.Sandbox = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetNamespace(v string) *ListTriggersResponseBodyTriggerListTriggers {
	s.Namespace = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetGmtModified(v int64) *ListTriggersResponseBodyTriggerListTriggers {
	s.GmtModified = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetSource(v string) *ListTriggersResponseBodyTriggerListTriggers {
	s.Source = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetChainedFunctionIds(v string) *ListTriggersResponseBodyTriggerListTriggers {
	s.ChainedFunctionIds = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetGmtCreate(v int64) *ListTriggersResponseBodyTriggerListTriggers {
	s.GmtCreate = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetInvocationMode(v int32) *ListTriggersResponseBodyTriggerListTriggers {
	s.InvocationMode = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggers) SetId(v int64) *ListTriggersResponseBodyTriggerListTriggers {
	s.Id = &v
	return s
}

type ListTriggersResponseBodyTriggerListTriggersFunctions struct {
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	FileId      *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s ListTriggersResponseBodyTriggerListTriggersFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggerListTriggersFunctions) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetGmtCreate(v int64) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.GmtCreate = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetFileName(v string) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.FileName = &v
	return s
}

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetName(v string) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.Name = &v
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

func (s *ListTriggersResponseBodyTriggerListTriggersFunctions) SetFileId(v int64) *ListTriggersResponseBodyTriggerListTriggersFunctions {
	s.FileId = &v
	return s
}

type ListTriggersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTriggersResponse) SetBody(v *ListTriggersResponseBody) *ListTriggersResponse {
	s.Body = v
	return s
}

type ListUpstreamAppKeyRelationsRequest struct {
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppServerId *int64  `json:"AppServerId,omitempty" xml:"AppServerId,omitempty"`
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s ListUpstreamAppKeyRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsRequest) SetPageSize(v int32) *ListUpstreamAppKeyRelationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsRequest) SetProjectId(v string) *ListUpstreamAppKeyRelationsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsRequest) SetAppServerId(v int64) *ListUpstreamAppKeyRelationsRequest {
	s.AppServerId = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsRequest) SetPageIndex(v int32) *ListUpstreamAppKeyRelationsRequest {
	s.PageIndex = &v
	return s
}

type ListUpstreamAppKeyRelationsResponseBody struct {
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RelationList *ListUpstreamAppKeyRelationsResponseBodyRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Struct"`
}

func (s ListUpstreamAppKeyRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBody) SetRequestId(v string) *ListUpstreamAppKeyRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBody) SetRelationList(v *ListUpstreamAppKeyRelationsResponseBodyRelationList) *ListUpstreamAppKeyRelationsResponseBody {
	s.RelationList = v
	return s
}

type ListUpstreamAppKeyRelationsResponseBodyRelationList struct {
	Pagination *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListUpstreamAppKeyRelationsResponseBodyRelationListList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationList) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationList) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationList) SetPagination(v *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) *ListUpstreamAppKeyRelationsResponseBodyRelationList {
	s.Pagination = v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationList) SetList(v []*ListUpstreamAppKeyRelationsResponseBodyRelationListList) *ListUpstreamAppKeyRelationsResponseBodyRelationList {
	s.List = v
	return s
}

type ListUpstreamAppKeyRelationsResponseBodyRelationListPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination) SetTotalPageCount(v int32) *ListUpstreamAppKeyRelationsResponseBodyRelationListPagination {
	s.TotalPageCount = &v
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

type ListUpstreamAppKeyRelationsResponseBodyRelationListList struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppPackage *string `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	PAppKey    *string `json:"PAppKey,omitempty" xml:"PAppKey,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationListList) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppKeyRelationsResponseBodyRelationListList) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetAppName(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.AppName = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetAppKey(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.AppKey = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetAppPackage(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.AppPackage = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetProjectId(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetGmtCreate(v int64) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.GmtCreate = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetPAppKey(v string) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.PAppKey = &v
	return s
}

func (s *ListUpstreamAppKeyRelationsResponseBodyRelationListList) SetId(v int64) *ListUpstreamAppKeyRelationsResponseBodyRelationListList {
	s.Id = &v
	return s
}

type ListUpstreamAppKeyRelationsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUpstreamAppKeyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUpstreamAppKeyRelationsResponse) SetBody(v *ListUpstreamAppKeyRelationsResponseBody) *ListUpstreamAppKeyRelationsResponse {
	s.Body = v
	return s
}

type ListUpstreamAppServersRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageIndex *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUpstreamAppServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersRequest) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersRequest) SetProjectId(v string) *ListUpstreamAppServersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamAppServersRequest) SetPageIndex(v string) *ListUpstreamAppServersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListUpstreamAppServersRequest) SetPageSize(v string) *ListUpstreamAppServersRequest {
	s.PageSize = &v
	return s
}

type ListUpstreamAppServersResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppServers *ListUpstreamAppServersResponseBodyAppServers `json:"AppServers,omitempty" xml:"AppServers,omitempty" type:"Struct"`
}

func (s ListUpstreamAppServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBody) SetRequestId(v string) *ListUpstreamAppServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUpstreamAppServersResponseBody) SetAppServers(v *ListUpstreamAppServersResponseBodyAppServers) *ListUpstreamAppServersResponseBody {
	s.AppServers = v
	return s
}

type ListUpstreamAppServersResponseBodyAppServers struct {
	Pagination *ListUpstreamAppServersResponseBodyAppServersPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	List       []*ListUpstreamAppServersResponseBodyAppServersList     `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListUpstreamAppServersResponseBodyAppServers) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBodyAppServers) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBodyAppServers) SetPagination(v *ListUpstreamAppServersResponseBodyAppServersPagination) *ListUpstreamAppServersResponseBodyAppServers {
	s.Pagination = v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServers) SetList(v []*ListUpstreamAppServersResponseBodyAppServersList) *ListUpstreamAppServersResponseBodyAppServers {
	s.List = v
	return s
}

type ListUpstreamAppServersResponseBodyAppServersPagination struct {
	PageIndex      *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListUpstreamAppServersResponseBodyAppServersPagination) SetTotalPageCount(v int32) *ListUpstreamAppServersResponseBodyAppServersPagination {
	s.TotalPageCount = &v
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

type ListUpstreamAppServersResponseBodyAppServersList struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	QueueNameList *string `json:"QueueNameList,omitempty" xml:"QueueNameList,omitempty"`
	PAppKey       *string `json:"PAppKey,omitempty" xml:"PAppKey,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUpstreamAppServersResponseBodyAppServersList) String() string {
	return tea.Prettify(s)
}

func (s ListUpstreamAppServersResponseBodyAppServersList) GoString() string {
	return s.String()
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetProjectId(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.ProjectId = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetGmtCreate(v int64) *ListUpstreamAppServersResponseBodyAppServersList {
	s.GmtCreate = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetTags(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.Tags = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetQueueNameList(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.QueueNameList = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetPAppKey(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.PAppKey = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetGmtModified(v int64) *ListUpstreamAppServersResponseBodyAppServersList {
	s.GmtModified = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetName(v string) *ListUpstreamAppServersResponseBodyAppServersList {
	s.Name = &v
	return s
}

func (s *ListUpstreamAppServersResponseBodyAppServersList) SetId(v int64) *ListUpstreamAppServersResponseBodyAppServersList {
	s.Id = &v
	return s
}

type ListUpstreamAppServersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUpstreamAppServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceGroupList []*ListVersionDeviceGroupsResponseBodyDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Repeated"`
}

func (s ListVersionDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVersionDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionDeviceGroupsResponseBody) SetRequestId(v string) *ListVersionDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBody) SetDeviceGroupList(v []*ListVersionDeviceGroupsResponseBodyDeviceGroupList) *ListVersionDeviceGroupsResponseBody {
	s.DeviceGroupList = v
	return s
}

type ListVersionDeviceGroupsResponseBodyDeviceGroupList struct {
	GmtModify   *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListVersionDeviceGroupsResponseBodyDeviceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListVersionDeviceGroupsResponseBodyDeviceGroupList) GoString() string {
	return s.String()
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetGmtModify(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.GmtModify = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetDescription(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Description = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetGmtCreate(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.GmtCreate = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetName(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Name = &v
	return s
}

func (s *ListVersionDeviceGroupsResponseBodyDeviceGroupList) SetId(v string) *ListVersionDeviceGroupsResponseBodyDeviceGroupList {
	s.Id = &v
	return s
}

type ListVersionDeviceGroupsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVersionDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListVersionDeviceGroupsResponse) SetBody(v *ListVersionDeviceGroupsResponseBody) *ListVersionDeviceGroupsResponse {
	s.Body = v
	return s
}

type PublishAppVersionRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	SendMessage *bool   `json:"SendMessage,omitempty" xml:"SendMessage,omitempty"`
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

func (s *PublishAppVersionRequest) SetVersionId(v string) *PublishAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *PublishAppVersionRequest) SetSendMessage(v bool) *PublishAppVersionRequest {
	s.SendMessage = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishAppVersionResponse) SetBody(v *PublishAppVersionResponseBody) *PublishAppVersionResponse {
	s.Body = v
	return s
}

type PublishMqttMessageRequest struct {
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	QoS       *int32  `json:"QoS,omitempty" xml:"QoS,omitempty"`
}

func (s PublishMqttMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishMqttMessageRequest) GoString() string {
	return s.String()
}

func (s *PublishMqttMessageRequest) SetAppKey(v string) *PublishMqttMessageRequest {
	s.AppKey = &v
	return s
}

func (s *PublishMqttMessageRequest) SetTopic(v string) *PublishMqttMessageRequest {
	s.Topic = &v
	return s
}

func (s *PublishMqttMessageRequest) SetProjectId(v string) *PublishMqttMessageRequest {
	s.ProjectId = &v
	return s
}

func (s *PublishMqttMessageRequest) SetMessage(v string) *PublishMqttMessageRequest {
	s.Message = &v
	return s
}

func (s *PublishMqttMessageRequest) SetQoS(v int32) *PublishMqttMessageRequest {
	s.QoS = &v
	return s
}

type PublishMqttMessageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Mid       *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s PublishMqttMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishMqttMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PublishMqttMessageResponseBody) SetRequestId(v string) *PublishMqttMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishMqttMessageResponseBody) SetMid(v string) *PublishMqttMessageResponseBody {
	s.Mid = &v
	return s
}

type PublishMqttMessageResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishMqttMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishMqttMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishMqttMessageResponse) GoString() string {
	return s.String()
}

func (s *PublishMqttMessageResponse) SetHeaders(v map[string]*string) *PublishMqttMessageResponse {
	s.Headers = v
	return s
}

func (s *PublishMqttMessageResponse) SetBody(v *PublishMqttMessageResponseBody) *PublishMqttMessageResponse {
	s.Body = v
	return s
}

type PublishOsVersionRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	SendMessage *bool   `json:"SendMessage,omitempty" xml:"SendMessage,omitempty"`
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

func (s *PublishOsVersionRequest) SetVersionId(v string) *PublishOsVersionRequest {
	s.VersionId = &v
	return s
}

func (s *PublishOsVersionRequest) SetSendMessage(v bool) *PublishOsVersionRequest {
	s.SendMessage = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishOsVersionResponse) SetBody(v *PublishOsVersionResponseBody) *PublishOsVersionResponse {
	s.Body = v
	return s
}

type PushMessageRequest struct {
	AppPackage     *string `json:"AppPackage,omitempty" xml:"AppPackage,omitempty"`
	Desc           *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Act            *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Uri            *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	PkgContent     *string `json:"PkgContent,omitempty" xml:"PkgContent,omitempty"`
	CustomContent  *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	ReceiverType   *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ReceiverValues *string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty"`
	ExpiredTime    *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppKey         *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PushMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMessageRequest) GoString() string {
	return s.String()
}

func (s *PushMessageRequest) SetAppPackage(v string) *PushMessageRequest {
	s.AppPackage = &v
	return s
}

func (s *PushMessageRequest) SetDesc(v string) *PushMessageRequest {
	s.Desc = &v
	return s
}

func (s *PushMessageRequest) SetAct(v string) *PushMessageRequest {
	s.Act = &v
	return s
}

func (s *PushMessageRequest) SetUri(v string) *PushMessageRequest {
	s.Uri = &v
	return s
}

func (s *PushMessageRequest) SetPkgContent(v string) *PushMessageRequest {
	s.PkgContent = &v
	return s
}

func (s *PushMessageRequest) SetCustomContent(v string) *PushMessageRequest {
	s.CustomContent = &v
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

func (s *PushMessageRequest) SetExpiredTime(v int64) *PushMessageRequest {
	s.ExpiredTime = &v
	return s
}

func (s *PushMessageRequest) SetTitle(v string) *PushMessageRequest {
	s.Title = &v
	return s
}

func (s *PushMessageRequest) SetProjectId(v string) *PushMessageRequest {
	s.ProjectId = &v
	return s
}

func (s *PushMessageRequest) SetAppKey(v string) *PushMessageRequest {
	s.AppKey = &v
	return s
}

func (s *PushMessageRequest) SetType(v int32) *PushMessageRequest {
	s.Type = &v
	return s
}

type PushMessageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Mid       *int64  `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s PushMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushMessageResponseBody) SetRequestId(v string) *PushMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMessageResponseBody) SetMid(v int64) *PushMessageResponseBody {
	s.Mid = &v
	return s
}

type PushMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushVersionMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PushVersionMessageResponse) SetBody(v *PushVersionMessageResponseBody) *PushVersionMessageResponse {
	s.Body = v
	return s
}

type QueryPrepublishPassedDeviceCountRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
}

func (s QueryPrepublishPassedDeviceCountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepublishPassedDeviceCountRequest) GoString() string {
	return s.String()
}

func (s *QueryPrepublishPassedDeviceCountRequest) SetProjectId(v string) *QueryPrepublishPassedDeviceCountRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryPrepublishPassedDeviceCountRequest) SetPrepublishId(v string) *QueryPrepublishPassedDeviceCountRequest {
	s.PrepublishId = &v
	return s
}

type QueryPrepublishPassedDeviceCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s QueryPrepublishPassedDeviceCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepublishPassedDeviceCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPrepublishPassedDeviceCountResponseBody) SetRequestId(v string) *QueryPrepublishPassedDeviceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPrepublishPassedDeviceCountResponseBody) SetCount(v int32) *QueryPrepublishPassedDeviceCountResponseBody {
	s.Count = &v
	return s
}

type QueryPrepublishPassedDeviceCountResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPrepublishPassedDeviceCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryPrepublishPassedDeviceCountResponse) SetBody(v *QueryPrepublishPassedDeviceCountResponseBody) *QueryPrepublishPassedDeviceCountResponse {
	s.Body = v
	return s
}

type SubmitAssistReportRequest struct {
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AssistId          *string `json:"AssistId,omitempty" xml:"AssistId,omitempty"`
	AssistDescription *string `json:"AssistDescription,omitempty" xml:"AssistDescription,omitempty"`
	AssistResult      *string `json:"AssistResult,omitempty" xml:"AssistResult,omitempty"`
	AssistTag         *string `json:"AssistTag,omitempty" xml:"AssistTag,omitempty"`
	AssistReason      *string `json:"AssistReason,omitempty" xml:"AssistReason,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
}

func (s SubmitAssistReportRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAssistReportRequest) GoString() string {
	return s.String()
}

func (s *SubmitAssistReportRequest) SetProjectId(v string) *SubmitAssistReportRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitAssistReportRequest) SetAssistId(v string) *SubmitAssistReportRequest {
	s.AssistId = &v
	return s
}

func (s *SubmitAssistReportRequest) SetAssistDescription(v string) *SubmitAssistReportRequest {
	s.AssistDescription = &v
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

func (s *SubmitAssistReportRequest) SetAssistReason(v string) *SubmitAssistReportRequest {
	s.AssistReason = &v
	return s
}

func (s *SubmitAssistReportRequest) SetDeviceModel(v string) *SubmitAssistReportRequest {
	s.DeviceModel = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAssistReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitAssistReportResponse) SetBody(v *SubmitAssistReportResponseBody) *SubmitAssistReportResponse {
	s.Body = v
	return s
}

type UpdateApiGatewayAppStatusRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GatewayAppId *string `json:"GatewayAppId,omitempty" xml:"GatewayAppId,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateApiGatewayAppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiGatewayAppStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiGatewayAppStatusRequest) SetProjectId(v string) *UpdateApiGatewayAppStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateApiGatewayAppStatusRequest) SetGatewayAppId(v string) *UpdateApiGatewayAppStatusRequest {
	s.GatewayAppId = &v
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApiGatewayAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateApiGatewayAppStatusResponse) SetBody(v *UpdateApiGatewayAppStatusResponseBody) *UpdateApiGatewayAppStatusResponse {
	s.Body = v
	return s
}

type UpdateAppBlackWhiteVersionsRequest struct {
	WhiteAppVersions *string `json:"WhiteAppVersions,omitempty" xml:"WhiteAppVersions,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	BlackAppVersions *string `json:"BlackAppVersions,omitempty" xml:"BlackAppVersions,omitempty"`
}

func (s UpdateAppBlackWhiteVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppBlackWhiteVersionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppBlackWhiteVersionsRequest) SetWhiteAppVersions(v string) *UpdateAppBlackWhiteVersionsRequest {
	s.WhiteAppVersions = &v
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

func (s *UpdateAppBlackWhiteVersionsRequest) SetBlackAppVersions(v string) *UpdateAppBlackWhiteVersionsRequest {
	s.BlackAppVersions = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppBlackWhiteVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAppBlackWhiteVersionsResponse) SetBody(v *UpdateAppBlackWhiteVersionsResponseBody) *UpdateAppBlackWhiteVersionsResponse {
	s.Body = v
	return s
}

type UpdateAppVersionRequest struct {
	BlackVersionList  *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	IsAllowNewInstall *string `json:"IsAllowNewInstall,omitempty" xml:"IsAllowNewInstall,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion        *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	InstallType       *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	IsForceUpgrade    *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	IsSilentUpgrade   *string `json:"IsSilentUpgrade,omitempty" xml:"IsSilentUpgrade,omitempty"`
	IsNeedRestart     *string `json:"IsNeedRestart,omitempty" xml:"IsNeedRestart,omitempty"`
	PackageUrl        *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	ReleaseNote       *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	WhiteVersionList  *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
	RestartType       *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
	RestartAppType    *string `json:"RestartAppType,omitempty" xml:"RestartAppType,omitempty"`
	RestartAppParam   *string `json:"RestartAppParam,omitempty" xml:"RestartAppParam,omitempty"`
	DeviceAdapterList *string `json:"DeviceAdapterList,omitempty" xml:"DeviceAdapterList,omitempty"`
	VersionId         *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ApkMd5            *string `json:"ApkMd5,omitempty" xml:"ApkMd5,omitempty"`
}

func (s UpdateAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionRequest) SetBlackVersionList(v string) *UpdateAppVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsAllowNewInstall(v string) *UpdateAppVersionRequest {
	s.IsAllowNewInstall = &v
	return s
}

func (s *UpdateAppVersionRequest) SetProjectId(v string) *UpdateAppVersionRequest {
	s.ProjectId = &v
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

func (s *UpdateAppVersionRequest) SetVersionCode(v string) *UpdateAppVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *UpdateAppVersionRequest) SetInstallType(v string) *UpdateAppVersionRequest {
	s.InstallType = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRemark(v string) *UpdateAppVersionRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsForceUpgrade(v string) *UpdateAppVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsSilentUpgrade(v string) *UpdateAppVersionRequest {
	s.IsSilentUpgrade = &v
	return s
}

func (s *UpdateAppVersionRequest) SetIsNeedRestart(v string) *UpdateAppVersionRequest {
	s.IsNeedRestart = &v
	return s
}

func (s *UpdateAppVersionRequest) SetPackageUrl(v string) *UpdateAppVersionRequest {
	s.PackageUrl = &v
	return s
}

func (s *UpdateAppVersionRequest) SetReleaseNote(v string) *UpdateAppVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *UpdateAppVersionRequest) SetWhiteVersionList(v string) *UpdateAppVersionRequest {
	s.WhiteVersionList = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRestartType(v string) *UpdateAppVersionRequest {
	s.RestartType = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRestartAppType(v string) *UpdateAppVersionRequest {
	s.RestartAppType = &v
	return s
}

func (s *UpdateAppVersionRequest) SetRestartAppParam(v string) *UpdateAppVersionRequest {
	s.RestartAppParam = &v
	return s
}

func (s *UpdateAppVersionRequest) SetDeviceAdapterList(v string) *UpdateAppVersionRequest {
	s.DeviceAdapterList = &v
	return s
}

func (s *UpdateAppVersionRequest) SetVersionId(v string) *UpdateAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateAppVersionRequest) SetApkMd5(v string) *UpdateAppVersionRequest {
	s.ApkMd5 = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAppVersionResponse) SetBody(v *UpdateAppVersionResponseBody) *UpdateAppVersionResponse {
	s.Body = v
	return s
}

type UpdateAppVersionReleaseNoteRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
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

func (s *UpdateAppVersionReleaseNoteRequest) SetVersionId(v string) *UpdateAppVersionReleaseNoteRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateAppVersionReleaseNoteRequest) SetReleaseNote(v string) *UpdateAppVersionReleaseNoteRequest {
	s.ReleaseNote = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppVersionReleaseNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAppVersionReleaseNoteResponse) SetBody(v *UpdateAppVersionReleaseNoteResponseBody) *UpdateAppVersionReleaseNoteResponse {
	s.Body = v
	return s
}

type UpdateAppVersionRemarkRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s *UpdateAppVersionRemarkRequest) SetVersionId(v string) *UpdateAppVersionRemarkRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateAppVersionRemarkRequest) SetRemark(v string) *UpdateAppVersionRemarkRequest {
	s.Remark = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppVersionRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAppVersionRemarkResponse) SetBody(v *UpdateAppVersionRemarkResponseBody) *UpdateAppVersionRemarkResponse {
	s.Body = v
	return s
}

type UpdateAppVersionStatusRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAppVersionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppVersionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppVersionStatusRequest) SetProjectId(v string) *UpdateAppVersionStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAppVersionStatusRequest) SetId(v string) *UpdateAppVersionStatusRequest {
	s.Id = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAppVersionStatusResponse) SetBody(v *UpdateAppVersionStatusResponseBody) *UpdateAppVersionStatusResponse {
	s.Body = v
	return s
}

type UpdateCustomizedFilterRequest struct {
	BlackWhiteType   *string `json:"BlackWhiteType,omitempty" xml:"BlackWhiteType,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueType        *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	ValueCompareType *string `json:"ValueCompareType,omitempty" xml:"ValueCompareType,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *UpdateCustomizedFilterRequest) SetProjectId(v string) *UpdateCustomizedFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetName(v string) *UpdateCustomizedFilterRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetValue(v string) *UpdateCustomizedFilterRequest {
	s.Value = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetValueType(v string) *UpdateCustomizedFilterRequest {
	s.ValueType = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetValueCompareType(v string) *UpdateCustomizedFilterRequest {
	s.ValueCompareType = &v
	return s
}

func (s *UpdateCustomizedFilterRequest) SetId(v int64) *UpdateCustomizedFilterRequest {
	s.Id = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCustomizedFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateCustomizedFilterResponse) SetBody(v *UpdateCustomizedFilterResponseBody) *UpdateCustomizedFilterResponse {
	s.Body = v
	return s
}

type UpdateDeviceModelRequest struct {
	InitUsageType     *string `json:"InitUsageType,omitempty" xml:"InitUsageType,omitempty"`
	ModelName         *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	CanCreateDeviceId *string `json:"CanCreateDeviceId,omitempty" xml:"CanCreateDeviceId,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	HardwareType      *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	SecurityChip      *string `json:"SecurityChip,omitempty" xml:"SecurityChip,omitempty"`
	OsPlatform        *string `json:"OsPlatform,omitempty" xml:"OsPlatform,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s UpdateDeviceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceModelRequest) SetInitUsageType(v string) *UpdateDeviceModelRequest {
	s.InitUsageType = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetModelName(v string) *UpdateDeviceModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetId(v string) *UpdateDeviceModelRequest {
	s.Id = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetBrandName(v string) *UpdateDeviceModelRequest {
	s.BrandName = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetDescription(v string) *UpdateDeviceModelRequest {
	s.Description = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetDeviceType(v string) *UpdateDeviceModelRequest {
	s.DeviceType = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetCanCreateDeviceId(v string) *UpdateDeviceModelRequest {
	s.CanCreateDeviceId = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetProjectId(v string) *UpdateDeviceModelRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetHardwareType(v string) *UpdateDeviceModelRequest {
	s.HardwareType = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetSecurityChip(v string) *UpdateDeviceModelRequest {
	s.SecurityChip = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetOsPlatform(v string) *UpdateDeviceModelRequest {
	s.OsPlatform = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetObjectKey(v string) *UpdateDeviceModelRequest {
	s.ObjectKey = &v
	return s
}

func (s *UpdateDeviceModelRequest) SetDeviceName(v string) *UpdateDeviceModelRequest {
	s.DeviceName = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDeviceModelResponse) SetBody(v *UpdateDeviceModelResponseBody) *UpdateDeviceModelResponse {
	s.Body = v
	return s
}

type UpdateNamespaceDataRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	AuthType     *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeviceIdType *string `json:"DeviceIdType,omitempty" xml:"DeviceIdType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	AccountType  *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
	OldData      *string `json:"OldData,omitempty" xml:"OldData,omitempty"`
	NewData      *string `json:"NewData,omitempty" xml:"NewData,omitempty"`
}

func (s UpdateNamespaceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceDataRequest) SetProjectId(v string) *UpdateNamespaceDataRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetNamespace(v string) *UpdateNamespaceDataRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetAuthType(v string) *UpdateNamespaceDataRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetDeviceIdType(v string) *UpdateNamespaceDataRequest {
	s.DeviceIdType = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetDeviceId(v string) *UpdateNamespaceDataRequest {
	s.DeviceId = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetAccountType(v string) *UpdateNamespaceDataRequest {
	s.AccountType = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetAccountId(v string) *UpdateNamespaceDataRequest {
	s.AccountId = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetPath(v string) *UpdateNamespaceDataRequest {
	s.Path = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetOldData(v string) *UpdateNamespaceDataRequest {
	s.OldData = &v
	return s
}

func (s *UpdateNamespaceDataRequest) SetNewData(v string) *UpdateNamespaceDataRequest {
	s.NewData = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateNamespaceDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateNamespaceDataResponse) SetBody(v *UpdateNamespaceDataResponseBody) *UpdateNamespaceDataResponse {
	s.Body = v
	return s
}

type UpdateOsBlackWhiteVersionsRequest struct {
	WhiteVersions *string `json:"WhiteVersions,omitempty" xml:"WhiteVersions,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	BlackVersions *string `json:"BlackVersions,omitempty" xml:"BlackVersions,omitempty"`
}

func (s UpdateOsBlackWhiteVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsBlackWhiteVersionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsBlackWhiteVersionsRequest) SetWhiteVersions(v string) *UpdateOsBlackWhiteVersionsRequest {
	s.WhiteVersions = &v
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

func (s *UpdateOsBlackWhiteVersionsRequest) SetBlackVersions(v string) *UpdateOsBlackWhiteVersionsRequest {
	s.BlackVersions = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOsBlackWhiteVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOsBlackWhiteVersionsResponse) SetBody(v *UpdateOsBlackWhiteVersionsResponseBody) *UpdateOsBlackWhiteVersionsResponse {
	s.Body = v
	return s
}

type UpdateOsVersionRequest struct {
	IsMilestone                 *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	IsForceNightUpgrade         *string `json:"IsForceNightUpgrade,omitempty" xml:"IsForceNightUpgrade,omitempty"`
	ProjectId                   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceModelId               *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	SystemVersion               *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	ReleaseNote                 *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Remark                      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	IsForceUpgrade              *string `json:"IsForceUpgrade,omitempty" xml:"IsForceUpgrade,omitempty"`
	BlackVersionList            *string `json:"BlackVersionList,omitempty" xml:"BlackVersionList,omitempty"`
	WhiteVersionList            *string `json:"WhiteVersionList,omitempty" xml:"WhiteVersionList,omitempty"`
	MaxClientVersion            *string `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	MinClientVersion            *string `json:"MinClientVersion,omitempty" xml:"MinClientVersion,omitempty"`
	NightUpgradeDownloadType    *string `json:"NightUpgradeDownloadType,omitempty" xml:"NightUpgradeDownloadType,omitempty"`
	NightUpgradeIsShowTip       *string `json:"NightUpgradeIsShowTip,omitempty" xml:"NightUpgradeIsShowTip,omitempty"`
	NightUpgradeIsAllowedCancel *string `json:"NightUpgradeIsAllowedCancel,omitempty" xml:"NightUpgradeIsAllowedCancel,omitempty"`
	RomList                     *string `json:"RomList,omitempty" xml:"RomList,omitempty"`
	Id                          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	EnableMobileDownload        *string `json:"EnableMobileDownload,omitempty" xml:"EnableMobileDownload,omitempty"`
	MobileDownloadMaxSize       *string `json:"MobileDownloadMaxSize,omitempty" xml:"MobileDownloadMaxSize,omitempty"`
}

func (s UpdateOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionRequest) SetIsMilestone(v string) *UpdateOsVersionRequest {
	s.IsMilestone = &v
	return s
}

func (s *UpdateOsVersionRequest) SetIsForceNightUpgrade(v string) *UpdateOsVersionRequest {
	s.IsForceNightUpgrade = &v
	return s
}

func (s *UpdateOsVersionRequest) SetProjectId(v string) *UpdateOsVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsVersionRequest) SetDeviceModelId(v string) *UpdateOsVersionRequest {
	s.DeviceModelId = &v
	return s
}

func (s *UpdateOsVersionRequest) SetSystemVersion(v string) *UpdateOsVersionRequest {
	s.SystemVersion = &v
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

func (s *UpdateOsVersionRequest) SetIsForceUpgrade(v string) *UpdateOsVersionRequest {
	s.IsForceUpgrade = &v
	return s
}

func (s *UpdateOsVersionRequest) SetBlackVersionList(v string) *UpdateOsVersionRequest {
	s.BlackVersionList = &v
	return s
}

func (s *UpdateOsVersionRequest) SetWhiteVersionList(v string) *UpdateOsVersionRequest {
	s.WhiteVersionList = &v
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

func (s *UpdateOsVersionRequest) SetNightUpgradeDownloadType(v string) *UpdateOsVersionRequest {
	s.NightUpgradeDownloadType = &v
	return s
}

func (s *UpdateOsVersionRequest) SetNightUpgradeIsShowTip(v string) *UpdateOsVersionRequest {
	s.NightUpgradeIsShowTip = &v
	return s
}

func (s *UpdateOsVersionRequest) SetNightUpgradeIsAllowedCancel(v string) *UpdateOsVersionRequest {
	s.NightUpgradeIsAllowedCancel = &v
	return s
}

func (s *UpdateOsVersionRequest) SetRomList(v string) *UpdateOsVersionRequest {
	s.RomList = &v
	return s
}

func (s *UpdateOsVersionRequest) SetId(v string) *UpdateOsVersionRequest {
	s.Id = &v
	return s
}

func (s *UpdateOsVersionRequest) SetEnableMobileDownload(v string) *UpdateOsVersionRequest {
	s.EnableMobileDownload = &v
	return s
}

func (s *UpdateOsVersionRequest) SetMobileDownloadMaxSize(v string) *UpdateOsVersionRequest {
	s.MobileDownloadMaxSize = &v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOsVersionResponse) SetBody(v *UpdateOsVersionResponseBody) *UpdateOsVersionResponse {
	s.Body = v
	return s
}

type UpdateOsVersionReleaseNoteRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
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

func (s *UpdateOsVersionReleaseNoteRequest) SetVersionId(v string) *UpdateOsVersionReleaseNoteRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateOsVersionReleaseNoteRequest) SetReleaseNote(v string) *UpdateOsVersionReleaseNoteRequest {
	s.ReleaseNote = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOsVersionReleaseNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOsVersionReleaseNoteResponse) SetBody(v *UpdateOsVersionReleaseNoteResponseBody) *UpdateOsVersionReleaseNoteResponse {
	s.Body = v
	return s
}

type UpdateOsVersionRemarkRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s *UpdateOsVersionRemarkRequest) SetVersionId(v string) *UpdateOsVersionRemarkRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateOsVersionRemarkRequest) SetRemark(v string) *UpdateOsVersionRemarkRequest {
	s.Remark = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOsVersionRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOsVersionRemarkResponse) SetBody(v *UpdateOsVersionRemarkResponseBody) *UpdateOsVersionRemarkResponse {
	s.Body = v
	return s
}

type UpdateOsVersionStatusRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateOsVersionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionStatusRequest) SetProjectId(v string) *UpdateOsVersionStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateOsVersionStatusRequest) SetId(v string) *UpdateOsVersionStatusRequest {
	s.Id = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOsVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOsVersionStatusResponse) SetBody(v *UpdateOsVersionStatusResponseBody) *UpdateOsVersionStatusResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateSchemaSubscribeRequest struct {
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	SubscribeList *string `json:"SubscribeList,omitempty" xml:"SubscribeList,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
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

func (s *UpdateSchemaSubscribeRequest) SetSubscribeList(v string) *UpdateSchemaSubscribeRequest {
	s.SubscribeList = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSchemaSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateSchemaSubscribeResponse) SetBody(v *UpdateSchemaSubscribeResponseBody) *UpdateSchemaSubscribeResponse {
	s.Body = v
	return s
}

type UpdateShadowSchemaRequest struct {
	DeviceModelId *string `json:"DeviceModelId,omitempty" xml:"DeviceModelId,omitempty"`
	AuthType      *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Schema        *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateShadowSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateShadowSchemaRequest) GoString() string {
	return s.String()
}

func (s *UpdateShadowSchemaRequest) SetDeviceModelId(v string) *UpdateShadowSchemaRequest {
	s.DeviceModelId = &v
	return s
}

func (s *UpdateShadowSchemaRequest) SetAuthType(v string) *UpdateShadowSchemaRequest {
	s.AuthType = &v
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

func (s *UpdateShadowSchemaRequest) SetId(v string) *UpdateShadowSchemaRequest {
	s.Id = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateShadowSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateShadowSchemaResponse) SetBody(v *UpdateShadowSchemaResponseBody) *UpdateShadowSchemaResponse {
	s.Body = v
	return s
}

type UpdateTriggerRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Sandbox    *int32  `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Production *int32  `json:"Production,omitempty" xml:"Production,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) SetProjectId(v string) *UpdateTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTriggerRequest) SetId(v int64) *UpdateTriggerRequest {
	s.Id = &v
	return s
}

func (s *UpdateTriggerRequest) SetSandbox(v int32) *UpdateTriggerRequest {
	s.Sandbox = &v
	return s
}

func (s *UpdateTriggerRequest) SetProduction(v int32) *UpdateTriggerRequest {
	s.Production = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateTriggerResponse) SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse {
	s.Body = v
	return s
}

type UpdateUpstreamAppServerRequest struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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

func (s *UpdateUpstreamAppServerRequest) SetTags(v string) *UpdateUpstreamAppServerRequest {
	s.Tags = &v
	return s
}

func (s *UpdateUpstreamAppServerRequest) SetProjectId(v string) *UpdateUpstreamAppServerRequest {
	s.ProjectId = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUpstreamAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateUpstreamAppServerResponse) SetBody(v *UpdateUpstreamAppServerResponseBody) *UpdateUpstreamAppServerResponse {
	s.Body = v
	return s
}

type UpdateVersionDeviceGroupRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateVersionDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateVersionDeviceGroupRequest) SetProjectId(v string) *UpdateVersionDeviceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateVersionDeviceGroupRequest) SetName(v string) *UpdateVersionDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateVersionDeviceGroupRequest) SetDescription(v string) *UpdateVersionDeviceGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateVersionDeviceGroupRequest) SetId(v string) *UpdateVersionDeviceGroupRequest {
	s.Id = &v
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVersionDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateVersionDeviceGroupResponse) SetBody(v *UpdateVersionDeviceGroupResponseBody) *UpdateVersionDeviceGroupResponse {
	s.Body = v
	return s
}

type UpdateVersionPrepublishActiveStatusRequest struct {
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PrepublishId *string `json:"PrepublishId,omitempty" xml:"PrepublishId,omitempty"`
	IsActive     *string `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
}

func (s UpdateVersionPrepublishActiveStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionPrepublishActiveStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateVersionPrepublishActiveStatusRequest) SetProjectId(v string) *UpdateVersionPrepublishActiveStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateVersionPrepublishActiveStatusRequest) SetPrepublishId(v string) *UpdateVersionPrepublishActiveStatusRequest {
	s.PrepublishId = &v
	return s
}

func (s *UpdateVersionPrepublishActiveStatusRequest) SetIsActive(v string) *UpdateVersionPrepublishActiveStatusRequest {
	s.IsActive = &v
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVersionPrepublishActiveStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUploadedFunctionFileInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUploadedFunctionFileInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVersionBlackDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVersionBlackDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVersionGroupDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVersionGroupDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVersionWhiteDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVersionWhiteDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVersionWhiteDevicesByDeviceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVersionWhiteDevicesByDeviceGroups"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConnectAssistDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConnectAssistDevice"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CountDeviceBrandsWithOptions(request *CountDeviceBrandsRequest, runtime *util.RuntimeOptions) (_result *CountDeviceBrandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CountDeviceBrandsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CountDeviceBrands"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &CountDeviceModelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CountDeviceModels"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &CountDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CountDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CountProjectsWithOptions(runtime *util.RuntimeOptions) (_result *CountProjectsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &CountProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CountProjects"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountProjects() (_result *CountProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountProjectsResponse{}
	_body, _err := client.CountProjectsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountYunIdInfoWithOptions(runtime *util.RuntimeOptions) (_result *CountYunIdInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &CountYunIdInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CountYunIdInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAppVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCustomizedFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCustomizedFilter"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCustomizedPropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCustomizedProperty"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateDeviceWithOptions(request *CreateDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevice"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateDeviceBrandWithOptions(request *CreateDeviceBrandRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceBrandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDeviceBrandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDeviceBrand"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDeviceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDeviceModel"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateMqttRootTopicWithOptions(request *CreateMqttRootTopicRequest, runtime *util.RuntimeOptions) (_result *CreateMqttRootTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMqttRootTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMqttRootTopic"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMqttRootTopic(request *CreateMqttRootTopicRequest) (_result *CreateMqttRootTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMqttRootTopicResponse{}
	_body, _err := client.CreateMqttRootTopicWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNamespace"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOsVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProjectApp"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRpcServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRpcService"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSchemaSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSchemaSubscribe"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateShadowSchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateShadowSchema"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTrigger"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUpstreamAppKeyRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUpstreamAppKeyRelation"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUpstreamAppKeyRelationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUpstreamAppKeyRelations"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUpstreamAppServerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUpstreamAppServer"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVersionDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVersionDeviceGroup"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVersionPrepublishResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVersionPrepublish"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVersionTestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVersionTest"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DelayPublishOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DelayPublishOsVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAllCustomizedFiltersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAllCustomizedFilters"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAllVersionGroupDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAllVersionGroupDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCustomizedFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCustomizedFilter"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCustomizedPropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCustomizedProperty"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevice"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFunctionFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFunctionFile"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteMqttRootTopicWithOptions(request *DeleteMqttRootTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteMqttRootTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMqttRootTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMqttRootTopic"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMqttRootTopic(request *DeleteMqttRootTopicRequest) (_result *DeleteMqttRootTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMqttRootTopicResponse{}
	_body, _err := client.DeleteMqttRootTopicWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNamespace"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &DeleteOpenAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOpenAccount"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProjectApp"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRpcServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRpcService"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSchemaSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSchemaSubscribe"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteShadowSchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteShadowSchema"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTrigger"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUpstreamAppKeyRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUpstreamAppKeyRelation"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUpstreamAppServerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUpstreamAppServer"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionAllBlackDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionAllBlackDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionAllWhiteDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionAllWhiteDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionBlackDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionBlackDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionBlackDevicesByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionBlackDevicesById"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionDeviceGroup"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionGroupDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionGroupDevice"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionGroupDeviceByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionGroupDeviceById"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionTestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionTest"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionWhiteDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionWhiteDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVersionWhiteDevicesByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVersionWhiteDevicesById"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployFunctionFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployFunctionFile"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeApiGatewayAppSecurityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeApiGatewayAppSecurity"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAppVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAssistReportWithOptions(request *DescribeAssistReportRequest, runtime *util.RuntimeOptions) (_result *DescribeAssistReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeAssistReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAssistReport"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAssistRTMPServerAddressWithOptions(request *DescribeAssistRTMPServerAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeAssistRTMPServerAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeAssistRTMPServerAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAssistRTMPServerAddress"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAssistWSServerAddressWithOptions(request *DescribeAssistWSServerAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeAssistWSServerAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeAssistWSServerAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAssistWSServerAddress"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomizedFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomizedFilter"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDefaultSchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDefaultSchema"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDeviceWithOptions(request *DescribeDeviceRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDevice"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDeviceBrandWithOptions(request *DescribeDeviceBrandRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceBrandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDeviceBrandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceBrand"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &DescribeDeviceIdByOuterInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceIdByOuterInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &DescribeDeviceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceModel"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeviceOnlineInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceOnlineInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeviceShadowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceShadow"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeviceValiditySchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceValiditySchema"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMessage"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeMqttClientStatusWithOptions(request *DescribeMqttClientStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeMqttClientStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMqttClientStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMqttClientStatus"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMqttClientStatus(request *DescribeMqttClientStatusRequest) (_result *DescribeMqttClientStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMqttClientStatusResponse{}
	_body, _err := client.DescribeMqttClientStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMqttMessageWithOptions(request *DescribeMqttMessageRequest, runtime *util.RuntimeOptions) (_result *DescribeMqttMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMqttMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMqttMessage"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMqttMessage(request *DescribeMqttMessageRequest) (_result *DescribeMqttMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMqttMessageResponse{}
	_body, _err := client.DescribeMqttMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMqttTopicSubscriptionWithOptions(request *DescribeMqttTopicSubscriptionRequest, runtime *util.RuntimeOptions) (_result *DescribeMqttTopicSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMqttTopicSubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMqttTopicSubscription"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMqttTopicSubscription(request *DescribeMqttTopicSubscriptionRequest) (_result *DescribeMqttTopicSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMqttTopicSubscriptionResponse{}
	_body, _err := client.DescribeMqttTopicSubscriptionWithOptions(request, runtime)
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
		Query: query,
	}
	_result = &DescribeOpenAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOpenAccount"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOsVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProject"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProjectAppSecurityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProjectAppSecurity"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeShadowSchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeShadowSchema"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVersionDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVersionDeviceGroup"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DiagnosisVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DiagnosisVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindAppVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindAppVersions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindCustomizedFiltersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindCustomizedFilters"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindCustomizedPropertiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindCustomizedProperties"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindOsVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindOsVersions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FindPrepublishesByParentIdWithOptions(request *FindPrepublishesByParentIdRequest, runtime *util.RuntimeOptions) (_result *FindPrepublishesByParentIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindPrepublishesByParentIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindPrepublishesByParentId"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindPrepublishesByVersionIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindPrepublishesByVersionId"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FindPrepublishPassedDevicesWithOptions(request *FindPrepublishPassedDevicesRequest, runtime *util.RuntimeOptions) (_result *FindPrepublishPassedDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindPrepublishPassedDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindPrepublishPassedDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FindVersionBlackDevicesWithOptions(request *FindVersionBlackDevicesRequest, runtime *util.RuntimeOptions) (_result *FindVersionBlackDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionBlackDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionBlackDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionDeviceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionDeviceGroups"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionGroupDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionGroupDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FindVersionMessagesWithOptions(request *FindVersionMessagesRequest, runtime *util.RuntimeOptions) (_result *FindVersionMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionMessagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionMessages"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FindVersionMessageSendRecordsWithOptions(request *FindVersionMessageSendRecordsRequest, runtime *util.RuntimeOptions) (_result *FindVersionMessageSendRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionMessageSendRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionMessageSendRecords"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FindVersionTestsWithOptions(request *FindVersionTestsRequest, runtime *util.RuntimeOptions) (_result *FindVersionTestsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionTestsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionTests"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FindVersionWhiteDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FindVersionWhiteDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GenerateAssistFileUploadUrlWithOptions(request *GenerateAssistFileUploadUrlRequest, runtime *util.RuntimeOptions) (_result *GenerateAssistFileUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateAssistFileUploadUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateAssistFileUploadUrl"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAssistFileUploadUrl(request *GenerateAssistFileUploadUrlRequest) (_result *GenerateAssistFileUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAssistFileUploadUrlResponse{}
	_body, _err := client.GenerateAssistFileUploadUrlWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateFunctionFileUploadMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateFunctionFileUploadMeta"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateOssPostPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateOssPostPolicy"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateOssUploadMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateOssUploadMeta"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateSdkDownloadInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateSdkDownloadInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateSysAppDownloadInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateSysAppDownloadInfo"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceAppUpdateFunnelEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceAppUpdateFunnelEvents"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceSystemUpdateFunnelEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceSystemUpdateFunnelEvents"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetNamespaceDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetNamespaceData"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetOssUploadMetaWithOptions(request *GetOssUploadMetaRequest, runtime *util.RuntimeOptions) (_result *GetOssUploadMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOssUploadMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOssUploadMeta"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InvokeFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InvokeFunction"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListApiGatewayAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApiGatewayApps"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApps"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListAssistActionDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAssistActionDetails"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListAssistDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAssistDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListAssistHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAssistHistories"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListAssistHistoryDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAssistHistoryDetails"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListClientPluginsWithOptions(request *ListClientPluginsRequest, runtime *util.RuntimeOptions) (_result *ListClientPluginsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClientPluginsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClientPlugins"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListClientPluginVersionsWithOptions(request *ListClientPluginVersionsRequest, runtime *util.RuntimeOptions) (_result *ListClientPluginVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClientPluginVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClientPluginVersions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListClientSdksWithOptions(request *ListClientSdksRequest, runtime *util.RuntimeOptions) (_result *ListClientSdksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClientSdksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClientSdks"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConnectLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConnectLogs"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeployedFunctionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeployedFunctions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListDeviceBrandsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceBrands"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDeviceModelWithOptions(request *ListDeviceModelRequest, runtime *util.RuntimeOptions) (_result *ListDeviceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDeviceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceModel"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceModel(request *ListDeviceModelRequest) (_result *ListDeviceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceModelResponse{}
	_body, _err := client.ListDeviceModelWithOptions(request, runtime)
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
		Query: query,
	}
	_result = &ListDeviceModelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceModels"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDeviceTypesWithOptions(request *ListDeviceTypesRequest, runtime *util.RuntimeOptions) (_result *ListDeviceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDeviceTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceTypes"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListFunctionExecuteLogWithOptions(request *ListFunctionExecuteLogRequest, runtime *util.RuntimeOptions) (_result *ListFunctionExecuteLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFunctionExecuteLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunctionExecuteLog"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFunctionFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunctionFiles"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFunctionFilesByProjectIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunctionFilesByProjectId"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMessageAcksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMessageAcks"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMessageReceiversResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMessageReceivers"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListMqttClientSubscriptionsWithOptions(request *ListMqttClientSubscriptionsRequest, runtime *util.RuntimeOptions) (_result *ListMqttClientSubscriptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMqttClientSubscriptionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMqttClientSubscriptions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqttClientSubscriptions(request *ListMqttClientSubscriptionsRequest) (_result *ListMqttClientSubscriptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqttClientSubscriptionsResponse{}
	_body, _err := client.ListMqttClientSubscriptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqttMessageLogsWithOptions(request *ListMqttMessageLogsRequest, runtime *util.RuntimeOptions) (_result *ListMqttMessageLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMqttMessageLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMqttMessageLogs"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqttMessageLogs(request *ListMqttMessageLogsRequest) (_result *ListMqttMessageLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqttMessageLogsResponse{}
	_body, _err := client.ListMqttMessageLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqttRootTopicsWithOptions(request *ListMqttRootTopicsRequest, runtime *util.RuntimeOptions) (_result *ListMqttRootTopicsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMqttRootTopicsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMqttRootTopics"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqttRootTopics(request *ListMqttRootTopicsRequest) (_result *ListMqttRootTopicsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqttRootTopicsResponse{}
	_body, _err := client.ListMqttRootTopicsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNamespaces"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOfflineMessagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOfflineMessages"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListOpenAccountLinksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOpenAccountLinks"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListOpenAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOpenAccounts"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_result = &ListPreChecksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPreChecks"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectApps"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_result = &ListProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjects"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRpcServicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRpcServices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListSchemaSubscribesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSchemaSubscribes"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_result = &ListServicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServices"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &ListShadowSchemaDeviceModelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListShadowSchemaDeviceModels"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListShadowSchemasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListShadowSchemas"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_result = &ListSupportFeaturesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSupportFeatures"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTriggers"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUpstreamAppKeyRelationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUpstreamAppKeyRelations"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUpstreamAppServersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUpstreamAppServers"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVersionDeviceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVersionDeviceGroups"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishAppVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishAppVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PublishMqttMessageWithOptions(request *PublishMqttMessageRequest, runtime *util.RuntimeOptions) (_result *PublishMqttMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishMqttMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishMqttMessage"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishMqttMessage(request *PublishMqttMessageRequest) (_result *PublishMqttMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishMqttMessageResponse{}
	_body, _err := client.PublishMqttMessageWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishOsVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PushMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PushMessage"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PushVersionMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PushVersionMessage"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPrepublishPassedDeviceCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPrepublishPassedDeviceCount"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAssistReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAssistReport"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateApiGatewayAppStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApiGatewayAppStatus"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppBlackWhiteVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppBlackWhiteVersions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppVersionReleaseNoteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppVersionReleaseNote"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppVersionRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppVersionRemark"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppVersionStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppVersionStatus"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCustomizedFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCustomizedFilter"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceModel"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateNamespaceDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateNamespaceData"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateOsBlackWhiteVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOsBlackWhiteVersions"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOsVersion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateOsVersionReleaseNoteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOsVersionReleaseNote"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateOsVersionRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOsVersionRemark"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateOsVersionStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOsVersionStatus"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSchemaSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSchemaSubscribe"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateShadowSchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateShadowSchema"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTrigger"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUpstreamAppServerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUpstreamAppServer"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVersionDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVersionDeviceGroup"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVersionPrepublishActiveStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVersionPrepublishActiveStatus"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
