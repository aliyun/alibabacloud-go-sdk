// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomSourceTopicAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomSourceTopicAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomSourceTopicAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomSourceTopicAnalysisTaskResponseBody) *GetCustomSourceTopicAnalysisTaskResponse
	GetBody() *GetCustomSourceTopicAnalysisTaskResponseBody
}

type GetCustomSourceTopicAnalysisTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomSourceTopicAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomSourceTopicAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomSourceTopicAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) GetBody() *GetCustomSourceTopicAnalysisTaskResponseBody {
	return s.Body
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetCustomSourceTopicAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) SetStatusCode(v int32) *GetCustomSourceTopicAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) SetBody(v *GetCustomSourceTopicAnalysisTaskResponseBody) *GetCustomSourceTopicAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponse) Validate() error {
	return dara.Validate(s)
}
