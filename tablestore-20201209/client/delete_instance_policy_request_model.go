// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeleteInstancePolicyRequest
	GetInstanceName() *string
	SetPolicyVersion(v int64) *DeleteInstancePolicyRequest
	GetPolicyVersion() *int64
}

type DeleteInstancePolicyRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-test-12345
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The version of the instance policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s DeleteInstancePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancePolicyRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteInstancePolicyRequest) GetPolicyVersion() *int64 {
	return s.PolicyVersion
}

func (s *DeleteInstancePolicyRequest) SetInstanceName(v string) *DeleteInstancePolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteInstancePolicyRequest) SetPolicyVersion(v int64) *DeleteInstancePolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *DeleteInstancePolicyRequest) Validate() error {
	return dara.Validate(s)
}
