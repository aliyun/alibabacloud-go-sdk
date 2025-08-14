// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaAiAnalysisJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaAiAnalysisJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaAiAnalysisJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaAiAnalysisJobResponseBody) *SubmitMediaAiAnalysisJobResponse
	GetBody() *SubmitMediaAiAnalysisJobResponseBody
}

type SubmitMediaAiAnalysisJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaAiAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaAiAnalysisJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaAiAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaAiAnalysisJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaAiAnalysisJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaAiAnalysisJobResponse) GetBody() *SubmitMediaAiAnalysisJobResponseBody {
	return s.Body
}

func (s *SubmitMediaAiAnalysisJobResponse) SetHeaders(v map[string]*string) *SubmitMediaAiAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaAiAnalysisJobResponse) SetStatusCode(v int32) *SubmitMediaAiAnalysisJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaAiAnalysisJobResponse) SetBody(v *SubmitMediaAiAnalysisJobResponseBody) *SubmitMediaAiAnalysisJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaAiAnalysisJobResponse) Validate() error {
	return dara.Validate(s)
}
