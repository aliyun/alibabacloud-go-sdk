// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsSlowSqlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsSlowSqlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsSlowSqlsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsSlowSqlsResponseBody) *DescribeDrdsSlowSqlsResponse
	GetBody() *DescribeDrdsSlowSqlsResponseBody
}

type DescribeDrdsSlowSqlsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsSlowSqlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsSlowSqlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsSlowSqlsResponse) GetBody() *DescribeDrdsSlowSqlsResponseBody {
	return s.Body
}

func (s *DescribeDrdsSlowSqlsResponse) SetHeaders(v map[string]*string) *DescribeDrdsSlowSqlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponse) SetStatusCode(v int32) *DescribeDrdsSlowSqlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponse) SetBody(v *DescribeDrdsSlowSqlsResponseBody) *DescribeDrdsSlowSqlsResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponse) Validate() error {
	return dara.Validate(s)
}
