// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnTagResourcesResponseBody) *DescribeDcdnTagResourcesResponse
	GetBody() *DescribeDcdnTagResourcesResponseBody
}

type DescribeDcdnTagResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnTagResourcesResponse) GetBody() *DescribeDcdnTagResourcesResponseBody {
	return s.Body
}

func (s *DescribeDcdnTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeDcdnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnTagResourcesResponse) SetStatusCode(v int32) *DescribeDcdnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponse) SetBody(v *DescribeDcdnTagResourcesResponseBody) *DescribeDcdnTagResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}
