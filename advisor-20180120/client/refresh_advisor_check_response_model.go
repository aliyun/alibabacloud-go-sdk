// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshAdvisorCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshAdvisorCheckResponse
	GetStatusCode() *int32
	SetBody(v *RefreshAdvisorCheckResponseBody) *RefreshAdvisorCheckResponse
	GetBody() *RefreshAdvisorCheckResponseBody
}

type RefreshAdvisorCheckResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAdvisorCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAdvisorCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCheckResponse) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshAdvisorCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshAdvisorCheckResponse) GetBody() *RefreshAdvisorCheckResponseBody {
	return s.Body
}

func (s *RefreshAdvisorCheckResponse) SetHeaders(v map[string]*string) *RefreshAdvisorCheckResponse {
	s.Headers = v
	return s
}

func (s *RefreshAdvisorCheckResponse) SetStatusCode(v int32) *RefreshAdvisorCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAdvisorCheckResponse) SetBody(v *RefreshAdvisorCheckResponseBody) *RefreshAdvisorCheckResponse {
	s.Body = v
	return s
}

func (s *RefreshAdvisorCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
