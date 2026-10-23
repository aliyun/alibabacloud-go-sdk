// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateILMPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateILMPolicyRequest
	GetClientToken() *string
	SetBody(v string) *CreateILMPolicyRequest
	GetBody() *string
}

type CreateILMPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name and phase configurations of the index lifecycle policy to be created.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateILMPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateILMPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateILMPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateILMPolicyRequest) GetBody() *string {
	return s.Body
}

func (s *CreateILMPolicyRequest) SetClientToken(v string) *CreateILMPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateILMPolicyRequest) SetBody(v string) *CreateILMPolicyRequest {
	s.Body = &v
	return s
}

func (s *CreateILMPolicyRequest) Validate() error {
	return dara.Validate(s)
}
