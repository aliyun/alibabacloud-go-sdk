// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackendInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackendInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackendInfoResponseBody) *DescribeBackendInfoResponse
	GetBody() *DescribeBackendInfoResponseBody
}

type DescribeBackendInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackendInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackendInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackendInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackendInfoResponse) GetBody() *DescribeBackendInfoResponseBody {
	return s.Body
}

func (s *DescribeBackendInfoResponse) SetHeaders(v map[string]*string) *DescribeBackendInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackendInfoResponse) SetStatusCode(v int32) *DescribeBackendInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackendInfoResponse) SetBody(v *DescribeBackendInfoResponseBody) *DescribeBackendInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeBackendInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
