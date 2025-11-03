// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualBorderBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVirtualBorderBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVirtualBorderBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVirtualBorderBandwidthResponseBody) *UpdateVirtualBorderBandwidthResponse
	GetBody() *UpdateVirtualBorderBandwidthResponseBody
}

type UpdateVirtualBorderBandwidthResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirtualBorderBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirtualBorderBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualBorderBandwidthResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirtualBorderBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVirtualBorderBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVirtualBorderBandwidthResponse) GetBody() *UpdateVirtualBorderBandwidthResponseBody {
	return s.Body
}

func (s *UpdateVirtualBorderBandwidthResponse) SetHeaders(v map[string]*string) *UpdateVirtualBorderBandwidthResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponse) SetStatusCode(v int32) *UpdateVirtualBorderBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponse) SetBody(v *UpdateVirtualBorderBandwidthResponseBody) *UpdateVirtualBorderBandwidthResponse {
	s.Body = v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
