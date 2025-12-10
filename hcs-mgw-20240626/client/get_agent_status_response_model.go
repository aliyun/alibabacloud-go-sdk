// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentStatusResponseBody) *GetAgentStatusResponse
	GetBody() *GetAgentStatusResponseBody
}

type GetAgentStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentStatusResponse) GetBody() *GetAgentStatusResponseBody {
	return s.Body
}

func (s *GetAgentStatusResponse) SetHeaders(v map[string]*string) *GetAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentStatusResponse) SetStatusCode(v int32) *GetAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentStatusResponse) SetBody(v *GetAgentStatusResponseBody) *GetAgentStatusResponse {
	s.Body = v
	return s
}

func (s *GetAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
