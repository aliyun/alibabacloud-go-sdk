// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageResourceUsageResponseBody) *DescribeStorageResourceUsageResponse
	GetBody() *DescribeStorageResourceUsageResponseBody
}

type DescribeStorageResourceUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageResourceUsageResponse) GetBody() *DescribeStorageResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeStorageResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeStorageResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageResourceUsageResponse) SetStatusCode(v int32) *DescribeStorageResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageResourceUsageResponse) SetBody(v *DescribeStorageResourceUsageResponseBody) *DescribeStorageResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
