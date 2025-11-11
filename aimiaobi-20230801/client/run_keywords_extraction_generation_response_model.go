// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunKeywordsExtractionGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunKeywordsExtractionGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunKeywordsExtractionGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunKeywordsExtractionGenerationResponseBody) *RunKeywordsExtractionGenerationResponse
	GetBody() *RunKeywordsExtractionGenerationResponseBody
}

type RunKeywordsExtractionGenerationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunKeywordsExtractionGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunKeywordsExtractionGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunKeywordsExtractionGenerationResponse) GetBody() *RunKeywordsExtractionGenerationResponseBody {
	return s.Body
}

func (s *RunKeywordsExtractionGenerationResponse) SetHeaders(v map[string]*string) *RunKeywordsExtractionGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponse) SetStatusCode(v int32) *RunKeywordsExtractionGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponse) SetBody(v *RunKeywordsExtractionGenerationResponseBody) *RunKeywordsExtractionGenerationResponse {
	s.Body = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
