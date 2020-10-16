// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetMonitorListRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	PageNo   *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *GetMonitorListRequest) SetPageNo(v int) *GetMonitorListRequest {
	s.PageNo = &v
	return s
}

func (s *GetMonitorListRequest) SetPageSize(v int) *GetMonitorListRequest {
	s.PageSize = &v
	return s
}

type GetMonitorListResponse struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetMonitorListResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetMonitorListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponse) SetCode(v string) *GetMonitorListResponse {
	s.Code = &v
	return s
}

func (s *GetMonitorListResponse) SetMessage(v string) *GetMonitorListResponse {
	s.Message = &v
	return s
}

func (s *GetMonitorListResponse) SetRequestId(v string) *GetMonitorListResponse {
	s.RequestId = &v
	return s
}

func (s *GetMonitorListResponse) SetData(v *GetMonitorListResponseData) *GetMonitorListResponse {
	s.Data = v
	return s
}

type GetMonitorListResponseData struct {
	PageNo     *int                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize   *int                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                                 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*GetMonitorListResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s GetMonitorListResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponseData) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponseData) SetPageNo(v int) *GetMonitorListResponseData {
	s.PageNo = &v
	return s
}

func (s *GetMonitorListResponseData) SetPageSize(v int) *GetMonitorListResponseData {
	s.PageSize = &v
	return s
}

func (s *GetMonitorListResponseData) SetTotalCount(v int) *GetMonitorListResponseData {
	s.TotalCount = &v
	return s
}

func (s *GetMonitorListResponseData) SetTotalPage(v int) *GetMonitorListResponseData {
	s.TotalPage = &v
	return s
}

func (s *GetMonitorListResponseData) SetRecords(v []*GetMonitorListResponseDataRecords) *GetMonitorListResponseData {
	s.Records = v
	return s
}

type GetMonitorListResponseDataRecords struct {
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	MonitorType     *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty" require:"true"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty" require:"true"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	ModifiedDate    *string `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty" require:"true"`
	DeviceList      *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" require:"true"`
	Attributes      *string `json:"Attributes,omitempty" xml:"Attributes,omitempty" require:"true"`
	RuleExpression  *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty" require:"true"`
	NotifierType    *string `json:"NotifierType,omitempty" xml:"NotifierType,omitempty" require:"true"`
	NotifierExtra   *string `json:"NotifierExtra,omitempty" xml:"NotifierExtra,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Expression      *string `json:"Expression,omitempty" xml:"Expression,omitempty" require:"true"`
	ImageMatch      *string `json:"ImageMatch,omitempty" xml:"ImageMatch,omitempty" require:"true"`
}

func (s GetMonitorListResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorListResponseDataRecords) GoString() string {
	return s.String()
}

func (s *GetMonitorListResponseDataRecords) SetTaskId(v string) *GetMonitorListResponseDataRecords {
	s.TaskId = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetStatus(v string) *GetMonitorListResponseDataRecords {
	s.Status = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetMonitorType(v string) *GetMonitorListResponseDataRecords {
	s.MonitorType = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetRuleName(v string) *GetMonitorListResponseDataRecords {
	s.RuleName = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetAlgorithmVendor(v string) *GetMonitorListResponseDataRecords {
	s.AlgorithmVendor = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetCreateDate(v string) *GetMonitorListResponseDataRecords {
	s.CreateDate = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetModifiedDate(v string) *GetMonitorListResponseDataRecords {
	s.ModifiedDate = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetDeviceList(v string) *GetMonitorListResponseDataRecords {
	s.DeviceList = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetAttributes(v string) *GetMonitorListResponseDataRecords {
	s.Attributes = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetRuleExpression(v string) *GetMonitorListResponseDataRecords {
	s.RuleExpression = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetNotifierType(v string) *GetMonitorListResponseDataRecords {
	s.NotifierType = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetNotifierExtra(v string) *GetMonitorListResponseDataRecords {
	s.NotifierExtra = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetDescription(v string) *GetMonitorListResponseDataRecords {
	s.Description = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetExpression(v string) *GetMonitorListResponseDataRecords {
	s.Expression = &v
	return s
}

func (s *GetMonitorListResponseDataRecords) SetImageMatch(v string) *GetMonitorListResponseDataRecords {
	s.ImageMatch = &v
	return s
}

type ListDeviceGroupsRequest struct {
	DeviceCodeList *string `json:"DeviceCodeList,omitempty" xml:"DeviceCodeList,omitempty"`
	CorpIdList     *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	IsPage         *int    `json:"IsPage,omitempty" xml:"IsPage,omitempty" require:"true"`
	PageNum        *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Group          *string `json:"Group,omitempty" xml:"Group,omitempty"`
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

func (s *ListDeviceGroupsRequest) SetIsPage(v int) *ListDeviceGroupsRequest {
	s.IsPage = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetPageNum(v int) *ListDeviceGroupsRequest {
	s.PageNum = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetPageSize(v int) *ListDeviceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetGroup(v string) *ListDeviceGroupsRequest {
	s.Group = &v
	return s
}

type ListDeviceGroupsResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      []*ListDeviceGroupsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListDeviceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponse) SetRequestId(v string) *ListDeviceGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGroupsResponse) SetCode(v string) *ListDeviceGroupsResponse {
	s.Code = &v
	return s
}

func (s *ListDeviceGroupsResponse) SetMessage(v string) *ListDeviceGroupsResponse {
	s.Message = &v
	return s
}

func (s *ListDeviceGroupsResponse) SetData(v []*ListDeviceGroupsResponseData) *ListDeviceGroupsResponse {
	s.Data = v
	return s
}

type ListDeviceGroupsResponseData struct {
	TotalCount *string                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	List       []*ListDeviceGroupsResponseDataList `json:"List,omitempty" xml:"List,omitempty" require:"true" type:"Repeated"`
}

func (s ListDeviceGroupsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponseData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseData) SetTotalCount(v string) *ListDeviceGroupsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceGroupsResponseData) SetList(v []*ListDeviceGroupsResponseDataList) *ListDeviceGroupsResponseData {
	s.List = v
	return s
}

type ListDeviceGroupsResponseDataList struct {
	DeviceGroup         *string `json:"DeviceGroup,omitempty" xml:"DeviceGroup,omitempty" require:"true"`
	DeviceName          *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty" require:"true"`
	DeviceCode          *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty" require:"true"`
	BitRate             *string `json:"BitRate,omitempty" xml:"BitRate,omitempty" require:"true"`
	CodingFormat        *string `json:"CodingFormat,omitempty" xml:"CodingFormat,omitempty" require:"true"`
	ResolvingPower      *string `json:"ResolvingPower,omitempty" xml:"ResolvingPower,omitempty" require:"true"`
	DataSourceType      *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty" require:"true"`
	RegionName          *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstallAddress      *string `json:"InstallAddress,omitempty" xml:"InstallAddress,omitempty" require:"true"`
	DeviceSn            *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty" require:"true"`
	DeviceStatus        *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty" require:"true"`
	DeviceStreamStatus  *string `json:"DeviceStreamStatus,omitempty" xml:"DeviceStreamStatus,omitempty" require:"true"`
	DeviceComputeStatus *string `json:"DeviceComputeStatus,omitempty" xml:"DeviceComputeStatus,omitempty" require:"true"`
}

func (s ListDeviceGroupsResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupsResponseDataList) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceGroup(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceGroup = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceName(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceCode(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceCode = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetBitRate(v string) *ListDeviceGroupsResponseDataList {
	s.BitRate = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetCodingFormat(v string) *ListDeviceGroupsResponseDataList {
	s.CodingFormat = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetResolvingPower(v string) *ListDeviceGroupsResponseDataList {
	s.ResolvingPower = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDataSourceType(v string) *ListDeviceGroupsResponseDataList {
	s.DataSourceType = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetRegionName(v string) *ListDeviceGroupsResponseDataList {
	s.RegionName = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetRegionId(v string) *ListDeviceGroupsResponseDataList {
	s.RegionId = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetInstallAddress(v string) *ListDeviceGroupsResponseDataList {
	s.InstallAddress = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceSn(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceSn = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceStatus(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceStatus = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceStreamStatus(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceStreamStatus = &v
	return s
}

func (s *ListDeviceGroupsResponseDataList) SetDeviceComputeStatus(v string) *ListDeviceGroupsResponseDataList {
	s.DeviceComputeStatus = &v
	return s
}

type SearchObjectRequest struct {
	CorpId        *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	ObjectType    *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty" require:"true"`
	StartTime     *int64                 `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber    *int                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize      *int                   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *SearchObjectRequest) SetPageNumber(v int) *SearchObjectRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchObjectRequest) SetPageSize(v int) *SearchObjectRequest {
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
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	ObjectType       *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty" require:"true"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber       *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize         *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *SearchObjectShrinkRequest) SetPageNumber(v int) *SearchObjectShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchObjectShrinkRequest) SetPageSize(v int) *SearchObjectShrinkRequest {
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

type SearchObjectResponse struct {
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SearchObjectResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SearchObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponse) GoString() string {
	return s.String()
}

func (s *SearchObjectResponse) SetCode(v string) *SearchObjectResponse {
	s.Code = &v
	return s
}

func (s *SearchObjectResponse) SetMessage(v string) *SearchObjectResponse {
	s.Message = &v
	return s
}

func (s *SearchObjectResponse) SetRequestId(v string) *SearchObjectResponse {
	s.RequestId = &v
	return s
}

func (s *SearchObjectResponse) SetData(v *SearchObjectResponseData) *SearchObjectResponse {
	s.Data = v
	return s
}

type SearchObjectResponseData struct {
	PageNumber *int                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                               `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*SearchObjectResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s SearchObjectResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseData) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseData) SetPageNumber(v int) *SearchObjectResponseData {
	s.PageNumber = &v
	return s
}

func (s *SearchObjectResponseData) SetPageSize(v int) *SearchObjectResponseData {
	s.PageSize = &v
	return s
}

func (s *SearchObjectResponseData) SetTotalCount(v int) *SearchObjectResponseData {
	s.TotalCount = &v
	return s
}

func (s *SearchObjectResponseData) SetTotalPage(v int) *SearchObjectResponseData {
	s.TotalPage = &v
	return s
}

func (s *SearchObjectResponseData) SetRecords(v []*SearchObjectResponseDataRecords) *SearchObjectResponseData {
	s.Records = v
	return s
}

type SearchObjectResponseDataRecords struct {
	CompareResult   *string  `json:"CompareResult,omitempty" xml:"CompareResult,omitempty" require:"true"`
	DeviceID        *string  `json:"DeviceID,omitempty" xml:"DeviceID,omitempty" require:"true"`
	ShotTime        *int64   `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	LeftTopX        *int     `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY        *int     `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	RightBtmX       *int     `json:"RightBtmX,omitempty" xml:"RightBtmX,omitempty" require:"true"`
	RightBtmY       *int     `json:"RightBtmY,omitempty" xml:"RightBtmY,omitempty" require:"true"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	SourceID        *string  `json:"SourceID,omitempty" xml:"SourceID,omitempty" require:"true"`
	SourceImagePath *string  `json:"SourceImagePath,omitempty" xml:"SourceImagePath,omitempty" require:"true"`
	SourceImageUrl  *string  `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty" require:"true"`
	TargetImagePath *string  `json:"TargetImagePath,omitempty" xml:"TargetImagePath,omitempty" require:"true"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty" require:"true"`
}

func (s SearchObjectResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchObjectResponseDataRecords) GoString() string {
	return s.String()
}

func (s *SearchObjectResponseDataRecords) SetCompareResult(v string) *SearchObjectResponseDataRecords {
	s.CompareResult = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetDeviceID(v string) *SearchObjectResponseDataRecords {
	s.DeviceID = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetShotTime(v int64) *SearchObjectResponseDataRecords {
	s.ShotTime = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetLeftTopX(v int) *SearchObjectResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetLeftTopY(v int) *SearchObjectResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetRightBtmX(v int) *SearchObjectResponseDataRecords {
	s.RightBtmX = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetRightBtmY(v int) *SearchObjectResponseDataRecords {
	s.RightBtmY = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetScore(v float32) *SearchObjectResponseDataRecords {
	s.Score = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetSourceID(v string) *SearchObjectResponseDataRecords {
	s.SourceID = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetSourceImagePath(v string) *SearchObjectResponseDataRecords {
	s.SourceImagePath = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetSourceImageUrl(v string) *SearchObjectResponseDataRecords {
	s.SourceImageUrl = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetTargetImagePath(v string) *SearchObjectResponseDataRecords {
	s.TargetImagePath = &v
	return s
}

func (s *SearchObjectResponseDataRecords) SetTargetImageUrl(v string) *SearchObjectResponseDataRecords {
	s.TargetImageUrl = &v
	return s
}

type DescribeDevicesRequest struct {
	PageNum    *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CorpIdList *string `json:"CorpIdList,omitempty" xml:"CorpIdList,omitempty" require:"true"`
}

func (s DescribeDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDevicesRequest) SetPageNum(v int) *DescribeDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageSize(v int) *DescribeDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesRequest) SetCorpIdList(v string) *DescribeDevicesRequest {
	s.CorpIdList = &v
	return s
}

type DescribeDevicesResponse struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DescribeDevicesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponse) SetCode(v string) *DescribeDevicesResponse {
	s.Code = &v
	return s
}

func (s *DescribeDevicesResponse) SetMessage(v string) *DescribeDevicesResponse {
	s.Message = &v
	return s
}

func (s *DescribeDevicesResponse) SetRequestId(v string) *DescribeDevicesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDevicesResponse) SetData(v *DescribeDevicesResponseData) *DescribeDevicesResponse {
	s.Data = v
	return s
}

type DescribeDevicesResponseData struct {
	PageNum    *int                                  `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize   *int                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                                  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*DescribeDevicesResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDevicesResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseData) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseData) SetPageNum(v int) *DescribeDevicesResponseData {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesResponseData) SetPageSize(v int) *DescribeDevicesResponseData {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesResponseData) SetTotalCount(v int) *DescribeDevicesResponseData {
	s.TotalCount = &v
	return s
}

func (s *DescribeDevicesResponseData) SetTotalPage(v int) *DescribeDevicesResponseData {
	s.TotalPage = &v
	return s
}

func (s *DescribeDevicesResponseData) SetRecords(v []*DescribeDevicesResponseDataRecords) *DescribeDevicesResponseData {
	s.Records = v
	return s
}

type DescribeDevicesResponseDataRecords struct {
	InProtocol    *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty" require:"true"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	DeviceAddress *string `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty" require:"true"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty" require:"true"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty" require:"true"`
	Latitude      *string `json:"Latitude,omitempty" xml:"Latitude,omitempty" require:"true"`
	Longitude     *string `json:"Longitude,omitempty" xml:"Longitude,omitempty" require:"true"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Vendor        *string `json:"Vendor,omitempty" xml:"Vendor,omitempty" require:"true"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
}

func (s DescribeDevicesResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseDataRecords) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseDataRecords) SetInProtocol(v string) *DescribeDevicesResponseDataRecords {
	s.InProtocol = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetCreateTime(v string) *DescribeDevicesResponseDataRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetDeviceAddress(v string) *DescribeDevicesResponseDataRecords {
	s.DeviceAddress = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetDeviceId(v string) *DescribeDevicesResponseDataRecords {
	s.DeviceId = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetDeviceName(v string) *DescribeDevicesResponseDataRecords {
	s.DeviceName = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetDeviceType(v string) *DescribeDevicesResponseDataRecords {
	s.DeviceType = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetLatitude(v string) *DescribeDevicesResponseDataRecords {
	s.Latitude = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetLongitude(v string) *DescribeDevicesResponseDataRecords {
	s.Longitude = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetStatus(v string) *DescribeDevicesResponseDataRecords {
	s.Status = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetVendor(v string) *DescribeDevicesResponseDataRecords {
	s.Vendor = &v
	return s
}

func (s *DescribeDevicesResponseDataRecords) SetCorpId(v string) *DescribeDevicesResponseDataRecords {
	s.CorpId = &v
	return s
}

type GetProfileListRequest struct {
	CorpId                *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId              *string                `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	Name                  *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId             *int64                 `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IdNumber              *string                `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl               *string                `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress           *string                `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender                *int                   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string                `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string                `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType             *string                `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId                 *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize              *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *GetProfileListRequest) SetGender(v int) *GetProfileListRequest {
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
	IsvSubId              *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId             *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IdNumber              *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl               *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress           *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender                *int    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType             *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *GetProfileListShrinkRequest) SetGender(v int) *GetProfileListShrinkRequest {
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

type GetProfileListResponse struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetProfileListResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetProfileListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponse) GoString() string {
	return s.String()
}

func (s *GetProfileListResponse) SetCode(v string) *GetProfileListResponse {
	s.Code = &v
	return s
}

func (s *GetProfileListResponse) SetMessage(v string) *GetProfileListResponse {
	s.Message = &v
	return s
}

func (s *GetProfileListResponse) SetRequestId(v string) *GetProfileListResponse {
	s.RequestId = &v
	return s
}

func (s *GetProfileListResponse) SetData(v *GetProfileListResponseData) *GetProfileListResponse {
	s.Data = v
	return s
}

type GetProfileListResponseData struct {
	PageNumber *int64                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                               `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Success    *bool                                `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Total      *int64                               `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Records    []*GetProfileListResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s GetProfileListResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponseData) GoString() string {
	return s.String()
}

func (s *GetProfileListResponseData) SetPageNumber(v int64) *GetProfileListResponseData {
	s.PageNumber = &v
	return s
}

func (s *GetProfileListResponseData) SetPageSize(v int64) *GetProfileListResponseData {
	s.PageSize = &v
	return s
}

func (s *GetProfileListResponseData) SetSuccess(v bool) *GetProfileListResponseData {
	s.Success = &v
	return s
}

func (s *GetProfileListResponseData) SetTotal(v int64) *GetProfileListResponseData {
	s.Total = &v
	return s
}

func (s *GetProfileListResponseData) SetRecords(v []*GetProfileListResponseDataRecords) *GetProfileListResponseData {
	s.Records = v
	return s
}

type GetProfileListResponseDataRecords struct {
	SceneType          *string `json:"SceneType,omitempty" xml:"SceneType,omitempty" require:"true"`
	BizId              *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	FaceUrl            *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty" require:"true"`
	Gender             *string `json:"Gender,omitempty" xml:"Gender,omitempty" require:"true"`
	IdNumber           *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" require:"true"`
	IsvSubId           *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	SearchMatchingRate *string `json:"SearchMatchingRate,omitempty" xml:"SearchMatchingRate,omitempty" require:"true"`
	PersonId           *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	CatalogId          *int    `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	ProfileId          *int    `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetProfileListResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetProfileListResponseDataRecords) GoString() string {
	return s.String()
}

func (s *GetProfileListResponseDataRecords) SetSceneType(v string) *GetProfileListResponseDataRecords {
	s.SceneType = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetBizId(v string) *GetProfileListResponseDataRecords {
	s.BizId = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetFaceUrl(v string) *GetProfileListResponseDataRecords {
	s.FaceUrl = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetGender(v string) *GetProfileListResponseDataRecords {
	s.Gender = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetIdNumber(v string) *GetProfileListResponseDataRecords {
	s.IdNumber = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetIsvSubId(v string) *GetProfileListResponseDataRecords {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetSearchMatchingRate(v string) *GetProfileListResponseDataRecords {
	s.SearchMatchingRate = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetPersonId(v string) *GetProfileListResponseDataRecords {
	s.PersonId = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetCatalogId(v int) *GetProfileListResponseDataRecords {
	s.CatalogId = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetProfileId(v int) *GetProfileListResponseDataRecords {
	s.ProfileId = &v
	return s
}

func (s *GetProfileListResponseDataRecords) SetName(v string) *GetProfileListResponseDataRecords {
	s.Name = &v
	return s
}

type GetProfileDetailRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	ProfileId *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
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

type GetProfileDetailResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetProfileDetailResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetProfileDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProfileDetailResponse) GoString() string {
	return s.String()
}

func (s *GetProfileDetailResponse) SetCode(v string) *GetProfileDetailResponse {
	s.Code = &v
	return s
}

func (s *GetProfileDetailResponse) SetMessage(v string) *GetProfileDetailResponse {
	s.Message = &v
	return s
}

func (s *GetProfileDetailResponse) SetRequestId(v string) *GetProfileDetailResponse {
	s.RequestId = &v
	return s
}

func (s *GetProfileDetailResponse) SetData(v *GetProfileDetailResponseData) *GetProfileDetailResponse {
	s.Data = v
	return s
}

type GetProfileDetailResponseData struct {
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty" require:"true"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty" require:"true"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty" require:"true"`
	Gender      *string `json:"Gender,omitempty" xml:"Gender,omitempty" require:"true"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" require:"true"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty" require:"true"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty" require:"true"`
	CatalogId   *int    `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	ProfileId   *int    `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	PersonId    *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
}

func (s GetProfileDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetProfileDetailResponseData) GoString() string {
	return s.String()
}

func (s *GetProfileDetailResponseData) SetLiveAddress(v string) *GetProfileDetailResponseData {
	s.LiveAddress = &v
	return s
}

func (s *GetProfileDetailResponseData) SetSceneType(v string) *GetProfileDetailResponseData {
	s.SceneType = &v
	return s
}

func (s *GetProfileDetailResponseData) SetBizId(v string) *GetProfileDetailResponseData {
	s.BizId = &v
	return s
}

func (s *GetProfileDetailResponseData) SetFaceUrl(v string) *GetProfileDetailResponseData {
	s.FaceUrl = &v
	return s
}

func (s *GetProfileDetailResponseData) SetGender(v string) *GetProfileDetailResponseData {
	s.Gender = &v
	return s
}

func (s *GetProfileDetailResponseData) SetIdNumber(v string) *GetProfileDetailResponseData {
	s.IdNumber = &v
	return s
}

func (s *GetProfileDetailResponseData) SetIsvSubId(v string) *GetProfileDetailResponseData {
	s.IsvSubId = &v
	return s
}

func (s *GetProfileDetailResponseData) SetPhoneNo(v string) *GetProfileDetailResponseData {
	s.PhoneNo = &v
	return s
}

func (s *GetProfileDetailResponseData) SetPlateNo(v string) *GetProfileDetailResponseData {
	s.PlateNo = &v
	return s
}

func (s *GetProfileDetailResponseData) SetCatalogId(v int) *GetProfileDetailResponseData {
	s.CatalogId = &v
	return s
}

func (s *GetProfileDetailResponseData) SetProfileId(v int) *GetProfileDetailResponseData {
	s.ProfileId = &v
	return s
}

func (s *GetProfileDetailResponseData) SetName(v string) *GetProfileDetailResponseData {
	s.Name = &v
	return s
}

func (s *GetProfileDetailResponseData) SetPersonId(v string) *GetProfileDetailResponseData {
	s.PersonId = &v
	return s
}

type DeleteProfileCatalogRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
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

type DeleteProfileCatalogResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteProfileCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileCatalogResponse) GoString() string {
	return s.String()
}

func (s *DeleteProfileCatalogResponse) SetCode(v string) *DeleteProfileCatalogResponse {
	s.Code = &v
	return s
}

func (s *DeleteProfileCatalogResponse) SetData(v bool) *DeleteProfileCatalogResponse {
	s.Data = &v
	return s
}

func (s *DeleteProfileCatalogResponse) SetMessage(v string) *DeleteProfileCatalogResponse {
	s.Message = &v
	return s
}

func (s *DeleteProfileCatalogResponse) SetRequestId(v string) *DeleteProfileCatalogResponse {
	s.RequestId = &v
	return s
}

type BindPersonRequest struct {
	CorpId             *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId           *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	PersonMatchingRate *string `json:"PersonMatchingRate,omitempty" xml:"PersonMatchingRate,omitempty" require:"true"`
	PersonId           *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	ProfileId          *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
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

type BindPersonResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BindPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s BindPersonResponse) GoString() string {
	return s.String()
}

func (s *BindPersonResponse) SetCode(v string) *BindPersonResponse {
	s.Code = &v
	return s
}

func (s *BindPersonResponse) SetData(v bool) *BindPersonResponse {
	s.Data = &v
	return s
}

func (s *BindPersonResponse) SetMessage(v string) *BindPersonResponse {
	s.Message = &v
	return s
}

func (s *BindPersonResponse) SetRequestId(v string) *BindPersonResponse {
	s.RequestId = &v
	return s
}

type UpdateProfileRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender      *int    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ProfileId   *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
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

func (s *UpdateProfileRequest) SetGender(v int) *UpdateProfileRequest {
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

type UpdateProfileResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateProfileResponse) SetCode(v string) *UpdateProfileResponse {
	s.Code = &v
	return s
}

func (s *UpdateProfileResponse) SetData(v string) *UpdateProfileResponse {
	s.Data = &v
	return s
}

func (s *UpdateProfileResponse) SetMessage(v string) *UpdateProfileResponse {
	s.Message = &v
	return s
}

func (s *UpdateProfileResponse) SetRequestId(v string) *UpdateProfileResponse {
	s.RequestId = &v
	return s
}

type UnbindPersonRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	ProfileId *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
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

type UnbindPersonResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnbindPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindPersonResponse) GoString() string {
	return s.String()
}

func (s *UnbindPersonResponse) SetCode(v string) *UnbindPersonResponse {
	s.Code = &v
	return s
}

func (s *UnbindPersonResponse) SetData(v bool) *UnbindPersonResponse {
	s.Data = &v
	return s
}

func (s *UnbindPersonResponse) SetMessage(v string) *UnbindPersonResponse {
	s.Message = &v
	return s
}

func (s *UnbindPersonResponse) SetRequestId(v string) *UnbindPersonResponse {
	s.RequestId = &v
	return s
}

type AddProfileRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty"`
	Gender      *int    `json:"Gender,omitempty" xml:"Gender,omitempty"`
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

func (s *AddProfileRequest) SetGender(v int) *AddProfileRequest {
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

type AddProfileResponse struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AddProfileResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponse) GoString() string {
	return s.String()
}

func (s *AddProfileResponse) SetCode(v string) *AddProfileResponse {
	s.Code = &v
	return s
}

func (s *AddProfileResponse) SetMessage(v string) *AddProfileResponse {
	s.Message = &v
	return s
}

func (s *AddProfileResponse) SetRequestId(v string) *AddProfileResponse {
	s.RequestId = &v
	return s
}

func (s *AddProfileResponse) SetData(v *AddProfileResponseData) *AddProfileResponse {
	s.Data = v
	return s
}

type AddProfileResponseData struct {
	LiveAddress *string `json:"LiveAddress,omitempty" xml:"LiveAddress,omitempty" require:"true"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty" require:"true"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	FaceUrl     *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty" require:"true"`
	Gender      *string `json:"Gender,omitempty" xml:"Gender,omitempty" require:"true"`
	IdNumber    *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" require:"true"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	PhoneNo     *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty" require:"true"`
	PlateNo     *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty" require:"true"`
	CatalogId   *int    `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	ProfileId   *int    `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s AddProfileResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponseData) GoString() string {
	return s.String()
}

func (s *AddProfileResponseData) SetLiveAddress(v string) *AddProfileResponseData {
	s.LiveAddress = &v
	return s
}

func (s *AddProfileResponseData) SetSceneType(v string) *AddProfileResponseData {
	s.SceneType = &v
	return s
}

func (s *AddProfileResponseData) SetBizId(v string) *AddProfileResponseData {
	s.BizId = &v
	return s
}

func (s *AddProfileResponseData) SetFaceUrl(v string) *AddProfileResponseData {
	s.FaceUrl = &v
	return s
}

func (s *AddProfileResponseData) SetGender(v string) *AddProfileResponseData {
	s.Gender = &v
	return s
}

func (s *AddProfileResponseData) SetIdNumber(v string) *AddProfileResponseData {
	s.IdNumber = &v
	return s
}

func (s *AddProfileResponseData) SetIsvSubId(v string) *AddProfileResponseData {
	s.IsvSubId = &v
	return s
}

func (s *AddProfileResponseData) SetPhoneNo(v string) *AddProfileResponseData {
	s.PhoneNo = &v
	return s
}

func (s *AddProfileResponseData) SetPlateNo(v string) *AddProfileResponseData {
	s.PlateNo = &v
	return s
}

func (s *AddProfileResponseData) SetCatalogId(v int) *AddProfileResponseData {
	s.CatalogId = &v
	return s
}

func (s *AddProfileResponseData) SetProfileId(v int) *AddProfileResponseData {
	s.ProfileId = &v
	return s
}

func (s *AddProfileResponseData) SetName(v string) *AddProfileResponseData {
	s.Name = &v
	return s
}

type UpdateProfileCatalogRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty" require:"true"`
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

type UpdateProfileCatalogResponse struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *UpdateProfileCatalogResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s UpdateProfileCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileCatalogResponse) GoString() string {
	return s.String()
}

func (s *UpdateProfileCatalogResponse) SetCode(v string) *UpdateProfileCatalogResponse {
	s.Code = &v
	return s
}

func (s *UpdateProfileCatalogResponse) SetMessage(v string) *UpdateProfileCatalogResponse {
	s.Message = &v
	return s
}

func (s *UpdateProfileCatalogResponse) SetRequestId(v string) *UpdateProfileCatalogResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateProfileCatalogResponse) SetData(v *UpdateProfileCatalogResponseData) *UpdateProfileCatalogResponse {
	s.Data = v
	return s
}

type UpdateProfileCatalogResponseData struct {
	IsvSubId        *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	ParentCatalogId *string `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty" require:"true"`
	ProfileCount    *int64  `json:"ProfileCount,omitempty" xml:"ProfileCount,omitempty" require:"true"`
	CatalogId       *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty" require:"true"`
}

func (s UpdateProfileCatalogResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateProfileCatalogResponseData) GoString() string {
	return s.String()
}

func (s *UpdateProfileCatalogResponseData) SetIsvSubId(v string) *UpdateProfileCatalogResponseData {
	s.IsvSubId = &v
	return s
}

func (s *UpdateProfileCatalogResponseData) SetParentCatalogId(v string) *UpdateProfileCatalogResponseData {
	s.ParentCatalogId = &v
	return s
}

func (s *UpdateProfileCatalogResponseData) SetProfileCount(v int64) *UpdateProfileCatalogResponseData {
	s.ProfileCount = &v
	return s
}

func (s *UpdateProfileCatalogResponseData) SetCatalogId(v int64) *UpdateProfileCatalogResponseData {
	s.CatalogId = &v
	return s
}

func (s *UpdateProfileCatalogResponseData) SetCatalogName(v string) *UpdateProfileCatalogResponseData {
	s.CatalogName = &v
	return s
}

type AddProfileCatalogRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId        *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty" require:"true"`
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

type AddProfileCatalogResponse struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AddProfileCatalogResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddProfileCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProfileCatalogResponse) GoString() string {
	return s.String()
}

func (s *AddProfileCatalogResponse) SetCode(v string) *AddProfileCatalogResponse {
	s.Code = &v
	return s
}

func (s *AddProfileCatalogResponse) SetMessage(v string) *AddProfileCatalogResponse {
	s.Message = &v
	return s
}

func (s *AddProfileCatalogResponse) SetRequestId(v string) *AddProfileCatalogResponse {
	s.RequestId = &v
	return s
}

func (s *AddProfileCatalogResponse) SetData(v *AddProfileCatalogResponseData) *AddProfileCatalogResponse {
	s.Data = v
	return s
}

type AddProfileCatalogResponseData struct {
	CatalogId   *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty" require:"true"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
}

func (s AddProfileCatalogResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddProfileCatalogResponseData) GoString() string {
	return s.String()
}

func (s *AddProfileCatalogResponseData) SetCatalogId(v int64) *AddProfileCatalogResponseData {
	s.CatalogId = &v
	return s
}

func (s *AddProfileCatalogResponseData) SetCatalogName(v string) *AddProfileCatalogResponseData {
	s.CatalogName = &v
	return s
}

func (s *AddProfileCatalogResponseData) SetIsvSubId(v string) *AddProfileCatalogResponseData {
	s.IsvSubId = &v
	return s
}

type GetCatalogListRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
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

type GetCatalogListResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      []*GetCatalogListResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s GetCatalogListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogListResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogListResponse) SetCode(v string) *GetCatalogListResponse {
	s.Code = &v
	return s
}

func (s *GetCatalogListResponse) SetMessage(v string) *GetCatalogListResponse {
	s.Message = &v
	return s
}

func (s *GetCatalogListResponse) SetRequestId(v string) *GetCatalogListResponse {
	s.RequestId = &v
	return s
}

func (s *GetCatalogListResponse) SetData(v []*GetCatalogListResponseData) *GetCatalogListResponse {
	s.Data = v
	return s
}

type GetCatalogListResponseData struct {
	IsvSubId        *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	ParentCatalogId *int64  `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty" require:"true"`
	ProfileCount    *int64  `json:"ProfileCount,omitempty" xml:"ProfileCount,omitempty" require:"true"`
	CatalogId       *int64  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty" require:"true"`
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty" require:"true"`
}

func (s GetCatalogListResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogListResponseData) GoString() string {
	return s.String()
}

func (s *GetCatalogListResponseData) SetIsvSubId(v string) *GetCatalogListResponseData {
	s.IsvSubId = &v
	return s
}

func (s *GetCatalogListResponseData) SetParentCatalogId(v int64) *GetCatalogListResponseData {
	s.ParentCatalogId = &v
	return s
}

func (s *GetCatalogListResponseData) SetProfileCount(v int64) *GetCatalogListResponseData {
	s.ProfileCount = &v
	return s
}

func (s *GetCatalogListResponseData) SetCatalogId(v int64) *GetCatalogListResponseData {
	s.CatalogId = &v
	return s
}

func (s *GetCatalogListResponseData) SetCatalogName(v string) *GetCatalogListResponseData {
	s.CatalogName = &v
	return s
}

type DeleteProfileRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId  *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	ProfileId *int64  `json:"ProfileId,omitempty" xml:"ProfileId,omitempty" require:"true"`
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

type DeleteProfileResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteProfileResponse) SetCode(v string) *DeleteProfileResponse {
	s.Code = &v
	return s
}

func (s *DeleteProfileResponse) SetData(v bool) *DeleteProfileResponse {
	s.Data = &v
	return s
}

func (s *DeleteProfileResponse) SetMessage(v string) *DeleteProfileResponse {
	s.Message = &v
	return s
}

func (s *DeleteProfileResponse) SetRequestId(v string) *DeleteProfileResponse {
	s.RequestId = &v
	return s
}

type UnbindCorpGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	CorpGroupId *string `json:"CorpGroupId,omitempty" xml:"CorpGroupId,omitempty" require:"true"`
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

type UnbindCorpGroupResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s UnbindCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *UnbindCorpGroupResponse) SetCode(v string) *UnbindCorpGroupResponse {
	s.Code = &v
	return s
}

func (s *UnbindCorpGroupResponse) SetMessage(v string) *UnbindCorpGroupResponse {
	s.Message = &v
	return s
}

func (s *UnbindCorpGroupResponse) SetRequestId(v string) *UnbindCorpGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UnbindCorpGroupResponse) SetSuccess(v bool) *UnbindCorpGroupResponse {
	s.Success = &v
	return s
}

type BindCorpGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	CorpGroupId *string `json:"CorpGroupId,omitempty" xml:"CorpGroupId,omitempty" require:"true"`
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

type BindCorpGroupResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s BindCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BindCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *BindCorpGroupResponse) SetCode(v string) *BindCorpGroupResponse {
	s.Code = &v
	return s
}

func (s *BindCorpGroupResponse) SetMessage(v string) *BindCorpGroupResponse {
	s.Message = &v
	return s
}

func (s *BindCorpGroupResponse) SetRequestId(v string) *BindCorpGroupResponse {
	s.RequestId = &v
	return s
}

func (s *BindCorpGroupResponse) SetSuccess(v bool) *BindCorpGroupResponse {
	s.Success = &v
	return s
}

type ListUserGroupsRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
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

type ListUserGroupsResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      []*ListUserGroupsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponse) SetCode(v string) *ListUserGroupsResponse {
	s.Code = &v
	return s
}

func (s *ListUserGroupsResponse) SetMessage(v string) *ListUserGroupsResponse {
	s.Message = &v
	return s
}

func (s *ListUserGroupsResponse) SetRequestId(v string) *ListUserGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponse) SetData(v []*ListUserGroupsResponseData) *ListUserGroupsResponse {
	s.Data = v
	return s
}

type ListUserGroupsResponseData struct {
	Creator           *string `json:"Creator,omitempty" xml:"Creator,omitempty" require:"true"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty" require:"true"`
	IsvSubId          *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserGroupId       *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	UserCount         *int64  `json:"UserCount,omitempty" xml:"UserCount,omitempty" require:"true"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime        *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	ParentUserGroupId *int64  `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty" require:"true"`
}

func (s ListUserGroupsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseData) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseData) SetCreator(v string) *ListUserGroupsResponseData {
	s.Creator = &v
	return s
}

func (s *ListUserGroupsResponseData) SetUserGroupName(v string) *ListUserGroupsResponseData {
	s.UserGroupName = &v
	return s
}

func (s *ListUserGroupsResponseData) SetIsvSubId(v string) *ListUserGroupsResponseData {
	s.IsvSubId = &v
	return s
}

func (s *ListUserGroupsResponseData) SetUserGroupId(v int64) *ListUserGroupsResponseData {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsResponseData) SetUserCount(v int64) *ListUserGroupsResponseData {
	s.UserCount = &v
	return s
}

func (s *ListUserGroupsResponseData) SetCreateTime(v string) *ListUserGroupsResponseData {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsResponseData) SetUpdateTime(v string) *ListUserGroupsResponseData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserGroupsResponseData) SetParentUserGroupId(v int64) *ListUserGroupsResponseData {
	s.ParentUserGroupId = &v
	return s
}

type GetPersonListRequest struct {
	PageNumber                *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize                  *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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
	PageNumber                *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize                  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type GetPersonListResponse struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetPersonListResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetPersonListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponse) GoString() string {
	return s.String()
}

func (s *GetPersonListResponse) SetCode(v string) *GetPersonListResponse {
	s.Code = &v
	return s
}

func (s *GetPersonListResponse) SetMessage(v string) *GetPersonListResponse {
	s.Message = &v
	return s
}

func (s *GetPersonListResponse) SetRequestId(v string) *GetPersonListResponse {
	s.RequestId = &v
	return s
}

func (s *GetPersonListResponse) SetData(v *GetPersonListResponseData) *GetPersonListResponse {
	s.Data = v
	return s
}

type GetPersonListResponseData struct {
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Records    []*GetPersonListResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s GetPersonListResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseData) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseData) SetPageNumber(v int64) *GetPersonListResponseData {
	s.PageNumber = &v
	return s
}

func (s *GetPersonListResponseData) SetPageSize(v int64) *GetPersonListResponseData {
	s.PageSize = &v
	return s
}

func (s *GetPersonListResponseData) SetTotalCount(v int64) *GetPersonListResponseData {
	s.TotalCount = &v
	return s
}

func (s *GetPersonListResponseData) SetRecords(v []*GetPersonListResponseDataRecords) *GetPersonListResponseData {
	s.Records = v
	return s
}

type GetPersonListResponseDataRecords struct {
	FaceUrl            *string                                            `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty" require:"true"`
	FirstShotTime      *int64                                             `json:"FirstShotTime,omitempty" xml:"FirstShotTime,omitempty" require:"true"`
	PersonId           *string                                            `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	SearchMatchingRate *string                                            `json:"SearchMatchingRate,omitempty" xml:"SearchMatchingRate,omitempty" require:"true"`
	LastShotTime       *int64                                             `json:"LastShotTime,omitempty" xml:"LastShotTime,omitempty" require:"true"`
	PropertyTagList    []*GetPersonListResponseDataRecordsPropertyTagList `json:"PropertyTagList,omitempty" xml:"PropertyTagList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPersonListResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseDataRecords) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseDataRecords) SetFaceUrl(v string) *GetPersonListResponseDataRecords {
	s.FaceUrl = &v
	return s
}

func (s *GetPersonListResponseDataRecords) SetFirstShotTime(v int64) *GetPersonListResponseDataRecords {
	s.FirstShotTime = &v
	return s
}

func (s *GetPersonListResponseDataRecords) SetPersonId(v string) *GetPersonListResponseDataRecords {
	s.PersonId = &v
	return s
}

func (s *GetPersonListResponseDataRecords) SetSearchMatchingRate(v string) *GetPersonListResponseDataRecords {
	s.SearchMatchingRate = &v
	return s
}

func (s *GetPersonListResponseDataRecords) SetLastShotTime(v int64) *GetPersonListResponseDataRecords {
	s.LastShotTime = &v
	return s
}

func (s *GetPersonListResponseDataRecords) SetPropertyTagList(v []*GetPersonListResponseDataRecordsPropertyTagList) *GetPersonListResponseDataRecords {
	s.PropertyTagList = v
	return s
}

type GetPersonListResponseDataRecordsPropertyTagList struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	TagCodeName *string `json:"TagCodeName,omitempty" xml:"TagCodeName,omitempty" require:"true"`
	TagName     *string `json:"TagName,omitempty" xml:"TagName,omitempty" require:"true"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetPersonListResponseDataRecordsPropertyTagList) String() string {
	return tea.Prettify(s)
}

func (s GetPersonListResponseDataRecordsPropertyTagList) GoString() string {
	return s.String()
}

func (s *GetPersonListResponseDataRecordsPropertyTagList) SetCode(v string) *GetPersonListResponseDataRecordsPropertyTagList {
	s.Code = &v
	return s
}

func (s *GetPersonListResponseDataRecordsPropertyTagList) SetTagCodeName(v string) *GetPersonListResponseDataRecordsPropertyTagList {
	s.TagCodeName = &v
	return s
}

func (s *GetPersonListResponseDataRecordsPropertyTagList) SetTagName(v string) *GetPersonListResponseDataRecordsPropertyTagList {
	s.TagName = &v
	return s
}

func (s *GetPersonListResponseDataRecordsPropertyTagList) SetValue(v string) *GetPersonListResponseDataRecordsPropertyTagList {
	s.Value = &v
	return s
}

type ListUsersRequest struct {
	CorpId                *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId              *string                `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserName              *string                `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId           *int64                 `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdNumber              *string                `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl          *string                `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Address               *string                `json:"Address,omitempty" xml:"Address,omitempty"`
	Age                   *int                   `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender                *int                   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string                `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string                `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment            *string                `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId                 *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize              *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *ListUsersRequest) SetAge(v int) *ListUsersRequest {
	s.Age = &v
	return s
}

func (s *ListUsersRequest) SetGender(v int) *ListUsersRequest {
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
	IsvSubId              *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId           *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdNumber              *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl          *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Age                   *int    `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender                *int    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo               *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo               *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment            *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *ListUsersShrinkRequest) SetAge(v int) *ListUsersShrinkRequest {
	s.Age = &v
	return s
}

func (s *ListUsersShrinkRequest) SetGender(v int) *ListUsersShrinkRequest {
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

type ListUsersResponse struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListUsersResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetCode(v string) *ListUsersResponse {
	s.Code = &v
	return s
}

func (s *ListUsersResponse) SetMessage(v string) *ListUsersResponse {
	s.Message = &v
	return s
}

func (s *ListUsersResponse) SetRequestId(v string) *ListUsersResponse {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponse) SetData(v *ListUsersResponseData) *ListUsersResponse {
	s.Data = v
	return s
}

type ListUsersResponseData struct {
	PageNumber *int64                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Success    *int64                          `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Total      *int64                          `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Records    []*ListUsersResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListUsersResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseData) SetPageNumber(v int64) *ListUsersResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseData) SetPageSize(v int64) *ListUsersResponseData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseData) SetSuccess(v int64) *ListUsersResponseData {
	s.Success = &v
	return s
}

func (s *ListUsersResponseData) SetTotal(v int64) *ListUsersResponseData {
	s.Total = &v
	return s
}

func (s *ListUsersResponseData) SetRecords(v []*ListUsersResponseDataRecords) *ListUsersResponseData {
	s.Records = v
	return s
}

type ListUsersResponseDataRecords struct {
	UserGroupId  *int    `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty" require:"true"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty" require:"true"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty" require:"true"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" require:"true"`
	UserId       *int    `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	MatchingRate *string `json:"MatchingRate,omitempty" xml:"MatchingRate,omitempty" require:"true"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
}

func (s ListUsersResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListUsersResponseDataRecords) SetUserGroupId(v int) *ListUsersResponseDataRecords {
	s.UserGroupId = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetAge(v string) *ListUsersResponseDataRecords {
	s.Age = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetAttachment(v string) *ListUsersResponseDataRecords {
	s.Attachment = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetBizId(v string) *ListUsersResponseDataRecords {
	s.BizId = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetFaceImageUrl(v string) *ListUsersResponseDataRecords {
	s.FaceImageUrl = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetGender(v string) *ListUsersResponseDataRecords {
	s.Gender = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetIdNumber(v string) *ListUsersResponseDataRecords {
	s.IdNumber = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetUserId(v int) *ListUsersResponseDataRecords {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetUserName(v string) *ListUsersResponseDataRecords {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetIsvSubId(v string) *ListUsersResponseDataRecords {
	s.IsvSubId = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetMatchingRate(v string) *ListUsersResponseDataRecords {
	s.MatchingRate = &v
	return s
}

func (s *ListUsersResponseDataRecords) SetPersonId(v string) *ListUsersResponseDataRecords {
	s.PersonId = &v
	return s
}

type CreateUserRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	UserGroupId  *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Age          *int    `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender       *int    `json:"Gender,omitempty" xml:"Gender,omitempty"`
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

func (s *CreateUserRequest) SetAge(v int) *CreateUserRequest {
	s.Age = &v
	return s
}

func (s *CreateUserRequest) SetGender(v int) *CreateUserRequest {
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

type CreateUserResponse struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *CreateUserResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetCode(v string) *CreateUserResponse {
	s.Code = &v
	return s
}

func (s *CreateUserResponse) SetMessage(v string) *CreateUserResponse {
	s.Message = &v
	return s
}

func (s *CreateUserResponse) SetRequestId(v string) *CreateUserResponse {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponse) SetData(v *CreateUserResponseData) *CreateUserResponse {
	s.Data = v
	return s
}

type CreateUserResponseData struct {
	UserId       *int    `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	UserGroupId  *int    `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" require:"true"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty" require:"true"`
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty" require:"true"`
	PlateNo      *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty" require:"true"`
	PhoneNo      *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty" require:"true"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty" require:"true"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
}

func (s CreateUserResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseData) SetUserId(v int) *CreateUserResponseData {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseData) SetIsvSubId(v string) *CreateUserResponseData {
	s.IsvSubId = &v
	return s
}

func (s *CreateUserResponseData) SetUserName(v string) *CreateUserResponseData {
	s.UserName = &v
	return s
}

func (s *CreateUserResponseData) SetUserGroupId(v int) *CreateUserResponseData {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserResponseData) SetIdNumber(v string) *CreateUserResponseData {
	s.IdNumber = &v
	return s
}

func (s *CreateUserResponseData) SetFaceImageUrl(v string) *CreateUserResponseData {
	s.FaceImageUrl = &v
	return s
}

func (s *CreateUserResponseData) SetAddress(v string) *CreateUserResponseData {
	s.Address = &v
	return s
}

func (s *CreateUserResponseData) SetAge(v string) *CreateUserResponseData {
	s.Age = &v
	return s
}

func (s *CreateUserResponseData) SetGender(v string) *CreateUserResponseData {
	s.Gender = &v
	return s
}

func (s *CreateUserResponseData) SetPlateNo(v string) *CreateUserResponseData {
	s.PlateNo = &v
	return s
}

func (s *CreateUserResponseData) SetPhoneNo(v string) *CreateUserResponseData {
	s.PhoneNo = &v
	return s
}

func (s *CreateUserResponseData) SetAttachment(v string) *CreateUserResponseData {
	s.Attachment = &v
	return s
}

func (s *CreateUserResponseData) SetBizId(v string) *CreateUserResponseData {
	s.BizId = &v
	return s
}

type BindUserRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	MatchingRate *string `json:"MatchingRate,omitempty" xml:"MatchingRate,omitempty" require:"true"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
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

type BindUserResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BindUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindUserResponse) GoString() string {
	return s.String()
}

func (s *BindUserResponse) SetCode(v string) *BindUserResponse {
	s.Code = &v
	return s
}

func (s *BindUserResponse) SetData(v bool) *BindUserResponse {
	s.Data = &v
	return s
}

func (s *BindUserResponse) SetMessage(v string) *BindUserResponse {
	s.Message = &v
	return s
}

func (s *BindUserResponse) SetRequestId(v string) *BindUserResponse {
	s.RequestId = &v
	return s
}

type GetUserDetailRequest struct {
	CorpId         *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId       *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserId         *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
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

type GetUserDetailResponse struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetUserDetailResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetUserDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserDetailResponse) GoString() string {
	return s.String()
}

func (s *GetUserDetailResponse) SetCode(v string) *GetUserDetailResponse {
	s.Code = &v
	return s
}

func (s *GetUserDetailResponse) SetMessage(v string) *GetUserDetailResponse {
	s.Message = &v
	return s
}

func (s *GetUserDetailResponse) SetRequestId(v string) *GetUserDetailResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserDetailResponse) SetData(v *GetUserDetailResponseData) *GetUserDetailResponse {
	s.Data = v
	return s
}

type GetUserDetailResponseData struct {
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Age          *string `json:"Age,omitempty" xml:"Age,omitempty" require:"true"`
	Attachment   *string `json:"Attachment,omitempty" xml:"Attachment,omitempty" require:"true"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	Gender       *string `json:"Gender,omitempty" xml:"Gender,omitempty" require:"true"`
	IdNumber     *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" require:"true"`
	PhoneNo      *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty" require:"true"`
	PlateNo      *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty" require:"true"`
	UserGroupId  *int    `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	UserId       *int    `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
}

func (s GetUserDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetUserDetailResponseData) GoString() string {
	return s.String()
}

func (s *GetUserDetailResponseData) SetAddress(v string) *GetUserDetailResponseData {
	s.Address = &v
	return s
}

func (s *GetUserDetailResponseData) SetAge(v string) *GetUserDetailResponseData {
	s.Age = &v
	return s
}

func (s *GetUserDetailResponseData) SetAttachment(v string) *GetUserDetailResponseData {
	s.Attachment = &v
	return s
}

func (s *GetUserDetailResponseData) SetBizId(v string) *GetUserDetailResponseData {
	s.BizId = &v
	return s
}

func (s *GetUserDetailResponseData) SetFaceImageUrl(v string) *GetUserDetailResponseData {
	s.FaceImageUrl = &v
	return s
}

func (s *GetUserDetailResponseData) SetGender(v string) *GetUserDetailResponseData {
	s.Gender = &v
	return s
}

func (s *GetUserDetailResponseData) SetIdNumber(v string) *GetUserDetailResponseData {
	s.IdNumber = &v
	return s
}

func (s *GetUserDetailResponseData) SetPhoneNo(v string) *GetUserDetailResponseData {
	s.PhoneNo = &v
	return s
}

func (s *GetUserDetailResponseData) SetPlateNo(v string) *GetUserDetailResponseData {
	s.PlateNo = &v
	return s
}

func (s *GetUserDetailResponseData) SetUserGroupId(v int) *GetUserDetailResponseData {
	s.UserGroupId = &v
	return s
}

func (s *GetUserDetailResponseData) SetUserId(v int) *GetUserDetailResponseData {
	s.UserId = &v
	return s
}

func (s *GetUserDetailResponseData) SetUserName(v string) *GetUserDetailResponseData {
	s.UserName = &v
	return s
}

func (s *GetUserDetailResponseData) SetIsvSubId(v string) *GetUserDetailResponseData {
	s.IsvSubId = &v
	return s
}

type UploadImageRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
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

type UploadImageResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s UploadImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponse) GoString() string {
	return s.String()
}

func (s *UploadImageResponse) SetCode(v string) *UploadImageResponse {
	s.Code = &v
	return s
}

func (s *UploadImageResponse) SetMessage(v string) *UploadImageResponse {
	s.Message = &v
	return s
}

func (s *UploadImageResponse) SetRequestId(v string) *UploadImageResponse {
	s.RequestId = &v
	return s
}

func (s *UploadImageResponse) SetData(v string) *UploadImageResponse {
	s.Data = &v
	return s
}

type UpdateUserGroupRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId      *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserGroupId   *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty" require:"true"`
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

type UpdateUserGroupResponse struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *UpdateUserGroupResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s UpdateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponse) SetCode(v string) *UpdateUserGroupResponse {
	s.Code = &v
	return s
}

func (s *UpdateUserGroupResponse) SetMessage(v string) *UpdateUserGroupResponse {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupResponse) SetRequestId(v string) *UpdateUserGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponse) SetData(v *UpdateUserGroupResponseData) *UpdateUserGroupResponse {
	s.Data = v
	return s
}

type UpdateUserGroupResponseData struct {
	UserGroupId       *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	IsvSubId          *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty" require:"true"`
	UserCount         *int64  `json:"UserCount,omitempty" xml:"UserCount,omitempty" require:"true"`
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty" require:"true"`
}

func (s UpdateUserGroupResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponseData) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseData) SetUserGroupId(v int64) *UpdateUserGroupResponseData {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupResponseData) SetIsvSubId(v string) *UpdateUserGroupResponseData {
	s.IsvSubId = &v
	return s
}

func (s *UpdateUserGroupResponseData) SetUserGroupName(v string) *UpdateUserGroupResponseData {
	s.UserGroupName = &v
	return s
}

func (s *UpdateUserGroupResponseData) SetUserCount(v int64) *UpdateUserGroupResponseData {
	s.UserCount = &v
	return s
}

func (s *UpdateUserGroupResponseData) SetParentUserGroupId(v string) *UpdateUserGroupResponseData {
	s.ParentUserGroupId = &v
	return s
}

type CreateUserGroupRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId          *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty" require:"true"`
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

type CreateUserGroupResponse struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *CreateUserGroupResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CreateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponse) SetCode(v string) *CreateUserGroupResponse {
	s.Code = &v
	return s
}

func (s *CreateUserGroupResponse) SetMessage(v string) *CreateUserGroupResponse {
	s.Message = &v
	return s
}

func (s *CreateUserGroupResponse) SetRequestId(v string) *CreateUserGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponse) SetData(v *CreateUserGroupResponseData) *CreateUserGroupResponse {
	s.Data = v
	return s
}

type CreateUserGroupResponseData struct {
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty" require:"true"`
	IsvSubId      *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserGroupId   *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
}

func (s CreateUserGroupResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseData) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseData) SetUserGroupName(v string) *CreateUserGroupResponseData {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupResponseData) SetIsvSubId(v string) *CreateUserGroupResponseData {
	s.IsvSubId = &v
	return s
}

func (s *CreateUserGroupResponseData) SetUserGroupId(v int64) *CreateUserGroupResponseData {
	s.UserGroupId = &v
	return s
}

type UnbindUserRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
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

type UnbindUserResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnbindUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserResponse) SetCode(v string) *UnbindUserResponse {
	s.Code = &v
	return s
}

func (s *UnbindUserResponse) SetData(v bool) *UnbindUserResponse {
	s.Data = &v
	return s
}

func (s *UnbindUserResponse) SetMessage(v string) *UnbindUserResponse {
	s.Message = &v
	return s
}

func (s *UnbindUserResponse) SetRequestId(v string) *UnbindUserResponse {
	s.RequestId = &v
	return s
}

type UpdateUserRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId         *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserGroupId      *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	IdNumber         *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	FaceImageUrl     *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceImageContent *string `json:"FaceImageContent,omitempty" xml:"FaceImageContent,omitempty"`
	Address          *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Age              *int    `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender           *int    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PlateNo          *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	PhoneNo          *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	Attachment       *string `json:"Attachment,omitempty" xml:"Attachment,omitempty"`
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UserId           *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
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

func (s *UpdateUserRequest) SetAge(v int) *UpdateUserRequest {
	s.Age = &v
	return s
}

func (s *UpdateUserRequest) SetGender(v int) *UpdateUserRequest {
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

type UpdateUserResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetCode(v string) *UpdateUserResponse {
	s.Code = &v
	return s
}

func (s *UpdateUserResponse) SetMessage(v string) *UpdateUserResponse {
	s.Message = &v
	return s
}

func (s *UpdateUserResponse) SetRequestId(v string) *UpdateUserResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponse) SetData(v string) *UpdateUserResponse {
	s.Data = &v
	return s
}

type DeleteUserRequest struct {
	CorpId   *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
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

type DeleteUserResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetCode(v string) *DeleteUserResponse {
	s.Code = &v
	return s
}

func (s *DeleteUserResponse) SetData(v bool) *DeleteUserResponse {
	s.Data = &v
	return s
}

func (s *DeleteUserResponse) SetMessage(v string) *DeleteUserResponse {
	s.Message = &v
	return s
}

func (s *DeleteUserResponse) SetRequestId(v string) *DeleteUserResponse {
	s.RequestId = &v
	return s
}

type DeleteUserGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	IsvSubId    *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
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

type DeleteUserGroupResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s DeleteUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) SetCode(v string) *DeleteUserGroupResponse {
	s.Code = &v
	return s
}

func (s *DeleteUserGroupResponse) SetMessage(v string) *DeleteUserGroupResponse {
	s.Message = &v
	return s
}

func (s *DeleteUserGroupResponse) SetRequestId(v string) *DeleteUserGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponse) SetData(v bool) *DeleteUserGroupResponse {
	s.Data = &v
	return s
}

type ListPersonVisitCountRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	PageNumber        *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize          *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	AggregateType     *string `json:"AggregateType,omitempty" xml:"AggregateType,omitempty" require:"true"`
	TagCode           *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TimeAggregateType *string `json:"TimeAggregateType,omitempty" xml:"TimeAggregateType,omitempty" require:"true"`
	MinVal            *int    `json:"MinVal,omitempty" xml:"MinVal,omitempty"`
	MaxVal            *int    `json:"MaxVal,omitempty" xml:"MaxVal,omitempty"`
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

func (s *ListPersonVisitCountRequest) SetPageNumber(v int) *ListPersonVisitCountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetPageSize(v int) *ListPersonVisitCountRequest {
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

func (s *ListPersonVisitCountRequest) SetMinVal(v int) *ListPersonVisitCountRequest {
	s.MinVal = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetMaxVal(v int) *ListPersonVisitCountRequest {
	s.MaxVal = &v
	return s
}

func (s *ListPersonVisitCountRequest) SetCountType(v string) *ListPersonVisitCountRequest {
	s.CountType = &v
	return s
}

type ListPersonVisitCountResponse struct {
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	PageNo     *string                             `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize   *string                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *string                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TotalCount *string                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*ListPersonVisitCountResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListPersonVisitCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonVisitCountResponse) GoString() string {
	return s.String()
}

func (s *ListPersonVisitCountResponse) SetCode(v string) *ListPersonVisitCountResponse {
	s.Code = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetMessage(v string) *ListPersonVisitCountResponse {
	s.Message = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetPageNo(v string) *ListPersonVisitCountResponse {
	s.PageNo = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetPageSize(v string) *ListPersonVisitCountResponse {
	s.PageSize = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetRequestId(v string) *ListPersonVisitCountResponse {
	s.RequestId = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetSuccess(v string) *ListPersonVisitCountResponse {
	s.Success = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetTotalCount(v string) *ListPersonVisitCountResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPersonVisitCountResponse) SetData(v []*ListPersonVisitCountResponseData) *ListPersonVisitCountResponse {
	s.Data = v
	return s
}

type ListPersonVisitCountResponseData struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	PersonId   *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TagMetrics *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty" require:"true"`
	HourId     *string `json:"HourId,omitempty" xml:"HourId,omitempty" require:"true"`
	DayId      *string `json:"DayId,omitempty" xml:"DayId,omitempty" require:"true"`
}

func (s ListPersonVisitCountResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonVisitCountResponseData) GoString() string {
	return s.String()
}

func (s *ListPersonVisitCountResponseData) SetCorpId(v string) *ListPersonVisitCountResponseData {
	s.CorpId = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetDeviceId(v string) *ListPersonVisitCountResponseData {
	s.DeviceId = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetGroupId(v string) *ListPersonVisitCountResponseData {
	s.GroupId = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetPersonId(v string) *ListPersonVisitCountResponseData {
	s.PersonId = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetTagCode(v string) *ListPersonVisitCountResponseData {
	s.TagCode = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetTagMetrics(v string) *ListPersonVisitCountResponseData {
	s.TagMetrics = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetHourId(v string) *ListPersonVisitCountResponseData {
	s.HourId = &v
	return s
}

func (s *ListPersonVisitCountResponseData) SetDayId(v string) *ListPersonVisitCountResponseData {
	s.DayId = &v
	return s
}

type ListEventAlgorithmDetailsRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *ListEventAlgorithmDetailsRequest) SetPageNumber(v int) *ListEventAlgorithmDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmDetailsRequest) SetPageSize(v int) *ListEventAlgorithmDetailsRequest {
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

type ListEventAlgorithmDetailsResponse struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	PageNumber *int                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *string                                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TotalCount *int                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*ListEventAlgorithmDetailsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListEventAlgorithmDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmDetailsResponse) SetCode(v string) *ListEventAlgorithmDetailsResponse {
	s.Code = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetMessage(v string) *ListEventAlgorithmDetailsResponse {
	s.Message = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetPageNumber(v int) *ListEventAlgorithmDetailsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetPageSize(v int) *ListEventAlgorithmDetailsResponse {
	s.PageSize = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetRequestId(v string) *ListEventAlgorithmDetailsResponse {
	s.RequestId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetSuccess(v string) *ListEventAlgorithmDetailsResponse {
	s.Success = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetTotalCount(v int) *ListEventAlgorithmDetailsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponse) SetData(v []*ListEventAlgorithmDetailsResponseData) *ListEventAlgorithmDetailsResponse {
	s.Data = v
	return s
}

type ListEventAlgorithmDetailsResponseData struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
	EventType        *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	EventValue       *string `json:"EventValue,omitempty" xml:"EventValue,omitempty" require:"true"`
	ExtendValue      *string `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty" require:"true"`
	ExtraExtendValue *string `json:"ExtraExtendValue,omitempty" xml:"ExtraExtendValue,omitempty" require:"true"`
	FaceCount        *string `json:"FaceCount,omitempty" xml:"FaceCount,omitempty" require:"true"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	PicUrlPath       *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty" require:"true"`
	PointX           *string `json:"PointX,omitempty" xml:"PointX,omitempty" require:"true"`
	PointY           *string `json:"PointY,omitempty" xml:"PointY,omitempty" require:"true"`
	RecordId         *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	ShotTime         *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	SourceId         *string `json:"SourceId,omitempty" xml:"SourceId,omitempty" require:"true"`
	TargetPicUrlPath *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty" require:"true"`
}

func (s ListEventAlgorithmDetailsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmDetailsResponseData) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmDetailsResponseData) SetCorpId(v string) *ListEventAlgorithmDetailsResponseData {
	s.CorpId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetDataSourceId(v string) *ListEventAlgorithmDetailsResponseData {
	s.DataSourceId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetEventType(v string) *ListEventAlgorithmDetailsResponseData {
	s.EventType = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetEventValue(v string) *ListEventAlgorithmDetailsResponseData {
	s.EventValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetExtendValue(v string) *ListEventAlgorithmDetailsResponseData {
	s.ExtendValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetExtraExtendValue(v string) *ListEventAlgorithmDetailsResponseData {
	s.ExtraExtendValue = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetFaceCount(v string) *ListEventAlgorithmDetailsResponseData {
	s.FaceCount = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetLeftTopX(v string) *ListEventAlgorithmDetailsResponseData {
	s.LeftTopX = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetLeftTopY(v string) *ListEventAlgorithmDetailsResponseData {
	s.LeftTopY = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetPicUrlPath(v string) *ListEventAlgorithmDetailsResponseData {
	s.PicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetPointX(v string) *ListEventAlgorithmDetailsResponseData {
	s.PointX = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetPointY(v string) *ListEventAlgorithmDetailsResponseData {
	s.PointY = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetRecordId(v string) *ListEventAlgorithmDetailsResponseData {
	s.RecordId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetRightBottomX(v string) *ListEventAlgorithmDetailsResponseData {
	s.RightBottomX = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetRightBottomY(v string) *ListEventAlgorithmDetailsResponseData {
	s.RightBottomY = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetShotTime(v string) *ListEventAlgorithmDetailsResponseData {
	s.ShotTime = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetSourceId(v string) *ListEventAlgorithmDetailsResponseData {
	s.SourceId = &v
	return s
}

func (s *ListEventAlgorithmDetailsResponseData) SetTargetPicUrlPath(v string) *ListEventAlgorithmDetailsResponseData {
	s.TargetPicUrlPath = &v
	return s
}

type ListCorpMetricsRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TagCode         *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListCorpMetricsResponse struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	PageNumber *int                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *string                        `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TotalCount *int                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*ListCorpMetricsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListCorpMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsResponse) SetCode(v string) *ListCorpMetricsResponse {
	s.Code = &v
	return s
}

func (s *ListCorpMetricsResponse) SetMessage(v string) *ListCorpMetricsResponse {
	s.Message = &v
	return s
}

func (s *ListCorpMetricsResponse) SetPageNumber(v int) *ListCorpMetricsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListCorpMetricsResponse) SetPageSize(v int) *ListCorpMetricsResponse {
	s.PageSize = &v
	return s
}

func (s *ListCorpMetricsResponse) SetRequestId(v string) *ListCorpMetricsResponse {
	s.RequestId = &v
	return s
}

func (s *ListCorpMetricsResponse) SetSuccess(v string) *ListCorpMetricsResponse {
	s.Success = &v
	return s
}

func (s *ListCorpMetricsResponse) SetTotalCount(v int) *ListCorpMetricsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListCorpMetricsResponse) SetData(v []*ListCorpMetricsResponseData) *ListCorpMetricsResponse {
	s.Data = v
	return s
}

type ListCorpMetricsResponseData struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TagMetrics    *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty" require:"true"`
	TagValue      *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty" require:"true"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	PersonId      *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	DateId        *string `json:"DateId,omitempty" xml:"DateId,omitempty" require:"true"`
}

func (s ListCorpMetricsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpMetricsResponseData) GoString() string {
	return s.String()
}

func (s *ListCorpMetricsResponseData) SetCorpId(v string) *ListCorpMetricsResponseData {
	s.CorpId = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetTagCode(v string) *ListCorpMetricsResponseData {
	s.TagCode = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetTagMetrics(v string) *ListCorpMetricsResponseData {
	s.TagMetrics = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetTagValue(v string) *ListCorpMetricsResponseData {
	s.TagValue = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetDeviceGroupId(v string) *ListCorpMetricsResponseData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetDeviceId(v string) *ListCorpMetricsResponseData {
	s.DeviceId = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetUserGroupId(v string) *ListCorpMetricsResponseData {
	s.UserGroupId = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetPersonId(v string) *ListCorpMetricsResponseData {
	s.PersonId = &v
	return s
}

func (s *ListCorpMetricsResponseData) SetDateId(v string) *ListCorpMetricsResponseData {
	s.DateId = &v
	return s
}

type ListPersonTraceRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	PersonId     *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
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

type ListPersonTraceResponse struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *string                        `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TotalCount *int                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber *int                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Data       []*ListPersonTraceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListPersonTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceResponse) GoString() string {
	return s.String()
}

func (s *ListPersonTraceResponse) SetCode(v string) *ListPersonTraceResponse {
	s.Code = &v
	return s
}

func (s *ListPersonTraceResponse) SetMessage(v string) *ListPersonTraceResponse {
	s.Message = &v
	return s
}

func (s *ListPersonTraceResponse) SetRequestId(v string) *ListPersonTraceResponse {
	s.RequestId = &v
	return s
}

func (s *ListPersonTraceResponse) SetSuccess(v string) *ListPersonTraceResponse {
	s.Success = &v
	return s
}

func (s *ListPersonTraceResponse) SetTotalCount(v int) *ListPersonTraceResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPersonTraceResponse) SetPageSize(v int) *ListPersonTraceResponse {
	s.PageSize = &v
	return s
}

func (s *ListPersonTraceResponse) SetPageNumber(v int) *ListPersonTraceResponse {
	s.PageNumber = &v
	return s
}

func (s *ListPersonTraceResponse) SetData(v []*ListPersonTraceResponseData) *ListPersonTraceResponse {
	s.Data = v
	return s
}

type ListPersonTraceResponseData struct {
	Date             *string `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	LastTime         *string `json:"LastTime,omitempty" xml:"LastTime,omitempty" require:"true"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndSourceImage   *string `json:"EndSourceImage,omitempty" xml:"EndSourceImage,omitempty" require:"true"`
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	StartTargetImage *string `json:"StartTargetImage,omitempty" xml:"StartTargetImage,omitempty" require:"true"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	PersonId         *string `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	StartSourceImage *string `json:"StartSourceImage,omitempty" xml:"StartSourceImage,omitempty" require:"true"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	EndTargetImage   *string `json:"EndTargetImage,omitempty" xml:"EndTargetImage,omitempty" require:"true"`
}

func (s ListPersonTraceResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonTraceResponseData) GoString() string {
	return s.String()
}

func (s *ListPersonTraceResponseData) SetDate(v string) *ListPersonTraceResponseData {
	s.Date = &v
	return s
}

func (s *ListPersonTraceResponseData) SetLastTime(v string) *ListPersonTraceResponseData {
	s.LastTime = &v
	return s
}

func (s *ListPersonTraceResponseData) SetStartTime(v string) *ListPersonTraceResponseData {
	s.StartTime = &v
	return s
}

func (s *ListPersonTraceResponseData) SetEndSourceImage(v string) *ListPersonTraceResponseData {
	s.EndSourceImage = &v
	return s
}

func (s *ListPersonTraceResponseData) SetDeviceId(v string) *ListPersonTraceResponseData {
	s.DeviceId = &v
	return s
}

func (s *ListPersonTraceResponseData) SetStartTargetImage(v string) *ListPersonTraceResponseData {
	s.StartTargetImage = &v
	return s
}

func (s *ListPersonTraceResponseData) SetGroupId(v string) *ListPersonTraceResponseData {
	s.GroupId = &v
	return s
}

func (s *ListPersonTraceResponseData) SetPersonId(v string) *ListPersonTraceResponseData {
	s.PersonId = &v
	return s
}

func (s *ListPersonTraceResponseData) SetStartSourceImage(v string) *ListPersonTraceResponseData {
	s.StartSourceImage = &v
	return s
}

func (s *ListPersonTraceResponseData) SetCorpId(v string) *ListPersonTraceResponseData {
	s.CorpId = &v
	return s
}

func (s *ListPersonTraceResponseData) SetEndTargetImage(v string) *ListPersonTraceResponseData {
	s.EndTargetImage = &v
	return s
}

type ListCorpGroupMetricsRequest struct {
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	TagCode     *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListCorpGroupMetricsResponse struct {
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Success    *string                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data       []*ListCorpGroupMetricsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListCorpGroupMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpGroupMetricsResponse) SetCode(v string) *ListCorpGroupMetricsResponse {
	s.Code = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetMessage(v string) *ListCorpGroupMetricsResponse {
	s.Message = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetRequestId(v string) *ListCorpGroupMetricsResponse {
	s.RequestId = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetPageNumber(v int) *ListCorpGroupMetricsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetPageSize(v int) *ListCorpGroupMetricsResponse {
	s.PageSize = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetTotalCount(v int) *ListCorpGroupMetricsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetSuccess(v string) *ListCorpGroupMetricsResponse {
	s.Success = &v
	return s
}

func (s *ListCorpGroupMetricsResponse) SetData(v []*ListCorpGroupMetricsResponseData) *ListCorpGroupMetricsResponse {
	s.Data = v
	return s
}

type ListCorpGroupMetricsResponseData struct {
	DateId        *string `json:"DateId,omitempty" xml:"DateId,omitempty" require:"true"`
	TagMetrics    *string `json:"TagMetrics,omitempty" xml:"TagMetrics,omitempty" require:"true"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TagValue      *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	CorpGroupId   *string `json:"CorpGroupId,omitempty" xml:"CorpGroupId,omitempty" require:"true"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty" require:"true"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty" require:"true"`
	PersonID      *string `json:"PersonID,omitempty" xml:"PersonID,omitempty" require:"true"`
}

func (s ListCorpGroupMetricsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupMetricsResponseData) GoString() string {
	return s.String()
}

func (s *ListCorpGroupMetricsResponseData) SetDateId(v string) *ListCorpGroupMetricsResponseData {
	s.DateId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetTagMetrics(v string) *ListCorpGroupMetricsResponseData {
	s.TagMetrics = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetTagCode(v string) *ListCorpGroupMetricsResponseData {
	s.TagCode = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetTagValue(v string) *ListCorpGroupMetricsResponseData {
	s.TagValue = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetCorpGroupId(v string) *ListCorpGroupMetricsResponseData {
	s.CorpGroupId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetCorpId(v string) *ListCorpGroupMetricsResponseData {
	s.CorpId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetDeviceGroupId(v string) *ListCorpGroupMetricsResponseData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetDeviceId(v string) *ListCorpGroupMetricsResponseData {
	s.DeviceId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetUserGroupId(v string) *ListCorpGroupMetricsResponseData {
	s.UserGroupId = &v
	return s
}

func (s *ListCorpGroupMetricsResponseData) SetPersonID(v string) *ListCorpGroupMetricsResponseData {
	s.PersonID = &v
	return s
}

type GetFaceModelResultRequest struct {
	PictureId      *string `json:"PictureId,omitempty" xml:"PictureId,omitempty" require:"true"`
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

type GetFaceModelResultResponse struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetFaceModelResultResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetFaceModelResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponse) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponse) SetCode(v string) *GetFaceModelResultResponse {
	s.Code = &v
	return s
}

func (s *GetFaceModelResultResponse) SetMessage(v string) *GetFaceModelResultResponse {
	s.Message = &v
	return s
}

func (s *GetFaceModelResultResponse) SetRequestId(v string) *GetFaceModelResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetFaceModelResultResponse) SetData(v *GetFaceModelResultResponseData) *GetFaceModelResultResponse {
	s.Data = v
	return s
}

type GetFaceModelResultResponseData struct {
	Records []*GetFaceModelResultResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s GetFaceModelResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponseData) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponseData) SetRecords(v []*GetFaceModelResultResponseDataRecords) *GetFaceModelResultResponseData {
	s.Records = v
	return s
}

type GetFaceModelResultResponseDataRecords struct {
	MustacheStyle              *string    `json:"MustacheStyle,omitempty" xml:"MustacheStyle,omitempty" require:"true"`
	FaceStyle                  *string    `json:"FaceStyle,omitempty" xml:"FaceStyle,omitempty" require:"true"`
	CapStyle                   *int       `json:"CapStyle,omitempty" xml:"CapStyle,omitempty" require:"true"`
	SkinColor                  *int       `json:"SkinColor,omitempty" xml:"SkinColor,omitempty" require:"true"`
	RightBottomY               *float32   `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	RightBottomX               *float32   `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	HairStyle                  *int       `json:"HairStyle,omitempty" xml:"HairStyle,omitempty" require:"true"`
	GenderCode                 *int       `json:"GenderCode,omitempty" xml:"GenderCode,omitempty" require:"true"`
	RespiratorColor            *int       `json:"RespiratorColor,omitempty" xml:"RespiratorColor,omitempty" require:"true"`
	EthicCode                  *int       `json:"EthicCode,omitempty" xml:"EthicCode,omitempty" require:"true"`
	AgeLowerLimit              *int       `json:"AgeLowerLimit,omitempty" xml:"AgeLowerLimit,omitempty" require:"true"`
	LeftTopY                   *float32   `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	LeftTopX                   *float32   `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	HairColor                  *int       `json:"HairColor,omitempty" xml:"HairColor,omitempty" require:"true"`
	AgeUpLimit                 *int       `json:"AgeUpLimit,omitempty" xml:"AgeUpLimit,omitempty" require:"true"`
	GlassStyle                 *int       `json:"GlassStyle,omitempty" xml:"GlassStyle,omitempty" require:"true"`
	GlassColor                 *int       `json:"GlassColor,omitempty" xml:"GlassColor,omitempty" require:"true"`
	CapColor                   *int       `json:"CapColor,omitempty" xml:"CapColor,omitempty" require:"true"`
	CapColorReliability        *string    `json:"CapColorReliability,omitempty" xml:"CapColorReliability,omitempty" require:"true"`
	RespiratorColorReliability *string    `json:"RespiratorColorReliability,omitempty" xml:"RespiratorColorReliability,omitempty" require:"true"`
	EthicCodeReliability       *string    `json:"EthicCodeReliability,omitempty" xml:"EthicCodeReliability,omitempty" require:"true"`
	GenderCodeReliability      *string    `json:"GenderCodeReliability,omitempty" xml:"GenderCodeReliability,omitempty" require:"true"`
	GlassColorReliability      *string    `json:"GlassColorReliability,omitempty" xml:"GlassColorReliability,omitempty" require:"true"`
	SkinColorReliability       *string    `json:"SkinColorReliability,omitempty" xml:"SkinColorReliability,omitempty" require:"true"`
	MustacheStyleReliability   *string    `json:"MustacheStyleReliability,omitempty" xml:"MustacheStyleReliability,omitempty" require:"true"`
	CapStyleReliability        *string    `json:"CapStyleReliability,omitempty" xml:"CapStyleReliability,omitempty" require:"true"`
	FaceStyleReliability       *string    `json:"FaceStyleReliability,omitempty" xml:"FaceStyleReliability,omitempty" require:"true"`
	GlassStyleReliability      *string    `json:"GlassStyleReliability,omitempty" xml:"GlassStyleReliability,omitempty" require:"true"`
	AgeUpLimitReliability      *string    `json:"AgeUpLimitReliability,omitempty" xml:"AgeUpLimitReliability,omitempty" require:"true"`
	HairStyleReliability       *string    `json:"HairStyleReliability,omitempty" xml:"HairStyleReliability,omitempty" require:"true"`
	AgeLowerLimitReliability   *string    `json:"AgeLowerLimitReliability,omitempty" xml:"AgeLowerLimitReliability,omitempty" require:"true"`
	HairColorReliability       *string    `json:"HairColorReliability,omitempty" xml:"HairColorReliability,omitempty" require:"true"`
	FeatureData                []*float32 `json:"FeatureData,omitempty" xml:"FeatureData,omitempty" require:"true" type:"Repeated"`
}

func (s GetFaceModelResultResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetFaceModelResultResponseDataRecords) GoString() string {
	return s.String()
}

func (s *GetFaceModelResultResponseDataRecords) SetMustacheStyle(v string) *GetFaceModelResultResponseDataRecords {
	s.MustacheStyle = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetFaceStyle(v string) *GetFaceModelResultResponseDataRecords {
	s.FaceStyle = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetCapStyle(v int) *GetFaceModelResultResponseDataRecords {
	s.CapStyle = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetSkinColor(v int) *GetFaceModelResultResponseDataRecords {
	s.SkinColor = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetRightBottomY(v float32) *GetFaceModelResultResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetRightBottomX(v float32) *GetFaceModelResultResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetHairStyle(v int) *GetFaceModelResultResponseDataRecords {
	s.HairStyle = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetGenderCode(v int) *GetFaceModelResultResponseDataRecords {
	s.GenderCode = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetRespiratorColor(v int) *GetFaceModelResultResponseDataRecords {
	s.RespiratorColor = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetEthicCode(v int) *GetFaceModelResultResponseDataRecords {
	s.EthicCode = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetAgeLowerLimit(v int) *GetFaceModelResultResponseDataRecords {
	s.AgeLowerLimit = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetLeftTopY(v float32) *GetFaceModelResultResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetLeftTopX(v float32) *GetFaceModelResultResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetHairColor(v int) *GetFaceModelResultResponseDataRecords {
	s.HairColor = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetAgeUpLimit(v int) *GetFaceModelResultResponseDataRecords {
	s.AgeUpLimit = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetGlassStyle(v int) *GetFaceModelResultResponseDataRecords {
	s.GlassStyle = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetGlassColor(v int) *GetFaceModelResultResponseDataRecords {
	s.GlassColor = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetCapColor(v int) *GetFaceModelResultResponseDataRecords {
	s.CapColor = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetCapColorReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.CapColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetRespiratorColorReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.RespiratorColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetEthicCodeReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.EthicCodeReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetGenderCodeReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.GenderCodeReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetGlassColorReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.GlassColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetSkinColorReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.SkinColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetMustacheStyleReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.MustacheStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetCapStyleReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.CapStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetFaceStyleReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.FaceStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetGlassStyleReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.GlassStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetAgeUpLimitReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.AgeUpLimitReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetHairStyleReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.HairStyleReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetAgeLowerLimitReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.AgeLowerLimitReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetHairColorReliability(v string) *GetFaceModelResultResponseDataRecords {
	s.HairColorReliability = &v
	return s
}

func (s *GetFaceModelResultResponseDataRecords) SetFeatureData(v []*float32) *GetFaceModelResultResponseDataRecords {
	s.FeatureData = v
	return s
}

type CreateCorpGroupRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
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

type CreateCorpGroupResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s CreateCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCorpGroupResponse) SetCode(v string) *CreateCorpGroupResponse {
	s.Code = &v
	return s
}

func (s *CreateCorpGroupResponse) SetMessage(v string) *CreateCorpGroupResponse {
	s.Message = &v
	return s
}

func (s *CreateCorpGroupResponse) SetRequestId(v string) *CreateCorpGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateCorpGroupResponse) SetSuccess(v bool) *CreateCorpGroupResponse {
	s.Success = &v
	return s
}

type ListCorpGroupsRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListCorpGroupsResponse struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListCorpGroupsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListCorpGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpGroupsResponse) SetCode(v string) *ListCorpGroupsResponse {
	s.Code = &v
	return s
}

func (s *ListCorpGroupsResponse) SetMessage(v string) *ListCorpGroupsResponse {
	s.Message = &v
	return s
}

func (s *ListCorpGroupsResponse) SetRequestId(v string) *ListCorpGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *ListCorpGroupsResponse) SetData(v *ListCorpGroupsResponseData) *ListCorpGroupsResponse {
	s.Data = v
	return s
}

type ListCorpGroupsResponseData struct {
	PageNumber *int64    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int64    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*string `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListCorpGroupsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpGroupsResponseData) GoString() string {
	return s.String()
}

func (s *ListCorpGroupsResponseData) SetPageNumber(v int64) *ListCorpGroupsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListCorpGroupsResponseData) SetPageSize(v int64) *ListCorpGroupsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListCorpGroupsResponseData) SetTotalCount(v int64) *ListCorpGroupsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListCorpGroupsResponseData) SetTotalPage(v int64) *ListCorpGroupsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListCorpGroupsResponseData) SetRecords(v []*string) *ListCorpGroupsResponseData {
	s.Records = v
	return s
}

type DeleteCorpGroupRequest struct {
	CorpId  *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
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

type DeleteCorpGroupResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeleteCorpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorpGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteCorpGroupResponse) SetCode(v string) *DeleteCorpGroupResponse {
	s.Code = &v
	return s
}

func (s *DeleteCorpGroupResponse) SetMessage(v string) *DeleteCorpGroupResponse {
	s.Message = &v
	return s
}

func (s *DeleteCorpGroupResponse) SetRequestId(v string) *DeleteCorpGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCorpGroupResponse) SetSuccess(v bool) *DeleteCorpGroupResponse {
	s.Success = &v
	return s
}

type InvokeMotorModelRequest struct {
	PicId   *string `json:"PicId,omitempty" xml:"PicId,omitempty" require:"true"`
	CorpId  *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
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

type InvokeMotorModelResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *InvokeMotorModelResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s InvokeMotorModelResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeMotorModelResponse) GoString() string {
	return s.String()
}

func (s *InvokeMotorModelResponse) SetCode(v string) *InvokeMotorModelResponse {
	s.Code = &v
	return s
}

func (s *InvokeMotorModelResponse) SetMessage(v string) *InvokeMotorModelResponse {
	s.Message = &v
	return s
}

func (s *InvokeMotorModelResponse) SetRequestId(v string) *InvokeMotorModelResponse {
	s.RequestId = &v
	return s
}

func (s *InvokeMotorModelResponse) SetData(v *InvokeMotorModelResponseData) *InvokeMotorModelResponse {
	s.Data = v
	return s
}

type InvokeMotorModelResponseData struct {
	StructList *string `json:"StructList,omitempty" xml:"StructList,omitempty" require:"true"`
}

func (s InvokeMotorModelResponseData) String() string {
	return tea.Prettify(s)
}

func (s InvokeMotorModelResponseData) GoString() string {
	return s.String()
}

func (s *InvokeMotorModelResponseData) SetStructList(v string) *InvokeMotorModelResponseData {
	s.StructList = &v
	return s
}

type GetDeviceConfigRequest struct {
	DeviceSn        *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty" require:"true"`
	DeviceTimeStamp *string `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty" require:"true"`
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

type GetDeviceConfigResponse struct {
	AudioEnable   *bool                             `json:"AudioEnable,omitempty" xml:"AudioEnable,omitempty" require:"true"`
	AudioFormat   *string                           `json:"AudioFormat,omitempty" xml:"AudioFormat,omitempty" require:"true"`
	BitRate       *string                           `json:"BitRate,omitempty" xml:"BitRate,omitempty" require:"true"`
	Code          *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DeviceAddress *string                           `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty" require:"true"`
	DeviceName    *string                           `json:"DeviceName,omitempty" xml:"DeviceName,omitempty" require:"true"`
	EncodeFormat  *string                           `json:"EncodeFormat,omitempty" xml:"EncodeFormat,omitempty" require:"true"`
	FrameRate     *string                           `json:"FrameRate,omitempty" xml:"FrameRate,omitempty" require:"true"`
	GovLength     *int                              `json:"GovLength,omitempty" xml:"GovLength,omitempty" require:"true"`
	Latitude      *string                           `json:"Latitude,omitempty" xml:"Latitude,omitempty" require:"true"`
	Longitude     *string                           `json:"Longitude,omitempty" xml:"Longitude,omitempty" require:"true"`
	Message       *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	OSDTimeEnable *string                           `json:"OSDTimeEnable,omitempty" xml:"OSDTimeEnable,omitempty" require:"true"`
	OSDTimeType   *string                           `json:"OSDTimeType,omitempty" xml:"OSDTimeType,omitempty" require:"true"`
	OSDTimeX      *string                           `json:"OSDTimeX,omitempty" xml:"OSDTimeX,omitempty" require:"true"`
	OSDTimeY      *string                           `json:"OSDTimeY,omitempty" xml:"OSDTimeY,omitempty" require:"true"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Resolution    *string                           `json:"Resolution,omitempty" xml:"Resolution,omitempty" require:"true"`
	RetryInterval *string                           `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty" require:"true"`
	DeviceId      *string                           `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	UserName      *string                           `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	PassWord      *string                           `json:"PassWord,omitempty" xml:"PassWord,omitempty" require:"true"`
	Protocol      *string                           `json:"Protocol,omitempty" xml:"Protocol,omitempty" require:"true"`
	ServerId      *string                           `json:"ServerId,omitempty" xml:"ServerId,omitempty" require:"true"`
	ServerPort    *string                           `json:"ServerPort,omitempty" xml:"ServerPort,omitempty" require:"true"`
	ServerIp      *string                           `json:"ServerIp,omitempty" xml:"ServerIp,omitempty" require:"true"`
	OSDList       []*GetDeviceConfigResponseOSDList `json:"OSDList,omitempty" xml:"OSDList,omitempty" require:"true" type:"Repeated"`
}

func (s GetDeviceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponse) SetAudioEnable(v bool) *GetDeviceConfigResponse {
	s.AudioEnable = &v
	return s
}

func (s *GetDeviceConfigResponse) SetAudioFormat(v string) *GetDeviceConfigResponse {
	s.AudioFormat = &v
	return s
}

func (s *GetDeviceConfigResponse) SetBitRate(v string) *GetDeviceConfigResponse {
	s.BitRate = &v
	return s
}

func (s *GetDeviceConfigResponse) SetCode(v string) *GetDeviceConfigResponse {
	s.Code = &v
	return s
}

func (s *GetDeviceConfigResponse) SetDeviceAddress(v string) *GetDeviceConfigResponse {
	s.DeviceAddress = &v
	return s
}

func (s *GetDeviceConfigResponse) SetDeviceName(v string) *GetDeviceConfigResponse {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceConfigResponse) SetEncodeFormat(v string) *GetDeviceConfigResponse {
	s.EncodeFormat = &v
	return s
}

func (s *GetDeviceConfigResponse) SetFrameRate(v string) *GetDeviceConfigResponse {
	s.FrameRate = &v
	return s
}

func (s *GetDeviceConfigResponse) SetGovLength(v int) *GetDeviceConfigResponse {
	s.GovLength = &v
	return s
}

func (s *GetDeviceConfigResponse) SetLatitude(v string) *GetDeviceConfigResponse {
	s.Latitude = &v
	return s
}

func (s *GetDeviceConfigResponse) SetLongitude(v string) *GetDeviceConfigResponse {
	s.Longitude = &v
	return s
}

func (s *GetDeviceConfigResponse) SetMessage(v string) *GetDeviceConfigResponse {
	s.Message = &v
	return s
}

func (s *GetDeviceConfigResponse) SetOSDTimeEnable(v string) *GetDeviceConfigResponse {
	s.OSDTimeEnable = &v
	return s
}

func (s *GetDeviceConfigResponse) SetOSDTimeType(v string) *GetDeviceConfigResponse {
	s.OSDTimeType = &v
	return s
}

func (s *GetDeviceConfigResponse) SetOSDTimeX(v string) *GetDeviceConfigResponse {
	s.OSDTimeX = &v
	return s
}

func (s *GetDeviceConfigResponse) SetOSDTimeY(v string) *GetDeviceConfigResponse {
	s.OSDTimeY = &v
	return s
}

func (s *GetDeviceConfigResponse) SetRequestId(v string) *GetDeviceConfigResponse {
	s.RequestId = &v
	return s
}

func (s *GetDeviceConfigResponse) SetResolution(v string) *GetDeviceConfigResponse {
	s.Resolution = &v
	return s
}

func (s *GetDeviceConfigResponse) SetRetryInterval(v string) *GetDeviceConfigResponse {
	s.RetryInterval = &v
	return s
}

func (s *GetDeviceConfigResponse) SetDeviceId(v string) *GetDeviceConfigResponse {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceConfigResponse) SetUserName(v string) *GetDeviceConfigResponse {
	s.UserName = &v
	return s
}

func (s *GetDeviceConfigResponse) SetPassWord(v string) *GetDeviceConfigResponse {
	s.PassWord = &v
	return s
}

func (s *GetDeviceConfigResponse) SetProtocol(v string) *GetDeviceConfigResponse {
	s.Protocol = &v
	return s
}

func (s *GetDeviceConfigResponse) SetServerId(v string) *GetDeviceConfigResponse {
	s.ServerId = &v
	return s
}

func (s *GetDeviceConfigResponse) SetServerPort(v string) *GetDeviceConfigResponse {
	s.ServerPort = &v
	return s
}

func (s *GetDeviceConfigResponse) SetServerIp(v string) *GetDeviceConfigResponse {
	s.ServerIp = &v
	return s
}

func (s *GetDeviceConfigResponse) SetOSDList(v []*GetDeviceConfigResponseOSDList) *GetDeviceConfigResponse {
	s.OSDList = v
	return s
}

type GetDeviceConfigResponseOSDList struct {
	Text     *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	LeftTopX *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
}

func (s GetDeviceConfigResponseOSDList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponseOSDList) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponseOSDList) SetText(v string) *GetDeviceConfigResponseOSDList {
	s.Text = &v
	return s
}

func (s *GetDeviceConfigResponseOSDList) SetLeftTopX(v string) *GetDeviceConfigResponseOSDList {
	s.LeftTopX = &v
	return s
}

func (s *GetDeviceConfigResponseOSDList) SetLeftTopY(v string) *GetDeviceConfigResponseOSDList {
	s.LeftTopY = &v
	return s
}

type SyncDeviceTimeRequest struct {
	DeviceSn        *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty" require:"true"`
	DeviceTimeStamp *string `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty" require:"true"`
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

type SyncDeviceTimeResponse struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	NTPServer     *string `json:"NTPServer,omitempty" xml:"NTPServer,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RetryInterval *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty" require:"true"`
	SyncInterval  *string `json:"SyncInterval,omitempty" xml:"SyncInterval,omitempty" require:"true"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
}

func (s SyncDeviceTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDeviceTimeResponse) GoString() string {
	return s.String()
}

func (s *SyncDeviceTimeResponse) SetCode(v string) *SyncDeviceTimeResponse {
	s.Code = &v
	return s
}

func (s *SyncDeviceTimeResponse) SetMessage(v string) *SyncDeviceTimeResponse {
	s.Message = &v
	return s
}

func (s *SyncDeviceTimeResponse) SetNTPServer(v string) *SyncDeviceTimeResponse {
	s.NTPServer = &v
	return s
}

func (s *SyncDeviceTimeResponse) SetRequestId(v string) *SyncDeviceTimeResponse {
	s.RequestId = &v
	return s
}

func (s *SyncDeviceTimeResponse) SetRetryInterval(v string) *SyncDeviceTimeResponse {
	s.RetryInterval = &v
	return s
}

func (s *SyncDeviceTimeResponse) SetSyncInterval(v string) *SyncDeviceTimeResponse {
	s.SyncInterval = &v
	return s
}

func (s *SyncDeviceTimeResponse) SetTimeStamp(v string) *SyncDeviceTimeResponse {
	s.TimeStamp = &v
	return s
}

type RegisterDeviceRequest struct {
	DeviceSn        *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty" require:"true"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ServerId        *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	DeviceTimeStamp *string `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty" require:"true"`
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

type RegisterDeviceResponse struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RetryInterval *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetCode(v string) *RegisterDeviceResponse {
	s.Code = &v
	return s
}

func (s *RegisterDeviceResponse) SetMessage(v string) *RegisterDeviceResponse {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponse) SetRequestId(v string) *RegisterDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceResponse) SetRetryInterval(v string) *RegisterDeviceResponse {
	s.RetryInterval = &v
	return s
}

type ReportDeviceCapacityRequest struct {
	Longitude        *string                                        `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude         *string                                        `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	AudioFormat      *string                                        `json:"AudioFormat,omitempty" xml:"AudioFormat,omitempty"`
	PresetNum        *string                                        `json:"PresetNum,omitempty" xml:"PresetNum,omitempty"`
	PTZCapacity      *string                                        `json:"PTZCapacity,omitempty" xml:"PTZCapacity,omitempty"`
	DeviceSn         *string                                        `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty" require:"true"`
	StreamCapacities []*ReportDeviceCapacityRequestStreamCapacities `json:"StreamCapacities,omitempty" xml:"StreamCapacities,omitempty" require:"true" type:"Repeated"`
	DeviceTimeStamp  *string                                        `json:"DeviceTimeStamp,omitempty" xml:"DeviceTimeStamp,omitempty" require:"true"`
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

func (s *ReportDeviceCapacityRequest) SetStreamCapacities(v []*ReportDeviceCapacityRequestStreamCapacities) *ReportDeviceCapacityRequest {
	s.StreamCapacities = v
	return s
}

func (s *ReportDeviceCapacityRequest) SetDeviceTimeStamp(v string) *ReportDeviceCapacityRequest {
	s.DeviceTimeStamp = &v
	return s
}

type ReportDeviceCapacityRequestStreamCapacities struct {
	EncodeFormat   *string `json:"EncodeFormat,omitempty" xml:"EncodeFormat,omitempty" require:"true"`
	Resolution     *string `json:"Resolution,omitempty" xml:"Resolution,omitempty" require:"true"`
	MaxFrameRate   *string `json:"MaxFrameRate,omitempty" xml:"MaxFrameRate,omitempty" require:"true"`
	MaxStream      *string `json:"MaxStream,omitempty" xml:"MaxStream,omitempty" require:"true"`
	BitrateRange   *string `json:"BitrateRange,omitempty" xml:"BitrateRange,omitempty" require:"true"`
	GovLengthRange *string `json:"GovLengthRange,omitempty" xml:"GovLengthRange,omitempty" require:"true"`
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

func (s *ReportDeviceCapacityRequestStreamCapacities) SetResolution(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.Resolution = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetMaxFrameRate(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.MaxFrameRate = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetMaxStream(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.MaxStream = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetBitrateRange(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.BitrateRange = &v
	return s
}

func (s *ReportDeviceCapacityRequestStreamCapacities) SetGovLengthRange(v string) *ReportDeviceCapacityRequestStreamCapacities {
	s.GovLengthRange = &v
	return s
}

type ReportDeviceCapacityResponse struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RetryInterval *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty" require:"true"`
}

func (s ReportDeviceCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceCapacityResponse) GoString() string {
	return s.String()
}

func (s *ReportDeviceCapacityResponse) SetCode(v string) *ReportDeviceCapacityResponse {
	s.Code = &v
	return s
}

func (s *ReportDeviceCapacityResponse) SetMessage(v string) *ReportDeviceCapacityResponse {
	s.Message = &v
	return s
}

func (s *ReportDeviceCapacityResponse) SetRequestId(v string) *ReportDeviceCapacityResponse {
	s.RequestId = &v
	return s
}

func (s *ReportDeviceCapacityResponse) SetRetryInterval(v string) *ReportDeviceCapacityResponse {
	s.RetryInterval = &v
	return s
}

type SaveVideoSummaryTaskVideoRequest struct {
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	SaveVideo *bool   `json:"SaveVideo,omitempty" xml:"SaveVideo,omitempty" require:"true"`
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

type SaveVideoSummaryTaskVideoResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SaveVideoSummaryTaskVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoSummaryTaskVideoResponse) GoString() string {
	return s.String()
}

func (s *SaveVideoSummaryTaskVideoResponse) SetCode(v string) *SaveVideoSummaryTaskVideoResponse {
	s.Code = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponse) SetData(v string) *SaveVideoSummaryTaskVideoResponse {
	s.Data = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponse) SetMessage(v string) *SaveVideoSummaryTaskVideoResponse {
	s.Message = &v
	return s
}

func (s *SaveVideoSummaryTaskVideoResponse) SetRequestId(v string) *SaveVideoSummaryTaskVideoResponse {
	s.RequestId = &v
	return s
}

type ListBodyAlgorithmResultsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty" require:"true"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListBodyAlgorithmResultsResponse struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListBodyAlgorithmResultsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListBodyAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponse) SetCode(v string) *ListBodyAlgorithmResultsResponse {
	s.Code = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponse) SetMessage(v string) *ListBodyAlgorithmResultsResponse {
	s.Message = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponse) SetRequestId(v string) *ListBodyAlgorithmResultsResponse {
	s.RequestId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponse) SetData(v *ListBodyAlgorithmResultsResponseData) *ListBodyAlgorithmResultsResponse {
	s.Data = v
	return s
}

type ListBodyAlgorithmResultsResponseData struct {
	PageNumber *int                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListBodyAlgorithmResultsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListBodyAlgorithmResultsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponseData) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponseData) SetPageNumber(v int) *ListBodyAlgorithmResultsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseData) SetPageSize(v int) *ListBodyAlgorithmResultsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseData) SetTotalCount(v int) *ListBodyAlgorithmResultsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseData) SetTotalPage(v int) *ListBodyAlgorithmResultsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseData) SetRecords(v []*ListBodyAlgorithmResultsResponseDataRecords) *ListBodyAlgorithmResultsResponseData {
	s.Records = v
	return s
}

type ListBodyAlgorithmResultsResponseDataRecords struct {
	CapStyle         *string  `json:"CapStyle,omitempty" xml:"CapStyle,omitempty" require:"true"`
	CorpId           *string  `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceId     *string  `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
	PersonId         *string  `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	GenderCode       *string  `json:"GenderCode,omitempty" xml:"GenderCode,omitempty" require:"true"`
	HairStyle        *string  `json:"HairStyle,omitempty" xml:"HairStyle,omitempty" require:"true"`
	LeftTopX         *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY         *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	MaxAge           *string  `json:"MaxAge,omitempty" xml:"MaxAge,omitempty" require:"true"`
	MinAge           *string  `json:"MinAge,omitempty" xml:"MinAge,omitempty" require:"true"`
	PicUrlPath       *string  `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty" require:"true"`
	RightBottomX     *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY     *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	ShotTime         *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	TargetPicUrlPath *string  `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty" require:"true"`
	CoatLength       *string  `json:"CoatLength,omitempty" xml:"CoatLength,omitempty" require:"true"`
	CoatStyle        *string  `json:"CoatStyle,omitempty" xml:"CoatStyle,omitempty" require:"true"`
	TrousersLength   *string  `json:"TrousersLength,omitempty" xml:"TrousersLength,omitempty" require:"true"`
	TrousersStyle    *string  `json:"TrousersStyle,omitempty" xml:"TrousersStyle,omitempty" require:"true"`
	CoatColor        *string  `json:"CoatColor,omitempty" xml:"CoatColor,omitempty" require:"true"`
	TrousersColor    *string  `json:"TrousersColor,omitempty" xml:"TrousersColor,omitempty" require:"true"`
	SourceId         *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty" require:"true"`
}

func (s ListBodyAlgorithmResultsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListBodyAlgorithmResultsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetCapStyle(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.CapStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetCorpId(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetDataSourceId(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetPersonId(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.PersonId = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetGenderCode(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.GenderCode = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetHairStyle(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.HairStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetLeftTopX(v float32) *ListBodyAlgorithmResultsResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetLeftTopY(v float32) *ListBodyAlgorithmResultsResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetMaxAge(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.MaxAge = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetMinAge(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.MinAge = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetPicUrlPath(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetRightBottomX(v float32) *ListBodyAlgorithmResultsResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetRightBottomY(v float32) *ListBodyAlgorithmResultsResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetShotTime(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetTargetPicUrlPath(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetCoatLength(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.CoatLength = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetCoatStyle(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.CoatStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetTrousersLength(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.TrousersLength = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetTrousersStyle(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.TrousersStyle = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetCoatColor(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.CoatColor = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetTrousersColor(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.TrousersColor = &v
	return s
}

func (s *ListBodyAlgorithmResultsResponseDataRecords) SetSourceId(v string) *ListBodyAlgorithmResultsResponseDataRecords {
	s.SourceId = &v
	return s
}

type AddDataSourceRequest struct {
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceName    *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty" require:"true"`
	DataSourceType    *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty" require:"true"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileRetentionDays *int    `json:"FileRetentionDays,omitempty" xml:"FileRetentionDays,omitempty"`
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

func (s *AddDataSourceRequest) SetFileRetentionDays(v int) *AddDataSourceRequest {
	s.FileRetentionDays = &v
	return s
}

type AddDataSourceResponse struct {
	Code    *string                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data    *AddDataSourceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponse) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponse) SetCode(v string) *AddDataSourceResponse {
	s.Code = &v
	return s
}

func (s *AddDataSourceResponse) SetMessage(v string) *AddDataSourceResponse {
	s.Message = &v
	return s
}

func (s *AddDataSourceResponse) SetData(v *AddDataSourceResponseData) *AddDataSourceResponse {
	s.Data = v
	return s
}

type AddDataSourceResponseData struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
	KafkaTopic   *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty" require:"true"`
	OssPath      *string `json:"OssPath,omitempty" xml:"OssPath,omitempty" require:"true"`
}

func (s AddDataSourceResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponseData) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseData) SetDataSourceId(v string) *AddDataSourceResponseData {
	s.DataSourceId = &v
	return s
}

func (s *AddDataSourceResponseData) SetKafkaTopic(v string) *AddDataSourceResponseData {
	s.KafkaTopic = &v
	return s
}

func (s *AddDataSourceResponseData) SetOssPath(v string) *AddDataSourceResponseData {
	s.OssPath = &v
	return s
}

type GetVideoComposeResultRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TaskRequestId *string `json:"TaskRequestId,omitempty" xml:"TaskRequestId,omitempty" require:"true"`
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

type GetVideoComposeResultResponse struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VideoUrl  *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s GetVideoComposeResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoComposeResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoComposeResultResponse) SetMessage(v string) *GetVideoComposeResultResponse {
	s.Message = &v
	return s
}

func (s *GetVideoComposeResultResponse) SetRequestId(v string) *GetVideoComposeResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetVideoComposeResultResponse) SetVideoUrl(v string) *GetVideoComposeResultResponse {
	s.VideoUrl = &v
	return s
}

func (s *GetVideoComposeResultResponse) SetCode(v string) *GetVideoComposeResultResponse {
	s.Code = &v
	return s
}

func (s *GetVideoComposeResultResponse) SetStatus(v string) *GetVideoComposeResultResponse {
	s.Status = &v
	return s
}

type CreateVideoComposeTaskRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty" require:"true"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ImageFileNames  *string `json:"ImageFileNames,omitempty" xml:"ImageFileNames,omitempty" require:"true"`
	AudioFileName   *string `json:"AudioFileName,omitempty" xml:"AudioFileName,omitempty" require:"true"`
	ImageParameters *string `json:"ImageParameters,omitempty" xml:"ImageParameters,omitempty" require:"true"`
	VideoFormat     *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	VideoFrameRate  *int    `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty"`
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

func (s *CreateVideoComposeTaskRequest) SetVideoFrameRate(v int) *CreateVideoComposeTaskRequest {
	s.VideoFrameRate = &v
	return s
}

type CreateVideoComposeTaskResponse struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty" require:"true"`
}

func (s CreateVideoComposeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoComposeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoComposeTaskResponse) SetCode(v string) *CreateVideoComposeTaskResponse {
	s.Code = &v
	return s
}

func (s *CreateVideoComposeTaskResponse) SetMessage(v string) *CreateVideoComposeTaskResponse {
	s.Message = &v
	return s
}

func (s *CreateVideoComposeTaskResponse) SetRequestId(v string) *CreateVideoComposeTaskResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVideoComposeTaskResponse) SetDomainName(v string) *CreateVideoComposeTaskResponse {
	s.DomainName = &v
	return s
}

func (s *CreateVideoComposeTaskResponse) SetBucketName(v string) *CreateVideoComposeTaskResponse {
	s.BucketName = &v
	return s
}

type DeleteDataSourceRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
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

type DeleteDataSourceResponse struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetCode(v string) *DeleteDataSourceResponse {
	s.Code = &v
	return s
}

func (s *DeleteDataSourceResponse) SetData(v string) *DeleteDataSourceResponse {
	s.Data = &v
	return s
}

func (s *DeleteDataSourceResponse) SetMessage(v string) *DeleteDataSourceResponse {
	s.Message = &v
	return s
}

type UploadFileRequest struct {
	FileType      *string `json:"FileType,omitempty" xml:"FileType,omitempty" require:"true"`
	MD5           *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	FileContent   *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	FileAliasName *string `json:"FileAliasName,omitempty" xml:"FileAliasName,omitempty"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
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

type UploadFileResponse struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *UploadFileResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s UploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponse) GoString() string {
	return s.String()
}

func (s *UploadFileResponse) SetCode(v string) *UploadFileResponse {
	s.Code = &v
	return s
}

func (s *UploadFileResponse) SetMessage(v string) *UploadFileResponse {
	s.Message = &v
	return s
}

func (s *UploadFileResponse) SetRequestId(v string) *UploadFileResponse {
	s.RequestId = &v
	return s
}

func (s *UploadFileResponse) SetData(v *UploadFileResponseData) *UploadFileResponse {
	s.Data = v
	return s
}

type UploadFileResponseData struct {
	Records []*UploadFileResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s UploadFileResponseData) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseData) GoString() string {
	return s.String()
}

func (s *UploadFileResponseData) SetRecords(v []*UploadFileResponseDataRecords) *UploadFileResponseData {
	s.Records = v
	return s
}

type UploadFileResponseDataRecords struct {
	OssPath  *string `json:"OssPath,omitempty" xml:"OssPath,omitempty" require:"true"`
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty" require:"true"`
}

func (s UploadFileResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseDataRecords) GoString() string {
	return s.String()
}

func (s *UploadFileResponseDataRecords) SetOssPath(v string) *UploadFileResponseDataRecords {
	s.OssPath = &v
	return s
}

func (s *UploadFileResponseDataRecords) SetSourceId(v string) *UploadFileResponseDataRecords {
	s.SourceId = &v
	return s
}

type ListEventAlgorithmResultsRequest struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListEventAlgorithmResultsResponse struct {
	Code        *string                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ExtendValue *string                                `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty" require:"true"`
	Data        *ListEventAlgorithmResultsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListEventAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponse) SetCode(v string) *ListEventAlgorithmResultsResponse {
	s.Code = &v
	return s
}

func (s *ListEventAlgorithmResultsResponse) SetMessage(v string) *ListEventAlgorithmResultsResponse {
	s.Message = &v
	return s
}

func (s *ListEventAlgorithmResultsResponse) SetRequestId(v string) *ListEventAlgorithmResultsResponse {
	s.RequestId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponse) SetExtendValue(v string) *ListEventAlgorithmResultsResponse {
	s.ExtendValue = &v
	return s
}

func (s *ListEventAlgorithmResultsResponse) SetData(v *ListEventAlgorithmResultsResponseData) *ListEventAlgorithmResultsResponse {
	s.Data = v
	return s
}

type ListEventAlgorithmResultsResponseData struct {
	PageNumber *int                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListEventAlgorithmResultsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListEventAlgorithmResultsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponseData) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponseData) SetPageNumber(v int) *ListEventAlgorithmResultsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseData) SetPageSize(v int) *ListEventAlgorithmResultsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseData) SetTotalCount(v int) *ListEventAlgorithmResultsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseData) SetTotalPage(v int) *ListEventAlgorithmResultsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseData) SetRecords(v []*ListEventAlgorithmResultsResponseDataRecords) *ListEventAlgorithmResultsResponseData {
	s.Records = v
	return s
}

type ListEventAlgorithmResultsResponseDataRecords struct {
	CapStyle         *string `json:"CapStyle,omitempty" xml:"CapStyle,omitempty" require:"true"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceId     *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
	EventType        *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	FaceCount        *string `json:"FaceCount,omitempty" xml:"FaceCount,omitempty" require:"true"`
	PicUrlPath       *string `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty" require:"true"`
	ShotTime         *string `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	TargetPicUrlPath *string `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty" require:"true"`
	RecordId         *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	ExtendValue      *string `json:"ExtendValue,omitempty" xml:"ExtendValue,omitempty" require:"true"`
	ExtendValueTwo   *string `json:"ExtendValueTwo,omitempty" xml:"ExtendValueTwo,omitempty" require:"true"`
	ExtendValueThree *string `json:"ExtendValueThree,omitempty" xml:"ExtendValueThree,omitempty" require:"true"`
}

func (s ListEventAlgorithmResultsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEventAlgorithmResultsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetCapStyle(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.CapStyle = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetCorpId(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetDataSourceId(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetEventType(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.EventType = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetFaceCount(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.FaceCount = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetPicUrlPath(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetShotTime(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetTargetPicUrlPath(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetRecordId(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.RecordId = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetExtendValue(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.ExtendValue = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetExtendValueTwo(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.ExtendValueTwo = &v
	return s
}

func (s *ListEventAlgorithmResultsResponseDataRecords) SetExtendValueThree(v string) *ListEventAlgorithmResultsResponseDataRecords {
	s.ExtendValueThree = &v
	return s
}

type DeleteVideoSummaryTaskRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
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

type DeleteVideoSummaryTaskResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVideoSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoSummaryTaskResponse) SetCode(v string) *DeleteVideoSummaryTaskResponse {
	s.Code = &v
	return s
}

func (s *DeleteVideoSummaryTaskResponse) SetData(v string) *DeleteVideoSummaryTaskResponse {
	s.Data = &v
	return s
}

func (s *DeleteVideoSummaryTaskResponse) SetMessage(v string) *DeleteVideoSummaryTaskResponse {
	s.Message = &v
	return s
}

func (s *DeleteVideoSummaryTaskResponse) SetRequestId(v string) *DeleteVideoSummaryTaskResponse {
	s.RequestId = &v
	return s
}

type GetVideoSummaryTaskResultRequest struct {
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
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

type GetVideoSummaryTaskResultResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetVideoSummaryTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoSummaryTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoSummaryTaskResultResponse) SetCode(v string) *GetVideoSummaryTaskResultResponse {
	s.Code = &v
	return s
}

func (s *GetVideoSummaryTaskResultResponse) SetData(v string) *GetVideoSummaryTaskResultResponse {
	s.Data = &v
	return s
}

func (s *GetVideoSummaryTaskResultResponse) SetMessage(v string) *GetVideoSummaryTaskResultResponse {
	s.Message = &v
	return s
}

func (s *GetVideoSummaryTaskResultResponse) SetRequestId(v string) *GetVideoSummaryTaskResultResponse {
	s.RequestId = &v
	return s
}

type CreateVideoSummaryTaskRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	StartTimeStamp   *int64  `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty" require:"true"`
	EndTimeStamp     *int64  `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty" require:"true"`
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

type CreateVideoSummaryTaskResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s CreateVideoSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoSummaryTaskResponse) SetCode(v string) *CreateVideoSummaryTaskResponse {
	s.Code = &v
	return s
}

func (s *CreateVideoSummaryTaskResponse) SetMessage(v string) *CreateVideoSummaryTaskResponse {
	s.Message = &v
	return s
}

func (s *CreateVideoSummaryTaskResponse) SetRequestId(v string) *CreateVideoSummaryTaskResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVideoSummaryTaskResponse) SetData(v string) *CreateVideoSummaryTaskResponse {
	s.Data = &v
	return s
}

type ListMotorAlgorithmResultsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty" require:"true"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListMotorAlgorithmResultsResponse struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListMotorAlgorithmResultsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListMotorAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponse) SetCode(v string) *ListMotorAlgorithmResultsResponse {
	s.Code = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponse) SetMessage(v string) *ListMotorAlgorithmResultsResponse {
	s.Message = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponse) SetRequestId(v string) *ListMotorAlgorithmResultsResponse {
	s.RequestId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponse) SetData(v *ListMotorAlgorithmResultsResponseData) *ListMotorAlgorithmResultsResponse {
	s.Data = v
	return s
}

type ListMotorAlgorithmResultsResponseData struct {
	PageNumber *int                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListMotorAlgorithmResultsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListMotorAlgorithmResultsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponseData) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponseData) SetPageNumber(v int) *ListMotorAlgorithmResultsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseData) SetPageSize(v int) *ListMotorAlgorithmResultsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseData) SetTotalCount(v int) *ListMotorAlgorithmResultsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseData) SetTotalPage(v int) *ListMotorAlgorithmResultsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseData) SetRecords(v []*ListMotorAlgorithmResultsResponseDataRecords) *ListMotorAlgorithmResultsResponseData {
	s.Records = v
	return s
}

type ListMotorAlgorithmResultsResponseDataRecords struct {
	CorpId           *string  `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceId     *string  `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
	LeftTopX         *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY         *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	MotorId          *string  `json:"MotorId,omitempty" xml:"MotorId,omitempty" require:"true"`
	PicUrlPath       *string  `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty" require:"true"`
	PlateNumber      *string  `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty" require:"true"`
	RightBottomX     *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY     *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	ShotTime         *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	TargetPicUrlPath *string  `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty" require:"true"`
	MotorStyle       *string  `json:"MotorStyle,omitempty" xml:"MotorStyle,omitempty" require:"true"`
	MotorModel       *string  `json:"MotorModel,omitempty" xml:"MotorModel,omitempty" require:"true"`
	MotorColor       *string  `json:"MotorColor,omitempty" xml:"MotorColor,omitempty" require:"true"`
	MotorClass       *string  `json:"MotorClass,omitempty" xml:"MotorClass,omitempty" require:"true"`
	MotorBrand       *string  `json:"MotorBrand,omitempty" xml:"MotorBrand,omitempty" require:"true"`
	PlateColor       *string  `json:"PlateColor,omitempty" xml:"PlateColor,omitempty" require:"true"`
	PlateClass       *string  `json:"PlateClass,omitempty" xml:"PlateClass,omitempty" require:"true"`
	SafetyBelt       *string  `json:"SafetyBelt,omitempty" xml:"SafetyBelt,omitempty" require:"true"`
	Calling          *string  `json:"Calling,omitempty" xml:"Calling,omitempty" require:"true"`
	SourceId         *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty" require:"true"`
}

func (s ListMotorAlgorithmResultsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListMotorAlgorithmResultsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetCorpId(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetDataSourceId(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetLeftTopX(v float32) *ListMotorAlgorithmResultsResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetLeftTopY(v float32) *ListMotorAlgorithmResultsResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetMotorId(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.MotorId = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetPicUrlPath(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetPlateNumber(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.PlateNumber = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetRightBottomX(v float32) *ListMotorAlgorithmResultsResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetRightBottomY(v float32) *ListMotorAlgorithmResultsResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetShotTime(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetTargetPicUrlPath(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetMotorStyle(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.MotorStyle = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetMotorModel(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.MotorModel = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetMotorColor(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.MotorColor = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetMotorClass(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.MotorClass = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetMotorBrand(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.MotorBrand = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetPlateColor(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.PlateColor = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetPlateClass(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.PlateClass = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetSafetyBelt(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.SafetyBelt = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetCalling(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.Calling = &v
	return s
}

func (s *ListMotorAlgorithmResultsResponseDataRecords) SetSourceId(v string) *ListMotorAlgorithmResultsResponseDataRecords {
	s.SourceId = &v
	return s
}

type ListFaceAlgorithmResultsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty" require:"true"`
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListFaceAlgorithmResultsResponse struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListFaceAlgorithmResultsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListFaceAlgorithmResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponse) SetCode(v string) *ListFaceAlgorithmResultsResponse {
	s.Code = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponse) SetMessage(v string) *ListFaceAlgorithmResultsResponse {
	s.Message = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponse) SetRequestId(v string) *ListFaceAlgorithmResultsResponse {
	s.RequestId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponse) SetData(v *ListFaceAlgorithmResultsResponseData) *ListFaceAlgorithmResultsResponse {
	s.Data = v
	return s
}

type ListFaceAlgorithmResultsResponseData struct {
	PageNumber *int                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListFaceAlgorithmResultsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListFaceAlgorithmResultsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponseData) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponseData) SetPageNumber(v int) *ListFaceAlgorithmResultsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseData) SetPageSize(v int) *ListFaceAlgorithmResultsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseData) SetTotalCount(v int) *ListFaceAlgorithmResultsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseData) SetTotalPage(v int) *ListFaceAlgorithmResultsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseData) SetRecords(v []*ListFaceAlgorithmResultsResponseDataRecords) *ListFaceAlgorithmResultsResponseData {
	s.Records = v
	return s
}

type ListFaceAlgorithmResultsResponseDataRecords struct {
	FaceId           *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty" require:"true"`
	CorpId           *string  `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	DataSourceId     *string  `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty" require:"true"`
	ShotTime         *string  `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	GenderCode       *string  `json:"GenderCode,omitempty" xml:"GenderCode,omitempty" require:"true"`
	MinAge           *string  `json:"MinAge,omitempty" xml:"MinAge,omitempty" require:"true"`
	MaxAge           *string  `json:"MaxAge,omitempty" xml:"MaxAge,omitempty" require:"true"`
	CapStyle         *string  `json:"CapStyle,omitempty" xml:"CapStyle,omitempty" require:"true"`
	HairStyle        *string  `json:"HairStyle,omitempty" xml:"HairStyle,omitempty" require:"true"`
	LeftTopX         *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY         *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	RightBottomX     *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY     *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	PicUrlPath       *string  `json:"PicUrlPath,omitempty" xml:"PicUrlPath,omitempty" require:"true"`
	TargetPicUrlPath *string  `json:"TargetPicUrlPath,omitempty" xml:"TargetPicUrlPath,omitempty" require:"true"`
	SourceId         *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty" require:"true"`
}

func (s ListFaceAlgorithmResultsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListFaceAlgorithmResultsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetFaceId(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.FaceId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetCorpId(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetDataSourceId(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.DataSourceId = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetShotTime(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.ShotTime = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetGenderCode(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.GenderCode = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetMinAge(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.MinAge = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetMaxAge(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.MaxAge = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetCapStyle(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.CapStyle = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetHairStyle(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.HairStyle = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetLeftTopX(v float32) *ListFaceAlgorithmResultsResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetLeftTopY(v float32) *ListFaceAlgorithmResultsResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetRightBottomX(v float32) *ListFaceAlgorithmResultsResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetRightBottomY(v float32) *ListFaceAlgorithmResultsResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetPicUrlPath(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.PicUrlPath = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetTargetPicUrlPath(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.TargetPicUrlPath = &v
	return s
}

func (s *ListFaceAlgorithmResultsResponseDataRecords) SetSourceId(v string) *ListFaceAlgorithmResultsResponseDataRecords {
	s.SourceId = &v
	return s
}

type ListMetricsRequest struct {
	CorpId        *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TagCode       *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	AggregateType *string `json:"AggregateType,omitempty" xml:"AggregateType,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

type ListMetricsResponse struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListMetricsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListMetricsResponse) SetCode(v string) *ListMetricsResponse {
	s.Code = &v
	return s
}

func (s *ListMetricsResponse) SetMessage(v string) *ListMetricsResponse {
	s.Message = &v
	return s
}

func (s *ListMetricsResponse) SetRequestId(v string) *ListMetricsResponse {
	s.RequestId = &v
	return s
}

func (s *ListMetricsResponse) SetData(v *ListMetricsResponseData) *ListMetricsResponse {
	s.Data = v
	return s
}

type ListMetricsResponseData struct {
	PageNumber *int                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListMetricsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListMetricsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseData) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseData) SetPageNumber(v int) *ListMetricsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListMetricsResponseData) SetPageSize(v int) *ListMetricsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListMetricsResponseData) SetTotalCount(v int) *ListMetricsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListMetricsResponseData) SetTotalPage(v int) *ListMetricsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListMetricsResponseData) SetRecords(v []*ListMetricsResponseDataRecords) *ListMetricsResponseData {
	s.Records = v
	return s
}

type ListMetricsResponseDataRecords struct {
	DateTime  *string `json:"DateTime,omitempty" xml:"DateTime,omitempty" require:"true"`
	TagCode   *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TagValue  *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	TagMetric *string `json:"TagMetric,omitempty" xml:"TagMetric,omitempty" require:"true"`
}

func (s ListMetricsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListMetricsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListMetricsResponseDataRecords) SetDateTime(v string) *ListMetricsResponseDataRecords {
	s.DateTime = &v
	return s
}

func (s *ListMetricsResponseDataRecords) SetTagCode(v string) *ListMetricsResponseDataRecords {
	s.TagCode = &v
	return s
}

func (s *ListMetricsResponseDataRecords) SetTagValue(v string) *ListMetricsResponseDataRecords {
	s.TagValue = &v
	return s
}

func (s *ListMetricsResponseDataRecords) SetTagMetric(v string) *ListMetricsResponseDataRecords {
	s.TagMetric = &v
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

type DeleteRecordsResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s DeleteRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordsResponse) SetCode(v string) *DeleteRecordsResponse {
	s.Code = &v
	return s
}

func (s *DeleteRecordsResponse) SetMessage(v string) *DeleteRecordsResponse {
	s.Message = &v
	return s
}

func (s *DeleteRecordsResponse) SetRequestId(v string) *DeleteRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordsResponse) SetData(v string) *DeleteRecordsResponse {
	s.Data = &v
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

type RecognizeFaceQualityResponse struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *RecognizeFaceQualityResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s RecognizeFaceQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponse) SetCode(v string) *RecognizeFaceQualityResponse {
	s.Code = &v
	return s
}

func (s *RecognizeFaceQualityResponse) SetMessage(v string) *RecognizeFaceQualityResponse {
	s.Message = &v
	return s
}

func (s *RecognizeFaceQualityResponse) SetRequestId(v string) *RecognizeFaceQualityResponse {
	s.RequestId = &v
	return s
}

func (s *RecognizeFaceQualityResponse) SetData(v *RecognizeFaceQualityResponseData) *RecognizeFaceQualityResponse {
	s.Data = v
	return s
}

type RecognizeFaceQualityResponseData struct {
	QualityScore *string                                     `json:"QualityScore,omitempty" xml:"QualityScore,omitempty" require:"true"`
	Description  *string                                     `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Attributes   *RecognizeFaceQualityResponseDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" require:"true" type:"Struct"`
}

func (s RecognizeFaceQualityResponseData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponseData) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponseData) SetQualityScore(v string) *RecognizeFaceQualityResponseData {
	s.QualityScore = &v
	return s
}

func (s *RecognizeFaceQualityResponseData) SetDescription(v string) *RecognizeFaceQualityResponseData {
	s.Description = &v
	return s
}

func (s *RecognizeFaceQualityResponseData) SetAttributes(v *RecognizeFaceQualityResponseDataAttributes) *RecognizeFaceQualityResponseData {
	s.Attributes = v
	return s
}

type RecognizeFaceQualityResponseDataAttributes struct {
	LeftTopX               *int    `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY               *int    `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	RightBottomX           *int    `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY           *int    `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	TargetImageStoragePath *string `json:"TargetImageStoragePath,omitempty" xml:"TargetImageStoragePath,omitempty" require:"true"`
	FaceStyle              *string `json:"FaceStyle,omitempty" xml:"FaceStyle,omitempty" require:"true"`
	FaceQuality            *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty" require:"true"`
	FaceScore              *string `json:"FaceScore,omitempty" xml:"FaceScore,omitempty" require:"true"`
}

func (s RecognizeFaceQualityResponseDataAttributes) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceQualityResponseDataAttributes) GoString() string {
	return s.String()
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetLeftTopX(v int) *RecognizeFaceQualityResponseDataAttributes {
	s.LeftTopX = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetLeftTopY(v int) *RecognizeFaceQualityResponseDataAttributes {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetRightBottomX(v int) *RecognizeFaceQualityResponseDataAttributes {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetRightBottomY(v int) *RecognizeFaceQualityResponseDataAttributes {
	s.RightBottomY = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetTargetImageStoragePath(v string) *RecognizeFaceQualityResponseDataAttributes {
	s.TargetImageStoragePath = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetFaceStyle(v string) *RecognizeFaceQualityResponseDataAttributes {
	s.FaceStyle = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetFaceQuality(v string) *RecognizeFaceQualityResponseDataAttributes {
	s.FaceQuality = &v
	return s
}

func (s *RecognizeFaceQualityResponseDataAttributes) SetFaceScore(v string) *RecognizeFaceQualityResponseDataAttributes {
	s.FaceScore = &v
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

type ListPersonsResponse struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListPersonsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListPersonsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponse) GoString() string {
	return s.String()
}

func (s *ListPersonsResponse) SetCode(v string) *ListPersonsResponse {
	s.Code = &v
	return s
}

func (s *ListPersonsResponse) SetMessage(v string) *ListPersonsResponse {
	s.Message = &v
	return s
}

func (s *ListPersonsResponse) SetRequestId(v string) *ListPersonsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPersonsResponse) SetData(v *ListPersonsResponseData) *ListPersonsResponse {
	s.Data = v
	return s
}

type ListPersonsResponseData struct {
	PageNo     *string                           `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize   *string                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *string                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *string                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListPersonsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListPersonsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseData) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseData) SetPageNo(v string) *ListPersonsResponseData {
	s.PageNo = &v
	return s
}

func (s *ListPersonsResponseData) SetPageSize(v string) *ListPersonsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListPersonsResponseData) SetTotalCount(v string) *ListPersonsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListPersonsResponseData) SetTotalPage(v string) *ListPersonsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListPersonsResponseData) SetRecords(v []*ListPersonsResponseDataRecords) *ListPersonsResponseData {
	s.Records = v
	return s
}

type ListPersonsResponseDataRecords struct {
	FirstAppearTime *string                                  `json:"FirstAppearTime,omitempty" xml:"FirstAppearTime,omitempty" require:"true"`
	PersonId        *string                                  `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	PicUrl          *string                                  `json:"PicUrl,omitempty" xml:"PicUrl,omitempty" require:"true"`
	TagList         []*ListPersonsResponseDataRecordsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPersonsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseDataRecords) SetFirstAppearTime(v string) *ListPersonsResponseDataRecords {
	s.FirstAppearTime = &v
	return s
}

func (s *ListPersonsResponseDataRecords) SetPersonId(v string) *ListPersonsResponseDataRecords {
	s.PersonId = &v
	return s
}

func (s *ListPersonsResponseDataRecords) SetPicUrl(v string) *ListPersonsResponseDataRecords {
	s.PicUrl = &v
	return s
}

func (s *ListPersonsResponseDataRecords) SetTagList(v []*ListPersonsResponseDataRecordsTagList) *ListPersonsResponseDataRecords {
	s.TagList = v
	return s
}

type ListPersonsResponseDataRecordsTagList struct {
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty" require:"true"`
	TagValue   *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	TagValueId *string `json:"TagValueId,omitempty" xml:"TagValueId,omitempty" require:"true"`
}

func (s ListPersonsResponseDataRecordsTagList) String() string {
	return tea.Prettify(s)
}

func (s ListPersonsResponseDataRecordsTagList) GoString() string {
	return s.String()
}

func (s *ListPersonsResponseDataRecordsTagList) SetTagCode(v string) *ListPersonsResponseDataRecordsTagList {
	s.TagCode = &v
	return s
}

func (s *ListPersonsResponseDataRecordsTagList) SetTagName(v string) *ListPersonsResponseDataRecordsTagList {
	s.TagName = &v
	return s
}

func (s *ListPersonsResponseDataRecordsTagList) SetTagValue(v string) *ListPersonsResponseDataRecordsTagList {
	s.TagValue = &v
	return s
}

func (s *ListPersonsResponseDataRecordsTagList) SetTagValueId(v string) *ListPersonsResponseDataRecordsTagList {
	s.TagValueId = &v
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

type GetPersonDetailResponse struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetPersonDetailResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetPersonDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponse) SetCode(v string) *GetPersonDetailResponse {
	s.Code = &v
	return s
}

func (s *GetPersonDetailResponse) SetMessage(v string) *GetPersonDetailResponse {
	s.Message = &v
	return s
}

func (s *GetPersonDetailResponse) SetRequestId(v string) *GetPersonDetailResponse {
	s.RequestId = &v
	return s
}

func (s *GetPersonDetailResponse) SetData(v *GetPersonDetailResponseData) *GetPersonDetailResponse {
	s.Data = v
	return s
}

type GetPersonDetailResponseData struct {
	PicUrl   *string                               `json:"PicUrl,omitempty" xml:"PicUrl,omitempty" require:"true"`
	PersonId *string                               `json:"PersonId,omitempty" xml:"PersonId,omitempty" require:"true"`
	TagList  []*GetPersonDetailResponseDataTagList `json:"TagList,omitempty" xml:"TagList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPersonDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponseData) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponseData) SetPicUrl(v string) *GetPersonDetailResponseData {
	s.PicUrl = &v
	return s
}

func (s *GetPersonDetailResponseData) SetPersonId(v string) *GetPersonDetailResponseData {
	s.PersonId = &v
	return s
}

func (s *GetPersonDetailResponseData) SetTagList(v []*GetPersonDetailResponseDataTagList) *GetPersonDetailResponseData {
	s.TagList = v
	return s
}

type GetPersonDetailResponseDataTagList struct {
	TagCode    *string `json:"TagCode,omitempty" xml:"TagCode,omitempty" require:"true"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty" require:"true"`
	TagValue   *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	TagValueId *string `json:"TagValueId,omitempty" xml:"TagValueId,omitempty" require:"true"`
}

func (s GetPersonDetailResponseDataTagList) String() string {
	return tea.Prettify(s)
}

func (s GetPersonDetailResponseDataTagList) GoString() string {
	return s.String()
}

func (s *GetPersonDetailResponseDataTagList) SetTagCode(v string) *GetPersonDetailResponseDataTagList {
	s.TagCode = &v
	return s
}

func (s *GetPersonDetailResponseDataTagList) SetTagName(v string) *GetPersonDetailResponseDataTagList {
	s.TagName = &v
	return s
}

func (s *GetPersonDetailResponseDataTagList) SetTagValue(v string) *GetPersonDetailResponseDataTagList {
	s.TagValue = &v
	return s
}

func (s *GetPersonDetailResponseDataTagList) SetTagValueId(v string) *GetPersonDetailResponseDataTagList {
	s.TagValueId = &v
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

type GetFaceOptionsResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      []*GetFaceOptionsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s GetFaceOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponse) SetCode(v string) *GetFaceOptionsResponse {
	s.Code = &v
	return s
}

func (s *GetFaceOptionsResponse) SetMessage(v string) *GetFaceOptionsResponse {
	s.Message = &v
	return s
}

func (s *GetFaceOptionsResponse) SetRequestId(v string) *GetFaceOptionsResponse {
	s.RequestId = &v
	return s
}

func (s *GetFaceOptionsResponse) SetData(v []*GetFaceOptionsResponseData) *GetFaceOptionsResponse {
	s.Data = v
	return s
}

type GetFaceOptionsResponseData struct {
	Key        *string                                 `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Name       *string                                 `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	OptionList []*GetFaceOptionsResponseDataOptionList `json:"OptionList,omitempty" xml:"OptionList,omitempty" require:"true" type:"Repeated"`
}

func (s GetFaceOptionsResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponseData) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponseData) SetKey(v string) *GetFaceOptionsResponseData {
	s.Key = &v
	return s
}

func (s *GetFaceOptionsResponseData) SetName(v string) *GetFaceOptionsResponseData {
	s.Name = &v
	return s
}

func (s *GetFaceOptionsResponseData) SetOptionList(v []*GetFaceOptionsResponseDataOptionList) *GetFaceOptionsResponseData {
	s.OptionList = v
	return s
}

type GetFaceOptionsResponseDataOptionList struct {
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetFaceOptionsResponseDataOptionList) String() string {
	return tea.Prettify(s)
}

func (s GetFaceOptionsResponseDataOptionList) GoString() string {
	return s.String()
}

func (s *GetFaceOptionsResponseDataOptionList) SetKey(v string) *GetFaceOptionsResponseDataOptionList {
	s.Key = &v
	return s
}

func (s *GetFaceOptionsResponseDataOptionList) SetName(v string) *GetFaceOptionsResponseDataOptionList {
	s.Name = &v
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

type GetBodyOptionsResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      []*GetBodyOptionsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBodyOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponse) SetCode(v string) *GetBodyOptionsResponse {
	s.Code = &v
	return s
}

func (s *GetBodyOptionsResponse) SetMessage(v string) *GetBodyOptionsResponse {
	s.Message = &v
	return s
}

func (s *GetBodyOptionsResponse) SetRequestId(v string) *GetBodyOptionsResponse {
	s.RequestId = &v
	return s
}

func (s *GetBodyOptionsResponse) SetData(v []*GetBodyOptionsResponseData) *GetBodyOptionsResponse {
	s.Data = v
	return s
}

type GetBodyOptionsResponseData struct {
	Key        *string                                 `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Name       *string                                 `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	OptionList []*GetBodyOptionsResponseDataOptionList `json:"OptionList,omitempty" xml:"OptionList,omitempty" require:"true" type:"Repeated"`
}

func (s GetBodyOptionsResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponseData) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponseData) SetKey(v string) *GetBodyOptionsResponseData {
	s.Key = &v
	return s
}

func (s *GetBodyOptionsResponseData) SetName(v string) *GetBodyOptionsResponseData {
	s.Name = &v
	return s
}

func (s *GetBodyOptionsResponseData) SetOptionList(v []*GetBodyOptionsResponseDataOptionList) *GetBodyOptionsResponseData {
	s.OptionList = v
	return s
}

type GetBodyOptionsResponseDataOptionList struct {
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetBodyOptionsResponseDataOptionList) String() string {
	return tea.Prettify(s)
}

func (s GetBodyOptionsResponseDataOptionList) GoString() string {
	return s.String()
}

func (s *GetBodyOptionsResponseDataOptionList) SetKey(v string) *GetBodyOptionsResponseDataOptionList {
	s.Key = &v
	return s
}

func (s *GetBodyOptionsResponseDataOptionList) SetName(v string) *GetBodyOptionsResponseDataOptionList {
	s.Name = &v
	return s
}

type StopMonitorRequest struct {
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	AlgorithmVendor *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty" require:"true"`
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

type StopMonitorResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMonitorResponse) GoString() string {
	return s.String()
}

func (s *StopMonitorResponse) SetCode(v string) *StopMonitorResponse {
	s.Code = &v
	return s
}

func (s *StopMonitorResponse) SetData(v string) *StopMonitorResponse {
	s.Data = &v
	return s
}

func (s *StopMonitorResponse) SetMessage(v string) *StopMonitorResponse {
	s.Message = &v
	return s
}

func (s *StopMonitorResponse) SetRequestId(v string) *StopMonitorResponse {
	s.RequestId = &v
	return s
}

type SearchBodyRequest struct {
	CorpId         *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	GbId           *string                `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp *int64                 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty" require:"true"`
	EndTimeStamp   *int64                 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty" require:"true"`
	PageNo         *int                   `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize       *int                   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *SearchBodyRequest) SetPageNo(v int) *SearchBodyRequest {
	s.PageNo = &v
	return s
}

func (s *SearchBodyRequest) SetPageSize(v int) *SearchBodyRequest {
	s.PageSize = &v
	return s
}

func (s *SearchBodyRequest) SetOptionList(v map[string]interface{}) *SearchBodyRequest {
	s.OptionList = v
	return s
}

type SearchBodyShrinkRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp   *int64  `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty" require:"true"`
	EndTimeStamp     *int64  `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty" require:"true"`
	PageNo           *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize         *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *SearchBodyShrinkRequest) SetPageNo(v int) *SearchBodyShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetPageSize(v int) *SearchBodyShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchBodyShrinkRequest) SetOptionListShrink(v string) *SearchBodyShrinkRequest {
	s.OptionListShrink = &v
	return s
}

type SearchBodyResponse struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SearchBodyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SearchBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponse) GoString() string {
	return s.String()
}

func (s *SearchBodyResponse) SetCode(v string) *SearchBodyResponse {
	s.Code = &v
	return s
}

func (s *SearchBodyResponse) SetMessage(v string) *SearchBodyResponse {
	s.Message = &v
	return s
}

func (s *SearchBodyResponse) SetRequestId(v string) *SearchBodyResponse {
	s.RequestId = &v
	return s
}

func (s *SearchBodyResponse) SetData(v *SearchBodyResponseData) *SearchBodyResponse {
	s.Data = v
	return s
}

type SearchBodyResponseData struct {
	PageNo     *int                             `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize   *int                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*SearchBodyResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s SearchBodyResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponseData) GoString() string {
	return s.String()
}

func (s *SearchBodyResponseData) SetPageNo(v int) *SearchBodyResponseData {
	s.PageNo = &v
	return s
}

func (s *SearchBodyResponseData) SetPageSize(v int) *SearchBodyResponseData {
	s.PageSize = &v
	return s
}

func (s *SearchBodyResponseData) SetTotalCount(v int) *SearchBodyResponseData {
	s.TotalCount = &v
	return s
}

func (s *SearchBodyResponseData) SetTotalPage(v int) *SearchBodyResponseData {
	s.TotalPage = &v
	return s
}

func (s *SearchBodyResponseData) SetRecords(v []*SearchBodyResponseDataRecords) *SearchBodyResponseData {
	s.Records = v
	return s
}

type SearchBodyResponseDataRecords struct {
	GbId           *string  `json:"GbId,omitempty" xml:"GbId,omitempty" require:"true"`
	ImageUrl       *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	LeftTopX       *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY       *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	RightBottomX   *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY   *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	TargetImageUrl *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty" require:"true"`
}

func (s SearchBodyResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchBodyResponseDataRecords) GoString() string {
	return s.String()
}

func (s *SearchBodyResponseDataRecords) SetGbId(v string) *SearchBodyResponseDataRecords {
	s.GbId = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetImageUrl(v string) *SearchBodyResponseDataRecords {
	s.ImageUrl = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetLeftTopX(v float32) *SearchBodyResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetLeftTopY(v float32) *SearchBodyResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetRightBottomX(v float32) *SearchBodyResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetRightBottomY(v float32) *SearchBodyResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetScore(v float32) *SearchBodyResponseDataRecords {
	s.Score = &v
	return s
}

func (s *SearchBodyResponseDataRecords) SetTargetImageUrl(v string) *SearchBodyResponseDataRecords {
	s.TargetImageUrl = &v
	return s
}

type AddMonitorRequest struct {
	CorpId               *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	MonitorType          *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty" require:"true"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	BatchIndicator       *int    `json:"BatchIndicator,omitempty" xml:"BatchIndicator,omitempty"`
	AlgorithmVendor      *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty" require:"true"`
	NotifierType         *string `json:"NotifierType,omitempty" xml:"NotifierType,omitempty"`
	NotifierUrl          *string `json:"NotifierUrl,omitempty" xml:"NotifierUrl,omitempty"`
	NotifierAppSecret    *string `json:"NotifierAppSecret,omitempty" xml:"NotifierAppSecret,omitempty"`
	NotifierTimeOut      *int    `json:"NotifierTimeOut,omitempty" xml:"NotifierTimeOut,omitempty"`
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

func (s *AddMonitorRequest) SetBatchIndicator(v int) *AddMonitorRequest {
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

func (s *AddMonitorRequest) SetNotifierTimeOut(v int) *AddMonitorRequest {
	s.NotifierTimeOut = &v
	return s
}

func (s *AddMonitorRequest) SetNotifierExtendValues(v string) *AddMonitorRequest {
	s.NotifierExtendValues = &v
	return s
}

type AddMonitorResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *AddMonitorResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddMonitorResponse) SetRequestId(v string) *AddMonitorResponse {
	s.RequestId = &v
	return s
}

func (s *AddMonitorResponse) SetCode(v string) *AddMonitorResponse {
	s.Code = &v
	return s
}

func (s *AddMonitorResponse) SetMessage(v string) *AddMonitorResponse {
	s.Message = &v
	return s
}

func (s *AddMonitorResponse) SetData(v *AddMonitorResponseData) *AddMonitorResponse {
	s.Data = v
	return s
}

type AddMonitorResponseData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s AddMonitorResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddMonitorResponseData) GoString() string {
	return s.String()
}

func (s *AddMonitorResponseData) SetTaskId(v string) *AddMonitorResponseData {
	s.TaskId = &v
	return s
}

type GetMonitorResultRequest struct {
	CorpId          *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	MinRecordId     *string `json:"MinRecordId,omitempty" xml:"MinRecordId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
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

type GetMonitorResultResponse struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetMonitorResultResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetMonitorResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponse) SetCode(v string) *GetMonitorResultResponse {
	s.Code = &v
	return s
}

func (s *GetMonitorResultResponse) SetMessage(v string) *GetMonitorResultResponse {
	s.Message = &v
	return s
}

func (s *GetMonitorResultResponse) SetRequestId(v string) *GetMonitorResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetMonitorResultResponse) SetData(v *GetMonitorResultResponseData) *GetMonitorResultResponse {
	s.Data = v
	return s
}

type GetMonitorResultResponseData struct {
	MaxId   *string                                `json:"MaxId,omitempty" xml:"MaxId,omitempty" require:"true"`
	Records []*GetMonitorResultResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s GetMonitorResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseData) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseData) SetMaxId(v string) *GetMonitorResultResponseData {
	s.MaxId = &v
	return s
}

func (s *GetMonitorResultResponseData) SetRecords(v []*GetMonitorResultResponseDataRecords) *GetMonitorResultResponseData {
	s.Records = v
	return s
}

type GetMonitorResultResponseDataRecords struct {
	RightBottomY  *string                                        `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	RightBottomX  *string                                        `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	LeftUpY       *string                                        `json:"LeftUpY,omitempty" xml:"LeftUpY,omitempty" require:"true"`
	LeftUpX       *string                                        `json:"LeftUpX,omitempty" xml:"LeftUpX,omitempty" require:"true"`
	GbId          *string                                        `json:"GbId,omitempty" xml:"GbId,omitempty" require:"true"`
	Score         *string                                        `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	PicUrl        *string                                        `json:"PicUrl,omitempty" xml:"PicUrl,omitempty" require:"true"`
	ShotTime      *string                                        `json:"ShotTime,omitempty" xml:"ShotTime,omitempty" require:"true"`
	MonitorPicUrl *string                                        `json:"MonitorPicUrl,omitempty" xml:"MonitorPicUrl,omitempty" require:"true"`
	TargetPicUrl  *string                                        `json:"TargetPicUrl,omitempty" xml:"TargetPicUrl,omitempty" require:"true"`
	TaskId        *string                                        `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	ExtendInfo    *GetMonitorResultResponseDataRecordsExtendInfo `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetMonitorResultResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseDataRecords) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseDataRecords) SetRightBottomY(v string) *GetMonitorResultResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetRightBottomX(v string) *GetMonitorResultResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetLeftUpY(v string) *GetMonitorResultResponseDataRecords {
	s.LeftUpY = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetLeftUpX(v string) *GetMonitorResultResponseDataRecords {
	s.LeftUpX = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetGbId(v string) *GetMonitorResultResponseDataRecords {
	s.GbId = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetScore(v string) *GetMonitorResultResponseDataRecords {
	s.Score = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetPicUrl(v string) *GetMonitorResultResponseDataRecords {
	s.PicUrl = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetShotTime(v string) *GetMonitorResultResponseDataRecords {
	s.ShotTime = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetMonitorPicUrl(v string) *GetMonitorResultResponseDataRecords {
	s.MonitorPicUrl = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetTargetPicUrl(v string) *GetMonitorResultResponseDataRecords {
	s.TargetPicUrl = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetTaskId(v string) *GetMonitorResultResponseDataRecords {
	s.TaskId = &v
	return s
}

func (s *GetMonitorResultResponseDataRecords) SetExtendInfo(v *GetMonitorResultResponseDataRecordsExtendInfo) *GetMonitorResultResponseDataRecords {
	s.ExtendInfo = v
	return s
}

type GetMonitorResultResponseDataRecordsExtendInfo struct {
	PlateNo *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty" require:"true"`
}

func (s GetMonitorResultResponseDataRecordsExtendInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorResultResponseDataRecordsExtendInfo) GoString() string {
	return s.String()
}

func (s *GetMonitorResultResponseDataRecordsExtendInfo) SetPlateNo(v string) *GetMonitorResultResponseDataRecordsExtendInfo {
	s.PlateNo = &v
	return s
}

type UpdateMonitorRequest struct {
	CorpId               *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
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
	AlgorithmVendor      *string `json:"AlgorithmVendor,omitempty" xml:"AlgorithmVendor,omitempty" require:"true"`
	NotifierType         *string `json:"NotifierType,omitempty" xml:"NotifierType,omitempty"`
	NotifierUrl          *string `json:"NotifierUrl,omitempty" xml:"NotifierUrl,omitempty"`
	NotifierAppSecret    *string `json:"NotifierAppSecret,omitempty" xml:"NotifierAppSecret,omitempty"`
	NotifierTimeOut      *int    `json:"NotifierTimeOut,omitempty" xml:"NotifierTimeOut,omitempty"`
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

func (s *UpdateMonitorRequest) SetNotifierTimeOut(v int) *UpdateMonitorRequest {
	s.NotifierTimeOut = &v
	return s
}

func (s *UpdateMonitorRequest) SetNotifierExtendValues(v string) *UpdateMonitorRequest {
	s.NotifierExtendValues = &v
	return s
}

type UpdateMonitorResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorResponse) SetCode(v string) *UpdateMonitorResponse {
	s.Code = &v
	return s
}

func (s *UpdateMonitorResponse) SetData(v string) *UpdateMonitorResponse {
	s.Data = &v
	return s
}

func (s *UpdateMonitorResponse) SetMessage(v string) *UpdateMonitorResponse {
	s.Message = &v
	return s
}

func (s *UpdateMonitorResponse) SetRequestId(v string) *UpdateMonitorResponse {
	s.RequestId = &v
	return s
}

type GetDeviceVideoUrlRequest struct {
	CorpId      *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
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

type GetDeviceVideoUrlResponse struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty" require:"true"`
}

func (s GetDeviceVideoUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceVideoUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceVideoUrlResponse) SetCode(v string) *GetDeviceVideoUrlResponse {
	s.Code = &v
	return s
}

func (s *GetDeviceVideoUrlResponse) SetMessage(v string) *GetDeviceVideoUrlResponse {
	s.Message = &v
	return s
}

func (s *GetDeviceVideoUrlResponse) SetRequestId(v string) *GetDeviceVideoUrlResponse {
	s.RequestId = &v
	return s
}

func (s *GetDeviceVideoUrlResponse) SetUrl(v string) *GetDeviceVideoUrlResponse {
	s.Url = &v
	return s
}

func (s *GetDeviceVideoUrlResponse) SetOutProtocol(v string) *GetDeviceVideoUrlResponse {
	s.OutProtocol = &v
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

type GetInventoryResponse struct {
	Success *bool                     `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data    *GetInventoryResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponse) GoString() string {
	return s.String()
}

func (s *GetInventoryResponse) SetSuccess(v bool) *GetInventoryResponse {
	s.Success = &v
	return s
}

func (s *GetInventoryResponse) SetData(v *GetInventoryResponseData) *GetInventoryResponse {
	s.Data = v
	return s
}

type GetInventoryResponseData struct {
	ResultObject []*GetInventoryResponseDataResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Repeated"`
}

func (s GetInventoryResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponseData) GoString() string {
	return s.String()
}

func (s *GetInventoryResponseData) SetResultObject(v []*GetInventoryResponseDataResultObject) *GetInventoryResponseData {
	s.ResultObject = v
	return s
}

type GetInventoryResponseDataResultObject struct {
	BuyerId          *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty" require:"true"`
	CommodityCode    *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty" require:"true"`
	CurrentInventory *string `json:"CurrentInventory,omitempty" xml:"CurrentInventory,omitempty" require:"true"`
	ValidEndTime     *string `json:"ValidEndTime,omitempty" xml:"ValidEndTime,omitempty" require:"true"`
	ValidStartTime   *string `json:"ValidStartTime,omitempty" xml:"ValidStartTime,omitempty" require:"true"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InventoryId      *string `json:"InventoryId,omitempty" xml:"InventoryId,omitempty" require:"true"`
}

func (s GetInventoryResponseDataResultObject) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryResponseDataResultObject) GoString() string {
	return s.String()
}

func (s *GetInventoryResponseDataResultObject) SetBuyerId(v string) *GetInventoryResponseDataResultObject {
	s.BuyerId = &v
	return s
}

func (s *GetInventoryResponseDataResultObject) SetCommodityCode(v string) *GetInventoryResponseDataResultObject {
	s.CommodityCode = &v
	return s
}

func (s *GetInventoryResponseDataResultObject) SetCurrentInventory(v string) *GetInventoryResponseDataResultObject {
	s.CurrentInventory = &v
	return s
}

func (s *GetInventoryResponseDataResultObject) SetValidEndTime(v string) *GetInventoryResponseDataResultObject {
	s.ValidEndTime = &v
	return s
}

func (s *GetInventoryResponseDataResultObject) SetValidStartTime(v string) *GetInventoryResponseDataResultObject {
	s.ValidStartTime = &v
	return s
}

func (s *GetInventoryResponseDataResultObject) SetInstanceId(v string) *GetInventoryResponseDataResultObject {
	s.InstanceId = &v
	return s
}

func (s *GetInventoryResponseDataResultObject) SetInventoryId(v string) *GetInventoryResponseDataResultObject {
	s.InventoryId = &v
	return s
}

type RecognizeImageRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	PicContent *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	PicFormat  *string `json:"PicFormat,omitempty" xml:"PicFormat,omitempty" require:"true"`
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

type RecognizeImageResponse struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *RecognizeImageResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s RecognizeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponse) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponse) SetCode(v string) *RecognizeImageResponse {
	s.Code = &v
	return s
}

func (s *RecognizeImageResponse) SetMessage(v string) *RecognizeImageResponse {
	s.Message = &v
	return s
}

func (s *RecognizeImageResponse) SetRequestId(v string) *RecognizeImageResponse {
	s.RequestId = &v
	return s
}

func (s *RecognizeImageResponse) SetData(v *RecognizeImageResponseData) *RecognizeImageResponse {
	s.Data = v
	return s
}

type RecognizeImageResponseData struct {
	BodyList []*RecognizeImageResponseDataBodyList `json:"BodyList,omitempty" xml:"BodyList,omitempty" require:"true" type:"Repeated"`
	FaceList []*RecognizeImageResponseDataFaceList `json:"FaceList,omitempty" xml:"FaceList,omitempty" require:"true" type:"Repeated"`
}

func (s RecognizeImageResponseData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseData) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseData) SetBodyList(v []*RecognizeImageResponseDataBodyList) *RecognizeImageResponseData {
	s.BodyList = v
	return s
}

func (s *RecognizeImageResponseData) SetFaceList(v []*RecognizeImageResponseDataFaceList) *RecognizeImageResponseData {
	s.FaceList = v
	return s
}

type RecognizeImageResponseDataBodyList struct {
	Feature          *string `json:"Feature,omitempty" xml:"Feature,omitempty" require:"true"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	ImageBaseSixFour *string `json:"ImageBaseSixFour,omitempty" xml:"ImageBaseSixFour,omitempty" require:"true"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	LocalFeature     *string `json:"LocalFeature,omitempty" xml:"LocalFeature,omitempty" require:"true"`
	RespiratorColor  *string `json:"RespiratorColor,omitempty" xml:"RespiratorColor,omitempty" require:"true"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
}

func (s RecognizeImageResponseDataBodyList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseDataBodyList) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseDataBodyList) SetFeature(v string) *RecognizeImageResponseDataBodyList {
	s.Feature = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetFileName(v string) *RecognizeImageResponseDataBodyList {
	s.FileName = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetImageBaseSixFour(v string) *RecognizeImageResponseDataBodyList {
	s.ImageBaseSixFour = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetLeftTopX(v string) *RecognizeImageResponseDataBodyList {
	s.LeftTopX = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetLeftTopY(v string) *RecognizeImageResponseDataBodyList {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetLocalFeature(v string) *RecognizeImageResponseDataBodyList {
	s.LocalFeature = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetRespiratorColor(v string) *RecognizeImageResponseDataBodyList {
	s.RespiratorColor = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetRightBottomX(v string) *RecognizeImageResponseDataBodyList {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeImageResponseDataBodyList) SetRightBottomY(v string) *RecognizeImageResponseDataBodyList {
	s.RightBottomY = &v
	return s
}

type RecognizeImageResponseDataFaceList struct {
	Feature          *string `json:"Feature,omitempty" xml:"Feature,omitempty" require:"true"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	ImageBaseSixFour *string `json:"ImageBaseSixFour,omitempty" xml:"ImageBaseSixFour,omitempty" require:"true"`
	LeftTopX         *string `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY         *string `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	LocalFeature     *string `json:"LocalFeature,omitempty" xml:"LocalFeature,omitempty" require:"true"`
	RespiratorColor  *string `json:"RespiratorColor,omitempty" xml:"RespiratorColor,omitempty" require:"true"`
	RightBottomX     *string `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY     *string `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
}

func (s RecognizeImageResponseDataFaceList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageResponseDataFaceList) GoString() string {
	return s.String()
}

func (s *RecognizeImageResponseDataFaceList) SetFeature(v string) *RecognizeImageResponseDataFaceList {
	s.Feature = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetFileName(v string) *RecognizeImageResponseDataFaceList {
	s.FileName = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetImageBaseSixFour(v string) *RecognizeImageResponseDataFaceList {
	s.ImageBaseSixFour = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetLeftTopX(v string) *RecognizeImageResponseDataFaceList {
	s.LeftTopX = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetLeftTopY(v string) *RecognizeImageResponseDataFaceList {
	s.LeftTopY = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetLocalFeature(v string) *RecognizeImageResponseDataFaceList {
	s.LocalFeature = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetRespiratorColor(v string) *RecognizeImageResponseDataFaceList {
	s.RespiratorColor = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetRightBottomX(v string) *RecognizeImageResponseDataFaceList {
	s.RightBottomX = &v
	return s
}

func (s *RecognizeImageResponseDataFaceList) SetRightBottomY(v string) *RecognizeImageResponseDataFaceList {
	s.RightBottomY = &v
	return s
}

type ListCorpsRequest struct {
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CorpName   *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
}

func (s ListCorpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsRequest) GoString() string {
	return s.String()
}

func (s *ListCorpsRequest) SetPageNumber(v int) *ListCorpsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpsRequest) SetPageSize(v int) *ListCorpsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpsRequest) SetCorpName(v string) *ListCorpsRequest {
	s.CorpName = &v
	return s
}

type ListCorpsResponse struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListCorpsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListCorpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponse) GoString() string {
	return s.String()
}

func (s *ListCorpsResponse) SetCode(v string) *ListCorpsResponse {
	s.Code = &v
	return s
}

func (s *ListCorpsResponse) SetMessage(v string) *ListCorpsResponse {
	s.Message = &v
	return s
}

func (s *ListCorpsResponse) SetRequestId(v string) *ListCorpsResponse {
	s.RequestId = &v
	return s
}

func (s *ListCorpsResponse) SetData(v *ListCorpsResponseData) *ListCorpsResponse {
	s.Data = v
	return s
}

type ListCorpsResponseData struct {
	PageNumber *int                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListCorpsResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListCorpsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponseData) GoString() string {
	return s.String()
}

func (s *ListCorpsResponseData) SetPageNumber(v int) *ListCorpsResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListCorpsResponseData) SetPageSize(v int) *ListCorpsResponseData {
	s.PageSize = &v
	return s
}

func (s *ListCorpsResponseData) SetTotalCount(v int) *ListCorpsResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListCorpsResponseData) SetTotalPage(v int) *ListCorpsResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListCorpsResponseData) SetRecords(v []*ListCorpsResponseDataRecords) *ListCorpsResponseData {
	s.Records = v
	return s
}

type ListCorpsResponseDataRecords struct {
	CorpId       *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	CorpName     *string `json:"CorpName,omitempty" xml:"CorpName,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	CreateDate   *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	ParentCorpId *string `json:"ParentCorpId,omitempty" xml:"ParentCorpId,omitempty" require:"true"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	DeviceCount  *int    `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty" require:"true"`
	IsvSubId     *string `json:"IsvSubId,omitempty" xml:"IsvSubId,omitempty" require:"true"`
	AcuUsed      *int    `json:"AcuUsed,omitempty" xml:"AcuUsed,omitempty" require:"true"`
	IconPath     *string `json:"IconPath,omitempty" xml:"IconPath,omitempty" require:"true"`
}

func (s ListCorpsResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListCorpsResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListCorpsResponseDataRecords) SetCorpId(v string) *ListCorpsResponseDataRecords {
	s.CorpId = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetCorpName(v string) *ListCorpsResponseDataRecords {
	s.CorpName = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetDescription(v string) *ListCorpsResponseDataRecords {
	s.Description = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetCreateDate(v string) *ListCorpsResponseDataRecords {
	s.CreateDate = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetParentCorpId(v string) *ListCorpsResponseDataRecords {
	s.ParentCorpId = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetAppName(v string) *ListCorpsResponseDataRecords {
	s.AppName = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetDeviceCount(v int) *ListCorpsResponseDataRecords {
	s.DeviceCount = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetIsvSubId(v string) *ListCorpsResponseDataRecords {
	s.IsvSubId = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetAcuUsed(v int) *ListCorpsResponseDataRecords {
	s.AcuUsed = &v
	return s
}

func (s *ListCorpsResponseDataRecords) SetIconPath(v string) *ListCorpsResponseDataRecords {
	s.IconPath = &v
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

type UpdateCorpResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s UpdateCorpResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCorpResponse) GoString() string {
	return s.String()
}

func (s *UpdateCorpResponse) SetCode(v string) *UpdateCorpResponse {
	s.Code = &v
	return s
}

func (s *UpdateCorpResponse) SetMessage(v string) *UpdateCorpResponse {
	s.Message = &v
	return s
}

func (s *UpdateCorpResponse) SetRequestId(v string) *UpdateCorpResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateCorpResponse) SetData(v string) *UpdateCorpResponse {
	s.Data = &v
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

type UpdateDeviceResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s UpdateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResponse) SetCode(v string) *UpdateDeviceResponse {
	s.Code = &v
	return s
}

func (s *UpdateDeviceResponse) SetMessage(v string) *UpdateDeviceResponse {
	s.Message = &v
	return s
}

func (s *UpdateDeviceResponse) SetRequestId(v string) *UpdateDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceResponse) SetData(v string) *UpdateDeviceResponse {
	s.Data = &v
	return s
}

type ListDevicesRequest struct {
	CorpId     *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GbId       *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListDevicesRequest) SetPageNumber(v int) *ListDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesRequest) SetPageSize(v int) *ListDevicesRequest {
	s.PageSize = &v
	return s
}

type ListDevicesResponse struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListDevicesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) SetCode(v string) *ListDevicesResponse {
	s.Code = &v
	return s
}

func (s *ListDevicesResponse) SetMessage(v string) *ListDevicesResponse {
	s.Message = &v
	return s
}

func (s *ListDevicesResponse) SetRequestId(v string) *ListDevicesResponse {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponse) SetData(v *ListDevicesResponseData) *ListDevicesResponse {
	s.Data = v
	return s
}

type ListDevicesResponseData struct {
	PageNumber *int                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*ListDevicesResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s ListDevicesResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseData) SetPageNumber(v int) *ListDevicesResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesResponseData) SetPageSize(v int) *ListDevicesResponseData {
	s.PageSize = &v
	return s
}

func (s *ListDevicesResponseData) SetTotalCount(v int) *ListDevicesResponseData {
	s.TotalCount = &v
	return s
}

func (s *ListDevicesResponseData) SetTotalPage(v int) *ListDevicesResponseData {
	s.TotalPage = &v
	return s
}

func (s *ListDevicesResponseData) SetRecords(v []*ListDevicesResponseDataRecords) *ListDevicesResponseData {
	s.Records = v
	return s
}

type ListDevicesResponseDataRecords struct {
	AccessProtocolType *string `json:"AccessProtocolType,omitempty" xml:"AccessProtocolType,omitempty" require:"true"`
	BitRate            *string `json:"BitRate,omitempty" xml:"BitRate,omitempty" require:"true"`
	CoverImageUrl      *string `json:"CoverImageUrl,omitempty" xml:"CoverImageUrl,omitempty" require:"true"`
	GbId               *string `json:"GbId,omitempty" xml:"GbId,omitempty" require:"true"`
	DeviceAddress      *string `json:"DeviceAddress,omitempty" xml:"DeviceAddress,omitempty" require:"true"`
	DeviceDirection    *string `json:"DeviceDirection,omitempty" xml:"DeviceDirection,omitempty" require:"true"`
	DeviceSite         *string `json:"DeviceSite,omitempty" xml:"DeviceSite,omitempty" require:"true"`
	Latitude           *string `json:"Latitude,omitempty" xml:"Latitude,omitempty" require:"true"`
	Longitude          *string `json:"Longitude,omitempty" xml:"Longitude,omitempty" require:"true"`
	DeviceName         *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty" require:"true"`
	Resolution         *string `json:"Resolution,omitempty" xml:"Resolution,omitempty" require:"true"`
	SipGBId            *string `json:"SipGBId,omitempty" xml:"SipGBId,omitempty" require:"true"`
	SipPassword        *string `json:"SipPassword,omitempty" xml:"SipPassword,omitempty" require:"true"`
	SipServerIp        *string `json:"SipServerIp,omitempty" xml:"SipServerIp,omitempty" require:"true"`
	SipServerPort      *string `json:"SipServerPort,omitempty" xml:"SipServerPort,omitempty" require:"true"`
	Status             *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DeviceType         *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty" require:"true"`
	Vendor             *string `json:"Vendor,omitempty" xml:"Vendor,omitempty" require:"true"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s ListDevicesResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseDataRecords) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseDataRecords) SetAccessProtocolType(v string) *ListDevicesResponseDataRecords {
	s.AccessProtocolType = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetBitRate(v string) *ListDevicesResponseDataRecords {
	s.BitRate = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetCoverImageUrl(v string) *ListDevicesResponseDataRecords {
	s.CoverImageUrl = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetGbId(v string) *ListDevicesResponseDataRecords {
	s.GbId = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetDeviceAddress(v string) *ListDevicesResponseDataRecords {
	s.DeviceAddress = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetDeviceDirection(v string) *ListDevicesResponseDataRecords {
	s.DeviceDirection = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetDeviceSite(v string) *ListDevicesResponseDataRecords {
	s.DeviceSite = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetLatitude(v string) *ListDevicesResponseDataRecords {
	s.Latitude = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetLongitude(v string) *ListDevicesResponseDataRecords {
	s.Longitude = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetDeviceName(v string) *ListDevicesResponseDataRecords {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetResolution(v string) *ListDevicesResponseDataRecords {
	s.Resolution = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetSipGBId(v string) *ListDevicesResponseDataRecords {
	s.SipGBId = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetSipPassword(v string) *ListDevicesResponseDataRecords {
	s.SipPassword = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetSipServerIp(v string) *ListDevicesResponseDataRecords {
	s.SipServerIp = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetSipServerPort(v string) *ListDevicesResponseDataRecords {
	s.SipServerPort = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetStatus(v int) *ListDevicesResponseDataRecords {
	s.Status = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetDeviceType(v string) *ListDevicesResponseDataRecords {
	s.DeviceType = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetVendor(v string) *ListDevicesResponseDataRecords {
	s.Vendor = &v
	return s
}

func (s *ListDevicesResponseDataRecords) SetCreateTime(v string) *ListDevicesResponseDataRecords {
	s.CreateTime = &v
	return s
}

type GetDeviceLiveUrlRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	StreamType  *int    `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
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

func (s *GetDeviceLiveUrlRequest) SetStreamType(v int) *GetDeviceLiveUrlRequest {
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

type GetDeviceLiveUrlResponse struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty" require:"true"`
	StreamType  *int    `json:"StreamType,omitempty" xml:"StreamType,omitempty" require:"true"`
}

func (s GetDeviceLiveUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceLiveUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceLiveUrlResponse) SetCode(v string) *GetDeviceLiveUrlResponse {
	s.Code = &v
	return s
}

func (s *GetDeviceLiveUrlResponse) SetMessage(v string) *GetDeviceLiveUrlResponse {
	s.Message = &v
	return s
}

func (s *GetDeviceLiveUrlResponse) SetRequestId(v string) *GetDeviceLiveUrlResponse {
	s.RequestId = &v
	return s
}

func (s *GetDeviceLiveUrlResponse) SetUrl(v string) *GetDeviceLiveUrlResponse {
	s.Url = &v
	return s
}

func (s *GetDeviceLiveUrlResponse) SetOutProtocol(v string) *GetDeviceLiveUrlResponse {
	s.OutProtocol = &v
	return s
}

func (s *GetDeviceLiveUrlResponse) SetStreamType(v int) *GetDeviceLiveUrlResponse {
	s.StreamType = &v
	return s
}

type SearchFaceRequest struct {
	CorpId         *string                `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	GbId           *string                `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp *int64                 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty" require:"true"`
	EndTimeStamp   *int64                 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty" require:"true"`
	PageNo         *int                   `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize       *int                   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *SearchFaceRequest) SetPageNo(v int) *SearchFaceRequest {
	s.PageNo = &v
	return s
}

func (s *SearchFaceRequest) SetPageSize(v int) *SearchFaceRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaceRequest) SetOptionList(v map[string]interface{}) *SearchFaceRequest {
	s.OptionList = v
	return s
}

type SearchFaceShrinkRequest struct {
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	StartTimeStamp   *int64  `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty" require:"true"`
	EndTimeStamp     *int64  `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty" require:"true"`
	PageNo           *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize         *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *SearchFaceShrinkRequest) SetPageNo(v int) *SearchFaceShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetPageSize(v int) *SearchFaceShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaceShrinkRequest) SetOptionListShrink(v string) *SearchFaceShrinkRequest {
	s.OptionListShrink = &v
	return s
}

type SearchFaceResponse struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SearchFaceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SearchFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchFaceResponse) SetCode(v string) *SearchFaceResponse {
	s.Code = &v
	return s
}

func (s *SearchFaceResponse) SetMessage(v string) *SearchFaceResponse {
	s.Message = &v
	return s
}

func (s *SearchFaceResponse) SetRequestId(v string) *SearchFaceResponse {
	s.RequestId = &v
	return s
}

func (s *SearchFaceResponse) SetData(v *SearchFaceResponseData) *SearchFaceResponse {
	s.Data = v
	return s
}

type SearchFaceResponseData struct {
	PageNo     *int                             `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize   *int                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TotalPage  *int                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	Records    []*SearchFaceResponseDataRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s SearchFaceResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseData) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseData) SetPageNo(v int) *SearchFaceResponseData {
	s.PageNo = &v
	return s
}

func (s *SearchFaceResponseData) SetPageSize(v int) *SearchFaceResponseData {
	s.PageSize = &v
	return s
}

func (s *SearchFaceResponseData) SetTotalCount(v int) *SearchFaceResponseData {
	s.TotalCount = &v
	return s
}

func (s *SearchFaceResponseData) SetTotalPage(v int) *SearchFaceResponseData {
	s.TotalPage = &v
	return s
}

func (s *SearchFaceResponseData) SetRecords(v []*SearchFaceResponseDataRecords) *SearchFaceResponseData {
	s.Records = v
	return s
}

type SearchFaceResponseDataRecords struct {
	GbId            *string  `json:"GbId,omitempty" xml:"GbId,omitempty" require:"true"`
	ImageUrl        *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	LeftTopX        *float32 `json:"LeftTopX,omitempty" xml:"LeftTopX,omitempty" require:"true"`
	LeftTopY        *float32 `json:"LeftTopY,omitempty" xml:"LeftTopY,omitempty" require:"true"`
	MatchSuggestion *string  `json:"MatchSuggestion,omitempty" xml:"MatchSuggestion,omitempty" require:"true"`
	RightBottomX    *float32 `json:"RightBottomX,omitempty" xml:"RightBottomX,omitempty" require:"true"`
	RightBottomY    *float32 `json:"RightBottomY,omitempty" xml:"RightBottomY,omitempty" require:"true"`
	Score           *float32 `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	TargetImageUrl  *string  `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty" require:"true"`
	SourceId        *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty" require:"true"`
}

func (s SearchFaceResponseDataRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseDataRecords) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseDataRecords) SetGbId(v string) *SearchFaceResponseDataRecords {
	s.GbId = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetImageUrl(v string) *SearchFaceResponseDataRecords {
	s.ImageUrl = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetLeftTopX(v float32) *SearchFaceResponseDataRecords {
	s.LeftTopX = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetLeftTopY(v float32) *SearchFaceResponseDataRecords {
	s.LeftTopY = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetMatchSuggestion(v string) *SearchFaceResponseDataRecords {
	s.MatchSuggestion = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetRightBottomX(v float32) *SearchFaceResponseDataRecords {
	s.RightBottomX = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetRightBottomY(v float32) *SearchFaceResponseDataRecords {
	s.RightBottomY = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetScore(v float32) *SearchFaceResponseDataRecords {
	s.Score = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetTargetImageUrl(v string) *SearchFaceResponseDataRecords {
	s.TargetImageUrl = &v
	return s
}

func (s *SearchFaceResponseDataRecords) SetSourceId(v string) *SearchFaceResponseDataRecords {
	s.SourceId = &v
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

type AddDeviceResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s AddDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceResponse) SetCode(v string) *AddDeviceResponse {
	s.Code = &v
	return s
}

func (s *AddDeviceResponse) SetMessage(v string) *AddDeviceResponse {
	s.Message = &v
	return s
}

func (s *AddDeviceResponse) SetRequestId(v string) *AddDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *AddDeviceResponse) SetData(v string) *AddDeviceResponse {
	s.Data = &v
	return s
}

type CreateCorpRequest struct {
	CorpName      *string `json:"CorpName,omitempty" xml:"CorpName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
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

type CreateCorpResponse struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CorpId    *string `json:"CorpId,omitempty" xml:"CorpId,omitempty" require:"true"`
}

func (s CreateCorpResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpResponse) GoString() string {
	return s.String()
}

func (s *CreateCorpResponse) SetCode(v string) *CreateCorpResponse {
	s.Code = &v
	return s
}

func (s *CreateCorpResponse) SetMessage(v string) *CreateCorpResponse {
	s.Message = &v
	return s
}

func (s *CreateCorpResponse) SetRequestId(v string) *CreateCorpResponse {
	s.RequestId = &v
	return s
}

func (s *CreateCorpResponse) SetCorpId(v string) *CreateCorpResponse {
	s.CorpId = &v
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

type DeleteDeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s DeleteDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponse) SetRequestId(v string) *DeleteDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceResponse) SetCode(v string) *DeleteDeviceResponse {
	s.Code = &v
	return s
}

func (s *DeleteDeviceResponse) SetData(v string) *DeleteDeviceResponse {
	s.Data = &v
	return s
}

func (s *DeleteDeviceResponse) SetMessage(v string) *DeleteDeviceResponse {
	s.Message = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) GetMonitorListWithOptions(request *GetMonitorListRequest, runtime *util.RuntimeOptions) (_result *GetMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetMonitorListResponse{}
	_body, _err := client.DoRequest(tea.String("GetMonitorList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListDeviceGroupsWithOptions(request *ListDeviceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDeviceGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("ListDeviceGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SearchObjectWithOptions(tmp *SearchObjectRequest, runtime *util.RuntimeOptions) (_result *SearchObjectResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &SearchObjectShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.DeviceList)) {
		request.DeviceListShrink = util.ToJSONString(tmp.DeviceList)
	}

	if !tea.BoolValue(util.IsUnset(tmp.Conditions)) {
		request.ConditionsShrink = util.ToJSONString(tmp.Conditions)
	}

	if !tea.BoolValue(util.IsUnset(tmp.ImagePath)) {
		request.ImagePathShrink = util.ToJSONString(tmp.ImagePath)
	}

	_result = &SearchObjectResponse{}
	_body, _err := client.DoRequest(tea.String("SearchObject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeDevicesWithOptions(request *DescribeDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDevices"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetProfileListWithOptions(tmp *GetProfileListRequest, runtime *util.RuntimeOptions) (_result *GetProfileListResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &GetProfileListShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.PersonIdList)) {
		request.PersonIdListShrink = util.ToJSONString(tmp.PersonIdList)
	}

	if !tea.BoolValue(util.IsUnset(tmp.ProfileIdList)) {
		request.ProfileIdListShrink = util.ToJSONString(tmp.ProfileIdList)
	}

	_result = &GetProfileListResponse{}
	_body, _err := client.DoRequest(tea.String("GetProfileList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetProfileDetailWithOptions(request *GetProfileDetailRequest, runtime *util.RuntimeOptions) (_result *GetProfileDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetProfileDetailResponse{}
	_body, _err := client.DoRequest(tea.String("GetProfileDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteProfileCatalogWithOptions(request *DeleteProfileCatalogRequest, runtime *util.RuntimeOptions) (_result *DeleteProfileCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteProfileCatalogResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteProfileCatalog"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) BindPersonWithOptions(request *BindPersonRequest, runtime *util.RuntimeOptions) (_result *BindPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindPersonResponse{}
	_body, _err := client.DoRequest(tea.String("BindPerson"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateProfileWithOptions(request *UpdateProfileRequest, runtime *util.RuntimeOptions) (_result *UpdateProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateProfileResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UnbindPersonWithOptions(request *UnbindPersonRequest, runtime *util.RuntimeOptions) (_result *UnbindPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnbindPersonResponse{}
	_body, _err := client.DoRequest(tea.String("UnbindPerson"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AddProfileWithOptions(request *AddProfileRequest, runtime *util.RuntimeOptions) (_result *AddProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddProfileResponse{}
	_body, _err := client.DoRequest(tea.String("AddProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateProfileCatalogWithOptions(request *UpdateProfileCatalogRequest, runtime *util.RuntimeOptions) (_result *UpdateProfileCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateProfileCatalogResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateProfileCatalog"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AddProfileCatalogWithOptions(request *AddProfileCatalogRequest, runtime *util.RuntimeOptions) (_result *AddProfileCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddProfileCatalogResponse{}
	_body, _err := client.DoRequest(tea.String("AddProfileCatalog"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetCatalogListWithOptions(request *GetCatalogListRequest, runtime *util.RuntimeOptions) (_result *GetCatalogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetCatalogListResponse{}
	_body, _err := client.DoRequest(tea.String("GetCatalogList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteProfileWithOptions(request *DeleteProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteProfileResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UnbindCorpGroupWithOptions(request *UnbindCorpGroupRequest, runtime *util.RuntimeOptions) (_result *UnbindCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnbindCorpGroupResponse{}
	_body, _err := client.DoRequest(tea.String("UnbindCorpGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) BindCorpGroupWithOptions(request *BindCorpGroupRequest, runtime *util.RuntimeOptions) (_result *BindCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindCorpGroupResponse{}
	_body, _err := client.DoRequest(tea.String("BindCorpGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("ListUserGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetPersonListWithOptions(tmp *GetPersonListRequest, runtime *util.RuntimeOptions) (_result *GetPersonListResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &GetPersonListShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.CorpIdList)) {
		request.CorpIdListShrink = util.ToJSONString(tmp.CorpIdList)
	}

	if !tea.BoolValue(util.IsUnset(tmp.PersonIdList)) {
		request.PersonIdListShrink = util.ToJSONString(tmp.PersonIdList)
	}

	_result = &GetPersonListResponse{}
	_body, _err := client.DoRequest(tea.String("GetPersonList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListUsersWithOptions(tmp *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &ListUsersShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.PersonList)) {
		request.PersonListShrink = util.ToJSONString(tmp.PersonList)
	}

	if !tea.BoolValue(util.IsUnset(tmp.UserList)) {
		request.UserListShrink = util.ToJSONString(tmp.UserList)
	}

	_result = &ListUsersResponse{}
	_body, _err := client.DoRequest(tea.String("ListUsers"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateUserResponse{}
	_body, _err := client.DoRequest(tea.String("CreateUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) BindUserWithOptions(request *BindUserRequest, runtime *util.RuntimeOptions) (_result *BindUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindUserResponse{}
	_body, _err := client.DoRequest(tea.String("BindUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetUserDetailWithOptions(request *GetUserDetailRequest, runtime *util.RuntimeOptions) (_result *GetUserDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetUserDetailResponse{}
	_body, _err := client.DoRequest(tea.String("GetUserDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UploadImageWithOptions(request *UploadImageRequest, runtime *util.RuntimeOptions) (_result *UploadImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadImageResponse{}
	_body, _err := client.DoRequest(tea.String("UploadImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateUserGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateUserGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UnbindUserWithOptions(request *UnbindUserRequest, runtime *util.RuntimeOptions) (_result *UnbindUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnbindUserResponse{}
	_body, _err := client.DoRequest(tea.String("UnbindUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteUserGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListPersonVisitCountWithOptions(request *ListPersonVisitCountRequest, runtime *util.RuntimeOptions) (_result *ListPersonVisitCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPersonVisitCountResponse{}
	_body, _err := client.DoRequest(tea.String("ListPersonVisitCount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListEventAlgorithmDetailsWithOptions(request *ListEventAlgorithmDetailsRequest, runtime *util.RuntimeOptions) (_result *ListEventAlgorithmDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListEventAlgorithmDetailsResponse{}
	_body, _err := client.DoRequest(tea.String("ListEventAlgorithmDetails"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListCorpMetricsWithOptions(request *ListCorpMetricsRequest, runtime *util.RuntimeOptions) (_result *ListCorpMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListCorpMetricsResponse{}
	_body, _err := client.DoRequest(tea.String("ListCorpMetrics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListPersonTraceWithOptions(request *ListPersonTraceRequest, runtime *util.RuntimeOptions) (_result *ListPersonTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPersonTraceResponse{}
	_body, _err := client.DoRequest(tea.String("ListPersonTrace"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListCorpGroupMetricsWithOptions(request *ListCorpGroupMetricsRequest, runtime *util.RuntimeOptions) (_result *ListCorpGroupMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListCorpGroupMetricsResponse{}
	_body, _err := client.DoRequest(tea.String("ListCorpGroupMetrics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetFaceModelResultWithOptions(request *GetFaceModelResultRequest, runtime *util.RuntimeOptions) (_result *GetFaceModelResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFaceModelResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetFaceModelResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateCorpGroupWithOptions(request *CreateCorpGroupRequest, runtime *util.RuntimeOptions) (_result *CreateCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCorpGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateCorpGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListCorpGroupsWithOptions(request *ListCorpGroupsRequest, runtime *util.RuntimeOptions) (_result *ListCorpGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListCorpGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("ListCorpGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteCorpGroupWithOptions(request *DeleteCorpGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteCorpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCorpGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCorpGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) InvokeMotorModelWithOptions(request *InvokeMotorModelRequest, runtime *util.RuntimeOptions) (_result *InvokeMotorModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InvokeMotorModelResponse{}
	_body, _err := client.DoRequest(tea.String("InvokeMotorModel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetDeviceConfigWithOptions(request *GetDeviceConfigRequest, runtime *util.RuntimeOptions) (_result *GetDeviceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDeviceConfigResponse{}
	_body, _err := client.DoRequest(tea.String("GetDeviceConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SyncDeviceTimeWithOptions(request *SyncDeviceTimeRequest, runtime *util.RuntimeOptions) (_result *SyncDeviceTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SyncDeviceTimeResponse{}
	_body, _err := client.DoRequest(tea.String("SyncDeviceTime"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.DoRequest(tea.String("RegisterDevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ReportDeviceCapacityResponse{}
	_body, _err := client.DoRequest(tea.String("ReportDeviceCapacity"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &SaveVideoSummaryTaskVideoResponse{}
	_body, _err := client.DoRequest(tea.String("SaveVideoSummaryTaskVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListBodyAlgorithmResultsWithOptions(request *ListBodyAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListBodyAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListBodyAlgorithmResultsResponse{}
	_body, _err := client.DoRequest(tea.String("ListBodyAlgorithmResults"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AddDataSourceWithOptions(request *AddDataSourceRequest, runtime *util.RuntimeOptions) (_result *AddDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddDataSourceResponse{}
	_body, _err := client.DoRequest(tea.String("AddDataSource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetVideoComposeResultWithOptions(request *GetVideoComposeResultRequest, runtime *util.RuntimeOptions) (_result *GetVideoComposeResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetVideoComposeResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetVideoComposeResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateVideoComposeTaskWithOptions(request *CreateVideoComposeTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoComposeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVideoComposeTaskResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVideoComposeTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDataSource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UploadFileWithOptions(request *UploadFileRequest, runtime *util.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadFileResponse{}
	_body, _err := client.DoRequest(tea.String("UploadFile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListEventAlgorithmResultsWithOptions(request *ListEventAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListEventAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListEventAlgorithmResultsResponse{}
	_body, _err := client.DoRequest(tea.String("ListEventAlgorithmResults"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteVideoSummaryTaskWithOptions(request *DeleteVideoSummaryTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVideoSummaryTaskResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVideoSummaryTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetVideoSummaryTaskResultWithOptions(request *GetVideoSummaryTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetVideoSummaryTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetVideoSummaryTaskResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetVideoSummaryTaskResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateVideoSummaryTaskWithOptions(request *CreateVideoSummaryTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVideoSummaryTaskResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVideoSummaryTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListMotorAlgorithmResultsWithOptions(request *ListMotorAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListMotorAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListMotorAlgorithmResultsResponse{}
	_body, _err := client.DoRequest(tea.String("ListMotorAlgorithmResults"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListFaceAlgorithmResultsWithOptions(request *ListFaceAlgorithmResultsRequest, runtime *util.RuntimeOptions) (_result *ListFaceAlgorithmResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListFaceAlgorithmResultsResponse{}
	_body, _err := client.DoRequest(tea.String("ListFaceAlgorithmResults"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListMetricsResponse{}
	_body, _err := client.DoRequest(tea.String("ListMetrics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteRecordsWithOptions(request *DeleteRecordsRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) RecognizeFaceQualityWithOptions(request *RecognizeFaceQualityRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RecognizeFaceQualityResponse{}
	_body, _err := client.DoRequest(tea.String("RecognizeFaceQuality"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListPersonsWithOptions(request *ListPersonsRequest, runtime *util.RuntimeOptions) (_result *ListPersonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPersonsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPersons"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetPersonDetailWithOptions(request *GetPersonDetailRequest, runtime *util.RuntimeOptions) (_result *GetPersonDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPersonDetailResponse{}
	_body, _err := client.DoRequest(tea.String("GetPersonDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetFaceOptionsWithOptions(request *GetFaceOptionsRequest, runtime *util.RuntimeOptions) (_result *GetFaceOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFaceOptionsResponse{}
	_body, _err := client.DoRequest(tea.String("GetFaceOptions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetBodyOptionsWithOptions(request *GetBodyOptionsRequest, runtime *util.RuntimeOptions) (_result *GetBodyOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetBodyOptionsResponse{}
	_body, _err := client.DoRequest(tea.String("GetBodyOptions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) StopMonitorWithOptions(request *StopMonitorRequest, runtime *util.RuntimeOptions) (_result *StopMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopMonitorResponse{}
	_body, _err := client.DoRequest(tea.String("StopMonitor"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SearchBodyWithOptions(tmp *SearchBodyRequest, runtime *util.RuntimeOptions) (_result *SearchBodyResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &SearchBodyShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.OptionList)) {
		request.OptionListShrink = util.ToJSONString(tmp.OptionList)
	}

	_result = &SearchBodyResponse{}
	_body, _err := client.DoRequest(tea.String("SearchBody"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AddMonitorWithOptions(request *AddMonitorRequest, runtime *util.RuntimeOptions) (_result *AddMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddMonitorResponse{}
	_body, _err := client.DoRequest(tea.String("AddMonitor"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetMonitorResultWithOptions(request *GetMonitorResultRequest, runtime *util.RuntimeOptions) (_result *GetMonitorResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetMonitorResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetMonitorResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateMonitorWithOptions(request *UpdateMonitorRequest, runtime *util.RuntimeOptions) (_result *UpdateMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateMonitorResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateMonitor"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetDeviceVideoUrlWithOptions(request *GetDeviceVideoUrlRequest, runtime *util.RuntimeOptions) (_result *GetDeviceVideoUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDeviceVideoUrlResponse{}
	_body, _err := client.DoRequest(tea.String("GetDeviceVideoUrl"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetInventoryWithOptions(request *GetInventoryRequest, runtime *util.RuntimeOptions) (_result *GetInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetInventoryResponse{}
	_body, _err := client.DoRequest(tea.String("GetInventory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) RecognizeImageWithOptions(request *RecognizeImageRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RecognizeImageResponse{}
	_body, _err := client.DoRequest(tea.String("RecognizeImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListCorpsWithOptions(request *ListCorpsRequest, runtime *util.RuntimeOptions) (_result *ListCorpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListCorpsResponse{}
	_body, _err := client.DoRequest(tea.String("ListCorps"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK,APP"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateCorpWithOptions(request *UpdateCorpRequest, runtime *util.RuntimeOptions) (_result *UpdateCorpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateCorpResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateCorp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &UpdateDeviceResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateDevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.DoRequest(tea.String("ListDevices"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetDeviceLiveUrlWithOptions(request *GetDeviceLiveUrlRequest, runtime *util.RuntimeOptions) (_result *GetDeviceLiveUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDeviceLiveUrlResponse{}
	_body, _err := client.DoRequest(tea.String("GetDeviceLiveUrl"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SearchFaceWithOptions(tmp *SearchFaceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &SearchFaceShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.OptionList)) {
		request.OptionListShrink = util.ToJSONString(tmp.OptionList)
	}

	_result = &SearchFaceResponse{}
	_body, _err := client.DoRequest(tea.String("SearchFace"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AddDeviceWithOptions(request *AddDeviceRequest, runtime *util.RuntimeOptions) (_result *AddDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddDeviceResponse{}
	_body, _err := client.DoRequest(tea.String("AddDevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateCorpWithOptions(request *CreateCorpRequest, runtime *util.RuntimeOptions) (_result *CreateCorpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCorpResponse{}
	_body, _err := client.DoRequest(tea.String("CreateCorp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-05-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
