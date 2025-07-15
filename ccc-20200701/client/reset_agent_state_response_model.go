// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgentStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAgentStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAgentStateResponse
	GetStatusCode() *int32
	SetBody(v *ResetAgentStateResponseBody) *ResetAgentStateResponse
	GetBody() *ResetAgentStateResponseBody
}

type ResetAgentStateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAgentStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAgentStateResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAgentStateResponse) GoString() string {
	return s.String()
}

func (s *ResetAgentStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAgentStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAgentStateResponse) GetBody() *ResetAgentStateResponseBody {
	return s.Body
}

func (s *ResetAgentStateResponse) SetHeaders(v map[string]*string) *ResetAgentStateResponse {
	s.Headers = v
	return s
}

func (s *ResetAgentStateResponse) SetStatusCode(v int32) *ResetAgentStateResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAgentStateResponse) SetBody(v *ResetAgentStateResponseBody) *ResetAgentStateResponse {
	s.Body = v
	return s
}

func (s *ResetAgentStateResponse) Validate() error {
	return dara.Validate(s)
}
