// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApSsidConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApSsidConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApSsidConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApSsidConfigResponseBody) *DeleteApSsidConfigResponse
	GetBody() *DeleteApSsidConfigResponseBody
}

type DeleteApSsidConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApSsidConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApSsidConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApSsidConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteApSsidConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApSsidConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApSsidConfigResponse) GetBody() *DeleteApSsidConfigResponseBody {
	return s.Body
}

func (s *DeleteApSsidConfigResponse) SetHeaders(v map[string]*string) *DeleteApSsidConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteApSsidConfigResponse) SetStatusCode(v int32) *DeleteApSsidConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApSsidConfigResponse) SetBody(v *DeleteApSsidConfigResponseBody) *DeleteApSsidConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteApSsidConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
