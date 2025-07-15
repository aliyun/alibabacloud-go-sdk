// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVcoRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVcoRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteVcoRouteEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9208DDD8-0930-3CE6-AF7F-732B4E67B3DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVcoRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVcoRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVcoRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVcoRouteEntryResponseBody) SetRequestId(v string) *DeleteVcoRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVcoRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
