// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenChildInstanceRouteEntryToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCenChildInstanceRouteEntryToCenResponseBody
	GetRequestId() *string
}

type CreateCenChildInstanceRouteEntryToCenResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 437ED236-BE47-5370-8695-15C58C7A8014
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenChildInstanceRouteEntryToCenResponseBody) SetRequestId(v string) *CreateCenChildInstanceRouteEntryToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenResponseBody) Validate() error {
	return dara.Validate(s)
}
