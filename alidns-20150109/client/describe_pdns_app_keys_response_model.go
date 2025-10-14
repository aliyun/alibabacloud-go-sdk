// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAppKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsAppKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsAppKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsAppKeysResponseBody) *DescribePdnsAppKeysResponse
	GetBody() *DescribePdnsAppKeysResponseBody
}

type DescribePdnsAppKeysResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsAppKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsAppKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsAppKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsAppKeysResponse) GetBody() *DescribePdnsAppKeysResponseBody {
	return s.Body
}

func (s *DescribePdnsAppKeysResponse) SetHeaders(v map[string]*string) *DescribePdnsAppKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsAppKeysResponse) SetStatusCode(v int32) *DescribePdnsAppKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsAppKeysResponse) SetBody(v *DescribePdnsAppKeysResponseBody) *DescribePdnsAppKeysResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsAppKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
