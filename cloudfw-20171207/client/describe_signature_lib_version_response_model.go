// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignatureLibVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSignatureLibVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSignatureLibVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSignatureLibVersionResponseBody) *DescribeSignatureLibVersionResponse
	GetBody() *DescribeSignatureLibVersionResponseBody
}

type DescribeSignatureLibVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSignatureLibVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSignatureLibVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignatureLibVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignatureLibVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSignatureLibVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSignatureLibVersionResponse) GetBody() *DescribeSignatureLibVersionResponseBody {
	return s.Body
}

func (s *DescribeSignatureLibVersionResponse) SetHeaders(v map[string]*string) *DescribeSignatureLibVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignatureLibVersionResponse) SetStatusCode(v int32) *DescribeSignatureLibVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSignatureLibVersionResponse) SetBody(v *DescribeSignatureLibVersionResponseBody) *DescribeSignatureLibVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeSignatureLibVersionResponse) Validate() error {
	return dara.Validate(s)
}
