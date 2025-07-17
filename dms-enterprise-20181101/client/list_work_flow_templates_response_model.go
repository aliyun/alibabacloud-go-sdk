// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkFlowTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkFlowTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkFlowTemplatesResponseBody) *ListWorkFlowTemplatesResponse
	GetBody() *ListWorkFlowTemplatesResponseBody
}

type ListWorkFlowTemplatesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkFlowTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkFlowTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkFlowTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkFlowTemplatesResponse) GetBody() *ListWorkFlowTemplatesResponseBody {
	return s.Body
}

func (s *ListWorkFlowTemplatesResponse) SetHeaders(v map[string]*string) *ListWorkFlowTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkFlowTemplatesResponse) SetStatusCode(v int32) *ListWorkFlowTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkFlowTemplatesResponse) SetBody(v *ListWorkFlowTemplatesResponseBody) *ListWorkFlowTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListWorkFlowTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
