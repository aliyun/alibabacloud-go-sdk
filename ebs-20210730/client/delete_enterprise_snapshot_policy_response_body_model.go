// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnterpriseSnapshotPolicyResponseBody
	GetRequestId() *string
}

type DeleteEnterpriseSnapshotPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B9F716DF-FAFD-50FD-B962-BCE0C837639A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnterpriseSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *DeleteEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
