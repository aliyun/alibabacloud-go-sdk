// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageQueryAgentListNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageQueryAgentListNewResponse
	GetStatusCode() *int32
	SetBody(v *PageQueryAgentListNewResponseBody) *PageQueryAgentListNewResponse
	GetBody() *PageQueryAgentListNewResponseBody
}

type PageQueryAgentListNewResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageQueryAgentListNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageQueryAgentListNewResponse) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListNewResponse) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageQueryAgentListNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageQueryAgentListNewResponse) GetBody() *PageQueryAgentListNewResponseBody {
	return s.Body
}

func (s *PageQueryAgentListNewResponse) SetHeaders(v map[string]*string) *PageQueryAgentListNewResponse {
	s.Headers = v
	return s
}

func (s *PageQueryAgentListNewResponse) SetStatusCode(v int32) *PageQueryAgentListNewResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryAgentListNewResponse) SetBody(v *PageQueryAgentListNewResponseBody) *PageQueryAgentListNewResponse {
	s.Body = v
	return s
}

func (s *PageQueryAgentListNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
