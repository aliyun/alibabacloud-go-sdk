// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForkReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateForkReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateForkReviewResponse
	GetStatusCode() *int32
	SetBody(v *ForkReviewCreateResult) *CreateForkReviewResponse
	GetBody() *ForkReviewCreateResult
}

type CreateForkReviewResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForkReviewCreateResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateForkReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateForkReviewResponse) GoString() string {
	return s.String()
}

func (s *CreateForkReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateForkReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateForkReviewResponse) GetBody() *ForkReviewCreateResult {
	return s.Body
}

func (s *CreateForkReviewResponse) SetHeaders(v map[string]*string) *CreateForkReviewResponse {
	s.Headers = v
	return s
}

func (s *CreateForkReviewResponse) SetStatusCode(v int32) *CreateForkReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateForkReviewResponse) SetBody(v *ForkReviewCreateResult) *CreateForkReviewResponse {
	s.Body = v
	return s
}

func (s *CreateForkReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
