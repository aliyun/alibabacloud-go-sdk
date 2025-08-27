// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureRefundDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureRefundDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureRefundDetailResponse
	GetStatusCode() *int32
	SetBody(v *InsureRefundDetailResponseBody) *InsureRefundDetailResponse
	GetBody() *InsureRefundDetailResponseBody
}

type InsureRefundDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureRefundDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureRefundDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponse) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureRefundDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureRefundDetailResponse) GetBody() *InsureRefundDetailResponseBody {
	return s.Body
}

func (s *InsureRefundDetailResponse) SetHeaders(v map[string]*string) *InsureRefundDetailResponse {
	s.Headers = v
	return s
}

func (s *InsureRefundDetailResponse) SetStatusCode(v int32) *InsureRefundDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureRefundDetailResponse) SetBody(v *InsureRefundDetailResponseBody) *InsureRefundDetailResponse {
	s.Body = v
	return s
}

func (s *InsureRefundDetailResponse) Validate() error {
	return dara.Validate(s)
}
