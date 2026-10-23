// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodeRenewPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudPhoneNodeRenewPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudPhoneNodeRenewPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudPhoneNodeRenewPriceResponseBody) *DescribeCloudPhoneNodeRenewPriceResponse
	GetBody() *DescribeCloudPhoneNodeRenewPriceResponseBody
}

type DescribeCloudPhoneNodeRenewPriceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudPhoneNodeRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) GetBody() *DescribeCloudPhoneNodeRenewPriceResponseBody {
	return s.Body
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) SetHeaders(v map[string]*string) *DescribeCloudPhoneNodeRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) SetStatusCode(v int32) *DescribeCloudPhoneNodeRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) SetBody(v *DescribeCloudPhoneNodeRenewPriceResponseBody) *DescribeCloudPhoneNodeRenewPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
