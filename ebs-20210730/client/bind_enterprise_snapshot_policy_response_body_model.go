// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEnterpriseSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindEnterpriseSnapshotPolicyResponseBody
	GetRequestId() *string
}

type BindEnterpriseSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF4CA176-3358-5B74-B317-B1908B4B1F7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindEnterpriseSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BindEnterpriseSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *BindEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
