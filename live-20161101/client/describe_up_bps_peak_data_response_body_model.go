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
	// The information about peak inbound bandwidth on each day.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeUpBpsPeakDataResponseBodyDescribeUpPeakTrafficsDescribeUpPeakTraffic struct {
	// The daily peak inbound bandwidth.
	//
	// example:
	//
	// 777.2727083333333
	BandWidth *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The time when the daily peak bandwidth is reached.
	//
	// example:
	//
	// 1522180000000
	PeakTime *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	// The time queried on the day.
	//
	// example:
	//
	// 1522080000000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The category of the statistical data. If the DomainSwitch parameter is set to on, the value of this parameter is the domain name. If the DomainSwitch parameter is set to off or not specified, the value of this parameter is the user ID.
	//
	// example:
	//
	// push-live.aliyuncs.com
	StatName *string `json:"StatName,omitempty" xml:"StatName,omitempty"`
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
