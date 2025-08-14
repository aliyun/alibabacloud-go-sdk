// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageChannelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLivePackageChannelGroupResponseBody
	GetRequestId() *string
}

type DeleteLivePackageChannelGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5D87B753-0250-5D9D-B248-D40C3271F864
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLivePackageChannelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageChannelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageChannelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePackageChannelGroupResponseBody) SetRequestId(v string) *DeleteLivePackageChannelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePackageChannelGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
