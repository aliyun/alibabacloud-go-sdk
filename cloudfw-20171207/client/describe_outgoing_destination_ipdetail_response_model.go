// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationIPDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingDestinationIPDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingDestinationIPDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingDestinationIPDetailResponseBody) *DescribeOutgoingDestinationIPDetailResponse
	GetBody() *DescribeOutgoingDestinationIPDetailResponseBody
}

type DescribeOutgoingDestinationIPDetailResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDestinationIPDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDestinationIPDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingDestinationIPDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingDestinationIPDetailResponse) GetBody() *DescribeOutgoingDestinationIPDetailResponseBody {
	return s.Body
}

func (s *DescribeOutgoingDestinationIPDetailResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDestinationIPDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponse) SetStatusCode(v int32) *DescribeOutgoingDestinationIPDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponse) SetBody(v *DescribeOutgoingDestinationIPDetailResponseBody) *DescribeOutgoingDestinationIPDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
