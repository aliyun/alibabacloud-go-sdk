// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpBpsPeakDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeUpPeakTraffics(v *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) *DescribeUpBpsPeakDataResponseBody
	GetDescribeUpPeakTraffics() *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics
	SetRequestId(v string) *DescribeUpBpsPeakDataResponseBody
	GetRequestId() *string
}

type DescribeUpBpsPeakDataResponseBody struct {
	DescribeUpPeakTraffics *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics `json:"DescribeUpPeakTraffics,omitempty" xml:"DescribeUpPeakTraffics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpBpsPeakDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponseBody) GetDescribeUpPeakTraffics() *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics {
	return s.DescribeUpPeakTraffics
}

func (s *DescribeUpBpsPeakDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpBpsPeakDataResponseBody) SetDescribeUpPeakTraffics(v *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) *DescribeUpBpsPeakDataResponseBody {
	s.DescribeUpPeakTraffics = v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBody) SetRequestId(v string) *DescribeUpBpsPeakDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBody) Validate() error {
	if s.DescribeUpPeakTraffics != nil {
		if err := s.DescribeUpPeakTraffics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics struct {
	DescribeUpPeakTraffic []*DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic `json:"DescribeUpPeakTraffic,omitempty" xml:"DescribeUpPeakTraffic,omitempty" type:"Repeated"`
}

func (s DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) GetDescribeUpPeakTraffic() []*DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	return s.DescribeUpPeakTraffic
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) SetDescribeUpPeakTraffic(v []*DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics {
	s.DescribeUpPeakTraffic = v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTraffics) Validate() error {
	if s.DescribeUpPeakTraffic != nil {
		for _, item := range s.DescribeUpPeakTraffic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic struct {
	BandWidth *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	PeakTime  *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	StatName  *string `json:"StatName,omitempty" xml:"StatName,omitempty"`
}

func (s DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) GetBandWidth() *string {
	return s.BandWidth
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) GetPeakTime() *string {
	return s.PeakTime
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) GetStatName() *string {
	return s.StatName
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetBandWidth(v string) *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.BandWidth = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetPeakTime(v string) *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.PeakTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetQueryTime(v string) *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.QueryTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetStatName(v string) *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.StatName = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic) Validate() error {
	return dara.Validate(s)
}
