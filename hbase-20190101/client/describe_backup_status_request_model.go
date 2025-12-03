// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeBackupStatusRequest
	GetClusterId() *string
}

type DescribeBackupStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupStatusRequest) SetClusterId(v string) *DescribeBackupStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupStatusRequest) Validate() error {
	return dara.Validate(s)
}
