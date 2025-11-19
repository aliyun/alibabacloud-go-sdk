// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStorageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodStorageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodStorageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodStorageDataResponseBody) *DescribeVodStorageDataResponse
	GetBody() *DescribeVodStorageDataResponseBody
}

type DescribeVodStorageDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodStorageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodStorageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStorageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodStorageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodStorageDataResponse) GetBody() *DescribeVodStorageDataResponseBody {
	return s.Body
}

func (s *DescribeVodStorageDataResponse) SetHeaders(v map[string]*string) *DescribeVodStorageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodStorageDataResponse) SetStatusCode(v int32) *DescribeVodStorageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodStorageDataResponse) SetBody(v *DescribeVodStorageDataResponseBody) *DescribeVodStorageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodStorageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
