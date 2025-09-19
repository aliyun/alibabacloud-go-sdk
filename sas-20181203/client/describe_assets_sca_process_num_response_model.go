// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetsScaProcessNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetsScaProcessNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetsScaProcessNumResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetsScaProcessNumResponseBody) *DescribeAssetsScaProcessNumResponse
	GetBody() *DescribeAssetsScaProcessNumResponseBody
}

type DescribeAssetsScaProcessNumResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetsScaProcessNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetsScaProcessNumResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsScaProcessNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetsScaProcessNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetsScaProcessNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetsScaProcessNumResponse) GetBody() *DescribeAssetsScaProcessNumResponseBody {
	return s.Body
}

func (s *DescribeAssetsScaProcessNumResponse) SetHeaders(v map[string]*string) *DescribeAssetsScaProcessNumResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetsScaProcessNumResponse) SetStatusCode(v int32) *DescribeAssetsScaProcessNumResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetsScaProcessNumResponse) SetBody(v *DescribeAssetsScaProcessNumResponseBody) *DescribeAssetsScaProcessNumResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetsScaProcessNumResponse) Validate() error {
	return dara.Validate(s)
}
