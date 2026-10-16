// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultPolicyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *SetDefaultPolicyVersionRequest
	GetPolicyName() *string
	SetVersionId(v string) *SetDefaultPolicyVersionRequest
	GetVersionId() *string
}

type SetDefaultPolicyVersionRequest struct {
	// Policy Name of the access policy.
	//
	// Policy Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The ID of the access policy version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v2
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s SetDefaultPolicyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultPolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *SetDefaultPolicyVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *SetDefaultPolicyVersionRequest) SetPolicyName(v string) *SetDefaultPolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *SetDefaultPolicyVersionRequest) SetVersionId(v string) *SetDefaultPolicyVersionRequest {
	s.VersionId = &v
	return s
}

func (s *SetDefaultPolicyVersionRequest) Validate() error {
	return dara.Validate(s)
}
