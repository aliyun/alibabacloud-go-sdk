// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowDraftResponse
	GetStatusCode() *int32
	SetBody(v *FlowResult) *GetFlowDraftResponse
	GetBody() *FlowResult
}

type GetFlowDraftResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowDraftResponse) GoString() string {
	return s.String()
}

func (s *GetFlowDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowDraftResponse) GetBody() *FlowResult {
	return s.Body
}

func (s *GetFlowDraftResponse) SetHeaders(v map[string]*string) *GetFlowDraftResponse {
	s.Headers = v
	return s
}

func (s *GetFlowDraftResponse) SetStatusCode(v int32) *GetFlowDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowDraftResponse) SetBody(v *FlowResult) *GetFlowDraftResponse {
	s.Body = v
	return s
}

func (s *GetFlowDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
