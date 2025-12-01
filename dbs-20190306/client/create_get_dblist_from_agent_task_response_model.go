// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGetDBListFromAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGetDBListFromAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGetDBListFromAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateGetDBListFromAgentTaskResponseBody) *CreateGetDBListFromAgentTaskResponse
	GetBody() *CreateGetDBListFromAgentTaskResponseBody
}

type CreateGetDBListFromAgentTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGetDBListFromAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGetDBListFromAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGetDBListFromAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGetDBListFromAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGetDBListFromAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGetDBListFromAgentTaskResponse) GetBody() *CreateGetDBListFromAgentTaskResponseBody {
	return s.Body
}

func (s *CreateGetDBListFromAgentTaskResponse) SetHeaders(v map[string]*string) *CreateGetDBListFromAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponse) SetStatusCode(v int32) *CreateGetDBListFromAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponse) SetBody(v *CreateGetDBListFromAgentTaskResponseBody) *CreateGetDBListFromAgentTaskResponse {
	s.Body = v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
