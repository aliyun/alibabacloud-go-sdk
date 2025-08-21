// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosCreditResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosCreditResponseBody) *DescribeDdosCreditResponse
	GetBody() *DescribeDdosCreditResponseBody
}

type DescribeDdosCreditResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCreditResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosCreditResponse) GetBody() *DescribeDdosCreditResponseBody {
	return s.Body
}

func (s *DescribeDdosCreditResponse) SetHeaders(v map[string]*string) *DescribeDdosCreditResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosCreditResponse) SetStatusCode(v int32) *DescribeDdosCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosCreditResponse) SetBody(v *DescribeDdosCreditResponseBody) *DescribeDdosCreditResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosCreditResponse) Validate() error {
	return dara.Validate(s)
}
