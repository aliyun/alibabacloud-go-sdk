// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentConfigResponseBody) *CreateAgentConfigResponse
	GetBody() *CreateAgentConfigResponseBody
}

type CreateAgentConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentConfigResponse) GetBody() *CreateAgentConfigResponseBody {
	return s.Body
}

func (s *CreateAgentConfigResponse) SetHeaders(v map[string]*string) *CreateAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentConfigResponse) SetStatusCode(v int32) *CreateAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentConfigResponse) SetBody(v *CreateAgentConfigResponseBody) *CreateAgentConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
