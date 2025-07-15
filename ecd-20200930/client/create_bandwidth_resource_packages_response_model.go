// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBandwidthResourcePackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBandwidthResourcePackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBandwidthResourcePackagesResponse
	GetStatusCode() *int32
	SetBody(v *CreateBandwidthResourcePackagesResponseBody) *CreateBandwidthResourcePackagesResponse
	GetBody() *CreateBandwidthResourcePackagesResponseBody
}

type CreateBandwidthResourcePackagesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBandwidthResourcePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBandwidthResourcePackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthResourcePackagesResponse) GoString() string {
	return s.String()
}

func (s *CreateBandwidthResourcePackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBandwidthResourcePackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBandwidthResourcePackagesResponse) GetBody() *CreateBandwidthResourcePackagesResponseBody {
	return s.Body
}

func (s *CreateBandwidthResourcePackagesResponse) SetHeaders(v map[string]*string) *CreateBandwidthResourcePackagesResponse {
	s.Headers = v
	return s
}

func (s *CreateBandwidthResourcePackagesResponse) SetStatusCode(v int32) *CreateBandwidthResourcePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBandwidthResourcePackagesResponse) SetBody(v *CreateBandwidthResourcePackagesResponseBody) *CreateBandwidthResourcePackagesResponse {
	s.Body = v
	return s
}

func (s *CreateBandwidthResourcePackagesResponse) Validate() error {
	return dara.Validate(s)
}
