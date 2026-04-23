// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIAgentConcurrencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIAgentConcurrencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIAgentConcurrencyResponse
	GetStatusCode() *int32
	SetBody(v *GetAIAgentConcurrencyResponseBody) *GetAIAgentConcurrencyResponse
	GetBody() *GetAIAgentConcurrencyResponseBody
}

type GetAIAgentConcurrencyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIAgentConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIAgentConcurrencyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIAgentConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *GetAIAgentConcurrencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIAgentConcurrencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIAgentConcurrencyResponse) GetBody() *GetAIAgentConcurrencyResponseBody {
	return s.Body
}

func (s *GetAIAgentConcurrencyResponse) SetHeaders(v map[string]*string) *GetAIAgentConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *GetAIAgentConcurrencyResponse) SetStatusCode(v int32) *GetAIAgentConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIAgentConcurrencyResponse) SetBody(v *GetAIAgentConcurrencyResponseBody) *GetAIAgentConcurrencyResponse {
	s.Body = v
	return s
}

func (s *GetAIAgentConcurrencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
