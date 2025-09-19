// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentlessTaskCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentlessTaskCountResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentlessTaskCountResponseBody) *GetAgentlessTaskCountResponse
	GetBody() *GetAgentlessTaskCountResponseBody
}

type GetAgentlessTaskCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentlessTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentlessTaskCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskCountResponse) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentlessTaskCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentlessTaskCountResponse) GetBody() *GetAgentlessTaskCountResponseBody {
	return s.Body
}

func (s *GetAgentlessTaskCountResponse) SetHeaders(v map[string]*string) *GetAgentlessTaskCountResponse {
	s.Headers = v
	return s
}

func (s *GetAgentlessTaskCountResponse) SetStatusCode(v int32) *GetAgentlessTaskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentlessTaskCountResponse) SetBody(v *GetAgentlessTaskCountResponseBody) *GetAgentlessTaskCountResponse {
	s.Body = v
	return s
}

func (s *GetAgentlessTaskCountResponse) Validate() error {
	return dara.Validate(s)
}
