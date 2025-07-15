// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveSnapshotNotifyConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveSnapshotNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 32A96B9A-F203-4EC5-8E43-CB92E68F4I79E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveSnapshotNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveSnapshotNotifyConfigResponseBody) SetRequestId(v string) *DeleteLiveSnapshotNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
