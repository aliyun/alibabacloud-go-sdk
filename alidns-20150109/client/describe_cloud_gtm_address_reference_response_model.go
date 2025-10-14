// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressReferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmAddressReferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmAddressReferenceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmAddressReferenceResponseBody) *DescribeCloudGtmAddressReferenceResponse
	GetBody() *DescribeCloudGtmAddressReferenceResponseBody
}

type DescribeCloudGtmAddressReferenceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmAddressReferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmAddressReferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmAddressReferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmAddressReferenceResponse) GetBody() *DescribeCloudGtmAddressReferenceResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmAddressReferenceResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmAddressReferenceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponse) SetStatusCode(v int32) *DescribeCloudGtmAddressReferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponse) SetBody(v *DescribeCloudGtmAddressReferenceResponseBody) *DescribeCloudGtmAddressReferenceResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
