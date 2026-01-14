// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthPackageAddAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BandwidthPackageAddAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BandwidthPackageAddAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *BandwidthPackageAddAcceleratorResponseBody) *BandwidthPackageAddAcceleratorResponse
	GetBody() *BandwidthPackageAddAcceleratorResponseBody
}

type BandwidthPackageAddAcceleratorResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BandwidthPackageAddAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BandwidthPackageAddAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BandwidthPackageAddAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BandwidthPackageAddAcceleratorResponse) GetBody() *BandwidthPackageAddAcceleratorResponseBody {
	return s.Body
}

func (s *BandwidthPackageAddAcceleratorResponse) SetHeaders(v map[string]*string) *BandwidthPackageAddAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponse) SetStatusCode(v int32) *BandwidthPackageAddAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponse) SetBody(v *BandwidthPackageAddAcceleratorResponseBody) *BandwidthPackageAddAcceleratorResponse {
	s.Body = v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
