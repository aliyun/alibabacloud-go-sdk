// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *CreateEnterpriseSnapshotPolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreateEnterpriseSnapshotPolicyResponseBody
	GetRequestId() *string
}

type CreateEnterpriseSnapshotPolicyResponseBody struct {
	// The id of a policy.
	//
	// example:
	//
	// esp-xxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A8959DA-1E04-5724-8288-58334031454E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) SetPolicyId(v string) *CreateEnterpriseSnapshotPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *CreateEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
