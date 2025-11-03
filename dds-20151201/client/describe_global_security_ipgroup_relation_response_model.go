// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalSecurityIPGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalSecurityIPGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalSecurityIPGroupRelationResponseBody) *DescribeGlobalSecurityIPGroupRelationResponse
	GetBody() *DescribeGlobalSecurityIPGroupRelationResponseBody
}

type DescribeGlobalSecurityIPGroupRelationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalSecurityIPGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) GetBody() *DescribeGlobalSecurityIPGroupRelationResponseBody {
	return s.Body
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) SetHeaders(v map[string]*string) *DescribeGlobalSecurityIPGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) SetStatusCode(v int32) *DescribeGlobalSecurityIPGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) SetBody(v *DescribeGlobalSecurityIPGroupRelationResponseBody) *DescribeGlobalSecurityIPGroupRelationResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
