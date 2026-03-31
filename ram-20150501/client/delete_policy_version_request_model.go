// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DeletePolicyVersionRequest
	GetPolicyName() *string
	SetVersionId(v string) *DeletePolicyVersionRequest
	GetVersionId() *string
}

type DeletePolicyVersionRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The ID of the policy version that you want to delete.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeletePolicyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyVersionRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DeletePolicyVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeletePolicyVersionRequest) SetPolicyName(v string) *DeletePolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *DeletePolicyVersionRequest) SetVersionId(v string) *DeletePolicyVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DeletePolicyVersionRequest) Validate() error {
	return dara.Validate(s)
}
