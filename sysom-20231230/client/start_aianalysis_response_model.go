// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAIAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAIAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *StartAIAnalysisResponseBody) *StartAIAnalysisResponse
	GetBody() *StartAIAnalysisResponseBody
}

type StartAIAnalysisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAIAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAIAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAIAnalysisResponse) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAIAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAIAnalysisResponse) GetBody() *StartAIAnalysisResponseBody {
	return s.Body
}

func (s *StartAIAnalysisResponse) SetHeaders(v map[string]*string) *StartAIAnalysisResponse {
	s.Headers = v
	return s
}

func (s *StartAIAnalysisResponse) SetStatusCode(v int32) *StartAIAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAIAnalysisResponse) SetBody(v *StartAIAnalysisResponseBody) *StartAIAnalysisResponse {
	s.Body = v
	return s
}

func (s *StartAIAnalysisResponse) Validate() error {
	return dara.Validate(s)
}
