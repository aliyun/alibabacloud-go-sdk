// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundDetailResponse
	GetStatusCode() *int32
	SetBody(v *RefundDetailResponseBody) *RefundDetailResponse
	GetBody() *RefundDetailResponseBody
}

type RefundDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponse) GoString() string {
	return s.String()
}

func (s *RefundDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundDetailResponse) GetBody() *RefundDetailResponseBody {
	return s.Body
}

func (s *RefundDetailResponse) SetHeaders(v map[string]*string) *RefundDetailResponse {
	s.Headers = v
	return s
}

func (s *RefundDetailResponse) SetStatusCode(v int32) *RefundDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundDetailResponse) SetBody(v *RefundDetailResponseBody) *RefundDetailResponse {
	s.Body = v
	return s
}

func (s *RefundDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
