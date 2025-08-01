// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentResponseBody) *GetAgentResponse
	GetBody() *GetAgentResponseBody
}

type GetAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponse) GoString() string {
	return s.String()
}

func (s *GetAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentResponse) GetBody() *GetAgentResponseBody {
	return s.Body
}

func (s *GetAgentResponse) SetHeaders(v map[string]*string) *GetAgentResponse {
	s.Headers = v
	return s
}

func (s *GetAgentResponse) SetStatusCode(v int32) *GetAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentResponse) SetBody(v *GetAgentResponseBody) *GetAgentResponse {
	s.Body = v
	return s
}

func (s *GetAgentResponse) Validate() error {
	return dara.Validate(s)
}
