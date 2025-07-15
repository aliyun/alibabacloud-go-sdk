// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneCenterPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupId(v string) *CloneCenterPolicyResponseBody
	GetPolicyGroupId() *string
	SetRequestId(v string) *CloneCenterPolicyResponseBody
	GetRequestId() *string
}

type CloneCenterPolicyResponseBody struct {
	// The ID of the duplicated cloud computer policy.
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

func (s CloneCenterPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneCenterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CloneCenterPolicyResponseBody) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CloneCenterPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneCenterPolicyResponseBody) SetPolicyGroupId(v string) *CloneCenterPolicyResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *CloneCenterPolicyResponseBody) SetRequestId(v string) *CloneCenterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneCenterPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
