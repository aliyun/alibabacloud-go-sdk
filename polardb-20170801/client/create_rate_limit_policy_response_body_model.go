// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRateLimitPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *CreateRateLimitPolicyResponseBody
	GetGwClusterId() *string
	SetPolicyId(v string) *CreateRateLimitPolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreateRateLimitPolicyResponseBody
	GetRequestId() *string
}

type CreateRateLimitPolicyResponseBody struct {
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 02eccf7c61cf4d05a543075ee907f3**
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRateLimitPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRateLimitPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRateLimitPolicyResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateRateLimitPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateRateLimitPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRateLimitPolicyResponseBody) SetGwClusterId(v string) *CreateRateLimitPolicyResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *CreateRateLimitPolicyResponseBody) SetPolicyId(v string) *CreateRateLimitPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateRateLimitPolicyResponseBody) SetRequestId(v string) *CreateRateLimitPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRateLimitPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
