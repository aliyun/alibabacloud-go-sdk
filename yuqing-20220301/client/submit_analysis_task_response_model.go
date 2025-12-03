// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAnalysisTaskResponseBody) *SubmitAnalysisTaskResponse
	GetBody() *SubmitAnalysisTaskResponseBody
}

type SubmitAnalysisTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAnalysisTaskResponse) GetBody() *SubmitAnalysisTaskResponseBody {
	return s.Body
}

func (s *SubmitAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAnalysisTaskResponse) SetStatusCode(v int32) *SubmitAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAnalysisTaskResponse) SetBody(v *SubmitAnalysisTaskResponseBody) *SubmitAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
