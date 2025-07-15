// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotDetectPornConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveSnapshotDetectPornConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveSnapshotDetectPornConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveSnapshotDetectPornConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotDetectPornConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotDetectPornConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveSnapshotDetectPornConfigResponseBody) SetRequestId(v string) *UpdateLiveSnapshotDetectPornConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
