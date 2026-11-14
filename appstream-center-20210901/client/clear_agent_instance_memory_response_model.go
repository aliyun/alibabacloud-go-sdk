// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAgentInstanceMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearAgentInstanceMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearAgentInstanceMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ClearAgentInstanceMemoryResponseBody) *ClearAgentInstanceMemoryResponse
	GetBody() *ClearAgentInstanceMemoryResponseBody
}

type ClearAgentInstanceMemoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearAgentInstanceMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearAgentInstanceMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearAgentInstanceMemoryResponse) GoString() string {
	return s.String()
}

func (s *ClearAgentInstanceMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearAgentInstanceMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearAgentInstanceMemoryResponse) GetBody() *ClearAgentInstanceMemoryResponseBody {
	return s.Body
}

func (s *ClearAgentInstanceMemoryResponse) SetHeaders(v map[string]*string) *ClearAgentInstanceMemoryResponse {
	s.Headers = v
	return s
}

func (s *ClearAgentInstanceMemoryResponse) SetStatusCode(v int32) *ClearAgentInstanceMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearAgentInstanceMemoryResponse) SetBody(v *ClearAgentInstanceMemoryResponseBody) *ClearAgentInstanceMemoryResponse {
	s.Body = v
	return s
}

func (s *ClearAgentInstanceMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
