// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataAgentMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataAgentMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataAgentMemoryResponse
	GetStatusCode() *int32
	SetBody(v *AddDataAgentMemoryResponseBody) *AddDataAgentMemoryResponse
	GetBody() *AddDataAgentMemoryResponseBody
}

type AddDataAgentMemoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataAgentMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataAgentMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataAgentMemoryResponse) GoString() string {
	return s.String()
}

func (s *AddDataAgentMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataAgentMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataAgentMemoryResponse) GetBody() *AddDataAgentMemoryResponseBody {
	return s.Body
}

func (s *AddDataAgentMemoryResponse) SetHeaders(v map[string]*string) *AddDataAgentMemoryResponse {
	s.Headers = v
	return s
}

func (s *AddDataAgentMemoryResponse) SetStatusCode(v int32) *AddDataAgentMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataAgentMemoryResponse) SetBody(v *AddDataAgentMemoryResponseBody) *AddDataAgentMemoryResponse {
	s.Body = v
	return s
}

func (s *AddDataAgentMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
