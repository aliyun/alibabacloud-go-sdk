// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEnterpriseVocAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunEnterpriseVocAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunEnterpriseVocAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *RunEnterpriseVocAnalysisResponseBody) *RunEnterpriseVocAnalysisResponse
	GetBody() *RunEnterpriseVocAnalysisResponseBody
}

type RunEnterpriseVocAnalysisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunEnterpriseVocAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunEnterpriseVocAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunEnterpriseVocAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunEnterpriseVocAnalysisResponse) GetBody() *RunEnterpriseVocAnalysisResponseBody {
	return s.Body
}

func (s *RunEnterpriseVocAnalysisResponse) SetHeaders(v map[string]*string) *RunEnterpriseVocAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) SetStatusCode(v int32) *RunEnterpriseVocAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) SetBody(v *RunEnterpriseVocAnalysisResponseBody) *RunEnterpriseVocAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
