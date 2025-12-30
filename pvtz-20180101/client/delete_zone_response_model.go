// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteZoneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteZoneResponseBody) *DeleteZoneResponse
	GetBody() *DeleteZoneResponseBody
}

type DeleteZoneResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteZoneResponse) GetBody() *DeleteZoneResponseBody {
	return s.Body
}

func (s *DeleteZoneResponse) SetHeaders(v map[string]*string) *DeleteZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteZoneResponse) SetStatusCode(v int32) *DeleteZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteZoneResponse) SetBody(v *DeleteZoneResponseBody) *DeleteZoneResponse {
	s.Body = v
	return s
}

func (s *DeleteZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
