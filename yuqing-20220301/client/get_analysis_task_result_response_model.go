// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnalysisTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAnalysisTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAnalysisTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAnalysisTaskResultResponseBody) *GetAnalysisTaskResultResponse
	GetBody() *GetAnalysisTaskResultResponseBody
}

type GetAnalysisTaskResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAnalysisTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnalysisTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAnalysisTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetAnalysisTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAnalysisTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAnalysisTaskResultResponse) GetBody() *GetAnalysisTaskResultResponseBody {
	return s.Body
}

func (s *GetAnalysisTaskResultResponse) SetHeaders(v map[string]*string) *GetAnalysisTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetAnalysisTaskResultResponse) SetStatusCode(v int32) *GetAnalysisTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnalysisTaskResultResponse) SetBody(v *GetAnalysisTaskResultResponseBody) *GetAnalysisTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetAnalysisTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
