// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcReviewResponse
	GetStatusCode() *int32
	SetBody(v *PbcReview) *CreatePbcReviewResponse
	GetBody() *PbcReview
}

type CreatePbcReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcReview         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcReviewResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcReviewResponse) GetBody() *PbcReview {
	return s.Body
}

func (s *CreatePbcReviewResponse) SetHeaders(v map[string]*string) *CreatePbcReviewResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcReviewResponse) SetStatusCode(v int32) *CreatePbcReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcReviewResponse) SetBody(v *PbcReview) *CreatePbcReviewResponse {
	s.Body = v
	return s
}

func (s *CreatePbcReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
