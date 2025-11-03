// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentTaskResponseBody) *GetAgentTaskResponse
	GetBody() *GetAgentTaskResponseBody
}

type GetAgentTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentTaskResponse) GetBody() *GetAgentTaskResponseBody {
	return s.Body
}

func (s *GetAgentTaskResponse) SetHeaders(v map[string]*string) *GetAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAgentTaskResponse) SetStatusCode(v int32) *GetAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentTaskResponse) SetBody(v *GetAgentTaskResponseBody) *GetAgentTaskResponse {
	s.Body = v
	return s
}

func (s *GetAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
