// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanResultFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScanResultFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScanResultFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScanResultFeedbackResponseBody) *UpdateScanResultFeedbackResponse
	GetBody() *UpdateScanResultFeedbackResponseBody
}

type UpdateScanResultFeedbackResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScanResultFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScanResultFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanResultFeedbackResponse) GoString() string {
	return s.String()
}

func (s *UpdateScanResultFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScanResultFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScanResultFeedbackResponse) GetBody() *UpdateScanResultFeedbackResponseBody {
	return s.Body
}

func (s *UpdateScanResultFeedbackResponse) SetHeaders(v map[string]*string) *UpdateScanResultFeedbackResponse {
	s.Headers = v
	return s
}

func (s *UpdateScanResultFeedbackResponse) SetStatusCode(v int32) *UpdateScanResultFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScanResultFeedbackResponse) SetBody(v *UpdateScanResultFeedbackResponseBody) *UpdateScanResultFeedbackResponse {
	s.Body = v
	return s
}

func (s *UpdateScanResultFeedbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
