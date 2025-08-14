// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLivePackageChannelResponseBody
	GetRequestId() *string
}

type DeleteLivePackageChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20B3A1B6-4BD2-5DE6-BCBC-098C9B4F4E91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLivePackageChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePackageChannelResponseBody) SetRequestId(v string) *DeleteLivePackageChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePackageChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
