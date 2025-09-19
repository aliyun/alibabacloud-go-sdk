// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogstoreStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogstoreStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogstoreStorageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogstoreStorageResponseBody) *DescribeLogstoreStorageResponse
	GetBody() *DescribeLogstoreStorageResponseBody
}

type DescribeLogstoreStorageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogstoreStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogstoreStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstoreStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogstoreStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogstoreStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogstoreStorageResponse) GetBody() *DescribeLogstoreStorageResponseBody {
	return s.Body
}

func (s *DescribeLogstoreStorageResponse) SetHeaders(v map[string]*string) *DescribeLogstoreStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogstoreStorageResponse) SetStatusCode(v int32) *DescribeLogstoreStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogstoreStorageResponse) SetBody(v *DescribeLogstoreStorageResponseBody) *DescribeLogstoreStorageResponse {
	s.Body = v
	return s
}

func (s *DescribeLogstoreStorageResponse) Validate() error {
	return dara.Validate(s)
}
