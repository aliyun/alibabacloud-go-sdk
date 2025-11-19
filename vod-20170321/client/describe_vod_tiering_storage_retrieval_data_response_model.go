// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTieringStorageRetrievalDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodTieringStorageRetrievalDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodTieringStorageRetrievalDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodTieringStorageRetrievalDataResponseBody) *DescribeVodTieringStorageRetrievalDataResponse
	GetBody() *DescribeVodTieringStorageRetrievalDataResponseBody
}

type DescribeVodTieringStorageRetrievalDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodTieringStorageRetrievalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodTieringStorageRetrievalDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageRetrievalDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) GetBody() *DescribeVodTieringStorageRetrievalDataResponseBody {
	return s.Body
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) SetHeaders(v map[string]*string) *DescribeVodTieringStorageRetrievalDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) SetStatusCode(v int32) *DescribeVodTieringStorageRetrievalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) SetBody(v *DescribeVodTieringStorageRetrievalDataResponseBody) *DescribeVodTieringStorageRetrievalDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
