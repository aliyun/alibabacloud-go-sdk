// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppsResponseBody) *DeleteAppsResponse
	GetBody() *DeleteAppsResponseBody
}

type DeleteAppsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppsResponse) GetBody() *DeleteAppsResponseBody {
	return s.Body
}

func (s *DeleteAppsResponse) SetHeaders(v map[string]*string) *DeleteAppsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppsResponse) SetStatusCode(v int32) *DeleteAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppsResponse) SetBody(v *DeleteAppsResponseBody) *DeleteAppsResponse {
	s.Body = v
	return s
}

func (s *DeleteAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
