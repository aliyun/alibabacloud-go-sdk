// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTimeShiftDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveDomainTimeShiftDataResponseBody
	GetRequestId() *string
	SetTimeShiftData(v *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) *DescribeLiveDomainTimeShiftDataResponseBody
	GetTimeShiftData() *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData
}

type DescribeLiveDomainTimeShiftDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8AE1CB3A-6510-442E-A6B9-EF03D05B3E09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time shifting usage data that was collected for each time interval.
	TimeShiftData *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData `json:"TimeShiftData,omitempty" xml:"TimeShiftData,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainTimeShiftDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainTimeShiftDataResponseBody) GetTimeShiftData() *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData {
	return s.TimeShiftData
}

func (s *DescribeLiveDomainTimeShiftDataResponseBody) SetRequestId(v string) *DescribeLiveDomainTimeShiftDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseBody) SetTimeShiftData(v *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) *DescribeLiveDomainTimeShiftDataResponseBody {
	s.TimeShiftData = v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseBody) Validate() error {
	if s.TimeShiftData != nil {
		if err := s.TimeShiftData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData struct {
	DataModule []*DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) GetDataModule() []*DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) SetDataModule(v []*DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftData) Validate() error {
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

type DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule struct {
	// The storage used for time shifting. Unit: bytes.
	//
	// example:
	//
	// 1664165660
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2021-03-03T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The type of time shifting. Examples: HLS_D1 and HLS_D7.
	//
	// example:
	//
	// HLS_D7
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) GetSize() *string {
	return s.Size
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) GetType() *string {
	return s.Type
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) SetSize(v string) *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule {
	s.Size = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) SetTimeStamp(v string) *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) SetType(v string) *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule {
	s.Type = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseBodyTimeShiftDataDataModule) Validate() error {
	return dara.Validate(s)
}
