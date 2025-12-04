// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *GetErRouteEntryResponseBody) *GetErRouteEntryResponse
	GetBody() *GetErRouteEntryResponseBody
}

type GetErRouteEntryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErRouteEntryResponse) GetBody() *GetErRouteEntryResponseBody {
	return s.Body
}

func (s *GetErRouteEntryResponse) SetHeaders(v map[string]*string) *GetErRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *GetErRouteEntryResponse) SetStatusCode(v int32) *GetErRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErRouteEntryResponse) SetBody(v *GetErRouteEntryResponseBody) *GetErRouteEntryResponse {
	s.Body = v
	return s
}

func (s *GetErRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
