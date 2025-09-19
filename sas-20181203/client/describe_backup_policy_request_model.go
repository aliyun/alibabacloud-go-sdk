// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeBackupPolicyRequest
	GetId() *int64
}

type DescribeBackupPolicyRequest struct {
	// The ID of the anti-ransomware policy.
	//
	// >  You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the IDs of anti-ransomware policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51880
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeBackupPolicyRequest) SetId(v int64) *DescribeBackupPolicyRequest {
	s.Id = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
