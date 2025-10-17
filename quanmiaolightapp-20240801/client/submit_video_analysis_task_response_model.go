// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoAnalysisTaskResponseBody) *SubmitVideoAnalysisTaskResponse
	GetBody() *SubmitVideoAnalysisTaskResponseBody
}

type SubmitVideoAnalysisTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoAnalysisTaskResponse) GetBody() *SubmitVideoAnalysisTaskResponseBody {
	return s.Body
}

func (s *SubmitVideoAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitVideoAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoAnalysisTaskResponse) SetStatusCode(v int32) *SubmitVideoAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoAnalysisTaskResponse) SetBody(v *SubmitVideoAnalysisTaskResponseBody) *SubmitVideoAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
