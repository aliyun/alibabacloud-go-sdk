// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCnameCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCnameCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCnameCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCnameCountResponseBody) *DescribeCnameCountResponse
	GetBody() *DescribeCnameCountResponseBody
}

type DescribeCnameCountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCnameCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCnameCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCnameCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCnameCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCnameCountResponse) GetBody() *DescribeCnameCountResponseBody {
	return s.Body
}

func (s *DescribeCnameCountResponse) SetHeaders(v map[string]*string) *DescribeCnameCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCnameCountResponse) SetStatusCode(v int32) *DescribeCnameCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCnameCountResponse) SetBody(v *DescribeCnameCountResponseBody) *DescribeCnameCountResponse {
	s.Body = v
	return s
}

func (s *DescribeCnameCountResponse) Validate() error {
	return dara.Validate(s)
}
