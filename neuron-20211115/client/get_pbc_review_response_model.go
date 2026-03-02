// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcReviewResponse
	GetStatusCode() *int32
	SetBody(v *PbcReview) *GetPbcReviewResponse
	GetBody() *PbcReview
}

type GetPbcReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcReview         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcReviewResponse) GoString() string {
	return s.String()
}

func (s *GetPbcReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcReviewResponse) GetBody() *PbcReview {
	return s.Body
}

func (s *GetPbcReviewResponse) SetHeaders(v map[string]*string) *GetPbcReviewResponse {
	s.Headers = v
	return s
}

func (s *GetPbcReviewResponse) SetStatusCode(v int32) *GetPbcReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcReviewResponse) SetBody(v *PbcReview) *GetPbcReviewResponse {
	s.Body = v
	return s
}

func (s *GetPbcReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
