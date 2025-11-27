// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDevicesDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevicesDataPerInterval(v *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) *DescribeVsDevicesDataResponseBody
	GetDevicesDataPerInterval() *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval
	SetRequestId(v string) *DescribeVsDevicesDataResponseBody
	GetRequestId() *string
}

type DescribeVsDevicesDataResponseBody struct {
	DevicesDataPerInterval *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval `json:"DevicesDataPerInterval,omitempty" xml:"DevicesDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDevicesDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDevicesDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponseBody) GetDevicesDataPerInterval() *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval {
	return s.DevicesDataPerInterval
}

func (s *DescribeVsDevicesDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDevicesDataResponseBody) SetDevicesDataPerInterval(v *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) *DescribeVsDevicesDataResponseBody {
	s.DevicesDataPerInterval = v
	return s
}

func (s *DescribeVsDevicesDataResponseBody) SetRequestId(v string) *DescribeVsDevicesDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDevicesDataResponseBody) Validate() error {
	if s.DevicesDataPerInterval != nil {
		if err := s.DevicesDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDevicesDataResponseBodyDevicesDataPerInterval struct {
	DataModule []*DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) GetDataModule() []*DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) SetDataModule(v []*DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerInterval) Validate() error {
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

type DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule struct {
	// example:
	//
	// 128
	DevicesDataValue *string `json:"DevicesDataValue,omitempty" xml:"DevicesDataValue,omitempty"`
	// example:
	//
	// 2022-01-04T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) GetDevicesDataValue() *string {
	return s.DevicesDataValue
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) SetDevicesDataValue(v string) *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule {
	s.DevicesDataValue = &v
	return s
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDevicesDataResponseBodyDevicesDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
