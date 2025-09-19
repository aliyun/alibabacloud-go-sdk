// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttestorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAttestorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAttestorsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAttestorsResponseBody) *DescribeAttestorsResponse
	GetBody() *DescribeAttestorsResponseBody
}

type DescribeAttestorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAttestorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAttestorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttestorsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttestorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAttestorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAttestorsResponse) GetBody() *DescribeAttestorsResponseBody {
	return s.Body
}

func (s *DescribeAttestorsResponse) SetHeaders(v map[string]*string) *DescribeAttestorsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttestorsResponse) SetStatusCode(v int32) *DescribeAttestorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttestorsResponse) SetBody(v *DescribeAttestorsResponseBody) *DescribeAttestorsResponse {
	s.Body = v
	return s
}

func (s *DescribeAttestorsResponse) Validate() error {
	return dara.Validate(s)
}
