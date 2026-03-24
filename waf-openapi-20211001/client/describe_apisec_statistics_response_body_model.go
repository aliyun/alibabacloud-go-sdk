// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeApisecStatisticsResponseBodyData) *DescribeApisecStatisticsResponseBody
	GetData() *DescribeApisecStatisticsResponseBodyData
	SetRequestId(v string) *DescribeApisecStatisticsResponseBody
	GetRequestId() *string
}

type DescribeApisecStatisticsResponseBody struct {
	// The statistical results of API security risks or events.
	Data *DescribeApisecStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 221F0F14-54C6-59A1-9967-72***81B61A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecStatisticsResponseBody) GetData() *DescribeApisecStatisticsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecStatisticsResponseBody) SetData(v *DescribeApisecStatisticsResponseBodyData) *DescribeApisecStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecStatisticsResponseBody) SetRequestId(v string) *DescribeApisecStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisecStatisticsResponseBodyData struct {
	// The number of accounts associated with the monitored APIs.
	//
	// example:
	//
	// 1
	Account *int64 `json:"Account,omitempty" xml:"Account,omitempty"`
	// The number of handled items.
	//
	// example:
	//
	// 1
	Actioned *int64 `json:"Actioned,omitempty" xml:"Actioned,omitempty"`
	// The number of APIs that are monitored by the API security module.
	//
	// example:
	//
	// 202
	Api *int64 `json:"Api,omitempty" xml:"Api,omitempty"`
	// The number of confirmed items.
	//
	// example:
	//
	// 11
	Confirmed *int64 `json:"Confirmed,omitempty" xml:"Confirmed,omitempty"`
	// The number of domain names that are monitored by the API security module.
	//
	// example:
	//
	// 22
	Domain *int64 `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of items that are manually verified as fixed.
	//
	// example:
	//
	// 13
	Fixed *int64 `json:"Fixed,omitempty" xml:"Fixed,omitempty"`
	// The number of high-risk items.
	//
	// example:
	//
	// 135
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of ignored items.
	//
	// example:
	//
	// 3
	Ignore *int64 `json:"Ignore,omitempty" xml:"Ignore,omitempty"`
	// The number of low-risk items.
	//
	// example:
	//
	// 160
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-risk items.
	//
	// example:
	//
	// 27
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// The number of items that are verified as not fixed.
	//
	// example:
	//
	// 0
	NotFixed *int64 `json:"NotFixed,omitempty" xml:"NotFixed,omitempty"`
	// The number of items that are verified as fixed by the system.
	//
	// example:
	//
	// 1
	SystemFixed *int64 `json:"SystemFixed,omitempty" xml:"SystemFixed,omitempty"`
	// The number of items to be confirmed.
	//
	// example:
	//
	// 295
	ToBeConfirmed *int64 `json:"ToBeConfirmed,omitempty" xml:"ToBeConfirmed,omitempty"`
	// The number of items to be fixed.
	//
	// example:
	//
	// 3
	ToBeFixed *int64 `json:"ToBeFixed,omitempty" xml:"ToBeFixed,omitempty"`
	// The number of items to be verified by the system.
	//
	// example:
	//
	// 2
	ToBeVerified *int64 `json:"ToBeVerified,omitempty" xml:"ToBeVerified,omitempty"`
	// The number of new high-risk items detected today.
	//
	// example:
	//
	// 3
	TodayHigh *string `json:"TodayHigh,omitempty" xml:"TodayHigh,omitempty"`
	// The number of new low-risk items detected today.
	//
	// example:
	//
	// 0
	TodayLow *int64 `json:"TodayLow,omitempty" xml:"TodayLow,omitempty"`
	// The number of new medium-risk items detected today.
	//
	// example:
	//
	// 0
	TodayMedium *string `json:"TodayMedium,omitempty" xml:"TodayMedium,omitempty"`
	// The total number of new items detected today.
	//
	// example:
	//
	// 3
	TodayTotal *string `json:"TodayTotal,omitempty" xml:"TodayTotal,omitempty"`
	// The total number of items.
	//
	// > This is the sum of the values of the **High**, **Medium**, and **Low*	- response parameters.
	//
	// example:
	//
	// 322
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeApisecStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecStatisticsResponseBodyData) GetAccount() *int64 {
	return s.Account
}

func (s *DescribeApisecStatisticsResponseBodyData) GetActioned() *int64 {
	return s.Actioned
}

func (s *DescribeApisecStatisticsResponseBodyData) GetApi() *int64 {
	return s.Api
}

func (s *DescribeApisecStatisticsResponseBodyData) GetConfirmed() *int64 {
	return s.Confirmed
}

func (s *DescribeApisecStatisticsResponseBodyData) GetDomain() *int64 {
	return s.Domain
}

func (s *DescribeApisecStatisticsResponseBodyData) GetFixed() *int64 {
	return s.Fixed
}

func (s *DescribeApisecStatisticsResponseBodyData) GetHigh() *int64 {
	return s.High
}

func (s *DescribeApisecStatisticsResponseBodyData) GetIgnore() *int64 {
	return s.Ignore
}

func (s *DescribeApisecStatisticsResponseBodyData) GetLow() *int64 {
	return s.Low
}

func (s *DescribeApisecStatisticsResponseBodyData) GetMedium() *int64 {
	return s.Medium
}

func (s *DescribeApisecStatisticsResponseBodyData) GetNotFixed() *int64 {
	return s.NotFixed
}

func (s *DescribeApisecStatisticsResponseBodyData) GetSystemFixed() *int64 {
	return s.SystemFixed
}

func (s *DescribeApisecStatisticsResponseBodyData) GetToBeConfirmed() *int64 {
	return s.ToBeConfirmed
}

func (s *DescribeApisecStatisticsResponseBodyData) GetToBeFixed() *int64 {
	return s.ToBeFixed
}

func (s *DescribeApisecStatisticsResponseBodyData) GetToBeVerified() *int64 {
	return s.ToBeVerified
}

func (s *DescribeApisecStatisticsResponseBodyData) GetTodayHigh() *string {
	return s.TodayHigh
}

func (s *DescribeApisecStatisticsResponseBodyData) GetTodayLow() *int64 {
	return s.TodayLow
}

func (s *DescribeApisecStatisticsResponseBodyData) GetTodayMedium() *string {
	return s.TodayMedium
}

func (s *DescribeApisecStatisticsResponseBodyData) GetTodayTotal() *string {
	return s.TodayTotal
}

func (s *DescribeApisecStatisticsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeApisecStatisticsResponseBodyData) SetAccount(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Account = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetActioned(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Actioned = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetApi(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Api = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetConfirmed(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Confirmed = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetDomain(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Domain = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetFixed(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Fixed = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetHigh(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.High = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetIgnore(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Ignore = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetLow(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Low = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetMedium(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Medium = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetNotFixed(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.NotFixed = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetSystemFixed(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.SystemFixed = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetToBeConfirmed(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.ToBeConfirmed = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetToBeFixed(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.ToBeFixed = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetToBeVerified(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.ToBeVerified = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetTodayHigh(v string) *DescribeApisecStatisticsResponseBodyData {
	s.TodayHigh = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetTodayLow(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.TodayLow = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetTodayMedium(v string) *DescribeApisecStatisticsResponseBodyData {
	s.TodayMedium = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetTodayTotal(v string) *DescribeApisecStatisticsResponseBodyData {
	s.TodayTotal = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) SetTotal(v int64) *DescribeApisecStatisticsResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeApisecStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
