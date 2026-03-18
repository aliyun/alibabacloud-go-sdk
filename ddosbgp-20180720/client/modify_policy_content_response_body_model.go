// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPolicyContentResponseBody
	GetRequestId() *string
}

type ModifyPolicyContentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3777EF25-940B-51F4-BB1D-99B5********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolicyContentResponseBody) SetRequestId(v string) *ModifyPolicyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolicyContentResponseBody) Validate() error {
	return dara.Validate(s)
}
