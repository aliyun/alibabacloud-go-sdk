// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCommonBandwidthPackageIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveCommonBandwidthPackageIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveCommonBandwidthPackageIpResponse
	GetStatusCode() *int32
	SetBody(v *RemoveCommonBandwidthPackageIpResponseBody) *RemoveCommonBandwidthPackageIpResponse
	GetBody() *RemoveCommonBandwidthPackageIpResponseBody
}

type RemoveCommonBandwidthPackageIpResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCommonBandwidthPackageIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCommonBandwidthPackageIpResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveCommonBandwidthPackageIpResponse) GoString() string {
	return s.String()
}

func (s *RemoveCommonBandwidthPackageIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveCommonBandwidthPackageIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveCommonBandwidthPackageIpResponse) GetBody() *RemoveCommonBandwidthPackageIpResponseBody {
	return s.Body
}

func (s *RemoveCommonBandwidthPackageIpResponse) SetHeaders(v map[string]*string) *RemoveCommonBandwidthPackageIpResponse {
	s.Headers = v
	return s
}

func (s *RemoveCommonBandwidthPackageIpResponse) SetStatusCode(v int32) *RemoveCommonBandwidthPackageIpResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpResponse) SetBody(v *RemoveCommonBandwidthPackageIpResponseBody) *RemoveCommonBandwidthPackageIpResponse {
	s.Body = v
	return s
}

func (s *RemoveCommonBandwidthPackageIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
