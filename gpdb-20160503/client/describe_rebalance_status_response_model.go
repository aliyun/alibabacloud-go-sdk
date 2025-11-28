// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRebalanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRebalanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRebalanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRebalanceStatusResponseBody) *DescribeRebalanceStatusResponse
	GetBody() *DescribeRebalanceStatusResponseBody
}

type DescribeRebalanceStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRebalanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRebalanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRebalanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRebalanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRebalanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRebalanceStatusResponse) GetBody() *DescribeRebalanceStatusResponseBody {
	return s.Body
}

func (s *DescribeRebalanceStatusResponse) SetHeaders(v map[string]*string) *DescribeRebalanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRebalanceStatusResponse) SetStatusCode(v int32) *DescribeRebalanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRebalanceStatusResponse) SetBody(v *DescribeRebalanceStatusResponseBody) *DescribeRebalanceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRebalanceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
