// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageIpBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCommonBandwidthPackageIpBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCommonBandwidthPackageIpBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCommonBandwidthPackageIpBandwidthResponseBody) *ModifyCommonBandwidthPackageIpBandwidthResponse
	GetBody() *ModifyCommonBandwidthPackageIpBandwidthResponseBody
}

type ModifyCommonBandwidthPackageIpBandwidthResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCommonBandwidthPackageIpBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCommonBandwidthPackageIpBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageIpBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) GetBody() *ModifyCommonBandwidthPackageIpBandwidthResponseBody {
	return s.Body
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) SetHeaders(v map[string]*string) *ModifyCommonBandwidthPackageIpBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) SetStatusCode(v int32) *ModifyCommonBandwidthPackageIpBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) SetBody(v *ModifyCommonBandwidthPackageIpBandwidthResponseBody) *ModifyCommonBandwidthPackageIpBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
