// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLocalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeBackupLocalRequest
	GetDrdsInstanceId() *string
}

type DescribeBackupLocalRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupLocalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLocalRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBackupLocalRequest) SetDrdsInstanceId(v string) *DescribeBackupLocalRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupLocalRequest) Validate() error {
	return dara.Validate(s)
}
