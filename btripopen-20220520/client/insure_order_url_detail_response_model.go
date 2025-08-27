// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderUrlDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderUrlDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderUrlDetailResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderUrlDetailResponseBody) *InsureOrderUrlDetailResponse
	GetBody() *InsureOrderUrlDetailResponseBody
}

type InsureOrderUrlDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderUrlDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderUrlDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderUrlDetailResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderUrlDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderUrlDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderUrlDetailResponse) GetBody() *InsureOrderUrlDetailResponseBody {
	return s.Body
}

func (s *InsureOrderUrlDetailResponse) SetHeaders(v map[string]*string) *InsureOrderUrlDetailResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderUrlDetailResponse) SetStatusCode(v int32) *InsureOrderUrlDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderUrlDetailResponse) SetBody(v *InsureOrderUrlDetailResponseBody) *InsureOrderUrlDetailResponse {
	s.Body = v
	return s
}

func (s *InsureOrderUrlDetailResponse) Validate() error {
	return dara.Validate(s)
}
