// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataAgentMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataAgentMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataAgentMcpResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataAgentMcpResponseBody) *ModifyDataAgentMcpResponse
	GetBody() *ModifyDataAgentMcpResponseBody
}

type ModifyDataAgentMcpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataAgentMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataAgentMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataAgentMcpResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataAgentMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataAgentMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataAgentMcpResponse) GetBody() *ModifyDataAgentMcpResponseBody {
	return s.Body
}

func (s *ModifyDataAgentMcpResponse) SetHeaders(v map[string]*string) *ModifyDataAgentMcpResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataAgentMcpResponse) SetStatusCode(v int32) *ModifyDataAgentMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataAgentMcpResponse) SetBody(v *ModifyDataAgentMcpResponseBody) *ModifyDataAgentMcpResponse {
	s.Body = v
	return s
}

func (s *ModifyDataAgentMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
