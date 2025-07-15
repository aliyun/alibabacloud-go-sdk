// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCommonBandwidthPackageSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCommonBandwidthPackageSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCommonBandwidthPackageSpecResponseBody) *ModifyCommonBandwidthPackageSpecResponse
	GetBody() *ModifyCommonBandwidthPackageSpecResponseBody
}

type ModifyCommonBandwidthPackageSpecResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCommonBandwidthPackageSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCommonBandwidthPackageSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCommonBandwidthPackageSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCommonBandwidthPackageSpecResponse) GetBody() *ModifyCommonBandwidthPackageSpecResponseBody {
	return s.Body
}

func (s *ModifyCommonBandwidthPackageSpecResponse) SetHeaders(v map[string]*string) *ModifyCommonBandwidthPackageSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecResponse) SetStatusCode(v int32) *ModifyCommonBandwidthPackageSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecResponse) SetBody(v *ModifyCommonBandwidthPackageSpecResponseBody) *ModifyCommonBandwidthPackageSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecResponse) Validate() error {
	return dara.Validate(s)
}
