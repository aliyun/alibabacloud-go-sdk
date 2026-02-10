// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRecordUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLiveDomainRecordUsageDataResponseBody
	GetEndTime() *string
	SetRecordUsageData(v *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) *DescribeLiveDomainRecordUsageDataResponseBody
	GetRecordUsageData() *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData
	SetRequestId(v string) *DescribeLiveDomainRecordUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainRecordUsageDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainRecordUsageDataResponseBody struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2021-05-10T21:00:00Z
	EndTime         *string                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RecordUsageData *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData `json:"RecordUsageData,omitempty" xml:"RecordUsageData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B460F8B-993C-4F48-B98A-910811DEBFEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2021-05-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRecordUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRecordUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) GetRecordUsageData() *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData {
	return s.RecordUsageData
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) SetEndTime(v string) *DescribeLiveDomainRecordUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) SetRecordUsageData(v *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) *DescribeLiveDomainRecordUsageDataResponseBody {
	s.RecordUsageData = v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) SetRequestId(v string) *DescribeLiveDomainRecordUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) SetStartTime(v string) *DescribeLiveDomainRecordUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBody) Validate() error {
	if s.RecordUsageData != nil {
		if err := s.RecordUsageData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData struct {
	DataModule []*DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) GetDataModule() []*DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) SetDataModule(v []*DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageData) Validate() error {
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

type DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule struct {
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) GetType() *string {
	return s.Type
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) SetCount(v int64) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	s.Count = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) SetDomain(v string) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) SetDuration(v int64) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	s.Duration = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) SetRegion(v string) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) SetTimeStamp(v string) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) SetType(v string) *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule {
	s.Type = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponseBodyRecordUsageDataDataModule) Validate() error {
	return dara.Validate(s)
}
