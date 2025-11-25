// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingDestinationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingDestinationResponseBody) *DescribeOutgoingDestinationResponse
	GetBody() *DescribeOutgoingDestinationResponseBody
}

type DescribeOutgoingDestinationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingDestinationResponse) GetBody() *DescribeOutgoingDestinationResponseBody {
	return s.Body
}

func (s *DescribeOutgoingDestinationResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDestinationResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDestinationResponse) SetStatusCode(v int32) *DescribeOutgoingDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDestinationResponse) SetBody(v *DescribeOutgoingDestinationResponseBody) *DescribeOutgoingDestinationResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
