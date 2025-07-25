// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsAppKeyResponseBody) *DescribePdnsAppKeyResponse
	GetBody() *DescribePdnsAppKeyResponseBody
}

type DescribePdnsAppKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsAppKeyResponse) GetBody() *DescribePdnsAppKeyResponseBody {
	return s.Body
}

func (s *DescribePdnsAppKeyResponse) SetHeaders(v map[string]*string) *DescribePdnsAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsAppKeyResponse) SetStatusCode(v int32) *DescribePdnsAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsAppKeyResponse) SetBody(v *DescribePdnsAppKeyResponseBody) *DescribePdnsAppKeyResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsAppKeyResponse) Validate() error {
	return dara.Validate(s)
}
