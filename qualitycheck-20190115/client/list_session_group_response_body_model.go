// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSessionGroupResponseBody
	GetCode() *string
	SetCount(v int32) *ListSessionGroupResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *ListSessionGroupResponseBody
	GetCurrentPage() *int32
	SetData(v *ListSessionGroupResponseBodyData) *ListSessionGroupResponseBody
	GetData() *ListSessionGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListSessionGroupResponseBody
	GetHttpStatusCode() *int32
	SetLastDataId(v string) *ListSessionGroupResponseBody
	GetLastDataId() *string
	SetMessage(v string) *ListSessionGroupResponseBody
	GetMessage() *string
	SetMessages(v *ListSessionGroupResponseBodyMessages) *ListSessionGroupResponseBody
	GetMessages() *ListSessionGroupResponseBodyMessages
	SetPageNumber(v int32) *ListSessionGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSessionGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSessionGroupResponseBody
	GetRequestId() *string
	SetResultCountId(v string) *ListSessionGroupResponseBody
	GetResultCountId() *string
	SetSuccess(v bool) *ListSessionGroupResponseBody
	GetSuccess() *bool
}

type ListSessionGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2228
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListSessionGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// xxx
	LastDataId *string `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	// example:
	//
	// successful
	Message  *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *ListSessionGroupResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSessionGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSessionGroupResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListSessionGroupResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSessionGroupResponseBody) GetData() *ListSessionGroupResponseBodyData {
	return s.Data
}

func (s *ListSessionGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSessionGroupResponseBody) GetLastDataId() *string {
	return s.LastDataId
}

func (s *ListSessionGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSessionGroupResponseBody) GetMessages() *ListSessionGroupResponseBodyMessages {
	return s.Messages
}

func (s *ListSessionGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSessionGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSessionGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionGroupResponseBody) GetResultCountId() *string {
	return s.ResultCountId
}

func (s *ListSessionGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSessionGroupResponseBody) SetCode(v string) *ListSessionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetCount(v int32) *ListSessionGroupResponseBody {
	s.Count = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetCurrentPage(v int32) *ListSessionGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetData(v *ListSessionGroupResponseBodyData) *ListSessionGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionGroupResponseBody) SetHttpStatusCode(v int32) *ListSessionGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetLastDataId(v string) *ListSessionGroupResponseBody {
	s.LastDataId = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetMessage(v string) *ListSessionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetMessages(v *ListSessionGroupResponseBodyMessages) *ListSessionGroupResponseBody {
	s.Messages = v
	return s
}

func (s *ListSessionGroupResponseBody) SetPageNumber(v int32) *ListSessionGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetPageSize(v int32) *ListSessionGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetRequestId(v string) *ListSessionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetResultCountId(v string) *ListSessionGroupResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListSessionGroupResponseBody) SetSuccess(v bool) *ListSessionGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListSessionGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyData struct {
	Data []*ListSessionGroupResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyData) GetData() []*ListSessionGroupResponseBodyDataData {
	return s.Data
}

func (s *ListSessionGroupResponseBodyData) SetData(v []*ListSessionGroupResponseBodyDataData) *ListSessionGroupResponseBodyData {
	s.Data = v
	return s
}

func (s *ListSessionGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataData struct {
	// example:
	//
	// 1
	AssignStatus *int32 `json:"AssignStatus,omitempty" xml:"AssignStatus,omitempty"`
	// example:
	//
	// 2022-09-26 10:09:14
	CallStartTime           *string                                                      `json:"CallStartTime,omitempty" xml:"CallStartTime,omitempty"`
	CallerList              *ListSessionGroupResponseBodyDataDataCallerList              `json:"CallerList,omitempty" xml:"CallerList,omitempty" type:"Struct"`
	CustomerIdList          *ListSessionGroupResponseBodyDataDataCustomerIdList          `json:"CustomerIdList,omitempty" xml:"CustomerIdList,omitempty" type:"Struct"`
	CustomerNameList        *ListSessionGroupResponseBodyDataDataCustomerNameList        `json:"CustomerNameList,omitempty" xml:"CustomerNameList,omitempty" type:"Struct"`
	CustomerServiceIdList   *ListSessionGroupResponseBodyDataDataCustomerServiceIdList   `json:"CustomerServiceIdList,omitempty" xml:"CustomerServiceIdList,omitempty" type:"Struct"`
	CustomerServiceNameList *ListSessionGroupResponseBodyDataDataCustomerServiceNameList `json:"CustomerServiceNameList,omitempty" xml:"CustomerServiceNameList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	HitSessionCount *int32 `json:"HitSessionCount,omitempty" xml:"HitSessionCount,omitempty"`
	// example:
	//
	// 4498420@a_z@93EAADF1-01D3-44BD-8AC9-F57F447EFCE8_1614*****
	LastDataId *string `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	// example:
	//
	// 1
	ReviewStatus *int32                                            `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	ReviewerList *ListSessionGroupResponseBodyDataDataReviewerList `json:"ReviewerList,omitempty" xml:"ReviewerList,omitempty" type:"Struct"`
	// example:
	//
	// 123
	SchemeTaskConfigId   *int64  `json:"SchemeTaskConfigId,omitempty" xml:"SchemeTaskConfigId,omitempty"`
	SchemeTaskConfigName *string `json:"SchemeTaskConfigName,omitempty" xml:"SchemeTaskConfigName,omitempty"`
	// example:
	//
	// 100
	Score *int64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 1
	SessionCount *int32 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// SessionGroupA
	SessionGroupId *string `json:"SessionGroupId,omitempty" xml:"SessionGroupId,omitempty"`
	// example:
	//
	// true
	SessionGroupReviewedOrComplained *bool                                                   `json:"SessionGroupReviewedOrComplained,omitempty" xml:"SessionGroupReviewedOrComplained,omitempty"`
	SkillGroupNameList               *ListSessionGroupResponseBodyDataDataSkillGroupNameList `json:"SkillGroupNameList,omitempty" xml:"SkillGroupNameList,omitempty" type:"Struct"`
}

func (s ListSessionGroupResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataData) GetAssignStatus() *int32 {
	return s.AssignStatus
}

func (s *ListSessionGroupResponseBodyDataData) GetCallStartTime() *string {
	return s.CallStartTime
}

func (s *ListSessionGroupResponseBodyDataData) GetCallerList() *ListSessionGroupResponseBodyDataDataCallerList {
	return s.CallerList
}

func (s *ListSessionGroupResponseBodyDataData) GetCustomerIdList() *ListSessionGroupResponseBodyDataDataCustomerIdList {
	return s.CustomerIdList
}

func (s *ListSessionGroupResponseBodyDataData) GetCustomerNameList() *ListSessionGroupResponseBodyDataDataCustomerNameList {
	return s.CustomerNameList
}

func (s *ListSessionGroupResponseBodyDataData) GetCustomerServiceIdList() *ListSessionGroupResponseBodyDataDataCustomerServiceIdList {
	return s.CustomerServiceIdList
}

func (s *ListSessionGroupResponseBodyDataData) GetCustomerServiceNameList() *ListSessionGroupResponseBodyDataDataCustomerServiceNameList {
	return s.CustomerServiceNameList
}

func (s *ListSessionGroupResponseBodyDataData) GetHitSessionCount() *int32 {
	return s.HitSessionCount
}

func (s *ListSessionGroupResponseBodyDataData) GetLastDataId() *string {
	return s.LastDataId
}

func (s *ListSessionGroupResponseBodyDataData) GetReviewStatus() *int32 {
	return s.ReviewStatus
}

func (s *ListSessionGroupResponseBodyDataData) GetReviewerList() *ListSessionGroupResponseBodyDataDataReviewerList {
	return s.ReviewerList
}

func (s *ListSessionGroupResponseBodyDataData) GetSchemeTaskConfigId() *int64 {
	return s.SchemeTaskConfigId
}

func (s *ListSessionGroupResponseBodyDataData) GetSchemeTaskConfigName() *string {
	return s.SchemeTaskConfigName
}

func (s *ListSessionGroupResponseBodyDataData) GetScore() *int64 {
	return s.Score
}

func (s *ListSessionGroupResponseBodyDataData) GetSessionCount() *int32 {
	return s.SessionCount
}

func (s *ListSessionGroupResponseBodyDataData) GetSessionGroupId() *string {
	return s.SessionGroupId
}

func (s *ListSessionGroupResponseBodyDataData) GetSessionGroupReviewedOrComplained() *bool {
	return s.SessionGroupReviewedOrComplained
}

func (s *ListSessionGroupResponseBodyDataData) GetSkillGroupNameList() *ListSessionGroupResponseBodyDataDataSkillGroupNameList {
	return s.SkillGroupNameList
}

func (s *ListSessionGroupResponseBodyDataData) SetAssignStatus(v int32) *ListSessionGroupResponseBodyDataData {
	s.AssignStatus = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCallStartTime(v string) *ListSessionGroupResponseBodyDataData {
	s.CallStartTime = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCallerList(v *ListSessionGroupResponseBodyDataDataCallerList) *ListSessionGroupResponseBodyDataData {
	s.CallerList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerIdList(v *ListSessionGroupResponseBodyDataDataCustomerIdList) *ListSessionGroupResponseBodyDataData {
	s.CustomerIdList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerNameList(v *ListSessionGroupResponseBodyDataDataCustomerNameList) *ListSessionGroupResponseBodyDataData {
	s.CustomerNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerServiceIdList(v *ListSessionGroupResponseBodyDataDataCustomerServiceIdList) *ListSessionGroupResponseBodyDataData {
	s.CustomerServiceIdList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetCustomerServiceNameList(v *ListSessionGroupResponseBodyDataDataCustomerServiceNameList) *ListSessionGroupResponseBodyDataData {
	s.CustomerServiceNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetHitSessionCount(v int32) *ListSessionGroupResponseBodyDataData {
	s.HitSessionCount = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetLastDataId(v string) *ListSessionGroupResponseBodyDataData {
	s.LastDataId = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetReviewStatus(v int32) *ListSessionGroupResponseBodyDataData {
	s.ReviewStatus = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetReviewerList(v *ListSessionGroupResponseBodyDataDataReviewerList) *ListSessionGroupResponseBodyDataData {
	s.ReviewerList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSchemeTaskConfigId(v int64) *ListSessionGroupResponseBodyDataData {
	s.SchemeTaskConfigId = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSchemeTaskConfigName(v string) *ListSessionGroupResponseBodyDataData {
	s.SchemeTaskConfigName = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetScore(v int64) *ListSessionGroupResponseBodyDataData {
	s.Score = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSessionCount(v int32) *ListSessionGroupResponseBodyDataData {
	s.SessionCount = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSessionGroupId(v string) *ListSessionGroupResponseBodyDataData {
	s.SessionGroupId = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSessionGroupReviewedOrComplained(v bool) *ListSessionGroupResponseBodyDataData {
	s.SessionGroupReviewedOrComplained = &v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) SetSkillGroupNameList(v *ListSessionGroupResponseBodyDataDataSkillGroupNameList) *ListSessionGroupResponseBodyDataData {
	s.SkillGroupNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataCallerList struct {
	CallerList []*string `json:"CallerList,omitempty" xml:"CallerList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCallerList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCallerList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCallerList) GetCallerList() []*string {
	return s.CallerList
}

func (s *ListSessionGroupResponseBodyDataDataCallerList) SetCallerList(v []*string) *ListSessionGroupResponseBodyDataDataCallerList {
	s.CallerList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataCallerList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataCustomerIdList struct {
	CustomerIdList []*string `json:"CustomerIdList,omitempty" xml:"CustomerIdList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerIdList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerIdList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerIdList) GetCustomerIdList() []*string {
	return s.CustomerIdList
}

func (s *ListSessionGroupResponseBodyDataDataCustomerIdList) SetCustomerIdList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerIdList {
	s.CustomerIdList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataCustomerIdList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataCustomerNameList struct {
	CustomerNameList []*string `json:"CustomerNameList,omitempty" xml:"CustomerNameList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerNameList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerNameList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerNameList) GetCustomerNameList() []*string {
	return s.CustomerNameList
}

func (s *ListSessionGroupResponseBodyDataDataCustomerNameList) SetCustomerNameList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerNameList {
	s.CustomerNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataCustomerNameList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataCustomerServiceIdList struct {
	CustomerServiceIdList []*string `json:"CustomerServiceIdList,omitempty" xml:"CustomerServiceIdList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceIdList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceIdList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceIdList) GetCustomerServiceIdList() []*string {
	return s.CustomerServiceIdList
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceIdList) SetCustomerServiceIdList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerServiceIdList {
	s.CustomerServiceIdList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceIdList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataCustomerServiceNameList struct {
	CustomerServiceNameList []*string `json:"CustomerServiceNameList,omitempty" xml:"CustomerServiceNameList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceNameList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataCustomerServiceNameList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceNameList) GetCustomerServiceNameList() []*string {
	return s.CustomerServiceNameList
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceNameList) SetCustomerServiceNameList(v []*string) *ListSessionGroupResponseBodyDataDataCustomerServiceNameList {
	s.CustomerServiceNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataCustomerServiceNameList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataReviewerList struct {
	ReviewerList []*string `json:"ReviewerList,omitempty" xml:"ReviewerList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataReviewerList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataReviewerList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataReviewerList) GetReviewerList() []*string {
	return s.ReviewerList
}

func (s *ListSessionGroupResponseBodyDataDataReviewerList) SetReviewerList(v []*string) *ListSessionGroupResponseBodyDataDataReviewerList {
	s.ReviewerList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataReviewerList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyDataDataSkillGroupNameList struct {
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" xml:"SkillGroupNameList,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyDataDataSkillGroupNameList) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyDataDataSkillGroupNameList) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyDataDataSkillGroupNameList) GetSkillGroupNameList() []*string {
	return s.SkillGroupNameList
}

func (s *ListSessionGroupResponseBodyDataDataSkillGroupNameList) SetSkillGroupNameList(v []*string) *ListSessionGroupResponseBodyDataDataSkillGroupNameList {
	s.SkillGroupNameList = v
	return s
}

func (s *ListSessionGroupResponseBodyDataDataSkillGroupNameList) Validate() error {
	return dara.Validate(s)
}

type ListSessionGroupResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s ListSessionGroupResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *ListSessionGroupResponseBodyMessages) SetMessage(v []*string) *ListSessionGroupResponseBodyMessages {
	s.Message = v
	return s
}

func (s *ListSessionGroupResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
