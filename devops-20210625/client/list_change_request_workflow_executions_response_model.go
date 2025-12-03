// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestWorkflowExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChangeRequestWorkflowExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChangeRequestWorkflowExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListChangeRequestWorkflowExecutionsResponseBody) *ListChangeRequestWorkflowExecutionsResponse
	GetBody() *ListChangeRequestWorkflowExecutionsResponseBody
}

type ListChangeRequestWorkflowExecutionsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChangeRequestWorkflowExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChangeRequestWorkflowExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestWorkflowExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListChangeRequestWorkflowExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChangeRequestWorkflowExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChangeRequestWorkflowExecutionsResponse) GetBody() *ListChangeRequestWorkflowExecutionsResponseBody {
	return s.Body
}

func (s *ListChangeRequestWorkflowExecutionsResponse) SetHeaders(v map[string]*string) *ListChangeRequestWorkflowExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponse) SetStatusCode(v int32) *ListChangeRequestWorkflowExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponse) SetBody(v *ListChangeRequestWorkflowExecutionsResponseBody) *ListChangeRequestWorkflowExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
