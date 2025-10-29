// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHttpApiRouteResponseBody
	GetCode() *string
	SetData(v *HttpRoute) *GetHttpApiRouteResponseBody
	GetData() *HttpRoute
	SetMessage(v string) *GetHttpApiRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHttpApiRouteResponseBody
	GetRequestId() *string
}

type GetHttpApiRouteResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The route details.
	Data *HttpRoute `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetHttpApiRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpApiRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHttpApiRouteResponseBody) GetData() *HttpRoute {
	return s.Data
}

func (s *GetHttpApiRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHttpApiRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpApiRouteResponseBody) SetCode(v string) *GetHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpApiRouteResponseBody) SetData(v *HttpRoute) *GetHttpApiRouteResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpApiRouteResponseBody) SetMessage(v string) *GetHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpApiRouteResponseBody) SetRequestId(v string) *GetHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpApiRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
