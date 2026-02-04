// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateCenBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateCenBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateCenBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *AssociateCenBandwidthPackageResponseBody) *AssociateCenBandwidthPackageResponse
	GetBody() *AssociateCenBandwidthPackageResponseBody
}

type AssociateCenBandwidthPackageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateCenBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateCenBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateCenBandwidthPackageResponse) GetBody() *AssociateCenBandwidthPackageResponseBody {
	return s.Body
}

func (s *AssociateCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *AssociateCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *AssociateCenBandwidthPackageResponse) SetStatusCode(v int32) *AssociateCenBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateCenBandwidthPackageResponse) SetBody(v *AssociateCenBandwidthPackageResponseBody) *AssociateCenBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *AssociateCenBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
