// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCorePackageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCorePackageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCorePackageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCorePackageListResponseBody) *DescribeCorePackageListResponse
	GetBody() *DescribeCorePackageListResponseBody
}

type DescribeCorePackageListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCorePackageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCorePackageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCorePackageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCorePackageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCorePackageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCorePackageListResponse) GetBody() *DescribeCorePackageListResponseBody {
	return s.Body
}

func (s *DescribeCorePackageListResponse) SetHeaders(v map[string]*string) *DescribeCorePackageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCorePackageListResponse) SetStatusCode(v int32) *DescribeCorePackageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCorePackageListResponse) SetBody(v *DescribeCorePackageListResponseBody) *DescribeCorePackageListResponse {
	s.Body = v
	return s
}

func (s *DescribeCorePackageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
