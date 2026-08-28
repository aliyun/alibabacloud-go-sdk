// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExternalAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExternalAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreateExternalAgentResponseBody) *CreateExternalAgentResponse
	GetBody() *CreateExternalAgentResponseBody
}

type CreateExternalAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExternalAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExternalAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExternalAgentResponse) GetBody() *CreateExternalAgentResponseBody {
	return s.Body
}

func (s *CreateExternalAgentResponse) SetHeaders(v map[string]*string) *CreateExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateExternalAgentResponse) SetStatusCode(v int32) *CreateExternalAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExternalAgentResponse) SetBody(v *CreateExternalAgentResponseBody) *CreateExternalAgentResponse {
	s.Body = v
	return s
}

func (s *CreateExternalAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
