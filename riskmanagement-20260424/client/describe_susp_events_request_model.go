// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeSuspEventsRequest
	GetRegionId() *string
	SetSdkRequest(v *DescribeSuspEventsRequestSdkRequest) *DescribeSuspEventsRequest
	GetSdkRequest() *DescribeSuspEventsRequestSdkRequest
}

type DescribeSuspEventsRequest struct {
	RegionId   *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequest *DescribeSuspEventsRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s DescribeSuspEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSuspEventsRequest) GetSdkRequest() *DescribeSuspEventsRequestSdkRequest {
	return s.SdkRequest
}

func (s *DescribeSuspEventsRequest) SetRegionId(v string) *DescribeSuspEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSdkRequest(v *DescribeSuspEventsRequestSdkRequest) *DescribeSuspEventsRequest {
	s.SdkRequest = v
	return s
}

func (s *DescribeSuspEventsRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSuspEventsRequestSdkRequest struct {
	AlarmUniqueInfo            *string   `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	AssetsTypeList             []*string `json:"AssetsTypeList,omitempty" xml:"AssetsTypeList,omitempty" type:"Repeated"`
	ClusterId                  *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerFieldName         *string   `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	ContainerFieldValue        *string   `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	CurrentPage                *string   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Dealed                     *string   `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	DetectSource               *string   `json:"DetectSource,omitempty" xml:"DetectSource,omitempty"`
	EventNames                 *string   `json:"EventNames,omitempty" xml:"EventNames,omitempty"`
	From                       *string   `json:"From,omitempty" xml:"From,omitempty"`
	GroupId                    *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id                         *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang                       *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Levels                     *string   `json:"Levels,omitempty" xml:"Levels,omitempty"`
	MultiAccountActionType     *int32    `json:"MultiAccountActionType,omitempty" xml:"MultiAccountActionType,omitempty"`
	Name                       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateErrorCodeList       []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
	OperateTimeEnd             *string   `json:"OperateTimeEnd,omitempty" xml:"OperateTimeEnd,omitempty"`
	OperateTimeStart           *string   `json:"OperateTimeStart,omitempty" xml:"OperateTimeStart,omitempty"`
	PageSize                   *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentEventTypes           *string   `json:"ParentEventTypes,omitempty" xml:"ParentEventTypes,omitempty"`
	Remark                     *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceDirectoryAccountId *int64    `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	SortColumn                 *string   `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortType                   *string   `json:"SortType,omitempty" xml:"SortType,omitempty"`
	Source                     *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceAliUids              []*int64  `json:"SourceAliUids,omitempty" xml:"SourceAliUids,omitempty" type:"Repeated"`
	SourceIp                   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StrictMode                 *string   `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	SupportOperateCodeList     []*string `json:"SupportOperateCodeList,omitempty" xml:"SupportOperateCodeList,omitempty" type:"Repeated"`
	TacticId                   *string   `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	TargetType                 *string   `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TimeEnd                    *string   `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	TimeStart                  *string   `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
	UniqueInfo                 *string   `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	Uuids                      *string   `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeSuspEventsRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequestSdkRequest) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeSuspEventsRequestSdkRequest) GetAssetsTypeList() []*string {
	return s.AssetsTypeList
}

func (s *DescribeSuspEventsRequestSdkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSuspEventsRequestSdkRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeSuspEventsRequestSdkRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeSuspEventsRequestSdkRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSuspEventsRequestSdkRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeSuspEventsRequestSdkRequest) GetDetectSource() *string {
	return s.DetectSource
}

func (s *DescribeSuspEventsRequestSdkRequest) GetEventNames() *string {
	return s.EventNames
}

func (s *DescribeSuspEventsRequestSdkRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeSuspEventsRequestSdkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeSuspEventsRequestSdkRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeSuspEventsRequestSdkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSuspEventsRequestSdkRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeSuspEventsRequestSdkRequest) GetMultiAccountActionType() *int32 {
	return s.MultiAccountActionType
}

func (s *DescribeSuspEventsRequestSdkRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSuspEventsRequestSdkRequest) GetOperateErrorCodeList() []*string {
	return s.OperateErrorCodeList
}

func (s *DescribeSuspEventsRequestSdkRequest) GetOperateTimeEnd() *string {
	return s.OperateTimeEnd
}

func (s *DescribeSuspEventsRequestSdkRequest) GetOperateTimeStart() *string {
	return s.OperateTimeStart
}

func (s *DescribeSuspEventsRequestSdkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSuspEventsRequestSdkRequest) GetParentEventTypes() *string {
	return s.ParentEventTypes
}

func (s *DescribeSuspEventsRequestSdkRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeSuspEventsRequestSdkRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeSuspEventsRequestSdkRequest) GetSortColumn() *string {
	return s.SortColumn
}

func (s *DescribeSuspEventsRequestSdkRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeSuspEventsRequestSdkRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeSuspEventsRequestSdkRequest) GetSourceAliUids() []*int64 {
	return s.SourceAliUids
}

func (s *DescribeSuspEventsRequestSdkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventsRequestSdkRequest) GetStrictMode() *string {
	return s.StrictMode
}

func (s *DescribeSuspEventsRequestSdkRequest) GetSupportOperateCodeList() []*string {
	return s.SupportOperateCodeList
}

func (s *DescribeSuspEventsRequestSdkRequest) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeSuspEventsRequestSdkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeSuspEventsRequestSdkRequest) GetTimeEnd() *string {
	return s.TimeEnd
}

func (s *DescribeSuspEventsRequestSdkRequest) GetTimeStart() *string {
	return s.TimeStart
}

func (s *DescribeSuspEventsRequestSdkRequest) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *DescribeSuspEventsRequestSdkRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeSuspEventsRequestSdkRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequestSdkRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetAssetsTypeList(v []*string) *DescribeSuspEventsRequestSdkRequest {
	s.AssetsTypeList = v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetClusterId(v string) *DescribeSuspEventsRequestSdkRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetContainerFieldName(v string) *DescribeSuspEventsRequestSdkRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetContainerFieldValue(v string) *DescribeSuspEventsRequestSdkRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetCurrentPage(v string) *DescribeSuspEventsRequestSdkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetDealed(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetDetectSource(v string) *DescribeSuspEventsRequestSdkRequest {
	s.DetectSource = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetEventNames(v string) *DescribeSuspEventsRequestSdkRequest {
	s.EventNames = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetFrom(v string) *DescribeSuspEventsRequestSdkRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetGroupId(v int64) *DescribeSuspEventsRequestSdkRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetId(v int64) *DescribeSuspEventsRequestSdkRequest {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetLang(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetLevels(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetMultiAccountActionType(v int32) *DescribeSuspEventsRequestSdkRequest {
	s.MultiAccountActionType = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetName(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetOperateErrorCodeList(v []*string) *DescribeSuspEventsRequestSdkRequest {
	s.OperateErrorCodeList = v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetOperateTimeEnd(v string) *DescribeSuspEventsRequestSdkRequest {
	s.OperateTimeEnd = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetOperateTimeStart(v string) *DescribeSuspEventsRequestSdkRequest {
	s.OperateTimeStart = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetPageSize(v string) *DescribeSuspEventsRequestSdkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetParentEventTypes(v string) *DescribeSuspEventsRequestSdkRequest {
	s.ParentEventTypes = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetRemark(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Remark = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetResourceDirectoryAccountId(v int64) *DescribeSuspEventsRequestSdkRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetSortColumn(v string) *DescribeSuspEventsRequestSdkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetSortType(v string) *DescribeSuspEventsRequestSdkRequest {
	s.SortType = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetSource(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Source = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetSourceAliUids(v []*int64) *DescribeSuspEventsRequestSdkRequest {
	s.SourceAliUids = v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetSourceIp(v string) *DescribeSuspEventsRequestSdkRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetStrictMode(v string) *DescribeSuspEventsRequestSdkRequest {
	s.StrictMode = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetSupportOperateCodeList(v []*string) *DescribeSuspEventsRequestSdkRequest {
	s.SupportOperateCodeList = v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetTacticId(v string) *DescribeSuspEventsRequestSdkRequest {
	s.TacticId = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetTargetType(v string) *DescribeSuspEventsRequestSdkRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetTimeEnd(v string) *DescribeSuspEventsRequestSdkRequest {
	s.TimeEnd = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetTimeStart(v string) *DescribeSuspEventsRequestSdkRequest {
	s.TimeStart = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetUniqueInfo(v string) *DescribeSuspEventsRequestSdkRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) SetUuids(v string) *DescribeSuspEventsRequestSdkRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeSuspEventsRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
