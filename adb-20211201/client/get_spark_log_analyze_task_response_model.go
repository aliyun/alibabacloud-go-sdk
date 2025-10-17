// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkLogAnalyzeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkLogAnalyzeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkLogAnalyzeTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkLogAnalyzeTaskResponseBody) *GetSparkLogAnalyzeTaskResponse
	GetBody() *GetSparkLogAnalyzeTaskResponseBody
}

type GetSparkLogAnalyzeTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkLogAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkLogAnalyzeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLogAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSparkLogAnalyzeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkLogAnalyzeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkLogAnalyzeTaskResponse) GetBody() *GetSparkLogAnalyzeTaskResponseBody {
	return s.Body
}

func (s *GetSparkLogAnalyzeTaskResponse) SetHeaders(v map[string]*string) *GetSparkLogAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponse) SetStatusCode(v int32) *GetSparkLogAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponse) SetBody(v *GetSparkLogAnalyzeTaskResponseBody) *GetSparkLogAnalyzeTaskResponse {
	s.Body = v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
