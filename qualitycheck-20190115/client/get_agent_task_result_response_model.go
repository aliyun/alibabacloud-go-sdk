// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentTaskResultResponseBody) *GetAgentTaskResultResponse
	GetBody() *GetAgentTaskResultResponseBody
}

type GetAgentTaskResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentTaskResultResponse) GetBody() *GetAgentTaskResultResponseBody {
	return s.Body
}

func (s *GetAgentTaskResultResponse) SetHeaders(v map[string]*string) *GetAgentTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetAgentTaskResultResponse) SetStatusCode(v int32) *GetAgentTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentTaskResultResponse) SetBody(v *GetAgentTaskResultResponseBody) *GetAgentTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetAgentTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
