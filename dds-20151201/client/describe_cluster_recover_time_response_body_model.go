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
	// The list of cluster backup sets. A cluster backup contains the backup set of each node.
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
	if s.RestoreRanges != nil {
		for _, item := range s.RestoreRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterRecoverTimeResponseBodyRestoreRanges struct {
	// The beginning of the restorable time range.
	//
	// example:
	//
	// 2023-10-16T19:33:20Z
	RestoreBeginTime *string `json:"RestoreBeginTime,omitempty" xml:"RestoreBeginTime,omitempty"`
	// The end of the restorable time range.
	//
	// example:
	//
	// 2023-10-16T19:43:20Z
	RestoreEndTime *string `json:"RestoreEndTime,omitempty" xml:"RestoreEndTime,omitempty"`
	// The restoration method. Valid values:
	//
	// 	- **PointInTime*	- (default): point-in-time restoration.
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
