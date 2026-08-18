// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSnapshotRequest struct {
}

func (s GetSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
