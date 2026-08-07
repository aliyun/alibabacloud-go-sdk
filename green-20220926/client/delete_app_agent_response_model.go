// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppAgentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppAgentResponseBody) *DeleteAppAgentResponse
	GetBody() *DeleteAppAgentResponseBody
}

type DeleteAppAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppAgentResponse) GetBody() *DeleteAppAgentResponseBody {
	return s.Body
}

func (s *DeleteAppAgentResponse) SetHeaders(v map[string]*string) *DeleteAppAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppAgentResponse) SetStatusCode(v int32) *DeleteAppAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppAgentResponse) SetBody(v *DeleteAppAgentResponseBody) *DeleteAppAgentResponse {
	s.Body = v
	return s
}

func (s *DeleteAppAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
