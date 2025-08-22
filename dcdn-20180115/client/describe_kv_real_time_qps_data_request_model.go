// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvRealTimeQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *DescribeKvRealTimeQpsDataRequest
	GetAccessType() *string
	SetEndTime(v string) *DescribeKvRealTimeQpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeKvRealTimeQpsDataRequest
	GetInterval() *string
	SetNamespaceId(v string) *DescribeKvRealTimeQpsDataRequest
	GetNamespaceId() *string
	SetSplitBy(v string) *DescribeKvRealTimeQpsDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeKvRealTimeQpsDataRequest
	GetStartTime() *string
}

type DescribeKvRealTimeQpsDataRequest struct {
	// example:
	//
	// get
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 2022-08-10T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 534167033424646***
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// type
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// example:
	//
	// 2022-08-10T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeKvRealTimeQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeKvRealTimeQpsDataRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeKvRealTimeQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeKvRealTimeQpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeKvRealTimeQpsDataRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeKvRealTimeQpsDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeKvRealTimeQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeKvRealTimeQpsDataRequest) SetAccessType(v string) *DescribeKvRealTimeQpsDataRequest {
	s.AccessType = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataRequest) SetEndTime(v string) *DescribeKvRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataRequest) SetInterval(v string) *DescribeKvRealTimeQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataRequest) SetNamespaceId(v string) *DescribeKvRealTimeQpsDataRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataRequest) SetSplitBy(v string) *DescribeKvRealTimeQpsDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataRequest) SetStartTime(v string) *DescribeKvRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
