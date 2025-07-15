// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse
	GetBody() *DescribeAppsResponseBody
}

type DescribeAppsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppsResponse) GetBody() *DescribeAppsResponseBody {
	return s.Body
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetStatusCode(v int32) *DescribeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

func (s *DescribeAppsResponse) Validate() error {
	return dara.Validate(s)
}
