// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenBandwidthPackageSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCenBandwidthPackageSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCenBandwidthPackageSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCenBandwidthPackageSpecResponseBody) *ModifyCenBandwidthPackageSpecResponse
	GetBody() *ModifyCenBandwidthPackageSpecResponseBody
}

type ModifyCenBandwidthPackageSpecResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCenBandwidthPackageSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCenBandwidthPackageSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCenBandwidthPackageSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCenBandwidthPackageSpecResponse) GetBody() *ModifyCenBandwidthPackageSpecResponseBody {
	return s.Body
}

func (s *ModifyCenBandwidthPackageSpecResponse) SetHeaders(v map[string]*string) *ModifyCenBandwidthPackageSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenBandwidthPackageSpecResponse) SetStatusCode(v int32) *ModifyCenBandwidthPackageSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecResponse) SetBody(v *ModifyCenBandwidthPackageSpecResponseBody) *ModifyCenBandwidthPackageSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyCenBandwidthPackageSpecResponse) Validate() error {
	return dara.Validate(s)
}
