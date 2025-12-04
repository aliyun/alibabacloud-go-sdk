// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOSSStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOSSStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOSSStorageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOSSStorageResponseBody) *DescribeOSSStorageResponse
	GetBody() *DescribeOSSStorageResponseBody
}

type DescribeOSSStorageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOSSStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOSSStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOSSStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeOSSStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOSSStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOSSStorageResponse) GetBody() *DescribeOSSStorageResponseBody {
	return s.Body
}

func (s *DescribeOSSStorageResponse) SetHeaders(v map[string]*string) *DescribeOSSStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeOSSStorageResponse) SetStatusCode(v int32) *DescribeOSSStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOSSStorageResponse) SetBody(v *DescribeOSSStorageResponseBody) *DescribeOSSStorageResponse {
	s.Body = v
	return s
}

func (s *DescribeOSSStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
