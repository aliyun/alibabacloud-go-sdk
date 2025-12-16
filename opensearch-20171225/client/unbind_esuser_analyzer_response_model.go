// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindESUserAnalyzerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindESUserAnalyzerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindESUserAnalyzerResponse
	GetStatusCode() *int32
	SetBody(v *UnbindESUserAnalyzerResponseBody) *UnbindESUserAnalyzerResponse
	GetBody() *UnbindESUserAnalyzerResponseBody
}

type UnbindESUserAnalyzerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindESUserAnalyzerResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindESUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindESUserAnalyzerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindESUserAnalyzerResponse) GetBody() *UnbindESUserAnalyzerResponseBody {
	return s.Body
}

func (s *UnbindESUserAnalyzerResponse) SetHeaders(v map[string]*string) *UnbindESUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *UnbindESUserAnalyzerResponse) SetStatusCode(v int32) *UnbindESUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindESUserAnalyzerResponse) SetBody(v *UnbindESUserAnalyzerResponseBody) *UnbindESUserAnalyzerResponse {
	s.Body = v
	return s
}

func (s *UnbindESUserAnalyzerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
