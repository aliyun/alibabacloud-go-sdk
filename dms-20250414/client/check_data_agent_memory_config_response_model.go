// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataAgentMemoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDataAgentMemoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDataAgentMemoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *CheckDataAgentMemoryConfigResponseBody) *CheckDataAgentMemoryConfigResponse
	GetBody() *CheckDataAgentMemoryConfigResponseBody
}

type CheckDataAgentMemoryConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataAgentMemoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataAgentMemoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDataAgentMemoryConfigResponse) GoString() string {
	return s.String()
}

func (s *CheckDataAgentMemoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDataAgentMemoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDataAgentMemoryConfigResponse) GetBody() *CheckDataAgentMemoryConfigResponseBody {
	return s.Body
}

func (s *CheckDataAgentMemoryConfigResponse) SetHeaders(v map[string]*string) *CheckDataAgentMemoryConfigResponse {
	s.Headers = v
	return s
}

func (s *CheckDataAgentMemoryConfigResponse) SetStatusCode(v int32) *CheckDataAgentMemoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataAgentMemoryConfigResponse) SetBody(v *CheckDataAgentMemoryConfigResponseBody) *CheckDataAgentMemoryConfigResponse {
	s.Body = v
	return s
}

func (s *CheckDataAgentMemoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
