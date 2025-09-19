// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMatchedMaliciousNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMatchedMaliciousNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMatchedMaliciousNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMatchedMaliciousNamesResponseBody) *DescribeMatchedMaliciousNamesResponse
	GetBody() *DescribeMatchedMaliciousNamesResponseBody
}

type DescribeMatchedMaliciousNamesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMatchedMaliciousNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMatchedMaliciousNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMatchedMaliciousNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMatchedMaliciousNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMatchedMaliciousNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMatchedMaliciousNamesResponse) GetBody() *DescribeMatchedMaliciousNamesResponseBody {
	return s.Body
}

func (s *DescribeMatchedMaliciousNamesResponse) SetHeaders(v map[string]*string) *DescribeMatchedMaliciousNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponse) SetStatusCode(v int32) *DescribeMatchedMaliciousNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponse) SetBody(v *DescribeMatchedMaliciousNamesResponseBody) *DescribeMatchedMaliciousNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponse) Validate() error {
	return dara.Validate(s)
}
