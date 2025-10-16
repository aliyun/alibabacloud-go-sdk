// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteTableResponseBody) *DeleteRouteTableResponse
	GetBody() *DeleteRouteTableResponseBody
}

type DeleteRouteTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteTableResponse) GetBody() *DeleteRouteTableResponseBody {
	return s.Body
}

func (s *DeleteRouteTableResponse) SetHeaders(v map[string]*string) *DeleteRouteTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteTableResponse) SetStatusCode(v int32) *DeleteRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteTableResponse) SetBody(v *DeleteRouteTableResponseBody) *DeleteRouteTableResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
