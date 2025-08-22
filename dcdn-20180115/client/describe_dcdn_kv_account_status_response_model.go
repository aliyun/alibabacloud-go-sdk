// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvAccountStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnKvAccountStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnKvAccountStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnKvAccountStatusResponseBody) *DescribeDcdnKvAccountStatusResponse
	GetBody() *DescribeDcdnKvAccountStatusResponseBody
}

type DescribeDcdnKvAccountStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnKvAccountStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnKvAccountStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvAccountStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvAccountStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnKvAccountStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnKvAccountStatusResponse) GetBody() *DescribeDcdnKvAccountStatusResponseBody {
	return s.Body
}

func (s *DescribeDcdnKvAccountStatusResponse) SetHeaders(v map[string]*string) *DescribeDcdnKvAccountStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnKvAccountStatusResponse) SetStatusCode(v int32) *DescribeDcdnKvAccountStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnKvAccountStatusResponse) SetBody(v *DescribeDcdnKvAccountStatusResponseBody) *DescribeDcdnKvAccountStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnKvAccountStatusResponse) Validate() error {
	return dara.Validate(s)
}
