// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatasetInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatasetInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatasetInfoResponseBody) *DescribeDatasetInfoResponse
	GetBody() *DescribeDatasetInfoResponseBody
}

type DescribeDatasetInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatasetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatasetInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatasetInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatasetInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatasetInfoResponse) GetBody() *DescribeDatasetInfoResponseBody {
	return s.Body
}

func (s *DescribeDatasetInfoResponse) SetHeaders(v map[string]*string) *DescribeDatasetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatasetInfoResponse) SetStatusCode(v int32) *DescribeDatasetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatasetInfoResponse) SetBody(v *DescribeDatasetInfoResponseBody) *DescribeDatasetInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDatasetInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
