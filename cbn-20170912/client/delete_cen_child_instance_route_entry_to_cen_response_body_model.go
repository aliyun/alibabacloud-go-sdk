// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenChildInstanceRouteEntryToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenChildInstanceRouteEntryToCenResponseBody
	GetRequestId() *string
}

type DeleteCenChildInstanceRouteEntryToCenResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C05000A4-2FC5-5B2C-9527-954044DE2CF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponseBody) SetRequestId(v string) *DeleteCenChildInstanceRouteEntryToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponseBody) Validate() error {
	return dara.Validate(s)
}
