// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyParameterLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModifyParameterLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModifyParameterLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModifyParameterLogResponseBody) *DescribeModifyParameterLogResponse
	GetBody() *DescribeModifyParameterLogResponseBody
}

type DescribeModifyParameterLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModifyParameterLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModifyParameterLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModifyParameterLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModifyParameterLogResponse) GetBody() *DescribeModifyParameterLogResponseBody {
	return s.Body
}

func (s *DescribeModifyParameterLogResponse) SetHeaders(v map[string]*string) *DescribeModifyParameterLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetStatusCode(v int32) *DescribeModifyParameterLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetBody(v *DescribeModifyParameterLogResponseBody) *DescribeModifyParameterLogResponse {
	s.Body = v
	return s
}

func (s *DescribeModifyParameterLogResponse) Validate() error {
	return dara.Validate(s)
}
