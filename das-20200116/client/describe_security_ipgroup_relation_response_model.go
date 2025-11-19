// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityIPGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityIPGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityIPGroupRelationResponseBody) *DescribeSecurityIPGroupRelationResponse
	GetBody() *DescribeSecurityIPGroupRelationResponseBody
}

type DescribeSecurityIPGroupRelationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityIPGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityIPGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityIPGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityIPGroupRelationResponse) GetBody() *DescribeSecurityIPGroupRelationResponseBody {
	return s.Body
}

func (s *DescribeSecurityIPGroupRelationResponse) SetHeaders(v map[string]*string) *DescribeSecurityIPGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponse) SetStatusCode(v int32) *DescribeSecurityIPGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponse) SetBody(v *DescribeSecurityIPGroupRelationResponseBody) *DescribeSecurityIPGroupRelationResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
