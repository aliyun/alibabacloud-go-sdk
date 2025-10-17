// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLicenseOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLicenseOrdersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLicenseOrdersResponseBody) *DescribeLicenseOrdersResponse
	GetBody() *DescribeLicenseOrdersResponseBody
}

type DescribeLicenseOrdersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLicenseOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLicenseOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrdersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLicenseOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLicenseOrdersResponse) GetBody() *DescribeLicenseOrdersResponseBody {
	return s.Body
}

func (s *DescribeLicenseOrdersResponse) SetHeaders(v map[string]*string) *DescribeLicenseOrdersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLicenseOrdersResponse) SetStatusCode(v int32) *DescribeLicenseOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLicenseOrdersResponse) SetBody(v *DescribeLicenseOrdersResponseBody) *DescribeLicenseOrdersResponse {
	s.Body = v
	return s
}

func (s *DescribeLicenseOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
