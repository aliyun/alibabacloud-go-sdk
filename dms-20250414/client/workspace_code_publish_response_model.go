// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceCodePublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WorkspaceCodePublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WorkspaceCodePublishResponse
	GetStatusCode() *int32
	SetBody(v *WorkspaceCodePublishResponseBody) *WorkspaceCodePublishResponse
	GetBody() *WorkspaceCodePublishResponseBody
}

type WorkspaceCodePublishResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WorkspaceCodePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WorkspaceCodePublishResponse) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceCodePublishResponse) GoString() string {
	return s.String()
}

func (s *WorkspaceCodePublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WorkspaceCodePublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WorkspaceCodePublishResponse) GetBody() *WorkspaceCodePublishResponseBody {
	return s.Body
}

func (s *WorkspaceCodePublishResponse) SetHeaders(v map[string]*string) *WorkspaceCodePublishResponse {
	s.Headers = v
	return s
}

func (s *WorkspaceCodePublishResponse) SetStatusCode(v int32) *WorkspaceCodePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *WorkspaceCodePublishResponse) SetBody(v *WorkspaceCodePublishResponseBody) *WorkspaceCodePublishResponse {
	s.Body = v
	return s
}

func (s *WorkspaceCodePublishResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
