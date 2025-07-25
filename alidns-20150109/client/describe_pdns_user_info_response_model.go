// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsUserInfoResponseBody) *DescribePdnsUserInfoResponse
	GetBody() *DescribePdnsUserInfoResponseBody
}

type DescribePdnsUserInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsUserInfoResponse) GetBody() *DescribePdnsUserInfoResponseBody {
	return s.Body
}

func (s *DescribePdnsUserInfoResponse) SetHeaders(v map[string]*string) *DescribePdnsUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsUserInfoResponse) SetStatusCode(v int32) *DescribePdnsUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsUserInfoResponse) SetBody(v *DescribePdnsUserInfoResponseBody) *DescribePdnsUserInfoResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
