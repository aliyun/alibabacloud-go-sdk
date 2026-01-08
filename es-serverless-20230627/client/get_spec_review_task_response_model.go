// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecReviewTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpecReviewTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpecReviewTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetSpecReviewTaskResponseBody) *GetSpecReviewTaskResponse
	GetBody() *GetSpecReviewTaskResponseBody
}

type GetSpecReviewTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpecReviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpecReviewTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpecReviewTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSpecReviewTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpecReviewTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpecReviewTaskResponse) GetBody() *GetSpecReviewTaskResponseBody {
	return s.Body
}

func (s *GetSpecReviewTaskResponse) SetHeaders(v map[string]*string) *GetSpecReviewTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSpecReviewTaskResponse) SetStatusCode(v int32) *GetSpecReviewTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpecReviewTaskResponse) SetBody(v *GetSpecReviewTaskResponseBody) *GetSpecReviewTaskResponse {
	s.Body = v
	return s
}

func (s *GetSpecReviewTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
