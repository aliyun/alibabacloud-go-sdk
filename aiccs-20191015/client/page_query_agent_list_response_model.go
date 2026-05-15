// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageQueryAgentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageQueryAgentListResponse
	GetStatusCode() *int32
	SetBody(v *PageQueryAgentListResponseBody) *PageQueryAgentListResponse
	GetBody() *PageQueryAgentListResponseBody
}

type PageQueryAgentListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageQueryAgentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageQueryAgentListResponse) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListResponse) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageQueryAgentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageQueryAgentListResponse) GetBody() *PageQueryAgentListResponseBody {
	return s.Body
}

func (s *PageQueryAgentListResponse) SetHeaders(v map[string]*string) *PageQueryAgentListResponse {
	s.Headers = v
	return s
}

func (s *PageQueryAgentListResponse) SetStatusCode(v int32) *PageQueryAgentListResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryAgentListResponse) SetBody(v *PageQueryAgentListResponseBody) *PageQueryAgentListResponse {
	s.Body = v
	return s
}

func (s *PageQueryAgentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
