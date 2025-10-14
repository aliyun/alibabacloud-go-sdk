// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShowStorageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeShowStorageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeShowStorageInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeShowStorageInfoResponseBody) *DescribeShowStorageInfoResponse
	GetBody() *DescribeShowStorageInfoResponseBody
}

type DescribeShowStorageInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeShowStorageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeShowStorageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowStorageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeShowStorageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeShowStorageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeShowStorageInfoResponse) GetBody() *DescribeShowStorageInfoResponseBody {
	return s.Body
}

func (s *DescribeShowStorageInfoResponse) SetHeaders(v map[string]*string) *DescribeShowStorageInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeShowStorageInfoResponse) SetStatusCode(v int32) *DescribeShowStorageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShowStorageInfoResponse) SetBody(v *DescribeShowStorageInfoResponseBody) *DescribeShowStorageInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeShowStorageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
