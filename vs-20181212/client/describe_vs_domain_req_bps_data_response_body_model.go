// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainReqBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainReqBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainReqBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainReqBpsDataResponseBody
	GetEndTime() *string
	SetReqBpsDataPerInterval(v *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) *DescribeVsDomainReqBpsDataResponseBody
	GetReqBpsDataPerInterval() *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval
	SetRequestId(v string) *DescribeVsDomainReqBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainReqBpsDataResponseBody
	GetStartTime() *string
}

type DescribeVsDomainReqBpsDataResponseBody struct {
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-09-24T03:30:46Z
	EndTime               *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReqBpsDataPerInterval *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval `json:"ReqBpsDataPerInterval,omitempty" xml:"ReqBpsDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-12-24T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainReqBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainReqBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainReqBpsDataResponseBody) GetReqBpsDataPerInterval() *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval {
	return s.ReqBpsDataPerInterval
}

func (s *DescribeVsDomainReqBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainReqBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetDataInterval(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetDomainName(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetEndTime(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetReqBpsDataPerInterval(v *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) *DescribeVsDomainReqBpsDataResponseBody {
	s.ReqBpsDataPerInterval = v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetRequestId(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) SetStartTime(v string) *DescribeVsDomainReqBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBody) Validate() error {
	if s.ReqBpsDataPerInterval != nil {
		if err := s.ReqBpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval struct {
	DataModule []*DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) GetDataModule() []*DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) SetDataModule(v []*DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerInterval) Validate() error {
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

type DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule struct {
	// example:
	//
	// 1000
	ReqBpsValue *string `json:"ReqBpsValue,omitempty" xml:"ReqBpsValue,omitempty"`
	// example:
	//
	// 2021-12-24T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) GetReqBpsValue() *string {
	return s.ReqBpsValue
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) SetReqBpsValue(v string) *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule {
	s.ReqBpsValue = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponseBodyReqBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
