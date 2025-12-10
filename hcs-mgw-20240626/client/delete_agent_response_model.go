// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentResponse
	GetStatusCode() *int32
}

type DeleteAgentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentResponse) SetHeaders(v map[string]*string) *DeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentResponse) SetStatusCode(v int32) *DeleteAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentResponse) Validate() error {
	return dara.Validate(s)
}
