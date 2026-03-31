// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *GetPolicyVersionRequest
	GetPolicyName() *string
	SetPolicyType(v string) *GetPolicyVersionRequest
	GetPolicyType() *string
	SetVersionId(v string) *GetPolicyVersionRequest
	GetVersionId() *string
}

type GetPolicyVersionRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetPolicyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyVersionRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetPolicyVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetPolicyVersionRequest) SetPolicyName(v string) *GetPolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyVersionRequest) SetPolicyType(v string) *GetPolicyVersionRequest {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyVersionRequest) SetVersionId(v string) *GetPolicyVersionRequest {
	s.VersionId = &v
	return s
}

func (s *GetPolicyVersionRequest) Validate() error {
	return dara.Validate(s)
}
