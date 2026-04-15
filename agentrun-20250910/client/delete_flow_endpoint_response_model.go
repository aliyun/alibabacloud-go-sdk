// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CommonResult) *DeleteFlowEndpointResponse
	GetBody() *CommonResult
}

type DeleteFlowEndpointResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommonResult      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowEndpointResponse) GetBody() *CommonResult {
	return s.Body
}

func (s *DeleteFlowEndpointResponse) SetHeaders(v map[string]*string) *DeleteFlowEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowEndpointResponse) SetStatusCode(v int32) *DeleteFlowEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowEndpointResponse) SetBody(v *CommonResult) *DeleteFlowEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
