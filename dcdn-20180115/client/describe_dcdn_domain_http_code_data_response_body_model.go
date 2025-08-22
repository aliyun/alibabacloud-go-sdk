// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDataPerInterval(v *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) *DescribeDcdnDomainHttpCodeDataResponseBody
	GetDataPerInterval() *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval
	SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainHttpCodeDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The proportions of HTTP status codes at each time interval.
	DataPerInterval *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval `json:"DataPerInterval,omitempty" xml:"DataPerInterval,omitempty" type:"Struct"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2018-03-01T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91FC2D9D-B042-4634-8A5C-7B8E7482C22D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-03-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) GetDataPerInterval() *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval {
	return s.DataPerInterval
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetDataPerInterval(v *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.DataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval struct {
	DataModule []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) GetDataModule() []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) SetDataModule(v []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule struct {
	// The proportions of the HTTP status codes.
	HttpCodeDataPerInterval *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval `json:"HttpCodeDataPerInterval,omitempty" xml:"HttpCodeDataPerInterval,omitempty" type:"Struct"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2019-03-01T13:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) GetHttpCodeDataPerInterval() *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval {
	return s.HttpCodeDataPerInterval
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) SetHttpCodeDataPerInterval(v *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	s.HttpCodeDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval struct {
	HttpCodeDataModule []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule `json:"HttpCodeDataModule,omitempty" xml:"HttpCodeDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) GetHttpCodeDataModule() []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	return s.HttpCodeDataModule
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) SetHttpCodeDataModule(v []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval {
	s.HttpCodeDataModule = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 404
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The count of each HTTP status code.
	//
	// example:
	//
	// 1
	Count *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 33.333333
	Proportion *float32 `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) GetCount() *float32 {
	return s.Count
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) GetProportion() *float32 {
	return s.Proportion
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetCode(v int32) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetCount(v float32) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetProportion(v float32) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) Validate() error {
	return dara.Validate(s)
}
