// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceActionLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WorkspaceActionLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WorkspaceActionLogResponse
	GetStatusCode() *int32
	SetBody(v *WorkspaceActionLogResponseBody) *WorkspaceActionLogResponse
	GetBody() *WorkspaceActionLogResponseBody
}

type WorkspaceActionLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WorkspaceActionLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WorkspaceActionLogResponse) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionLogResponse) GoString() string {
	return s.String()
}

func (s *WorkspaceActionLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WorkspaceActionLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WorkspaceActionLogResponse) GetBody() *WorkspaceActionLogResponseBody {
	return s.Body
}

func (s *WorkspaceActionLogResponse) SetHeaders(v map[string]*string) *WorkspaceActionLogResponse {
	s.Headers = v
	return s
}

func (s *WorkspaceActionLogResponse) SetStatusCode(v int32) *WorkspaceActionLogResponse {
	s.StatusCode = &v
	return s
}

func (s *WorkspaceActionLogResponse) SetBody(v *WorkspaceActionLogResponseBody) *WorkspaceActionLogResponse {
	s.Body = v
	return s
}

func (s *WorkspaceActionLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
