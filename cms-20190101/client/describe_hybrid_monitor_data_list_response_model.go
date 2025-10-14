// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorDataListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridMonitorDataListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridMonitorDataListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridMonitorDataListResponseBody) *DescribeHybridMonitorDataListResponse
	GetBody() *DescribeHybridMonitorDataListResponseBody
}

type DescribeHybridMonitorDataListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridMonitorDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridMonitorDataListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorDataListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorDataListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridMonitorDataListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridMonitorDataListResponse) GetBody() *DescribeHybridMonitorDataListResponseBody {
	return s.Body
}

func (s *DescribeHybridMonitorDataListResponse) SetHeaders(v map[string]*string) *DescribeHybridMonitorDataListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridMonitorDataListResponse) SetStatusCode(v int32) *DescribeHybridMonitorDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponse) SetBody(v *DescribeHybridMonitorDataListResponseBody) *DescribeHybridMonitorDataListResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridMonitorDataListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
