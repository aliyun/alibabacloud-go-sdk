// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkflowResponseBody) *AddWorkflowResponse
	GetBody() *AddWorkflowResponseBody
}

type AddWorkflowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkflowResponse) GoString() string {
	return s.String()
}

func (s *AddWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkflowResponse) GetBody() *AddWorkflowResponseBody {
	return s.Body
}

func (s *AddWorkflowResponse) SetHeaders(v map[string]*string) *AddWorkflowResponse {
	s.Headers = v
	return s
}

func (s *AddWorkflowResponse) SetStatusCode(v int32) *AddWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkflowResponse) SetBody(v *AddWorkflowResponseBody) *AddWorkflowResponse {
	s.Body = v
	return s
}

func (s *AddWorkflowResponse) Validate() error {
	return dara.Validate(s)
}
