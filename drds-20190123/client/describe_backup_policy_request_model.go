// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeBackupPolicyRequest
	GetDrdsInstanceId() *string
}

type DescribeBackupPolicyRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga71nn****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBackupPolicyRequest) SetDrdsInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
