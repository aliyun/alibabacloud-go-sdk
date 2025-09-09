// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTimesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBackupTimesResponseBody
	GetRequestId() *string
	SetRestoreTime(v *DescribeBackupTimesResponseBodyRestoreTime) *DescribeBackupTimesResponseBody
	GetRestoreTime() *DescribeBackupTimesResponseBodyRestoreTime
	SetSuccess(v bool) *DescribeBackupTimesResponseBody
	GetSuccess() *bool
}

type DescribeBackupTimesResponseBody struct {
	// Indicates the ID of the request.
	//
	// example:
	//
	// 4780A19F-5ECB-4C56-AD20-966A3FF9DE5R
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates the information about the time range within which the data of the instance can be restored to a point in time.
	RestoreTime *DescribeBackupTimesResponseBodyRestoreTime `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupTimesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTimesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupTimesResponseBody) GetRestoreTime() *DescribeBackupTimesResponseBodyRestoreTime {
	return s.RestoreTime
}

func (s *DescribeBackupTimesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupTimesResponseBody) SetRequestId(v string) *DescribeBackupTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetRestoreTime(v *DescribeBackupTimesResponseBodyRestoreTime) *DescribeBackupTimesResponseBody {
	s.RestoreTime = v
	return s
}

func (s *DescribeBackupTimesResponseBody) SetSuccess(v bool) *DescribeBackupTimesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupTimesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTimesResponseBodyRestoreTime struct {
	// Indicates the end time. The time is in the UNIX timestamp format. The time is in UTC. Unit: ms.
	//
	// example:
	//
	// 1568636922671
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates the start time. The time is in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// example:
	//
	// 1568632853000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupTimesResponseBodyRestoreTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTimesResponseBodyRestoreTime) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) SetEndTime(v string) *DescribeBackupTimesResponseBodyRestoreTime {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) SetStartTime(v string) *DescribeBackupTimesResponseBodyRestoreTime {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTimesResponseBodyRestoreTime) Validate() error {
	return dara.Validate(s)
}
