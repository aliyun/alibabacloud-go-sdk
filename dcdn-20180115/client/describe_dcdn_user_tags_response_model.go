// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserTagsResponseBody) *DescribeDcdnUserTagsResponse
	GetBody() *DescribeDcdnUserTagsResponseBody
}

type DescribeDcdnUserTagsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserTagsResponse) GetBody() *DescribeDcdnUserTagsResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserTagsResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserTagsResponse) SetStatusCode(v int32) *DescribeDcdnUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserTagsResponse) SetBody(v *DescribeDcdnUserTagsResponseBody) *DescribeDcdnUserTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserTagsResponse) Validate() error {
	return dara.Validate(s)
}
