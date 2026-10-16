// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetPolicyRequest
	GetLanguage() *string
	SetPolicyName(v string) *GetPolicyRequest
	GetPolicyName() *string
	SetPolicyType(v string) *GetPolicyRequest
	GetPolicyType() *string
}

type GetPolicyRequest struct {
	// The language type. The description of the system access policy is returned in this language. Valid values:
	//
	// - en: English.
	//
	// - zh-CN: Chinese.
	//
	// - ja: Japanese.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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
	// The type of the access policy. Valid values:
	//
	// - Custom: custom policy.
	//
	// - System: system policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetPolicyRequest) SetLanguage(v string) *GetPolicyRequest {
	s.Language = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyName(v string) *GetPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyType(v string) *GetPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
