// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitReviewInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitReviewInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitReviewInfoResponse
	GetStatusCode() *int32
	SetBody(v *SubmitReviewInfoResponseBody) *SubmitReviewInfoResponse
	GetBody() *SubmitReviewInfoResponseBody
}

type SubmitReviewInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitReviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitReviewInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitReviewInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitReviewInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitReviewInfoResponse) GetBody() *SubmitReviewInfoResponseBody {
	return s.Body
}

func (s *SubmitReviewInfoResponse) SetHeaders(v map[string]*string) *SubmitReviewInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitReviewInfoResponse) SetStatusCode(v int32) *SubmitReviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitReviewInfoResponse) SetBody(v *SubmitReviewInfoResponseBody) *SubmitReviewInfoResponse {
	s.Body = v
	return s
}

func (s *SubmitReviewInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
