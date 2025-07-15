// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClonePolicyGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupId(v string) *ClonePolicyGroupResponseBody
	GetPolicyGroupId() *string
	SetRequestId(v string) *ClonePolicyGroupResponseBody
	GetRequestId() *string
}

type ClonePolicyGroupResponseBody struct {
	// The ID of the new cloud computer policy.
	//
	// example:
	//
	// pg-7jcaznnx6go6n****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClonePolicyGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClonePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponseBody) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ClonePolicyGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClonePolicyGroupResponseBody) SetPolicyGroupId(v string) *ClonePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupResponseBody) SetRequestId(v string) *ClonePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClonePolicyGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
