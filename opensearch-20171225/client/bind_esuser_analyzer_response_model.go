// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindESUserAnalyzerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindESUserAnalyzerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindESUserAnalyzerResponse
	GetStatusCode() *int32
	SetBody(v *BindESUserAnalyzerResponseBody) *BindESUserAnalyzerResponse
	GetBody() *BindESUserAnalyzerResponseBody
}

type BindESUserAnalyzerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindESUserAnalyzerResponse) String() string {
	return dara.Prettify(s)
}

func (s BindESUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindESUserAnalyzerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindESUserAnalyzerResponse) GetBody() *BindESUserAnalyzerResponseBody {
	return s.Body
}

func (s *BindESUserAnalyzerResponse) SetHeaders(v map[string]*string) *BindESUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *BindESUserAnalyzerResponse) SetStatusCode(v int32) *BindESUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *BindESUserAnalyzerResponse) SetBody(v *BindESUserAnalyzerResponseBody) *BindESUserAnalyzerResponse {
	s.Body = v
	return s
}

func (s *BindESUserAnalyzerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
