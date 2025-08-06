// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainUvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainUvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainUvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainUvDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainUvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainUvDataResponseBody
	GetStartTime() *string
	SetUvDataInterval(v *DescribeVodDomainUvDataResponseBodyUvDataInterval) *DescribeVodDomainUvDataResponseBody
	GetUvDataInterval() *DescribeVodDomainUvDataResponseBodyUvDataInterval
}

type DescribeVodDomainUvDataResponseBody struct {
	DataInterval   *string                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UvDataInterval *DescribeVodDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeVodDomainUvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainUvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainUvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainUvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainUvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainUvDataResponseBody) GetUvDataInterval() *DescribeVodDomainUvDataResponseBodyUvDataInterval {
	return s.UvDataInterval
}

func (s *DescribeVodDomainUvDataResponseBody) SetDataInterval(v string) *DescribeVodDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBody) SetDomainName(v string) *DescribeVodDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBody) SetEndTime(v string) *DescribeVodDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBody) SetRequestId(v string) *DescribeVodDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBody) SetStartTime(v string) *DescribeVodDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBody) SetUvDataInterval(v *DescribeVodDomainUvDataResponseBodyUvDataInterval) *DescribeVodDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

func (s *DescribeVodDomainUvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainUvDataResponseBodyUvDataInterval struct {
	UvDataInterval []*DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainUvDataResponseBodyUvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataInterval) GetUvDataInterval() []*DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval {
	return s.UvDataInterval
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataInterval) SetUvDataInterval(v []*DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) *DescribeVodDomainUvDataResponseBodyUvDataInterval {
	s.UvDataInterval = v
	return s
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) SetTimeStamp(v string) *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) SetValue(v string) *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainUvDataResponseBodyUvDataIntervalUvDataInterval) Validate() error {
	return dara.Validate(s)
}
