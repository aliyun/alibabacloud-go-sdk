// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteServiceInCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteServiceInCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteServiceInCenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteServiceInCenResponseBody) *DeleteRouteServiceInCenResponse
	GetBody() *DeleteRouteServiceInCenResponseBody
}

type DeleteRouteServiceInCenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteServiceInCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteServiceInCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteServiceInCenResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteServiceInCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteServiceInCenResponse) GetBody() *DeleteRouteServiceInCenResponseBody {
	return s.Body
}

func (s *DeleteRouteServiceInCenResponse) SetHeaders(v map[string]*string) *DeleteRouteServiceInCenResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteServiceInCenResponse) SetStatusCode(v int32) *DeleteRouteServiceInCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteServiceInCenResponse) SetBody(v *DeleteRouteServiceInCenResponseBody) *DeleteRouteServiceInCenResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteServiceInCenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
