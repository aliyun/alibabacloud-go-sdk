// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityAnalysisResponseBody) *GetDataQualityAnalysisResponse
	GetBody() *GetDataQualityAnalysisResponseBody
}

type GetDataQualityAnalysisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityAnalysisResponse) GetBody() *GetDataQualityAnalysisResponseBody {
	return s.Body
}

func (s *GetDataQualityAnalysisResponse) SetHeaders(v map[string]*string) *GetDataQualityAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityAnalysisResponse) SetStatusCode(v int32) *GetDataQualityAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityAnalysisResponse) SetBody(v *GetDataQualityAnalysisResponseBody) *GetDataQualityAnalysisResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
