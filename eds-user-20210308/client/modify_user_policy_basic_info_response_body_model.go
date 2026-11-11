// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPolicyBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserPolicyBasicInfoResponseBody
	GetRequestId() *string
}

type ModifyUserPolicyBasicInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserPolicyBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPolicyBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPolicyBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserPolicyBasicInfoResponseBody) SetRequestId(v string) *ModifyUserPolicyBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserPolicyBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
