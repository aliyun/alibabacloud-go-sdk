// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeKvUsageDataResponseBody
	GetEndTime() *string
	SetKvUsageData(v []*DescribeKvUsageDataResponseBodyKvUsageData) *DescribeKvUsageDataResponseBody
	GetKvUsageData() []*DescribeKvUsageDataResponseBodyKvUsageData
	SetRequestId(v string) *DescribeKvUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeKvUsageDataResponseBody
	GetStartTime() *string
}

type DescribeKvUsageDataResponseBody struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2022-11-18T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The usage details.
	KvUsageData []*DescribeKvUsageDataResponseBodyKvUsageData `json:"KvUsageData,omitempty" xml:"KvUsageData,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2022-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeKvUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKvUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeKvUsageDataResponseBody) GetKvUsageData() []*DescribeKvUsageDataResponseBodyKvUsageData {
	return s.KvUsageData
}

func (s *DescribeKvUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKvUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeKvUsageDataResponseBody) SetEndTime(v string) *DescribeKvUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeKvUsageDataResponseBody) SetKvUsageData(v []*DescribeKvUsageDataResponseBodyKvUsageData) *DescribeKvUsageDataResponseBody {
	s.KvUsageData = v
	return s
}

func (s *DescribeKvUsageDataResponseBody) SetRequestId(v string) *DescribeKvUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKvUsageDataResponseBody) SetStartTime(v string) *DescribeKvUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeKvUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeKvUsageDataResponseBodyKvUsageData struct {
	// The number of visits.
	//
	// example:
	//
	// 1340000
	Acc *int64 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The request method. This parameter is available only when the **SplitBy*	- parameter is set to **type**.
	//
	// example:
	//
	// get
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The namespace ID. This parameter is available only when the **SplitBy*	- parameter is set to **namespace**.
	//
	// example:
	//
	// 534167033424646144
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2022-11-14T15:00:03Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeKvUsageDataResponseBodyKvUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvUsageDataResponseBodyKvUsageData) GoString() string {
	return s.String()
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) GetAcc() *int64 {
	return s.Acc
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) SetAcc(v int64) *DescribeKvUsageDataResponseBodyKvUsageData {
	s.Acc = &v
	return s
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) SetAccessType(v string) *DescribeKvUsageDataResponseBodyKvUsageData {
	s.AccessType = &v
	return s
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) SetNamespaceId(v string) *DescribeKvUsageDataResponseBodyKvUsageData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) SetTimeStamp(v string) *DescribeKvUsageDataResponseBodyKvUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeKvUsageDataResponseBodyKvUsageData) Validate() error {
	return dara.Validate(s)
}
