// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommonBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCommonBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCommonBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateCommonBandwidthPackageResponseBody) *CreateCommonBandwidthPackageResponse
	GetBody() *CreateCommonBandwidthPackageResponseBody
}

type CreateCommonBandwidthPackageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommonBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommonBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCommonBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateCommonBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCommonBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCommonBandwidthPackageResponse) GetBody() *CreateCommonBandwidthPackageResponseBody {
	return s.Body
}

func (s *CreateCommonBandwidthPackageResponse) SetHeaders(v map[string]*string) *CreateCommonBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateCommonBandwidthPackageResponse) SetStatusCode(v int32) *CreateCommonBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommonBandwidthPackageResponse) SetBody(v *CreateCommonBandwidthPackageResponseBody) *CreateCommonBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *CreateCommonBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
