// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTieringStorageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodTieringStorageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodTieringStorageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodTieringStorageDataResponseBody) *DescribeVodTieringStorageDataResponse
	GetBody() *DescribeVodTieringStorageDataResponseBody
}

type DescribeVodTieringStorageDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodTieringStorageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodTieringStorageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodTieringStorageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodTieringStorageDataResponse) GetBody() *DescribeVodTieringStorageDataResponseBody {
	return s.Body
}

func (s *DescribeVodTieringStorageDataResponse) SetHeaders(v map[string]*string) *DescribeVodTieringStorageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTieringStorageDataResponse) SetStatusCode(v int32) *DescribeVodTieringStorageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponse) SetBody(v *DescribeVodTieringStorageDataResponseBody) *DescribeVodTieringStorageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodTieringStorageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
