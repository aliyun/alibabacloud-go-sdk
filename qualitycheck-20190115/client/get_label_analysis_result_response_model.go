// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelAnalysisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLabelAnalysisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLabelAnalysisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetLabelAnalysisResultResponseBody) *GetLabelAnalysisResultResponse
	GetBody() *GetLabelAnalysisResultResponseBody
}

type GetLabelAnalysisResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLabelAnalysisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLabelAnalysisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponse) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLabelAnalysisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLabelAnalysisResultResponse) GetBody() *GetLabelAnalysisResultResponseBody {
	return s.Body
}

func (s *GetLabelAnalysisResultResponse) SetHeaders(v map[string]*string) *GetLabelAnalysisResultResponse {
	s.Headers = v
	return s
}

func (s *GetLabelAnalysisResultResponse) SetStatusCode(v int32) *GetLabelAnalysisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelAnalysisResultResponse) SetBody(v *GetLabelAnalysisResultResponseBody) *GetLabelAnalysisResultResponse {
	s.Body = v
	return s
}

func (s *GetLabelAnalysisResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
