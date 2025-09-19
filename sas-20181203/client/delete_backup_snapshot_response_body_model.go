// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackupSnapshotResponseBody
	GetRequestId() *string
}

type DeleteBackupSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupSnapshotResponseBody) SetRequestId(v string) *DeleteBackupSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
