// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdminInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdminInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdminInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdminInfoResponseBody) *DescribeAdminInfoResponse
	GetBody() *DescribeAdminInfoResponseBody
}

type DescribeAdminInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdminInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdminInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdminInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdminInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdminInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdminInfoResponse) GetBody() *DescribeAdminInfoResponseBody {
	return s.Body
}

func (s *DescribeAdminInfoResponse) SetHeaders(v map[string]*string) *DescribeAdminInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdminInfoResponse) SetStatusCode(v int32) *DescribeAdminInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdminInfoResponse) SetBody(v *DescribeAdminInfoResponseBody) *DescribeAdminInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeAdminInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
