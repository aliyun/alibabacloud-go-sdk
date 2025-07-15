// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody
	GetRequestId() *string
}

type CreateAutoSnapshotPolicyResponseBody struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-3hpa78d0qyt99****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateAutoSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
