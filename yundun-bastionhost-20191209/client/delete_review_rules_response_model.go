// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReviewRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReviewRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReviewRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReviewRulesResponseBody) *DeleteReviewRulesResponse
	GetBody() *DeleteReviewRulesResponseBody
}

type DeleteReviewRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReviewRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReviewRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReviewRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteReviewRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReviewRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReviewRulesResponse) GetBody() *DeleteReviewRulesResponseBody {
	return s.Body
}

func (s *DeleteReviewRulesResponse) SetHeaders(v map[string]*string) *DeleteReviewRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteReviewRulesResponse) SetStatusCode(v int32) *DeleteReviewRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReviewRulesResponse) SetBody(v *DeleteReviewRulesResponseBody) *DeleteReviewRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteReviewRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
