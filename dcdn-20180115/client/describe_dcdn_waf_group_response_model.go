// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafGroupResponseBody) *DescribeDcdnWafGroupResponse
	GetBody() *DescribeDcdnWafGroupResponseBody
}

type DescribeDcdnWafGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafGroupResponse) GetBody() *DescribeDcdnWafGroupResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafGroupResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafGroupResponse) SetStatusCode(v int32) *DescribeDcdnWafGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafGroupResponse) SetBody(v *DescribeDcdnWafGroupResponseBody) *DescribeDcdnWafGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafGroupResponse) Validate() error {
	return dara.Validate(s)
}
