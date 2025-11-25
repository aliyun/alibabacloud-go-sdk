// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpsPrivateAssocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpsPrivateAssocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpsPrivateAssocResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpsPrivateAssocResponseBody) *DescribeIpsPrivateAssocResponse
	GetBody() *DescribeIpsPrivateAssocResponseBody
}

type DescribeIpsPrivateAssocResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpsPrivateAssocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpsPrivateAssocResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpsPrivateAssocResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpsPrivateAssocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpsPrivateAssocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpsPrivateAssocResponse) GetBody() *DescribeIpsPrivateAssocResponseBody {
	return s.Body
}

func (s *DescribeIpsPrivateAssocResponse) SetHeaders(v map[string]*string) *DescribeIpsPrivateAssocResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpsPrivateAssocResponse) SetStatusCode(v int32) *DescribeIpsPrivateAssocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponse) SetBody(v *DescribeIpsPrivateAssocResponseBody) *DescribeIpsPrivateAssocResponse {
	s.Body = v
	return s
}

func (s *DescribeIpsPrivateAssocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
