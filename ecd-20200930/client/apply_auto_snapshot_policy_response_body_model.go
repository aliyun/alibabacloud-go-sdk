// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAutoSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyAutoSnapshotPolicyResponseBody
	GetRequestId() *string
}

type ApplyAutoSnapshotPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 63740E03-1B4B-5A18-AC27-2745A4F2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ApplyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
