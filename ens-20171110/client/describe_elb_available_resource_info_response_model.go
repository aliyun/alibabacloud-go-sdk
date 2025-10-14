// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElbAvailableResourceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElbAvailableResourceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElbAvailableResourceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElbAvailableResourceInfoResponseBody) *DescribeElbAvailableResourceInfoResponse
	GetBody() *DescribeElbAvailableResourceInfoResponseBody
}

type DescribeElbAvailableResourceInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElbAvailableResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElbAvailableResourceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElbAvailableResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeElbAvailableResourceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElbAvailableResourceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElbAvailableResourceInfoResponse) GetBody() *DescribeElbAvailableResourceInfoResponseBody {
	return s.Body
}

func (s *DescribeElbAvailableResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeElbAvailableResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponse) SetStatusCode(v int32) *DescribeElbAvailableResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponse) SetBody(v *DescribeElbAvailableResourceInfoResponseBody) *DescribeElbAvailableResourceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
