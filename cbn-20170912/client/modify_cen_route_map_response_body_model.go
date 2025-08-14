// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCenRouteMapResponseBody
	GetRequestId() *string
}

type ModifyCenRouteMapResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCenRouteMapResponseBody) SetRequestId(v string) *ModifyCenRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCenRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
