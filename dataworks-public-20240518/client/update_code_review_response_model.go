// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCodeReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCodeReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCodeReviewResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCodeReviewResponseBody) *UpdateCodeReviewResponse
	GetBody() *UpdateCodeReviewResponseBody
}

type UpdateCodeReviewResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCodeReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCodeReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCodeReviewResponse) GoString() string {
	return s.String()
}

func (s *UpdateCodeReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCodeReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCodeReviewResponse) GetBody() *UpdateCodeReviewResponseBody {
	return s.Body
}

func (s *UpdateCodeReviewResponse) SetHeaders(v map[string]*string) *UpdateCodeReviewResponse {
	s.Headers = v
	return s
}

func (s *UpdateCodeReviewResponse) SetStatusCode(v int32) *UpdateCodeReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCodeReviewResponse) SetBody(v *UpdateCodeReviewResponseBody) *UpdateCodeReviewResponse {
	s.Body = v
	return s
}

func (s *UpdateCodeReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
