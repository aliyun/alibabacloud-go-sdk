// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthPackageRemoveAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerators(v []*string) *BandwidthPackageRemoveAcceleratorResponseBody
	GetAccelerators() []*string
	SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorResponseBody
	GetBandwidthPackageId() *string
	SetRequestId(v string) *BandwidthPackageRemoveAcceleratorResponseBody
	GetRequestId() *string
}

type BandwidthPackageRemoveAcceleratorResponseBody struct {
	// The ID of the GA instance.
	Accelerators []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B7770CB9-9745-4FE5-9EDA-D14B01A12A50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) GetAccelerators() []*string {
	return s.Accelerators
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetAccelerators(v []*string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.Accelerators = v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetRequestId(v string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
