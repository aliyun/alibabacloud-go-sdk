// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateRouteTableResponseBody) *UnassociateRouteTableResponse
	GetBody() *UnassociateRouteTableResponseBody
}

type UnassociateRouteTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateRouteTableResponse) GoString() string {
	return s.String()
}

func (s *UnassociateRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateRouteTableResponse) GetBody() *UnassociateRouteTableResponseBody {
	return s.Body
}

func (s *UnassociateRouteTableResponse) SetHeaders(v map[string]*string) *UnassociateRouteTableResponse {
	s.Headers = v
	return s
}

func (s *UnassociateRouteTableResponse) SetStatusCode(v int32) *UnassociateRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateRouteTableResponse) SetBody(v *UnassociateRouteTableResponseBody) *UnassociateRouteTableResponse {
	s.Body = v
	return s
}

func (s *UnassociateRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
