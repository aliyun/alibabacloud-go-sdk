// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreExistStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogStoreExistStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogStoreExistStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogStoreExistStatusResponseBody) *DescribeLogStoreExistStatusResponse
	GetBody() *DescribeLogStoreExistStatusResponseBody
}

type DescribeLogStoreExistStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogStoreExistStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogStoreExistStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogStoreExistStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogStoreExistStatusResponse) GetBody() *DescribeLogStoreExistStatusResponseBody {
	return s.Body
}

func (s *DescribeLogStoreExistStatusResponse) SetHeaders(v map[string]*string) *DescribeLogStoreExistStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) SetStatusCode(v int32) *DescribeLogStoreExistStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) SetBody(v *DescribeLogStoreExistStatusResponseBody) *DescribeLogStoreExistStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
