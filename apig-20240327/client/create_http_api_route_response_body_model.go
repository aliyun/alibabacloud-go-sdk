// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHttpApiRouteResponseBody
	GetCode() *string
	SetData(v *CreateHttpApiRouteResponseBodyData) *CreateHttpApiRouteResponseBody
	GetData() *CreateHttpApiRouteResponseBodyData
	SetMessage(v string) *CreateHttpApiRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHttpApiRouteResponseBody
	GetRequestId() *string
}

type CreateHttpApiRouteResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateHttpApiRouteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHttpApiRouteResponseBody) GetData() *CreateHttpApiRouteResponseBodyData {
	return s.Data
}

func (s *CreateHttpApiRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHttpApiRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpApiRouteResponseBody) SetCode(v string) *CreateHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiRouteResponseBody) SetData(v *CreateHttpApiRouteResponseBodyData) *CreateHttpApiRouteResponseBody {
	s.Data = v
	return s
}

func (s *CreateHttpApiRouteResponseBody) SetMessage(v string) *CreateHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiRouteResponseBody) SetRequestId(v string) *CreateHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpApiRouteResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateHttpApiRouteResponseBodyData struct {
	// The route ID.
	//
	// example:
	//
	// hr-cr82undlhtgrlej***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s CreateHttpApiRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteResponseBodyData) GetRouteId() *string {
	return s.RouteId
}

func (s *CreateHttpApiRouteResponseBodyData) SetRouteId(v string) *CreateHttpApiRouteResponseBodyData {
	s.RouteId = &v
	return s
}

func (s *CreateHttpApiRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
