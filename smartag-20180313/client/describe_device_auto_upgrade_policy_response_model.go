// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceAutoUpgradePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceAutoUpgradePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceAutoUpgradePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceAutoUpgradePolicyResponseBody) *DescribeDeviceAutoUpgradePolicyResponse
	GetBody() *DescribeDeviceAutoUpgradePolicyResponseBody
}

type DescribeDeviceAutoUpgradePolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceAutoUpgradePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceAutoUpgradePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceAutoUpgradePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) GetBody() *DescribeDeviceAutoUpgradePolicyResponseBody {
	return s.Body
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) SetHeaders(v map[string]*string) *DescribeDeviceAutoUpgradePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) SetStatusCode(v int32) *DescribeDeviceAutoUpgradePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) SetBody(v *DescribeDeviceAutoUpgradePolicyResponseBody) *DescribeDeviceAutoUpgradePolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
