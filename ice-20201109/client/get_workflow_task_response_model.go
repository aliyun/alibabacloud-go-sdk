// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowTaskResponseBody) *GetWorkflowTaskResponse
	GetBody() *GetWorkflowTaskResponseBody
}

type GetWorkflowTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTaskResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowTaskResponse) GetBody() *GetWorkflowTaskResponseBody {
	return s.Body
}

func (s *GetWorkflowTaskResponse) SetHeaders(v map[string]*string) *GetWorkflowTaskResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowTaskResponse) SetStatusCode(v int32) *GetWorkflowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowTaskResponse) SetBody(v *GetWorkflowTaskResponseBody) *GetWorkflowTaskResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
