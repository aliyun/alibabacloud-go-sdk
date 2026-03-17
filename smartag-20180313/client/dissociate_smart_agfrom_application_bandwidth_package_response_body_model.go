// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateSmartAGFromApplicationBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateSmartAGFromApplicationBandwidthPackageResponseBody
	GetRequestId() *string
}

type DissociateSmartAGFromApplicationBandwidthPackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateSmartAGFromApplicationBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateSmartAGFromApplicationBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponseBody) SetRequestId(v string) *DissociateSmartAGFromApplicationBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
