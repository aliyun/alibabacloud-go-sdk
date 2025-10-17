// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoAnalysisTaskResponseBody) *GetVideoAnalysisTaskResponse
	GetBody() *GetVideoAnalysisTaskResponseBody
}

type GetVideoAnalysisTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoAnalysisTaskResponse) GetBody() *GetVideoAnalysisTaskResponseBody {
	return s.Body
}

func (s *GetVideoAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetVideoAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVideoAnalysisTaskResponse) SetStatusCode(v int32) *GetVideoAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoAnalysisTaskResponse) SetBody(v *GetVideoAnalysisTaskResponseBody) *GetVideoAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *GetVideoAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
