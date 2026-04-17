// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelAgentTaskResponseBody) *CancelAgentTaskResponse
	GetBody() *CancelAgentTaskResponseBody
}

type CancelAgentTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAgentTaskResponse) GetBody() *CancelAgentTaskResponseBody {
	return s.Body
}

func (s *CancelAgentTaskResponse) SetHeaders(v map[string]*string) *CancelAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelAgentTaskResponse) SetStatusCode(v int32) *CancelAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAgentTaskResponse) SetBody(v *CancelAgentTaskResponseBody) *CancelAgentTaskResponse {
	s.Body = v
	return s
}

func (s *CancelAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
