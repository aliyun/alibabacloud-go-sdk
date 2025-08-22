// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafGroupsResponseBody) *DescribeDcdnWafGroupsResponse
	GetBody() *DescribeDcdnWafGroupsResponseBody
}

type DescribeDcdnWafGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafGroupsResponse) GetBody() *DescribeDcdnWafGroupsResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafGroupsResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafGroupsResponse) SetStatusCode(v int32) *DescribeDcdnWafGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponse) SetBody(v *DescribeDcdnWafGroupsResponseBody) *DescribeDcdnWafGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafGroupsResponse) Validate() error {
	return dara.Validate(s)
}
