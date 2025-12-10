// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentResponse
	GetStatusCode() *int32
}

type CreateAgentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentResponse) SetHeaders(v map[string]*string) *CreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentResponse) SetStatusCode(v int32) *CreateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentResponse) Validate() error {
	return dara.Validate(s)
}
