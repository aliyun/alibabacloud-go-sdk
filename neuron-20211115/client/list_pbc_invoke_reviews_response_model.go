// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcInvokeReviewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcInvokeReviewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcInvokeReviewsResponse
	GetStatusCode() *int32
	SetBody(v *PbcListResult) *ListPbcInvokeReviewsResponse
	GetBody() *PbcListResult
}

type ListPbcInvokeReviewsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcListResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcInvokeReviewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcInvokeReviewsResponse) GoString() string {
	return s.String()
}

func (s *ListPbcInvokeReviewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcInvokeReviewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcInvokeReviewsResponse) GetBody() *PbcListResult {
	return s.Body
}

func (s *ListPbcInvokeReviewsResponse) SetHeaders(v map[string]*string) *ListPbcInvokeReviewsResponse {
	s.Headers = v
	return s
}

func (s *ListPbcInvokeReviewsResponse) SetStatusCode(v int32) *ListPbcInvokeReviewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcInvokeReviewsResponse) SetBody(v *PbcListResult) *ListPbcInvokeReviewsResponse {
	s.Body = v
	return s
}

func (s *ListPbcInvokeReviewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
