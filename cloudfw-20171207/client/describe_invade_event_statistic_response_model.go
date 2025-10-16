// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvadeEventStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvadeEventStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvadeEventStatisticResponseBody) *DescribeInvadeEventStatisticResponse
	GetBody() *DescribeInvadeEventStatisticResponseBody
}

type DescribeInvadeEventStatisticResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEventStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEventStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvadeEventStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvadeEventStatisticResponse) GetBody() *DescribeInvadeEventStatisticResponseBody {
	return s.Body
}

func (s *DescribeInvadeEventStatisticResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventStatisticResponse) SetStatusCode(v int32) *DescribeInvadeEventStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventStatisticResponse) SetBody(v *DescribeInvadeEventStatisticResponseBody) *DescribeInvadeEventStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeInvadeEventStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
