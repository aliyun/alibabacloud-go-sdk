// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCommonBandwidthPackageAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCommonBandwidthPackageAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCommonBandwidthPackageAttributeResponseBody) *ModifyCommonBandwidthPackageAttributeResponse
	GetBody() *ModifyCommonBandwidthPackageAttributeResponseBody
}

type ModifyCommonBandwidthPackageAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCommonBandwidthPackageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCommonBandwidthPackageAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) GetBody() *ModifyCommonBandwidthPackageAttributeResponseBody {
	return s.Body
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) SetHeaders(v map[string]*string) *ModifyCommonBandwidthPackageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) SetStatusCode(v int32) *ModifyCommonBandwidthPackageAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) SetBody(v *ModifyCommonBandwidthPackageAttributeResponseBody) *ModifyCommonBandwidthPackageAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeResponse) Validate() error {
	return dara.Validate(s)
}
