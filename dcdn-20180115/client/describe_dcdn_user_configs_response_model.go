// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserConfigsResponseBody) *DescribeDcdnUserConfigsResponse
	GetBody() *DescribeDcdnUserConfigsResponseBody
}

type DescribeDcdnUserConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserConfigsResponse) GetBody() *DescribeDcdnUserConfigsResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserConfigsResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserConfigsResponse) SetStatusCode(v int32) *DescribeDcdnUserConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserConfigsResponse) SetBody(v *DescribeDcdnUserConfigsResponseBody) *DescribeDcdnUserConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserConfigsResponse) Validate() error {
	return dara.Validate(s)
}
