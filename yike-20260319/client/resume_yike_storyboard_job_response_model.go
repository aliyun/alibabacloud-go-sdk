// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeYikeStoryboardJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeYikeStoryboardJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeYikeStoryboardJobResponse
	GetStatusCode() *int32
	SetBody(v *ResumeYikeStoryboardJobResponseBody) *ResumeYikeStoryboardJobResponse
	GetBody() *ResumeYikeStoryboardJobResponseBody
}

type ResumeYikeStoryboardJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeYikeStoryboardJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeYikeStoryboardJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeYikeStoryboardJobResponse) GoString() string {
	return s.String()
}

func (s *ResumeYikeStoryboardJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeYikeStoryboardJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeYikeStoryboardJobResponse) GetBody() *ResumeYikeStoryboardJobResponseBody {
	return s.Body
}

func (s *ResumeYikeStoryboardJobResponse) SetHeaders(v map[string]*string) *ResumeYikeStoryboardJobResponse {
	s.Headers = v
	return s
}

func (s *ResumeYikeStoryboardJobResponse) SetStatusCode(v int32) *ResumeYikeStoryboardJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeYikeStoryboardJobResponse) SetBody(v *ResumeYikeStoryboardJobResponseBody) *ResumeYikeStoryboardJobResponse {
	s.Body = v
	return s
}

func (s *ResumeYikeStoryboardJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
