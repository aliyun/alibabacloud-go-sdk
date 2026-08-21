// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAiAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaAiAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaAiAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaAiAnalysisResponseBody) *GetMediaAiAnalysisResponse
	GetBody() *GetMediaAiAnalysisResponseBody
}

type GetMediaAiAnalysisResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaAiAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaAiAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAiAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAiAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaAiAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaAiAnalysisResponse) GetBody() *GetMediaAiAnalysisResponseBody {
	return s.Body
}

func (s *GetMediaAiAnalysisResponse) SetHeaders(v map[string]*string) *GetMediaAiAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAiAnalysisResponse) SetStatusCode(v int32) *GetMediaAiAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaAiAnalysisResponse) SetBody(v *GetMediaAiAnalysisResponseBody) *GetMediaAiAnalysisResponse {
	s.Body = v
	return s
}

func (s *GetMediaAiAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
