// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultiModalAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultiModalAgentResponse
	GetStatusCode() *int32
	SetBody(v *MultiModalAgentResponseBody) *MultiModalAgentResponse
	GetBody() *MultiModalAgentResponseBody
}

type MultiModalAgentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiModalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentResponse) GoString() string {
	return s.String()
}

func (s *MultiModalAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultiModalAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultiModalAgentResponse) GetBody() *MultiModalAgentResponseBody {
	return s.Body
}

func (s *MultiModalAgentResponse) SetHeaders(v map[string]*string) *MultiModalAgentResponse {
	s.Headers = v
	return s
}

func (s *MultiModalAgentResponse) SetStatusCode(v int32) *MultiModalAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiModalAgentResponse) SetBody(v *MultiModalAgentResponseBody) *MultiModalAgentResponse {
	s.Body = v
	return s
}

func (s *MultiModalAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
