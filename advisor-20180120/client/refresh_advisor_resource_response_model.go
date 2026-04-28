// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshAdvisorResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshAdvisorResourceResponse
	GetStatusCode() *int32
	SetBody(v *RefreshAdvisorResourceResponseBody) *RefreshAdvisorResourceResponse
	GetBody() *RefreshAdvisorResourceResponseBody
}

type RefreshAdvisorResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAdvisorResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAdvisorResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorResourceResponse) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshAdvisorResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshAdvisorResourceResponse) GetBody() *RefreshAdvisorResourceResponseBody {
	return s.Body
}

func (s *RefreshAdvisorResourceResponse) SetHeaders(v map[string]*string) *RefreshAdvisorResourceResponse {
	s.Headers = v
	return s
}

func (s *RefreshAdvisorResourceResponse) SetStatusCode(v int32) *RefreshAdvisorResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAdvisorResourceResponse) SetBody(v *RefreshAdvisorResourceResponseBody) *RefreshAdvisorResourceResponse {
	s.Body = v
	return s
}

func (s *RefreshAdvisorResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
