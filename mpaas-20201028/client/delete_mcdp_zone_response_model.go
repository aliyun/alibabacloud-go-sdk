// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcdpZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcdpZoneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcdpZoneResponseBody) *DeleteMcdpZoneResponse
	GetBody() *DeleteMcdpZoneResponseBody
}

type DeleteMcdpZoneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcdpZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcdpZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcdpZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcdpZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcdpZoneResponse) GetBody() *DeleteMcdpZoneResponseBody {
	return s.Body
}

func (s *DeleteMcdpZoneResponse) SetHeaders(v map[string]*string) *DeleteMcdpZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcdpZoneResponse) SetStatusCode(v int32) *DeleteMcdpZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcdpZoneResponse) SetBody(v *DeleteMcdpZoneResponseBody) *DeleteMcdpZoneResponse {
	s.Body = v
	return s
}

func (s *DeleteMcdpZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
