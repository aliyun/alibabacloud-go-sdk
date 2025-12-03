// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeBackupPlanConfigRequest
	GetClusterId() *string
}

type DescribeBackupPlanConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupPlanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupPlanConfigRequest) SetClusterId(v string) *DescribeBackupPlanConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPlanConfigRequest) Validate() error {
	return dara.Validate(s)
}
