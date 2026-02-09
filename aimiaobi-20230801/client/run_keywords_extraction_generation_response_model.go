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
	SetId(v string) *RunKeywordsExtractionGenerationResponse
	GetId() *string
	SetEvent(v string) *RunKeywordsExtractionGenerationResponse
	GetEvent() *string
	SetBody(v *RunKeywordsExtractionGenerationResponseBody) *RunKeywordsExtractionGenerationResponse
	GetBody() *RunKeywordsExtractionGenerationResponseBody
}

type RunKeywordsExtractionGenerationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                                      `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                                      `json:"event,omitempty" xml:"event,omitempty"`
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

func (s *RunKeywordsExtractionGenerationResponse) GetId() *string {
	return s.Id
}

func (s *RunKeywordsExtractionGenerationResponse) GetEvent() *string {
	return s.Event
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

func (s *RunKeywordsExtractionGenerationResponse) SetId(v string) *RunKeywordsExtractionGenerationResponse {
	s.Id = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponse) SetEvent(v string) *RunKeywordsExtractionGenerationResponse {
	s.Event = &v
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
