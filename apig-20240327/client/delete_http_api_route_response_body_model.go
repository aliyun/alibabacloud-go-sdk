// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHttpApiRouteResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHttpApiRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHttpApiRouteResponseBody
	GetRequestId() *string
}

type DeleteHttpApiRouteResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F138FFC-6E2B-56C1-9BAB-A67462E339D1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteHttpApiRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHttpApiRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHttpApiRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpApiRouteResponseBody) SetCode(v string) *DeleteHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpApiRouteResponseBody) SetMessage(v string) *DeleteHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpApiRouteResponseBody) SetRequestId(v string) *DeleteHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpApiRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
