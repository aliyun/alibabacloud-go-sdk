// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIngressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIngressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIngressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIngressResponseBody) *DeleteIngressResponse
	GetBody() *DeleteIngressResponseBody
}

type DeleteIngressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIngressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIngressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIngressResponse) GoString() string {
	return s.String()
}

func (s *DeleteIngressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIngressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIngressResponse) GetBody() *DeleteIngressResponseBody {
	return s.Body
}

func (s *DeleteIngressResponse) SetHeaders(v map[string]*string) *DeleteIngressResponse {
	s.Headers = v
	return s
}

func (s *DeleteIngressResponse) SetStatusCode(v int32) *DeleteIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIngressResponse) SetBody(v *DeleteIngressResponseBody) *DeleteIngressResponse {
	s.Body = v
	return s
}

func (s *DeleteIngressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
