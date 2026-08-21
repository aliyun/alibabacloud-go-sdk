// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceQueueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceQueueResponseBody) *DeleteWorkspaceQueueResponse
	GetBody() *DeleteWorkspaceQueueResponseBody
}

type DeleteWorkspaceQueueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceQueueResponse) GetBody() *DeleteWorkspaceQueueResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceQueueResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceQueueResponse) SetStatusCode(v int32) *DeleteWorkspaceQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceQueueResponse) SetBody(v *DeleteWorkspaceQueueResponseBody) *DeleteWorkspaceQueueResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
