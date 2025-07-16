// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDomainBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeDomainBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeDomainBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainBpsDataResponseBody
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDomainBpsDataResponseBody
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainBpsDataResponseBody
	GetLocationNameEn() *string
	SetRequestId(v string) *DescribeDomainBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDomainBpsDataResponseBody struct {
	// The list of bandwidth data entries returned at each interval.
	BpsDataPerInterval *DescribeDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
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
	// 2015-12-10T20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the ISP.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2020-05-14T09:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBody) GetBpsDataPerInterval() *DescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeDomainBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainBpsDataResponseBody) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainBpsDataResponseBody) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDomainName(v string) *DescribeDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetEndTime(v string) *DescribeDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetIspNameEn(v string) *DescribeDomainBpsDataResponseBody {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetLocationNameEn(v string) *DescribeDomainBpsDataResponseBody {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetRequestId(v string) *DescribeDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetStartTime(v string) *DescribeDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	// The bandwidth value in the Chinese mainland. When the bandwidth data is queried by ISP, this parameter is empty.
	//
	// example:
	//
	// 11286111
	DomesticValue *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	// The bandwidth data for HTTPS requests in the Chinese mainland. When the bandwidth data is queried by ISP, this parameter is empty.
	//
	// example:
	//
	// 11286111
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	// The bandwidth data for HTTPS requests in regions outside the Chinese mainland. When the bandwidth data is queried by ISP, this parameter is empty.
	//
	// example:
	//
	// 2000
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	// The bandwidth value for HTTPS requests. Unit: bit/s.
	//
	// example:
	//
	// 11288111
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The bandwidth data in regions outside the Chinese mainland. When the bandwidth data is queried by ISP, this parameter is empty.
	//
	// example:
	//
	// 2000
	OverseasValue *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 11288111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsDomesticValue() *string {
	return s.HttpsDomesticValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsOverseasValue() *string {
	return s.HttpsOverseasValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
