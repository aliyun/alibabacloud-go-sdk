// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainPvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainPvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainPvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainPvDataResponseBody
	GetEndTime() *string
	SetPvDataInterval(v *DescribeVodDomainPvDataResponseBodyPvDataInterval) *DescribeVodDomainPvDataResponseBody
	GetPvDataInterval() *DescribeVodDomainPvDataResponseBodyPvDataInterval
	SetRequestId(v string) *DescribeVodDomainPvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainPvDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainPvDataResponseBody struct {
	DataInterval   *string                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvDataInterval *DescribeVodDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainPvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainPvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainPvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainPvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainPvDataResponseBody) GetPvDataInterval() *DescribeVodDomainPvDataResponseBodyPvDataInterval {
	return s.PvDataInterval
}

func (s *DescribeVodDomainPvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainPvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainPvDataResponseBody) SetDataInterval(v string) *DescribeVodDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBody) SetDomainName(v string) *DescribeVodDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBody) SetEndTime(v string) *DescribeVodDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBody) SetPvDataInterval(v *DescribeVodDomainPvDataResponseBodyPvDataInterval) *DescribeVodDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeVodDomainPvDataResponseBody) SetRequestId(v string) *DescribeVodDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBody) SetStartTime(v string) *DescribeVodDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainPvDataResponseBodyPvDataInterval struct {
	PvDataInterval []*DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainPvDataResponseBodyPvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataInterval) GetPvDataInterval() []*DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval {
	return s.PvDataInterval
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataInterval) SetPvDataInterval(v []*DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) *DescribeVodDomainPvDataResponseBodyPvDataInterval {
	s.PvDataInterval = v
	return s
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) SetTimeStamp(v string) *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) SetValue(v string) *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainPvDataResponseBodyPvDataIntervalPvDataInterval) Validate() error {
	return dara.Validate(s)
}
