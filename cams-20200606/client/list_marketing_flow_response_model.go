// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketingFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMarketingFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMarketingFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListMarketingFlowResponseBody) *ListMarketingFlowResponse
	GetBody() *ListMarketingFlowResponseBody
}

type ListMarketingFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMarketingFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMarketingFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMarketingFlowResponse) GoString() string {
	return s.String()
}

func (s *ListMarketingFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMarketingFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMarketingFlowResponse) GetBody() *ListMarketingFlowResponseBody {
	return s.Body
}

func (s *ListMarketingFlowResponse) SetHeaders(v map[string]*string) *ListMarketingFlowResponse {
	s.Headers = v
	return s
}

func (s *ListMarketingFlowResponse) SetStatusCode(v int32) *ListMarketingFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMarketingFlowResponse) SetBody(v *ListMarketingFlowResponseBody) *ListMarketingFlowResponse {
	s.Body = v
	return s
}

func (s *ListMarketingFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
