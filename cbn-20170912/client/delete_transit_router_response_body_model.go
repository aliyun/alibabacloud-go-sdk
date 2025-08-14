// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6DE3EE92-39C8-4BBD-A3AD-F568D74741BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterResponseBody) SetRequestId(v string) *DeleteTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
