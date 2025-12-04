// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateErRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateErRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *CreateErRouteMapResponseBody) *CreateErRouteMapResponse
	GetBody() *CreateErRouteMapResponseBody
}

type CreateErRouteMapResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateErRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateErRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateErRouteMapResponse) GetBody() *CreateErRouteMapResponseBody {
	return s.Body
}

func (s *CreateErRouteMapResponse) SetHeaders(v map[string]*string) *CreateErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *CreateErRouteMapResponse) SetStatusCode(v int32) *CreateErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErRouteMapResponse) SetBody(v *CreateErRouteMapResponseBody) *CreateErRouteMapResponse {
	s.Body = v
	return s
}

func (s *CreateErRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
