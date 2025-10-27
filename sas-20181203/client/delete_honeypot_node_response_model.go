// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHoneypotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHoneypotNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHoneypotNodeResponseBody) *DeleteHoneypotNodeResponse
	GetBody() *DeleteHoneypotNodeResponseBody
}

type DeleteHoneypotNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoneypotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoneypotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHoneypotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHoneypotNodeResponse) GetBody() *DeleteHoneypotNodeResponseBody {
	return s.Body
}

func (s *DeleteHoneypotNodeResponse) SetHeaders(v map[string]*string) *DeleteHoneypotNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoneypotNodeResponse) SetStatusCode(v int32) *DeleteHoneypotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoneypotNodeResponse) SetBody(v *DeleteHoneypotNodeResponseBody) *DeleteHoneypotNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteHoneypotNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
