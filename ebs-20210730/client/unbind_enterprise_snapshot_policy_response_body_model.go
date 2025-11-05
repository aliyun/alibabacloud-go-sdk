// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEnterpriseSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindEnterpriseSnapshotPolicyResponseBody
	GetRequestId() *string
}

type UnbindEnterpriseSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 061DE1AB-08BA-5ACD-A03A-440117C6939A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindEnterpriseSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEnterpriseSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *UnbindEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
