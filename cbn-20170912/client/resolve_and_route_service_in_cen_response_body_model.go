// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveAndRouteServiceInCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResolveAndRouteServiceInCenResponseBody
	GetRequestId() *string
}

type ResolveAndRouteServiceInCenResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C0245BEF-52AC-44A8-A776-EF96FD26A5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResolveAndRouteServiceInCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResolveAndRouteServiceInCenResponseBody) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResolveAndRouteServiceInCenResponseBody) SetRequestId(v string) *ResolveAndRouteServiceInCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenResponseBody) Validate() error {
	return dara.Validate(s)
}
