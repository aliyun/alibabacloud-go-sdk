// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDynamicRouteResponseBody
	GetRequestId() *string
}

type DeleteDynamicRouteResponseBody struct {
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDynamicRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDynamicRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDynamicRouteResponseBody) SetRequestId(v string) *DeleteDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDynamicRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
