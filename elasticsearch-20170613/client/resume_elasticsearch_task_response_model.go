// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeElasticsearchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeElasticsearchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeElasticsearchTaskResponse
	GetStatusCode() *int32
	SetBody(v *ResumeElasticsearchTaskResponseBody) *ResumeElasticsearchTaskResponse
	GetBody() *ResumeElasticsearchTaskResponseBody
}

type ResumeElasticsearchTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeElasticsearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeElasticsearchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeElasticsearchTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeElasticsearchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeElasticsearchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeElasticsearchTaskResponse) GetBody() *ResumeElasticsearchTaskResponseBody {
	return s.Body
}

func (s *ResumeElasticsearchTaskResponse) SetHeaders(v map[string]*string) *ResumeElasticsearchTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeElasticsearchTaskResponse) SetStatusCode(v int32) *ResumeElasticsearchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeElasticsearchTaskResponse) SetBody(v *ResumeElasticsearchTaskResponseBody) *ResumeElasticsearchTaskResponse {
	s.Body = v
	return s
}

func (s *ResumeElasticsearchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
