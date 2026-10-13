// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScriptPreReviewResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitScriptPreReviewResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitScriptPreReviewResultResponse
	GetStatusCode() *int32
	SetBody(v *SubmitScriptPreReviewResultResponseBody) *SubmitScriptPreReviewResultResponse
	GetBody() *SubmitScriptPreReviewResultResponseBody
}

type SubmitScriptPreReviewResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitScriptPreReviewResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitScriptPreReviewResultResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptPreReviewResultResponse) GoString() string {
	return s.String()
}

func (s *SubmitScriptPreReviewResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitScriptPreReviewResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitScriptPreReviewResultResponse) GetBody() *SubmitScriptPreReviewResultResponseBody {
	return s.Body
}

func (s *SubmitScriptPreReviewResultResponse) SetHeaders(v map[string]*string) *SubmitScriptPreReviewResultResponse {
	s.Headers = v
	return s
}

func (s *SubmitScriptPreReviewResultResponse) SetStatusCode(v int32) *SubmitScriptPreReviewResultResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponse) SetBody(v *SubmitScriptPreReviewResultResponseBody) *SubmitScriptPreReviewResultResponse {
	s.Body = v
	return s
}

func (s *SubmitScriptPreReviewResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
