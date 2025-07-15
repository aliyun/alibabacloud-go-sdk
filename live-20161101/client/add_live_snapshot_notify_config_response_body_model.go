// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveSnapshotNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveSnapshotNotifyConfigResponseBody
	GetRequestId() *string
}

type AddLiveSnapshotNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveSnapshotNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveSnapshotNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveSnapshotNotifyConfigResponseBody) SetRequestId(v string) *AddLiveSnapshotNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
