// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSpecReviewTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSpecReviewTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSpecReviewTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelSpecReviewTaskResponseBody) *CancelSpecReviewTaskResponse
	GetBody() *CancelSpecReviewTaskResponseBody
}

type CancelSpecReviewTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSpecReviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSpecReviewTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSpecReviewTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelSpecReviewTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSpecReviewTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSpecReviewTaskResponse) GetBody() *CancelSpecReviewTaskResponseBody {
	return s.Body
}

func (s *CancelSpecReviewTaskResponse) SetHeaders(v map[string]*string) *CancelSpecReviewTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelSpecReviewTaskResponse) SetStatusCode(v int32) *CancelSpecReviewTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSpecReviewTaskResponse) SetBody(v *CancelSpecReviewTaskResponseBody) *CancelSpecReviewTaskResponse {
	s.Body = v
	return s
}

func (s *CancelSpecReviewTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
