// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExternalAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExternalAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExternalAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExternalAgentResponseBody) *UpdateExternalAgentResponse
	GetBody() *UpdateExternalAgentResponseBody
}

type UpdateExternalAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExternalAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExternalAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExternalAgentResponse) GetBody() *UpdateExternalAgentResponseBody {
	return s.Body
}

func (s *UpdateExternalAgentResponse) SetHeaders(v map[string]*string) *UpdateExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExternalAgentResponse) SetStatusCode(v int32) *UpdateExternalAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExternalAgentResponse) SetBody(v *UpdateExternalAgentResponseBody) *UpdateExternalAgentResponse {
	s.Body = v
	return s
}

func (s *UpdateExternalAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
