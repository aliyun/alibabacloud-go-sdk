// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePoolInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStoragePoolInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStoragePoolInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStoragePoolInfoResponseBody) *DescribeStoragePoolInfoResponse
	GetBody() *DescribeStoragePoolInfoResponseBody
}

type DescribeStoragePoolInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoragePoolInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoragePoolInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePoolInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoragePoolInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStoragePoolInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStoragePoolInfoResponse) GetBody() *DescribeStoragePoolInfoResponseBody {
	return s.Body
}

func (s *DescribeStoragePoolInfoResponse) SetHeaders(v map[string]*string) *DescribeStoragePoolInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoragePoolInfoResponse) SetStatusCode(v int32) *DescribeStoragePoolInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoragePoolInfoResponse) SetBody(v *DescribeStoragePoolInfoResponseBody) *DescribeStoragePoolInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeStoragePoolInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
