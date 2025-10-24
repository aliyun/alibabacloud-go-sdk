// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnalysisJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAnalysisJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAnalysisJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryAnalysisJobListResponseBody) *QueryAnalysisJobListResponse
	GetBody() *QueryAnalysisJobListResponseBody
}

type QueryAnalysisJobListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAnalysisJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAnalysisJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAnalysisJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAnalysisJobListResponse) GetBody() *QueryAnalysisJobListResponseBody {
	return s.Body
}

func (s *QueryAnalysisJobListResponse) SetHeaders(v map[string]*string) *QueryAnalysisJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryAnalysisJobListResponse) SetStatusCode(v int32) *QueryAnalysisJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAnalysisJobListResponse) SetBody(v *QueryAnalysisJobListResponseBody) *QueryAnalysisJobListResponse {
	s.Body = v
	return s
}

func (s *QueryAnalysisJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
