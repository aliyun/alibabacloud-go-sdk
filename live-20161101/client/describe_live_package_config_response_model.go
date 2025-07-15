// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePackageConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePackageConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePackageConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePackageConfigResponseBody) *DescribeLivePackageConfigResponse
	GetBody() *DescribeLivePackageConfigResponseBody
}

type DescribeLivePackageConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePackageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePackageConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePackageConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePackageConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePackageConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePackageConfigResponse) GetBody() *DescribeLivePackageConfigResponseBody {
	return s.Body
}

func (s *DescribeLivePackageConfigResponse) SetHeaders(v map[string]*string) *DescribeLivePackageConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePackageConfigResponse) SetStatusCode(v int32) *DescribeLivePackageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePackageConfigResponse) SetBody(v *DescribeLivePackageConfigResponseBody) *DescribeLivePackageConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePackageConfigResponse) Validate() error {
	return dara.Validate(s)
}
