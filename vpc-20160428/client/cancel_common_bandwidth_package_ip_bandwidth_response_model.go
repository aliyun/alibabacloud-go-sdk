// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCommonBandwidthPackageIpBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCommonBandwidthPackageIpBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCommonBandwidthPackageIpBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *CancelCommonBandwidthPackageIpBandwidthResponseBody) *CancelCommonBandwidthPackageIpBandwidthResponse
	GetBody() *CancelCommonBandwidthPackageIpBandwidthResponseBody
}

type CancelCommonBandwidthPackageIpBandwidthResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCommonBandwidthPackageIpBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCommonBandwidthPackageIpBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCommonBandwidthPackageIpBandwidthResponse) GoString() string {
	return s.String()
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) GetBody() *CancelCommonBandwidthPackageIpBandwidthResponseBody {
	return s.Body
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) SetHeaders(v map[string]*string) *CancelCommonBandwidthPackageIpBandwidthResponse {
	s.Headers = v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) SetStatusCode(v int32) *CancelCommonBandwidthPackageIpBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) SetBody(v *CancelCommonBandwidthPackageIpBandwidthResponseBody) *CancelCommonBandwidthPackageIpBandwidthResponse {
	s.Body = v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
