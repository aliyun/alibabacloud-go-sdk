// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetControlPolicyRequest
	GetLanguage() *string
	SetPolicyId(v string) *GetControlPolicyRequest
	GetPolicyId() *string
}

type GetControlPolicyRequest struct {
	// The language in which you want to return the description of the access control policy. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// > This parameter is valid only for system access control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-SImPt8GCEwiq****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetControlPolicyRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetControlPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetControlPolicyRequest) SetLanguage(v string) *GetControlPolicyRequest {
	s.Language = &v
	return s
}

func (s *GetControlPolicyRequest) SetPolicyId(v string) *GetControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
