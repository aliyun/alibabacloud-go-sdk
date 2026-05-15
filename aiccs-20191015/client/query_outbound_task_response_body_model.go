// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryOutboundTaskResponseBody
	GetCode() *string
	SetData(v *QueryOutboundTaskResponseBodyData) *QueryOutboundTaskResponseBody
	GetData() *QueryOutboundTaskResponseBodyData
	SetHttpStatusCode(v string) *QueryOutboundTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *QueryOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *QueryOutboundTaskResponseBody
	GetSuccess() *string
}

type QueryOutboundTaskResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryOutboundTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryOutboundTaskResponseBody) GetData() *QueryOutboundTaskResponseBodyData {
	return s.Data
}

func (s *QueryOutboundTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *QueryOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOutboundTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QueryOutboundTaskResponseBody) SetCode(v string) *QueryOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetData(v *QueryOutboundTaskResponseBodyData) *QueryOutboundTaskResponseBody {
	s.Data = v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetHttpStatusCode(v string) *QueryOutboundTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetMessage(v string) *QueryOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetRequestId(v string) *QueryOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) SetSuccess(v string) *QueryOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOutboundTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOutboundTaskResponseBodyData struct {
	CurrentPage  *string                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List         []*QueryOutboundTaskResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize     *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalResults *string                                  `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s QueryOutboundTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOutboundTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponseBodyData) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *QueryOutboundTaskResponseBodyData) GetList() []*QueryOutboundTaskResponseBodyDataList {
	return s.List
}

func (s *QueryOutboundTaskResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryOutboundTaskResponseBodyData) GetTotalResults() *string {
	return s.TotalResults
}

func (s *QueryOutboundTaskResponseBodyData) SetCurrentPage(v string) *QueryOutboundTaskResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) SetList(v []*QueryOutboundTaskResponseBodyDataList) *QueryOutboundTaskResponseBodyData {
	s.List = v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) SetPageSize(v string) *QueryOutboundTaskResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) SetTotalResults(v string) *QueryOutboundTaskResponseBodyData {
	s.TotalResults = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryOutboundTaskResponseBodyDataList struct {
	BuId          *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	CallerNum     *string `json:"CallerNum,omitempty" xml:"CallerNum,omitempty"`
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DepartmentId  *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtAttrs      *string `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Model         *int32  `json:"Model,omitempty" xml:"Model,omitempty"`
	Modifier      *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RetryInterval *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryTime     *int32  `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	SkillGroup    *int64  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type          *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryOutboundTaskResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryOutboundTaskResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponseBodyDataList) GetBuId() *int64 {
	return s.BuId
}

func (s *QueryOutboundTaskResponseBodyDataList) GetCallerNum() *string {
	return s.CallerNum
}

func (s *QueryOutboundTaskResponseBodyDataList) GetCreator() *string {
	return s.Creator
}

func (s *QueryOutboundTaskResponseBodyDataList) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *QueryOutboundTaskResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *QueryOutboundTaskResponseBodyDataList) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryOutboundTaskResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryOutboundTaskResponseBodyDataList) GetExtAttrs() *string {
	return s.ExtAttrs
}

func (s *QueryOutboundTaskResponseBodyDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryOutboundTaskResponseBodyDataList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryOutboundTaskResponseBodyDataList) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryOutboundTaskResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *QueryOutboundTaskResponseBodyDataList) GetModel() *int32 {
	return s.Model
}

func (s *QueryOutboundTaskResponseBodyDataList) GetModifier() *string {
	return s.Modifier
}

func (s *QueryOutboundTaskResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *QueryOutboundTaskResponseBodyDataList) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *QueryOutboundTaskResponseBodyDataList) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *QueryOutboundTaskResponseBodyDataList) GetSkillGroup() *int64 {
	return s.SkillGroup
}

func (s *QueryOutboundTaskResponseBodyDataList) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryOutboundTaskResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryOutboundTaskResponseBodyDataList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryOutboundTaskResponseBodyDataList) GetType() *int32 {
	return s.Type
}

func (s *QueryOutboundTaskResponseBodyDataList) SetBuId(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.BuId = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetCallerNum(v string) *QueryOutboundTaskResponseBodyDataList {
	s.CallerNum = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetCreator(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetDepartmentId(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.DepartmentId = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetDescription(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetEndDate(v string) *QueryOutboundTaskResponseBodyDataList {
	s.EndDate = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetEndTime(v string) *QueryOutboundTaskResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetExtAttrs(v string) *QueryOutboundTaskResponseBodyDataList {
	s.ExtAttrs = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetGmtCreate(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetGmtModified(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetGroupName(v string) *QueryOutboundTaskResponseBodyDataList {
	s.GroupName = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetId(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetModel(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.Model = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetModifier(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Modifier = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetName(v string) *QueryOutboundTaskResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetRetryInterval(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.RetryInterval = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetRetryTime(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.RetryTime = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetSkillGroup(v int64) *QueryOutboundTaskResponseBodyDataList {
	s.SkillGroup = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetStartDate(v string) *QueryOutboundTaskResponseBodyDataList {
	s.StartDate = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetStartTime(v string) *QueryOutboundTaskResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetStatus(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) SetType(v int32) *QueryOutboundTaskResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *QueryOutboundTaskResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
