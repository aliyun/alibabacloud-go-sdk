// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAppSnapshotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveAppSnapshotConfigResponseBody
	GetRequestId() *string
}

type AddLiveAppSnapshotConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveAppSnapshotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppSnapshotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveAppSnapshotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveAppSnapshotConfigResponseBody) SetRequestId(v string) *AddLiveAppSnapshotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveAppSnapshotConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
