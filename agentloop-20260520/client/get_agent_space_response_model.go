// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSpaceResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSpaceResponseBody) *GetAgentSpaceResponse
	GetBody() *GetAgentSpaceResponseBody
}

type GetAgentSpaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSpaceResponse) GetBody() *GetAgentSpaceResponseBody {
	return s.Body
}

func (s *GetAgentSpaceResponse) SetHeaders(v map[string]*string) *GetAgentSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSpaceResponse) SetStatusCode(v int32) *GetAgentSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSpaceResponse) SetBody(v *GetAgentSpaceResponseBody) *GetAgentSpaceResponse {
	s.Body = v
	return s
}

func (s *GetAgentSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
