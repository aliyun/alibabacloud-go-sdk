// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainSnapshotDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVsDomainSnapshotDataResponseBody
	GetRequestId() *string
	SetSnapshotDataPerInterval(v *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) *DescribeVsDomainSnapshotDataResponseBody
	GetSnapshotDataPerInterval() *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval
}

type DescribeVsDomainSnapshotDataResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId               *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotDataPerInterval *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval `json:"SnapshotDataPerInterval,omitempty" xml:"SnapshotDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVsDomainSnapshotDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainSnapshotDataResponseBody) GetSnapshotDataPerInterval() *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval {
	return s.SnapshotDataPerInterval
}

func (s *DescribeVsDomainSnapshotDataResponseBody) SetRequestId(v string) *DescribeVsDomainSnapshotDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBody) SetSnapshotDataPerInterval(v *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) *DescribeVsDomainSnapshotDataResponseBody {
	s.SnapshotDataPerInterval = v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval struct {
	DataModule []*DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) GetDataModule() []*DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) SetDataModule(v []*DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule struct {
	// example:
	//
	// 1
	SnapshotValue *string `json:"SnapshotValue,omitempty" xml:"SnapshotValue,omitempty"`
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) GetSnapshotValue() *string {
	return s.SnapshotValue
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) SetSnapshotValue(v string) *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule {
	s.SnapshotValue = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponseBodySnapshotDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
