// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentServiceStatusResponseBody) *GetAgentServiceStatusResponse
	GetBody() *GetAgentServiceStatusResponseBody
}

type GetAgentServiceStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentServiceStatusResponse) GetBody() *GetAgentServiceStatusResponseBody {
	return s.Body
}

func (s *GetAgentServiceStatusResponse) SetHeaders(v map[string]*string) *GetAgentServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentServiceStatusResponse) SetStatusCode(v int32) *GetAgentServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentServiceStatusResponse) SetBody(v *GetAgentServiceStatusResponseBody) *GetAgentServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetAgentServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
