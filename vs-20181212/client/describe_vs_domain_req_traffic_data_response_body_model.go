// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainReqTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainReqTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainReqTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainReqTrafficDataResponseBody
	GetEndTime() *string
	SetReqTrafficDataPerInterval(v *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) *DescribeVsDomainReqTrafficDataResponseBody
	GetReqTrafficDataPerInterval() *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval
	SetRequestId(v string) *DescribeVsDomainReqTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainReqTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeVsDomainReqTrafficDataResponseBody struct {
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
	// 2021-09-22T03:40:41Z
	EndTime                   *string                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReqTrafficDataPerInterval *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval `json:"ReqTrafficDataPerInterval,omitempty" xml:"ReqTrafficDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// 9BEC5E85-C76B-56EF-A922-860EFDB8B64B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-09-21T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) GetReqTrafficDataPerInterval() *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval {
	return s.ReqTrafficDataPerInterval
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetDataInterval(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetDomainName(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetEndTime(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetReqTrafficDataPerInterval(v *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) *DescribeVsDomainReqTrafficDataResponseBody {
	s.ReqTrafficDataPerInterval = v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetRequestId(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) SetStartTime(v string) *DescribeVsDomainReqTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBody) Validate() error {
	if s.ReqTrafficDataPerInterval != nil {
		if err := s.ReqTrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval struct {
	DataModule []*DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) GetDataModule() []*DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) SetDataModule(v []*DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerInterval) Validate() error {
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

type DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule struct {
	// example:
	//
	// 10000
	ReqTrafficValue *string `json:"ReqTrafficValue,omitempty" xml:"ReqTrafficValue,omitempty"`
	// example:
	//
	// 2021-09-22T03:40:41Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) GetReqTrafficValue() *string {
	return s.ReqTrafficValue
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) SetReqTrafficValue(v string) *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule {
	s.ReqTrafficValue = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponseBodyReqTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
