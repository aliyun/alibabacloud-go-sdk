// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsBatchFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOssCheckResultsBatchFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOssCheckResultsBatchFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOssCheckResultsBatchFeedbackResponseBody) *UpdateOssCheckResultsBatchFeedbackResponse
	GetBody() *UpdateOssCheckResultsBatchFeedbackResponseBody
}

type UpdateOssCheckResultsBatchFeedbackResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssCheckResultsBatchFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssCheckResultsBatchFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsBatchFeedbackResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) GetBody() *UpdateOssCheckResultsBatchFeedbackResponseBody {
	return s.Body
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) SetHeaders(v map[string]*string) *UpdateOssCheckResultsBatchFeedbackResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) SetStatusCode(v int32) *UpdateOssCheckResultsBatchFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) SetBody(v *UpdateOssCheckResultsBatchFeedbackResponseBody) *UpdateOssCheckResultsBatchFeedbackResponse {
	s.Body = v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponse) Validate() error {
	return dara.Validate(s)
}
