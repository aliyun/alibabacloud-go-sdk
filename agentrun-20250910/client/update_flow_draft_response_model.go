// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowDraftResponse
	GetStatusCode() *int32
	SetBody(v *FlowResult) *UpdateFlowDraftResponse
	GetBody() *FlowResult
}

type UpdateFlowDraftResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowDraftResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowDraftResponse) GetBody() *FlowResult {
	return s.Body
}

func (s *UpdateFlowDraftResponse) SetHeaders(v map[string]*string) *UpdateFlowDraftResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowDraftResponse) SetStatusCode(v int32) *UpdateFlowDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowDraftResponse) SetBody(v *FlowResult) *UpdateFlowDraftResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
