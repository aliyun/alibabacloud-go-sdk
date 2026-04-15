// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowEndpointResponse
	GetStatusCode() *int32
	SetBody(v *FlowEndpointResult) *CreateFlowEndpointResponse
	GetBody() *FlowEndpointResult
}

type CreateFlowEndpointResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowEndpointResponse) GetBody() *FlowEndpointResult {
	return s.Body
}

func (s *CreateFlowEndpointResponse) SetHeaders(v map[string]*string) *CreateFlowEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowEndpointResponse) SetStatusCode(v int32) *CreateFlowEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowEndpointResponse) SetBody(v *FlowEndpointResult) *CreateFlowEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateFlowEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
