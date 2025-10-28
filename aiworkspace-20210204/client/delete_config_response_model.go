// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigResponseBody) *DeleteConfigResponse
	GetBody() *DeleteConfigResponseBody
}

type DeleteConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigResponse) GetBody() *DeleteConfigResponseBody {
	return s.Body
}

func (s *DeleteConfigResponse) SetHeaders(v map[string]*string) *DeleteConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigResponse) SetStatusCode(v int32) *DeleteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigResponse) SetBody(v *DeleteConfigResponseBody) *DeleteConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
