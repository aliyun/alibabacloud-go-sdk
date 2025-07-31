// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRecoverTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceRecoverTimeResponseBody
	GetRequestId() *string
	SetRestoreRanges(v []*DescribeInstanceRecoverTimeResponseBodyRestoreRanges) *DescribeInstanceRecoverTimeResponseBody
	GetRestoreRanges() []*DescribeInstanceRecoverTimeResponseBodyRestoreRanges
}

type DescribeInstanceRecoverTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8CA8312-530A-413A-9129-F2BB32A8D404
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time ranges to which data can be restored. The time ranges include those used for point-in-time data restoration.
	RestoreRanges []*DescribeInstanceRecoverTimeResponseBodyRestoreRanges `json:"RestoreRanges,omitempty" xml:"RestoreRanges,omitempty" type:"Repeated"`
}

func (s DescribeInstanceRecoverTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRecoverTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRecoverTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRecoverTimeResponseBody) GetRestoreRanges() []*DescribeInstanceRecoverTimeResponseBodyRestoreRanges {
	return s.RestoreRanges
}

func (s *DescribeInstanceRecoverTimeResponseBody) SetRequestId(v string) *DescribeInstanceRecoverTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRecoverTimeResponseBody) SetRestoreRanges(v []*DescribeInstanceRecoverTimeResponseBodyRestoreRanges) *DescribeInstanceRecoverTimeResponseBody {
	s.RestoreRanges = v
	return s
}

func (s *DescribeInstanceRecoverTimeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceRecoverTimeResponseBodyRestoreRanges struct {
	// The beginning of the time range to which data can be restored.
	//
	// example:
	//
	// 2023-10-16T19:33:20Z
	RestoreBeginTime *string `json:"RestoreBeginTime,omitempty" xml:"RestoreBeginTime,omitempty"`
	// The end of the time range to which data can be restored.
	//
	// example:
	//
	// 2023-10-16T19:43:20Z
	RestoreEndTime *string `json:"RestoreEndTime,omitempty" xml:"RestoreEndTime,omitempty"`
	// The method used to restore data. Valid value:
	//
	// 	- PointInTime (default): Data is restored to a point in time.
	//
	// example:
	//
	// PointInTime
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
}

func (s DescribeInstanceRecoverTimeResponseBodyRestoreRanges) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRecoverTimeResponseBodyRestoreRanges) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) GetRestoreBeginTime() *string {
	return s.RestoreBeginTime
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) GetRestoreEndTime() *string {
	return s.RestoreEndTime
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) SetRestoreBeginTime(v string) *DescribeInstanceRecoverTimeResponseBodyRestoreRanges {
	s.RestoreBeginTime = &v
	return s
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) SetRestoreEndTime(v string) *DescribeInstanceRecoverTimeResponseBodyRestoreRanges {
	s.RestoreEndTime = &v
	return s
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) SetRestoreType(v string) *DescribeInstanceRecoverTimeResponseBodyRestoreRanges {
	s.RestoreType = &v
	return s
}

func (s *DescribeInstanceRecoverTimeResponseBodyRestoreRanges) Validate() error {
	return dara.Validate(s)
}
