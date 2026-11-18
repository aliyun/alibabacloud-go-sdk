// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSessionResponseBody) *GetAgentSessionResponse
	GetBody() *GetAgentSessionResponseBody
}

type GetAgentSessionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSessionResponse) GetBody() *GetAgentSessionResponseBody {
	return s.Body
}

func (s *GetAgentSessionResponse) SetHeaders(v map[string]*string) *GetAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSessionResponse) SetStatusCode(v int32) *GetAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSessionResponse) SetBody(v *GetAgentSessionResponseBody) *GetAgentSessionResponse {
	s.Body = v
	return s
}

func (s *GetAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
