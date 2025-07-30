// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitReviewInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSubmitReviewInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSubmitReviewInfoResponse
	GetStatusCode() *int32
	SetBody(v *BatchSubmitReviewInfoResponseBody) *BatchSubmitReviewInfoResponse
	GetBody() *BatchSubmitReviewInfoResponseBody
}

type BatchSubmitReviewInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSubmitReviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSubmitReviewInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitReviewInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSubmitReviewInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSubmitReviewInfoResponse) GetBody() *BatchSubmitReviewInfoResponseBody {
	return s.Body
}

func (s *BatchSubmitReviewInfoResponse) SetHeaders(v map[string]*string) *BatchSubmitReviewInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchSubmitReviewInfoResponse) SetStatusCode(v int32) *BatchSubmitReviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSubmitReviewInfoResponse) SetBody(v *BatchSubmitReviewInfoResponseBody) *BatchSubmitReviewInfoResponse {
	s.Body = v
	return s
}

func (s *BatchSubmitReviewInfoResponse) Validate() error {
	return dara.Validate(s)
}
