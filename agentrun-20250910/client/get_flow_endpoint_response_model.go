// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowEndpointResponse
	GetStatusCode() *int32
	SetBody(v *FlowEndpointResult) *GetFlowEndpointResponse
	GetBody() *FlowEndpointResult
}

type GetFlowEndpointResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetFlowEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowEndpointResponse) GetBody() *FlowEndpointResult {
	return s.Body
}

func (s *GetFlowEndpointResponse) SetHeaders(v map[string]*string) *GetFlowEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetFlowEndpointResponse) SetStatusCode(v int32) *GetFlowEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowEndpointResponse) SetBody(v *FlowEndpointResult) *GetFlowEndpointResponse {
	s.Body = v
	return s
}

func (s *GetFlowEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
