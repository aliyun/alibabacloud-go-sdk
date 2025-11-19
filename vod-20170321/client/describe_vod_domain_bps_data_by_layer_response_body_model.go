// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainBpsDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataInterval(v *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeVodDomainBpsDataByLayerResponseBody
	GetBpsDataInterval() *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval
	SetDataInterval(v int32) *DescribeVodDomainBpsDataByLayerResponseBody
	GetDataInterval() *int32
	SetRequestId(v string) *DescribeVodDomainBpsDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeVodDomainBpsDataByLayerResponseBody struct {
	// The bandwidth returned at each time interval. Unit: bit/s.
	BpsDataInterval *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval `json:"BpsDataInterval,omitempty" xml:"BpsDataInterval,omitempty" type:"Struct"`
	// The time interval between the entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *int32 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainBpsDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) GetBpsDataInterval() *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval {
	return s.BpsDataInterval
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) GetDataInterval() *int32 {
	return s.DataInterval
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) SetBpsDataInterval(v *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeVodDomainBpsDataByLayerResponseBody {
	s.BpsDataInterval = v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) SetDataInterval(v int32) *DescribeVodDomainBpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) SetRequestId(v string) *DescribeVodDomainBpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBody) Validate() error {
	if s.BpsDataInterval != nil {
		if err := s.BpsDataInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval struct {
	DataModule []*DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) GetDataModule() []*DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) SetDataModule(v []*DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataInterval) Validate() error {
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

type DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule struct {
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-02-08T10:09:19Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 1000
	TrafficValue *int64 `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
	// The peak bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 75.33
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetTrafficValue() *int64 {
	return s.TrafficValue
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GetValue() *float64 {
	return s.Value
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTrafficValue(v int64) *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetValue(v float64) *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
