// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundDetailListResponse
	GetStatusCode() *int32
	SetBody(v *RefundDetailListResponseBody) *RefundDetailListResponse
	GetBody() *RefundDetailListResponseBody
}

type RefundDetailListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailListResponse) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundDetailListResponse) GetBody() *RefundDetailListResponseBody {
	return s.Body
}

func (s *RefundDetailListResponse) SetHeaders(v map[string]*string) *RefundDetailListResponse {
	s.Headers = v
	return s
}

func (s *RefundDetailListResponse) SetStatusCode(v int32) *RefundDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundDetailListResponse) SetBody(v *RefundDetailListResponseBody) *RefundDetailListResponse {
	s.Body = v
	return s
}

func (s *RefundDetailListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
