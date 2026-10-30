// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteMapResponseBody) *CreateRouteMapResponse
	GetBody() *CreateRouteMapResponseBody
}

type CreateRouteMapResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteMapResponse) GetBody() *CreateRouteMapResponseBody {
	return s.Body
}

func (s *CreateRouteMapResponse) SetHeaders(v map[string]*string) *CreateRouteMapResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteMapResponse) SetStatusCode(v int32) *CreateRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteMapResponse) SetBody(v *CreateRouteMapResponseBody) *CreateRouteMapResponse {
	s.Body = v
	return s
}

func (s *CreateRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
