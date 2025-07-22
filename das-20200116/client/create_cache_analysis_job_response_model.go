// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheAnalysisJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCacheAnalysisJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCacheAnalysisJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateCacheAnalysisJobResponseBody) *CreateCacheAnalysisJobResponse
	GetBody() *CreateCacheAnalysisJobResponseBody
}

type CreateCacheAnalysisJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCacheAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCacheAnalysisJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCacheAnalysisJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCacheAnalysisJobResponse) GetBody() *CreateCacheAnalysisJobResponseBody {
	return s.Body
}

func (s *CreateCacheAnalysisJobResponse) SetHeaders(v map[string]*string) *CreateCacheAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheAnalysisJobResponse) SetStatusCode(v int32) *CreateCacheAnalysisJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCacheAnalysisJobResponse) SetBody(v *CreateCacheAnalysisJobResponseBody) *CreateCacheAnalysisJobResponse {
	s.Body = v
	return s
}

func (s *CreateCacheAnalysisJobResponse) Validate() error {
	return dara.Validate(s)
}
