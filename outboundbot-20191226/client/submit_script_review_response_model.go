// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScriptReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitScriptReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitScriptReviewResponse
	GetStatusCode() *int32
	SetBody(v *SubmitScriptReviewResponseBody) *SubmitScriptReviewResponse
	GetBody() *SubmitScriptReviewResponseBody
}

type SubmitScriptReviewResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitScriptReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitScriptReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptReviewResponse) GoString() string {
	return s.String()
}

func (s *SubmitScriptReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitScriptReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitScriptReviewResponse) GetBody() *SubmitScriptReviewResponseBody {
	return s.Body
}

func (s *SubmitScriptReviewResponse) SetHeaders(v map[string]*string) *SubmitScriptReviewResponse {
	s.Headers = v
	return s
}

func (s *SubmitScriptReviewResponse) SetStatusCode(v int32) *SubmitScriptReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitScriptReviewResponse) SetBody(v *SubmitScriptReviewResponseBody) *SubmitScriptReviewResponse {
	s.Body = v
	return s
}

func (s *SubmitScriptReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
