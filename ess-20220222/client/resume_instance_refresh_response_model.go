// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeInstanceRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeInstanceRefreshResponse
	GetStatusCode() *int32
	SetBody(v *ResumeInstanceRefreshResponseBody) *ResumeInstanceRefreshResponse
	GetBody() *ResumeInstanceRefreshResponseBody
}

type ResumeInstanceRefreshResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeInstanceRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeInstanceRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceRefreshResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeInstanceRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeInstanceRefreshResponse) GetBody() *ResumeInstanceRefreshResponseBody {
	return s.Body
}

func (s *ResumeInstanceRefreshResponse) SetHeaders(v map[string]*string) *ResumeInstanceRefreshResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceRefreshResponse) SetStatusCode(v int32) *ResumeInstanceRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceRefreshResponse) SetBody(v *ResumeInstanceRefreshResponseBody) *ResumeInstanceRefreshResponse {
	s.Body = v
	return s
}

func (s *ResumeInstanceRefreshResponse) Validate() error {
	return dara.Validate(s)
}
