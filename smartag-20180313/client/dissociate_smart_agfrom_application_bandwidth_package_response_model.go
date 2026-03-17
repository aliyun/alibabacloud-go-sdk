// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateSmartAGFromApplicationBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateSmartAGFromApplicationBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateSmartAGFromApplicationBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *DissociateSmartAGFromApplicationBandwidthPackageResponseBody) *DissociateSmartAGFromApplicationBandwidthPackageResponse
	GetBody() *DissociateSmartAGFromApplicationBandwidthPackageResponseBody
}

type DissociateSmartAGFromApplicationBandwidthPackageResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateSmartAGFromApplicationBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateSmartAGFromApplicationBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateSmartAGFromApplicationBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) GetBody() *DissociateSmartAGFromApplicationBandwidthPackageResponseBody {
	return s.Body
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) SetHeaders(v map[string]*string) *DissociateSmartAGFromApplicationBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) SetStatusCode(v int32) *DissociateSmartAGFromApplicationBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) SetBody(v *DissociateSmartAGFromApplicationBandwidthPackageResponseBody) *DissociateSmartAGFromApplicationBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
