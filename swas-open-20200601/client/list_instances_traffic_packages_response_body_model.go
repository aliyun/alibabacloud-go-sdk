// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesTrafficPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTrafficPackageUsages(v []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) *ListInstancesTrafficPackagesResponseBody
	GetInstanceTrafficPackageUsages() []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages
	SetRequestId(v string) *ListInstancesTrafficPackagesResponseBody
	GetRequestId() *string
}

type ListInstancesTrafficPackagesResponseBody struct {
	// The data transfers that exceed the quota of the data transfer plan in the current month. Unit: bytes.
	InstanceTrafficPackageUsages []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages `json:"InstanceTrafficPackageUsages,omitempty" xml:"InstanceTrafficPackageUsages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesTrafficPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponseBody) GetInstanceTrafficPackageUsages() []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	return s.InstanceTrafficPackageUsages
}

func (s *ListInstancesTrafficPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesTrafficPackagesResponseBody) SetInstanceTrafficPackageUsages(v []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) *ListInstancesTrafficPackagesResponseBody {
	s.InstanceTrafficPackageUsages = v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBody) SetRequestId(v string) *ListInstancesTrafficPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBody) Validate() error {
	if s.InstanceTrafficPackageUsages != nil {
		for _, item := range s.InstanceTrafficPackageUsages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// ccscqwqwqqqw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The data transfers that exceeds the quota of the data transfer plan in the current month. Unit: Byte.
	//
	// example:
	//
	// 0
	TrafficOverflow *int64 `json:"TrafficOverflow,omitempty" xml:"TrafficOverflow,omitempty"`
	// The unused quota of the data transfer plan in the current month. Unit: Byte.
	//
	// example:
	//
	// 10000
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitempty" xml:"TrafficPackageRemaining,omitempty"`
	// The quota of the data transfer plan in the current month. Unit: Byte.
	//
	// > TrafficPackageTotal = TrafficUsed + TrafficPackageRemaining
	//
	// example:
	//
	// 20000
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitempty" xml:"TrafficPackageTotal,omitempty"`
	// The used quota of the data transfer plan in the current month. Unit: Byte.
	//
	// example:
	//
	// 10000
	TrafficUsed *int64 `json:"TrafficUsed,omitempty" xml:"TrafficUsed,omitempty"`
}

func (s ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GetTrafficOverflow() *int64 {
	return s.TrafficOverflow
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GetTrafficPackageRemaining() *int64 {
	return s.TrafficPackageRemaining
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GetTrafficPackageTotal() *int64 {
	return s.TrafficPackageTotal
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GetTrafficUsed() *int64 {
	return s.TrafficUsed
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetInstanceId(v string) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficOverflow(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficOverflow = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficPackageRemaining(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficPackageRemaining = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficPackageTotal(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficPackageTotal = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficUsed(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficUsed = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) Validate() error {
	return dara.Validate(s)
}
