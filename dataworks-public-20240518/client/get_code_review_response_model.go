// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCodeReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCodeReviewResponse
	GetStatusCode() *int32
	SetBody(v *GetCodeReviewResponseBody) *GetCodeReviewResponse
	GetBody() *GetCodeReviewResponseBody
}

type GetCodeReviewResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCodeReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCodeReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCodeReviewResponse) GoString() string {
	return s.String()
}

func (s *GetCodeReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCodeReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCodeReviewResponse) GetBody() *GetCodeReviewResponseBody {
	return s.Body
}

func (s *GetCodeReviewResponse) SetHeaders(v map[string]*string) *GetCodeReviewResponse {
	s.Headers = v
	return s
}

func (s *GetCodeReviewResponse) SetStatusCode(v int32) *GetCodeReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeReviewResponse) SetBody(v *GetCodeReviewResponseBody) *GetCodeReviewResponse {
	s.Body = v
	return s
}

func (s *GetCodeReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
