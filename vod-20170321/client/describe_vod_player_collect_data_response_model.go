// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodPlayerCollectDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodPlayerCollectDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodPlayerCollectDataResponseBody) *DescribeVodPlayerCollectDataResponse
	GetBody() *DescribeVodPlayerCollectDataResponseBody
}

type DescribeVodPlayerCollectDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodPlayerCollectDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodPlayerCollectDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodPlayerCollectDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodPlayerCollectDataResponse) GetBody() *DescribeVodPlayerCollectDataResponseBody {
	return s.Body
}

func (s *DescribeVodPlayerCollectDataResponse) SetHeaders(v map[string]*string) *DescribeVodPlayerCollectDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodPlayerCollectDataResponse) SetStatusCode(v int32) *DescribeVodPlayerCollectDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodPlayerCollectDataResponse) SetBody(v *DescribeVodPlayerCollectDataResponseBody) *DescribeVodPlayerCollectDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodPlayerCollectDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
