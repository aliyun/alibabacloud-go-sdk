// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnWafUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnWafUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnWafUsageDataResponseBody
	GetStartTime() *string
	SetWafUsageData(v *DescribeDcdnWafUsageDataResponseBodyWafUsageData) *DescribeDcdnWafUsageDataResponseBody
	GetWafUsageData() *DescribeDcdnWafUsageDataResponseBodyWafUsageData
}

type DescribeDcdnWafUsageDataResponseBody struct {
	// The operation that you want to perform. Set the value to **DescribeDcdnWafUsageData**.
	//
	// example:
	//
	// 2018-10-01T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies how query results are grouped. By default, this parameter is empty. Valid values:
	//
	// 	- domain: Query results are grouped by accelerated domain name.
	//
	// 	- An empty string: Query results are not grouped.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-802B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z",
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of monitored requests.
	WafUsageData *DescribeDcdnWafUsageDataResponseBodyWafUsageData `json:"WafUsageData,omitempty" xml:"WafUsageData,omitempty" type:"Struct"`
}

func (s DescribeDcdnWafUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnWafUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnWafUsageDataResponseBody) GetWafUsageData() *DescribeDcdnWafUsageDataResponseBodyWafUsageData {
	return s.WafUsageData
}

func (s *DescribeDcdnWafUsageDataResponseBody) SetEndTime(v string) *DescribeDcdnWafUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBody) SetRequestId(v string) *DescribeDcdnWafUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBody) SetStartTime(v string) *DescribeDcdnWafUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBody) SetWafUsageData(v *DescribeDcdnWafUsageDataResponseBodyWafUsageData) *DescribeDcdnWafUsageDataResponseBody {
	s.WafUsageData = v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafUsageDataResponseBodyWafUsageData struct {
	WafUsageDataItem []*DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem `json:"WafUsageDataItem,omitempty" xml:"WafUsageDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnWafUsageDataResponseBodyWafUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafUsageDataResponseBodyWafUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageData) GetWafUsageDataItem() []*DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	return s.WafUsageDataItem
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageData) SetWafUsageDataItem(v []*DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) *DescribeDcdnWafUsageDataResponseBodyWafUsageData {
	s.WafUsageDataItem = v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem struct {
	// The number of blocked requests.
	//
	// example:
	//
	// 600
	AccessCnt *int64 `json:"AccessCnt,omitempty" xml:"AccessCnt,omitempty"`
	// The number of allowed requests.
	//
	// example:
	//
	// 300
	BlockCnt *int64 `json:"BlockCnt,omitempty" xml:"BlockCnt,omitempty"`
	// The domain name that you want to query. If you do not specify an accelerated domain name, all accelerated domain names are queried by default.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 300
	ObserveCnt *int64 `json:"ObserveCnt,omitempty" xml:"ObserveCnt,omitempty"`
	// The time granularity for a query. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day).
	//
	// example:
	//
	// 50
	SecCu *int64 `json:"SecCu,omitempty" xml:"SecCu,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GetAccessCnt() *int64 {
	return s.AccessCnt
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GetBlockCnt() *int64 {
	return s.BlockCnt
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GetObserveCnt() *int64 {
	return s.ObserveCnt
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GetSecCu() *int64 {
	return s.SecCu
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) SetAccessCnt(v int64) *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	s.AccessCnt = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) SetBlockCnt(v int64) *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	s.BlockCnt = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) SetDomain(v string) *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) SetObserveCnt(v int64) *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	s.ObserveCnt = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) SetSecCu(v int64) *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	s.SecCu = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) SetTimeStamp(v string) *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponseBodyWafUsageDataWafUsageDataItem) Validate() error {
	return dara.Validate(s)
}
