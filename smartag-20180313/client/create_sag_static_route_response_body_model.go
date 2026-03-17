// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSagStaticRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSagStaticRouteResponseBody
	GetRequestId() *string
}

type CreateSagStaticRouteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1120228A-E5E1-4E9C-B56D-96887E1A2B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSagStaticRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSagStaticRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSagStaticRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSagStaticRouteResponseBody) SetRequestId(v string) *CreateSagStaticRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSagStaticRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
