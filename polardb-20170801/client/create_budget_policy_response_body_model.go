// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetPolicyId(v string) *CreateBudgetPolicyResponseBody
	GetBudgetPolicyId() *string
	SetGwClusterId(v string) *CreateBudgetPolicyResponseBody
	GetGwClusterId() *string
	SetRequestId(v string) *CreateBudgetPolicyResponseBody
	GetRequestId() *string
}

type CreateBudgetPolicyResponseBody struct {
	// example:
	//
	// 05a5a8603df444a8a605af712ffexxx
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBudgetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBudgetPolicyResponseBody) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *CreateBudgetPolicyResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateBudgetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBudgetPolicyResponseBody) SetBudgetPolicyId(v string) *CreateBudgetPolicyResponseBody {
	s.BudgetPolicyId = &v
	return s
}

func (s *CreateBudgetPolicyResponseBody) SetGwClusterId(v string) *CreateBudgetPolicyResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *CreateBudgetPolicyResponseBody) SetRequestId(v string) *CreateBudgetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBudgetPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
