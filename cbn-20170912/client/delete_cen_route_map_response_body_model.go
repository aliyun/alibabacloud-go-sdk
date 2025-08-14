// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenRouteMapResponseBody
	GetRequestId() *string
}

type DeleteCenRouteMapResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5903EE99-D542-4E14-BC65-AAC1CB2D3D03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenRouteMapResponseBody) SetRequestId(v string) *DeleteCenRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
