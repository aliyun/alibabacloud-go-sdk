// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheAnalysisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCacheAnalysisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCacheAnalysisTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCacheAnalysisTaskResponseBody) *CreateCacheAnalysisTaskResponse
	GetBody() *CreateCacheAnalysisTaskResponseBody
}

type CreateCacheAnalysisTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCacheAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCacheAnalysisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCacheAnalysisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCacheAnalysisTaskResponse) GetBody() *CreateCacheAnalysisTaskResponseBody {
	return s.Body
}

func (s *CreateCacheAnalysisTaskResponse) SetHeaders(v map[string]*string) *CreateCacheAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheAnalysisTaskResponse) SetStatusCode(v int32) *CreateCacheAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCacheAnalysisTaskResponse) SetBody(v *CreateCacheAnalysisTaskResponseBody) *CreateCacheAnalysisTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCacheAnalysisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
