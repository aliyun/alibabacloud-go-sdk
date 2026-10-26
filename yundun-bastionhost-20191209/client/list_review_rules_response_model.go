// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReviewRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReviewRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReviewRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListReviewRulesResponseBody) *ListReviewRulesResponse
	GetBody() *ListReviewRulesResponseBody
}

type ListReviewRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReviewRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReviewRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReviewRulesResponse) GoString() string {
	return s.String()
}

func (s *ListReviewRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReviewRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReviewRulesResponse) GetBody() *ListReviewRulesResponseBody {
	return s.Body
}

func (s *ListReviewRulesResponse) SetHeaders(v map[string]*string) *ListReviewRulesResponse {
	s.Headers = v
	return s
}

func (s *ListReviewRulesResponse) SetStatusCode(v int32) *ListReviewRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReviewRulesResponse) SetBody(v *ListReviewRulesResponseBody) *ListReviewRulesResponse {
	s.Body = v
	return s
}

func (s *ListReviewRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
