// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEnterpriseSnapshotPolicyResponseBody
	GetRequestId() *string
}

type UpdateEnterpriseSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FED145A8-7D5F-5C60-B054-4EB2899A5996
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *UpdateEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
