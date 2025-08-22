// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnAclFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnAclFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnAclFieldsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnAclFieldsResponseBody) *DescribeDcdnAclFieldsResponse
	GetBody() *DescribeDcdnAclFieldsResponseBody
}

type DescribeDcdnAclFieldsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnAclFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnAclFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnAclFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnAclFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnAclFieldsResponse) GetBody() *DescribeDcdnAclFieldsResponseBody {
	return s.Body
}

func (s *DescribeDcdnAclFieldsResponse) SetHeaders(v map[string]*string) *DescribeDcdnAclFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnAclFieldsResponse) SetStatusCode(v int32) *DescribeDcdnAclFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnAclFieldsResponse) SetBody(v *DescribeDcdnAclFieldsResponseBody) *DescribeDcdnAclFieldsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnAclFieldsResponse) Validate() error {
	return dara.Validate(s)
}
