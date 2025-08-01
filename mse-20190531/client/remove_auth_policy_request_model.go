// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAuthPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *RemoveAuthPolicyRequest
	GetAcceptLanguage() *string
	SetPolicyId(v string) *RemoveAuthPolicyRequest
	GetPolicyId() *string
}

type RemoveAuthPolicyRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 259
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s RemoveAuthPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAuthPolicyRequest) GoString() string {
	return s.String()
}

func (s *RemoveAuthPolicyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *RemoveAuthPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *RemoveAuthPolicyRequest) SetAcceptLanguage(v string) *RemoveAuthPolicyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *RemoveAuthPolicyRequest) SetPolicyId(v string) *RemoveAuthPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *RemoveAuthPolicyRequest) Validate() error {
	return dara.Validate(s)
}
