// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRepairPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAutoRepairPolicyResponseBody
	GetRequestId() *string
}

type ModifyAutoRepairPolicyResponseBody struct {
	// example:
	//
	// db82195b-75a8-40e5-9be4-16f182******
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s ModifyAutoRepairPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoRepairPolicyResponseBody) SetRequestId(v string) *ModifyAutoRepairPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoRepairPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
