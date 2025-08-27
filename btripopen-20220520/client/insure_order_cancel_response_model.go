// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderCancelResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderCancelResponseBody) *InsureOrderCancelResponse
	GetBody() *InsureOrderCancelResponseBody
}

type InsureOrderCancelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCancelResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderCancelResponse) GetBody() *InsureOrderCancelResponseBody {
	return s.Body
}

func (s *InsureOrderCancelResponse) SetHeaders(v map[string]*string) *InsureOrderCancelResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderCancelResponse) SetStatusCode(v int32) *InsureOrderCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderCancelResponse) SetBody(v *InsureOrderCancelResponseBody) *InsureOrderCancelResponse {
	s.Body = v
	return s
}

func (s *InsureOrderCancelResponse) Validate() error {
	return dara.Validate(s)
}
