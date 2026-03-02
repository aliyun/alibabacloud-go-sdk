// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcInvokeReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcInvokeReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcInvokeReviewResponse
	GetStatusCode() *int32
	SetBody(v *PbcInvokeReview) *GetPbcInvokeReviewResponse
	GetBody() *PbcInvokeReview
}

type GetPbcInvokeReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcInvokeReview   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcInvokeReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcInvokeReviewResponse) GoString() string {
	return s.String()
}

func (s *GetPbcInvokeReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcInvokeReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcInvokeReviewResponse) GetBody() *PbcInvokeReview {
	return s.Body
}

func (s *GetPbcInvokeReviewResponse) SetHeaders(v map[string]*string) *GetPbcInvokeReviewResponse {
	s.Headers = v
	return s
}

func (s *GetPbcInvokeReviewResponse) SetStatusCode(v int32) *GetPbcInvokeReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcInvokeReviewResponse) SetBody(v *PbcInvokeReview) *GetPbcInvokeReviewResponse {
	s.Body = v
	return s
}

func (s *GetPbcInvokeReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
