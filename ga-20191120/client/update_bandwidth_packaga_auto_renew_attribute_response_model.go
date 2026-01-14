// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBandwidthPackagaAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBandwidthPackagaAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBandwidthPackagaAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) *UpdateBandwidthPackagaAutoRenewAttributeResponse
	GetBody() *UpdateBandwidthPackagaAutoRenewAttributeResponseBody
}

type UpdateBandwidthPackagaAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBandwidthPackagaAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) GetBody() *UpdateBandwidthPackagaAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *UpdateBandwidthPackagaAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) SetStatusCode(v int32) *UpdateBandwidthPackagaAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) SetBody(v *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) *UpdateBandwidthPackagaAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
