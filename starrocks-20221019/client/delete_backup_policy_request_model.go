// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteBackupPolicyRequest
	GetInstanceId() *string
	SetPolicyId(v string) *DeleteBackupPolicyRequest
	GetPolicyId() *string
}

type DeleteBackupPolicyRequest struct {
	// example:
	//
	// c-0104730e9d*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// bp-298ahiu289
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBackupPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeleteBackupPolicyRequest) SetInstanceId(v string) *DeleteBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetPolicyId(v string) *DeleteBackupPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
