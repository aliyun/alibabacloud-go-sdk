// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowEndpointsResult) *ListFlowEndpointsResponse
	GetBody() *ListFlowEndpointsResult
}

type ListFlowEndpointsResponse struct {
	Headers    map[string]*string       `json:"headers" xml:"headers"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowEndpointsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowEndpointsResponse) GetBody() *ListFlowEndpointsResult {
	return s.Body
}

func (s *ListFlowEndpointsResponse) SetHeaders(v map[string]*string) *ListFlowEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowEndpointsResponse) SetStatusCode(v int32) *ListFlowEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowEndpointsResponse) SetBody(v *ListFlowEndpointsResult) *ListFlowEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListFlowEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
