// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDBsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDBsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDBsResponseBody) *DescribeDrdsDBsResponse
	GetBody() *DescribeDrdsDBsResponseBody
}

type DescribeDrdsDBsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDBsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDBsResponse) GetBody() *DescribeDrdsDBsResponseBody {
	return s.Body
}

func (s *DescribeDrdsDBsResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBsResponse) SetStatusCode(v int32) *DescribeDrdsDBsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBsResponse) SetBody(v *DescribeDrdsDBsResponseBody) *DescribeDrdsDBsResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDBsResponse) Validate() error {
	return dara.Validate(s)
}
