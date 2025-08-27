// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderRefundResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderRefundResponseBody) *InsureOrderRefundResponse
	GetBody() *InsureOrderRefundResponseBody
}

type InsureOrderRefundResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderRefundResponse) GetBody() *InsureOrderRefundResponseBody {
	return s.Body
}

func (s *InsureOrderRefundResponse) SetHeaders(v map[string]*string) *InsureOrderRefundResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderRefundResponse) SetStatusCode(v int32) *InsureOrderRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderRefundResponse) SetBody(v *InsureOrderRefundResponseBody) *InsureOrderRefundResponse {
	s.Body = v
	return s
}

func (s *InsureOrderRefundResponse) Validate() error {
	return dara.Validate(s)
}
