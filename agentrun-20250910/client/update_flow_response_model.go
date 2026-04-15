// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowResponse
	GetStatusCode() *int32
	SetBody(v *FlowResult) *UpdateFlowResponse
	GetBody() *FlowResult
}

type UpdateFlowResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowResponse) GetBody() *FlowResult {
	return s.Body
}

func (s *UpdateFlowResponse) SetHeaders(v map[string]*string) *UpdateFlowResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowResponse) SetStatusCode(v int32) *UpdateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowResponse) SetBody(v *FlowResult) *UpdateFlowResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
