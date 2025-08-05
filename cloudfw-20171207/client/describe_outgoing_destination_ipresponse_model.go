// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingDestinationIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingDestinationIPResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingDestinationIPResponseBody) *DescribeOutgoingDestinationIPResponse
	GetBody() *DescribeOutgoingDestinationIPResponseBody
}

type DescribeOutgoingDestinationIPResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDestinationIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingDestinationIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingDestinationIPResponse) GetBody() *DescribeOutgoingDestinationIPResponseBody {
	return s.Body
}

func (s *DescribeOutgoingDestinationIPResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDestinationIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) SetStatusCode(v int32) *DescribeOutgoingDestinationIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) SetBody(v *DescribeOutgoingDestinationIPResponseBody) *DescribeOutgoingDestinationIPResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) Validate() error {
	return dara.Validate(s)
}
