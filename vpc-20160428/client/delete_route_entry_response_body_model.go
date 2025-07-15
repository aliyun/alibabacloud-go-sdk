// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteRouteEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteEntryResponseBody) SetRequestId(v string) *DeleteRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
