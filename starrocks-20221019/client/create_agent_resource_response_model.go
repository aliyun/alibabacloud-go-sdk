// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentResourceResponseBody) *CreateAgentResourceResponse
	GetBody() *CreateAgentResourceResponseBody
}

type CreateAgentResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentResourceResponse) GetBody() *CreateAgentResourceResponseBody {
	return s.Body
}

func (s *CreateAgentResourceResponse) SetHeaders(v map[string]*string) *CreateAgentResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentResourceResponse) SetStatusCode(v int32) *CreateAgentResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentResourceResponse) SetBody(v *CreateAgentResourceResponseBody) *CreateAgentResourceResponse {
	s.Body = v
	return s
}

func (s *CreateAgentResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
