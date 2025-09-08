// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertTypeResponseBody) *DescribeAlertTypeResponse
	GetBody() *DescribeAlertTypeResponseBody
}

type DescribeAlertTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertTypeResponse) GetBody() *DescribeAlertTypeResponseBody {
	return s.Body
}

func (s *DescribeAlertTypeResponse) SetHeaders(v map[string]*string) *DescribeAlertTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertTypeResponse) SetStatusCode(v int32) *DescribeAlertTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertTypeResponse) SetBody(v *DescribeAlertTypeResponseBody) *DescribeAlertTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertTypeResponse) Validate() error {
	return dara.Validate(s)
}
