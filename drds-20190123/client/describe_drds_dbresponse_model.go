// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDBResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDBResponseBody) *DescribeDrdsDBResponse
	GetBody() *DescribeDrdsDBResponseBody
}

type DescribeDrdsDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDBResponse) GetBody() *DescribeDrdsDBResponseBody {
	return s.Body
}

func (s *DescribeDrdsDBResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBResponse) SetStatusCode(v int32) *DescribeDrdsDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBResponse) SetBody(v *DescribeDrdsDBResponseBody) *DescribeDrdsDBResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDBResponse) Validate() error {
	return dara.Validate(s)
}
