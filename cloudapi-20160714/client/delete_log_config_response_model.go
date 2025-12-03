// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogConfigResponseBody) *DeleteLogConfigResponse
	GetBody() *DeleteLogConfigResponseBody
}

type DeleteLogConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogConfigResponse) GetBody() *DeleteLogConfigResponseBody {
	return s.Body
}

func (s *DeleteLogConfigResponse) SetHeaders(v map[string]*string) *DeleteLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogConfigResponse) SetStatusCode(v int32) *DeleteLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogConfigResponse) SetBody(v *DeleteLogConfigResponseBody) *DeleteLogConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
