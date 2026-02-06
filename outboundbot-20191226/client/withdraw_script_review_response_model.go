// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawScriptReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawScriptReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawScriptReviewResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawScriptReviewResponseBody) *WithdrawScriptReviewResponse
	GetBody() *WithdrawScriptReviewResponseBody
}

type WithdrawScriptReviewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawScriptReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawScriptReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawScriptReviewResponse) GoString() string {
	return s.String()
}

func (s *WithdrawScriptReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawScriptReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawScriptReviewResponse) GetBody() *WithdrawScriptReviewResponseBody {
	return s.Body
}

func (s *WithdrawScriptReviewResponse) SetHeaders(v map[string]*string) *WithdrawScriptReviewResponse {
	s.Headers = v
	return s
}

func (s *WithdrawScriptReviewResponse) SetStatusCode(v int32) *WithdrawScriptReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawScriptReviewResponse) SetBody(v *WithdrawScriptReviewResponseBody) *WithdrawScriptReviewResponse {
	s.Body = v
	return s
}

func (s *WithdrawScriptReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
