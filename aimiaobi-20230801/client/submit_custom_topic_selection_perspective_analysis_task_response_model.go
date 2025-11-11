// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
	GetBody() *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) GetBody() *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	return s.Body
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
