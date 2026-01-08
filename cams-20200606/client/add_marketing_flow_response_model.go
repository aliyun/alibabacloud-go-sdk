// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMarketingFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMarketingFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMarketingFlowResponse
	GetStatusCode() *int32
	SetBody(v *AddMarketingFlowResponseBody) *AddMarketingFlowResponse
	GetBody() *AddMarketingFlowResponseBody
}

type AddMarketingFlowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMarketingFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMarketingFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMarketingFlowResponse) GoString() string {
	return s.String()
}

func (s *AddMarketingFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMarketingFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMarketingFlowResponse) GetBody() *AddMarketingFlowResponseBody {
	return s.Body
}

func (s *AddMarketingFlowResponse) SetHeaders(v map[string]*string) *AddMarketingFlowResponse {
	s.Headers = v
	return s
}

func (s *AddMarketingFlowResponse) SetStatusCode(v int32) *AddMarketingFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMarketingFlowResponse) SetBody(v *AddMarketingFlowResponseBody) *AddMarketingFlowResponse {
	s.Body = v
	return s
}

func (s *AddMarketingFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
