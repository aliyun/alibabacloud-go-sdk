// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupPolicyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpgradeBackupPolicyVersionRequest
	GetId() *int64
}

type UpgradeBackupPolicyVersionRequest struct {
	// The ID of the anti-ransomware mitigation policies.
	//
	// >You can invoke [DescribeBackupPolicies](~~DescribeBackupPolicies~~) to obtain this parameter.
	//
	// Only policy IDs with PolicyVersion set to 1.0.0 and UpgradeStatus set to NotUpgraded are supported. Otherwise, the API returns an InvalidParam fault.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpgradeBackupPolicyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupPolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPolicyVersionRequest) GetId() *int64 {
	return s.Id
}

func (s *UpgradeBackupPolicyVersionRequest) SetId(v int64) *UpgradeBackupPolicyVersionRequest {
	s.Id = &v
	return s
}

func (s *UpgradeBackupPolicyVersionRequest) Validate() error {
	return dara.Validate(s)
}
