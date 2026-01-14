// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateBandwidthPackageResponseBody) *CreateBandwidthPackageResponse
	GetBody() *CreateBandwidthPackageResponseBody
}

type CreateBandwidthPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBandwidthPackageResponse) GetBody() *CreateBandwidthPackageResponseBody {
	return s.Body
}

func (s *CreateBandwidthPackageResponse) SetHeaders(v map[string]*string) *CreateBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateBandwidthPackageResponse) SetStatusCode(v int32) *CreateBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBandwidthPackageResponse) SetBody(v *CreateBandwidthPackageResponseBody) *CreateBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *CreateBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
