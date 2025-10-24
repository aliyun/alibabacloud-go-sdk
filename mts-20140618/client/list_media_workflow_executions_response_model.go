// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaWorkflowExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaWorkflowExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaWorkflowExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaWorkflowExecutionsResponseBody) *ListMediaWorkflowExecutionsResponse
	GetBody() *ListMediaWorkflowExecutionsResponseBody
}

type ListMediaWorkflowExecutionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaWorkflowExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaWorkflowExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaWorkflowExecutionsResponse) GetBody() *ListMediaWorkflowExecutionsResponseBody {
	return s.Body
}

func (s *ListMediaWorkflowExecutionsResponse) SetHeaders(v map[string]*string) *ListMediaWorkflowExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponse) SetStatusCode(v int32) *ListMediaWorkflowExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponse) SetBody(v *ListMediaWorkflowExecutionsResponseBody) *ListMediaWorkflowExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
