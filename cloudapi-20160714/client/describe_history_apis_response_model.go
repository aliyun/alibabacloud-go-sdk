// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHistoryApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHistoryApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHistoryApisResponseBody) *DescribeHistoryApisResponse
	GetBody() *DescribeHistoryApisResponseBody
}

type DescribeHistoryApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHistoryApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHistoryApisResponse) GetBody() *DescribeHistoryApisResponseBody {
	return s.Body
}

func (s *DescribeHistoryApisResponse) SetHeaders(v map[string]*string) *DescribeHistoryApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryApisResponse) SetStatusCode(v int32) *DescribeHistoryApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryApisResponse) SetBody(v *DescribeHistoryApisResponseBody) *DescribeHistoryApisResponse {
	s.Body = v
	return s
}

func (s *DescribeHistoryApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
