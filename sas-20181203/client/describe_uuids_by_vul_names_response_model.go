// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUuidsByVulNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUuidsByVulNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUuidsByVulNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUuidsByVulNamesResponseBody) *DescribeUuidsByVulNamesResponse
	GetBody() *DescribeUuidsByVulNamesResponseBody
}

type DescribeUuidsByVulNamesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUuidsByVulNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUuidsByVulNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidsByVulNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUuidsByVulNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUuidsByVulNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUuidsByVulNamesResponse) GetBody() *DescribeUuidsByVulNamesResponseBody {
	return s.Body
}

func (s *DescribeUuidsByVulNamesResponse) SetHeaders(v map[string]*string) *DescribeUuidsByVulNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUuidsByVulNamesResponse) SetStatusCode(v int32) *DescribeUuidsByVulNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponse) SetBody(v *DescribeUuidsByVulNamesResponseBody) *DescribeUuidsByVulNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeUuidsByVulNamesResponse) Validate() error {
	return dara.Validate(s)
}
