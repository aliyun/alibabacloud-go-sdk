// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTagMiningAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTagMiningAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTagMiningAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTagMiningAnalysisTaskResponseBody) *SubmitTagMiningAnalysisTaskResponse
	GetBody() *SubmitTagMiningAnalysisTaskResponseBody
}

type SubmitTagMiningAnalysisTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTagMiningAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTagMiningAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTagMiningAnalysisTaskResponse) GetBody() *SubmitTagMiningAnalysisTaskResponseBody {
	return s.Body
}

func (s *SubmitTagMiningAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitTagMiningAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponse) SetStatusCode(v int32) *SubmitTagMiningAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponse) SetBody(v *SubmitTagMiningAnalysisTaskResponseBody) *SubmitTagMiningAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
