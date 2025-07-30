// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResultToReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResultToReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResultToReviewResponse
	GetStatusCode() *int32
	SetBody(v *GetResultToReviewResponseBody) *GetResultToReviewResponse
	GetBody() *GetResultToReviewResponseBody
}

type GetResultToReviewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResultToReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResultToReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewResponse) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResultToReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResultToReviewResponse) GetBody() *GetResultToReviewResponseBody {
	return s.Body
}

func (s *GetResultToReviewResponse) SetHeaders(v map[string]*string) *GetResultToReviewResponse {
	s.Headers = v
	return s
}

func (s *GetResultToReviewResponse) SetStatusCode(v int32) *GetResultToReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultToReviewResponse) SetBody(v *GetResultToReviewResponseBody) *GetResultToReviewResponse {
	s.Body = v
	return s
}

func (s *GetResultToReviewResponse) Validate() error {
	return dara.Validate(s)
}
