// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceHAConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceHAConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceHAConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceHAConfigResponseBody) *DescribeDBInstanceHAConfigResponse
	GetBody() *DescribeDBInstanceHAConfigResponseBody
}

type DescribeDBInstanceHAConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceHAConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceHAConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceHAConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceHAConfigResponse) GetBody() *DescribeDBInstanceHAConfigResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceHAConfigResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceHAConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceHAConfigResponse) SetStatusCode(v int32) *DescribeDBInstanceHAConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponse) SetBody(v *DescribeDBInstanceHAConfigResponseBody) *DescribeDBInstanceHAConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceHAConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
