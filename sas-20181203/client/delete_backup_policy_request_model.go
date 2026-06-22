// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteBackupPolicyRequest
	GetId() *int64
	SetPolicyVersion(v string) *DeleteBackupPolicyRequest
	GetPolicyVersion() *string
}

type DeleteBackupPolicyRequest struct {
	// The ID of the mitigation policies that you want to delete.
	//
	// >Invoke the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The version of the mitigation policies that you want to delete. You can invoke the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query this parameter. Valid values:
	//
	// - **1.0.0**: The mitigation policies version is 1.0.0.
	//
	// - **2.0.0**: The mitigation policies version is 2.0.0.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s DeleteBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteBackupPolicyRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *DeleteBackupPolicyRequest) SetId(v int64) *DeleteBackupPolicyRequest {
	s.Id = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetPolicyVersion(v string) *DeleteBackupPolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *DeleteBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
