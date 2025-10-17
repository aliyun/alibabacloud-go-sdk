// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbClusterAttributeZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDbClusterAttributeZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDbClusterAttributeZonalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDbClusterAttributeZonalResponseBody) *DescribeDbClusterAttributeZonalResponse
	GetBody() *DescribeDbClusterAttributeZonalResponseBody
}

type DescribeDbClusterAttributeZonalResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbClusterAttributeZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbClusterAttributeZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbClusterAttributeZonalResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbClusterAttributeZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDbClusterAttributeZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDbClusterAttributeZonalResponse) GetBody() *DescribeDbClusterAttributeZonalResponseBody {
	return s.Body
}

func (s *DescribeDbClusterAttributeZonalResponse) SetHeaders(v map[string]*string) *DescribeDbClusterAttributeZonalResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponse) SetStatusCode(v int32) *DescribeDbClusterAttributeZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponse) SetBody(v *DescribeDbClusterAttributeZonalResponseBody) *DescribeDbClusterAttributeZonalResponse {
	s.Body = v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
