// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteServiceInCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteServiceInCenResponseBody
	GetRequestId() *string
}

type DeleteRouteServiceInCenResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteServiceInCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteServiceInCenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteServiceInCenResponseBody) SetRequestId(v string) *DeleteRouteServiceInCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteServiceInCenResponseBody) Validate() error {
	return dara.Validate(s)
}
