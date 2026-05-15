// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTaskDetailResponseBody
	GetCode() *string
	SetData(v *QueryTaskDetailResponseBodyData) *QueryTaskDetailResponseBody
	GetData() *QueryTaskDetailResponseBodyData
	SetHttpStatusCode(v string) *QueryTaskDetailResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *QueryTaskDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v string) *QueryTaskDetailResponseBody
	GetSuccess() *string
}

type QueryTaskDetailResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTaskDetailResponseBody) GetData() *QueryTaskDetailResponseBodyData {
	return s.Data
}

func (s *QueryTaskDetailResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *QueryTaskDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QueryTaskDetailResponseBody) SetCode(v string) *QueryTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetData(v *QueryTaskDetailResponseBodyData) *QueryTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskDetailResponseBody) SetHttpStatusCode(v string) *QueryTaskDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetMessage(v string) *QueryTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetRequestId(v string) *QueryTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskDetailResponseBody) SetSuccess(v string) *QueryTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskDetailResponseBodyData struct {
	CurrentPage  *string                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List         []*QueryTaskDetailResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize     *string                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalResults *string                                `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s QueryTaskDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponseBodyData) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *QueryTaskDetailResponseBodyData) GetList() []*QueryTaskDetailResponseBodyDataList {
	return s.List
}

func (s *QueryTaskDetailResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryTaskDetailResponseBodyData) GetTotalResults() *string {
	return s.TotalResults
}

func (s *QueryTaskDetailResponseBodyData) SetCurrentPage(v string) *QueryTaskDetailResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryTaskDetailResponseBodyData) SetList(v []*QueryTaskDetailResponseBodyDataList) *QueryTaskDetailResponseBodyData {
	s.List = v
	return s
}

func (s *QueryTaskDetailResponseBodyData) SetPageSize(v string) *QueryTaskDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailResponseBodyData) SetTotalResults(v string) *QueryTaskDetailResponseBodyData {
	s.TotalResults = &v
	return s
}

func (s *QueryTaskDetailResponseBodyData) Validate() error {
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

type QueryTaskDetailResponseBodyDataList struct {
	Ani            *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	BuId           *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	DepartmentId   *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Dnis           *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	EndReason      *int32  `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	ExtAttrs       *string `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	MemberId       *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName     *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	OutboundNum    *int32  `json:"OutboundNum,omitempty" xml:"OutboundNum,omitempty"`
	OutboundTaskId *int64  `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
	Priority       *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RetryTime      *string `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	ServicerId     *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ServicerName   *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	SkillGroup     *int32  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTaskDetailResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponseBodyDataList) GetAni() *string {
	return s.Ani
}

func (s *QueryTaskDetailResponseBodyDataList) GetBuId() *int64 {
	return s.BuId
}

func (s *QueryTaskDetailResponseBodyDataList) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *QueryTaskDetailResponseBodyDataList) GetDnis() *string {
	return s.Dnis
}

func (s *QueryTaskDetailResponseBodyDataList) GetEndReason() *int32 {
	return s.EndReason
}

func (s *QueryTaskDetailResponseBodyDataList) GetExtAttrs() *string {
	return s.ExtAttrs
}

func (s *QueryTaskDetailResponseBodyDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryTaskDetailResponseBodyDataList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryTaskDetailResponseBodyDataList) GetId() *int32 {
	return s.Id
}

func (s *QueryTaskDetailResponseBodyDataList) GetMemberId() *int64 {
	return s.MemberId
}

func (s *QueryTaskDetailResponseBodyDataList) GetMemberName() *string {
	return s.MemberName
}

func (s *QueryTaskDetailResponseBodyDataList) GetOutboundNum() *int32 {
	return s.OutboundNum
}

func (s *QueryTaskDetailResponseBodyDataList) GetOutboundTaskId() *int64 {
	return s.OutboundTaskId
}

func (s *QueryTaskDetailResponseBodyDataList) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryTaskDetailResponseBodyDataList) GetRetryTime() *string {
	return s.RetryTime
}

func (s *QueryTaskDetailResponseBodyDataList) GetServicerId() *int64 {
	return s.ServicerId
}

func (s *QueryTaskDetailResponseBodyDataList) GetServicerName() *string {
	return s.ServicerName
}

func (s *QueryTaskDetailResponseBodyDataList) GetSkillGroup() *int32 {
	return s.SkillGroup
}

func (s *QueryTaskDetailResponseBodyDataList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTaskDetailResponseBodyDataList) SetAni(v string) *QueryTaskDetailResponseBodyDataList {
	s.Ani = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetBuId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.BuId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetDepartmentId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.DepartmentId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetDnis(v string) *QueryTaskDetailResponseBodyDataList {
	s.Dnis = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetEndReason(v int32) *QueryTaskDetailResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetExtAttrs(v string) *QueryTaskDetailResponseBodyDataList {
	s.ExtAttrs = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetGmtCreate(v int64) *QueryTaskDetailResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetGmtModified(v int64) *QueryTaskDetailResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetId(v int32) *QueryTaskDetailResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetMemberId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.MemberId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetMemberName(v string) *QueryTaskDetailResponseBodyDataList {
	s.MemberName = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetOutboundNum(v int32) *QueryTaskDetailResponseBodyDataList {
	s.OutboundNum = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetOutboundTaskId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.OutboundTaskId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetPriority(v int32) *QueryTaskDetailResponseBodyDataList {
	s.Priority = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetRetryTime(v string) *QueryTaskDetailResponseBodyDataList {
	s.RetryTime = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetServicerId(v int64) *QueryTaskDetailResponseBodyDataList {
	s.ServicerId = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetServicerName(v string) *QueryTaskDetailResponseBodyDataList {
	s.ServicerName = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetSkillGroup(v int32) *QueryTaskDetailResponseBodyDataList {
	s.SkillGroup = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) SetStatus(v int32) *QueryTaskDetailResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryTaskDetailResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
