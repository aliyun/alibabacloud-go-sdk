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

type AddDataSourceRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	DataSourceName    *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceType    *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileRetentionDays *int32  `json:"FileRetentionDays,omitempty" xml:"FileRetentionDays,omitempty"`
}

func (s AddDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceRequest) SetCorpId(v string) *AddDataSourceRequest {
	s.CorpId = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceName(v string) *AddDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceType(v string) *AddDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *AddDataSourceRequest) SetDescription(v string) *AddDataSourceRequest {
	s.Description = &v
	return s
}

func (s *AddDataSourceRequest) SetFileRetentionDays(v int32) *AddDataSourceRequest {
	s.FileRetentionDays = &v
	return s
}

type AddDataSourceResponseBody struct {
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    *AddDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code    *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBody) SetMessage(v string) *AddDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *AddDataSourceResponseBody) SetData(v *AddDataSourceResponseBodyData) *AddDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *AddDataSourceResponseBody) SetCode(v string) *AddDataSourceResponseBody {
	s.Code = &v
	return s
}

type AddDataSourceResponseBodyData struct {
	KafkaTopic   *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	OssPath      *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s AddDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBodyData) SetKafkaTopic(v string) *AddDataSourceResponseBodyData {
	s.KafkaTopic = &v
	return s
}

func (s *AddDataSourceResponseBodyData) SetDataSourceId(v string) *AddDataSourceResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *AddDataSourceResponseBodyData) SetOssPath(v string) *AddDataSourceResponseBodyData {
	s.OssPath = &v
	return s
}

type AddDataSourceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponse) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponse) SetHeaders(v map[string]*string) *AddDataSourceResponse {
	s.Headers = v
	return s
}

func (s *AddDataSourceResponse) SetBody(v *AddDataSourceResponseBody) *AddDataSourceResponse {
	s.Body = v
	return s
}

type AddDeviceRequest struct {
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType       *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceAddress    *string `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty"`
	DeviceSite       *string `json:"DeviceSite,omitempty" xml:"DeviceSite,omitempty"`
	DeviceDirection  *string `json:"DeviceDirection,omitempty" xml:"DeviceDirection,omitempty"`
	DeviceResolution *string `json:"DeviceResolution,omitempty" xml:"DeviceResolution,omitempty"`
	BitRate          *string `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Vendor           *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s AddDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceRequest) SetGbId(v string) *AddDeviceRequest {
	s.GbId = &v
	return s
}

func (s *AddDeviceRequest) SetDeviceName(v string) *AddDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *AddDeviceRequest) SetDeviceType(v string) *AddDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *AddDeviceRequest) SetDeviceAddress(v string) *AddDeviceRequest {
	s.DeviceAddress = &v
	return s
}

func (s *AddDeviceRequest) SetDeviceSite(v string) *AddDeviceRequest {
	s.DeviceSite = &v
	return s
}

func (s *AddDeviceRequest) SetDeviceDirection(v string) *AddDeviceRequest {
	s.DeviceDirection = &v
	return s
}

func (s *AddDeviceRequest) SetDeviceResolution(v string) *AddDeviceRequest {
	s.DeviceResolution = &v
	return s
}

func (s *AddDeviceRequest) SetBitRate(v string) *AddDeviceRequest {
	s.BitRate = &v
	return s
}

func (s *AddDeviceRequest) SetCorpId(v string) *AddDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *AddDeviceRequest) SetVendor(v string) *AddDeviceRequest {
	s.Vendor = &v
	return s
}

type AddDeviceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceResponseBody) SetMessage(v string) *AddDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *AddDeviceResponseBody) SetRequestId(v string) *AddDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDeviceResponseBody) SetData(v string) *AddDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *AddDeviceResponseBody) SetCode(v string) *AddDeviceResponseBody {
	s.Code = &v
	return s
}

type AddDeviceResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceResponse) SetHeaders(v map[string]*string) *AddDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceResponse) SetBody(v *AddDeviceResponseBody) *AddDeviceResponse {
	s.Body = v
	return s
}

type AddMonitorRequest struct {
	CorpId               *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	MonitorType          *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	BatchIndicator       *int32  `json:"BatchIndicator,omitempty" xml:"BatchIndicator,omitempty"`
	AlgorithmVendor      *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
	NotifierType         *string `json:"NotifierType,omitempty" xml:"NotifierType,omitempty"`
	NotifierUrl          *string `json:"NotifierUrl,omitempty" xml:"NotifierUrl,omitempty"`
	NotifierAppSecret    *string `json:"NotifierAppSecret,omitempty" xml:"NotifierAppSecret,omitempty"`
	NotifierTimeOut      *int32  `json:"NotifierTimeOut,omitempty" xml:"NotifierTimeOut,omitempty"`
	NotifierExtendValues *string `json:"NotifierExtendValues,omitempty" xml:"NotifierExtendValues,omitempty"`
}

func (s AddMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMonitorRequest) GoString() string {
	return s.String()
}

func (s *AddMonitorRequest) SetCorpId(v string) *AddMonitorRequest {
	s.CorpId = &v
	return s
}

func (s *AddMonitorRequest) SetMonitorType(v string) *AddMonitorRequest {
	s.MonitorType = &v
	return s
}

func (s *AddMonitorRequest) SetDescription(v string) *AddMonitorRequest {
	s.Description = &v
	return s
}

func (s *AddMonitorRequest) SetBatchIndicator(v int32) *AddMonitorRequest {
	s.BatchIndicator = &v
	return s
}

func (s *AddMonitorRequest) SetAlgorithmVendor(v string) *AddMonitorRequest {
	s.AlgorithmVendor = &v
	return s
}

func (s *AddMonitorRequest) SetNotifierType(v string) *AddMonitorRequest {
	s.NotifierType = &v
	return s
}

func (s *AddMonitorRequest) SetNotifierUrl(v string) *AddMonitorRequest {
	s.NotifierUrl = &v
	return s
}

func (s *AddMonitorRequest) SetNotifierAppSecret(v string) *AddMonitorRequest {
	s.NotifierAppSecret = &v
	return s
}

func (s *AddMonitorRequest) SetNotifierTimeOut(v int32) *AddMonitorRequest {
	s.NotifierTimeOut = &v
	return s
}

func (s *AddMonitorRequest) SetNotifierExtendValues(v string) *AddMonitorRequest {
	s.NotifierExtendValues = &v
	return s
}

type AddMonitorResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AddMonitorResponseBody) SetMessage(v string) *AddMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *AddMonitorResponseBody) SetRequestId(v string) *AddMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMonitorResponseBody) SetData(v *AddMonitorResponseBodyData) *AddMonitorResponseBody {
	s.Data = v
	return s
}

func (s *AddMonitorResponseBody) SetCode(v string) *AddMonitorResponseBody {
	s.Code = &v
	return s
}

type AddMonitorResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddMonitorResponseBodyData) SetTaskId(v string) *AddMonitorResponseBodyData {
	s.TaskId = &v
	return s
}

type AddMonitorResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddMonitorResponse) SetHeaders(v map[string]*string) *AddMonitorResponse {
	s.Headers = v
	return s
}

func (s *AddMonitorResponse) SetBody(v *AddMonitorResponseBody) *AddMonitorResponse {
	s.Body = v
	return s
}

type AddProfileRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender      *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s AddProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProfileRequest) GoString() string {
	return s.String()
}

func (s *AddProfileRequest) SetCorpId(v string) *AddProfileRequest {
	s.CorpId = &v
	return s
}

func (s *AddProfileRequest) SetIsvSubId(v string) *AddProfileRequest {
	s.IsvSubId = &v
	return s
}

func (s *AddProfileRequest) SetName(v string) *AddProfileRequest {
	s.Name = &v
	return s
}

func (s *AddProfileRequest) SetCatalogId(v int64) *AddProfileRequest {
	s.CatalogId = &v
	return s
}

func (s *AddProfileRequest) SetIdNumber(v string) *AddProfileRequest {
	s.IdNumber = &v
	return s
}

func (s *AddProfileRequest) SetFaceUrl(v string) *AddProfileRequest {
	s.FaceUrl = &v
	return s
}

func (s *AddProfileRequest) SetLiveAddress(v string) *AddProfileRequest {
	s.LiveAddress = &v
	return s
}

func (s *AddProfileRequest) SetGender(v int32) *AddProfileRequest {
	s.Gender = &v
	return s
}

func (s *AddProfileRequest) SetPlateNo(v string) *AddProfileRequest {
	s.PlateNo = &v
	return s
}

func (s *AddProfileRequest) SetPhoneNo(v string) *AddProfileRequest {
	s.PhoneNo = &v
	return s
}

func (s *AddProfileRequest) SetSceneType(v string) *AddProfileRequest {
	s.SceneType = &v
	return s
}

func (s *AddProfileRequest) SetBizId(v string) *AddProfileRequest {
	s.BizId = &v
	return s
}

type AddProfileResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponseBody) GoString() string {
	return s.String()
}

func (s *AddProfileResponseBody) SetMessage(v string) *AddProfileResponseBody {
	s.Message = &v
	return s
}

func (s *AddProfileResponseBody) SetRequestId(v string) *AddProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProfileResponseBody) SetData(v *AddProfileResponseBodyData) *AddProfileResponseBody {
	s.Data = v
	return s
}

func (s *AddProfileResponseBody) SetCode(v string) *AddProfileResponseBody {
	s.Code = &v
	return s
}

type AddProfileResponseBodyData struct {
	CatalogId   *int32  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ProfileId   *int32  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Gender      *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
}

func (s AddProfileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddProfileResponseBodyData) SetCatalogId(v int32) *AddProfileResponseBodyData {
	s.CatalogId = &v
	return s
}

func (s *AddProfileResponseBodyData) SetProfileId(v int32) *AddProfileResponseBodyData {
	s.ProfileId = &v
	return s
}

func (s *AddProfileResponseBodyData) SetIsvSubId(v string) *AddProfileResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *AddProfileResponseBodyData) SetGender(v string) *AddProfileResponseBodyData {
	s.Gender = &v
	return s
}

func (s *AddProfileResponseBodyData) SetBizId(v string) *AddProfileResponseBodyData {
	s.BizId = &v
	return s
}

func (s *AddProfileResponseBodyData) SetIdNumber(v string) *AddProfileResponseBodyData {
	s.IdNumber = &v
	return s
}

func (s *AddProfileResponseBodyData) SetSceneType(v string) *AddProfileResponseBodyData {
	s.SceneType = &v
	return s
}

func (s *AddProfileResponseBodyData) SetPhoneNo(v string) *AddProfileResponseBodyData {
	s.PhoneNo = &v
	return s
}

func (s *AddProfileResponseBodyData) SetFaceUrl(v string) *AddProfileResponseBodyData {
	s.FaceUrl = &v
	return s
}

func (s *AddProfileResponseBodyData) SetLiveAddress(v string) *AddProfileResponseBodyData {
	s.LiveAddress = &v
	return s
}

func (s *AddProfileResponseBodyData) SetName(v string) *AddProfileResponseBodyData {
	s.Name = &v
	return s
}

func (s *AddProfileResponseBodyData) SetPlateNo(v string) *AddProfileResponseBodyData {
	s.PlateNo = &v
	return s
}

type AddProfileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponse) GoString() string {
	return s.String()
}

func (s *AddProfileResponse) SetHeaders(v map[string]*string) *AddProfileResponse {
	s.Headers = v
	return s
}

func (s *AddProfileResponse) SetBody(v *AddProfileResponseBody) *AddProfileResponse {
	s.Body = v
	return s
}

type AddProfileCatalogRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId        *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	ParentCatalogId *int64  `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty"`
}

func (s AddProfileCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProfileCatalogRequest) GoString() string {
	return s.String()
}

func (s *AddProfileCatalogRequest) SetCorpId(v string) *AddProfileCatalogRequest {
	s.CorpId = &v
	return s
}

func (s *AddProfileCatalogRequest) SetIsvSubId(v string) *AddProfileCatalogRequest {
	s.IsvSubId = &v
	return s
}

func (s *AddProfileCatalogRequest) SetCatalogName(v string) *AddProfileCatalogRequest {
	s.CatalogName = &v
	return s
}

func (s *AddProfileCatalogRequest) SetParentCatalogId(v int64) *AddProfileCatalogRequest {
	s.ParentCatalogId = &v
	return s
}

type AddProfileCatalogResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddProfileCatalogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddProfileCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProfileCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *AddProfileCatalogResponseBody) SetMessage(v string) *AddProfileCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *AddProfileCatalogResponseBody) SetRequestId(v string) *AddProfileCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProfileCatalogResponseBody) SetData(v *AddProfileCatalogResponseBodyData) *AddProfileCatalogResponseBody {
	s.Data = v
	return s
}

func (s *AddProfileCatalogResponseBody) SetCode(v string) *AddProfileCatalogResponseBody {
	s.Code = &v
	return s
}

type AddProfileCatalogResponseBodyData struct {
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
}

func (s AddProfileCatalogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddProfileCatalogResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddProfileCatalogResponseBodyData) SetCatalogId(v int64) *AddProfileCatalogResponseBodyData {
	s.CatalogId = &v
	return s
}

func (s *AddProfileCatalogResponseBodyData) SetCatalogName(v string) *AddProfileCatalogResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *AddProfileCatalogResponseBodyData) SetIsvSubId(v string) *AddProfileCatalogResponseBodyData {
	s.IsvSubId = &v
	return s
}

type AddProfileCatalogResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProfileCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProfileCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProfileCatalogResponse) GoString() string {
	return s.String()
}

func (s *AddProfileCatalogResponse) SetHeaders(v map[string]*string) *AddProfileCatalogResponse {
	s.Headers = v
	return s
}

func (s *AddProfileCatalogResponse) SetBody(v *AddProfileCatalogResponseBody) *AddProfileCatalogResponse {
	s.Body = v
	return s
}

type BindCorpGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	CorpGroupId *string `json:"CorpGroupId,omitempty" xml:"CorpGroupId,omitempty"`
}

func (s BindCorpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s BindCorpGroupRequest) GoString() string {
	return s.String()
}

func (s *BindCorpGroupRequest) SetCorpId(v string) *BindCorpGroupRequest {
	s.CorpId = &v
	return s
}

func (s *BindCorpGroupRequest) SetCorpGroupId(v string) *BindCorpGroupRequest {
	s.CorpGroupId = &v
	return s
}

type BindCorpGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindCorpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindCorpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BindCorpGroupResponseBody) SetMessage(v string) *BindCorpGroupResponseBody {
	s.Message = &v
	return s
}

func (s *BindCorpGroupResponseBody) SetRequestId(v string) *BindCorpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindCorpGroupResponseBody) SetCode(v string) *BindCorpGroupResponseBody {
	s.Code = &v
	return s
}

func (s *BindCorpGroupResponseBody) SetSuccess(v bool) *BindCorpGroupResponseBody {
	s.Success = &v
	return s
}

type BindCorpGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindCorpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BindCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *BindCorpGroupResponse) SetHeaders(v map[string]*string) *BindCorpGroupResponse {
	s.Headers = v
	return s
}

func (s *BindCorpGroupResponse) SetBody(v *BindCorpGroupResponseBody) *BindCorpGroupResponse {
	s.Body = v
	return s
}

type BindPersonRequest struct {
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId           *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	PersonMatchingRate *string `json:"PersonMatchingRate,omitempty" xml:"PersonMatchingRate,omitempty"`
	PersonId           *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	ProfileId          *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
}

func (s BindPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s BindPersonRequest) GoString() string {
	return s.String()
}

func (s *BindPersonRequest) SetCorpId(v string) *BindPersonRequest {
	s.CorpId = &v
	return s
}

func (s *BindPersonRequest) SetIsvSubId(v string) *BindPersonRequest {
	s.IsvSubId = &v
	return s
}

func (s *BindPersonRequest) SetPersonMatchingRate(v string) *BindPersonRequest {
	s.PersonMatchingRate = &v
	return s
}

func (s *BindPersonRequest) SetPersonId(v string) *BindPersonRequest {
	s.PersonId = &v
	return s
}

func (s *BindPersonRequest) SetProfileId(v int64) *BindPersonRequest {
	s.ProfileId = &v
	return s
}

type BindPersonResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s BindPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindPersonResponseBody) GoString() string {
	return s.String()
}

func (s *BindPersonResponseBody) SetMessage(v string) *BindPersonResponseBody {
	s.Message = &v
	return s
}

func (s *BindPersonResponseBody) SetRequestId(v string) *BindPersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPersonResponseBody) SetData(v bool) *BindPersonResponseBody {
	s.Data = &v
	return s
}

func (s *BindPersonResponseBody) SetCode(v string) *BindPersonResponseBody {
	s.Code = &v
	return s
}

type BindPersonResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindPersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s BindPersonResponse) GoString() string {
	return s.String()
}

func (s *BindPersonResponse) SetHeaders(v map[string]*string) *BindPersonResponse {
	s.Headers = v
	return s
}

func (s *BindPersonResponse) SetBody(v *BindPersonResponseBody) *BindPersonResponse {
	s.Body = v
	return s
}

type BindUserRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	MatchingRate *string `json:"MatchingRate,omitempty" xml:"MatchingRate,omitempty"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BindUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindUserRequest) GoString() string {
	return s.String()
}

func (s *BindUserRequest) SetCorpId(v string) *BindUserRequest {
	s.CorpId = &v
	return s
}

func (s *BindUserRequest) SetIsvSubId(v string) *BindUserRequest {
	s.IsvSubId = &v
	return s
}

func (s *BindUserRequest) SetMatchingRate(v string) *BindUserRequest {
	s.MatchingRate = &v
	return s
}

func (s *BindUserRequest) SetPersonId(v string) *BindUserRequest {
	s.PersonId = &v
	return s
}

func (s *BindUserRequest) SetUserId(v int64) *BindUserRequest {
	s.UserId = &v
	return s
}

type BindUserResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s BindUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindUserResponseBody) SetMessage(v string) *BindUserResponseBody {
	s.Message = &v
	return s
}

func (s *BindUserResponseBody) SetRequestId(v string) *BindUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindUserResponseBody) SetData(v bool) *BindUserResponseBody {
	s.Data = &v
	return s
}

func (s *BindUserResponseBody) SetCode(v string) *BindUserResponseBody {
	s.Code = &v
	return s
}

type BindUserResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindUserResponse) GoString() string {
	return s.String()
}

func (s *BindUserResponse) SetHeaders(v map[string]*string) *BindUserResponse {
	s.Headers = v
	return s
}

func (s *BindUserResponse) SetBody(v *BindUserResponseBody) *BindUserResponse {
	s.Body = v
	return s
}

type CreateCorpRequest struct {
	CorpName      *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ParentCorpId  *string `json:"ParentCorpId,omitempty" xml:"ParentCorpId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	IsvSubId      *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	IconPath      *string `json:"IconPath,omitempty" xml:"IconPath,omitempty"`
}

func (s CreateCorpRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpRequest) GoString() string {
	return s.String()
}

func (s *CreateCorpRequest) SetCorpName(v string) *CreateCorpRequest {
	s.CorpName = &v
	return s
}

func (s *CreateCorpRequest) SetAppName(v string) *CreateCorpRequest {
	s.AppName = &v
	return s
}

func (s *CreateCorpRequest) SetParentCorpId(v string) *CreateCorpRequest {
	s.ParentCorpId = &v
	return s
}

func (s *CreateCorpRequest) SetDescription(v string) *CreateCorpRequest {
	s.Description = &v
	return s
}

func (s *CreateCorpRequest) SetAlgorithmType(v string) *CreateCorpRequest {
	s.AlgorithmType = &v
	return s
}

func (s *CreateCorpRequest) SetIsvSubId(v string) *CreateCorpRequest {
	s.IsvSubId = &v
	return s
}

func (s *CreateCorpRequest) SetIconPath(v string) *CreateCorpRequest {
	s.IconPath = &v
	return s
}

type CreateCorpResponseBody struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateCorpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCorpResponseBody) SetCorpId(v string) *CreateCorpResponseBody {
	s.CorpId = &v
	return s
}

func (s *CreateCorpResponseBody) SetMessage(v string) *CreateCorpResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCorpResponseBody) SetRequestId(v string) *CreateCorpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCorpResponseBody) SetCode(v string) *CreateCorpResponseBody {
	s.Code = &v
	return s
}

type CreateCorpResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCorpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCorpResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpResponse) GoString() string {
	return s.String()
}

func (s *CreateCorpResponse) SetHeaders(v map[string]*string) *CreateCorpResponse {
	s.Headers = v
	return s
}

func (s *CreateCorpResponse) SetBody(v *CreateCorpResponseBody) *CreateCorpResponse {
	s.Body = v
	return s
}

type CreateCorpGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateCorpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCorpGroupRequest) SetCorpId(v string) *CreateCorpGroupRequest {
	s.CorpId = &v
	return s
}

func (s *CreateCorpGroupRequest) SetGroupId(v string) *CreateCorpGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateCorpGroupRequest) SetClientToken(v string) *CreateCorpGroupRequest {
	s.ClientToken = &v
	return s
}

type CreateCorpGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCorpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCorpGroupResponseBody) SetMessage(v string) *CreateCorpGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCorpGroupResponseBody) SetRequestId(v string) *CreateCorpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCorpGroupResponseBody) SetCode(v string) *CreateCorpGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCorpGroupResponseBody) SetSuccess(v bool) *CreateCorpGroupResponseBody {
	s.Success = &v
	return s
}

type CreateCorpGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCorpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCorpGroupResponse) SetHeaders(v map[string]*string) *CreateCorpGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateCorpGroupResponse) SetBody(v *CreateCorpGroupResponseBody) *CreateCorpGroupResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId  *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Age          *int32  `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender       *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo      *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo      *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetCorpId(v string) *CreateUserRequest {
	s.CorpId = &v
	return s
}

func (s *CreateUserRequest) SetIsvSubId(v string) *CreateUserRequest {
	s.IsvSubId = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) SetUserGroupId(v int64) *CreateUserRequest {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserRequest) SetIdNumber(v string) *CreateUserRequest {
	s.IdNumber = &v
	return s
}

func (s *CreateUserRequest) SetFaceImageUrl(v string) *CreateUserRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *CreateUserRequest) SetAddress(v string) *CreateUserRequest {
	s.Address = &v
	return s
}

func (s *CreateUserRequest) SetAge(v int32) *CreateUserRequest {
	s.Age = &v
	return s
}

func (s *CreateUserRequest) SetGender(v int32) *CreateUserRequest {
	s.Gender = &v
	return s
}

func (s *CreateUserRequest) SetPlateNo(v string) *CreateUserRequest {
	s.PlateNo = &v
	return s
}

func (s *CreateUserRequest) SetPhoneNo(v string) *CreateUserRequest {
	s.PhoneNo = &v
	return s
}

func (s *CreateUserRequest) SetAttachment(v string) *CreateUserRequest {
	s.Attachment = &v
	return s
}

func (s *CreateUserRequest) SetBizId(v string) *CreateUserRequest {
	s.BizId = &v
	return s
}

type CreateUserResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

type CreateUserResponseBodyData struct {
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId  *int32  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId       *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	PhoneNo      *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	PlateNo      *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
}

func (s CreateUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyData) SetGender(v string) *CreateUserResponseBodyData {
	s.Gender = &v
	return s
}

func (s *CreateUserResponseBodyData) SetFaceImageUrl(v string) *CreateUserResponseBodyData {
	s.FaceImageUrl = &v
	return s
}

func (s *CreateUserResponseBodyData) SetIsvSubId(v string) *CreateUserResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetUserGroupId(v int32) *CreateUserResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetUserId(v int32) *CreateUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetBizId(v string) *CreateUserResponseBodyData {
	s.BizId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetAttachment(v string) *CreateUserResponseBodyData {
	s.Attachment = &v
	return s
}

func (s *CreateUserResponseBodyData) SetAge(v string) *CreateUserResponseBodyData {
	s.Age = &v
	return s
}

func (s *CreateUserResponseBodyData) SetIdNumber(v string) *CreateUserResponseBodyData {
	s.IdNumber = &v
	return s
}

func (s *CreateUserResponseBodyData) SetPhoneNo(v string) *CreateUserResponseBodyData {
	s.PhoneNo = &v
	return s
}

func (s *CreateUserResponseBodyData) SetAddress(v string) *CreateUserResponseBodyData {
	s.Address = &v
	return s
}

func (s *CreateUserResponseBodyData) SetUserName(v string) *CreateUserResponseBodyData {
	s.UserName = &v
	return s
}

func (s *CreateUserResponseBodyData) SetPlateNo(v string) *CreateUserResponseBodyData {
	s.PlateNo = &v
	return s
}

type CreateUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateUserGroupRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId          *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	ParentUserGroupId *int64  `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) SetCorpId(v string) *CreateUserGroupRequest {
	s.CorpId = &v
	return s
}

func (s *CreateUserGroupRequest) SetIsvSubId(v string) *CreateUserGroupRequest {
	s.IsvSubId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupRequest) SetParentUserGroupId(v int64) *CreateUserGroupRequest {
	s.ParentUserGroupId = &v
	return s
}

type CreateUserGroupResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) SetMessage(v string) *CreateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetData(v *CreateUserGroupResponseBodyData) *CreateUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserGroupResponseBody) SetCode(v string) *CreateUserGroupResponseBody {
	s.Code = &v
	return s
}

type CreateUserGroupResponseBodyData struct {
	IsvSubId      *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId   *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s CreateUserGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBodyData) SetIsvSubId(v string) *CreateUserGroupResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *CreateUserGroupResponseBodyData) SetUserGroupId(v int64) *CreateUserGroupResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserGroupResponseBodyData) SetUserGroupName(v string) *CreateUserGroupResponseBodyData {
	s.UserGroupName = &v
	return s
}

type CreateUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponse) SetHeaders(v map[string]*string) *CreateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupResponse) SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse {
	s.Body = v
	return s
}

type CreateVideoComposeTaskRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ImageFileNames  *string `json:"ImageFileNames,omitempty" xml:"ImageFileNames,omitempty"`
	AudioFileName   *string `json:"AudioFileName,omitempty" xml:"AudioFileName,omitempty"`
	ImageParameters *string `json:"ImageParameters,omitempty" xml:"ImageParameters,omitempty"`
	VideoFormat     *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	VideoFrameRate  *int32  `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty"`
}

func (s CreateVideoComposeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoComposeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoComposeTaskRequest) SetCorpId(v string) *CreateVideoComposeTaskRequest {
	s.CorpId = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetBucketName(v string) *CreateVideoComposeTaskRequest {
	s.BucketName = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetDomainName(v string) *CreateVideoComposeTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetImageFileNames(v string) *CreateVideoComposeTaskRequest {
	s.ImageFileNames = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetAudioFileName(v string) *CreateVideoComposeTaskRequest {
	s.AudioFileName = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetImageParameters(v string) *CreateVideoComposeTaskRequest {
	s.ImageParameters = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetVideoFormat(v string) *CreateVideoComposeTaskRequest {
	s.VideoFormat = &v
	return s
}

func (s *CreateVideoComposeTaskRequest) SetVideoFrameRate(v int32) *CreateVideoComposeTaskRequest {
	s.VideoFrameRate = &v
	return s
}

type CreateVideoComposeTaskResponseBody struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateVideoComposeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoComposeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoComposeTaskResponseBody) SetBucketName(v string) *CreateVideoComposeTaskResponseBody {
	s.BucketName = &v
	return s
}

func (s *CreateVideoComposeTaskResponseBody) SetMessage(v string) *CreateVideoComposeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVideoComposeTaskResponseBody) SetRequestId(v string) *CreateVideoComposeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoComposeTaskResponseBody) SetDomainName(v string) *CreateVideoComposeTaskResponseBody {
	s.DomainName = &v
	return s
}

func (s *CreateVideoComposeTaskResponseBody) SetCode(v string) *CreateVideoComposeTaskResponseBody {
	s.Code = &v
	return s
}

type CreateVideoComposeTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoComposeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoComposeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoComposeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoComposeTaskResponse) SetHeaders(v map[string]*string) *CreateVideoComposeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoComposeTaskResponse) SetBody(v *CreateVideoComposeTaskResponseBody) *CreateVideoComposeTaskResponse {
	s.Body = v
	return s
}

type CreateVideoSummaryTaskRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	StartTimeStamp   *int64  `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	EndTimeStamp     *int64  `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	OptionList       *string `json:"OptionList,omitempty" xml:"OptionList,omitempty"`
	LiveVideoSummary *string `json:"LiveVideoSummary,omitempty" xml:"LiveVideoSummary,omitempty"`
}

func (s CreateVideoSummaryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoSummaryTaskRequest) SetCorpId(v string) *CreateVideoSummaryTaskRequest {
	s.CorpId = &v
	return s
}

func (s *CreateVideoSummaryTaskRequest) SetDeviceId(v string) *CreateVideoSummaryTaskRequest {
	s.DeviceId = &v
	return s
}

func (s *CreateVideoSummaryTaskRequest) SetStartTimeStamp(v int64) *CreateVideoSummaryTaskRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *CreateVideoSummaryTaskRequest) SetEndTimeStamp(v int64) *CreateVideoSummaryTaskRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *CreateVideoSummaryTaskRequest) SetOptionList(v string) *CreateVideoSummaryTaskRequest {
	s.OptionList = &v
	return s
}

func (s *CreateVideoSummaryTaskRequest) SetLiveVideoSummary(v string) *CreateVideoSummaryTaskRequest {
	s.LiveVideoSummary = &v
	return s
}

type CreateVideoSummaryTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateVideoSummaryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoSummaryTaskResponseBody) SetMessage(v string) *CreateVideoSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVideoSummaryTaskResponseBody) SetRequestId(v string) *CreateVideoSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoSummaryTaskResponseBody) SetData(v string) *CreateVideoSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVideoSummaryTaskResponseBody) SetCode(v string) *CreateVideoSummaryTaskResponseBody {
	s.Code = &v
	return s
}

type CreateVideoSummaryTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateVideoSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoSummaryTaskResponse) SetBody(v *CreateVideoSummaryTaskResponseBody) *CreateVideoSummaryTaskResponse {
	s.Body = v
	return s
}

type DeleteChannelRequest struct {
	DeviceCodes *string `json:"DeviceCodes,omitempty" xml:"DeviceCodes,omitempty"`
}

func (s DeleteChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteChannelRequest) SetDeviceCodes(v string) *DeleteChannelRequest {
	s.DeviceCodes = &v
	return s
}

type DeleteChannelResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponseBody) SetRequestId(v string) *DeleteChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChannelResponseBody) SetCode(v string) *DeleteChannelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChannelResponseBody) SetMessage(v string) *DeleteChannelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChannelResponseBody) SetData(v string) *DeleteChannelResponseBody {
	s.Data = &v
	return s
}

type DeleteChannelResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponse) SetHeaders(v map[string]*string) *DeleteChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteChannelResponse) SetBody(v *DeleteChannelResponseBody) *DeleteChannelResponse {
	s.Body = v
	return s
}

type DeleteCorpGroupRequest struct {
	CorpId  *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteCorpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorpGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteCorpGroupRequest) SetCorpId(v string) *DeleteCorpGroupRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteCorpGroupRequest) SetGroupId(v string) *DeleteCorpGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteCorpGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCorpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCorpGroupResponseBody) SetMessage(v string) *DeleteCorpGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCorpGroupResponseBody) SetRequestId(v string) *DeleteCorpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCorpGroupResponseBody) SetCode(v string) *DeleteCorpGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCorpGroupResponseBody) SetSuccess(v bool) *DeleteCorpGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteCorpGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCorpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteCorpGroupResponse) SetHeaders(v map[string]*string) *DeleteCorpGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteCorpGroupResponse) SetBody(v *DeleteCorpGroupResponseBody) *DeleteCorpGroupResponse {
	s.Body = v
	return s
}

type DeleteDataSourceRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) SetCorpId(v string) *DeleteDataSourceRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetDataSourceId(v string) *DeleteDataSourceRequest {
	s.DataSourceId = &v
	return s
}

type DeleteDataSourceResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetMessage(v string) *DeleteDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetData(v string) *DeleteDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetCode(v string) *DeleteDataSourceResponseBody {
	s.Code = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type DeleteDeviceRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId   *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) SetCorpId(v string) *DeleteDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteDeviceRequest) SetGbId(v string) *DeleteDeviceRequest {
	s.GbId = &v
	return s
}

type DeleteDeviceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponseBody) SetMessage(v string) *DeleteDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDeviceResponseBody) SetRequestId(v string) *DeleteDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceResponseBody) SetData(v string) *DeleteDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDeviceResponseBody) SetCode(v string) *DeleteDeviceResponseBody {
	s.Code = &v
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

type DeleteDeviceForInstanceRequest struct {
	InstanceId         *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Devices            []*DeleteDeviceForInstanceRequestDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	ProjectId          *string                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AlgorithmId        *string                                  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	DeleteInstanceFlag *bool                                    `json:"DeleteInstanceFlag,omitempty" xml:"DeleteInstanceFlag,omitempty"`
	DeviceCount        *string                                  `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
}

func (s DeleteDeviceForInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceForInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceForInstanceRequest) SetInstanceId(v string) *DeleteDeviceForInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDeviceForInstanceRequest) SetDevices(v []*DeleteDeviceForInstanceRequestDevices) *DeleteDeviceForInstanceRequest {
	s.Devices = v
	return s
}

func (s *DeleteDeviceForInstanceRequest) SetProjectId(v string) *DeleteDeviceForInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDeviceForInstanceRequest) SetAlgorithmId(v string) *DeleteDeviceForInstanceRequest {
	s.AlgorithmId = &v
	return s
}

func (s *DeleteDeviceForInstanceRequest) SetDeleteInstanceFlag(v bool) *DeleteDeviceForInstanceRequest {
	s.DeleteInstanceFlag = &v
	return s
}

func (s *DeleteDeviceForInstanceRequest) SetDeviceCount(v string) *DeleteDeviceForInstanceRequest {
	s.DeviceCount = &v
	return s
}

type DeleteDeviceForInstanceRequestDevices struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDeviceForInstanceRequestDevices) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceForInstanceRequestDevices) GoString() string {
	return s.String()
}

func (s *DeleteDeviceForInstanceRequestDevices) SetDeviceId(v string) *DeleteDeviceForInstanceRequestDevices {
	s.DeviceId = &v
	return s
}

func (s *DeleteDeviceForInstanceRequestDevices) SetRegionId(v string) *DeleteDeviceForInstanceRequestDevices {
	s.RegionId = &v
	return s
}

type DeleteDeviceForInstanceShrinkRequest struct {
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DevicesShrink      *string `json:"Devices,omitempty" xml:"Devices,omitempty"`
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AlgorithmId        *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	DeleteInstanceFlag *bool   `json:"DeleteInstanceFlag,omitempty" xml:"DeleteInstanceFlag,omitempty"`
	DeviceCount        *string `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
}

func (s DeleteDeviceForInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceForInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceForInstanceShrinkRequest) SetInstanceId(v string) *DeleteDeviceForInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDeviceForInstanceShrinkRequest) SetDevicesShrink(v string) *DeleteDeviceForInstanceShrinkRequest {
	s.DevicesShrink = &v
	return s
}

func (s *DeleteDeviceForInstanceShrinkRequest) SetProjectId(v string) *DeleteDeviceForInstanceShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDeviceForInstanceShrinkRequest) SetAlgorithmId(v string) *DeleteDeviceForInstanceShrinkRequest {
	s.AlgorithmId = &v
	return s
}

func (s *DeleteDeviceForInstanceShrinkRequest) SetDeleteInstanceFlag(v bool) *DeleteDeviceForInstanceShrinkRequest {
	s.DeleteInstanceFlag = &v
	return s
}

func (s *DeleteDeviceForInstanceShrinkRequest) SetDeviceCount(v string) *DeleteDeviceForInstanceShrinkRequest {
	s.DeviceCount = &v
	return s
}

type DeleteDeviceForInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDeviceForInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceForInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceForInstanceResponseBody) SetRequestId(v string) *DeleteDeviceForInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceForInstanceResponseBody) SetMessage(v string) *DeleteDeviceForInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDeviceForInstanceResponseBody) SetCode(v string) *DeleteDeviceForInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDeviceForInstanceResponseBody) SetSuccess(v bool) *DeleteDeviceForInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteDeviceForInstanceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeviceForInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceForInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceForInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceForInstanceResponse) SetHeaders(v map[string]*string) *DeleteDeviceForInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceForInstanceResponse) SetBody(v *DeleteDeviceForInstanceResponseBody) *DeleteDeviceForInstanceResponse {
	s.Body = v
	return s
}

type DeleteIPCDeviceRequest struct {
	DeviceCodes *string `json:"DeviceCodes,omitempty" xml:"DeviceCodes,omitempty"`
}

func (s DeleteIPCDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIPCDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteIPCDeviceRequest) SetDeviceCodes(v string) *DeleteIPCDeviceRequest {
	s.DeviceCodes = &v
	return s
}

type DeleteIPCDeviceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteIPCDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIPCDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIPCDeviceResponseBody) SetRequestId(v string) *DeleteIPCDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIPCDeviceResponseBody) SetCode(v string) *DeleteIPCDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIPCDeviceResponseBody) SetMessage(v string) *DeleteIPCDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIPCDeviceResponseBody) SetData(v string) *DeleteIPCDeviceResponseBody {
	s.Data = &v
	return s
}

type DeleteIPCDeviceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIPCDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIPCDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIPCDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteIPCDeviceResponse) SetHeaders(v map[string]*string) *DeleteIPCDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteIPCDeviceResponse) SetBody(v *DeleteIPCDeviceResponseBody) *DeleteIPCDeviceResponse {
	s.Body = v
	return s
}

type DeleteNVRDeviceRequest struct {
	DeviceCodes *string `json:"DeviceCodes,omitempty" xml:"DeviceCodes,omitempty"`
}

func (s DeleteNVRDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNVRDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNVRDeviceRequest) SetDeviceCodes(v string) *DeleteNVRDeviceRequest {
	s.DeviceCodes = &v
	return s
}

type DeleteNVRDeviceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteNVRDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNVRDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNVRDeviceResponseBody) SetRequestId(v string) *DeleteNVRDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNVRDeviceResponseBody) SetCode(v string) *DeleteNVRDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNVRDeviceResponseBody) SetMessage(v string) *DeleteNVRDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNVRDeviceResponseBody) SetData(v string) *DeleteNVRDeviceResponseBody {
	s.Data = &v
	return s
}

type DeleteNVRDeviceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNVRDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNVRDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNVRDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNVRDeviceResponse) SetHeaders(v map[string]*string) *DeleteNVRDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNVRDeviceResponse) SetBody(v *DeleteNVRDeviceResponseBody) *DeleteNVRDeviceResponse {
	s.Body = v
	return s
}

type DeleteProfileRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	ProfileId *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
}

func (s DeleteProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteProfileRequest) SetCorpId(v string) *DeleteProfileRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteProfileRequest) SetIsvSubId(v string) *DeleteProfileRequest {
	s.IsvSubId = &v
	return s
}

func (s *DeleteProfileRequest) SetProfileId(v int64) *DeleteProfileRequest {
	s.ProfileId = &v
	return s
}

type DeleteProfileResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProfileResponseBody) SetMessage(v string) *DeleteProfileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProfileResponseBody) SetRequestId(v string) *DeleteProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProfileResponseBody) SetData(v bool) *DeleteProfileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteProfileResponseBody) SetCode(v string) *DeleteProfileResponseBody {
	s.Code = &v
	return s
}

type DeleteProfileResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteProfileResponse) SetHeaders(v map[string]*string) *DeleteProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteProfileResponse) SetBody(v *DeleteProfileResponseBody) *DeleteProfileResponse {
	s.Body = v
	return s
}

type DeleteProfileCatalogRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
}

func (s DeleteProfileCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileCatalogRequest) GoString() string {
	return s.String()
}

func (s *DeleteProfileCatalogRequest) SetCorpId(v string) *DeleteProfileCatalogRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteProfileCatalogRequest) SetIsvSubId(v string) *DeleteProfileCatalogRequest {
	s.IsvSubId = &v
	return s
}

func (s *DeleteProfileCatalogRequest) SetCatalogId(v string) *DeleteProfileCatalogRequest {
	s.CatalogId = &v
	return s
}

type DeleteProfileCatalogResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteProfileCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProfileCatalogResponseBody) SetMessage(v string) *DeleteProfileCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProfileCatalogResponseBody) SetRequestId(v string) *DeleteProfileCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProfileCatalogResponseBody) SetData(v bool) *DeleteProfileCatalogResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteProfileCatalogResponseBody) SetCode(v string) *DeleteProfileCatalogResponseBody {
	s.Code = &v
	return s
}

type DeleteProfileCatalogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProfileCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProfileCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileCatalogResponse) GoString() string {
	return s.String()
}

func (s *DeleteProfileCatalogResponse) SetHeaders(v map[string]*string) *DeleteProfileCatalogResponse {
	s.Headers = v
	return s
}

func (s *DeleteProfileCatalogResponse) SetBody(v *DeleteProfileCatalogResponseBody) *DeleteProfileCatalogResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// 项目id,多个以”,“隔开
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectIds(v string) *DeleteProjectRequest {
	s.ProjectIds = &v
	return s
}

type DeleteProjectResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetMessage(v string) *DeleteProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetCode(v string) *DeleteProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteRecordsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	OperatorType  *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordsRequest) SetCorpId(v string) *DeleteRecordsRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteRecordsRequest) SetAlgorithmType(v string) *DeleteRecordsRequest {
	s.AlgorithmType = &v
	return s
}

func (s *DeleteRecordsRequest) SetAttributeName(v string) *DeleteRecordsRequest {
	s.AttributeName = &v
	return s
}

func (s *DeleteRecordsRequest) SetOperatorType(v string) *DeleteRecordsRequest {
	s.OperatorType = &v
	return s
}

func (s *DeleteRecordsRequest) SetValue(v string) *DeleteRecordsRequest {
	s.Value = &v
	return s
}

type DeleteRecordsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordsResponseBody) SetMessage(v string) *DeleteRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRecordsResponseBody) SetRequestId(v string) *DeleteRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordsResponseBody) SetData(v string) *DeleteRecordsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRecordsResponseBody) SetCode(v string) *DeleteRecordsResponseBody {
	s.Code = &v
	return s
}

type DeleteRecordsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordsResponse) SetHeaders(v map[string]*string) *DeleteRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordsResponse) SetBody(v *DeleteRecordsResponseBody) *DeleteRecordsResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetCorpId(v string) *DeleteUserRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteUserRequest) SetIsvSubId(v string) *DeleteUserRequest {
	s.IsvSubId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v int64) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetData(v bool) *DeleteUserResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

type DeleteUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteUserGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) SetCorpId(v string) *DeleteUserGroupRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetIsvSubId(v string) *DeleteUserGroupRequest {
	s.IsvSubId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type DeleteUserGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) SetMessage(v string) *DeleteUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetData(v bool) *DeleteUserGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetCode(v string) *DeleteUserGroupResponseBody {
	s.Code = &v
	return s
}

type DeleteUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) SetHeaders(v map[string]*string) *DeleteUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

type DeleteVideoSummaryTaskRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteVideoSummaryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoSummaryTaskRequest) SetCorpId(v string) *DeleteVideoSummaryTaskRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteVideoSummaryTaskRequest) SetTaskId(v string) *DeleteVideoSummaryTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteVideoSummaryTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteVideoSummaryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoSummaryTaskResponseBody) SetMessage(v string) *DeleteVideoSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVideoSummaryTaskResponseBody) SetRequestId(v string) *DeleteVideoSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVideoSummaryTaskResponseBody) SetData(v string) *DeleteVideoSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVideoSummaryTaskResponseBody) SetCode(v string) *DeleteVideoSummaryTaskResponseBody {
	s.Code = &v
	return s
}

type DeleteVideoSummaryTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVideoSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVideoSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoSummaryTaskResponse) SetHeaders(v map[string]*string) *DeleteVideoSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoSummaryTaskResponse) SetBody(v *DeleteVideoSummaryTaskResponseBody) *DeleteVideoSummaryTaskResponse {
	s.Body = v
	return s
}

type DescribeDevicesRequest struct {
	PageNum    *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpIdList *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
}

func (s DescribeDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDevicesRequest) SetPageNum(v int32) *DescribeDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageSize(v int32) *DescribeDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesRequest) SetCorpIdList(v string) *DescribeDevicesRequest {
	s.CorpIdList = &v
	return s
}

type DescribeDevicesResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBody) SetMessage(v string) *DescribeDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetRequestId(v string) *DescribeDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetData(v *DescribeDevicesResponseBodyData) *DescribeDevicesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDevicesResponseBody) SetCode(v string) *DescribeDevicesResponseBody {
	s.Code = &v
	return s
}

type DescribeDevicesResponseBodyData struct {
	PageNum    *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	Records    []*DescribeDevicesResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyData) SetPageNum(v int32) *DescribeDevicesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetRecords(v []*DescribeDevicesResponseBodyDataRecords) *DescribeDevicesResponseBodyData {
	s.Records = v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetTotalPage(v int32) *DescribeDevicesResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetPageSize(v int32) *DescribeDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetTotalCount(v int32) *DescribeDevicesResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeDevicesResponseBodyDataRecords struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceAddress *string `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Longitude     *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	InProtocol    *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	Latitude      *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Vendor        *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeDevicesResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDataRecords) SetStatus(v string) *DescribeDevicesResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetDeviceName(v string) *DescribeDevicesResponseBodyDataRecords {
	s.DeviceName = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetDeviceType(v string) *DescribeDevicesResponseBodyDataRecords {
	s.DeviceType = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetDeviceId(v string) *DescribeDevicesResponseBodyDataRecords {
	s.DeviceId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetDeviceAddress(v string) *DescribeDevicesResponseBodyDataRecords {
	s.DeviceAddress = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetCreateTime(v string) *DescribeDevicesResponseBodyDataRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetCorpId(v string) *DescribeDevicesResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetLongitude(v string) *DescribeDevicesResponseBodyDataRecords {
	s.Longitude = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetInProtocol(v string) *DescribeDevicesResponseBodyDataRecords {
	s.InProtocol = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetLatitude(v string) *DescribeDevicesResponseBodyDataRecords {
	s.Latitude = &v
	return s
}

func (s *DescribeDevicesResponseBodyDataRecords) SetVendor(v string) *DescribeDevicesResponseBodyDataRecords {
	s.Vendor = &v
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

type GetBodyOptionsRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s GetBodyOptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsRequest) SetCorpId(v string) *GetBodyOptionsRequest {
	s.CorpId = &v
	return s
}

type GetBodyOptionsResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetBodyOptionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetBodyOptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponseBody) SetMessage(v string) *GetBodyOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetBodyOptionsResponseBody) SetRequestId(v string) *GetBodyOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBodyOptionsResponseBody) SetData(v []*GetBodyOptionsResponseBodyData) *GetBodyOptionsResponseBody {
	s.Data = v
	return s
}

func (s *GetBodyOptionsResponseBody) SetCode(v string) *GetBodyOptionsResponseBody {
	s.Code = &v
	return s
}

type GetBodyOptionsResponseBodyData struct {
	Key        *string                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	OptionList []*GetBodyOptionsResponseBodyDataOptionList `json:"OptionList,omitempty" xml:"OptionList,omitempty" type:"Repeated"`
	Name       *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetBodyOptionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponseBodyData) SetKey(v string) *GetBodyOptionsResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetBodyOptionsResponseBodyData) SetOptionList(v []*GetBodyOptionsResponseBodyDataOptionList) *GetBodyOptionsResponseBodyData {
	s.OptionList = v
	return s
}

func (s *GetBodyOptionsResponseBodyData) SetName(v string) *GetBodyOptionsResponseBodyData {
	s.Name = &v
	return s
}

type GetBodyOptionsResponseBodyDataOptionList struct {
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetBodyOptionsResponseBodyDataOptionList) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponseBodyDataOptionList) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponseBodyDataOptionList) SetKey(v string) *GetBodyOptionsResponseBodyDataOptionList {
	s.Key = &v
	return s
}

func (s *GetBodyOptionsResponseBodyDataOptionList) SetName(v string) *GetBodyOptionsResponseBodyDataOptionList {
	s.Name = &v
	return s
}

type GetBodyOptionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBodyOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBodyOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponse) SetHeaders(v map[string]*string) *GetBodyOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetBodyOptionsResponse) SetBody(v *GetBodyOptionsResponseBody) *GetBodyOptionsResponse {
	s.Body = v
	return s
}

type GetCatalogListRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
}

func (s GetCatalogListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogListRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogListRequest) SetCorpId(v string) *GetCatalogListRequest {
	s.CorpId = &v
	return s
}

func (s *GetCatalogListRequest) SetIsvSubId(v string) *GetCatalogListRequest {
	s.IsvSubId = &v
	return s
}

type GetCatalogListResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetCatalogListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetCatalogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogListResponseBody) SetMessage(v string) *GetCatalogListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCatalogListResponseBody) SetRequestId(v string) *GetCatalogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogListResponseBody) SetData(v []*GetCatalogListResponseBodyData) *GetCatalogListResponseBody {
	s.Data = v
	return s
}

func (s *GetCatalogListResponseBody) SetCode(v string) *GetCatalogListResponseBody {
	s.Code = &v
	return s
}

type GetCatalogListResponseBodyData struct {
	CatalogId       *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	IsvSubId        *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	ParentCatalogId *int64  `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty"`
	ProfileCount    *int64  `json:"ProfileCount,omitempty" xml:"ProfileCount,omitempty"`
}

func (s GetCatalogListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCatalogListResponseBodyData) SetCatalogId(v int64) *GetCatalogListResponseBodyData {
	s.CatalogId = &v
	return s
}

func (s *GetCatalogListResponseBodyData) SetCatalogName(v string) *GetCatalogListResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *GetCatalogListResponseBodyData) SetIsvSubId(v string) *GetCatalogListResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *GetCatalogListResponseBodyData) SetParentCatalogId(v int64) *GetCatalogListResponseBodyData {
	s.ParentCatalogId = &v
	return s
}

func (s *GetCatalogListResponseBodyData) SetProfileCount(v int64) *GetCatalogListResponseBodyData {
	s.ProfileCount = &v
	return s
}

type GetCatalogListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCatalogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCatalogListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogListResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogListResponse) SetHeaders(v map[string]*string) *GetCatalogListResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogListResponse) SetBody(v *GetCatalogListResponseBody) *GetCatalogListResponse {
	s.Body = v
	return s
}

type GetDeviceCaptureStrategyRequest struct {
	// 设备通道号
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	// 设备类型
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
}

func (s GetDeviceCaptureStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCaptureStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceCaptureStrategyRequest) SetDeviceCode(v string) *GetDeviceCaptureStrategyRequest {
	s.DeviceCode = &v
	return s
}

func (s *GetDeviceCaptureStrategyRequest) SetDeviceType(v string) *GetDeviceCaptureStrategyRequest {
	s.DeviceType = &v
	return s
}

type GetDeviceCaptureStrategyResponseBody struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 响应数据内容
	Data *GetDeviceCaptureStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceCaptureStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCaptureStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceCaptureStrategyResponseBody) SetCode(v string) *GetDeviceCaptureStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBody) SetMessage(v string) *GetDeviceCaptureStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBody) SetData(v *GetDeviceCaptureStrategyResponseBodyData) *GetDeviceCaptureStrategyResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBody) SetRequestId(v string) *GetDeviceCaptureStrategyResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceCaptureStrategyResponseBodyData struct {
	// 设备类型
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// 设备通道
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	// 星期一抓取策略
	MondayCaptureStrategy *string `json:"MondayCaptureStrategy,omitempty" xml:"MondayCaptureStrategy,omitempty"`
	// 星期二抓取策略
	TuesdayCaptureStrategy *string `json:"TuesdayCaptureStrategy,omitempty" xml:"TuesdayCaptureStrategy,omitempty"`
	// 星期三抓取策略
	WednesdayCaptureStrategy *string `json:"WednesdayCaptureStrategy,omitempty" xml:"WednesdayCaptureStrategy,omitempty"`
	// 星期四抓取策略
	ThursdayCaptureStrategy *string `json:"ThursdayCaptureStrategy,omitempty" xml:"ThursdayCaptureStrategy,omitempty"`
	// 星期五抓取策略
	FridayCaptureStrategy *string `json:"FridayCaptureStrategy,omitempty" xml:"FridayCaptureStrategy,omitempty"`
	// 星期六抓取策略
	SaturdayCaptureStrategy *string `json:"SaturdayCaptureStrategy,omitempty" xml:"SaturdayCaptureStrategy,omitempty"`
	// 星期日抓取策略
	SundayCaptureStrategy *string `json:"SundayCaptureStrategy,omitempty" xml:"SundayCaptureStrategy,omitempty"`
}

func (s GetDeviceCaptureStrategyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCaptureStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetDeviceType(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetDeviceCode(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.DeviceCode = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetMondayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.MondayCaptureStrategy = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetTuesdayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.TuesdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetWednesdayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.WednesdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetThursdayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.ThursdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetFridayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.FridayCaptureStrategy = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetSaturdayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.SaturdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceCaptureStrategyResponseBodyData) SetSundayCaptureStrategy(v string) *GetDeviceCaptureStrategyResponseBodyData {
	s.SundayCaptureStrategy = &v
	return s
}

type GetDeviceCaptureStrategyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceCaptureStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceCaptureStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCaptureStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceCaptureStrategyResponse) SetHeaders(v map[string]*string) *GetDeviceCaptureStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceCaptureStrategyResponse) SetBody(v *GetDeviceCaptureStrategyResponseBody) *GetDeviceCaptureStrategyResponse {
	s.Body = v
	return s
}

type GetDeviceConfigRequest struct {
	DeviceSn        *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	DeviceTimeStamp *string `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty"`
}

func (s GetDeviceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigRequest) SetDeviceSn(v string) *GetDeviceConfigRequest {
	s.DeviceSn = &v
	return s
}

func (s *GetDeviceConfigRequest) SetDeviceTimeStamp(v string) *GetDeviceConfigRequest {
	s.DeviceTimeStamp = &v
	return s
}

type GetDeviceConfigResponseBody struct {
	// Id of the request
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message       *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RetryInterval *string                                   `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	Code          *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	AudioEnable   *string                                   `json:"AudioEnable,omitempty" xml:"AudioEnable,omitempty"`
	AudioFormat   *string                                   `json:"AudioFormat,omitempty" xml:"AudioFormat,omitempty"`
	BitRate       *string                                   `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	DeviceAddress *string                                   `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty"`
	DeviceName    *string                                   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EncodeFormat  *string                                   `json:"EncodeFormat,omitempty" xml:"EncodeFormat,omitempty"`
	FrameRate     *string                                   `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	GovLength     *int64                                    `json:"GovLength,omitempty" xml:"GovLength,omitempty"`
	Latitude      *string                                   `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude     *string                                   `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	OSDList       []*GetDeviceConfigResponseBodyOSDList     `json:"OSDList,omitempty" xml:"OSDList,omitempty" type:"Repeated"`
	OSDTimeEnable *string                                   `json:"OSDTimeEnable,omitempty" xml:"OSDTimeEnable,omitempty"`
	OSDTimeType   *string                                   `json:"OSDTimeType,omitempty" xml:"OSDTimeType,omitempty"`
	OSDTimeX      *string                                   `json:"OSDTimeX,omitempty" xml:"OSDTimeX,omitempty"`
	OSDTimeY      *string                                   `json:"OSDTimeY,omitempty" xml:"OSDTimeY,omitempty"`
	Resolution    *string                                   `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	DeviceId      *string                                   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserName      *string                                   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	PassWord      *string                                   `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	Protocol      *string                                   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ServerId      *string                                   `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerPort    *string                                   `json:"ServerPort,omitempty" xml:"ServerPort,omitempty"`
	ServerIp      *string                                   `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	ChannelList   []*GetDeviceConfigResponseBodyChannelList `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Repeated"`
}

func (s GetDeviceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponseBody) SetRequestId(v string) *GetDeviceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetMessage(v string) *GetDeviceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetRetryInterval(v string) *GetDeviceConfigResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetCode(v string) *GetDeviceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetAudioEnable(v string) *GetDeviceConfigResponseBody {
	s.AudioEnable = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetAudioFormat(v string) *GetDeviceConfigResponseBody {
	s.AudioFormat = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetBitRate(v string) *GetDeviceConfigResponseBody {
	s.BitRate = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetDeviceAddress(v string) *GetDeviceConfigResponseBody {
	s.DeviceAddress = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetDeviceName(v string) *GetDeviceConfigResponseBody {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetEncodeFormat(v string) *GetDeviceConfigResponseBody {
	s.EncodeFormat = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetFrameRate(v string) *GetDeviceConfigResponseBody {
	s.FrameRate = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetGovLength(v int64) *GetDeviceConfigResponseBody {
	s.GovLength = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetLatitude(v string) *GetDeviceConfigResponseBody {
	s.Latitude = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetLongitude(v string) *GetDeviceConfigResponseBody {
	s.Longitude = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetOSDList(v []*GetDeviceConfigResponseBodyOSDList) *GetDeviceConfigResponseBody {
	s.OSDList = v
	return s
}

func (s *GetDeviceConfigResponseBody) SetOSDTimeEnable(v string) *GetDeviceConfigResponseBody {
	s.OSDTimeEnable = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetOSDTimeType(v string) *GetDeviceConfigResponseBody {
	s.OSDTimeType = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetOSDTimeX(v string) *GetDeviceConfigResponseBody {
	s.OSDTimeX = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetOSDTimeY(v string) *GetDeviceConfigResponseBody {
	s.OSDTimeY = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetResolution(v string) *GetDeviceConfigResponseBody {
	s.Resolution = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetDeviceId(v string) *GetDeviceConfigResponseBody {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetUserName(v string) *GetDeviceConfigResponseBody {
	s.UserName = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetPassWord(v string) *GetDeviceConfigResponseBody {
	s.PassWord = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetProtocol(v string) *GetDeviceConfigResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetServerId(v string) *GetDeviceConfigResponseBody {
	s.ServerId = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetServerPort(v string) *GetDeviceConfigResponseBody {
	s.ServerPort = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetServerIp(v string) *GetDeviceConfigResponseBody {
	s.ServerIp = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetChannelList(v []*GetDeviceConfigResponseBodyChannelList) *GetDeviceConfigResponseBody {
	s.ChannelList = v
	return s
}

type GetDeviceConfigResponseBodyOSDList struct {
	LeftTopX *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	LeftTopY *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Text     *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDeviceConfigResponseBodyOSDList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponseBodyOSDList) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponseBodyOSDList) SetLeftTopX(v string) *GetDeviceConfigResponseBodyOSDList {
	s.LeftTopX = &v
	return s
}

func (s *GetDeviceConfigResponseBodyOSDList) SetLeftTopY(v string) *GetDeviceConfigResponseBodyOSDList {
	s.LeftTopY = &v
	return s
}

func (s *GetDeviceConfigResponseBodyOSDList) SetText(v string) *GetDeviceConfigResponseBodyOSDList {
	s.Text = &v
	return s
}

type GetDeviceConfigResponseBodyChannelList struct {
	ChannelGbId              *string `json:"ChannelGbId,omitempty" xml:"ChannelGbId,omitempty"`
	MondayCaptureStrategy    *string `json:"MondayCaptureStrategy,omitempty" xml:"MondayCaptureStrategy,omitempty"`
	TuesdayCaptureStrategy   *string `json:"TuesdayCaptureStrategy,omitempty" xml:"TuesdayCaptureStrategy,omitempty"`
	WednesdayCaptureStrategy *string `json:"WednesdayCaptureStrategy,omitempty" xml:"WednesdayCaptureStrategy,omitempty"`
	ThursdayCaptureStrategy  *string `json:"ThursdayCaptureStrategy,omitempty" xml:"ThursdayCaptureStrategy,omitempty"`
	FridayCaptureStrategy    *string `json:"FridayCaptureStrategy,omitempty" xml:"FridayCaptureStrategy,omitempty"`
	SaturdayCaptureStrategy  *string `json:"SaturdayCaptureStrategy,omitempty" xml:"SaturdayCaptureStrategy,omitempty"`
	SundayCaptureStrategy    *string `json:"SundayCaptureStrategy,omitempty" xml:"SundayCaptureStrategy,omitempty"`
}

func (s GetDeviceConfigResponseBodyChannelList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponseBodyChannelList) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponseBodyChannelList) SetChannelGbId(v string) *GetDeviceConfigResponseBodyChannelList {
	s.ChannelGbId = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetMondayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.MondayCaptureStrategy = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetTuesdayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.TuesdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetWednesdayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.WednesdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetThursdayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.ThursdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetFridayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.FridayCaptureStrategy = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetSaturdayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.SaturdayCaptureStrategy = &v
	return s
}

func (s *GetDeviceConfigResponseBodyChannelList) SetSundayCaptureStrategy(v string) *GetDeviceConfigResponseBodyChannelList {
	s.SundayCaptureStrategy = &v
	return s
}

type GetDeviceConfigResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponse) SetHeaders(v map[string]*string) *GetDeviceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceConfigResponse) SetBody(v *GetDeviceConfigResponseBody) *GetDeviceConfigResponse {
	s.Body = v
	return s
}

type GetDeviceLiveUrlRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId        *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
}

func (s GetDeviceLiveUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceLiveUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceLiveUrlRequest) SetDeviceId(v string) *GetDeviceLiveUrlRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceLiveUrlRequest) SetStreamType(v int32) *GetDeviceLiveUrlRequest {
	s.StreamType = &v
	return s
}

func (s *GetDeviceLiveUrlRequest) SetOutProtocol(v string) *GetDeviceLiveUrlRequest {
	s.OutProtocol = &v
	return s
}

func (s *GetDeviceLiveUrlRequest) SetCorpId(v string) *GetDeviceLiveUrlRequest {
	s.CorpId = &v
	return s
}

func (s *GetDeviceLiveUrlRequest) SetGbId(v string) *GetDeviceLiveUrlRequest {
	s.GbId = &v
	return s
}

type GetDeviceLiveUrlResponseBody struct {
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDeviceLiveUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceLiveUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceLiveUrlResponseBody) SetStreamType(v int32) *GetDeviceLiveUrlResponseBody {
	s.StreamType = &v
	return s
}

func (s *GetDeviceLiveUrlResponseBody) SetRequestId(v string) *GetDeviceLiveUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceLiveUrlResponseBody) SetMessage(v string) *GetDeviceLiveUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceLiveUrlResponseBody) SetCode(v string) *GetDeviceLiveUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceLiveUrlResponseBody) SetOutProtocol(v string) *GetDeviceLiveUrlResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *GetDeviceLiveUrlResponseBody) SetUrl(v string) *GetDeviceLiveUrlResponseBody {
	s.Url = &v
	return s
}

type GetDeviceLiveUrlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceLiveUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceLiveUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceLiveUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceLiveUrlResponse) SetHeaders(v map[string]*string) *GetDeviceLiveUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceLiveUrlResponse) SetBody(v *GetDeviceLiveUrlResponseBody) *GetDeviceLiveUrlResponse {
	s.Body = v
	return s
}

type GetDeviceVideoUrlRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId        *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
}

func (s GetDeviceVideoUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceVideoUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceVideoUrlRequest) SetCorpId(v string) *GetDeviceVideoUrlRequest {
	s.CorpId = &v
	return s
}

func (s *GetDeviceVideoUrlRequest) SetGbId(v string) *GetDeviceVideoUrlRequest {
	s.GbId = &v
	return s
}

func (s *GetDeviceVideoUrlRequest) SetStartTime(v int64) *GetDeviceVideoUrlRequest {
	s.StartTime = &v
	return s
}

func (s *GetDeviceVideoUrlRequest) SetEndTime(v int64) *GetDeviceVideoUrlRequest {
	s.EndTime = &v
	return s
}

func (s *GetDeviceVideoUrlRequest) SetDeviceId(v string) *GetDeviceVideoUrlRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceVideoUrlRequest) SetOutProtocol(v string) *GetDeviceVideoUrlRequest {
	s.OutProtocol = &v
	return s
}

type GetDeviceVideoUrlResponseBody struct {
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDeviceVideoUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceVideoUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceVideoUrlResponseBody) SetMessage(v string) *GetDeviceVideoUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceVideoUrlResponseBody) SetRequestId(v string) *GetDeviceVideoUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceVideoUrlResponseBody) SetCode(v string) *GetDeviceVideoUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceVideoUrlResponseBody) SetOutProtocol(v string) *GetDeviceVideoUrlResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *GetDeviceVideoUrlResponseBody) SetUrl(v string) *GetDeviceVideoUrlResponseBody {
	s.Url = &v
	return s
}

type GetDeviceVideoUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceVideoUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceVideoUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceVideoUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceVideoUrlResponse) SetHeaders(v map[string]*string) *GetDeviceVideoUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceVideoUrlResponse) SetBody(v *GetDeviceVideoUrlResponseBody) *GetDeviceVideoUrlResponse {
	s.Body = v
	return s
}

type GetFaceModelResultRequest struct {
	PictureId      *string `json:"PictureId,omitempty" xml:"PictureId,omitempty"`
	PictureContent *string `json:"PictureContent,omitempty" xml:"PictureContent,omitempty"`
	PictureUrl     *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
}

func (s GetFaceModelResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultRequest) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultRequest) SetPictureId(v string) *GetFaceModelResultRequest {
	s.PictureId = &v
	return s
}

func (s *GetFaceModelResultRequest) SetPictureContent(v string) *GetFaceModelResultRequest {
	s.PictureContent = &v
	return s
}

func (s *GetFaceModelResultRequest) SetPictureUrl(v string) *GetFaceModelResultRequest {
	s.PictureUrl = &v
	return s
}

type GetFaceModelResultResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetFaceModelResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetFaceModelResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponseBody) SetMessage(v string) *GetFaceModelResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetFaceModelResultResponseBody) SetRequestId(v string) *GetFaceModelResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFaceModelResultResponseBody) SetData(v *GetFaceModelResultResponseBodyData) *GetFaceModelResultResponseBody {
	s.Data = v
	return s
}

func (s *GetFaceModelResultResponseBody) SetCode(v string) *GetFaceModelResultResponseBody {
	s.Code = &v
	return s
}

type GetFaceModelResultResponseBodyData struct {
	Records []*GetFaceModelResultResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s GetFaceModelResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponseBodyData) SetRecords(v []*GetFaceModelResultResponseBodyDataRecords) *GetFaceModelResultResponseBodyData {
	s.Records = v
	return s
}

type GetFaceModelResultResponseBodyDataRecords struct {
	RightBottomY               *float32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	HairColorReliability       *string    `json:"HairColorReliability,omitempty" xml:"HairColorReliability,omitempty"`
	HairColor                  *int32     `json:"HairColor,omitempty" xml:"HairColor,omitempty"`
	FaceStyle                  *string    `json:"FaceStyle,omitempty" xml:"FaceStyle,omitempty"`
	GlassStyleReliability      *string    `json:"GlassStyleReliability,omitempty" xml:"GlassStyleReliability,omitempty"`
	MustacheStyleReliability   *string    `json:"MustacheStyleReliability,omitempty" xml:"MustacheStyleReliability,omitempty"`
	RespiratorColorReliability *string    `json:"RespiratorColorReliability,omitempty" xml:"RespiratorColorReliability,omitempty"`
	RightBottomX               *float32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	AgeUpLimit                 *int32     `json:"AgeUpLimit,omitempty" xml:"AgeUpLimit,omitempty"`
	AgeUpLimitReliability      *string    `json:"AgeUpLimitReliability,omitempty" xml:"AgeUpLimitReliability,omitempty"`
	HairStyle                  *int32     `json:"HairStyle,omitempty" xml:"HairStyle,omitempty"`
	AgeLowerLimit              *int32     `json:"AgeLowerLimit,omitempty" xml:"AgeLowerLimit,omitempty"`
	LeftTopY                   *float32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	FeatureData                []*float32 `json:"FeatureData,omitempty" xml:"FeatureData,omitempty" type:"Repeated"`
	SkinColorReliability       *string    `json:"SkinColorReliability,omitempty" xml:"SkinColorReliability,omitempty"`
	CapColor                   *int32     `json:"CapColor,omitempty" xml:"CapColor,omitempty"`
	FaceStyleReliability       *string    `json:"FaceStyleReliability,omitempty" xml:"FaceStyleReliability,omitempty"`
	CapStyleReliability        *string    `json:"CapStyleReliability,omitempty" xml:"CapStyleReliability,omitempty"`
	GenderCodeReliability      *string    `json:"GenderCodeReliability,omitempty" xml:"GenderCodeReliability,omitempty"`
	HairStyleReliability       *string    `json:"HairStyleReliability,omitempty" xml:"HairStyleReliability,omitempty"`
	GlassColorReliability      *string    `json:"GlassColorReliability,omitempty" xml:"GlassColorReliability,omitempty"`
	EthicCode                  *int32     `json:"EthicCode,omitempty" xml:"EthicCode,omitempty"`
	RespiratorColor            *int32     `json:"RespiratorColor,omitempty" xml:"RespiratorColor,omitempty"`
	MustacheStyle              *string    `json:"MustacheStyle,omitempty" xml:"MustacheStyle,omitempty"`
	GlassColor                 *int32     `json:"GlassColor,omitempty" xml:"GlassColor,omitempty"`
	GlassStyle                 *int32     `json:"GlassStyle,omitempty" xml:"GlassStyle,omitempty"`
	SkinColor                  *int32     `json:"SkinColor,omitempty" xml:"SkinColor,omitempty"`
	CapColorReliability        *string    `json:"CapColorReliability,omitempty" xml:"CapColorReliability,omitempty"`
	CapStyle                   *int32     `json:"CapStyle,omitempty" xml:"CapStyle,omitempty"`
	GenderCode                 *int32     `json:"GenderCode,omitempty" xml:"GenderCode,omitempty"`
	LeftTopX                   *float32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	AgeLowerLimitReliability   *string    `json:"AgeLowerLimitReliability,omitempty" xml:"AgeLowerLimitReliability,omitempty"`
	EthicCodeReliability       *string    `json:"EthicCodeReliability,omitempty" xml:"EthicCodeReliability,omitempty"`
}

func (s GetFaceModelResultResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetRightBottomY(v float32) *GetFaceModelResultResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetHairColorReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.HairColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetHairColor(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.HairColor = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetFaceStyle(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.FaceStyle = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetGlassStyleReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.GlassStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetMustacheStyleReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.MustacheStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetRespiratorColorReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.RespiratorColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetRightBottomX(v float32) *GetFaceModelResultResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetAgeUpLimit(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.AgeUpLimit = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetAgeUpLimitReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.AgeUpLimitReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetHairStyle(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.HairStyle = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetAgeLowerLimit(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.AgeLowerLimit = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetLeftTopY(v float32) *GetFaceModelResultResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetFeatureData(v []*float32) *GetFaceModelResultResponseBodyDataRecords {
	s.FeatureData = v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetSkinColorReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.SkinColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetCapColor(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.CapColor = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetFaceStyleReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.FaceStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetCapStyleReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.CapStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetGenderCodeReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.GenderCodeReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetHairStyleReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.HairStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetGlassColorReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.GlassColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetEthicCode(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.EthicCode = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetRespiratorColor(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.RespiratorColor = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetMustacheStyle(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.MustacheStyle = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetGlassColor(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.GlassColor = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetGlassStyle(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.GlassStyle = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetSkinColor(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.SkinColor = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetCapColorReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.CapColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetCapStyle(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.CapStyle = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetGenderCode(v int32) *GetFaceModelResultResponseBodyDataRecords {
	s.GenderCode = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetLeftTopX(v float32) *GetFaceModelResultResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetAgeLowerLimitReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.AgeLowerLimitReliability = &v
	return s
}

func (s *GetFaceModelResultResponseBodyDataRecords) SetEthicCodeReliability(v string) *GetFaceModelResultResponseBodyDataRecords {
	s.EthicCodeReliability = &v
	return s
}

type GetFaceModelResultResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFaceModelResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFaceModelResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponse) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponse) SetHeaders(v map[string]*string) *GetFaceModelResultResponse {
	s.Headers = v
	return s
}

func (s *GetFaceModelResultResponse) SetBody(v *GetFaceModelResultResponseBody) *GetFaceModelResultResponse {
	s.Body = v
	return s
}

type GetFaceOptionsRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s GetFaceOptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsRequest) SetCorpId(v string) *GetFaceOptionsRequest {
	s.CorpId = &v
	return s
}

type GetFaceOptionsResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetFaceOptionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetFaceOptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponseBody) SetMessage(v string) *GetFaceOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetFaceOptionsResponseBody) SetRequestId(v string) *GetFaceOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFaceOptionsResponseBody) SetData(v []*GetFaceOptionsResponseBodyData) *GetFaceOptionsResponseBody {
	s.Data = v
	return s
}

func (s *GetFaceOptionsResponseBody) SetCode(v string) *GetFaceOptionsResponseBody {
	s.Code = &v
	return s
}

type GetFaceOptionsResponseBodyData struct {
	Key        *string                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	OptionList []*GetFaceOptionsResponseBodyDataOptionList `json:"OptionList,omitempty" xml:"OptionList,omitempty" type:"Repeated"`
	Name       *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetFaceOptionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponseBodyData) SetKey(v string) *GetFaceOptionsResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetFaceOptionsResponseBodyData) SetOptionList(v []*GetFaceOptionsResponseBodyDataOptionList) *GetFaceOptionsResponseBodyData {
	s.OptionList = v
	return s
}

func (s *GetFaceOptionsResponseBodyData) SetName(v string) *GetFaceOptionsResponseBodyData {
	s.Name = &v
	return s
}

type GetFaceOptionsResponseBodyDataOptionList struct {
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetFaceOptionsResponseBodyDataOptionList) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponseBodyDataOptionList) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponseBodyDataOptionList) SetKey(v string) *GetFaceOptionsResponseBodyDataOptionList {
	s.Key = &v
	return s
}

func (s *GetFaceOptionsResponseBodyDataOptionList) SetName(v string) *GetFaceOptionsResponseBodyDataOptionList {
	s.Name = &v
	return s
}

type GetFaceOptionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFaceOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFaceOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponse) SetHeaders(v map[string]*string) *GetFaceOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetFaceOptionsResponse) SetBody(v *GetFaceOptionsResponseBody) *GetFaceOptionsResponse {
	s.Body = v
	return s
}

type GetInventoryRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s GetInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryRequest) GoString() string {
	return s.String()
}

func (s *GetInventoryRequest) SetCommodityCode(v string) *GetInventoryRequest {
	s.CommodityCode = &v
	return s
}

type GetInventoryResponseBody struct {
	Data    *GetInventoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventoryResponseBody) SetData(v *GetInventoryResponseBodyData) *GetInventoryResponseBody {
	s.Data = v
	return s
}

func (s *GetInventoryResponseBody) SetSuccess(v bool) *GetInventoryResponseBody {
	s.Success = &v
	return s
}

type GetInventoryResponseBodyData struct {
	ResultObject []*GetInventoryResponseBodyDataResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
}

func (s GetInventoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInventoryResponseBodyData) SetResultObject(v []*GetInventoryResponseBodyDataResultObject) *GetInventoryResponseBodyData {
	s.ResultObject = v
	return s
}

type GetInventoryResponseBodyDataResultObject struct {
	CommodityCode    *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CurrentInventory *string `json:"CurrentInventory,omitempty" xml:"CurrentInventory,omitempty"`
	InventoryId      *string `json:"InventoryId,omitempty" xml:"InventoryId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BuyerId          *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	ValidStartTime   *string `json:"ValidStartTime,omitempty" xml:"ValidStartTime,omitempty"`
	ValidEndTime     *string `json:"ValidEndTime,omitempty" xml:"ValidEndTime,omitempty"`
}

func (s GetInventoryResponseBodyDataResultObject) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponseBodyDataResultObject) GoString() string {
	return s.String()
}

func (s *GetInventoryResponseBodyDataResultObject) SetCommodityCode(v string) *GetInventoryResponseBodyDataResultObject {
	s.CommodityCode = &v
	return s
}

func (s *GetInventoryResponseBodyDataResultObject) SetCurrentInventory(v string) *GetInventoryResponseBodyDataResultObject {
	s.CurrentInventory = &v
	return s
}

func (s *GetInventoryResponseBodyDataResultObject) SetInventoryId(v string) *GetInventoryResponseBodyDataResultObject {
	s.InventoryId = &v
	return s
}

func (s *GetInventoryResponseBodyDataResultObject) SetInstanceId(v string) *GetInventoryResponseBodyDataResultObject {
	s.InstanceId = &v
	return s
}

func (s *GetInventoryResponseBodyDataResultObject) SetBuyerId(v string) *GetInventoryResponseBodyDataResultObject {
	s.BuyerId = &v
	return s
}

func (s *GetInventoryResponseBodyDataResultObject) SetValidStartTime(v string) *GetInventoryResponseBodyDataResultObject {
	s.ValidStartTime = &v
	return s
}

func (s *GetInventoryResponseBodyDataResultObject) SetValidEndTime(v string) *GetInventoryResponseBodyDataResultObject {
	s.ValidEndTime = &v
	return s
}

type GetInventoryResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponse) GoString() string {
	return s.String()
}

func (s *GetInventoryResponse) SetHeaders(v map[string]*string) *GetInventoryResponse {
	s.Headers = v
	return s
}

func (s *GetInventoryResponse) SetBody(v *GetInventoryResponseBody) *GetInventoryResponse {
	s.Body = v
	return s
}

type GetMonitorListRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetMonitorListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListRequest) GoString() string {
	return s.String()
}

func (s *GetMonitorListRequest) SetCorpId(v string) *GetMonitorListRequest {
	s.CorpId = &v
	return s
}

func (s *GetMonitorListRequest) SetPageNo(v int32) *GetMonitorListRequest {
	s.PageNo = &v
	return s
}

func (s *GetMonitorListRequest) SetPageSize(v int32) *GetMonitorListRequest {
	s.PageSize = &v
	return s
}

type GetMonitorListResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetMonitorListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetMonitorListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponseBody) SetMessage(v string) *GetMonitorListResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonitorListResponseBody) SetRequestId(v string) *GetMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonitorListResponseBody) SetData(v *GetMonitorListResponseBodyData) *GetMonitorListResponseBody {
	s.Data = v
	return s
}

func (s *GetMonitorListResponseBody) SetCode(v string) *GetMonitorListResponseBody {
	s.Code = &v
	return s
}

type GetMonitorListResponseBodyData struct {
	Records    []*GetMonitorListResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNo     *int32                                   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TotalPage  *int32                                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMonitorListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponseBodyData) SetRecords(v []*GetMonitorListResponseBodyDataRecords) *GetMonitorListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetMonitorListResponseBodyData) SetPageNo(v int32) *GetMonitorListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetMonitorListResponseBodyData) SetTotalPage(v int32) *GetMonitorListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *GetMonitorListResponseBodyData) SetPageSize(v int32) *GetMonitorListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMonitorListResponseBodyData) SetTotalCount(v int32) *GetMonitorListResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetMonitorListResponseBodyDataRecords struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RuleExpression  *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	ImageMatch      *string `json:"ImageMatch,omitempty" xml:"ImageMatch,omitempty"`
	MonitorType     *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	NotifierType    *string `json:"NotifierType,omitempty" xml:"NotifierType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Expression      *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	NotifierExtra   *string `json:"NotifierExtra,omitempty" xml:"NotifierExtra,omitempty"`
	Attributes      *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	DeviceList      *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ModifiedDate    *string `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
}

func (s GetMonitorListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponseBodyDataRecords) SetStatus(v string) *GetMonitorListResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetRuleExpression(v string) *GetMonitorListResponseBodyDataRecords {
	s.RuleExpression = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetImageMatch(v string) *GetMonitorListResponseBodyDataRecords {
	s.ImageMatch = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetMonitorType(v string) *GetMonitorListResponseBodyDataRecords {
	s.MonitorType = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetCreateDate(v string) *GetMonitorListResponseBodyDataRecords {
	s.CreateDate = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetRuleName(v string) *GetMonitorListResponseBodyDataRecords {
	s.RuleName = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetNotifierType(v string) *GetMonitorListResponseBodyDataRecords {
	s.NotifierType = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetDescription(v string) *GetMonitorListResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetExpression(v string) *GetMonitorListResponseBodyDataRecords {
	s.Expression = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetNotifierExtra(v string) *GetMonitorListResponseBodyDataRecords {
	s.NotifierExtra = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetAttributes(v string) *GetMonitorListResponseBodyDataRecords {
	s.Attributes = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetDeviceList(v string) *GetMonitorListResponseBodyDataRecords {
	s.DeviceList = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetTaskId(v string) *GetMonitorListResponseBodyDataRecords {
	s.TaskId = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetModifiedDate(v string) *GetMonitorListResponseBodyDataRecords {
	s.ModifiedDate = &v
	return s
}

func (s *GetMonitorListResponseBodyDataRecords) SetAlgorithmVendor(v string) *GetMonitorListResponseBodyDataRecords {
	s.AlgorithmVendor = &v
	return s
}

type GetMonitorListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMonitorListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponse) SetHeaders(v map[string]*string) *GetMonitorListResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorListResponse) SetBody(v *GetMonitorListResponseBody) *GetMonitorListResponse {
	s.Body = v
	return s
}

type GetMonitorResultRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	MinRecordId     *string `json:"MinRecordId,omitempty" xml:"MinRecordId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
}

func (s GetMonitorResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultRequest) GoString() string {
	return s.String()
}

func (s *GetMonitorResultRequest) SetCorpId(v string) *GetMonitorResultRequest {
	s.CorpId = &v
	return s
}

func (s *GetMonitorResultRequest) SetTaskId(v string) *GetMonitorResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetMonitorResultRequest) SetMinRecordId(v string) *GetMonitorResultRequest {
	s.MinRecordId = &v
	return s
}

func (s *GetMonitorResultRequest) SetStartTime(v int64) *GetMonitorResultRequest {
	s.StartTime = &v
	return s
}

func (s *GetMonitorResultRequest) SetEndTime(v int64) *GetMonitorResultRequest {
	s.EndTime = &v
	return s
}

func (s *GetMonitorResultRequest) SetAlgorithmVendor(v string) *GetMonitorResultRequest {
	s.AlgorithmVendor = &v
	return s
}

type GetMonitorResultResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetMonitorResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetMonitorResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseBody) SetMessage(v string) *GetMonitorResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonitorResultResponseBody) SetRequestId(v string) *GetMonitorResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonitorResultResponseBody) SetData(v *GetMonitorResultResponseBodyData) *GetMonitorResultResponseBody {
	s.Data = v
	return s
}

func (s *GetMonitorResultResponseBody) SetCode(v string) *GetMonitorResultResponseBody {
	s.Code = &v
	return s
}

type GetMonitorResultResponseBodyData struct {
	MaxId   *string                                    `json:"MaxId,omitempty" xml:"MaxId,omitempty"`
	Records []*GetMonitorResultResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s GetMonitorResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseBodyData) SetMaxId(v string) *GetMonitorResultResponseBodyData {
	s.MaxId = &v
	return s
}

func (s *GetMonitorResultResponseBodyData) SetRecords(v []*GetMonitorResultResponseBodyDataRecords) *GetMonitorResultResponseBodyData {
	s.Records = v
	return s
}

type GetMonitorResultResponseBodyDataRecords struct {
	PicUrl        *string                                            `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	RightBottomY  *string                                            `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Score         *string                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	MonitorPicUrl *string                                            `json:"MonitorPicUrl,omitempty" xml:"MonitorPicUrl,omitempty"`
	RightBottomX  *string                                            `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	ExtendInfo    *GetMonitorResultResponseBodyDataRecordsExtendInfo `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty" type:"Struct"`
	GbId          *string                                            `json:"GbId,omitempty" xml:"GbId,omitempty"`
	LeftUpY       *string                                            `json:"LeftUpY,omitempty" xml:"LeftUpY,omitempty"`
	LeftUpX       *string                                            `json:"LeftUpX,omitempty" xml:"LeftUpX,omitempty"`
	ShotTime      *string                                            `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	TaskId        *string                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TargetPicUrl  *string                                            `json:"TargetPicUrl,omitempty" xml:"TargetPicUrl,omitempty"`
}

func (s GetMonitorResultResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseBodyDataRecords) SetPicUrl(v string) *GetMonitorResultResponseBodyDataRecords {
	s.PicUrl = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetRightBottomY(v string) *GetMonitorResultResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetScore(v string) *GetMonitorResultResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetMonitorPicUrl(v string) *GetMonitorResultResponseBodyDataRecords {
	s.MonitorPicUrl = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetRightBottomX(v string) *GetMonitorResultResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetExtendInfo(v *GetMonitorResultResponseBodyDataRecordsExtendInfo) *GetMonitorResultResponseBodyDataRecords {
	s.ExtendInfo = v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetGbId(v string) *GetMonitorResultResponseBodyDataRecords {
	s.GbId = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetLeftUpY(v string) *GetMonitorResultResponseBodyDataRecords {
	s.LeftUpY = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetLeftUpX(v string) *GetMonitorResultResponseBodyDataRecords {
	s.LeftUpX = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetShotTime(v string) *GetMonitorResultResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetTaskId(v string) *GetMonitorResultResponseBodyDataRecords {
	s.TaskId = &v
	return s
}

func (s *GetMonitorResultResponseBodyDataRecords) SetTargetPicUrl(v string) *GetMonitorResultResponseBodyDataRecords {
	s.TargetPicUrl = &v
	return s
}

type GetMonitorResultResponseBodyDataRecordsExtendInfo struct {
	PlateNo *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
}

func (s GetMonitorResultResponseBodyDataRecordsExtendInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseBodyDataRecordsExtendInfo) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseBodyDataRecordsExtendInfo) SetPlateNo(v string) *GetMonitorResultResponseBodyDataRecordsExtendInfo {
	s.PlateNo = &v
	return s
}

type GetMonitorResultResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMonitorResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMonitorResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponse) SetHeaders(v map[string]*string) *GetMonitorResultResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorResultResponse) SetBody(v *GetMonitorResultResponseBody) *GetMonitorResultResponse {
	s.Body = v
	return s
}

type GetPersonDetailRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonID      *string `json:"PersonID,omitempty" xml:"PersonID,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
}

func (s GetPersonDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailRequest) GoString() string {
	return s.String()
}

func (s *GetPersonDetailRequest) SetCorpId(v string) *GetPersonDetailRequest {
	s.CorpId = &v
	return s
}

func (s *GetPersonDetailRequest) SetPersonID(v string) *GetPersonDetailRequest {
	s.PersonID = &v
	return s
}

func (s *GetPersonDetailRequest) SetAlgorithmType(v string) *GetPersonDetailRequest {
	s.AlgorithmType = &v
	return s
}

type GetPersonDetailResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetPersonDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPersonDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponseBody) SetMessage(v string) *GetPersonDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetPersonDetailResponseBody) SetRequestId(v string) *GetPersonDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPersonDetailResponseBody) SetData(v *GetPersonDetailResponseBodyData) *GetPersonDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetPersonDetailResponseBody) SetCode(v string) *GetPersonDetailResponseBody {
	s.Code = &v
	return s
}

type GetPersonDetailResponseBodyData struct {
	PicUrl   *string                                   `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	TagList  []*GetPersonDetailResponseBodyDataTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	PersonId *string                                   `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s GetPersonDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponseBodyData) SetPicUrl(v string) *GetPersonDetailResponseBodyData {
	s.PicUrl = &v
	return s
}

func (s *GetPersonDetailResponseBodyData) SetTagList(v []*GetPersonDetailResponseBodyDataTagList) *GetPersonDetailResponseBodyData {
	s.TagList = v
	return s
}

func (s *GetPersonDetailResponseBodyData) SetPersonId(v string) *GetPersonDetailResponseBodyData {
	s.PersonId = &v
	return s
}

type GetPersonDetailResponseBodyDataTagList struct {
	TagValueId *string `json:"TagValueId,omitempty" xml:"TagValueId,omitempty"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagValue   *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetPersonDetailResponseBodyDataTagList) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponseBodyDataTagList) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponseBodyDataTagList) SetTagValueId(v string) *GetPersonDetailResponseBodyDataTagList {
	s.TagValueId = &v
	return s
}

func (s *GetPersonDetailResponseBodyDataTagList) SetTagName(v string) *GetPersonDetailResponseBodyDataTagList {
	s.TagName = &v
	return s
}

func (s *GetPersonDetailResponseBodyDataTagList) SetTagCode(v string) *GetPersonDetailResponseBodyDataTagList {
	s.TagCode = &v
	return s
}

func (s *GetPersonDetailResponseBodyDataTagList) SetTagValue(v string) *GetPersonDetailResponseBodyDataTagList {
	s.TagValue = &v
	return s
}

type GetPersonDetailResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPersonDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPersonDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponse) SetHeaders(v map[string]*string) *GetPersonDetailResponse {
	s.Headers = v
	return s
}

func (s *GetPersonDetailResponse) SetBody(v *GetPersonDetailResponseBody) *GetPersonDetailResponse {
	s.Body = v
	return s
}

type GetPersonListRequest struct {
	PageNumber                *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	FaceUrl                   *string                `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	CorpIdList                map[string]interface{} `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
	FaceMatchingRateThreshold *string                `json:"FaceMatchingRateThreshold,omitempty" xml:"FaceMatchingRateThreshold,omitempty"`
	CorpId                    *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonIdList              map[string]interface{} `json:"PersonIdList,omitempty" xml:"PersonIdList,omitempty"`
}

func (s GetPersonListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListRequest) GoString() string {
	return s.String()
}

func (s *GetPersonListRequest) SetPageNumber(v int64) *GetPersonListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPersonListRequest) SetPageSize(v int64) *GetPersonListRequest {
	s.PageSize = &v
	return s
}

func (s *GetPersonListRequest) SetFaceUrl(v string) *GetPersonListRequest {
	s.FaceUrl = &v
	return s
}

func (s *GetPersonListRequest) SetCorpIdList(v map[string]interface{}) *GetPersonListRequest {
	s.CorpIdList = v
	return s
}

func (s *GetPersonListRequest) SetFaceMatchingRateThreshold(v string) *GetPersonListRequest {
	s.FaceMatchingRateThreshold = &v
	return s
}

func (s *GetPersonListRequest) SetCorpId(v string) *GetPersonListRequest {
	s.CorpId = &v
	return s
}

func (s *GetPersonListRequest) SetPersonIdList(v map[string]interface{}) *GetPersonListRequest {
	s.PersonIdList = v
	return s
}

type GetPersonListShrinkRequest struct {
	PageNumber                *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	FaceUrl                   *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	CorpIdListShrink          *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
	FaceMatchingRateThreshold *string `json:"FaceMatchingRateThreshold,omitempty" xml:"FaceMatchingRateThreshold,omitempty"`
	CorpId                    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonIdListShrink        *string `json:"PersonIdList,omitempty" xml:"PersonIdList,omitempty"`
}

func (s GetPersonListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPersonListShrinkRequest) SetPageNumber(v int64) *GetPersonListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPersonListShrinkRequest) SetPageSize(v int64) *GetPersonListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetPersonListShrinkRequest) SetFaceUrl(v string) *GetPersonListShrinkRequest {
	s.FaceUrl = &v
	return s
}

func (s *GetPersonListShrinkRequest) SetCorpIdListShrink(v string) *GetPersonListShrinkRequest {
	s.CorpIdListShrink = &v
	return s
}

func (s *GetPersonListShrinkRequest) SetFaceMatchingRateThreshold(v string) *GetPersonListShrinkRequest {
	s.FaceMatchingRateThreshold = &v
	return s
}

func (s *GetPersonListShrinkRequest) SetCorpId(v string) *GetPersonListShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *GetPersonListShrinkRequest) SetPersonIdListShrink(v string) *GetPersonListShrinkRequest {
	s.PersonIdListShrink = &v
	return s
}

type GetPersonListResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetPersonListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPersonListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseBody) SetMessage(v string) *GetPersonListResponseBody {
	s.Message = &v
	return s
}

func (s *GetPersonListResponseBody) SetRequestId(v string) *GetPersonListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPersonListResponseBody) SetData(v *GetPersonListResponseBodyData) *GetPersonListResponseBody {
	s.Data = v
	return s
}

func (s *GetPersonListResponseBody) SetCode(v string) *GetPersonListResponseBody {
	s.Code = &v
	return s
}

type GetPersonListResponseBodyData struct {
	Records    []*GetPersonListResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetPersonListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseBodyData) SetRecords(v []*GetPersonListResponseBodyDataRecords) *GetPersonListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetPersonListResponseBodyData) SetPageNumber(v int64) *GetPersonListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetPersonListResponseBodyData) SetPageSize(v int64) *GetPersonListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetPersonListResponseBodyData) SetTotalCount(v int64) *GetPersonListResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetPersonListResponseBodyDataRecords struct {
	PropertyTagList    []*GetPersonListResponseBodyDataRecordsPropertyTagList `json:"PropertyTagList,omitempty" xml:"PropertyTagList,omitempty" type:"Repeated"`
	FaceUrl            *string                                                `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	SearchMatchingRate *string                                                `json:"SearchMatchingRate,omitempty" xml:"SearchMatchingRate,omitempty"`
	PersonId           *string                                                `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	LastShotTime       *int64                                                 `json:"LastShotTime,omitempty" xml:"LastShotTime,omitempty"`
	FirstShotTime      *int64                                                 `json:"FirstShotTime,omitempty" xml:"FirstShotTime,omitempty"`
}

func (s GetPersonListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseBodyDataRecords) SetPropertyTagList(v []*GetPersonListResponseBodyDataRecordsPropertyTagList) *GetPersonListResponseBodyDataRecords {
	s.PropertyTagList = v
	return s
}

func (s *GetPersonListResponseBodyDataRecords) SetFaceUrl(v string) *GetPersonListResponseBodyDataRecords {
	s.FaceUrl = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecords) SetSearchMatchingRate(v string) *GetPersonListResponseBodyDataRecords {
	s.SearchMatchingRate = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecords) SetPersonId(v string) *GetPersonListResponseBodyDataRecords {
	s.PersonId = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecords) SetLastShotTime(v int64) *GetPersonListResponseBodyDataRecords {
	s.LastShotTime = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecords) SetFirstShotTime(v int64) *GetPersonListResponseBodyDataRecords {
	s.FirstShotTime = &v
	return s
}

type GetPersonListResponseBodyDataRecordsPropertyTagList struct {
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TagName     *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagCodeName *string `json:"TagCodeName,omitempty" xml:"TagCodeName,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPersonListResponseBodyDataRecordsPropertyTagList) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseBodyDataRecordsPropertyTagList) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseBodyDataRecordsPropertyTagList) SetValue(v string) *GetPersonListResponseBodyDataRecordsPropertyTagList {
	s.Value = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecordsPropertyTagList) SetTagName(v string) *GetPersonListResponseBodyDataRecordsPropertyTagList {
	s.TagName = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecordsPropertyTagList) SetTagCodeName(v string) *GetPersonListResponseBodyDataRecordsPropertyTagList {
	s.TagCodeName = &v
	return s
}

func (s *GetPersonListResponseBodyDataRecordsPropertyTagList) SetCode(v string) *GetPersonListResponseBodyDataRecordsPropertyTagList {
	s.Code = &v
	return s
}

type GetPersonListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPersonListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPersonListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponse) GoString() string {
	return s.String()
}

func (s *GetPersonListResponse) SetHeaders(v map[string]*string) *GetPersonListResponse {
	s.Headers = v
	return s
}

func (s *GetPersonListResponse) SetBody(v *GetPersonListResponseBody) *GetPersonListResponse {
	s.Body = v
	return s
}

type GetProfileDetailRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	ProfileId *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
}

func (s GetProfileDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProfileDetailRequest) GoString() string {
	return s.String()
}

func (s *GetProfileDetailRequest) SetCorpId(v string) *GetProfileDetailRequest {
	s.CorpId = &v
	return s
}

func (s *GetProfileDetailRequest) SetIsvSubId(v string) *GetProfileDetailRequest {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileDetailRequest) SetProfileId(v int64) *GetProfileDetailRequest {
	s.ProfileId = &v
	return s
}

type GetProfileDetailResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetProfileDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetProfileDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProfileDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetProfileDetailResponseBody) SetMessage(v string) *GetProfileDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetProfileDetailResponseBody) SetRequestId(v string) *GetProfileDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProfileDetailResponseBody) SetData(v *GetProfileDetailResponseBodyData) *GetProfileDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetProfileDetailResponseBody) SetCode(v string) *GetProfileDetailResponseBody {
	s.Code = &v
	return s
}

type GetProfileDetailResponseBodyData struct {
	CatalogId   *int32  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ProfileId   *int32  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Gender      *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonId    *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
}

func (s GetProfileDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProfileDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProfileDetailResponseBodyData) SetCatalogId(v int32) *GetProfileDetailResponseBodyData {
	s.CatalogId = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetProfileId(v int32) *GetProfileDetailResponseBodyData {
	s.ProfileId = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetIsvSubId(v string) *GetProfileDetailResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetGender(v string) *GetProfileDetailResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetBizId(v string) *GetProfileDetailResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetIdNumber(v string) *GetProfileDetailResponseBodyData {
	s.IdNumber = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetSceneType(v string) *GetProfileDetailResponseBodyData {
	s.SceneType = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetPhoneNo(v string) *GetProfileDetailResponseBodyData {
	s.PhoneNo = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetFaceUrl(v string) *GetProfileDetailResponseBodyData {
	s.FaceUrl = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetLiveAddress(v string) *GetProfileDetailResponseBodyData {
	s.LiveAddress = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetName(v string) *GetProfileDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetPersonId(v string) *GetProfileDetailResponseBodyData {
	s.PersonId = &v
	return s
}

func (s *GetProfileDetailResponseBodyData) SetPlateNo(v string) *GetProfileDetailResponseBodyData {
	s.PlateNo = &v
	return s
}

type GetProfileDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProfileDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProfileDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProfileDetailResponse) GoString() string {
	return s.String()
}

func (s *GetProfileDetailResponse) SetHeaders(v map[string]*string) *GetProfileDetailResponse {
	s.Headers = v
	return s
}

func (s *GetProfileDetailResponse) SetBody(v *GetProfileDetailResponseBody) *GetProfileDetailResponse {
	s.Body = v
	return s
}

type GetProfileListRequest struct {
	CorpId                *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId              *string                `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Name                  *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId             *int64                 `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IdNumber              *string                `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl               *string                `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress           *string                `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender                *int32                 `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string                `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string                `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType             *string                `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId                 *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PersonIdList          map[string]interface{} `json:"PersonIdList,omitempty" xml:"PersonIdList,omitempty"`
	ProfileIdList         map[string]interface{} `json:"ProfileIdList,omitempty" xml:"ProfileIdList,omitempty"`
	MatchingRateThreshold *string                `json:"MatchingRateThreshold,omitempty" xml:"MatchingRateThreshold,omitempty"`
	FaceImageId           *string                `json:"FaceImageId,omitempty" xml:"FaceImageId,omitempty"`
}

func (s GetProfileListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListRequest) GoString() string {
	return s.String()
}

func (s *GetProfileListRequest) SetCorpId(v string) *GetProfileListRequest {
	s.CorpId = &v
	return s
}

func (s *GetProfileListRequest) SetIsvSubId(v string) *GetProfileListRequest {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileListRequest) SetName(v string) *GetProfileListRequest {
	s.Name = &v
	return s
}

func (s *GetProfileListRequest) SetCatalogId(v int64) *GetProfileListRequest {
	s.CatalogId = &v
	return s
}

func (s *GetProfileListRequest) SetIdNumber(v string) *GetProfileListRequest {
	s.IdNumber = &v
	return s
}

func (s *GetProfileListRequest) SetFaceUrl(v string) *GetProfileListRequest {
	s.FaceUrl = &v
	return s
}

func (s *GetProfileListRequest) SetLiveAddress(v string) *GetProfileListRequest {
	s.LiveAddress = &v
	return s
}

func (s *GetProfileListRequest) SetGender(v int32) *GetProfileListRequest {
	s.Gender = &v
	return s
}

func (s *GetProfileListRequest) SetPlateNo(v string) *GetProfileListRequest {
	s.PlateNo = &v
	return s
}

func (s *GetProfileListRequest) SetPhoneNo(v string) *GetProfileListRequest {
	s.PhoneNo = &v
	return s
}

func (s *GetProfileListRequest) SetSceneType(v string) *GetProfileListRequest {
	s.SceneType = &v
	return s
}

func (s *GetProfileListRequest) SetBizId(v string) *GetProfileListRequest {
	s.BizId = &v
	return s
}

func (s *GetProfileListRequest) SetPageNumber(v int64) *GetProfileListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetProfileListRequest) SetPageSize(v int64) *GetProfileListRequest {
	s.PageSize = &v
	return s
}

func (s *GetProfileListRequest) SetPersonIdList(v map[string]interface{}) *GetProfileListRequest {
	s.PersonIdList = v
	return s
}

func (s *GetProfileListRequest) SetProfileIdList(v map[string]interface{}) *GetProfileListRequest {
	s.ProfileIdList = v
	return s
}

func (s *GetProfileListRequest) SetMatchingRateThreshold(v string) *GetProfileListRequest {
	s.MatchingRateThreshold = &v
	return s
}

func (s *GetProfileListRequest) SetFaceImageId(v string) *GetProfileListRequest {
	s.FaceImageId = &v
	return s
}

type GetProfileListShrinkRequest struct {
	CorpId                *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId              *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId             *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IdNumber              *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl               *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress           *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender                *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType             *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PersonIdListShrink    *string `json:"PersonIdList,omitempty" xml:"PersonIdList,omitempty"`
	ProfileIdListShrink   *string `json:"ProfileIdList,omitempty" xml:"ProfileIdList,omitempty"`
	MatchingRateThreshold *string `json:"MatchingRateThreshold,omitempty" xml:"MatchingRateThreshold,omitempty"`
	FaceImageId           *string `json:"FaceImageId,omitempty" xml:"FaceImageId,omitempty"`
}

func (s GetProfileListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetProfileListShrinkRequest) SetCorpId(v string) *GetProfileListShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetIsvSubId(v string) *GetProfileListShrinkRequest {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetName(v string) *GetProfileListShrinkRequest {
	s.Name = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetCatalogId(v int64) *GetProfileListShrinkRequest {
	s.CatalogId = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetIdNumber(v string) *GetProfileListShrinkRequest {
	s.IdNumber = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetFaceUrl(v string) *GetProfileListShrinkRequest {
	s.FaceUrl = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetLiveAddress(v string) *GetProfileListShrinkRequest {
	s.LiveAddress = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetGender(v int32) *GetProfileListShrinkRequest {
	s.Gender = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetPlateNo(v string) *GetProfileListShrinkRequest {
	s.PlateNo = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetPhoneNo(v string) *GetProfileListShrinkRequest {
	s.PhoneNo = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetSceneType(v string) *GetProfileListShrinkRequest {
	s.SceneType = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetBizId(v string) *GetProfileListShrinkRequest {
	s.BizId = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetPageNumber(v int64) *GetProfileListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetPageSize(v int64) *GetProfileListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetPersonIdListShrink(v string) *GetProfileListShrinkRequest {
	s.PersonIdListShrink = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetProfileIdListShrink(v string) *GetProfileListShrinkRequest {
	s.ProfileIdListShrink = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetMatchingRateThreshold(v string) *GetProfileListShrinkRequest {
	s.MatchingRateThreshold = &v
	return s
}

func (s *GetProfileListShrinkRequest) SetFaceImageId(v string) *GetProfileListShrinkRequest {
	s.FaceImageId = &v
	return s
}

type GetProfileListResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetProfileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetProfileListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetProfileListResponseBody) SetMessage(v string) *GetProfileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetProfileListResponseBody) SetRequestId(v string) *GetProfileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProfileListResponseBody) SetData(v *GetProfileListResponseBodyData) *GetProfileListResponseBody {
	s.Data = v
	return s
}

func (s *GetProfileListResponseBody) SetCode(v string) *GetProfileListResponseBody {
	s.Code = &v
	return s
}

type GetProfileListResponseBodyData struct {
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Records    []*GetProfileListResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNumber *int64                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetProfileListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProfileListResponseBodyData) SetSuccess(v bool) *GetProfileListResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetProfileListResponseBodyData) SetRecords(v []*GetProfileListResponseBodyDataRecords) *GetProfileListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetProfileListResponseBodyData) SetPageNumber(v int64) *GetProfileListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetProfileListResponseBodyData) SetPageSize(v int64) *GetProfileListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetProfileListResponseBodyData) SetTotal(v int64) *GetProfileListResponseBodyData {
	s.Total = &v
	return s
}

type GetProfileListResponseBodyDataRecords struct {
	CatalogId          *int32  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ProfileId          *int32  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
	IdNumber           *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	SceneType          *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	IsvSubId           *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Gender             *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceUrl            *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	BizId              *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SearchMatchingRate *string `json:"SearchMatchingRate,omitempty" xml:"SearchMatchingRate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonId           *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s GetProfileListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetProfileListResponseBodyDataRecords) SetCatalogId(v int32) *GetProfileListResponseBodyDataRecords {
	s.CatalogId = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetProfileId(v int32) *GetProfileListResponseBodyDataRecords {
	s.ProfileId = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetIdNumber(v string) *GetProfileListResponseBodyDataRecords {
	s.IdNumber = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetSceneType(v string) *GetProfileListResponseBodyDataRecords {
	s.SceneType = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetIsvSubId(v string) *GetProfileListResponseBodyDataRecords {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetGender(v string) *GetProfileListResponseBodyDataRecords {
	s.Gender = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetFaceUrl(v string) *GetProfileListResponseBodyDataRecords {
	s.FaceUrl = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetBizId(v string) *GetProfileListResponseBodyDataRecords {
	s.BizId = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetSearchMatchingRate(v string) *GetProfileListResponseBodyDataRecords {
	s.SearchMatchingRate = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetName(v string) *GetProfileListResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *GetProfileListResponseBodyDataRecords) SetPersonId(v string) *GetProfileListResponseBodyDataRecords {
	s.PersonId = &v
	return s
}

type GetProfileListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProfileListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProfileListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponse) GoString() string {
	return s.String()
}

func (s *GetProfileListResponse) SetHeaders(v map[string]*string) *GetProfileListResponse {
	s.Headers = v
	return s
}

func (s *GetProfileListResponse) SetBody(v *GetProfileListResponseBody) *GetProfileListResponse {
	s.Body = v
	return s
}

type GetUserDetailRequest struct {
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId       *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserId         *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	NeedFaceDetail *bool   `json:"NeedFaceDetail,omitempty" xml:"NeedFaceDetail,omitempty"`
}

func (s GetUserDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserDetailRequest) GoString() string {
	return s.String()
}

func (s *GetUserDetailRequest) SetCorpId(v string) *GetUserDetailRequest {
	s.CorpId = &v
	return s
}

func (s *GetUserDetailRequest) SetIsvSubId(v string) *GetUserDetailRequest {
	s.IsvSubId = &v
	return s
}

func (s *GetUserDetailRequest) SetUserId(v int64) *GetUserDetailRequest {
	s.UserId = &v
	return s
}

func (s *GetUserDetailRequest) SetNeedFaceDetail(v bool) *GetUserDetailRequest {
	s.NeedFaceDetail = &v
	return s
}

type GetUserDetailResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetUserDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetUserDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDetailResponseBody) SetMessage(v string) *GetUserDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserDetailResponseBody) SetRequestId(v string) *GetUserDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDetailResponseBody) SetData(v *GetUserDetailResponseBodyData) *GetUserDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetUserDetailResponseBody) SetCode(v string) *GetUserDetailResponseBody {
	s.Code = &v
	return s
}

type GetUserDetailResponseBodyData struct {
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId  *int32  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId       *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	PhoneNo      *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	PlateNo      *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
}

func (s GetUserDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserDetailResponseBodyData) SetGender(v string) *GetUserDetailResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetFaceImageUrl(v string) *GetUserDetailResponseBodyData {
	s.FaceImageUrl = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetIsvSubId(v string) *GetUserDetailResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetUserGroupId(v int32) *GetUserDetailResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetUserId(v int32) *GetUserDetailResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetBizId(v string) *GetUserDetailResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetAttachment(v string) *GetUserDetailResponseBodyData {
	s.Attachment = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetAge(v string) *GetUserDetailResponseBodyData {
	s.Age = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetIdNumber(v string) *GetUserDetailResponseBodyData {
	s.IdNumber = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetPhoneNo(v string) *GetUserDetailResponseBodyData {
	s.PhoneNo = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetAddress(v string) *GetUserDetailResponseBodyData {
	s.Address = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetUserName(v string) *GetUserDetailResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetUserDetailResponseBodyData) SetPlateNo(v string) *GetUserDetailResponseBodyData {
	s.PlateNo = &v
	return s
}

type GetUserDetailResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserDetailResponse) GoString() string {
	return s.String()
}

func (s *GetUserDetailResponse) SetHeaders(v map[string]*string) *GetUserDetailResponse {
	s.Headers = v
	return s
}

func (s *GetUserDetailResponse) SetBody(v *GetUserDetailResponseBody) *GetUserDetailResponse {
	s.Body = v
	return s
}

type GetVideoComposeResultRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskRequestId *string `json:"TaskRequestId,omitempty" xml:"TaskRequestId,omitempty"`
}

func (s GetVideoComposeResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoComposeResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoComposeResultRequest) SetCorpId(v string) *GetVideoComposeResultRequest {
	s.CorpId = &v
	return s
}

func (s *GetVideoComposeResultRequest) SetTaskRequestId(v string) *GetVideoComposeResultRequest {
	s.TaskRequestId = &v
	return s
}

type GetVideoComposeResultResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	VideoUrl  *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetVideoComposeResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoComposeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoComposeResultResponseBody) SetStatus(v string) *GetVideoComposeResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoComposeResultResponseBody) SetMessage(v string) *GetVideoComposeResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoComposeResultResponseBody) SetRequestId(v string) *GetVideoComposeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoComposeResultResponseBody) SetCode(v string) *GetVideoComposeResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoComposeResultResponseBody) SetVideoUrl(v string) *GetVideoComposeResultResponseBody {
	s.VideoUrl = &v
	return s
}

type GetVideoComposeResultResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoComposeResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoComposeResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoComposeResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoComposeResultResponse) SetHeaders(v map[string]*string) *GetVideoComposeResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoComposeResultResponse) SetBody(v *GetVideoComposeResultResponseBody) *GetVideoComposeResultResponse {
	s.Body = v
	return s
}

type GetVideoSummaryTaskResultRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetVideoSummaryTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoSummaryTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoSummaryTaskResultRequest) SetCorpId(v string) *GetVideoSummaryTaskResultRequest {
	s.CorpId = &v
	return s
}

func (s *GetVideoSummaryTaskResultRequest) SetTaskId(v string) *GetVideoSummaryTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetVideoSummaryTaskResultResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetVideoSummaryTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoSummaryTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoSummaryTaskResultResponseBody) SetMessage(v string) *GetVideoSummaryTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoSummaryTaskResultResponseBody) SetRequestId(v string) *GetVideoSummaryTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoSummaryTaskResultResponseBody) SetData(v string) *GetVideoSummaryTaskResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetVideoSummaryTaskResultResponseBody) SetCode(v string) *GetVideoSummaryTaskResultResponseBody {
	s.Code = &v
	return s
}

type GetVideoSummaryTaskResultResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoSummaryTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoSummaryTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoSummaryTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoSummaryTaskResultResponse) SetHeaders(v map[string]*string) *GetVideoSummaryTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoSummaryTaskResultResponse) SetBody(v *GetVideoSummaryTaskResultResponseBody) *GetVideoSummaryTaskResultResponse {
	s.Body = v
	return s
}

type InvokeMotorModelRequest struct {
	PicId   *string `json:"PicId,omitempty" xml:"PicId,omitempty"`
	CorpId  *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PicPath *string `json:"PicPath,omitempty" xml:"PicPath,omitempty"`
	PicUrl  *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s InvokeMotorModelRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeMotorModelRequest) GoString() string {
	return s.String()
}

func (s *InvokeMotorModelRequest) SetPicId(v string) *InvokeMotorModelRequest {
	s.PicId = &v
	return s
}

func (s *InvokeMotorModelRequest) SetCorpId(v string) *InvokeMotorModelRequest {
	s.CorpId = &v
	return s
}

func (s *InvokeMotorModelRequest) SetPicPath(v string) *InvokeMotorModelRequest {
	s.PicPath = &v
	return s
}

func (s *InvokeMotorModelRequest) SetPicUrl(v string) *InvokeMotorModelRequest {
	s.PicUrl = &v
	return s
}

type InvokeMotorModelResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *InvokeMotorModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s InvokeMotorModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeMotorModelResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeMotorModelResponseBody) SetMessage(v string) *InvokeMotorModelResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeMotorModelResponseBody) SetRequestId(v string) *InvokeMotorModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeMotorModelResponseBody) SetData(v *InvokeMotorModelResponseBodyData) *InvokeMotorModelResponseBody {
	s.Data = v
	return s
}

func (s *InvokeMotorModelResponseBody) SetCode(v string) *InvokeMotorModelResponseBody {
	s.Code = &v
	return s
}

type InvokeMotorModelResponseBodyData struct {
	StructList *string `json:"StructList,omitempty" xml:"StructList,omitempty"`
}

func (s InvokeMotorModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InvokeMotorModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvokeMotorModelResponseBodyData) SetStructList(v string) *InvokeMotorModelResponseBodyData {
	s.StructList = &v
	return s
}

type InvokeMotorModelResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InvokeMotorModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeMotorModelResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeMotorModelResponse) GoString() string {
	return s.String()
}

func (s *InvokeMotorModelResponse) SetHeaders(v map[string]*string) *InvokeMotorModelResponse {
	s.Headers = v
	return s
}

func (s *InvokeMotorModelResponse) SetBody(v *InvokeMotorModelResponseBody) *InvokeMotorModelResponse {
	s.Body = v
	return s
}

type ListAccessNumberRequest struct {
	CorpIdList *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
}

func (s ListAccessNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessNumberRequest) GoString() string {
	return s.String()
}

func (s *ListAccessNumberRequest) SetCorpIdList(v string) *ListAccessNumberRequest {
	s.CorpIdList = &v
	return s
}

type ListAccessNumberResponseBody struct {
	// Id of the request
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      []*ListAccessNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListAccessNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessNumberResponseBody) SetRequestId(v string) *ListAccessNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessNumberResponseBody) SetCode(v string) *ListAccessNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccessNumberResponseBody) SetMessage(v string) *ListAccessNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccessNumberResponseBody) SetData(v []*ListAccessNumberResponseBodyData) *ListAccessNumberResponseBody {
	s.Data = v
	return s
}

type ListAccessNumberResponseBodyData struct {
	Item    *string `json:"Item,omitempty" xml:"Item,omitempty"`
	Count   *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
}

func (s ListAccessNumberResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAccessNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccessNumberResponseBodyData) SetItem(v string) *ListAccessNumberResponseBodyData {
	s.Item = &v
	return s
}

func (s *ListAccessNumberResponseBodyData) SetCount(v string) *ListAccessNumberResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListAccessNumberResponseBodyData) SetPercent(v string) *ListAccessNumberResponseBodyData {
	s.Percent = &v
	return s
}

type ListAccessNumberResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAccessNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccessNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessNumberResponse) GoString() string {
	return s.String()
}

func (s *ListAccessNumberResponse) SetHeaders(v map[string]*string) *ListAccessNumberResponse {
	s.Headers = v
	return s
}

func (s *ListAccessNumberResponse) SetBody(v *ListAccessNumberResponseBody) *ListAccessNumberResponse {
	s.Body = v
	return s
}

type ListBodyAlgorithmResultsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CapStyle      *string `json:"CapStyle,omitempty" xml:"CapStyle,omitempty"`
}

func (s ListBodyAlgorithmResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsRequest) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsRequest) SetCorpId(v string) *ListBodyAlgorithmResultsRequest {
	s.CorpId = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetAlgorithmType(v string) *ListBodyAlgorithmResultsRequest {
	s.AlgorithmType = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetDataSourceId(v string) *ListBodyAlgorithmResultsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetStartTime(v string) *ListBodyAlgorithmResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetEndTime(v string) *ListBodyAlgorithmResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetPageNumber(v string) *ListBodyAlgorithmResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetPageSize(v string) *ListBodyAlgorithmResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBodyAlgorithmResultsRequest) SetCapStyle(v string) *ListBodyAlgorithmResultsRequest {
	s.CapStyle = &v
	return s
}

type ListBodyAlgorithmResultsResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListBodyAlgorithmResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListBodyAlgorithmResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponseBody) SetMessage(v string) *ListBodyAlgorithmResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBody) SetRequestId(v string) *ListBodyAlgorithmResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBody) SetData(v *ListBodyAlgorithmResultsResponseBodyData) *ListBodyAlgorithmResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBody) SetCode(v string) *ListBodyAlgorithmResultsResponseBody {
	s.Code = &v
	return s
}

type ListBodyAlgorithmResultsResponseBodyData struct {
	Records    []*ListBodyAlgorithmResultsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBodyAlgorithmResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponseBodyData) SetRecords(v []*ListBodyAlgorithmResultsResponseBodyDataRecords) *ListBodyAlgorithmResultsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyData) SetTotalPage(v int32) *ListBodyAlgorithmResultsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyData) SetPageNumber(v int32) *ListBodyAlgorithmResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyData) SetPageSize(v int32) *ListBodyAlgorithmResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyData) SetTotalCount(v int32) *ListBodyAlgorithmResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListBodyAlgorithmResultsResponseBodyDataRecords struct {
	RightBottomY     *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceId     *string  `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PicUrlPath       *string  `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	TrousersColor    *string  `json:"TrousersColor,omitempty" xml:"TrousersColor,omitempty"`
	RightBottomX     *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	CoatColor        *string  `json:"CoatColor,omitempty" xml:"CoatColor,omitempty"`
	SourceId         *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	MaxAge           *string  `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	CoatLength       *string  `json:"CoatLength,omitempty" xml:"CoatLength,omitempty"`
	TargetPicUrlPath *string  `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	HairStyle        *string  `json:"HairStyle,omitempty" xml:"HairStyle,omitempty"`
	CoatStyle        *string  `json:"CoatStyle,omitempty" xml:"CoatStyle,omitempty"`
	LeftTopY         *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	MinAge           *string  `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	CorpId           *string  `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TrousersLength   *string  `json:"TrousersLength,omitempty" xml:"TrousersLength,omitempty"`
	TrousersStyle    *string  `json:"TrousersStyle,omitempty" xml:"TrousersStyle,omitempty"`
	ShotTime         *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	LeftTopX         *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	GenderCode       *string  `json:"GenderCode,omitempty" xml:"GenderCode,omitempty"`
	PersonId         *string  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	CapStyle         *string  `json:"CapStyle,omitempty" xml:"CapStyle,omitempty"`
}

func (s ListBodyAlgorithmResultsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetRightBottomY(v float32) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetDataSourceId(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetPicUrlPath(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetTrousersColor(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.TrousersColor = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetRightBottomX(v float32) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetCoatColor(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.CoatColor = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetSourceId(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.SourceId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetMaxAge(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.MaxAge = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetCoatLength(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.CoatLength = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetTargetPicUrlPath(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetHairStyle(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.HairStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetCoatStyle(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.CoatStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetLeftTopY(v float32) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetMinAge(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.MinAge = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetCorpId(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetTrousersLength(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.TrousersLength = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetTrousersStyle(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.TrousersStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetShotTime(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetLeftTopX(v float32) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetGenderCode(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.GenderCode = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetPersonId(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.PersonId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseBodyDataRecords) SetCapStyle(v string) *ListBodyAlgorithmResultsResponseBodyDataRecords {
	s.CapStyle = &v
	return s
}

type ListBodyAlgorithmResultsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBodyAlgorithmResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBodyAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponse) SetHeaders(v map[string]*string) *ListBodyAlgorithmResultsResponse {
	s.Headers = v
	return s
}

func (s *ListBodyAlgorithmResultsResponse) SetBody(v *ListBodyAlgorithmResultsResponseBody) *ListBodyAlgorithmResultsResponse {
	s.Body = v
	return s
}

type ListCorpGroupMetricsRequest struct {
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagCode     *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	UserGroup   *string `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
	DeviceGroup *string `json:"DeviceGroup,omitempty" xml:"DeviceGroup,omitempty"`
}

func (s ListCorpGroupMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListCorpGroupMetricsRequest) SetStartTime(v string) *ListCorpGroupMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetTagCode(v string) *ListCorpGroupMetricsRequest {
	s.TagCode = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetEndTime(v string) *ListCorpGroupMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetGroupId(v string) *ListCorpGroupMetricsRequest {
	s.GroupId = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetPageNumber(v string) *ListCorpGroupMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetPageSize(v string) *ListCorpGroupMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetDeviceId(v string) *ListCorpGroupMetricsRequest {
	s.DeviceId = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetCorpId(v string) *ListCorpGroupMetricsRequest {
	s.CorpId = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetUserGroup(v string) *ListCorpGroupMetricsRequest {
	s.UserGroup = &v
	return s
}

func (s *ListCorpGroupMetricsRequest) SetDeviceGroup(v string) *ListCorpGroupMetricsRequest {
	s.DeviceGroup = &v
	return s
}

type ListCorpGroupMetricsResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCorpGroupMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCorpGroupMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorpGroupMetricsResponseBody) SetTotalCount(v int32) *ListCorpGroupMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetRequestId(v string) *ListCorpGroupMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetMessage(v string) *ListCorpGroupMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetPageSize(v int32) *ListCorpGroupMetricsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetPageNumber(v int32) *ListCorpGroupMetricsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetData(v []*ListCorpGroupMetricsResponseBodyData) *ListCorpGroupMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetCode(v string) *ListCorpGroupMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBody) SetSuccess(v string) *ListCorpGroupMetricsResponseBody {
	s.Success = &v
	return s
}

type ListCorpGroupMetricsResponseBodyData struct {
	DateId        *string `json:"DateId,omitempty" xml:"DateId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	CorpGroupId   *string `json:"CorpGroupId,omitempty" xml:"CorpGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagMetrics    *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty"`
	TagValue      *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	PersonID      *string `json:"PersonID,omitempty" xml:"PersonID,omitempty"`
}

func (s ListCorpGroupMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCorpGroupMetricsResponseBodyData) SetDateId(v string) *ListCorpGroupMetricsResponseBodyData {
	s.DateId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetDeviceGroupId(v string) *ListCorpGroupMetricsResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetCorpGroupId(v string) *ListCorpGroupMetricsResponseBodyData {
	s.CorpGroupId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetDeviceId(v string) *ListCorpGroupMetricsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetUserGroupId(v string) *ListCorpGroupMetricsResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetTagCode(v string) *ListCorpGroupMetricsResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetCorpId(v string) *ListCorpGroupMetricsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetTagMetrics(v string) *ListCorpGroupMetricsResponseBodyData {
	s.TagMetrics = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetTagValue(v string) *ListCorpGroupMetricsResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListCorpGroupMetricsResponseBodyData) SetPersonID(v string) *ListCorpGroupMetricsResponseBodyData {
	s.PersonID = &v
	return s
}

type ListCorpGroupMetricsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorpGroupMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorpGroupMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpGroupMetricsResponse) SetHeaders(v map[string]*string) *ListCorpGroupMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetBody(v *ListCorpGroupMetricsResponseBody) *ListCorpGroupMetricsResponse {
	s.Body = v
	return s
}

type ListCorpGroupsRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCorpGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCorpGroupsRequest) SetCorpId(v string) *ListCorpGroupsRequest {
	s.CorpId = &v
	return s
}

func (s *ListCorpGroupsRequest) SetPageNumber(v int64) *ListCorpGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpGroupsRequest) SetPageSize(v int64) *ListCorpGroupsRequest {
	s.PageSize = &v
	return s
}

type ListCorpGroupsResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListCorpGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCorpGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorpGroupsResponseBody) SetMessage(v string) *ListCorpGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCorpGroupsResponseBody) SetRequestId(v string) *ListCorpGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorpGroupsResponseBody) SetData(v *ListCorpGroupsResponseBodyData) *ListCorpGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListCorpGroupsResponseBody) SetCode(v string) *ListCorpGroupsResponseBody {
	s.Code = &v
	return s
}

type ListCorpGroupsResponseBodyData struct {
	Records    []*string `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int64    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int64    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCorpGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCorpGroupsResponseBodyData) SetRecords(v []*string) *ListCorpGroupsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListCorpGroupsResponseBodyData) SetTotalPage(v int64) *ListCorpGroupsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListCorpGroupsResponseBodyData) SetPageNumber(v int64) *ListCorpGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCorpGroupsResponseBodyData) SetPageSize(v int64) *ListCorpGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCorpGroupsResponseBodyData) SetTotalCount(v int64) *ListCorpGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCorpGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorpGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorpGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpGroupsResponse) SetHeaders(v map[string]*string) *ListCorpGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListCorpGroupsResponse) SetBody(v *ListCorpGroupsResponseBody) *ListCorpGroupsResponse {
	s.Body = v
	return s
}

type ListCorpMetricsRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode         *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserGroupList   *string `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty"`
	DeviceGroupList *string `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty"`
	DeviceIdList    *string `json:"DeviceIdList,omitempty" xml:"DeviceIdList,omitempty"`
}

func (s ListCorpMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsRequest) SetCorpId(v string) *ListCorpMetricsRequest {
	s.CorpId = &v
	return s
}

func (s *ListCorpMetricsRequest) SetTagCode(v string) *ListCorpMetricsRequest {
	s.TagCode = &v
	return s
}

func (s *ListCorpMetricsRequest) SetStartTime(v string) *ListCorpMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCorpMetricsRequest) SetEndTime(v string) *ListCorpMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCorpMetricsRequest) SetPageNumber(v string) *ListCorpMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpMetricsRequest) SetPageSize(v string) *ListCorpMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpMetricsRequest) SetUserGroupList(v string) *ListCorpMetricsRequest {
	s.UserGroupList = &v
	return s
}

func (s *ListCorpMetricsRequest) SetDeviceGroupList(v string) *ListCorpMetricsRequest {
	s.DeviceGroupList = &v
	return s
}

func (s *ListCorpMetricsRequest) SetDeviceIdList(v string) *ListCorpMetricsRequest {
	s.DeviceIdList = &v
	return s
}

type ListCorpMetricsResponseBody struct {
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCorpMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCorpMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsResponseBody) SetTotalCount(v int32) *ListCorpMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCorpMetricsResponseBody) SetRequestId(v string) *ListCorpMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorpMetricsResponseBody) SetMessage(v string) *ListCorpMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCorpMetricsResponseBody) SetPageSize(v int32) *ListCorpMetricsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCorpMetricsResponseBody) SetPageNumber(v int32) *ListCorpMetricsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCorpMetricsResponseBody) SetData(v []*ListCorpMetricsResponseBodyData) *ListCorpMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ListCorpMetricsResponseBody) SetCode(v string) *ListCorpMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCorpMetricsResponseBody) SetSuccess(v string) *ListCorpMetricsResponseBody {
	s.Success = &v
	return s
}

type ListCorpMetricsResponseBodyData struct {
	DateId        *string `json:"DateId,omitempty" xml:"DateId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagMetrics    *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty"`
	TagValue      *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	PersonId      *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListCorpMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsResponseBodyData) SetDateId(v string) *ListCorpMetricsResponseBodyData {
	s.DateId = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetDeviceGroupId(v string) *ListCorpMetricsResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetDeviceId(v string) *ListCorpMetricsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetUserGroupId(v string) *ListCorpMetricsResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetTagCode(v string) *ListCorpMetricsResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetCorpId(v string) *ListCorpMetricsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetTagMetrics(v string) *ListCorpMetricsResponseBodyData {
	s.TagMetrics = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetTagValue(v string) *ListCorpMetricsResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListCorpMetricsResponseBodyData) SetPersonId(v string) *ListCorpMetricsResponseBodyData {
	s.PersonId = &v
	return s
}

type ListCorpMetricsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorpMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorpMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsResponse) SetHeaders(v map[string]*string) *ListCorpMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListCorpMetricsResponse) SetBody(v *ListCorpMetricsResponseBody) *ListCorpMetricsResponse {
	s.Body = v
	return s
}

type ListCorpsRequest struct {
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
}

func (s ListCorpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsRequest) GoString() string {
	return s.String()
}

func (s *ListCorpsRequest) SetPageNumber(v int32) *ListCorpsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpsRequest) SetPageSize(v int32) *ListCorpsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpsRequest) SetCorpName(v string) *ListCorpsRequest {
	s.CorpName = &v
	return s
}

type ListCorpsResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListCorpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCorpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorpsResponseBody) SetMessage(v string) *ListCorpsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCorpsResponseBody) SetRequestId(v string) *ListCorpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorpsResponseBody) SetData(v *ListCorpsResponseBodyData) *ListCorpsResponseBody {
	s.Data = v
	return s
}

func (s *ListCorpsResponseBody) SetCode(v string) *ListCorpsResponseBody {
	s.Code = &v
	return s
}

type ListCorpsResponseBodyData struct {
	Records    []*ListCorpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCorpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCorpsResponseBodyData) SetRecords(v []*ListCorpsResponseBodyDataRecords) *ListCorpsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListCorpsResponseBodyData) SetTotalPage(v int32) *ListCorpsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListCorpsResponseBodyData) SetPageNumber(v int32) *ListCorpsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCorpsResponseBodyData) SetPageSize(v int32) *ListCorpsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCorpsResponseBodyData) SetTotalCount(v int32) *ListCorpsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCorpsResponseBodyDataRecords struct {
	ParentCorpId *string `json:"ParentCorpId,omitempty" xml:"ParentCorpId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CorpName     *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	AcuUsed      *int32  `json:"AcuUsed,omitempty" xml:"AcuUsed,omitempty"`
	CreateDate   *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	IconPath     *string `json:"IconPath,omitempty" xml:"IconPath,omitempty"`
	DeviceCount  *int32  `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
}

func (s ListCorpsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListCorpsResponseBodyDataRecords) SetParentCorpId(v string) *ListCorpsResponseBodyDataRecords {
	s.ParentCorpId = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetAppName(v string) *ListCorpsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetIsvSubId(v string) *ListCorpsResponseBodyDataRecords {
	s.IsvSubId = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetDescription(v string) *ListCorpsResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetCorpName(v string) *ListCorpsResponseBodyDataRecords {
	s.CorpName = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetCorpId(v string) *ListCorpsResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetAcuUsed(v int32) *ListCorpsResponseBodyDataRecords {
	s.AcuUsed = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetCreateDate(v string) *ListCorpsResponseBodyDataRecords {
	s.CreateDate = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetIconPath(v string) *ListCorpsResponseBodyDataRecords {
	s.IconPath = &v
	return s
}

func (s *ListCorpsResponseBodyDataRecords) SetDeviceCount(v int32) *ListCorpsResponseBodyDataRecords {
	s.DeviceCount = &v
	return s
}

type ListCorpsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpsResponse) SetHeaders(v map[string]*string) *ListCorpsResponse {
	s.Headers = v
	return s
}

func (s *ListCorpsResponse) SetBody(v *ListCorpsResponseBody) *ListCorpsResponse {
	s.Body = v
	return s
}

type ListDeviceGroupsRequest struct {
	DeviceCodeList *string `json:"DeviceCodeList,omitempty" xml:"DeviceCodeList,omitempty"`
	CorpIdList     *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	IsPage         *int32  `json:"IsPage,omitempty" xml:"IsPage,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Group          *string `json:"Group,omitempty" xml:"Group,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
}

func (s ListDeviceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsRequest) SetDeviceCodeList(v string) *ListDeviceGroupsRequest {
	s.DeviceCodeList = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetCorpIdList(v string) *ListDeviceGroupsRequest {
	s.CorpIdList = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetName(v string) *ListDeviceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetIsPage(v int32) *ListDeviceGroupsRequest {
	s.IsPage = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetPageNum(v int32) *ListDeviceGroupsRequest {
	s.PageNum = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetPageSize(v int32) *ListDeviceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetGroup(v string) *ListDeviceGroupsRequest {
	s.Group = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetDataSourceType(v string) *ListDeviceGroupsRequest {
	s.DataSourceType = &v
	return s
}

type ListDeviceGroupsResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListDeviceGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDeviceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseBody) SetMessage(v string) *ListDeviceGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceGroupsResponseBody) SetRequestId(v string) *ListDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGroupsResponseBody) SetData(v []*ListDeviceGroupsResponseBodyData) *ListDeviceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceGroupsResponseBody) SetCode(v string) *ListDeviceGroupsResponseBody {
	s.Code = &v
	return s
}

type ListDeviceGroupsResponseBodyData struct {
	List       []*ListDeviceGroupsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *string                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseBodyData) SetList(v []*ListDeviceGroupsResponseBodyDataList) *ListDeviceGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListDeviceGroupsResponseBodyData) SetTotalCount(v string) *ListDeviceGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListDeviceGroupsResponseBodyDataList struct {
	DeviceStreamStatus  *string `json:"DeviceStreamStatus,omitempty" xml:"DeviceStreamStatus,omitempty"`
	DeviceName          *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceStatus        *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DeviceSn            *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	DeviceComputeStatus *string `json:"DeviceComputeStatus,omitempty" xml:"DeviceComputeStatus,omitempty"`
	InstallAddress      *string `json:"InstallAddress,omitempty" xml:"InstallAddress,omitempty"`
	DeviceGroup         *string `json:"DeviceGroup,omitempty" xml:"DeviceGroup,omitempty"`
	RegionName          *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	DataSourceType      *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	ResolvingPower      *string `json:"ResolvingPower,omitempty" xml:"ResolvingPower,omitempty"`
	DeviceCode          *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	BitRate             *string `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	CodingFormat        *string `json:"CodingFormat,omitempty" xml:"CodingFormat,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDeviceGroupsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceStreamStatus(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceStreamStatus = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceName(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceStatus(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceStatus = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetRegionId(v string) *ListDeviceGroupsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceSn(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceSn = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceComputeStatus(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceComputeStatus = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetInstallAddress(v string) *ListDeviceGroupsResponseBodyDataList {
	s.InstallAddress = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceGroup(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceGroup = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetRegionName(v string) *ListDeviceGroupsResponseBodyDataList {
	s.RegionName = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDataSourceType(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DataSourceType = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetResolvingPower(v string) *ListDeviceGroupsResponseBodyDataList {
	s.ResolvingPower = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetDeviceCode(v string) *ListDeviceGroupsResponseBodyDataList {
	s.DeviceCode = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetBitRate(v string) *ListDeviceGroupsResponseBodyDataList {
	s.BitRate = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetCodingFormat(v string) *ListDeviceGroupsResponseBodyDataList {
	s.CodingFormat = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDataList) SetType(v string) *ListDeviceGroupsResponseBodyDataList {
	s.Type = &v
	return s
}

type ListDeviceGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponse) SetHeaders(v map[string]*string) *ListDeviceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceGroupsResponse) SetBody(v *ListDeviceGroupsResponseBody) *ListDeviceGroupsResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId       *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetCorpId(v string) *ListDevicesRequest {
	s.CorpId = &v
	return s
}

func (s *ListDevicesRequest) SetGbId(v string) *ListDevicesRequest {
	s.GbId = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceName(v string) *ListDevicesRequest {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesRequest) SetPageNumber(v int32) *ListDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesRequest) SetPageSize(v int32) *ListDevicesRequest {
	s.PageSize = &v
	return s
}

type ListDevicesResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetMessage(v string) *ListDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetData(v *ListDevicesResponseBodyData) *ListDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicesResponseBody) SetCode(v string) *ListDevicesResponseBody {
	s.Code = &v
	return s
}

type ListDevicesResponseBodyData struct {
	Records    []*ListDevicesResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyData) SetRecords(v []*ListDevicesResponseBodyDataRecords) *ListDevicesResponseBodyData {
	s.Records = v
	return s
}

func (s *ListDevicesResponseBodyData) SetTotalPage(v int32) *ListDevicesResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetPageNumber(v int32) *ListDevicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetPageSize(v int32) *ListDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetTotalCount(v int32) *ListDevicesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListDevicesResponseBodyDataRecords struct {
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SipGBId            *string `json:"SipGBId,omitempty" xml:"SipGBId,omitempty"`
	DeviceDirection    *string `json:"DeviceDirection,omitempty" xml:"DeviceDirection,omitempty"`
	DeviceName         *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceAddress      *string `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty"`
	DeviceType         *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SipPassword        *string `json:"SipPassword,omitempty" xml:"SipPassword,omitempty"`
	SipServerPort      *string `json:"SipServerPort,omitempty" xml:"SipServerPort,omitempty"`
	Vendor             *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	GbId               *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	CoverImageUrl      *string `json:"CoverImageUrl,omitempty" xml:"CoverImageUrl,omitempty"`
	AccessProtocolType *string `json:"AccessProtocolType,omitempty" xml:"AccessProtocolType,omitempty"`
	DeviceSite         *string `json:"DeviceSite,omitempty" xml:"DeviceSite,omitempty"`
	Longitude          *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude           *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Resolution         *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	SipServerIp        *string `json:"SipServerIp,omitempty" xml:"SipServerIp,omitempty"`
	BitRate            *string `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
}

func (s ListDevicesResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataRecords) SetStatus(v int32) *ListDevicesResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetSipGBId(v string) *ListDevicesResponseBodyDataRecords {
	s.SipGBId = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetDeviceDirection(v string) *ListDevicesResponseBodyDataRecords {
	s.DeviceDirection = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetDeviceName(v string) *ListDevicesResponseBodyDataRecords {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetDeviceAddress(v string) *ListDevicesResponseBodyDataRecords {
	s.DeviceAddress = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetDeviceType(v string) *ListDevicesResponseBodyDataRecords {
	s.DeviceType = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetCreateTime(v string) *ListDevicesResponseBodyDataRecords {
	s.CreateTime = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetSipPassword(v string) *ListDevicesResponseBodyDataRecords {
	s.SipPassword = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetSipServerPort(v string) *ListDevicesResponseBodyDataRecords {
	s.SipServerPort = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetVendor(v string) *ListDevicesResponseBodyDataRecords {
	s.Vendor = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetGbId(v string) *ListDevicesResponseBodyDataRecords {
	s.GbId = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetCoverImageUrl(v string) *ListDevicesResponseBodyDataRecords {
	s.CoverImageUrl = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetAccessProtocolType(v string) *ListDevicesResponseBodyDataRecords {
	s.AccessProtocolType = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetDeviceSite(v string) *ListDevicesResponseBodyDataRecords {
	s.DeviceSite = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetLongitude(v string) *ListDevicesResponseBodyDataRecords {
	s.Longitude = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetLatitude(v string) *ListDevicesResponseBodyDataRecords {
	s.Latitude = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetResolution(v string) *ListDevicesResponseBodyDataRecords {
	s.Resolution = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetSipServerIp(v string) *ListDevicesResponseBodyDataRecords {
	s.SipServerIp = &v
	return s
}

func (s *ListDevicesResponseBodyDataRecords) SetBitRate(v string) *ListDevicesResponseBodyDataRecords {
	s.BitRate = &v
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

type ListEventAlgorithmDetailsRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceId     *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	RecordId     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	EventValue   *string `json:"EventValue,omitempty" xml:"EventValue,omitempty"`
	ExtendValue  *string `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty"`
}

func (s ListEventAlgorithmDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmDetailsRequest) SetCorpId(v string) *ListEventAlgorithmDetailsRequest {
	s.CorpId = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetEventType(v string) *ListEventAlgorithmDetailsRequest {
	s.EventType = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetDataSourceId(v string) *ListEventAlgorithmDetailsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetStartTime(v string) *ListEventAlgorithmDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetEndTime(v string) *ListEventAlgorithmDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetPageNumber(v int32) *ListEventAlgorithmDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetPageSize(v int32) *ListEventAlgorithmDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetSourceId(v string) *ListEventAlgorithmDetailsRequest {
	s.SourceId = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetRecordId(v string) *ListEventAlgorithmDetailsRequest {
	s.RecordId = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetEventValue(v string) *ListEventAlgorithmDetailsRequest {
	s.EventValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetExtendValue(v string) *ListEventAlgorithmDetailsRequest {
	s.ExtendValue = &v
	return s
}

type ListEventAlgorithmDetailsResponseBody struct {
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListEventAlgorithmDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventAlgorithmDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmDetailsResponseBody) SetTotalCount(v int32) *ListEventAlgorithmDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetRequestId(v string) *ListEventAlgorithmDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetMessage(v string) *ListEventAlgorithmDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetPageSize(v int32) *ListEventAlgorithmDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetPageNumber(v int32) *ListEventAlgorithmDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetData(v []*ListEventAlgorithmDetailsResponseBodyData) *ListEventAlgorithmDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetCode(v string) *ListEventAlgorithmDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBody) SetSuccess(v string) *ListEventAlgorithmDetailsResponseBody {
	s.Success = &v
	return s
}

type ListEventAlgorithmDetailsResponseBodyData struct {
	RightBottomY       *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceId       *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PicUrlPath         *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	RecordId           *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ExtendValue        *string `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty"`
	FaceCount          *string `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	RightBottomX       *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	SourceId           *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	ExtraExtendValue   *string `json:"ExtraExtendValue,omitempty" xml:"ExtraExtendValue,omitempty"`
	TargetPicUrlPath   *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	EventType          *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	PointY             *string `json:"PointY,omitempty" xml:"PointY,omitempty"`
	LeftTopY           *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	PointX             *string `json:"PointX,omitempty" xml:"PointX,omitempty"`
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	EventValue         *string `json:"EventValue,omitempty" xml:"EventValue,omitempty"`
	ShotTime           *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	LeftTopX           *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	TagCode            *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagCodeReliability *string `json:"TagCodeReliability,omitempty" xml:"TagCodeReliability,omitempty"`
	UuidCode           *string `json:"UuidCode,omitempty" xml:"UuidCode,omitempty"`
}

func (s ListEventAlgorithmDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetRightBottomY(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.RightBottomY = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetDataSourceId(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetPicUrlPath(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.PicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetRecordId(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.RecordId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetExtendValue(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.ExtendValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetFaceCount(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetRightBottomX(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.RightBottomX = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetSourceId(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetExtraExtendValue(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.ExtraExtendValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetTargetPicUrlPath(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetEventType(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.EventType = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetPointY(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.PointY = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetLeftTopY(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.LeftTopY = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetPointX(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.PointX = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetCorpId(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetEventValue(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.EventValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetShotTime(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.ShotTime = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetLeftTopX(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.LeftTopX = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetTagCode(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetTagCodeReliability(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.TagCodeReliability = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseBodyData) SetUuidCode(v string) *ListEventAlgorithmDetailsResponseBodyData {
	s.UuidCode = &v
	return s
}

type ListEventAlgorithmDetailsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEventAlgorithmDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventAlgorithmDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmDetailsResponse) SetHeaders(v map[string]*string) *ListEventAlgorithmDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetBody(v *ListEventAlgorithmDetailsResponseBody) *ListEventAlgorithmDetailsResponse {
	s.Body = v
	return s
}

type ListEventAlgorithmResultsRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ExtendValue  *string `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty"`
}

func (s ListEventAlgorithmResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsRequest) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsRequest) SetCorpId(v string) *ListEventAlgorithmResultsRequest {
	s.CorpId = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetEventType(v string) *ListEventAlgorithmResultsRequest {
	s.EventType = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetDataSourceId(v string) *ListEventAlgorithmResultsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetStartTime(v string) *ListEventAlgorithmResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetEndTime(v string) *ListEventAlgorithmResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetPageNumber(v string) *ListEventAlgorithmResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetPageSize(v string) *ListEventAlgorithmResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventAlgorithmResultsRequest) SetExtendValue(v string) *ListEventAlgorithmResultsRequest {
	s.ExtendValue = &v
	return s
}

type ListEventAlgorithmResultsResponseBody struct {
	ExtendValue *string                                    `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty"`
	Message     *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data        *ListEventAlgorithmResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code        *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListEventAlgorithmResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponseBody) SetExtendValue(v string) *ListEventAlgorithmResultsResponseBody {
	s.ExtendValue = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBody) SetMessage(v string) *ListEventAlgorithmResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBody) SetRequestId(v string) *ListEventAlgorithmResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBody) SetData(v *ListEventAlgorithmResultsResponseBodyData) *ListEventAlgorithmResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListEventAlgorithmResultsResponseBody) SetCode(v string) *ListEventAlgorithmResultsResponseBody {
	s.Code = &v
	return s
}

type ListEventAlgorithmResultsResponseBodyData struct {
	Records    []*ListEventAlgorithmResultsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEventAlgorithmResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponseBodyData) SetRecords(v []*ListEventAlgorithmResultsResponseBodyDataRecords) *ListEventAlgorithmResultsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyData) SetTotalPage(v int32) *ListEventAlgorithmResultsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyData) SetPageNumber(v int32) *ListEventAlgorithmResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyData) SetPageSize(v int32) *ListEventAlgorithmResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyData) SetTotalCount(v int32) *ListEventAlgorithmResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListEventAlgorithmResultsResponseBodyDataRecords struct {
	ExtendValueTwo     *string `json:"ExtendValueTwo,omitempty" xml:"ExtendValueTwo,omitempty"`
	RecordId           *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	PicUrlPath         *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	DataSourceId       *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	ExtendValue        *string `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty"`
	ExtendValueThree   *string `json:"ExtendValueThree,omitempty" xml:"ExtendValueThree,omitempty"`
	FaceCount          *string `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	TargetPicUrlPath   *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	EventType          *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ShotTime           *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	CapStyle           *string `json:"CapStyle,omitempty" xml:"CapStyle,omitempty"`
	TagCode            *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagCodeReliability *string `json:"TagCodeReliability,omitempty" xml:"TagCodeReliability,omitempty"`
	UuidCode           *string `json:"UuidCode,omitempty" xml:"UuidCode,omitempty"`
}

func (s ListEventAlgorithmResultsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetExtendValueTwo(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.ExtendValueTwo = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetRecordId(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.RecordId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetPicUrlPath(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetDataSourceId(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetExtendValue(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.ExtendValue = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetExtendValueThree(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.ExtendValueThree = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetFaceCount(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.FaceCount = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetTargetPicUrlPath(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetEventType(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.EventType = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetCorpId(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetShotTime(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetCapStyle(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.CapStyle = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetTagCode(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.TagCode = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetTagCodeReliability(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.TagCodeReliability = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseBodyDataRecords) SetUuidCode(v string) *ListEventAlgorithmResultsResponseBodyDataRecords {
	s.UuidCode = &v
	return s
}

type ListEventAlgorithmResultsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEventAlgorithmResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponse) SetHeaders(v map[string]*string) *ListEventAlgorithmResultsResponse {
	s.Headers = v
	return s
}

func (s *ListEventAlgorithmResultsResponse) SetBody(v *ListEventAlgorithmResultsResponseBody) *ListEventAlgorithmResultsResponse {
	s.Body = v
	return s
}

type ListFaceAlgorithmResultsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFaceAlgorithmResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsRequest) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsRequest) SetCorpId(v string) *ListFaceAlgorithmResultsRequest {
	s.CorpId = &v
	return s
}

func (s *ListFaceAlgorithmResultsRequest) SetAlgorithmType(v string) *ListFaceAlgorithmResultsRequest {
	s.AlgorithmType = &v
	return s
}

func (s *ListFaceAlgorithmResultsRequest) SetDataSourceId(v string) *ListFaceAlgorithmResultsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListFaceAlgorithmResultsRequest) SetStartTime(v string) *ListFaceAlgorithmResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListFaceAlgorithmResultsRequest) SetEndTime(v string) *ListFaceAlgorithmResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListFaceAlgorithmResultsRequest) SetPageNumber(v string) *ListFaceAlgorithmResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFaceAlgorithmResultsRequest) SetPageSize(v string) *ListFaceAlgorithmResultsRequest {
	s.PageSize = &v
	return s
}

type ListFaceAlgorithmResultsResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListFaceAlgorithmResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListFaceAlgorithmResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponseBody) SetMessage(v string) *ListFaceAlgorithmResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBody) SetRequestId(v string) *ListFaceAlgorithmResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBody) SetData(v *ListFaceAlgorithmResultsResponseBodyData) *ListFaceAlgorithmResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBody) SetCode(v string) *ListFaceAlgorithmResultsResponseBody {
	s.Code = &v
	return s
}

type ListFaceAlgorithmResultsResponseBodyData struct {
	Records    []*ListFaceAlgorithmResultsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFaceAlgorithmResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponseBodyData) SetRecords(v []*ListFaceAlgorithmResultsResponseBodyDataRecords) *ListFaceAlgorithmResultsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyData) SetTotalPage(v int32) *ListFaceAlgorithmResultsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyData) SetPageNumber(v int32) *ListFaceAlgorithmResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyData) SetPageSize(v int32) *ListFaceAlgorithmResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyData) SetTotalCount(v int32) *ListFaceAlgorithmResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListFaceAlgorithmResultsResponseBodyDataRecords struct {
	RightBottomY     *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceId     *string  `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PicUrlPath       *string  `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	FaceId           *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	RightBottomX     *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	SourceId         *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	MaxAge           *string  `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	TargetPicUrlPath *string  `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	HairStyle        *string  `json:"HairStyle,omitempty" xml:"HairStyle,omitempty"`
	LeftTopY         *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	MinAge           *string  `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	CorpId           *string  `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ShotTime         *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	GenderCode       *string  `json:"GenderCode,omitempty" xml:"GenderCode,omitempty"`
	CapStyle         *string  `json:"CapStyle,omitempty" xml:"CapStyle,omitempty"`
	LeftTopX         *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s ListFaceAlgorithmResultsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetRightBottomY(v float32) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetDataSourceId(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetPicUrlPath(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetFaceId(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.FaceId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetRightBottomX(v float32) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetSourceId(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.SourceId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetMaxAge(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.MaxAge = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetTargetPicUrlPath(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetHairStyle(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.HairStyle = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetLeftTopY(v float32) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetMinAge(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.MinAge = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetCorpId(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetShotTime(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetGenderCode(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.GenderCode = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetCapStyle(v string) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.CapStyle = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseBodyDataRecords) SetLeftTopX(v float32) *ListFaceAlgorithmResultsResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

type ListFaceAlgorithmResultsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceAlgorithmResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponse) SetHeaders(v map[string]*string) *ListFaceAlgorithmResultsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceAlgorithmResultsResponse) SetBody(v *ListFaceAlgorithmResultsResponseBody) *ListFaceAlgorithmResultsResponse {
	s.Body = v
	return s
}

type ListMetricsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	AggregateType *string `json:"AggregateType,omitempty" xml:"AggregateType,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListMetricsRequest) SetCorpId(v string) *ListMetricsRequest {
	s.CorpId = &v
	return s
}

func (s *ListMetricsRequest) SetTagCode(v string) *ListMetricsRequest {
	s.TagCode = &v
	return s
}

func (s *ListMetricsRequest) SetAggregateType(v string) *ListMetricsRequest {
	s.AggregateType = &v
	return s
}

func (s *ListMetricsRequest) SetStartTime(v string) *ListMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMetricsRequest) SetEndTime(v string) *ListMetricsRequest {
	s.EndTime = &v
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

type ListMetricsResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseBody) SetMessage(v string) *ListMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMetricsResponseBody) SetRequestId(v string) *ListMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetricsResponseBody) SetData(v *ListMetricsResponseBodyData) *ListMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ListMetricsResponseBody) SetCode(v string) *ListMetricsResponseBody {
	s.Code = &v
	return s
}

type ListMetricsResponseBodyData struct {
	Records    []*ListMetricsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseBodyData) SetRecords(v []*ListMetricsResponseBodyDataRecords) *ListMetricsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListMetricsResponseBodyData) SetTotalPage(v int32) *ListMetricsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListMetricsResponseBodyData) SetPageNumber(v int32) *ListMetricsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMetricsResponseBodyData) SetPageSize(v int32) *ListMetricsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMetricsResponseBodyData) SetTotalCount(v int32) *ListMetricsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListMetricsResponseBodyDataRecords struct {
	TagMetric *string `json:"TagMetric,omitempty" xml:"TagMetric,omitempty"`
	TagCode   *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagValue  *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	DateTime  *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s ListMetricsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseBodyDataRecords) SetTagMetric(v string) *ListMetricsResponseBodyDataRecords {
	s.TagMetric = &v
	return s
}

func (s *ListMetricsResponseBodyDataRecords) SetTagCode(v string) *ListMetricsResponseBodyDataRecords {
	s.TagCode = &v
	return s
}

func (s *ListMetricsResponseBodyDataRecords) SetTagValue(v string) *ListMetricsResponseBodyDataRecords {
	s.TagValue = &v
	return s
}

func (s *ListMetricsResponseBodyDataRecords) SetDateTime(v string) *ListMetricsResponseBodyDataRecords {
	s.DateTime = &v
	return s
}

type ListMetricsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListMetricsResponse) SetBody(v *ListMetricsResponseBody) *ListMetricsResponse {
	s.Body = v
	return s
}

type ListMotorAlgorithmResultsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlateNumber   *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
}

func (s ListMotorAlgorithmResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsRequest) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsRequest) SetCorpId(v string) *ListMotorAlgorithmResultsRequest {
	s.CorpId = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetAlgorithmType(v string) *ListMotorAlgorithmResultsRequest {
	s.AlgorithmType = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetDataSourceId(v string) *ListMotorAlgorithmResultsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetStartTime(v string) *ListMotorAlgorithmResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetEndTime(v string) *ListMotorAlgorithmResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetPageNumber(v string) *ListMotorAlgorithmResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetPageSize(v string) *ListMotorAlgorithmResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMotorAlgorithmResultsRequest) SetPlateNumber(v string) *ListMotorAlgorithmResultsRequest {
	s.PlateNumber = &v
	return s
}

type ListMotorAlgorithmResultsResponseBody struct {
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListMotorAlgorithmResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListMotorAlgorithmResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponseBody) SetMessage(v string) *ListMotorAlgorithmResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBody) SetRequestId(v string) *ListMotorAlgorithmResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBody) SetData(v *ListMotorAlgorithmResultsResponseBodyData) *ListMotorAlgorithmResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBody) SetCode(v string) *ListMotorAlgorithmResultsResponseBody {
	s.Code = &v
	return s
}

type ListMotorAlgorithmResultsResponseBodyData struct {
	Records    []*ListMotorAlgorithmResultsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMotorAlgorithmResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponseBodyData) SetRecords(v []*ListMotorAlgorithmResultsResponseBodyDataRecords) *ListMotorAlgorithmResultsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyData) SetTotalPage(v int32) *ListMotorAlgorithmResultsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyData) SetPageNumber(v int32) *ListMotorAlgorithmResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyData) SetPageSize(v int32) *ListMotorAlgorithmResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyData) SetTotalCount(v int32) *ListMotorAlgorithmResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListMotorAlgorithmResultsResponseBodyDataRecords struct {
	MotorClass       *string  `json:"MotorClass,omitempty" xml:"MotorClass,omitempty"`
	RightBottomY     *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceId     *string  `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PicUrlPath       *string  `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	PlateClass       *string  `json:"PlateClass,omitempty" xml:"PlateClass,omitempty"`
	PlateColor       *string  `json:"PlateColor,omitempty" xml:"PlateColor,omitempty"`
	RightBottomX     *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	SourceId         *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SafetyBelt       *string  `json:"SafetyBelt,omitempty" xml:"SafetyBelt,omitempty"`
	MotorStyle       *string  `json:"MotorStyle,omitempty" xml:"MotorStyle,omitempty"`
	TargetPicUrlPath *string  `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	LeftTopY         *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	MotorColor       *string  `json:"MotorColor,omitempty" xml:"MotorColor,omitempty"`
	PlateNumber      *string  `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	CorpId           *string  `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ShotTime         *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	Calling          *string  `json:"Calling,omitempty" xml:"Calling,omitempty"`
	LeftTopX         *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	MotorBrand       *string  `json:"MotorBrand,omitempty" xml:"MotorBrand,omitempty"`
	MotorModel       *string  `json:"MotorModel,omitempty" xml:"MotorModel,omitempty"`
	MotorId          *string  `json:"MotorId,omitempty" xml:"MotorId,omitempty"`
}

func (s ListMotorAlgorithmResultsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetMotorClass(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.MotorClass = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetRightBottomY(v float32) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetDataSourceId(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetPicUrlPath(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetPlateClass(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.PlateClass = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetPlateColor(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.PlateColor = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetRightBottomX(v float32) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetSourceId(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.SourceId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetSafetyBelt(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.SafetyBelt = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetMotorStyle(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.MotorStyle = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetTargetPicUrlPath(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetLeftTopY(v float32) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetMotorColor(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.MotorColor = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetPlateNumber(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.PlateNumber = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetCorpId(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetShotTime(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetCalling(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.Calling = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetLeftTopX(v float32) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetMotorBrand(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.MotorBrand = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetMotorModel(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.MotorModel = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseBodyDataRecords) SetMotorId(v string) *ListMotorAlgorithmResultsResponseBodyDataRecords {
	s.MotorId = &v
	return s
}

type ListMotorAlgorithmResultsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMotorAlgorithmResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMotorAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponse) SetHeaders(v map[string]*string) *ListMotorAlgorithmResultsResponse {
	s.Headers = v
	return s
}

func (s *ListMotorAlgorithmResultsResponse) SetBody(v *ListMotorAlgorithmResultsResponseBody) *ListMotorAlgorithmResultsResponse {
	s.Body = v
	return s
}

type ListNVRChannelDeviceRequest struct {
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	IsPage     *string `json:"IsPage,omitempty" xml:"IsPage,omitempty"`
	PageNum    *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNVRChannelDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNVRChannelDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListNVRChannelDeviceRequest) SetDeviceCode(v string) *ListNVRChannelDeviceRequest {
	s.DeviceCode = &v
	return s
}

func (s *ListNVRChannelDeviceRequest) SetIsPage(v string) *ListNVRChannelDeviceRequest {
	s.IsPage = &v
	return s
}

func (s *ListNVRChannelDeviceRequest) SetPageNum(v string) *ListNVRChannelDeviceRequest {
	s.PageNum = &v
	return s
}

func (s *ListNVRChannelDeviceRequest) SetPageSize(v string) *ListNVRChannelDeviceRequest {
	s.PageSize = &v
	return s
}

type ListNVRChannelDeviceResponseBody struct {
	// Id of the request
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *string                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	Data      []*ListNVRChannelDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListNVRChannelDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNVRChannelDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListNVRChannelDeviceResponseBody) SetRequestId(v string) *ListNVRChannelDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBody) SetSuccess(v bool) *ListNVRChannelDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBody) SetTotal(v string) *ListNVRChannelDeviceResponseBody {
	s.Total = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBody) SetData(v []*ListNVRChannelDeviceResponseBodyData) *ListNVRChannelDeviceResponseBody {
	s.Data = v
	return s
}

type ListNVRChannelDeviceResponseBodyData struct {
	DeviceCode     *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	DeviceName     *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType     *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	DeviceStatus   *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	StreamStatus   *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
	ComptureStatus *string `json:"ComptureStatus,omitempty" xml:"ComptureStatus,omitempty"`
	DeviceSn       *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	SampleName     *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	RegionName     *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListNVRChannelDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNVRChannelDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNVRChannelDeviceResponseBodyData) SetDeviceCode(v string) *ListNVRChannelDeviceResponseBodyData {
	s.DeviceCode = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetDeviceName(v string) *ListNVRChannelDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetDeviceType(v string) *ListNVRChannelDeviceResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetDatasourceType(v string) *ListNVRChannelDeviceResponseBodyData {
	s.DatasourceType = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetDeviceStatus(v string) *ListNVRChannelDeviceResponseBodyData {
	s.DeviceStatus = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetStreamStatus(v string) *ListNVRChannelDeviceResponseBodyData {
	s.StreamStatus = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetComptureStatus(v string) *ListNVRChannelDeviceResponseBodyData {
	s.ComptureStatus = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetDeviceSn(v string) *ListNVRChannelDeviceResponseBodyData {
	s.DeviceSn = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetSampleName(v string) *ListNVRChannelDeviceResponseBodyData {
	s.SampleName = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetRegionName(v string) *ListNVRChannelDeviceResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListNVRChannelDeviceResponseBodyData) SetCorpId(v string) *ListNVRChannelDeviceResponseBodyData {
	s.CorpId = &v
	return s
}

type ListNVRChannelDeviceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNVRChannelDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNVRChannelDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNVRChannelDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListNVRChannelDeviceResponse) SetHeaders(v map[string]*string) *ListNVRChannelDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListNVRChannelDeviceResponse) SetBody(v *ListNVRChannelDeviceResponseBody) *ListNVRChannelDeviceResponse {
	s.Body = v
	return s
}

type ListNVRDeviceRequest struct {
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	CorpIdList *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
	IsPage     *int64  `json:"IsPage,omitempty" xml:"IsPage,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNVRDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNVRDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListNVRDeviceRequest) SetDeviceCode(v string) *ListNVRDeviceRequest {
	s.DeviceCode = &v
	return s
}

func (s *ListNVRDeviceRequest) SetCorpIdList(v string) *ListNVRDeviceRequest {
	s.CorpIdList = &v
	return s
}

func (s *ListNVRDeviceRequest) SetIsPage(v int64) *ListNVRDeviceRequest {
	s.IsPage = &v
	return s
}

func (s *ListNVRDeviceRequest) SetPageNum(v int64) *ListNVRDeviceRequest {
	s.PageNum = &v
	return s
}

func (s *ListNVRDeviceRequest) SetPageSize(v int64) *ListNVRDeviceRequest {
	s.PageSize = &v
	return s
}

type ListNVRDeviceResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *string                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Data      []*ListNVRDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListNVRDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNVRDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListNVRDeviceResponseBody) SetRequestId(v string) *ListNVRDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNVRDeviceResponseBody) SetSuccess(v bool) *ListNVRDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *ListNVRDeviceResponseBody) SetTotal(v string) *ListNVRDeviceResponseBody {
	s.Total = &v
	return s
}

func (s *ListNVRDeviceResponseBody) SetData(v []*ListNVRDeviceResponseBodyData) *ListNVRDeviceResponseBody {
	s.Data = v
	return s
}

type ListNVRDeviceResponseBodyData struct {
	DeviceCode       *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType       *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DatasourceType   *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	DeviceStatus     *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	StreamStatus     *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
	ComptureStatus   *string `json:"ComptureStatus,omitempty" xml:"ComptureStatus,omitempty"`
	RegionName       *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ProjectName      *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegistrationTime *string `json:"RegistrationTime,omitempty" xml:"RegistrationTime,omitempty"`
	AccessQuota      *string `json:"AccessQuota,omitempty" xml:"AccessQuota,omitempty"`
	Channel          *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	DeviceSn         *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListNVRDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNVRDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNVRDeviceResponseBodyData) SetDeviceCode(v string) *ListNVRDeviceResponseBodyData {
	s.DeviceCode = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetDeviceName(v string) *ListNVRDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetDeviceType(v string) *ListNVRDeviceResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetDatasourceType(v string) *ListNVRDeviceResponseBodyData {
	s.DatasourceType = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetDeviceStatus(v string) *ListNVRDeviceResponseBodyData {
	s.DeviceStatus = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetStreamStatus(v string) *ListNVRDeviceResponseBodyData {
	s.StreamStatus = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetComptureStatus(v string) *ListNVRDeviceResponseBodyData {
	s.ComptureStatus = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetRegionName(v string) *ListNVRDeviceResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetProjectName(v string) *ListNVRDeviceResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetRegistrationTime(v string) *ListNVRDeviceResponseBodyData {
	s.RegistrationTime = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetAccessQuota(v string) *ListNVRDeviceResponseBodyData {
	s.AccessQuota = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetChannel(v string) *ListNVRDeviceResponseBodyData {
	s.Channel = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetDeviceSn(v string) *ListNVRDeviceResponseBodyData {
	s.DeviceSn = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetType(v string) *ListNVRDeviceResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListNVRDeviceResponseBodyData) SetCorpId(v string) *ListNVRDeviceResponseBodyData {
	s.CorpId = &v
	return s
}

type ListNVRDeviceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNVRDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNVRDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNVRDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListNVRDeviceResponse) SetHeaders(v map[string]*string) *ListNVRDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListNVRDeviceResponse) SetBody(v *ListNVRDeviceResponseBody) *ListNVRDeviceResponse {
	s.Body = v
	return s
}

type ListPersonsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNo        *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
}

func (s ListPersonsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsRequest) GoString() string {
	return s.String()
}

func (s *ListPersonsRequest) SetCorpId(v string) *ListPersonsRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonsRequest) SetPageNo(v string) *ListPersonsRequest {
	s.PageNo = &v
	return s
}

func (s *ListPersonsRequest) SetPageSize(v string) *ListPersonsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonsRequest) SetStartTime(v string) *ListPersonsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonsRequest) SetEndTime(v string) *ListPersonsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonsRequest) SetAlgorithmType(v string) *ListPersonsRequest {
	s.AlgorithmType = &v
	return s
}

type ListPersonsResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListPersonsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseBody) SetMessage(v string) *ListPersonsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonsResponseBody) SetRequestId(v string) *ListPersonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonsResponseBody) SetData(v *ListPersonsResponseBodyData) *ListPersonsResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonsResponseBody) SetCode(v string) *ListPersonsResponseBody {
	s.Code = &v
	return s
}

type ListPersonsResponseBodyData struct {
	Records    []*ListPersonsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNo     *string                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TotalPage  *string                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *string                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPersonsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseBodyData) SetRecords(v []*ListPersonsResponseBodyDataRecords) *ListPersonsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListPersonsResponseBodyData) SetPageNo(v string) *ListPersonsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListPersonsResponseBodyData) SetTotalPage(v string) *ListPersonsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListPersonsResponseBodyData) SetPageSize(v string) *ListPersonsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPersonsResponseBodyData) SetTotalCount(v string) *ListPersonsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListPersonsResponseBodyDataRecords struct {
	PicUrl          *string                                      `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	TagList         []*ListPersonsResponseBodyDataRecordsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	FirstAppearTime *string                                      `json:"FirstAppearTime,omitempty" xml:"FirstAppearTime,omitempty"`
	PersonId        *string                                      `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListPersonsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseBodyDataRecords) SetPicUrl(v string) *ListPersonsResponseBodyDataRecords {
	s.PicUrl = &v
	return s
}

func (s *ListPersonsResponseBodyDataRecords) SetTagList(v []*ListPersonsResponseBodyDataRecordsTagList) *ListPersonsResponseBodyDataRecords {
	s.TagList = v
	return s
}

func (s *ListPersonsResponseBodyDataRecords) SetFirstAppearTime(v string) *ListPersonsResponseBodyDataRecords {
	s.FirstAppearTime = &v
	return s
}

func (s *ListPersonsResponseBodyDataRecords) SetPersonId(v string) *ListPersonsResponseBodyDataRecords {
	s.PersonId = &v
	return s
}

type ListPersonsResponseBodyDataRecordsTagList struct {
	TagValueId *string `json:"TagValueId,omitempty" xml:"TagValueId,omitempty"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagValue   *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPersonsResponseBodyDataRecordsTagList) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseBodyDataRecordsTagList) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseBodyDataRecordsTagList) SetTagValueId(v string) *ListPersonsResponseBodyDataRecordsTagList {
	s.TagValueId = &v
	return s
}

func (s *ListPersonsResponseBodyDataRecordsTagList) SetTagName(v string) *ListPersonsResponseBodyDataRecordsTagList {
	s.TagName = &v
	return s
}

func (s *ListPersonsResponseBodyDataRecordsTagList) SetTagCode(v string) *ListPersonsResponseBodyDataRecordsTagList {
	s.TagCode = &v
	return s
}

func (s *ListPersonsResponseBodyDataRecordsTagList) SetTagValue(v string) *ListPersonsResponseBodyDataRecordsTagList {
	s.TagValue = &v
	return s
}

type ListPersonsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponse) GoString() string {
	return s.String()
}

func (s *ListPersonsResponse) SetHeaders(v map[string]*string) *ListPersonsResponse {
	s.Headers = v
	return s
}

func (s *ListPersonsResponse) SetBody(v *ListPersonsResponseBody) *ListPersonsResponse {
	s.Body = v
	return s
}

type ListPersonTraceRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListPersonTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceRequest) GoString() string {
	return s.String()
}

func (s *ListPersonTraceRequest) SetStartTime(v string) *ListPersonTraceRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonTraceRequest) SetCorpId(v string) *ListPersonTraceRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonTraceRequest) SetPageNumber(v string) *ListPersonTraceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTraceRequest) SetPageSize(v string) *ListPersonTraceRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonTraceRequest) SetEndTime(v string) *ListPersonTraceRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonTraceRequest) SetDataSourceId(v string) *ListPersonTraceRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListPersonTraceRequest) SetPersonId(v string) *ListPersonTraceRequest {
	s.PersonId = &v
	return s
}

func (s *ListPersonTraceRequest) SetGroupId(v string) *ListPersonTraceRequest {
	s.GroupId = &v
	return s
}

type ListPersonTraceResponseBody struct {
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPersonTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonTraceResponseBody) SetTotalCount(v int32) *ListPersonTraceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonTraceResponseBody) SetRequestId(v string) *ListPersonTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonTraceResponseBody) SetMessage(v string) *ListPersonTraceResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonTraceResponseBody) SetPageSize(v int32) *ListPersonTraceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonTraceResponseBody) SetPageNumber(v int32) *ListPersonTraceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTraceResponseBody) SetData(v []*ListPersonTraceResponseBodyData) *ListPersonTraceResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonTraceResponseBody) SetCode(v string) *ListPersonTraceResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersonTraceResponseBody) SetSuccess(v string) *ListPersonTraceResponseBody {
	s.Success = &v
	return s
}

type ListPersonTraceResponseBodyData struct {
	EndTargetImage   *string `json:"EndTargetImage,omitempty" xml:"EndTargetImage,omitempty"`
	LastTime         *string `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	EndSourceImage   *string `json:"EndSourceImage,omitempty" xml:"EndSourceImage,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	StartSourceImage *string `json:"StartSourceImage,omitempty" xml:"StartSourceImage,omitempty"`
	Date             *string `json:"Date,omitempty" xml:"Date,omitempty"`
	PersonId         *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	StartTargetImage *string `json:"StartTargetImage,omitempty" xml:"StartTargetImage,omitempty"`
}

func (s ListPersonTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonTraceResponseBodyData) SetEndTargetImage(v string) *ListPersonTraceResponseBodyData {
	s.EndTargetImage = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetLastTime(v string) *ListPersonTraceResponseBodyData {
	s.LastTime = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetStartTime(v string) *ListPersonTraceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetGroupId(v string) *ListPersonTraceResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetDeviceId(v string) *ListPersonTraceResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetEndSourceImage(v string) *ListPersonTraceResponseBodyData {
	s.EndSourceImage = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetCorpId(v string) *ListPersonTraceResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetStartSourceImage(v string) *ListPersonTraceResponseBodyData {
	s.StartSourceImage = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetDate(v string) *ListPersonTraceResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetPersonId(v string) *ListPersonTraceResponseBodyData {
	s.PersonId = &v
	return s
}

func (s *ListPersonTraceResponseBodyData) SetStartTargetImage(v string) *ListPersonTraceResponseBodyData {
	s.StartTargetImage = &v
	return s
}

type ListPersonTraceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceResponse) GoString() string {
	return s.String()
}

func (s *ListPersonTraceResponse) SetHeaders(v map[string]*string) *ListPersonTraceResponse {
	s.Headers = v
	return s
}

func (s *ListPersonTraceResponse) SetBody(v *ListPersonTraceResponseBody) *ListPersonTraceResponse {
	s.Body = v
	return s
}

type ListPersonTraceDetailsRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubId        *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s ListPersonTraceDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListPersonTraceDetailsRequest) SetCorpId(v string) *ListPersonTraceDetailsRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetPageNumber(v int64) *ListPersonTraceDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetPageSize(v int64) *ListPersonTraceDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetEndTime(v string) *ListPersonTraceDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetPersonId(v string) *ListPersonTraceDetailsRequest {
	s.PersonId = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetStartTime(v string) *ListPersonTraceDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetSubId(v string) *ListPersonTraceDetailsRequest {
	s.SubId = &v
	return s
}

func (s *ListPersonTraceDetailsRequest) SetDataSourceId(v string) *ListPersonTraceDetailsRequest {
	s.DataSourceId = &v
	return s
}

type ListPersonTraceDetailsResponseBody struct {
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonTraceDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonTraceDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonTraceDetailsResponseBody) SetTotalCount(v int64) *ListPersonTraceDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBody) SetRequestId(v string) *ListPersonTraceDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBody) SetMessage(v string) *ListPersonTraceDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBody) SetPageSize(v int64) *ListPersonTraceDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBody) SetPageNumber(v int64) *ListPersonTraceDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBody) SetData(v []*ListPersonTraceDetailsResponseBodyData) *ListPersonTraceDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonTraceDetailsResponseBody) SetCode(v string) *ListPersonTraceDetailsResponseBody {
	s.Code = &v
	return s
}

type ListPersonTraceDetailsResponseBodyData struct {
	TargetPicUrlPath *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	SubId            *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	PicUrlPath       *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ShotTime         *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	PersonId         *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListPersonTraceDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonTraceDetailsResponseBodyData) SetTargetPicUrlPath(v string) *ListPersonTraceDetailsResponseBodyData {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetSubId(v string) *ListPersonTraceDetailsResponseBodyData {
	s.SubId = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetRightBottomY(v string) *ListPersonTraceDetailsResponseBodyData {
	s.RightBottomY = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetLeftTopY(v string) *ListPersonTraceDetailsResponseBodyData {
	s.LeftTopY = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetPicUrlPath(v string) *ListPersonTraceDetailsResponseBodyData {
	s.PicUrlPath = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetDataSourceId(v string) *ListPersonTraceDetailsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetCorpId(v string) *ListPersonTraceDetailsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetShotTime(v string) *ListPersonTraceDetailsResponseBodyData {
	s.ShotTime = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetLeftTopX(v string) *ListPersonTraceDetailsResponseBodyData {
	s.LeftTopX = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetRightBottomX(v string) *ListPersonTraceDetailsResponseBodyData {
	s.RightBottomX = &v
	return s
}

func (s *ListPersonTraceDetailsResponseBodyData) SetPersonId(v string) *ListPersonTraceDetailsResponseBodyData {
	s.PersonId = &v
	return s
}

type ListPersonTraceDetailsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonTraceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonTraceDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListPersonTraceDetailsResponse) SetHeaders(v map[string]*string) *ListPersonTraceDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListPersonTraceDetailsResponse) SetBody(v *ListPersonTraceDetailsResponseBody) *ListPersonTraceDetailsResponse {
	s.Body = v
	return s
}

type ListPersonVisitCountRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AggregateType     *string `json:"AggregateType,omitempty" xml:"AggregateType,omitempty"`
	TagCode           *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TimeAggregateType *string `json:"TimeAggregateType,omitempty" xml:"TimeAggregateType,omitempty"`
	MinVal            *int32  `json:"MinVal,omitempty" xml:"MinVal,omitempty"`
	MaxVal            *int32  `json:"MaxVal,omitempty" xml:"MaxVal,omitempty"`
	CountType         *string `json:"CountType,omitempty" xml:"CountType,omitempty"`
}

func (s ListPersonVisitCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonVisitCountRequest) GoString() string {
	return s.String()
}

func (s *ListPersonVisitCountRequest) SetCorpId(v string) *ListPersonVisitCountRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetPageNumber(v int32) *ListPersonVisitCountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetPageSize(v int32) *ListPersonVisitCountRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetStartTime(v string) *ListPersonVisitCountRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetEndTime(v string) *ListPersonVisitCountRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetAggregateType(v string) *ListPersonVisitCountRequest {
	s.AggregateType = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetTagCode(v string) *ListPersonVisitCountRequest {
	s.TagCode = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetTimeAggregateType(v string) *ListPersonVisitCountRequest {
	s.TimeAggregateType = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetMinVal(v int32) *ListPersonVisitCountRequest {
	s.MinVal = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetMaxVal(v int32) *ListPersonVisitCountRequest {
	s.MaxVal = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetCountType(v string) *ListPersonVisitCountRequest {
	s.CountType = &v
	return s
}

type ListPersonVisitCountResponseBody struct {
	TotalCount *string                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo     *string                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Data       []*ListPersonVisitCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPersonVisitCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonVisitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonVisitCountResponseBody) SetTotalCount(v string) *ListPersonVisitCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetRequestId(v string) *ListPersonVisitCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetMessage(v string) *ListPersonVisitCountResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetPageSize(v string) *ListPersonVisitCountResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetPageNo(v string) *ListPersonVisitCountResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetData(v []*ListPersonVisitCountResponseBodyData) *ListPersonVisitCountResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetCode(v string) *ListPersonVisitCountResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersonVisitCountResponseBody) SetSuccess(v string) *ListPersonVisitCountResponseBody {
	s.Success = &v
	return s
}

type ListPersonVisitCountResponseBodyData struct {
	DayId      *string `json:"DayId,omitempty" xml:"DayId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagMetrics *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty"`
	HourId     *string `json:"HourId,omitempty" xml:"HourId,omitempty"`
	PersonId   *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListPersonVisitCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonVisitCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonVisitCountResponseBodyData) SetDayId(v string) *ListPersonVisitCountResponseBodyData {
	s.DayId = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetGroupId(v string) *ListPersonVisitCountResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetDeviceId(v string) *ListPersonVisitCountResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetTagCode(v string) *ListPersonVisitCountResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetCorpId(v string) *ListPersonVisitCountResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetTagMetrics(v string) *ListPersonVisitCountResponseBodyData {
	s.TagMetrics = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetHourId(v string) *ListPersonVisitCountResponseBodyData {
	s.HourId = &v
	return s
}

func (s *ListPersonVisitCountResponseBodyData) SetPersonId(v string) *ListPersonVisitCountResponseBodyData {
	s.PersonId = &v
	return s
}

type ListPersonVisitCountResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonVisitCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonVisitCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonVisitCountResponse) GoString() string {
	return s.String()
}

func (s *ListPersonVisitCountResponse) SetHeaders(v map[string]*string) *ListPersonVisitCountResponse {
	s.Headers = v
	return s
}

func (s *ListPersonVisitCountResponse) SetBody(v *ListPersonVisitCountResponseBody) *ListPersonVisitCountResponse {
	s.Body = v
	return s
}

type ListSubscribeDeviceRequest struct {
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubscribeDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscribeDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListSubscribeDeviceRequest) SetPageNum(v int32) *ListSubscribeDeviceRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubscribeDeviceRequest) SetPageSize(v int32) *ListSubscribeDeviceRequest {
	s.PageSize = &v
	return s
}

type ListSubscribeDeviceResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListSubscribeDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListSubscribeDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscribeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscribeDeviceResponseBody) SetMessage(v string) *ListSubscribeDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscribeDeviceResponseBody) SetRequestId(v string) *ListSubscribeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscribeDeviceResponseBody) SetData(v *ListSubscribeDeviceResponseBodyData) *ListSubscribeDeviceResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscribeDeviceResponseBody) SetCode(v string) *ListSubscribeDeviceResponseBody {
	s.Code = &v
	return s
}

type ListSubscribeDeviceResponseBodyData struct {
	SubscribeList []*ListSubscribeDeviceResponseBodyDataSubscribeList `json:"SubscribeList,omitempty" xml:"SubscribeList,omitempty" type:"Repeated"`
	TotalCount    *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubscribeDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscribeDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscribeDeviceResponseBodyData) SetSubscribeList(v []*ListSubscribeDeviceResponseBodyDataSubscribeList) *ListSubscribeDeviceResponseBodyData {
	s.SubscribeList = v
	return s
}

func (s *ListSubscribeDeviceResponseBodyData) SetTotalCount(v int32) *ListSubscribeDeviceResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSubscribeDeviceResponseBodyDataSubscribeList struct {
	PushConfig *string `json:"PushConfig,omitempty" xml:"PushConfig,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSubscribeDeviceResponseBodyDataSubscribeList) String() string {
	return tea.Prettify(s)
}

func (s ListSubscribeDeviceResponseBodyDataSubscribeList) GoString() string {
	return s.String()
}

func (s *ListSubscribeDeviceResponseBodyDataSubscribeList) SetPushConfig(v string) *ListSubscribeDeviceResponseBodyDataSubscribeList {
	s.PushConfig = &v
	return s
}

func (s *ListSubscribeDeviceResponseBodyDataSubscribeList) SetUpdateTime(v string) *ListSubscribeDeviceResponseBodyDataSubscribeList {
	s.UpdateTime = &v
	return s
}

func (s *ListSubscribeDeviceResponseBodyDataSubscribeList) SetDeviceId(v string) *ListSubscribeDeviceResponseBodyDataSubscribeList {
	s.DeviceId = &v
	return s
}

func (s *ListSubscribeDeviceResponseBodyDataSubscribeList) SetCreateTime(v string) *ListSubscribeDeviceResponseBodyDataSubscribeList {
	s.CreateTime = &v
	return s
}

func (s *ListSubscribeDeviceResponseBodyDataSubscribeList) SetUserId(v string) *ListSubscribeDeviceResponseBodyDataSubscribeList {
	s.UserId = &v
	return s
}

type ListSubscribeDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubscribeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubscribeDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscribeDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListSubscribeDeviceResponse) SetHeaders(v map[string]*string) *ListSubscribeDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListSubscribeDeviceResponse) SetBody(v *ListSubscribeDeviceResponseBody) *ListSubscribeDeviceResponse {
	s.Body = v
	return s
}

type ListUserGroupsRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) SetCorpId(v string) *ListUserGroupsRequest {
	s.CorpId = &v
	return s
}

func (s *ListUserGroupsRequest) SetIsvSubId(v string) *ListUserGroupsRequest {
	s.IsvSubId = &v
	return s
}

type ListUserGroupsResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListUserGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) SetMessage(v string) *ListUserGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetData(v []*ListUserGroupsResponseBodyData) *ListUserGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserGroupsResponseBody) SetCode(v string) *ListUserGroupsResponseBody {
	s.Code = &v
	return s
}

type ListUserGroupsResponseBodyData struct {
	UpdateTime        *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	IsvSubId          *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId       *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	UserCount         *int64  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	ParentUserGroupId *int64  `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	Creator           *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
}

func (s ListUserGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyData) SetUpdateTime(v string) *ListUserGroupsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetIsvSubId(v string) *ListUserGroupsResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetUserGroupId(v int64) *ListUserGroupsResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetCreateTime(v string) *ListUserGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetUserGroupName(v string) *ListUserGroupsResponseBodyData {
	s.UserGroupName = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetUserCount(v int64) *ListUserGroupsResponseBodyData {
	s.UserCount = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetParentUserGroupId(v int64) *ListUserGroupsResponseBodyData {
	s.ParentUserGroupId = &v
	return s
}

func (s *ListUserGroupsResponseBodyData) SetCreator(v string) *ListUserGroupsResponseBodyData {
	s.Creator = &v
	return s
}

type ListUserGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponse) SetHeaders(v map[string]*string) *ListUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsResponse) SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	CorpId                *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId              *string                `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserName              *string                `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId           *int64                 `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdNumber              *string                `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl          *string                `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Address               *string                `json:"Address,omitempty" xml:"Address,omitempty"`
	Age                   *int32                 `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender                *int32                 `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string                `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string                `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment            *string                `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId                 *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PersonList            map[string]interface{} `json:"PersonList,omitempty" xml:"PersonList,omitempty"`
	UserList              map[string]interface{} `json:"UserList,omitempty" xml:"UserList,omitempty"`
	MatchingRateThreshold *string                `json:"MatchingRateThreshold,omitempty" xml:"MatchingRateThreshold,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetCorpId(v string) *ListUsersRequest {
	s.CorpId = &v
	return s
}

func (s *ListUsersRequest) SetIsvSubId(v string) *ListUsersRequest {
	s.IsvSubId = &v
	return s
}

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

func (s *ListUsersRequest) SetUserGroupId(v int64) *ListUsersRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListUsersRequest) SetIdNumber(v string) *ListUsersRequest {
	s.IdNumber = &v
	return s
}

func (s *ListUsersRequest) SetFaceImageUrl(v string) *ListUsersRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *ListUsersRequest) SetAddress(v string) *ListUsersRequest {
	s.Address = &v
	return s
}

func (s *ListUsersRequest) SetAge(v int32) *ListUsersRequest {
	s.Age = &v
	return s
}

func (s *ListUsersRequest) SetGender(v int32) *ListUsersRequest {
	s.Gender = &v
	return s
}

func (s *ListUsersRequest) SetPlateNo(v string) *ListUsersRequest {
	s.PlateNo = &v
	return s
}

func (s *ListUsersRequest) SetPhoneNo(v string) *ListUsersRequest {
	s.PhoneNo = &v
	return s
}

func (s *ListUsersRequest) SetAttachment(v string) *ListUsersRequest {
	s.Attachment = &v
	return s
}

func (s *ListUsersRequest) SetBizId(v string) *ListUsersRequest {
	s.BizId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int64) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetPersonList(v map[string]interface{}) *ListUsersRequest {
	s.PersonList = v
	return s
}

func (s *ListUsersRequest) SetUserList(v map[string]interface{}) *ListUsersRequest {
	s.UserList = v
	return s
}

func (s *ListUsersRequest) SetMatchingRateThreshold(v string) *ListUsersRequest {
	s.MatchingRateThreshold = &v
	return s
}

type ListUsersShrinkRequest struct {
	CorpId                *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId              *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId           *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdNumber              *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl          *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Age                   *int32  `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender                *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment            *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PersonListShrink      *string `json:"PersonList,omitempty" xml:"PersonList,omitempty"`
	UserListShrink        *string `json:"UserList,omitempty" xml:"UserList,omitempty"`
	MatchingRateThreshold *string `json:"MatchingRateThreshold,omitempty" xml:"MatchingRateThreshold,omitempty"`
}

func (s ListUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUsersShrinkRequest) SetCorpId(v string) *ListUsersShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *ListUsersShrinkRequest) SetIsvSubId(v string) *ListUsersShrinkRequest {
	s.IsvSubId = &v
	return s
}

func (s *ListUsersShrinkRequest) SetUserName(v string) *ListUsersShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ListUsersShrinkRequest) SetUserGroupId(v int64) *ListUsersShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListUsersShrinkRequest) SetIdNumber(v string) *ListUsersShrinkRequest {
	s.IdNumber = &v
	return s
}

func (s *ListUsersShrinkRequest) SetFaceImageUrl(v string) *ListUsersShrinkRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *ListUsersShrinkRequest) SetAddress(v string) *ListUsersShrinkRequest {
	s.Address = &v
	return s
}

func (s *ListUsersShrinkRequest) SetAge(v int32) *ListUsersShrinkRequest {
	s.Age = &v
	return s
}

func (s *ListUsersShrinkRequest) SetGender(v int32) *ListUsersShrinkRequest {
	s.Gender = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPlateNo(v string) *ListUsersShrinkRequest {
	s.PlateNo = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPhoneNo(v string) *ListUsersShrinkRequest {
	s.PhoneNo = &v
	return s
}

func (s *ListUsersShrinkRequest) SetAttachment(v string) *ListUsersShrinkRequest {
	s.Attachment = &v
	return s
}

func (s *ListUsersShrinkRequest) SetBizId(v string) *ListUsersShrinkRequest {
	s.BizId = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPageNumber(v int64) *ListUsersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPageSize(v int64) *ListUsersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersShrinkRequest) SetPersonListShrink(v string) *ListUsersShrinkRequest {
	s.PersonListShrink = &v
	return s
}

func (s *ListUsersShrinkRequest) SetUserListShrink(v string) *ListUsersShrinkRequest {
	s.UserListShrink = &v
	return s
}

func (s *ListUsersShrinkRequest) SetMatchingRateThreshold(v string) *ListUsersShrinkRequest {
	s.MatchingRateThreshold = &v
	return s
}

type ListUsersResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

type ListUsersResponseBodyData struct {
	Success    *int64                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Records    []*ListUsersResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetSuccess(v int64) *ListUsersResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBodyData) SetRecords(v []*ListUsersResponseBodyDataRecords) *ListUsersResponseBodyData {
	s.Records = v
	return s
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int64) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int64) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotal(v int64) *ListUsersResponseBodyData {
	s.Total = &v
	return s
}

type ListUsersResponseBodyDataRecords struct {
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	UserGroupId  *int32  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId       *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	MatchingRate *string `json:"MatchingRate,omitempty" xml:"MatchingRate,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataRecords) SetIsvSubId(v string) *ListUsersResponseBodyDataRecords {
	s.IsvSubId = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetGender(v string) *ListUsersResponseBodyDataRecords {
	s.Gender = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetFaceImageUrl(v string) *ListUsersResponseBodyDataRecords {
	s.FaceImageUrl = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetUserGroupId(v int32) *ListUsersResponseBodyDataRecords {
	s.UserGroupId = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetUserId(v int32) *ListUsersResponseBodyDataRecords {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetMatchingRate(v string) *ListUsersResponseBodyDataRecords {
	s.MatchingRate = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetBizId(v string) *ListUsersResponseBodyDataRecords {
	s.BizId = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetAttachment(v string) *ListUsersResponseBodyDataRecords {
	s.Attachment = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetAge(v string) *ListUsersResponseBodyDataRecords {
	s.Age = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetIdNumber(v string) *ListUsersResponseBodyDataRecords {
	s.IdNumber = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetPersonId(v string) *ListUsersResponseBodyDataRecords {
	s.PersonId = &v
	return s
}

func (s *ListUsersResponseBodyDataRecords) SetUserName(v string) *ListUsersResponseBodyDataRecords {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type RecognizeFaceQualityRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PicContent *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicFormat  *string `json:"PicFormat,omitempty" xml:"PicFormat,omitempty"`
	PicUrl     *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s RecognizeFaceQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityRequest) SetCorpId(v string) *RecognizeFaceQualityRequest {
	s.CorpId = &v
	return s
}

func (s *RecognizeFaceQualityRequest) SetPicContent(v string) *RecognizeFaceQualityRequest {
	s.PicContent = &v
	return s
}

func (s *RecognizeFaceQualityRequest) SetPicFormat(v string) *RecognizeFaceQualityRequest {
	s.PicFormat = &v
	return s
}

func (s *RecognizeFaceQualityRequest) SetPicUrl(v string) *RecognizeFaceQualityRequest {
	s.PicUrl = &v
	return s
}

type RecognizeFaceQualityResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeFaceQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RecognizeFaceQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponseBody) SetMessage(v string) *RecognizeFaceQualityResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeFaceQualityResponseBody) SetRequestId(v string) *RecognizeFaceQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeFaceQualityResponseBody) SetData(v *RecognizeFaceQualityResponseBodyData) *RecognizeFaceQualityResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeFaceQualityResponseBody) SetCode(v string) *RecognizeFaceQualityResponseBody {
	s.Code = &v
	return s
}

type RecognizeFaceQualityResponseBodyData struct {
	QualityScore *string                                         `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	Description  *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Attributes   *RecognizeFaceQualityResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Struct"`
}

func (s RecognizeFaceQualityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponseBodyData) SetQualityScore(v string) *RecognizeFaceQualityResponseBodyData {
	s.QualityScore = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyData) SetDescription(v string) *RecognizeFaceQualityResponseBodyData {
	s.Description = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyData) SetAttributes(v *RecognizeFaceQualityResponseBodyDataAttributes) *RecognizeFaceQualityResponseBodyData {
	s.Attributes = v
	return s
}

type RecognizeFaceQualityResponseBodyDataAttributes struct {
	FaceScore              *string `json:"FaceScore,omitempty" xml:"FaceScore,omitempty"`
	RightBottomY           *int32  `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY               *int32  `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	FaceStyle              *string `json:"FaceStyle,omitempty" xml:"FaceStyle,omitempty"`
	FaceQuality            *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	LeftTopX               *int32  `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	RightBottomX           *int32  `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	TargetImageStoragePath *string `json:"TargetImageStoragePath,omitempty" xml:"TargetImageStoragePath,omitempty"`
}

func (s RecognizeFaceQualityResponseBodyDataAttributes) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetFaceScore(v string) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.FaceScore = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetRightBottomY(v int32) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.RightBottomY = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetLeftTopY(v int32) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetFaceStyle(v string) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.FaceStyle = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetFaceQuality(v string) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.FaceQuality = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetLeftTopX(v int32) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.LeftTopX = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetRightBottomX(v int32) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeFaceQualityResponseBodyDataAttributes) SetTargetImageStoragePath(v string) *RecognizeFaceQualityResponseBodyDataAttributes {
	s.TargetImageStoragePath = &v
	return s
}

type RecognizeFaceQualityResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeFaceQualityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFaceQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponse) SetHeaders(v map[string]*string) *RecognizeFaceQualityResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFaceQualityResponse) SetBody(v *RecognizeFaceQualityResponseBody) *RecognizeFaceQualityResponse {
	s.Body = v
	return s
}

type RecognizeImageRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PicContent *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicFormat  *string `json:"PicFormat,omitempty" xml:"PicFormat,omitempty"`
	PicUrl     *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s RecognizeImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageRequest) SetCorpId(v string) *RecognizeImageRequest {
	s.CorpId = &v
	return s
}

func (s *RecognizeImageRequest) SetPicContent(v string) *RecognizeImageRequest {
	s.PicContent = &v
	return s
}

func (s *RecognizeImageRequest) SetPicFormat(v string) *RecognizeImageRequest {
	s.PicFormat = &v
	return s
}

func (s *RecognizeImageRequest) SetPicUrl(v string) *RecognizeImageRequest {
	s.PicUrl = &v
	return s
}

type RecognizeImageResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RecognizeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseBody) SetMessage(v string) *RecognizeImageResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeImageResponseBody) SetRequestId(v string) *RecognizeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeImageResponseBody) SetData(v *RecognizeImageResponseBodyData) *RecognizeImageResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeImageResponseBody) SetCode(v string) *RecognizeImageResponseBody {
	s.Code = &v
	return s
}

type RecognizeImageResponseBodyData struct {
	BodyList []*RecognizeImageResponseBodyDataBodyList `json:"BodyList,omitempty" xml:"BodyList,omitempty" type:"Repeated"`
	FaceList []*RecognizeImageResponseBodyDataFaceList `json:"FaceList,omitempty" xml:"FaceList,omitempty" type:"Repeated"`
}

func (s RecognizeImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseBodyData) SetBodyList(v []*RecognizeImageResponseBodyDataBodyList) *RecognizeImageResponseBodyData {
	s.BodyList = v
	return s
}

func (s *RecognizeImageResponseBodyData) SetFaceList(v []*RecognizeImageResponseBodyDataFaceList) *RecognizeImageResponseBodyData {
	s.FaceList = v
	return s
}

type RecognizeImageResponseBodyDataBodyList struct {
	RespiratorColor  *string `json:"RespiratorColor,omitempty" xml:"RespiratorColor,omitempty"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Feature          *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ImageBaseSixFour *string `json:"ImageBaseSixFour,omitempty" xml:"ImageBaseSixFour,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LocalFeature     *string `json:"LocalFeature,omitempty" xml:"LocalFeature,omitempty"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s RecognizeImageResponseBodyDataBodyList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseBodyDataBodyList) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseBodyDataBodyList) SetRespiratorColor(v string) *RecognizeImageResponseBodyDataBodyList {
	s.RespiratorColor = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetRightBottomY(v string) *RecognizeImageResponseBodyDataBodyList {
	s.RightBottomY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetFeature(v string) *RecognizeImageResponseBodyDataBodyList {
	s.Feature = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetLeftTopY(v string) *RecognizeImageResponseBodyDataBodyList {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetImageBaseSixFour(v string) *RecognizeImageResponseBodyDataBodyList {
	s.ImageBaseSixFour = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetFileName(v string) *RecognizeImageResponseBodyDataBodyList {
	s.FileName = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetRightBottomX(v string) *RecognizeImageResponseBodyDataBodyList {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetLocalFeature(v string) *RecognizeImageResponseBodyDataBodyList {
	s.LocalFeature = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetLeftTopX(v string) *RecognizeImageResponseBodyDataBodyList {
	s.LeftTopX = &v
	return s
}

type RecognizeImageResponseBodyDataFaceList struct {
	Quality          *float32 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	RespiratorColor  *string  `json:"RespiratorColor,omitempty" xml:"RespiratorColor,omitempty"`
	KeyPointQuality  *float32 `json:"KeyPointQuality,omitempty" xml:"KeyPointQuality,omitempty"`
	RightBottomY     *string  `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Feature          *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	LeftTopY         *string  `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ImageBaseSixFour *string  `json:"ImageBaseSixFour,omitempty" xml:"ImageBaseSixFour,omitempty"`
	FileName         *string  `json:"FileName,omitempty" xml:"FileName,omitempty"`
	RightBottomX     *string  `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LocalFeature     *string  `json:"LocalFeature,omitempty" xml:"LocalFeature,omitempty"`
	LeftTopX         *string  `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s RecognizeImageResponseBodyDataFaceList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseBodyDataFaceList) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseBodyDataFaceList) SetQuality(v float32) *RecognizeImageResponseBodyDataFaceList {
	s.Quality = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetRespiratorColor(v string) *RecognizeImageResponseBodyDataFaceList {
	s.RespiratorColor = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetKeyPointQuality(v float32) *RecognizeImageResponseBodyDataFaceList {
	s.KeyPointQuality = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetRightBottomY(v string) *RecognizeImageResponseBodyDataFaceList {
	s.RightBottomY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetFeature(v string) *RecognizeImageResponseBodyDataFaceList {
	s.Feature = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetLeftTopY(v string) *RecognizeImageResponseBodyDataFaceList {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetImageBaseSixFour(v string) *RecognizeImageResponseBodyDataFaceList {
	s.ImageBaseSixFour = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetFileName(v string) *RecognizeImageResponseBodyDataFaceList {
	s.FileName = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetRightBottomX(v string) *RecognizeImageResponseBodyDataFaceList {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetLocalFeature(v string) *RecognizeImageResponseBodyDataFaceList {
	s.LocalFeature = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetLeftTopX(v string) *RecognizeImageResponseBodyDataFaceList {
	s.LeftTopX = &v
	return s
}

type RecognizeImageResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponse) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponse) SetHeaders(v map[string]*string) *RecognizeImageResponse {
	s.Headers = v
	return s
}

func (s *RecognizeImageResponse) SetBody(v *RecognizeImageResponseBody) *RecognizeImageResponse {
	s.Body = v
	return s
}

type RegisterDeviceRequest struct {
	DeviceSn        *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ServerId        *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	DeviceTimeStamp *string `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetDeviceSn(v string) *RegisterDeviceRequest {
	s.DeviceSn = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceId(v string) *RegisterDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterDeviceRequest) SetServerId(v string) *RegisterDeviceRequest {
	s.ServerId = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceTimeStamp(v string) *RegisterDeviceRequest {
	s.DeviceTimeStamp = &v
	return s
}

type RegisterDeviceResponseBody struct {
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetryInterval *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRetryInterval(v string) *RegisterDeviceResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetCode(v string) *RegisterDeviceResponseBody {
	s.Code = &v
	return s
}

type RegisterDeviceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type ReportDeviceCapacityRequest struct {
	Longitude        *string                                        `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude         *string                                        `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	AudioFormat      *string                                        `json:"AudioFormat,omitempty" xml:"AudioFormat,omitempty"`
	PresetNum        *string                                        `json:"PresetNum,omitempty" xml:"PresetNum,omitempty"`
	PTZCapacity      *string                                        `json:"PTZCapacity,omitempty" xml:"PTZCapacity,omitempty"`
	DeviceSn         *string                                        `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	DeviceTimeStamp  *string                                        `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty"`
	StreamCapacities []*ReportDeviceCapacityRequestStreamCapacities `json:"StreamCapacities,omitempty" xml:"StreamCapacities,omitempty" type:"Repeated"`
}

func (s ReportDeviceCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceCapacityRequest) GoString() string {
	return s.String()
}

func (s *ReportDeviceCapacityRequest) SetLongitude(v string) *ReportDeviceCapacityRequest {
	s.Longitude = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetLatitude(v string) *ReportDeviceCapacityRequest {
	s.Latitude = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetAudioFormat(v string) *ReportDeviceCapacityRequest {
	s.AudioFormat = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetPresetNum(v string) *ReportDeviceCapacityRequest {
	s.PresetNum = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetPTZCapacity(v string) *ReportDeviceCapacityRequest {
	s.PTZCapacity = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetDeviceSn(v string) *ReportDeviceCapacityRequest {
	s.DeviceSn = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetDeviceTimeStamp(v string) *ReportDeviceCapacityRequest {
	s.DeviceTimeStamp = &v
	return s
}

func (s *ReportDeviceCapacityRequest) SetStreamCapacities(v []*ReportDeviceCapacityRequestStreamCapacities) *ReportDeviceCapacityRequest {
	s.StreamCapacities = v
	return s
}

type ReportDeviceCapacityRequestStreamCapacities struct {
	EncodeFormat   *string `json:"EncodeFormat,omitempty" xml:"EncodeFormat,omitempty"`
	GovLengthRange *string `json:"GovLengthRange,omitempty" xml:"GovLengthRange,omitempty"`
	MaxFrameRate   *string `json:"MaxFrameRate,omitempty" xml:"MaxFrameRate,omitempty"`
	BitrateRange   *string `json:"BitrateRange,omitempty" xml:"BitrateRange,omitempty"`
	MaxStream      *string `json:"MaxStream,omitempty" xml:"MaxStream,omitempty"`
	Resolution     *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s ReportDeviceCapacityRequestStreamCapacities) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceCapacityRequestStreamCapacities) GoString() string {
	return s.String()
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetEncodeFormat(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.EncodeFormat = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetGovLengthRange(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.GovLengthRange = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetMaxFrameRate(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.MaxFrameRate = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetBitrateRange(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.BitrateRange = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetMaxStream(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.MaxStream = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetResolution(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.Resolution = &v
	return s
}

type ReportDeviceCapacityResponseBody struct {
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetryInterval *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ReportDeviceCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *ReportDeviceCapacityResponseBody) SetMessage(v string) *ReportDeviceCapacityResponseBody {
	s.Message = &v
	return s
}

func (s *ReportDeviceCapacityResponseBody) SetRequestId(v string) *ReportDeviceCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportDeviceCapacityResponseBody) SetRetryInterval(v string) *ReportDeviceCapacityResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *ReportDeviceCapacityResponseBody) SetCode(v string) *ReportDeviceCapacityResponseBody {
	s.Code = &v
	return s
}

type ReportDeviceCapacityResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportDeviceCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportDeviceCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceCapacityResponse) GoString() string {
	return s.String()
}

func (s *ReportDeviceCapacityResponse) SetHeaders(v map[string]*string) *ReportDeviceCapacityResponse {
	s.Headers = v
	return s
}

func (s *ReportDeviceCapacityResponse) SetBody(v *ReportDeviceCapacityResponseBody) *ReportDeviceCapacityResponse {
	s.Body = v
	return s
}

type SaveVideoSummaryTaskVideoRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SaveVideo *bool   `json:"SaveVideo,omitempty" xml:"SaveVideo,omitempty"`
}

func (s SaveVideoSummaryTaskVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoSummaryTaskVideoRequest) GoString() string {
	return s.String()
}

func (s *SaveVideoSummaryTaskVideoRequest) SetCorpId(v string) *SaveVideoSummaryTaskVideoRequest {
	s.CorpId = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoRequest) SetTaskId(v int64) *SaveVideoSummaryTaskVideoRequest {
	s.TaskId = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoRequest) SetSaveVideo(v bool) *SaveVideoSummaryTaskVideoRequest {
	s.SaveVideo = &v
	return s
}

type SaveVideoSummaryTaskVideoResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SaveVideoSummaryTaskVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoSummaryTaskVideoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveVideoSummaryTaskVideoResponseBody) SetMessage(v string) *SaveVideoSummaryTaskVideoResponseBody {
	s.Message = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponseBody) SetRequestId(v string) *SaveVideoSummaryTaskVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponseBody) SetData(v string) *SaveVideoSummaryTaskVideoResponseBody {
	s.Data = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponseBody) SetCode(v string) *SaveVideoSummaryTaskVideoResponseBody {
	s.Code = &v
	return s
}

type SaveVideoSummaryTaskVideoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveVideoSummaryTaskVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveVideoSummaryTaskVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoSummaryTaskVideoResponse) GoString() string {
	return s.String()
}

func (s *SaveVideoSummaryTaskVideoResponse) SetHeaders(v map[string]*string) *SaveVideoSummaryTaskVideoResponse {
	s.Headers = v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponse) SetBody(v *SaveVideoSummaryTaskVideoResponseBody) *SaveVideoSummaryTaskVideoResponse {
	s.Body = v
	return s
}

type SearchBodyRequest struct {
	CorpId         *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId           *string                `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp *int64                 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	EndTimeStamp   *int64                 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	PageNo         *int32                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OptionList     map[string]interface{} `json:"OptionList,omitempty" xml:"OptionList,omitempty"`
}

func (s SearchBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyRequest) GoString() string {
	return s.String()
}

func (s *SearchBodyRequest) SetCorpId(v string) *SearchBodyRequest {
	s.CorpId = &v
	return s
}

func (s *SearchBodyRequest) SetGbId(v string) *SearchBodyRequest {
	s.GbId = &v
	return s
}

func (s *SearchBodyRequest) SetStartTimeStamp(v int64) *SearchBodyRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *SearchBodyRequest) SetEndTimeStamp(v int64) *SearchBodyRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *SearchBodyRequest) SetPageNo(v int32) *SearchBodyRequest {
	s.PageNo = &v
	return s
}

func (s *SearchBodyRequest) SetPageSize(v int32) *SearchBodyRequest {
	s.PageSize = &v
	return s
}

func (s *SearchBodyRequest) SetOptionList(v map[string]interface{}) *SearchBodyRequest {
	s.OptionList = v
	return s
}

type SearchBodyShrinkRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp   *int64  `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	EndTimeStamp     *int64  `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	PageNo           *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OptionListShrink *string `json:"OptionList,omitempty" xml:"OptionList,omitempty"`
}

func (s SearchBodyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchBodyShrinkRequest) SetCorpId(v string) *SearchBodyShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetGbId(v string) *SearchBodyShrinkRequest {
	s.GbId = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetStartTimeStamp(v int64) *SearchBodyShrinkRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetEndTimeStamp(v int64) *SearchBodyShrinkRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetPageNo(v int32) *SearchBodyShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetPageSize(v int32) *SearchBodyShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetOptionListShrink(v string) *SearchBodyShrinkRequest {
	s.OptionListShrink = &v
	return s
}

type SearchBodyResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SearchBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SearchBodyResponseBody) SetMessage(v string) *SearchBodyResponseBody {
	s.Message = &v
	return s
}

func (s *SearchBodyResponseBody) SetRequestId(v string) *SearchBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchBodyResponseBody) SetData(v *SearchBodyResponseBodyData) *SearchBodyResponseBody {
	s.Data = v
	return s
}

func (s *SearchBodyResponseBody) SetCode(v string) *SearchBodyResponseBody {
	s.Code = &v
	return s
}

type SearchBodyResponseBodyData struct {
	Records    []*SearchBodyResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNo     *int32                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TotalPage  *int32                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchBodyResponseBodyData) SetRecords(v []*SearchBodyResponseBodyDataRecords) *SearchBodyResponseBodyData {
	s.Records = v
	return s
}

func (s *SearchBodyResponseBodyData) SetPageNo(v int32) *SearchBodyResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *SearchBodyResponseBodyData) SetTotalPage(v int32) *SearchBodyResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *SearchBodyResponseBodyData) SetPageSize(v int32) *SearchBodyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *SearchBodyResponseBodyData) SetTotalCount(v int32) *SearchBodyResponseBodyData {
	s.TotalCount = &v
	return s
}

type SearchBodyResponseBodyDataRecords struct {
	GbId           *string  `json:"GbId,omitempty" xml:"GbId,omitempty"`
	TargetImageUrl *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	RightBottomY   *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	ImageUrl       *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	LeftTopY       *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	LeftTopX       *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	RightBottomX   *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
}

func (s SearchBodyResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *SearchBodyResponseBodyDataRecords) SetGbId(v string) *SearchBodyResponseBodyDataRecords {
	s.GbId = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetTargetImageUrl(v string) *SearchBodyResponseBodyDataRecords {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetRightBottomY(v float32) *SearchBodyResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetImageUrl(v string) *SearchBodyResponseBodyDataRecords {
	s.ImageUrl = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetLeftTopY(v float32) *SearchBodyResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetScore(v float32) *SearchBodyResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetLeftTopX(v float32) *SearchBodyResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *SearchBodyResponseBodyDataRecords) SetRightBottomX(v float32) *SearchBodyResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

type SearchBodyResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponse) GoString() string {
	return s.String()
}

func (s *SearchBodyResponse) SetHeaders(v map[string]*string) *SearchBodyResponse {
	s.Headers = v
	return s
}

func (s *SearchBodyResponse) SetBody(v *SearchBodyResponseBody) *SearchBodyResponse {
	s.Body = v
	return s
}

type SearchFaceRequest struct {
	CorpId         *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId           *string                `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp *int64                 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	EndTimeStamp   *int64                 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	PageNo         *int32                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OptionList     map[string]interface{} `json:"OptionList,omitempty" xml:"OptionList,omitempty"`
}

func (s SearchFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceRequest) SetCorpId(v string) *SearchFaceRequest {
	s.CorpId = &v
	return s
}

func (s *SearchFaceRequest) SetGbId(v string) *SearchFaceRequest {
	s.GbId = &v
	return s
}

func (s *SearchFaceRequest) SetStartTimeStamp(v int64) *SearchFaceRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *SearchFaceRequest) SetEndTimeStamp(v int64) *SearchFaceRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *SearchFaceRequest) SetPageNo(v int32) *SearchFaceRequest {
	s.PageNo = &v
	return s
}

func (s *SearchFaceRequest) SetPageSize(v int32) *SearchFaceRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaceRequest) SetOptionList(v map[string]interface{}) *SearchFaceRequest {
	s.OptionList = v
	return s
}

type SearchFaceShrinkRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp   *int64  `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	EndTimeStamp     *int64  `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	PageNo           *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OptionListShrink *string `json:"OptionList,omitempty" xml:"OptionList,omitempty"`
}

func (s SearchFaceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceShrinkRequest) SetCorpId(v string) *SearchFaceShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetGbId(v string) *SearchFaceShrinkRequest {
	s.GbId = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetStartTimeStamp(v int64) *SearchFaceShrinkRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetEndTimeStamp(v int64) *SearchFaceShrinkRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetPageNo(v int32) *SearchFaceShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetPageSize(v int32) *SearchFaceShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetOptionListShrink(v string) *SearchFaceShrinkRequest {
	s.OptionListShrink = &v
	return s
}

type SearchFaceResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SearchFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBody) SetMessage(v string) *SearchFaceResponseBody {
	s.Message = &v
	return s
}

func (s *SearchFaceResponseBody) SetRequestId(v string) *SearchFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFaceResponseBody) SetData(v *SearchFaceResponseBodyData) *SearchFaceResponseBody {
	s.Data = v
	return s
}

func (s *SearchFaceResponseBody) SetCode(v string) *SearchFaceResponseBody {
	s.Code = &v
	return s
}

type SearchFaceResponseBodyData struct {
	Records    []*SearchFaceResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNo     *int32                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TotalPage  *int32                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyData) SetRecords(v []*SearchFaceResponseBodyDataRecords) *SearchFaceResponseBodyData {
	s.Records = v
	return s
}

func (s *SearchFaceResponseBodyData) SetPageNo(v int32) *SearchFaceResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *SearchFaceResponseBodyData) SetTotalPage(v int32) *SearchFaceResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *SearchFaceResponseBodyData) SetPageSize(v int32) *SearchFaceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *SearchFaceResponseBodyData) SetTotalCount(v int32) *SearchFaceResponseBodyData {
	s.TotalCount = &v
	return s
}

type SearchFaceResponseBodyDataRecords struct {
	GbId            *string  `json:"GbId,omitempty" xml:"GbId,omitempty"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	RightBottomY    *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	ImageUrl        *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	LeftTopY        *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	SourceId        *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	RightBottomX    *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LeftTopX        *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	MatchSuggestion *string  `json:"MatchSuggestion,omitempty" xml:"MatchSuggestion,omitempty"`
}

func (s SearchFaceResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataRecords) SetGbId(v string) *SearchFaceResponseBodyDataRecords {
	s.GbId = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetTargetImageUrl(v string) *SearchFaceResponseBodyDataRecords {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetRightBottomY(v float32) *SearchFaceResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetImageUrl(v string) *SearchFaceResponseBodyDataRecords {
	s.ImageUrl = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetLeftTopY(v float32) *SearchFaceResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetScore(v float32) *SearchFaceResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetSourceId(v string) *SearchFaceResponseBodyDataRecords {
	s.SourceId = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetRightBottomX(v float32) *SearchFaceResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetLeftTopX(v float32) *SearchFaceResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *SearchFaceResponseBodyDataRecords) SetMatchSuggestion(v string) *SearchFaceResponseBodyDataRecords {
	s.MatchSuggestion = &v
	return s
}

type SearchFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchFaceResponse) SetHeaders(v map[string]*string) *SearchFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchFaceResponse) SetBody(v *SearchFaceResponseBody) *SearchFaceResponse {
	s.Body = v
	return s
}

type SearchObjectRequest struct {
	CorpId        *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ObjectType    *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	StartTime     *int64                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DeviceList    map[string]interface{} `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	PicUrl        *string                `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Conditions    map[string]interface{} `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	AlgorithmType *string                `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	ImagePath     map[string]interface{} `json:"ImagePath,omitempty" xml:"ImagePath,omitempty"`
}

func (s SearchObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectRequest) GoString() string {
	return s.String()
}

func (s *SearchObjectRequest) SetCorpId(v string) *SearchObjectRequest {
	s.CorpId = &v
	return s
}

func (s *SearchObjectRequest) SetObjectType(v string) *SearchObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *SearchObjectRequest) SetStartTime(v int64) *SearchObjectRequest {
	s.StartTime = &v
	return s
}

func (s *SearchObjectRequest) SetEndTime(v int64) *SearchObjectRequest {
	s.EndTime = &v
	return s
}

func (s *SearchObjectRequest) SetPageNumber(v int32) *SearchObjectRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchObjectRequest) SetPageSize(v int32) *SearchObjectRequest {
	s.PageSize = &v
	return s
}

func (s *SearchObjectRequest) SetDeviceList(v map[string]interface{}) *SearchObjectRequest {
	s.DeviceList = v
	return s
}

func (s *SearchObjectRequest) SetPicUrl(v string) *SearchObjectRequest {
	s.PicUrl = &v
	return s
}

func (s *SearchObjectRequest) SetConditions(v map[string]interface{}) *SearchObjectRequest {
	s.Conditions = v
	return s
}

func (s *SearchObjectRequest) SetAlgorithmType(v string) *SearchObjectRequest {
	s.AlgorithmType = &v
	return s
}

func (s *SearchObjectRequest) SetImagePath(v map[string]interface{}) *SearchObjectRequest {
	s.ImagePath = v
	return s
}

type SearchObjectShrinkRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ObjectType       *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DeviceListShrink *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	PicUrl           *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ConditionsShrink *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	AlgorithmType    *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	ImagePathShrink  *string `json:"ImagePath,omitempty" xml:"ImagePath,omitempty"`
}

func (s SearchObjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchObjectShrinkRequest) SetCorpId(v string) *SearchObjectShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetObjectType(v string) *SearchObjectShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetStartTime(v int64) *SearchObjectShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetEndTime(v int64) *SearchObjectShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetPageNumber(v int32) *SearchObjectShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetPageSize(v int32) *SearchObjectShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetDeviceListShrink(v string) *SearchObjectShrinkRequest {
	s.DeviceListShrink = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetPicUrl(v string) *SearchObjectShrinkRequest {
	s.PicUrl = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetConditionsShrink(v string) *SearchObjectShrinkRequest {
	s.ConditionsShrink = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetAlgorithmType(v string) *SearchObjectShrinkRequest {
	s.AlgorithmType = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetImagePathShrink(v string) *SearchObjectShrinkRequest {
	s.ImagePathShrink = &v
	return s
}

type SearchObjectResponseBody struct {
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SearchObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBody) SetMessage(v string) *SearchObjectResponseBody {
	s.Message = &v
	return s
}

func (s *SearchObjectResponseBody) SetRequestId(v string) *SearchObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchObjectResponseBody) SetData(v *SearchObjectResponseBodyData) *SearchObjectResponseBody {
	s.Data = v
	return s
}

func (s *SearchObjectResponseBody) SetCode(v string) *SearchObjectResponseBody {
	s.Code = &v
	return s
}

type SearchObjectResponseBodyData struct {
	Records    []*SearchObjectResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyData) SetRecords(v []*SearchObjectResponseBodyDataRecords) *SearchObjectResponseBodyData {
	s.Records = v
	return s
}

func (s *SearchObjectResponseBodyData) SetTotalPage(v int32) *SearchObjectResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *SearchObjectResponseBodyData) SetPageNumber(v int32) *SearchObjectResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *SearchObjectResponseBodyData) SetPageSize(v int32) *SearchObjectResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *SearchObjectResponseBodyData) SetTotalCount(v int32) *SearchObjectResponseBodyData {
	s.TotalCount = &v
	return s
}

type SearchObjectResponseBodyDataRecords struct {
	DeviceID        *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	CompareResult   *string  `json:"CompareResult,omitempty" xml:"CompareResult,omitempty"`
	RightBtmX       *int32   `json:"RightBtmX,omitempty" xml:"RightBtmX,omitempty"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	SourceImageUrl  *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	SourceID        *string  `json:"SourceID,omitempty" xml:"SourceID,omitempty"`
	RightBtmY       *int32   `json:"RightBtmY,omitempty" xml:"RightBtmY,omitempty"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	LeftTopY        *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	TargetImagePath *string  `json:"TargetImagePath,omitempty" xml:"TargetImagePath,omitempty"`
	ShotTime        *int64   `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	LeftTopX        *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	SourceImagePath *string  `json:"SourceImagePath,omitempty" xml:"SourceImagePath,omitempty"`
}

func (s SearchObjectResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyDataRecords) SetDeviceID(v string) *SearchObjectResponseBodyDataRecords {
	s.DeviceID = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetCompareResult(v string) *SearchObjectResponseBodyDataRecords {
	s.CompareResult = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetRightBtmX(v int32) *SearchObjectResponseBodyDataRecords {
	s.RightBtmX = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetScore(v float32) *SearchObjectResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetSourceImageUrl(v string) *SearchObjectResponseBodyDataRecords {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetSourceID(v string) *SearchObjectResponseBodyDataRecords {
	s.SourceID = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetRightBtmY(v int32) *SearchObjectResponseBodyDataRecords {
	s.RightBtmY = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetTargetImageUrl(v string) *SearchObjectResponseBodyDataRecords {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetLeftTopY(v int32) *SearchObjectResponseBodyDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetTargetImagePath(v string) *SearchObjectResponseBodyDataRecords {
	s.TargetImagePath = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetShotTime(v int64) *SearchObjectResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetLeftTopX(v int32) *SearchObjectResponseBodyDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *SearchObjectResponseBodyDataRecords) SetSourceImagePath(v string) *SearchObjectResponseBodyDataRecords {
	s.SourceImagePath = &v
	return s
}

type SearchObjectResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponse) GoString() string {
	return s.String()
}

func (s *SearchObjectResponse) SetHeaders(v map[string]*string) *SearchObjectResponse {
	s.Headers = v
	return s
}

func (s *SearchObjectResponse) SetBody(v *SearchObjectResponseBody) *SearchObjectResponse {
	s.Body = v
	return s
}

type StopMonitorRequest struct {
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
}

func (s StopMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMonitorRequest) GoString() string {
	return s.String()
}

func (s *StopMonitorRequest) SetTaskId(v string) *StopMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *StopMonitorRequest) SetAlgorithmVendor(v string) *StopMonitorRequest {
	s.AlgorithmVendor = &v
	return s
}

type StopMonitorResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StopMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *StopMonitorResponseBody) SetMessage(v string) *StopMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *StopMonitorResponseBody) SetRequestId(v string) *StopMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMonitorResponseBody) SetData(v string) *StopMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *StopMonitorResponseBody) SetCode(v string) *StopMonitorResponseBody {
	s.Code = &v
	return s
}

type StopMonitorResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMonitorResponse) GoString() string {
	return s.String()
}

func (s *StopMonitorResponse) SetHeaders(v map[string]*string) *StopMonitorResponse {
	s.Headers = v
	return s
}

func (s *StopMonitorResponse) SetBody(v *StopMonitorResponseBody) *StopMonitorResponse {
	s.Body = v
	return s
}

type SubscribeDeviceEventRequest struct {
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	PushConfig *string `json:"PushConfig,omitempty" xml:"PushConfig,omitempty"`
}

func (s SubscribeDeviceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeDeviceEventRequest) GoString() string {
	return s.String()
}

func (s *SubscribeDeviceEventRequest) SetDeviceId(v string) *SubscribeDeviceEventRequest {
	s.DeviceId = &v
	return s
}

func (s *SubscribeDeviceEventRequest) SetPushConfig(v string) *SubscribeDeviceEventRequest {
	s.PushConfig = &v
	return s
}

type SubscribeDeviceEventResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SubscribeDeviceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeDeviceEventResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeDeviceEventResponseBody) SetMessage(v string) *SubscribeDeviceEventResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeDeviceEventResponseBody) SetRequestId(v string) *SubscribeDeviceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeDeviceEventResponseBody) SetCode(v string) *SubscribeDeviceEventResponseBody {
	s.Code = &v
	return s
}

type SubscribeDeviceEventResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubscribeDeviceEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscribeDeviceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeDeviceEventResponse) GoString() string {
	return s.String()
}

func (s *SubscribeDeviceEventResponse) SetHeaders(v map[string]*string) *SubscribeDeviceEventResponse {
	s.Headers = v
	return s
}

func (s *SubscribeDeviceEventResponse) SetBody(v *SubscribeDeviceEventResponseBody) *SubscribeDeviceEventResponse {
	s.Body = v
	return s
}

type SubscribeSpaceEventRequest struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	PushConfig *string `json:"PushConfig,omitempty" xml:"PushConfig,omitempty"`
}

func (s SubscribeSpaceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeSpaceEventRequest) GoString() string {
	return s.String()
}

func (s *SubscribeSpaceEventRequest) SetSpaceId(v string) *SubscribeSpaceEventRequest {
	s.SpaceId = &v
	return s
}

func (s *SubscribeSpaceEventRequest) SetPushConfig(v string) *SubscribeSpaceEventRequest {
	s.PushConfig = &v
	return s
}

type SubscribeSpaceEventResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SubscribeSpaceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeSpaceEventResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeSpaceEventResponseBody) SetMessage(v string) *SubscribeSpaceEventResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeSpaceEventResponseBody) SetRequestId(v string) *SubscribeSpaceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeSpaceEventResponseBody) SetCode(v string) *SubscribeSpaceEventResponseBody {
	s.Code = &v
	return s
}

type SubscribeSpaceEventResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubscribeSpaceEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscribeSpaceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeSpaceEventResponse) GoString() string {
	return s.String()
}

func (s *SubscribeSpaceEventResponse) SetHeaders(v map[string]*string) *SubscribeSpaceEventResponse {
	s.Headers = v
	return s
}

func (s *SubscribeSpaceEventResponse) SetBody(v *SubscribeSpaceEventResponseBody) *SubscribeSpaceEventResponse {
	s.Body = v
	return s
}

type SyncDeviceTimeRequest struct {
	DeviceSn        *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	DeviceTimeStamp *string `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty"`
}

func (s SyncDeviceTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceTimeRequest) GoString() string {
	return s.String()
}

func (s *SyncDeviceTimeRequest) SetDeviceSn(v string) *SyncDeviceTimeRequest {
	s.DeviceSn = &v
	return s
}

func (s *SyncDeviceTimeRequest) SetDeviceTimeStamp(v string) *SyncDeviceTimeRequest {
	s.DeviceTimeStamp = &v
	return s
}

type SyncDeviceTimeResponseBody struct {
	SyncInterval  *string `json:"SyncInterval,omitempty" xml:"SyncInterval,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RetryInterval *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	NTPServer     *string `json:"NTPServer,omitempty" xml:"NTPServer,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SyncDeviceTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceTimeResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDeviceTimeResponseBody) SetSyncInterval(v string) *SyncDeviceTimeResponseBody {
	s.SyncInterval = &v
	return s
}

func (s *SyncDeviceTimeResponseBody) SetRequestId(v string) *SyncDeviceTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDeviceTimeResponseBody) SetMessage(v string) *SyncDeviceTimeResponseBody {
	s.Message = &v
	return s
}

func (s *SyncDeviceTimeResponseBody) SetRetryInterval(v string) *SyncDeviceTimeResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *SyncDeviceTimeResponseBody) SetNTPServer(v string) *SyncDeviceTimeResponseBody {
	s.NTPServer = &v
	return s
}

func (s *SyncDeviceTimeResponseBody) SetCode(v string) *SyncDeviceTimeResponseBody {
	s.Code = &v
	return s
}

func (s *SyncDeviceTimeResponseBody) SetTimeStamp(v string) *SyncDeviceTimeResponseBody {
	s.TimeStamp = &v
	return s
}

type SyncDeviceTimeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncDeviceTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncDeviceTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceTimeResponse) GoString() string {
	return s.String()
}

func (s *SyncDeviceTimeResponse) SetHeaders(v map[string]*string) *SyncDeviceTimeResponse {
	s.Headers = v
	return s
}

func (s *SyncDeviceTimeResponse) SetBody(v *SyncDeviceTimeResponseBody) *SyncDeviceTimeResponse {
	s.Body = v
	return s
}

type UnbindCorpGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	CorpGroupId *string `json:"CorpGroupId,omitempty" xml:"CorpGroupId,omitempty"`
}

func (s UnbindCorpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindCorpGroupRequest) GoString() string {
	return s.String()
}

func (s *UnbindCorpGroupRequest) SetCorpId(v string) *UnbindCorpGroupRequest {
	s.CorpId = &v
	return s
}

func (s *UnbindCorpGroupRequest) SetCorpGroupId(v string) *UnbindCorpGroupRequest {
	s.CorpGroupId = &v
	return s
}

type UnbindCorpGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindCorpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindCorpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindCorpGroupResponseBody) SetMessage(v string) *UnbindCorpGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindCorpGroupResponseBody) SetRequestId(v string) *UnbindCorpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindCorpGroupResponseBody) SetCode(v string) *UnbindCorpGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindCorpGroupResponseBody) SetSuccess(v bool) *UnbindCorpGroupResponseBody {
	s.Success = &v
	return s
}

type UnbindCorpGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindCorpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *UnbindCorpGroupResponse) SetHeaders(v map[string]*string) *UnbindCorpGroupResponse {
	s.Headers = v
	return s
}

func (s *UnbindCorpGroupResponse) SetBody(v *UnbindCorpGroupResponseBody) *UnbindCorpGroupResponse {
	s.Body = v
	return s
}

type UnbindPersonRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	ProfileId *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
}

func (s UnbindPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindPersonRequest) GoString() string {
	return s.String()
}

func (s *UnbindPersonRequest) SetCorpId(v string) *UnbindPersonRequest {
	s.CorpId = &v
	return s
}

func (s *UnbindPersonRequest) SetIsvSubId(v string) *UnbindPersonRequest {
	s.IsvSubId = &v
	return s
}

func (s *UnbindPersonRequest) SetProfileId(v int64) *UnbindPersonRequest {
	s.ProfileId = &v
	return s
}

type UnbindPersonResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnbindPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindPersonResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPersonResponseBody) SetMessage(v string) *UnbindPersonResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindPersonResponseBody) SetRequestId(v string) *UnbindPersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPersonResponseBody) SetData(v bool) *UnbindPersonResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindPersonResponseBody) SetCode(v string) *UnbindPersonResponseBody {
	s.Code = &v
	return s
}

type UnbindPersonResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindPersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindPersonResponse) GoString() string {
	return s.String()
}

func (s *UnbindPersonResponse) SetHeaders(v map[string]*string) *UnbindPersonResponse {
	s.Headers = v
	return s
}

func (s *UnbindPersonResponse) SetBody(v *UnbindPersonResponseBody) *UnbindPersonResponse {
	s.Body = v
	return s
}

type UnbindUserRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnbindUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindUserRequest) SetCorpId(v string) *UnbindUserRequest {
	s.CorpId = &v
	return s
}

func (s *UnbindUserRequest) SetIsvSubId(v string) *UnbindUserRequest {
	s.IsvSubId = &v
	return s
}

func (s *UnbindUserRequest) SetUserId(v int64) *UnbindUserRequest {
	s.UserId = &v
	return s
}

type UnbindUserResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnbindUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindUserResponseBody) SetMessage(v string) *UnbindUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindUserResponseBody) SetRequestId(v string) *UnbindUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindUserResponseBody) SetData(v bool) *UnbindUserResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindUserResponseBody) SetCode(v string) *UnbindUserResponseBody {
	s.Code = &v
	return s
}

type UnbindUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserResponse) SetHeaders(v map[string]*string) *UnbindUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindUserResponse) SetBody(v *UnbindUserResponseBody) *UnbindUserResponse {
	s.Body = v
	return s
}

type UnsubscribeDeviceEventRequest struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s UnsubscribeDeviceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeDeviceEventRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeDeviceEventRequest) SetDeviceId(v string) *UnsubscribeDeviceEventRequest {
	s.DeviceId = &v
	return s
}

type UnsubscribeDeviceEventResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnsubscribeDeviceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeDeviceEventResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeDeviceEventResponseBody) SetMessage(v string) *UnsubscribeDeviceEventResponseBody {
	s.Message = &v
	return s
}

func (s *UnsubscribeDeviceEventResponseBody) SetRequestId(v string) *UnsubscribeDeviceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeDeviceEventResponseBody) SetCode(v string) *UnsubscribeDeviceEventResponseBody {
	s.Code = &v
	return s
}

type UnsubscribeDeviceEventResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnsubscribeDeviceEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsubscribeDeviceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeDeviceEventResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeDeviceEventResponse) SetHeaders(v map[string]*string) *UnsubscribeDeviceEventResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeDeviceEventResponse) SetBody(v *UnsubscribeDeviceEventResponseBody) *UnsubscribeDeviceEventResponse {
	s.Body = v
	return s
}

type UnsubscribeSpaceEventRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UnsubscribeSpaceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeSpaceEventRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeSpaceEventRequest) SetSpaceId(v string) *UnsubscribeSpaceEventRequest {
	s.SpaceId = &v
	return s
}

type UnsubscribeSpaceEventResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnsubscribeSpaceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeSpaceEventResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeSpaceEventResponseBody) SetMessage(v string) *UnsubscribeSpaceEventResponseBody {
	s.Message = &v
	return s
}

func (s *UnsubscribeSpaceEventResponseBody) SetRequestId(v string) *UnsubscribeSpaceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeSpaceEventResponseBody) SetCode(v string) *UnsubscribeSpaceEventResponseBody {
	s.Code = &v
	return s
}

type UnsubscribeSpaceEventResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnsubscribeSpaceEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsubscribeSpaceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeSpaceEventResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeSpaceEventResponse) SetHeaders(v map[string]*string) *UnsubscribeSpaceEventResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeSpaceEventResponse) SetBody(v *UnsubscribeSpaceEventResponseBody) *UnsubscribeSpaceEventResponse {
	s.Body = v
	return s
}

type UpdateCorpRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	CorpName     *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ParentCorpId *string `json:"ParentCorpId,omitempty" xml:"ParentCorpId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	IconPath     *string `json:"IconPath,omitempty" xml:"IconPath,omitempty"`
}

func (s UpdateCorpRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCorpRequest) GoString() string {
	return s.String()
}

func (s *UpdateCorpRequest) SetCorpId(v string) *UpdateCorpRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateCorpRequest) SetCorpName(v string) *UpdateCorpRequest {
	s.CorpName = &v
	return s
}

func (s *UpdateCorpRequest) SetAppName(v string) *UpdateCorpRequest {
	s.AppName = &v
	return s
}

func (s *UpdateCorpRequest) SetParentCorpId(v string) *UpdateCorpRequest {
	s.ParentCorpId = &v
	return s
}

func (s *UpdateCorpRequest) SetDescription(v string) *UpdateCorpRequest {
	s.Description = &v
	return s
}

func (s *UpdateCorpRequest) SetIsvSubId(v string) *UpdateCorpRequest {
	s.IsvSubId = &v
	return s
}

func (s *UpdateCorpRequest) SetIconPath(v string) *UpdateCorpRequest {
	s.IconPath = &v
	return s
}

type UpdateCorpResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateCorpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCorpResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCorpResponseBody) SetMessage(v string) *UpdateCorpResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCorpResponseBody) SetRequestId(v string) *UpdateCorpResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCorpResponseBody) SetData(v string) *UpdateCorpResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCorpResponseBody) SetCode(v string) *UpdateCorpResponseBody {
	s.Code = &v
	return s
}

type UpdateCorpResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCorpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCorpResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCorpResponse) GoString() string {
	return s.String()
}

func (s *UpdateCorpResponse) SetHeaders(v map[string]*string) *UpdateCorpResponse {
	s.Headers = v
	return s
}

func (s *UpdateCorpResponse) SetBody(v *UpdateCorpResponseBody) *UpdateCorpResponse {
	s.Body = v
	return s
}

type UpdateDeviceRequest struct {
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType       *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceAddress    *string `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty"`
	DeviceSite       *string `json:"DeviceSite,omitempty" xml:"DeviceSite,omitempty"`
	DeviceDirection  *string `json:"DeviceDirection,omitempty" xml:"DeviceDirection,omitempty"`
	DeviceResolution *string `json:"DeviceResolution,omitempty" xml:"DeviceResolution,omitempty"`
	BitRate          *string `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Vendor           *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s UpdateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceRequest) SetGbId(v string) *UpdateDeviceRequest {
	s.GbId = &v
	return s
}

func (s *UpdateDeviceRequest) SetDeviceName(v string) *UpdateDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateDeviceRequest) SetDeviceType(v string) *UpdateDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *UpdateDeviceRequest) SetDeviceAddress(v string) *UpdateDeviceRequest {
	s.DeviceAddress = &v
	return s
}

func (s *UpdateDeviceRequest) SetDeviceSite(v string) *UpdateDeviceRequest {
	s.DeviceSite = &v
	return s
}

func (s *UpdateDeviceRequest) SetDeviceDirection(v string) *UpdateDeviceRequest {
	s.DeviceDirection = &v
	return s
}

func (s *UpdateDeviceRequest) SetDeviceResolution(v string) *UpdateDeviceRequest {
	s.DeviceResolution = &v
	return s
}

func (s *UpdateDeviceRequest) SetBitRate(v string) *UpdateDeviceRequest {
	s.BitRate = &v
	return s
}

func (s *UpdateDeviceRequest) SetCorpId(v string) *UpdateDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateDeviceRequest) SetVendor(v string) *UpdateDeviceRequest {
	s.Vendor = &v
	return s
}

type UpdateDeviceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResponseBody) SetMessage(v string) *UpdateDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeviceResponseBody) SetRequestId(v string) *UpdateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceResponseBody) SetData(v string) *UpdateDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDeviceResponseBody) SetCode(v string) *UpdateDeviceResponseBody {
	s.Code = &v
	return s
}

type UpdateDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResponse) SetHeaders(v map[string]*string) *UpdateDeviceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceResponse) SetBody(v *UpdateDeviceResponseBody) *UpdateDeviceResponse {
	s.Body = v
	return s
}

type UpdateDeviceCaptureStrategyRequest struct {
	// 设备类型
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// 设备通道
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	// 周一图片抓去模式
	MondayCaptureStrategy *string `json:"MondayCaptureStrategy,omitempty" xml:"MondayCaptureStrategy,omitempty"`
}

func (s UpdateDeviceCaptureStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCaptureStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCaptureStrategyRequest) SetDeviceType(v string) *UpdateDeviceCaptureStrategyRequest {
	s.DeviceType = &v
	return s
}

func (s *UpdateDeviceCaptureStrategyRequest) SetDeviceCode(v string) *UpdateDeviceCaptureStrategyRequest {
	s.DeviceCode = &v
	return s
}

func (s *UpdateDeviceCaptureStrategyRequest) SetMondayCaptureStrategy(v string) *UpdateDeviceCaptureStrategyRequest {
	s.MondayCaptureStrategy = &v
	return s
}

type UpdateDeviceCaptureStrategyResponseBody struct {
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateDeviceCaptureStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCaptureStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCaptureStrategyResponseBody) SetRequestId(v string) *UpdateDeviceCaptureStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceCaptureStrategyResponseBody) SetCode(v string) *UpdateDeviceCaptureStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDeviceCaptureStrategyResponseBody) SetMessage(v string) *UpdateDeviceCaptureStrategyResponseBody {
	s.Message = &v
	return s
}

type UpdateDeviceCaptureStrategyResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceCaptureStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceCaptureStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCaptureStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCaptureStrategyResponse) SetHeaders(v map[string]*string) *UpdateDeviceCaptureStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceCaptureStrategyResponse) SetBody(v *UpdateDeviceCaptureStrategyResponseBody) *UpdateDeviceCaptureStrategyResponse {
	s.Body = v
	return s
}

type UpdateMonitorRequest struct {
	CorpId               *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	DeviceOperateType    *string `json:"DeviceOperateType,omitempty" xml:"DeviceOperateType,omitempty"`
	DeviceList           *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	PicOperateType       *string `json:"PicOperateType,omitempty" xml:"PicOperateType,omitempty"`
	PicList              *string `json:"PicList,omitempty" xml:"PicList,omitempty"`
	AttributeOperateType *string `json:"AttributeOperateType,omitempty" xml:"AttributeOperateType,omitempty"`
	AttributeName        *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	AttributeValueList   *string `json:"AttributeValueList,omitempty" xml:"AttributeValueList,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleExpression       *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	AlgorithmVendor      *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
	NotifierType         *string `json:"NotifierType,omitempty" xml:"NotifierType,omitempty"`
	NotifierUrl          *string `json:"NotifierUrl,omitempty" xml:"NotifierUrl,omitempty"`
	NotifierAppSecret    *string `json:"NotifierAppSecret,omitempty" xml:"NotifierAppSecret,omitempty"`
	NotifierTimeOut      *int32  `json:"NotifierTimeOut,omitempty" xml:"NotifierTimeOut,omitempty"`
	NotifierExtendValues *string `json:"NotifierExtendValues,omitempty" xml:"NotifierExtendValues,omitempty"`
}

func (s UpdateMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorRequest) SetCorpId(v string) *UpdateMonitorRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateMonitorRequest) SetTaskId(v string) *UpdateMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateMonitorRequest) SetRuleName(v string) *UpdateMonitorRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateMonitorRequest) SetDeviceOperateType(v string) *UpdateMonitorRequest {
	s.DeviceOperateType = &v
	return s
}

func (s *UpdateMonitorRequest) SetDeviceList(v string) *UpdateMonitorRequest {
	s.DeviceList = &v
	return s
}

func (s *UpdateMonitorRequest) SetPicOperateType(v string) *UpdateMonitorRequest {
	s.PicOperateType = &v
	return s
}

func (s *UpdateMonitorRequest) SetPicList(v string) *UpdateMonitorRequest {
	s.PicList = &v
	return s
}

func (s *UpdateMonitorRequest) SetAttributeOperateType(v string) *UpdateMonitorRequest {
	s.AttributeOperateType = &v
	return s
}

func (s *UpdateMonitorRequest) SetAttributeName(v string) *UpdateMonitorRequest {
	s.AttributeName = &v
	return s
}

func (s *UpdateMonitorRequest) SetAttributeValueList(v string) *UpdateMonitorRequest {
	s.AttributeValueList = &v
	return s
}

func (s *UpdateMonitorRequest) SetDescription(v string) *UpdateMonitorRequest {
	s.Description = &v
	return s
}

func (s *UpdateMonitorRequest) SetRuleExpression(v string) *UpdateMonitorRequest {
	s.RuleExpression = &v
	return s
}

func (s *UpdateMonitorRequest) SetAlgorithmVendor(v string) *UpdateMonitorRequest {
	s.AlgorithmVendor = &v
	return s
}

func (s *UpdateMonitorRequest) SetNotifierType(v string) *UpdateMonitorRequest {
	s.NotifierType = &v
	return s
}

func (s *UpdateMonitorRequest) SetNotifierUrl(v string) *UpdateMonitorRequest {
	s.NotifierUrl = &v
	return s
}

func (s *UpdateMonitorRequest) SetNotifierAppSecret(v string) *UpdateMonitorRequest {
	s.NotifierAppSecret = &v
	return s
}

func (s *UpdateMonitorRequest) SetNotifierTimeOut(v int32) *UpdateMonitorRequest {
	s.NotifierTimeOut = &v
	return s
}

func (s *UpdateMonitorRequest) SetNotifierExtendValues(v string) *UpdateMonitorRequest {
	s.NotifierExtendValues = &v
	return s
}

type UpdateMonitorResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMonitorResponseBody) SetMessage(v string) *UpdateMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMonitorResponseBody) SetRequestId(v string) *UpdateMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMonitorResponseBody) SetData(v string) *UpdateMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMonitorResponseBody) SetCode(v string) *UpdateMonitorResponseBody {
	s.Code = &v
	return s
}

type UpdateMonitorResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorResponse) SetHeaders(v map[string]*string) *UpdateMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorResponse) SetBody(v *UpdateMonitorResponseBody) *UpdateMonitorResponse {
	s.Body = v
	return s
}

type UpdateProfileRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender      *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ProfileId   *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
}

func (s UpdateProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateProfileRequest) SetCorpId(v string) *UpdateProfileRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateProfileRequest) SetIsvSubId(v string) *UpdateProfileRequest {
	s.IsvSubId = &v
	return s
}

func (s *UpdateProfileRequest) SetName(v string) *UpdateProfileRequest {
	s.Name = &v
	return s
}

func (s *UpdateProfileRequest) SetCatalogId(v int64) *UpdateProfileRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateProfileRequest) SetIdNumber(v string) *UpdateProfileRequest {
	s.IdNumber = &v
	return s
}

func (s *UpdateProfileRequest) SetFaceUrl(v string) *UpdateProfileRequest {
	s.FaceUrl = &v
	return s
}

func (s *UpdateProfileRequest) SetLiveAddress(v string) *UpdateProfileRequest {
	s.LiveAddress = &v
	return s
}

func (s *UpdateProfileRequest) SetGender(v int32) *UpdateProfileRequest {
	s.Gender = &v
	return s
}

func (s *UpdateProfileRequest) SetPlateNo(v string) *UpdateProfileRequest {
	s.PlateNo = &v
	return s
}

func (s *UpdateProfileRequest) SetPhoneNo(v string) *UpdateProfileRequest {
	s.PhoneNo = &v
	return s
}

func (s *UpdateProfileRequest) SetSceneType(v string) *UpdateProfileRequest {
	s.SceneType = &v
	return s
}

func (s *UpdateProfileRequest) SetBizId(v string) *UpdateProfileRequest {
	s.BizId = &v
	return s
}

func (s *UpdateProfileRequest) SetProfileId(v int64) *UpdateProfileRequest {
	s.ProfileId = &v
	return s
}

type UpdateProfileResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProfileResponseBody) SetMessage(v string) *UpdateProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProfileResponseBody) SetRequestId(v string) *UpdateProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProfileResponseBody) SetData(v string) *UpdateProfileResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProfileResponseBody) SetCode(v string) *UpdateProfileResponseBody {
	s.Code = &v
	return s
}

type UpdateProfileResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateProfileResponse) SetHeaders(v map[string]*string) *UpdateProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateProfileResponse) SetBody(v *UpdateProfileResponseBody) *UpdateProfileResponse {
	s.Body = v
	return s
}

type UpdateProfileCatalogRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
}

func (s UpdateProfileCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileCatalogRequest) GoString() string {
	return s.String()
}

func (s *UpdateProfileCatalogRequest) SetCorpId(v string) *UpdateProfileCatalogRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateProfileCatalogRequest) SetIsvSubId(v string) *UpdateProfileCatalogRequest {
	s.IsvSubId = &v
	return s
}

func (s *UpdateProfileCatalogRequest) SetCatalogId(v int64) *UpdateProfileCatalogRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateProfileCatalogRequest) SetCatalogName(v string) *UpdateProfileCatalogRequest {
	s.CatalogName = &v
	return s
}

type UpdateProfileCatalogResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateProfileCatalogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateProfileCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProfileCatalogResponseBody) SetMessage(v string) *UpdateProfileCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProfileCatalogResponseBody) SetRequestId(v string) *UpdateProfileCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProfileCatalogResponseBody) SetData(v *UpdateProfileCatalogResponseBodyData) *UpdateProfileCatalogResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProfileCatalogResponseBody) SetCode(v string) *UpdateProfileCatalogResponseBody {
	s.Code = &v
	return s
}

type UpdateProfileCatalogResponseBodyData struct {
	CatalogId       *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	IsvSubId        *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	ParentCatalogId *string `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty"`
	ProfileCount    *int64  `json:"ProfileCount,omitempty" xml:"ProfileCount,omitempty"`
}

func (s UpdateProfileCatalogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileCatalogResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProfileCatalogResponseBodyData) SetCatalogId(v int64) *UpdateProfileCatalogResponseBodyData {
	s.CatalogId = &v
	return s
}

func (s *UpdateProfileCatalogResponseBodyData) SetCatalogName(v string) *UpdateProfileCatalogResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *UpdateProfileCatalogResponseBodyData) SetIsvSubId(v string) *UpdateProfileCatalogResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *UpdateProfileCatalogResponseBodyData) SetParentCatalogId(v string) *UpdateProfileCatalogResponseBodyData {
	s.ParentCatalogId = &v
	return s
}

func (s *UpdateProfileCatalogResponseBodyData) SetProfileCount(v int64) *UpdateProfileCatalogResponseBodyData {
	s.ProfileCount = &v
	return s
}

type UpdateProfileCatalogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProfileCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProfileCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileCatalogResponse) GoString() string {
	return s.String()
}

func (s *UpdateProfileCatalogResponse) SetHeaders(v map[string]*string) *UpdateProfileCatalogResponse {
	s.Headers = v
	return s
}

func (s *UpdateProfileCatalogResponse) SetBody(v *UpdateProfileCatalogResponseBody) *UpdateProfileCatalogResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId         *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId      *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdNumber         *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl     *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceImageContent *string `json:"FaceImageContent,omitempty" xml:"FaceImageContent,omitempty"`
	Address          *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Age              *int32  `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender           *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo          *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo          *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment       *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UserId           *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetCorpId(v string) *UpdateUserRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateUserRequest) SetIsvSubId(v string) *UpdateUserRequest {
	s.IsvSubId = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserRequest) SetUserGroupId(v int64) *UpdateUserRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserRequest) SetIdNumber(v string) *UpdateUserRequest {
	s.IdNumber = &v
	return s
}

func (s *UpdateUserRequest) SetFaceImageUrl(v string) *UpdateUserRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *UpdateUserRequest) SetFaceImageContent(v string) *UpdateUserRequest {
	s.FaceImageContent = &v
	return s
}

func (s *UpdateUserRequest) SetAddress(v string) *UpdateUserRequest {
	s.Address = &v
	return s
}

func (s *UpdateUserRequest) SetAge(v int32) *UpdateUserRequest {
	s.Age = &v
	return s
}

func (s *UpdateUserRequest) SetGender(v int32) *UpdateUserRequest {
	s.Gender = &v
	return s
}

func (s *UpdateUserRequest) SetPlateNo(v string) *UpdateUserRequest {
	s.PlateNo = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNo(v string) *UpdateUserRequest {
	s.PhoneNo = &v
	return s
}

func (s *UpdateUserRequest) SetAttachment(v string) *UpdateUserRequest {
	s.Attachment = &v
	return s
}

func (s *UpdateUserRequest) SetBizId(v string) *UpdateUserRequest {
	s.BizId = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v int64) *UpdateUserRequest {
	s.UserId = &v
	return s
}

type UpdateUserResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetData(v string) *UpdateUserResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

type UpdateUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateUserGroupRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId      *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId   *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) SetCorpId(v string) *UpdateUserGroupRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetIsvSubId(v string) *UpdateUserGroupRequest {
	s.IsvSubId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupId(v int64) *UpdateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupName(v string) *UpdateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

type UpdateUserGroupResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) SetMessage(v string) *UpdateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetData(v *UpdateUserGroupResponseBodyData) *UpdateUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUserGroupResponseBody) SetCode(v string) *UpdateUserGroupResponseBody {
	s.Code = &v
	return s
}

type UpdateUserGroupResponseBodyData struct {
	IsvSubId          *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty"`
	UserGroupId       *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	UserCount         *int64  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
}

func (s UpdateUserGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBodyData) SetIsvSubId(v string) *UpdateUserGroupResponseBodyData {
	s.IsvSubId = &v
	return s
}

func (s *UpdateUserGroupResponseBodyData) SetUserGroupId(v int64) *UpdateUserGroupResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupResponseBodyData) SetUserGroupName(v string) *UpdateUserGroupResponseBodyData {
	s.UserGroupName = &v
	return s
}

func (s *UpdateUserGroupResponseBodyData) SetUserCount(v int64) *UpdateUserGroupResponseBodyData {
	s.UserCount = &v
	return s
}

func (s *UpdateUserGroupResponseBodyData) SetParentUserGroupId(v string) *UpdateUserGroupResponseBodyData {
	s.ParentUserGroupId = &v
	return s
}

type UpdateUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponse) SetHeaders(v map[string]*string) *UpdateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupResponse) SetBody(v *UpdateUserGroupResponseBody) *UpdateUserGroupResponse {
	s.Body = v
	return s
}

type UploadFileRequest struct {
	FileType      *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	MD5           *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	FileContent   *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileAliasName *string `json:"FileAliasName,omitempty" xml:"FileAliasName,omitempty"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	FilePath      *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s UploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFileRequest) GoString() string {
	return s.String()
}

func (s *UploadFileRequest) SetFileType(v string) *UploadFileRequest {
	s.FileType = &v
	return s
}

func (s *UploadFileRequest) SetMD5(v string) *UploadFileRequest {
	s.MD5 = &v
	return s
}

func (s *UploadFileRequest) SetCorpId(v string) *UploadFileRequest {
	s.CorpId = &v
	return s
}

func (s *UploadFileRequest) SetFileContent(v string) *UploadFileRequest {
	s.FileContent = &v
	return s
}

func (s *UploadFileRequest) SetFileName(v string) *UploadFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadFileRequest) SetFileAliasName(v string) *UploadFileRequest {
	s.FileAliasName = &v
	return s
}

func (s *UploadFileRequest) SetDataSourceId(v string) *UploadFileRequest {
	s.DataSourceId = &v
	return s
}

func (s *UploadFileRequest) SetFilePath(v string) *UploadFileRequest {
	s.FilePath = &v
	return s
}

type UploadFileResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UploadFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBody) SetMessage(v string) *UploadFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadFileResponseBody) SetRequestId(v string) *UploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFileResponseBody) SetData(v *UploadFileResponseBodyData) *UploadFileResponseBody {
	s.Data = v
	return s
}

func (s *UploadFileResponseBody) SetCode(v string) *UploadFileResponseBody {
	s.Code = &v
	return s
}

type UploadFileResponseBodyData struct {
	Records []*UploadFileResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s UploadFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBodyData) SetRecords(v []*UploadFileResponseBodyDataRecords) *UploadFileResponseBodyData {
	s.Records = v
	return s
}

type UploadFileResponseBodyDataRecords struct {
	OssPath  *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s UploadFileResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBodyDataRecords) SetOssPath(v string) *UploadFileResponseBodyDataRecords {
	s.OssPath = &v
	return s
}

func (s *UploadFileResponseBodyDataRecords) SetSourceId(v string) *UploadFileResponseBodyDataRecords {
	s.SourceId = &v
	return s
}

type UploadFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponse) GoString() string {
	return s.String()
}

func (s *UploadFileResponse) SetHeaders(v map[string]*string) *UploadFileResponse {
	s.Headers = v
	return s
}

func (s *UploadFileResponse) SetBody(v *UploadFileResponseBody) *UploadFileResponse {
	s.Body = v
	return s
}

type UploadImageRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s UploadImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadImageRequest) GoString() string {
	return s.String()
}

func (s *UploadImageRequest) SetImageUrl(v string) *UploadImageRequest {
	s.ImageUrl = &v
	return s
}

type UploadImageResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UploadImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponseBody) GoString() string {
	return s.String()
}

func (s *UploadImageResponseBody) SetMessage(v string) *UploadImageResponseBody {
	s.Message = &v
	return s
}

func (s *UploadImageResponseBody) SetRequestId(v string) *UploadImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadImageResponseBody) SetData(v string) *UploadImageResponseBody {
	s.Data = &v
	return s
}

func (s *UploadImageResponseBody) SetCode(v string) *UploadImageResponseBody {
	s.Code = &v
	return s
}

type UploadImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponse) GoString() string {
	return s.String()
}

func (s *UploadImageResponse) SetHeaders(v map[string]*string) *UploadImageResponse {
	s.Headers = v
	return s
}

func (s *UploadImageResponse) SetBody(v *UploadImageResponseBody) *UploadImageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("vcs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddDataSourceWithOptions(request *AddDataSourceRequest, runtime *util.RuntimeOptions) (_result *AddDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDataSource"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDataSource(request *AddDataSourceRequest) (_result *AddDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDataSourceResponse{}
	_body, _err := client.AddDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDeviceWithOptions(request *AddDeviceRequest, runtime *util.RuntimeOptions) (_result *AddDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDevice(request *AddDeviceRequest) (_result *AddDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDeviceResponse{}
	_body, _err := client.AddDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMonitorWithOptions(request *AddMonitorRequest, runtime *util.RuntimeOptions) (_result *AddMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddMonitor"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMonitor(request *AddMonitorRequest) (_result *AddMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMonitorResponse{}
	_body, _err := client.AddMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProfileWithOptions(request *AddProfileRequest, runtime *util.RuntimeOptions) (_result *AddProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddProfileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddProfile"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProfile(request *AddProfileRequest) (_result *AddProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddProfileResponse{}
	_body, _err := client.AddProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProfileCatalogWithOptions(request *AddProfileCatalogRequest, runtime *util.RuntimeOptions) (_result *AddProfileCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddProfileCatalogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddProfileCatalog"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProfileCatalog(request *AddProfileCatalogRequest) (_result *AddProfileCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddProfileCatalogResponse{}
	_body, _err := client.AddProfileCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindCorpGroupWithOptions(request *BindCorpGroupRequest, runtime *util.RuntimeOptions) (_result *BindCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindCorpGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindCorpGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindCorpGroup(request *BindCorpGroupRequest) (_result *BindCorpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindCorpGroupResponse{}
	_body, _err := client.BindCorpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindPersonWithOptions(request *BindPersonRequest, runtime *util.RuntimeOptions) (_result *BindPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindPersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindPerson"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindPerson(request *BindPersonRequest) (_result *BindPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindPersonResponse{}
	_body, _err := client.BindPersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindUserWithOptions(request *BindUserRequest, runtime *util.RuntimeOptions) (_result *BindUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindUser"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindUser(request *BindUserRequest) (_result *BindUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindUserResponse{}
	_body, _err := client.BindUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCorpWithOptions(request *CreateCorpRequest, runtime *util.RuntimeOptions) (_result *CreateCorpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCorpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCorp"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCorp(request *CreateCorpRequest) (_result *CreateCorpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCorpResponse{}
	_body, _err := client.CreateCorpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCorpGroupWithOptions(request *CreateCorpGroupRequest, runtime *util.RuntimeOptions) (_result *CreateCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCorpGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCorpGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCorpGroup(request *CreateCorpGroupRequest) (_result *CreateCorpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCorpGroupResponse{}
	_body, _err := client.CreateCorpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUser"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUserGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoComposeTaskWithOptions(request *CreateVideoComposeTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoComposeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVideoComposeTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVideoComposeTask"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoComposeTask(request *CreateVideoComposeTaskRequest) (_result *CreateVideoComposeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoComposeTaskResponse{}
	_body, _err := client.CreateVideoComposeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoSummaryTaskWithOptions(request *CreateVideoSummaryTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVideoSummaryTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVideoSummaryTask"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoSummaryTask(request *CreateVideoSummaryTaskRequest) (_result *CreateVideoSummaryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoSummaryTaskResponse{}
	_body, _err := client.CreateVideoSummaryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteChannelWithOptions(request *DeleteChannelRequest, runtime *util.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteChannel"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteChannel(request *DeleteChannelRequest) (_result *DeleteChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChannelResponse{}
	_body, _err := client.DeleteChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCorpGroupWithOptions(request *DeleteCorpGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCorpGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCorpGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCorpGroup(request *DeleteCorpGroupRequest) (_result *DeleteCorpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCorpGroupResponse{}
	_body, _err := client.DeleteCorpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDataSource"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteDeviceForInstanceWithOptions(tmpReq *DeleteDeviceForInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceForInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteDeviceForInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Devices)) {
		request.DevicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Devices, tea.String("Devices"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDeviceForInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDeviceForInstance"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceForInstance(request *DeleteDeviceForInstanceRequest) (_result *DeleteDeviceForInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceForInstanceResponse{}
	_body, _err := client.DeleteDeviceForInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIPCDeviceWithOptions(request *DeleteIPCDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteIPCDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIPCDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIPCDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIPCDevice(request *DeleteIPCDeviceRequest) (_result *DeleteIPCDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIPCDeviceResponse{}
	_body, _err := client.DeleteIPCDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNVRDeviceWithOptions(request *DeleteNVRDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteNVRDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNVRDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNVRDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNVRDevice(request *DeleteNVRDeviceRequest) (_result *DeleteNVRDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNVRDeviceResponse{}
	_body, _err := client.DeleteNVRDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProfileWithOptions(request *DeleteProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProfileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProfile"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProfile(request *DeleteProfileRequest) (_result *DeleteProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProfileResponse{}
	_body, _err := client.DeleteProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProfileCatalogWithOptions(request *DeleteProfileCatalogRequest, runtime *util.RuntimeOptions) (_result *DeleteProfileCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProfileCatalogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProfileCatalog"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProfileCatalog(request *DeleteProfileCatalogRequest) (_result *DeleteProfileCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProfileCatalogResponse{}
	_body, _err := client.DeleteProfileCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecordsWithOptions(request *DeleteRecordsRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRecords"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecords(request *DeleteRecordsRequest) (_result *DeleteRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordsResponse{}
	_body, _err := client.DeleteRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUser"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoSummaryTaskWithOptions(request *DeleteVideoSummaryTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVideoSummaryTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVideoSummaryTask"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideoSummaryTask(request *DeleteVideoSummaryTaskRequest) (_result *DeleteVideoSummaryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVideoSummaryTaskResponse{}
	_body, _err := client.DeleteVideoSummaryTaskWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDevices"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetBodyOptionsWithOptions(request *GetBodyOptionsRequest, runtime *util.RuntimeOptions) (_result *GetBodyOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBodyOptionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBodyOptions"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBodyOptions(request *GetBodyOptionsRequest) (_result *GetBodyOptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBodyOptionsResponse{}
	_body, _err := client.GetBodyOptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCatalogListWithOptions(request *GetCatalogListRequest, runtime *util.RuntimeOptions) (_result *GetCatalogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCatalogListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCatalogList"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCatalogList(request *GetCatalogListRequest) (_result *GetCatalogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCatalogListResponse{}
	_body, _err := client.GetCatalogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceCaptureStrategyWithOptions(request *GetDeviceCaptureStrategyRequest, runtime *util.RuntimeOptions) (_result *GetDeviceCaptureStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceCaptureStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceCaptureStrategy"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceCaptureStrategy(request *GetDeviceCaptureStrategyRequest) (_result *GetDeviceCaptureStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceCaptureStrategyResponse{}
	_body, _err := client.GetDeviceCaptureStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceConfigWithOptions(request *GetDeviceConfigRequest, runtime *util.RuntimeOptions) (_result *GetDeviceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceConfig"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceConfig(request *GetDeviceConfigRequest) (_result *GetDeviceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceConfigResponse{}
	_body, _err := client.GetDeviceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceLiveUrlWithOptions(request *GetDeviceLiveUrlRequest, runtime *util.RuntimeOptions) (_result *GetDeviceLiveUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceLiveUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceLiveUrl"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceLiveUrl(request *GetDeviceLiveUrlRequest) (_result *GetDeviceLiveUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceLiveUrlResponse{}
	_body, _err := client.GetDeviceLiveUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceVideoUrlWithOptions(request *GetDeviceVideoUrlRequest, runtime *util.RuntimeOptions) (_result *GetDeviceVideoUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceVideoUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceVideoUrl"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceVideoUrl(request *GetDeviceVideoUrlRequest) (_result *GetDeviceVideoUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceVideoUrlResponse{}
	_body, _err := client.GetDeviceVideoUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFaceModelResultWithOptions(request *GetFaceModelResultRequest, runtime *util.RuntimeOptions) (_result *GetFaceModelResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFaceModelResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFaceModelResult"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaceModelResult(request *GetFaceModelResultRequest) (_result *GetFaceModelResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFaceModelResultResponse{}
	_body, _err := client.GetFaceModelResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFaceOptionsWithOptions(request *GetFaceOptionsRequest, runtime *util.RuntimeOptions) (_result *GetFaceOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFaceOptionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFaceOptions"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaceOptions(request *GetFaceOptionsRequest) (_result *GetFaceOptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFaceOptionsResponse{}
	_body, _err := client.GetFaceOptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInventoryWithOptions(request *GetInventoryRequest, runtime *util.RuntimeOptions) (_result *GetInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInventoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInventory"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInventory(request *GetInventoryRequest) (_result *GetInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInventoryResponse{}
	_body, _err := client.GetInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMonitorListWithOptions(request *GetMonitorListRequest, runtime *util.RuntimeOptions) (_result *GetMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMonitorListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMonitorList"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMonitorList(request *GetMonitorListRequest) (_result *GetMonitorListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMonitorListResponse{}
	_body, _err := client.GetMonitorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMonitorResultWithOptions(request *GetMonitorResultRequest, runtime *util.RuntimeOptions) (_result *GetMonitorResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMonitorResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMonitorResult"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMonitorResult(request *GetMonitorResultRequest) (_result *GetMonitorResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMonitorResultResponse{}
	_body, _err := client.GetMonitorResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPersonDetailWithOptions(request *GetPersonDetailRequest, runtime *util.RuntimeOptions) (_result *GetPersonDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPersonDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPersonDetail"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPersonDetail(request *GetPersonDetailRequest) (_result *GetPersonDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPersonDetailResponse{}
	_body, _err := client.GetPersonDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPersonListWithOptions(tmpReq *GetPersonListRequest, runtime *util.RuntimeOptions) (_result *GetPersonListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPersonListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CorpIdList)) {
		request.CorpIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CorpIdList, tea.String("CorpIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PersonIdList)) {
		request.PersonIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonIdList, tea.String("PersonIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPersonListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPersonList"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPersonList(request *GetPersonListRequest) (_result *GetPersonListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPersonListResponse{}
	_body, _err := client.GetPersonListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProfileDetailWithOptions(request *GetProfileDetailRequest, runtime *util.RuntimeOptions) (_result *GetProfileDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProfileDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProfileDetail"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProfileDetail(request *GetProfileDetailRequest) (_result *GetProfileDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProfileDetailResponse{}
	_body, _err := client.GetProfileDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProfileListWithOptions(tmpReq *GetProfileListRequest, runtime *util.RuntimeOptions) (_result *GetProfileListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetProfileListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PersonIdList)) {
		request.PersonIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonIdList, tea.String("PersonIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ProfileIdList)) {
		request.ProfileIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProfileIdList, tea.String("ProfileIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProfileListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProfileList"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProfileList(request *GetProfileListRequest) (_result *GetProfileListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProfileListResponse{}
	_body, _err := client.GetProfileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserDetailWithOptions(request *GetUserDetailRequest, runtime *util.RuntimeOptions) (_result *GetUserDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserDetail"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserDetail(request *GetUserDetailRequest) (_result *GetUserDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserDetailResponse{}
	_body, _err := client.GetUserDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoComposeResultWithOptions(request *GetVideoComposeResultRequest, runtime *util.RuntimeOptions) (_result *GetVideoComposeResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoComposeResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoComposeResult"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoComposeResult(request *GetVideoComposeResultRequest) (_result *GetVideoComposeResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoComposeResultResponse{}
	_body, _err := client.GetVideoComposeResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoSummaryTaskResultWithOptions(request *GetVideoSummaryTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetVideoSummaryTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoSummaryTaskResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoSummaryTaskResult"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoSummaryTaskResult(request *GetVideoSummaryTaskResultRequest) (_result *GetVideoSummaryTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoSummaryTaskResultResponse{}
	_body, _err := client.GetVideoSummaryTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeMotorModelWithOptions(request *InvokeMotorModelRequest, runtime *util.RuntimeOptions) (_result *InvokeMotorModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InvokeMotorModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InvokeMotorModel"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeMotorModel(request *InvokeMotorModelRequest) (_result *InvokeMotorModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeMotorModelResponse{}
	_body, _err := client.InvokeMotorModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessNumberWithOptions(request *ListAccessNumberRequest, runtime *util.RuntimeOptions) (_result *ListAccessNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAccessNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAccessNumber"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessNumber(request *ListAccessNumberRequest) (_result *ListAccessNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessNumberResponse{}
	_body, _err := client.ListAccessNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBodyAlgorithmResultsWithOptions(request *ListBodyAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListBodyAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBodyAlgorithmResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBodyAlgorithmResults"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBodyAlgorithmResults(request *ListBodyAlgorithmResultsRequest) (_result *ListBodyAlgorithmResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBodyAlgorithmResultsResponse{}
	_body, _err := client.ListBodyAlgorithmResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorpGroupMetricsWithOptions(request *ListCorpGroupMetricsRequest, runtime *util.RuntimeOptions) (_result *ListCorpGroupMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCorpGroupMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCorpGroupMetrics"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorpGroupMetrics(request *ListCorpGroupMetricsRequest) (_result *ListCorpGroupMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorpGroupMetricsResponse{}
	_body, _err := client.ListCorpGroupMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorpGroupsWithOptions(request *ListCorpGroupsRequest, runtime *util.RuntimeOptions) (_result *ListCorpGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCorpGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCorpGroups"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorpGroups(request *ListCorpGroupsRequest) (_result *ListCorpGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorpGroupsResponse{}
	_body, _err := client.ListCorpGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorpMetricsWithOptions(request *ListCorpMetricsRequest, runtime *util.RuntimeOptions) (_result *ListCorpMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCorpMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCorpMetrics"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorpMetrics(request *ListCorpMetricsRequest) (_result *ListCorpMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorpMetricsResponse{}
	_body, _err := client.ListCorpMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorpsWithOptions(request *ListCorpsRequest, runtime *util.RuntimeOptions) (_result *ListCorpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCorpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCorps"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorps(request *ListCorpsRequest) (_result *ListCorpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorpsResponse{}
	_body, _err := client.ListCorpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceGroupsWithOptions(request *ListDeviceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeviceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceGroups"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceGroups(request *ListDeviceGroupsRequest) (_result *ListDeviceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceGroupsResponse{}
	_body, _err := client.ListDeviceGroupsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevices"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListEventAlgorithmDetailsWithOptions(request *ListEventAlgorithmDetailsRequest, runtime *util.RuntimeOptions) (_result *ListEventAlgorithmDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEventAlgorithmDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEventAlgorithmDetails"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventAlgorithmDetails(request *ListEventAlgorithmDetailsRequest) (_result *ListEventAlgorithmDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventAlgorithmDetailsResponse{}
	_body, _err := client.ListEventAlgorithmDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventAlgorithmResultsWithOptions(request *ListEventAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListEventAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEventAlgorithmResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEventAlgorithmResults"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventAlgorithmResults(request *ListEventAlgorithmResultsRequest) (_result *ListEventAlgorithmResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventAlgorithmResultsResponse{}
	_body, _err := client.ListEventAlgorithmResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceAlgorithmResultsWithOptions(request *ListFaceAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListFaceAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFaceAlgorithmResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaceAlgorithmResults"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceAlgorithmResults(request *ListFaceAlgorithmResultsRequest) (_result *ListFaceAlgorithmResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceAlgorithmResultsResponse{}
	_body, _err := client.ListFaceAlgorithmResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMetricsWithOptions(request *ListMetricsRequest, runtime *util.RuntimeOptions) (_result *ListMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMetrics"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMetrics(request *ListMetricsRequest) (_result *ListMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMetricsResponse{}
	_body, _err := client.ListMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMotorAlgorithmResultsWithOptions(request *ListMotorAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListMotorAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMotorAlgorithmResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMotorAlgorithmResults"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMotorAlgorithmResults(request *ListMotorAlgorithmResultsRequest) (_result *ListMotorAlgorithmResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMotorAlgorithmResultsResponse{}
	_body, _err := client.ListMotorAlgorithmResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNVRChannelDeviceWithOptions(request *ListNVRChannelDeviceRequest, runtime *util.RuntimeOptions) (_result *ListNVRChannelDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListNVRChannelDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNVRChannelDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNVRChannelDevice(request *ListNVRChannelDeviceRequest) (_result *ListNVRChannelDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNVRChannelDeviceResponse{}
	_body, _err := client.ListNVRChannelDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNVRDeviceWithOptions(request *ListNVRDeviceRequest, runtime *util.RuntimeOptions) (_result *ListNVRDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListNVRDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNVRDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNVRDevice(request *ListNVRDeviceRequest) (_result *ListNVRDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNVRDeviceResponse{}
	_body, _err := client.ListNVRDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonsWithOptions(request *ListPersonsRequest, runtime *util.RuntimeOptions) (_result *ListPersonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersons"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersons(request *ListPersonsRequest) (_result *ListPersonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonsResponse{}
	_body, _err := client.ListPersonsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonTraceWithOptions(request *ListPersonTraceRequest, runtime *util.RuntimeOptions) (_result *ListPersonTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonTrace"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonTrace(request *ListPersonTraceRequest) (_result *ListPersonTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonTraceResponse{}
	_body, _err := client.ListPersonTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonTraceDetailsWithOptions(request *ListPersonTraceDetailsRequest, runtime *util.RuntimeOptions) (_result *ListPersonTraceDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonTraceDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonTraceDetails"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonTraceDetails(request *ListPersonTraceDetailsRequest) (_result *ListPersonTraceDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonTraceDetailsResponse{}
	_body, _err := client.ListPersonTraceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonVisitCountWithOptions(request *ListPersonVisitCountRequest, runtime *util.RuntimeOptions) (_result *ListPersonVisitCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonVisitCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonVisitCount"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonVisitCount(request *ListPersonVisitCountRequest) (_result *ListPersonVisitCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonVisitCountResponse{}
	_body, _err := client.ListPersonVisitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubscribeDeviceWithOptions(request *ListSubscribeDeviceRequest, runtime *util.RuntimeOptions) (_result *ListSubscribeDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSubscribeDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSubscribeDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubscribeDevice(request *ListSubscribeDeviceRequest) (_result *ListSubscribeDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubscribeDeviceResponse{}
	_body, _err := client.ListSubscribeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserGroups"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(tmpReq *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PersonList)) {
		request.PersonListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonList, tea.String("PersonList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserList)) {
		request.UserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserList, tea.String("UserList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFaceQualityWithOptions(request *RecognizeFaceQualityRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeFaceQualityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeFaceQuality"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFaceQuality(request *RecognizeFaceQualityRequest) (_result *RecognizeFaceQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFaceQualityResponse{}
	_body, _err := client.RecognizeFaceQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeImageWithOptions(request *RecognizeImageRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeImage"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeImage(request *RecognizeImageRequest) (_result *RecognizeImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeImageResponse{}
	_body, _err := client.RecognizeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportDeviceCapacityWithOptions(request *ReportDeviceCapacityRequest, runtime *util.RuntimeOptions) (_result *ReportDeviceCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReportDeviceCapacityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReportDeviceCapacity"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportDeviceCapacity(request *ReportDeviceCapacityRequest) (_result *ReportDeviceCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportDeviceCapacityResponse{}
	_body, _err := client.ReportDeviceCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveVideoSummaryTaskVideoWithOptions(request *SaveVideoSummaryTaskVideoRequest, runtime *util.RuntimeOptions) (_result *SaveVideoSummaryTaskVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveVideoSummaryTaskVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveVideoSummaryTaskVideo"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveVideoSummaryTaskVideo(request *SaveVideoSummaryTaskVideoRequest) (_result *SaveVideoSummaryTaskVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveVideoSummaryTaskVideoResponse{}
	_body, _err := client.SaveVideoSummaryTaskVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchBodyWithOptions(tmpReq *SearchBodyRequest, runtime *util.RuntimeOptions) (_result *SearchBodyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchBodyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptionList)) {
		request.OptionListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptionList, tea.String("OptionList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchBodyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchBody"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchBody(request *SearchBodyRequest) (_result *SearchBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchBodyResponse{}
	_body, _err := client.SearchBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFaceWithOptions(tmpReq *SearchFaceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchFaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptionList)) {
		request.OptionListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptionList, tea.String("OptionList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchFace"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFace(request *SearchFaceRequest) (_result *SearchFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchFaceResponse{}
	_body, _err := client.SearchFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchObjectWithOptions(tmpReq *SearchObjectRequest, runtime *util.RuntimeOptions) (_result *SearchObjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchObjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceList)) {
		request.DeviceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceList, tea.String("DeviceList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Conditions)) {
		request.ConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditions, tea.String("Conditions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ImagePath)) {
		request.ImagePathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImagePath, tea.String("ImagePath"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchObject"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchObject(request *SearchObjectRequest) (_result *SearchObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchObjectResponse{}
	_body, _err := client.SearchObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMonitorWithOptions(request *StopMonitorRequest, runtime *util.RuntimeOptions) (_result *StopMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopMonitor"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMonitor(request *StopMonitorRequest) (_result *StopMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopMonitorResponse{}
	_body, _err := client.StopMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeDeviceEventWithOptions(request *SubscribeDeviceEventRequest, runtime *util.RuntimeOptions) (_result *SubscribeDeviceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubscribeDeviceEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubscribeDeviceEvent"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeDeviceEvent(request *SubscribeDeviceEventRequest) (_result *SubscribeDeviceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscribeDeviceEventResponse{}
	_body, _err := client.SubscribeDeviceEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeSpaceEventWithOptions(request *SubscribeSpaceEventRequest, runtime *util.RuntimeOptions) (_result *SubscribeSpaceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubscribeSpaceEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubscribeSpaceEvent"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeSpaceEvent(request *SubscribeSpaceEventRequest) (_result *SubscribeSpaceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscribeSpaceEventResponse{}
	_body, _err := client.SubscribeSpaceEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDeviceTimeWithOptions(request *SyncDeviceTimeRequest, runtime *util.RuntimeOptions) (_result *SyncDeviceTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncDeviceTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncDeviceTime"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncDeviceTime(request *SyncDeviceTimeRequest) (_result *SyncDeviceTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncDeviceTimeResponse{}
	_body, _err := client.SyncDeviceTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindCorpGroupWithOptions(request *UnbindCorpGroupRequest, runtime *util.RuntimeOptions) (_result *UnbindCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindCorpGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindCorpGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindCorpGroup(request *UnbindCorpGroupRequest) (_result *UnbindCorpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindCorpGroupResponse{}
	_body, _err := client.UnbindCorpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindPersonWithOptions(request *UnbindPersonRequest, runtime *util.RuntimeOptions) (_result *UnbindPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindPersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindPerson"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindPerson(request *UnbindPersonRequest) (_result *UnbindPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindPersonResponse{}
	_body, _err := client.UnbindPersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindUserWithOptions(request *UnbindUserRequest, runtime *util.RuntimeOptions) (_result *UnbindUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindUser"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindUser(request *UnbindUserRequest) (_result *UnbindUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindUserResponse{}
	_body, _err := client.UnbindUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsubscribeDeviceEventWithOptions(request *UnsubscribeDeviceEventRequest, runtime *util.RuntimeOptions) (_result *UnsubscribeDeviceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnsubscribeDeviceEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnsubscribeDeviceEvent"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsubscribeDeviceEvent(request *UnsubscribeDeviceEventRequest) (_result *UnsubscribeDeviceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnsubscribeDeviceEventResponse{}
	_body, _err := client.UnsubscribeDeviceEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsubscribeSpaceEventWithOptions(request *UnsubscribeSpaceEventRequest, runtime *util.RuntimeOptions) (_result *UnsubscribeSpaceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnsubscribeSpaceEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnsubscribeSpaceEvent"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsubscribeSpaceEvent(request *UnsubscribeSpaceEventRequest) (_result *UnsubscribeSpaceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnsubscribeSpaceEventResponse{}
	_body, _err := client.UnsubscribeSpaceEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCorpWithOptions(request *UpdateCorpRequest, runtime *util.RuntimeOptions) (_result *UpdateCorpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCorpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCorp"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCorp(request *UpdateCorpRequest) (_result *UpdateCorpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCorpResponse{}
	_body, _err := client.UpdateCorpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceWithOptions(request *UpdateDeviceRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevice"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevice(request *UpdateDeviceRequest) (_result *UpdateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceResponse{}
	_body, _err := client.UpdateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceCaptureStrategyWithOptions(request *UpdateDeviceCaptureStrategyRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceCaptureStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceCaptureStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceCaptureStrategy"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceCaptureStrategy(request *UpdateDeviceCaptureStrategyRequest) (_result *UpdateDeviceCaptureStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceCaptureStrategyResponse{}
	_body, _err := client.UpdateDeviceCaptureStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMonitorWithOptions(request *UpdateMonitorRequest, runtime *util.RuntimeOptions) (_result *UpdateMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMonitor"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMonitor(request *UpdateMonitorRequest) (_result *UpdateMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMonitorResponse{}
	_body, _err := client.UpdateMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProfileWithOptions(request *UpdateProfileRequest, runtime *util.RuntimeOptions) (_result *UpdateProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProfileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProfile"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProfile(request *UpdateProfileRequest) (_result *UpdateProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProfileResponse{}
	_body, _err := client.UpdateProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProfileCatalogWithOptions(request *UpdateProfileCatalogRequest, runtime *util.RuntimeOptions) (_result *UpdateProfileCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProfileCatalogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProfileCatalog"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProfileCatalog(request *UpdateProfileCatalogRequest) (_result *UpdateProfileCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProfileCatalogResponse{}
	_body, _err := client.UpdateProfileCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUser"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUserGroup"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFileWithOptions(request *UploadFileRequest, runtime *util.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadFile"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadFile(request *UploadFileRequest) (_result *UploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadFileResponse{}
	_body, _err := client.UploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadImageWithOptions(request *UploadImageRequest, runtime *util.RuntimeOptions) (_result *UploadImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadImage"), tea.String("2020-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadImage(request *UploadImageRequest) (_result *UploadImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadImageResponse{}
	_body, _err := client.UploadImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
