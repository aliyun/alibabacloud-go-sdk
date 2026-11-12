// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSnapshotsResponseBody
	GetRequestId() *string
}

type DeleteSnapshotsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C2DE174B-7196-5778-A00D-6EA2601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnapshotsResponseBody) SetRequestId(v string) *DeleteSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}
