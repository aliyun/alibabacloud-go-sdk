// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkLogAnalyzeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillSparkLogAnalyzeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillSparkLogAnalyzeTaskResponse
	GetStatusCode() *int32
	SetBody(v *KillSparkLogAnalyzeTaskResponseBody) *KillSparkLogAnalyzeTaskResponse
	GetBody() *KillSparkLogAnalyzeTaskResponseBody
}

type KillSparkLogAnalyzeTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkLogAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSparkLogAnalyzeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s KillSparkLogAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *KillSparkLogAnalyzeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillSparkLogAnalyzeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillSparkLogAnalyzeTaskResponse) GetBody() *KillSparkLogAnalyzeTaskResponseBody {
	return s.Body
}

func (s *KillSparkLogAnalyzeTaskResponse) SetHeaders(v map[string]*string) *KillSparkLogAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponse) SetStatusCode(v int32) *KillSparkLogAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponse) SetBody(v *KillSparkLogAnalyzeTaskResponseBody) *KillSparkLogAnalyzeTaskResponse {
	s.Body = v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
