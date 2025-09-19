// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBuildRiskByKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageBuildRiskByKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageBuildRiskByKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageBuildRiskByKeyResponseBody) *DescribeImageBuildRiskByKeyResponse
	GetBody() *DescribeImageBuildRiskByKeyResponseBody
}

type DescribeImageBuildRiskByKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageBuildRiskByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageBuildRiskByKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskByKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskByKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageBuildRiskByKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageBuildRiskByKeyResponse) GetBody() *DescribeImageBuildRiskByKeyResponseBody {
	return s.Body
}

func (s *DescribeImageBuildRiskByKeyResponse) SetHeaders(v map[string]*string) *DescribeImageBuildRiskByKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponse) SetStatusCode(v int32) *DescribeImageBuildRiskByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponse) SetBody(v *DescribeImageBuildRiskByKeyResponseBody) *DescribeImageBuildRiskByKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponse) Validate() error {
	return dara.Validate(s)
}
