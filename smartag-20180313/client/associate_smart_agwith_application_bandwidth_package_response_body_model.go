// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateSmartAGWithApplicationBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateSmartAGWithApplicationBandwidthPackageResponseBody
	GetRequestId() *string
}

type AssociateSmartAGWithApplicationBandwidthPackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFF2FD9D-66BD-4DD4-8152-C4508119D7B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateSmartAGWithApplicationBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateSmartAGWithApplicationBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponseBody) SetRequestId(v string) *AssociateSmartAGWithApplicationBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
