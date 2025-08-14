// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageOriginEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLivePackageOriginEndpointResponseBody
	GetRequestId() *string
}

type DeleteLivePackageOriginEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5D87B753-0250-5D9D-B248-D40C3271F864
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLivePackageOriginEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageOriginEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageOriginEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePackageOriginEndpointResponseBody) SetRequestId(v string) *DeleteLivePackageOriginEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePackageOriginEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
