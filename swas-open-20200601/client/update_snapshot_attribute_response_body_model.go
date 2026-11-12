// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSnapshotAttributeResponseBody
	GetRequestId() *string
}

type UpdateSnapshotAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSnapshotAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSnapshotAttributeResponseBody) SetRequestId(v string) *UpdateSnapshotAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSnapshotAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
