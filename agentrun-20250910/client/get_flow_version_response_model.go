// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *FlowVersionResult) *GetFlowVersionResponse
	GetBody() *FlowVersionResult
}

type GetFlowVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowVersionResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowVersionResponse) GetBody() *FlowVersionResult {
	return s.Body
}

func (s *GetFlowVersionResponse) SetHeaders(v map[string]*string) *GetFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFlowVersionResponse) SetStatusCode(v int32) *GetFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowVersionResponse) SetBody(v *FlowVersionResult) *GetFlowVersionResponse {
	s.Body = v
	return s
}

func (s *GetFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
