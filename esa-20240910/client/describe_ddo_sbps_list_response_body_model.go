// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v int32) *DescribeDDoSBpsListResponseBody
	GetDataInterval() *int32
	SetDataModule(v []*DescribeDDoSBpsListResponseBodyDataModule) *DescribeDDoSBpsListResponseBody
	GetDataModule() []*DescribeDDoSBpsListResponseBodyDataModule
	SetEndTime(v string) *DescribeDDoSBpsListResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDDoSBpsListResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDDoSBpsListResponseBody
	GetStartTime() *string
}

type DescribeDDoSBpsListResponseBody struct {
	DataInterval *int32                                       `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DataModule   []*DescribeDDoSBpsListResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
	EndTime      *string                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSBpsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListResponseBody) GetDataInterval() *int32 {
	return s.DataInterval
}

func (s *DescribeDDoSBpsListResponseBody) GetDataModule() []*DescribeDDoSBpsListResponseBodyDataModule {
	return s.DataModule
}

func (s *DescribeDDoSBpsListResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSBpsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSBpsListResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSBpsListResponseBody) SetDataInterval(v int32) *DescribeDDoSBpsListResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetDataModule(v []*DescribeDDoSBpsListResponseBodyDataModule) *DescribeDDoSBpsListResponseBody {
	s.DataModule = v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetEndTime(v string) *DescribeDDoSBpsListResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetRequestId(v string) *DescribeDDoSBpsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetStartTime(v string) *DescribeDDoSBpsListResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) Validate() error {
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

type DescribeDDoSBpsListResponseBodyDataModule struct {
	AttackBps *int64  `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	AttackPps *int64  `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	NormalBps *int64  `json:"NormalBps,omitempty" xml:"NormalBps,omitempty"`
	NormalPps *int64  `json:"NormalPps,omitempty" xml:"NormalPps,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TotalBps  *int64  `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
	TotalPps  *int64  `json:"TotalPps,omitempty" xml:"TotalPps,omitempty"`
}

func (s DescribeDDoSBpsListResponseBodyDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListResponseBodyDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetAttackBps() *int64 {
	return s.AttackBps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetAttackPps() *int64 {
	return s.AttackPps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetNormalBps() *int64 {
	return s.NormalBps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetNormalPps() *int64 {
	return s.NormalPps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetTotalPps() *int64 {
	return s.TotalPps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetAttackBps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.AttackBps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetAttackPps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.AttackPps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetNormalBps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.NormalBps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetNormalPps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.NormalPps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetTimeStamp(v string) *DescribeDDoSBpsListResponseBodyDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetTotalBps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.TotalBps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetTotalPps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.TotalPps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) Validate() error {
	return dara.Validate(s)
}
