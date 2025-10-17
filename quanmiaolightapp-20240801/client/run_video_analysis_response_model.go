// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunVideoAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunVideoAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *RunVideoAnalysisResponseBody) *RunVideoAnalysisResponse
	GetBody() *RunVideoAnalysisResponseBody
}

type RunVideoAnalysisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunVideoAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunVideoAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunVideoAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunVideoAnalysisResponse) GetBody() *RunVideoAnalysisResponseBody {
	return s.Body
}

func (s *RunVideoAnalysisResponse) SetHeaders(v map[string]*string) *RunVideoAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunVideoAnalysisResponse) SetStatusCode(v int32) *RunVideoAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunVideoAnalysisResponse) SetBody(v *RunVideoAnalysisResponseBody) *RunVideoAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunVideoAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
