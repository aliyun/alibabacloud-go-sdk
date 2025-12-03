// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemWorkFlowStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkItemWorkFlowStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkItemWorkFlowStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkItemWorkFlowStatusResponseBody) *ListWorkItemWorkFlowStatusResponse
	GetBody() *ListWorkItemWorkFlowStatusResponseBody
}

type ListWorkItemWorkFlowStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkItemWorkFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkItemWorkFlowStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusResponse) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkItemWorkFlowStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkItemWorkFlowStatusResponse) GetBody() *ListWorkItemWorkFlowStatusResponseBody {
	return s.Body
}

func (s *ListWorkItemWorkFlowStatusResponse) SetHeaders(v map[string]*string) *ListWorkItemWorkFlowStatusResponse {
	s.Headers = v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponse) SetStatusCode(v int32) *ListWorkItemWorkFlowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponse) SetBody(v *ListWorkItemWorkFlowStatusResponseBody) *ListWorkItemWorkFlowStatusResponse {
	s.Body = v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
