// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitWorkflowJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitWorkflowJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitWorkflowJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitWorkflowJobResponseBody) *SubmitWorkflowJobResponse
	GetBody() *SubmitWorkflowJobResponseBody
}

type SubmitWorkflowJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitWorkflowJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitWorkflowJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitWorkflowJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitWorkflowJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitWorkflowJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitWorkflowJobResponse) GetBody() *SubmitWorkflowJobResponseBody {
	return s.Body
}

func (s *SubmitWorkflowJobResponse) SetHeaders(v map[string]*string) *SubmitWorkflowJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitWorkflowJobResponse) SetStatusCode(v int32) *SubmitWorkflowJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitWorkflowJobResponse) SetBody(v *SubmitWorkflowJobResponseBody) *SubmitWorkflowJobResponse {
	s.Body = v
	return s
}

func (s *SubmitWorkflowJobResponse) Validate() error {
	return dara.Validate(s)
}
