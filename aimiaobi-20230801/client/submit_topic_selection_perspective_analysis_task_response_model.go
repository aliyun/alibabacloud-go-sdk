// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTopicSelectionPerspectiveAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse
	GetBody() *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody
}

type SubmitTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) GetBody() *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	return s.Body
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
