// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateILMPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateILMPolicyRequest
	GetClientToken() *string
	SetBody(v string) *UpdateILMPolicyRequest
	GetBody() *string
}

type UpdateILMPolicyRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateILMPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateILMPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateILMPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateILMPolicyRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateILMPolicyRequest) SetClientToken(v string) *UpdateILMPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateILMPolicyRequest) SetBody(v string) *UpdateILMPolicyRequest {
	s.Body = &v
	return s
}

func (s *UpdateILMPolicyRequest) Validate() error {
	return dara.Validate(s)
}
