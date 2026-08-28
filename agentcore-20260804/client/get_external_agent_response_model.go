// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExternalAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExternalAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetExternalAgentResponseBody) *GetExternalAgentResponse
	GetBody() *GetExternalAgentResponseBody
}

type GetExternalAgentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExternalAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExternalAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExternalAgentResponse) GetBody() *GetExternalAgentResponseBody {
	return s.Body
}

func (s *GetExternalAgentResponse) SetHeaders(v map[string]*string) *GetExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *GetExternalAgentResponse) SetStatusCode(v int32) *GetExternalAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalAgentResponse) SetBody(v *GetExternalAgentResponseBody) *GetExternalAgentResponse {
	s.Body = v
	return s
}

func (s *GetExternalAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
