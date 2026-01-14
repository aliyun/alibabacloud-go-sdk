// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthPackageRemoveAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BandwidthPackageRemoveAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BandwidthPackageRemoveAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *BandwidthPackageRemoveAcceleratorResponseBody) *BandwidthPackageRemoveAcceleratorResponse
	GetBody() *BandwidthPackageRemoveAcceleratorResponseBody
}

type BandwidthPackageRemoveAcceleratorResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BandwidthPackageRemoveAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BandwidthPackageRemoveAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BandwidthPackageRemoveAcceleratorResponse) GetBody() *BandwidthPackageRemoveAcceleratorResponseBody {
	return s.Body
}

func (s *BandwidthPackageRemoveAcceleratorResponse) SetHeaders(v map[string]*string) *BandwidthPackageRemoveAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponse) SetStatusCode(v int32) *BandwidthPackageRemoveAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponse) SetBody(v *BandwidthPackageRemoveAcceleratorResponseBody) *BandwidthPackageRemoveAcceleratorResponse {
	s.Body = v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
