// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceOwnerUidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceOwnerUidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceOwnerUidResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceOwnerUidResponseBody) *DescribeDefenseResourceOwnerUidResponse
	GetBody() *DescribeDefenseResourceOwnerUidResponseBody
}

type DescribeDefenseResourceOwnerUidResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceOwnerUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceOwnerUidResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceOwnerUidResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceOwnerUidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceOwnerUidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceOwnerUidResponse) GetBody() *DescribeDefenseResourceOwnerUidResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceOwnerUidResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceOwnerUidResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponse) SetStatusCode(v int32) *DescribeDefenseResourceOwnerUidResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponse) SetBody(v *DescribeDefenseResourceOwnerUidResponseBody) *DescribeDefenseResourceOwnerUidResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponse) Validate() error {
	return dara.Validate(s)
}
