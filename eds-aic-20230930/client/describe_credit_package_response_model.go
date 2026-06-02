// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreditPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreditPackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreditPackageResponseBody) *DescribeCreditPackageResponse
	GetBody() *DescribeCreditPackageResponseBody
}

type DescribeCreditPackageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreditPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreditPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreditPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreditPackageResponse) GetBody() *DescribeCreditPackageResponseBody {
	return s.Body
}

func (s *DescribeCreditPackageResponse) SetHeaders(v map[string]*string) *DescribeCreditPackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreditPackageResponse) SetStatusCode(v int32) *DescribeCreditPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreditPackageResponse) SetBody(v *DescribeCreditPackageResponseBody) *DescribeCreditPackageResponse {
	s.Body = v
	return s
}

func (s *DescribeCreditPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
