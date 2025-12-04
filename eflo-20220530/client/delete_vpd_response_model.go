// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpdResponseBody) *DeleteVpdResponse
	GetBody() *DeleteVpdResponseBody
}

type DeleteVpdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpdResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpdResponse) GetBody() *DeleteVpdResponseBody {
	return s.Body
}

func (s *DeleteVpdResponse) SetHeaders(v map[string]*string) *DeleteVpdResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpdResponse) SetStatusCode(v int32) *DeleteVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpdResponse) SetBody(v *DeleteVpdResponseBody) *DeleteVpdResponse {
	s.Body = v
	return s
}

func (s *DeleteVpdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
