// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentMcpResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentMcpResponseBody) *DeleteDataAgentMcpResponse
	GetBody() *DeleteDataAgentMcpResponseBody
}

type DeleteDataAgentMcpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMcpResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentMcpResponse) GetBody() *DeleteDataAgentMcpResponseBody {
	return s.Body
}

func (s *DeleteDataAgentMcpResponse) SetHeaders(v map[string]*string) *DeleteDataAgentMcpResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentMcpResponse) SetStatusCode(v int32) *DeleteDataAgentMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentMcpResponse) SetBody(v *DeleteDataAgentMcpResponseBody) *DeleteDataAgentMcpResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
