// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenBandwidthPackageAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCenBandwidthPackageAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCenBandwidthPackageAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCenBandwidthPackageAttributeResponseBody) *ModifyCenBandwidthPackageAttributeResponse
	GetBody() *ModifyCenBandwidthPackageAttributeResponseBody
}

type ModifyCenBandwidthPackageAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCenBandwidthPackageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCenBandwidthPackageAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCenBandwidthPackageAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCenBandwidthPackageAttributeResponse) GetBody() *ModifyCenBandwidthPackageAttributeResponseBody {
	return s.Body
}

func (s *ModifyCenBandwidthPackageAttributeResponse) SetHeaders(v map[string]*string) *ModifyCenBandwidthPackageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeResponse) SetStatusCode(v int32) *ModifyCenBandwidthPackageAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeResponse) SetBody(v *ModifyCenBandwidthPackageAttributeResponseBody) *ModifyCenBandwidthPackageAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeResponse) Validate() error {
	return dara.Validate(s)
}
