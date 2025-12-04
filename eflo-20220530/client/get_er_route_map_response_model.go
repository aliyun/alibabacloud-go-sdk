// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *GetErRouteMapResponseBody) *GetErRouteMapResponse
	GetBody() *GetErRouteMapResponseBody
}

type GetErRouteMapResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErRouteMapResponse) GetBody() *GetErRouteMapResponseBody {
	return s.Body
}

func (s *GetErRouteMapResponse) SetHeaders(v map[string]*string) *GetErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *GetErRouteMapResponse) SetStatusCode(v int32) *GetErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErRouteMapResponse) SetBody(v *GetErRouteMapResponseBody) *GetErRouteMapResponse {
	s.Body = v
	return s
}

func (s *GetErRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
