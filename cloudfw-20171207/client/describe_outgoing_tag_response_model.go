// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingTagResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingTagResponseBody) *DescribeOutgoingTagResponse
	GetBody() *DescribeOutgoingTagResponseBody
}

type DescribeOutgoingTagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingTagResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingTagResponse) GetBody() *DescribeOutgoingTagResponseBody {
	return s.Body
}

func (s *DescribeOutgoingTagResponse) SetHeaders(v map[string]*string) *DescribeOutgoingTagResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingTagResponse) SetStatusCode(v int32) *DescribeOutgoingTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingTagResponse) SetBody(v *DescribeOutgoingTagResponseBody) *DescribeOutgoingTagResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
