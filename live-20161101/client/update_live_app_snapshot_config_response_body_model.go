// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAppSnapshotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveAppSnapshotConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveAppSnapshotConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveAppSnapshotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppSnapshotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppSnapshotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveAppSnapshotConfigResponseBody) SetRequestId(v string) *UpdateLiveAppSnapshotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
