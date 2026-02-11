// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageResponseBody) *DescribeStorageResponse
	GetBody() *DescribeStorageResponseBody
}

type DescribeStorageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageResponse) GetBody() *DescribeStorageResponseBody {
	return s.Body
}

func (s *DescribeStorageResponse) SetHeaders(v map[string]*string) *DescribeStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageResponse) SetStatusCode(v int32) *DescribeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageResponse) SetBody(v *DescribeStorageResponseBody) *DescribeStorageResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
