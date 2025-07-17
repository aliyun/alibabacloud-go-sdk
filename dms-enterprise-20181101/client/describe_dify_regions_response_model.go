// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDifyRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDifyRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDifyRegionsResponseBody) *DescribeDifyRegionsResponse
	GetBody() *DescribeDifyRegionsResponseBody
}

type DescribeDifyRegionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDifyRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDifyRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDifyRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDifyRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDifyRegionsResponse) GetBody() *DescribeDifyRegionsResponseBody {
	return s.Body
}

func (s *DescribeDifyRegionsResponse) SetHeaders(v map[string]*string) *DescribeDifyRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDifyRegionsResponse) SetStatusCode(v int32) *DescribeDifyRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDifyRegionsResponse) SetBody(v *DescribeDifyRegionsResponseBody) *DescribeDifyRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDifyRegionsResponse) Validate() error {
	return dara.Validate(s)
}
