// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendManagedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecommendManagedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecommendManagedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecommendManagedRulesResponseBody) *ListRecommendManagedRulesResponse
	GetBody() *ListRecommendManagedRulesResponseBody
}

type ListRecommendManagedRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecommendManagedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecommendManagedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendManagedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRecommendManagedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecommendManagedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecommendManagedRulesResponse) GetBody() *ListRecommendManagedRulesResponseBody {
	return s.Body
}

func (s *ListRecommendManagedRulesResponse) SetHeaders(v map[string]*string) *ListRecommendManagedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRecommendManagedRulesResponse) SetStatusCode(v int32) *ListRecommendManagedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecommendManagedRulesResponse) SetBody(v *ListRecommendManagedRulesResponseBody) *ListRecommendManagedRulesResponse {
	s.Body = v
	return s
}

func (s *ListRecommendManagedRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
