// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveSnapshotNotifyConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveSnapshotNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A3136B58-5876-5A93-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveSnapshotNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveSnapshotNotifyConfigResponseBody) SetRequestId(v string) *UpdateLiveSnapshotNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
