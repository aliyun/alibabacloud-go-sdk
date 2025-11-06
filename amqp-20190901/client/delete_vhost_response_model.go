// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVhostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVhostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVhostResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVhostResponseBody) *DeleteVhostResponse
	GetBody() *DeleteVhostResponseBody
}

type DeleteVhostResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVhostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVhostResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVhostResponse) GoString() string {
	return s.String()
}

func (s *DeleteVhostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVhostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVhostResponse) GetBody() *DeleteVhostResponseBody {
	return s.Body
}

func (s *DeleteVhostResponse) SetHeaders(v map[string]*string) *DeleteVhostResponse {
	s.Headers = v
	return s
}

func (s *DeleteVhostResponse) SetStatusCode(v int32) *DeleteVhostResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVhostResponse) SetBody(v *DeleteVhostResponseBody) *DeleteVhostResponse {
	s.Body = v
	return s
}

func (s *DeleteVhostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
