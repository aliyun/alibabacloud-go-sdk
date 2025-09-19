// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentlessSensitiveFileByKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgentlessSensitiveFileByKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgentlessSensitiveFileByKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgentlessSensitiveFileByKeyResponseBody) *DescribeAgentlessSensitiveFileByKeyResponse
	GetBody() *DescribeAgentlessSensitiveFileByKeyResponseBody
}

type DescribeAgentlessSensitiveFileByKeyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgentlessSensitiveFileByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgentlessSensitiveFileByKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentlessSensitiveFileByKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) GetBody() *DescribeAgentlessSensitiveFileByKeyResponseBody {
	return s.Body
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) SetHeaders(v map[string]*string) *DescribeAgentlessSensitiveFileByKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) SetStatusCode(v int32) *DescribeAgentlessSensitiveFileByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) SetBody(v *DescribeAgentlessSensitiveFileByKeyResponseBody) *DescribeAgentlessSensitiveFileByKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponse) Validate() error {
	return dara.Validate(s)
}
