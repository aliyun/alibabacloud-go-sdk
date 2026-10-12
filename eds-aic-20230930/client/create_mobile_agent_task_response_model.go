// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMobileAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMobileAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMobileAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMobileAgentTaskResponseBody) *CreateMobileAgentTaskResponse
	GetBody() *CreateMobileAgentTaskResponseBody
}

type CreateMobileAgentTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMobileAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMobileAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMobileAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMobileAgentTaskResponse) GetBody() *CreateMobileAgentTaskResponseBody {
	return s.Body
}

func (s *CreateMobileAgentTaskResponse) SetHeaders(v map[string]*string) *CreateMobileAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMobileAgentTaskResponse) SetStatusCode(v int32) *CreateMobileAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMobileAgentTaskResponse) SetBody(v *CreateMobileAgentTaskResponseBody) *CreateMobileAgentTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMobileAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
