// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateEdgeMobileAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateEdgeMobileAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateEdgeMobileAgentResponse
	GetStatusCode() *int32
	SetBody(v *ActivateEdgeMobileAgentResponseBody) *ActivateEdgeMobileAgentResponse
	GetBody() *ActivateEdgeMobileAgentResponseBody
}

type ActivateEdgeMobileAgentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateEdgeMobileAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateEdgeMobileAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateEdgeMobileAgentResponse) GoString() string {
	return s.String()
}

func (s *ActivateEdgeMobileAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateEdgeMobileAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateEdgeMobileAgentResponse) GetBody() *ActivateEdgeMobileAgentResponseBody {
	return s.Body
}

func (s *ActivateEdgeMobileAgentResponse) SetHeaders(v map[string]*string) *ActivateEdgeMobileAgentResponse {
	s.Headers = v
	return s
}

func (s *ActivateEdgeMobileAgentResponse) SetStatusCode(v int32) *ActivateEdgeMobileAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateEdgeMobileAgentResponse) SetBody(v *ActivateEdgeMobileAgentResponseBody) *ActivateEdgeMobileAgentResponse {
	s.Body = v
	return s
}

func (s *ActivateEdgeMobileAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
