// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIDiffAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAIDiffAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAIDiffAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *StartAIDiffAnalysisResponseBody) *StartAIDiffAnalysisResponse
	GetBody() *StartAIDiffAnalysisResponseBody
}

type StartAIDiffAnalysisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAIDiffAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAIDiffAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAIDiffAnalysisResponse) GoString() string {
	return s.String()
}

func (s *StartAIDiffAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAIDiffAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAIDiffAnalysisResponse) GetBody() *StartAIDiffAnalysisResponseBody {
	return s.Body
}

func (s *StartAIDiffAnalysisResponse) SetHeaders(v map[string]*string) *StartAIDiffAnalysisResponse {
	s.Headers = v
	return s
}

func (s *StartAIDiffAnalysisResponse) SetStatusCode(v int32) *StartAIDiffAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAIDiffAnalysisResponse) SetBody(v *StartAIDiffAnalysisResponseBody) *StartAIDiffAnalysisResponse {
	s.Body = v
	return s
}

func (s *StartAIDiffAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
