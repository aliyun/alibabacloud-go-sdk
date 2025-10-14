// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressPoolReferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmAddressPoolReferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmAddressPoolReferenceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmAddressPoolReferenceResponseBody) *DescribeCloudGtmAddressPoolReferenceResponse
	GetBody() *DescribeCloudGtmAddressPoolReferenceResponseBody
}

type DescribeCloudGtmAddressPoolReferenceResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmAddressPoolReferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmAddressPoolReferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolReferenceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) GetBody() *DescribeCloudGtmAddressPoolReferenceResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmAddressPoolReferenceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) SetStatusCode(v int32) *DescribeCloudGtmAddressPoolReferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) SetBody(v *DescribeCloudGtmAddressPoolReferenceResponseBody) *DescribeCloudGtmAddressPoolReferenceResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
