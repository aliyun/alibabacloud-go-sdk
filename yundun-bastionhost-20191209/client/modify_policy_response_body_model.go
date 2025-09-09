// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPolicyResponseBody
	GetRequestId() *string
}

type ModifyPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A34F9731-A4EA-5D1D-8445-BAEF782BFFE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolicyResponseBody) SetRequestId(v string) *ModifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
