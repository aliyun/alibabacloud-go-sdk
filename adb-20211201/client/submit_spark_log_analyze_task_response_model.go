// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSparkLogAnalyzeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSparkLogAnalyzeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSparkLogAnalyzeTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSparkLogAnalyzeTaskResponseBody) *SubmitSparkLogAnalyzeTaskResponse
	GetBody() *SubmitSparkLogAnalyzeTaskResponseBody
}

type SubmitSparkLogAnalyzeTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSparkLogAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSparkLogAnalyzeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkLogAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkLogAnalyzeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSparkLogAnalyzeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSparkLogAnalyzeTaskResponse) GetBody() *SubmitSparkLogAnalyzeTaskResponseBody {
	return s.Body
}

func (s *SubmitSparkLogAnalyzeTaskResponse) SetHeaders(v map[string]*string) *SubmitSparkLogAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponse) SetStatusCode(v int32) *SubmitSparkLogAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponse) SetBody(v *SubmitSparkLogAnalyzeTaskResponseBody) *SubmitSparkLogAnalyzeTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
