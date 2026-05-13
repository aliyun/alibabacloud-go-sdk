// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceActionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WorkspaceActionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WorkspaceActionStatusResponse
	GetStatusCode() *int32
	SetBody(v *WorkspaceActionStatusResponseBody) *WorkspaceActionStatusResponse
	GetBody() *WorkspaceActionStatusResponseBody
}

type WorkspaceActionStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WorkspaceActionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WorkspaceActionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionStatusResponse) GoString() string {
	return s.String()
}

func (s *WorkspaceActionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WorkspaceActionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WorkspaceActionStatusResponse) GetBody() *WorkspaceActionStatusResponseBody {
	return s.Body
}

func (s *WorkspaceActionStatusResponse) SetHeaders(v map[string]*string) *WorkspaceActionStatusResponse {
	s.Headers = v
	return s
}

func (s *WorkspaceActionStatusResponse) SetStatusCode(v int32) *WorkspaceActionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *WorkspaceActionStatusResponse) SetBody(v *WorkspaceActionStatusResponseBody) *WorkspaceActionStatusResponse {
	s.Body = v
	return s
}

func (s *WorkspaceActionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
