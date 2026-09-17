// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkspaceQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkspaceQueueResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkspaceQueueResponseBody) *CreateWorkspaceQueueResponse
	GetBody() *CreateWorkspaceQueueResponseBody
}

type CreateWorkspaceQueueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkspaceQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkspaceQueueResponse) GetBody() *CreateWorkspaceQueueResponseBody {
	return s.Body
}

func (s *CreateWorkspaceQueueResponse) SetHeaders(v map[string]*string) *CreateWorkspaceQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceQueueResponse) SetStatusCode(v int32) *CreateWorkspaceQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceQueueResponse) SetBody(v *CreateWorkspaceQueueResponseBody) *CreateWorkspaceQueueResponse {
	s.Body = v
	return s
}

func (s *CreateWorkspaceQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
