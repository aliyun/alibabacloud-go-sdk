// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataReDistributeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataReDistributeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataReDistributeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataReDistributeInfoResponseBody) *DescribeDataReDistributeInfoResponse
	GetBody() *DescribeDataReDistributeInfoResponseBody
}

type DescribeDataReDistributeInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataReDistributeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataReDistributeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataReDistributeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataReDistributeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataReDistributeInfoResponse) GetBody() *DescribeDataReDistributeInfoResponseBody {
	return s.Body
}

func (s *DescribeDataReDistributeInfoResponse) SetHeaders(v map[string]*string) *DescribeDataReDistributeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataReDistributeInfoResponse) SetStatusCode(v int32) *DescribeDataReDistributeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponse) SetBody(v *DescribeDataReDistributeInfoResponseBody) *DescribeDataReDistributeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDataReDistributeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
