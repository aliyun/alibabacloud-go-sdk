// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAppSnapshotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAppSnapshotConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveAppSnapshotConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAppSnapshotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAppSnapshotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppSnapshotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAppSnapshotConfigResponseBody) SetRequestId(v string) *DeleteLiveAppSnapshotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
