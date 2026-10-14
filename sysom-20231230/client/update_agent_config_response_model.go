// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentConfigResponseBody) *UpdateAgentConfigResponse
	GetBody() *UpdateAgentConfigResponseBody
}

type UpdateAgentConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentConfigResponse) GetBody() *UpdateAgentConfigResponseBody {
	return s.Body
}

func (s *UpdateAgentConfigResponse) SetHeaders(v map[string]*string) *UpdateAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentConfigResponse) SetStatusCode(v int32) *UpdateAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentConfigResponse) SetBody(v *UpdateAgentConfigResponseBody) *UpdateAgentConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
