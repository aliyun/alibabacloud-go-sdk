// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserLoginPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserLoginPoliciesResponseBody
	GetRequestId() *string
}

type RemoveUserLoginPoliciesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserLoginPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserLoginPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserLoginPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserLoginPoliciesResponseBody) SetRequestId(v string) *RemoveUserLoginPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserLoginPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
