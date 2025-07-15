// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupId(v string) *CreatePolicyGroupResponseBody
	GetPolicyGroupId() *string
	SetRequestId(v string) *CreatePolicyGroupResponseBody
	GetRequestId() *string
}

type CreatePolicyGroupResponseBody struct {
	// The cloud computer policy ID.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponseBody) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreatePolicyGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyGroupResponseBody) SetPolicyGroupId(v string) *CreatePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *CreatePolicyGroupResponseBody) SetRequestId(v string) *CreatePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
