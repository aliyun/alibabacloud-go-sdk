// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDbInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDbInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDbInstanceResponseBody) *DescribeDrdsDbInstanceResponse
	GetBody() *DescribeDrdsDbInstanceResponseBody
}

type DescribeDrdsDbInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDbInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDbInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDbInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDbInstanceResponse) GetBody() *DescribeDrdsDbInstanceResponseBody {
	return s.Body
}

func (s *DescribeDrdsDbInstanceResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) SetStatusCode(v int32) *DescribeDrdsDbInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) SetBody(v *DescribeDrdsDbInstanceResponseBody) *DescribeDrdsDbInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDbInstanceResponse) Validate() error {
	return dara.Validate(s)
}
