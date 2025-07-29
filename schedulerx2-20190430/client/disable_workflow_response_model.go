// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *DisableWorkflowResponseBody) *DisableWorkflowResponse
	GetBody() *DisableWorkflowResponseBody
}

type DisableWorkflowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DisableWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableWorkflowResponse) GetBody() *DisableWorkflowResponseBody {
	return s.Body
}

func (s *DisableWorkflowResponse) SetHeaders(v map[string]*string) *DisableWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DisableWorkflowResponse) SetStatusCode(v int32) *DisableWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableWorkflowResponse) SetBody(v *DisableWorkflowResponseBody) *DisableWorkflowResponse {
	s.Body = v
	return s
}

func (s *DisableWorkflowResponse) Validate() error {
	return dara.Validate(s)
}
