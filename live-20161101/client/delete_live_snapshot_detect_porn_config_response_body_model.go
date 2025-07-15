// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotDetectPornConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveSnapshotDetectPornConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveSnapshotDetectPornConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E6*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveSnapshotDetectPornConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotDetectPornConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotDetectPornConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveSnapshotDetectPornConfigResponseBody) SetRequestId(v string) *DeleteLiveSnapshotDetectPornConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
