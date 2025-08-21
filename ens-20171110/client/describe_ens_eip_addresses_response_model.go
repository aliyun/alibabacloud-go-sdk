// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsEipAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsEipAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsEipAddressesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsEipAddressesResponseBody) *DescribeEnsEipAddressesResponse
	GetBody() *DescribeEnsEipAddressesResponseBody
}

type DescribeEnsEipAddressesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsEipAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsEipAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsEipAddressesResponse) GetBody() *DescribeEnsEipAddressesResponseBody {
	return s.Body
}

func (s *DescribeEnsEipAddressesResponse) SetHeaders(v map[string]*string) *DescribeEnsEipAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsEipAddressesResponse) SetStatusCode(v int32) *DescribeEnsEipAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsEipAddressesResponse) SetBody(v *DescribeEnsEipAddressesResponseBody) *DescribeEnsEipAddressesResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsEipAddressesResponse) Validate() error {
	return dara.Validate(s)
}
