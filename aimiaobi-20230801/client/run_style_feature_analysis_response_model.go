// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleFeatureAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunStyleFeatureAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunStyleFeatureAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *RunStyleFeatureAnalysisResponseBody) *RunStyleFeatureAnalysisResponse
	GetBody() *RunStyleFeatureAnalysisResponseBody
}

type RunStyleFeatureAnalysisResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunStyleFeatureAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStyleFeatureAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunStyleFeatureAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunStyleFeatureAnalysisResponse) GetBody() *RunStyleFeatureAnalysisResponseBody {
	return s.Body
}

func (s *RunStyleFeatureAnalysisResponse) SetHeaders(v map[string]*string) *RunStyleFeatureAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetStatusCode(v int32) *RunStyleFeatureAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetBody(v *RunStyleFeatureAnalysisResponseBody) *RunStyleFeatureAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
