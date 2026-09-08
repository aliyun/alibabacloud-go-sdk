// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentJobResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentJobResponseBody) *GetAgentJobResponse
	GetBody() *GetAgentJobResponseBody
}

type GetAgentJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentJobResponse) GoString() string {
	return s.String()
}

func (s *GetAgentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentJobResponse) GetBody() *GetAgentJobResponseBody {
	return s.Body
}

func (s *GetAgentJobResponse) SetHeaders(v map[string]*string) *GetAgentJobResponse {
	s.Headers = v
	return s
}

func (s *GetAgentJobResponse) SetStatusCode(v int32) *GetAgentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentJobResponse) SetBody(v *GetAgentJobResponseBody) *GetAgentJobResponse {
	s.Body = v
	return s
}

func (s *GetAgentJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
