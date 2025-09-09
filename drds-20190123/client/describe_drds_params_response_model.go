// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsParamsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsParamsResponseBody) *DescribeDrdsParamsResponse
	GetBody() *DescribeDrdsParamsResponseBody
}

type DescribeDrdsParamsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsParamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsParamsResponse) GetBody() *DescribeDrdsParamsResponseBody {
	return s.Body
}

func (s *DescribeDrdsParamsResponse) SetHeaders(v map[string]*string) *DescribeDrdsParamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsParamsResponse) SetStatusCode(v int32) *DescribeDrdsParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsParamsResponse) SetBody(v *DescribeDrdsParamsResponseBody) *DescribeDrdsParamsResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsParamsResponse) Validate() error {
	return dara.Validate(s)
}
