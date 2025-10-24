// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAnalysisJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAnalysisJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAnalysisJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAnalysisJobResponseBody) *SubmitAnalysisJobResponse
	GetBody() *SubmitAnalysisJobResponseBody
}

type SubmitAnalysisJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAnalysisJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAnalysisJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAnalysisJobResponse) GetBody() *SubmitAnalysisJobResponseBody {
	return s.Body
}

func (s *SubmitAnalysisJobResponse) SetHeaders(v map[string]*string) *SubmitAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAnalysisJobResponse) SetStatusCode(v int32) *SubmitAnalysisJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAnalysisJobResponse) SetBody(v *SubmitAnalysisJobResponseBody) *SubmitAnalysisJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAnalysisJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
