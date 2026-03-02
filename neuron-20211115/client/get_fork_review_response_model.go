// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForkReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetForkReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetForkReviewResponse
	GetStatusCode() *int32
	SetBody(v *ForkReview) *GetForkReviewResponse
	GetBody() *ForkReview
}

type GetForkReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForkReview        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetForkReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetForkReviewResponse) GoString() string {
	return s.String()
}

func (s *GetForkReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetForkReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetForkReviewResponse) GetBody() *ForkReview {
	return s.Body
}

func (s *GetForkReviewResponse) SetHeaders(v map[string]*string) *GetForkReviewResponse {
	s.Headers = v
	return s
}

func (s *GetForkReviewResponse) SetStatusCode(v int32) *GetForkReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetForkReviewResponse) SetBody(v *ForkReview) *GetForkReviewResponse {
	s.Body = v
	return s
}

func (s *GetForkReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
