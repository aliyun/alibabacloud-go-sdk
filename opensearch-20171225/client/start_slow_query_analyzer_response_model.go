// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSlowQueryAnalyzerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSlowQueryAnalyzerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSlowQueryAnalyzerResponse
	GetStatusCode() *int32
	SetBody(v *StartSlowQueryAnalyzerResponseBody) *StartSlowQueryAnalyzerResponse
	GetBody() *StartSlowQueryAnalyzerResponseBody
}

type StartSlowQueryAnalyzerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSlowQueryAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSlowQueryAnalyzerResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSlowQueryAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *StartSlowQueryAnalyzerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSlowQueryAnalyzerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSlowQueryAnalyzerResponse) GetBody() *StartSlowQueryAnalyzerResponseBody {
	return s.Body
}

func (s *StartSlowQueryAnalyzerResponse) SetHeaders(v map[string]*string) *StartSlowQueryAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *StartSlowQueryAnalyzerResponse) SetStatusCode(v int32) *StartSlowQueryAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSlowQueryAnalyzerResponse) SetBody(v *StartSlowQueryAnalyzerResponseBody) *StartSlowQueryAnalyzerResponse {
	s.Body = v
	return s
}

func (s *StartSlowQueryAnalyzerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
