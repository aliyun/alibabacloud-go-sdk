// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcpResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcpResponseBody) *DeleteMcpResponse
	GetBody() *DeleteMcpResponseBody
}

type DeleteMcpResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcpResponse) GetBody() *DeleteMcpResponseBody {
	return s.Body
}

func (s *DeleteMcpResponse) SetHeaders(v map[string]*string) *DeleteMcpResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpResponse) SetStatusCode(v int32) *DeleteMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpResponse) SetBody(v *DeleteMcpResponseBody) *DeleteMcpResponse {
	s.Body = v
	return s
}

func (s *DeleteMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
