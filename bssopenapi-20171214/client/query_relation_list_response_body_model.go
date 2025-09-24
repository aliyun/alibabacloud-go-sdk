// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRelationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRelationListResponseBody
	GetCode() *string
	SetData(v *QueryRelationListResponseBodyData) *QueryRelationListResponseBody
	GetData() *QueryRelationListResponseBodyData
	SetMessage(v string) *QueryRelationListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRelationListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRelationListResponseBody
	GetSuccess() *bool
}

type QueryRelationListResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryRelationListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7FC5D662-37FD-40A6-85B1-33442D815184
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRelationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRelationListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRelationListResponseBody) GetData() *QueryRelationListResponseBodyData {
	return s.Data
}

func (s *QueryRelationListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRelationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRelationListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRelationListResponseBody) SetCode(v string) *QueryRelationListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRelationListResponseBody) SetData(v *QueryRelationListResponseBodyData) *QueryRelationListResponseBody {
	s.Data = v
	return s
}

func (s *QueryRelationListResponseBody) SetMessage(v string) *QueryRelationListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRelationListResponseBody) SetRequestId(v string) *QueryRelationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRelationListResponseBody) SetSuccess(v bool) *QueryRelationListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRelationListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryRelationListResponseBodyData struct {
	// The relationships.
	FinancialRelationInfoList []*QueryRelationListResponseBodyDataFinancialRelationInfoList `json:"FinancialRelationInfoList,omitempty" xml:"FinancialRelationInfoList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRelationListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRelationListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponseBodyData) GetFinancialRelationInfoList() []*QueryRelationListResponseBodyDataFinancialRelationInfoList {
	return s.FinancialRelationInfoList
}

func (s *QueryRelationListResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRelationListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRelationListResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryRelationListResponseBodyData) SetFinancialRelationInfoList(v []*QueryRelationListResponseBodyDataFinancialRelationInfoList) *QueryRelationListResponseBodyData {
	s.FinancialRelationInfoList = v
	return s
}

func (s *QueryRelationListResponseBodyData) SetPageNum(v int32) *QueryRelationListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryRelationListResponseBodyData) SetPageSize(v int32) *QueryRelationListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRelationListResponseBodyData) SetTotalCount(v int32) *QueryRelationListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryRelationListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryRelationListResponseBodyDataFinancialRelationInfoList struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1851253838840762
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// caiwuyun_test4
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// test4
	AccountNickName *string `json:"AccountNickName,omitempty" xml:"AccountNickName,omitempty"`
	// The type of the account. Valid values: MASTER and MEMBER.
	//
	// example:
	//
	// MEMBER
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The time when the relationship became invalid. If no value is returned, the relationship is still valid.
	//
	// example:
	//
	// 2021-03-08T15:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the relationship.
	//
	// example:
	//
	// 51463
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The type of the relationship. Valid values: FinancialManagement and FinancialTrusteeship.
	//
	// example:
	//
	// FinancialManagement
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The time when the relationship was established. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC. Example: 2016-05-23T12:00:00Z.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	SetupTime *string `json:"SetupTime,omitempty" xml:"SetupTime,omitempty"`
	// The time when the relationship became valid. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC. Example: 2016-05-23T12:00:00Z.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the relationship. One of the enumeration members of the RelationshipStatusEnum data type is returned.
	//
	// example:
	//
	// RELATED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryRelationListResponseBodyDataFinancialRelationInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryRelationListResponseBodyDataFinancialRelationInfoList) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetAccountNickName() *string {
	return s.AccountNickName
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetAccountType() *string {
	return s.AccountType
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetRelationId() *int64 {
	return s.RelationId
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetRelationType() *string {
	return s.RelationType
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetSetupTime() *string {
	return s.SetupTime
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) GetState() *string {
	return s.State
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountId(v int64) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountId = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountName(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountName = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountNickName(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountNickName = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountType(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountType = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetEndTime(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.EndTime = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetRelationId(v int64) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.RelationId = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetRelationType(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.RelationType = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetSetupTime(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.SetupTime = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetStartTime(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.StartTime = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetState(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.State = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) Validate() error {
	return dara.Validate(s)
}
