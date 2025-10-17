// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoAnalysisTaskResponseBody) *UpdateVideoAnalysisTaskResponse
	GetBody() *UpdateVideoAnalysisTaskResponseBody
}

type UpdateVideoAnalysisTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoAnalysisTaskResponse) GetBody() *UpdateVideoAnalysisTaskResponseBody {
	return s.Body
}

func (s *UpdateVideoAnalysisTaskResponse) SetHeaders(v map[string]*string) *UpdateVideoAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoAnalysisTaskResponse) SetStatusCode(v int32) *UpdateVideoAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponse) SetBody(v *UpdateVideoAnalysisTaskResponseBody) *UpdateVideoAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
