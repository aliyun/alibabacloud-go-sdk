// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *ResumeWorkItemResponseBody) *ResumeWorkItemResponse
	GetBody() *ResumeWorkItemResponseBody
}

type ResumeWorkItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeWorkItemResponse) GoString() string {
	return s.String()
}

func (s *ResumeWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeWorkItemResponse) GetBody() *ResumeWorkItemResponseBody {
	return s.Body
}

func (s *ResumeWorkItemResponse) SetHeaders(v map[string]*string) *ResumeWorkItemResponse {
	s.Headers = v
	return s
}

func (s *ResumeWorkItemResponse) SetStatusCode(v int32) *ResumeWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeWorkItemResponse) SetBody(v *ResumeWorkItemResponseBody) *ResumeWorkItemResponse {
	s.Body = v
	return s
}

func (s *ResumeWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
