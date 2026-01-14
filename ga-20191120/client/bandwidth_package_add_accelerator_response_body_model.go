// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthPackageAddAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerators(v []*string) *BandwidthPackageAddAcceleratorResponseBody
	GetAccelerators() []*string
	SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorResponseBody
	GetBandwidthPackageId() *string
	SetRequestId(v string) *BandwidthPackageAddAcceleratorResponseBody
	GetRequestId() *string
}

type BandwidthPackageAddAcceleratorResponseBody struct {
	// The GA instance IDs.
	Accelerators []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The bandwidth plan ID.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B7770CB9-9745-4FE5-9EDA-D14B01A12A50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BandwidthPackageAddAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorResponseBody) GetAccelerators() []*string {
	return s.Accelerators
}

func (s *BandwidthPackageAddAcceleratorResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *BandwidthPackageAddAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetAccelerators(v []*string) *BandwidthPackageAddAcceleratorResponseBody {
	s.Accelerators = v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetRequestId(v string) *BandwidthPackageAddAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
