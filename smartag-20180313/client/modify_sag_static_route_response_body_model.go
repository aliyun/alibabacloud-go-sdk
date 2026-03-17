// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagStaticRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagStaticRouteResponseBody
	GetRequestId() *string
}

type ModifySagStaticRouteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ACA5A9FE-77FA-4C09-980B-5C353160FA4A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagStaticRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagStaticRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagStaticRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagStaticRouteResponseBody) SetRequestId(v string) *ModifySagStaticRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagStaticRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
