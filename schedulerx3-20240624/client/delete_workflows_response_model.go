// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkflowsResponseBody) *DeleteWorkflowsResponse
	GetBody() *DeleteWorkflowsResponseBody
}

type DeleteWorkflowsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkflowsResponse) GetBody() *DeleteWorkflowsResponseBody {
	return s.Body
}

func (s *DeleteWorkflowsResponse) SetHeaders(v map[string]*string) *DeleteWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowsResponse) SetStatusCode(v int32) *DeleteWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowsResponse) SetBody(v *DeleteWorkflowsResponseBody) *DeleteWorkflowsResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkflowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
