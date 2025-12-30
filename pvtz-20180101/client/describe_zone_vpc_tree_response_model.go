// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneVpcTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZoneVpcTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZoneVpcTreeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZoneVpcTreeResponseBody) *DescribeZoneVpcTreeResponse
	GetBody() *DescribeZoneVpcTreeResponseBody
}

type DescribeZoneVpcTreeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneVpcTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneVpcTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZoneVpcTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZoneVpcTreeResponse) GetBody() *DescribeZoneVpcTreeResponseBody {
	return s.Body
}

func (s *DescribeZoneVpcTreeResponse) SetHeaders(v map[string]*string) *DescribeZoneVpcTreeResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneVpcTreeResponse) SetStatusCode(v int32) *DescribeZoneVpcTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneVpcTreeResponse) SetBody(v *DescribeZoneVpcTreeResponseBody) *DescribeZoneVpcTreeResponse {
	s.Body = v
	return s
}

func (s *DescribeZoneVpcTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
