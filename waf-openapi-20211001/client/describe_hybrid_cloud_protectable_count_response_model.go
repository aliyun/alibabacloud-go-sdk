// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudProtectableCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudProtectableCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudProtectableCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudProtectableCountResponseBody) *DescribeHybridCloudProtectableCountResponse
	GetBody() *DescribeHybridCloudProtectableCountResponseBody
}

type DescribeHybridCloudProtectableCountResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudProtectableCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudProtectableCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProtectableCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProtectableCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudProtectableCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudProtectableCountResponse) GetBody() *DescribeHybridCloudProtectableCountResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudProtectableCountResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudProtectableCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudProtectableCountResponse) SetStatusCode(v int32) *DescribeHybridCloudProtectableCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudProtectableCountResponse) SetBody(v *DescribeHybridCloudProtectableCountResponseBody) *DescribeHybridCloudProtectableCountResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudProtectableCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
