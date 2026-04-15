// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowEndpointResponse
	GetStatusCode() *int32
	SetBody(v *FlowEndpointResult) *UpdateFlowEndpointResponse
	GetBody() *FlowEndpointResult
}

type UpdateFlowEndpointResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowEndpointResponse) GetBody() *FlowEndpointResult {
	return s.Body
}

func (s *UpdateFlowEndpointResponse) SetHeaders(v map[string]*string) *UpdateFlowEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowEndpointResponse) SetStatusCode(v int32) *UpdateFlowEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowEndpointResponse) SetBody(v *FlowEndpointResult) *UpdateFlowEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
