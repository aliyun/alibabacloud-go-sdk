// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingStatisticResponseBody) *DescribeOutgoingStatisticResponse
	GetBody() *DescribeOutgoingStatisticResponseBody
}

type DescribeOutgoingStatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingStatisticResponse) GetBody() *DescribeOutgoingStatisticResponseBody {
	return s.Body
}

func (s *DescribeOutgoingStatisticResponse) SetHeaders(v map[string]*string) *DescribeOutgoingStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingStatisticResponse) SetStatusCode(v int32) *DescribeOutgoingStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingStatisticResponse) SetBody(v *DescribeOutgoingStatisticResponseBody) *DescribeOutgoingStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
