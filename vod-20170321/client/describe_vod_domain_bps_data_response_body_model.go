// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeVodDomainBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeVodDomainBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainBpsDataResponseBody
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodDomainBpsDataResponseBody
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainBpsDataResponseBody
	GetLocationNameEn() *string
	SetRequestId(v string) *DescribeVodDomainBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainBpsDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// The time interval between the returned entries. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The domain name for CDN.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range in which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the ISP. By default, the data of all ISPs is returned.
	//
	// example:
	//
	// Alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. By default, the data in all regions is returned.
	//
	// example:
	//
	// cn-shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-****-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range in which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponseBody) GetBpsDataPerInterval() *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeVodDomainBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainBpsDataResponseBody) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainBpsDataResponseBody) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeVodDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeVodDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetDomainName(v string) *DescribeVodDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetEndTime(v string) *DescribeVodDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetIspNameEn(v string) *DescribeVodDomainBpsDataResponseBody {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetLocationNameEn(v string) *DescribeVodDomainBpsDataResponseBody {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetStartTime(v string) *DescribeVodDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) Validate() error {
	if s.BpsDataPerInterval != nil {
		if err := s.BpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	DomesticValue      *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	HttpsValue         *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	OverseasValue      *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDomesticValue() *string {
	return s.DomesticValue
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsDomesticValue() *string {
	return s.HttpsDomesticValue
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsOverseasValue() *string {
	return s.HttpsOverseasValue
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetOverseasValue() *string {
	return s.OverseasValue
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
