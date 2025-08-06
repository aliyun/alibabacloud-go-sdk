// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoSummaryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoSummaryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoSummaryJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoSummaryJobResponseBody) *SubmitAIVideoSummaryJobResponse
	GetBody() *SubmitAIVideoSummaryJobResponseBody
}

type SubmitAIVideoSummaryJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoSummaryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoSummaryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoSummaryJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoSummaryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoSummaryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoSummaryJobResponse) GetBody() *SubmitAIVideoSummaryJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoSummaryJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoSummaryJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoSummaryJobResponse) SetStatusCode(v int32) *SubmitAIVideoSummaryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoSummaryJobResponse) SetBody(v *SubmitAIVideoSummaryJobResponseBody) *SubmitAIVideoSummaryJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoSummaryJobResponse) Validate() error {
	return dara.Validate(s)
}
