// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceAgentConfigResponseBody) *UpdateInstanceAgentConfigResponse
	GetBody() *UpdateInstanceAgentConfigResponseBody
}

type UpdateInstanceAgentConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceAgentConfigResponse) GetBody() *UpdateInstanceAgentConfigResponseBody {
	return s.Body
}

func (s *UpdateInstanceAgentConfigResponse) SetHeaders(v map[string]*string) *UpdateInstanceAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAgentConfigResponse) SetStatusCode(v int32) *UpdateInstanceAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAgentConfigResponse) SetBody(v *UpdateInstanceAgentConfigResponseBody) *UpdateInstanceAgentConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
