// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeBackupSetsRequest
	GetDrdsInstanceId() *string
	SetEndTime(v string) *DescribeBackupSetsRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeBackupSetsRequest
	GetStartTime() *string
}

type DescribeBackupSetsRequest struct {
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end of the query time which is in timestamp format (measured in millisecond) .
	//
	// >  The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1591326000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the query time which is in timestamp format (measured in millisecond).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1591327800000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBackupSetsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupSetsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupSetsRequest) SetDrdsInstanceId(v string) *DescribeBackupSetsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetEndTime(v string) *DescribeBackupSetsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetsRequest) SetStartTime(v string) *DescribeBackupSetsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupSetsRequest) Validate() error {
	return dara.Validate(s)
}
