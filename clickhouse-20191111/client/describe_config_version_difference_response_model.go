// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigVersionDifferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfigVersionDifferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfigVersionDifferenceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfigVersionDifferenceResponseBody) *DescribeConfigVersionDifferenceResponse
	GetBody() *DescribeConfigVersionDifferenceResponseBody
}

type DescribeConfigVersionDifferenceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigVersionDifferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigVersionDifferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigVersionDifferenceResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigVersionDifferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfigVersionDifferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfigVersionDifferenceResponse) GetBody() *DescribeConfigVersionDifferenceResponseBody {
	return s.Body
}

func (s *DescribeConfigVersionDifferenceResponse) SetHeaders(v map[string]*string) *DescribeConfigVersionDifferenceResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigVersionDifferenceResponse) SetStatusCode(v int32) *DescribeConfigVersionDifferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponse) SetBody(v *DescribeConfigVersionDifferenceResponseBody) *DescribeConfigVersionDifferenceResponse {
	s.Body = v
	return s
}

func (s *DescribeConfigVersionDifferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
