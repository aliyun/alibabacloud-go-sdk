// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainRecordDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordDataPerInterval(v *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) *DescribeVsDomainRecordDataResponseBody
	GetRecordDataPerInterval() *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval
	SetRequestId(v string) *DescribeVsDomainRecordDataResponseBody
	GetRequestId() *string
}

type DescribeVsDomainRecordDataResponseBody struct {
	RecordDataPerInterval *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval `json:"RecordDataPerInterval,omitempty" xml:"RecordDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainRecordDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponseBody) GetRecordDataPerInterval() *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval {
	return s.RecordDataPerInterval
}

func (s *DescribeVsDomainRecordDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainRecordDataResponseBody) SetRecordDataPerInterval(v *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) *DescribeVsDomainRecordDataResponseBody {
	s.RecordDataPerInterval = v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBody) SetRequestId(v string) *DescribeVsDomainRecordDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval struct {
	DataModule []*DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) GetDataModule() []*DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) SetDataModule(v []*DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule struct {
	// example:
	//
	// 100
	RecordValue *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	// example:
	//
	// 1
	StreamCountValue *string `json:"StreamCountValue,omitempty" xml:"StreamCountValue,omitempty"`
	// example:
	//
	// 2021-11-19T15:59:59Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) GetRecordValue() *string {
	return s.RecordValue
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) GetStreamCountValue() *string {
	return s.StreamCountValue
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) SetRecordValue(v string) *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	s.RecordValue = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) SetStreamCountValue(v string) *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	s.StreamCountValue = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponseBodyRecordDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
