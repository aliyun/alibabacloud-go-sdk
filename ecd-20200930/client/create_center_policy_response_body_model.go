// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenterPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupId(v string) *CreateCenterPolicyResponseBody
	GetPolicyGroupId() *string
	SetRequestId(v string) *CreateCenterPolicyResponseBody
	GetRequestId() *string
}

type CreateCenterPolicyResponseBody struct {
	// The cloud computer policy ID.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenterPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyResponseBody) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateCenterPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenterPolicyResponseBody) SetPolicyGroupId(v string) *CreateCenterPolicyResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateCenterPolicyResponseBody) SetRequestId(v string) *CreateCenterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenterPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
