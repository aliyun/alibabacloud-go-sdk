// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkFlowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkFlowsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkFlowsResponseBody) *ListWorkFlowsResponse
	GetBody() *ListWorkFlowsResponseBody
}

type ListWorkFlowsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkFlowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkFlowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkFlowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkFlowsResponse) GetBody() *ListWorkFlowsResponseBody {
	return s.Body
}

func (s *ListWorkFlowsResponse) SetHeaders(v map[string]*string) *ListWorkFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkFlowsResponse) SetStatusCode(v int32) *ListWorkFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkFlowsResponse) SetBody(v *ListWorkFlowsResponseBody) *ListWorkFlowsResponse {
	s.Body = v
	return s
}

func (s *ListWorkFlowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
