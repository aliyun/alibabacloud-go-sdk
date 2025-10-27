// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentlessScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentlessScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentlessScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentlessScanTaskResponseBody) *CreateAgentlessScanTaskResponse
	GetBody() *CreateAgentlessScanTaskResponseBody
}

type CreateAgentlessScanTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentlessScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentlessScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentlessScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentlessScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentlessScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentlessScanTaskResponse) GetBody() *CreateAgentlessScanTaskResponseBody {
	return s.Body
}

func (s *CreateAgentlessScanTaskResponse) SetHeaders(v map[string]*string) *CreateAgentlessScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentlessScanTaskResponse) SetStatusCode(v int32) *CreateAgentlessScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentlessScanTaskResponse) SetBody(v *CreateAgentlessScanTaskResponseBody) *CreateAgentlessScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAgentlessScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
