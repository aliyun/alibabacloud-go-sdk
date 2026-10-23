// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudPhoneNodePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudPhoneNodePriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudPhoneNodePriceResponseBody) *DescribeCloudPhoneNodePriceResponse
	GetBody() *DescribeCloudPhoneNodePriceResponseBody
}

type DescribeCloudPhoneNodePriceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudPhoneNodePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudPhoneNodePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudPhoneNodePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudPhoneNodePriceResponse) GetBody() *DescribeCloudPhoneNodePriceResponseBody {
	return s.Body
}

func (s *DescribeCloudPhoneNodePriceResponse) SetHeaders(v map[string]*string) *DescribeCloudPhoneNodePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponse) SetStatusCode(v int32) *DescribeCloudPhoneNodePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponse) SetBody(v *DescribeCloudPhoneNodePriceResponseBody) *DescribeCloudPhoneNodePriceResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
