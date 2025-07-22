// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteTableResponseBody) *CreateRouteTableResponse
	GetBody() *CreateRouteTableResponseBody
}

type CreateRouteTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTableResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteTableResponse) GetBody() *CreateRouteTableResponseBody {
	return s.Body
}

func (s *CreateRouteTableResponse) SetHeaders(v map[string]*string) *CreateRouteTableResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteTableResponse) SetStatusCode(v int32) *CreateRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteTableResponse) SetBody(v *CreateRouteTableResponseBody) *CreateRouteTableResponse {
	s.Body = v
	return s
}

func (s *CreateRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
