// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaWorkflowExecutionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaWorkflowExecutionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaWorkflowExecutionListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaWorkflowExecutionListResponseBody) *QueryMediaWorkflowExecutionListResponse
	GetBody() *QueryMediaWorkflowExecutionListResponseBody
}

type QueryMediaWorkflowExecutionListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaWorkflowExecutionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaWorkflowExecutionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaWorkflowExecutionListResponse) GetBody() *QueryMediaWorkflowExecutionListResponseBody {
	return s.Body
}

func (s *QueryMediaWorkflowExecutionListResponse) SetHeaders(v map[string]*string) *QueryMediaWorkflowExecutionListResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponse) SetStatusCode(v int32) *QueryMediaWorkflowExecutionListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponse) SetBody(v *QueryMediaWorkflowExecutionListResponseBody) *QueryMediaWorkflowExecutionListResponse {
	s.Body = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
