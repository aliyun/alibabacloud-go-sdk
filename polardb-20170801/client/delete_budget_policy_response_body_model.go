// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBudgetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBudgetPolicyResponseBody
	GetRequestId() *string
}

type DeleteBudgetPolicyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBudgetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBudgetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBudgetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBudgetPolicyResponseBody) SetRequestId(v string) *DeleteBudgetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBudgetPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
