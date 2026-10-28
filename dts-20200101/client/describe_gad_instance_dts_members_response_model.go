// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDtsMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGadInstanceDtsMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGadInstanceDtsMembersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGadInstanceDtsMembersResponseBody) *DescribeGadInstanceDtsMembersResponse
	GetBody() *DescribeGadInstanceDtsMembersResponseBody
}

type DescribeGadInstanceDtsMembersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGadInstanceDtsMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGadInstanceDtsMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGadInstanceDtsMembersResponse) GetBody() *DescribeGadInstanceDtsMembersResponseBody {
	return s.Body
}

func (s *DescribeGadInstanceDtsMembersResponse) SetHeaders(v map[string]*string) *DescribeGadInstanceDtsMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponse) SetStatusCode(v int32) *DescribeGadInstanceDtsMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponse) SetBody(v *DescribeGadInstanceDtsMembersResponseBody) *DescribeGadInstanceDtsMembersResponse {
	s.Body = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
