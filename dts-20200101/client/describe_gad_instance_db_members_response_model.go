// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDbMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGadInstanceDbMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGadInstanceDbMembersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGadInstanceDbMembersResponseBody) *DescribeGadInstanceDbMembersResponse
	GetBody() *DescribeGadInstanceDbMembersResponseBody
}

type DescribeGadInstanceDbMembersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGadInstanceDbMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGadInstanceDbMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDbMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDbMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGadInstanceDbMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGadInstanceDbMembersResponse) GetBody() *DescribeGadInstanceDbMembersResponseBody {
	return s.Body
}

func (s *DescribeGadInstanceDbMembersResponse) SetHeaders(v map[string]*string) *DescribeGadInstanceDbMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeGadInstanceDbMembersResponse) SetStatusCode(v int32) *DescribeGadInstanceDbMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponse) SetBody(v *DescribeGadInstanceDbMembersResponseBody) *DescribeGadInstanceDbMembersResponse {
	s.Body = v
	return s
}

func (s *DescribeGadInstanceDbMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
