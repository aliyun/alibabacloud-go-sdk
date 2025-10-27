// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentlessRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentlessRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentlessRegionResponseBody) *ListAgentlessRegionResponse
	GetBody() *ListAgentlessRegionResponseBody
}

type ListAgentlessRegionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentlessRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentlessRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRegionResponse) GoString() string {
	return s.String()
}

func (s *ListAgentlessRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentlessRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentlessRegionResponse) GetBody() *ListAgentlessRegionResponseBody {
	return s.Body
}

func (s *ListAgentlessRegionResponse) SetHeaders(v map[string]*string) *ListAgentlessRegionResponse {
	s.Headers = v
	return s
}

func (s *ListAgentlessRegionResponse) SetStatusCode(v int32) *ListAgentlessRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentlessRegionResponse) SetBody(v *ListAgentlessRegionResponseBody) *ListAgentlessRegionResponse {
	s.Body = v
	return s
}

func (s *ListAgentlessRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
