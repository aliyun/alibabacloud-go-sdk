// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHttpApiRouteResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHttpApiRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHttpApiRouteResponseBody
	GetRequestId() *string
}

type UpdateHttpApiRouteResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBEEB8C1-108E-50F0-9BEA-DED79553C309
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateHttpApiRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHttpApiRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHttpApiRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpApiRouteResponseBody) SetCode(v string) *UpdateHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpApiRouteResponseBody) SetMessage(v string) *UpdateHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpApiRouteResponseBody) SetRequestId(v string) *UpdateHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpApiRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
