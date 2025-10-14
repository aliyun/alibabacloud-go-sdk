// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecursionZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecursionZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecursionZoneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecursionZoneResponseBody) *DeleteRecursionZoneResponse
	GetBody() *DeleteRecursionZoneResponseBody
}

type DeleteRecursionZoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecursionZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecursionZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecursionZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecursionZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecursionZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecursionZoneResponse) GetBody() *DeleteRecursionZoneResponseBody {
	return s.Body
}

func (s *DeleteRecursionZoneResponse) SetHeaders(v map[string]*string) *DeleteRecursionZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecursionZoneResponse) SetStatusCode(v int32) *DeleteRecursionZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecursionZoneResponse) SetBody(v *DeleteRecursionZoneResponseBody) *DeleteRecursionZoneResponse {
	s.Body = v
	return s
}

func (s *DeleteRecursionZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
