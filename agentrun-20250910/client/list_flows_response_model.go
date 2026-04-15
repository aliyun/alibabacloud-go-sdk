// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowsResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowsResult) *ListFlowsResponse
	GetBody() *ListFlowsResult
}

type ListFlowsResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowsResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowsResponse) GetBody() *ListFlowsResult {
	return s.Body
}

func (s *ListFlowsResponse) SetHeaders(v map[string]*string) *ListFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowsResponse) SetStatusCode(v int32) *ListFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowsResponse) SetBody(v *ListFlowsResult) *ListFlowsResponse {
	s.Body = v
	return s
}

func (s *ListFlowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
