// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomAgentResponseBody) *DeleteCustomAgentResponse
	GetBody() *DeleteCustomAgentResponseBody
}

type DeleteCustomAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomAgentResponse) GetBody() *DeleteCustomAgentResponseBody {
	return s.Body
}

func (s *DeleteCustomAgentResponse) SetHeaders(v map[string]*string) *DeleteCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomAgentResponse) SetStatusCode(v int32) *DeleteCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomAgentResponse) SetBody(v *DeleteCustomAgentResponseBody) *DeleteCustomAgentResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
