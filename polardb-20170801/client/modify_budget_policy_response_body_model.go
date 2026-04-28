// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBudgetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *ModifyBudgetPolicyResponseBody
	GetGwClusterId() *string
	SetRequestId(v string) *ModifyBudgetPolicyResponseBody
	GetRequestId() *string
}

type ModifyBudgetPolicyResponseBody struct {
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBudgetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBudgetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBudgetPolicyResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyBudgetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBudgetPolicyResponseBody) SetGwClusterId(v string) *ModifyBudgetPolicyResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *ModifyBudgetPolicyResponseBody) SetRequestId(v string) *ModifyBudgetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBudgetPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
