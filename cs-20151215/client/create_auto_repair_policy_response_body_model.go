// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoRepairPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *CreateAutoRepairPolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreateAutoRepairPolicyResponseBody
	GetRequestId() *string
}

type CreateAutoRepairPolicyResponseBody struct {
	// example:
	//
	// r-xxxxxxx
	PolicyId *string `json:"policy_id,omitempty" xml:"policy_id,omitempty"`
	// example:
	//
	// E368C761-F8F6-4A36-9B58-BD53D5******
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s CreateAutoRepairPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateAutoRepairPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoRepairPolicyResponseBody) SetPolicyId(v string) *CreateAutoRepairPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateAutoRepairPolicyResponseBody) SetRequestId(v string) *CreateAutoRepairPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoRepairPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
