// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowDagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowDagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowDagResponseBody) *UpdateWorkflowDagResponse
	GetBody() *UpdateWorkflowDagResponseBody
}

type UpdateWorkflowDagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDagResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowDagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowDagResponse) GetBody() *UpdateWorkflowDagResponseBody {
	return s.Body
}

func (s *UpdateWorkflowDagResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDagResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDagResponse) SetStatusCode(v int32) *UpdateWorkflowDagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDagResponse) SetBody(v *UpdateWorkflowDagResponseBody) *UpdateWorkflowDagResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowDagResponse) Validate() error {
	return dara.Validate(s)
}
