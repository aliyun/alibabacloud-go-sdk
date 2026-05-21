// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSlaCouponApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSlaCouponApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSlaCouponApplyResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSlaCouponApplyResponseBody) *SubmitSlaCouponApplyResponse
	GetBody() *SubmitSlaCouponApplyResponseBody
}

type SubmitSlaCouponApplyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSlaCouponApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSlaCouponApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSlaCouponApplyResponse) GoString() string {
	return s.String()
}

func (s *SubmitSlaCouponApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSlaCouponApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSlaCouponApplyResponse) GetBody() *SubmitSlaCouponApplyResponseBody {
	return s.Body
}

func (s *SubmitSlaCouponApplyResponse) SetHeaders(v map[string]*string) *SubmitSlaCouponApplyResponse {
	s.Headers = v
	return s
}

func (s *SubmitSlaCouponApplyResponse) SetStatusCode(v int32) *SubmitSlaCouponApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSlaCouponApplyResponse) SetBody(v *SubmitSlaCouponApplyResponseBody) *SubmitSlaCouponApplyResponse {
	s.Body = v
	return s
}

func (s *SubmitSlaCouponApplyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
