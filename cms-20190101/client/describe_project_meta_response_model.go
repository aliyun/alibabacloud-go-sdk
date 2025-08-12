// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectMetaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectMetaResponseBody) *DescribeProjectMetaResponse
	GetBody() *DescribeProjectMetaResponseBody
}

type DescribeProjectMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectMetaResponse) GetBody() *DescribeProjectMetaResponseBody {
	return s.Body
}

func (s *DescribeProjectMetaResponse) SetHeaders(v map[string]*string) *DescribeProjectMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectMetaResponse) SetStatusCode(v int32) *DescribeProjectMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectMetaResponse) SetBody(v *DescribeProjectMetaResponseBody) *DescribeProjectMetaResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectMetaResponse) Validate() error {
	return dara.Validate(s)
}
