// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSTrafficResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSTrafficResponseBody) *DescribeDDoSTrafficResponse
	GetBody() *DescribeDDoSTrafficResponseBody
}

type DescribeDDoSTrafficResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSTrafficResponse) GetBody() *DescribeDDoSTrafficResponseBody {
	return s.Body
}

func (s *DescribeDDoSTrafficResponse) SetHeaders(v map[string]*string) *DescribeDDoSTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetStatusCode(v int32) *DescribeDDoSTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetBody(v *DescribeDDoSTrafficResponseBody) *DescribeDDoSTrafficResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSTrafficResponse) Validate() error {
	return dara.Validate(s)
}
