// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoAnalysisTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoAnalysisTasksResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoAnalysisTasksResponseBody) *UpdateVideoAnalysisTasksResponse
	GetBody() *UpdateVideoAnalysisTasksResponseBody
}

type UpdateVideoAnalysisTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoAnalysisTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoAnalysisTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTasksResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoAnalysisTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoAnalysisTasksResponse) GetBody() *UpdateVideoAnalysisTasksResponseBody {
	return s.Body
}

func (s *UpdateVideoAnalysisTasksResponse) SetHeaders(v map[string]*string) *UpdateVideoAnalysisTasksResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoAnalysisTasksResponse) SetStatusCode(v int32) *UpdateVideoAnalysisTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponse) SetBody(v *UpdateVideoAnalysisTasksResponseBody) *UpdateVideoAnalysisTasksResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoAnalysisTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
