// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceSystemCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmInstanceSystemCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmInstanceSystemCnameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmInstanceSystemCnameResponseBody) *DescribeGtmInstanceSystemCnameResponse
	GetBody() *DescribeGtmInstanceSystemCnameResponseBody
}

type DescribeGtmInstanceSystemCnameResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmInstanceSystemCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmInstanceSystemCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceSystemCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceSystemCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmInstanceSystemCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmInstanceSystemCnameResponse) GetBody() *DescribeGtmInstanceSystemCnameResponseBody {
	return s.Body
}

func (s *DescribeGtmInstanceSystemCnameResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceSystemCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponse) SetStatusCode(v int32) *DescribeGtmInstanceSystemCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponse) SetBody(v *DescribeGtmInstanceSystemCnameResponseBody) *DescribeGtmInstanceSystemCnameResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
