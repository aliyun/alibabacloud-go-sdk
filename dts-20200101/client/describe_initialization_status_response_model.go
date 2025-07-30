// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInitializationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInitializationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInitializationStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInitializationStatusResponseBody) *DescribeInitializationStatusResponse
	GetBody() *DescribeInitializationStatusResponseBody
}

type DescribeInitializationStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInitializationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInitializationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInitializationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInitializationStatusResponse) GetBody() *DescribeInitializationStatusResponseBody {
	return s.Body
}

func (s *DescribeInitializationStatusResponse) SetHeaders(v map[string]*string) *DescribeInitializationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInitializationStatusResponse) SetStatusCode(v int32) *DescribeInitializationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInitializationStatusResponse) SetBody(v *DescribeInitializationStatusResponseBody) *DescribeInitializationStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeInitializationStatusResponse) Validate() error {
	return dara.Validate(s)
}
