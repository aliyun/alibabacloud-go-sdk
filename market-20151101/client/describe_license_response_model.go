// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLicenseResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLicenseResponseBody) *DescribeLicenseResponse
	GetBody() *DescribeLicenseResponseBody
}

type DescribeLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseResponse) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLicenseResponse) GetBody() *DescribeLicenseResponseBody {
	return s.Body
}

func (s *DescribeLicenseResponse) SetHeaders(v map[string]*string) *DescribeLicenseResponse {
	s.Headers = v
	return s
}

func (s *DescribeLicenseResponse) SetStatusCode(v int32) *DescribeLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLicenseResponse) SetBody(v *DescribeLicenseResponseBody) *DescribeLicenseResponse {
	s.Body = v
	return s
}

func (s *DescribeLicenseResponse) Validate() error {
	return dara.Validate(s)
}
