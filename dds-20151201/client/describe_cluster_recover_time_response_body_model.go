// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterRecoverTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeClusterRecoverTimeResponseBody
	GetRequestId() *string
	SetRestoreRanges(v []*DescribeClusterRecoverTimeResponseBodyRestoreRanges) *DescribeClusterRecoverTimeResponseBody
	GetRestoreRanges() []*DescribeClusterRecoverTimeResponseBodyRestoreRanges
}

type DescribeClusterRecoverTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 60EEBD77-227C-5B39-86EA-D89163C5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The cluster backup sets of the instance. A cluster backup file contains the backup sets of each node.
	RestoreRanges []*DescribeClusterRecoverTimeResponseBodyRestoreRanges `json:"RestoreRanges,omitempty" xml:"RestoreRanges,omitempty" type:"Repeated"`
}

func (s DescribeClusterRecoverTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterRecoverTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterRecoverTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterRecoverTimeResponseBody) GetRestoreRanges() []*DescribeClusterRecoverTimeResponseBodyRestoreRanges {
	return s.RestoreRanges
}

func (s *DescribeClusterRecoverTimeResponseBody) SetRequestId(v string) *DescribeClusterRecoverTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterRecoverTimeResponseBody) SetRestoreRanges(v []*DescribeClusterRecoverTimeResponseBodyRestoreRanges) *DescribeClusterRecoverTimeResponseBody {
	s.RestoreRanges = v
	return s
}

func (s *DescribeClusterRecoverTimeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterRecoverTimeResponseBodyRestoreRanges struct {
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
	// The method used to restore data. Valid values:
	//
	// 	- **PointInTime*	- (default): Data is restored based on point in time
	//
	// example:
	//
	// PointInTime
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
}

func (s DescribeClusterRecoverTimeResponseBodyRestoreRanges) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterRecoverTimeResponseBodyRestoreRanges) GoString() string {
	return s.String()
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) GetRestoreBeginTime() *string {
	return s.RestoreBeginTime
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) GetRestoreEndTime() *string {
	return s.RestoreEndTime
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) SetRestoreBeginTime(v string) *DescribeClusterRecoverTimeResponseBodyRestoreRanges {
	s.RestoreBeginTime = &v
	return s
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) SetRestoreEndTime(v string) *DescribeClusterRecoverTimeResponseBodyRestoreRanges {
	s.RestoreEndTime = &v
	return s
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) SetRestoreType(v string) *DescribeClusterRecoverTimeResponseBodyRestoreRanges {
	s.RestoreType = &v
	return s
}

func (s *DescribeClusterRecoverTimeResponseBodyRestoreRanges) Validate() error {
	return dara.Validate(s)
}
