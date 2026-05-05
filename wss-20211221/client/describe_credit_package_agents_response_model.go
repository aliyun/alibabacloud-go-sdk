// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackageAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreditPackageAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreditPackageAgentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreditPackageAgentsResponseBody) *DescribeCreditPackageAgentsResponse
	GetBody() *DescribeCreditPackageAgentsResponseBody
}

type DescribeCreditPackageAgentsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreditPackageAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreditPackageAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageAgentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreditPackageAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreditPackageAgentsResponse) GetBody() *DescribeCreditPackageAgentsResponseBody {
	return s.Body
}

func (s *DescribeCreditPackageAgentsResponse) SetHeaders(v map[string]*string) *DescribeCreditPackageAgentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreditPackageAgentsResponse) SetStatusCode(v int32) *DescribeCreditPackageAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponse) SetBody(v *DescribeCreditPackageAgentsResponseBody) *DescribeCreditPackageAgentsResponse {
	s.Body = v
	return s
}

func (s *DescribeCreditPackageAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
