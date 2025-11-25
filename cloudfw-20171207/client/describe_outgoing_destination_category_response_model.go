// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingDestinationCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingDestinationCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingDestinationCategoryResponseBody) *DescribeOutgoingDestinationCategoryResponse
	GetBody() *DescribeOutgoingDestinationCategoryResponseBody
}

type DescribeOutgoingDestinationCategoryResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDestinationCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDestinationCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationCategoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingDestinationCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingDestinationCategoryResponse) GetBody() *DescribeOutgoingDestinationCategoryResponseBody {
	return s.Body
}

func (s *DescribeOutgoingDestinationCategoryResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDestinationCategoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponse) SetStatusCode(v int32) *DescribeOutgoingDestinationCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponse) SetBody(v *DescribeOutgoingDestinationCategoryResponseBody) *DescribeOutgoingDestinationCategoryResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
