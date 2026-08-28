// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExternalAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExternalAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExternalAgentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExternalAgentResponseBody) *DeleteExternalAgentResponse
	GetBody() *DeleteExternalAgentResponseBody
}

type DeleteExternalAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExternalAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteExternalAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExternalAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExternalAgentResponse) GetBody() *DeleteExternalAgentResponseBody {
	return s.Body
}

func (s *DeleteExternalAgentResponse) SetHeaders(v map[string]*string) *DeleteExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteExternalAgentResponse) SetStatusCode(v int32) *DeleteExternalAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExternalAgentResponse) SetBody(v *DeleteExternalAgentResponseBody) *DeleteExternalAgentResponse {
	s.Body = v
	return s
}

func (s *DeleteExternalAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
