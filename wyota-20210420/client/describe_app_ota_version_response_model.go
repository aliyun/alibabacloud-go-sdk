// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppOtaVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppOtaVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppOtaVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppOtaVersionResponseBody) *DescribeAppOtaVersionResponse
	GetBody() *DescribeAppOtaVersionResponseBody
}

type DescribeAppOtaVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppOtaVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppOtaVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppOtaVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppOtaVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppOtaVersionResponse) GetBody() *DescribeAppOtaVersionResponseBody {
	return s.Body
}

func (s *DescribeAppOtaVersionResponse) SetHeaders(v map[string]*string) *DescribeAppOtaVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppOtaVersionResponse) SetStatusCode(v int32) *DescribeAppOtaVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppOtaVersionResponse) SetBody(v *DescribeAppOtaVersionResponseBody) *DescribeAppOtaVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeAppOtaVersionResponse) Validate() error {
	return dara.Validate(s)
}
