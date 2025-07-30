// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAdditionalBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAdditionalBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAdditionalBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *RenewAdditionalBandwidthResponseBody) *RenewAdditionalBandwidthResponse
	GetBody() *RenewAdditionalBandwidthResponseBody
}

type RenewAdditionalBandwidthResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAdditionalBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAdditionalBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAdditionalBandwidthResponse) GoString() string {
	return s.String()
}

func (s *RenewAdditionalBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAdditionalBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAdditionalBandwidthResponse) GetBody() *RenewAdditionalBandwidthResponseBody {
	return s.Body
}

func (s *RenewAdditionalBandwidthResponse) SetHeaders(v map[string]*string) *RenewAdditionalBandwidthResponse {
	s.Headers = v
	return s
}

func (s *RenewAdditionalBandwidthResponse) SetStatusCode(v int32) *RenewAdditionalBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAdditionalBandwidthResponse) SetBody(v *RenewAdditionalBandwidthResponseBody) *RenewAdditionalBandwidthResponse {
	s.Body = v
	return s
}

func (s *RenewAdditionalBandwidthResponse) Validate() error {
	return dara.Validate(s)
}
