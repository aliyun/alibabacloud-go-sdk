// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmInstanceResponseBody) *DescribeGtmInstanceResponse
	GetBody() *DescribeGtmInstanceResponseBody
}

type DescribeGtmInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmInstanceResponse) GetBody() *DescribeGtmInstanceResponseBody {
	return s.Body
}

func (s *DescribeGtmInstanceResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceResponse) SetStatusCode(v int32) *DescribeGtmInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmInstanceResponse) SetBody(v *DescribeGtmInstanceResponseBody) *DescribeGtmInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
