// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateSmartAGWithApplicationBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateSmartAGWithApplicationBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateSmartAGWithApplicationBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *AssociateSmartAGWithApplicationBandwidthPackageResponseBody) *AssociateSmartAGWithApplicationBandwidthPackageResponse
	GetBody() *AssociateSmartAGWithApplicationBandwidthPackageResponseBody
}

type AssociateSmartAGWithApplicationBandwidthPackageResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateSmartAGWithApplicationBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateSmartAGWithApplicationBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateSmartAGWithApplicationBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) GetBody() *AssociateSmartAGWithApplicationBandwidthPackageResponseBody {
	return s.Body
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) SetHeaders(v map[string]*string) *AssociateSmartAGWithApplicationBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) SetStatusCode(v int32) *AssociateSmartAGWithApplicationBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) SetBody(v *AssociateSmartAGWithApplicationBandwidthPackageResponseBody) *AssociateSmartAGWithApplicationBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
