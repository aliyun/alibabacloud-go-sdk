// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderDetailResponseBody) *InsureOrderDetailResponse
	GetBody() *InsureOrderDetailResponseBody
}

type InsureOrderDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderDetailResponse) GetBody() *InsureOrderDetailResponseBody {
	return s.Body
}

func (s *InsureOrderDetailResponse) SetHeaders(v map[string]*string) *InsureOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderDetailResponse) SetStatusCode(v int32) *InsureOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderDetailResponse) SetBody(v *InsureOrderDetailResponseBody) *InsureOrderDetailResponse {
	s.Body = v
	return s
}

func (s *InsureOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
