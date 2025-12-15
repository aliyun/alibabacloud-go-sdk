// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstanceConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDbInstanceConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDbInstanceConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDbInstanceConnectivityResponseBody) *DescribeDbInstanceConnectivityResponse
	GetBody() *DescribeDbInstanceConnectivityResponseBody
}

type DescribeDbInstanceConnectivityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbInstanceConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbInstanceConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceConnectivityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDbInstanceConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDbInstanceConnectivityResponse) GetBody() *DescribeDbInstanceConnectivityResponseBody {
	return s.Body
}

func (s *DescribeDbInstanceConnectivityResponse) SetHeaders(v map[string]*string) *DescribeDbInstanceConnectivityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstanceConnectivityResponse) SetStatusCode(v int32) *DescribeDbInstanceConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstanceConnectivityResponse) SetBody(v *DescribeDbInstanceConnectivityResponseBody) *DescribeDbInstanceConnectivityResponse {
	s.Body = v
	return s
}

func (s *DescribeDbInstanceConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
