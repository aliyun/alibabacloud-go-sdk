// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPvUvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainPvUvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainPvUvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPvUvDataResponseBody
	GetEndTime() *string
	SetPvUvDataInfos(v *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) *DescribeLiveDomainPvUvDataResponseBody
	GetPvUvDataInfos() *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos
	SetRequestId(v string) *DescribeLiveDomainPvUvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainPvUvDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainPvUvDataResponseBody struct {
	// The time interval between the entries returned. Unit: seconds. Default value: 3600.
	//
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-03-20T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data of PVs and UVs.
	PvUvDataInfos *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos `json:"PvUvDataInfos,omitempty" xml:"PvUvDataInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E9D3257A-1B7C-414C-90C1-8D07AC47BCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-03-17T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainPvUvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPvUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPvUvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainPvUvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPvUvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPvUvDataResponseBody) GetPvUvDataInfos() *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos {
	return s.PvUvDataInfos
}

func (s *DescribeLiveDomainPvUvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainPvUvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPvUvDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainPvUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBody) SetDomainName(v string) *DescribeLiveDomainPvUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBody) SetEndTime(v string) *DescribeLiveDomainPvUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBody) SetPvUvDataInfos(v *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) *DescribeLiveDomainPvUvDataResponseBody {
	s.PvUvDataInfos = v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBody) SetRequestId(v string) *DescribeLiveDomainPvUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBody) SetStartTime(v string) *DescribeLiveDomainPvUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBody) Validate() error {
	if s.PvUvDataInfos != nil {
		if err := s.PvUvDataInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos struct {
	PvUvDataInfo []*DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo `json:"PvUvDataInfo,omitempty" xml:"PvUvDataInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) GetPvUvDataInfo() []*DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	return s.PvUvDataInfo
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) SetPvUvDataInfo(v []*DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos {
	s.PvUvDataInfo = v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfos) Validate() error {
	if s.PvUvDataInfo != nil {
		for _, item := range s.PvUvDataInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo struct {
	// The number of PVs.
	//
	// example:
	//
	// 3036
	PV *string `json:"PV,omitempty" xml:"PV,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-03-19T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The number of UVs.
	//
	// example:
	//
	// 2
	UV *string `json:"UV,omitempty" xml:"UV,omitempty"`
}

func (s DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GetPV() *string {
	return s.PV
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) GetUV() *string {
	return s.UV
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetPV(v string) *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.PV = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetTimeStamp(v string) *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) SetUV(v string) *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo {
	s.UV = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponseBodyPvUvDataInfosPvUvDataInfo) Validate() error {
	return dara.Validate(s)
}
