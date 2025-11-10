// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiMcpServerResponseBody) *DeleteApiMcpServerResponse
	GetBody() *DeleteApiMcpServerResponseBody
}

type DeleteApiMcpServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiMcpServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiMcpServerResponse) GetBody() *DeleteApiMcpServerResponseBody {
	return s.Body
}

func (s *DeleteApiMcpServerResponse) SetHeaders(v map[string]*string) *DeleteApiMcpServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiMcpServerResponse) SetStatusCode(v int32) *DeleteApiMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiMcpServerResponse) SetBody(v *DeleteApiMcpServerResponseBody) *DeleteApiMcpServerResponse {
	s.Body = v
	return s
}

func (s *DeleteApiMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
