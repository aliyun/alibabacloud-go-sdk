// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAgentMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAgentMcpResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAgentMcpResponseBody) *GetDataAgentMcpResponse
	GetBody() *GetDataAgentMcpResponseBody
}

type GetDataAgentMcpResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAgentMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAgentMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentMcpResponse) GoString() string {
	return s.String()
}

func (s *GetDataAgentMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAgentMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAgentMcpResponse) GetBody() *GetDataAgentMcpResponseBody {
	return s.Body
}

func (s *GetDataAgentMcpResponse) SetHeaders(v map[string]*string) *GetDataAgentMcpResponse {
	s.Headers = v
	return s
}

func (s *GetDataAgentMcpResponse) SetStatusCode(v int32) *GetDataAgentMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAgentMcpResponse) SetBody(v *GetDataAgentMcpResponseBody) *GetDataAgentMcpResponse {
	s.Body = v
	return s
}

func (s *GetDataAgentMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
