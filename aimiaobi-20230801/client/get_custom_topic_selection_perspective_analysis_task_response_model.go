// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTopicSelectionPerspectiveAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
	GetBody() *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) GetBody() *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	return s.Body
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) Validate() error {
	return dara.Validate(s)
}
