// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsInstanceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsInstanceVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsInstanceVersionResponseBody) *DescribeDrdsInstanceVersionResponse
	GetBody() *DescribeDrdsInstanceVersionResponseBody
}

type DescribeDrdsInstanceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsInstanceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsInstanceVersionResponse) GetBody() *DescribeDrdsInstanceVersionResponseBody {
	return s.Body
}

func (s *DescribeDrdsInstanceVersionResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) SetStatusCode(v int32) *DescribeDrdsInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) SetBody(v *DescribeDrdsInstanceVersionResponseBody) *DescribeDrdsInstanceVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponse) Validate() error {
	return dara.Validate(s)
}
