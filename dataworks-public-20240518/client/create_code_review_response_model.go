// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCodeReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCodeReviewResponse
	GetStatusCode() *int32
	SetBody(v *CreateCodeReviewResponseBody) *CreateCodeReviewResponse
	GetBody() *CreateCodeReviewResponseBody
}

type CreateCodeReviewResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCodeReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCodeReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeReviewResponse) GoString() string {
	return s.String()
}

func (s *CreateCodeReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCodeReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCodeReviewResponse) GetBody() *CreateCodeReviewResponseBody {
	return s.Body
}

func (s *CreateCodeReviewResponse) SetHeaders(v map[string]*string) *CreateCodeReviewResponse {
	s.Headers = v
	return s
}

func (s *CreateCodeReviewResponse) SetStatusCode(v int32) *CreateCodeReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCodeReviewResponse) SetBody(v *CreateCodeReviewResponseBody) *CreateCodeReviewResponse {
	s.Body = v
	return s
}

func (s *CreateCodeReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
