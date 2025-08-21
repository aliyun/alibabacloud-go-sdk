// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainPvUvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainPvUvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainPvUvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainPvUvDataResponseBody
	GetEndTime() *string
	SetPvUvDataInfos(v *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) *DescribeVsDomainPvUvDataResponseBody
	GetPvUvDataInfos() *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos
	SetRequestId(v string) *DescribeVsDomainPvUvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainPvUvDataResponseBody
	GetStartTime() *string
}

type DescribeVsDomainPvUvDataResponseBody struct {
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
	// 2021-11-24T00:00:00Z
	EndTime       *string                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvUvDataInfos *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos `json:"PvUvDataInfos,omitempty" xml:"PvUvDataInfos,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-12-12T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvUvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainPvUvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainPvUvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainPvUvDataResponseBody) GetPvUvDataInfos() *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos {
	return s.PvUvDataInfos
}

func (s *DescribeVsDomainPvUvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainPvUvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetDataInterval(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetDomainName(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetEndTime(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetPvUvDataInfos(v *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) *DescribeVsDomainPvUvDataResponseBody {
	s.PvUvDataInfos = v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetRequestId(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) SetStartTime(v string) *DescribeVsDomainPvUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos struct {
	PvUvDataInfo []*DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo `json:"PvUvDataInfo,omitempty" xml:"PvUvDataInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) GetPvUvDataInfo() []*DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	return s.PvUvDataInfo
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) SetPvUvDataInfo(v []*DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos {
	s.PvUvDataInfo = v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo struct {
	// example:
	//
	// 100
	PV *string `json:"PV,omitempty" xml:"PV,omitempty"`
	// example:
	//
	// 2021-10-14T23:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 100
	UV *string `json:"UV,omitempty" xml:"UV,omitempty"`
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GetPV() *string {
	return s.PV
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GetUV() *string {
	return s.UV
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetPV(v string) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.PV = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetTimeStamp(v string) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetUV(v string) *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.UV = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) Validate() error {
	return dara.Validate(s)
}
