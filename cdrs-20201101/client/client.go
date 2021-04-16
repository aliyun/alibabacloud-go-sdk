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

type SearchObjectRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Vendor        *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	Feature       *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	ImageContent  *string `json:"ImageContent,omitempty" xml:"ImageContent,omitempty"`
	ImageUrl      *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	DeviceList    *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	Attributes    *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	ShotTimeStart *string `json:"ShotTimeStart,omitempty" xml:"ShotTimeStart,omitempty"`
	ShotTimeEnd   *string `json:"ShotTimeEnd,omitempty" xml:"ShotTimeEnd,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *SearchObjectRequest) SetVendor(v string) *SearchObjectRequest {
	s.Vendor = &v
	return s
}

func (s *SearchObjectRequest) SetFeature(v string) *SearchObjectRequest {
	s.Feature = &v
	return s
}

func (s *SearchObjectRequest) SetImageContent(v string) *SearchObjectRequest {
	s.ImageContent = &v
	return s
}

func (s *SearchObjectRequest) SetImageUrl(v string) *SearchObjectRequest {
	s.ImageUrl = &v
	return s
}

func (s *SearchObjectRequest) SetDeviceList(v string) *SearchObjectRequest {
	s.DeviceList = &v
	return s
}

func (s *SearchObjectRequest) SetAttributes(v string) *SearchObjectRequest {
	s.Attributes = &v
	return s
}

func (s *SearchObjectRequest) SetShotTimeStart(v string) *SearchObjectRequest {
	s.ShotTimeStart = &v
	return s
}

func (s *SearchObjectRequest) SetShotTimeEnd(v string) *SearchObjectRequest {
	s.ShotTimeEnd = &v
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

type SearchObjectResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize  *int64                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                        `json:"Total,omitempty" xml:"Total,omitempty"`
	Data      *SearchObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBody) SetRequestId(v string) *SearchObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchObjectResponseBody) SetMessage(v string) *SearchObjectResponseBody {
	s.Message = &v
	return s
}

func (s *SearchObjectResponseBody) SetPageSize(v int64) *SearchObjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchObjectResponseBody) SetTotal(v int32) *SearchObjectResponseBody {
	s.Total = &v
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

func (s *SearchObjectResponseBody) SetSuccess(v bool) *SearchObjectResponseBody {
	s.Success = &v
	return s
}

type SearchObjectResponseBodyData struct {
	BodyList     []*SearchObjectResponseBodyDataBodyList     `json:"BodyList,omitempty" xml:"BodyList,omitempty" type:"Repeated"`
	FaceList     []*SearchObjectResponseBodyDataFaceList     `json:"FaceList,omitempty" xml:"FaceList,omitempty" type:"Repeated"`
	MotorList    []*SearchObjectResponseBodyDataMotorList    `json:"MotorList,omitempty" xml:"MotorList,omitempty" type:"Repeated"`
	NonMotorList []*SearchObjectResponseBodyDataNonMotorList `json:"NonMotorList,omitempty" xml:"NonMotorList,omitempty" type:"Repeated"`
}

func (s SearchObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyData) SetBodyList(v []*SearchObjectResponseBodyDataBodyList) *SearchObjectResponseBodyData {
	s.BodyList = v
	return s
}

func (s *SearchObjectResponseBodyData) SetFaceList(v []*SearchObjectResponseBodyDataFaceList) *SearchObjectResponseBodyData {
	s.FaceList = v
	return s
}

func (s *SearchObjectResponseBodyData) SetMotorList(v []*SearchObjectResponseBodyDataMotorList) *SearchObjectResponseBodyData {
	s.MotorList = v
	return s
}

func (s *SearchObjectResponseBodyData) SetNonMotorList(v []*SearchObjectResponseBodyDataNonMotorList) *SearchObjectResponseBodyData {
	s.NonMotorList = v
	return s
}

type SearchObjectResponseBodyDataBodyList struct {
	SourceImageUrl *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	DeviceID       *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType     *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	TargetImageUrl *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	RightBottomY   *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY       *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	ShotTime       *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	RightBottomX   *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LeftTopX       *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s SearchObjectResponseBodyDataBodyList) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyDataBodyList) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyDataBodyList) SetSourceImageUrl(v string) *SearchObjectResponseBodyDataBodyList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetDeviceID(v string) *SearchObjectResponseBodyDataBodyList {
	s.DeviceID = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetObjectType(v string) *SearchObjectResponseBodyDataBodyList {
	s.ObjectType = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetTargetImageUrl(v string) *SearchObjectResponseBodyDataBodyList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetRightBottomY(v int32) *SearchObjectResponseBodyDataBodyList {
	s.RightBottomY = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetLeftTopY(v int32) *SearchObjectResponseBodyDataBodyList {
	s.LeftTopY = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetScore(v float32) *SearchObjectResponseBodyDataBodyList {
	s.Score = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetShotTime(v string) *SearchObjectResponseBodyDataBodyList {
	s.ShotTime = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetRightBottomX(v int32) *SearchObjectResponseBodyDataBodyList {
	s.RightBottomX = &v
	return s
}

func (s *SearchObjectResponseBodyDataBodyList) SetLeftTopX(v int32) *SearchObjectResponseBodyDataBodyList {
	s.LeftTopX = &v
	return s
}

type SearchObjectResponseBodyDataFaceList struct {
	SourceImageUrl *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	DeviceID       *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType     *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	TargetImageUrl *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	RightBottomY   *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY       *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	ShotTime       *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	RightBottomX   *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LeftTopX       *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s SearchObjectResponseBodyDataFaceList) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyDataFaceList) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyDataFaceList) SetSourceImageUrl(v string) *SearchObjectResponseBodyDataFaceList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetDeviceID(v string) *SearchObjectResponseBodyDataFaceList {
	s.DeviceID = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetObjectType(v string) *SearchObjectResponseBodyDataFaceList {
	s.ObjectType = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetTargetImageUrl(v string) *SearchObjectResponseBodyDataFaceList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetRightBottomY(v int32) *SearchObjectResponseBodyDataFaceList {
	s.RightBottomY = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetLeftTopY(v int32) *SearchObjectResponseBodyDataFaceList {
	s.LeftTopY = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetScore(v float32) *SearchObjectResponseBodyDataFaceList {
	s.Score = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetShotTime(v string) *SearchObjectResponseBodyDataFaceList {
	s.ShotTime = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetRightBottomX(v int32) *SearchObjectResponseBodyDataFaceList {
	s.RightBottomX = &v
	return s
}

func (s *SearchObjectResponseBodyDataFaceList) SetLeftTopX(v int32) *SearchObjectResponseBodyDataFaceList {
	s.LeftTopX = &v
	return s
}

type SearchObjectResponseBodyDataMotorList struct {
	SourceImageUrl *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	DeviceID       *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType     *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	TargetImageUrl *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	RightBottomY   *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY       *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	ShotTime       *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	RightBottomX   *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LeftTopX       *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s SearchObjectResponseBodyDataMotorList) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyDataMotorList) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyDataMotorList) SetSourceImageUrl(v string) *SearchObjectResponseBodyDataMotorList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetDeviceID(v string) *SearchObjectResponseBodyDataMotorList {
	s.DeviceID = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetObjectType(v string) *SearchObjectResponseBodyDataMotorList {
	s.ObjectType = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetTargetImageUrl(v string) *SearchObjectResponseBodyDataMotorList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetRightBottomY(v int32) *SearchObjectResponseBodyDataMotorList {
	s.RightBottomY = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetLeftTopY(v int32) *SearchObjectResponseBodyDataMotorList {
	s.LeftTopY = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetScore(v float32) *SearchObjectResponseBodyDataMotorList {
	s.Score = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetShotTime(v string) *SearchObjectResponseBodyDataMotorList {
	s.ShotTime = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetRightBottomX(v int32) *SearchObjectResponseBodyDataMotorList {
	s.RightBottomX = &v
	return s
}

func (s *SearchObjectResponseBodyDataMotorList) SetLeftTopX(v int32) *SearchObjectResponseBodyDataMotorList {
	s.LeftTopX = &v
	return s
}

type SearchObjectResponseBodyDataNonMotorList struct {
	SourceImageUrl *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	DeviceID       *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType     *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	TargetImageUrl *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	RightBottomY   *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY       *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	ShotTime       *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	RightBottomX   *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LeftTopX       *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
}

func (s SearchObjectResponseBodyDataNonMotorList) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseBodyDataNonMotorList) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetSourceImageUrl(v string) *SearchObjectResponseBodyDataNonMotorList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetDeviceID(v string) *SearchObjectResponseBodyDataNonMotorList {
	s.DeviceID = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetObjectType(v string) *SearchObjectResponseBodyDataNonMotorList {
	s.ObjectType = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetTargetImageUrl(v string) *SearchObjectResponseBodyDataNonMotorList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetRightBottomY(v int32) *SearchObjectResponseBodyDataNonMotorList {
	s.RightBottomY = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetLeftTopY(v int32) *SearchObjectResponseBodyDataNonMotorList {
	s.LeftTopY = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetScore(v float32) *SearchObjectResponseBodyDataNonMotorList {
	s.Score = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetShotTime(v string) *SearchObjectResponseBodyDataNonMotorList {
	s.ShotTime = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetRightBottomX(v int32) *SearchObjectResponseBodyDataNonMotorList {
	s.RightBottomX = &v
	return s
}

func (s *SearchObjectResponseBodyDataNonMotorList) SetLeftTopX(v int32) *SearchObjectResponseBodyDataNonMotorList {
	s.LeftTopX = &v
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

type ListAreaHotSpotMetricsRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonId   *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAreaHotSpotMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAreaHotSpotMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListAreaHotSpotMetricsRequest) SetCorpId(v string) *ListAreaHotSpotMetricsRequest {
	s.CorpId = &v
	return s
}

func (s *ListAreaHotSpotMetricsRequest) SetPersonId(v string) *ListAreaHotSpotMetricsRequest {
	s.PersonId = &v
	return s
}

func (s *ListAreaHotSpotMetricsRequest) SetDeviceId(v string) *ListAreaHotSpotMetricsRequest {
	s.DeviceId = &v
	return s
}

func (s *ListAreaHotSpotMetricsRequest) SetStartTime(v string) *ListAreaHotSpotMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAreaHotSpotMetricsRequest) SetEndTime(v string) *ListAreaHotSpotMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAreaHotSpotMetricsRequest) SetPageNumber(v string) *ListAreaHotSpotMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAreaHotSpotMetricsRequest) SetPageSize(v string) *ListAreaHotSpotMetricsRequest {
	s.PageSize = &v
	return s
}

type ListAreaHotSpotMetricsResponseBody struct {
	TotalCount *string                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListAreaHotSpotMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListAreaHotSpotMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAreaHotSpotMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAreaHotSpotMetricsResponseBody) SetTotalCount(v string) *ListAreaHotSpotMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBody) SetRequestId(v string) *ListAreaHotSpotMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBody) SetMessage(v string) *ListAreaHotSpotMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBody) SetPageSize(v string) *ListAreaHotSpotMetricsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBody) SetPageNumber(v string) *ListAreaHotSpotMetricsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBody) SetData(v []*ListAreaHotSpotMetricsResponseBodyData) *ListAreaHotSpotMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBody) SetCode(v string) *ListAreaHotSpotMetricsResponseBody {
	s.Code = &v
	return s
}

type ListAreaHotSpotMetricsResponseBodyData struct {
	Coordinates  *string `json:"Coordinates,omitempty" xml:"Coordinates,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Times        *string `json:"Times,omitempty" xml:"Times,omitempty"`
	IntervalTime *string `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListAreaHotSpotMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAreaHotSpotMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAreaHotSpotMetricsResponseBodyData) SetCoordinates(v string) *ListAreaHotSpotMetricsResponseBodyData {
	s.Coordinates = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBodyData) SetDeviceId(v string) *ListAreaHotSpotMetricsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBodyData) SetTimes(v string) *ListAreaHotSpotMetricsResponseBodyData {
	s.Times = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBodyData) SetIntervalTime(v string) *ListAreaHotSpotMetricsResponseBodyData {
	s.IntervalTime = &v
	return s
}

func (s *ListAreaHotSpotMetricsResponseBodyData) SetPersonId(v string) *ListAreaHotSpotMetricsResponseBodyData {
	s.PersonId = &v
	return s
}

type ListAreaHotSpotMetricsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAreaHotSpotMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAreaHotSpotMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAreaHotSpotMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListAreaHotSpotMetricsResponse) SetHeaders(v map[string]*string) *ListAreaHotSpotMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListAreaHotSpotMetricsResponse) SetBody(v *ListAreaHotSpotMetricsResponseBody) *ListAreaHotSpotMetricsResponse {
	s.Body = v
	return s
}

type BindDeviceRequest struct {
	CorpId  *string                     `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Devices []*BindDeviceRequestDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
}

func (s BindDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindDeviceRequest) SetCorpId(v string) *BindDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *BindDeviceRequest) SetDevices(v []*BindDeviceRequestDevices) *BindDeviceRequest {
	s.Devices = v
	return s
}

type BindDeviceRequestDevices struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s BindDeviceRequestDevices) String() string {
	return tea.Prettify(s)
}

func (s BindDeviceRequestDevices) GoString() string {
	return s.String()
}

func (s *BindDeviceRequestDevices) SetDeviceId(v string) *BindDeviceRequestDevices {
	s.DeviceId = &v
	return s
}

func (s *BindDeviceRequestDevices) SetCorpId(v string) *BindDeviceRequestDevices {
	s.CorpId = &v
	return s
}

type BindDeviceResponseBody struct {
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*BindDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s BindDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindDeviceResponseBody) SetMessage(v string) *BindDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *BindDeviceResponseBody) SetRequestId(v string) *BindDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindDeviceResponseBody) SetData(v []*BindDeviceResponseBodyData) *BindDeviceResponseBody {
	s.Data = v
	return s
}

func (s *BindDeviceResponseBody) SetCode(v string) *BindDeviceResponseBody {
	s.Code = &v
	return s
}

type BindDeviceResponseBodyData struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s BindDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindDeviceResponseBodyData) SetDeviceId(v string) *BindDeviceResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *BindDeviceResponseBodyData) SetSuccess(v bool) *BindDeviceResponseBodyData {
	s.Success = &v
	return s
}

func (s *BindDeviceResponseBodyData) SetCode(v string) *BindDeviceResponseBodyData {
	s.Code = &v
	return s
}

func (s *BindDeviceResponseBodyData) SetMessage(v string) *BindDeviceResponseBodyData {
	s.Message = &v
	return s
}

type BindDeviceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindDeviceResponse) SetHeaders(v map[string]*string) *BindDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindDeviceResponse) SetBody(v *BindDeviceResponseBody) *BindDeviceResponse {
	s.Body = v
	return s
}

type GetCdrsMonitorResultRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	MinRecordId     *string `json:"MinRecordId,omitempty" xml:"MinRecordId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
}

func (s GetCdrsMonitorResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorResultRequest) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorResultRequest) SetCorpId(v string) *GetCdrsMonitorResultRequest {
	s.CorpId = &v
	return s
}

func (s *GetCdrsMonitorResultRequest) SetTaskId(v string) *GetCdrsMonitorResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetCdrsMonitorResultRequest) SetMinRecordId(v string) *GetCdrsMonitorResultRequest {
	s.MinRecordId = &v
	return s
}

func (s *GetCdrsMonitorResultRequest) SetStartTime(v int64) *GetCdrsMonitorResultRequest {
	s.StartTime = &v
	return s
}

func (s *GetCdrsMonitorResultRequest) SetEndTime(v int64) *GetCdrsMonitorResultRequest {
	s.EndTime = &v
	return s
}

func (s *GetCdrsMonitorResultRequest) SetAlgorithmVendor(v string) *GetCdrsMonitorResultRequest {
	s.AlgorithmVendor = &v
	return s
}

type GetCdrsMonitorResultResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetCdrsMonitorResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetCdrsMonitorResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorResultResponseBody) SetMessage(v string) *GetCdrsMonitorResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBody) SetRequestId(v string) *GetCdrsMonitorResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBody) SetData(v *GetCdrsMonitorResultResponseBodyData) *GetCdrsMonitorResultResponseBody {
	s.Data = v
	return s
}

func (s *GetCdrsMonitorResultResponseBody) SetCode(v string) *GetCdrsMonitorResultResponseBody {
	s.Code = &v
	return s
}

type GetCdrsMonitorResultResponseBodyData struct {
	MaxId   *string                                        `json:"MaxId,omitempty" xml:"MaxId,omitempty"`
	Records []*GetCdrsMonitorResultResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s GetCdrsMonitorResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorResultResponseBodyData) SetMaxId(v string) *GetCdrsMonitorResultResponseBodyData {
	s.MaxId = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyData) SetRecords(v []*GetCdrsMonitorResultResponseBodyDataRecords) *GetCdrsMonitorResultResponseBodyData {
	s.Records = v
	return s
}

type GetCdrsMonitorResultResponseBodyDataRecords struct {
	PicUrl        *string                                                `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	RightBottomY  *string                                                `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Score         *string                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	MonitorPicUrl *string                                                `json:"MonitorPicUrl,omitempty" xml:"MonitorPicUrl,omitempty"`
	RightBottomX  *string                                                `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	ExtendInfo    *GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty" type:"Struct"`
	GbId          *string                                                `json:"GbId,omitempty" xml:"GbId,omitempty"`
	LeftUpY       *string                                                `json:"LeftUpY,omitempty" xml:"LeftUpY,omitempty"`
	LeftUpX       *string                                                `json:"LeftUpX,omitempty" xml:"LeftUpX,omitempty"`
	ShotTime      *string                                                `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	TaskId        *string                                                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TargetPicUrl  *string                                                `json:"TargetPicUrl,omitempty" xml:"TargetPicUrl,omitempty"`
}

func (s GetCdrsMonitorResultResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorResultResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetPicUrl(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.PicUrl = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetRightBottomY(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetScore(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetMonitorPicUrl(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.MonitorPicUrl = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetRightBottomX(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetExtendInfo(v *GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.ExtendInfo = v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetGbId(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.GbId = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetLeftUpY(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.LeftUpY = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetLeftUpX(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.LeftUpX = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetShotTime(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.ShotTime = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetTaskId(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.TaskId = &v
	return s
}

func (s *GetCdrsMonitorResultResponseBodyDataRecords) SetTargetPicUrl(v string) *GetCdrsMonitorResultResponseBodyDataRecords {
	s.TargetPicUrl = &v
	return s
}

type GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo struct {
	PlateNo *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
}

func (s GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo) SetPlateNo(v string) *GetCdrsMonitorResultResponseBodyDataRecordsExtendInfo {
	s.PlateNo = &v
	return s
}

type GetCdrsMonitorResultResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCdrsMonitorResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCdrsMonitorResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorResultResponse) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorResultResponse) SetHeaders(v map[string]*string) *GetCdrsMonitorResultResponse {
	s.Headers = v
	return s
}

func (s *GetCdrsMonitorResultResponse) SetBody(v *GetCdrsMonitorResultResponseBody) *GetCdrsMonitorResultResponse {
	s.Body = v
	return s
}

type ListVehicleDetailsRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PlateId    *string `json:"PlateId,omitempty" xml:"PlateId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVehicleDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListVehicleDetailsRequest) SetCorpId(v string) *ListVehicleDetailsRequest {
	s.CorpId = &v
	return s
}

func (s *ListVehicleDetailsRequest) SetPlateId(v string) *ListVehicleDetailsRequest {
	s.PlateId = &v
	return s
}

func (s *ListVehicleDetailsRequest) SetStartTime(v string) *ListVehicleDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListVehicleDetailsRequest) SetEndTime(v string) *ListVehicleDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListVehicleDetailsRequest) SetPageNumber(v string) *ListVehicleDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleDetailsRequest) SetPageSize(v string) *ListVehicleDetailsRequest {
	s.PageSize = &v
	return s
}

type ListVehicleDetailsResponseBody struct {
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListVehicleDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListVehicleDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVehicleDetailsResponseBody) SetTotalCount(v int64) *ListVehicleDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVehicleDetailsResponseBody) SetRequestId(v string) *ListVehicleDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVehicleDetailsResponseBody) SetMessage(v string) *ListVehicleDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVehicleDetailsResponseBody) SetPageSize(v int64) *ListVehicleDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVehicleDetailsResponseBody) SetPageNumber(v int64) *ListVehicleDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleDetailsResponseBody) SetData(v []*ListVehicleDetailsResponseBodyData) *ListVehicleDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListVehicleDetailsResponseBody) SetCode(v string) *ListVehicleDetailsResponseBody {
	s.Code = &v
	return s
}

type ListVehicleDetailsResponseBodyData struct {
	VehicleApplication     *string `json:"VehicleApplication,omitempty" xml:"VehicleApplication,omitempty"`
	SourceUrl              *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	Gender                 *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	VehicleColor           *string `json:"VehicleColor,omitempty" xml:"VehicleColor,omitempty"`
	VehicleId              *string `json:"VehicleId,omitempty" xml:"VehicleId,omitempty"`
	SourceImageStoragePath *string `json:"SourceImageStoragePath,omitempty" xml:"SourceImageStoragePath,omitempty"`
	PersonType             *string `json:"PersonType,omitempty" xml:"PersonType,omitempty"`
	PopularPoi             *string `json:"PopularPoi,omitempty" xml:"PopularPoi,omitempty"`
	PopularAddress         *string `json:"PopularAddress,omitempty" xml:"PopularAddress,omitempty"`
	PlateId                *string `json:"PlateId,omitempty" xml:"PlateId,omitempty"`
	TargetUrl              *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	VehicleClass           *string `json:"VehicleClass,omitempty" xml:"VehicleClass,omitempty"`
	PrefOutTime            *string `json:"PrefOutTime,omitempty" xml:"PrefOutTime,omitempty"`
	PersonId               *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	TargetImageStoragePath *string `json:"TargetImageStoragePath,omitempty" xml:"TargetImageStoragePath,omitempty"`
}

func (s ListVehicleDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVehicleDetailsResponseBodyData) SetVehicleApplication(v string) *ListVehicleDetailsResponseBodyData {
	s.VehicleApplication = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetSourceUrl(v string) *ListVehicleDetailsResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetGender(v string) *ListVehicleDetailsResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetVehicleColor(v string) *ListVehicleDetailsResponseBodyData {
	s.VehicleColor = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetVehicleId(v string) *ListVehicleDetailsResponseBodyData {
	s.VehicleId = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetSourceImageStoragePath(v string) *ListVehicleDetailsResponseBodyData {
	s.SourceImageStoragePath = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetPersonType(v string) *ListVehicleDetailsResponseBodyData {
	s.PersonType = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetPopularPoi(v string) *ListVehicleDetailsResponseBodyData {
	s.PopularPoi = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetPopularAddress(v string) *ListVehicleDetailsResponseBodyData {
	s.PopularAddress = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetPlateId(v string) *ListVehicleDetailsResponseBodyData {
	s.PlateId = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetTargetUrl(v string) *ListVehicleDetailsResponseBodyData {
	s.TargetUrl = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetVehicleClass(v string) *ListVehicleDetailsResponseBodyData {
	s.VehicleClass = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetPrefOutTime(v string) *ListVehicleDetailsResponseBodyData {
	s.PrefOutTime = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetPersonId(v string) *ListVehicleDetailsResponseBodyData {
	s.PersonId = &v
	return s
}

func (s *ListVehicleDetailsResponseBodyData) SetTargetImageStoragePath(v string) *ListVehicleDetailsResponseBodyData {
	s.TargetImageStoragePath = &v
	return s
}

type ListVehicleDetailsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVehicleDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVehicleDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListVehicleDetailsResponse) SetHeaders(v map[string]*string) *ListVehicleDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListVehicleDetailsResponse) SetBody(v *ListVehicleDetailsResponseBody) *ListVehicleDetailsResponse {
	s.Body = v
	return s
}

type GetCdrsMonitorListRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetCdrsMonitorListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorListRequest) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorListRequest) SetCorpId(v string) *GetCdrsMonitorListRequest {
	s.CorpId = &v
	return s
}

func (s *GetCdrsMonitorListRequest) SetPageNo(v int32) *GetCdrsMonitorListRequest {
	s.PageNo = &v
	return s
}

func (s *GetCdrsMonitorListRequest) SetPageSize(v int32) *GetCdrsMonitorListRequest {
	s.PageSize = &v
	return s
}

type GetCdrsMonitorListResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetCdrsMonitorListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetCdrsMonitorListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorListResponseBody) SetMessage(v string) *GetCdrsMonitorListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCdrsMonitorListResponseBody) SetRequestId(v string) *GetCdrsMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCdrsMonitorListResponseBody) SetData(v *GetCdrsMonitorListResponseBodyData) *GetCdrsMonitorListResponseBody {
	s.Data = v
	return s
}

func (s *GetCdrsMonitorListResponseBody) SetCode(v string) *GetCdrsMonitorListResponseBody {
	s.Code = &v
	return s
}

type GetCdrsMonitorListResponseBodyData struct {
	Records    []*GetCdrsMonitorListResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNo     *int32                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TotalPage  *int32                                       `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCdrsMonitorListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorListResponseBodyData) SetRecords(v []*GetCdrsMonitorListResponseBodyDataRecords) *GetCdrsMonitorListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetCdrsMonitorListResponseBodyData) SetPageNo(v int32) *GetCdrsMonitorListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyData) SetTotalPage(v int32) *GetCdrsMonitorListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyData) SetPageSize(v int32) *GetCdrsMonitorListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyData) SetTotalCount(v int32) *GetCdrsMonitorListResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetCdrsMonitorListResponseBodyDataRecords struct {
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

func (s GetCdrsMonitorListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetStatus(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetRuleExpression(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.RuleExpression = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetImageMatch(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.ImageMatch = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetMonitorType(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.MonitorType = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetCreateDate(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.CreateDate = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetRuleName(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.RuleName = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetNotifierType(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.NotifierType = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetDescription(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetExpression(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.Expression = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetNotifierExtra(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.NotifierExtra = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetAttributes(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.Attributes = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetDeviceList(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.DeviceList = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetTaskId(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.TaskId = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetModifiedDate(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.ModifiedDate = &v
	return s
}

func (s *GetCdrsMonitorListResponseBodyDataRecords) SetAlgorithmVendor(v string) *GetCdrsMonitorListResponseBodyDataRecords {
	s.AlgorithmVendor = &v
	return s
}

type GetCdrsMonitorListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCdrsMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCdrsMonitorListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCdrsMonitorListResponse) GoString() string {
	return s.String()
}

func (s *GetCdrsMonitorListResponse) SetHeaders(v map[string]*string) *GetCdrsMonitorListResponse {
	s.Headers = v
	return s
}

func (s *GetCdrsMonitorListResponse) SetBody(v *GetCdrsMonitorListResponseBody) *GetCdrsMonitorListResponse {
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
	PicExtendList        *string `json:"PicExtendList,omitempty" xml:"PicExtendList,omitempty"`
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

func (s *UpdateMonitorRequest) SetPicExtendList(v string) *UpdateMonitorRequest {
	s.PicExtendList = &v
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

type ListDataStatisticsRequest struct {
	BackCategory *string `json:"BackCategory,omitempty" xml:"BackCategory,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ListDataStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsRequest) SetBackCategory(v string) *ListDataStatisticsRequest {
	s.BackCategory = &v
	return s
}

func (s *ListDataStatisticsRequest) SetSchema(v string) *ListDataStatisticsRequest {
	s.Schema = &v
	return s
}

type ListDataStatisticsResponseBody struct {
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListDataStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDataStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsResponseBody) SetTotalCount(v int64) *ListDataStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataStatisticsResponseBody) SetRequestId(v string) *ListDataStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataStatisticsResponseBody) SetMessage(v string) *ListDataStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataStatisticsResponseBody) SetPageSize(v int64) *ListDataStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataStatisticsResponseBody) SetPageNumber(v int64) *ListDataStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataStatisticsResponseBody) SetData(v []*ListDataStatisticsResponseBodyData) *ListDataStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataStatisticsResponseBody) SetCode(v string) *ListDataStatisticsResponseBody {
	s.Code = &v
	return s
}

type ListDataStatisticsResponseBodyData struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ListDataStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsResponseBodyData) SetCorpId(v string) *ListDataStatisticsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListDataStatisticsResponseBodyData) SetNumber(v string) *ListDataStatisticsResponseBodyData {
	s.Number = &v
	return s
}

type ListDataStatisticsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsResponse) SetHeaders(v map[string]*string) *ListDataStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListDataStatisticsResponse) SetBody(v *ListDataStatisticsResponseBody) *ListDataStatisticsResponse {
	s.Body = v
	return s
}

type UnbindDeviceRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	DeviceIds *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
}

func (s UnbindDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequest) SetCorpId(v string) *UnbindDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *UnbindDeviceRequest) SetDeviceIds(v string) *UnbindDeviceRequest {
	s.DeviceIds = &v
	return s
}

type UnbindDeviceResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*UnbindDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnbindDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponseBody) SetMessage(v string) *UnbindDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetRequestId(v string) *UnbindDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetData(v []*UnbindDeviceResponseBodyData) *UnbindDeviceResponseBody {
	s.Data = v
	return s
}

func (s *UnbindDeviceResponseBody) SetCode(v string) *UnbindDeviceResponseBody {
	s.Code = &v
	return s
}

type UnbindDeviceResponseBodyData struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UnbindDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponseBodyData) SetDeviceId(v string) *UnbindDeviceResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *UnbindDeviceResponseBodyData) SetSuccess(v bool) *UnbindDeviceResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnbindDeviceResponseBodyData) SetCode(v string) *UnbindDeviceResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnbindDeviceResponseBodyData) SetMessage(v string) *UnbindDeviceResponseBodyData {
	s.Message = &v
	return s
}

type UnbindDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponse) SetHeaders(v map[string]*string) *UnbindDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindDeviceResponse) SetBody(v *UnbindDeviceResponseBody) *UnbindDeviceResponse {
	s.Body = v
	return s
}

type ListPersonDetailsRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonId   *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Schema     *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ListPersonDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListPersonDetailsRequest) SetCorpId(v string) *ListPersonDetailsRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonDetailsRequest) SetPersonId(v string) *ListPersonDetailsRequest {
	s.PersonId = &v
	return s
}

func (s *ListPersonDetailsRequest) SetStartTime(v string) *ListPersonDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonDetailsRequest) SetEndTime(v string) *ListPersonDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonDetailsRequest) SetPageNumber(v int64) *ListPersonDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonDetailsRequest) SetPageSize(v int64) *ListPersonDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonDetailsRequest) SetSchema(v string) *ListPersonDetailsRequest {
	s.Schema = &v
	return s
}

type ListPersonDetailsResponseBody struct {
	TotalCount *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonDetailsResponseBody) SetTotalCount(v int64) *ListPersonDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonDetailsResponseBody) SetRequestId(v string) *ListPersonDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonDetailsResponseBody) SetMessage(v string) *ListPersonDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonDetailsResponseBody) SetPageSize(v int64) *ListPersonDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonDetailsResponseBody) SetPageNumber(v int64) *ListPersonDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonDetailsResponseBody) SetData(v []*ListPersonDetailsResponseBodyData) *ListPersonDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonDetailsResponseBody) SetCode(v string) *ListPersonDetailsResponseBody {
	s.Code = &v
	return s
}

type ListPersonDetailsResponseBodyData struct {
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Gender          *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	BodyTargetImage *string `json:"BodyTargetImage,omitempty" xml:"BodyTargetImage,omitempty"`
	PreferredColor  *string `json:"PreferredColor,omitempty" xml:"PreferredColor,omitempty"`
	HotSpotAddress  *string `json:"HotSpotAddress,omitempty" xml:"HotSpotAddress,omitempty"`
	Age             *string `json:"Age,omitempty" xml:"Age,omitempty"`
	PersonType      *string `json:"PersonType,omitempty" xml:"PersonType,omitempty"`
	Transportation  *string `json:"Transportation,omitempty" xml:"Transportation,omitempty"`
	Profession      *string `json:"Profession,omitempty" xml:"Profession,omitempty"`
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	FaceSourceImage *string `json:"FaceSourceImage,omitempty" xml:"FaceSourceImage,omitempty"`
	FaceTargetImage *string `json:"FaceTargetImage,omitempty" xml:"FaceTargetImage,omitempty"`
	PrefOutTime     *string `json:"PrefOutTime,omitempty" xml:"PrefOutTime,omitempty"`
	BodySourceImage *string `json:"BodySourceImage,omitempty" xml:"BodySourceImage,omitempty"`
	PersonId        *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListPersonDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonDetailsResponseBodyData) SetUpdateTime(v string) *ListPersonDetailsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetGender(v string) *ListPersonDetailsResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetBodyTargetImage(v string) *ListPersonDetailsResponseBodyData {
	s.BodyTargetImage = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetPreferredColor(v string) *ListPersonDetailsResponseBodyData {
	s.PreferredColor = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetHotSpotAddress(v string) *ListPersonDetailsResponseBodyData {
	s.HotSpotAddress = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetAge(v string) *ListPersonDetailsResponseBodyData {
	s.Age = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetPersonType(v string) *ListPersonDetailsResponseBodyData {
	s.PersonType = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetTransportation(v string) *ListPersonDetailsResponseBodyData {
	s.Transportation = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetProfession(v string) *ListPersonDetailsResponseBodyData {
	s.Profession = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetAddress(v string) *ListPersonDetailsResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetFaceSourceImage(v string) *ListPersonDetailsResponseBodyData {
	s.FaceSourceImage = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetFaceTargetImage(v string) *ListPersonDetailsResponseBodyData {
	s.FaceTargetImage = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetPrefOutTime(v string) *ListPersonDetailsResponseBodyData {
	s.PrefOutTime = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetBodySourceImage(v string) *ListPersonDetailsResponseBodyData {
	s.BodySourceImage = &v
	return s
}

func (s *ListPersonDetailsResponseBodyData) SetPersonId(v string) *ListPersonDetailsResponseBodyData {
	s.PersonId = &v
	return s
}

type ListPersonDetailsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListPersonDetailsResponse) SetHeaders(v map[string]*string) *ListPersonDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListPersonDetailsResponse) SetBody(v *ListPersonDetailsResponseBody) *ListPersonDetailsResponse {
	s.Body = v
	return s
}

type ListVehicleTagDistributeRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode   *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListVehicleTagDistributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTagDistributeRequest) GoString() string {
	return s.String()
}

func (s *ListVehicleTagDistributeRequest) SetCorpId(v string) *ListVehicleTagDistributeRequest {
	s.CorpId = &v
	return s
}

func (s *ListVehicleTagDistributeRequest) SetTagCode(v string) *ListVehicleTagDistributeRequest {
	s.TagCode = &v
	return s
}

func (s *ListVehicleTagDistributeRequest) SetStartTime(v string) *ListVehicleTagDistributeRequest {
	s.StartTime = &v
	return s
}

func (s *ListVehicleTagDistributeRequest) SetEndTime(v string) *ListVehicleTagDistributeRequest {
	s.EndTime = &v
	return s
}

type ListVehicleTagDistributeResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListVehicleTagDistributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListVehicleTagDistributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTagDistributeResponseBody) GoString() string {
	return s.String()
}

func (s *ListVehicleTagDistributeResponseBody) SetTotalCount(v int64) *ListVehicleTagDistributeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBody) SetRequestId(v string) *ListVehicleTagDistributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBody) SetMessage(v string) *ListVehicleTagDistributeResponseBody {
	s.Message = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBody) SetPageSize(v int64) *ListVehicleTagDistributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBody) SetPageNumber(v int64) *ListVehicleTagDistributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBody) SetData(v []*ListVehicleTagDistributeResponseBodyData) *ListVehicleTagDistributeResponseBody {
	s.Data = v
	return s
}

func (s *ListVehicleTagDistributeResponseBody) SetCode(v string) *ListVehicleTagDistributeResponseBody {
	s.Code = &v
	return s
}

type ListVehicleTagDistributeResponseBodyData struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVehicleTagDistributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTagDistributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVehicleTagDistributeResponseBodyData) SetValue(v string) *ListVehicleTagDistributeResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBodyData) SetCorpId(v string) *ListVehicleTagDistributeResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListVehicleTagDistributeResponseBodyData) SetTagValue(v string) *ListVehicleTagDistributeResponseBodyData {
	s.TagValue = &v
	return s
}

type ListVehicleTagDistributeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVehicleTagDistributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVehicleTagDistributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTagDistributeResponse) GoString() string {
	return s.String()
}

func (s *ListVehicleTagDistributeResponse) SetHeaders(v map[string]*string) *ListVehicleTagDistributeResponse {
	s.Headers = v
	return s
}

func (s *ListVehicleTagDistributeResponse) SetBody(v *ListVehicleTagDistributeResponseBody) *ListVehicleTagDistributeResponse {
	s.Body = v
	return s
}

type ListDevicePersonStatisticsRequest struct {
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListDevicePersonStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListDevicePersonStatisticsRequest) SetDataSourceId(v string) *ListDevicePersonStatisticsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDevicePersonStatisticsRequest) SetStatisticsType(v string) *ListDevicePersonStatisticsRequest {
	s.StatisticsType = &v
	return s
}

func (s *ListDevicePersonStatisticsRequest) SetStartTime(v string) *ListDevicePersonStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDevicePersonStatisticsRequest) SetEndTime(v string) *ListDevicePersonStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDevicePersonStatisticsRequest) SetCorpId(v string) *ListDevicePersonStatisticsRequest {
	s.CorpId = &v
	return s
}

type ListDevicePersonStatisticsResponseBody struct {
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       []*ListDevicePersonStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDevicePersonStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicePersonStatisticsResponseBody) SetTotalCount(v int64) *ListDevicePersonStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDevicePersonStatisticsResponseBody) SetMessage(v string) *ListDevicePersonStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicePersonStatisticsResponseBody) SetRequestId(v string) *ListDevicePersonStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicePersonStatisticsResponseBody) SetData(v []*ListDevicePersonStatisticsResponseBodyData) *ListDevicePersonStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicePersonStatisticsResponseBody) SetCode(v string) *ListDevicePersonStatisticsResponseBody {
	s.Code = &v
	return s
}

type ListDevicePersonStatisticsResponseBodyData struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	Number       *string `json:"Number,omitempty" xml:"Number,omitempty"`
	ShotTime     *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
}

func (s ListDevicePersonStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicePersonStatisticsResponseBodyData) SetDataSourceId(v string) *ListDevicePersonStatisticsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListDevicePersonStatisticsResponseBodyData) SetNumber(v string) *ListDevicePersonStatisticsResponseBodyData {
	s.Number = &v
	return s
}

func (s *ListDevicePersonStatisticsResponseBodyData) SetShotTime(v string) *ListDevicePersonStatisticsResponseBodyData {
	s.ShotTime = &v
	return s
}

type ListDevicePersonStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicePersonStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicePersonStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListDevicePersonStatisticsResponse) SetHeaders(v map[string]*string) *ListDevicePersonStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListDevicePersonStatisticsResponse) SetBody(v *ListDevicePersonStatisticsResponseBody) *ListDevicePersonStatisticsResponse {
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

type PaginateDeviceRequest struct {
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CountTotalNum *bool   `json:"CountTotalNum,omitempty" xml:"CountTotalNum,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s PaginateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s PaginateDeviceRequest) GoString() string {
	return s.String()
}

func (s *PaginateDeviceRequest) SetPageNumber(v int64) *PaginateDeviceRequest {
	s.PageNumber = &v
	return s
}

func (s *PaginateDeviceRequest) SetPageSize(v int64) *PaginateDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *PaginateDeviceRequest) SetCountTotalNum(v bool) *PaginateDeviceRequest {
	s.CountTotalNum = &v
	return s
}

func (s *PaginateDeviceRequest) SetCorpId(v string) *PaginateDeviceRequest {
	s.CorpId = &v
	return s
}

type PaginateDeviceResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *PaginateDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s PaginateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PaginateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *PaginateDeviceResponseBody) SetMessage(v string) *PaginateDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *PaginateDeviceResponseBody) SetRequestId(v string) *PaginateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PaginateDeviceResponseBody) SetData(v *PaginateDeviceResponseBodyData) *PaginateDeviceResponseBody {
	s.Data = v
	return s
}

func (s *PaginateDeviceResponseBody) SetCode(v string) *PaginateDeviceResponseBody {
	s.Code = &v
	return s
}

type PaginateDeviceResponseBodyData struct {
	Records    []*PaginateDeviceResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PaginateDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PaginateDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PaginateDeviceResponseBodyData) SetRecords(v []*PaginateDeviceResponseBodyDataRecords) *PaginateDeviceResponseBodyData {
	s.Records = v
	return s
}

func (s *PaginateDeviceResponseBodyData) SetPageNumber(v int32) *PaginateDeviceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *PaginateDeviceResponseBodyData) SetPageSize(v int32) *PaginateDeviceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PaginateDeviceResponseBodyData) SetTotalCount(v int32) *PaginateDeviceResponseBodyData {
	s.TotalCount = &v
	return s
}

type PaginateDeviceResponseBodyDataRecords struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s PaginateDeviceResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s PaginateDeviceResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *PaginateDeviceResponseBodyDataRecords) SetDeviceId(v string) *PaginateDeviceResponseBodyDataRecords {
	s.DeviceId = &v
	return s
}

type PaginateDeviceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PaginateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PaginateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s PaginateDeviceResponse) GoString() string {
	return s.String()
}

func (s *PaginateDeviceResponse) SetHeaders(v map[string]*string) *PaginateDeviceResponse {
	s.Headers = v
	return s
}

func (s *PaginateDeviceResponse) SetBody(v *PaginateDeviceResponseBody) *PaginateDeviceResponse {
	s.Body = v
	return s
}

type StopCdrsMonitorRequest struct {
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s StopCdrsMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCdrsMonitorRequest) GoString() string {
	return s.String()
}

func (s *StopCdrsMonitorRequest) SetTaskId(v string) *StopCdrsMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *StopCdrsMonitorRequest) SetAlgorithmVendor(v string) *StopCdrsMonitorRequest {
	s.AlgorithmVendor = &v
	return s
}

func (s *StopCdrsMonitorRequest) SetCorpId(v string) *StopCdrsMonitorRequest {
	s.CorpId = &v
	return s
}

type StopCdrsMonitorResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StopCdrsMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCdrsMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *StopCdrsMonitorResponseBody) SetMessage(v string) *StopCdrsMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *StopCdrsMonitorResponseBody) SetRequestId(v string) *StopCdrsMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCdrsMonitorResponseBody) SetData(v string) *StopCdrsMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *StopCdrsMonitorResponseBody) SetCode(v string) *StopCdrsMonitorResponseBody {
	s.Code = &v
	return s
}

type StopCdrsMonitorResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCdrsMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCdrsMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCdrsMonitorResponse) GoString() string {
	return s.String()
}

func (s *StopCdrsMonitorResponse) SetHeaders(v map[string]*string) *StopCdrsMonitorResponse {
	s.Headers = v
	return s
}

func (s *StopCdrsMonitorResponse) SetBody(v *StopCdrsMonitorResponseBody) *StopCdrsMonitorResponse {
	s.Body = v
	return s
}

type RecallTrajectoryByCoordinateTimeRequest struct {
	StartTime        *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CorpId           *string                                             `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	OutputIdCount    *int32                                              `json:"OutputIdCount,omitempty" xml:"OutputIdCount,omitempty"`
	PointList        []*RecallTrajectoryByCoordinateTimeRequestPointList `json:"PointList,omitempty" xml:"PointList,omitempty" type:"Repeated"`
	OutputIdTypeList []*string                                           `json:"OutputIdTypeList,omitempty" xml:"OutputIdTypeList,omitempty" type:"Repeated"`
}

func (s RecallTrajectoryByCoordinateTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallTrajectoryByCoordinateTimeRequest) GoString() string {
	return s.String()
}

func (s *RecallTrajectoryByCoordinateTimeRequest) SetStartTime(v string) *RecallTrajectoryByCoordinateTimeRequest {
	s.StartTime = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequest) SetEndTime(v string) *RecallTrajectoryByCoordinateTimeRequest {
	s.EndTime = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequest) SetCorpId(v string) *RecallTrajectoryByCoordinateTimeRequest {
	s.CorpId = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequest) SetOutputIdCount(v int32) *RecallTrajectoryByCoordinateTimeRequest {
	s.OutputIdCount = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequest) SetPointList(v []*RecallTrajectoryByCoordinateTimeRequestPointList) *RecallTrajectoryByCoordinateTimeRequest {
	s.PointList = v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequest) SetOutputIdTypeList(v []*string) *RecallTrajectoryByCoordinateTimeRequest {
	s.OutputIdTypeList = v
	return s
}

type RecallTrajectoryByCoordinateTimeRequestPointList struct {
	EndTime       *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime     *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Longitude     *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude      *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	DeltaDistance *float32 `json:"DeltaDistance,omitempty" xml:"DeltaDistance,omitempty"`
}

func (s RecallTrajectoryByCoordinateTimeRequestPointList) String() string {
	return tea.Prettify(s)
}

func (s RecallTrajectoryByCoordinateTimeRequestPointList) GoString() string {
	return s.String()
}

func (s *RecallTrajectoryByCoordinateTimeRequestPointList) SetEndTime(v string) *RecallTrajectoryByCoordinateTimeRequestPointList {
	s.EndTime = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequestPointList) SetStartTime(v string) *RecallTrajectoryByCoordinateTimeRequestPointList {
	s.StartTime = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequestPointList) SetLongitude(v float32) *RecallTrajectoryByCoordinateTimeRequestPointList {
	s.Longitude = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequestPointList) SetLatitude(v float32) *RecallTrajectoryByCoordinateTimeRequestPointList {
	s.Latitude = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeRequestPointList) SetDeltaDistance(v float32) *RecallTrajectoryByCoordinateTimeRequestPointList {
	s.DeltaDistance = &v
	return s
}

type RecallTrajectoryByCoordinateTimeResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecallTrajectoryByCoordinateTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallTrajectoryByCoordinateTimeResponseBody) GoString() string {
	return s.String()
}

func (s *RecallTrajectoryByCoordinateTimeResponseBody) SetMessage(v string) *RecallTrajectoryByCoordinateTimeResponseBody {
	s.Message = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeResponseBody) SetRequestId(v string) *RecallTrajectoryByCoordinateTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeResponseBody) SetData(v string) *RecallTrajectoryByCoordinateTimeResponseBody {
	s.Data = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeResponseBody) SetCode(v int32) *RecallTrajectoryByCoordinateTimeResponseBody {
	s.Code = &v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeResponseBody) SetSuccess(v bool) *RecallTrajectoryByCoordinateTimeResponseBody {
	s.Success = &v
	return s
}

type RecallTrajectoryByCoordinateTimeResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecallTrajectoryByCoordinateTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecallTrajectoryByCoordinateTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallTrajectoryByCoordinateTimeResponse) GoString() string {
	return s.String()
}

func (s *RecallTrajectoryByCoordinateTimeResponse) SetHeaders(v map[string]*string) *RecallTrajectoryByCoordinateTimeResponse {
	s.Headers = v
	return s
}

func (s *RecallTrajectoryByCoordinateTimeResponse) SetBody(v *RecallTrajectoryByCoordinateTimeResponseBody) *RecallTrajectoryByCoordinateTimeResponse {
	s.Body = v
	return s
}

type ListCityMapPersonFlowRequest struct {
	OriginDataSourceIdList map[string]interface{} `json:"OriginDataSourceIdList,omitempty" xml:"OriginDataSourceIdList,omitempty"`
	PageNumber             *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TargetDataSourceIdList map[string]interface{} `json:"TargetDataSourceIdList,omitempty" xml:"TargetDataSourceIdList,omitempty"`
	EndTime                *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime              *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Range                  *string                `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s ListCityMapPersonFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapPersonFlowRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapPersonFlowRequest) SetOriginDataSourceIdList(v map[string]interface{}) *ListCityMapPersonFlowRequest {
	s.OriginDataSourceIdList = v
	return s
}

func (s *ListCityMapPersonFlowRequest) SetPageNumber(v int64) *ListCityMapPersonFlowRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapPersonFlowRequest) SetPageSize(v int64) *ListCityMapPersonFlowRequest {
	s.PageSize = &v
	return s
}

func (s *ListCityMapPersonFlowRequest) SetTargetDataSourceIdList(v map[string]interface{}) *ListCityMapPersonFlowRequest {
	s.TargetDataSourceIdList = v
	return s
}

func (s *ListCityMapPersonFlowRequest) SetEndTime(v string) *ListCityMapPersonFlowRequest {
	s.EndTime = &v
	return s
}

func (s *ListCityMapPersonFlowRequest) SetStartTime(v string) *ListCityMapPersonFlowRequest {
	s.StartTime = &v
	return s
}

func (s *ListCityMapPersonFlowRequest) SetRange(v string) *ListCityMapPersonFlowRequest {
	s.Range = &v
	return s
}

type ListCityMapPersonFlowShrinkRequest struct {
	OriginDataSourceIdListShrink *string `json:"OriginDataSourceIdList,omitempty" xml:"OriginDataSourceIdList,omitempty"`
	PageNumber                   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TargetDataSourceIdListShrink *string `json:"TargetDataSourceIdList,omitempty" xml:"TargetDataSourceIdList,omitempty"`
	EndTime                      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime                    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Range                        *string `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s ListCityMapPersonFlowShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapPersonFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapPersonFlowShrinkRequest) SetOriginDataSourceIdListShrink(v string) *ListCityMapPersonFlowShrinkRequest {
	s.OriginDataSourceIdListShrink = &v
	return s
}

func (s *ListCityMapPersonFlowShrinkRequest) SetPageNumber(v int64) *ListCityMapPersonFlowShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapPersonFlowShrinkRequest) SetPageSize(v int64) *ListCityMapPersonFlowShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCityMapPersonFlowShrinkRequest) SetTargetDataSourceIdListShrink(v string) *ListCityMapPersonFlowShrinkRequest {
	s.TargetDataSourceIdListShrink = &v
	return s
}

func (s *ListCityMapPersonFlowShrinkRequest) SetEndTime(v string) *ListCityMapPersonFlowShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListCityMapPersonFlowShrinkRequest) SetStartTime(v string) *ListCityMapPersonFlowShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListCityMapPersonFlowShrinkRequest) SetRange(v string) *ListCityMapPersonFlowShrinkRequest {
	s.Range = &v
	return s
}

type ListCityMapPersonFlowResponseBody struct {
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCityMapPersonFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCityMapPersonFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapPersonFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListCityMapPersonFlowResponseBody) SetTotalCount(v int64) *ListCityMapPersonFlowResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBody) SetRequestId(v string) *ListCityMapPersonFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBody) SetMessage(v string) *ListCityMapPersonFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBody) SetPageSize(v int64) *ListCityMapPersonFlowResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBody) SetPageNumber(v int64) *ListCityMapPersonFlowResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBody) SetData(v []*ListCityMapPersonFlowResponseBodyData) *ListCityMapPersonFlowResponseBody {
	s.Data = v
	return s
}

func (s *ListCityMapPersonFlowResponseBody) SetCode(v string) *ListCityMapPersonFlowResponseBody {
	s.Code = &v
	return s
}

type ListCityMapPersonFlowResponseBodyData struct {
	FlowNumber         *string `json:"FlowNumber,omitempty" xml:"FlowNumber,omitempty"`
	OriginDataSourceId *string `json:"OriginDataSourceId,omitempty" xml:"OriginDataSourceId,omitempty"`
	TargetDataSourceId *string `json:"TargetDataSourceId,omitempty" xml:"TargetDataSourceId,omitempty"`
	FlowDirection      *string `json:"FlowDirection,omitempty" xml:"FlowDirection,omitempty"`
}

func (s ListCityMapPersonFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapPersonFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCityMapPersonFlowResponseBodyData) SetFlowNumber(v string) *ListCityMapPersonFlowResponseBodyData {
	s.FlowNumber = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBodyData) SetOriginDataSourceId(v string) *ListCityMapPersonFlowResponseBodyData {
	s.OriginDataSourceId = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBodyData) SetTargetDataSourceId(v string) *ListCityMapPersonFlowResponseBodyData {
	s.TargetDataSourceId = &v
	return s
}

func (s *ListCityMapPersonFlowResponseBodyData) SetFlowDirection(v string) *ListCityMapPersonFlowResponseBodyData {
	s.FlowDirection = &v
	return s
}

type ListCityMapPersonFlowResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCityMapPersonFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCityMapPersonFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapPersonFlowResponse) GoString() string {
	return s.String()
}

func (s *ListCityMapPersonFlowResponse) SetHeaders(v map[string]*string) *ListCityMapPersonFlowResponse {
	s.Headers = v
	return s
}

func (s *ListCityMapPersonFlowResponse) SetBody(v *ListCityMapPersonFlowResponseBody) *ListCityMapPersonFlowResponse {
	s.Body = v
	return s
}

type AddCdrsMonitorRequest struct {
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

func (s AddCdrsMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCdrsMonitorRequest) GoString() string {
	return s.String()
}

func (s *AddCdrsMonitorRequest) SetCorpId(v string) *AddCdrsMonitorRequest {
	s.CorpId = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetMonitorType(v string) *AddCdrsMonitorRequest {
	s.MonitorType = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetDescription(v string) *AddCdrsMonitorRequest {
	s.Description = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetBatchIndicator(v int32) *AddCdrsMonitorRequest {
	s.BatchIndicator = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetAlgorithmVendor(v string) *AddCdrsMonitorRequest {
	s.AlgorithmVendor = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetNotifierType(v string) *AddCdrsMonitorRequest {
	s.NotifierType = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetNotifierUrl(v string) *AddCdrsMonitorRequest {
	s.NotifierUrl = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetNotifierAppSecret(v string) *AddCdrsMonitorRequest {
	s.NotifierAppSecret = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetNotifierTimeOut(v int32) *AddCdrsMonitorRequest {
	s.NotifierTimeOut = &v
	return s
}

func (s *AddCdrsMonitorRequest) SetNotifierExtendValues(v string) *AddCdrsMonitorRequest {
	s.NotifierExtendValues = &v
	return s
}

type AddCdrsMonitorResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddCdrsMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddCdrsMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCdrsMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AddCdrsMonitorResponseBody) SetMessage(v string) *AddCdrsMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *AddCdrsMonitorResponseBody) SetRequestId(v string) *AddCdrsMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCdrsMonitorResponseBody) SetData(v *AddCdrsMonitorResponseBodyData) *AddCdrsMonitorResponseBody {
	s.Data = v
	return s
}

func (s *AddCdrsMonitorResponseBody) SetCode(v string) *AddCdrsMonitorResponseBody {
	s.Code = &v
	return s
}

type AddCdrsMonitorResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddCdrsMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddCdrsMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddCdrsMonitorResponseBodyData) SetTaskId(v string) *AddCdrsMonitorResponseBodyData {
	s.TaskId = &v
	return s
}

type AddCdrsMonitorResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCdrsMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCdrsMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCdrsMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddCdrsMonitorResponse) SetHeaders(v map[string]*string) *AddCdrsMonitorResponse {
	s.Headers = v
	return s
}

func (s *AddCdrsMonitorResponse) SetBody(v *AddCdrsMonitorResponseBody) *AddCdrsMonitorResponse {
	s.Body = v
	return s
}

type ListMapRouteDetailsRequest struct {
	RouteList map[string]interface{} `json:"RouteList,omitempty" xml:"RouteList,omitempty"`
}

func (s ListMapRouteDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMapRouteDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListMapRouteDetailsRequest) SetRouteList(v map[string]interface{}) *ListMapRouteDetailsRequest {
	s.RouteList = v
	return s
}

type ListMapRouteDetailsShrinkRequest struct {
	RouteListShrink *string `json:"RouteList,omitempty" xml:"RouteList,omitempty"`
}

func (s ListMapRouteDetailsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMapRouteDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMapRouteDetailsShrinkRequest) SetRouteListShrink(v string) *ListMapRouteDetailsShrinkRequest {
	s.RouteListShrink = &v
	return s
}

type ListMapRouteDetailsResponseBody struct {
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListMapRouteDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListMapRouteDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMapRouteDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMapRouteDetailsResponseBody) SetTotalCount(v int64) *ListMapRouteDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMapRouteDetailsResponseBody) SetRequestId(v string) *ListMapRouteDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMapRouteDetailsResponseBody) SetMessage(v string) *ListMapRouteDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMapRouteDetailsResponseBody) SetPageSize(v int64) *ListMapRouteDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMapRouteDetailsResponseBody) SetPageNumber(v int64) *ListMapRouteDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMapRouteDetailsResponseBody) SetData(v []*ListMapRouteDetailsResponseBodyData) *ListMapRouteDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListMapRouteDetailsResponseBody) SetCode(v string) *ListMapRouteDetailsResponseBody {
	s.Code = &v
	return s
}

type ListMapRouteDetailsResponseBodyData struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Origin      *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Route       *string `json:"Route,omitempty" xml:"Route,omitempty"`
}

func (s ListMapRouteDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMapRouteDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMapRouteDetailsResponseBodyData) SetType(v string) *ListMapRouteDetailsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListMapRouteDetailsResponseBodyData) SetDestination(v string) *ListMapRouteDetailsResponseBodyData {
	s.Destination = &v
	return s
}

func (s *ListMapRouteDetailsResponseBodyData) SetOrigin(v string) *ListMapRouteDetailsResponseBodyData {
	s.Origin = &v
	return s
}

func (s *ListMapRouteDetailsResponseBodyData) SetRoute(v string) *ListMapRouteDetailsResponseBodyData {
	s.Route = &v
	return s
}

type ListMapRouteDetailsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMapRouteDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMapRouteDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMapRouteDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListMapRouteDetailsResponse) SetHeaders(v map[string]*string) *ListMapRouteDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListMapRouteDetailsResponse) SetBody(v *ListMapRouteDetailsResponseBody) *ListMapRouteDetailsResponse {
	s.Body = v
	return s
}

type ListPersonTopRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonId  *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListPersonTopRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTopRequest) GoString() string {
	return s.String()
}

func (s *ListPersonTopRequest) SetCorpId(v string) *ListPersonTopRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonTopRequest) SetPersonId(v string) *ListPersonTopRequest {
	s.PersonId = &v
	return s
}

func (s *ListPersonTopRequest) SetStartTime(v string) *ListPersonTopRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonTopRequest) SetEndTime(v string) *ListPersonTopRequest {
	s.EndTime = &v
	return s
}

type ListPersonTopResponseBody struct {
	TotalCount *int64                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTopResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonTopResponseBody) SetTotalCount(v int64) *ListPersonTopResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonTopResponseBody) SetRequestId(v string) *ListPersonTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonTopResponseBody) SetMessage(v string) *ListPersonTopResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonTopResponseBody) SetPageSize(v int64) *ListPersonTopResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonTopResponseBody) SetPageNumber(v int64) *ListPersonTopResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTopResponseBody) SetData(v []*ListPersonTopResponseBodyData) *ListPersonTopResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonTopResponseBody) SetCode(v string) *ListPersonTopResponseBody {
	s.Code = &v
	return s
}

type ListPersonTopResponseBodyData struct {
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	PoiId     *string `json:"PoiId,omitempty" xml:"PoiId,omitempty"`
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PassHour  *string `json:"PassHour,omitempty" xml:"PassHour,omitempty"`
	PoiName   *string `json:"PoiName,omitempty" xml:"PoiName,omitempty"`
	PersonId  *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListPersonTopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonTopResponseBodyData) SetFrequency(v string) *ListPersonTopResponseBodyData {
	s.Frequency = &v
	return s
}

func (s *ListPersonTopResponseBodyData) SetPoiId(v string) *ListPersonTopResponseBodyData {
	s.PoiId = &v
	return s
}

func (s *ListPersonTopResponseBodyData) SetCorpId(v string) *ListPersonTopResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListPersonTopResponseBodyData) SetPassHour(v string) *ListPersonTopResponseBodyData {
	s.PassHour = &v
	return s
}

func (s *ListPersonTopResponseBodyData) SetPoiName(v string) *ListPersonTopResponseBodyData {
	s.PoiName = &v
	return s
}

func (s *ListPersonTopResponseBodyData) SetPersonId(v string) *ListPersonTopResponseBodyData {
	s.PersonId = &v
	return s
}

type ListPersonTopResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonTopResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTopResponse) GoString() string {
	return s.String()
}

func (s *ListPersonTopResponse) SetHeaders(v map[string]*string) *ListPersonTopResponse {
	s.Headers = v
	return s
}

func (s *ListPersonTopResponse) SetBody(v *ListPersonTopResponseBody) *ListPersonTopResponse {
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

type ListCityMapAoisRequest struct {
	Radius    *int32  `json:"Radius,omitempty" xml:"Radius,omitempty"`
	Latitude  *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s ListCityMapAoisRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapAoisRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapAoisRequest) SetRadius(v int32) *ListCityMapAoisRequest {
	s.Radius = &v
	return s
}

func (s *ListCityMapAoisRequest) SetLatitude(v string) *ListCityMapAoisRequest {
	s.Latitude = &v
	return s
}

func (s *ListCityMapAoisRequest) SetLongitude(v string) *ListCityMapAoisRequest {
	s.Longitude = &v
	return s
}

type ListCityMapAoisResponseBody struct {
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       []*ListCityMapAoisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCityMapAoisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapAoisResponseBody) GoString() string {
	return s.String()
}

func (s *ListCityMapAoisResponseBody) SetTotalCount(v int64) *ListCityMapAoisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCityMapAoisResponseBody) SetMessage(v string) *ListCityMapAoisResponseBody {
	s.Message = &v
	return s
}

func (s *ListCityMapAoisResponseBody) SetRequestId(v string) *ListCityMapAoisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCityMapAoisResponseBody) SetData(v []*ListCityMapAoisResponseBodyData) *ListCityMapAoisResponseBody {
	s.Data = v
	return s
}

func (s *ListCityMapAoisResponseBody) SetCode(v string) *ListCityMapAoisResponseBody {
	s.Code = &v
	return s
}

type ListCityMapAoisResponseBodyData struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListCityMapAoisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapAoisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCityMapAoisResponseBodyData) SetValue(v string) *ListCityMapAoisResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListCityMapAoisResponseBodyData) SetId(v string) *ListCityMapAoisResponseBodyData {
	s.Id = &v
	return s
}

type ListCityMapAoisResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCityMapAoisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCityMapAoisResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapAoisResponse) GoString() string {
	return s.String()
}

func (s *ListCityMapAoisResponse) SetHeaders(v map[string]*string) *ListCityMapAoisResponse {
	s.Headers = v
	return s
}

func (s *ListCityMapAoisResponse) SetBody(v *ListCityMapAoisResponseBody) *ListCityMapAoisResponse {
	s.Body = v
	return s
}

type RecognizeImageRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Vendor           *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	ImageContent     *string `json:"ImageContent,omitempty" xml:"ImageContent,omitempty"`
	ImageUrl         *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	RecognizeType    *string `json:"RecognizeType,omitempty" xml:"RecognizeType,omitempty"`
	RequireCropImage *bool   `json:"RequireCropImage,omitempty" xml:"RequireCropImage,omitempty"`
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

func (s *RecognizeImageRequest) SetVendor(v string) *RecognizeImageRequest {
	s.Vendor = &v
	return s
}

func (s *RecognizeImageRequest) SetImageContent(v string) *RecognizeImageRequest {
	s.ImageContent = &v
	return s
}

func (s *RecognizeImageRequest) SetImageUrl(v string) *RecognizeImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *RecognizeImageRequest) SetRecognizeType(v string) *RecognizeImageRequest {
	s.RecognizeType = &v
	return s
}

func (s *RecognizeImageRequest) SetRequireCropImage(v bool) *RecognizeImageRequest {
	s.RequireCropImage = &v
	return s
}

type RecognizeImageResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *RecognizeImageResponseBody) SetSuccess(v string) *RecognizeImageResponseBody {
	s.Success = &v
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
	CropAlgorithmCode  *string `json:"CropAlgorithmCode,omitempty" xml:"CropAlgorithmCode,omitempty"`
	RightBottomY       *int32  `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Feature            *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	LeftTopY           *int32  `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	TargetImageContent *string `json:"TargetImageContent,omitempty" xml:"TargetImageContent,omitempty"`
	LeftTopX           *int32  `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	RightBottomX       *int32  `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
}

func (s RecognizeImageResponseBodyDataBodyList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseBodyDataBodyList) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseBodyDataBodyList) SetCropAlgorithmCode(v string) *RecognizeImageResponseBodyDataBodyList {
	s.CropAlgorithmCode = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetRightBottomY(v int32) *RecognizeImageResponseBodyDataBodyList {
	s.RightBottomY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetFeature(v string) *RecognizeImageResponseBodyDataBodyList {
	s.Feature = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetLeftTopY(v int32) *RecognizeImageResponseBodyDataBodyList {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetTargetImageContent(v string) *RecognizeImageResponseBodyDataBodyList {
	s.TargetImageContent = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetLeftTopX(v int32) *RecognizeImageResponseBodyDataBodyList {
	s.LeftTopX = &v
	return s
}

func (s *RecognizeImageResponseBodyDataBodyList) SetRightBottomX(v int32) *RecognizeImageResponseBodyDataBodyList {
	s.RightBottomX = &v
	return s
}

type RecognizeImageResponseBodyDataFaceList struct {
	CropAlgorithmCode   *string  `json:"CropAlgorithmCode,omitempty" xml:"CropAlgorithmCode,omitempty"`
	Feature             *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	RightBottomY        *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	LeftTopY            *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	TargetImageContent  *string  `json:"TargetImageContent,omitempty" xml:"TargetImageContent,omitempty"`
	FaceQuality         *float32 `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	RightBottomX        *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	LeftTopX            *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	FaceKeyPointQuality *float32 `json:"FaceKeyPointQuality,omitempty" xml:"FaceKeyPointQuality,omitempty"`
}

func (s RecognizeImageResponseBodyDataFaceList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseBodyDataFaceList) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseBodyDataFaceList) SetCropAlgorithmCode(v string) *RecognizeImageResponseBodyDataFaceList {
	s.CropAlgorithmCode = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetFeature(v string) *RecognizeImageResponseBodyDataFaceList {
	s.Feature = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetRightBottomY(v int32) *RecognizeImageResponseBodyDataFaceList {
	s.RightBottomY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetLeftTopY(v int32) *RecognizeImageResponseBodyDataFaceList {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetTargetImageContent(v string) *RecognizeImageResponseBodyDataFaceList {
	s.TargetImageContent = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetFaceQuality(v float32) *RecognizeImageResponseBodyDataFaceList {
	s.FaceQuality = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetRightBottomX(v int32) *RecognizeImageResponseBodyDataFaceList {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetLeftTopX(v int32) *RecognizeImageResponseBodyDataFaceList {
	s.LeftTopX = &v
	return s
}

func (s *RecognizeImageResponseBodyDataFaceList) SetFaceKeyPointQuality(v float32) *RecognizeImageResponseBodyDataFaceList {
	s.FaceKeyPointQuality = &v
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

type GetMonitorListRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *GetMonitorListRequest) SetPageNumber(v int32) *GetMonitorListRequest {
	s.PageNumber = &v
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
	TotalPage  *int32                                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *GetMonitorListResponseBodyData) SetTotalPage(v int32) *GetMonitorListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *GetMonitorListResponseBodyData) SetPageNumber(v int32) *GetMonitorListResponseBodyData {
	s.PageNumber = &v
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

type ListDeviceRelationRequest struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s ListDeviceRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceRelationRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceRelationRequest) SetDeviceId(v string) *ListDeviceRelationRequest {
	s.DeviceId = &v
	return s
}

type ListDeviceRelationResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListDeviceRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDeviceRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceRelationResponseBody) SetMessage(v string) *ListDeviceRelationResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceRelationResponseBody) SetRequestId(v string) *ListDeviceRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceRelationResponseBody) SetData(v []*ListDeviceRelationResponseBodyData) *ListDeviceRelationResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceRelationResponseBody) SetCode(v string) *ListDeviceRelationResponseBody {
	s.Code = &v
	return s
}

type ListDeviceRelationResponseBodyData struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListDeviceRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceRelationResponseBodyData) SetCorpId(v string) *ListDeviceRelationResponseBodyData {
	s.CorpId = &v
	return s
}

type ListDeviceRelationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceRelationResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceRelationResponse) SetHeaders(v map[string]*string) *ListDeviceRelationResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceRelationResponse) SetBody(v *ListDeviceRelationResponseBody) *ListDeviceRelationResponse {
	s.Body = v
	return s
}

type ListPersonTrackRequest struct {
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PersonId           *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber         *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ImageSourceType    *string `json:"ImageSourceType,omitempty" xml:"ImageSourceType,omitempty"`
	AggregateDimension *string `json:"AggregateDimension,omitempty" xml:"AggregateDimension,omitempty"`
	QualityScore       *string `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
}

func (s ListPersonTrackRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTrackRequest) GoString() string {
	return s.String()
}

func (s *ListPersonTrackRequest) SetCorpId(v string) *ListPersonTrackRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonTrackRequest) SetPersonId(v string) *ListPersonTrackRequest {
	s.PersonId = &v
	return s
}

func (s *ListPersonTrackRequest) SetStartTime(v string) *ListPersonTrackRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonTrackRequest) SetEndTime(v string) *ListPersonTrackRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonTrackRequest) SetPageNumber(v int64) *ListPersonTrackRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTrackRequest) SetPageSize(v int64) *ListPersonTrackRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonTrackRequest) SetImageSourceType(v string) *ListPersonTrackRequest {
	s.ImageSourceType = &v
	return s
}

func (s *ListPersonTrackRequest) SetAggregateDimension(v string) *ListPersonTrackRequest {
	s.AggregateDimension = &v
	return s
}

func (s *ListPersonTrackRequest) SetQualityScore(v string) *ListPersonTrackRequest {
	s.QualityScore = &v
	return s
}

type ListPersonTrackResponseBody struct {
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonTrackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonTrackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTrackResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonTrackResponseBody) SetTotalCount(v int64) *ListPersonTrackResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonTrackResponseBody) SetRequestId(v string) *ListPersonTrackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonTrackResponseBody) SetMessage(v string) *ListPersonTrackResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonTrackResponseBody) SetPageSize(v int64) *ListPersonTrackResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonTrackResponseBody) SetPageNumber(v int64) *ListPersonTrackResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTrackResponseBody) SetData(v []*ListPersonTrackResponseBodyData) *ListPersonTrackResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonTrackResponseBody) SetCode(v string) *ListPersonTrackResponseBody {
	s.Code = &v
	return s
}

type ListPersonTrackResponseBodyData struct {
	SourceUrl        *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceName   *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	PicUrlPath       *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	TargetPicUrlPath *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	TargetUrl        *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Longitude        *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude         *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	ShotTime         *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	PersonId         *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	TagGender        *string `json:"TagGender,omitempty" xml:"TagGender,omitempty"`
	TagAge           *string `json:"TagAge,omitempty" xml:"TagAge,omitempty"`
}

func (s ListPersonTrackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTrackResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonTrackResponseBodyData) SetSourceUrl(v string) *ListPersonTrackResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetRightBottomY(v string) *ListPersonTrackResponseBodyData {
	s.RightBottomY = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetDataSourceName(v string) *ListPersonTrackResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetPicUrlPath(v string) *ListPersonTrackResponseBodyData {
	s.PicUrlPath = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetDataSourceId(v string) *ListPersonTrackResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetRightBottomX(v string) *ListPersonTrackResponseBodyData {
	s.RightBottomX = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetTargetPicUrlPath(v string) *ListPersonTrackResponseBodyData {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetLeftTopY(v string) *ListPersonTrackResponseBodyData {
	s.LeftTopY = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetTargetUrl(v string) *ListPersonTrackResponseBodyData {
	s.TargetUrl = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetCorpId(v string) *ListPersonTrackResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetLongitude(v string) *ListPersonTrackResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetLatitude(v string) *ListPersonTrackResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetShotTime(v string) *ListPersonTrackResponseBodyData {
	s.ShotTime = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetPersonId(v string) *ListPersonTrackResponseBodyData {
	s.PersonId = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetLeftTopX(v string) *ListPersonTrackResponseBodyData {
	s.LeftTopX = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetTagGender(v string) *ListPersonTrackResponseBodyData {
	s.TagGender = &v
	return s
}

func (s *ListPersonTrackResponseBodyData) SetTagAge(v string) *ListPersonTrackResponseBodyData {
	s.TagAge = &v
	return s
}

type ListPersonTrackResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonTrackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonTrackResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTrackResponse) GoString() string {
	return s.String()
}

func (s *ListPersonTrackResponse) SetHeaders(v map[string]*string) *ListPersonTrackResponse {
	s.Headers = v
	return s
}

func (s *ListPersonTrackResponse) SetBody(v *ListPersonTrackResponseBody) *ListPersonTrackResponse {
	s.Body = v
	return s
}

type ListCityMapCameraResultsRequest struct {
	DataSourceIdList map[string]interface{} `json:"DataSourceIdList,omitempty" xml:"DataSourceIdList,omitempty"`
	PageNum          *int64                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCityMapCameraResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraResultsRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraResultsRequest) SetDataSourceIdList(v map[string]interface{}) *ListCityMapCameraResultsRequest {
	s.DataSourceIdList = v
	return s
}

func (s *ListCityMapCameraResultsRequest) SetPageNum(v int64) *ListCityMapCameraResultsRequest {
	s.PageNum = &v
	return s
}

func (s *ListCityMapCameraResultsRequest) SetPageSize(v int64) *ListCityMapCameraResultsRequest {
	s.PageSize = &v
	return s
}

type ListCityMapCameraResultsShrinkRequest struct {
	DataSourceIdListShrink *string `json:"DataSourceIdList,omitempty" xml:"DataSourceIdList,omitempty"`
	PageNum                *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize               *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCityMapCameraResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraResultsShrinkRequest) SetDataSourceIdListShrink(v string) *ListCityMapCameraResultsShrinkRequest {
	s.DataSourceIdListShrink = &v
	return s
}

func (s *ListCityMapCameraResultsShrinkRequest) SetPageNum(v int64) *ListCityMapCameraResultsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListCityMapCameraResultsShrinkRequest) SetPageSize(v int64) *ListCityMapCameraResultsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListCityMapCameraResultsResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNum    *string                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data       []*ListCityMapCameraResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCityMapCameraResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraResultsResponseBody) SetTotalCount(v int64) *ListCityMapCameraResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBody) SetPageNum(v string) *ListCityMapCameraResultsResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBody) SetRequestId(v string) *ListCityMapCameraResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBody) SetMessage(v string) *ListCityMapCameraResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBody) SetPageSize(v string) *ListCityMapCameraResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBody) SetData(v []*ListCityMapCameraResultsResponseBodyData) *ListCityMapCameraResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListCityMapCameraResultsResponseBody) SetCode(v string) *ListCityMapCameraResultsResponseBody {
	s.Code = &v
	return s
}

type ListCityMapCameraResultsResponseBodyData struct {
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Longitude      *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude       *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	DataSourcePoi  *string `json:"DataSourcePoi,omitempty" xml:"DataSourcePoi,omitempty"`
	NearPoi        *string `json:"NearPoi,omitempty" xml:"NearPoi,omitempty"`
}

func (s ListCityMapCameraResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraResultsResponseBodyData) SetDataSourceName(v string) *ListCityMapCameraResultsResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBodyData) SetDataSourceId(v string) *ListCityMapCameraResultsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBodyData) SetCorpId(v string) *ListCityMapCameraResultsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBodyData) SetLongitude(v string) *ListCityMapCameraResultsResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBodyData) SetLatitude(v string) *ListCityMapCameraResultsResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBodyData) SetDataSourcePoi(v string) *ListCityMapCameraResultsResponseBodyData {
	s.DataSourcePoi = &v
	return s
}

func (s *ListCityMapCameraResultsResponseBodyData) SetNearPoi(v string) *ListCityMapCameraResultsResponseBodyData {
	s.NearPoi = &v
	return s
}

type ListCityMapCameraResultsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCityMapCameraResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCityMapCameraResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraResultsResponse) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraResultsResponse) SetHeaders(v map[string]*string) *ListCityMapCameraResultsResponse {
	s.Headers = v
	return s
}

func (s *ListCityMapCameraResultsResponse) SetBody(v *ListCityMapCameraResultsResponseBody) *ListCityMapCameraResultsResponse {
	s.Body = v
	return s
}

type QueryTrajectoryByIdRequest struct {
	StartTime  *string                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CorpId     *string                                 `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IdList     []*QueryTrajectoryByIdRequestIdList     `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	DeviceList []*QueryTrajectoryByIdRequestDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s QueryTrajectoryByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTrajectoryByIdRequest) GoString() string {
	return s.String()
}

func (s *QueryTrajectoryByIdRequest) SetStartTime(v string) *QueryTrajectoryByIdRequest {
	s.StartTime = &v
	return s
}

func (s *QueryTrajectoryByIdRequest) SetEndTime(v string) *QueryTrajectoryByIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTrajectoryByIdRequest) SetCorpId(v string) *QueryTrajectoryByIdRequest {
	s.CorpId = &v
	return s
}

func (s *QueryTrajectoryByIdRequest) SetIdList(v []*QueryTrajectoryByIdRequestIdList) *QueryTrajectoryByIdRequest {
	s.IdList = v
	return s
}

func (s *QueryTrajectoryByIdRequest) SetDeviceList(v []*QueryTrajectoryByIdRequestDeviceList) *QueryTrajectoryByIdRequest {
	s.DeviceList = v
	return s
}

type QueryTrajectoryByIdRequestIdList struct {
	IdValue *string `json:"IdValue,omitempty" xml:"IdValue,omitempty"`
	IdType  *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s QueryTrajectoryByIdRequestIdList) String() string {
	return tea.Prettify(s)
}

func (s QueryTrajectoryByIdRequestIdList) GoString() string {
	return s.String()
}

func (s *QueryTrajectoryByIdRequestIdList) SetIdValue(v string) *QueryTrajectoryByIdRequestIdList {
	s.IdValue = &v
	return s
}

func (s *QueryTrajectoryByIdRequestIdList) SetIdType(v string) *QueryTrajectoryByIdRequestIdList {
	s.IdType = &v
	return s
}

type QueryTrajectoryByIdRequestDeviceList struct {
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s QueryTrajectoryByIdRequestDeviceList) String() string {
	return tea.Prettify(s)
}

func (s QueryTrajectoryByIdRequestDeviceList) GoString() string {
	return s.String()
}

func (s *QueryTrajectoryByIdRequestDeviceList) SetDeviceId(v string) *QueryTrajectoryByIdRequestDeviceList {
	s.DeviceId = &v
	return s
}

type QueryTrajectoryByIdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTrajectoryByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTrajectoryByIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrajectoryByIdResponseBody) SetMessage(v string) *QueryTrajectoryByIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTrajectoryByIdResponseBody) SetRequestId(v string) *QueryTrajectoryByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrajectoryByIdResponseBody) SetData(v string) *QueryTrajectoryByIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTrajectoryByIdResponseBody) SetCode(v int32) *QueryTrajectoryByIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTrajectoryByIdResponseBody) SetSuccess(v bool) *QueryTrajectoryByIdResponseBody {
	s.Success = &v
	return s
}

type QueryTrajectoryByIdResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTrajectoryByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTrajectoryByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTrajectoryByIdResponse) GoString() string {
	return s.String()
}

func (s *QueryTrajectoryByIdResponse) SetHeaders(v map[string]*string) *QueryTrajectoryByIdResponse {
	s.Headers = v
	return s
}

func (s *QueryTrajectoryByIdResponse) SetBody(v *QueryTrajectoryByIdResponseBody) *QueryTrajectoryByIdResponse {
	s.Body = v
	return s
}

type ListCityMapImageDetailsRequest struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	RecordNumber *int64  `json:"RecordNumber,omitempty" xml:"RecordNumber,omitempty"`
	TimeInterval *string `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s ListCityMapImageDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapImageDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapImageDetailsRequest) SetDataSourceId(v string) *ListCityMapImageDetailsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListCityMapImageDetailsRequest) SetRecordNumber(v int64) *ListCityMapImageDetailsRequest {
	s.RecordNumber = &v
	return s
}

func (s *ListCityMapImageDetailsRequest) SetTimeInterval(v string) *ListCityMapImageDetailsRequest {
	s.TimeInterval = &v
	return s
}

type ListCityMapImageDetailsResponseBody struct {
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCityMapImageDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCityMapImageDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapImageDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCityMapImageDetailsResponseBody) SetTotalCount(v int64) *ListCityMapImageDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBody) SetRequestId(v string) *ListCityMapImageDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBody) SetMessage(v string) *ListCityMapImageDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBody) SetPageSize(v int64) *ListCityMapImageDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBody) SetPageNumber(v int64) *ListCityMapImageDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBody) SetData(v []*ListCityMapImageDetailsResponseBodyData) *ListCityMapImageDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListCityMapImageDetailsResponseBody) SetCode(v string) *ListCityMapImageDetailsResponseBody {
	s.Code = &v
	return s
}

type ListCityMapImageDetailsResponseBodyData struct {
	MotorTargetImageStoragePath  *string `json:"MotorTargetImageStoragePath,omitempty" xml:"MotorTargetImageStoragePath,omitempty"`
	RightBottomY                 *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceId                 *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	RecordId                     *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	VehicleColor                 *string `json:"VehicleColor,omitempty" xml:"VehicleColor,omitempty"`
	SourceImageStoragePath       *string `json:"SourceImageStoragePath,omitempty" xml:"SourceImageStoragePath,omitempty"`
	AgeUpLimit                   *string `json:"AgeUpLimit,omitempty" xml:"AgeUpLimit,omitempty"`
	CoatColor                    *string `json:"CoatColor,omitempty" xml:"CoatColor,omitempty"`
	RightBottomX                 *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	TrousersColorReliability     *string `json:"TrousersColorReliability,omitempty" xml:"TrousersColorReliability,omitempty"`
	AgeLowerLimit                *string `json:"AgeLowerLimit,omitempty" xml:"AgeLowerLimit,omitempty"`
	LeftTopY                     *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ShotTime                     *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	PersonTargetImageStoragePath *string `json:"PersonTargetImageStoragePath,omitempty" xml:"PersonTargetImageStoragePath,omitempty"`
	VehicleClassReliability      *string `json:"VehicleClassReliability,omitempty" xml:"VehicleClassReliability,omitempty"`
	GenderCodeReliability        *string `json:"GenderCodeReliability,omitempty" xml:"GenderCodeReliability,omitempty"`
	Gender                       *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	TrousersColor                *string `json:"TrousersColor,omitempty" xml:"TrousersColor,omitempty"`
	AgeCodeReliability           *string `json:"AgeCodeReliability,omitempty" xml:"AgeCodeReliability,omitempty"`
	FaceTargetImageStoragePath   *string `json:"FaceTargetImageStoragePath,omitempty" xml:"FaceTargetImageStoragePath,omitempty"`
	VehicleClass                 *string `json:"VehicleClass,omitempty" xml:"VehicleClass,omitempty"`
	VehicleColorReliability      *string `json:"VehicleColorReliability,omitempty" xml:"VehicleColorReliability,omitempty"`
	LeftTopX                     *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	AgeLowerLimitReliability     *string `json:"AgeLowerLimitReliability,omitempty" xml:"AgeLowerLimitReliability,omitempty"`
	CoatColorReliability         *string `json:"CoatColorReliability,omitempty" xml:"CoatColorReliability,omitempty"`
}

func (s ListCityMapImageDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapImageDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCityMapImageDetailsResponseBodyData) SetMotorTargetImageStoragePath(v string) *ListCityMapImageDetailsResponseBodyData {
	s.MotorTargetImageStoragePath = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetRightBottomY(v string) *ListCityMapImageDetailsResponseBodyData {
	s.RightBottomY = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetDataSourceId(v string) *ListCityMapImageDetailsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetRecordId(v string) *ListCityMapImageDetailsResponseBodyData {
	s.RecordId = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetVehicleColor(v string) *ListCityMapImageDetailsResponseBodyData {
	s.VehicleColor = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetSourceImageStoragePath(v string) *ListCityMapImageDetailsResponseBodyData {
	s.SourceImageStoragePath = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetAgeUpLimit(v string) *ListCityMapImageDetailsResponseBodyData {
	s.AgeUpLimit = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetCoatColor(v string) *ListCityMapImageDetailsResponseBodyData {
	s.CoatColor = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetRightBottomX(v string) *ListCityMapImageDetailsResponseBodyData {
	s.RightBottomX = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetTrousersColorReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.TrousersColorReliability = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetAgeLowerLimit(v string) *ListCityMapImageDetailsResponseBodyData {
	s.AgeLowerLimit = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetLeftTopY(v string) *ListCityMapImageDetailsResponseBodyData {
	s.LeftTopY = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetShotTime(v string) *ListCityMapImageDetailsResponseBodyData {
	s.ShotTime = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetPersonTargetImageStoragePath(v string) *ListCityMapImageDetailsResponseBodyData {
	s.PersonTargetImageStoragePath = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetVehicleClassReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.VehicleClassReliability = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetGenderCodeReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.GenderCodeReliability = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetGender(v string) *ListCityMapImageDetailsResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetTrousersColor(v string) *ListCityMapImageDetailsResponseBodyData {
	s.TrousersColor = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetAgeCodeReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.AgeCodeReliability = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetFaceTargetImageStoragePath(v string) *ListCityMapImageDetailsResponseBodyData {
	s.FaceTargetImageStoragePath = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetVehicleClass(v string) *ListCityMapImageDetailsResponseBodyData {
	s.VehicleClass = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetVehicleColorReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.VehicleColorReliability = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetLeftTopX(v string) *ListCityMapImageDetailsResponseBodyData {
	s.LeftTopX = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetAgeLowerLimitReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.AgeLowerLimitReliability = &v
	return s
}

func (s *ListCityMapImageDetailsResponseBodyData) SetCoatColorReliability(v string) *ListCityMapImageDetailsResponseBodyData {
	s.CoatColorReliability = &v
	return s
}

type ListCityMapImageDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCityMapImageDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCityMapImageDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapImageDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListCityMapImageDetailsResponse) SetHeaders(v map[string]*string) *ListCityMapImageDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListCityMapImageDetailsResponse) SetBody(v *ListCityMapImageDetailsResponseBody) *ListCityMapImageDetailsResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Icon               *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AggregateSceneCode *string `json:"AggregateSceneCode,omitempty" xml:"AggregateSceneCode,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetIcon(v string) *CreateProjectRequest {
	s.Icon = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetAggregateSceneCode(v string) *CreateProjectRequest {
	s.AggregateSceneCode = &v
	return s
}

type CreateProjectResponseBody struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetCorpId(v string) *CreateProjectResponseBody {
	s.CorpId = &v
	return s
}

func (s *CreateProjectResponseBody) SetMessage(v string) *CreateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetCode(v string) *CreateProjectResponseBody {
	s.Code = &v
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

type ListVehicleTopRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PlateId   *string `json:"PlateId,omitempty" xml:"PlateId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum   *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s ListVehicleTopRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTopRequest) GoString() string {
	return s.String()
}

func (s *ListVehicleTopRequest) SetCorpId(v string) *ListVehicleTopRequest {
	s.CorpId = &v
	return s
}

func (s *ListVehicleTopRequest) SetPlateId(v string) *ListVehicleTopRequest {
	s.PlateId = &v
	return s
}

func (s *ListVehicleTopRequest) SetStartTime(v string) *ListVehicleTopRequest {
	s.StartTime = &v
	return s
}

func (s *ListVehicleTopRequest) SetEndTime(v string) *ListVehicleTopRequest {
	s.EndTime = &v
	return s
}

func (s *ListVehicleTopRequest) SetPageSize(v string) *ListVehicleTopRequest {
	s.PageSize = &v
	return s
}

func (s *ListVehicleTopRequest) SetPageNum(v string) *ListVehicleTopRequest {
	s.PageNum = &v
	return s
}

type ListVehicleTopResponseBody struct {
	TotalCount *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListVehicleTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListVehicleTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTopResponseBody) GoString() string {
	return s.String()
}

func (s *ListVehicleTopResponseBody) SetTotalCount(v int64) *ListVehicleTopResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVehicleTopResponseBody) SetRequestId(v string) *ListVehicleTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVehicleTopResponseBody) SetMessage(v string) *ListVehicleTopResponseBody {
	s.Message = &v
	return s
}

func (s *ListVehicleTopResponseBody) SetPageSize(v int64) *ListVehicleTopResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVehicleTopResponseBody) SetPageNumber(v int64) *ListVehicleTopResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleTopResponseBody) SetData(v []*ListVehicleTopResponseBodyData) *ListVehicleTopResponseBody {
	s.Data = v
	return s
}

func (s *ListVehicleTopResponseBody) SetCode(v string) *ListVehicleTopResponseBody {
	s.Code = &v
	return s
}

type ListVehicleTopResponseBodyData struct {
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	PoiId     *string `json:"PoiId,omitempty" xml:"PoiId,omitempty"`
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	VehicleId *string `json:"VehicleId,omitempty" xml:"VehicleId,omitempty"`
	PassHour  *string `json:"PassHour,omitempty" xml:"PassHour,omitempty"`
	PoiName   *string `json:"PoiName,omitempty" xml:"PoiName,omitempty"`
}

func (s ListVehicleTopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVehicleTopResponseBodyData) SetFrequency(v string) *ListVehicleTopResponseBodyData {
	s.Frequency = &v
	return s
}

func (s *ListVehicleTopResponseBodyData) SetPoiId(v string) *ListVehicleTopResponseBodyData {
	s.PoiId = &v
	return s
}

func (s *ListVehicleTopResponseBodyData) SetCorpId(v string) *ListVehicleTopResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListVehicleTopResponseBodyData) SetVehicleId(v string) *ListVehicleTopResponseBodyData {
	s.VehicleId = &v
	return s
}

func (s *ListVehicleTopResponseBodyData) SetPassHour(v string) *ListVehicleTopResponseBodyData {
	s.PassHour = &v
	return s
}

func (s *ListVehicleTopResponseBodyData) SetPoiName(v string) *ListVehicleTopResponseBodyData {
	s.PoiName = &v
	return s
}

type ListVehicleTopResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVehicleTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVehicleTopResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTopResponse) GoString() string {
	return s.String()
}

func (s *ListVehicleTopResponse) SetHeaders(v map[string]*string) *ListVehicleTopResponse {
	s.Headers = v
	return s
}

func (s *ListVehicleTopResponse) SetBody(v *ListVehicleTopResponseBody) *ListVehicleTopResponse {
	s.Body = v
	return s
}

type ListDataStatisticsByDayRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListDataStatisticsByDayRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsByDayRequest) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsByDayRequest) SetCorpId(v string) *ListDataStatisticsByDayRequest {
	s.CorpId = &v
	return s
}

func (s *ListDataStatisticsByDayRequest) SetStartTime(v string) *ListDataStatisticsByDayRequest {
	s.StartTime = &v
	return s
}

func (s *ListDataStatisticsByDayRequest) SetEndTime(v string) *ListDataStatisticsByDayRequest {
	s.EndTime = &v
	return s
}

type ListDataStatisticsByDayResponseBody struct {
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListDataStatisticsByDayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDataStatisticsByDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsByDayResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsByDayResponseBody) SetTotalCount(v int64) *ListDataStatisticsByDayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBody) SetRequestId(v string) *ListDataStatisticsByDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBody) SetMessage(v string) *ListDataStatisticsByDayResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBody) SetPageSize(v int64) *ListDataStatisticsByDayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBody) SetPageNumber(v int64) *ListDataStatisticsByDayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBody) SetData(v []*ListDataStatisticsByDayResponseBodyData) *ListDataStatisticsByDayResponseBody {
	s.Data = v
	return s
}

func (s *ListDataStatisticsByDayResponseBody) SetCode(v string) *ListDataStatisticsByDayResponseBody {
	s.Code = &v
	return s
}

type ListDataStatisticsByDayResponseBodyData struct {
	NonMotorNumber *int32  `json:"NonMotorNumber,omitempty" xml:"NonMotorNumber,omitempty"`
	FaceNumber     *int32  `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	MotorNumber    *int32  `json:"MotorNumber,omitempty" xml:"MotorNumber,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Date           *string `json:"Date,omitempty" xml:"Date,omitempty"`
	BodyNumber     *int32  `json:"BodyNumber,omitempty" xml:"BodyNumber,omitempty"`
	TotalNumber    *int32  `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s ListDataStatisticsByDayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsByDayResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsByDayResponseBodyData) SetNonMotorNumber(v int32) *ListDataStatisticsByDayResponseBodyData {
	s.NonMotorNumber = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBodyData) SetFaceNumber(v int32) *ListDataStatisticsByDayResponseBodyData {
	s.FaceNumber = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBodyData) SetMotorNumber(v int32) *ListDataStatisticsByDayResponseBodyData {
	s.MotorNumber = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBodyData) SetCorpId(v string) *ListDataStatisticsByDayResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBodyData) SetDate(v string) *ListDataStatisticsByDayResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBodyData) SetBodyNumber(v int32) *ListDataStatisticsByDayResponseBodyData {
	s.BodyNumber = &v
	return s
}

func (s *ListDataStatisticsByDayResponseBodyData) SetTotalNumber(v int32) *ListDataStatisticsByDayResponseBodyData {
	s.TotalNumber = &v
	return s
}

type ListDataStatisticsByDayResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataStatisticsByDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataStatisticsByDayResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataStatisticsByDayResponse) GoString() string {
	return s.String()
}

func (s *ListDataStatisticsByDayResponse) SetHeaders(v map[string]*string) *ListDataStatisticsByDayResponse {
	s.Headers = v
	return s
}

func (s *ListDataStatisticsByDayResponse) SetBody(v *ListDataStatisticsByDayResponseBody) *ListDataStatisticsByDayResponse {
	s.Body = v
	return s
}

type ListVehicleResultsRequest struct {
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	VehicleColor       *string `json:"VehicleColor,omitempty" xml:"VehicleColor,omitempty"`
	VehicleClass       *string `json:"VehicleClass,omitempty" xml:"VehicleClass,omitempty"`
	VehicleApplication *string `json:"VehicleApplication,omitempty" xml:"VehicleApplication,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber         *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVehicleResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleResultsRequest) GoString() string {
	return s.String()
}

func (s *ListVehicleResultsRequest) SetCorpId(v string) *ListVehicleResultsRequest {
	s.CorpId = &v
	return s
}

func (s *ListVehicleResultsRequest) SetVehicleColor(v string) *ListVehicleResultsRequest {
	s.VehicleColor = &v
	return s
}

func (s *ListVehicleResultsRequest) SetVehicleClass(v string) *ListVehicleResultsRequest {
	s.VehicleClass = &v
	return s
}

func (s *ListVehicleResultsRequest) SetVehicleApplication(v string) *ListVehicleResultsRequest {
	s.VehicleApplication = &v
	return s
}

func (s *ListVehicleResultsRequest) SetStartTime(v string) *ListVehicleResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListVehicleResultsRequest) SetEndTime(v string) *ListVehicleResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListVehicleResultsRequest) SetPageNumber(v int64) *ListVehicleResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleResultsRequest) SetPageSize(v int64) *ListVehicleResultsRequest {
	s.PageSize = &v
	return s
}

type ListVehicleResultsResponseBody struct {
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListVehicleResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListVehicleResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVehicleResultsResponseBody) SetTotalCount(v int64) *ListVehicleResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVehicleResultsResponseBody) SetRequestId(v string) *ListVehicleResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVehicleResultsResponseBody) SetMessage(v string) *ListVehicleResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVehicleResultsResponseBody) SetPageSize(v int64) *ListVehicleResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVehicleResultsResponseBody) SetPageNumber(v int64) *ListVehicleResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleResultsResponseBody) SetData(v []*ListVehicleResultsResponseBodyData) *ListVehicleResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListVehicleResultsResponseBody) SetCode(v string) *ListVehicleResultsResponseBody {
	s.Code = &v
	return s
}

type ListVehicleResultsResponseBodyData struct {
	VehicleApplication *string `json:"VehicleApplication,omitempty" xml:"VehicleApplication,omitempty"`
	Profession         *string `json:"Profession,omitempty" xml:"Profession,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Gender             *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateId            *string `json:"PlateId,omitempty" xml:"PlateId,omitempty"`
	VehicleClass       *string `json:"VehicleClass,omitempty" xml:"VehicleClass,omitempty"`
	LiveAddress        *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	VehicleId          *string `json:"VehicleId,omitempty" xml:"VehicleId,omitempty"`
	PersonId           *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListVehicleResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVehicleResultsResponseBodyData) SetVehicleApplication(v string) *ListVehicleResultsResponseBodyData {
	s.VehicleApplication = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetProfession(v string) *ListVehicleResultsResponseBodyData {
	s.Profession = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetUpdateTime(v string) *ListVehicleResultsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetGender(v string) *ListVehicleResultsResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetPlateId(v string) *ListVehicleResultsResponseBodyData {
	s.PlateId = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetVehicleClass(v string) *ListVehicleResultsResponseBodyData {
	s.VehicleClass = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetLiveAddress(v string) *ListVehicleResultsResponseBodyData {
	s.LiveAddress = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetVehicleId(v string) *ListVehicleResultsResponseBodyData {
	s.VehicleId = &v
	return s
}

func (s *ListVehicleResultsResponseBodyData) SetPersonId(v string) *ListVehicleResultsResponseBodyData {
	s.PersonId = &v
	return s
}

type ListVehicleResultsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVehicleResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVehicleResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleResultsResponse) GoString() string {
	return s.String()
}

func (s *ListVehicleResultsResponse) SetHeaders(v map[string]*string) *ListVehicleResultsResponse {
	s.Headers = v
	return s
}

func (s *ListVehicleResultsResponse) SetBody(v *ListVehicleResultsResponseBody) *ListVehicleResultsResponse {
	s.Body = v
	return s
}

type SearchAggregateObjectRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	ObjectType        *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Vendor            *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	Feature           *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	ImageContent      *string `json:"ImageContent,omitempty" xml:"ImageContent,omitempty"`
	ImageUrl          *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	DeviceList        *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	Attributes        *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	ShotTimeStart     *string `json:"ShotTimeStart,omitempty" xml:"ShotTimeStart,omitempty"`
	ShotTimeEnd       *string `json:"ShotTimeEnd,omitempty" xml:"ShotTimeEnd,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequireTotalCount *bool   `json:"RequireTotalCount,omitempty" xml:"RequireTotalCount,omitempty"`
}

func (s SearchAggregateObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectRequest) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectRequest) SetCorpId(v string) *SearchAggregateObjectRequest {
	s.CorpId = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetObjectType(v string) *SearchAggregateObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetVendor(v string) *SearchAggregateObjectRequest {
	s.Vendor = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetFeature(v string) *SearchAggregateObjectRequest {
	s.Feature = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetImageContent(v string) *SearchAggregateObjectRequest {
	s.ImageContent = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetImageUrl(v string) *SearchAggregateObjectRequest {
	s.ImageUrl = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetDeviceList(v string) *SearchAggregateObjectRequest {
	s.DeviceList = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetAttributes(v string) *SearchAggregateObjectRequest {
	s.Attributes = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetShotTimeStart(v string) *SearchAggregateObjectRequest {
	s.ShotTimeStart = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetShotTimeEnd(v string) *SearchAggregateObjectRequest {
	s.ShotTimeEnd = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetPageNumber(v int32) *SearchAggregateObjectRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetPageSize(v int32) *SearchAggregateObjectRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAggregateObjectRequest) SetRequireTotalCount(v bool) *SearchAggregateObjectRequest {
	s.RequireTotalCount = &v
	return s
}

type SearchAggregateObjectResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize  *int64                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	Data      *SearchAggregateObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchAggregateObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponseBody) SetRequestId(v string) *SearchAggregateObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAggregateObjectResponseBody) SetMessage(v string) *SearchAggregateObjectResponseBody {
	s.Message = &v
	return s
}

func (s *SearchAggregateObjectResponseBody) SetPageSize(v int64) *SearchAggregateObjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchAggregateObjectResponseBody) SetTotal(v int32) *SearchAggregateObjectResponseBody {
	s.Total = &v
	return s
}

func (s *SearchAggregateObjectResponseBody) SetData(v *SearchAggregateObjectResponseBodyData) *SearchAggregateObjectResponseBody {
	s.Data = v
	return s
}

func (s *SearchAggregateObjectResponseBody) SetCode(v string) *SearchAggregateObjectResponseBody {
	s.Code = &v
	return s
}

func (s *SearchAggregateObjectResponseBody) SetSuccess(v bool) *SearchAggregateObjectResponseBody {
	s.Success = &v
	return s
}

type SearchAggregateObjectResponseBodyData struct {
	BodyList     []*SearchAggregateObjectResponseBodyDataBodyList     `json:"BodyList,omitempty" xml:"BodyList,omitempty" type:"Repeated"`
	FaceList     []*SearchAggregateObjectResponseBodyDataFaceList     `json:"FaceList,omitempty" xml:"FaceList,omitempty" type:"Repeated"`
	MotorList    []*SearchAggregateObjectResponseBodyDataMotorList    `json:"MotorList,omitempty" xml:"MotorList,omitempty" type:"Repeated"`
	NonMotorList []*SearchAggregateObjectResponseBodyDataNonMotorList `json:"NonMotorList,omitempty" xml:"NonMotorList,omitempty" type:"Repeated"`
}

func (s SearchAggregateObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponseBodyData) SetBodyList(v []*SearchAggregateObjectResponseBodyDataBodyList) *SearchAggregateObjectResponseBodyData {
	s.BodyList = v
	return s
}

func (s *SearchAggregateObjectResponseBodyData) SetFaceList(v []*SearchAggregateObjectResponseBodyDataFaceList) *SearchAggregateObjectResponseBodyData {
	s.FaceList = v
	return s
}

func (s *SearchAggregateObjectResponseBodyData) SetMotorList(v []*SearchAggregateObjectResponseBodyDataMotorList) *SearchAggregateObjectResponseBodyData {
	s.MotorList = v
	return s
}

func (s *SearchAggregateObjectResponseBodyData) SetNonMotorList(v []*SearchAggregateObjectResponseBodyDataNonMotorList) *SearchAggregateObjectResponseBodyData {
	s.NonMotorList = v
	return s
}

type SearchAggregateObjectResponseBodyDataBodyList struct {
	DeviceID        *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType      *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	DeviceName      *string  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	RightBottomY    *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	RightBottomX    *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	DeviceLongitude *float32 `json:"DeviceLongitude,omitempty" xml:"DeviceLongitude,omitempty"`
	SourceImageUrl  *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	LeftTopY        *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ShotTime        *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	PersonId        *string  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	LeftTopX        *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	DeviceLatitude  *float32 `json:"DeviceLatitude,omitempty" xml:"DeviceLatitude,omitempty"`
}

func (s SearchAggregateObjectResponseBodyDataBodyList) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponseBodyDataBodyList) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetDeviceID(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.DeviceID = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetObjectType(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.ObjectType = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetDeviceName(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.DeviceName = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetRightBottomY(v int32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.RightBottomY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetScore(v float32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.Score = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetRightBottomX(v int32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.RightBottomX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetDeviceLongitude(v float32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.DeviceLongitude = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetSourceImageUrl(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetTargetImageUrl(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetLeftTopY(v int32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.LeftTopY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetShotTime(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.ShotTime = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetPersonId(v string) *SearchAggregateObjectResponseBodyDataBodyList {
	s.PersonId = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetLeftTopX(v int32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.LeftTopX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataBodyList) SetDeviceLatitude(v float32) *SearchAggregateObjectResponseBodyDataBodyList {
	s.DeviceLatitude = &v
	return s
}

type SearchAggregateObjectResponseBodyDataFaceList struct {
	DeviceID        *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType      *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	DeviceName      *string  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	RightBottomY    *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	RightBottomX    *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	DeviceLongitude *float32 `json:"DeviceLongitude,omitempty" xml:"DeviceLongitude,omitempty"`
	SourceImageUrl  *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	LeftTopY        *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ShotTime        *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	PersonId        *string  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	LeftTopX        *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	DeviceLatitude  *float32 `json:"DeviceLatitude,omitempty" xml:"DeviceLatitude,omitempty"`
}

func (s SearchAggregateObjectResponseBodyDataFaceList) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponseBodyDataFaceList) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetDeviceID(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.DeviceID = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetObjectType(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.ObjectType = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetDeviceName(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.DeviceName = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetRightBottomY(v int32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.RightBottomY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetScore(v float32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.Score = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetRightBottomX(v int32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.RightBottomX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetDeviceLongitude(v float32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.DeviceLongitude = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetSourceImageUrl(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetTargetImageUrl(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetLeftTopY(v int32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.LeftTopY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetShotTime(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.ShotTime = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetPersonId(v string) *SearchAggregateObjectResponseBodyDataFaceList {
	s.PersonId = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetLeftTopX(v int32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.LeftTopX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataFaceList) SetDeviceLatitude(v float32) *SearchAggregateObjectResponseBodyDataFaceList {
	s.DeviceLatitude = &v
	return s
}

type SearchAggregateObjectResponseBodyDataMotorList struct {
	DeviceID        *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType      *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	DeviceName      *string  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	RightBottomY    *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	RightBottomX    *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	DeviceLongitude *float32 `json:"DeviceLongitude,omitempty" xml:"DeviceLongitude,omitempty"`
	SourceImageUrl  *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	LeftTopY        *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ShotTime        *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	PersonId        *string  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	LeftTopX        *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	DeviceLatitude  *float32 `json:"DeviceLatitude,omitempty" xml:"DeviceLatitude,omitempty"`
}

func (s SearchAggregateObjectResponseBodyDataMotorList) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponseBodyDataMotorList) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetDeviceID(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.DeviceID = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetObjectType(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.ObjectType = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetDeviceName(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.DeviceName = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetRightBottomY(v int32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.RightBottomY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetScore(v float32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.Score = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetRightBottomX(v int32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.RightBottomX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetDeviceLongitude(v float32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.DeviceLongitude = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetSourceImageUrl(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetTargetImageUrl(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetLeftTopY(v int32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.LeftTopY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetShotTime(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.ShotTime = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetPersonId(v string) *SearchAggregateObjectResponseBodyDataMotorList {
	s.PersonId = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetLeftTopX(v int32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.LeftTopX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataMotorList) SetDeviceLatitude(v float32) *SearchAggregateObjectResponseBodyDataMotorList {
	s.DeviceLatitude = &v
	return s
}

type SearchAggregateObjectResponseBodyDataNonMotorList struct {
	DeviceID        *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty"`
	ObjectType      *string  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	DeviceName      *string  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	RightBottomY    *int32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	RightBottomX    *int32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	DeviceLongitude *float32 `json:"DeviceLongitude,omitempty" xml:"DeviceLongitude,omitempty"`
	SourceImageUrl  *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
	LeftTopY        *int32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	ShotTime        *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty"`
	PersonId        *string  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	LeftTopX        *int32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	DeviceLatitude  *float32 `json:"DeviceLatitude,omitempty" xml:"DeviceLatitude,omitempty"`
}

func (s SearchAggregateObjectResponseBodyDataNonMotorList) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponseBodyDataNonMotorList) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetDeviceID(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.DeviceID = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetObjectType(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.ObjectType = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetDeviceName(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.DeviceName = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetRightBottomY(v int32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.RightBottomY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetScore(v float32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.Score = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetRightBottomX(v int32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.RightBottomX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetDeviceLongitude(v float32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.DeviceLongitude = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetSourceImageUrl(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetTargetImageUrl(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetLeftTopY(v int32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.LeftTopY = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetShotTime(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.ShotTime = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetPersonId(v string) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.PersonId = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetLeftTopX(v int32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.LeftTopX = &v
	return s
}

func (s *SearchAggregateObjectResponseBodyDataNonMotorList) SetDeviceLatitude(v float32) *SearchAggregateObjectResponseBodyDataNonMotorList {
	s.DeviceLatitude = &v
	return s
}

type SearchAggregateObjectResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchAggregateObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAggregateObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAggregateObjectResponse) GoString() string {
	return s.String()
}

func (s *SearchAggregateObjectResponse) SetHeaders(v map[string]*string) *SearchAggregateObjectResponse {
	s.Headers = v
	return s
}

func (s *SearchAggregateObjectResponse) SetBody(v *SearchAggregateObjectResponseBody) *SearchAggregateObjectResponse {
	s.Body = v
	return s
}

type ListCorpMetricsStatisticRequest struct {
	CorpId          *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode         *string                `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	StartTime       *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber      *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserGroupList   map[string]interface{} `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty"`
	DeviceGroupList map[string]interface{} `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty"`
	DeviceIdList    map[string]interface{} `json:"DeviceIdList,omitempty" xml:"DeviceIdList,omitempty"`
	QualitScore     *string                `json:"QualitScore,omitempty" xml:"QualitScore,omitempty"`
}

func (s ListCorpMetricsStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsStatisticRequest) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsStatisticRequest) SetCorpId(v string) *ListCorpMetricsStatisticRequest {
	s.CorpId = &v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetTagCode(v string) *ListCorpMetricsStatisticRequest {
	s.TagCode = &v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetStartTime(v string) *ListCorpMetricsStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetEndTime(v string) *ListCorpMetricsStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetPageNumber(v int64) *ListCorpMetricsStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetPageSize(v int64) *ListCorpMetricsStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetUserGroupList(v map[string]interface{}) *ListCorpMetricsStatisticRequest {
	s.UserGroupList = v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetDeviceGroupList(v map[string]interface{}) *ListCorpMetricsStatisticRequest {
	s.DeviceGroupList = v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetDeviceIdList(v map[string]interface{}) *ListCorpMetricsStatisticRequest {
	s.DeviceIdList = v
	return s
}

func (s *ListCorpMetricsStatisticRequest) SetQualitScore(v string) *ListCorpMetricsStatisticRequest {
	s.QualitScore = &v
	return s
}

type ListCorpMetricsStatisticShrinkRequest struct {
	CorpId                *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode               *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserGroupListShrink   *string `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty"`
	DeviceGroupListShrink *string `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty"`
	DeviceIdListShrink    *string `json:"DeviceIdList,omitempty" xml:"DeviceIdList,omitempty"`
	QualitScore           *string `json:"QualitScore,omitempty" xml:"QualitScore,omitempty"`
}

func (s ListCorpMetricsStatisticShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsStatisticShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetCorpId(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetTagCode(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.TagCode = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetStartTime(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetEndTime(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetPageNumber(v int64) *ListCorpMetricsStatisticShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetPageSize(v int64) *ListCorpMetricsStatisticShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetUserGroupListShrink(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.UserGroupListShrink = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetDeviceGroupListShrink(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.DeviceGroupListShrink = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetDeviceIdListShrink(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.DeviceIdListShrink = &v
	return s
}

func (s *ListCorpMetricsStatisticShrinkRequest) SetQualitScore(v string) *ListCorpMetricsStatisticShrinkRequest {
	s.QualitScore = &v
	return s
}

type ListCorpMetricsStatisticResponseBody struct {
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCorpMetricsStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *string                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCorpMetricsStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsStatisticResponseBody) SetTotalCount(v int32) *ListCorpMetricsStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetRequestId(v string) *ListCorpMetricsStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetMessage(v string) *ListCorpMetricsStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetPageSize(v int32) *ListCorpMetricsStatisticResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetPageNumber(v int32) *ListCorpMetricsStatisticResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetData(v []*ListCorpMetricsStatisticResponseBodyData) *ListCorpMetricsStatisticResponseBody {
	s.Data = v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetCode(v string) *ListCorpMetricsStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBody) SetSuccess(v string) *ListCorpMetricsStatisticResponseBody {
	s.Success = &v
	return s
}

type ListCorpMetricsStatisticResponseBodyData struct {
	DateId        *string `json:"DateId,omitempty" xml:"DateId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagMetrics    *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty"`
	TagValue      *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	PersonId      *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListCorpMetricsStatisticResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetDateId(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.DateId = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetDeviceGroupId(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetDeviceId(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetTagCode(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetUserGroupId(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetCorpId(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetTagMetrics(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.TagMetrics = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetTagValue(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListCorpMetricsStatisticResponseBodyData) SetPersonId(v string) *ListCorpMetricsStatisticResponseBodyData {
	s.PersonId = &v
	return s
}

type ListCorpMetricsStatisticResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorpMetricsStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorpMetricsStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsStatisticResponse) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsStatisticResponse) SetHeaders(v map[string]*string) *ListCorpMetricsStatisticResponse {
	s.Headers = v
	return s
}

func (s *ListCorpMetricsStatisticResponse) SetBody(v *ListCorpMetricsStatisticResponseBody) *ListCorpMetricsStatisticResponse {
	s.Body = v
	return s
}

type DetectTrajectoryRegularPatternRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IdType      *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	IdValue     *string `json:"IdValue,omitempty" xml:"IdValue,omitempty"`
	PredictDate *string `json:"PredictDate,omitempty" xml:"PredictDate,omitempty"`
}

func (s DetectTrajectoryRegularPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectTrajectoryRegularPatternRequest) GoString() string {
	return s.String()
}

func (s *DetectTrajectoryRegularPatternRequest) SetCorpId(v string) *DetectTrajectoryRegularPatternRequest {
	s.CorpId = &v
	return s
}

func (s *DetectTrajectoryRegularPatternRequest) SetIdType(v string) *DetectTrajectoryRegularPatternRequest {
	s.IdType = &v
	return s
}

func (s *DetectTrajectoryRegularPatternRequest) SetIdValue(v string) *DetectTrajectoryRegularPatternRequest {
	s.IdValue = &v
	return s
}

func (s *DetectTrajectoryRegularPatternRequest) SetPredictDate(v string) *DetectTrajectoryRegularPatternRequest {
	s.PredictDate = &v
	return s
}

type DetectTrajectoryRegularPatternResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectTrajectoryRegularPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectTrajectoryRegularPatternResponseBody) GoString() string {
	return s.String()
}

func (s *DetectTrajectoryRegularPatternResponseBody) SetMessage(v string) *DetectTrajectoryRegularPatternResponseBody {
	s.Message = &v
	return s
}

func (s *DetectTrajectoryRegularPatternResponseBody) SetRequestId(v string) *DetectTrajectoryRegularPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectTrajectoryRegularPatternResponseBody) SetData(v string) *DetectTrajectoryRegularPatternResponseBody {
	s.Data = &v
	return s
}

func (s *DetectTrajectoryRegularPatternResponseBody) SetCode(v int32) *DetectTrajectoryRegularPatternResponseBody {
	s.Code = &v
	return s
}

func (s *DetectTrajectoryRegularPatternResponseBody) SetSuccess(v bool) *DetectTrajectoryRegularPatternResponseBody {
	s.Success = &v
	return s
}

type DetectTrajectoryRegularPatternResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectTrajectoryRegularPatternResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectTrajectoryRegularPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectTrajectoryRegularPatternResponse) GoString() string {
	return s.String()
}

func (s *DetectTrajectoryRegularPatternResponse) SetHeaders(v map[string]*string) *DetectTrajectoryRegularPatternResponse {
	s.Headers = v
	return s
}

func (s *DetectTrajectoryRegularPatternResponse) SetBody(v *DetectTrajectoryRegularPatternResponseBody) *DetectTrajectoryRegularPatternResponse {
	s.Body = v
	return s
}

type ListVehicleTrackRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PlateId    *string `json:"PlateId,omitempty" xml:"PlateId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListVehicleTrackRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTrackRequest) GoString() string {
	return s.String()
}

func (s *ListVehicleTrackRequest) SetCorpId(v string) *ListVehicleTrackRequest {
	s.CorpId = &v
	return s
}

func (s *ListVehicleTrackRequest) SetPlateId(v string) *ListVehicleTrackRequest {
	s.PlateId = &v
	return s
}

func (s *ListVehicleTrackRequest) SetStartTime(v string) *ListVehicleTrackRequest {
	s.StartTime = &v
	return s
}

func (s *ListVehicleTrackRequest) SetEndTime(v string) *ListVehicleTrackRequest {
	s.EndTime = &v
	return s
}

func (s *ListVehicleTrackRequest) SetPageSize(v int64) *ListVehicleTrackRequest {
	s.PageSize = &v
	return s
}

func (s *ListVehicleTrackRequest) SetPageNumber(v int64) *ListVehicleTrackRequest {
	s.PageNumber = &v
	return s
}

type ListVehicleTrackResponseBody struct {
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListVehicleTrackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListVehicleTrackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTrackResponseBody) GoString() string {
	return s.String()
}

func (s *ListVehicleTrackResponseBody) SetTotalCount(v int64) *ListVehicleTrackResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVehicleTrackResponseBody) SetRequestId(v string) *ListVehicleTrackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVehicleTrackResponseBody) SetMessage(v string) *ListVehicleTrackResponseBody {
	s.Message = &v
	return s
}

func (s *ListVehicleTrackResponseBody) SetPageSize(v int64) *ListVehicleTrackResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVehicleTrackResponseBody) SetPageNumber(v int64) *ListVehicleTrackResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleTrackResponseBody) SetData(v []*ListVehicleTrackResponseBodyData) *ListVehicleTrackResponseBody {
	s.Data = v
	return s
}

func (s *ListVehicleTrackResponseBody) SetCode(v string) *ListVehicleTrackResponseBody {
	s.Code = &v
	return s
}

type ListVehicleTrackResponseBodyData struct {
	SourceUrl        *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty"`
	DataSourceName   *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	PicUrlPath       *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty"`
	TargetPicUrlPath *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	PlateId          *string `json:"PlateId,omitempty" xml:"PlateId,omitempty"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty"`
	TargetUrl        *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Longitude        *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude         *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty"`
	PassTime         *string `json:"PassTime,omitempty" xml:"PassTime,omitempty"`
}

func (s ListVehicleTrackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTrackResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVehicleTrackResponseBodyData) SetSourceUrl(v string) *ListVehicleTrackResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetRightBottomY(v string) *ListVehicleTrackResponseBodyData {
	s.RightBottomY = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetDataSourceName(v string) *ListVehicleTrackResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetPicUrlPath(v string) *ListVehicleTrackResponseBodyData {
	s.PicUrlPath = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetDataSourceId(v string) *ListVehicleTrackResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetRightBottomX(v string) *ListVehicleTrackResponseBodyData {
	s.RightBottomX = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetTargetPicUrlPath(v string) *ListVehicleTrackResponseBodyData {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetPlateId(v string) *ListVehicleTrackResponseBodyData {
	s.PlateId = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetLeftTopY(v string) *ListVehicleTrackResponseBodyData {
	s.LeftTopY = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetTargetUrl(v string) *ListVehicleTrackResponseBodyData {
	s.TargetUrl = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetCorpId(v string) *ListVehicleTrackResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetLongitude(v string) *ListVehicleTrackResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetLatitude(v string) *ListVehicleTrackResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetLeftTopX(v string) *ListVehicleTrackResponseBodyData {
	s.LeftTopX = &v
	return s
}

func (s *ListVehicleTrackResponseBodyData) SetPassTime(v string) *ListVehicleTrackResponseBodyData {
	s.PassTime = &v
	return s
}

type ListVehicleTrackResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVehicleTrackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVehicleTrackResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleTrackResponse) GoString() string {
	return s.String()
}

func (s *ListVehicleTrackResponse) SetHeaders(v map[string]*string) *ListVehicleTrackResponse {
	s.Headers = v
	return s
}

func (s *ListVehicleTrackResponse) SetBody(v *ListVehicleTrackResponseBody) *ListVehicleTrackResponse {
	s.Body = v
	return s
}

type ListStructureStatisticsRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	BackCategory *string `json:"BackCategory,omitempty" xml:"BackCategory,omitempty"`
}

func (s ListStructureStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStructureStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListStructureStatisticsRequest) SetCorpId(v string) *ListStructureStatisticsRequest {
	s.CorpId = &v
	return s
}

func (s *ListStructureStatisticsRequest) SetBackCategory(v string) *ListStructureStatisticsRequest {
	s.BackCategory = &v
	return s
}

type ListStructureStatisticsResponseBody struct {
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListStructureStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListStructureStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStructureStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStructureStatisticsResponseBody) SetTotalCount(v int64) *ListStructureStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStructureStatisticsResponseBody) SetRequestId(v string) *ListStructureStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStructureStatisticsResponseBody) SetMessage(v string) *ListStructureStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStructureStatisticsResponseBody) SetPageSize(v int64) *ListStructureStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStructureStatisticsResponseBody) SetPageNumber(v int64) *ListStructureStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStructureStatisticsResponseBody) SetData(v []*ListStructureStatisticsResponseBodyData) *ListStructureStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListStructureStatisticsResponseBody) SetCode(v string) *ListStructureStatisticsResponseBody {
	s.Code = &v
	return s
}

type ListStructureStatisticsResponseBodyData struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ListStructureStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListStructureStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStructureStatisticsResponseBodyData) SetCorpId(v string) *ListStructureStatisticsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListStructureStatisticsResponseBodyData) SetNumber(v string) *ListStructureStatisticsResponseBodyData {
	s.Number = &v
	return s
}

type ListStructureStatisticsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStructureStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStructureStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStructureStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListStructureStatisticsResponse) SetHeaders(v map[string]*string) *ListStructureStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListStructureStatisticsResponse) SetBody(v *ListStructureStatisticsResponseBody) *ListStructureStatisticsResponse {
	s.Body = v
	return s
}

type StopMonitorRequest struct {
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty"`
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
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

func (s *StopMonitorRequest) SetCorpId(v string) *StopMonitorRequest {
	s.CorpId = &v
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

type PredictTrajectoryDestinationRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IdType          *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	IdValue         *string `json:"IdValue,omitempty" xml:"IdValue,omitempty"`
	PredictTimeSpan *int32  `json:"PredictTimeSpan,omitempty" xml:"PredictTimeSpan,omitempty"`
}

func (s PredictTrajectoryDestinationRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictTrajectoryDestinationRequest) GoString() string {
	return s.String()
}

func (s *PredictTrajectoryDestinationRequest) SetCorpId(v string) *PredictTrajectoryDestinationRequest {
	s.CorpId = &v
	return s
}

func (s *PredictTrajectoryDestinationRequest) SetIdType(v string) *PredictTrajectoryDestinationRequest {
	s.IdType = &v
	return s
}

func (s *PredictTrajectoryDestinationRequest) SetIdValue(v string) *PredictTrajectoryDestinationRequest {
	s.IdValue = &v
	return s
}

func (s *PredictTrajectoryDestinationRequest) SetPredictTimeSpan(v int32) *PredictTrajectoryDestinationRequest {
	s.PredictTimeSpan = &v
	return s
}

type PredictTrajectoryDestinationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PredictTrajectoryDestinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictTrajectoryDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *PredictTrajectoryDestinationResponseBody) SetMessage(v string) *PredictTrajectoryDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *PredictTrajectoryDestinationResponseBody) SetRequestId(v string) *PredictTrajectoryDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictTrajectoryDestinationResponseBody) SetData(v string) *PredictTrajectoryDestinationResponseBody {
	s.Data = &v
	return s
}

func (s *PredictTrajectoryDestinationResponseBody) SetCode(v int32) *PredictTrajectoryDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *PredictTrajectoryDestinationResponseBody) SetSuccess(v bool) *PredictTrajectoryDestinationResponseBody {
	s.Success = &v
	return s
}

type PredictTrajectoryDestinationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PredictTrajectoryDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredictTrajectoryDestinationResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictTrajectoryDestinationResponse) GoString() string {
	return s.String()
}

func (s *PredictTrajectoryDestinationResponse) SetHeaders(v map[string]*string) *PredictTrajectoryDestinationResponse {
	s.Headers = v
	return s
}

func (s *PredictTrajectoryDestinationResponse) SetBody(v *PredictTrajectoryDestinationResponseBody) *PredictTrajectoryDestinationResponse {
	s.Body = v
	return s
}

type ListRangeDeviceRequest struct {
	Radius       *int32  `json:"Radius,omitempty" xml:"Radius,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListRangeDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRangeDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListRangeDeviceRequest) SetRadius(v int32) *ListRangeDeviceRequest {
	s.Radius = &v
	return s
}

func (s *ListRangeDeviceRequest) SetDataSourceId(v string) *ListRangeDeviceRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListRangeDeviceRequest) SetPageNumber(v int64) *ListRangeDeviceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRangeDeviceRequest) SetPageSize(v int64) *ListRangeDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *ListRangeDeviceRequest) SetCorpId(v string) *ListRangeDeviceRequest {
	s.CorpId = &v
	return s
}

type ListRangeDeviceResponseBody struct {
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListRangeDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListRangeDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRangeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListRangeDeviceResponseBody) SetTotalCount(v int64) *ListRangeDeviceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRangeDeviceResponseBody) SetRequestId(v string) *ListRangeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRangeDeviceResponseBody) SetMessage(v string) *ListRangeDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListRangeDeviceResponseBody) SetPageSize(v int64) *ListRangeDeviceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRangeDeviceResponseBody) SetPageNumber(v int64) *ListRangeDeviceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRangeDeviceResponseBody) SetData(v []*ListRangeDeviceResponseBodyData) *ListRangeDeviceResponseBody {
	s.Data = v
	return s
}

func (s *ListRangeDeviceResponseBody) SetCode(v string) *ListRangeDeviceResponseBody {
	s.Code = &v
	return s
}

type ListRangeDeviceResponseBodyData struct {
	DataSourceIdPoi  *string `json:"DataSourceIdPoi,omitempty" xml:"DataSourceIdPoi,omitempty"`
	Distance         *string `json:"Distance,omitempty" xml:"Distance,omitempty"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Longitude        *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	DataSourceIdName *string `json:"DataSourceIdName,omitempty" xml:"DataSourceIdName,omitempty"`
	Latitude         *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	NearPoi          *string `json:"NearPoi,omitempty" xml:"NearPoi,omitempty"`
}

func (s ListRangeDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRangeDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRangeDeviceResponseBodyData) SetDataSourceIdPoi(v string) *ListRangeDeviceResponseBodyData {
	s.DataSourceIdPoi = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetDistance(v string) *ListRangeDeviceResponseBodyData {
	s.Distance = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetDataSourceId(v string) *ListRangeDeviceResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetCorpId(v string) *ListRangeDeviceResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetLongitude(v string) *ListRangeDeviceResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetDataSourceIdName(v string) *ListRangeDeviceResponseBodyData {
	s.DataSourceIdName = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetLatitude(v string) *ListRangeDeviceResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListRangeDeviceResponseBodyData) SetNearPoi(v string) *ListRangeDeviceResponseBodyData {
	s.NearPoi = &v
	return s
}

type ListRangeDeviceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRangeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRangeDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRangeDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListRangeDeviceResponse) SetHeaders(v map[string]*string) *ListRangeDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListRangeDeviceResponse) SetBody(v *ListRangeDeviceResponseBody) *ListRangeDeviceResponse {
	s.Body = v
	return s
}

type ListCityMapRangeStatisticRequest struct {
	Radius     *int32  `json:"Radius,omitempty" xml:"Radius,omitempty"`
	Latitude   *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude  *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCityMapRangeStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapRangeStatisticRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapRangeStatisticRequest) SetRadius(v int32) *ListCityMapRangeStatisticRequest {
	s.Radius = &v
	return s
}

func (s *ListCityMapRangeStatisticRequest) SetLatitude(v string) *ListCityMapRangeStatisticRequest {
	s.Latitude = &v
	return s
}

func (s *ListCityMapRangeStatisticRequest) SetLongitude(v string) *ListCityMapRangeStatisticRequest {
	s.Longitude = &v
	return s
}

func (s *ListCityMapRangeStatisticRequest) SetEndTime(v string) *ListCityMapRangeStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *ListCityMapRangeStatisticRequest) SetPageNumber(v int64) *ListCityMapRangeStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapRangeStatisticRequest) SetPageSize(v int64) *ListCityMapRangeStatisticRequest {
	s.PageSize = &v
	return s
}

type ListCityMapRangeStatisticResponseBody struct {
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCityMapRangeStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCityMapRangeStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapRangeStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *ListCityMapRangeStatisticResponseBody) SetTotalCount(v int64) *ListCityMapRangeStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBody) SetRequestId(v string) *ListCityMapRangeStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBody) SetMessage(v string) *ListCityMapRangeStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBody) SetPageSize(v int64) *ListCityMapRangeStatisticResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBody) SetPageNumber(v int64) *ListCityMapRangeStatisticResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBody) SetData(v []*ListCityMapRangeStatisticResponseBodyData) *ListCityMapRangeStatisticResponseBody {
	s.Data = v
	return s
}

func (s *ListCityMapRangeStatisticResponseBody) SetCode(v string) *ListCityMapRangeStatisticResponseBody {
	s.Code = &v
	return s
}

type ListCityMapRangeStatisticResponseBodyData struct {
	AdultValue     *string `json:"AdultValue,omitempty" xml:"AdultValue,omitempty"`
	ChildValue     *string `json:"ChildValue,omitempty" xml:"ChildValue,omitempty"`
	OldValue       *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	ManValue       *string `json:"ManValue,omitempty" xml:"ManValue,omitempty"`
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	MotorValue     *string `json:"MotorValue,omitempty" xml:"MotorValue,omitempty"`
	Longitude      *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude       *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	WomanValue     *string `json:"WomanValue,omitempty" xml:"WomanValue,omitempty"`
}

func (s ListCityMapRangeStatisticResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapRangeStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetAdultValue(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.AdultValue = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetChildValue(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.ChildValue = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetOldValue(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.OldValue = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetManValue(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.ManValue = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetDataSourceName(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetDataSourceId(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetCorpId(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetMotorValue(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.MotorValue = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetLongitude(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetLatitude(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListCityMapRangeStatisticResponseBodyData) SetWomanValue(v string) *ListCityMapRangeStatisticResponseBodyData {
	s.WomanValue = &v
	return s
}

type ListCityMapRangeStatisticResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCityMapRangeStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCityMapRangeStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapRangeStatisticResponse) GoString() string {
	return s.String()
}

func (s *ListCityMapRangeStatisticResponse) SetHeaders(v map[string]*string) *ListCityMapRangeStatisticResponse {
	s.Headers = v
	return s
}

func (s *ListCityMapRangeStatisticResponse) SetBody(v *ListCityMapRangeStatisticResponseBody) *ListCityMapRangeStatisticResponse {
	s.Body = v
	return s
}

type ListStorageStatisticsRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListStorageStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStorageStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListStorageStatisticsRequest) SetCorpId(v string) *ListStorageStatisticsRequest {
	s.CorpId = &v
	return s
}

type ListStorageStatisticsResponseBody struct {
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       []*ListStorageStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListStorageStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStorageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStorageStatisticsResponseBody) SetTotalCount(v int64) *ListStorageStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStorageStatisticsResponseBody) SetMessage(v string) *ListStorageStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStorageStatisticsResponseBody) SetRequestId(v string) *ListStorageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStorageStatisticsResponseBody) SetData(v []*ListStorageStatisticsResponseBodyData) *ListStorageStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListStorageStatisticsResponseBody) SetCode(v string) *ListStorageStatisticsResponseBody {
	s.Code = &v
	return s
}

type ListStorageStatisticsResponseBodyData struct {
	UsedStore   *string `json:"UsedStore,omitempty" xml:"UsedStore,omitempty"`
	UnusedStore *string `json:"UnusedStore,omitempty" xml:"UnusedStore,omitempty"`
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Number      *string `json:"Number,omitempty" xml:"Number,omitempty"`
	TotalStore  *string `json:"TotalStore,omitempty" xml:"TotalStore,omitempty"`
}

func (s ListStorageStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListStorageStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStorageStatisticsResponseBodyData) SetUsedStore(v string) *ListStorageStatisticsResponseBodyData {
	s.UsedStore = &v
	return s
}

func (s *ListStorageStatisticsResponseBodyData) SetUnusedStore(v string) *ListStorageStatisticsResponseBodyData {
	s.UnusedStore = &v
	return s
}

func (s *ListStorageStatisticsResponseBodyData) SetCorpId(v string) *ListStorageStatisticsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListStorageStatisticsResponseBodyData) SetNumber(v string) *ListStorageStatisticsResponseBodyData {
	s.Number = &v
	return s
}

func (s *ListStorageStatisticsResponseBodyData) SetTotalStore(v string) *ListStorageStatisticsResponseBodyData {
	s.TotalStore = &v
	return s
}

type ListStorageStatisticsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStorageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStorageStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStorageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListStorageStatisticsResponse) SetHeaders(v map[string]*string) *ListStorageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListStorageStatisticsResponse) SetBody(v *ListStorageStatisticsResponseBody) *ListStorageStatisticsResponse {
	s.Body = v
	return s
}

type PaginateProjectRequest struct {
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CountTotalNum *bool   `json:"CountTotalNum,omitempty" xml:"CountTotalNum,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	NameLike      *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
}

func (s PaginateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PaginateProjectRequest) GoString() string {
	return s.String()
}

func (s *PaginateProjectRequest) SetPageNumber(v int64) *PaginateProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *PaginateProjectRequest) SetPageSize(v int64) *PaginateProjectRequest {
	s.PageSize = &v
	return s
}

func (s *PaginateProjectRequest) SetCountTotalNum(v bool) *PaginateProjectRequest {
	s.CountTotalNum = &v
	return s
}

func (s *PaginateProjectRequest) SetType(v string) *PaginateProjectRequest {
	s.Type = &v
	return s
}

func (s *PaginateProjectRequest) SetNameLike(v string) *PaginateProjectRequest {
	s.NameLike = &v
	return s
}

type PaginateProjectResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *PaginateProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s PaginateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PaginateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PaginateProjectResponseBody) SetMessage(v string) *PaginateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PaginateProjectResponseBody) SetRequestId(v string) *PaginateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PaginateProjectResponseBody) SetData(v *PaginateProjectResponseBodyData) *PaginateProjectResponseBody {
	s.Data = v
	return s
}

func (s *PaginateProjectResponseBody) SetCode(v string) *PaginateProjectResponseBody {
	s.Code = &v
	return s
}

type PaginateProjectResponseBodyData struct {
	Records    []*PaginateProjectResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalPage  *int32                                    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PaginateProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PaginateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PaginateProjectResponseBodyData) SetRecords(v []*PaginateProjectResponseBodyDataRecords) *PaginateProjectResponseBodyData {
	s.Records = v
	return s
}

func (s *PaginateProjectResponseBodyData) SetTotalPage(v int32) *PaginateProjectResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *PaginateProjectResponseBodyData) SetPageNumber(v int32) *PaginateProjectResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *PaginateProjectResponseBodyData) SetPageSize(v int32) *PaginateProjectResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PaginateProjectResponseBodyData) SetTotalCount(v int32) *PaginateProjectResponseBodyData {
	s.TotalCount = &v
	return s
}

type PaginateProjectResponseBodyDataRecords struct {
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ModifiedTime       *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AggregateSceneCode *string `json:"AggregateSceneCode,omitempty" xml:"AggregateSceneCode,omitempty"`
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Icon               *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CreatedTime        *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
}

func (s PaginateProjectResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s PaginateProjectResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *PaginateProjectResponseBodyDataRecords) SetType(v string) *PaginateProjectResponseBodyDataRecords {
	s.Type = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetModifiedTime(v string) *PaginateProjectResponseBodyDataRecords {
	s.ModifiedTime = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetDescription(v string) *PaginateProjectResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetAggregateSceneCode(v string) *PaginateProjectResponseBodyDataRecords {
	s.AggregateSceneCode = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetCorpId(v string) *PaginateProjectResponseBodyDataRecords {
	s.CorpId = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetUserId(v string) *PaginateProjectResponseBodyDataRecords {
	s.UserId = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetIcon(v string) *PaginateProjectResponseBodyDataRecords {
	s.Icon = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetName(v string) *PaginateProjectResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *PaginateProjectResponseBodyDataRecords) SetCreatedTime(v string) *PaginateProjectResponseBodyDataRecords {
	s.CreatedTime = &v
	return s
}

type PaginateProjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PaginateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PaginateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PaginateProjectResponse) GoString() string {
	return s.String()
}

func (s *PaginateProjectResponse) SetHeaders(v map[string]*string) *PaginateProjectResponse {
	s.Headers = v
	return s
}

func (s *PaginateProjectResponse) SetBody(v *PaginateProjectResponseBody) *PaginateProjectResponse {
	s.Body = v
	return s
}

type ListCityMapCameraStatisticsRequest struct {
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DataSourceIdList *string `json:"DataSourceIdList,omitempty" xml:"DataSourceIdList,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListCityMapCameraStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraStatisticsRequest) SetEndTime(v string) *ListCityMapCameraStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCityMapCameraStatisticsRequest) SetStartTime(v string) *ListCityMapCameraStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCityMapCameraStatisticsRequest) SetDataSourceIdList(v string) *ListCityMapCameraStatisticsRequest {
	s.DataSourceIdList = &v
	return s
}

func (s *ListCityMapCameraStatisticsRequest) SetPageSize(v int64) *ListCityMapCameraStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCityMapCameraStatisticsRequest) SetPageNumber(v int64) *ListCityMapCameraStatisticsRequest {
	s.PageNumber = &v
	return s
}

type ListCityMapCameraStatisticsResponseBody struct {
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListCityMapCameraStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListCityMapCameraStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraStatisticsResponseBody) SetTotalCount(v int64) *ListCityMapCameraStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBody) SetRequestId(v string) *ListCityMapCameraStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBody) SetMessage(v string) *ListCityMapCameraStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBody) SetPageSize(v int64) *ListCityMapCameraStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBody) SetPageNumber(v int64) *ListCityMapCameraStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBody) SetData(v []*ListCityMapCameraStatisticsResponseBodyData) *ListCityMapCameraStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBody) SetCode(v string) *ListCityMapCameraStatisticsResponseBody {
	s.Code = &v
	return s
}

type ListCityMapCameraStatisticsResponseBodyData struct {
	OldValue       *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StatisticTime  *string `json:"StatisticTime,omitempty" xml:"StatisticTime,omitempty"`
	AdultValue     *string `json:"AdultValue,omitempty" xml:"AdultValue,omitempty"`
	ChildValue     *string `json:"ChildValue,omitempty" xml:"ChildValue,omitempty"`
	ManValue       *string `json:"ManValue,omitempty" xml:"ManValue,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	MotorValue     *string `json:"MotorValue,omitempty" xml:"MotorValue,omitempty"`
	Longitude      *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude       *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	WomanValue     *string `json:"WomanValue,omitempty" xml:"WomanValue,omitempty"`
}

func (s ListCityMapCameraStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetOldValue(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.OldValue = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetDataSourceName(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetDataSourceId(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetStatisticTime(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.StatisticTime = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetAdultValue(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.AdultValue = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetChildValue(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.ChildValue = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetManValue(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.ManValue = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetCorpId(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetMotorValue(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.MotorValue = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetLongitude(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetLatitude(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListCityMapCameraStatisticsResponseBodyData) SetWomanValue(v string) *ListCityMapCameraStatisticsResponseBodyData {
	s.WomanValue = &v
	return s
}

type ListCityMapCameraStatisticsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCityMapCameraStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCityMapCameraStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCityMapCameraStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListCityMapCameraStatisticsResponse) SetHeaders(v map[string]*string) *ListCityMapCameraStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListCityMapCameraStatisticsResponse) SetBody(v *ListCityMapCameraStatisticsResponseBody) *ListCityMapCameraStatisticsResponse {
	s.Body = v
	return s
}

type UpdateCdrsMonitorRequest struct {
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

func (s UpdateCdrsMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdrsMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateCdrsMonitorRequest) SetCorpId(v string) *UpdateCdrsMonitorRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetTaskId(v string) *UpdateCdrsMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetRuleName(v string) *UpdateCdrsMonitorRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetDeviceOperateType(v string) *UpdateCdrsMonitorRequest {
	s.DeviceOperateType = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetDeviceList(v string) *UpdateCdrsMonitorRequest {
	s.DeviceList = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetPicOperateType(v string) *UpdateCdrsMonitorRequest {
	s.PicOperateType = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetPicList(v string) *UpdateCdrsMonitorRequest {
	s.PicList = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetAttributeOperateType(v string) *UpdateCdrsMonitorRequest {
	s.AttributeOperateType = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetAttributeName(v string) *UpdateCdrsMonitorRequest {
	s.AttributeName = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetAttributeValueList(v string) *UpdateCdrsMonitorRequest {
	s.AttributeValueList = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetDescription(v string) *UpdateCdrsMonitorRequest {
	s.Description = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetRuleExpression(v string) *UpdateCdrsMonitorRequest {
	s.RuleExpression = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetAlgorithmVendor(v string) *UpdateCdrsMonitorRequest {
	s.AlgorithmVendor = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetNotifierType(v string) *UpdateCdrsMonitorRequest {
	s.NotifierType = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetNotifierUrl(v string) *UpdateCdrsMonitorRequest {
	s.NotifierUrl = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetNotifierAppSecret(v string) *UpdateCdrsMonitorRequest {
	s.NotifierAppSecret = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetNotifierTimeOut(v int32) *UpdateCdrsMonitorRequest {
	s.NotifierTimeOut = &v
	return s
}

func (s *UpdateCdrsMonitorRequest) SetNotifierExtendValues(v string) *UpdateCdrsMonitorRequest {
	s.NotifierExtendValues = &v
	return s
}

type UpdateCdrsMonitorResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateCdrsMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdrsMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCdrsMonitorResponseBody) SetMessage(v string) *UpdateCdrsMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCdrsMonitorResponseBody) SetRequestId(v string) *UpdateCdrsMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCdrsMonitorResponseBody) SetData(v string) *UpdateCdrsMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCdrsMonitorResponseBody) SetCode(v string) *UpdateCdrsMonitorResponseBody {
	s.Code = &v
	return s
}

type UpdateCdrsMonitorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCdrsMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCdrsMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdrsMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateCdrsMonitorResponse) SetHeaders(v map[string]*string) *UpdateCdrsMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateCdrsMonitorResponse) SetBody(v *UpdateCdrsMonitorResponseBody) *UpdateCdrsMonitorResponse {
	s.Body = v
	return s
}

type ListPersonResultRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Profession   *string `json:"Profession,omitempty" xml:"Profession,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QualityScore *string `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
}

func (s ListPersonResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonResultRequest) GoString() string {
	return s.String()
}

func (s *ListPersonResultRequest) SetCorpId(v string) *ListPersonResultRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonResultRequest) SetAge(v string) *ListPersonResultRequest {
	s.Age = &v
	return s
}

func (s *ListPersonResultRequest) SetGender(v string) *ListPersonResultRequest {
	s.Gender = &v
	return s
}

func (s *ListPersonResultRequest) SetProfession(v string) *ListPersonResultRequest {
	s.Profession = &v
	return s
}

func (s *ListPersonResultRequest) SetStartTime(v string) *ListPersonResultRequest {
	s.StartTime = &v
	return s
}

func (s *ListPersonResultRequest) SetEndTime(v string) *ListPersonResultRequest {
	s.EndTime = &v
	return s
}

func (s *ListPersonResultRequest) SetPageNumber(v int64) *ListPersonResultRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonResultRequest) SetPageSize(v int64) *ListPersonResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonResultRequest) SetQualityScore(v string) *ListPersonResultRequest {
	s.QualityScore = &v
	return s
}

type ListPersonResultResponseBody struct {
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonResultResponseBody) SetTotalCount(v int64) *ListPersonResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonResultResponseBody) SetRequestId(v string) *ListPersonResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonResultResponseBody) SetMessage(v string) *ListPersonResultResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonResultResponseBody) SetPageSize(v int64) *ListPersonResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonResultResponseBody) SetPageNumber(v int64) *ListPersonResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonResultResponseBody) SetData(v []*ListPersonResultResponseBodyData) *ListPersonResultResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonResultResponseBody) SetCode(v string) *ListPersonResultResponseBody {
	s.Code = &v
	return s
}

type ListPersonResultResponseBodyData struct {
	SourceUrl      *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	Profession     *string `json:"Profession,omitempty" xml:"Profession,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Gender         *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	TargetUrl      *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	Address        *string `json:"Address,omitempty" xml:"Address,omitempty"`
	HotSpotAddress *string `json:"HotSpotAddress,omitempty" xml:"HotSpotAddress,omitempty"`
	Age            *string `json:"Age,omitempty" xml:"Age,omitempty"`
	PersonId       *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	PersonType     *string `json:"PersonType,omitempty" xml:"PersonType,omitempty"`
	Transportation *string `json:"Transportation,omitempty" xml:"Transportation,omitempty"`
}

func (s ListPersonResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonResultResponseBodyData) SetSourceUrl(v string) *ListPersonResultResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetProfession(v string) *ListPersonResultResponseBodyData {
	s.Profession = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetUpdateTime(v string) *ListPersonResultResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetGender(v string) *ListPersonResultResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetTargetUrl(v string) *ListPersonResultResponseBodyData {
	s.TargetUrl = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetAddress(v string) *ListPersonResultResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetHotSpotAddress(v string) *ListPersonResultResponseBodyData {
	s.HotSpotAddress = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetAge(v string) *ListPersonResultResponseBodyData {
	s.Age = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetPersonId(v string) *ListPersonResultResponseBodyData {
	s.PersonId = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetPersonType(v string) *ListPersonResultResponseBodyData {
	s.PersonType = &v
	return s
}

func (s *ListPersonResultResponseBodyData) SetTransportation(v string) *ListPersonResultResponseBodyData {
	s.Transportation = &v
	return s
}

type ListPersonResultResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonResultResponse) GoString() string {
	return s.String()
}

func (s *ListPersonResultResponse) SetHeaders(v map[string]*string) *ListPersonResultResponse {
	s.Headers = v
	return s
}

func (s *ListPersonResultResponse) SetBody(v *ListPersonResultResponseBody) *ListPersonResultResponse {
	s.Body = v
	return s
}

type ListTagMetricsRequest struct {
	CorpId        *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode       map[string]interface{} `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	AggregateType *string                `json:"AggregateType,omitempty" xml:"AggregateType,omitempty"`
	StartTime     *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *string                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTagMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListTagMetricsRequest) SetCorpId(v string) *ListTagMetricsRequest {
	s.CorpId = &v
	return s
}

func (s *ListTagMetricsRequest) SetTagCode(v map[string]interface{}) *ListTagMetricsRequest {
	s.TagCode = v
	return s
}

func (s *ListTagMetricsRequest) SetAggregateType(v string) *ListTagMetricsRequest {
	s.AggregateType = &v
	return s
}

func (s *ListTagMetricsRequest) SetStartTime(v string) *ListTagMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTagMetricsRequest) SetEndTime(v string) *ListTagMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTagMetricsRequest) SetPageNumber(v string) *ListTagMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagMetricsRequest) SetPageSize(v string) *ListTagMetricsRequest {
	s.PageSize = &v
	return s
}

type ListTagMetricsShrinkRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCodeShrink *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	AggregateType *string `json:"AggregateType,omitempty" xml:"AggregateType,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTagMetricsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagMetricsShrinkRequest) SetCorpId(v string) *ListTagMetricsShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *ListTagMetricsShrinkRequest) SetTagCodeShrink(v string) *ListTagMetricsShrinkRequest {
	s.TagCodeShrink = &v
	return s
}

func (s *ListTagMetricsShrinkRequest) SetAggregateType(v string) *ListTagMetricsShrinkRequest {
	s.AggregateType = &v
	return s
}

func (s *ListTagMetricsShrinkRequest) SetStartTime(v string) *ListTagMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListTagMetricsShrinkRequest) SetEndTime(v string) *ListTagMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListTagMetricsShrinkRequest) SetPageNumber(v string) *ListTagMetricsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagMetricsShrinkRequest) SetPageSize(v string) *ListTagMetricsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListTagMetricsResponseBody struct {
	TotalCount *string                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListTagMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTagMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagMetricsResponseBody) SetTotalCount(v string) *ListTagMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagMetricsResponseBody) SetRequestId(v string) *ListTagMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagMetricsResponseBody) SetMessage(v string) *ListTagMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagMetricsResponseBody) SetPageSize(v string) *ListTagMetricsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagMetricsResponseBody) SetPageNumber(v string) *ListTagMetricsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagMetricsResponseBody) SetData(v []*ListTagMetricsResponseBodyData) *ListTagMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ListTagMetricsResponseBody) SetCode(v string) *ListTagMetricsResponseBody {
	s.Code = &v
	return s
}

type ListTagMetricsResponseBodyData struct {
	TagMetric *string `json:"TagMetric,omitempty" xml:"TagMetric,omitempty"`
	TagCode   *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagValue  *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	DateTime  *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s ListTagMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTagMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagMetricsResponseBodyData) SetTagMetric(v string) *ListTagMetricsResponseBodyData {
	s.TagMetric = &v
	return s
}

func (s *ListTagMetricsResponseBodyData) SetTagCode(v string) *ListTagMetricsResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *ListTagMetricsResponseBodyData) SetCorpId(v string) *ListTagMetricsResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListTagMetricsResponseBodyData) SetTagValue(v string) *ListTagMetricsResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListTagMetricsResponseBodyData) SetDateTime(v string) *ListTagMetricsResponseBodyData {
	s.DateTime = &v
	return s
}

type ListTagMetricsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListTagMetricsResponse) SetHeaders(v map[string]*string) *ListTagMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListTagMetricsResponse) SetBody(v *ListTagMetricsResponseBody) *ListTagMetricsResponse {
	s.Body = v
	return s
}

type ListPersonTagRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPersonTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTagRequest) GoString() string {
	return s.String()
}

func (s *ListPersonTagRequest) SetCorpId(v string) *ListPersonTagRequest {
	s.CorpId = &v
	return s
}

func (s *ListPersonTagRequest) SetTagCode(v string) *ListPersonTagRequest {
	s.TagCode = &v
	return s
}

func (s *ListPersonTagRequest) SetPageNumber(v int64) *ListPersonTagRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTagRequest) SetPageSize(v int64) *ListPersonTagRequest {
	s.PageSize = &v
	return s
}

type ListPersonTagResponseBody struct {
	TotalCount *int64                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListPersonTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPersonTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonTagResponseBody) SetTotalCount(v int64) *ListPersonTagResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonTagResponseBody) SetRequestId(v string) *ListPersonTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonTagResponseBody) SetMessage(v string) *ListPersonTagResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonTagResponseBody) SetPageSize(v int64) *ListPersonTagResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonTagResponseBody) SetPageNumber(v int64) *ListPersonTagResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTagResponseBody) SetData(v []*ListPersonTagResponseBodyData) *ListPersonTagResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonTagResponseBody) SetCode(v string) *ListPersonTagResponseBody {
	s.Code = &v
	return s
}

type ListPersonTagResponseBodyData struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPersonTagResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonTagResponseBodyData) SetValue(v string) *ListPersonTagResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListPersonTagResponseBodyData) SetCorpId(v string) *ListPersonTagResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListPersonTagResponseBodyData) SetTagValue(v string) *ListPersonTagResponseBodyData {
	s.TagValue = &v
	return s
}

type ListPersonTagResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTagResponse) GoString() string {
	return s.String()
}

func (s *ListPersonTagResponse) SetHeaders(v map[string]*string) *ListPersonTagResponse {
	s.Headers = v
	return s
}

func (s *ListPersonTagResponse) SetBody(v *ListPersonTagResponseBody) *ListPersonTagResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Icon               *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AggregateSceneCode *string `json:"AggregateSceneCode,omitempty" xml:"AggregateSceneCode,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetCorpId(v string) *UpdateProjectRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateProjectRequest) SetIcon(v string) *UpdateProjectRequest {
	s.Icon = &v
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

func (s *UpdateProjectRequest) SetAggregateSceneCode(v string) *UpdateProjectRequest {
	s.AggregateSceneCode = &v
	return s
}

type UpdateProjectResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetMessage(v string) *UpdateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetCode(v string) *UpdateProjectResponseBody {
	s.Code = &v
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

type ListDevicePersonRequest struct {
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber     *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListDevicePersonRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonRequest) GoString() string {
	return s.String()
}

func (s *ListDevicePersonRequest) SetDataSourceId(v string) *ListDevicePersonRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDevicePersonRequest) SetStatisticsType(v string) *ListDevicePersonRequest {
	s.StatisticsType = &v
	return s
}

func (s *ListDevicePersonRequest) SetStartTime(v string) *ListDevicePersonRequest {
	s.StartTime = &v
	return s
}

func (s *ListDevicePersonRequest) SetEndTime(v string) *ListDevicePersonRequest {
	s.EndTime = &v
	return s
}

func (s *ListDevicePersonRequest) SetPageNumber(v int64) *ListDevicePersonRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevicePersonRequest) SetPageSize(v int64) *ListDevicePersonRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevicePersonRequest) SetCorpId(v string) *ListDevicePersonRequest {
	s.CorpId = &v
	return s
}

type ListDevicePersonResponseBody struct {
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListDevicePersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDevicePersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicePersonResponseBody) SetTotalCount(v int64) *ListDevicePersonResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDevicePersonResponseBody) SetRequestId(v string) *ListDevicePersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicePersonResponseBody) SetMessage(v string) *ListDevicePersonResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicePersonResponseBody) SetPageSize(v int64) *ListDevicePersonResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDevicePersonResponseBody) SetPageNumber(v int64) *ListDevicePersonResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDevicePersonResponseBody) SetData(v []*ListDevicePersonResponseBodyData) *ListDevicePersonResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicePersonResponseBody) SetCode(v string) *ListDevicePersonResponseBody {
	s.Code = &v
	return s
}

type ListDevicePersonResponseBodyData struct {
	TargetPicUrlPath *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty"`
	Gender           *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	FreqNum          *string `json:"FreqNum,omitempty" xml:"FreqNum,omitempty"`
	PersonId         *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s ListDevicePersonResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicePersonResponseBodyData) SetTargetPicUrlPath(v string) *ListDevicePersonResponseBodyData {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListDevicePersonResponseBodyData) SetGender(v string) *ListDevicePersonResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListDevicePersonResponseBodyData) SetDataSourceId(v string) *ListDevicePersonResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListDevicePersonResponseBodyData) SetFreqNum(v string) *ListDevicePersonResponseBodyData {
	s.FreqNum = &v
	return s
}

func (s *ListDevicePersonResponseBodyData) SetPersonId(v string) *ListDevicePersonResponseBodyData {
	s.PersonId = &v
	return s
}

type ListDevicePersonResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicePersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicePersonResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePersonResponse) GoString() string {
	return s.String()
}

func (s *ListDevicePersonResponse) SetHeaders(v map[string]*string) *ListDevicePersonResponse {
	s.Headers = v
	return s
}

func (s *ListDevicePersonResponse) SetBody(v *ListDevicePersonResponseBody) *ListDevicePersonResponse {
	s.Body = v
	return s
}

type ListDeviceDetailRequest struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListDeviceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceDetailRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceDetailRequest) SetDataSourceId(v string) *ListDeviceDetailRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDeviceDetailRequest) SetPageNumber(v int64) *ListDeviceDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeviceDetailRequest) SetPageSize(v int64) *ListDeviceDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceDetailRequest) SetCorpId(v string) *ListDeviceDetailRequest {
	s.CorpId = &v
	return s
}

type ListDeviceDetailResponseBody struct {
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Data       []*ListDeviceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDeviceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceDetailResponseBody) SetTotalCount(v int64) *ListDeviceDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceDetailResponseBody) SetRequestId(v string) *ListDeviceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceDetailResponseBody) SetMessage(v string) *ListDeviceDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceDetailResponseBody) SetPageSize(v int64) *ListDeviceDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeviceDetailResponseBody) SetPageNumber(v int64) *ListDeviceDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeviceDetailResponseBody) SetData(v []*ListDeviceDetailResponseBodyData) *ListDeviceDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceDetailResponseBody) SetCode(v string) *ListDeviceDetailResponseBody {
	s.Code = &v
	return s
}

type ListDeviceDetailResponseBodyData struct {
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	Longitude      *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude       *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	DataSourcePoi  *string `json:"DataSourcePoi,omitempty" xml:"DataSourcePoi,omitempty"`
	NearPoi        *string `json:"NearPoi,omitempty" xml:"NearPoi,omitempty"`
}

func (s ListDeviceDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceDetailResponseBodyData) SetDataSourceName(v string) *ListDeviceDetailResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListDeviceDetailResponseBodyData) SetDataSourceId(v string) *ListDeviceDetailResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListDeviceDetailResponseBodyData) SetCorpId(v string) *ListDeviceDetailResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListDeviceDetailResponseBodyData) SetLongitude(v string) *ListDeviceDetailResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListDeviceDetailResponseBodyData) SetLatitude(v string) *ListDeviceDetailResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListDeviceDetailResponseBodyData) SetDataSourcePoi(v string) *ListDeviceDetailResponseBodyData {
	s.DataSourcePoi = &v
	return s
}

func (s *ListDeviceDetailResponseBodyData) SetNearPoi(v string) *ListDeviceDetailResponseBodyData {
	s.NearPoi = &v
	return s
}

type ListDeviceDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceDetailResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceDetailResponse) SetHeaders(v map[string]*string) *ListDeviceDetailResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceDetailResponse) SetBody(v *ListDeviceDetailResponseBody) *ListDeviceDetailResponse {
	s.Body = v
	return s
}

type ListDeviceGenderStatisticsRequest struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
}

func (s ListDeviceGenderStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGenderStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceGenderStatisticsRequest) SetDataSourceId(v string) *ListDeviceGenderStatisticsRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDeviceGenderStatisticsRequest) SetStartTime(v string) *ListDeviceGenderStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDeviceGenderStatisticsRequest) SetEndTime(v string) *ListDeviceGenderStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDeviceGenderStatisticsRequest) SetCorpId(v string) *ListDeviceGenderStatisticsRequest {
	s.CorpId = &v
	return s
}

type ListDeviceGenderStatisticsResponseBody struct {
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       []*ListDeviceGenderStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDeviceGenderStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGenderStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceGenderStatisticsResponseBody) SetTotalCount(v int64) *ListDeviceGenderStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceGenderStatisticsResponseBody) SetMessage(v string) *ListDeviceGenderStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceGenderStatisticsResponseBody) SetRequestId(v string) *ListDeviceGenderStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGenderStatisticsResponseBody) SetData(v []*ListDeviceGenderStatisticsResponseBodyData) *ListDeviceGenderStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceGenderStatisticsResponseBody) SetCode(v string) *ListDeviceGenderStatisticsResponseBody {
	s.Code = &v
	return s
}

type ListDeviceGenderStatisticsResponseBodyData struct {
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	Number       *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ListDeviceGenderStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGenderStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceGenderStatisticsResponseBodyData) SetGender(v string) *ListDeviceGenderStatisticsResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ListDeviceGenderStatisticsResponseBodyData) SetDataSourceId(v string) *ListDeviceGenderStatisticsResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *ListDeviceGenderStatisticsResponseBodyData) SetNumber(v string) *ListDeviceGenderStatisticsResponseBodyData {
	s.Number = &v
	return s
}

type ListDeviceGenderStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceGenderStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceGenderStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGenderStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGenderStatisticsResponse) SetHeaders(v map[string]*string) *ListDeviceGenderStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceGenderStatisticsResponse) SetBody(v *ListDeviceGenderStatisticsResponseBody) *ListDeviceGenderStatisticsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cdrs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) SearchObjectWithOptions(request *SearchObjectRequest, runtime *util.RuntimeOptions) (_result *SearchObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchObject"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListAreaHotSpotMetricsWithOptions(request *ListAreaHotSpotMetricsRequest, runtime *util.RuntimeOptions) (_result *ListAreaHotSpotMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAreaHotSpotMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAreaHotSpotMetrics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAreaHotSpotMetrics(request *ListAreaHotSpotMetricsRequest) (_result *ListAreaHotSpotMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAreaHotSpotMetricsResponse{}
	_body, _err := client.ListAreaHotSpotMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindDeviceWithOptions(request *BindDeviceRequest, runtime *util.RuntimeOptions) (_result *BindDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindDevice"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindDevice(request *BindDeviceRequest) (_result *BindDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindDeviceResponse{}
	_body, _err := client.BindDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCdrsMonitorResultWithOptions(request *GetCdrsMonitorResultRequest, runtime *util.RuntimeOptions) (_result *GetCdrsMonitorResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCdrsMonitorResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCdrsMonitorResult"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCdrsMonitorResult(request *GetCdrsMonitorResultRequest) (_result *GetCdrsMonitorResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCdrsMonitorResultResponse{}
	_body, _err := client.GetCdrsMonitorResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVehicleDetailsWithOptions(request *ListVehicleDetailsRequest, runtime *util.RuntimeOptions) (_result *ListVehicleDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVehicleDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVehicleDetails"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVehicleDetails(request *ListVehicleDetailsRequest) (_result *ListVehicleDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVehicleDetailsResponse{}
	_body, _err := client.ListVehicleDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCdrsMonitorListWithOptions(request *GetCdrsMonitorListRequest, runtime *util.RuntimeOptions) (_result *GetCdrsMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCdrsMonitorListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCdrsMonitorList"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCdrsMonitorList(request *GetCdrsMonitorListRequest) (_result *GetCdrsMonitorListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCdrsMonitorListResponse{}
	_body, _err := client.GetCdrsMonitorListWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UpdateMonitor"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDataStatisticsWithOptions(request *ListDataStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListDataStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDataStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataStatistics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataStatistics(request *ListDataStatisticsRequest) (_result *ListDataStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataStatisticsResponse{}
	_body, _err := client.ListDataStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDeviceWithOptions(request *UnbindDeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindDevice"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDevice(request *UnbindDeviceRequest) (_result *UnbindDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.UnbindDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonDetailsWithOptions(request *ListPersonDetailsRequest, runtime *util.RuntimeOptions) (_result *ListPersonDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonDetails"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonDetails(request *ListPersonDetailsRequest) (_result *ListPersonDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonDetailsResponse{}
	_body, _err := client.ListPersonDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVehicleTagDistributeWithOptions(request *ListVehicleTagDistributeRequest, runtime *util.RuntimeOptions) (_result *ListVehicleTagDistributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVehicleTagDistributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVehicleTagDistribute"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVehicleTagDistribute(request *ListVehicleTagDistributeRequest) (_result *ListVehicleTagDistributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVehicleTagDistributeResponse{}
	_body, _err := client.ListVehicleTagDistributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevicePersonStatisticsWithOptions(request *ListDevicePersonStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListDevicePersonStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevicePersonStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevicePersonStatistics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevicePersonStatistics(request *ListDevicePersonStatisticsRequest) (_result *ListDevicePersonStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicePersonStatisticsResponse{}
	_body, _err := client.ListDevicePersonStatisticsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("AddMonitor"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PaginateDeviceWithOptions(request *PaginateDeviceRequest, runtime *util.RuntimeOptions) (_result *PaginateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PaginateDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PaginateDevice"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PaginateDevice(request *PaginateDeviceRequest) (_result *PaginateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PaginateDeviceResponse{}
	_body, _err := client.PaginateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCdrsMonitorWithOptions(request *StopCdrsMonitorRequest, runtime *util.RuntimeOptions) (_result *StopCdrsMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopCdrsMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopCdrsMonitor"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCdrsMonitor(request *StopCdrsMonitorRequest) (_result *StopCdrsMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCdrsMonitorResponse{}
	_body, _err := client.StopCdrsMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallTrajectoryByCoordinateTimeWithOptions(request *RecallTrajectoryByCoordinateTimeRequest, runtime *util.RuntimeOptions) (_result *RecallTrajectoryByCoordinateTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecallTrajectoryByCoordinateTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecallTrajectoryByCoordinateTime"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallTrajectoryByCoordinateTime(request *RecallTrajectoryByCoordinateTimeRequest) (_result *RecallTrajectoryByCoordinateTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecallTrajectoryByCoordinateTimeResponse{}
	_body, _err := client.RecallTrajectoryByCoordinateTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCityMapPersonFlowWithOptions(tmpReq *ListCityMapPersonFlowRequest, runtime *util.RuntimeOptions) (_result *ListCityMapPersonFlowResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCityMapPersonFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OriginDataSourceIdList)) {
		request.OriginDataSourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OriginDataSourceIdList, tea.String("OriginDataSourceIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TargetDataSourceIdList)) {
		request.TargetDataSourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetDataSourceIdList, tea.String("TargetDataSourceIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCityMapPersonFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCityMapPersonFlow"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCityMapPersonFlow(request *ListCityMapPersonFlowRequest) (_result *ListCityMapPersonFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCityMapPersonFlowResponse{}
	_body, _err := client.ListCityMapPersonFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCdrsMonitorWithOptions(request *AddCdrsMonitorRequest, runtime *util.RuntimeOptions) (_result *AddCdrsMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddCdrsMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddCdrsMonitor"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCdrsMonitor(request *AddCdrsMonitorRequest) (_result *AddCdrsMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCdrsMonitorResponse{}
	_body, _err := client.AddCdrsMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMapRouteDetailsWithOptions(tmpReq *ListMapRouteDetailsRequest, runtime *util.RuntimeOptions) (_result *ListMapRouteDetailsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMapRouteDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RouteList)) {
		request.RouteListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RouteList, tea.String("RouteList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMapRouteDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMapRouteDetails"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMapRouteDetails(request *ListMapRouteDetailsRequest) (_result *ListMapRouteDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMapRouteDetailsResponse{}
	_body, _err := client.ListMapRouteDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonTopWithOptions(request *ListPersonTopRequest, runtime *util.RuntimeOptions) (_result *ListPersonTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonTopResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonTop"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonTop(request *ListPersonTopRequest) (_result *ListPersonTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonTopResponse{}
	_body, _err := client.ListPersonTopWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("GetMonitorResult"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListCityMapAoisWithOptions(request *ListCityMapAoisRequest, runtime *util.RuntimeOptions) (_result *ListCityMapAoisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCityMapAoisResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCityMapAois"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCityMapAois(request *ListCityMapAoisRequest) (_result *ListCityMapAoisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCityMapAoisResponse{}
	_body, _err := client.ListCityMapAoisWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("RecognizeImage"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetMonitorListWithOptions(request *GetMonitorListRequest, runtime *util.RuntimeOptions) (_result *GetMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMonitorListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMonitorList"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDeviceRelationWithOptions(request *ListDeviceRelationRequest, runtime *util.RuntimeOptions) (_result *ListDeviceRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeviceRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceRelation"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceRelation(request *ListDeviceRelationRequest) (_result *ListDeviceRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceRelationResponse{}
	_body, _err := client.ListDeviceRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonTrackWithOptions(request *ListPersonTrackRequest, runtime *util.RuntimeOptions) (_result *ListPersonTrackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonTrackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonTrack"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonTrack(request *ListPersonTrackRequest) (_result *ListPersonTrackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonTrackResponse{}
	_body, _err := client.ListPersonTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCityMapCameraResultsWithOptions(tmpReq *ListCityMapCameraResultsRequest, runtime *util.RuntimeOptions) (_result *ListCityMapCameraResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCityMapCameraResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DataSourceIdList)) {
		request.DataSourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceIdList, tea.String("DataSourceIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCityMapCameraResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCityMapCameraResults"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCityMapCameraResults(request *ListCityMapCameraResultsRequest) (_result *ListCityMapCameraResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCityMapCameraResultsResponse{}
	_body, _err := client.ListCityMapCameraResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTrajectoryByIdWithOptions(request *QueryTrajectoryByIdRequest, runtime *util.RuntimeOptions) (_result *QueryTrajectoryByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTrajectoryByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTrajectoryById"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTrajectoryById(request *QueryTrajectoryByIdRequest) (_result *QueryTrajectoryByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTrajectoryByIdResponse{}
	_body, _err := client.QueryTrajectoryByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCityMapImageDetailsWithOptions(request *ListCityMapImageDetailsRequest, runtime *util.RuntimeOptions) (_result *ListCityMapImageDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCityMapImageDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCityMapImageDetails"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCityMapImageDetails(request *ListCityMapImageDetailsRequest) (_result *ListCityMapImageDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCityMapImageDetailsResponse{}
	_body, _err := client.ListCityMapImageDetailsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListVehicleTopWithOptions(request *ListVehicleTopRequest, runtime *util.RuntimeOptions) (_result *ListVehicleTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVehicleTopResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVehicleTop"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVehicleTop(request *ListVehicleTopRequest) (_result *ListVehicleTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVehicleTopResponse{}
	_body, _err := client.ListVehicleTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataStatisticsByDayWithOptions(request *ListDataStatisticsByDayRequest, runtime *util.RuntimeOptions) (_result *ListDataStatisticsByDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDataStatisticsByDayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataStatisticsByDay"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataStatisticsByDay(request *ListDataStatisticsByDayRequest) (_result *ListDataStatisticsByDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataStatisticsByDayResponse{}
	_body, _err := client.ListDataStatisticsByDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVehicleResultsWithOptions(request *ListVehicleResultsRequest, runtime *util.RuntimeOptions) (_result *ListVehicleResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVehicleResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVehicleResults"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVehicleResults(request *ListVehicleResultsRequest) (_result *ListVehicleResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVehicleResultsResponse{}
	_body, _err := client.ListVehicleResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAggregateObjectWithOptions(request *SearchAggregateObjectRequest, runtime *util.RuntimeOptions) (_result *SearchAggregateObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchAggregateObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchAggregateObject"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAggregateObject(request *SearchAggregateObjectRequest) (_result *SearchAggregateObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAggregateObjectResponse{}
	_body, _err := client.SearchAggregateObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorpMetricsStatisticWithOptions(tmpReq *ListCorpMetricsStatisticRequest, runtime *util.RuntimeOptions) (_result *ListCorpMetricsStatisticResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCorpMetricsStatisticShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserGroupList)) {
		request.UserGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupList, tea.String("UserGroupList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceGroupList)) {
		request.DeviceGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceGroupList, tea.String("DeviceGroupList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceIdList)) {
		request.DeviceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceIdList, tea.String("DeviceIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCorpMetricsStatisticResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCorpMetricsStatistic"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorpMetricsStatistic(request *ListCorpMetricsStatisticRequest) (_result *ListCorpMetricsStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorpMetricsStatisticResponse{}
	_body, _err := client.ListCorpMetricsStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectTrajectoryRegularPatternWithOptions(request *DetectTrajectoryRegularPatternRequest, runtime *util.RuntimeOptions) (_result *DetectTrajectoryRegularPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectTrajectoryRegularPatternResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectTrajectoryRegularPattern"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectTrajectoryRegularPattern(request *DetectTrajectoryRegularPatternRequest) (_result *DetectTrajectoryRegularPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectTrajectoryRegularPatternResponse{}
	_body, _err := client.DetectTrajectoryRegularPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVehicleTrackWithOptions(request *ListVehicleTrackRequest, runtime *util.RuntimeOptions) (_result *ListVehicleTrackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVehicleTrackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVehicleTrack"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVehicleTrack(request *ListVehicleTrackRequest) (_result *ListVehicleTrackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVehicleTrackResponse{}
	_body, _err := client.ListVehicleTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStructureStatisticsWithOptions(request *ListStructureStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListStructureStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStructureStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStructureStatistics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStructureStatistics(request *ListStructureStatisticsRequest) (_result *ListStructureStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStructureStatisticsResponse{}
	_body, _err := client.ListStructureStatisticsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("StopMonitor"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PredictTrajectoryDestinationWithOptions(request *PredictTrajectoryDestinationRequest, runtime *util.RuntimeOptions) (_result *PredictTrajectoryDestinationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PredictTrajectoryDestinationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PredictTrajectoryDestination"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictTrajectoryDestination(request *PredictTrajectoryDestinationRequest) (_result *PredictTrajectoryDestinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictTrajectoryDestinationResponse{}
	_body, _err := client.PredictTrajectoryDestinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRangeDeviceWithOptions(request *ListRangeDeviceRequest, runtime *util.RuntimeOptions) (_result *ListRangeDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRangeDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRangeDevice"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRangeDevice(request *ListRangeDeviceRequest) (_result *ListRangeDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRangeDeviceResponse{}
	_body, _err := client.ListRangeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCityMapRangeStatisticWithOptions(request *ListCityMapRangeStatisticRequest, runtime *util.RuntimeOptions) (_result *ListCityMapRangeStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCityMapRangeStatisticResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCityMapRangeStatistic"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCityMapRangeStatistic(request *ListCityMapRangeStatisticRequest) (_result *ListCityMapRangeStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCityMapRangeStatisticResponse{}
	_body, _err := client.ListCityMapRangeStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStorageStatisticsWithOptions(request *ListStorageStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListStorageStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStorageStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStorageStatistics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStorageStatistics(request *ListStorageStatisticsRequest) (_result *ListStorageStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStorageStatisticsResponse{}
	_body, _err := client.ListStorageStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PaginateProjectWithOptions(request *PaginateProjectRequest, runtime *util.RuntimeOptions) (_result *PaginateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PaginateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PaginateProject"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PaginateProject(request *PaginateProjectRequest) (_result *PaginateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PaginateProjectResponse{}
	_body, _err := client.PaginateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCityMapCameraStatisticsWithOptions(request *ListCityMapCameraStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListCityMapCameraStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCityMapCameraStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCityMapCameraStatistics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCityMapCameraStatistics(request *ListCityMapCameraStatisticsRequest) (_result *ListCityMapCameraStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCityMapCameraStatisticsResponse{}
	_body, _err := client.ListCityMapCameraStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCdrsMonitorWithOptions(request *UpdateCdrsMonitorRequest, runtime *util.RuntimeOptions) (_result *UpdateCdrsMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCdrsMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCdrsMonitor"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCdrsMonitor(request *UpdateCdrsMonitorRequest) (_result *UpdateCdrsMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCdrsMonitorResponse{}
	_body, _err := client.UpdateCdrsMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonResultWithOptions(request *ListPersonResultRequest, runtime *util.RuntimeOptions) (_result *ListPersonResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonResult"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonResult(request *ListPersonResultRequest) (_result *ListPersonResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonResultResponse{}
	_body, _err := client.ListPersonResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagMetricsWithOptions(tmpReq *ListTagMetricsRequest, runtime *util.RuntimeOptions) (_result *ListTagMetricsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TagCode)) {
		request.TagCodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagCode, tea.String("TagCode"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagMetrics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagMetrics(request *ListTagMetricsRequest) (_result *ListTagMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagMetricsResponse{}
	_body, _err := client.ListTagMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonTagWithOptions(request *ListPersonTagRequest, runtime *util.RuntimeOptions) (_result *ListPersonTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonTag"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonTag(request *ListPersonTagRequest) (_result *ListPersonTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonTagResponse{}
	_body, _err := client.ListPersonTagWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDevicePersonWithOptions(request *ListDevicePersonRequest, runtime *util.RuntimeOptions) (_result *ListDevicePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevicePersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevicePerson"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevicePerson(request *ListDevicePersonRequest) (_result *ListDevicePersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicePersonResponse{}
	_body, _err := client.ListDevicePersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceDetailWithOptions(request *ListDeviceDetailRequest, runtime *util.RuntimeOptions) (_result *ListDeviceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeviceDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceDetail"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceDetail(request *ListDeviceDetailRequest) (_result *ListDeviceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceDetailResponse{}
	_body, _err := client.ListDeviceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceGenderStatisticsWithOptions(request *ListDeviceGenderStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceGenderStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeviceGenderStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceGenderStatistics"), tea.String("2020-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceGenderStatistics(request *ListDeviceGenderStatisticsRequest) (_result *ListDeviceGenderStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceGenderStatisticsResponse{}
	_body, _err := client.ListDeviceGenderStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
