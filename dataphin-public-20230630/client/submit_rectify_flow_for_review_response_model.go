// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRectifyFlowForReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitRectifyFlowForReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitRectifyFlowForReviewResponse
	GetStatusCode() *int32
	SetBody(v *SubmitRectifyFlowForReviewResponseBody) *SubmitRectifyFlowForReviewResponse
	GetBody() *SubmitRectifyFlowForReviewResponseBody
}

type SubmitRectifyFlowForReviewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitRectifyFlowForReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitRectifyFlowForReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitRectifyFlowForReviewResponse) GoString() string {
	return s.String()
}

func (s *SubmitRectifyFlowForReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitRectifyFlowForReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitRectifyFlowForReviewResponse) GetBody() *SubmitRectifyFlowForReviewResponseBody {
	return s.Body
}

func (s *SubmitRectifyFlowForReviewResponse) SetHeaders(v map[string]*string) *SubmitRectifyFlowForReviewResponse {
	s.Headers = v
	return s
}

func (s *SubmitRectifyFlowForReviewResponse) SetStatusCode(v int32) *SubmitRectifyFlowForReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitRectifyFlowForReviewResponse) SetBody(v *SubmitRectifyFlowForReviewResponseBody) *SubmitRectifyFlowForReviewResponse {
	s.Body = v
	return s
}

func (s *SubmitRectifyFlowForReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
