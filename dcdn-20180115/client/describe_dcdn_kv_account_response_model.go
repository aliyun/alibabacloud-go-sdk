// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnKvAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnKvAccountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnKvAccountResponseBody) *DescribeDcdnKvAccountResponse
	GetBody() *DescribeDcdnKvAccountResponseBody
}

type DescribeDcdnKvAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnKvAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnKvAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvAccountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnKvAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnKvAccountResponse) GetBody() *DescribeDcdnKvAccountResponseBody {
	return s.Body
}

func (s *DescribeDcdnKvAccountResponse) SetHeaders(v map[string]*string) *DescribeDcdnKvAccountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnKvAccountResponse) SetStatusCode(v int32) *DescribeDcdnKvAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnKvAccountResponse) SetBody(v *DescribeDcdnKvAccountResponseBody) *DescribeDcdnKvAccountResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnKvAccountResponse) Validate() error {
	return dara.Validate(s)
}
