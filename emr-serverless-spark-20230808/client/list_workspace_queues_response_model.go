// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceQueuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspaceQueuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspaceQueuesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspaceQueuesResponseBody) *ListWorkspaceQueuesResponse
	GetBody() *ListWorkspaceQueuesResponseBody
}

type ListWorkspaceQueuesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceQueuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspaceQueuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspaceQueuesResponse) GetBody() *ListWorkspaceQueuesResponseBody {
	return s.Body
}

func (s *ListWorkspaceQueuesResponse) SetHeaders(v map[string]*string) *ListWorkspaceQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceQueuesResponse) SetStatusCode(v int32) *ListWorkspaceQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceQueuesResponse) SetBody(v *ListWorkspaceQueuesResponseBody) *ListWorkspaceQueuesResponse {
	s.Body = v
	return s
}

func (s *ListWorkspaceQueuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
