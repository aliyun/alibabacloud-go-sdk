// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeBackupTimesRequest
	GetDrdsInstanceId() *string
}

type DescribeBackupTimesRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga71nn****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupTimesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTimesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBackupTimesRequest) SetDrdsInstanceId(v string) *DescribeBackupTimesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupTimesRequest) Validate() error {
	return dara.Validate(s)
}
