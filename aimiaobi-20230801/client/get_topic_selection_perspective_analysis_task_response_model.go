// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicSelectionPerspectiveAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicSelectionPerspectiveAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicSelectionPerspectiveAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) *GetTopicSelectionPerspectiveAnalysisTaskResponse
	GetBody() *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
}

type GetTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) GetBody() *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	return s.Body
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *GetTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) *GetTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) Validate() error {
	return dara.Validate(s)
}
