// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsInstanceResponseBody) *DescribeDrdsInstanceResponse
	GetBody() *DescribeDrdsInstanceResponseBody
}

type DescribeDrdsInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsInstanceResponse) GetBody() *DescribeDrdsInstanceResponseBody {
	return s.Body
}

func (s *DescribeDrdsInstanceResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceResponse) SetStatusCode(v int32) *DescribeDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceResponse) SetBody(v *DescribeDrdsInstanceResponseBody) *DescribeDrdsInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
