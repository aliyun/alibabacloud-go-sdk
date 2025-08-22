// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
	GetEndTime() *string
	SetHttpCodeDataPerInterval(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
	GetHttpCodeDataPerInterval() *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
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
	// The HTTP status code.
	HttpCodeDataPerInterval *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval `json:"HttpCodeDataPerInterval,omitempty" xml:"HttpCodeDataPerInterval,omitempty" type:"Struct"`
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
	// 2018-03-01T05:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GetHttpCodeDataPerInterval() *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval {
	return s.HttpCodeDataPerInterval
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetHttpCodeDataPerInterval(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.HttpCodeDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval struct {
	DataModule []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) GetDataModule() []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) SetDataModule(v []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2018-03-01T13:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The proportions of the HTTP status codes.
	WebsocketHttpCode *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode `json:"WebsocketHttpCode,omitempty" xml:"WebsocketHttpCode,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) GetWebsocketHttpCode() *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode {
	return s.WebsocketHttpCode
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) SetWebsocketHttpCode(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule {
	s.WebsocketHttpCode = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode struct {
	HttpCodeDataModule []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule `json:"HttpCodeDataModule,omitempty" xml:"HttpCodeDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) GetHttpCodeDataModule() []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	return s.HttpCodeDataModule
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) SetHttpCodeDataModule(v []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode {
	s.HttpCodeDataModule = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 404
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	Count *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 100
	Proportion *float32 `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) GetCount() *float32 {
	return s.Count
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) GetProportion() *float32 {
	return s.Proportion
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) SetCode(v int32) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) SetCount(v float32) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) SetProportion(v float32) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) Validate() error {
	return dara.Validate(s)
}
