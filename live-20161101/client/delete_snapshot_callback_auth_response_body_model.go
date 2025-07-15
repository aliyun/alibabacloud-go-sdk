// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotCallbackAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSnapshotCallbackAuthResponseBody
	GetRequestId() *string
}

type DeleteSnapshotCallbackAuthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotCallbackAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotCallbackAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotCallbackAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnapshotCallbackAuthResponseBody) SetRequestId(v string) *DeleteSnapshotCallbackAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotCallbackAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
