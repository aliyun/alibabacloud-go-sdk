// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewRectifyFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReviewRectifyFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReviewRectifyFlowResponse
	GetStatusCode() *int32
	SetBody(v *ReviewRectifyFlowResponseBody) *ReviewRectifyFlowResponse
	GetBody() *ReviewRectifyFlowResponseBody
}

type ReviewRectifyFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReviewRectifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReviewRectifyFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ReviewRectifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ReviewRectifyFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReviewRectifyFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReviewRectifyFlowResponse) GetBody() *ReviewRectifyFlowResponseBody {
	return s.Body
}

func (s *ReviewRectifyFlowResponse) SetHeaders(v map[string]*string) *ReviewRectifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ReviewRectifyFlowResponse) SetStatusCode(v int32) *ReviewRectifyFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ReviewRectifyFlowResponse) SetBody(v *ReviewRectifyFlowResponseBody) *ReviewRectifyFlowResponse {
	s.Body = v
	return s
}

func (s *ReviewRectifyFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
