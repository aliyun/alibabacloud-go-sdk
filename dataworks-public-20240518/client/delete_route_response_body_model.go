// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRouteResponseBody
	GetSuccess() *bool
}

type DeleteRouteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRouteResponseBody) SetRequestId(v string) *DeleteRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteResponseBody) SetSuccess(v bool) *DeleteRouteResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
