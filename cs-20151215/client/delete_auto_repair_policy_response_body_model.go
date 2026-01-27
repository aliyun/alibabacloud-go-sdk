// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoRepairPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAutoRepairPolicyResponseBody
	GetRequestId() *string
}

type DeleteAutoRepairPolicyResponseBody struct {
	// example:
	//
	// A9891419-D125-4D89-AFCA-68846675****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s DeleteAutoRepairPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoRepairPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoRepairPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoRepairPolicyResponseBody) SetRequestId(v string) *DeleteAutoRepairPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoRepairPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
