// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHostShareKeyResponseBody) *DeleteHostShareKeyResponse
	GetBody() *DeleteHostShareKeyResponseBody
}

type DeleteHostShareKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHostShareKeyResponse) GetBody() *DeleteHostShareKeyResponseBody {
	return s.Body
}

func (s *DeleteHostShareKeyResponse) SetHeaders(v map[string]*string) *DeleteHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostShareKeyResponse) SetStatusCode(v int32) *DeleteHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostShareKeyResponse) SetBody(v *DeleteHostShareKeyResponseBody) *DeleteHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteHostShareKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
