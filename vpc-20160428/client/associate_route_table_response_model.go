// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *AssociateRouteTableResponseBody) *AssociateRouteTableResponse
	GetBody() *AssociateRouteTableResponseBody
}

type AssociateRouteTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableResponse) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateRouteTableResponse) GetBody() *AssociateRouteTableResponseBody {
	return s.Body
}

func (s *AssociateRouteTableResponse) SetHeaders(v map[string]*string) *AssociateRouteTableResponse {
	s.Headers = v
	return s
}

func (s *AssociateRouteTableResponse) SetStatusCode(v int32) *AssociateRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateRouteTableResponse) SetBody(v *AssociateRouteTableResponseBody) *AssociateRouteTableResponse {
	s.Body = v
	return s
}

func (s *AssociateRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
