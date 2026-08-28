// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOrganizationMembersResponseBody
	GetCode() *string
	SetData(v []*ListOrganizationMembersResponseBodyData) *ListOrganizationMembersResponseBody
	GetData() []*ListOrganizationMembersResponseBodyData
	SetMessage(v string) *ListOrganizationMembersResponseBody
	GetMessage() *string
	SetPageNo(v int32) *ListOrganizationMembersResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListOrganizationMembersResponseBody
	GetPageSize() *int32
	SetSuccess(v bool) *ListOrganizationMembersResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListOrganizationMembersResponseBody
	GetTotal() *int32
}

type ListOrganizationMembersResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data []*ListOrganizationMembersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 18
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOrganizationMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOrganizationMembersResponseBody) GetData() []*ListOrganizationMembersResponseBodyData {
	return s.Data
}

func (s *ListOrganizationMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOrganizationMembersResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListOrganizationMembersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrganizationMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrganizationMembersResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListOrganizationMembersResponseBody) SetCode(v string) *ListOrganizationMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetData(v []*ListOrganizationMembersResponseBodyData) *ListOrganizationMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetMessage(v string) *ListOrganizationMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetPageNo(v int32) *ListOrganizationMembersResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetPageSize(v int32) *ListOrganizationMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetSuccess(v bool) *ListOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetTotal(v int32) *ListOrganizationMembersResponseBody {
	s.Total = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationMembersResponseBodyData struct {
	// The member business ID.
	//
	// example:
	//
	// 112233
	AccountBizId *string `json:"AccountBizId,omitempty" xml:"AccountBizId,omitempty"`
	// The ID of the member accounts.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member accounts.
	//
	// example:
	//
	// test_001
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// API Key ID
	//
	// example:
	//
	// key_123456789
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// The member email address.
	//
	// example:
	//
	// test@email.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The time when the member joined.
	//
	// example:
	//
	// 2026-06-10T11:57:42.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The masked API key.
	//
	// example:
	//
	// prefix.abc****456
	MaskedApiKey *string `json:"MaskedApiKey,omitempty" xml:"MaskedApiKey,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_123456789
	OrgId         *string                                               `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PackLimitInfo *ListOrganizationMembersResponseBodyDataPackLimitInfo `json:"PackLimitInfo,omitempty" xml:"PackLimitInfo,omitempty" type:"Struct"`
	// The list of member roles.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The seat resource allocate ID.
	//
	// example:
	//
	// seat_123456
	SeatId *string `json:"SeatId,omitempty" xml:"SeatId,omitempty"`
	// The seat specification type. Valid values:
	//
	// - standard: Standard seat.
	//
	// - pro: Pro seat.
	//
	// - max: Max seat.
	//
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The member status.
	//
	// example:
	//
	// ACTIVE
	Status           *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscriptionInfo *ListOrganizationMembersResponseBodyDataSubscriptionInfo `json:"SubscriptionInfo,omitempty" xml:"SubscriptionInfo,omitempty" type:"Struct"`
}

func (s ListOrganizationMembersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyData) GetAccountBizId() *string {
	return s.AccountBizId
}

func (s *ListOrganizationMembersResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListOrganizationMembersResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListOrganizationMembersResponseBodyData) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ListOrganizationMembersResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ListOrganizationMembersResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListOrganizationMembersResponseBodyData) GetMaskedApiKey() *string {
	return s.MaskedApiKey
}

func (s *ListOrganizationMembersResponseBodyData) GetOrgId() *string {
	return s.OrgId
}

func (s *ListOrganizationMembersResponseBodyData) GetPackLimitInfo() *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	return s.PackLimitInfo
}

func (s *ListOrganizationMembersResponseBodyData) GetRoles() []*string {
	return s.Roles
}

func (s *ListOrganizationMembersResponseBodyData) GetSeatId() *string {
	return s.SeatId
}

func (s *ListOrganizationMembersResponseBodyData) GetSpecType() *string {
	return s.SpecType
}

func (s *ListOrganizationMembersResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListOrganizationMembersResponseBodyData) GetSubscriptionInfo() *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	return s.SubscriptionInfo
}

func (s *ListOrganizationMembersResponseBodyData) SetAccountBizId(v string) *ListOrganizationMembersResponseBodyData {
	s.AccountBizId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetAccountId(v string) *ListOrganizationMembersResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetAccountName(v string) *ListOrganizationMembersResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetApiKeyId(v string) *ListOrganizationMembersResponseBodyData {
	s.ApiKeyId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetEmail(v string) *ListOrganizationMembersResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetGmtCreate(v string) *ListOrganizationMembersResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetMaskedApiKey(v string) *ListOrganizationMembersResponseBodyData {
	s.MaskedApiKey = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetOrgId(v string) *ListOrganizationMembersResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetPackLimitInfo(v *ListOrganizationMembersResponseBodyDataPackLimitInfo) *ListOrganizationMembersResponseBodyData {
	s.PackLimitInfo = v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetRoles(v []*string) *ListOrganizationMembersResponseBodyData {
	s.Roles = v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetSeatId(v string) *ListOrganizationMembersResponseBodyData {
	s.SeatId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetSpecType(v string) *ListOrganizationMembersResponseBodyData {
	s.SpecType = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetStatus(v string) *ListOrganizationMembersResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetSubscriptionInfo(v *ListOrganizationMembersResponseBodyDataSubscriptionInfo) *ListOrganizationMembersResponseBodyData {
	s.SubscriptionInfo = v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) Validate() error {
	if s.PackLimitInfo != nil {
		if err := s.PackLimitInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionInfo != nil {
		if err := s.SubscriptionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOrganizationMembersResponseBodyDataPackLimitInfo struct {
	AvailableLimit    *float64 `json:"AvailableLimit,omitempty" xml:"AvailableLimit,omitempty"`
	CycleEndTime      *int64   `json:"CycleEndTime,omitempty" xml:"CycleEndTime,omitempty"`
	CycleStartTime    *int64   `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	FrozenCredits     *float64 `json:"FrozenCredits,omitempty" xml:"FrozenCredits,omitempty"`
	HasShareLimit     *bool    `json:"HasShareLimit,omitempty" xml:"HasShareLimit,omitempty"`
	IsAvailable       *bool    `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	LastConfirmedTime *int64   `json:"LastConfirmedTime,omitempty" xml:"LastConfirmedTime,omitempty"`
	UpperLimit        *float64 `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
	UsedCredits       *float64 `json:"UsedCredits,omitempty" xml:"UsedCredits,omitempty"`
}

func (s ListOrganizationMembersResponseBodyDataPackLimitInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyDataPackLimitInfo) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetAvailableLimit() *float64 {
	return s.AvailableLimit
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetCycleEndTime() *int64 {
	return s.CycleEndTime
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetCycleStartTime() *int64 {
	return s.CycleStartTime
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetFrozenCredits() *float64 {
	return s.FrozenCredits
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetHasShareLimit() *bool {
	return s.HasShareLimit
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetLastConfirmedTime() *int64 {
	return s.LastConfirmedTime
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetUpperLimit() *float64 {
	return s.UpperLimit
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) GetUsedCredits() *float64 {
	return s.UsedCredits
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetAvailableLimit(v float64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.AvailableLimit = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetCycleEndTime(v int64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.CycleEndTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetCycleStartTime(v int64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.CycleStartTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetFrozenCredits(v float64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.FrozenCredits = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetHasShareLimit(v bool) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.HasShareLimit = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetIsAvailable(v bool) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.IsAvailable = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetLastConfirmedTime(v int64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.LastConfirmedTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetUpperLimit(v float64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.UpperLimit = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) SetUsedCredits(v float64) *ListOrganizationMembersResponseBodyDataPackLimitInfo {
	s.UsedCredits = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataPackLimitInfo) Validate() error {
	return dara.Validate(s)
}

type ListOrganizationMembersResponseBodyDataSubscriptionInfo struct {
	EndTime      *int64                                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EquityList   []*ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList `json:"EquityList,omitempty" xml:"EquityList,omitempty" type:"Repeated"`
	InstanceCode *string                                                              `json:"InstanceCode,omitempty" xml:"InstanceCode,omitempty"`
	PayMode      *string                                                              `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	ProductCode  *string                                                              `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SpecType     *string                                                              `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	StartTime    *int64                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOrganizationMembersResponseBodyDataSubscriptionInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyDataSubscriptionInfo) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetEquityList() []*ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	return s.EquityList
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetInstanceCode() *string {
	return s.InstanceCode
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetPayMode() *string {
	return s.PayMode
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetSpecType() *string {
	return s.SpecType
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) GetStatus() *string {
	return s.Status
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetEndTime(v int64) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.EndTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetEquityList(v []*ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.EquityList = v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetInstanceCode(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.InstanceCode = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetPayMode(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.PayMode = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetProductCode(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.ProductCode = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetSpecType(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.SpecType = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetStartTime(v int64) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.StartTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) SetStatus(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfo {
	s.Status = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfo) Validate() error {
	if s.EquityList != nil {
		for _, item := range s.EquityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList struct {
	CycleEndTime      *int64   `json:"CycleEndTime,omitempty" xml:"CycleEndTime,omitempty"`
	CycleStartTime    *int64   `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	CycleSurplusValue *float64 `json:"CycleSurplusValue,omitempty" xml:"CycleSurplusValue,omitempty"`
	CycleTotalValue   *float64 `json:"CycleTotalValue,omitempty" xml:"CycleTotalValue,omitempty"`
	EquityType        *string  `json:"EquityType,omitempty" xml:"EquityType,omitempty"`
	EquityUnit        *string  `json:"EquityUnit,omitempty" xml:"EquityUnit,omitempty"`
}

func (s ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GetCycleEndTime() *int64 {
	return s.CycleEndTime
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GetCycleStartTime() *int64 {
	return s.CycleStartTime
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GetCycleSurplusValue() *float64 {
	return s.CycleSurplusValue
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GetCycleTotalValue() *float64 {
	return s.CycleTotalValue
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GetEquityType() *string {
	return s.EquityType
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) GetEquityUnit() *string {
	return s.EquityUnit
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) SetCycleEndTime(v int64) *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	s.CycleEndTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) SetCycleStartTime(v int64) *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	s.CycleStartTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) SetCycleSurplusValue(v float64) *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	s.CycleSurplusValue = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) SetCycleTotalValue(v float64) *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	s.CycleTotalValue = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) SetEquityType(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	s.EquityType = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) SetEquityUnit(v string) *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList {
	s.EquityUnit = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyDataSubscriptionInfoEquityList) Validate() error {
	return dara.Validate(s)
}
