// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvRealTimeQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregateData(v []*DescribeKvRealTimeQpsDataResponseBodyAggregateData) *DescribeKvRealTimeQpsDataResponseBody
	GetAggregateData() []*DescribeKvRealTimeQpsDataResponseBodyAggregateData
	SetEndTime(v string) *DescribeKvRealTimeQpsDataResponseBody
	GetEndTime() *string
	SetKvQpsData(v []*DescribeKvRealTimeQpsDataResponseBodyKvQpsData) *DescribeKvRealTimeQpsDataResponseBody
	GetKvQpsData() []*DescribeKvRealTimeQpsDataResponseBodyKvQpsData
	SetRequestId(v string) *DescribeKvRealTimeQpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeKvRealTimeQpsDataResponseBody
	GetStartTime() *string
}

type DescribeKvRealTimeQpsDataResponseBody struct {
	AggregateData []*DescribeKvRealTimeQpsDataResponseBodyAggregateData `json:"AggregateData,omitempty" xml:"AggregateData,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-01-18T15:59:59Z
	EndTime   *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	KvQpsData []*DescribeKvRealTimeQpsDataResponseBodyKvQpsData `json:"KvQpsData,omitempty" xml:"KvQpsData,omitempty" type:"Repeated"`
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92C***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-01-10T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeKvRealTimeQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKvRealTimeQpsDataResponseBody) GetAggregateData() []*DescribeKvRealTimeQpsDataResponseBodyAggregateData {
	return s.AggregateData
}

func (s *DescribeKvRealTimeQpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeKvRealTimeQpsDataResponseBody) GetKvQpsData() []*DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	return s.KvQpsData
}

func (s *DescribeKvRealTimeQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKvRealTimeQpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeKvRealTimeQpsDataResponseBody) SetAggregateData(v []*DescribeKvRealTimeQpsDataResponseBodyAggregateData) *DescribeKvRealTimeQpsDataResponseBody {
	s.AggregateData = v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBody) SetEndTime(v string) *DescribeKvRealTimeQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBody) SetKvQpsData(v []*DescribeKvRealTimeQpsDataResponseBodyKvQpsData) *DescribeKvRealTimeQpsDataResponseBody {
	s.KvQpsData = v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeKvRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBody) SetStartTime(v string) *DescribeKvRealTimeQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeKvRealTimeQpsDataResponseBodyAggregateData struct {
	// example:
	//
	// 123
	Acc *int64 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// example:
	//
	// get
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 1234
	KeyAcc *int64 `json:"KeyAcc,omitempty" xml:"KeyAcc,omitempty"`
	// example:
	//
	// 1233
	KeySuccAcc *int64 `json:"KeySuccAcc,omitempty" xml:"KeySuccAcc,omitempty"`
}

func (s DescribeKvRealTimeQpsDataResponseBodyAggregateData) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvRealTimeQpsDataResponseBodyAggregateData) GoString() string {
	return s.String()
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) GetAcc() *int64 {
	return s.Acc
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) GetKeyAcc() *int64 {
	return s.KeyAcc
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) GetKeySuccAcc() *int64 {
	return s.KeySuccAcc
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) SetAcc(v int64) *DescribeKvRealTimeQpsDataResponseBodyAggregateData {
	s.Acc = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) SetAccessType(v string) *DescribeKvRealTimeQpsDataResponseBodyAggregateData {
	s.AccessType = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) SetKeyAcc(v int64) *DescribeKvRealTimeQpsDataResponseBodyAggregateData {
	s.KeyAcc = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) SetKeySuccAcc(v int64) *DescribeKvRealTimeQpsDataResponseBodyAggregateData {
	s.KeySuccAcc = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyAggregateData) Validate() error {
	return dara.Validate(s)
}

type DescribeKvRealTimeQpsDataResponseBodyKvQpsData struct {
	// example:
	//
	// get
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 1234
	KeyQps *int64 `json:"KeyQps,omitempty" xml:"KeyQps,omitempty"`
	// example:
	//
	// 1233
	KeySuccQps *int64 `json:"KeySuccQps,omitempty" xml:"KeySuccQps,omitempty"`
	// example:
	//
	// 534167033424646***
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// 5236
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// example:
	//
	// 2023-01-10T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeKvRealTimeQpsDataResponseBodyKvQpsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GoString() string {
	return s.String()
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GetKeyQps() *int64 {
	return s.KeyQps
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GetKeySuccQps() *int64 {
	return s.KeySuccQps
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GetQps() *int64 {
	return s.Qps
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) SetAccessType(v string) *DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	s.AccessType = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) SetKeyQps(v int64) *DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	s.KeyQps = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) SetKeySuccQps(v int64) *DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	s.KeySuccQps = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) SetNamespaceId(v string) *DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) SetQps(v int64) *DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	s.Qps = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) SetTimeStamp(v string) *DescribeKvRealTimeQpsDataResponseBodyKvQpsData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponseBodyKvQpsData) Validate() error {
	return dara.Validate(s)
}
