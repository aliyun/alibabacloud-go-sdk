// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientUserDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientUserDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientUserDNSResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientUserDNSResponseBody) *DescribeClientUserDNSResponse
	GetBody() *DescribeClientUserDNSResponseBody
}

type DescribeClientUserDNSResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientUserDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientUserDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientUserDNSResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientUserDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientUserDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientUserDNSResponse) GetBody() *DescribeClientUserDNSResponseBody {
	return s.Body
}

func (s *DescribeClientUserDNSResponse) SetHeaders(v map[string]*string) *DescribeClientUserDNSResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientUserDNSResponse) SetStatusCode(v int32) *DescribeClientUserDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientUserDNSResponse) SetBody(v *DescribeClientUserDNSResponseBody) *DescribeClientUserDNSResponse {
	s.Body = v
	return s
}

func (s *DescribeClientUserDNSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
